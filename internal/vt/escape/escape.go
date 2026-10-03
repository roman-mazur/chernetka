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

type ConfigFunc func(out io.Writer) Restore

func SyncOutput(out io.Writer) Restore {
	return applyPair(out, "\x1b[?2026h", "\x1b[?2026l")
}

func EnableAlternativeBuffer(out io.Writer) Restore {
	return applyPair(out, "\x1b[?1049h", "\x1b[?1049l")
}

func EnableMouse(out io.Writer) Restore {
	return applyPair(out, "\x1b[?1002h\x1b[?1006h", "\x1b[?1002l\x1b[?1006l")
}

func ReverseVideo(out io.Writer) Restore {
	return applyPair(out, "\x1b[7m", "\x1b[0m")
}

func HideCursor(out io.Writer, thinCursor bool) Restore {
	restoreCode := "\x1b[1 q\x1b[?25h"
	if thinCursor {
		restoreCode = "\x1b[5 q\x1b[?25h"
	}
	return applyPair(out, "\x1b[?25l", restoreCode)
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

// StyleTextSet starts the defined TextStyle for the text written after it.
// It writes nothing and returns false for the zero style: StyleTextReset is needed only if it returns true.
func StyleTextSet(out io.Writer, style styles.TextStyle) bool {
	if style == (styles.TextStyle{}) {
		return false
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
	return true
}

// StyleText prints the provided text with the defined TextStyle.
func StyleText(out io.Writer, text string, style styles.TextStyle) {
	styleSet := StyleTextSet(out, style)
	_, _ = io.WriteString(out, text)
	if styleSet {
		StyleTextReset(out)
	}
}

//go:generate go run ./mklookup zlookup.go

func write8bitColor(out io.Writer, c color.Color) {
	r, g, b, _ := c.RGBA()
	_, _ = io.WriteString(out, numbersLookup[r>>8])
	_, _ = io.WriteString(out, ";")
	_, _ = io.WriteString(out, numbersLookup[g>>8])
	_, _ = io.WriteString(out, ";")
	_, _ = io.WriteString(out, numbersLookup[b>>8])
}

// StyleTextReset ends the style started with StyleTextSet.
func StyleTextReset(out io.Writer) {
	_, _ = io.WriteString(out, "\x1b[0m")
}

func DisableLineWrapping(out io.Writer) Restore {
	return applyPair(out, "\x1b[?7l", "\x1b[?7h")
}

func MouseShape(out io.Writer, shape string) {
	_, _ = io.WriteString(out, "\x1b]22;")
	_, _ = io.WriteString(out, shape)
	_, _ = io.WriteString(out, "\x07")
}

// EnableFocusReporting asks the terminal to send focus in/out events.
func EnableFocusReporting(out io.Writer) Restore {
	return applyPair(out, "\x1b[?1004h", "\x1b[?1004l")
}

func EnableBracketedPasteMode(out io.Writer) Restore {
	return applyPair(out, "\x1b[?2004h", "\x1b[?2004l")
}

func applyPair(out io.Writer, action, revert string) Restore {
	_, err := io.WriteString(out, action)
	if err == nil {
		return Restore{out: out, value: revert}
	}
	return Restore{}
}

// Restore undoes a terminal setting applied by one of the functions in this package.
// The zero value does nothing.
type Restore struct {
	out   io.Writer
	value string
}

// Undo writes the sequence that reverts the setting.
func (r Restore) Undo() {
	if r.out == nil {
		return
	}
	_, _ = io.WriteString(r.out, r.value)
}

var (
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
