package editor

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/inputs"
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
			buf := &Buffer{Content: &ft, c: content.Position{tc.cx, tc.cy}}
			RelMove{Dx: tc.d}.DoOnBuffer(buf, RenderPrefs{TabSize: 4})
			if buf.c.Col != tc.wantCx {
				t.Errorf("cx = %d, want %d (line %q, cx=%d, d=%d)",
					buf.c.Col, tc.wantCx, tc.lines[tc.cy], tc.cx, tc.d)
			}
		})
	}
}

func TestScreenMove_DoOnBuffer(t *testing.T) {
	tenLines := slices.Repeat(content.FullText{content.TextLine("line")}, 10)
	hundredLines := slices.Repeat(content.FullText{content.TextLine("line")}, 100)

	for _, tc := range []struct {
		dy  int // input
		buf Buffer
		// Expected buffer state.
		cx, cy, offset int
	}{
		{
			dy:     1,
			buf:    Buffer{Content: &tenLines, w: 80, h: 40},
			cx:     0,
			cy:     0,
			offset: 0,
		},
		{
			dy:     1,
			buf:    Buffer{Content: &hundredLines, w: 80, h: 40, c: content.Position{2, 0}},
			cx:     2,
			cy:     38,
			offset: 38,
		},
		{
			dy:     2,
			buf:    Buffer{Content: &hundredLines, w: 80, h: 40, c: content.Position{3, 5}},
			cx:     3,
			cy:     82,
			offset: 77,
		},
		{
			dy:     -1,
			buf:    Buffer{Content: &hundredLines, w: 80, h: 40, c: content.Position{1, 60}, offset: 40},
			cx:     1,
			cy:     22,
			offset: 2,
		},
		{
			dy:     -2,
			buf:    Buffer{Content: &hundredLines, w: 80, h: 40, c: content.Position{1, 60}, offset: 40},
			cx:     1,
			cy:     0,
			offset: 0,
		},
	} {
		name := fmt.Sprintf("content=%d/w=%d/h=%d/cx=%d/cy=%d/offset=%d",
			tc.buf.Content.Len(), tc.buf.w, tc.buf.h, tc.cx, tc.cx, tc.offset)
		t.Run(name, func(t *testing.T) {
			cmd := ScreenMove{ScreenD: tc.dy}
			cmd.DoOnBuffer(&tc.buf, RenderPrefs{TabSize: 4})
			if tc.buf.c.Col != tc.cx {
				t.Errorf("buf.cx = %d, want %d", tc.buf.c.Col, tc.cx)
			}
			if tc.buf.c.Line != tc.cy {
				t.Errorf("buf.cy = %d, want %d", tc.buf.c.Line, tc.cy)
			}
			if tc.buf.offset != tc.offset {
				t.Errorf("buf.offset = %d, want %d", tc.buf.offset, tc.offset)
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

func TestScroll(t *testing.T) {
	ftContent, err := content.LoadFullText(strings.NewReader("test content\nline 2\nline 3\nline 4"))
	if err != nil {
		t.Fatal(err)
	}

	buf := &Buffer{
		Content: &ftContent,
		w:       10,
		h:       2,
	}

	prefs := RenderPrefs{TabSize: 4}

	Scroll(inputs.ScrollDirectionDown).DoOnBuffer(buf, prefs)
	if buf.offset != 1 {
		t.Errorf("scroll not applied, buf.offset=%d", buf.offset)
	}
	Scroll(inputs.ScrollDirectionDown).DoOnBuffer(buf, prefs)
	if buf.offset != 2 {
		t.Errorf("second scroll not applied, buf.offset=%d", buf.offset)
	}
	Scroll(inputs.ScrollDirectionDown).DoOnBuffer(buf, prefs)
	if buf.offset != 2 {
		t.Errorf("third scroll should be skipped, buf.offset=%d", buf.offset)
	}
}

func TestSwitchMode(t *testing.T) {
	var buf Buffer

	if buf.canEdit() {
		t.Fatal("can edit buffer without content")
	}

	SwitchMode(ModeInsert).DoOnBuffer(&buf, RenderPrefs{TabSize: 4})
	if buf.mode != ModeNormal {
		t.Errorf("switched from normal to %s for non-editable content", buf.mode)
	}

	textContent := content.FullText{content.TextLine("a line")}
	buf.Content = &textContent
	if !buf.canEdit() {
		t.Fatal("can NOT edit text content")
	}

	SwitchMode(ModeInsert).DoOnBuffer(&buf, RenderPrefs{TabSize: 4})
	if buf.mode != ModeInsert {
		t.Errorf("insert mode not activated: %s", buf.mode)
	}
}

func TestSave_WritesFileAndClearsDirty(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.txt")

	ft := content.FullText{
		content.TextLine("first"),
		content.TextLine("second"),
		content.TextLine(""),
	}
	buf := &Buffer{Path: path, Content: &ft, dirty: true}

	(&Save{path}).DoOnBuffer(buf, RenderPrefs{})

	if buf.dirty {
		t.Errorf("dirty = true, want false after save")
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if got, want := string(data), "first\nsecond\n"; got != want {
		t.Errorf("file content = %q, want %q", got, want)
	}
}

func TestSave_NoPath(t *testing.T) {
	ft := content.FullText{content.TextLine("x")}
	buf := &Buffer{Content: &ft, dirty: true}

	(&Save{""}).DoOnBuffer(buf, RenderPrefs{}) // must not panic

	if !buf.dirty {
		t.Errorf("dirty cleared despite missing path")
	}
}

func TestSave_NotMutable(t *testing.T) {
	buf := &Buffer{
		Path:    "irrelevant",
		Content: &content.ErrorContent{Error: os.ErrNotExist},
		dirty:   true,
	}

	(&Save{""}).DoOnBuffer(buf, RenderPrefs{}) // must not panic

	if !buf.dirty {
		t.Errorf("dirty cleared despite non-mutable content")
	}
}
