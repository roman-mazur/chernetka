package editor

import (
	"bufio"
	"fmt"
	"image/color"
	"io"
	"strconv"
	"strings"

	"rmazur.io/chernetka/internal/editor/escape"
)

// printFmt renders lines [i, j) with line numbers and tab expansion.
// When suggestion is non-empty it is drawn as dim text at byte offset cx of
// the cursor line cy.
func (b *Buffer) printFmt(out *bufio.Writer, i, j int, prefs *RenderPrefs) {
	lines := b.Content.Lines()
	_ = lines[i:j] // boundary check

	tab := strings.Repeat(" ", prefs.TabSize)
	renderText := func(s string) string { return strings.ReplaceAll(s, "\t", tab) }

	nl := nlDigitsLen(j)
	nlFmt := "%" + strconv.Itoa(nl) + "d "

	var suggestion string
	lspExt := b.ExtensionData("lsp")
	if lspExt != nil {
		suggestion = lspExt.TextSuggestion()
	}

	renderLine := func(out io.Writer, x int, line string) {
		if x == b.cy && suggestion != "" && b.cx <= len(line) {
			_, _ = fmt.Fprint(out, renderText(line[:b.cx]))
			escape.ColorText(out, suggestion, colors.Suggestion)
			_, _ = fmt.Fprint(out, renderText(line[b.cx:]))
		} else {
			_, _ = fmt.Fprint(out, renderText(line))
		}
	}

	for x := i; x < j; x++ {
		raw := lines[x].String()
		escape.ClearLine(out)

		if !b.hideLineNumbers {
			nlColor := colors.Suggestion
			if x == b.cy {
				nlColor = colors.Selected
			}
			escape.ColorText(out, fmt.Sprintf(nlFmt, x+1), nlColor)
		}

		if x == b.cy && b.highlightCurrentLine {
			var lineOut strings.Builder
			renderLine(&lineOut, x, raw)
			rightPad := b.w - byteToScreenCol(lineOut.String(), lineOut.Len(), prefs.TabSize)
			lineOut.WriteString(strings.Repeat(" ", rightPad))
			escape.ColorBackground(out, lineOut.String(), colors.SelectedBg)
		} else {
			renderLine(out, x, raw)
		}

		out.WriteString("\r\n")
	}
}

func nlDigitsLen(x int) int {
	l := 0
	for x > 0 {
		x /= 10
		l++
	}
	return l
}

var colors = struct {
	Suggestion color.Color
	Selected   color.Color
	SelectedBg color.Color
}{
	Suggestion: color.Gray{Y: 100},
	Selected:   color.Gray{Y: 200},
	SelectedBg: color.Gray{Y: 150},
}
