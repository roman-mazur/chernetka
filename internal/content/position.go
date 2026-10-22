package content

import (
	"fmt"
	"slices"
)

// Position indicates a symbol position in the content.
type Position struct {
	Col  int // byte offset in the line
	Line int // 0-indexed line number
}

func (p Position) String() string {
	return fmt.Sprintf("l=%d/c=%d", p.Line, p.Col)
}

// Compare returns -1 if p is before q in the content, 1 if it's after, and 0 if they're equal.
func (p Position) Compare(q Position) int {
	if p.Line < q.Line {
		return -1
	}
	if p.Line > q.Line {
		return 1
	}
	if p.Col == q.Col {
		return 0
	}
	if p.Col < q.Col {
		return -1
	}
	return 1
}

// Span represents a pair of positions to refer to a piece of content.
type Span struct {
	Start, End Position
}

func (s *Span) Min() Position {
	if s.Start.Line != s.End.Line {
		if s.Start.Line < s.End.Line {
			return s.Start
		}
		return s.End
	}
	return Position{min(s.Start.Col, s.End.Col), s.Start.Line}
}

func (s *Span) Max() Position {
	if s.Start.Line != s.End.Line {
		if s.Start.Line < s.End.Line {
			return s.End
		}
		return s.Start
	}
	return Position{max(s.Start.Col, s.End.Col), s.Start.Line}
}

func (s *Span) ContainsLine(line int) bool {
	return s.Min().Line <= line && line <= s.Max().Line
}

func (s *Span) LineProjection(idx int, lineLen int) Span {
	res := Span{s.Min(), s.Max()}
	if res.Start.Line < idx {
		res.Start = Position{0, idx}
	}
	if res.End.Line > idx {
		res.End = Position{lineLen, idx}
	}
	return res
}

// NormalizeSpans turns spans into non-empty spans going forward (Start before End) and sorted
// by Start, merging the ones that overlap or touch, so they never share a position but their ends.
// Like slices.Compact, it reuses and modifies spans, and returns the shortened slice.
func NormalizeSpans(spans []Span) []Span {
	res := spans[:0]
	for _, s := range spans {
		if s.Start != s.End {
			res = append(res, Span{Start: s.Min(), End: s.Max()})
		}
	}
	slices.SortFunc(res, func(a, b Span) int { return a.Start.Compare(b.Start) })
	merged := res[:0]
	for _, s := range res {
		if n := len(merged); n > 0 && s.Start.Compare(merged[n-1].End) <= 0 {
			if s.End.Compare(merged[n-1].End) > 0 {
				merged[n-1].End = s.End
			}
			continue
		}
		merged = append(merged, s)
	}
	clear(spans[len(merged):])
	return merged
}
