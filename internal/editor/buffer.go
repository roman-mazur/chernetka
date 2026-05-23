package editor

import (
	"bufio"
	"fmt"
	"strings"

	"rmazur.io/x/edit/internal/content"
	"rmazur.io/x/edit/internal/editor/escape"
)

type Buffer struct {
	path    string
	content content.Interface
	mode    Mode
	dirty   bool

	cx, cy int // cursor column and row within the file (0-indexed)
	offset int // first visible row (scroll)
	w, h   int // terminal window dimensions
}

func NewScratchBuffer() *Buffer {
	return &Buffer{
		content: content.Empty(),
	}
}

func (b *Buffer) clampCursor() {
	lines := b.content.Lines()

	b.cy = max(0, min(b.cy, len(lines)-1))

	maxX := 0
	if len(lines) > 0 {
		maxX = len(lines[b.cy].String())
		if b.mode == ModeNormal && maxX > 0 {
			maxX-- // Normal mode: cursor sits on a character, not past the last one.
		}
	}
	b.cx = max(0, min(b.cx, maxX))

	// Adjust scroll so cursor is visible.
	if b.cy < b.offset {
		b.offset = b.cy
	}
	if b.cy >= b.offset+b.viewHeight() {
		b.offset = b.cy - b.viewHeight() + 1
	}
}

// viewHeight is the number of text rows (terminal height minus the status bar).
func (b *Buffer) viewHeight() int { return b.h - 1 }

func (b *Buffer) render(out *bufio.Writer, prefs *RenderPrefs) {
	showCursor := escape.HideCursor(out)
	escape.MoveTopLeft(out)

	lines := b.content.Lines()

	printableCount := min(b.viewHeight(), len(lines)-b.offset)
	printFmt(out, b.content, b.offset, b.offset+printableCount, prefs.TabSize)

	for row := printableCount; row < b.viewHeight(); row++ {
		escape.ClearLine(out)
		out.WriteString("\r\n")
	}

	// Status bar (reverse colors).
	restoreColors := escape.ReverseVideo(out)
	modeLabel := b.mode.String()
	dirtyMark := ""
	if b.dirty {
		dirtyMark = " [*]"
	}
	status := fmt.Sprintf(" %s  %s%s", modeLabel, b.path, dirtyMark)
	pos := fmt.Sprintf("%d:%d ", b.cy+1, b.cx+1)
	padding := max(b.w-len(status)-len(pos), 0)
	out.WriteString(status)
	out.WriteString(strings.Repeat(" ", padding))
	out.WriteString(pos)
	restoreColors()

	// Reposition and show cursor.
	screenRow := b.cy - b.offset + 1
	screenCol := b.cx + 1
	escape.SetCursorPosition(out, screenRow, screenCol)
	showCursor()

	_ = out.Flush()
}

func (b *Buffer) canEdit() bool {
	_, ok := b.content.(content.Mutable)
	return ok
}

func (b *Buffer) handleCursor(input []byte, prefs *RenderPrefs) {
	switch input[2] {
	case 'A':
		RelMove{Dy: -1}.DoOnBuffer(b, *prefs)
	case 'B':
		RelMove{Dy: 1}.DoOnBuffer(b, *prefs)
	case 'C':
		RelMove{Dx: 1}.DoOnBuffer(b, *prefs)
	case 'D':
		RelMove{Dx: -1}.DoOnBuffer(b, *prefs)
	}
}

func (b *Buffer) mutate() content.Mutable {
	m, _ := b.content.(content.Mutable)
	return bufMutator{
		Mutable: m,
		Buffer:  b,
	}
}

type bufMutator struct {
	content.Mutable
	*Buffer
}

func (bm bufMutator) Insert(pos int, line content.Line) {
	if bm.Mutable == nil {
		return
	}
	bm.Mutable.Insert(pos, line)
	bm.dirty = true
}
func (bm bufMutator) Update(pos int, line content.Line) {
	if bm.Mutable == nil {
		return
	}
	bm.Mutable.Update(pos, line)
	bm.dirty = true
}
func (bm bufMutator) Delete(pos int) {
	if bm.Mutable == nil {
		return
	}
	bm.Mutable.Delete(pos)
	bm.dirty = true
}
