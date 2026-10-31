package exeproxy

import (
	"os"
	"os/exec"
	"syscall"

	"golang.org/x/sys/windows"
)

// processTree is a command together with the processes it starts: they run in a job object.
// The processes the command starts before it is assigned to the job are not tracked.
type processTree struct {
	p   *os.Process
	job windows.Handle
}

// newProcessTree configures the command before it starts.
func newProcessTree(*exec.Cmd) *processTree {
	return &processTree{}
}

// track puts the started command process to a new job.
// If it fails, kill stops only the process itself.
func (pt *processTree) track(p *os.Process) error {
	pt.p = p
	job, err := windows.CreateJobObject(nil, nil)
	if err != nil {
		return err
	}
	h, err := windows.OpenProcess(windows.PROCESS_SET_QUOTA|windows.PROCESS_TERMINATE, false, uint32(p.Pid))
	if err != nil {
		_ = windows.CloseHandle(job)
		return err
	}
	defer windows.CloseHandle(h)
	if err := windows.AssignProcessToJobObject(job, h); err != nil {
		_ = windows.CloseHandle(job)
		return err
	}
	pt.job = job
	return nil
}

// interrupt kills the command together with the processes it started:
// the console control events cannot target them alone.
func (pt *processTree) interrupt() {
	pt.kill()
}

// kill kills the command together with the processes it started.
func (pt *processTree) kill() {
	if pt.job == 0 {
		_ = pt.p.Kill()
		return
	}
	_ = windows.TerminateJobObject(pt.job, 1)
	_ = windows.CloseHandle(pt.job)
	pt.job = 0
}

// exited is called when the command process exits.
func (pt *processTree) exited() {}
