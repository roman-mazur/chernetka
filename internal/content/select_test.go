package content

import "testing"

func TestSelect(t *testing.T) {
	ft := FullText{
		TextLine("0123456789ABCDEFGHIJ"), // line 0, len 20
		TextLine("middle"),               // line 1
		TextLine("012"),                  // line 2, len 3
	}

	tests := []struct {
		name  string
		spans []Span
		want  string
	}{
		{
			name:  "no spans",
			spans: nil,
			want:  "",
		},
		{
			name:  "single line",
			spans: []Span{{Start: Position{Col: 0, Line: 1}, End: Position{Col: 3, Line: 1}}},
			want:  "mid",
		},
		{
			name: "down-and-left drag normalizes via Min/Max",
			spans: []Span{{
				Start: Position{Col: 15, Line: 0},
				End:   Position{Col: 2, Line: 2},
			}},
			want: "FGHIJ\nmiddle\n01",
		},
		{
			name: "multiple spans concatenate in order",
			spans: []Span{
				{Start: Position{Col: 0, Line: 1}, End: Position{Col: 3, Line: 1}},
				{Start: Position{Col: 0, Line: 2}, End: Position{Col: 3, Line: 2}},
			},
			want: "mid012",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Select(&ft, tt.spans); got != tt.want {
				t.Errorf("Select() = %q, want %q", got, tt.want)
			}
		})
	}
}
