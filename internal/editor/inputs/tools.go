package inputs

import "strconv"

type CursorArrow byte

const (
	CursorArrowUp CursorArrow = iota
	CursorArrowDown
	CursorArrowLeft
	CursorArrowRight
)

// Modifier is a bit mask giving more context about the Mouse or keyboard event.
// bit	meaning
// 0	Shift
// 1	Alt
// 2	Ctrl
// 3	Motion
// 4	Scroll wheel
type Modifier byte

func (m Modifier) HasShift() bool  { return m&1 != 0 }
func (m Modifier) HasAlt() bool    { return m&2 != 0 }
func (m Modifier) HasCtrl() bool   { return m&4 != 0 }
func (m Modifier) HasMotion() bool { return m&8 != 0 }
func (m Modifier) HasWheel() bool  { return m&16 != 0 }

func (m Modifier) SrollDirection(mouse Mouse) ScrollDirection { return ScrollDirection(mouse.Button) }

func (m Modifier) String() string {
	var data [5]byte
	for i := range data {
		data[i] = '-'
	}
	if m.HasShift() {
		data[0] = 's'
	}
	if m.HasAlt() {
		data[1] = 'a'
	}
	if m.HasCtrl() {
		data[2] = 'c'
	}
	if m.HasMotion() {
		data[3] = 'm'
	}
	if m.HasWheel() {
		data[4] = 'w'
	}
	return string(data[:])
}

func IsArrow(b []byte, arrowType *CursorArrow, mod *Modifier) bool {
	if len(b) < 3 || b[0] != Escape || b[1] != '[' {
		return false
	}
	if len(b) == 3 {
		return cursorArrow(b[2], arrowType)
	}
	if len(b) < 6 {
		return false
	}
	if b[2] == '1' && b[3] == ';' {
		mn, err := strconv.ParseInt(string(b[4:5]), 10, 64)
		if err != nil {
			return false
		}
		*mod = Modifier(mn - 1)
		return cursorArrow(b[5], arrowType)
	}
	return false
}

func cursorArrow(b byte, c *CursorArrow) bool {
	switch b {
	case 'A':
		*c = CursorArrowUp
	case 'B':
		*c = CursorArrowDown
	case 'C':
		*c = CursorArrowRight
	case 'D':
		*c = CursorArrowLeft
	default:
		return false
	}
	return true
}

const Escape = 0x1b

func IsEscape(b []byte) bool { return len(b) == 1 && b[0] == Escape }

func IsTab(b []byte) bool { return len(b) == 1 && b[0] == '\t' }
