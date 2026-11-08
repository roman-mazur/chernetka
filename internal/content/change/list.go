package change

import (
	"fmt"
	"slices"
	"time"

	"rmazur.io/chernetka/internal/content"
)

// Operation represents one of the edit actions like insert, delete, and update.
type Operation byte

const (
	OperationUpdate Operation = iota
	OperationDelete
	OperationInsert
)

// Item holds the details about change in a document.
type Item struct {
	Operation Operation
	// Span is where the change happened: the inserted lines for an insert, the deleted lines for a delete,
	// and the changed part of the line for an update, where Start.Col is in both the old and the new
	// line, and End.Col is in the new line.
	Span     content.Span
	Old, New []content.Line // old/new lines within the change span; old is empty for inserts, new is empty for deletes
}

// Insertion describes inserting the lines at pos. There must be at least one line.
func Insertion(pos int, lines []content.Line) Item {
	return Item{
		Operation: OperationInsert,
		New:       lines,
		Span: content.Span{
			Start: content.Position{Line: pos},
			End:   content.Position{Line: pos + len(lines) - 1, Col: lines[len(lines)-1].Len()},
		},
	}
}

// Replacement describes replacing the line at pos, which was prev, with line.
// For the text lines, the span leaves out the prefix and the suffix the lines have in common.
func Replacement(pos int, prev, line content.Line) Item {
	prefixLen, suffixLen := 0, 0
	if prev.MimeType() == content.MimeTypeTextPlain {
		// TODO: this could happen concurrently.
		ps := prev.String()
		cs := line.String()
		for prefixLen < min(len(ps), len(cs)) && ps[prefixLen] == cs[prefixLen] {
			prefixLen++
		}
		for suffixLen < min(len(ps), len(cs))-prefixLen && ps[len(ps)-1-suffixLen] == cs[len(cs)-1-suffixLen] {
			suffixLen++
		}
	}
	return Item{
		Operation: OperationUpdate,
		Span: content.Span{
			Start: content.Position{Line: pos, Col: prefixLen},
			End:   content.Position{Line: pos, Col: line.Len() - suffixLen},
		},
		Old: []content.Line{prev},
		New: []content.Line{line},
	}
}

// Deletion describes deleting the lines starting at pos. There must be at least one line.
func Deletion(pos int, lines []content.Line) Item {
	return Item{
		Operation: OperationDelete,
		Old:       lines,
		Span: content.Span{
			Start: content.Position{Line: pos},
			End:   content.Position{Line: pos + len(lines) - 1, Col: lines[len(lines)-1].Len()},
		},
	}
}

func (it *Item) apply(m content.Mutable) content.Position {
	switch it.Operation {
	case OperationDelete:
		m.Delete(it.Span.Start.Line, it.Span.End.Line+1)
		return it.Span.Start
	case OperationInsert:
		m.Insert(it.Span.Start.Line, it.New...)
		return it.Span.End
	default:
		for i, line := range it.New {
			m.Update(it.Span.Start.Line+i, line)
		}
		return it.Span.End
	}
}

func (it *Item) revert(m content.Mutable) content.Position {
	switch it.Operation {
	case OperationDelete:
		m.Insert(it.Span.Start.Line, it.Old...)
		return it.Span.End
	case OperationInsert:
		m.Delete(it.Span.Start.Line, it.Span.End.Line+1)
		return it.Span.Start
	default:
		for i, line := range it.Old {
			m.Update(it.Span.Start.Line+i, line)
		}
		return it.Span.Start
	}
}

// Group combines multiple change items that done by the user at the same time (logically).
type Group struct {
	Time    time.Time
	Changes []Item
}

// List is a changes list on a base content.
type List struct {
	Base  content.Document
	Edits []Group
}

func (l *List) iterateChanges(yield func(Item) bool) {
	for _, group := range l.Edits {
		for _, item := range group.Changes {
			if !yield(item) {
				return
			}
		}
	}
}

// Replay applies the changes in the List to its Base copy and produces a new document.
// nil is returned if copying the base does not produce a mutable document.
func (l *List) Replay() content.Document {
	m, ok := content.Copy(l.Base).(content.Mutable)
	if !ok {
		return nil
	}

	for item := range l.iterateChanges {
		item.apply(m)
	}
	return m
}

// History represents the tracked change lists for some document.
type History []*List

const historyCap = 10

// Track wraps the mutable content returning implementation that tracks the changes.
func (h *History) Track(m content.Mutable) content.Mutable {
	base := content.Copy(m)
	if base == nil {
		// cannot be tracked
		return m
	}

	if cap(*h) == 0 {
		// ensure one allocation
		*h = make(History, 0, historyCap)
	}
	*h = append((*h)[max(len(*h)-historyCap, 0):], &List{Base: base})

	return &tracker{
		Mutable: m,
		list:    (*h)[len(*h)-1],
	}
}

type tracker struct {
	list *List
	content.Mutable
	currentPos     int
	groupRequested bool
}

const maxGroupDelay = 2 * time.Second

func (t *tracker) trackChange(change Item) {
	if l := len(t.list.Edits); l == 0 || t.currentPos != l-1 || t.groupRequested {
		t.startNewGroup(change)
		return
	}

	lastGroup := &t.list.Edits[t.currentPos]
	if time.Since(lastGroup.Time) > maxGroupDelay {
		t.startNewGroup(change)
		return
	}

	lastGroup.Time = time.Now()

	if change.Operation == OperationUpdate {
		lastChange := &lastGroup.Changes[len(lastGroup.Changes)-1]
		if lastChange.Operation == OperationUpdate {
			if lastChange.Span.Start.Line == change.Span.Start.Line && lastChange.Span.End.Line == change.Span.End.Line {
				// The span is found for the merged lines: the columns of the two spans
				// are in different versions of the line.
				*lastChange = Replacement(change.Span.Start.Line, lastChange.Old[0], change.New[0])
				return
			}
		}
	}

	lastGroup.Changes = append(lastGroup.Changes, change)
}

func (t *tracker) startNewGroup(change Item) {
	index := min(len(t.list.Edits)-1, t.currentPos)
	t.list.Edits = append(t.list.Edits[:index+1], Group{
		Time:    time.Now(),
		Changes: []Item{change},
	})
	t.currentPos = len(t.list.Edits) - 1
	t.groupRequested = false
}

func (t *tracker) Insert(pos int, lines ...content.Line) {
	if len(lines) == 0 {
		return
	}

	newLines := make([]content.Line, len(lines))
	for i, line := range lines {
		newLines[i] = content.CopyLine(line)
	}

	t.Mutable.Insert(pos, lines...)
	t.trackChange(Insertion(pos, newLines))
}

func (t *tracker) Update(pos int, line content.Line) {
	prevLine := t.Mutable.Lines()[pos]
	t.Mutable.Update(pos, line)
	t.trackChange(Replacement(pos, prevLine, content.CopyLine(line)))
}

func (t *tracker) Delete(startPos, endPos int) {
	if startPos == endPos {
		return
	}
	if startPos > endPos {
		panic(fmt.Errorf("startPos > endPos: %d > %d", startPos, endPos))
	}

	deletedLines := slices.Clone(t.Mutable.Lines()[startPos:endPos])
	t.Mutable.Delete(startPos, endPos)
	t.trackChange(Deletion(startPos, deletedLines))
}

func (t *tracker) HasUndo() bool { return len(t.list.Edits) > 0 && t.currentPos >= 0 }
func (t *tracker) HasRedo() bool { return t.currentPos < len(t.list.Edits)-1 }

func (t *tracker) Undo() (cursor content.Position) {
	group := t.list.Edits[t.currentPos]
	for _, change := range slices.Backward(group.Changes) {
		cursor = change.revert(t.Mutable)
	}
	t.currentPos--
	return
}

func (t *tracker) Redo() (cursor content.Position) {
	group := t.list.Edits[t.currentPos+1]
	for _, change := range group.Changes {
		cursor = change.apply(t.Mutable)
	}
	t.currentPos++
	return
}

func (t *tracker) SplitChangeGroup() {
	t.groupRequested = true
}
