package navigation

import (
	"fmt"
	"slices"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/debugflags"
)

// History allows recording and navigating the jumps between different locations
// in the current editor session.
type History struct {
	records []Location

	insertPos int // next insert position in the ring buffer
	ringEnd   int // the end of the ring buffer

	disabled bool
}

var Nowhere = Location{}

var maxSize = debugflags.ValueInt("HISTORY_RING_SIZE", 100)

// Location points to a particular place in a project.
type Location struct {
	Path     string
	Position content.Position
}

func (l Location) String() string {
	if l == Nowhere {
		return "<nowhere>"
	}
	return fmt.Sprintf("%s@%d:%d", l.Path, l.Position.Line, l.Position.Col)
}

func (h *History) RecordJump(from, to Location) {
	if h.disabled {
		return
	}

	if from != Nowhere {
		if l, _ := h.last(); l != nil && l.Path == from.Path {
			l.Position = from.Position
		} else {
			h.record(from)
		}
	}

	if to == Nowhere {
		return
	}

	last := Nowhere
	if h.insertPos > 0 {
		last = h.records[h.insertPos-1]
	}
	if last == to {
		return
	}
	h.record(to)
}

func (h *History) last() (*Location, int) {
	l := len(h.records)
	if l == 0 {
		return nil, -1
	}
	if l < maxSize && h.insertPos == 0 {
		return nil, -1
	}
	i := (h.insertPos - 1 + l) % l
	if h.insertPos > h.ringEnd && i <= h.ringEnd {
		return nil, -1
	}
	return &h.records[i], i
}

func (h *History) record(l Location) {
	if h.records == nil {
		h.records = make([]Location, 0, maxSize)
	}
	if len(h.records) < maxSize {
		// keep growing
		if prev, pos := h.last(); prev != nil {
			prevLoc := *prev
			h.records = append(slices.Delete(h.records, pos, pos+1), prevLoc)
		}
		h.records = append(h.records, l)
		h.advanceRing()
		return
	}

	last, pos := h.last()
	if last == nil {
		panic("records with max size and no last")
	}
	h.rotateRecords(pos)
	h.records[h.ringEnd] = l
	h.advanceRing()
}

func (h *History) advanceRing() {
	h.ringEnd = (h.ringEnd + 1) % maxSize
	h.insertPos = h.ringEnd
}

// rotateRecords puts the element the specified position at the end of ring shifting
// elements after it.
func (h *History) rotateRecords(pos int) {
	data := h.records[pos]
	if pos == h.ringEnd {
		panic("pos == h.ringEnd")
	}
	if pos > h.ringEnd {
		if pos < len(h.records)-1 {
			copy(h.records[pos:], h.records[pos+1:])
		}
		if h.ringEnd > 0 {
			h.records[len(h.records)-1] = h.records[0]
		} else {
			h.records[len(h.records)-1] = data
		}
		pos = 0
	}
	if pos < h.ringEnd {
		copy(h.records[pos:h.ringEnd], h.records[pos+1:h.ringEnd])
		h.records[h.ringEnd-1] = data
	}
}

func (h *History) MovePrev(from Location) Location {
	l, pos := h.last()
	if l == nil {
		return Nowhere
	}

	var oldInsert int
	h.insertPos, oldInsert = pos, h.insertPos
	res, _ := h.last()
	if res == nil {
		h.insertPos = oldInsert
		return Nowhere
	}

	if l.Path == from.Path {
		l.Position = from.Position
	}
	return *res
}

func (h *History) MoveNext(from Location) Location {
	if h.insertPos >= len(h.records) || h.insertPos == h.ringEnd {
		return Nowhere
	}
	if l, _ := h.last(); l != nil && l.Path == from.Path {
		l.Position = from.Position
	}
	res := h.records[h.insertPos]
	h.insertPos = (h.insertPos + 1) % maxSize
	return res
}

func (h *History) EnableRecording(enabled bool) {
	h.disabled = !enabled
}
