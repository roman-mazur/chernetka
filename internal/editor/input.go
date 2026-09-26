package editor

import (
	"bufio"
	"context"
	"errors"
	"path/filepath"
	"sync"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/inputs"
)

func (e *Editor) readAndHandleInput(ctx context.Context, in *bufio.Reader) {
	const bufSize = 64
	bPool := sync.Pool{New: func() any { return make([]byte, bufSize) }}
	var inBuf [bufSize]byte

	for ctx.Err() == nil {
		// Get input.
		n, err := in.Read(inBuf[:])
		if err != nil {
			e.handleInputError(err)
			break
		}
		if n == 0 {
			continue
		}
		buf := bPool.Get().([]byte)
		copy(buf, inBuf[:n])
		input := buf[:n]
		if debugInput {
			e.Logf("input: %v", input)
		}

		// Check for mouse input.
		_, k, err := e.readAndHandleMouse(input)
		if err != nil {
			e.handleInputError(err)
			break
		}
		input = input[k:]
		if len(input) == 0 {
			bPool.Put(buf)
			continue
		}

		// Check for clipboard paste.
		clipboardContent, detected, err := inputs.ConsumeClipboardPaste(input, in)
		if err != nil {
			e.handleInputError(err)
			break
		}
		if clipboardContent != "" {
			e.SendBufferCmd(PasteText(clipboardContent))
		}
		if detected {
			bPool.Put(buf)
			continue
		}

		// Handle input.
		e.Send(CommandFunc(func(e *Editor) {
			quit := e.handleInput(input)
			if quit {
				commandQuit(e)
			} else {
				e.renderRequested = true
			}
			bPool.Put(buf)
		}))
	}
}

func (e *Editor) handleInputError(err error) {
	e.Logf("input error, quitting: %s", err)
	e.cmdChannel <- commandQuit
}

func (e *Editor) readAndHandleMouse(input []byte) (handled bool, n int, err error) {
	var data inputs.Mouse
	data, n, err = inputs.ReadMouse(input)

	if err != nil {
		if errors.Is(err, inputs.ErrorNotMouse) {
			err = nil
		}
		return
	}

	handled = true
	e.cmdChannel <- CommandFunc(func(e *Editor) { e.handleMouse(data) })
	return
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

func (e *Editor) handleInput(input []byte) (quit bool) {
	buf := e.Top()
	// Any input may edit the buffer: the extensions are notified about it once.
	defer e.handleAfterEdit(buf)

	// Save the changes when the terminal loses focus, so the file system is in sync
	// for the tools used outside the editor.
	if inputs.IsFocusOut(input) {
		e.saveTop()
		return false
	}
	if inputs.IsFocusIn(input) {
		return false
	}

	// Ctrl+S saves the current buffer in any mode.
	if inputs.IsSaveCommand(input) {
		e.execBufferCmd(&Save{buf.Path}) // Formatting may change it.
		return false
	}

	// Ctrl+R re-runs the last engaged line action in any mode.
	if inputs.IsRerunCommand(input) {
		if e.lastAction != nil {
			e.lastAction.Engage()
		}
		return false
	}

	// Ctrl+O shows the file picker in any mode.
	if inputs.IsQuickOpenCommand(input) && e.status.quick == nil {
		e.startQuickOpen(buf)
		return false
	}

	// Ctrl+F starts the search in any mode, or moves to the next match while the pattern is typed.
	if inputs.IsFindCommand(input) {
		if buf.search.typing {
			buf.searchMove(1)
		} else {
			e.closeQuickOpen()
			buf.startSearch()
		}
		return false
	}

	// Ctrl+D shows the changes of the current file in any mode.
	if inputs.IsShowDiffCommand(input) {
		e.showDiff(buf)
		return false
	}

	var clipboardOp inputs.ClipboardOp
	if inputs.IsClipboardOp(input, &clipboardOp) {
		if cmd := ClipboardCommand(clipboardOp); cmd != nil {
			e.execBufferCmd(cmd)
		}
		return false
	}

	switch buf.mode {
	case ModeNormal:
		quit = normalInput(buf, input, &e.rPrefs)
		if buf.engaged != nil {
			// TODO: Consider different ownership.
			e.lastAction, buf.engaged = buf.engaged, nil
		}
		return

	case ModeInsert:
		handled := e.extHandleInsert(buf, input)
		if !handled {
			insertInput(buf, input, &e.rPrefs)
		}
		return false

	case ModeCommand:
		if e.status.quick != nil && e.quickOpenInput(input) {
			return false
		}
		if buf.search.typing && searchInput(buf, input) {
			return false
		}
		quit = commandInput(buf, input, &e.rPrefs)
		e.syncQuickOpen(buf)
		buf.syncSearch()
		return quit
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

func (e *Editor) extHandleInsert(buf *Buffer, b []byte) (handled bool) {
	for _, ext := range e.x {
		handled = ext.HandleInsertInput(buf, &e.rPrefs, b)
		if handled {
			return
		}
	}
	return
}
