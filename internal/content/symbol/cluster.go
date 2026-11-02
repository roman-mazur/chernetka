// Package symbol measures and walks the text the way a person sees it: by grapheme clusters
// rather than bytes or code points, and in terminal columns.
package symbol

import (
	"iter"
	"strings"
	"unicode/utf8"

	"github.com/rivo/uniseg"
)

// NoTabExpansion is the tab size for text that isn't laid out with tab stops, like the status
// bar, or when only the clusters matter and not their width: a tab takes one column.
const NoTabExpansion = 1

// Clusters iterates over the grapheme clusters of s, what a person sees as one character,
// with the terminal columns each one takes. A mermaid emoji is 4 code points, 13 bytes, but
// a single cluster of 2 columns. A tab takes tabSize columns.
//
// The text between two ASCII characters is not looked up: they are always separate clusters.
func Clusters(s string, tabSize int) iter.Seq2[string, int] {
	return func(yield func(cluster string, width int) bool) {
		for len(s) > 0 {
			var (
				cluster string
				width   int
			)
			if s[0] < utf8.RuneSelf && (len(s) == 1 || s[1] < utf8.RuneSelf) {
				cluster, s, width = s[:1], s[1:], 1
			} else {
				cluster, s, width, _ = uniseg.FirstGraphemeClusterInString(s, -1)
			}
			if cluster == "\t" {
				width = tabSize
			}
			if !yield(cluster, width) {
				return
			}
		}
	}
}

// DisplayWidth returns the terminal columns taken by s.
func DisplayWidth(s string, tabSize int) int {
	if isASCII(s) {
		return len(s) + strings.Count(s, "\t")*(tabSize-1)
	}
	// Repeats the logic in Clusters, so s can stay on the stack.
	total := 0
	for len(s) > 0 {
		var cluster string
		width := 1
		if s[0] < utf8.RuneSelf && (len(s) == 1 || s[1] < utf8.RuneSelf) {
			cluster, s = s[:1], s[1:]
		} else {
			cluster, s, width, _ = uniseg.FirstGraphemeClusterInString(s, -1)
		}
		if cluster == "\t" {
			width = tabSize
		}
		total += width
	}
	return total
}

// isASCII reports whether s can skip looking for clusters, which matters for very long lines.
func isASCII(s string) bool {
	for i := range len(s) {
		if s[i] >= utf8.RuneSelf {
			return false
		}
	}
	return true
}

// NextClusterEnd returns the byte index of the end of the cluster at col.
func NextClusterEnd(line string, col int) int {
	end := 0
	for cluster := range Clusters(line, NoTabExpansion) {
		end += len(cluster)
		if end > col {
			return end
		}
	}
	return len(line)
}

// PrevClusterStart returns the byte index of the start of the cluster before col,
// or of the cluster around it if col is in the middle of one.
func PrevClusterStart(line string, col int) int {
	start := 0
	for cluster := range Clusters(line, NoTabExpansion) {
		end := start + len(cluster)
		if end >= col {
			return start
		}
		start = end
	}
	return start
}

// ClusterStart returns the byte index of the start of the cluster around col.
func ClusterStart(line string, col int) int {
	if col <= 0 || col >= len(line) {
		return col
	}
	if line[col-1] < utf8.RuneSelf && line[col] < utf8.RuneSelf {
		return col
	}
	start := 0
	for cluster := range Clusters(line, NoTabExpansion) {
		end := start + len(cluster)
		if end > col {
			return start
		}
		start = end
	}
	return col
}

// ClusterEnd returns the byte index of the end of the cluster around col,
// or col if a cluster starts there.
func ClusterEnd(line string, col int) int {
	if ClusterStart(line, col) == col {
		return col
	}
	return NextClusterEnd(line, col)
}

// ClipColumns returns the part of s visible in the terminal columns from..to, s starting
// at column x. A wide character cut by the edge can't be shown: spaces take its visible part.
func ClipColumns(s string, x, from, to int) string {
	var visible strings.Builder
	for cluster, width := range Clusters(s, NoTabExpansion) {
		end := x + width
		switch {
		case x >= from && end <= to:
			visible.WriteString(cluster)
		case end > from && x < to:
			visible.WriteString(strings.Repeat(" ", min(end, to)-max(x, from)))
		}
		x = end
	}
	return visible.String()
}
