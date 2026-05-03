package editor

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"rmazur.io/x/edit/internal/content"
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

func NewFullTextBuffer(path string, in io.Reader) (*Buffer, error) {
	data, err := content.LoadFullText(in)
	if err != nil {
		return nil, err
	}

	return &Buffer{
		path:    path,
		content: data,
	}, nil
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
		maxX := len(lines[b.cy].String())
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

func (b *Buffer) render(out *bufio.Writer) {
	out.WriteString("\x1b[?25l") // hide cursor while drawing
	out.WriteString("\x1b[H")    // move to top-left

	lines := b.content.Lines()

	for row := 0; row < b.viewHeight(); row++ {
		contentRow := row + b.offset
		out.WriteString("\x1b[K") // clear line
		if contentRow < len(lines) {
			line := lines[contentRow].String()
			if len(line) > b.w {
				line = line[:b.w]
			}
			out.WriteString(line)
		} else {
			out.WriteString("~")
		}
		out.WriteString("\r\n")
	}

	// status bar (reverse video)
	out.WriteString("\x1b[7m")
	modeLabel := b.mode.String()
	dirtyMark := ""
	if b.dirty {
		dirtyMark = " [*]"
	}
	status := fmt.Sprintf(" %s  %s%s", modeLabel, b.path, dirtyMark)
	pos := fmt.Sprintf("%d:%d ", b.cy+1, b.cx+1)
	padding := b.w - len(status) - len(pos)
	if padding < 0 {
		padding = 0
	}
	out.WriteString(status)
	out.WriteString(strings.Repeat(" ", padding))
	out.WriteString(pos)
	out.WriteString("\x1b[0m")

	// reposition and show cursor
	screenRow := b.cy - b.offset + 1
	screenCol := b.cx + 1
	fmt.Fprintf(out, "\x1b[%d;%dH", screenRow, screenCol)
	out.WriteString("\x1b[?25h")
	out.Flush()
}

func (b *Buffer) canEdit() bool {
	_, ok := b.content.(content.Mutable)
	return ok
}

func (b *Buffer) handleCursor(input []byte) {
	switch input[2] {
	case 'A':
		b.cy--
	case 'B':
		b.cy++
	case 'C':
		b.cx++
	case 'D':
		b.cx--
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
	if !bm.canEdit() {
		return
	}
	bm.Mutable.Insert(pos, line)
	bm.dirty = true
}
func (bm bufMutator) Update(pos int, line content.Line) {
	if !bm.canEdit() {
		return
	}
	bm.Mutable.Update(pos, line)
	bm.dirty = true
}
func (bm bufMutator) Delete(pos int) {
	if !bm.canEdit() {
		return
	}
	bm.Mutable.Delete(pos)
	bm.dirty = true
}
