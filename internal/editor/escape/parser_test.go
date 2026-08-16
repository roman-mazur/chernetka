package escape

import (
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
				{Start: Position{0, 5}, End: Position{0, 12}},
			},
		},
		{
			name:  "two lines",
			input: "line \x1b[123m1\x1b[0m\nline \x1b[123m2\x1b[0m",
			start: "123", end: "0",
			expected: []Span{
				{Start: Position{0, 5}, End: Position{0, 12}},
				{Start: Position{1, 5}, End: Position{1, 12}},
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
				{Start: Position{0, 5}, End: Position{0, 12}},
				{Start: Position{1, 0}, End: Position{1, 18}},
			},
		},
		{
			name:  "no escape",
			input: "line 1\nline 2",
			start: "123", end: "0",
			expected: nil,
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			res := ScanPositions(testCase.input, testCase.start, testCase.end)
			if diff := cmp.Diff(testCase.expected, res); diff != "" {
				t.Errorf("(-want, +got)\n%s", diff)
			}
		})
	}
}
