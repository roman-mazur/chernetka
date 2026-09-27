package editor

import (
	"bufio"
	"context"
	"errors"
	"io"
	"path/filepath"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/inputs"
)

func (e *Editor) readAndHandleInput(ctx context.Context, in *bufio.Reader) {
	var (
		inBuf   [64]byte
		pending []byte // an incomplete sequence at the end of the previous read
	)
	for ctx.Err() == nil {
		n, err := in.Read(inBuf[:])
		if err != nil {
			e.handleInputError(err)
			break
		}
		input := append(pending, inBuf[:n]...)
		if debugInput {
			e.Logf("input: %v", input)
		}
		pending, err = e.sendInput(input, in)
		if err != nil {
			e.handleInputError(err)
			break
		}
		if len(pending) > maxPendingInput {
			e.Logf("dropping unknown input: %v", pending)
			pending = nil
		}
	}
}

// maxPendingInput limits the incomplete sequence kept until the next read.
const maxPendingInput = 256

// sendInput sends the keys, mouse events, and pastes read from the terminal to the editor loop.
// It returns the incomplete sequence at the end of the input to be completed by the next read.
// A paste is read until its end from in.
func (e *Editor) sendInput(input []byte, in io.Reader) (rest []byte, err error) {
	for len(input) > 0 {
		keys, n := inputs.Keys(input)
		if len(keys) > 0 {
			e.Send(CommandFunc(func(e *Editor) { e.handleKeys(keys) }))
		}
		input = input[n:]
		if len(input) == 0 {
			break
		}

		// The keys stopped before a mouse event, a paste, or an incomplete sequence.
		if inputs.IsMouseInput(input) {
			data, n, err := inputs.ReadMouse(input)
			if errors.Is(err, io.EOF) {
				return input, nil // Incomplete.
			}
			if err != nil {
				e.Logf("dropping bad mouse input %v: %s", input, err)
				return nil, nil
			}
			e.Send(CommandFunc(func(e *Editor) { e.handleMouse(data) }))
			input = input[n:]
			continue
		}
		pasted, detected, err := inputs.ConsumeClipboardPaste(input, in)
		if err != nil {
			return nil, err
		}
		if detected {
			if pasted != "" {
				e.SendBufferCmd(PasteText(pasted))
			}
			return nil, nil // The rest of the input is consumed by the paste.
		}
		return input, nil // Incomplete.
	}
	return nil, nil
}

func (e *Editor) handleInputError(err error) {
	e.Logf("input error, quitting: %s", err)
	e.cmdChannel <- commandQuit
}

// handleKeys handles the keys read at once. The keys after the one quitting the buffer are ignored.
func (e *Editor) handleKeys(keys []inputs.Key) {
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

func (e *Editor) handleMouse(data inputs.Mouse) {
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
		if buf.CheckContentCoordinates(event.Y, event.X) {
			e.execBufferCmd(StartTextSelection)
		}
		e.renderRequested = true

	case mouseEventTypeDragEnd:
		e.execBufferCmd(StopTextSelection)

	case mouseEventTypeDoubleClick:
		e.execBufferCmd(SelectWord)

	case mouseEventTypeTripleClick:
		e.execBufferCmd(SelectLine)

	case mouseEventTypeRaw:
		overContent := buf.CheckContentCoordinates(event.Y, event.X)
		if e.ensureMouseTextShape(overContent) {
			e.renderRequested = true
		}
		if event.Button != inputs.MouseButtonLeft || !event.Pressed {
			return
		}

		buf.c = buf.screenToContentPosition(data.Y-1, data.X-1, e.rPrefs.TabSize)
		if buf.selecting && event.Mod.HasMotion() && overContent {
			buf.sel[len(buf.sel)-1].End = buf.c
		} else if !buf.selecting {
			buf.sel = nil
		}
		e.renderRequested = true
	}

}

func (e *Editor) handleKey(k inputs.Key) (quit bool) {
	buf := e.Top()
	// Any key may edit the buffer: the extensions are notified about it once.
	defer e.handleAfterEdit(buf)

	// The keys working in any mode.
	switch k {
	// Save the changes when the terminal loses focus, so the file system is in sync
	// for the tools used outside the editor.
	case inputs.Of(inputs.FocusOut):
		e.saveTop()
		return false
	case inputs.Of(inputs.FocusIn):
		return false

	case inputs.Ctrl('s'):
		e.execBufferCmd(&Save{buf.Path}) // Formatting may change it.
		return false

	// Re-run the last engaged line action.
	case inputs.Ctrl('r'):
		if e.lastAction != nil {
			e.lastAction.Engage()
		}
		return false

	// Show the file picker. The picker itself selects the next match with it.
	case inputs.Ctrl('o'):
		if !prompting[*quickOpen](e) {
			e.startQuickOpen(buf)
			return false
		}

	// Start the search. The search itself moves to the next match with it.
	case inputs.Ctrl('f'):
		if !prompting[*searchPrompt](e) {
			e.startSearch(buf)
			return false
		}

	// Show the changes of the current file.
	case inputs.Ctrl('d'):
		e.showDiff(buf)
		return false

	// Clipboard.
	case inputs.Ctrl('c'):
		e.execBufferCmd(ClipboardCopy)
		return false
	case inputs.Ctrl('v'):
		e.execBufferCmd(ClipboardPaste)
		return false
	case inputs.Ctrl('x'):
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
		case inputs.Rune(':'):
			e.openCmdLine(buf, "", newExPrompt)
			return false
		case inputs.Rune('/'):
			e.startSearch(buf)
			return false
		}
		quit = normalInput(buf, k, &e.rPrefs)
		if buf.engaged != nil {
			// TODO: Consider different ownership.
			e.lastAction, buf.engaged = buf.engaged, nil
		}
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

func (e *Editor) extHandleInsert(buf *Buffer, k inputs.Key) (handled bool) {
	for _, ext := range e.x {
		handled = ext.HandleInsertInput(buf, &e.rPrefs, k)
		if handled {
			return
		}
	}
	return
}
