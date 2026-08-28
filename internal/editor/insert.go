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
		if len(buf.sel) > 0 {
			buf.cancelSelection()
			return
		}

		buf.mode = ModeNormal
		if buf.c.Col > 0 {
			buf.c.Col-- // Land on the last typed character.
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
	line := lines[buf.c.Line].String()
	mut := buf.Mutate()

	switch ch := b[0]; ch {
	// Backspace.
	case 0x7f, 0x08:
		if buf.c.Col > 0 {
			_, sz := utf8.DecodeLastRuneInString(line[:buf.c.Col])
			mut.Update(buf.c.Line, content.TextLine(line[:buf.c.Col-sz]+line[buf.c.Col:]))
			buf.c.Col -= sz
		} else if buf.c.Line > 0 {
			prev := lines[buf.c.Line-1].String()
			buf.c.Col = len(prev)
			mut.Update(buf.c.Line-1, content.TextLine(prev+line))
			mut.Delete(buf.c.Line)
			buf.c.Line--
		}

	// Enter.
	case '\r':
		mut.Update(buf.c.Line, content.TextLine(line[:buf.c.Col]))
		mut.Insert(buf.c.Line+1, content.TextLine(line[buf.c.Col:]))
		buf.c.Line++
		buf.c.Col = 0

	// Printable ASCII.
	default:
		if inputs.IsTab(b) || ch >= 0x20 {
			mut.Update(buf.c.Line, content.TextLine(line[:buf.c.Col]+string(ch)+line[buf.c.Col:]))
			buf.c.Col++
		}
	}

	return
}
