package editor

import "rmazur.io/x/edit/internal/content"

func normalInput(buf *Buffer, b []byte) (quit bool) {
	lines := buf.content.Lines()
	switch {
	case len(b) == 1:
		switch b[0] {
		case 'h':
			buf.cx--
		case 'l':
			buf.cx++
		case 'j':
			buf.cy++
		case 'k':
			buf.cy--
		case '0':
			buf.cx = 0
		case '$':
			buf.cx = lines[buf.cy].Len()
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
		case 'x':
			if !buf.canEdit() {
				return false
			}
			line := lines[buf.cy].String()
			if buf.cx < len(line) {
				buf.mutate().Update(buf.cy, content.TextLine(line[:buf.cx]+line[buf.cx+1:]))
			}
		case 'q':
			return true
		}
	}
	return false
}
