package content

// Position indicates a symbol position in the content.
type Position struct {
	Col  int // rune offset in the line
	Line int // 0-indexed line number
}

// Span represents a pair of positions to refer to a piece of content.
type Span struct {
	Start, End Position
}

func (s *Span) Min() Position {
	if s.Start.Line <= s.End.Line {
		return Position{min(s.Start.Col, s.End.Col), s.Start.Line}
	}
	return s.End
}

func (s *Span) Max() Position {
	if s.End.Line >= s.Start.Line {
		return Position{max(s.Start.Col, s.End.Col), s.End.Line}
	}
	return s.Start
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
