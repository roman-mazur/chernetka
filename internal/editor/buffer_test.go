package editor

import (
	"bufio"
	"bytes"
	"regexp"
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/content"
)

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
				cy:   1,
				cx:   3,
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
				cx:      3,
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
	}

	var ansiRE = regexp.MustCompile(`\x1b\[[0-9;?]*[A-Za-z]`)

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.buf.xData = make(map[string]BufferExtData)
			tc.buf.xData["lsp"] = testSuggestionExt(tc.suggestions)

			var out bytes.Buffer
			bw := bufio.NewWriter(&out)
			tc.buf.render(bw, &tc.prefs)
			got := ansiRE.ReplaceAllString(out.String(), "")

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
}

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
		cx:      3,
	}

	buf.AcceptSuggestion("ntln")
	if got := buf.Content.Lines()[0].String(); got != "Println" {
		t.Errorf("line = %q, want %q", got, "Println")
	}
	if buf.cx != 7 {
		t.Errorf("cx = %d, want 7", buf.cx)
	}
}
