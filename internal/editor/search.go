package editor

import (
	"regexp"
	"slices"
	"strings"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/inputs"
)

// searchPrefix starts the command line of the search: "/pattern" or "/pattern/replacement".
const searchPrefix = "/"

// bufSearch is the state of the search in a Buffer. The matches of re are highlighted
// while it's set, and the cursor can be moved between them.
type bufSearch struct {
	re *regexp.Regexp

	typing   bool             // the pattern is being typed in the command line
	prevMode Mode             // the mode to return to when typing is done
	origin   content.Position // the cursor position before the search was started
	offset   int              // the scroll offset before the search was started

	replace   bool   // the replacement is being typed
	template  string // the replacement template for regexp.Expand
	failure   string // why the pattern cannot be compiled or "no matches"
	lastInput string // the command line the state was computed for
}

// startSearch switches the buffer to the command mode to type the search pattern.
func (b *Buffer) startSearch() {
	prevMode := b.mode
	if prevMode == ModeCommand {
		prevMode = ModeNormal
	}
	b.search = bufSearch{
		typing:   true,
		prevMode: prevMode,
		origin:   b.c,
		offset:   b.offset,
	}
	b.mode = ModeCommand
	b.cmdline = searchPrefix
}

// searchInput handles the keys specific to typing the search pattern.
// Other input edits the command line as usual.
func searchInput(b *Buffer, input []byte) (handled bool) {
	var (
		arrow inputs.Cursor
		mod   inputs.Modifier
	)
	switch {
	case inputs.IsCursor(input, &arrow, &mod):
		switch arrow {
		case inputs.CursorArrowUp:
			b.searchMove(-1)
		case inputs.CursorArrowDown:
			b.searchMove(1)
		}
	case inputs.IsTab(input), inputs.IsFindCommand(input):
		b.searchMove(1)
	case inputs.IsBacktab(input):
		b.searchMove(-1)
	case inputs.IsEscape(input):
		b.cancelSearch()
	case len(input) == 1 && input[0] == '\r':
		b.finishSearch()
	default:
		return false
	}
	return true
}

// syncSearch updates the search state to follow the typed command line.
// Deleting the whole pattern including the "/" prefix cancels the search.
func (b *Buffer) syncSearch() {
	s := &b.search
	if !s.typing {
		return
	}
	if b.mode != ModeCommand || !strings.HasPrefix(b.cmdline, searchPrefix) {
		b.cancelSearch()
		return
	}
	if s.lastInput == b.cmdline {
		return
	}
	s.lastInput = b.cmdline

	pattern, template, replace := parseSearch(strings.TrimPrefix(b.cmdline, searchPrefix))
	s.replace, s.template, s.failure = replace, template, ""
	if pattern == "" {
		s.re = nil
		b.c, b.offset = s.origin, s.offset
		return
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		s.failure = "invalid pattern"
		return // Keep the last valid pattern and its matches.
	}
	s.re = re
	if span, ok := b.findMatch(s.origin, 1, true); ok {
		b.c = span.Start
	} else {
		b.c, b.offset = s.origin, s.offset
		s.failure = "no matches"
	}
}

// cancelSearch stops the search returning the cursor to where it was before.
func (b *Buffer) cancelSearch() {
	s := b.search
	if s.typing {
		b.c, b.offset = s.origin, s.offset
		b.mode = s.prevMode
		b.cmdline = ""
	}
	b.search = bufSearch{}
}

// finishSearch completes typing: the cursor stays at the current match, and
// the matches remain highlighted. If the replacement is typed, all the matches
// are replaced instead.
func (b *Buffer) finishSearch() {
	s := &b.search
	b.mode, b.cmdline = s.prevMode, ""
	s.typing = false
	if s.re == nil || s.failure != "" {
		b.search = bufSearch{}
		return
	}
	if s.replace {
		b.replaceAll()
		b.search = bufSearch{}
	}
}

// searchMove moves the cursor to the next (dir > 0) or previous (dir < 0) match.
func (b *Buffer) searchMove(dir int) {
	if b.search.re == nil {
		return
	}
	if span, ok := b.findMatch(b.c, dir, false); ok {
		b.c = span.Start
		b.updateSelection()
	}
}

// findMatch looks for the match closest to from in the direction dir, wrapping around the content end.
// A match at from is only accepted if inclusive is set.
func (b *Buffer) findMatch(from content.Position, dir int, inclusive bool) (content.Span, bool) {
	lines := b.Content.Lines()
	n := len(lines)
	if n == 0 || b.search.re == nil {
		return content.Span{}, false
	}
	from.Line = max(0, min(from.Line, n-1))

	accept := func(m []int, first bool) bool {
		if !first {
			return true
		}
		switch {
		case dir > 0 && inclusive:
			return m[0] >= from.Col
		case dir > 0:
			return m[0] > from.Col
		default:
			return m[0] < from.Col
		}
	}
	// Check lines starting with the one of the cursor, and finish with it again
	// to find the matches on the other side of the cursor after wrapping.
	for i := 0; i <= n; i++ {
		ln := (from.Line + dir*i + n*(n+1)) % n
		matches := b.search.re.FindAllStringIndex(lines[ln].String(), -1)
		if dir < 0 {
			slices.Reverse(matches)
		}
		for _, m := range matches {
			if i == n || accept(m, i == 0) {
				return content.Span{
					Start: content.Position{Line: ln, Col: m[0]},
					End:   content.Position{Line: ln, Col: m[1]},
				}, true
			}
		}
	}
	return content.Span{}, false
}

// searchMatchesOnLine returns the non-empty matches on the line to highlight them.
func (b *Buffer) searchMatchesOnLine(ln int, line string) []content.Span {
	if b.search.re == nil {
		return nil
	}
	var res []content.Span
	for _, m := range b.search.re.FindAllStringIndex(line, -1) {
		if m[0] == m[1] {
			continue
		}
		res = append(res, content.Span{
			Start: content.Position{Line: ln, Col: m[0]},
			End:   content.Position{Line: ln, Col: m[1]},
		})
	}
	return res
}

// replaceAll replaces all the matches in the content with the expanded template.
// The cursor is placed at the start of the first replacement after the cursor.
func (b *Buffer) replaceAll() {
	if !b.canEdit() {
		return
	}
	s := b.search
	first, found := b.findMatch(b.c, 1, true)

	lines := b.Content.Lines()
	for ln := len(lines) - 1; ln >= 0; ln-- {
		line := lines[ln].String()
		replaced := s.re.ReplaceAllString(line, s.template)
		if replaced == line {
			continue
		}
		b.ReplaceText(content.Span{
			Start: content.Position{Line: ln},
			End:   content.Position{Line: ln, Col: len(line)},
		}, replaced)
	}
	if found {
		b.c = first.Start
	}
}

// parseSearch splits the command line input (without the "/" prefix) into the pattern
// and the replacement template. The replacement follows the second unescaped "/".
// Like in sed, "\/" is a slash in the pattern or replacement, an optional trailing
// "/" finishes the replacement, and the template can refer to the whole match with "&"
// and to the groups with "\1".."\9". "\n" and "\t" insert a new line and a tab.
func parseSearch(input string) (pattern, template string, replace bool) {
	var p strings.Builder
	i := 0
	for ; i < len(input); i++ {
		c := input[i]
		if c == '\\' && i+1 < len(input) && input[i+1] == '/' {
			p.WriteByte('/')
			i++
			continue
		}
		if c == '\\' && i+1 < len(input) {
			p.WriteString(input[i : i+2])
			i++
			continue
		}
		if c == '/' {
			replace = true
			i++
			break
		}
		p.WriteByte(c)
	}
	if !replace {
		return p.String(), "", false
	}

	var t strings.Builder
	repl := input[i:]
	for i := 0; i < len(repl); i++ {
		c := repl[i]
		switch {
		case c == '\\' && i+1 < len(repl):
			i++
			switch e := repl[i]; {
			case '0' <= e && e <= '9':
				t.WriteString("${" + string(e) + "}")
			case e == 'n':
				t.WriteByte('\n')
			case e == 't':
				t.WriteByte('\t')
			case e == '$':
				t.WriteString("$$")
			default:
				t.WriteByte(e)
			}
		case c == '/' && i == len(repl)-1:
			// The closing slash.
		case c == '&':
			t.WriteString("${0}")
		case c == '$':
			t.WriteString("$$")
		default:
			t.WriteByte(c)
		}
	}
	return p.String(), t.String(), true
}
