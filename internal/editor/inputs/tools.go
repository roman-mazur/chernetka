package inputs

import (
	"bytes"
)

type Cursor byte

//go:generate go run golang.org/x/tools/cmd/stringer -type=Cursor -trimprefix=Cursor

const (
	CursorArrowUp Cursor = iota
	CursorArrowDown
	CursorArrowLeft
	CursorArrowRight
	CursorHome
	CursorEnd
)

// Modifier is a bit mask giving more context about the Mouse or keyboard event.
// bit	meaning
// 0	Shift
// 1	Alt
// 2	Ctrl
// 3	Motion / command
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

func IsCursor(b []byte, cursorType *Cursor, mod *Modifier) bool {
	if len(b) == 1 {
		switch b[0] {
		case 1:
			*cursorType = CursorHome
			return true
		case 5:
			*cursorType = CursorEnd
			return true
		}
		return false
	}

	if len(b) < 3 || b[0] != Escape || b[1] != '[' {
		return false
	}
	if len(b) == 3 {
		return cursorArrow(b[2], cursorType)
	}
	if len(b) < 6 {
		return false
	}
	if b[2] == '1' && b[3] == ';' {
		mn, typC, _, err := termParseNextInt(bytes.NewReader(b[4:]))
		if err != nil {
			return false
		}
		*mod = Modifier(mn - 1)
		if cursorArrow(typC, cursorType) {
			if *cursorType == CursorArrowLeft && mod.HasMotion() {
				*cursorType = CursorHome
			}
			if *cursorType == CursorArrowRight && mod.HasMotion() {
				*cursorType = CursorEnd
			}
			return true
		}
	}
	return false
}

func cursorArrow(b byte, c *Cursor) bool {
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

func IsSaveCommand(b []byte) bool {
	return len(b) == 1 && b[0] == 0x13 // Ctrl+S
}

// ClipboardOp encodes one of the clipboard operations (copy/paste/cut).
type ClipboardOp byte

const (
	ClipboardOpNone ClipboardOp = iota
	ClipboardOpCopy
	ClipboardOpPaste
	ClipboardOpCut
)

func IsClipboardOp(b []byte, op *ClipboardOp) bool {
	*op = ClipboardOpNone
	if len(b) != 1 {
		return false
	}
	switch b[0] {
	case 0x3:
		*op = ClipboardOpCopy
		return true
	case 0x16:
		*op = ClipboardOpPaste
		return true
	case 0x18:
		*op = ClipboardOpCut
		return true
	}
	return false
}
