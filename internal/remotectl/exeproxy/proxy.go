// Package exeproxy implements a server that can accept remotectl commands
// to execute some command piping it output. It can be used to maintain a single process
// for a helper pane of the editor: running tests or git diff, other executables in general
// that don't implement remotectl protocol themself.
package exeproxy

import (
	"cmp"
	"io"
	"os"
	"os/exec"
	"sync"

	"rmazur.io/chernetka/internal/logger"
	"rmazur.io/chernetka/internal/remotectl"
)

type CommandProxy struct {
	Logf logger.Func

	// Stdout and Stderr receive the output of the executed commands.
	// os.Stdout and os.Stderr are used if nil.
	Stdout, Stderr io.Writer

	pWatcher chan (<-chan struct{})

	srv *remotectl.Server

	mu    sync.Mutex
	p     *os.Process
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

func (cp *CommandProxy) exec(cmd remotectl.CommandData) {
	if cmd.Action == "quit" {
		// This may result in an error on Send on the client side as they may not get an ack.
		// It's fine as it's used in tests mainly for cleaning up.
		cp.replaceProcess(nil)
		_ = cp.srv.Close()
		return
	}

	if cmd.Action != "exe" {
		cp.logf("unknown action %s", cmd.Action)
		return
	}
	if len(cmd.Args) == 0 {
		cp.logf("exe command without args")
		return
	}

	exeCmd := exec.Command(cmd.Args[0], cmd.Args[1:]...)
	exeCmd.Stdin = os.Stdin
	exeCmd.Stdout = cmp.Or[io.Writer](cp.Stdout, os.Stdout)
	exeCmd.Stderr = cmp.Or[io.Writer](cp.Stderr, os.Stderr)
	cp.replaceProcess(exeCmd)
}

func (cp *CommandProxy) replaceProcess(exeCmd *exec.Cmd) {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	if cp.p != nil {
		_ = cp.p.Kill()
		<-cp.pDone
		cp.p, cp.pDone = nil, nil
	}

	if exeCmd == nil {
		return
	}
	if err := exeCmd.Start(); err != nil {
		cp.logf("cannot run command %q: %s", exeCmd.Path, err)
		return
	}
	p, pDone := exeCmd.Process, make(chan struct{})
	go func() {
		_ = exeCmd.Wait()
		close(pDone)
	}()
	cp.p, cp.pDone = p, pDone
	if cp.pWatcher != nil {
		go func() {
			cp.pWatcher <- pDone
		}()
	}
}

type exeFunc func(cmd remotectl.CommandData)

func (f exeFunc) ExecuteCommand(cmd remotectl.CommandData) { f(cmd) }
