package inputs

type CursorArrow byte

const (
	CursorArrowUp CursorArrow = iota
	CursorArrowDown
	CursorArrowLeft
	CursorArrowRight
)

func IsArrow(b []byte, arrowType *CursorArrow) bool {
	if len(b) != 3 || b[0] != Escape || b[1] != '[' {
		return false
	}
	switch b[2] {
	case 'A':
		*arrowType = CursorArrowUp
	case 'B':
		*arrowType = CursorArrowDown
	case 'C':
		*arrowType = CursorArrowRight
	case 'D':
		*arrowType = CursorArrowLeft
	default:
		return false
	}
	return true
}

const Escape = 0x1b

func IsEscape(b []byte) bool { return len(b) == 1 && b[0] == Escape }

func IsTab(b []byte) bool { return len(b) == 1 && b[0] == '\t' }
