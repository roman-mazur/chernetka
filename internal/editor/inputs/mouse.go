package inputs

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/go-errors/errors"
)

func IsMouseInput(b []byte) bool {
	if len(b) < 3 {
		return false
	}
	return b[0] == Escape && b[1] == '[' && b[2] == '<'
}

// Mouse click info.
type Mouse struct {
	Button  MouseButton
	X, Y    int
	Pressed bool // False means button is released.
	Mod     MouseModifier
}

type MouseButton byte

const (
	MouseButtonLeft MouseButton = iota
	MouseButtonMiddle
	MouseButtonRight
	MouseButtonNone // used for hover
)

// MouseModifier is a bit mask giving more context about the Mouse event.
// bit	meaning
// 0	Shift
// 1	Alt
// 2	Ctrl
// 3	Motion
type MouseModifier byte

func (mm MouseModifier) HasShift() bool  { return mm&1 != 0 }
func (mm MouseModifier) HasAlt() bool    { return mm&2 != 0 }
func (mm MouseModifier) HasCtrl() bool   { return mm&4 != 0 }
func (mm MouseModifier) HasMotion() bool { return mm&8 != 0 }

var ErrorNotMouse = errors.New("not a mouse input")

func ReadMouse(in *bufio.Reader) (data Mouse, err error) {
	var prefix []byte
	prefix, err = in.Peek(1)
	if err != nil {
		return
	}
	if !IsEscape(prefix) {
		err = ErrorNotMouse
		return
	}

	prefix, err = in.Peek(3)
	if err != nil {
		return
	}
	if !IsMouseInput(prefix[:]) {
		err = ErrorNotMouse
		return
	}
	err = discard(in, 3)
	if err != nil {
		return
	}

	var (
		B   int
		sep byte
	)

	B, sep, err = mouseParseNextInt(in)
	if err != nil {
		return
	}
	if sep != ';' {
		err = fmt.Errorf("expected ';', got '%c'", sep)
		return
	}
	data.Button = MouseButton(B & 3)
	data.Mod = MouseModifier((B >> 2) & 0xf)

	data.X, sep, err = mouseParseNextInt(in)
	if err != nil {
		return
	}
	if sep != ';' {
		err = fmt.Errorf("expected ';', got '%c'", sep)
		return
	}

	data.Y, sep, err = mouseParseNextInt(in)
	if err != nil {
		return
	}
	if sep != 'm' && sep != 'M' {
		err = fmt.Errorf("expected 'm' or 'M' in the end of mouse input, got '%c'", sep)
		return
	}

	data.Pressed = sep == 'M'
	return
}

func discard(in *bufio.Reader, x int) (err error) {
	for x > 0 {
		var n int
		n, err = in.Discard(x)
		if err != nil {
			return
		}
		x -= n
	}
	return
}

func mouseParseNextInt(in *bufio.Reader) (int, byte, error) {
	var digits []byte
	for {
		b, err := in.ReadByte()
		if err != nil {
			return 0, 0, err
		}
		if b < '0' || b > '9' {
			n, err := strconv.Atoi(string(digits))
			return n, b, err
		}
		digits = append(digits, b)
	}
}
