package editor

import (
	"unicode"
	"unicode/utf8"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/input"
)

func insertInput(buf *Buffer, k input.Key, prefs *RenderPrefs) {
	switch k.Special {
	case input.CursorMove:
		buf.handleCursor(k.Cursor, k.Mod, prefs)
		return

	case input.Esc:
		if buf.hasSelection() {
			buf.cancelSelection()
			return
		}
		buf.cancelSelection()
		buf.mode = ModeNormal
		if c := buf.c(); c.Col > 0 {
			c.Col-- // Land on the last typed character.
			buf.updateCursor(c)
		}
		return

	case input.Text:
		if k.Mod != 0 {
			return // Not a typed character.
		}
	case input.Enter, input.Tab, input.Backspace:
	default:
		return
	}

	if !buf.canEdit() {
		return
	}
	if buf.hasSelection() {
		// The input replaces the selection. The cursor may move to another line.
		DeleteSelection.DoOnBuffer(buf, *prefs)
		if k.Special == input.Backspace {
			return
		}
	}
	buf.cancelSelection() // An empty selection is not replaced, and the edit invalidates it.
	lines := buf.Content.Lines()
	c := buf.c()
	line := lines[c.Line].String()
	mut := buf.Mutate()

	switch k.Special {
	case input.Backspace:
		if c.Col > 0 {
			_, sz := utf8.DecodeLastRuneInString(line[:c.Col])
			mut.Update(c.Line, content.TextLine(line[:c.Col-sz]+line[c.Col:]))
			buf.updateCursor(content.Position{Line: c.Line, Col: c.Col - sz})
		} else if c.Line > 0 {
			prev := lines[c.Line-1].String()
			mut.Update(c.Line-1, content.TextLine(prev+line))
			mut.Delete(c.Line)
			buf.updateCursor(content.Position{Line: c.Line - 1, Col: len(prev)})
		}
		return

	case input.Enter:
		mut.Update(c.Line, content.TextLine(line[:c.Col]))
		mut.Insert(c.Line+1, content.TextLine(line[c.Col:]))
		buf.updateCursor(content.Position{Line: c.Line + 1})
		return

	case input.Tab:
		insertContent(buf, []byte{'\t'}, mut, line, 1)
		return
	}

	switch ch := k.Rune; ch {
	// Brackets.
	case '{', '(', '[':
		insertContent(buf, []byte{byte(ch), bracketPair(byte(ch))}, mut, line, 1)
	case '}', ')', ']':
		if isRepeatedBracket(buf, line, byte(ch)) {
			buf.updateCursor(content.Position{Line: c.Line, Col: c.Col + 1})
		} else {
			insertContent(buf, []byte{byte(ch)}, mut, line, 1)
		}
	case '"', '\'', '`':
		if isRepeatedBracket(buf, line, byte(ch)) {
			buf.updateCursor(content.Position{Line: c.Line, Col: c.Col + 1})
		} else {
			insertContent(buf, []byte{byte(ch), bracketPair(byte(ch))}, mut, line, 1)
		}

	default:
		if unicode.IsPrint(ch) {
			text := []byte(string(ch))
			insertContent(buf, text, mut, line, len(text))
		}
	}
}

func isRepeatedBracket(buf *Buffer, line string, ch byte) bool {
	c := buf.c()
	return 0 < c.Col && c.Col < len(line) &&
		line[c.Col] == ch && line[c.Col-1] == bracketPair(ch)
}

func insertContent(buf *Buffer, b []byte, mut content.Mutable, line string, advanceCursor int) {
	c := buf.c()
	if len(b) > 0 {
		mut.Update(c.Line, content.TextLine(line[:c.Col]+string(b)+line[c.Col:]))
	}
	buf.updateCursor(content.Position{Line: c.Line, Col: c.Col + advanceCursor})
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
