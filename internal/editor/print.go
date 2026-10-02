package editor

import (
	"fmt"
	"image/color"
	"io"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/content/code"
	"rmazur.io/chernetka/internal/editor/styles"
	"rmazur.io/chernetka/internal/vt/escape"
)

type contentPrinter struct {
	CodeAssist
	SyntaxHighlighter

	b *Buffer

	lines []content.Line
	i, j  int // lines range

	suggestion code.Suggestion
	diags      map[int]code.Diagnostic // the most severe problem of the lines that have them

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

	cr.CodeAssist, _ = FindExtData[CodeAssist](b)
	cr.SyntaxHighlighter, _ = FindExtData[SyntaxHighlighter](b)

	if cr.CodeAssist != nil && b.mode == ModeInsert {
		cr.suggestion = cr.TextSuggestion()
	}
	if dp, ok := FindExtData[DiagnosticsProvider](b); ok {
		cr.diags = worstDiagnostics(dp.Diagnostics(), i, j)
	}
}

// worstDiagnostics returns the most severe of diags for every line in [i, j).
func worstDiagnostics(diags []code.Diagnostic, i, j int) map[int]code.Diagnostic {
	var res map[int]code.Diagnostic
	for _, d := range diags {
		if d.Line < i || d.Line >= j {
			continue
		}
		if prev, ok := res[d.Line]; ok && prev.Severity <= d.Severity {
			continue
		}
		if res == nil {
			res = make(map[int]code.Diagnostic)
		}
		res[d.Line] = d
	}
	return res
}

func severityColor(s code.Severity) color.Color {
	if s == code.SeverityError {
		return styles.DefaultColors.Error
	}
	return styles.DefaultColors.Warning
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
			style := styles.TextStyle{TextColor: styles.DefaultColors.Suggestion}
			if ln == cr.b.c().Line {
				style.TextColor = styles.DefaultColors.LineSelected
			}
			if d, ok := cr.diags[ln]; ok {
				style.TextColor = severityColor(d.Severity)
			}
			escape.StyleText(out, cr.lineNumber(ln+1), style)
		}

		lineHL := ln == cr.b.c().Line && !cr.b.noCurrentLineHL
		cr.renderLine(out, ln, raw, lineHL)
		cr.renderLineTail(out, ln, raw, lineHL)

		if !cr.b.hideLineActions && cr.b.lineAction(ln) != nil {
			cr.renderActionMarker(out, lineHL, ln)
		}

		printLineEnding(out)
	}

}

// renderActionMarker shows that the line can be engaged placing a marker in the last column.
// The marker overwrites the line content if it's too long to fit the window.
func (cr *contentPrinter) renderActionMarker(out io.Writer, lineHL bool, ln int) {
	style := styles.TextStyle{TextColor: styles.DefaultColors.LineAction}
	if lineHL {
		style.BgColor = styles.DefaultColors.LineSelectedBg
	}
	screenRow := ln - cr.i + 1 // lines are rendered starting from the buffer offset
	escape.SetCursorPosition(out, screenRow, cr.b.w)
	escape.StyleText(out, actionMarker, style)
}

const actionMarker = "▶"

func (cr *contentPrinter) renderLine(out io.Writer, ln int, line string, hlLine bool) {
	p := colorLinePrinter{
		out:  out,
		tab:  cr.tab,
		line: line,
		ln:   ln,
		bg:   cr.buildBgSpans(ln, len(line), hlLine),
		from: cr.b.xoff,
		to:   cr.b.xoff + cr.b.textWidth(),
	}

	// The suggestion is shown as ghost text at the cursor, in the middle of
	// the highlighted segments.
	ghostAt := -1
	if cr.hasGhost(ln, line) {
		ghostAt = cr.b.c().Col
	}
	printSegment := func(from, to int, style styles.TextStyle) {
		if from <= ghostAt && ghostAt < to {
			if from < ghostAt {
				p.print(line[from:ghostAt], style)
			}
			p.printSuggestion(cr.suggestion.Text, styles.DefaultColors.Suggestion)
			from, ghostAt = ghostAt, -1
		}
		p.print(line[from:to], style)
	}

	lastIndex := 0
	if cr.SyntaxHighlighter != nil {
		for _, span := range cr.SyntaxSpans(ln, line) {
			if span.Start > lastIndex {
				printSegment(lastIndex, span.Start, styles.TextStyle{})
			}
			if span.End > len(line) {
				panic(fmt.Errorf("line %d %q, span %s out of range", ln, line, span))
			}
			printSegment(span.Start, span.End, styles.ResolveTokenStyle(span.TokenType))
			lastIndex = span.End
		}
	}
	if lastIndex < len(line) {
		printSegment(lastIndex, len(line), styles.TextStyle{})
	}
	if ghostAt != -1 {
		p.printSuggestion(cr.suggestion.Text, styles.DefaultColors.Suggestion)
	}
}

func (cr *contentPrinter) hasGhost(ln int, line string) bool {
	return ln == cr.b.c().Line && cr.suggestion.Text != "" && cr.b.c().Col <= len(line)
}

// renderLineTail fills the rest of the line after its content: the current
// line background, and the info aligned to the right: about the suggestion, or
// the problem found on the line.
func (cr *contentPrinter) renderLineTail(out io.Writer, ln int, raw string, lineHL bool) {
	used := runeToScreenCol(raw, len(raw), len(cr.tab))
	var (
		info      string
		infoColor = styles.DefaultColors.Suggestion
		shorten   bool // the info may be cut to fit
	)
	if cr.hasGhost(ln, raw) {
		used += utf8.RuneCountInString(cr.suggestion.Text)
		info = cr.suggestion.Info
	} else if d, ok := cr.diags[ln]; ok && ln == cr.b.c().Line {
		info, _, _ = strings.Cut(d.Message, "\n")
		infoColor = severityColor(d.Severity)
		shorten = true
	}
	used = max(0, used-cr.b.xoff) // only the visible part takes the space
	free := cr.b.w - used
	if !cr.b.hideLineNumbers {
		free -= cr.lnDigits + 1
	}

	var bg styles.TextStyle
	if lineHL {
		bg.BgColor = styles.DefaultColors.LineSelectedBg
	}
	const gap, margin = 4, 2 // keep the info off the text and the action marker
	if shorten {
		info = shortenText(info, free-gap-margin)
	}
	if infoLen := utf8.RuneCountInString(info); info != "" && free >= infoLen+gap+margin {
		escape.StyleText(out, strings.Repeat(" ", free-infoLen-margin), bg)
		infoStyle := bg
		infoStyle.TextColor = infoColor
		escape.StyleText(out, info, infoStyle)
		free = margin
	}
	if lineHL && free > 0 {
		escape.StyleText(out, strings.Repeat(" ", free), bg)
	}
}

func (cr *contentPrinter) buildBgSpans(line int, lineLen int, hlLine bool) []colorSpan {
	var spans []colorSpan
	sel := cr.b.selectionsOnLine(line)
	for _, span := range sel {
		spans = append(spans, colorSpan{Span: span, color: styles.DefaultColors.TextSelectedBg})
	}
	// The selection takes priority over the search matches.
	for _, match := range cr.b.searchMatchesOnLine(line, cr.lines[line].String()) {
		overlaps := slices.ContainsFunc(sel, func(s content.Span) bool {
			return s.Start.Col < match.End.Col && match.Start.Col < s.End.Col
		})
		if overlaps {
			continue
		}
		bg := styles.DefaultColors.SearchMatchBg
		if match.Start == cr.b.c() {
			bg = styles.DefaultColors.SearchCursorBg
		}
		spans = append(spans, colorSpan{Span: match, color: bg})
	}

	if len(spans) == 0 {
		if !hlLine {
			return nil
		}
		cs := colorSpan{color: styles.DefaultColors.LineSelectedBg,
			Start: content.Position{Col: 0, Line: line},
			End:   content.Position{Col: lineLen, Line: line}}
		return []colorSpan{cs}
	}
	slices.SortFunc(spans, func(a, b colorSpan) int { return a.Start.Col - b.Start.Col })

	var defBg color.Color
	if hlLine {
		defBg = styles.DefaultColors.LineSelectedBg
	}
	gap := func(from, to int) colorSpan {
		return colorSpan{
			Start: content.Position{Col: from, Line: line},
			End:   content.Position{Col: to, Line: line},
			color: defBg,
		}
	}

	res := make([]colorSpan, 0, 2*len(spans)+1)
	last := 0
	for _, span := range spans {
		if span.Start.Col != last {
			res = append(res, gap(last, span.Start.Col))
		}
		res = append(res, span)
		last = span.End.Col
	}
	if last != lineLen {
		res = append(res, gap(last, lineLen))
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

	x        int // screen column of the printed text
	from, to int // visible screen columns range, the text outside of it is scrolled away
}

func (clp *colorLinePrinter) currentBgIdx() int {
	return slices.IndexFunc(clp.bg, func(cs colorSpan) bool {
		return cs.Start.Col <= clp.li && cs.End.Col >= clp.li
	})
}

func (clp *colorLinePrinter) print(s string, style styles.TextStyle) {
	appliedStyle := style
	bgIdx := clp.currentBgIdx()
	if bgIdx == -1 {
		appliedStyle.BgColor = nil
		clp.emitTabSplit(s, appliedStyle)
		return
	}
	for start := 0; start < len(s); {
		bg := clp.bg[bgIdx]
		end := min(len(s), bg.End.Col-clp.li)
		appliedStyle.BgColor = bg.color
		clp.emitTabSplit(s[start:end], appliedStyle)
		start = end
		if clp.li+start >= bg.End.Col {
			bgIdx++
		}
	}
	clp.li += len(s)
}

func (clp *colorLinePrinter) printSuggestion(txt string, fg color.Color) {
	style := styles.TextStyle{TextColor: fg}
	if bgIdx := clp.currentBgIdx(); bgIdx != -1 {
		style.BgColor = clp.bg[bgIdx].color
	}

	clp.emitTabSplit(txt, style)
}

func (clp *colorLinePrinter) emitTabSplit(s string, style styles.TextStyle) {
	i := 0
	for s := range strings.SplitSeq(s, "\t") {
		if i > 0 {
			clp.emit(clp.tab, style)
		}
		i++
		clp.emit(s, style)
	}
}

// emit writes the part of the printed text that falls into the visible columns range.
func (clp *colorLinePrinter) emit(s string, style styles.TextStyle) {
	x := clp.x
	clp.x += utf8.RuneCountInString(s)
	if clp.x <= clp.from || x >= clp.to {
		return
	}
	for ; x < clp.from; x++ {
		_, sz := utf8.DecodeRuneInString(s)
		s = s[sz:]
	}
	if clp.x > clp.to {
		end := 0
		for ; x < clp.to; x++ {
			_, sz := utf8.DecodeRuneInString(s[end:])
			end += sz
		}
		s = s[:end]
	}
	escape.StyleText(clp.out, s, style)
}

// minShortened is the width text is not shortened below: less isn't readable.
const minShortened = 12

// shortenText cuts s to be no wider than w, marking the cut with an ellipsis.
// s is returned as is if it fits or w is too small to show anything useful.
func shortenText(s string, w int) string {
	if w < minShortened || utf8.RuneCountInString(s) <= w {
		return s
	}
	i := 0
	for range w - 1 {
		_, sz := utf8.DecodeRuneInString(s[i:])
		i += sz
	}
	return s[:i] + "…"
}

func nlDigitsLen(x int) int {
	l := 0
	for x > 0 {
		x /= 10
		l++
	}
	return l
}

func printLineEnding(out io.Writer) {
	_, _ = io.WriteString(out, "\r\n")
}
