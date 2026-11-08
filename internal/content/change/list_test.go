package change

import (
	"bytes"
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
	"rmazur.io/chernetka/internal/content"
)

func TestHistory_Track(t *testing.T) {
	var h History

	m := h.Track(content.Empty().(content.Mutable))
	m.Insert(0, content.TextLine("line 1"))
	if len(h) != 1 {
		t.Fatal("history didn't start on fist edit")
	}

	m.Insert(1, content.TextLine("line 2"))
	m.Update(1, content.TextLine("update line 2"))
	m.Delete(0, 1)
	m.Update(0, content.TextLine("update line 2 v2"))
	m.Update(0, content.TextLine("update line 2 v3")) // modifies the last change

	allChanges := slices.Collect[Item](h[0].iterateChanges)
	if editsLen := len(allChanges); editsLen != 5 {
		t.Errorf("expected history of 5 items, got %d", editsLen)
	}

	var out bytes.Buffer
	if err := content.SaveToWriter(m, &out); err != nil {
		t.Fatal(err)
	}

	var replayOut bytes.Buffer
	if err := content.SaveToWriter(h[0].Replay(), &replayOut); err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(out.String(), replayOut.String()); diff != "" {
		t.Errorf("replay mismatch (-want +got):\n%s", diff)
	}

	ur, ok := m.(content.UndoRedo)
	if !ok {
		t.Fatal("no UndoRedo support")
	}

	undoCount := 0
	for ur.HasUndo() {
		ur.Undo()
		undoCount++
	}
	if undoCount == 0 {
		t.Error("no undo actions")
	}
	if ur.Len() != 1 {
		t.Error("unexpected doc after all updates, len=", ur.Len())
	}

	redoCount := 0
	for ur.HasRedo() {
		ur.Redo()
		redoCount++
	}
	if undoCount != redoCount {
		t.Errorf("unexpected redo count %d, undo count %d", redoCount, undoCount)
	}

	var redoOut bytes.Buffer
	if err := content.SaveToWriter(ur, &redoOut); err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(out.String(), redoOut.String()); diff != "" {
		t.Errorf("after-redo mismatch (-want +got):\n%s", diff)
	}
}

// TestHistory_MergedUpdateSpan checks the span of the updates of a line merged into one change:
// it's the part of the line the updates changed together, so redo puts the cursor after it.
func TestHistory_MergedUpdateSpan(t *testing.T) {
	for _, tc := range []struct {
		name    string
		updates []string
		want    content.Span
	}{
		{
			name:    "typing before the previous insertion",
			updates: []string{"01234a56789", "01b234a56789"},
			want:    content.Span{Start: content.Position{Col: 2}, End: content.Position{Col: 7}},
		},
		{
			// The lines have nothing in common at either end. The span was 0..11 before,
			// past the end of the line.
			name:    "deleting before the previous insertion",
			updates: []string{"0123456789a", "56789a"},
			want:    content.Span{Start: content.Position{Col: 0}, End: content.Position{Col: 6}},
		},
		{
			name:    "typing after the previous insertion",
			updates: []string{"0a123456789", "0a1234b56789"},
			want:    content.Span{Start: content.Position{Col: 1}, End: content.Position{Col: 7}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var h History
			doc := content.FullText{content.TextLine("0123456789")}
			m := h.Track(&doc)
			for _, u := range tc.updates {
				m.Update(0, content.TextLine(u))
			}
			changes := slices.Collect[Item](h[0].iterateChanges)
			if len(changes) != 1 {
				t.Fatalf("got %d changes, want the updates merged into 1", len(changes))
			}
			if got := changes[0].Span; got != tc.want {
				t.Errorf("span = %v, want %v", got, tc.want)
			}

			ur := m.(content.UndoRedo)
			ur.Undo()
			if got := doc[0].String(); got != "0123456789" {
				t.Errorf("after undo, the line is %q", got)
			}
			if got, want := ur.Redo(), tc.want.End; got != want {
				t.Errorf("redo cursor = %v, want %v", got, want)
			}
			if got, want := doc[0].String(), tc.updates[len(tc.updates)-1]; got != want {
				t.Errorf("after redo, the line is %q, want %q", got, want)
			}
		})
	}
}
