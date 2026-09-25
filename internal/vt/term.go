package vt

import (
	"fmt"
	"io"
	"os"
	"slices"

	"golang.org/x/term"
	"rmazur.io/chernetka/internal/vt/escape"
)

type Terminal interface {
	io.ReadWriteCloser
	WindowSizeChanges() <-chan struct{}
	Size() (width, height int, err error)
	Configure(f ...escape.ConfigFunc)
}

func SystemTerminal() (Terminal, error) {
	tf := findTerminalFile(os.Stderr, os.Stdout, os.Stdin)
	shouldClose := false
	if tf == nil {
		ttyFile, err := os.Open("/dev/tty")
		if err != nil {
			return nil, err
		}
		tf = findTerminalFile(ttyFile)
		if tf == nil {
			_ = ttyFile.Close()
			return nil, fmt.Errorf("/dev/tty is not a terminal")
		}
		shouldClose = true
	}

	sysTerm := &sysTerminal{File: tf, fd: int(tf.Fd())}
	if shouldClose {
		sysTerm.cleanup = append(sysTerm.cleanup, func() {
			_ = tf.Close()
		})
	}

	// Enable raw mode.
	sysTerm.Configure(func(io.Writer) (restore func()) {
		state, err := term.MakeRaw(sysTerm.fd)
		if err != nil {
			_ = sysTerm.Close()
			return
		}
		restore = func() {
			_ = term.Restore(sysTerm.fd, state)
		}
		return
	})

	return sysTerm, nil
}

func findTerminalFile(files ...*os.File) *os.File {
	for _, f := range files {
		if term.IsTerminal(int(f.Fd())) {
			return f
		}
	}
	return nil
}

type sysTerminal struct {
	*os.File
	cleanup []func()
	fd      int
}

func (st *sysTerminal) Configure(opts ...escape.ConfigFunc) {
	for _, opt := range opts {
		st.cleanup = append(st.cleanup, opt(st))
	}
}

func (st *sysTerminal) WindowSizeChanges() <-chan struct{} {
	return windowChangeSignal()
}

func (st *sysTerminal) Size() (width, height int, err error) {
	return term.GetSize(st.fd)
}

func (st *sysTerminal) Close() error {
	for _, f := range slices.Backward(st.cleanup) {
		if f != nil {
			f()
		}
	}
	return nil
}

// TestTerminal produces a test Terminal that never emits a window change event.
// Reader and writer are provided in the arguments. If it's also an io.Closer, it will be closed in Close.
func TestTerminal(w, h int, rw io.ReadWriter) *MockTerminal {
	return &MockTerminal{
		ReadWriter: rw,
		SizeFunc: func() (int, int, error) {
			return w, h, nil
		},
	}
}

type MockTerminal struct {
	io.ReadWriter
	SizeFunc func() (int, int, error)
}

func (t *MockTerminal) WindowSizeChanges() <-chan struct{}   { return nil }
func (t *MockTerminal) Size() (width, height int, err error) { return t.SizeFunc() }
func (t *MockTerminal) Configure(...escape.ConfigFunc)       {}

func (t *MockTerminal) Close() error {
	if c, ok := t.ReadWriter.(io.Closer); ok {
		return c.Close()
	}
	return nil
}
