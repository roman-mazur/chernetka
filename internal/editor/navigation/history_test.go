package navigation

import (
	"slices"
	"testing"

	"rmazur.io/chernetka/internal/content"
)

func TestHistory(t *testing.T) {
	var h History

	confirmNoHistory := func() {
		if l := h.MovePrev(Nowhere); l != Nowhere {
			t.Fatalf("h.Prev() = %s, want Nowhere", l)
		}
		if l := h.MoveNext(Nowhere); l != Nowhere {
			t.Fatalf("h.Next() = %s, want Nowhere", l)
		}
	}
	confirmNoHistory()

	t.Log("initialize with A")
	a := Location{Path: "a.txt"}
	h.RecordJump(Nowhere, a)
	confirmNoHistory()

	t.Log("go to B")
	b := Location{Path: "b.txt"}
	h.RecordJump(a, b)
	if l := h.MovePrev(b); l != a {
		t.Errorf("h.Prev() = %s, want %s", l, a)
	}

	t.Log("restore and go to B with pos")
	h.RecordJump(a, b)
	bWithPos := Location{Path: "b.txt", Position: content.Position{Line: 3, Col: 2}}
	h.RecordJump(b, bWithPos)
	if l := h.MovePrev(bWithPos); l != b {
		t.Errorf("h.Prev() = %s, want %s", l, b)
	}
	if l := h.MoveNext(b); l != bWithPos {
		t.Errorf("h.Next() = %s, want %s", l, bWithPos)
	}

	t.Log("back to A and go to C")
	h.MovePrev(bWithPos)
	h.MovePrev(b)
	c := Location{Path: "c.txt"}
	h.RecordJump(a, c)
	if l := h.MovePrev(c); l != a {
		t.Errorf("h.Prev() = %s, want %s", l, a)
	}
	if l := h.MoveNext(a); l != c {
		t.Errorf("h.Next() = %s, want %s", l, c)
	}
	if l := h.MoveNext(c); l != Nowhere {
		t.Errorf("h.Next() = %s, want Nowhere", l)
	}

	t.Log("recording the same navigation")
	h.RecordJump(a, c)
	if l := h.MovePrev(c); l != a {
		t.Errorf("h.Prev() = %s, want %s", l, a)
	}
	if l := h.MoveNext(a); l != c {
		t.Errorf("h.Next() = %s, want %s", l, c)
	}

	t.Log("disable recording, go to B")
	h.EnableRecording(false)
	h.RecordJump(c, b)
	if l := h.MovePrev(c); l != a {
		t.Errorf("h.Prev() = %s, want %s", l, a)
	}
	if l := h.MoveNext(a); l != c {
		t.Errorf("h.Next() = %s, want %s", l, c)
	}
}

func TestHistory_Records(t *testing.T) {
	loc := func(path string, line int) Location {
		return Location{Path: path, Position: content.Position{Line: line}}
	}
	a, b, c := loc("a.txt", 0), loc("b.txt", 0), loc("c.txt", 0)

	for _, tc := range []struct {
		name    string
		steps   func(h *History)
		want    []Location
		wantPos int
	}{
		{
			name:    "first jump records both ends",
			steps:   func(h *History) { h.RecordJump(a, b) },
			want:    []Location{a, b},
			wantPos: 2,
		},
		{
			name: "departure updates the position in the same file",
			steps: func(h *History) {
				h.RecordJump(Nowhere, a)
				h.RecordJump(loc("a.txt", 50), b)
			},
			want:    []Location{loc("a.txt", 50), b},
			wantPos: 2,
		},
		{
			name: "departure from another file is inserted",
			steps: func(h *History) {
				h.RecordJump(Nowhere, a)
				h.RecordJump(loc("b.txt", 5), c)
			},
			want:    []Location{a, loc("b.txt", 5), c},
			wantPos: 3,
		},
		{
			name: "jump within the same file",
			steps: func(h *History) {
				h.RecordJump(Nowhere, a)
				h.RecordJump(loc("a.txt", 50), loc("a.txt", 10))
			},
			want:    []Location{loc("a.txt", 50), loc("a.txt", 10)},
			wantPos: 2,
		},
		{
			name: "jump to the departure location is ignored",
			steps: func(h *History) {
				h.RecordJump(Nowhere, a)
				h.RecordJump(loc("a.txt", 50), loc("a.txt", 50))
			},
			want:    []Location{loc("a.txt", 50)},
			wantPos: 1,
		},
		{
			name: "jump to nowhere only records the departure",
			steps: func(h *History) {
				h.RecordJump(Nowhere, a)
				h.RecordJump(loc("a.txt", 5), Nowhere)
			},
			want:    []Location{loc("a.txt", 5)},
			wantPos: 1,
		},
		{
			name: "jump after going back drops the forward history",
			steps: func(h *History) {
				h.RecordJump(a, b)
				h.RecordJump(b, c)
				h.MovePrev(c)
				h.MovePrev(b)
				h.RecordJump(loc("a.txt", 7), loc("d.txt", 0))
			},
			want:    []Location{loc("a.txt", 7), loc("d.txt", 0)},
			wantPos: 2,
		},
		{
			name: "ignored jump after going back keeps the forward history",
			steps: func(h *History) {
				h.RecordJump(a, b)
				h.MovePrev(b)
				h.RecordJump(loc("a.txt", 7), loc("a.txt", 7))
			},
			want:    []Location{loc("a.txt", 7), b},
			wantPos: 1,
		},
		{
			name: "moving back saves the position being left",
			steps: func(h *History) {
				h.RecordJump(a, b)
				h.MovePrev(loc("b.txt", 20))
			},
			want:    []Location{a, loc("b.txt", 20)},
			wantPos: 1,
		},
		{
			name: "moving forward saves the position being left",
			steps: func(h *History) {
				h.RecordJump(a, b)
				h.MovePrev(b)
				h.MoveNext(loc("a.txt", 9))
			},
			want:    []Location{loc("a.txt", 9), b},
			wantPos: 2,
		},
		{
			name: "moving from another file keeps the records",
			steps: func(h *History) {
				h.RecordJump(a, b)
				h.MovePrev(loc("c.txt", 3))
				h.MoveNext(Nowhere)
			},
			want:    []Location{a, b},
			wantPos: 2,
		},
		{
			name: "moving past the ends keeps the records",
			steps: func(h *History) {
				h.RecordJump(a, b)
				h.MoveNext(loc("b.txt", 4))
				h.MovePrev(b)
				h.MovePrev(loc("a.txt", 4))
			},
			want:    []Location{a, b},
			wantPos: 1,
		},
		{
			name: "disabled recording ignores the departure",
			steps: func(h *History) {
				h.RecordJump(Nowhere, a)
				h.EnableRecording(false)
				h.RecordJump(loc("a.txt", 5), b)
				h.EnableRecording(true)
			},
			want:    []Location{a},
			wantPos: 1,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var h History
			tc.steps(&h)
			if !slices.Equal(h.records, tc.want) || h.pos != tc.wantPos {
				t.Errorf("records %v at %d, want %v at %d", h.records, h.pos, tc.want, tc.wantPos)
			}
		})
	}
}
