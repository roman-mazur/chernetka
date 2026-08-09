package inputs

import (
	"bufio"
	"bytes"
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

func (m *Mouse) String() string {
	return fmt.Sprintf("{b:%d, coords:(%d,%d), p:%t, m:%s}", m.Button, m.X, m.Y, m.Pressed, m.Mod)
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
// 4	Scroll wheel
type MouseModifier byte

func (mm MouseModifier) HasShift() bool  { return mm&1 != 0 }
func (mm MouseModifier) HasAlt() bool    { return mm&2 != 0 }
func (mm MouseModifier) HasCtrl() bool   { return mm&4 != 0 }
func (mm MouseModifier) HasMotion() bool { return mm&8 != 0 }
func (mm MouseModifier) HasWheel() bool  { return mm&16 != 0 }

func (mm MouseModifier) SrollDirection(m Mouse) ScrollDirection { return ScrollDirection(m.Button) }

func (mm MouseModifier) String() string {
	var data [5]byte
	for i := range data {
		data[i] = '-'
	}
	if mm.HasShift() {
		data[0] = 's'
	}
	if mm.HasAlt() {
		data[1] = 'a'
	}
	if mm.HasCtrl() {
		data[2] = 'c'
	}
	if mm.HasMotion() {
		data[3] = 'm'
	}
	if mm.HasWheel() {
		data[4] = 'w'
	}
	return string(data[:])
}

type ScrollDirection byte

const (
	ScrollDirectionUp ScrollDirection = iota
	ScrollDirectionDown
	S
	ScrollDirectionLeft
	ScrollDirectionRight
)

var ErrorNotMouse = errors.New("not a mouse input")

func ReadMouse(inData []byte) (data Mouse, n int, err error) {
	if !IsMouseInput(inData) {
		err = ErrorNotMouse
		return
	}
	n = 3
	in := bufio.NewReader(bytes.NewReader(inData[n:]))

	var (
		B   int
		sep byte
		k   int
	)

	B, sep, k, err = mouseParseNextInt(in)
	n += k
	if err != nil {
		return
	}
	if sep != ';' {
		err = fmt.Errorf("expected ';', got '%c'", sep)
		return
	}
	data.Button = MouseButton(B & 3)
	data.Mod = MouseModifier((B >> 2) & 0xff)

	data.X, sep, k, err = mouseParseNextInt(in)
	n += k
	if err != nil {
		return
	}
	if sep != ';' {
		err = fmt.Errorf("expected ';', got '%c'", sep)
		return
	}

	data.Y, sep, k, err = mouseParseNextInt(in)
	n += k
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

func mouseParseNextInt(in *bufio.Reader) (int, byte, int, error) {
	var digits []byte
	for {
		b, err := in.ReadByte()
		if err != nil {
			return 0, 0, 0, err
		}
		if b < '0' || b > '9' {
			n, err := strconv.Atoi(string(digits))
			return n, b, len(digits) + 1, err
		}
		digits = append(digits, b)
	}
}
