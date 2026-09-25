package escape

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestScanPositions(t *testing.T) {
	for _, testCase := range []struct {
		name       string
		input      string
		start, end string
		expected   []Span
	}{
		{
			name:  "empty",
			input: "",
			start: "1", end: "2",
			expected: nil,
		},
		{
			name:  "one line",
			input: "line \x1b[123m1\x1b[0m",
			start: "123", end: "0",
			expected: []Span{
				{Start: Position{0, 5}, End: Position{0, 6}},
			},
		},
		{
			name:  "two lines",
			input: "line \x1b[123m1\x1b[0m\nline \x1b[123m2\x1b[0m",
			start: "123", end: "0",
			expected: []Span{
				{Start: Position{0, 5}, End: Position{0, 6}},
				{Start: Position{1, 5}, End: Position{1, 6}},
			},
		},
		{
			name:  "not closed",
			input: "line \x1b[123m1",
			start: "123", end: "0",
			expected: nil,
		},
		{
			name:  "not closed inside",
			input: "line \x1b[123m1\x1b[0m\n\x1b[123mline \x1b[123m2\x1b[0m",
			start: "123", end: "0",
			expected: []Span{
				{Start: Position{0, 5}, End: Position{0, 6}},
				{Start: Position{1, 0}, End: Position{1, 6}},
			},
		},
		{
			name:  "no escape",
			input: "line 1\nline 2",
			start: "123", end: "0",
			expected: nil,
		},
		{
			name:  "span over lines",
			input: "line \x1b[123m1\nline 2\nline\x1b[0m 3",
			start: "123", end: "0",
			expected: []Span{
				{Start: Position{0, 5}, End: Position{2, 4}},
			},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			res := ScanPositions(testCase.input, testCase.start, testCase.end)
			if diff := cmp.Diff(testCase.expected, res); diff != "" {
				t.Errorf("(-want, +got)\n%s", diff)
			}

			noEscape := strings.Split(Clean(testCase.input), "\n")
			for _, span := range res {
				line := noEscape[span.Start.Line]
				if span.Start.Offset > len(line) || span.End.Offset > len(line) {
					t.Errorf("span %v out of bound for %q", span, line)
				}
			}
		})
	}
}
