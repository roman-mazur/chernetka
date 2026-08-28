package content

import "testing"

func TestSpan_Min(t *testing.T) {
	tests := []struct {
		name string
		span Span
		want Position
	}{
		{
			name: "start before end",
			span: Span{Start: Position{Col: 1, Line: 0}, End: Position{Col: 5, Line: 2}},
			want: Position{Col: 1, Line: 0},
		},
		{
			name: "end before start",
			span: Span{Start: Position{Col: 5, Line: 2}, End: Position{Col: 1, Line: 0}},
			want: Position{Col: 1, Line: 0},
		},
		{
			name: "same line, start col greater",
			span: Span{Start: Position{Col: 5, Line: 1}, End: Position{Col: 2, Line: 1}},
			want: Position{Col: 2, Line: 1},
		},
		{
			name: "same line, end col greater",
			span: Span{Start: Position{Col: 2, Line: 1}, End: Position{Col: 5, Line: 1}},
			want: Position{Col: 2, Line: 1},
		},
		{
			name: "equal positions",
			span: Span{Start: Position{Col: 3, Line: 4}, End: Position{Col: 3, Line: 4}},
			want: Position{Col: 3, Line: 4},
		},
		{
			// Down-and-left drag: Start is on an earlier line but at a
			// higher column than End. Min must be Start verbatim, not a
			// column mixed in from End's (unrelated) line.
			name: "different lines, start col greater than end col",
			span: Span{Start: Position{Col: 20, Line: 1}, End: Position{Col: 3, Line: 4}},
			want: Position{Col: 20, Line: 1},
		},
		{
			// Same as above with Start/End swapped.
			name: "different lines (reversed), end col greater than start col",
			span: Span{Start: Position{Col: 3, Line: 4}, End: Position{Col: 20, Line: 1}},
			want: Position{Col: 20, Line: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.span.Min(); got != tt.want {
				t.Errorf("Min() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestSpan_Max(t *testing.T) {
	tests := []struct {
		name string
		span Span
		want Position
	}{
		{
			name: "start before end",
			span: Span{Start: Position{Col: 1, Line: 0}, End: Position{Col: 5, Line: 2}},
			want: Position{Col: 5, Line: 2},
		},
		{
			name: "end before start",
			span: Span{Start: Position{Col: 5, Line: 2}, End: Position{Col: 1, Line: 0}},
			want: Position{Col: 5, Line: 2},
		},
		{
			name: "same line, start col greater",
			span: Span{Start: Position{Col: 5, Line: 1}, End: Position{Col: 2, Line: 1}},
			want: Position{Col: 5, Line: 1},
		},
		{
			name: "same line, end col greater",
			span: Span{Start: Position{Col: 2, Line: 1}, End: Position{Col: 5, Line: 1}},
			want: Position{Col: 5, Line: 1},
		},
		{
			name: "equal positions",
			span: Span{Start: Position{Col: 3, Line: 4}, End: Position{Col: 3, Line: 4}},
			want: Position{Col: 3, Line: 4},
		},
		{
			// Down-and-left drag: End is on the later line but at a lower
			// column than Start. Max must be End verbatim, not a column
			// mixed in from Start's (unrelated) line.
			name: "different lines, start col greater than end col",
			span: Span{Start: Position{Col: 20, Line: 1}, End: Position{Col: 3, Line: 4}},
			want: Position{Col: 3, Line: 4},
		},
		{
			// Same as above with Start/End swapped.
			name: "different lines (reversed), end col greater than start col",
			span: Span{Start: Position{Col: 3, Line: 4}, End: Position{Col: 20, Line: 1}},
			want: Position{Col: 3, Line: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.span.Max(); got != tt.want {
				t.Errorf("Max() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestSpan_ContainsLine(t *testing.T) {
	tests := []struct {
		name string
		span Span
		line int
		want bool
	}{
		{
			name: "within range, normal order",
			span: Span{Start: Position{Line: 1}, End: Position{Line: 5}},
			line: 3,
			want: true,
		},
		{
			name: "within range, reversed order",
			span: Span{Start: Position{Line: 5}, End: Position{Line: 1}},
			line: 3,
			want: true,
		},
		{
			name: "at min boundary",
			span: Span{Start: Position{Line: 1}, End: Position{Line: 5}},
			line: 1,
			want: true,
		},
		{
			name: "at max boundary",
			span: Span{Start: Position{Line: 1}, End: Position{Line: 5}},
			line: 5,
			want: true,
		},
		{
			name: "below range",
			span: Span{Start: Position{Line: 1}, End: Position{Line: 5}},
			line: 0,
			want: false,
		},
		{
			name: "above range",
			span: Span{Start: Position{Line: 1}, End: Position{Line: 5}},
			line: 6,
			want: false,
		},
		{
			name: "single line span, matching",
			span: Span{Start: Position{Line: 2}, End: Position{Line: 2}},
			line: 2,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.span.ContainsLine(tt.line); got != tt.want {
				t.Errorf("ContainsLine(%d) = %v, want %v", tt.line, got, tt.want)
			}
		})
	}
}

func TestSpan_LineProjection(t *testing.T) {
	tests := []struct {
		name    string
		span    Span
		idx     int
		lineLen int
		want    Span
	}{
		{
			name:    "single line span matches idx",
			span:    Span{Start: Position{Col: 1, Line: 2}, End: Position{Col: 4, Line: 2}},
			idx:     2,
			lineLen: 10,
			want:    Span{Start: Position{Col: 1, Line: 2}, End: Position{Col: 4, Line: 2}},
		},
		{
			name:    "multi-line span, first line",
			span:    Span{Start: Position{Col: 3, Line: 1}, End: Position{Col: 4, Line: 3}},
			idx:     1,
			lineLen: 10,
			want:    Span{Start: Position{Col: 3, Line: 1}, End: Position{Col: 10, Line: 1}},
		},
		{
			name:    "multi-line span, middle line",
			span:    Span{Start: Position{Col: 3, Line: 1}, End: Position{Col: 4, Line: 3}},
			idx:     2,
			lineLen: 10,
			want:    Span{Start: Position{Col: 0, Line: 2}, End: Position{Col: 10, Line: 2}},
		},
		{
			name:    "multi-line span, last line",
			span:    Span{Start: Position{Col: 3, Line: 1}, End: Position{Col: 4, Line: 3}},
			idx:     3,
			lineLen: 10,
			want:    Span{Start: Position{Col: 0, Line: 3}, End: Position{Col: 4, Line: 3}},
		},
		{
			name:    "reversed span is normalized before projection",
			span:    Span{Start: Position{Col: 4, Line: 3}, End: Position{Col: 3, Line: 1}},
			idx:     2,
			lineLen: 10,
			want:    Span{Start: Position{Col: 0, Line: 2}, End: Position{Col: 10, Line: 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.span.LineProjection(tt.idx, tt.lineLen); got != tt.want {
				t.Errorf("LineProjection(%d, %d) = %+v, want %+v", tt.idx, tt.lineLen, got, tt.want)
			}
		})
	}
}
