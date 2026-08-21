package editor

import (
	"unicode/utf8"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/inputs"
)

func insertInput(buf *Buffer, b []byte, prefs *RenderPrefs) {
	var (
		arrow inputs.CursorArrow
		mod   inputs.Modifier
	)
	if inputs.IsArrow(b, &arrow, &mod) {
		buf.handleCursor(arrow, mod, prefs)
		return
	}

	// Esc.
	if inputs.IsEscape(b) {
		buf.mode = ModeNormal
		if buf.c.x > 0 {
			buf.c.x-- // Land on the last typed character.
		}
		return
	}

	if len(b) != 1 {
		return
	}

	if !buf.canEdit() {
		return
	}
	lines := buf.Content.Lines()
	line := lines[buf.c.y].String()
	mut := buf.Mutate()

	switch ch := b[0]; ch {
	// Backspace.
	case 0x7f, 0x08:
		if buf.c.x > 0 {
			_, sz := utf8.DecodeLastRuneInString(line[:buf.c.x])
			mut.Update(buf.c.y, content.TextLine(line[:buf.c.x-sz]+line[buf.c.x:]))
			buf.c.x -= sz
		} else if buf.c.y > 0 {
			prev := lines[buf.c.y-1].String()
			buf.c.x = len(prev)
			mut.Update(buf.c.y-1, content.TextLine(prev+line))
			mut.Delete(buf.c.y)
			buf.c.y--
		}

	// Enter.
	case '\r':
		mut.Update(buf.c.y, content.TextLine(line[:buf.c.x]))
		mut.Insert(buf.c.y+1, content.TextLine(line[buf.c.x:]))
		buf.c.y++
		buf.c.x = 0

	// Printable ASCII.
	default:
		if inputs.IsTab(b) || ch >= 0x20 {
			mut.Update(buf.c.y, content.TextLine(line[:buf.c.x]+string(ch)+line[buf.c.x:]))
			buf.c.x++
		}
	}

	return
}
