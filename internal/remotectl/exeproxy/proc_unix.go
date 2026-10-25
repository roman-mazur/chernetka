//go:build unix

package exeproxy

import (
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"golang.org/x/sys/unix"
	"golang.org/x/term"
)

// processTree is a command together with the processes it starts: they run in a separate process group.
// If stdin is a terminal, the group is put to the terminal foreground: the command can read the terminal
// (e.g., a pager) and receives signals like Ctrl+C instead of the proxy.
type processTree struct {
	pid        int
	foreground bool
}

// newProcessTree configures the command before it starts.
func newProcessTree(cmd *exec.Cmd) *processTree {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return &processTree{}
	}
	cmd.SysProcAttr.Foreground = true
	cmd.SysProcAttr.Ctty = int(os.Stdin.Fd())
	return &processTree{foreground: true}
}

// track remembers the started command process.
func (pt *processTree) track(p *os.Process) error {
	pt.pid = p.Pid
	return nil
}

// kill kills the command together with the processes it started.
func (pt *processTree) kill() {
	_ = syscall.Kill(-pt.pid, syscall.SIGKILL)
}

// exited is called when the command process exits.
// It puts the proxy process group back to the terminal foreground.
func (pt *processTree) exited() {
	if !pt.foreground {
		return
	}
	// A background process changing the foreground group gets SIGTTOU.
	signal.Ignore(syscall.SIGTTOU)
	defer signal.Reset(syscall.SIGTTOU)
	_ = unix.IoctlSetPointerInt(int(os.Stdin.Fd()), unix.TIOCSPGRP, unix.Getpgrp())
}
