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

func (h *History) RecordJump(to Location) {
	if h.disabled {
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

func (h *History) MovePrev() Location {
	if h.pos > len(h.records) || h.pos <= 1 {
		return Nowhere
	}
	h.pos--
	return h.records[h.pos-1]
}

func (h *History) MoveNext() Location {
	if h.pos >= len(h.records) || h.pos == 0 {
		return Nowhere
	}
	res := h.records[h.pos]
	h.pos++
	return res
}

func (h *History) EnableRecording(enabled bool) {
	h.disabled = !enabled
}
