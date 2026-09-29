package extcomment

import (
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor"
)

func TestIntegration_MakeBufferData(t *testing.T) {
	for _, tc := range []struct {
		path string
		want string // "" for no data
	}{
		{path: "main.go", want: "//"},
		{path: "DB.SQL", want: "--"},
		{path: "config.yml", want: "#"},
		{path: "/home/me/.zshrc", want: "#"},
		{path: "README.md"},
		{path: "notes.txt"},
		{path: ""},
	} {
		t.Run(tc.path, func(t *testing.T) {
			buf := &editor.Buffer{Path: tc.path, Content: &content.FullText{content.TextLine("x")}}
			data := new(Integration).MakeBufferData(buf)
			if tc.want == "" {
				if data != nil {
					t.Errorf("data = %v, want nil", data)
				}
				return
			}
			d, ok := data.(*document)
			if !ok {
				t.Fatalf("data = %v, want a document", data)
			}
			if d.prefix != tc.want {
				t.Errorf("prefix = %q, want %q", d.prefix, tc.want)
			}
		})
	}
}

const (
	ctrlSlash  = "\x1f"
	right      = "\x1b[C"
	shiftDown  = "\x1b[1;2B"
	shiftRight = "\x1b[1;2C"
	shiftEnd   = "\x1b[1;2F"
)

func TestToggle(t *testing.T) {
	for _, tc := range []struct {
		name  string
		path  string
		text  string
		input []string // each one is sent at once, the text is checked after each
		want  []string
	}{
		{
			name:  "cursor line",
			path:  "a.go",
			text:  "\tx := 1\n\ty := 2",
			input: []string{"i" + ctrlSlash, ctrlSlash},
			want:  []string{"\t// x := 1\n\ty := 2", "\tx := 1\n\ty := 2"},
		},
		{
			name:  "uncomment without space",
			path:  "a.go",
			text:  "//x",
			input: []string{"i" + ctrlSlash},
			want:  []string{"x"},
		},
		{
			name:  "selection at smallest indent",
			path:  "a.yaml",
			text:  "a:\n  b: 1\n\nc: 2\nd: 3",
			input: []string{"i" + strings.Repeat(shiftDown, 4) + ctrlSlash, ctrlSlash},
			want:  []string{"# a:\n#   b: 1\n\n# c: 2\nd: 3", "a:\n  b: 1\n\nc: 2\nd: 3"},
		},
		{
			name:  "mixed selection is commented",
			path:  "a.nix",
			text:  "# a\nb",
			input: []string{"i" + shiftDown + shiftEnd + ctrlSlash},
			want:  []string{"# # a\n# b"},
		},
		{
			name:  "blank line",
			path:  "a.go",
			text:  "\t\nx",
			input: []string{"i" + ctrlSlash},
			want:  []string{"\t\nx"},
		},
		{
			name:  "normal mode",
			path:  "a.go",
			text:  "x",
			input: []string{ctrlSlash},
			want:  []string{"x"},
		},
		{
			name:  "unknown file type",
			path:  "a.txt",
			text:  "x",
			input: []string{"i" + ctrlSlash},
			want:  []string{"x"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			h := editor.NewTestHarness()
			h.Extend(new(Integration))
			if err := h.OpenReader(tc.path, strings.NewReader(tc.text)); err != nil {
				t.Fatal(err)
			}
			h.Run(t)
			for i, in := range tc.input {
				h.SendInput(t, []byte(in))
				if got := topText(t, h); got != tc.want[i] {
					t.Errorf("text after %q = %q, want %q", in, got, tc.want[i])
				}
			}
		})
	}
}

func TestToggle_CursorAndSelection(t *testing.T) {
	h := editor.NewTestHarness()
	h.Extend(new(Integration))
	if err := h.OpenReader("a.go", strings.NewReader("\tx := 1\n\ty := 2")); err != nil {
		t.Fatal(err)
	}
	h.Run(t)
	pos := func(line, col int) content.Position { return content.Position{Line: line, Col: col} }

	h.SendInput(t, []byte("i"+right+right+shiftDown+shiftRight+ctrlSlash))
	onLoop(t, h, func(buf *editor.Buffer) {
		if cx, cy := buf.Pos(); cx != 6 || cy != 1 {
			t.Errorf("cursor = %d:%d, want 1:6", cy, cx)
		}
		want := content.Span{Start: pos(0, 5), End: pos(1, 6)}
		if got := buf.Selection(); len(got) != 1 || got[0] != want {
			t.Errorf("selection = %v, want [%v]", got, want)
		}
	})

	h.SendInput(t, []byte(ctrlSlash))
	onLoop(t, h, func(buf *editor.Buffer) {
		if cx, cy := buf.Pos(); cx != 3 || cy != 1 {
			t.Errorf("cursor after uncommenting = %d:%d, want 1:3", cy, cx)
		}
	})

	// The input still replaces the selection that followed the edits.
	h.SendInput(t, []byte(ctrlSlash+"z"))
	onLoop(t, h, func(buf *editor.Buffer) {
		if got, want := buf.Text(), "\t// xz:= 2"; got != want {
			t.Errorf("text after typing = %q, want %q", got, want)
		}
		if got := buf.Selection(); len(got) != 0 {
			t.Errorf("selection after typing = %v, want none", got)
		}
	})
}

// onLoop runs f with the top buffer on the editor loop.
func onLoop(t *testing.T, h *editor.TestHarness, f func(buf *editor.Buffer)) {
	t.Helper()
	h.Post(t, editor.CommandFunc(func(e *editor.Editor) { f(e.Top()) }))
}

func topText(t *testing.T, h *editor.TestHarness) (text string) {
	t.Helper()
	onLoop(t, h, func(buf *editor.Buffer) { text = buf.Text() })
	return text
}
