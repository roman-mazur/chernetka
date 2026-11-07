// Package markdown scans Markdown documents with a hand written scanner.
//
// Markdown suits a hand written scanner well: almost every construct is
// anchored to the start of a line, and the two that are not - fenced code
// blocks and setext headings - only need the lines immediately around them.
// The scanner is deliberately pragmatic rather than a complete CommonMark
// implementation. It aims to colorize what someone writing notes actually
// types, and it never fails: text it does not recognize is simply left
// unhighlighted.
package markdown

import (
	"slices"
	"strings"
	"unicode/utf8"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/content/code"
)

// Span is a range of a line, [Start, End) in bytes, recognized as the Token.
type Span struct {
	Line, Start, End int
	Token            code.TokenType
}

// Scan walks the document lines reporting the recognized ranges to emit. The spans
// of the nested constructs follow the ones they are nested in, and a span may be
// reported for a line preceding the current one (e.g., the setext heading text).
// It returns the fenced code blocks of the document.
func Scan(lines []string, emit func(Span)) []code.Block {
	s := mdScanner{lines: lines, emit: emit}
	s.run()
	return s.blocks
}

// TaskCheckbox returns the index of the mark within the checkbox of a task list item
// line, like "- [ ] task" or "1. [x] done", or -1 if the line is not one.
// The item may be nested in block quotes. The line is checked on its own: the document
// context, like being inside a code block, is not taken into account.
func TaskCheckbox(line string) int {
	i, _ := skipQuotes(line, len(line))
	if i = skipSpace(line, i); i == len(line) {
		return -1
	}
	item, ok := parseListItem(line, i)
	if !ok {
		return -1
	}
	return item.checkbox
}

// mdScanner walks a document line by line, carrying the little bit of state
// that spills across line boundaries.
type mdScanner struct {
	lines []string
	emit  func(Span)

	// Open code fence, tracked so its body is colorized as raw text.
	inFence     bool
	fenceChar   byte
	fenceLen    int
	fenceIndent int // columns between the fence and the content start of its block quotes
	fenceQuotes int // the number of the block quotes the fence is nested in

	// The block quote markers of the current line scanned so far.
	quotes   int // their number
	quoteEnd int // the index after the last of them

	para        paragraph // the paragraph line found on the current line
	above       paragraph // the paragraph line found on the line above
	inTable     bool      // the previous line was part of a table
	listContent int       // content column of the innermost open list item, or 0

	blocks []code.Block // fenced code blocks found so far

	// The inline constructs whose closer is not found are remembered for the line, so
	// that the other openers of the same kind after them fail without rescanning the
	// rest of it: rescanning it from each unclosed opener is quadratic.
	noEmphasis   [3][4]failure // by the delimiter ('*', '_', '~') and the run length
	noLinkTarget [2]failure    // by the closing byte (')' or ']')
	noAngle      failure

	// The code spans and the brackets of a line are matched once: searching for the
	// closer of each opener is quadratic.
	matchedLine  int   // the line number of the matches below plus one, or 0 if none
	codeEnd      []int // the end of the code span opened at each index of the line, or -1
	brackets     []int // the index of the ']' matching each '[' of the line, or -1
	bracketStack []int
	ticks        []tickRun
	lastTicks    map[int]int

	linkDepth int // the number of the link labels the scanned text is nested in
}

// paragraph describes a paragraph line, the only line a setext underline may follow.
type paragraph struct {
	line   int // the line number, or -1 if there is no paragraph line
	col    int // the start of the paragraph text
	quotes int // the number of the block quotes the paragraph is nested in
}

// tickRun is a maximal run of backticks, line[start:end].
type tickRun struct {
	start, end int
	// The next run after this one of the same length, and of one backtick less, or -1.
	same, shorter int
}

// maxLinkDepth is how deep the links nest: a link may hold an image, like a badge
// "[![build](badge.svg)](url)". The inner links of a deeper nesting are not
// highlighted, which keeps the text scanned again for each level within a limit.
const maxLinkDepth = 2

// failure records that a closer was not found in line[from:to]. Then it's not found
// in a narrower range either: the delimiters are matched in the same runs.
type failure struct{ ln, from, to int }

// covers reports whether the closer is known to be missing in line[from:to].
func (f failure) covers(ln, from, to int) bool {
	return f.ln == ln && f.from <= from && to <= f.to
}

func (s *mdScanner) run() {
	s.para.line = -1
	for ln, line := range s.lines {
		s.above, s.para = s.para, paragraph{line: -1}
		s.quotes, s.quoteEnd = 0, 0
		if s.inFence && s.scanFenceLine(ln, line) {
			continue
		}
		s.scanBlock(ln, line, 0)
	}
	// An unterminated fence just runs to the end of the document.
	if s.inFence {
		last := len(s.lines) - 1
		s.closeBlock(last, len(s.lines[last]), false)
	}
}

// closeBlock ends the last found code block at the given position.
func (s *mdScanner) closeBlock(ln, col int, closed bool) {
	block := &s.blocks[len(s.blocks)-1]
	block.End = content.Position{Line: ln, Col: col}
	block.Closed = closed
}

// scanFenceLine colorizes a line of the open fence. It reports false if the line
// does not continue the block quotes the fence is nested in: then the quotes end,
// and the fence along with them.
func (s *mdScanner) scanFenceLine(ln int, line string) bool {
	start, quotes := skipQuotes(line, s.fenceQuotes)
	if quotes < s.fenceQuotes {
		s.inFence = false
		s.closeBlock(ln-1, len(s.lines[ln-1]), false)
		return false
	}
	for i := range start {
		if line[i] == '>' {
			s.mark(ln, i, i+1, code.TtPunctuation)
		}
	}
	if s.closesFence(line[start:]) {
		s.mark(ln, start, len(line), code.TtPunctuation)
		s.inFence = false
		s.closeBlock(ln, len(line), true)
	} else {
		s.mark(ln, start, len(line), code.TtRawText)
	}
	return true
}

// scanBlock highlights the block construct that starts at or after col. col is
// what lets blocks nest: a block quote colorizes its own '>' and then hands the
// rest of the line back here, so "> ## Heading" is still a heading.
func (s *mdScanner) scanBlock(ln int, line string, col int) {
	indent := skipSpace(line, col)
	if indent == len(line) {
		// A blank line closes the table above it, and the paragraph, which is
		// not set for the line.
		s.inTable = false
		return
	}

	if col == 0 {
		if indent == 0 {
			s.listContent = 0
		}
		// An indented code block: four columns of indentation, following a blank
		// line, and not the continuation of a list item.
		if width(line[:indent]) >= 4 && s.above.line < 0 && s.listContent == 0 {
			s.mark(ln, indent, len(line), code.TtRawText)
			return
		}
	}

	switch {
	case s.scanFenceOpen(ln, line, indent):
	case s.scanHeading(ln, line, indent):
	// A setext underline has to be tried before a thematic break, because '---'
	// under a paragraph line underlines it instead of breaking the page.
	case s.scanSetext(ln, line, indent):
	case s.scanThematicBreak(ln, line, indent):
	case s.scanBlockQuote(ln, line, col, indent):
	case s.scanTableRow(ln, line, indent):
	case s.scanListItem(ln, line, indent):
	case s.scanLinkDefinition(ln, line, indent):
	default:
		s.scanInline(ln, line, indent, len(line))
		s.para = paragraph{line: ln, col: indent, quotes: s.quotes}
	}
}

// scanFenceOpen recognizes the ``` or ~~~ line that opens a code fence.
func (s *mdScanner) scanFenceOpen(ln int, line string, indent int) bool {
	c := line[indent]
	if c != '`' && c != '~' {
		return false
	}
	i := skipRun(line, indent, len(line), c)
	if i-indent < 3 {
		return false
	}
	// Whatever follows the fence is the info string naming the language.
	info := line[i:]
	if c == '`' && strings.IndexByte(info, '`') >= 0 {
		return false // Inline code rather than a fence.
	}

	s.mark(ln, indent, i, code.TtPunctuation)
	s.mark(ln, i, len(line), code.TtKeyword)

	s.inFence, s.fenceChar, s.fenceLen = true, c, i-indent
	s.fenceIndent, s.fenceQuotes = width(line[s.quoteEnd:indent]), s.quotes
	var lang string
	if fields := strings.Fields(info); len(fields) > 0 {
		lang = fields[0]
	}
	s.blocks = append(s.blocks, code.Block{
		Start: content.Position{Line: ln, Col: indent},
		Lang:  lang,
	})
	return true
}

// closesFence reports whether line, stripped of the block quote markers, ends the open
// fence: a run of at least as many of the fence's own characters, with nothing else on
// the line.
func (s *mdScanner) closesFence(line string) bool {
	i := skipSpace(line, 0)
	if width(line[:i]) > s.fenceIndent+3 {
		return false
	}
	run := skipRun(line, i, len(line), s.fenceChar)
	return run-i >= s.fenceLen && isBlank(line[run:])
}

// scanHeading recognizes an ATX heading, "## Like this ##".
func (s *mdScanner) scanHeading(ln int, line string, indent int) bool {
	i := skipRun(line, indent, len(line), '#')
	if level := i - indent; level == 0 || level > 6 {
		return false
	}
	// The marker has to be followed by a space, or be the whole line.
	if i < len(line) && !isSpaceByte(line[i]) {
		return false
	}
	s.mark(ln, indent, i, code.TtPunctuation)

	end := max(i, len(strings.TrimRight(line, " \t")))
	// A closed heading ends with its own run of '#', which is punctuation too.
	closing := max(i, len(strings.TrimRight(line[:end], "#")))
	if closing < end && (closing == i || isSpaceByte(line[closing-1])) {
		s.mark(ln, closing, end, code.TtPunctuation)
		end = closing
	}

	// The heading text is colorized as a whole, then the inline markup inside it
	// overrides the narrower ranges it covers.
	s.mark(ln, i, end, code.TtHeading)
	s.scanInline(ln, line, i, end)
	return true
}

// scanSetext recognizes the '===' or '---' line that underlines the paragraph
// line above it, turning it into a heading. The underline has to be in the same
// block quotes as the paragraph: an underline can't continue a quote lazily.
func (s *mdScanner) scanSetext(ln int, line string, indent int) bool {
	p := s.above
	if p.line < 0 || p.line != ln-1 || p.quotes != s.quotes {
		return false
	}
	c := line[indent]
	if c != '=' && c != '-' {
		return false
	}
	if !isBlank(line[skipRun(line, indent, len(line), c):]) {
		return false
	}

	s.mark(ln, indent, len(line), code.TtPunctuation)
	// Promote the line above after the fact. Spans are collected before they are
	// flattened, so emitting one for an earlier line is fine, and the inline
	// spans already reported for it are narrower and still win.
	s.mark(p.line, p.col, len(s.lines[p.line]), code.TtHeading)
	return true
}

// scanThematicBreak recognizes a '---', '***' or '___' page break.
func (s *mdScanner) scanThematicBreak(ln int, line string, indent int) bool {
	c := line[indent]
	if c != '-' && c != '*' && c != '_' {
		return false
	}
	count := 0
	for i := indent; i < len(line); i++ {
		switch {
		case line[i] == c:
			count++
		case isSpaceByte(line[i]):
		default:
			return false
		}
	}
	if count < 3 {
		return false
	}
	s.mark(ln, indent, len(line), code.TtListMarker)
	return true
}

// scanBlockQuote colorizes a '>' marker and re-dispatches the rest of the line,
// so the blocks nested in a quote are highlighted as themselves.
func (s *mdScanner) scanBlockQuote(ln int, line string, col, indent int) bool {
	if line[indent] != '>' {
		return false
	}
	s.mark(ln, indent, indent+1, code.TtPunctuation)
	s.quotes, s.quoteEnd = s.quotes+1, indent+1
	if col == 0 {
		// The outermost quote covers the nested ones: marking the rest of the line
		// for each of them makes the spans of "> > > ..." quadratic in its length.
		s.mark(ln, indent+1, len(line), code.TtQuote)
	}
	s.scanBlock(ln, line, indent+1)
	return true
}

// scanTableRow colorizes a GFM table row. A row only counts as one next to the
// '|---|---|' delimiter row that defines the table, so a stray pipe in a
// sentence is not mistaken for a table.
func (s *mdScanner) scanTableRow(ln int, line string, indent int) bool {
	if !strings.Contains(line[indent:], "|") {
		s.inTable = false
		return false
	}
	delimiter := isTableDelimiterRow(line[indent:])
	// The row above the delimiter is the header, and reads as bold. The delimiter
	// row is in the same block quotes as the header.
	next := s.line(ln + 1)
	start, quotes := skipQuotes(next, s.quotes)
	header := quotes == s.quotes && isTableDelimiterRow(next[start:])
	if !s.inTable && !delimiter && !header {
		return false
	}
	s.inTable = true

	if delimiter {
		s.mark(ln, indent, len(line), code.TtPunctuation)
		return true
	}
	cellToken := code.TtNothing
	if header {
		cellToken = code.TtStrong
	}
	cell := indent
	markCell := func(end int) {
		s.mark(ln, cell, end, cellToken)
		s.scanInline(ln, line, cell, end)
	}
	for i := indent; i < len(line); i++ {
		if line[i] != '|' {
			continue
		}
		markCell(i)
		s.mark(ln, i, i+1, code.TtPunctuation)
		cell = i + 1
	}
	markCell(len(line))
	return true
}

// isTableDelimiterRow reports whether s looks like '| --- | :---: |'.
func isTableDelimiterRow(s string) bool {
	cells, dashes := 0, 0
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] == '|':
			cells++
		case s[i] == '-':
			dashes++
		case s[i] == ':' || isSpaceByte(s[i]):
		default:
			return false
		}
	}
	return cells > 0 && dashes >= 3
}

// scanListItem colorizes a bullet or ordered list marker and the task list
// checkbox that may follow it.
func (s *mdScanner) scanListItem(ln int, line string, indent int) bool {
	item, ok := parseListItem(line, indent)
	if !ok {
		return false
	}
	s.mark(ln, indent, item.marker, code.TtListMarker)
	s.listContent = item.content

	text := item.content
	if item.checkbox >= 0 {
		s.mark(ln, item.checkbox-1, item.checkbox+2, code.TtConstant)
		text = item.checkbox + 2
	}
	s.scanInline(ln, line, text, len(line))
	return true
}

// listItem describes a list item line.
type listItem struct {
	marker   int // end of the bullet or the number with its delimiter
	content  int // start of the text after the marker
	checkbox int // index of the mark within the task list checkbox, or -1
}

// parseListItem recognizes the list item starting at indent: a bullet (-, +, *), or a number
// of up to 9 digits followed by . or ), then a space or the end of the line. The item text
// may start with a task list checkbox: [ ], [x], or [X] followed by a space or the end of the line.
func parseListItem(line string, indent int) (listItem, bool) {
	i := indent
	switch line[i] {
	case '-', '+', '*':
		i++
	default:
		for i < len(line) && line[i] >= '0' && line[i] <= '9' {
			i++
		}
		if i == indent || i-indent > 9 || i >= len(line) || (line[i] != '.' && line[i] != ')') {
			return listItem{}, false
		}
		i++
	}
	// The marker has to be followed by a space, or be the whole line. This is
	// also what keeps "*emphasis*" from looking like a bullet.
	if i < len(line) && !isSpaceByte(line[i]) {
		return listItem{}, false
	}
	item := listItem{marker: i, content: skipSpace(line, i), checkbox: -1}

	if rest := line[item.content:]; len(rest) >= 3 && rest[0] == '[' && rest[2] == ']' &&
		(rest[1] == ' ' || rest[1] == 'x' || rest[1] == 'X') &&
		(len(rest) == 3 || isSpaceByte(rest[3])) {
		item.checkbox = item.content + 1
	}
	return item, true
}

// scanLinkDefinition colorizes '[label]: https://example.com "Title"'.
func (s *mdScanner) scanLinkDefinition(ln int, line string, indent int) bool {
	if line[indent] != '[' {
		return false
	}
	label := strings.Index(line[indent:], "]:")
	if label < 0 {
		return false
	}
	label += indent

	s.mark(ln, indent, indent+1, code.TtPunctuation)
	s.mark(ln, indent+1, label, code.TtLink)
	s.mark(ln, label, label+2, code.TtPunctuation)

	dest := skipSpace(line, label+2)
	destEnd := skipNonSpace(line, dest, len(line))
	s.mark(ln, dest, destEnd, code.TtURL)
	s.mark(ln, skipSpace(line, destEnd), len(line), code.TtStringLiteral)
	return true
}

// scanInline highlights the inline markup in line[from:to].
func (s *mdScanner) scanInline(ln int, line string, from, to int) {
	s.match(ln, line)
	for i := from; i < to; {
		var next int
		switch line[i] {
		case '\\':
			// A backslash escape hides the punctuation character after it.
			if i+1 < to && isMarkdownPunct(line[i+1]) {
				s.mark(ln, i, i+2, code.TtEscape)
				next = i + 2
			}
		case '`':
			next = s.scanCodeSpan(ln, line, i, to)
		case '*', '_', '~':
			next = s.scanEmphasis(ln, line, i, to)
		case '[', '!':
			next = s.scanLink(ln, line, i, to)
		case '<':
			next = s.scanAngle(ln, line, i, to)
		case 'h':
			next = s.scanBareURL(ln, line, i, to)
		}
		if next > i {
			i = next
		} else {
			i++
		}
	}
}

// scanCodeSpan colorizes a `code span` matched by matchCodeSpans.
func (s *mdScanner) scanCodeSpan(ln int, line string, i, to int) int {
	open := i
	i = skipRun(line, i, to, '`')
	end := s.codeEnd[open]
	if end < 0 || end > to {
		// The run is literal text as a whole, and its suffixes are not scanned again.
		return i
	}
	closing := end - (i - open)
	s.mark(ln, open, i, code.TtPunctuation)
	s.mark(ln, i, closing, code.TtRawText)
	s.mark(ln, closing, end, code.TtPunctuation)
	return end
}

// scanEmphasis colorizes *italic*, **bold** and ~~struck~~ runs. It follows the
// spirit of the CommonMark flanking rules rather than their letter: a delimiter
// opens when non-space text follows it and closes when non-space text precedes
// it, and '_' additionally has to sit on a word boundary so that snake_case
// identifiers survive.
func (s *mdScanner) scanEmphasis(ln int, line string, i, to int) int {
	c := line[i]
	open := i
	i = skipRun(line, i, to, c)
	n := i - open

	// The runs that can't open are skipped as a whole: rescanning each suffix
	// of a long run is quadratic.
	switch {
	case c == '~' && n != 2: // only ~~strikethrough~~ exists
		return i
	case n > 3:
		return i
	case i >= to || isSpaceByte(line[i]):
		return open // a closing or a lone run, not an opener
	case c == '_' && isWordByte(byteAt(line, open-1)):
		return open
	}

	kind := strings.IndexByte("*_~", c)
	for m := 1; m <= n; m++ {
		// A closer of a longer run closes a shorter one too.
		if s.noEmphasis[kind][m].covers(ln, i, to) {
			return open
		}
	}

	token := code.TtStrong
	switch {
	case c == '~':
		// The text styles have no strikethrough, so struck text borrows the
		// dimmed color comments use.
		token = code.TtComment
	case n == 1:
		token = code.TtEmphasis
	}

	for j := i; j < to; j++ {
		// The delimiters in the code spans and the escaped ones are literal text.
		if end := s.skipLiteral(line, j, to); end > j {
			j = end - 1
			continue
		}
		if line[j] != c {
			continue
		}
		run := skipRun(line, j, to, c) - j
		if run >= n && !isSpaceByte(line[j-1]) &&
			(c != '_' || !isWordByte(byteAt(line, j+run))) {
			s.mark(ln, open, i, code.TtPunctuation)
			s.mark(ln, i, j, token)
			s.mark(ln, j, j+n, code.TtPunctuation)
			s.scanInline(ln, line, i, j)
			return j + n
		}
		j += run - 1
	}
	s.noEmphasis[kind][n] = failure{ln, i, to}
	return open
}

// scanLink colorizes [label](dest "Title") and ![image](dest), along with the
// reference forms [label][ref] and [ref].
func (s *mdScanner) scanLink(ln int, line string, i, to int) int {
	start := i
	if s.linkDepth >= maxLinkDepth {
		return start
	}
	if line[i] == '!' {
		if i+1 >= to || line[i+1] != '[' {
			return start
		}
		i++
	}

	label := i + 1
	end := s.brackets[i]
	if end < 0 || end >= to {
		return start
	}

	s.mark(ln, start, label, code.TtPunctuation) // the '!' and the '['
	s.mark(ln, label, end, code.TtLink)
	s.linkDepth++
	s.scanInline(ln, line, label, end)
	s.linkDepth--
	s.mark(ln, end, end+1, code.TtPunctuation)
	i = end + 1

	if i >= to {
		return i
	}
	switch line[i] {
	case '(':
		return s.scanLinkTarget(ln, line, i, to, ')')
	case '[':
		return s.scanLinkTarget(ln, line, i, to, ']')
	}
	return i // a shortcut reference link, '[ref]' on its own
}

// match matches the code spans and the brackets of the line, once for each line.
func (s *mdScanner) match(ln int, line string) {
	if s.matchedLine == ln+1 {
		return
	}
	s.matchedLine = ln + 1
	s.codeEnd = fill(s.codeEnd, len(line), -1)
	s.brackets = fill(s.brackets, len(line), -1)
	s.matchCodeSpans(line)

	// The brackets in the code spans and the escaped ones are literal text.
	stack := s.bracketStack[:0]
	for j := 0; j < len(line); j++ {
		if end := s.skipLiteral(line, j, len(line)); end > j {
			j = end - 1
			continue
		}
		switch line[j] {
		case '[':
			stack = append(stack, j)
		case ']':
			if len(stack) > 0 {
				s.brackets[stack[len(stack)-1]] = j
				stack = stack[:len(stack)-1]
			}
		}
	}
	s.bracketStack = stack
}

// matchCodeSpans finds the code spans of the line in linear time. A run of backticks
// opens a code span closed by the next run of exactly the same length, and is literal
// text if there is none. A backslash escapes the first backtick of an opening run, but
// not of a closing one: the text of a code span is literal. Code spans are treated as
// line local, which is a simplification: CommonMark lets them wrap.
func (s *mdScanner) matchCodeSpans(line string) {
	runs := s.ticks[:0]
	for i := 0; i < len(line); {
		if line[i] != '`' {
			i++
			continue
		}
		end := skipRun(line, i, len(line), '`')
		runs = append(runs, tickRun{start: i, end: end})
		i = end
	}
	s.ticks = runs
	if len(runs) == 0 {
		return
	}

	if s.lastTicks == nil {
		s.lastTicks = make(map[int]int)
	}
	last := s.lastTicks // the first run of each length after the current one
	clear(last)
	following := func(n int) int {
		if r, ok := last[n]; ok {
			return r
		}
		return -1
	}
	for r := len(runs) - 1; r >= 0; r-- {
		n := runs[r].end - runs[r].start
		runs[r].same, runs[r].shorter = following(n), following(n-1)
		last[n] = r
	}

	r := 0
	for i := 0; i < len(line); i++ {
		switch line[i] {
		case '\\':
			if i+1 < len(line) && isMarkdownPunct(line[i+1]) {
				i++
			}
		case '`':
			for runs[r].end <= i {
				r++
			}
			// The opener is the whole run, or the run with its first backtick escaped.
			next := runs[r].same
			if i > runs[r].start {
				next = runs[r].shorter
			}
			if next < 0 {
				i = runs[r].end - 1
				continue
			}
			s.codeEnd[i] = runs[next].end
			i = runs[next].end - 1
		}
	}
}

// skipLiteral returns the end of the backslash escape or the code span at i, if it ends
// by to, or i otherwise.
func (s *mdScanner) skipLiteral(line string, i, to int) int {
	switch line[i] {
	case '\\':
		if i+1 < to && isMarkdownPunct(line[i+1]) {
			return i + 2
		}
	case '`':
		if end := s.codeEnd[i]; end >= 0 && end <= to {
			return end
		}
	}
	return i
}

// scanLinkTarget colorizes the '(destination "Title")' or '[reference]' that
// follows a link label, up to the matching close byte.
func (s *mdScanner) scanLinkTarget(ln int, line string, i, to int, close byte) int {
	missing := &s.noLinkTarget[strings.IndexByte(")]", close)]
	if missing.covers(ln, i+1, to) {
		return i
	}
	end := -1
	for j := i + 1; j < to && end < 0; j++ {
		switch line[j] {
		case '\\':
			j++
		case close:
			end = j
		}
	}
	if end < 0 {
		*missing = failure{ln, i + 1, to}
		return i
	}

	s.mark(ln, i, i+1, code.TtPunctuation)
	destEnd := skipNonSpace(line, i+1, end)
	s.mark(ln, i+1, destEnd, code.TtURL)
	// Anything after the destination is the optional quoted title.
	s.mark(ln, destEnd, end, code.TtStringLiteral)
	s.mark(ln, end, end+1, code.TtPunctuation)
	return end + 1
}

// scanAngle colorizes an <https://autolink> or an inline HTML tag.
func (s *mdScanner) scanAngle(ln int, line string, i, to int) int {
	if s.noAngle.covers(ln, i, to) {
		return i
	}
	end := strings.IndexByte(line[i:to], '>')
	if end < 0 {
		s.noAngle = failure{ln, i, to}
		return i
	}
	if end == 1 { // an empty '<>'
		return i
	}
	end += i

	body := line[i+1 : end]
	token := code.TtURL
	if !strings.Contains(body, "://") && !strings.Contains(body, "@") {
		token = code.TtPunctuation // an HTML tag rather than an autolink
	}
	s.mark(ln, i, i+1, code.TtPunctuation)
	s.mark(ln, i+1, end, token)
	s.mark(ln, end, end+1, code.TtPunctuation)
	return end + 1
}

// scanBareURL colorizes a URL written with no markup around it.
func (s *mdScanner) scanBareURL(ln int, line string, i, to int) int {
	if isWordByte(byteAt(line, i-1)) {
		return i
	}
	var scheme int
	switch {
	case strings.HasPrefix(line[i:to], "https://"):
		scheme = len("https://")
	case strings.HasPrefix(line[i:to], "http://"):
		scheme = len("http://")
	default:
		return i
	}

	end := i + scheme
	for end < to && !isSpaceByte(line[end]) && line[end] != '>' && line[end] != ')' {
		end++
	}
	// Sentence punctuation that happens to trail a URL is not part of it.
	for end > i+scheme && strings.IndexByte(".,;:!?", line[end-1]) >= 0 {
		end--
	}
	s.mark(ln, i, end, code.TtURL)
	return end
}

func (s *mdScanner) mark(ln, start, end int, token code.TokenType) {
	if start >= end {
		return
	}
	s.emit(Span{Line: ln, Start: start, End: end, Token: token})
}

// line returns the line with the given number, or an empty one if it's outside the document.
func (s *mdScanner) line(n int) string {
	if n < 0 || n >= len(s.lines) {
		return ""
	}
	return s.lines[n]
}

func isSpaceByte(b byte) bool { return b == ' ' || b == '\t' }

// skipSpace returns the index of the first byte at or after i that is not a space.
func skipSpace(line string, i int) int {
	for i < len(line) && isSpaceByte(line[i]) {
		i++
	}
	return i
}

// skipNonSpace returns the index of the first space at or after i, or to.
func skipNonSpace(line string, i, to int) int {
	for i < to && !isSpaceByte(line[i]) {
		i++
	}
	return i
}

// skipRun returns the end of the run of c starting at i, up to to.
func skipRun(line string, i, to int, c byte) int {
	for i < to && line[i] == c {
		i++
	}
	return i
}

// skipQuotes skips up to max block quote markers, each a '>' with the spaces before it,
// returning the index after the last of them and their number.
func skipQuotes(line string, max int) (end, n int) {
	for n < max {
		i := skipSpace(line, end)
		if i == len(line) || line[i] != '>' {
			break
		}
		end, n = i+1, n+1
	}
	return end, n
}

// isBlank reports whether s holds nothing but spaces.
func isBlank(s string) bool { return strings.TrimRight(s, " \t") == "" }

// width returns the number of columns s takes, with the tab stops every 4 columns.
func width(s string) int {
	w := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\t' {
			w += 4 - w%4
		} else {
			w++
		}
	}
	return w
}

// fill returns buf resized to n elements, each set to v.
func fill(buf []int, n, v int) []int {
	buf = slices.Grow(buf[:0], n)[:n]
	for i := range buf {
		buf[i] = v
	}
	return buf
}

// isWordByte reports whether b can be part of a word, which is what decides
// whether an '_' delimiter or a bare URL sits on a word boundary. Every byte of
// a multi-byte rune counts as a word byte.
func isWordByte(b byte) bool {
	return b == '_' || b >= '0' && b <= '9' ||
		b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b >= utf8.RuneSelf
}

// byteAt returns the byte at i, or 0 for an index outside the line, which reads
// as a word boundary.
func byteAt(line string, i int) byte {
	if i < 0 || i >= len(line) {
		return 0
	}
	return line[i]
}

func isMarkdownPunct(b byte) bool {
	return strings.IndexByte(`\`+"`"+`*_{}[]()#+-.!<>|~"'$%&,/:;=?@^`, b) >= 0
}
