package navigation

import (
	"fmt"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/debugflags"
)

// History allows recording and navigating the jumps between different locations
// in the current editor session.
type History struct {
	records []Location

	pos     int // current history element position
	ringEnd int // the end of the ring buffer

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

func (l *Location) updateIfMatch(from Location) bool {
	if l == nil {
		return false
	}
	if l.Path == from.Path {
		l.Position = from.Position
		return true
	}
	return false
}

func (h *History) RecordJump(from, to Location) {
	if h.disabled {
		return
	}

	if from != Nowhere {
		if l, _ := h.last(); !l.updateIfMatch(from) {
			h.record(from)
		}
	}

	if to == Nowhere {
		return
	}
	if last, _ := h.last(); last != nil && *last == to {
		return
	}

	h.record(to)
}

func (h *History) last() (*Location, int) {
	l := len(h.records)
	if l == 0 {
		return nil, -1
	}
	return &h.records[h.pos], h.pos
}

func (h *History) record(l Location) {
	if h.records == nil {
		h.records = make([]Location, 0, maxSize)
	}
	last, pos := h.last()
	if last != nil {
		h.rotateRecords(pos)
	}
	h.ringPush(l)
}

func (h *History) ringPush(l Location) {
	if len(h.records) < maxSize {
		h.records = append(h.records, l)
	} else {
		h.records[h.ringEnd] = l
	}
	h.pos = h.ringEnd
	h.ringEnd = (h.ringEnd + 1) % maxSize
}

// rotateRecords puts the element the specified position at the end of ring shifting
// elements after it.
func (h *History) rotateRecords(pos int) {
	data := h.records[pos]
	if pos >= h.ringEnd {
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

func (h *History) canMove(forward bool) bool {
	l, pos := h.last()
	if l == nil {
		return false // no history, no move
	}
	if forward {
		next := (pos + 1) % maxSize
		return next != h.ringEnd && next < len(h.records)
	}
	next := (pos - 1 + maxSize) % maxSize
	return pos != h.ringEnd && next < len(h.records)
}

func (h *History) MovePrev(from Location) Location {
	if !h.canMove(false) {
		return Nowhere
	}
	l, _ := h.last()
	l.updateIfMatch(from)
	h.pos = (h.pos - 1 + maxSize) % maxSize
	return h.records[h.pos]
}

func (h *History) MoveNext(from Location) Location {
	if !h.canMove(true) {
		return Nowhere
	}
	l, _ := h.last()
	l.updateIfMatch(from)
	h.pos = (h.pos + 1) % maxSize
	return h.records[h.pos]
}

func (h *History) EnableRecording(enabled bool) {
	h.disabled = !enabled
}
