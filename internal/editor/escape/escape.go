// Package escape exposes functions to write escape sequences supported by the terminal apps.
package escape

import (
	"fmt"
	"image/color"
	"io"
	"strings"
)

func EnableAlternativeBuffer(out io.Writer) (restore func()) {
	return applyPair(out, "\x1b[?1049h", "\x1b[?1049l")
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
	_, _ = fmt.Fprint(out, "\x1b[H")
}

func SetCursorPosition(out io.Writer, row, col int) {
	_, _ = fmt.Fprintf(out, "\x1b[%d;%dH", row, col)
}

func ClearLine(out io.Writer) {
	_, _ = fmt.Fprint(out, "\x1b[2K")
}

func ColorText(out io.Writer, text string, fg, bg color.Color) {
	_, _ = fmt.Fprintf(out, colorTemplate(fg, bg), text)
}

func colorTemplate(fg, bg color.Color) string {
	if fg == nil && bg == nil {
		return "%s"
	}

	var res strings.Builder
	res.WriteString("\x1b[")
	if fg != nil {
		r, g, b, _ := fg.RGBA()
		_, _ = fmt.Fprintf(&res, "38;2;%d;%d;%d", r, g, b)
		if bg != nil {
			res.WriteString(";")
		}
	}
	if bg != nil {
		r, g, b, _ := bg.RGBA()
		_, _ = fmt.Fprintf(&res, "48;2;%d;%d;%d", r, g, b)
	}

	res.WriteString("m%s\x1b[0m")
	return res.String()
}

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
