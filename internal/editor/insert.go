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
		if len(buf.sel) > 0 {
			buf.cancelSelection()
			return
		}
		buf.mode = ModeNormal
		if buf.c.Col > 0 {
			buf.c.Col-- // Land on the last typed character.
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
	if len(buf.sel) > 0 {
		// The input replaces the selection. The cursor may move to another line.
		DeleteSelection.DoOnBuffer(buf, *prefs)
		if k.Special == input.Backspace {
			return
		}
	}
	lines := buf.Content.Lines()
	line := lines[buf.c.Line].String()
	mut := buf.Mutate()

	switch k.Special {
	case input.Backspace:
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
		return

	case input.Enter:
		mut.Update(buf.c.Line, content.TextLine(line[:buf.c.Col]))
		mut.Insert(buf.c.Line+1, content.TextLine(line[buf.c.Col:]))
		buf.c.Line++
		buf.c.Col = 0
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
			buf.c.Col++
		} else {
			insertContent(buf, []byte{byte(ch)}, mut, line, 1)
		}
	case '"', '\'', '`':
		if isRepeatedBracket(buf, line, byte(ch)) {
			buf.c.Col++
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
