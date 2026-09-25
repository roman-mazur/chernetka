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
	Size() (WindowSize, error)
	Configure(f ...escape.ConfigFunc)
}

// WindowSize describes the terminal window dimensions.
// The size in pixels is zero if the terminal does not report it.
type WindowSize struct {
	Cols, Rows     int
	XPixel, YPixel int
}

// CellSize returns the size of one character cell in pixels.
// ok is false if the terminal does not report its size in pixels.
func (ws WindowSize) CellSize() (w, h float64, ok bool) {
	if ws.Cols <= 0 || ws.Rows <= 0 || ws.XPixel <= 0 || ws.YPixel <= 0 {
		return 0, 0, false
	}
	return float64(ws.XPixel) / float64(ws.Cols), float64(ws.YPixel) / float64(ws.Rows), true
}

// SystemTerminal returns the terminal the process is attached to.
// Input is read from stdin and output is written to stderr or stdout, when they are terminals.
// Otherwise, the controlling terminal device is opened.
// The input side is switched to raw mode until the terminal is closed.
func SystemTerminal() (Terminal, error) {
	st := &sysTerminal{
		in:  findTerminalFile(os.Stdin),
		out: findTerminalFile(os.Stderr, os.Stdout),
	}
	if st.in == nil || st.out == nil {
		in, out, closeTTY, err := openTTY()
		if err != nil {
			return nil, err
		}
		st.cleanup = append(st.cleanup, closeTTY)
		if st.in == nil {
			st.in = in
		}
		if st.out == nil {
			st.out = out
		}
	}

	// Enable raw mode.
	fd := int(st.in.Fd())
	state, err := term.MakeRaw(fd)
	if err != nil {
		_ = st.Close()
		return nil, fmt.Errorf("enable raw mode: %w", err)
	}
	st.cleanup = append(st.cleanup, func() {
		_ = term.Restore(fd, state)
	})

	return st, nil
}

func findTerminalFile(files ...*os.File) *os.File {
	for _, f := range files {
		if term.IsTerminal(int(f.Fd())) {
			return f
		}
	}
	return nil
}

// sysTerminal reads from in and writes to out.
// On Unix both are usually the same tty, but on Windows the console input and output are separate handles:
// raw mode applies to the input handle, while the window size is queried from the output one.
type sysTerminal struct {
	in, out *os.File
	cleanup []func()
}

func (st *sysTerminal) Read(p []byte) (int, error)  { return st.in.Read(p) }
func (st *sysTerminal) Write(p []byte) (int, error) { return st.out.Write(p) }

func (st *sysTerminal) Configure(opts ...escape.ConfigFunc) {
	for _, opt := range opts {
		st.cleanup = append(st.cleanup, opt(st))
	}
}

func (st *sysTerminal) WindowSizeChanges() <-chan struct{} {
	return windowChangeSignal()
}

func (st *sysTerminal) Size() (WindowSize, error) {
	return windowSize(int(st.out.Fd()))
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
// The terminal does not report its size in pixels, set SizeFunc to change it.
func TestTerminal(w, h int, rw io.ReadWriter) *MockTerminal {
	return &MockTerminal{
		ReadWriter: rw,
		SizeFunc: func() (WindowSize, error) {
			return WindowSize{Cols: w, Rows: h}, nil
		},
	}
}

type MockTerminal struct {
	io.ReadWriter
	SizeFunc func() (WindowSize, error)
}

func (t *MockTerminal) WindowSizeChanges() <-chan struct{} { return nil }
func (t *MockTerminal) Size() (WindowSize, error)          { return t.SizeFunc() }
func (t *MockTerminal) Configure(...escape.ConfigFunc)     {}

func (t *MockTerminal) Close() error {
	if c, ok := t.ReadWriter.(io.Closer); ok {
		return c.Close()
	}
	return nil
}
