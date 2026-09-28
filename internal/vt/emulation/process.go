//go:build darwin || linux

package emulation

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"sync/atomic"
	"syscall"
	"testing"
	"time"
)

// Process is a terminal app running in a pseudo-terminal.
type Process struct {
	name   string
	cmd    *exec.Cmd
	pty    *os.File
	stderr syncBuffer

	firstOutput chan struct{}
	lastOutput  atomic.Int64 // Unix nanoseconds

	exited chan struct{}
	err    error // the exit error, set before exited is closed
}

// Start runs the command in a new pseudo-terminal of the given size, and in a new session:
// it cannot reach the terminal of the test. Its stderr is kept apart from the terminal,
// the panics are written there. The process with its session is killed when the test ends.
func Start(t testing.TB, cmd *exec.Cmd, cols, rows int) *Process {
	t.Helper()
	pty, tty, err := OpenPTY(cols, rows)
	if err != nil {
		t.Fatalf("cannot open a pseudo-terminal: %s", err)
	}
	p := &Process{
		name:        filepath.Base(cmd.Path),
		cmd:         cmd,
		pty:         pty,
		firstOutput: make(chan struct{}),
		exited:      make(chan struct{}),
	}
	cmd.Stdin, cmd.Stdout, cmd.Stderr = tty, tty, &p.stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true, Setctty: true, Ctty: 0}
	if cmd.WaitDelay == 0 {
		cmd.WaitDelay = time.Second // The children may keep stderr open.
	}
	err = cmd.Start()
	_ = tty.Close() // The process has it: reading the output fails once the process exits.
	if err != nil {
		_ = pty.Close()
		t.Fatalf("cannot start %s: %s", p.name, err)
	}

	go func() {
		p.err = cmd.Wait()
		close(p.exited)
	}()
	go p.readOutput()
	t.Cleanup(func() {
		_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		<-p.exited
		_ = pty.Close()
	})
	return p
}

// readOutput discards the terminal output, tracking when it's written.
func (p *Process) readOutput() {
	buf := make([]byte, 32<<10)
	for first := true; ; first = false {
		if _, err := p.pty.Read(buf); err != nil {
			return
		}
		p.lastOutput.Store(time.Now().UnixNano())
		if first {
			close(p.firstOutput)
		}
	}
}

// Write writes the input to the terminal.
func (p *Process) Write(input []byte) (int, error) { return p.pty.Write(input) }

// Stderr returns what the process has written to stderr so far.
func (p *Process) Stderr() string { return p.stderr.String() }

// WaitReady waits for the process to write to the terminal, which it does once it's ready.
func (p *Process) WaitReady(t testing.TB, timeout time.Duration) {
	t.Helper()
	select {
	case <-p.firstOutput:
	case <-p.exited:
		t.Fatalf("%s exited: %v\n%s", p.name, p.err, p.Stderr())
	case <-time.After(timeout):
		t.Fatalf("%s does not start", p.name)
	}
}

// CheckRunning fails the test if the process has exited, reporting its stderr.
func (p *Process) CheckRunning(t testing.TB) {
	t.Helper()
	select {
	case <-p.exited:
		t.Logf("%s stderr:\n%s", p.name, p.Stderr())
		t.Fatalf("%s exited: %v", p.name, p.err)
	default:
	}
}

// Quit writes the input repeatedly until the process exits. It fails the test if the process
// does not exit within the timeout or exits with an error.
func (p *Process) Quit(t testing.TB, input string, timeout time.Duration) {
	t.Helper()
	p.CheckRunning(t)
	if err := p.quit(t, input, timeout); err != nil {
		t.Fatal(err)
	}
}

// quit writes the input until the process exits, reporting its stderr.
// A hung process is made to report its goroutines.
func (p *Process) quit(t testing.TB, input string, timeout time.Duration) error {
	t.Helper()
	go func() {
		for {
			if _, err := io.WriteString(p.pty, input); err != nil {
				return // The process has exited.
			}
			select {
			case <-p.exited:
				return
			case <-time.After(10 * time.Millisecond):
			}
		}
	}()
	select {
	case <-p.exited:
	case <-time.After(timeout):
		p.dump(t)
		return fmt.Errorf("%s does not quit for %s", p.name, timeout)
	}
	if stderr := p.Stderr(); stderr != "" {
		t.Logf("%s stderr:\n%s", p.name, stderr)
	}
	if p.err != nil {
		return fmt.Errorf("%s quit with an error: %w", p.name, p.err)
	}
	return nil
}

// dump makes a Go process write its goroutines to stderr, and reports them.
func (p *Process) dump(t testing.TB) {
	t.Helper()
	_ = p.cmd.Process.Signal(syscall.SIGQUIT)
	p.waitExit(5 * time.Second)
	t.Logf("%s stderr:\n%s", p.name, p.Stderr())
}

// waitExit waits for the process to exit for up to d.
func (p *Process) waitExit(d time.Duration) {
	select {
	case <-p.exited:
	case <-time.After(d):
	}
}

func (p *Process) sinceOutput() time.Duration {
	return time.Since(time.Unix(0, p.lastOutput.Load()))
}

// syncBuffer is a bytes.Buffer safe to read while the process writes to it.
type syncBuffer struct {
	mu  sync.Mutex
	buf bytes.Buffer
}

func (b *syncBuffer) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.Write(p)
}

func (b *syncBuffer) String() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.String()
}
