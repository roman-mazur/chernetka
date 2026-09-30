package editor

import (
	"bufio"
	"context"
	"errors"
	"io"
	"path/filepath"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/input"
)

func (e *Editor) readAndHandleInput(ctx context.Context, terminal *bufio.Reader) {
	var (
		inBuf   [64]byte
		pending []byte // an incomplete sequence at the end of the previous read
		in      = &countingReader{r: terminal}
	)
	for ctx.Err() == nil {
		n, err := in.Read(inBuf[:])
		if err != nil {
			e.handleInputError(err)
			break
		}
		data := append(pending, inBuf[:n]...)
		if debugInput {
			e.Logf("input: %v", data)
		}
		pending, err = e.sendInput(data, in)
		if err != nil {
			e.handleInputError(err)
			break
		}
		e.inputSent.Store(in.n)
		if len(pending) > maxPendingInput {
			e.Logf("dropping unknown input: %v", pending)
			pending = nil
		}
	}
}

// countingReader counts the bytes read from r.
type countingReader struct {
	r io.Reader
	n int64
}

func (cr *countingReader) Read(p []byte) (int, error) {
	n, err := cr.r.Read(p)
	cr.n += int64(n)
	return n, err
}

// maxPendingInput limits the incomplete sequence kept until the next read.
const maxPendingInput = 256

// sendInput sends the keys, mouse events, and pastes read from the terminal to the editor loop.
// It returns the incomplete sequence at the end of the input to be completed by the next read.
// A paste is read until its end from in.
func (e *Editor) sendInput(data []byte, in io.Reader) (rest []byte, err error) {
	for len(data) > 0 {
		keys, n := input.Keys(data)
		if len(keys) > 0 {
			e.Send(CommandFunc(func(e *Editor) { e.handleKeys(keys) }))
		}
		data = data[n:]
		if len(data) == 0 {
			break
		}

		// The keys stopped before a mouse event, a paste, or an incomplete sequence.
		if input.IsMouseInput(data) {
			mouse, n, err := input.ReadMouse(data)
			if errors.Is(err, input.ErrIncomplete) {
				return data, nil
			}
			if err != nil {
				e.Logf("dropping bad mouse input %v: %s", data, err)
				return nil, nil
			}
			e.Send(CommandFunc(func(e *Editor) { e.handleMouse(mouse) }))
			data = data[n:]
			continue
		}
		pasted, detected, err := input.ConsumeClipboardPaste(data, in)
		if err != nil {
			return nil, err
		}
		if detected {
			if pasted != "" {
				e.SendBufferCmd(PasteText(pasted))
			}
			return nil, nil // The rest of the input is consumed by the paste.
		}
		return data, nil // Incomplete.
	}
	return nil, nil
}

func (e *Editor) handleInputError(err error) {
	e.Logf("input error, quitting: %s", err)
	e.Send(commandQuit)
}

// handleKeys handles the keys read at once. The keys after the one quitting the buffer are ignored.
func (e *Editor) handleKeys(keys []input.Key) {
	for _, k := range keys {
		if debugInput {
			e.Logf("key: %s", k)
		}
		if e.Top() == nil {
			return
		}
		if e.handleKey(k) {
			commandQuit(e)
			return
		}
	}
	e.renderRequested = true
}

func (e *Editor) handleMouse(data input.Mouse) {
	event := e.transformInput(data)
	if debugInput {
		e.Logf("mouse: %s", event)
	}

	buf := e.Top()
	if buf == nil {
		return
	}
	buf.noKeyboard = true

	switch event.eventType {
	case mouseEventTypeScroll:
		dir := event.Mod.SrollDirection(event.Mouse)
		e.execBufferCmd(Scroll(dir))

	case mouseEventTypeDragStart:
		if buf.CheckContentCoordinates(event.Y-1, event.X-1) {
			e.execBufferCmd(StartTextSelection)
		}
		e.renderRequested = true

	case mouseEventTypeDragEnd:
		e.execBufferCmd(StopTextSelection)

	case mouseEventTypeDoubleClick:
		if buf.engageOnDoubleClick && e.maybeEngage(buf) {
			e.renderRequested = true
			e.handleAfterEdit(buf)
			return
		}
		e.execBufferCmd(SelectWord)

	case mouseEventTypeTripleClick:
		e.execBufferCmd(SelectLine)

	case mouseEventTypeRaw:
		overContent := buf.CheckContentCoordinates(event.Y-1, event.X-1)
		if e.ensureMouseTextShape(overContent) {
			e.renderRequested = true
		}
		// Some terminals on macOS report Ctrl+click as a right click with Ctrl.
		ctrlClick := event.Pressed && event.Mod.HasCtrl() && !event.Mod.HasMotion() &&
			(event.Button == input.MouseButtonLeft || event.Button == input.MouseButtonRight)
		if !ctrlClick && (event.Button != input.MouseButtonLeft || !event.Pressed) {
			return
		}

		buf.c = buf.screenToContentPosition(data.Y-1, data.X-1, e.rPrefs.TabSize)
		if buf.selecting && event.Mod.HasMotion() {
			// Dragging over the line numbers or below the text still selects the closest content.
			buf.sel[len(buf.sel)-1].End = buf.c
		} else if !buf.selecting {
			buf.sel = nil
		}
		e.renderRequested = true
		if ctrlClick && overContent {
			e.goToDefinition(buf)
		}
	}

}

// goToDefinition moves the cursor to the definition of the symbol under it
// if an extension can find it.
func (e *Editor) goToDefinition(buf *Buffer) {
	if f, ok := FindExtData[DefinitionFinder](buf); ok {
		f.FindDefinition(e, buf.c)
	}
}

func (e *Editor) handleKey(k input.Key) (quit bool) {
	buf := e.Top()
	// Any key may edit the buffer: the extensions are notified about it once.
	defer e.handleAfterEdit(buf)

	// The keys working in any mode.
	switch k {
	// Save the changes when the terminal loses focus, so the file system is in sync
	// for the tools used outside the editor.
	case input.Of(input.FocusOut):
		e.saveTop()
		return false
	case input.Of(input.FocusIn):
		return false

	case input.Ctrl('s'):
		e.execBufferCmd(&Save{buf.Path})
		return false

	// Re-run the last engaged line action.
	case input.Ctrl('r'):
		if e.lastAction.action != nil {
			e.engage(e.lastAction.action)
		}
		return false

	// Show the file picker. The picker itself selects the next match with it.
	case input.Ctrl('o'):
		if !prompting[*quickOpen](e) {
			e.startQuickOpen(buf)
			return false
		}

	// Open the command line to type the line number to go to.
	case input.Ctrl('l'):
		e.openCmdLine(buf, "", newExPrompt)
		return false

	// Start the search. The search itself moves to the next match with it.
	case input.Ctrl('f'):
		if !prompting[*searchPrompt](e) {
			e.startSearch(buf)
			return false
		}

	// Show the changes of the current file.
	case input.Ctrl('d'):
		e.showDiff(buf)
		return false

	// Clipboard.
	case input.Ctrl('c'):
		e.execBufferCmd(ClipboardCopy)
		return false
	case input.Ctrl('v'):
		e.execBufferCmd(ClipboardPaste)
		return false
	case input.Ctrl('x'):
		e.execBufferCmd(ClipboardCut)
		return false
	}

	if e.status.cmd != nil {
		return e.cmdLineInput(k)
	}

	switch buf.mode {
	case ModeNormal:
		// Open the command line.
		switch k {
		case input.Rune(':'):
			e.openCmdLine(buf, "", newExPrompt)
			return false
		case input.Rune('/'):
			e.startSearch(buf)
			return false
		case input.Of(input.Enter):
			if fired := e.maybeEngage(buf); fired {
				return false
			}
		}
		quit = normalInput(buf, k, &e.rPrefs)
		return

	case ModeInsert:
		if !e.extHandleInsert(buf, k) {
			insertInput(buf, k, &e.rPrefs)
		}
		return false

	default:
		return false
	}
}

func (e *Editor) maybeEngage(buf *Buffer) bool {
	if action := buf.lineAction(buf.c.Line); action != nil {
		e.engage(action)
		if _, single := action.(content.LineActionSingleShot); !single {
			e.lastAction = bufferAction{buf: buf, action: action}
		}
		return true
	}
	return false
}

// engage saves the changes of the top buffer and runs the action: the actions may work
// with the files rather than the buffers, for example, running a program.
func (e *Editor) engage(action content.LineAction) {
	e.saveTop()
	action.Engage()
}

// showDiff saves the changes in buf and passes its file to the ShowDiff hook.
// Buffers without a file (scratch, piped content, directories) are ignored.
func (e *Editor) showDiff(buf *Buffer) {
	if e.ShowDiff == nil || buf.Path == "" {
		return
	}
	if _, isDir := buf.Content.(*content.FsContent); isDir {
		return // Its path is only a display name.
	}
	e.saveTop() // The diff is taken from the saved file.
	path := buf.Path
	if abs, err := filepath.Abs(path); err == nil {
		path = abs
	}
	e.ShowDiff(path)
}

func (e *Editor) extHandleInsert(buf *Buffer, k input.Key) (handled bool) {
	for _, ext := range e.x {
		if h, ok := ext.(InsertKeyHandler); ok && h.HandleInsertKey(buf, k) {
			return true
		}
	}
	return false
}
