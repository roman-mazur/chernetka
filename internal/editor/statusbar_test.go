package editor

import (
	"bytes"
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/content"
)

func TestStatusBar_Render(t *testing.T) {
	cases := []struct {
		name     string
		buf      Buffer
		contains []string
		absent   []string
	}{
		{
			name: "normal mode shows status bar",
			buf: Buffer{
				Path: "test.txt",
				Content: &content.FullText{
					content.TextLine("hello"),
					content.TextLine("world"),
				},
			},
			contains: []string{"NORMAL", "test.txt", "1:1"},
			absent:   []string{"[*]"},
		},
		{
			name: "dirty buffer shows marker",
			buf: Buffer{
				Path:    "test.txt",
				Content: &content.FullText{content.TextLine("x")},
				mode:    ModeNormal,
				dirty:   true,
			},
			contains: []string{"[*]"},
		},
		{
			name: "cursor position reflected in status bar",
			buf: Buffer{
				Path: "test.txt",
				Content: &content.FullText{
					content.TextLine("first"),
					content.TextLine("second"),
					content.TextLine("third"),
				},
				mode: ModeNormal,
				c:    content.Position{Col: 3, Line: 1},
			},
			contains: []string{"2:4"},
		},
		{
			name: "command mode shows cmdline and hides status",
			buf: Buffer{
				Path:    "test.txt",
				Content: &content.FullText{content.TextLine("x")},
				mode:    ModeCommand,
				cmdline: "wq",
			},
			contains: []string{":wq"},
			absent:   []string{"NORMAL", "COMMAND"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.buf.w, tc.buf.h = 40, tc.buf.Content.Len()+2
			sb := StatusBar{buf: &tc.buf}
			var out bytes.Buffer
			sb.Render(&out)
			res := out.String()

			t.Log(res)
			for _, want := range tc.contains {
				if !strings.Contains(res, want) {
					t.Errorf("%q not found", want)
				}
			}
			for _, dontWant := range tc.absent {
				if strings.Contains(res, dontWant) {
					t.Errorf("%q found but should not be there", dontWant)
				}
			}
		})
	}

}
