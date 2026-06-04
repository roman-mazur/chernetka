package editor

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"rmazur.io/x/edit/internal/content"
)

func TestRelMove_Dx(t *testing.T) {
	cases := []struct {
		name   string
		lines  []string // one entry per line in the buffer
		cy     int      // which line the cursor is on
		cx     int      // starting byte offset
		d      int      // direction: +1 right, -1 left
		wantCx int
	}{
		// Moving right — plain characters.
		{name: "right/plain first", lines: s("abc"), cx: 0, d: 1, wantCx: 1},
		{name: "right/plain middle", lines: s("abc"), cx: 1, d: 1, wantCx: 2},
		{name: "right/at end no-op", lines: s("abc"), cx: 3, d: 1, wantCx: 3},

		// Moving right — tab (one byte in byte-index world).
		{name: "right/over leading tab", lines: s("\tabc"), cx: 0, d: 1, wantCx: 1},
		{name: "right/plain after tab", lines: s("\tabc"), cx: 1, d: 1, wantCx: 2},
		{name: "right/over tab in middle", lines: s("a\tb"), cx: 1, d: 1, wantCx: 2},
		{name: "right/two tabs second", lines: s("\t\t"), cx: 1, d: 1, wantCx: 2},

		// Moving left — plain characters.
		{name: "left/plain middle", lines: s("abc"), cx: 1, d: -1, wantCx: 0},
		{name: "left/at start no-op", lines: s("abc"), cx: 0, d: -1, wantCx: 0},

		// Moving left — tab.
		{name: "left/over leading tab", lines: s("\tabc"), cx: 1, d: -1, wantCx: 0},
		{name: "left/plain before tab", lines: s("\tabc"), cx: 2, d: -1, wantCx: 1},
		{name: "left/over mid tab", lines: s("a\tb"), cx: 2, d: -1, wantCx: 1},
		{name: "left/two tabs", lines: s("\t\t"), cx: 2, d: -1, wantCx: 1},

		// Multi-byte runes step by rune width.
		{name: "right/over 3-byte rune", lines: s("a日b"), cx: 1, d: 1, wantCx: 4},
		{name: "left/over 3-byte rune", lines: s("a日b"), cx: 4, d: -1, wantCx: 1},

		// Empty line.
		{name: "empty/right no-op", lines: s(""), cx: 0, d: 1, wantCx: 0},
		{name: "empty/left no-op", lines: s(""), cx: 0, d: -1, wantCx: 0},

		// Correct line is used when cy > 0.
		{name: "cy>0/right plain", lines: s("tab line: \t!", "abc"), cy: 1, cx: 0, d: 1, wantCx: 1},
		{name: "cy>0/right tab", lines: s("abc", "\tabc"), cy: 1, cx: 0, d: 1, wantCx: 1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ft := make(content.FullText, len(tc.lines))
			for i, l := range tc.lines {
				ft[i] = content.TextLine(l)
			}
			buf := &Buffer{Content: &ft, cy: tc.cy, cx: tc.cx}
			RelMove{Dx: tc.d}.DoOnBuffer(buf, RenderPrefs{TabSize: 4})
			if buf.cx != tc.wantCx {
				t.Errorf("cx = %d, want %d (line %q, cx=%d, d=%d)",
					buf.cx, tc.wantCx, tc.lines[tc.cy], tc.cx, tc.d)
			}
		})
	}
}

// s is a shorthand for constructing a string slice in table literals.
func s(lines ...string) []string { return lines }

func TestSave(t *testing.T) {
	ftContent, err := content.LoadFullText(strings.NewReader("test content\nline 2"))
	if err != nil {
		t.Fatal(err)
	}

	buf := &Buffer{
		Content: &ftContent,
		dirty:   true,
	}
	dstPath := filepath.Join(t.TempDir(), "editor-save")
	save := &Save{DstPath: dstPath}
	save.DoOnBuffer(buf, RenderPrefs{TabSize: 4})

	if buf.dirty {
		t.Error("buffer remains dirty after save")
	}

	outFile, err := os.Open(dstPath)
	if err != nil {
		t.Fatal("cannot open saved file:", err)
	}
	t.Cleanup(func() { _ = outFile.Close() })
	outContent, err := content.LoadFullText(outFile)
	if err != nil {
		t.Fatal("cannot read saved file:", err)
	}
	t.Logf("saved content (len=%d): %s", outContent.Len(), outContent)
	if outContent.Len() != ftContent.Len() {
		t.Errorf("wanted %d, got %d lines", ftContent.Len(), outContent.Len())
	}
}
