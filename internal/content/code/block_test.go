package code

import (
	"testing"

	"rmazur.io/chernetka/internal/content"
)

func TestBlock_ContentLines(t *testing.T) {
	block := func(open, end int, closed bool) Block {
		return Block{
			Span:   content.Span{Start: content.Position{Line: open}, End: content.Position{Line: end}},
			Closed: closed,
		}
	}
	for _, tc := range []struct {
		name       string
		block      Block
		start, end int
	}{
		{name: "closed", block: block(2, 5, true), start: 3, end: 5},
		{name: "closed empty", block: block(2, 3, true), start: 3, end: 3},
		{name: "not closed", block: block(2, 5, false), start: 3, end: 6},
		{name: "not closed at the fence", block: block(2, 2, false), start: 3, end: 3},
	} {
		t.Run(tc.name, func(t *testing.T) {
			start, end := tc.block.ContentLines()
			if start != tc.start || end != tc.end {
				t.Errorf("ContentLines() = %d, %d; want %d, %d", start, end, tc.start, tc.end)
			}
		})
	}
}
