package editor

import (
	"bytes"
	"io"
	"testing"
	"time"

	"rmazur.io/chernetka/internal/editor/escape"
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
	edit := &Editor{cmdChannel: make(chan Command, 1), rPrefs: newRenderPrefs()}

	r, w := io.Pipe()
	return &TestHarness{Editor: edit, pipeReader: r, pipeWriter: w}
}

// Run initializes the test IO and launches the Editor UI loop in a new go routine, then exits.
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

// Post sends a command to the Editor queue and blocks until it's executed.
func (h *TestHarness) Post(t *testing.T, cmd Command) {
	t.Helper()

	done := make(chan struct{})
	h.Editor.Post(CommandFunc(func(e *Editor) {
		cmd.DoOnEditor(e)
		close(done)
	}))

	select {
	case <-done:
		return
	case <-time.After(time.Second):
		t.Fatalf("Post timed out")
	}
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

	// Wait for commands to drain after this.
	h.Post(t, CommandFunc(func(e *Editor) {}))
}

// RenderBuffer renders the active buffer the way the run loop would and returns
// the terminal output, escape sequences included. The window is sized to fit
// the whole content, and line numbers and the current line background are
// turned off so that nothing competes with the syntax colors. It lets an
// extension package test print colorized output for inspection by eye.
func (h *TestHarness) RenderBuffer(width int) string {
	b := h.Top()
	if b == nil {
		return ""
	}
	b.w, b.h = width, b.Content.Len()+1
	b.hideLineNumbers = true
	b.noCurrentLineHL = true

	var out bytes.Buffer
	b.Render(&out, &h.rPrefs)
	return out.String()
}

// TokenLegend renders the name of every token type in the color the theme gives
// it, so a test log can show what each color in the rendered output means.
func TokenLegend() string {
	var out bytes.Buffer
	for token := TtNothing; token <= TtQuote; token++ {
		escape.ColorText(&out, token.String(), colors.ColorForTokenType(token), nil)
		out.WriteString("  ")
	}
	return out.String()
}
