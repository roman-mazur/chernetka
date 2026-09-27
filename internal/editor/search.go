package editor

import (
	"io"
	"regexp"
	"slices"
	"strings"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor/input"
)

// searchPrompt is shown while the search pattern is typed after "/". The cursor moves
// to the first match as the pattern is typed, and the matches are highlighted.
// Like in sed, "/pattern/replacement" replaces the matches when it's submitted.
type searchPrompt struct {
	c *cmdLine

	origin content.Position // the cursor position before the search was started
	offset int              // the scroll offset before the search was started
	xoff   int              // the horizontal scroll offset before the search was started
	prev   *regexp.Regexp   // the search highlighted before

	replace  bool   // the replacement is being typed
	template string // the replacement template for regexp.Expand
	failure  string // why the pattern cannot be compiled or "no matches"
}

// startSearch shows the command line to type the search pattern.
func (e *Editor) startSearch(buf *Buffer) {
	e.openCmdLine(buf, "", func(c *cmdLine) prompt {
		return &searchPrompt{c: c, origin: buf.c, offset: buf.offset, xoff: buf.xoff, prev: buf.search}
	})
}

func (p *searchPrompt) prefix() string { return "/" }

// input moves between the matches.
func (p *searchPrompt) input(_ *Editor, k input.Key) (handled bool) {
	switch {
	case k == input.Move(input.CursorArrowUp, 0), k == input.Of(input.Backtab):
		p.c.buf.searchMove(-1)
	case k == input.Move(input.CursorArrowDown, 0), k == input.Of(input.Tab), k == input.Ctrl('f'):
		p.c.buf.searchMove(1)
	case k.Special == input.CursorMove:
		// Other cursor keys are ignored.
	default:
		return false
	}
	return true
}

// changed searches for the typed pattern from the origin.
func (p *searchPrompt) changed(*Editor) {
	b := p.c.buf
	pattern, template, replace := parseSearch(p.c.text)
	p.replace, p.template, p.failure = replace, template, ""
	if pattern == "" {
		b.search = nil
		b.c, b.offset, b.xoff = p.origin, p.offset, p.xoff
		return
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		p.failure = "invalid pattern"
		return // Keep the last valid pattern and its matches.
	}
	b.search = re
	if span, ok := b.findMatch(p.origin, 1, true); ok {
		b.c = span.Start
	} else {
		b.c, b.offset, b.xoff = p.origin, p.offset, p.xoff
		p.failure = "no matches"
	}
}

// submit keeps the cursor at the current match, and the matches remain highlighted.
// If the replacement is typed, all the matches are replaced instead.
func (p *searchPrompt) submit(*Editor) (quit bool) {
	b := p.c.buf
	switch {
	case b.search == nil || p.failure != "":
		b.search = nil
	case p.replace:
		b.replaceAll(p.template)
		b.search = nil
	}
	return false
}

// cancel returns the cursor where it was before, and restores the previous search.
func (p *searchPrompt) cancel(*Editor) {
	b := p.c.buf
	b.c, b.offset, b.xoff = p.origin, p.offset, p.xoff
	b.search = p.prev
}

func (p *searchPrompt) height(int) int { return 1 }

func (p *searchPrompt) render(out io.Writer, w int) {
	var info string
	switch {
	case p.failure != "":
		info = p.failure
	case p.replace && p.c.buf.search != nil:
		info = "Enter: replace all"
	}
	renderCmdline(out, w, p.prefix()+p.c.text, info)
}

// searchMove moves the cursor to the next (dir > 0) or previous (dir < 0) match.
func (b *Buffer) searchMove(dir int) {
	if b.search == nil {
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
	if n == 0 || b.search == nil {
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
		matches := b.search.FindAllStringIndex(lines[ln].String(), -1)
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
	if b.search == nil {
		return nil
	}
	var res []content.Span
	for _, m := range b.search.FindAllStringIndex(line, -1) {
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

// replaceAll replaces all the search matches in the content with the expanded template.
// The cursor is placed at the start of the first replacement after the cursor.
func (b *Buffer) replaceAll(template string) {
	if !b.canEdit() {
		return
	}
	first, found := b.findMatch(b.c, 1, true)

	lines := b.Content.Lines()
	for ln, line := range slices.Backward(lines) {
		line := line.String()
		replaced := b.search.ReplaceAllString(line, template)
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

// parseSearch splits the command line input (after the "/" prefix) into the pattern
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
