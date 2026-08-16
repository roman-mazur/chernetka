package escape

import "strings"

type Position struct {
	Line   int
	Offset int
}

type Span struct {
	Start Position
	End   Position
}

// ScanPositions searches the input string the occurrences of the defined matching tags.
func ScanPositions(input string, startTag, endTag string) (res []Span) {
	s := scanner{
		lines: strings.Split(input, "\n"),
	}
	for {
		if !s.Next(startTag) {
			return
		}
		span := Span{
			Start: s.Pos(),
		}
		if !s.Next(endTag) {
			return
		}
		span.End = s.Pos()
		res = append(res, span)
	}
}

type scanner struct {
	lines []string

	cursor  Position
	lastPos Position
}

func (s *scanner) Next(tag string) bool {
	for {
		if ok, _ := s.advance("\x1b[", ""); !ok {
			return false
		}
		lastPos := s.cursor

		ok, found := s.advance("m", tag)
		if !ok {
			return false
		}
		if found {
			s.lastPos = lastPos
			return true
		}
	}
}

func (s *scanner) Pos() Position { return s.lastPos }

func (s *scanner) advance(delim, search string) (ok, found bool) {
	found = search == ""
	for s.cursor.Line < len(s.lines) {
		line := s.lines[s.cursor.Line]
		for s.cursor.Offset < len(line)-len(delim)+1 {
			if !found && s.cursor.Offset < len(line)-len(search)+1 {
				found = line[s.cursor.Offset:s.cursor.Offset+len(search)] == search
			}
			if line[s.cursor.Offset:s.cursor.Offset+len(delim)] == delim {
				ok = true
				return
			}
			s.cursor.Offset++
		}

		s.cursor.Line++
		s.cursor.Offset = 0
	}
	return
}
