// Package escape exposes functions to write escape sequences supported by the terminal apps.
package escape

import (
	"fmt"
	"image/color"
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

func ColorText(out io.Writer, text string, color color.Color) {
	_, _ = fmt.Fprintf(out, colorTemplate(color, colorPurposeForeground), text)
}

func ColorBackground(out io.Writer, text string, color color.Color) {
	_, _ = fmt.Fprintf(out, colorTemplate(color, colorPurposeBackground), text)
}

func colorTemplate(c color.Color, purpose colorPurpose) string {
	if c == nil {
		return "%s"
	}
	r, g, b, _ := c.RGBA()
	code := 38 // foreground color
	if purpose == colorPurposeBackground {
		code = 48
	}
	return fmt.Sprintf("\x1b[%d;2;%d;%d;%dm%%s\x1b[0m", code, r, g, b)
}

type colorPurpose byte

const (
	colorPurposeForeground colorPurpose = iota
	colorPurposeBackground
)

func DisableLineWrapping(out io.Writer) (restore func()) {
	return applyPair(out, "\x1b[?7l", "\x1b[?7h")
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
