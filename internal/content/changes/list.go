package changes

import (
	"time"

	"rmazur.io/chernetka/internal/content"
)

// EditItem represents a change on a content.
type EditItem struct {
	Time   time.Time
	Span   content.Span
	Line   content.Line // nil represents a deletion of the Span
	Insert bool
}

// List is a changes list on a base content.
type List struct {
	Base  content.Document
	Items []EditItem
}

// Replay applies the changes in the List to its Base copy and produces a new document.
// nil is returned if copying the base does not produce a mutable document.
func (l *List) Replay() content.Document {
	m, ok := content.Copy(l.Base).(content.Mutable)
	if !ok {
		return nil
	}
	for _, item := range l.Items {
		switch {
		case item.Insert:
			m.Insert(item.Span.Start.Line, item.Line)
		case item.Line == nil:
			m.Delete(item.Span.Start.Line)
		default:
			m.Update(item.Span.Start.Line, item.Line)
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

func (t *tracker) Insert(pos int, line content.Line) {
	t.Mutable.Insert(pos, line)

	t.list.Items = append(t.list.Items, EditItem{
		Time: time.Now(),
		Span: content.Span{
			Start: content.Position{Line: pos},
			End:   content.Position{Line: pos, Col: line.Len()},
		},
		Line:   content.CopyLine(line),
		Insert: true,
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

	t.list.Items = append(t.list.Items, EditItem{
		Time: time.Now(),
		Span: content.Span{
			Start: content.Position{Line: pos, Col: chStart},
			End:   content.Position{Line: pos, Col: chEnd},
		},
		Line: content.CopyLine(line),
	})
}

func (t *tracker) Delete(pos int) {
	prevLine := t.Mutable.Lines()[pos]
	t.Mutable.Delete(pos)
	t.list.Items = append(t.list.Items, EditItem{
		Time: time.Now(),
		Span: content.Span{
			Start: content.Position{Line: pos},
			End:   content.Position{Line: pos, Col: prevLine.Len()},
		},
	})
}
