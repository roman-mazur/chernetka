package extsyntaxhl

import (
	"strings"

	treesitter "rmazur.io/chernetka/internal/editor/extsyntaxhl/gotreesitter"
)

// textEdit describes turning old into cur as a single edit of the text between their
// common prefix and suffix, telling a parser what was edited in the tree it built for old.
func textEdit(old, cur string) treesitter.InputEdit {
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
	return treesitter.InputEdit{
		StartByte:      uint(start),
		OldEndByte:     uint(oldEnd),
		NewEndByte:     uint(newEnd),
		StartPosition:  pointAt(old, start),
		OldEndPosition: pointAt(old, oldEnd),
		NewEndPosition: pointAt(cur, newEnd),
	}
}

// pointAt returns the row and the byte column of the offset in the text.
func pointAt(text string, offset int) treesitter.Point {
	before := text[:offset]
	row := strings.Count(before, "\n")
	col := offset - (strings.LastIndexByte(before, '\n') + 1)
	return treesitter.NewPoint(uint(row), uint(col))
}
