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

	lnDigits int      // number of digits for line numbers
	lnBuf    [10]byte // buffer for line numbers
	tab      string   // what to render for \t
}

func (cr *contentPrinter) prepare(b *Buffer, i, j int, prefs *RenderPrefs) {
	cr.b = b
	cr.lines = b.Content.Lines()
	_ = cr.lines[i:j] // boundary check
	cr.i, cr.j = i, j
	cr.tab = strings.Repeat(" ", prefs.TabSize)
	cr.lnDigits = nlDigitsLen(j)

	findExtData(b, &cr.CodeAssist)
	findExtData(b, &cr.SyntaxHighlighter)

	if cr.CodeAssist != nil {
		cr.suggestion = cr.TextSuggestion()
	}
}

func (cr *contentPrinter) lineNumber(n int) string {
	if cr.lnDigits > len(cr.lnBuf)-1 {
		return fmt.Sprintf("%"+strconv.Itoa(cr.lnDigits)+"d", n)
	}
	i := 0
	for range cr.lnDigits - nlDigitsLen(n) {
		cr.lnBuf[i] = ' '
		i++
	}
	strconv.AppendInt(cr.lnBuf[i:i], int64(n), 10)
	cr.lnBuf[cr.lnDigits] = ' '
	return string(cr.lnBuf[:cr.lnDigits+1])
}

func (cr *contentPrinter) render(out *bufio.Writer) {
	for ln := cr.i; ln < cr.j; ln++ {
		raw := cr.lines[ln].String()
		escape.ClearLine(out)

		if !cr.b.hideLineNumbers {
			nlColor := colors.Suggestion
			if ln == cr.b.c.y {
				nlColor = colors.Selected
			}
			escape.ColorText(out, cr.lineNumber(ln+1), nlColor, nil)
		}

		lineHL := ln == cr.b.c.y && !cr.b.noCurrentLineHL
		cr.renderLine(out, ln, raw, lineHL)

		if lineHL {
			rightPad := cr.b.w - runeToScreenCol(raw, len(raw), len(cr.tab))
			if rightPad > 0 {
				escape.ColorText(out, strings.Repeat(" ", rightPad), nil, colors.SelectedBg)
			}
		}

		out.WriteString("\r\n")
	}

}

func (cr *contentPrinter) normalizeText(s string) string {
	return strings.ReplaceAll(s, "\t", cr.tab)
}

func (cr *contentPrinter) renderLine(out io.Writer, ln int, line string, hlLine bool) {
	var bgColor color.Color
	if hlLine {
		bgColor = colors.SelectedBg
	}

	if ln == cr.b.c.y && cr.suggestion != "" && cr.b.c.x <= len(line) {
		_, _ = fmt.Fprint(out, cr.normalizeText(line[:cr.b.c.x]))
		escape.ColorText(out, cr.suggestion, colors.Suggestion, bgColor)
		_, _ = fmt.Fprint(out, cr.normalizeText(line[cr.b.c.x:]))
		// TODO: use syntax HL
		return
	}
	if cr.SyntaxHighlighter == nil {
		escape.ColorText(out, cr.normalizeText(line), nil, bgColor)
		return
	}

	lastIndex := 0
	for _, span := range cr.SyntaxSpans(ln, line) {
		if span.Start > lastIndex {
			escape.ColorText(out, cr.normalizeText(line[lastIndex:span.Start]), nil, bgColor)
		}
		if span.End > len(line) {
			panic(fmt.Errorf("line %d %q, span %s out of range", ln, line, span))
		}
		text := cr.normalizeText(line[span.Start:span.End])
		escape.ColorText(out, text, colors.ColorForTokenType(span.TokenType), bgColor)
		lastIndex = span.End
	}
	if lastIndex < len(line) {
		escape.ColorText(out, cr.normalizeText(line[lastIndex:]), nil, bgColor)
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
		TtKeyword:         parseColor("CF8E6D"),
		TtTypeRef:         parseColor("BCBEC4"),
		TtStringLiteral:   parseColor("6AAB73"),
		TtNumberLiteral:   parseColor("2AACB8"),
		TtFuncDeclaration: parseColor("56A8F5"),
		TtComment:         parseColor("7A7E85"),
	},
}

func parseColor(s string) color.Color {
	if len(s) != 6 {
		panic(fmt.Errorf("expected 6 bytes, got %d", len(s)))
	}
	r, g, b := hex2i(s[:2]), hex2i(s[2:4]), hex2i(s[4:])
	return color.RGBA{R: r, G: g, B: b, A: 0xff}
}

func hex2i(s string) uint8 {
	res, err := strconv.ParseInt(s, 16, 32)
	if err != nil {
		panic(err)
	}
	return uint8(res)
}
