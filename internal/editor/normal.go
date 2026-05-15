package editor

import "rmazur.io/x/edit/internal/content"

func normalInput(buf *Buffer, b []byte, prefs *RenderPrefs) (quit bool) {
	lines := buf.content.Lines()
	switch {
	case len(b) == 1:
		switch b[0] {
		// Quit.
		case 'q':
			return true

		// Switch mode.
		case 'i':
			buf.mode = ModeInsert
		case 'a':
			buf.cx++
			buf.mode = ModeInsert
		case 'A':
			buf.cx = buf.content.Lines()[buf.cy].Len()
			buf.mode = ModeInsert
		case 'o':
			if !buf.canEdit() {
				return false
			}
			buf.cy++
			buf.mutate().Insert(buf.cy, content.TextLine(""))
			buf.cx = 0
			buf.mode = ModeInsert

		// Engage.
		case '\r':
			if action, ok := lines[buf.cy].(content.LineAction); ok {
				action.Engage()
			} else {
				RelMove{Dy: 1}.Do(buf, *prefs)
			}

		// Delete.
		case 'x':
			if !buf.canEdit() {
				return false
			}
			line := lines[buf.cy].String()
			if buf.cx < len(line) {
				buf.mutate().Update(buf.cy, content.TextLine(line[:buf.cx]+line[buf.cx+1:]))
			}

		// Commands.
		case 'h':
			RelMove{Dx: -1}.Do(buf, *prefs)
		case 'l':
			RelMove{Dx: 1}.Do(buf, *prefs)
		case 'j':
			RelMove{Dy: 1}.Do(buf, *prefs)
		case 'k':
			RelMove{Dy: -1}.Do(buf, *prefs)
		case '0':
			MoveHome.Do(buf, *prefs)
		case '$':
			MoveEnd.Do(buf, *prefs)
		}
	}
	return false
}
