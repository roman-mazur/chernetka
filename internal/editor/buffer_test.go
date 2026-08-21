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
				c:    position{3, 1},
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
				c:       position{3, 0},
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
			selections []span
			render     []escape.Span

			enableCurrentHighlight bool
		}{
			{
				name:    "one line",
				content: "line 1",
				selections: []span{
					{position{1, 0}, position{3, 0}},
				},
			},
			{
				name:    "one line reversed",
				content: "line 1",
				selections: []span{
					{position{3, 0}, position{1, 0}},
				},
				render: []escape.Span{
					{escape.Position{Offset: 1}, escape.Position{Offset: 3}},
				},
			},
			{
				name:    "one line with hl",
				content: "line 1",
				selections: []span{
					{position{1, 0}, position{4, 0}},
				},
				render: []escape.Span{
					{escape.Position{Line: 0, Offset: 0}, escape.Position{Line: 0, Offset: 1}},
					{escape.Position{Line: 0, Offset: 1}, escape.Position{Line: 0, Offset: 4}},
					{escape.Position{Line: 0, Offset: 4}, escape.Position{Line: 0, Offset: 6}},
					{escape.Position{Line: 0, Offset: 6}, escape.Position{Line: 0, Offset: 40}},
				},
				enableCurrentHighlight: true,
			},
			{
				name:    "multiple lines",
				content: "line 1\nline 2\nline 3",
				selections: []span{
					{position{3, 2}, position{1, 0}},
				},
				render: []escape.Span{
					{escape.Position{Line: 0, Offset: 1}, escape.Position{Line: 0, Offset: 6}},
					{escape.Position{Line: 1, Offset: 0}, escape.Position{Line: 1, Offset: 6}},
					{escape.Position{Line: 2, Offset: 0}, escape.Position{Line: 2, Offset: 3}},
				},
			},
			{
				name:    "multiple lines reversed",
				content: "line 1\nline 2\nline 3",
				selections: []span{
					{position{1, 0}, position{3, 2}},
				},
				render: []escape.Span{
					{escape.Position{Line: 0, Offset: 1}, escape.Position{Line: 0, Offset: 6}},
					{escape.Position{Line: 1, Offset: 0}, escape.Position{Line: 1, Offset: 6}},
					{escape.Position{Line: 2, Offset: 0}, escape.Position{Line: 2, Offset: 3}},
				},
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
			})
		}
	})
}

func makePos(p position) escape.Position { return escape.Position{Line: p.y, Offset: p.x} }
func makeSpan(s span) escape.Span        { return escape.Span{makePos(s.start), makePos(s.end)} }

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
		c:       position{3, 0},
	}

	buf.AcceptSuggestion("ntln")
	if got := buf.Content.Lines()[0].String(); got != "Println" {
		t.Errorf("line = %q, want %q", got, "Println")
	}
	if buf.c.x != 7 {
		t.Errorf("cx = %d, want 7", buf.c.x)
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
