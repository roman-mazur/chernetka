// Package exeproxy implements a server that can accept remotectl commands
// to execute some command piping its output. It can be used to maintain a single process
// for a helper pane of the editor: running the Go tests and programs of the line actions,
// other executables in general that don't implement remotectl protocol themselves.
package exeproxy

import (
	"cmp"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"rmazur.io/chernetka/internal/editor/styles"
	"rmazur.io/chernetka/internal/logger"
	"rmazur.io/chernetka/internal/remotectl"
	"rmazur.io/chernetka/internal/vt"
	"rmazur.io/chernetka/internal/vt/escape"
)

// Endpoint returns the endpoint of the proxy serving the purpose, like running the Go tests and programs:
// each purpose has its own proxy process, so the commands of one don't replace the commands of another.
func Endpoint(purpose string) remotectl.Endpoint {
	return remotectl.Endpoint("exeproxy").Scoped(purpose)
}

type CommandProxy struct {
	Logf logger.Func

	// Stdout and Stderr receive the output of the executed commands.
	// os.Stdout and os.Stderr are used if nil.
	Stdout, Stderr io.Writer

	// Header enables printing the command before it runs and its exit status after, to Stdout.
	Header bool
	// Clear clears the screen before each command, so that its output starts at the top.
	// A horizontal line is drawn before clearing, so the outputs are delimited in the scrollback.
	Clear bool
	// Grace limits how long a replaced command may take to exit after an interrupt before it's killed
	// (2s by default).
	Grace time.Duration

	pWatcher chan (<-chan struct{})

	srv *remotectl.Server

	mu      sync.Mutex
	p       *processTree
	pDone   chan struct{}
	started bool // Whether any command has been started.
}

func (cp *CommandProxy) Run(endpoint remotectl.Endpoint) error {
	srv, err := remotectl.NewServer(endpoint)
	if err != nil {
		return err
	}
	cp.srv = srv
	srv.Run(exeFunc(cp.exec), cp.logf)
	return nil
}

func (cp *CommandProxy) logf(fmt string, args ...any) {
	if cp.Logf != nil {
		cp.Logf(fmt, args...)
	}
}

func (cp *CommandProxy) stdout() io.Writer { return cmp.Or[io.Writer](cp.Stdout, os.Stdout) }

func (cp *CommandProxy) exec(cmd remotectl.CommandData) {
	if cmd.Action == "quit" {
		// This may result in an error on Send on the client side as they may not get an ack.
		// It's fine as it's used in tests mainly for cleaning up.
		cp.replaceProcess(nil, "")
		_ = cp.srv.Close()
		return
	}

	if len(cmd.Args) == 0 {
		cp.logf("exe command without args")
		return
	}

	exeCmd := exec.Command(cmd.Args[0], cmd.Args[1:]...)
	exeCmd.Dir = cmd.WorkDir
	exeCmd.Stdin = os.Stdin
	exeCmd.Stdout = cp.stdout()
	exeCmd.Stderr = cmp.Or[io.Writer](cp.Stderr, os.Stderr)
	cp.replaceProcess(exeCmd, cmdTitle(exeCmd))
}

func cmdTitle(cmd *exec.Cmd) string {
	p, err := filepath.Rel(".", cmd.Dir)
	if err != nil {
		p = cmd.Dir
	}
	return fmt.Sprintf("%s> %s", p, strings.Join(append([]string{filepath.Base(cmd.Path)}, cmd.Args[1:]...), " "))
}

// stop interrupts the running command and waits for it to exit, killing it after the grace period.
// It must be called with cp.mu held.
func (cp *CommandProxy) stop() {
	if cp.p == nil {
		return
	}
	cp.p.interrupt()
	select {
	case <-cp.pDone:
	case <-time.After(cmp.Or(cp.Grace, 2*time.Second)):
		cp.p.kill()
		<-cp.pDone
	}
	cp.p, cp.pDone = nil, nil
}

func (cp *CommandProxy) replaceProcess(exeCmd *exec.Cmd, title string) {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	cp.stop()

	if exeCmd == nil {
		return
	}

	p := newProcessTree(exeCmd)
	// The header is written before the command starts: once it runs, its output is copied to stdout concurrently.
	cp.publishNewProcessStart(title)
	start := time.Now()
	if err := exeCmd.Start(); err != nil {
		cp.publishStartError(err)
		cp.logf("cannot run command %q: %s", exeCmd.Path, err)
		return
	}
	if err := p.track(exeCmd.Process); err != nil {
		cp.logf("cannot track the processes started by %q: %s", exeCmd.Path, err)
	}

	pDone := make(chan struct{})
	go func() {
		err := exeCmd.Wait()
		p.exited()
		cp.publishProcessFinish(time.Since(start), err)
		close(pDone)
	}()

	cp.p, cp.pDone = p, pDone
	if cp.pWatcher != nil {
		go func() {
			cp.pWatcher <- pDone
		}()
	}
}

func (cp *CommandProxy) publishNewProcessStart(title string) {
	out := cp.stdout()
	if cp.Clear {
		if cp.started {
			escape.StyleText(out, strings.Repeat("─", outputWidth(out)), styles.TextStyle{TextColor: styles.DefaultColors.Suggestion})
			fmt.Fprintln(out)
		}
		escape.ClearScreen(out)
	}
	cp.started = true
	if cp.Header {
		escape.StyleText(out, title, styles.TextStyle{Bold: true})
		escape.UpdateProgress(out, escape.ProgressStateIndeterminate, 0)
		fmt.Fprintln(out)
	}
}

// outputWidth returns the number of columns of the terminal the output is written to, 80 if it's not a terminal.
func outputWidth(out io.Writer) int {
	if f, ok := out.(*os.File); ok {
		if size, err := vt.TerminalWinSize(f); err == nil && size.Cols > 0 {
			return size.Cols
		}
	}
	return 80
}

func (cp *CommandProxy) publishStartError(err error) {
	if !cp.Header {
		return
	}
	out := cp.stdout()
	escape.UpdateProgress(out, escape.ProgressStateDefault, 0)
	escape.StyleText(
		out,
		fmt.Sprintf("cannot start: %s", err),
		styles.TextStyle{TextColor: styles.DefaultColors.Error},
	)
	fmt.Fprintln(out)
}

func (cp *CommandProxy) publishProcessFinish(duration time.Duration, err error) {
	if !cp.Header {
		return
	}
	out := cp.stdout()
	escape.UpdateProgress(out, escape.ProgressStateDefault, 0)
	if err == nil {
		escape.StyleText(
			out,
			fmt.Sprintf("✓ done in %s", duration),
			styles.TextStyle{TextColor: styles.DefaultColors.LineAction},
		)
	} else {
		escape.StyleText(
			out,
			fmt.Sprintf("✗ %s after %s", err, duration),
			styles.TextStyle{TextColor: styles.DefaultColors.Error},
		)
	}
	fmt.Fprintln(out)
}

type exeFunc func(cmd remotectl.CommandData)

func (f exeFunc) ExecuteCommand(cmd remotectl.CommandData) { f(cmd) }
