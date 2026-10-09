package navigation

import (
	"fmt"
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

	fillRingBuffer := func(h *History) {
		h.RecordJump(Nowhere, loc("l0.txt", 0))
		for i := range maxSize - 1 {
			h.RecordJump(Nowhere, loc(fmt.Sprintf("l%d.txt", i+1), i+1))
		}
	}

	for _, tc := range []struct {
		name    string
		steps   func(h *History)
		want    []Location
		wantF   func() []Location
		wantPos int
	}{
		{
			name:    "first jump records both ends",
			steps:   func(h *History) { h.RecordJump(a, b) },
			want:    []Location{a, b},
			wantPos: 1,
		},
		{
			name: "departure updates the position in the same file",
			steps: func(h *History) {
				h.RecordJump(Nowhere, a)
				h.RecordJump(loc("a.txt", 50), b)
			},
			want:    []Location{loc("a.txt", 50), b},
			wantPos: 1,
		},
		{
			name: "departure from another file is inserted",
			steps: func(h *History) {
				h.RecordJump(Nowhere, a)
				h.RecordJump(loc("b.txt", 5), c)
			},
			want:    []Location{a, loc("b.txt", 5), c},
			wantPos: 2,
		},
		{
			name: "jump within the same file",
			steps: func(h *History) {
				h.RecordJump(Nowhere, a)
				h.RecordJump(loc("a.txt", 50), loc("a.txt", 10))
			},
			want:    []Location{loc("a.txt", 50), loc("a.txt", 10)},
			wantPos: 1,
		},
		{
			name: "jump to the departure location is ignored",
			steps: func(h *History) {
				h.RecordJump(Nowhere, a)
				h.RecordJump(loc("a.txt", 50), loc("a.txt", 50))
			},
			want:    []Location{loc("a.txt", 50)},
			wantPos: 0,
		},
		{
			name: "jump to nowhere only records the departure",
			steps: func(h *History) {
				h.RecordJump(Nowhere, a)
				h.RecordJump(loc("a.txt", 5), Nowhere)
			},
			want:    []Location{loc("a.txt", 5)},
			wantPos: 0,
		},
		{
			name: "jump after going back retains the forward history",
			steps: func(h *History) {
				h.RecordJump(a, b)
				h.RecordJump(b, c)
				h.MovePrev(c)
				h.MovePrev(b)
				h.RecordJump(loc("a.txt", 7), loc("d.txt", 0))
			},
			want:    []Location{b, c, loc("a.txt", 7), loc("d.txt", 0)},
			wantPos: 3,
		},
		{
			name: "ignored jump after going back keeps the forward history",
			steps: func(h *History) {
				h.RecordJump(a, b)
				h.MovePrev(b)
				h.RecordJump(loc("a.txt", 7), loc("a.txt", 7))
			},
			want:    []Location{loc("a.txt", 7), b},
			wantPos: 0,
		},
		{
			name: "moving back saves the position being left",
			steps: func(h *History) {
				h.RecordJump(a, b)
				h.MovePrev(loc("b.txt", 20))
			},
			want:    []Location{a, loc("b.txt", 20)},
			wantPos: 0,
		},
		{
			name: "moving forward saves the position being left",
			steps: func(h *History) {
				h.RecordJump(a, b)
				h.MovePrev(b)
				h.MoveNext(loc("a.txt", 9))
			},
			want:    []Location{loc("a.txt", 9), b},
			wantPos: 1,
		},
		{
			name: "moving from another file keeps the records",
			steps: func(h *History) {
				h.RecordJump(a, b)
				h.MovePrev(loc("c.txt", 3))
				h.MoveNext(Nowhere)
			},
			want:    []Location{a, b},
			wantPos: 1,
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
			wantPos: 0,
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
			wantPos: 0,
		},
		{
			name: "limit two maxSize records",
			steps: func(h *History) {
				fillRingBuffer(h)
				h.RecordJump(loc(fmt.Sprintf("l%d.txt", maxSize-1), maxSize-1), a)
				h.RecordJump(a, b)
				h.RecordJump(b, c)
			},
			wantF: func() []Location {
				res := make([]Location, maxSize)
				res[0], res[1], res[2] = a, b, c
				for i := 3; i < len(res); i++ {
					res[i] = loc(fmt.Sprintf("l%d.txt", i), i)
				}
				return res[:]
			},
			wantPos: 2,
		},
		{
			name: "rotate ring buffer",
			steps: func(h *History) {
				fillRingBuffer(h)
				l99 := loc(fmt.Sprintf("l%d.txt", maxSize-1), maxSize-1)
				l98 := loc(fmt.Sprintf("l%d.txt", maxSize-2), maxSize-2)
				h.RecordJump(l99, a)
				h.RecordJump(a, b)
				h.MovePrev(b)
				h.MovePrev(a)
				h.MovePrev(l99)
				h.RecordJump(l98, c)
			},
			wantF: func() []Location {
				l99 := loc(fmt.Sprintf("l%d.txt", maxSize-1), maxSize-1)
				l98 := loc(fmt.Sprintf("l%d.txt", maxSize-2), maxSize-2)

				res := make([]Location, maxSize)
				res[0], res[1], res[2] = b, l98, c
				for i := 3; i < len(res)-2; i++ {
					res[i] = loc(fmt.Sprintf("l%d.txt", i), i)
				}
				res[maxSize-2], res[maxSize-1] = l99, a
				return res[:]
			},
			wantPos: 2,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var h History
			tc.steps(&h)
			if tc.wantF != nil {
				tc.want = tc.wantF()
			}
			if !slices.Equal(h.records, tc.want) || h.pos != tc.wantPos {
				t.Errorf("records %v at %d, want %v at %d", h.records, h.pos, tc.want, tc.wantPos)
			}
		})
	}
}

// setMaxSize changes the ring buffer size for the duration of the test.
func setMaxSize(t testing.TB, size int) {
	old := maxSize
	maxSize = size
	t.Cleanup(func() { maxSize = old })
}

func TestHistory_FullRing(t *testing.T) {
	setMaxSize(t, 3)
	loc := func(i int) Location { return Location{Path: fmt.Sprintf("l%d.txt", i)} }

	// Cover every ring end position after the buffer becomes full.
	for jumps := maxSize - 1; jumps <= 3*maxSize; jumps++ {
		t.Run(fmt.Sprintf("%d jumps", jumps), func(t *testing.T) {
			var h History
			h.RecordJump(Nowhere, loc(0))
			for i := range jumps {
				h.RecordJump(loc(i), loc(i+1))
			}

			// All the kept records except the current one are reachable going back,
			// the oldest one being the last.
			var want []Location
			for i := jumps - 1; i >= max(0, jumps+1-maxSize); i-- {
				want = append(want, loc(i))
			}
			var got []Location
			for range 2 * maxSize {
				l := h.MovePrev(Nowhere)
				if l == Nowhere {
					break
				}
				got = append(got, l)
			}
			if !slices.Equal(got, want) {
				t.Fatalf("moving back reaches %v, want %v", got, want)
			}

			// And then going forward returns to the current location.
			got = got[:0]
			for range 2 * maxSize {
				l := h.MoveNext(Nowhere)
				if l == Nowhere {
					break
				}
				got = append(got, l)
			}
			want = want[:0]
			for i := max(0, jumps+2-maxSize); i <= jumps; i++ {
				want = append(want, loc(i))
			}
			if !slices.Equal(got, want) {
				t.Fatalf("moving forward reaches %v, want %v", got, want)
			}
		})
	}
}

func TestHistory_JumpFromOldest(t *testing.T) {
	setMaxSize(t, 3)
	loc := func(i int) Location { return Location{Path: fmt.Sprintf("l%d.txt", i)} }

	var h History
	h.RecordJump(Nowhere, loc(0))
	for i := range 4 {
		h.RecordJump(loc(i), loc(i+1))
	}
	// Records are l3, l4, l2 with the ring end at the last one.
	h.MovePrev(Nowhere)
	if l := h.MovePrev(Nowhere); l != loc(2) {
		t.Fatalf("h.MovePrev() = %s, want %s", l, loc(2))
	}

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("jump from the oldest location panics: %v", r)
		}
	}()
	h.RecordJump(loc(2), loc(5))
	if l := h.MovePrev(loc(5)); l != loc(2) {
		t.Errorf("h.MovePrev() = %s, want %s", l, loc(2))
	}
}

// FuzzHistory interprets every input byte as a step: the lower 2 bits select an action
// (a jump, moving back or forward), and the higher bits the locations involved.
func FuzzHistory(f *testing.F) {
	setMaxSize(f, 5)
	f.Add([]byte{0b0001_0000, 0b0010_0100, 1, 1, 2, 0b0011_1000})
	f.Add([]byte{0b0001_0000, 0b0010_0100, 0b0011_1000, 0b0000_1100, 0b0001_0000, 0b0010_0100, 1, 1, 1, 1, 1, 0b0011_1000})

	loc := func(b byte) Location { return Location{Path: fmt.Sprintf("%d.txt", b&3)} }
	f.Fuzz(func(t *testing.T, steps []byte) {
		var h History
		// Moving in one direction must stop within the ring buffer size.
		checkEnds := func(step int) {
			t.Helper()
			saved := h
			saved.records = slices.Clone(h.records)
			for _, move := range []func(Location) Location{h.MovePrev, h.MoveNext} {
				n := 0
				for move(Nowhere) != Nowhere {
					if n++; n >= maxSize {
						t.Fatalf("step %d: no end after %d moves, records %v", step, n, h.records)
					}
				}
			}
			h = saved
		}

		for i, b := range steps {
			from, to := loc(b>>2), loc(b>>4)
			switch b & 3 {
			case 0, 3:
				h.RecordJump(from, to)
			case 1:
				h.MovePrev(from)
			case 2:
				h.MoveNext(from)
			}
			checkEnds(i)
		}
	})
}
