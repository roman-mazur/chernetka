package editor

import (
	"bufio"
	"fmt"
	"image/color"
	"io"
	"strconv"
	"strings"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/escape"
)

type contentPrinter struct {
	CodeAssist
	SyntaxHighlighter

	b *Buffer

	lines []content.Line
	i, j  int // lines range

	suggestion string

	lnFmt string // format to render line numbers
	tab   string // what to render for \t
}

func (cr *contentPrinter) prepare(b *Buffer, i, j int, prefs *RenderPrefs) {
	cr.b = b
	cr.lines = b.Content.Lines()
	_ = cr.lines[i:j] // boundary check
	cr.i, cr.j = i, j
	cr.tab = strings.Repeat(" ", prefs.TabSize)
	nl := nlDigitsLen(j)
	cr.lnFmt = "%" + strconv.Itoa(nl) + "d "

	findExtData(b, &cr.CodeAssist)
	findExtData(b, &cr.SyntaxHighlighter)

	if cr.CodeAssist != nil {
		cr.suggestion = cr.TextSuggestion()
	}
}

func (cr *contentPrinter) render(out *bufio.Writer) {
	for ln := cr.i; ln < cr.j; ln++ {
		raw := cr.lines[ln].String()
		escape.ClearLine(out)

		if !cr.b.hideLineNumbers {
			nlColor := colors.Suggestion
			if ln == cr.b.cy {
				nlColor = colors.Selected
			}
			escape.ColorText(out, fmt.Sprintf(cr.lnFmt, ln+1), nlColor)
		}

		if ln == cr.b.cy && !cr.b.noCurrentLineHL {
			var lineOut strings.Builder
			cr.renderLine(&lineOut, ln, raw)
			rightPad := cr.b.w - byteToScreenCol(lineOut.String(), lineOut.Len(), len(cr.tab))
			if rightPad > 0 {
				lineOut.WriteString(strings.Repeat(" ", rightPad))
			}
			escape.ColorBackground(out, lineOut.String(), colors.SelectedBg)
		} else {
			cr.renderLine(out, ln, raw)
		}

		out.WriteString("\r\n")
	}

}

func (cr *contentPrinter) normalizeText(s string) string {
	return strings.ReplaceAll(s, "\t", cr.tab)
}

func (cr *contentPrinter) renderLine(out io.Writer, ln int, line string) {
	if ln == cr.b.cy && cr.suggestion != "" && cr.b.cx <= len(line) {
		_, _ = fmt.Fprint(out, cr.normalizeText(line[:cr.b.cx]))
		escape.ColorText(out, cr.suggestion, colors.Suggestion)
		_, _ = fmt.Fprint(out, cr.normalizeText(line[cr.b.cx:]))
		// TODO: use syntax HL
	} else {
		if cr.SyntaxHighlighter != nil {
			for _, span := range cr.SyntaxSpans(ln) {
				text := cr.normalizeText(line[span.Start:span.End])
				escape.ColorText(out, text, colors.ColorForTokenType(span.TokenType))
			}
		} else {
			_, _ = fmt.Fprint(out, cr.normalizeText(line))
		}
	}
}

func findExtData[T BufferExtData](b *Buffer, out *T) {
	if b.xData == nil {
		return
	}
	for _, data := range b.xData {
		if res, ok := data.(T); ok {
			*out = res
			return
		}
	}
	return
}

func nlDigitsLen(x int) int {
	l := 0
	for x > 0 {
		x /= 10
		l++
	}
	return l
}

type ColorTheme struct {
	Suggestion color.Color
	Selected   color.Color
	SelectedBg color.Color

	syntaxColors map[TokenType]color.Color
}

func (ct *ColorTheme) ColorForTokenType(t TokenType) color.Color {
	if ct.syntaxColors == nil {
		return nil
	}
	return ct.syntaxColors[t]
}

var colors = ColorTheme{
	Suggestion: color.Gray{Y: 100},
	Selected:   color.Gray{Y: 200},
	SelectedBg: color.Gray{Y: 70},

	syntaxColors: map[TokenType]color.Color{
		TokenTypeCall: color.RGBA{R: 0xaa, G: 0xaa, B: 0, A: 0xff},
	},
}
