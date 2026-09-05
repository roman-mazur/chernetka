package extsyntaxhl

import (
	"slices"
	"strings"

	"rmazur.io/chernetka/internal/editor"
)

// source is the snapshot of a document that highlighters work on.
// Its lines are substrings of text, so splitting it costs no per-line copy.
type source struct {
	text  string
	lines []string
}

func newSource(text string) *source {
	// Buffer.Text joins the content lines with "\n", so splitting on it gives
	// back exactly the lines the editor renders and numbers.
	return &source{text: text, lines: strings.Split(text, "\n")}
}

func (s *source) line(n int) string {
	if n < 0 || n >= len(s.lines) {
		return ""
	}
	return s.lines[n]
}

// rawSpan is a region of the document reported by a highlighter. Unlike the
// editor.SyntaxSpan it ends up as, it may cover several lines, may be nested in
// or overlap another span, and may reach past the end of a line.
type rawSpan struct {
	StartLine, StartCol int
	EndLine, EndCol     int
	TokenType           editor.TokenType
}

func lineSpan(ln, start, end int, t editor.TokenType) rawSpan {
	return rawSpan{StartLine: ln, StartCol: start, EndLine: ln, EndCol: end, TokenType: t}
}

// piece is a rawSpan cut down to a single line and clipped to that line.
type piece struct {
	line       int
	start, end int
	token      editor.TokenType
	seq        int // emission order, used to break ties between equal ranges
}

// spanBuilder flattens the raw spans a highlighter reports into the per-line,
// sorted, non-overlapping spans that the editor renders.
//
// Spans are painted onto a per-line canvas of token types, widest first, so a
// nested span overrides the one enclosing it: an escape sequence wins over the
// string literal holding it, and bold text wins over the heading around it. For
// two spans covering the exact same range the one reported first wins, which is
// how a highlight query puts its specific patterns ahead of its general ones.
// The painted canvas is then run length encoded, which drops the unhighlighted
// gaps and merges neighbours that ended up with the same token type.
type spanBuilder struct {
	pieces []piece
	canvas []editor.TokenType
}

func (sb *spanBuilder) reset() {
	sb.pieces = sb.pieces[:0]
}

// add cuts a raw span into per-line pieces, dropping whatever falls outside the
// document or outside its line.
func (sb *spanBuilder) add(src *source, rs rawSpan) {
	if rs.TokenType == editor.TtNothing || rs.EndLine < rs.StartLine {
		return
	}
	for ln := max(rs.StartLine, 0); ln <= min(rs.EndLine, len(src.lines)-1); ln++ {
		start, end := 0, len(src.line(ln))
		if ln == rs.StartLine {
			start = rs.StartCol
		}
		if ln == rs.EndLine {
			end = rs.EndCol
		}
		sb.addPiece(src, ln, start, end, rs.TokenType)
	}
}

func (sb *spanBuilder) addPiece(src *source, ln, start, end int, t editor.TokenType) {
	lineLen := len(src.line(ln))
	start, end = max(start, 0), min(end, lineLen)
	if start >= end {
		return
	}
	sb.pieces = append(sb.pieces, piece{
		line: ln, start: start, end: end, token: t, seq: len(sb.pieces),
	})
}

// build paints every piece and returns the resulting spans, sorted by line and
// then by start offset. dst is reused to hold the result.
func (sb *spanBuilder) build(dst []editor.SyntaxSpan) []editor.SyntaxSpan {
	// Widest first within a line, and for an identical range the piece reported
	// first is painted last so that it wins.
	slices.SortFunc(sb.pieces, func(a, b piece) int {
		switch {
		case a.line != b.line:
			return a.line - b.line
		case a.start != b.start:
			return a.start - b.start
		case a.end != b.end:
			return b.end - a.end
		default:
			return b.seq - a.seq
		}
	})

	out := dst[:0]
	for i := 0; i < len(sb.pieces); {
		ln := sb.pieces[i].line
		j := i
		width := 0
		for j < len(sb.pieces) && sb.pieces[j].line == ln {
			width = max(width, sb.pieces[j].end)
			j++
		}

		sb.canvas = slices.Grow(sb.canvas[:0], width)[:width]
		clear(sb.canvas)
		for _, p := range sb.pieces[i:j] {
			for k := p.start; k < p.end; k++ {
				sb.canvas[k] = p.token
			}
		}
		out = encodeRuns(out, ln, sb.canvas)

		i = j
	}
	return out
}

// encodeRuns appends one span per run of equal token types on the canvas,
// skipping the runs that need no highlight.
func encodeRuns(dst []editor.SyntaxSpan, ln int, canvas []editor.TokenType) []editor.SyntaxSpan {
	for i := 0; i < len(canvas); {
		t := canvas[i]
		j := i
		for j < len(canvas) && canvas[j] == t {
			j++
		}
		if t != editor.TtNothing {
			dst = append(dst, editor.SyntaxSpan{LineNumber: ln, Start: i, End: j, TokenType: t})
		}
		i = j
	}
	return dst
}

// spansOfLine returns the range of spans belonging to line ln. spans must be
// sorted by line number, as produced by spanBuilder.build.
func spansOfLine(spans []editor.SyntaxSpan, ln int) []editor.SyntaxSpan {
	// BinarySearchFunc returns the smallest matching index, so i starts the run.
	i, _ := slices.BinarySearchFunc(spans, ln, func(s editor.SyntaxSpan, ln int) int {
		return s.LineNumber - ln
	})
	j := i
	for j < len(spans) && spans[j].LineNumber == ln {
		j++
	}
	return spans[i:j]
}

// clipSpans keeps spans within the bounds of the line about to be rendered. The
// cached spans describe the document as of the last parse, which can lag behind
// what the editor is rendering, so without this the renderer could be handed an
// offset past the end of the line.
func clipSpans(spans []editor.SyntaxSpan, line string) []editor.SyntaxSpan {
	n := len(line)
	// Spans are sorted and never overlap, so anything out of bounds is a suffix.
	keep := len(spans)
	for keep > 0 && spans[keep-1].End > n {
		keep--
	}
	if keep == len(spans) {
		return spans
	}
	out := slices.Clone(spans[:keep])
	if spans[keep].Start < n {
		trimmed := spans[keep]
		trimmed.End = n
		out = append(out, trimmed)
	}
	return out
}
