package content

import "strings"

// DeleteSpan removes the text delimited by the span from the mutable content
// and returns the position where the deleted text used to start.
func DeleteSpan(m Mutable, span Span) Position {
	start, stop := span.Min(), span.Max()
	lines := m.Lines()

	first := lines[start.Line].String()
	if start.Line == stop.Line {
		m.Update(start.Line, TextLine(first[:start.Col]+first[stop.Col:]))
		return start
	}

	last := lines[stop.Line].String()
	m.Update(start.Line, TextLine(first[:start.Col]+last[stop.Col:]))
	for i := stop.Line; i > start.Line; i-- {
		m.Delete(i)
	}
	return start
}

// InsertText inserts text, which may span multiple lines separated by "\n", at pos.
// It returns the position right after the inserted text.
func InsertText(m Mutable, pos Position, text string) Position {
	line := m.Lines()[pos.Line].String()
	before, after := line[:pos.Col], line[pos.Col:]

	parts := strings.Split(text, "\n")
	if len(parts) == 1 {
		m.Update(pos.Line, TextLine(before+parts[0]+after))
		return Position{Col: pos.Col + len(parts[0]), Line: pos.Line}
	}

	m.Update(pos.Line, TextLine(before+parts[0]))
	for i := 1; i < len(parts)-1; i++ {
		m.Insert(pos.Line+i, TextLine(parts[i]))
	}
	last := len(parts) - 1
	m.Insert(pos.Line+last, TextLine(parts[last]+after))

	return Position{Col: len(parts[last]), Line: pos.Line + last}
}
