package editor

import (
	"bufio"
	"fmt"
	"strings"
	"unicode/utf8"

	"rmazur.io/x/edit/internal/content"
	"rmazur.io/x/edit/internal/editor/escape"
)

type Buffer struct {
	path    string
	content content.Interface
	mode    Mode
	dirty   bool
	cmdline string // text typed after ':' while in ModeCommand

	cx, cy int // cursor position: cy is the line index, cx is the byte offset within that line
	offset int // first visible row (scroll)
	w, h   int // terminal window dimensions
}

func NewScratchBuffer() *Buffer {
	return &Buffer{
		content: content.Empty(),
	}
}

func (b *Buffer) clampCursor(_ *RenderPrefs) {
	lines := b.content.Lines()

	b.cy = max(0, min(b.cy, len(lines)-1))

	var line string
	if len(lines) > 0 {
		line = lines[b.cy].String()
	}
	maxX := len(line)
	if b.mode == ModeNormal && maxX > 0 {
		_, sz := utf8.DecodeLastRuneInString(line)
		maxX -= sz // Normal mode: cursor sits on a character, not past the last one.
	}
	b.cx = max(0, min(b.cx, maxX))
	// Vertical movement may land cx mid-rune; snap back to the rune start.
	for b.cx > 0 && b.cx < len(line) && !utf8.RuneStart(line[b.cx]) {
		b.cx--
	}

	// Adjust scroll so cursor is visible.
	if b.cy < b.offset {
		b.offset = b.cy
	}
	if b.cy >= b.offset+b.viewHeight() {
		b.offset = b.cy - b.viewHeight() + 1
	}
}

// byteToScreenCol returns the visual column at the given byte index, expanding
// tabs by tabSize. Used once per render to project cx onto the terminal.
func byteToScreenCol(line string, byteIdx, tabSize int) int {
	if byteIdx > len(line) {
		byteIdx = len(line)
	}
	col := 0
	for i, r := range line {
		if i >= byteIdx {
			break
		}
		if r == '\t' {
			col += tabSize
		} else {
			col++
		}
	}
	return col
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

	var cursorLine string
	if len(lines) > 0 {
		cursorLine = lines[b.cy].String()
	}
	screenCol := byteToScreenCol(cursorLine, b.cx, prefs.TabSize)

	// Status bar (reverse colors).
	restoreColors := escape.ReverseVideo(out)

	if b.mode == ModeCommand {
		// In command mode, the status bar becomes the command line.
		escape.ClearLine(out)
		out.WriteString(":")
		out.WriteString(b.cmdline)
		escape.SetCursorPosition(out, b.h, len(b.cmdline)+2)
		showCursor()
		restoreColors()
		_ = out.Flush()
		return
	}

	modeLabel := b.mode.String()
	dirtyMark := ""
	if b.dirty {
		dirtyMark = " [*]"
	}
	status := fmt.Sprintf(" %s  %s%s", modeLabel, b.path, dirtyMark)
	pos := fmt.Sprintf("%d:%d ", b.cy+1, screenCol+1)
	padding := max(b.w-len(status)-len(pos), 0)
	out.WriteString(status)
	out.WriteString(strings.Repeat(" ", padding))
	out.WriteString(pos)
	restoreColors()

	// Reposition and show cursor.
	screenRow := b.cy - b.offset + 1
	numDisplayWidth := nlDigitsLen(b.content.Len())
	escape.SetCursorPosition(out, screenRow, screenCol+numDisplayWidth+1)
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
