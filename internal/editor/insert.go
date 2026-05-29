package editor

import (
	"unicode/utf8"

	"rmazur.io/x/edit/internal/content"
)

func insertInput(buf *Buffer, b []byte) {
	// Esc.
	if len(b) == 1 && b[0] == 0x1b {
		buf.mode = ModeNormal
		if buf.cx > 0 {
			buf.cx-- // Land on the last typed character.
		}
		return
	}

	if !buf.canEdit() {
		return
	}
	lines := buf.content.Lines()
	line := lines[buf.cy].String()
	mut := buf.mutate()

	switch {
	// Backspace.
	case len(b) == 1 && (b[0] == 0x7f || b[0] == 0x08):
		if buf.cx > 0 {
			_, sz := utf8.DecodeLastRuneInString(line[:buf.cx])
			mut.Update(buf.cy, content.TextLine(line[:buf.cx-sz]+line[buf.cx:]))
			buf.cx -= sz
		} else if buf.cy > 0 {
			prev := lines[buf.cy-1].String()
			buf.cx = len(prev)
			mut.Update(buf.cy-1, content.TextLine(prev+line))
			mut.Delete(buf.cy)
			buf.cy--
		}

	// Enter.
	case len(b) == 1 && b[0] == '\r':
		mut.Update(buf.cy, content.TextLine(line[:buf.cx]))
		mut.Insert(buf.cy+1, content.TextLine(line[buf.cx:]))
		buf.cy++
		buf.cx = 0

	// Printable ASCII.
	case len(b) == 1 && b[0] >= 0x20:
		mut.Update(buf.cy, content.TextLine(line[:buf.cx]+string(b[0])+line[buf.cx:]))
		buf.cx++
	}
}
