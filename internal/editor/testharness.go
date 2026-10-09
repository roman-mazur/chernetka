package editor

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"io"
	"path/filepath"
	"testing"
	"time"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/logger"
	"rmazur.io/chernetka/internal/vt"
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
	written     int64 // bytes written to the editor input
}

// NewTestHarness returns a harness around a fresh editor with an initialized
// command queue, so Post never blocks even without a running event loop.
func NewTestHarness() *TestHarness {
	edit := &Editor{cmdChannel: make(chan Command, 1)}
	edit.rPrefs = newRenderPrefs()

	r, w := io.Pipe()
	return &TestHarness{Editor: edit, pipeReader: r, pipeWriter: w}
}

// Run initializes the test IO and launches the Editor UI loop in a new go routine, then exits.
func (h *TestHarness) Run(t *testing.T) {
	type inOut struct {
		io.Reader
		io.Writer
	}
	term := vt.TestTerminal(80, 40, inOut{
		Reader: h.pipeReader,
		Writer: io.Discard,
	})
	t.Cleanup(func() {
		_ = term.Close()
	})

	runFinished := make(chan struct{})
	t.Cleanup(func() {
		_ = h.pipeWriter.CloseWithError(io.EOF)
		_ = h.pipeReader.CloseWithError(io.EOF)
		<-runFinished
	})
	// Extensions may already log from their goroutines: keep a logger
	// that is set, and set one before the loop starts otherwise.
	if h.Editor.Func == nil {
		h.Editor.LogEmbed = logger.Embed(logger.Prefix(t.Logf, "editor: "))
	}
	go func() {
		defer close(runFinished)
		h.Editor.Run(term)
	}()
}

// Post sends a command to the Editor queue and blocks until it's executed.
func (h *TestHarness) Post(t *testing.T, cmd Command) {
	t.Helper()

	ctx, done := context.WithTimeout(context.Background(), time.Second)
	h.Editor.Send(CommandFunc(func(e *Editor) {
		cmd.DoOnEditor(e)
		done()
	}))

	<-ctx.Done()
	if err := ctx.Err(); !errors.Is(err, context.Canceled) {
		t.Fatal("post cmd wait error:", err)
	}
}

func (h *TestHarness) DrainCommands(t *testing.T) {
	t.Helper()
	h.Post(t, CommandFunc(func(*Editor) {}))
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
// the run loop would, and waits until the commands for them are run.
func (h *TestHarness) SendInput(t *testing.T, b []byte) {
	t.Helper()
	h.WriteInput(t, b)

	// The write returns once the input is read: wait for the reader to send
	// the commands for it, then for the commands to drain.
	deadline := time.Now().Add(time.Second)
	for h.inputSent.Load() < h.written {
		if time.Now().After(deadline) {
			t.Fatalf("input is not handled: %d of %d bytes", h.inputSent.Load(), h.written)
		}
		time.Sleep(time.Millisecond)
	}
	h.Post(t, CommandFunc(func(e *Editor) {}))
}

// WriteInput writes the input without waiting for it to be handled, like a part of
// a paste the editor waits the rest of. The next SendInput waits for both.
func (h *TestHarness) WriteInput(t *testing.T, b []byte) {
	t.Helper()
	if _, err := h.pipeWriter.Write(b); err != nil {
		t.Fatal(err)
	}
	h.written += int64(len(b))
}

// SendInputSequence sends each byte of the string as a dedicated input.
func (h *TestHarness) SendInputSequence(t *testing.T, s string) {
	t.Helper()
	for _, b := range []byte(s) {
		h.SendInput(t, []byte{b})
	}
}

// RenderFrame renders the whole editor UI into out the way the run loop does on
// every change. It lets benchmarks measure a frame without the loop.
func (h *TestHarness) RenderFrame(out *bufio.Writer) {
	h.render(out)
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

// HandleInput handles the terminal input like the editor loop, but without it, so Run
// must not be called: the input is read in another goroutine the way the loop reads
// the terminal, and the commands sent for it are run in the calling goroutine, each
// followed by a render. The commands of the same input are run in the same order every
// time, so fuzzing can repeat an input. It returns once the commands sent for the input
// are run, or all the buffers are closed. The editor is shut down when the test ends,
// so the input of a harness is handled once.
func (h *TestHarness) HandleInput(t testing.TB, input []byte) {
	t.Cleanup(h.shutdown)
	read := make(chan struct{})
	go func() {
		defer close(read)
		h.readAndHandleInput(context.Background(), bufio.NewReader(bytes.NewReader(input)))
	}()

	out := bufio.NewWriter(io.Discard)
	if h.update(out) {
		return
	}
	run := func(cmd Command) (done bool) {
		cmd.DoOnEditor(h.Editor)
		return h.update(out)
	}
	for {
		select {
		case cmd := <-h.cmdChannel:
			if run(cmd) {
				return
			}
		case <-read:
			// The input is over: run the commands left in the queue.
			for {
				select {
				case cmd := <-h.cmdChannel:
					if run(cmd) {
						return
					}
				default:
					return
				}
			}
		}
	}
}

// IsolateEffects keeps the clipboard in memory, and lets the files be saved only in dir,
// until the test ends. It allows typing any input, like in fuzz tests, without
// changing the system clipboard and the files outside dir.
func IsolateEffects(t testing.TB, dir string) {
	prevClipboard, prevSave := clipboard, saveContent
	t.Cleanup(func() { clipboard, saveContent = prevClipboard, prevSave })

	clipboard = new(memoryClipboard)
	saveContent = func(doc content.Document, dst string) error {
		abs, err := filepath.Abs(dst)
		if err != nil {
			return err
		}
		if rel, err := filepath.Rel(dir, abs); err != nil || !filepath.IsLocal(rel) {
			return errors.New("saving outside the test directory")
		}
		return content.Save(doc, dst)
	}
}

type memoryClipboard struct{ text string }

func (c *memoryClipboard) Write(s string) { c.text = s }
func (c *memoryClipboard) Read() string   { return c.text }
