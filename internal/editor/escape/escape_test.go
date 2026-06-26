package escape

import (
	"bytes"
	"image/color"
	"testing"
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
			ColorText(out, tt.args.text, tt.args.fg, tt.args.bg)
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("ColorText() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
