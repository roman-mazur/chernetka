package editor

import (
	"unicode/utf8"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/inputs"
)

func normalInput(buf *Buffer, b []byte, prefs *RenderPrefs) (quit bool) {
	var arrow inputs.CursorArrow
	if inputs.IsArrow(b, &arrow) {
		buf.handleCursor(arrow, prefs)
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

		// Switch mode.
		case 'i':
			buf.mode = ModeInsert
		case 'a':
			line := lines[buf.c.y].String()
			if buf.c.x < len(line) {
				_, sz := utf8.DecodeRuneInString(line[buf.c.x:])
				buf.c.x += sz
			}
			buf.mode = ModeInsert
		case 'A':
			buf.c.x = buf.Content.Lines()[buf.c.y].Len()
			buf.mode = ModeInsert
		case 'o':
			if !buf.canEdit() {
				return false
			}
			buf.c.y++
			buf.Mutate().Insert(buf.c.y, content.TextLine(""))
			buf.c.x = 0
			buf.mode = ModeInsert

		// Enter command-line mode.
		case ':':
			buf.cmdline = ""
			buf.mode = ModeCommand

		// Engage.
		case '\r':
			if action, ok := lines[buf.c.y].(content.LineAction); ok {
				action.Engage()
			} else {
				RelMove{Dy: 1}.DoOnBuffer(buf, *prefs)
			}

		// Delete.
		case 'x':
			if !buf.canEdit() {
				return false
			}
			line := lines[buf.c.y].String()
			if buf.c.x < len(line) {
				_, sz := utf8.DecodeRuneInString(line[buf.c.x:])
				buf.Mutate().Update(buf.c.y, content.TextLine(line[:buf.c.x]+line[buf.c.x+sz:]))
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
