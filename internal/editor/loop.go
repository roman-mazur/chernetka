package editor

import (
	"bufio"
	"context"
	"io"
	"iter"
	"slices"
	"time"

	"rmazur.io/chernetka/internal/vt"
	"rmazur.io/chernetka/internal/vt/escape"
)

func (e *Editor) Run(t vt.Terminal) {
	start := time.Now()
	defer func() {
		e.Logf("session done %s", time.Since(start))
		close(e.loopDone())
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

		// Handle commands, including input.
		e.processCommands(time.Now())
	}
}

// processCommands runs the commands as they come until the render delay since lastRenderTime
// passes, or until a quit is requested. With renderEachCommand, it waits for a single command
// and runs it, so the loop renders after each command in the same order every time.
func (e *Editor) processCommands(lastRenderTime time.Time) {
	if e.renderEachCommand {
		(<-e.cmdChannel).DoOnEditor(e)
		return
	}

	const renderDelay = 10 * time.Millisecond

	// Waiting for the commands instead of sleeping keeps the input flowing: the queue has
	// a single slot, and the input reader blocks until the loop takes the queued command.
	// The timer is reused, so the loop does not allocate a new one for every render.
	delay := renderDelay - time.Since(lastRenderTime)
	if e.renderTimer == nil {
		e.renderTimer = time.NewTimer(delay)
	} else {
		e.renderTimer.Reset(delay)
	}
	for !e.quitRequested {
		select {
		case cmd := <-e.cmdChannel:
			cmd.DoOnEditor(e)
		case <-e.renderTimer.C:
			return
		}
	}
}

func (e *Editor) render(out *bufio.Writer) {
	topBuf := e.Top()
	if topBuf == nil {
		return
	}
	defer func() { _ = out.Flush() }()

	resetSyncOutput := escape.SyncOutput(out)
	defer resetSyncOutput.Undo()

	hideCursorReset := escape.HideCursor(out, topBuf.mode == ModeInsert || e.status.cmd != nil)
	defer hideCursorReset.Undo()

	// The layout needs the status bar height.
	e.status.buf = topBuf

	for buf := range e.layout() {
		buf.clampCursor(e.rPrefs.TabSize)
		escape.SetCursorPosition(out, buf.y+1, 1)
		buf.Render(out, &e.rPrefs)
		buf.noKeyboard = false
	}

	escape.SetCursorPosition(out, topBuf.y+topBuf.h+1, 1)
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
		timer     *time.Timer // created on the first use: Reset panics on a zero Timer
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
				if timer == nil {
					timer = time.NewTimer(maxFrequency)
				} else {
					timer.Reset(maxFrequency)
				}
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
	w, h := lps.resolveWindowSize()

	var (
		lBufs [2]*Buffer
		res   = lBufs[:0]
	)
	availableHeight := func() int {
		rh := h - lps.editor.status.Height()
		if rh <= 0 {
			return 0
		}
		for i := range res {
			rh -= res[i].h
		}
		return max(rh, 0)
	}

	// The height is resolved before the buffer is added to res: its h is from the previous pass.
	if buf := lps.editor.toolBuf; buf != nil {
		buf.w = w
		buf.h = min(availableHeight()/4, 7)
		res = append(res, buf)
	}
	if buf := lps.editor.Top(); buf != nil {
		buf.w = w
		buf.h = availableHeight()
		buf.y = 0
		res = append(res, buf)

		// The tool buffer is at the bottom, below the status bar of the top buffer.
		if tb := lps.editor.toolBuf; tb != nil {
			tb.y = buf.h + lps.editor.status.Height()
		}
	}

	return func(yield func(*Buffer) bool) {
		for _, re := range slices.Backward(res) {
			if !yield(re) {
				return
			}
		}
	}
}
