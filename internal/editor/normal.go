package editor

import (
	"unicode/utf8"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/inputs"
)

func normalInput(buf *Buffer, b []byte, prefs *RenderPrefs) (quit bool) {
	var (
		arrow inputs.CursorArrow
		mod   inputs.Modifier
	)
	if inputs.IsArrow(b, &arrow, &mod) {
		buf.handleCursor(arrow, mod, prefs)
		return false
	}

	lines := buf.Content.Lines()
	switch {
	case len(b) == 1:
		switch b[0] {
		// Quit.
		case 'q':
			return !buf.dirty

		// Tab size.
		case '+', '=':
			prefs.tabsScaleUp()
		case '-':
			prefs.tabsScaleDown()

		// Esc: clear selection.
		case '\x1b':
			buf.cancelSelection()

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
			buf.c.Line++
			buf.Mutate().Insert(buf.c.Line, content.TextLine(""))
			buf.c.Col = 0
			buf.mode = ModeInsert

		// Enter command-line mode.
		case ':':
			buf.cmdline = ""
			buf.mode = ModeCommand

		// Engage.
		case '\r':
			if action, ok := lines[buf.c.Line].(content.LineAction); ok {
				action.Engage()
			} else {
				RelMove{Dy: 1}.DoOnBuffer(buf, *prefs)
			}

		// Delete.
		case 'x':
			if !buf.canEdit() {
				return false
			}
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
	}
	return false
}
