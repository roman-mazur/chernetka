package editor

import (
	"bufio"
	"bytes"
	"regexp"
	"strings"
	"testing"

	"rmazur.io/x/edit/internal/content"
)

func TestBuffer_Render(t *testing.T) {
	cases := []struct {
		name     string
		buf      *Buffer
		prefs    RenderPrefs
		contains []string
		absent   []string
	}{
		{
			name: "normal mode shows numbered lines and status bar",
			buf: &Buffer{
				path: "test.txt",
				content: &content.FullText{
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
				path:    "test.txt",
				content: &content.FullText{content.TextLine("x")},
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
				path: "test.txt",
				content: &content.FullText{
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
				path:    "test.txt",
				content: &content.FullText{content.TextLine("a\tb")},
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
				path:    "test.go",
				content: &content.FullText{content.TextLine("Pri")},
				mode:    ModeInsert,
				cx:      3,
				lsp: lspBufferData{
					suggestions: []string{"ntln"},
				},
				w: 40,
				h: 3,
			},
			prefs:    RenderPrefs{TabSize: 4},
			contains: []string{"1 Println"},
		},
		{
			name: "command mode shows cmdline and hides status",
			buf: &Buffer{
				path:    "test.txt",
				content: &content.FullText{content.TextLine("x")},
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

func TestBuffer_AcceptSuggestion(t *testing.T) {
	buf := &Buffer{
		content: &content.FullText{content.TextLine("Pri")},
		cx:      3,
		lsp:     lspBufferData{suggestions: []string{"ntln"}},
	}
	buf.acceptSuggestion(buf.lsp.suggestions[0])
	if got := buf.content.Lines()[0].String(); got != "Println" {
		t.Errorf("line = %q, want %q", got, "Println")
	}
	if buf.cx != 7 {
		t.Errorf("cx = %d, want 7", buf.cx)
	}
	if buf.lsp.suggestions != nil {
		t.Errorf("suggestions = %v, want cleared", buf.lsp.suggestions)
	}
}
