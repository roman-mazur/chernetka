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
	Mod     Modifier
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
	data.Mod = Modifier((B >> 2) & 0xff)

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
