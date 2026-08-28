package editor

import (
	"fmt"
	"image/color"
	"io"
	"slices"
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

func (cr *contentPrinter) render(out io.Writer) {
	for ln := cr.i; ln < cr.j; ln++ {
		raw := cr.lines[ln].String()
		escape.ClearLine(out)

		if !cr.b.hideLineNumbers {
			nlColor := colors.Suggestion
			if ln == cr.b.c.Line {
				nlColor = colors.LineSelected
			}
			escape.ColorText(out, cr.lineNumber(ln+1), nlColor, nil)
		}

		lineHL := ln == cr.b.c.Line && !cr.b.noCurrentLineHL
		cr.renderLine(out, ln, raw, lineHL)

		if lineHL {
			rightPad := cr.b.w - runeToScreenCol(raw, len(raw), len(cr.tab))
			if rightPad > 0 {
				escape.ColorText(out, strings.Repeat(" ", rightPad), nil, colors.LineSelectedBg)
			}
		}

		printLineEnding(out)
	}

}

func (cr *contentPrinter) renderLine(out io.Writer, ln int, line string, hlLine bool) {
	p := colorLinePrinter{
		out:  out,
		tab:  cr.tab,
		line: line,
		ln:   ln,
		bg:   cr.buildBgSpans(ln, len(line), hlLine),
	}

	if ln == cr.b.c.Line && cr.suggestion != "" && cr.b.c.Col <= len(line) {
		p.print(line[:cr.b.c.Col], nil)
		p.printSuggestion(cr.suggestion, colors.Suggestion)
		p.print(line[cr.b.c.Col:], nil)
		// TODO: use syntax HL
		return
	}

	if cr.SyntaxHighlighter == nil {
		p.print(line, nil)
		return
	}

	lastIndex := 0
	for _, span := range cr.SyntaxSpans(ln, line) {
		if span.Start > lastIndex {
			p.print(line[lastIndex:span.Start], nil)
		}
		if span.End > len(line) {
			panic(fmt.Errorf("line %d %q, span %s out of range", ln, line, span))
		}
		p.print(line[span.Start:span.End], colors.ColorForTokenType(span.TokenType))
		lastIndex = span.End
	}
	if lastIndex < len(line) {
		p.print(line[lastIndex:], nil)
	}
}

func (cr *contentPrinter) buildBgSpans(line int, lineLen int, hlLine bool) []colorSpan {
	spans := cr.b.selectionsOnLine(line)
	if len(spans) == 0 {
		if !hlLine {
			return nil
		}
		cs := colorSpan{color: colors.LineSelectedBg,
			Start: content.Position{Col: 0, Line: line},
			End:   content.Position{Col: lineLen, Line: line}}
		return []colorSpan{cs}
	}

	defBg := func() color.Color {
		if hlLine {
			return colors.LineSelectedBg
		}
		return nil
	}

	res := make([]colorSpan, 0, len(spans)+2)

	if spans[0].Start.Col != 0 {
		res = append(res, colorSpan{
			Start: content.Position{0, line},
			End:   content.Position{spans[0].Start.Col, line},
			color: defBg(),
		})
	}

	for i, selSpan := range spans {
		res = append(res, colorSpan{
			Span:  selSpan,
			color: colors.TextSelectedBg,
		})
		if i < len(spans)-1 && selSpan.End.Col != spans[i+1].Start.Col {
			res = append(res, colorSpan{
				Start: content.Position{selSpan.End.Col, line},
				End:   content.Position{spans[i+1].Start.Col, line},
				color: defBg(),
			})
		}
	}

	if spans[len(spans)-1].End.Col != lineLen {
		res = append(res, colorSpan{
			Start: content.Position{spans[len(spans)-1].End.Col, line},
			End:   content.Position{lineLen, line},
			color: defBg(),
		})
	}

	return res
}

type colorSpan struct {
	content.Span
	color color.Color
}

type colorLinePrinter struct {
	out  io.Writer
	tab  string
	line string
	ln   int
	bg   []colorSpan

	li int // what part of the line text has been printed
}

func (clp *colorLinePrinter) printedText(s string) string {
	return strings.ReplaceAll(s, "\t", clp.tab)
}

func (clp *colorLinePrinter) currentBgIdx() int {
	return slices.IndexFunc(clp.bg, func(cs colorSpan) bool {
		return cs.Start.Col <= clp.li && cs.End.Col >= clp.li
	})
}

func (clp *colorLinePrinter) print(s string, fg color.Color) {
	bgIdx := clp.currentBgIdx()
	if bgIdx == -1 {
		escape.ColorText(clp.out, clp.printedText(s), fg, nil)
		return
	}
	for start := 0; start < len(s); {
		bg := clp.bg[bgIdx]
		end := min(len(s), bg.End.Col-clp.li)
		escape.ColorText(clp.out, clp.printedText(s[start:end]), fg, bg.color)
		start = end
		if clp.li+start >= bg.End.Col {
			bgIdx++
		}
	}
	clp.li += len(s)
}

func (clp *colorLinePrinter) printSuggestion(txt string, fg color.Color) {
	var bgColor color.Color
	if bgIdx := clp.currentBgIdx(); bgIdx != -1 {
		bgColor = clp.bg[bgIdx].color
	}
	escape.ColorText(clp.out, clp.printedText(txt), fg, bgColor)
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
	Suggestion     color.Color
	LineSelected   color.Color
	LineSelectedBg color.Color
	TextSelected   color.Color
	TextSelectedBg color.Color

	syntaxColors map[TokenType]color.Color
}

func (ct *ColorTheme) ColorForTokenType(t TokenType) color.Color {
	if ct.syntaxColors == nil {
		return nil
	}
	return ct.syntaxColors[t]
}

var colors = ColorTheme{
	Suggestion:     color.Gray{Y: 100},
	LineSelected:   color.Gray{Y: 200},
	LineSelectedBg: color.Gray{Y: 70},
	TextSelected:   color.Gray{Y: 200},
	TextSelectedBg: parseColor("1010FF"),

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

func printLineEnding(out io.Writer) {
	_, _ = io.WriteString(out, "\r\n")
}
