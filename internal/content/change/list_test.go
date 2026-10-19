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

	allChanges := slices.Collect((&h[0]).iterateChanges)
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
