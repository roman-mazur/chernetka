package content

import (
	"slices"
	"strings"
	"testing"
)

func TestDeleteSpan(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		span  Span
		want  []string
		wantP Position
	}{
		{
			name:  "within a single line",
			lines: []string{"hello world"},
			span:  Span{Start: Position{Col: 5, Line: 0}, End: Position{Col: 11, Line: 0}},
			want:  []string{"hello"},
			wantP: Position{Col: 5, Line: 0},
		},
		{
			name:  "reversed span normalizes",
			lines: []string{"hello world"},
			span:  Span{Start: Position{Col: 11, Line: 0}, End: Position{Col: 5, Line: 0}},
			want:  []string{"hello"},
			wantP: Position{Col: 5, Line: 0},
		},
		{
			name:  "spans multiple lines",
			lines: []string{"abc", "def", "ghi"},
			span:  Span{Start: Position{Col: 1, Line: 0}, End: Position{Col: 2, Line: 2}},
			want:  []string{"ai"},
			wantP: Position{Col: 1, Line: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ft := make(FullText, len(tt.lines))
			for i, l := range tt.lines {
				ft[i] = TextLine(l)
			}

			got := DeleteSpan(&ft, tt.span)
			if got != tt.wantP {
				t.Errorf("DeleteSpan() position = %+v, want %+v", got, tt.wantP)
			}
			assertLines(t, &ft, tt.want)
		})
	}
}

func TestInsertText(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		pos   Position
		text  string
		want  []string
		wantP Position
	}{
		{
			name:  "single line insert",
			lines: []string{"ac"},
			pos:   Position{Col: 1, Line: 0},
			text:  "b",
			want:  []string{"abc"},
			wantP: Position{Col: 2, Line: 0},
		},
		{
			name:  "two line insert",
			lines: []string{"ad"},
			pos:   Position{Col: 1, Line: 0},
			text:  "b\nc",
			want:  []string{"ab", "cd"},
			wantP: Position{Col: 1, Line: 1},
		},
		{
			name:  "three line insert",
			lines: []string{"ae"},
			pos:   Position{Col: 1, Line: 0},
			text:  "b\nc\nd",
			want:  []string{"ab", "c", "de"},
			wantP: Position{Col: 1, Line: 2},
		},
		{
			name:  "lines before and after",
			lines: []string{"0", "ae", "2", "3"},
			pos:   Position{Col: 1, Line: 1},
			text:  "b\nc\nd",
			want:  []string{"0", "ab", "c", "de", "2", "3"},
			wantP: Position{Col: 1, Line: 3},
		},
		{
			name:  "trailing new line",
			lines: []string{"ab", "c"},
			pos:   Position{Col: 1, Line: 0},
			text:  "x\ny\n",
			want:  []string{"ax", "y", "b", "c"},
			wantP: Position{Col: 0, Line: 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ft := make(FullText, len(tt.lines))
			for i, l := range tt.lines {
				ft[i] = TextLine(l)
			}

			got := InsertText(&ft, tt.pos, tt.text)
			if got != tt.wantP {
				t.Errorf("InsertText() position = %+v, want %+v", got, tt.wantP)
			}
			assertLines(t, &ft, tt.want)
		})
	}
}

// BenchmarkInsertText measures a paste of many lines in the middle of a big document.
func BenchmarkInsertText(b *testing.B) {
	pasted := strings.Repeat("\tfmt.Println(\"a pasted line\", x)\n", 20000)
	base := make(FullText, 10000)
	for i := range base {
		base[i] = TextLine("an existing line")
	}
	b.SetBytes(int64(len(pasted)))
	for b.Loop() {
		ft := slices.Clone(base)
		InsertText(&ft, Position{Line: len(base) / 2}, pasted)
	}
}

func assertLines(t *testing.T, ft *FullText, want []string) {
	t.Helper()
	var got []string
	for _, l := range ft.Lines() {
		got = append(got, l.String())
	}
	if strings.Join(got, "|") != strings.Join(want, "|") {
		t.Errorf("lines = %q, want %q", got, want)
	}
}
