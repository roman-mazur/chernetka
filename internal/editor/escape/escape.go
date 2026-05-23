// Package escape exposes functions to write escape sequences supported by the terminal apps.
package escape

import (
	"fmt"
	"io"
)

func EnableAlternativeBuffer(out io.Writer) (restore func()) {
	return applyPair(out, "\x1b[?1049h", "\x1b[?1049l")
}

func ReverseVideo(out io.Writer) (restore func()) {
	return applyPair(out, "\x1b[7m", "\x1b[0m")
}

func HideCursor(out io.Writer) (restore func()) {
	return applyPair(out, "\x1b[?25l", "\x1b[?25h")
}

func MoveTopLeft(out io.Writer) {
	_, _ = fmt.Fprint(out, "\x1b[H")
}

func SetCursorPosition(out io.Writer, row, col int) {
	_, _ = fmt.Fprintf(out, "\x1b[%d;%dH", row, col)
}

func ClearLine(out io.Writer) {
	_, _ = fmt.Fprint(out, "\x1b[2K")
}

func applyPair(out io.Writer, action, revert string) (restore func()) {
	_, err := fmt.Fprint(out, action)
	if err == nil {
		return func() {
			_, _ = fmt.Fprint(out, revert)
		}
	}
	return noop
}

var noop = func() {}
