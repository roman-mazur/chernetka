//go:build !unix

package exeproxy

import (
	"os"
	"os/exec"
)

// setProcessGroup does nothing: the command children are not tracked on this platform.
func setProcessGroup(*exec.Cmd) (foreground bool) { return false }

// killProcess kills only the process itself on this platform.
func killProcess(p *os.Process) {
	_ = p.Kill()
}

func reclaimForeground() {}
