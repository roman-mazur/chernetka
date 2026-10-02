package editor

import (
	"bytes"
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/content/code"
	"rmazur.io/chernetka/internal/vt/escape"
)

func TestStatusBar_Render(t *testing.T) {
	cases := []struct {
		name     string
		buf      Buffer
		cmd      string // the ex command typed in the command line
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
				_c:   content.Position{Col: 3, Line: 1},
			},
			contains: []string{"2:4"},
		},
		{
			name: "long path is shortened",
			buf: Buffer{
				Path:    "very/long/path/to/some/deeply/nested/dir/file.txt",
				Content: &content.FullText{content.TextLine("x")},
				dirty:   true,
			},
			contains: []string{"…", "nested/dir/file.txt [*]", "1:1"},
			absent:   []string{"very/long"},
		},
		{
			name: "command mode shows cmdline and hides status",
			buf: Buffer{
				Path:    "test.txt",
				Content: &content.FullText{content.TextLine("x")},
			},
			cmd:      "wq",
			contains: []string{":wq"},
			absent:   []string{"NORMAL", "COMMAND"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.buf.w, tc.buf.h = 40, tc.buf.Content.Len()+2
			sb := StatusBar{buf: &tc.buf}
			if tc.cmd != "" {
				sb.cmd = &cmdLine{buf: &tc.buf, text: tc.cmd}
				sb.cmd.prompt = newExPrompt(sb.cmd)
			}
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

func TestStatusBar_Diagnostics(t *testing.T) {
	for _, tc := range []struct {
		name  string
		diags []code.Diagnostic
		want  string
	}{
		{name: "none", want: ""},
		{name: "errors", diags: []code.Diagnostic{{Line: 1}, {Line: 1}}, want: "E2"},
		{name: "warnings", diags: []code.Diagnostic{{Severity: code.SeverityWarning}}, want: "W1"},
		{name: "both", diags: []code.Diagnostic{{Line: 3}, {Severity: code.SeverityWarning}}, want: "E1 W1"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			buf := newDiagnosticsTestBuffer("x", tc.diags...)
			sb := StatusBar{buf: buf}
			var out bytes.Buffer
			sb.Render(&out)
			got := escape.Clean(out.String())
			t.Log(got)
			if tc.want == "" {
				if strings.Contains(got, " E") || strings.Contains(got, " W") {
					t.Errorf("problems shown: %q", got)
				}
				return
			}
			if !strings.Contains(got, "test.go  "+tc.want) {
				t.Errorf("status %q, want %q", got, tc.want)
			}
		})
	}
}
