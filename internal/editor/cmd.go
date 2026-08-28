package editor

import (
	"os"
	"unicode/utf8"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/clipb"
	"rmazur.io/chernetka/internal/editor/inputs"
)

// BufferCommand performs some action on a Buffer.
type BufferCommand interface {
	DoOnBuffer(buf *Buffer, prefs RenderPrefs)
}

// Command performs an action on an Editor.
type Command interface {
	DoOnEditor(e *Editor)
}

// RelMove changes the cursor by Dx and Dy.
type RelMove struct {
	Dx, Dy int
}

func (r RelMove) DoOnBuffer(buf *Buffer, _ RenderPrefs) {
	if r.Dx != 0 {
		r.moveDx(buf)
	}
	if r.Dy != 0 {
		buf.c.Line += r.Dy
	}
	if buf.selecting {
		buf.sel[len(buf.sel)-1].End = buf.c
	}
}

func (r RelMove) moveDx(b *Buffer) {
	line := b.Content.Lines()[b.c.Line].String()
	d := r.Dx
	for d > 0 && b.c.Col < len(line) {
		_, sz := utf8.DecodeRuneInString(line[b.c.Col:])
		b.c.Col += sz
		d--
	}
	for d < 0 && b.c.Col > 0 {
		_, sz := utf8.DecodeLastRuneInString(line[:b.c.Col])
		b.c.Col -= sz
		d++
	}
}

// ScreenMove adjusts Buffer offset and cursor position to scroll by the screen height.
// ScreenD configured the number screen to scroll (buffer offset will be capped by the content length).
// Offset calculation also does not change the buffer if the content end is already visible.
// It also tries to keep the bottom or top line visible when scrolling by one screen.
type ScreenMove struct {
	ScreenD int
}

func (sm ScreenMove) DoOnBuffer(buf *Buffer, _ RenderPrefs) {
	screenH := buf.viewHeight()
	if screenH <= 0 {
		return
	}
	contentLen := buf.Content.Len()

	keepVisibleLineOffset := -1
	if sm.ScreenD < 0 {
		keepVisibleLineOffset = 1
	}
	dy := screenH*sm.ScreenD + keepVisibleLineOffset

	if sm.ScreenD == 1 && buf.offset+dy >= contentLen {
		return
	}

	buf.offset = max(0, min(buf.offset+dy, contentLen-1))
	buf.c.Line = max(0, min(buf.c.Line+dy, contentLen-1))
	clampBufferCx(buf)
}

type Scroll inputs.ScrollDirection

func (s Scroll) DoOnBuffer(buf *Buffer, _ RenderPrefs) {
	switch inputs.ScrollDirection(s) {
	case inputs.ScrollDirectionUp:
		buf.offset--
	case inputs.ScrollDirectionDown:
		buf.offset++
	}
	buf.offset = max(0, min(buf.offset, buf.Content.Len()-buf.viewHeight()-1))
}

func clampBufferCx(buf *Buffer) {
	if buf.Content.Len() > 0 {
		buf.c.Col = min(buf.c.Col, buf.Content.Lines()[buf.c.Line].Len()-1)
	}
}

type BufferCommandFunc func(b *Buffer, prefs RenderPrefs)

func (f BufferCommandFunc) DoOnBuffer(buf *Buffer, prefs RenderPrefs) { f(buf, prefs) }

var (
	// MoveHome moves the cursor to the beginning of the line.
	MoveHome = BufferCommandFunc(func(b *Buffer, _ RenderPrefs) { b.c.Col = 0 })
	// MoveEnd moves the cursor to the end of the line.
	MoveEnd = BufferCommandFunc(func(b *Buffer, _ RenderPrefs) { b.c.Col = b.Content.Lines()[b.c.Line].Len() })
	// MoveContentStart moves the cursor to the first line.
	MoveContentStart = BufferCommandFunc(func(b *Buffer, _ RenderPrefs) {
		b.c.Line = 0
		b.offset = 0
		clampBufferCx(b)
	})
	// MoveContentEnd moves the cursor to the last line.
	MoveContentEnd = BufferCommandFunc(func(b *Buffer, _ RenderPrefs) {
		linesCnt := b.Content.Len()
		if linesCnt == 0 {
			return
		}
		b.c.Line = linesCnt - 1
		b.offset = max(0, b.c.Line-b.viewHeight())
		clampBufferCx(b)
	})
	// StartTextSelection begins selecting text at the current cursor position.
	StartTextSelection = BufferCommandFunc(func(b *Buffer, _ RenderPrefs) {
		if b.selecting {
			return
		}
		b.selecting = true
		b.sel = []content.Span{{b.c, b.c}}
	})
	// StopTextSelection finishes selecting text at the current cursor position.
	StopTextSelection = BufferCommandFunc(func(b *Buffer, _ RenderPrefs) {
		if !b.selecting {
			return
		}
		b.selecting = false
		b.sel[len(b.sel)-1].End = b.c
	})
	// ClipboardCopy copies selected text to the clipboard.
	ClipboardCopy = BufferCommandFunc(func(b *Buffer, _ RenderPrefs) {
		clipboard.Write(b.SelectedText())
	})
	// ClipboardCut copies selected text to the clipboard and removes it from the buffer.
	ClipboardCut = BufferCommandFunc(func(b *Buffer, _ RenderPrefs) {
		if len(b.sel) == 0 || !b.canEdit() {
			return
		}
		clipboard.Write(b.SelectedText())
		b.c = content.DeleteSpan(b.Mutate(), b.sel[len(b.sel)-1])
		b.cancelSelection()
	})
	// ClipboardPaste inserts the clipboard text at the cursor position.
	ClipboardPaste = BufferCommandFunc(func(b *Buffer, _ RenderPrefs) {
		if !b.canEdit() {
			return
		}
		b.c = content.InsertText(b.Mutate(), b.c, clipboard.Read())
	})
)

var clipboard clipb.Clipboard

// Save stores the boffer content in the destination path.
type Save struct {
	DstPath string
}

func (s *Save) DoOnBuffer(buf *Buffer, _ RenderPrefs) {
	if s.DstPath == "" {
		return
	}
	err := content.Save(buf.Content, s.DstPath)
	if err == nil {
		buf.dirty = false
	}
	// TODO: Visualize the error.
}

// OpenFile opens a new file via Editor.OpenReader.
type OpenFile struct {
	Path string
}

func (of *OpenFile) DoOnEditor(e *Editor) {
	e.renderRequested = true

	f, err := os.Open(of.Path)
	if err != nil {
		of.handleError(e, err)
		return
	}
	defer func() { _ = f.Close() }()

	if err := e.OpenReader(of.Path, f); err != nil {
		of.handleError(e, err)
	}
}

func (of *OpenFile) handleError(e *Editor, err error) {
	e.push(&Buffer{
		Path:    of.Path,
		Content: &content.ErrorContent{Error: err},
	})
}

type CommandFunc func(e *Editor)

func (f CommandFunc) DoOnEditor(e *Editor) { f(e) }

var (
	commandQuit          = CommandFunc(func(e *Editor) { e.quitRequested = true })
	commandRequestLayout = CommandFunc(func(e *Editor) { e.renderRequested = true })
)

type SwitchMode Mode

func (s SwitchMode) DoOnBuffer(buf *Buffer, _ RenderPrefs) {
	if Mode(s) == ModeInsert && !buf.canEdit() {
		return
	}
	buf.mode = Mode(s)
}
