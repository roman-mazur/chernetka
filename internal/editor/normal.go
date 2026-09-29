package editor

import (
	"unicode/utf8"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/input"
)

func normalInput(buf *Buffer, k input.Key, prefs *RenderPrefs) (quit bool) {
	switch k.Special {
	case input.CursorMove:
		buf.handleCursor(k.Cursor, k.Mod, prefs)

	// Esc: clear selection and search highlights.
	case input.Esc:
		buf.cancelSelection()
		buf.search = nil

	// Engage.
	case input.Enter:
		RelMove{Dy: 1}.DoOnBuffer(buf, *prefs)

	case input.Text:
		if k.Mod == 0 {
			return normalCommand(buf, k.Rune, prefs)
		}
	}
	return false
}

// normalCommand runs the command typed with a single character.
func normalCommand(buf *Buffer, r rune, prefs *RenderPrefs) (quit bool) {
	lines := buf.Content.Lines()
	switch r {
	// Quit.
	case 'q':
		return !buf.dirty

	// Tab size.
	case '+', '=':
		prefs.tabsScaleUp()
	case '-':
		prefs.tabsScaleDown()

	// Search.
	case 'n':
		buf.searchMove(1)
	case 'N':
		buf.searchMove(-1)

	// Copy to clipboard.
	case 'y':
		ClipboardCopy.DoOnBuffer(buf, *prefs)

	// Switch mode.
	case 'i':
		buf.mode = ModeInsert
	case 'a':
		line := lines[buf.c.Line].String()
		if buf.c.Col < len(line) {
			_, sz := utf8.DecodeRuneInString(line[buf.c.Col:])
			buf.c.Col += sz
		}
		buf.mode = ModeInsert
	case 'A':
		buf.c.Col = buf.Content.Lines()[buf.c.Line].Len()
		buf.mode = ModeInsert
	case 'o':
		if !buf.canEdit() {
			return false
		}
		buf.cancelSelection()
		buf.c.Line++
		buf.Mutate().Insert(buf.c.Line, content.TextLine(""))
		buf.c.Col = 0
		buf.mode = ModeInsert

	// Delete.
	case 'x':
		if !buf.canEdit() {
			return false
		}
		if buf.hasSelection() {
			DeleteSelection.DoOnBuffer(buf, *prefs)
			return false
		}
		buf.cancelSelection()
		line := lines[buf.c.Line].String()
		if buf.c.Col < len(line) {
			_, sz := utf8.DecodeRuneInString(line[buf.c.Col:])
			buf.Mutate().Update(buf.c.Line, content.TextLine(line[:buf.c.Col]+line[buf.c.Col+sz:]))
		}

	// Commands.
	case 'h':
		RelMove{Dx: -1}.DoOnBuffer(buf, *prefs)
	case 'l':
		RelMove{Dx: 1}.DoOnBuffer(buf, *prefs)
	case 'j':
		RelMove{Dy: 1}.DoOnBuffer(buf, *prefs)
	case 'k':
		RelMove{Dy: -1}.DoOnBuffer(buf, *prefs)
	case '0':
		MoveHome.DoOnBuffer(buf, *prefs)
	case '$':
		MoveEnd.DoOnBuffer(buf, *prefs)
	case 'g':
		MoveContentStart.DoOnBuffer(buf, *prefs)
	case 'G':
		MoveContentEnd.DoOnBuffer(buf, *prefs)
	// Page scrolling forward and backward. This will keep top or bottom line visible.
	case ' ':
		ScreenMove{ScreenD: 1}.DoOnBuffer(buf, *prefs)
	case 'b':
		ScreenMove{ScreenD: -1}.DoOnBuffer(buf, *prefs)
	}
	return false
}
