package change

import (
	"bytes"
	"fmt"
	"math/rand/v2"
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
	"rmazur.io/chernetka/internal/content"
)

func TestHistory_Track(t *testing.T) {
	var h History

	m := h.Track(content.Empty().(content.Mutable), false)
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

func TestHistory_ConsumeTrackedChanges(t *testing.T) {
	var h History

	m := h.Track(content.Empty().(content.Mutable), true)
	ur := m.(content.UndoRedo)

	m.Insert(0, content.TextLine("line 1"))
	m.Insert(1, content.TextLine("line 2"))
	ur.SplitChangeGroup()

	changes := h.ConsumeTrackedChanges(m)
	if len(changes) != 2 {
		t.Error("bad changes list after 2 inserts:", changes)
	}

	m.Update(1, content.TextLine("update line 2"))
	m.Delete(0, 1)
	ur.SplitChangeGroup()

	changes = h.ConsumeTrackedChanges(m)
	if len(changes) != 2 {
		t.Error("bad changes list after update/delete", changes)
	}

	ur.Undo()
	changes = h.ConsumeTrackedChanges(m)
	if len(changes) != 2 {
		t.Errorf("bad changes list after undo, len=%d: %v", len(changes), changes)
	}

	ur.Undo()
	ur.Redo()
	changes = h.ConsumeTrackedChanges(m)
	if len(changes) != 4 {
		t.Errorf("bad changes list after 2*undo + redo, len=%d: %v", len(changes), changes)
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
			m := h.Track(&doc, true)
			for _, u := range tc.updates {
				m.Update(0, content.TextLine(u))
			}
			changes := slices.Collect[Item](h[0].iterateChanges)
			if len(changes) != 1 {
				t.Fatalf("got %d changes, want the updates merged into 1", len(changes))
			}
			if recentChanges := h.ConsumeTrackedChanges(m); len(changes) >= len(recentChanges) {
				t.Errorf("len(changes) < len(h.ConsumeTrackedChanges(&doc)): %d < %d / %v",
					len(changes), len(recentChanges), recentChanges)
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

// TestHistory_UndoRedoCursor checks the cursor positions returned by Undo and Redo.
// Undo puts the cursor where the first change of the group was reverted, redo puts it after
// the last change of the group was applied.
func TestHistory_UndoRedoCursor(t *testing.T) {
	type group struct {
		edit       func(m content.Mutable)
		undo, redo content.Position
	}
	pos := func(line, col int) content.Position { return content.Position{Line: line, Col: col} }

	for _, tc := range []struct {
		name   string
		doc    []string
		groups []group
	}{
		{
			name: "insert",
			doc:  []string{"a"},
			groups: []group{{
				edit: func(m content.Mutable) { m.Insert(1, content.TextLine("x"), content.TextLine("yy")) },
				undo: pos(1, 0), // start of the removed lines
				redo: pos(2, 2), // end of the inserted lines
			}},
		},
		{
			name: "delete",
			doc:  []string{"a", "bb", "ccc", "d"},
			groups: []group{{
				edit: func(m content.Mutable) { m.Delete(1, 3) },
				undo: pos(2, 3), // end of the restored lines
				redo: pos(1, 0), // where the lines were deleted
			}},
		},
		{
			name: "update",
			doc:  []string{"a", "hello world"},
			groups: []group{{
				edit: func(m content.Mutable) { m.Update(1, content.TextLine("hello there world")) },
				undo: pos(1, 6),  // start of the changed part
				redo: pos(1, 12), // end of the changed part in the new line
			}},
		},
		{
			name: "group of changes",
			doc:  []string{"a"},
			groups: []group{{
				edit: func(m content.Mutable) {
					m.Insert(1, content.TextLine("x"))
					m.Update(0, content.TextLine("ab"))
				},
				undo: pos(1, 0), // reverting the insert is the last undo step
				redo: pos(0, 2), // applying the update is the last redo step
			}},
		},
		{
			// The editor splits a line by truncating it and inserting the rest below. Undo must put
			// the cursor at the split point, not after the restored text.
			name: "line split",
			doc:  []string{"abcd"},
			groups: []group{{
				edit: func(m content.Mutable) {
					m.Update(0, content.TextLine("ab"))
					m.Insert(1, content.TextLine("cd"))
				},
				undo: pos(0, 2),
				redo: pos(1, 2),
			}},
		},
		{
			name: "multiple groups",
			doc:  []string{"a", "b"},
			groups: []group{
				{
					edit: func(m content.Mutable) { m.Update(0, content.TextLine("a1")) },
					undo: pos(0, 1),
					redo: pos(0, 2),
				},
				{
					edit: func(m content.Mutable) { m.Delete(1, 2) },
					undo: pos(1, 1),
					redo: pos(1, 0),
				},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var doc content.FullText
			for _, l := range tc.doc {
				doc = append(doc, content.TextLine(l))
			}
			var h History
			m := h.Track(&doc, false)
			ur := m.(content.UndoRedo)
			for _, g := range tc.groups {
				g.edit(m)
				ur.SplitChangeGroup()
			}

			for i, g := range slices.Backward(tc.groups) {
				if !ur.HasUndo() {
					t.Fatalf("no undo for group %d", i)
				}
				if got := ur.Undo(); got != g.undo {
					t.Errorf("group %d: undo cursor = %v, want %v", i, got, g.undo)
				}
			}
			if ur.HasUndo() {
				t.Error("unexpected undo after reverting all the groups")
			}
			var got []string
			for _, l := range doc {
				got = append(got, l.String())
			}
			if diff := cmp.Diff(tc.doc, got); diff != "" {
				t.Errorf("doc after undo mismatch (-want +got):\n%s", diff)
			}

			for i, g := range tc.groups {
				if !ur.HasRedo() {
					t.Fatalf("no redo for group %d", i)
				}
				if got := ur.Redo(); got != g.redo {
					t.Errorf("group %d: redo cursor = %v, want %v", i, got, g.redo)
				}
			}
			if ur.HasRedo() {
				t.Error("unexpected redo after applying all the groups")
			}
		})
	}
}

// TestHistory_ConsumeTrackedChanges_Replay makes random edits, undoes and redoes them, and checks
// that the changes consumed turn the document of the previous consumption into the current one.
func TestHistory_ConsumeTrackedChanges_Replay(t *testing.T) {
	var h History
	doc := content.FullText{content.TextLine("one"), content.TextLine("two")}
	m := h.Track(&doc, true)
	ur := m.(content.UndoRedo)

	rnd := rand.New(rand.NewPCG(1, 2))
	consumed := content.Copy(&doc).(content.Mutable)
	for i := range 2000 {
		lines := doc.Lines()
		switch op := rnd.IntN(8); {
		case op < 3 && len(lines) > 0: // Typing, merged into a change of the history.
			ln := rnd.IntN(len(lines))
			text := lines[ln].String()
			col := rnd.IntN(len(text) + 1)
			m.Update(ln, content.TextLine(text[:col]+"x"+text[col:]))
		case op == 3 && len(lines) > 0:
			start := rnd.IntN(len(lines))
			m.Delete(start, min(len(lines), start+1+rnd.IntN(2)))
		case op == 4:
			m.Insert(rnd.IntN(len(lines)+1), content.TextLine(fmt.Sprint(i)), content.TextLine(""))
		case op == 5 && ur.HasUndo():
			ur.Undo()
		case op == 6 && ur.HasRedo():
			ur.Redo()
		case op == 7:
			ur.SplitChangeGroup()
		}

		if rnd.IntN(3) == 0 {
			changes := h.ConsumeTrackedChanges(m)
			for _, ch := range changes {
				ch.apply(consumed)
			}
			if got, want := content.Sprint(consumed, 0, consumed.Len()), content.Sprint(&doc, 0, doc.Len()); got != want {
				t.Fatalf("edit %d: the changes %+v make\n%q\nwant\n%q", i, changes, got, want)
			}
		}
	}
}
