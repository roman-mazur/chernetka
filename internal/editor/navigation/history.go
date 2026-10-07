package navigation

import (
	"fmt"

	"rmazur.io/chernetka/internal/content"
)

// History allows recording and navigating the jumps between different locations
// in the current editor session.
type History struct {
	records []Location
	pos     int

	disabled bool
}

var Nowhere = Location{}

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
		if h.pos > 0 && h.records[h.pos-1].Path == from.Path {
			h.records[h.pos-1].Position = from.Position
		} else {
			h.records = append(h.records[:h.pos], from)
			h.pos++
		}
	}

	if to == Nowhere {
		return
	}

	last := Nowhere
	if h.pos > 0 {
		last = h.records[h.pos-1]
	}
	if last == to {
		return
	}

	h.records = append(h.records[:h.pos], to)
	h.pos++
}

func (h *History) MovePrev(from Location) Location {
	if h.pos > len(h.records) || h.pos <= 1 {
		return Nowhere
	}
	h.pos--
	if h.records[h.pos].Path == from.Path {
		h.records[h.pos].Position = from.Position
	}
	return h.records[h.pos-1]
}

func (h *History) MoveNext(from Location) Location {
	if h.pos >= len(h.records) || h.pos == 0 {
		return Nowhere
	}
	if h.records[h.pos-1].Path == from.Path {
		h.records[h.pos-1].Position = from.Position
	}
	res := h.records[h.pos]
	h.pos++
	return res
}

func (h *History) EnableRecording(enabled bool) {
	h.disabled = !enabled
}
