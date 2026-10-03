package escape

import (
	"bytes"
	"image/color"
	"io"
	"testing"

	"rmazur.io/chernetka/internal/editor/styles"
)

func TestColorText(t *testing.T) {
	type args struct {
		text string
		fg   color.Color
		bg   color.Color
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			name: "8-bit fg color",
			args: args{
				text: "test",
				fg:   color.White,
			},
			wantOut: "\x1b[38;2;255;255;255mtest\x1B[0m",
		},
		{
			name: "8-bit bg color",
			args: args{
				text: "test",
				bg:   color.White,
			},
			wantOut: "\x1b[48;2;255;255;255mtest\x1B[0m",
		},
		{
			name: "8-bit colors",
			args: args{
				text: "test",
				bg:   color.White,
				fg:   color.Black,
			},
			wantOut: "\x1b[38;2;0;0;0;48;2;255;255;255mtest\x1B[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			StyleText(out, tt.args.text, styles.TextStyle{TextColor: tt.args.fg, BgColor: tt.args.bg})
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("ColorText() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestStyleText(t *testing.T) {
	var buf bytes.Buffer
	buf.WriteString("\n")
	StyleText(&buf, "bold", styles.TextStyle{
		Bold: true,
	})
	buf.WriteString("\n")
	StyleText(&buf, "italic", styles.TextStyle{
		Italic: true,
	})
	buf.WriteString("\n")
	StyleText(&buf, "both", styles.TextStyle{
		Bold:   true,
		Italic: true,
	})
	buf.WriteString("\n")
	StyleText(&buf, "normal", styles.TextStyle{})
	buf.WriteString("\n")
	StyleText(&buf, "bold blue", styles.TextStyle{
		Bold:      true,
		TextColor: color.RGBA{B: 255, A: 255},
	})
	t.Log(buf.String())
}

func TestHideCursor(t *testing.T) {
	for _, tc := range []struct {
		name        string
		thinCursor  bool
		wantHide    string
		wantRestore string
	}{
		{name: "block cursor", wantHide: "\x1b[?25l", wantRestore: "\x1b[1 q\x1b[?25h"},
		{name: "thin cursor", thinCursor: true, wantHide: "\x1b[?25l", wantRestore: "\x1b[5 q\x1b[?25h"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var out bytes.Buffer
			restore := HideCursor(&out, tc.thinCursor)
			if got := out.String(); got != tc.wantHide {
				t.Errorf("hide = %q, want %q", got, tc.wantHide)
			}
			out.Reset()
			restore.Undo()
			if got := out.String(); got != tc.wantRestore {
				t.Errorf("restore = %q, want %q", got, tc.wantRestore)
			}
		})
	}
}

func BenchmarkStyleText(b *testing.B) {
	style := styles.TextStyle{
		Bold:      true,
		Italic:    true,
		TextColor: color.White,
		BgColor:   color.Gray{Y: 50},
	}
	for b.Loop() {
		StyleText(io.Discard, "some text examples", style)
	}
}
