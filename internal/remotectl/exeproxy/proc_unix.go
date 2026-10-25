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

// setProcessGroup makes the command run in its own process group, so that killProcess stops its children too.
// If stdin is a terminal, the group is put to the terminal foreground: the command can read the terminal
// (e.g., a pager) and receives signals like Ctrl+C instead of the proxy.
func setProcessGroup(cmd *exec.Cmd) (foreground bool) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return false
	}
	cmd.SysProcAttr.Foreground = true
	cmd.SysProcAttr.Ctty = int(os.Stdin.Fd())
	return true
}

// killProcess kills the process started with setProcessGroup together with its children.
func killProcess(p *os.Process) {
	_ = syscall.Kill(-p.Pid, syscall.SIGKILL)
}

// reclaimForeground puts the proxy process group back to the terminal foreground after the command exits.
func reclaimForeground() {
	// A background process changing the foreground group gets SIGTTOU.
	signal.Ignore(syscall.SIGTTOU)
	defer signal.Reset(syscall.SIGTTOU)
	_ = unix.IoctlSetPointerInt(int(os.Stdin.Fd()), unix.TIOCSPGRP, unix.Getpgrp())
}
