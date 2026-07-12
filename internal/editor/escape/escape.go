// Package escape exposes functions to write escape sequences supported by the terminal apps.
package escape

import (
	"fmt"
	"image/color"
	"io"
	"strconv"
)

func SyncOutput(out io.Writer) (restore func()) {
	return applyPair(out, "\x1b[?2026h", "\x1b[?2026l")
}

func EnableAlternativeBuffer(out io.Writer) (restore func()) {
	return applyPair(out, "\x1b[?1049h", "\x1b[?1049l")
}

func EnableMouse(out io.Writer) (restore func()) {
	return applyPair(out, "\x1b[?1002h\x1b[?1006h", "\x1b[?1002l\x1b[?1006l")
}

func ReverseVideo(out io.Writer) (restore func()) {
	return applyPair(out, "\x1b[7m", "\x1b[0m")
}

func HideCursor(out io.Writer, thinCursor bool) (restore func()) {
	styleCode := 1
	if thinCursor {
		styleCode = 5
	}
	restoreSeq := fmt.Sprintf("\x1b[%d q\x1b[?25h", styleCode)
	return applyPair(out, "\x1b[?25l", restoreSeq)
}

func MoveTopLeft(out io.Writer) {
	_, _ = io.WriteString(out, "\x1b[H")
}

func SetCursorPosition(out io.Writer, row, col int) {
	_, _ = fmt.Fprintf(out, "\x1b[%d;%dH", row, col)
}

func ClearLine(out io.Writer) {
	_, _ = io.WriteString(out, "\x1b[2K")
}

func ColorText(out io.Writer, text string, fg, bg color.Color) {
	if fg == nil && bg == nil {
		_, _ = io.WriteString(out, text)
		return
	}

	_, _ = io.WriteString(out, "\x1b[")
	if fg != nil {
		_, _ = io.WriteString(out, "38;2;")
		write8bitColor(out, fg)
		if bg != nil {
			_, _ = io.WriteString(out, ";")
		}
	}
	if bg != nil {
		_, _ = io.WriteString(out, "48;2;")
		write8bitColor(out, bg)
	}
	_, _ = io.WriteString(out, "m")
	_, _ = io.WriteString(out, text)
	_, _ = io.WriteString(out, "\x1b[0m")
}

func write8bitColor(out io.Writer, c color.Color) {
	r, g, b, _ := c.RGBA()
	_, _ = io.WriteString(out, strconv.Itoa(int(r>>8)))
	_, _ = io.WriteString(out, ";")
	_, _ = io.WriteString(out, strconv.Itoa(int(g>>8)))
	_, _ = io.WriteString(out, ";")
	_, _ = io.WriteString(out, strconv.Itoa(int(b>>8)))
}

func DisableLineWrapping(out io.Writer) (restore func()) {
	return applyPair(out, "\x1b[?7l", "\x1b[?7h")
}

func applyPair(out io.Writer, action, revert string) (restore func()) {
	_, err := io.WriteString(out, action)
	if err == nil {
		return func() {
			_, _ = io.WriteString(out, revert)
		}
	}
	return noop
}

var noop = func() {}
