//go:build !unix && !windows

package exeproxy

import (
	"os"
	"os/exec"
)

// processTree is a command process: the processes it starts are not tracked on this platform.
type processTree struct {
	p *os.Process
}

// newProcessTree configures the command before it starts.
func newProcessTree(*exec.Cmd) *processTree {
	return &processTree{}
}

// track remembers the started command process.
func (pt *processTree) track(p *os.Process) error {
	pt.p = p
	return nil
}

// kill kills only the command process itself.
func (pt *processTree) kill() {
	_ = pt.p.Kill()
}

// exited is called when the command process exits.
func (pt *processTree) exited() {}
