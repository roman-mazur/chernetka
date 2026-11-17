package editor

import (
	"slices"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/input"
)

// The tool buffer is shown at the bottom of the screen for the output of the tools,
// like the usages of a symbol found by a language server (see ShowTool). It's a
// read-only list of lines that can be engaged: Enter on a line jumps to what it
// lists. Ctrl+T switches the keyboard focus between the main buffer and the tool
// buffer, and so does a click. With the focus in the tool buffer, the cursor keys
// and the movement keys of the review mode move around it, Esc returns the focus
// to the main buffer, and q closes the tool buffer.

// ShowTool shows the document in the tool buffer, replacing what was shown there,
// and gives it the focus.
type ShowTool struct {
	Content content.Document
}

func (st *ShowTool) DoOnEditor(e *Editor) { e.showTool(st.Content) }

// CloseTool closes the tool buffer.
var CloseTool = CommandFunc(func(e *Editor) { e.closeTool() })

func (e *Editor) showTool(doc content.Document) {
	e.closeTool()
	buf := &Buffer{
		Content:         doc,
		hideLineNumbers: true,
		hideLineActions: true, // Most of the lines can be engaged: marking them is noise.
	}
	e.prepareExt(buf)
	e.toolBuf = buf
	e.toolFocused = true
	e.renderRequested = true
}

func (e *Editor) closeTool() {
	if e.toolBuf == nil {
		return
	}
	if e.lastAction.buf == e.toolBuf {
		e.lastAction = bufferAction{}
	}
	_ = e.toolBuf.Close()
	e.toolBuf = nil
	e.toolFocused = false
	e.renderRequested = true
}

// focused returns the buffer receiving the keyboard input: the tool buffer if it has
// the focus, the top buffer otherwise.
func (e *Editor) focused() *Buffer {
	if e.toolFocused && e.toolBuf != nil {
		return e.toolBuf
	}
	return e.Top()
}

// focusTool gives the keyboard focus to the tool buffer if there is one, or to the main buffer.
func (e *Editor) focusTool(focus bool) {
	e.toolFocused = focus && e.toolBuf != nil
	e.renderRequested = true
}

// toolInput handles the key with the tool buffer focused.
func (e *Editor) toolInput(k input.Key) {
	buf := e.toolBuf
	switch k.Special {
	case input.CursorMove:
		buf.handleCursor(k.Cursor, k.Mod, &e.rPrefs)
	case input.Esc:
		if buf.HasSelection() {
			buf.cancelSelection()
			return
		}
		e.focusTool(false)
	case input.Enter:
		if action := buf.lineAction(buf.c().Line); action != nil {
			e.engage(buf, action)
		} else {
			RelMove{Dy: 1}.DoOnBuffer(buf, e.rPrefs)
		}
	case input.Text:
		if k.Mod != 0 {
			return
		}
		switch k.Rune {
		case 'q':
			e.closeTool()
		case 'h', 'j', 'k', 'l', '0', '$', 'g', 'G', ' ', 'b', 'y', 'n', 'N':
			// Moving around and copying only: the tool buffer is not edited.
			reviewCommand(buf, k.Rune, &e.rPrefs)
		}
	}
}

// toolAt returns the tool buffer if the screen row (0-based) is in its area.
func (e *Editor) toolAt(row int) *Buffer {
	if tb := e.toolBuf; tb != nil && row >= tb.y {
		return tb
	}
	return nil
}

// FindBuffer returns the open buffer of the file at path, or nil.
func (e *Editor) FindBuffer(path string) *Buffer {
	i := slices.IndexFunc(e.bufs, func(b *Buffer) bool { return b.Path != "" && samePath(b.Path, path) })
	if i < 0 {
		return nil
	}
	return e.bufs[i]
}

// EditBuffer changes an open buffer with edit, which may not be the active one,
// and notifies the extensions about the changes like about the user edits.
// The changes of a buffer that is not active are saved to its file right away,
// like they are when the user switches to another buffer. It must be called on
// the editor loop.
func (e *Editor) EditBuffer(buf *Buffer, edit func(buf *Buffer)) {
	edit(buf)
	e.renderRequested = true
	e.handleAfterEdit(buf)
	if buf != e.Top() && buf.dirty && buf.Path != "" {
		e.SaveBuffer(buf)
	}
}

// SaveBuffer saves the buffer to its file, formatting it first like Ctrl+S does.
// It must be called on the editor loop.
func (e *Editor) SaveBuffer(buf *Buffer) {
	if buf.Path == "" {
		return
	}
	if _, isDir := buf.Content.(*content.FsContent); isDir {
		return // Its path is only a display name.
	}
	(&Save{buf.Path}).DoOnBuffer(buf, e.rPrefs)
	e.renderRequested = true
	e.handleAfterEdit(buf) // Formatting may change it.
}

// CloseBuffer closes the buffer, which may not be the active one, unless it's the
// only one: closing the last buffer ends the editor. It reports whether the buffer
// is closed. It must be called on the editor loop.
func (e *Editor) CloseBuffer(buf *Buffer) bool {
	i := slices.Index(e.bufs, buf)
	if i < 0 || len(e.bufs) == 1 {
		return false
	}
	if buf == e.Top() {
		e.pop()
		return true
	}
	e.cancelCmdLine()
	if e.lastAction.buf == buf {
		e.lastAction = bufferAction{}
	}
	_ = buf.Close()
	e.bufs = slices.Delete(e.bufs, i, i+1)
	e.renderRequested = true
	return true
}
