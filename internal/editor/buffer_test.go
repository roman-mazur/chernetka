package editor

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/escape"
)

var thousandLines = slices.Repeat(content.FullText{content.TextLine("test")}, 1000)

func TestBuffer_Render(t *testing.T) {
	cases := []struct {
		name        string
		buf         *Buffer
		suggestions []string
		prefs       RenderPrefs
		contains    []string
		absent      []string
	}{
		{
			name: "normal mode shows numbered lines and status bar",
			buf: &Buffer{
				Path: "test.txt",
				Content: &content.FullText{
					content.TextLine("hello"),
					content.TextLine("world"),
				},
				mode: ModeNormal,
				w:    40,
				h:    5,
			},
			prefs:    RenderPrefs{TabSize: 4},
			contains: []string{"1 hello", "2 world", "NORMAL", "test.txt", "1:1"},
			absent:   []string{"[*]"},
		},
		{
			name: "hide line numbers",
			buf: &Buffer{
				Path: "test.txt",
				Content: &content.FullText{
					content.TextLine("hello"),
					content.TextLine("world"),
				},

				hideLineNumbers: true,
				mode:            ModeNormal,
				w:               40,
				h:               5,
			},
			prefs:    RenderPrefs{TabSize: 4},
			contains: []string{"hello", "world", "NORMAL", "test.txt", "1:1"},
			absent:   []string{"[*]", "1 hello", "2 world"},
		},
		{
			name: "dirty buffer shows marker",
			buf: &Buffer{
				Path:    "test.txt",
				Content: &content.FullText{content.TextLine("x")},
				mode:    ModeNormal,
				dirty:   true,
				w:       40,
				h:       3,
			},
			prefs:    RenderPrefs{TabSize: 4},
			contains: []string{"[*]"},
		},
		{
			name: "cursor position reflected in status bar",
			buf: &Buffer{
				Path: "test.txt",
				Content: &content.FullText{
					content.TextLine("first"),
					content.TextLine("second"),
					content.TextLine("third"),
				},
				mode: ModeNormal,
				c:    content.Position{3, 1},
				w:    40,
				h:    5,
			},
			prefs:    RenderPrefs{TabSize: 4},
			contains: []string{"2:4"},
		},
		{
			name: "tab expands to TabSize spaces",
			buf: &Buffer{
				Path:    "test.txt",
				Content: &content.FullText{content.TextLine("a\tb")},
				mode:    ModeNormal,
				w:       40,
				h:       3,
			},
			prefs:    RenderPrefs{TabSize: 4},
			contains: []string{"1 a    b"},
		},
		{
			name: "inline suggestion rendered as ghost text after cursor",
			buf: &Buffer{
				Path:    "test.go",
				Content: &content.FullText{content.TextLine("Pri")},
				mode:    ModeInsert,
				c:       content.Position{3, 0},
				w:       40,
				h:       3,
			},
			suggestions: []string{"ntln"},
			prefs:       RenderPrefs{TabSize: 4},
			contains:    []string{"1 Println"},
		},
		{
			name: "command mode shows cmdline and hides status",
			buf: &Buffer{
				Path:    "test.txt",
				Content: &content.FullText{content.TextLine("x")},
				mode:    ModeCommand,
				cmdline: "wq",
				w:       40,
				h:       3,
			},
			prefs:    RenderPrefs{TabSize: 4},
			contains: []string{":wq"},
			absent:   []string{"NORMAL", "COMMAND", "test.txt"},
		},
		{
			name: "2 digits line numbers",
			buf: &Buffer{
				Path:    "test.txt",
				Content: &thousandLines,
				w:       40,
				h:       10,
				offset:  5,
			},
			contains: []string{" 9 test", "10 test"},
		},
		{
			name: "3 digits line numbers",
			buf: &Buffer{
				Path:    "test.txt",
				Content: &thousandLines,
				w:       40,
				h:       10,
				offset:  95,
			},
			contains: []string{" 99 test", "100 test"},
		},
		{
			name: "4 digits line numbers",
			buf: &Buffer{
				Path:    "test.txt",
				Content: &thousandLines,
				w:       40,
				h:       10,
				offset:  995,
			},
			contains: []string{" 999 test", "1000 test"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.buf.xData = make(map[string]BufferExtData)
			tc.buf.xData["lsp"] = testSuggestionExt(tc.suggestions)

			var out bytes.Buffer
			tc.buf.Render(&out, &tc.prefs)
			got := escape.Clean(out.String())

			for _, want := range tc.contains {
				if !strings.Contains(got, want) {
					t.Errorf("output missing %q; got:\n%s", want, got)
				}
			}
			for _, unwanted := range tc.absent {
				if strings.Contains(got, unwanted) {
					t.Errorf("output should not contain %q; got:\n%s", unwanted, got)
				}
			}
		})
	}

	t.Run("selections", func(t *testing.T) {
		for n, tc := range []struct {
			name       string
			content    string
			selections []content.Span
			render     []escape.Span
			selText    string

			enableCurrentHighlight bool
		}{
			{
				name:    "selection start",
				content: "line 1",
				selections: []content.Span{
					{content.Position{1, 0}, content.Position{1, 0}},
				},
			},
			{
				name:    "one line",
				content: "line 1",
				selections: []content.Span{
					{content.Position{1, 0}, content.Position{3, 0}},
				},
				selText: "in",
			},
			{
				name:    "one line reversed",
				content: "line 1",
				selections: []content.Span{
					{content.Position{3, 0}, content.Position{1, 0}},
				},
				render: []escape.Span{
					{escape.Position{Offset: 1}, escape.Position{Offset: 3}},
				},
				selText: "in",
			},
			{
				name:    "one line with hl",
				content: "line 1",
				selections: []content.Span{
					{content.Position{1, 0}, content.Position{4, 0}},
				},
				render: []escape.Span{
					{escape.Position{Line: 0, Offset: 0}, escape.Position{Line: 0, Offset: 1}},
					{escape.Position{Line: 0, Offset: 1}, escape.Position{Line: 0, Offset: 4}},
					{escape.Position{Line: 0, Offset: 4}, escape.Position{Line: 0, Offset: 6}},
					{escape.Position{Line: 0, Offset: 6}, escape.Position{Line: 0, Offset: 40}},
				},
				enableCurrentHighlight: true,
				selText:                "ine",
			},
			{
				name:    "multiple lines",
				content: "line 1\nline 2\nline 3",
				selections: []content.Span{
					{content.Position{3, 2}, content.Position{1, 0}},
				},
				render: []escape.Span{
					{escape.Position{Line: 0, Offset: 1}, escape.Position{Line: 0, Offset: 6}},
					{escape.Position{Line: 1, Offset: 0}, escape.Position{Line: 1, Offset: 6}},
					{escape.Position{Line: 2, Offset: 0}, escape.Position{Line: 2, Offset: 3}},
				},
				selText: "ine 1\nline 2\nlin",
			},
			{
				name:    "multiple lines reversed",
				content: "line 1\nline 2\nline 3",
				selections: []content.Span{
					{content.Position{1, 0}, content.Position{3, 2}},
				},
				render: []escape.Span{
					{escape.Position{Line: 0, Offset: 1}, escape.Position{Line: 0, Offset: 6}},
					{escape.Position{Line: 1, Offset: 0}, escape.Position{Line: 1, Offset: 6}},
					{escape.Position{Line: 2, Offset: 0}, escape.Position{Line: 2, Offset: 3}},
				},
				selText: "ine 1\nline 2\nlin",
			},
		} {
			t.Run(fmt.Sprintf("%d/%s", n, tc.name), func(t *testing.T) {
				var buf Buffer
				testContent, err := content.LoadFullText(strings.NewReader(tc.content))
				if err != nil {
					t.Fatal(err)
				}
				buf.Content = &testContent

				buf.sel = tc.selections
				buf.w = 40                    // not too long for the test output
				buf.h = testContent.Len() + 1 // print all test input
				buf.hideLineNumbers = true
				buf.noCurrentLineHL = !tc.enableCurrentHighlight

				var out bytes.Buffer
				buf.Render(&out, &RenderPrefs{TabSize: 2})
				sOut := out.String()
				t.Log("output:\n" + sOut)
				t.Log("raw:\n" + strings.ReplaceAll(sOut, "\x1b", "^"))

				actuals := escape.ScanPositions(sOut, "48;2;", "0")
				expected := tc.render
				if expected == nil {
					expected = make([]escape.Span, len(tc.selections))
					for i := range expected {
						expected[i] = makeSpan(tc.selections[i])
					}
				}

				if diff := cmp.Diff(expected, actuals); diff != "" {
					t.Errorf("rendering mismatch (-want +got):\n%s", diff)
				}

				if diff := cmp.Diff(tc.selText, buf.SelectedText()); diff != "" {
					t.Errorf("SelectedText() mismatch (-want +got):\n%s", diff)
				}
			})
		}
	})
}

func makePos(p content.Position) escape.Position { return escape.Position{Line: p.Line, Offset: p.Col} }
func makeSpan(s content.Span) escape.Span        { return escape.Span{makePos(s.Start), makePos(s.End)} }

type testSuggestionExt []string

func (tse testSuggestionExt) TextSuggestion() string {
	if len(tse) == 0 {
		return ""
	}
	return tse[0]
}

func TestBuffer_AcceptSuggestion(t *testing.T) {
	buf := &Buffer{
		Content: &content.FullText{content.TextLine("Pri")},
		c:       content.Position{3, 0},
	}

	buf.AcceptSuggestion("ntln")
	if got := buf.Content.Lines()[0].String(); got != "Println" {
		t.Errorf("line = %q, want %q", got, "Println")
	}
	if buf.c.Col != 7 {
		t.Errorf("cx = %d, want 7", buf.c.Col)
	}
}

func TestBuffer_Text(t *testing.T) {
	buf := &Buffer{
		Content: &content.FullText{content.TextLine("Hello")},
	}
	doubleCheckBufferText(t, buf, "Hello")
	buf.Mutate().Insert(1, content.TextLine("World"))
	doubleCheckBufferText(t, buf, "Hello\nWorld")
}

func TestBuffer_SelectedText_MultiLineDownAndLeft(t *testing.T) {
	// Selection starts at a high column on an early, long line and ends at
	// a low column on a later, short line (a normal down-and-left drag).
	// This exercises Span.Min()/Max() with Start.Line < End.Line and
	// Start.Col > End.Col, which must not mix columns across lines.
	buf := &Buffer{
		Content: &content.FullText{
			content.TextLine("0123456789ABCDEFGHIJ"), // line 0, len 20
			content.TextLine("middle"),               // line 1
			content.TextLine("012"),                  // line 2, len 3
		},
		sel: []content.Span{{
			Start: content.Position{Col: 15, Line: 0},
			End:   content.Position{Col: 2, Line: 2},
		}},
	}

	want := "FGHIJ\nmiddle\n01"
	if got := buf.SelectedText(); got != want {
		t.Errorf("SelectedText() = %q, want %q", got, want)
	}
}

func BenchmarkBuffer_Text(b *testing.B) {
	f, err := os.Open("buffer_test.go")
	if err != nil {
		b.Fatal(err)
	}
	b.Cleanup(func() {
		_ = f.Close()
	})
	data, err := io.ReadAll(f)
	if err != nil {
		b.Fatal(err)
	}

	cnt, err := content.LoadFullText(bytes.NewReader(data))
	if err != nil {
		b.Fatal(err)
	}
	buf := Buffer{Content: &cnt}

	for b.Loop() {
		doubleCheckBufferText(b, &buf, string(data))
	}
}

func doubleCheckBufferText[T TB](t T, buf *Buffer, want string) {
	t.Helper()
	for i := range 2 {
		if res := buf.Text(); res != want {
			t.Errorf("check %d: text = %q, want %q", i, res, want)
		}
	}
}

type TB interface {
	*testing.T | *testing.B

	Helper()
	Errorf(string, ...any)
}

// TestBuffer_Render_FillsExactlyWindowHeight pins the invariant that keeps the buffer
// visible in a short terminal pane: Render must emit exactly h rows, i.e. h-1 line
// endings plus the status bar with no trailing newline. Emitting more rows than the
// pane holds scrolls the top of the content out of view.
func TestBuffer_Render_FillsExactlyWindowHeight(t *testing.T) {
	for _, tc := range []struct {
		name  string
		lines int
		mode  Mode
		h     int
	}{
		{name: "content shorter than pane", lines: 3, h: 21},
		{name: "content longer than pane", lines: 1000, h: 21},
		{name: "single row pane", lines: 1000, h: 2},
		{name: "command mode", lines: 1000, mode: ModeCommand, h: 21},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var buf Buffer
			testContent := content.FullText(slices.Repeat(content.FullText{content.TextLine("test")}, tc.lines))
			buf.Content = &testContent
			buf.mode = tc.mode
			buf.cmdline = "w"
			buf.w, buf.h = 40, tc.h

			var out bytes.Buffer
			buf.Render(&out, &RenderPrefs{TabSize: 2})

			if got, want := strings.Count(out.String(), "\r\n"), tc.h-1; got != want {
				t.Errorf("Render emitted %d line endings, want %d (window height %d)", got, want, tc.h)
			}
			if strings.HasSuffix(out.String(), "\r\n") {
				t.Error("Render ended with a line ending; the last row must not wrap to the next one")
			}
		})
	}
}
