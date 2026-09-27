package editor

import (
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/clipb"
	"rmazur.io/chernetka/internal/editor/input"
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
	buf.clampPos()
	buf.updateSelection()
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
	clampBufferCx(b)
}

// ScreenMove adjusts Buffer offset and cursor position to scroll by the screen height.
// ScreenD configured the number screen to scroll (buffer offset will be capped by the content length).
// Offset calculation also does not change the buffer if the content end is already visible.
// It also tries to keep the bottom or top line visible when scrolling by one screen.
type ScreenMove struct {
	ScreenD int
}

func (sm ScreenMove) DoOnBuffer(buf *Buffer, _ RenderPrefs) {
	screenH := buf.h
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
	buf.updateSelection()
}

type Scroll input.ScrollDirection

func (s Scroll) DoOnBuffer(buf *Buffer, prefs RenderPrefs) {
	switch input.ScrollDirection(s) {
	case input.ScrollDirectionUp:
		buf.offset--
	case input.ScrollDirectionDown:
		buf.offset++
	case input.ScrollDirectionLeft:
		buf.xoff = max(0, buf.xoff-scrollStepX)
	case input.ScrollDirectionRight:
		// Stop when the end of the widest visible line is shown,
		// but don't jump back if the cursor has scrolled further.
		buf.xoff = min(buf.xoff+scrollStepX, max(buf.xoff, buf.maxScrollX(prefs.TabSize)))
	}
	buf.offset = max(0, min(buf.offset, buf.Content.Len()-buf.h-1))
}

// scrollStepX is how many columns the text is scrolled horizontally by one wheel event.
const scrollStepX = 2

// maxScrollX returns the horizontal scroll that shows the end of the widest visible line.
func (b *Buffer) maxScrollX(tabSize int) int {
	lines := b.Content.Lines()
	widest := 0
	for ln := b.offset; ln < b.offset+b.printableLinesCount(); ln++ {
		line := lines[ln].String()
		widest = max(widest, runeToScreenCol(line, len(line), tabSize))
	}
	return max(0, widest-b.textWidth()+1)
}

func clampBufferCx(buf *Buffer) {
	curLen := buf.Content.Len()
	if curLen > 0 && buf.c.Line < curLen {
		buf.c.Col = max(0, min(buf.c.Col, buf.Content.Lines()[buf.c.Line].Len()))
	} else {
		buf.c.Col = 0
	}
}

type BufferCommandFunc func(b *Buffer, prefs RenderPrefs)

func (f BufferCommandFunc) DoOnBuffer(buf *Buffer, prefs RenderPrefs) { f(buf, prefs) }

var (
	// MoveHome moves the cursor to the first non-whitespace symbol of the line.
	// If the cursor is already within the leading whitespace (or right after it),
	// it moves to the beginning of the line.
	MoveHome = BufferCommandFunc(func(b *Buffer, _ RenderPrefs) {
		if b.Content.Len() == 0 {
			return
		}
		line := b.Content.Lines()[b.c.Line].String()
		indent := len(line) - len(strings.TrimLeftFunc(line, unicode.IsSpace))
		if b.c.Col <= indent {
			b.c.Col = 0
		} else {
			b.c.Col = indent
		}
		b.updateSelection()
	})
	// MoveEnd moves the cursor to the end of the line.
	MoveEnd = BufferCommandFunc(func(b *Buffer, _ RenderPrefs) {
		b.c.Col = b.Content.Lines()[b.c.Line].Len()
		b.updateSelection()
	})
	// MoveContentStart moves the cursor to the first line.
	MoveContentStart = BufferCommandFunc(func(b *Buffer, _ RenderPrefs) {
		b.c.Line = 0
		b.offset = 0
		clampBufferCx(b)
		b.updateSelection()
	})
	// MoveContentEnd moves the cursor to the last line.
	MoveContentEnd = BufferCommandFunc(func(b *Buffer, _ RenderPrefs) {
		linesCnt := b.Content.Len()
		if linesCnt == 0 {
			return
		}
		b.c.Line = linesCnt - 1
		b.offset = max(0, b.c.Line-b.h)
		clampBufferCx(b)
		b.updateSelection()
	})
	// StartTextSelection begins selecting text at the current cursor position.
	// A selection ending at the cursor (e.g., a selected word) is extended instead.
	StartTextSelection = BufferCommandFunc(func(b *Buffer, _ RenderPrefs) {
		if b.selecting {
			return
		}
		b.clampPos()
		b.selecting = true
		if n := len(b.sel); n > 0 && b.sel[n-1].End == b.c {
			return
		}
		b.sel = []content.Span{{b.c, b.c}}
	})
	// StopTextSelection finishes selecting text at the current cursor position.
	StopTextSelection = BufferCommandFunc(func(b *Buffer, _ RenderPrefs) {
		if !b.selecting {
			return
		}
		b.updateSelection()
		b.selecting = false
	})
	// SelectWord adjusts the buffer selection to select the word at the current cursor.
	SelectWord = BufferCommandFunc(func(b *Buffer, _ RenderPrefs) {
		b.cancelSelection()
		if b.Content.Len() == 0 {
			return
		}
		line := b.Content.Lines()[b.c.Line].String()

		start := b.c.Col
		for start > 0 {
			r, size := utf8.DecodeLastRuneInString(line[:start])
			if unicode.IsLetter(r) {
				start -= size
			} else {
				break
			}
		}

		end := b.c.Col
		for end < len(line) {
			r, size := utf8.DecodeRuneInString(line[end:])
			if unicode.IsLetter(r) {
				end += size
			} else {
				break
			}
		}

		if end > start {
			b.sel = append(b.sel, content.Span{
				Start: content.Position{Line: b.c.Line, Col: start},
				End:   content.Position{Line: b.c.Line, Col: end},
			})
			b.c.Col = end
		}
	})
	// SelectLine adjusts the buffer selection to select the whole line at the current cursor,
	// including its line break, and moves the cursor to the selection end.
	SelectLine = BufferCommandFunc(func(b *Buffer, _ RenderPrefs) {
		b.cancelSelection()
		if b.Content.Len() == 0 {
			return
		}
		b.clampPos()
		end := content.Position{Line: b.c.Line + 1}
		if end.Line == b.Content.Len() {
			end = content.Position{Line: b.c.Line, Col: b.Content.Lines()[b.c.Line].Len()}
		}
		b.sel = []content.Span{{Start: content.Position{Line: b.c.Line}, End: end}}
		b.c = end
	})
	// DeleteSelection command deletes currently selected content in the buffer adjusting the cursor position.
	DeleteSelection = BufferCommandFunc(func(b *Buffer, _ RenderPrefs) {
		m := b.Mutate()
		for _, span := range slices.Backward(b.sel) {
			b.c = content.DeleteSpan(m, span)
		}
		b.cancelSelection()
	})
	// ClipboardCopy copies selected text to the clipboard.
	ClipboardCopy = BufferCommandFunc(func(b *Buffer, _ RenderPrefs) {
		clipboard.Write(b.SelectedText())
	})
	// ClipboardCut copies selected text to the clipboard and removes it from the buffer.
	ClipboardCut = BufferCommandFunc(func(b *Buffer, prefs RenderPrefs) {
		if !b.hasSelection() || !b.canEdit() {
			return
		}
		clipboard.Write(b.SelectedText())
		DeleteSelection.DoOnBuffer(b, prefs)
	})
	// ClipboardPaste inserts the clipboard text at the cursor position.
	ClipboardPaste = BufferCommandFunc(func(b *Buffer, prefs RenderPrefs) {
		PasteText(clipboard.Read()).DoOnBuffer(b, prefs)
	})
)

var clipboard clipb.Clipboard

type PasteText string

func (pt PasteText) DoOnBuffer(b *Buffer, prefs RenderPrefs) {
	if !b.canEdit() {
		return
	}
	if b.hasSelection() {
		// The pasted text replaces the selection.
		DeleteSelection.DoOnBuffer(b, prefs)
	}
	b.cancelSelection()
	b.c = content.InsertText(b.Mutate(), b.c, string(pt))
}

// Save stores the buffer content in the destination path, formatting it first
// with the extensions implementing Formatter.
type Save struct {
	DstPath string
}

func (s *Save) DoOnBuffer(buf *Buffer, prefs RenderPrefs) {
	if s.DstPath == "" {
		return
	}
	if f, ok := FindExtData[Formatter](buf); ok {
		f.Format(prefs)
	}
	err := content.Save(buf.Content, s.DstPath)
	if err == nil {
		buf.dirty = false
		if samePath(s.DstPath, buf.Path) {
			buf.fileText = buf.Text()
		}
	}
	// TODO: Visualize the error.
}

type CommandFunc func(e *Editor)

func (f CommandFunc) DoOnEditor(e *Editor) { f(e) }

var (
	commandQuit          = CommandFunc(func(e *Editor) { e.quitRequested = true })
	commandRequestRender = CommandFunc(func(e *Editor) { e.renderRequested = true })
)

type SwitchMode Mode

func (s SwitchMode) DoOnBuffer(buf *Buffer, _ RenderPrefs) {
	if Mode(s) == ModeInsert && !buf.canEdit() {
		return
	}
	buf.mode = Mode(s)
}
