package editor

import (
	"io"
	"testing"

	"rmazur.io/chernetka/internal/logger"
)

// TestHarness wraps an Editor so tests — including those in extension packages —
// can drive it without the terminal event loop. It exposes the internals (the
// active buffer, the command queue, input dispatch) that the production Editor
// keeps private to Run.
type TestHarness struct {
	*Editor

	runFinished chan struct{}
	pipeReader  *io.PipeReader
	pipeWriter  *io.PipeWriter
}

// NewTestHarness returns a harness around a fresh editor with an initialized
// command queue, so Post never blocks even without a running event loop.
func NewTestHarness() *TestHarness {
	edit := &Editor{cmdChannel: make(chan Command, 1)}

	r, w := io.Pipe()
	return &TestHarness{Editor: edit, pipeReader: r, pipeWriter: w}
}

func (h *TestHarness) Run(t *testing.T) {
	inOut := InOut{
		Writer: io.Discard,
		Reader: h.pipeReader,
	}

	runFinished := make(chan struct{})
	t.Cleanup(func() {
		_ = h.pipeWriter.CloseWithError(io.EOF)
		_ = h.pipeReader.CloseWithError(io.EOF)
		<-runFinished
	})
	go func() {
		defer close(runFinished)
		h.Editor.Run(&inOut, logger.Prefix(t.Logf, "editor: "))
	}()
}

// Commands returns the channel onto which Post delivers commands. Receiving from
// it lets a test observe commands the editor posts asynchronously.
func (h *TestHarness) Commands() <-chan Command { return h.cmdChannel }

// SetMode switches the active buffer to the given mode.
func (h *TestHarness) SetMode(m Mode) {
	if b := h.Top(); b != nil {
		SwitchMode(m).DoOnBuffer(b, h.rPrefs)
	}
}

// MoveCursorToLineEnd places the cursor just past the last character of the
// current line, where a completion request would typically originate.
func (h *TestHarness) MoveCursorToLineEnd() {
	if b := h.Top(); b != nil {
		MoveEnd.DoOnBuffer(b, h.rPrefs)
	}
}

// SendInput feeds raw input bytes through the editor's input handler, exactly as
// the run loop would.
func (h *TestHarness) SendInput(t *testing.T, b []byte) {
	t.Helper()
	_, err := h.pipeWriter.Write(b)
	if err != nil {
		t.Fatal(err)
	}
}
