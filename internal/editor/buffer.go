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
		content: &data,
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

func (b *Buffer) render(out *bufio.Writer, prefs *renderPrefs) {
	out.WriteString("\x1b[?25l") // hide cursor while drawing
	out.WriteString("\x1b[H")    // move to top-left

	lines := b.content.Lines()

	printableCount := min(b.viewHeight(), len(lines)-b.offset)
	printContent(out, b.content, b.offset, b.offset+printableCount, prefs)

	for row := printableCount; row < b.viewHeight(); row++ {
		out.WriteString("~\r\n")
	}

	// Status bar (reverse colors).
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

	// Reposition and show cursor.
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

func (b *Buffer) handleCursor(input []byte, prefs *renderPrefs) {
	switch input[2] {
	case 'A':
		b.cy--
	case 'B':
		b.cy++
	case 'C':
		b.moveHorizontal(1, prefs.tabSize)
	case 'D':
		b.moveHorizontal(-1, prefs.tabSize)
	}
}

func (b *Buffer) moveHorizontal(d int, tabSize int) {
	line := b.content.Lines()[b.cy].String()
	ti := screenToTextIdx(line, b.cx, tabSize)
	switch {
	case d > 0:
		if ti >= len(line) {
			return
		}
		if line[ti] == '\t' {
			b.cx += tabSize
		} else {
			b.cx++
		}
	case d < 0:
		if ti == 0 {
			return
		}
		// '\t' is single-byte ASCII (0x09); continuation bytes (0x80–0xBF) in
		// UTF-8 can never equal 0x09, so checking line[ti-1] is safe here.
		if line[ti-1] == '\t' {
			b.cx -= tabSize
		} else {
			b.cx--
		}
	}
}

// screenToTextIdx returns the byte index of the character displayed at the
// given visual column. Each tab counts as tabSize columns. Returns len(line)
// when screenCol is at or past the visual end of the line.
func screenToTextIdx(line string, screenCol int, tabSize int) int {
	sc := 0
	for i, c := range line {
		w := 1
		if c == '\t' {
			w = tabSize
		}
		sc += w
		if sc > screenCol {
			return i
		}
	}
	return len(line)
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
