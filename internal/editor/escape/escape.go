// Package escape exposes functions to write escape sequences supported by the terminal apps.
package escape

import (
	"fmt"
	"image/color"
	"io"
	"regexp"
	"strconv"

	"rmazur.io/chernetka/internal/editor/styles"
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

// StyleText prints the provided text with the defined TextStyle.
func StyleText(out io.Writer, text string, style styles.TextStyle) {
	if style == (styles.TextStyle{}) {
		_, _ = io.WriteString(out, text)
		return
	}

	styleSet := false
	writeStyle := func(code string) {
		if styleSet {
			_, _ = io.WriteString(out, ";")
		}
		_, _ = io.WriteString(out, code)
		styleSet = true
	}

	_, _ = io.WriteString(out, "\x1b[")

	if style.Bold {
		writeStyle("1")
	}
	if style.Italic {
		writeStyle("3")
	}
	if style.TextColor != nil {
		writeStyle("38;2;")
		write8bitColor(out, style.TextColor)
	}
	if style.BgColor != nil {
		writeStyle("48;2;")
		write8bitColor(out, style.BgColor)
	}

	_, _ = io.WriteString(out, "m")
	_, _ = io.WriteString(out, text)
	_, _ = io.WriteString(out, "\x1b[0m")
}

func write8bitColor(out io.Writer, c color.Color) {
	var buf [11]byte // 3 numbers of 3 digits max + 2*;
	r, g, b, _ := c.RGBA()
	tail := strconv.AppendInt(buf[0:0], int64(r>>8), 10)
	tail = append(tail, ';')
	tail = strconv.AppendInt(tail, int64(g>>8), 10)
	tail = append(tail, ';')
	tail = strconv.AppendInt(tail, int64(b>>8), 10)
	_, _ = out.Write(tail)
}

func DisableLineWrapping(out io.Writer) (restore func()) {
	return applyPair(out, "\x1b[?7l", "\x1b[?7h")
}

func MouseShape(out io.Writer, shape string) {
	_, _ = io.WriteString(out, "\x1b]22;")
	_, _ = io.WriteString(out, shape)
	_, _ = io.WriteString(out, "\x07")
}

func EnableBracketedPasteMode(out io.Writer) (restore func()) {
	return applyPair(out, "\x1b[?2004h", "\x1b[?2004l")
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

var (
	noop = func() {}

	ansiEscRE = regexp.MustCompile(`\x1b\[[0-9;?]*[A-Za-z]`)
)

// Clean removes escape symbols from the input string.
func Clean(s string) string {
	return ansiEscRE.ReplaceAllString(s, "")
}

// SetTermTitle uses escape sequence to instruct the terminal what tab/window title should be.
func SetTermTitle(out io.Writer, title string) {
	_, _ = io.WriteString(out, "\x1b]0;")
	_, _ = io.WriteString(out, title)
	_, _ = io.WriteString(out, "\x07")
}

type ProgressState int

const (
	ProgressStateHidden ProgressState = iota
	ProgressStateDefault
	ProgressStateError
	ProgressStateIndeterminate
	ProgressStateWarning
)

func UpdateProgress(out io.Writer, state ProgressState, progress int) {
	p := max(0, min(progress, 100))
	var data [8]byte
	_, _ = io.WriteString(out, "\x1b]9;4;")
	stateData := strconv.AppendInt(data[:], int64(state), 10)
	stateData = append(stateData, ';')
	stateData = strconv.AppendInt(stateData, int64(p), 10)
	_, _ = out.Write(stateData[:])
	_, _ = io.WriteString(out, "\x07")
}

// ClearScreen erases the screen content and moves the cursor to the top left corner.
func ClearScreen(out io.Writer) {
	_, _ = io.WriteString(out, "\x1b[2J\x1b[H")
}
