package editor

import (
	"bytes"
	"image/color"
	"testing"

	"github.com/google/go-cmp/cmp"
	"rmazur.io/chernetka/internal/editor/styles"
)

func TestColorLinePrinter_EmitTabSplit(t *testing.T) {
	bg := styles.TextStyle{BgColor: color.Gray{Y: 50}}
	const bgSet, reset = "\x1b[48;2;50;50;50m", "\x1b[0m"

	for _, tc := range []struct {
		name     string
		text     string
		style    styles.TextStyle
		from, to int
		want     string
	}{
		{name: "plain", text: "\t\tfoo", to: 100, want: "        foo"},
		{name: "styled", text: "\t\tfoo", style: bg, to: 100, want: bgSet + "        foo" + reset},
		{name: "styled without tabs", text: "foo", style: bg, to: 100, want: bgSet + "foo" + reset},
		{name: "styled tabs only", text: "\t\t", style: bg, to: 100, want: bgSet + "        " + reset},
		{name: "clipped left", text: "\t\tfoo", style: bg, from: 6, to: 100, want: bgSet + "  foo" + reset},
		{name: "clipped right", text: "\t\tfoo", style: bg, to: 6, want: bgSet + "      " + reset},
		{name: "scrolled away left", text: "\t\tfoo", style: bg, from: 50, to: 100, want: ""},
		{name: "scrolled away right", text: "\t\tfoo", style: bg, from: 0, to: 0, want: ""},
		{name: "empty", text: "", style: bg, to: 100, want: ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var out bytes.Buffer
			clp := colorLinePrinter{out: &out, tab: "    ", from: tc.from, to: tc.to}
			clp.emitTabSplit(tc.text, tc.style)
			if diff := cmp.Diff(out.String(), tc.want); diff != "" {
				t.Errorf("diff %s, want %q", diff, tc.want)
			}
		})
	}
}
