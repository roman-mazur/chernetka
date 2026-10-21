package editor

import (
	"strings"
	"testing"
	"testing/synctest"
	"time"
)

func TestEditor_UndoRedo(t *testing.T) {
	const (
		undo = "\x1a"        // Ctrl+Z
		redo = "\x1b[122;6u" // Ctrl+Shift+Z with the kitty keyboard protocol
		// pause ends the group of changes undone at once.
		pause = "pause"
	)

	for _, tc := range []struct {
		name   string
		text   string
		inputs []string
		want   string
	}{
		{
			name:   "nothing to undo",
			text:   "one",
			inputs: []string{undo, redo},
			want:   "one",
		},
		{
			name:   "kitty undo",
			text:   "one",
			inputs: []string{"Atwo\x1b[27u", "\x1b[122;5u"},
			want:   "one",
		},
		{
			name:   "undo typing",
			text:   "one",
			inputs: []string{"Atwo\x1b", undo},
			want:   "one",
		},
		{
			name:   "undo in insert mode",
			text:   "one",
			inputs: []string{"Atwo", undo, "x"},
			want:   "onex",
		},
		{
			name:   "undo the groups one by one",
			text:   "one",
			inputs: []string{"A two\x1b", pause, "A three\x1b", undo},
			want:   "one two",
		},
		{
			name:   "undo past the first group",
			text:   "one",
			inputs: []string{"A two\x1b", pause, "A three\x1b", undo, undo, undo},
			want:   "one",
		},
		{
			name:   "redo",
			text:   "one",
			inputs: []string{"A two\x1b", pause, "A three\x1b", undo, undo, redo},
			want:   "one two",
		},
		{
			name:   "redo all",
			text:   "one",
			inputs: []string{"A two\x1b", pause, "A three\x1b", undo, undo, redo, redo, redo},
			want:   "one two three",
		},
		{
			name:   "edit after undo drops the redo",
			text:   "one",
			inputs: []string{"A two\x1b", pause, "A three\x1b", undo, pause, "A four\x1b", redo},
			want:   "one two four",
		},
		{
			name:   "leaving insert mode ends the group",
			text:   "one",
			inputs: []string{"A two\x1b", "A three\x1b", undo},
			want:   "one two",
		},
		{
			name:   "entering insert mode ends the group",
			text:   "one",
			inputs: []string{"x", "Atwo\x1b", undo},
			want:   "ne",
		},
		{
			name:   "new line is undone with its text",
			text:   "one\ntwo",
			inputs: []string{"x", "otext\x1b", undo},
			want:   "ne\ntwo",
		},
		{
			name:   "command line in insert mode ends the group",
			text:   "one",
			inputs: []string{"Atwo", "\x0c\x1b", "three\x1b", undo},
			want:   "onetwo",
		},
		{
			name:   "split and join lines",
			text:   "one two\nthree",
			inputs: []string{"fi\r\x1b", pause, "jA\x7f\x7f\x1b", undo, undo},
			want:   "one two\nthree",
		},
		{
			name:   "deleted lines",
			text:   "one\ntwo\nthree",
			inputs: []string{"i\x1b[1;2B\x1b[1;2B\x7f\x1b", undo},
			want:   "one\ntwo\nthree",
		},
		{
			name:   "paste",
			text:   "one",
			inputs: []string{"A\x1b[200~\npasted\nlines\x1b[201~\x1b", undo},
			want:   "one",
		},
		{
			name:   "x and o",
			text:   "one\ntwo",
			inputs: []string{"x", pause, "otext\x1b", undo},
			want:   "ne\ntwo",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				e := &Editor{cmdChannel: make(chan Command, 16)}
				if err := e.OpenReader("", strings.NewReader(tc.text)); err != nil {
					t.Fatal(err)
				}
				b := e.Top()
				for _, in := range tc.inputs {
					switch in {
					case pause:
						time.Sleep(time.Minute)
					default:
						sendChunks(t, e, in)
					}
					b.clampCursor(e.rPrefs.TabSize) // Like the render between the inputs.
				}
				if got := b.Text(); got != tc.want {
					t.Errorf("text %q, want %q", got, tc.want)
				}
			})
		})
	}
}

func TestEditor_UndoNotifiesExtensions(t *testing.T) {
	e := &Editor{cmdChannel: make(chan Command, 16)}
	var edits countingEditExt
	e.Extend(&edits)
	if err := e.OpenReader("", strings.NewReader("one")); err != nil {
		t.Fatal(err)
	}
	sendChunks(t, e, "x")
	before := edits
	sendChunks(t, e, "\x1a")
	if edits != before+1 {
		t.Errorf("extensions notified %d times after the undo, want 1", edits-before)
	}
	if !e.Top().dirty {
		t.Error("the buffer is not dirty after the undo")
	}
}

type countingEditExt int

func (c *countingEditExt) ID() string                                   { return "counting edits" }
func (c *countingEditExt) AfterEdit(Sender, *Buffer)                    { *c++ }
func (c *countingEditExt) MakeBufferData(Sender, *Buffer) BufferExtData { return nil }
