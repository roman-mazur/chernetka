package change

import (
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
	Span      content.Span
	Old, New  []content.Line // old/new lines within the change span; old is empty for inserts, new is empty for inserts
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
		switch item.Operation {
		case OperationDelete:
			m.Delete(item.Span.Start.Line, item.Span.End.Line+1)
		case OperationInsert:
			m.Insert(item.Span.Start.Line, item.New...)
		default:
			for i, line := range item.New {
				m.Update(item.Span.Start.Line+i, line)
			}
		}
	}
	return m
}

// History represents the tracked change lists for some document.
type History []List

// Track wraps the mutable content returning implementation that tracks the changes.
func (h *History) Track(m content.Mutable) content.Mutable {
	base := content.Copy(m)
	if base == nil {
		// cannot be tracked
		return m
	}
	*h = append(*h, List{Base: base})
	return &tracker{
		Mutable: m,
		list:    &(*h)[len(*h)-1],
	}
}

type tracker struct {
	list *List
	content.Mutable
}

const maxGroupDelay = 2 * time.Second

func (t *tracker) trackChange(change Item) {
	if len(t.list.Edits) == 0 {
		t.startNewGroup(change)
		return
	}

	lastGroup := &t.list.Edits[len(t.list.Edits)-1]
	if time.Since(lastGroup.Time) > maxGroupDelay {
		t.startNewGroup(change)
		return
	}

	if change.Operation == OperationUpdate {
		lastChange := &lastGroup.Changes[len(lastGroup.Changes)-1]
		if lastChange.Operation == OperationUpdate {
			if lastChange.Span.Start.Line == change.Span.Start.Line && lastChange.Span.End.Line == change.Span.End.Line {
				lastChange.New = change.New
				return
			}
		}
	}

	lastGroup.Changes = append(lastGroup.Changes, change)
}

func (t *tracker) startNewGroup(change Item) {
	t.list.Edits = append(t.list.Edits, Group{
		Time:    time.Now(),
		Changes: []Item{change},
	})
}

func (t *tracker) Insert(pos int, lines ...content.Line) {
	newLines := make([]content.Line, len(lines))
	for i, line := range lines {
		newLines[i] = content.CopyLine(line)
	}

	t.Mutable.Insert(pos, lines...)

	t.trackChange(Item{
		Operation: OperationInsert,
		New:       newLines,
		Span: content.Span{
			Start: content.Position{Line: pos},
			End:   content.Position{Line: pos + len(newLines) - 1, Col: newLines[len(newLines)-1].Len()},
		},
	})
}

func (t *tracker) Update(pos int, line content.Line) {
	prevLine := t.Mutable.Lines()[pos]
	t.Mutable.Update(pos, line)

	chStart, chEnd := 0, prevLine.Len()
	if prevLine.MimeType() == content.MimeTypeTextPlain {
		// TODO: this could happen concurrently.
		ps := prevLine.String()
		cs := line.String()
		for chStart < min(len(ps), len(cs)) && ps[chStart] == cs[chStart] {
			chStart++
		}
		for chEnd > 0 {
			if len(cs) < chEnd {
				break
			}
			if ps[chEnd-1] != cs[chEnd-1] {
				break
			}
			chEnd--
		}
	}

	t.trackChange(Item{
		Operation: OperationUpdate,
		Span: content.Span{
			Start: content.Position{Line: pos, Col: chStart},
			End:   content.Position{Line: pos, Col: chEnd},
		},
		Old: []content.Line{prevLine},
		New: []content.Line{content.CopyLine(line)},
	})
}

func (t *tracker) Delete(startPos, endPos int) {
	deletedLines := slices.Clone(t.Mutable.Lines()[startPos:endPos])
	t.Mutable.Delete(startPos, endPos)
	t.trackChange(Item{
		Operation: OperationDelete,
		Old:       deletedLines,
		Span: content.Span{
			Start: content.Position{Line: startPos},
			End:   content.Position{Line: endPos - 1, Col: deletedLines[len(deletedLines)-1].Len()},
		},
	})
}
