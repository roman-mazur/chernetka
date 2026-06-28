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

		// Page scrolling forward and backward. This will keep top or bottom line visible.
		case ' ':
			ScreenMove{ScreenD: 1}.DoOnBuffer(buf, *prefs)
		case 'b':
			ScreenMove{ScreenD: -1}.DoOnBuffer(buf, *prefs)

		// Tab size.
		case '+', '=':
			prefs.tabsScaleUp()
		case '-':
			prefs.tabsScaleDown()

		// Switch mode.
		case 'i':
			buf.mode = ModeInsert
		case 'a':
			line := lines[buf.cy].String()
			if buf.cx < len(line) {
				_, sz := utf8.DecodeRuneInString(line[buf.cx:])
				buf.cx += sz
			}
			buf.mode = ModeInsert
		case 'A':
			buf.cx = buf.Content.Lines()[buf.cy].Len()
			buf.mode = ModeInsert
		case 'o':
			if !buf.canEdit() {
				return false
			}
			buf.cy++
			buf.Mutate().Insert(buf.cy, content.TextLine(""))
			buf.cx = 0
			buf.mode = ModeInsert

		// Enter command-line mode.
		case ':':
			buf.cmdline = ""
			buf.mode = ModeCommand

		// Engage.
		case '\r':
			if action, ok := lines[buf.cy].(content.LineAction); ok {
				action.Engage()
			} else {
				RelMove{Dy: 1}.DoOnBuffer(buf, *prefs)
			}

		// Delete.
		case 'x':
			if !buf.canEdit() {
				return false
			}
			line := lines[buf.cy].String()
			if buf.cx < len(line) {
				_, sz := utf8.DecodeRuneInString(line[buf.cx:])
				buf.Mutate().Update(buf.cy, content.TextLine(line[:buf.cx]+line[buf.cx+sz:]))
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
		}
	}
	return false
}
