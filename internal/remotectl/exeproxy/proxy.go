// Package exeproxy implements a server that can accept remotectl commands
// to execute some command piping it output. It can be used to maintain a single process
// for a helper pane of the editor: running tests or git diff, other executables in general
// that don't implement remotectl protocol themself.
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
	"rmazur.io/chernetka/internal/vt/escape"
)

// Endpoint returns the endpoint of the proxy serving the purpose, like running tests or showing diffs:
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
	Clear bool
	// Grace limits how long a replaced command may take to exit after an interrupt before it's killed
	// (2s by default).
	Grace time.Duration

	pWatcher chan (<-chan struct{})

	srv *remotectl.Server

	mu    sync.Mutex
	p     *processTree
	pDone chan struct{}
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
	return fmt.Sprintf("%s> %s %s", p, filepath.Base(cmd.Path), strings.Join(cmd.Args, " "))
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
	start := time.Now()
	err := exeCmd.Start()
	cp.publishNewProcessStart(title, err)
	if err != nil {
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

func (cp *CommandProxy) publishNewProcessStart(title string, err error) {
	out := cp.stdout()
	if cp.Clear {
		escape.ClearScreen(out)
	}
	if cp.Header {
		if err == nil {
			escape.StyleText(out, title, styles.TextStyle{Bold: true})
		} else {
			escape.StyleText(
				out,
				fmt.Sprintf("cannot start: %s", err),
				styles.TextStyle{TextColor: styles.DefaultColors.Error},
			)
		}
	}
}

func (cp *CommandProxy) publishProcessFinish(duration time.Duration, err error) {
	if !cp.Header {
		return
	}
	if err == nil {
		escape.StyleText(
			cp.stdout(),
			fmt.Sprintf("✓ done in %s", duration),
			styles.TextStyle{TextColor: styles.DefaultColors.LineAction},
		)
	} else {
		escape.StyleText(
			cp.stdout(),
			fmt.Sprintf("✗ %s after %s", err, duration),
			styles.TextStyle{TextColor: styles.DefaultColors.Error},
		)
	}
}

type exeFunc func(cmd remotectl.CommandData)

func (f exeFunc) ExecuteCommand(cmd remotectl.CommandData) { f(cmd) }
