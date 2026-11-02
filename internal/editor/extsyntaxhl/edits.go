package extsyntaxhl

import (
	"slices"

	"rmazur.io/chernetka/internal/content/change"
	treesitter "rmazur.io/chernetka/internal/editor/extsyntaxhl/gotreesitter"
)

// maxIncrementalChanges limits the number of changes turned into edits of the parsed tree.
// Each one is located in the text from its start, and more changes than that come from
// replacing much of the text at once, which is parsed from scratch about as fast.
const maxIncrementalChanges = 256

// textEdits converts the changes of the document into the edits of its text, telling a parser
// what was edited in the tree it built for the old revision. The old revision is given by the
// lengths of its lines, which are changed. The changes must be in the order they were made,
// and they must turn the old revision into the new one: ok is false otherwise, and the new
// revision is to be parsed from scratch. There are no edits if there are no changes,
// and the line lengths match the new revision.
//
// The text of a document is its lines joined with "\n" (see Buffer.Text), so a change of whole
// lines takes the line breaks around them too. An edit is located with the lengths of the lines
// left by the previous edits, which are checked against the lines of the new revision at the end.
func textEdits(oldLens []int, cur *source, changes []change.Item) (edits []treesitter.InputEdit, ok bool) {
	if len(changes) > maxIncrementalChanges {
		return nil, false
	}
	emptyText := len(oldLens) == 1 && oldLens[0] == 0
	if edits, ok = lineEdits(oldLens, cur, changes); !ok && emptyText {
		// An empty text is also the text of a document without lines.
		edits, ok = lineEdits(nil, cur, changes)
	}
	return edits, ok
}

// lineEdits implements textEdits for the old revision with the lines of the given lengths.
// The lengths are changed.
func lineEdits(lens []int, cur *source, changes []change.Item) (edits []treesitter.InputEdit, ok bool) {
	// lineStart returns the offset of the start of the line, which may be the end of the text.
	lineStart := func(ln int) (offset int) {
		for _, l := range lens[:ln] {
			offset += l + 1
		}
		return offset
	}
	point := func(ln, col int) treesitter.Point { return treesitter.NewPoint(uint(ln), uint(col)) }

	edits = make([]treesitter.InputEdit, 0, len(changes))
	for _, ch := range changes {
		n := len(lens)
		var e treesitter.InputEdit
		switch ch.Operation {
		case change.OperationInsert:
			pos := ch.Span.Start.Line
			if len(ch.New) == 0 || pos < 0 || pos > n {
				return nil, false
			}
			added := len(ch.New) - 1 // line breaks between the new lines
			newLens := make([]int, len(ch.New))
			for i, line := range ch.New {
				if line == nil {
					return nil, false
				}
				newLens[i] = line.Len()
				added += newLens[i]
			}
			last := len(ch.New) - 1
			switch {
			case pos < n: // The lines and a line break go before the line at pos.
				start := lineStart(pos)
				e = insertEdit(start, point(pos, 0), added+1, point(pos+last+1, 0))
			case n > 0: // A line break and the lines go after the last line.
				start := lineStart(n-1) + lens[n-1]
				e = insertEdit(start, point(n-1, lens[n-1]), added+1, point(n+last, newLens[last]))
			default: // The lines are all the document has.
				e = insertEdit(0, point(0, 0), added, point(last, newLens[last]))
			}
			lens = slices.Insert(lens, pos, newLens...)

		case change.OperationDelete:
			startLn, endLn := ch.Span.Start.Line, ch.Span.Start.Line+len(ch.Old)
			if len(ch.Old) == 0 || startLn < 0 || endLn > n {
				return nil, false
			}
			removed := len(ch.Old) // the line break of every line but the last one of the document
			for i, line := range ch.Old {
				if line == nil || line.Len() != lens[startLn+i] {
					return nil, false
				}
				removed += lens[startLn+i]
			}
			switch {
			case endLn < n: // The lines go with the line breaks after them.
				start := lineStart(startLn)
				e = deleteEdit(start, point(startLn, 0), removed, point(endLn, 0))
			case startLn > 0: // The last lines go with the line breaks before them.
				start := lineStart(startLn) - 1
				e = deleteEdit(start, point(startLn-1, lens[startLn-1]), removed, point(n-1, lens[n-1]))
			default: // The whole text goes, there's one line break less than lines.
				e = deleteEdit(0, point(0, 0), removed-1, point(n-1, lens[n-1]))
			}
			lens = slices.Delete(lens, startLn, endLn)

		case change.OperationUpdate:
			ln := ch.Span.Start.Line
			if len(ch.Old) != 1 || len(ch.New) != 1 || ch.Old[0] == nil || ch.New[0] == nil ||
				ln < 0 || ln >= n || ch.Old[0].Len() != lens[ln] {
				return nil, false
			}
			// The span leaves out the common prefix and suffix of the lines.
			oldLen, newLen := lens[ln], ch.New[0].Len()
			startCol := ch.Span.Start.Col
			suffix := newLen - ch.Span.End.Col
			oldEndCol, newEndCol := oldLen-suffix, newLen-suffix
			if startCol < 0 || suffix < 0 || startCol > oldEndCol || startCol > newEndCol {
				return nil, false
			}
			lineOffset := lineStart(ln)
			e = treesitter.InputEdit{
				StartByte:      uint(lineOffset + startCol),
				OldEndByte:     uint(lineOffset + oldEndCol),
				NewEndByte:     uint(lineOffset + newEndCol),
				StartPosition:  point(ln, startCol),
				OldEndPosition: point(ln, oldEndCol),
				NewEndPosition: point(ln, newEndCol),
			}
			lens[ln] = newLen

		default:
			return nil, false
		}
		edits = append(edits, e)
	}

	// An empty text is one empty line, while the document may have no lines at all.
	if len(lens) == 0 {
		lens = append(lens, 0)
	}
	if len(lens) != len(cur.lines) {
		return nil, false
	}
	for i, line := range cur.lines {
		if len(line) != lens[i] {
			return nil, false
		}
	}
	return edits, true
}

// insertEdit describes inserting size bytes at start, ending at the end point.
func insertEdit(start int, startPoint treesitter.Point, size int, end treesitter.Point) treesitter.InputEdit {
	return treesitter.InputEdit{
		StartByte:      uint(start),
		OldEndByte:     uint(start),
		NewEndByte:     uint(start + size),
		StartPosition:  startPoint,
		OldEndPosition: startPoint,
		NewEndPosition: end,
	}
}

// deleteEdit describes deleting size bytes at start, which ended at the end point.
func deleteEdit(start int, startPoint treesitter.Point, size int, end treesitter.Point) treesitter.InputEdit {
	return treesitter.InputEdit{
		StartByte:      uint(start),
		OldEndByte:     uint(start + size),
		NewEndByte:     uint(start),
		StartPosition:  startPoint,
		OldEndPosition: end,
		NewEndPosition: startPoint,
	}
}
