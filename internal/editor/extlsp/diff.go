package extlsp

import (
	"strings"
	"unicode/utf16"
	"unicode/utf8"

	"go.lsp.dev/protocol"
	"rmazur.io/chernetka/internal/lsp"
)

// diff returns a single change that turns before into after: the part between
// their common prefix and suffix is replaced. The range is in LSP coordinates
// (lines and UTF-16 code units) of before.
func diff(before, after string) lsp.TextChange {
	prefix := commonPrefix(before, after)
	suffix := commonSuffix(before[prefix:], after[prefix:])
	// Don't split a multi-byte rune: positions are counted in whole runes.
	for prefix > 0 && prefix < len(before) && !utf8.RuneStart(before[prefix]) {
		prefix--
	}
	for suffix > 0 && !utf8.RuneStart(before[len(before)-suffix]) {
		suffix--
	}

	start := position(before, prefix)
	end := start
	if removed := before[prefix : len(before)-suffix]; removed != "" {
		end = advance(start, removed)
	}
	return lsp.TextChange{
		Range: &protocol.Range{Start: start, End: end},
		Text:  after[prefix : len(after)-suffix],
	}
}

// chunk is how many bytes are compared at once when looking for the common
// prefix and suffix: comparing strings is vectorized, unlike a byte loop.
const chunk = 64

// commonPrefix returns the length of the common prefix of a and b.
func commonPrefix(a, b string) int {
	n := min(len(a), len(b))
	i := 0
	for i+chunk <= n && a[i:i+chunk] == b[i:i+chunk] {
		i += chunk
	}
	for i < n && a[i] == b[i] {
		i++
	}
	return i
}

// commonSuffix returns the length of the common suffix of a and b.
func commonSuffix(a, b string) int {
	n := min(len(a), len(b))
	i := 0
	for i+chunk <= n && a[len(a)-i-chunk:len(a)-i] == b[len(b)-i-chunk:len(b)-i] {
		i += chunk
	}
	for i < n && a[len(a)-1-i] == b[len(b)-1-i] {
		i++
	}
	return i
}

// position converts a byte offset in text to an LSP position.
func position(text string, offset int) protocol.Position {
	return advance(protocol.Position{}, text[:offset])
}

// advance moves p past s.
func advance(p protocol.Position, s string) protocol.Position {
	if i := strings.LastIndexByte(s, '\n'); i >= 0 {
		p.Line += uint32(strings.Count(s, "\n"))
		p.Character = 0
		s = s[i+1:]
	}
	for _, r := range s {
		p.Character += uint32(utf16.RuneLen(r))
	}
	return p
}

// applyChange applies change to text, as a server does.
func applyChange(text string, change lsp.TextChange) string {
	if change.Range == nil {
		return change.Text
	}
	start, end := offset(text, change.Range.Start), offset(text, change.Range.End)
	return text[:start] + change.Text + text[end:]
}

func offset(text string, p protocol.Position) int {
	i := 0
	for line := uint32(0); line < p.Line; line++ {
		i += strings.IndexByte(text[i:], '\n') + 1
	}
	var units uint32
	for j, r := range text[i:] {
		if units >= p.Character || r == '\n' {
			return i + j
		}
		units += uint32(utf16.RuneLen(r))
	}
	return len(text)
}
