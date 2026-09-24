package editor

import (
	"unicode/utf8"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/inputs"
)

func insertInput(buf *Buffer, b []byte, prefs *RenderPrefs) {
	if len(b) == 0 {
		return
	}

	var (
		arrow inputs.Cursor
		mod   inputs.Modifier
	)
	if inputs.IsCursor(b, &arrow, &mod) {
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

	if !buf.canEdit() {
		return
	}
	lines := buf.Content.Lines()
	line := lines[buf.c.Line].String()
	mut := buf.Mutate()

	if len(buf.sel) > 0 {
		DeleteSelection.DoOnBuffer(buf, *prefs)
	}

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

	// Brackets.
	case '{', '(', '[':
		insertContent(buf, []byte{ch, bracketPair(ch)}, mut, line, 1)
	case '}', ')', ']':
		if isRepeatedBracket(buf, line, ch) {
			buf.c.Col++
		} else {
			insertContent(buf, b, mut, line, len(b))
		}
	case '"', '\'', '`':
		if isRepeatedBracket(buf, line, ch) {
			buf.c.Col++
		} else {
			insertContent(buf, []byte{ch, bracketPair(ch)}, mut, line, 1)
		}

	// Printable ASCII.
	default:
		if inputs.IsTab(b) || ch >= 0x20 {
			insertContent(buf, b, mut, line, len(b))
		}
	}

	return
}

func isRepeatedBracket(buf *Buffer, line string, ch byte) bool {
	return 0 < buf.c.Col && buf.c.Col < len(line) &&
		line[buf.c.Col] == ch && line[buf.c.Col-1] == bracketPair(ch)
}

func insertContent(buf *Buffer, b []byte, mut content.Mutable, line string, advanceCursor int) {
	if len(b) > 0 {
		mut.Update(buf.c.Line, content.TextLine(line[:buf.c.Col]+string(b)+line[buf.c.Col:]))
	}
	buf.c.Col += advanceCursor
}

func bracketPair(b byte) byte {
	switch b {
	case '{':
		return '}'
	case '[':
		return ']'
	case '(':
		return ')'
	case '}':
		return '{'
	case ']':
		return '['
	case ')':
		return '('
	}
	return b
}
