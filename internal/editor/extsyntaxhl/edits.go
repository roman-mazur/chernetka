package extsyntaxhl

import (
	"strings"

	"rmazur.io/chernetka/internal/content"
)

// textEdit is an edit of the text of a document, telling a highlighter what changed since
// the revision it parsed: the bytes [StartByte, OldEndByte) of the old text are replaced with
// [StartByte, NewEndByte) of the new one. A position is the line and the byte column of an offset.
type textEdit struct {
	StartByte, OldEndByte, NewEndByte             int
	StartPosition, OldEndPosition, NewEndPosition content.Position
}

// diffEdit describes turning old into cur as a single edit of the text between their
// common prefix and suffix.
func diffEdit(old, cur string) textEdit {
	n := min(len(old), len(cur))
	start := 0
	for start < n && old[start] == cur[start] {
		start++
	}
	suffix := 0
	for suffix < n-start && old[len(old)-1-suffix] == cur[len(cur)-1-suffix] {
		suffix++
	}
	oldEnd, newEnd := len(old)-suffix, len(cur)-suffix
	return textEdit{
		StartByte:      start,
		OldEndByte:     oldEnd,
		NewEndByte:     newEnd,
		StartPosition:  pointAt(old, start),
		OldEndPosition: pointAt(old, oldEnd),
		NewEndPosition: pointAt(cur, newEnd),
	}
}

// pointAt returns the row and the byte column of the offset in the text.
func pointAt(text string, offset int) content.Position {
	before := text[:offset]
	return content.Position{
		Line: strings.Count(before, "\n"),
		Col:  offset - (strings.LastIndexByte(before, '\n') + 1),
	}
}
