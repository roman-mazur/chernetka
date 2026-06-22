package editor

import (
	"unicode/utf8"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/inputs"
)

func insertInput(buf *Buffer, b []byte, prefs *RenderPrefs) {
	var arrow inputs.CursorArrow
	if inputs.IsArrow(b, &arrow) {
		buf.handleCursor(arrow, prefs)
		return
	}

	if len(b) != 1 {
		return
	}

	// Esc.
	if b[0] == 0x1b {
		buf.mode = ModeNormal
		if buf.cx > 0 {
			buf.cx-- // Land on the last typed character.
		}
		return
	}

	if !buf.canEdit() {
		return
	}
	lines := buf.Content.Lines()
	line := lines[buf.cy].String()
	mut := buf.mutate()

	switch ch := b[0]; ch {
	// Backspace.
	case 0x7f, 0x08:
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
	case '\r':
		mut.Update(buf.cy, content.TextLine(line[:buf.cx]))
		mut.Insert(buf.cy+1, content.TextLine(line[buf.cx:]))
		buf.cy++
		buf.cx = 0

	// Printable ASCII.
	default:
		if ch >= 0x20 {
			mut.Update(buf.cy, content.TextLine(line[:buf.cx]+string(ch)+line[buf.cx:]))
			buf.cx++
		}
	}

	return
}
