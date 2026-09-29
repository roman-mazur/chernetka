package input

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
)

func IsMouseInput(b []byte) bool {
	if len(b) < 3 {
		return false
	}
	return b[0] == escByte && b[1] == '[' && b[2] == '<'
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

// Press returns the press of the button at the 1-based screen column x and row y.
func Press(b MouseButton, x, y int) Mouse { return Mouse{Button: b, X: x, Y: y, Pressed: true} }

// Release returns the release of the button at the 1-based screen column x and row y.
func Release(b MouseButton, x, y int) Mouse { return Mouse{Button: b, X: x, Y: y} }

// Drag returns the motion with the button pressed at the 1-based screen column x and row y.
func Drag(b MouseButton, x, y int) Mouse {
	return Mouse{Button: b, X: x, Y: y, Pressed: true, Mod: ModMotion}
}

// Hover returns the motion with no buttons pressed at the 1-based screen column x and row y.
func Hover(x, y int) Mouse { return Mouse{Button: MouseButtonNone, X: x, Y: y, Mod: ModMotion} }

// Scroll returns the wheel event at the 1-based screen column x and row y.
func Scroll(dir ScrollDirection, x, y int) Mouse {
	return Mouse{Button: MouseButton(dir), X: x, Y: y, Pressed: true, Mod: ModWheel}
}

// Click returns the terminal input of the button press and release at the 1-based screen column x and row y.
func Click(b MouseButton, x, y int) string {
	return Press(b, x, y).Encode() + Release(b, x, y).Encode()
}

// With returns the event with the modifiers added, like ModCtrl for Ctrl+click.
func (m Mouse) With(mod Modifier) Mouse {
	m.Mod |= mod
	return m
}

// Encode returns the terminal input of the event in the SGR encoding decoded by ReadMouse.
func (m Mouse) Encode() string {
	end := 'M'
	if !m.Pressed {
		end = 'm'
	}
	return fmt.Sprintf("\x1b[<%d;%d;%d%c", int(m.Button)|int(m.Mod)<<2, m.X, m.Y, end)
}

type MouseButton byte

const (
	MouseButtonLeft MouseButton = iota
	MouseButtonMiddle
	MouseButtonRight
	MouseButtonNone // used for hover
)

// ScrollDirection is the wheel "button" reported with the wheel modifier: 64-67 in the SGR encoding.
type ScrollDirection byte

const (
	ScrollDirectionUp ScrollDirection = iota
	ScrollDirectionDown
	ScrollDirectionLeft
	ScrollDirectionRight
)

var (
	ErrorNotMouse = errors.New("not a mouse input")
	// ErrIncomplete is returned when the input ends in the middle of a sequence,
	// which is to be completed by the next read.
	ErrIncomplete = errors.New("incomplete input")
)

// ReadMouse reads the mouse event at the start of the input returning the number of its bytes.
func ReadMouse(inData []byte) (data Mouse, n int, err error) {
	if !IsMouseInput(inData) {
		err = ErrorNotMouse
		return
	}
	defer func() {
		if errors.Is(err, io.EOF) {
			err = ErrIncomplete
		}
	}()
	n = 3
	in := bufio.NewReader(bytes.NewReader(inData[n:]))

	var (
		B   int
		sep byte
		k   int
	)

	B, sep, k, err = termParseNextInt(in)
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

	data.X, sep, k, err = termParseNextInt(in)
	n += k
	if err != nil {
		return
	}
	if sep != ';' {
		err = fmt.Errorf("expected ';', got '%c'", sep)
		return
	}

	data.Y, sep, k, err = termParseNextInt(in)
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

func termParseNextInt(in io.Reader) (int, byte, int, error) {
	var (
		digits []byte
		buf    [1]byte
	)
	for {
		bc, err := in.Read(buf[:])
		if err != nil {
			return 0, 0, len(digits) + bc, err
		}
		if bc == 0 {
			continue
		}

		if buf[0] < '0' || buf[0] > '9' {
			n, err := strconv.Atoi(string(digits))
			return n, buf[0], len(digits) + 1, err
		}
		digits = append(digits, buf[0])
	}
}
