package editor

import (
	"testing"

	"rmazur.io/x/edit/internal/content"
)

func TestScreenToTextIdx(t *testing.T) {
	cases := []struct {
		name      string
		line      string
		screenCol int
		tabSize   int
		want      int
	}{
		// Plain text — each character is one column wide.
		{name: "plain/first", line: "abc", screenCol: 0, tabSize: 4, want: 0},
		{name: "plain/middle", line: "abc", screenCol: 1, tabSize: 4, want: 1},
		{name: "plain/last", line: "abc", screenCol: 2, tabSize: 4, want: 2},
		{name: "plain/past end", line: "abc", screenCol: 3, tabSize: 4, want: 3},

		// Tab at start.
		{name: "leading tab/col 0", line: "\tabc", screenCol: 0, tabSize: 4, want: 0},
		{name: "leading tab/inside tab", line: "\tabc", screenCol: 2, tabSize: 4, want: 0},
		{name: "leading tab/last tab col", line: "\tabc", screenCol: 3, tabSize: 4, want: 0},
		{name: "leading tab/after tab", line: "\tabc", screenCol: 4, tabSize: 4, want: 1},

		// Tab in the middle.
		{name: "mid tab/before", line: "a\tb", screenCol: 0, tabSize: 4, want: 0},
		{name: "mid tab/on tab", line: "a\tb", screenCol: 1, tabSize: 4, want: 1},
		{name: "mid tab/inside tab", line: "a\tb", screenCol: 3, tabSize: 4, want: 1},
		{name: "mid tab/after tab", line: "a\tb", screenCol: 5, tabSize: 4, want: 2},

		// Two tabs.
		{name: "two tabs/second tab", line: "\t\t", screenCol: 4, tabSize: 4, want: 1},
		{name: "two tabs/past end", line: "\t\t", screenCol: 8, tabSize: 4, want: 2},

		// Different tab size.
		{name: "tabsize2/on tab", line: "\ta", screenCol: 0, tabSize: 2, want: 0},
		{name: "tabsize2/after tab", line: "\ta", screenCol: 2, tabSize: 2, want: 1},

		// Empty line.
		{name: "empty", line: "", screenCol: 0, tabSize: 4, want: 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := screenToTextIdx(tc.line, tc.screenCol, tc.tabSize)
			if got != tc.want {
				t.Errorf("screenToTextIdx(%q, %d, %d) = %d, want %d",
					tc.line, tc.screenCol, tc.tabSize, got, tc.want)
			}
		})
	}
}

func TestBuffer_MoveHorizontal(t *testing.T) {
	cases := []struct {
		name    string
		lines   []string // one entry per line in the buffer
		cy      int      // which line the cursor is on
		cx      int      // starting screen column
		d       int      // direction: +1 right, -1 left
		tabSize int
		wantCx  int
	}{
		// Moving right — plain characters.
		{name: "right/plain first", lines: s("abc"), cx: 0, d: 1, tabSize: 4, wantCx: 1},
		{name: "right/plain middle", lines: s("abc"), cx: 1, d: 1, tabSize: 4, wantCx: 2},
		{name: "right/at end no-op", lines: s("abc"), cx: 3, d: 1, tabSize: 4, wantCx: 3},

		// Moving right — tab.
		{name: "right/over leading tab", lines: s("\tabc"), cx: 0, d: 1, tabSize: 4, wantCx: 4},
		{name: "right/plain after tab", lines: s("\tabc"), cx: 4, d: 1, tabSize: 4, wantCx: 5},
		{name: "right/over tab in middle", lines: s("a\tb"), cx: 1, d: 1, tabSize: 4, wantCx: 5},
		{name: "right/two tabs second", lines: s("\t\t"), cx: 4, d: 1, tabSize: 4, wantCx: 8},

		// Moving left — plain characters.
		{name: "left/plain middle", lines: s("abc"), cx: 1, d: -1, tabSize: 4, wantCx: 0},
		{name: "left/at start no-op", lines: s("abc"), cx: 0, d: -1, tabSize: 4, wantCx: 0},

		// Moving left — tab.
		{name: "left/over leading tab", lines: s("\tabc"), cx: 4, d: -1, tabSize: 4, wantCx: 0},
		{name: "left/plain before tab", lines: s("\tabc"), cx: 5, d: -1, tabSize: 4, wantCx: 4},
		{name: "left/over mid tab", lines: s("a\tb"), cx: 5, d: -1, tabSize: 4, wantCx: 1},
		{name: "left/two tabs", lines: s("\t\t"), cx: 8, d: -1, tabSize: 4, wantCx: 4},

		// Different tab size.
		{name: "tabsize2/right over tab", lines: s("\ta"), cx: 0, d: 1, tabSize: 2, wantCx: 2},
		{name: "tabsize2/left over tab", lines: s("\ta"), cx: 2, d: -1, tabSize: 2, wantCx: 0},

		// Empty line.
		{name: "empty/right no-op", lines: s(""), cx: 0, d: 1, tabSize: 4, wantCx: 0},
		{name: "empty/left no-op", lines: s(""), cx: 0, d: -1, tabSize: 4, wantCx: 0},

		// Correct line is used when cy > 0.
		{name: "cy>0/right plain", lines: s("tab line: \t!", "abc"), cy: 1, cx: 0, d: 1, tabSize: 4, wantCx: 1},
		{name: "cy>0/right tab", lines: s("abc", "\tabc"), cy: 1, cx: 0, d: 1, tabSize: 4, wantCx: 4},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ft := make(content.FullText, len(tc.lines))
			for i, l := range tc.lines {
				ft[i] = content.TextLine(l)
			}
			buf := &Buffer{content: &ft, cy: tc.cy, cx: tc.cx}
			buf.moveHorizontal(tc.d, tc.tabSize)
			if buf.cx != tc.wantCx {
				t.Errorf("cx = %d, want %d (line %q, cx=%d, d=%d, TabSize=%d)",
					buf.cx, tc.wantCx, tc.lines[tc.cy], tc.cx, tc.d, tc.tabSize)
			}
		})
	}
}

// s is a shorthand for constructing a string slice in table literals.
func s(lines ...string) []string { return lines }
