package extlsp

import (
	"strings"
	"unicode"
)

// expandSnippet converts an LSP snippet like "Println(${1:})" into plain text
// ("Println()") and the offset in it where the cursor should land after the
// insertion (right after the opening bracket): the first tab stop, or the
// final one ($0) if there are no others, or the end of the text.
//
// Placeholders keep their default text, and choices keep their first option.
// Nested placeholders are flattened.
func expandSnippet(snippet string) (text string, cursor int) {
	var (
		out       strings.Builder
		firstStop = -1 // the smallest tab stop number seen so far
		cursorAt  = -1
		finalAt   = -1
	)
	markStop := func(n int) {
		if n == 0 {
			if finalAt == -1 {
				finalAt = out.Len()
			}
			return
		}
		if firstStop == -1 || n < firstStop {
			firstStop, cursorAt = n, out.Len()
		}
	}

	depth := 0 // nesting level of ${...}
	for i := 0; i < len(snippet); i++ {
		ch := snippet[i]
		switch {
		case ch == '\\' && i+1 < len(snippet) && strings.IndexByte(`\$}`, snippet[i+1]) >= 0:
			i++
			out.WriteByte(snippet[i])
		case ch == '$' && i+1 < len(snippet) && isDigit(snippet[i+1]):
			n, j := readNumber(snippet, i+1)
			markStop(n)
			i = j - 1
		case ch == '$' && i+2 < len(snippet) && snippet[i+1] == '{' && isDigit(snippet[i+2]):
			n, j := readNumber(snippet, i+2)
			markStop(n)
			switch {
			case j < len(snippet) && snippet[j] == ':':
				depth++
				i = j // the placeholder text follows
			case j < len(snippet) && snippet[j] == '|':
				// A choice: keep the first option.
				end := strings.IndexAny(snippet[j+1:], ",|")
				if end < 0 {
					return snippet, len(snippet) // malformed
				}
				out.WriteString(snippet[j+1 : j+1+end])
				close := strings.Index(snippet[j+1:], "|}")
				if close < 0 {
					return snippet, len(snippet)
				}
				i = j + 1 + close + 1
			default: // ${1}
				i = j
			}
		case ch == '}' && depth > 0:
			depth--
		default:
			out.WriteByte(ch)
		}
	}

	text = out.String()
	switch {
	case cursorAt >= 0:
		return text, cursorAt
	case finalAt >= 0:
		return text, finalAt
	default:
		return text, len(text)
	}
}

func isDigit(b byte) bool { return unicode.IsDigit(rune(b)) }

func readNumber(s string, i int) (n, end int) {
	for end = i; end < len(s) && isDigit(s[end]); end++ {
		n = n*10 + int(s[end]-'0')
	}
	return n, end
}
