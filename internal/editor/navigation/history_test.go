package navigation

import (
	"testing"

	"rmazur.io/chernetka/internal/content"
)

func TestHistory(t *testing.T) {
	var h History

	confirmNoHistory := func() {
		if l := h.MovePrev(); l != Nowhere {
			t.Fatalf("h.Prev() = %s, want Nowhere", l)
		}
		if l := h.MoveNext(); l != Nowhere {
			t.Fatalf("h.Next() = %s, want Nowhere", l)
		}
	}
	confirmNoHistory()

	t.Log("initialize with A")
	a := Location{Path: "a.txt"}
	h.RecordJump(a)
	confirmNoHistory()

	t.Log("go to B")
	b := Location{Path: "b.txt"}
	h.RecordJump(b)
	if l := h.MovePrev(); l != a {
		t.Errorf("h.Prev() = %s, want %s", l, a)
	}

	t.Log("restore and go to B with pos")
	h.RecordJump(b)
	bWithPos := Location{Path: "b.txt", Position: content.Position{Line: 3, Col: 2}}
	h.RecordJump(bWithPos)
	if l := h.MovePrev(); l != b {
		t.Errorf("h.Prev() = %s, want %s", l, b)
	}
	if l := h.MoveNext(); l != bWithPos {
		t.Errorf("h.Next() = %s, want %s", l, bWithPos)
	}

	t.Log("back to A and go to C")
	h.MovePrev()
	h.MovePrev()
	c := Location{Path: "c.txt"}
	h.RecordJump(c)
	if l := h.MovePrev(); l != a {
		t.Errorf("h.Prev() = %s, want %s", l, a)
	}
	if l := h.MoveNext(); l != c {
		t.Errorf("h.Next() = %s, want %s", l, c)
	}
	if l := h.MoveNext(); l != Nowhere {
		t.Errorf("h.Next() = %s, want Nowhere", l)
	}

	t.Log("recording the same navigation")
	h.RecordJump(c)
	if l := h.MovePrev(); l != a {
		t.Errorf("h.Prev() = %s, want %s", l, a)
	}
	if l := h.MoveNext(); l != c {
		t.Errorf("h.Next() = %s, want %s", l, c)
	}

	t.Log("disable recording, go to B")
	h.EnableRecording(false)
	h.RecordJump(b)
	if l := h.MovePrev(); l != a {
		t.Errorf("h.Prev() = %s, want %s", l, a)
	}
	if l := h.MoveNext(); l != c {
		t.Errorf("h.Next() = %s, want %s", l, c)
	}
}
