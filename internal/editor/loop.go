package editor

import (
	"bufio"
	"context"
	"io"
	"iter"
	"time"

	"rmazur.io/chernetka/internal/vt"
	"rmazur.io/chernetka/internal/vt/escape"
)

func (e *Editor) Run(t vt.Terminal) {
	start := time.Now()
	defer func() {
		e.Logf("session done %s", time.Since(start))
		for _, ext := range e.x {
			if c, ok := ext.(io.Closer); ok {
				_ = c.Close()
			}
		}
		if e.watcher != nil {
			_ = e.watcher.Close()
		}
	}()

	e.term = t
	configureTerminal(t)
	defer e.CloseAndLog(t, "term")

	e.rPrefs = newRenderPrefs()

	if e.cmdChannel == nil {
		e.cmdChannel = make(chan Command, 1)
	}

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	go e.readAndHandleInput(ctx, bufio.NewReader(t))
	go e.handleWindowChange(ctx, t.WindowSizeChanges())

	out := bufio.NewWriter(t)

	var lastRenderTime time.Time
	const renderDelay = 10 * time.Millisecond

	e.renderRequested = true
	for {
		// Close the current buffer if necessary.
		if e.quitRequested {
			if e.pop() {
				break // All done.
			}
			e.quitRequested = false
			e.renderRequested = true
		}

		// Render.
		if e.renderRequested {
			e.render(out)
			e.renderRequested = false
		}
		lastRenderTime = time.Now()

		// Handle commands, including input.
	loop:
		for {
			select {
			case cmd := <-e.cmdChannel:
				cmd.DoOnEditor(e)
			default:
				if passed := time.Since(lastRenderTime); passed < renderDelay {
					time.Sleep(renderDelay - passed)
				}
				break loop
			}
		}
	}
}

func (e *Editor) render(out *bufio.Writer) {
	topBuf := e.Top()
	if topBuf == nil {
		return
	}
	defer out.Flush()

	resetSyncOutput := escape.SyncOutput(out)
	defer resetSyncOutput()

	showCursor := escape.HideCursor(out, topBuf.mode == ModeInsert || e.status.cmd != nil)
	defer showCursor()

	// The layout needs the status bar height.
	e.status.buf = topBuf
	for buf := range e.layout() {
		buf.clampCursor(e.rPrefs.TabSize)
		escape.MoveTopLeft(out) // TODO: this works with one active buffer on top.
		buf.Render(out, &e.rPrefs)
		buf.noKeyboard = false
	}

	e.status.Render(out)

	if e.status.cmd != nil {
		e.status.RenderCursorPosition(out)
	} else {
		topBuf.RenderCursorPosition(out, &e.rPrefs)
	}

	e.mouseHandler.render(out)
}

func configureTerminal(t vt.Terminal) {
	t.Configure(
		escape.EnableAlternativeBuffer,
		escape.DisableLineWrapping,
		escape.EnableBracketedPasteMode,
		escape.EnableMouse,
		escape.EnableFocusReporting,
	)
}

// handleWindowChange reads signals typically wired to SIGWINCH and propagates a layout request.
// It ensures these signals are propagated at most every 10ms.
func (e *Editor) handleWindowChange(ctx context.Context, s <-chan struct{}) {
	const maxFrequency = 10 * time.Millisecond
	var (
		lastTime  time.Time
		timerChan <-chan time.Time
		timer     time.Timer
	)
	for {
		select {
		case _, ok := <-s:
			if !ok {
				return
			}
			if time.Since(lastTime) > maxFrequency {
				lastTime = time.Now()
				e.Send(commandRequestRender)
			} else if timerChan == nil {
				timer.Reset(maxFrequency)
				timerChan = timer.C
			}

		case <-timerChan:
			e.Send(commandRequestRender)
			timerChan = nil

		case <-ctx.Done():
			return
		}
	}
}

func (e *Editor) RequestRender() {
	e.renderRequested = true
}

// layout selects the next buffer to render configuring its dimensions.
func (e *Editor) layout() iter.Seq[*Buffer] {
	state := layoutState{e}
	return state.Pass()
}

type layoutState struct {
	editor *Editor
}

func (lps *layoutState) resolveWindowSize() (w int, h int) {
	if lps.editor.term == nil {
		return 80, 40 // Tests only.
	}
	size, err := lps.editor.term.Size()
	if err != nil || (size.Cols == 0 && size.Rows == 0) {
		return 80, 42 // TODO: Resolve window size issues on Windows.
	}
	return size.Cols, size.Rows
}

func (lps *layoutState) Pass() iter.Seq[*Buffer] {
	done := false
	w, h := lps.resolveWindowSize()
	return func(yield func(*Buffer) bool) {
		if done {
			return
		}
		// TODO: consider rendering multiple buffers.
		buf := lps.editor.Top()
		buf.w = w
		buf.h = max(h-lps.editor.status.Height(), 0)

		done = true
		yield(buf)
	}
}
