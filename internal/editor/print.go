package editor

import (
	"bufio"
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"rmazur.io/x/edit/internal/editor/escape"
)

// printFmt renders lines [i, j) with line numbers and tab expansion.
// When suggestion is non-empty it is drawn as dim text at byte offset cx of
// the cursor line cy.
func (b *Buffer) printFmt(out *bufio.Writer, i, j int, prefs *RenderPrefs) {
	lines := b.content.Lines()
	_ = lines[i:j] // boundary check

	tab := strings.Repeat(" ", prefs.TabSize)
	renderText := func(s string) string { return strings.ReplaceAll(s, "\t", tab) }

	nl := nlDigitsLen(j)
	nlFmt := "%" + strconv.Itoa(nl) + "d "

	for x := i; x < j; x++ {
		raw := lines[x].String()
		escape.ClearLine(out)

		nlColor := colors.Suggestion
		if x == b.cy {
			nlColor = colors.Selected
		}
		escape.ColorText(out, fmt.Sprintf(nlFmt, x+1), nlColor)

		if x == b.cy && len(b.lsp.suggestions) > 0 && b.cx <= len(raw) {
			out.WriteString(renderText(raw[:b.cx]))
			sugText := b.lsp.suggestions[0] // TODO: maintain suggestions index
			escape.ColorText(out, sugText, colors.Suggestion)
			out.WriteString(renderText(raw[b.cx:]))
		} else {
			out.WriteString(renderText(raw))
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
}{
	Suggestion: color.Gray{Y: 100},
	Selected:   color.Gray{Y: 200},
}
