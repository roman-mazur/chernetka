//go:build darwin || linux

package emulation

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"
	"testing"
	"time"
)

// Monkey types random input into a terminal app, checking that it neither crashes nor hangs.
// The input is decided by Choices, made from the input of a fuzz test, so a failing input
// can be typed again (the timing of the app still differs from run to run).
// The app is expected to write to the terminal in response to the input: when it does not
// for HangTimeout, it's considered hung.
type Monkey struct {
	Cmd        *exec.Cmd // the app to run, its standard streams are set by the monkey
	Cols, Rows int       // the terminal size, 80x40 by default

	Choices     *Choices      // decide the input, typing stops when they're exhausted
	Duration    time.Duration // the longest time to type, unlimited by default
	HangTimeout time.Duration // 10s by default

	// Snippets are the inputs meaningful for the app, like its commands, typed along with
	// the random ones. They must not have the banned characters.
	Snippets []string
	Ban      Ban
	// Quit is the input quitting the app after the run. It's written until the app exits.
	Quit string

	// Watch lists the processes that must keep running, like the helpers of the app.
	Watch []*Process
	// OnFailure is called on a failure to report more, like the log of the app.
	OnFailure func(t testing.TB)

	mu     sync.Mutex
	recent []string // the latest inputs for the failure reports
	count  int
}

const recentInputs = 50

// Run starts the app and types random input until the Choices are exhausted
// or for the Duration, then quits the app.
// It fails the test if the app or a watched process crashes, the app hangs, or it quits
// with an error. The failures report the stderr of the process, which has the panic or
// the goroutines of a hung Go app, and the latest input.
func (m *Monkey) Run(t testing.TB) {
	t.Helper()
	cols, rows := m.Cols, m.Rows
	if cols == 0 || rows == 0 {
		cols, rows = 80, 40
	}
	hangTimeout := m.HangTimeout
	if hangTimeout == 0 {
		hangTimeout = 10 * time.Second
	}
	if m.Choices == nil {
		t.Fatal("no choices to make the input")
	}
	for _, s := range m.Snippets {
		if m.Ban.contains(s) {
			t.Fatalf("snippet %q has a banned character", s)
		}
	}

	p := Start(t, m.Cmd, cols, rows)
	p.WaitReady(t, hangTimeout)

	fail := func(format string, args ...any) {
		t.Helper()
		t.Logf("recent input (oldest first):\n%s", m.recentInput())
		if m.OnFailure != nil {
			m.OnFailure(t)
		}
		t.Fatalf(format, args...)
	}
	checkRunning := func(p *Process) {
		t.Helper()
		select {
		case <-p.exited:
			t.Logf("%s stderr:\n%s", p.name, p.Stderr())
			fail("%s exited: %v", p.name, p.err)
		default:
		}
	}
	hung := func(what string) {
		t.Helper()
		p.dump(t)
		fail("%s %s for %s", p.name, what, hangTimeout)
	}

	in := newInputs(m.Choices, m.Ban, cols, rows, m.Snippets)
	done := make(chan error, 1)
	stop := make(chan struct{})
	go func() { done <- m.typeInput(p.pty, in, stop) }()

	var deadline <-chan time.Time
	if m.Duration > 0 {
		deadline = time.After(m.Duration)
	}
	typed := false
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()
loop:
	for {
		select {
		case <-ticker.C:
			checkRunning(p)
			for _, w := range m.Watch {
				checkRunning(w)
			}
			if p.sinceOutput() > hangTimeout {
				hung("does not respond")
			}
		case err := <-done:
			if err != nil {
				// The input cannot be written once the app exits: wait for its exit to be reported.
				p.waitExit(2 * time.Second)
				checkRunning(p)
				fail("cannot type: %v", err)
			}
			typed = true
			break loop
		case <-deadline:
			break loop
		}
	}
	close(stop)
	if !typed {
		select {
		case err := <-done:
			if err != nil {
				p.waitExit(2 * time.Second)
				checkRunning(p)
				fail("cannot type: %v", err)
			}
		case <-time.After(hangTimeout):
			hung("does not read the input")
		}
	}
	m.mu.Lock()
	t.Logf("typed %d inputs", m.count)
	m.mu.Unlock()

	for _, w := range m.Watch {
		checkRunning(w)
	}
	checkRunning(p)
	m.record(m.Quit)
	if err := p.quit(t, m.Quit, hangTimeout); err != nil {
		fail("%s", err)
	}
}

// typeInput writes random input to w until the choices are exhausted or stop is closed.
func (m *Monkey) typeInput(w io.Writer, in *inputs, stop <-chan struct{}) error {
	for !in.rnd.Exhausted() {
		select {
		case <-stop:
			return nil
		default:
		}
		input := in.next()
		m.record(input)
		if _, err := io.WriteString(w, input); err != nil {
			return err
		}
		// Type fast, but let the app render sometimes and tell the clicks apart.
		if in.rnd.IntN(5) == 0 {
			time.Sleep(time.Duration(in.rnd.IntN(50)) * time.Millisecond)
		}
	}
	return nil
}

func (m *Monkey) record(input string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.recent = append(m.recent, input)
	if len(m.recent) > recentInputs {
		m.recent = m.recent[1:]
	}
	m.count++
}

func (m *Monkey) recentInput() string {
	m.mu.Lock()
	defer m.mu.Unlock()
	var sb strings.Builder
	for _, s := range m.recent {
		fmt.Fprintf(&sb, "%q\n", s)
	}
	return sb.String()
}
