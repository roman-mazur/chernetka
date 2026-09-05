package extsyntaxhl

import (
	"strings"
	"unicode/utf8"

	"rmazur.io/chernetka/internal/editor"
)

// markdown highlights Markdown with a hand written scanner.
//
// Markdown suits a hand written scanner well: almost every construct is
// anchored to the start of a line, and the two that are not - fenced code
// blocks and setext headings - only need the lines immediately around them.
// The scanner is deliberately pragmatic rather than a complete CommonMark
// implementation. It aims to colorize what someone writing notes actually
// types, and it never fails: text it does not recognize is simply left
// unhighlighted.
type markdown struct{}

func newMarkdown() highlighter { return markdown{} }

// reparse does nothing: scanning a document is cheap enough that all the work
// happens in spans, which also runs once per revision.
func (markdown) reparse(*source) {}

func (markdown) Close() error { return nil }

func (markdown) spans(src *source, emit func(rawSpan)) {
	s := mdScanner{src: src, emit: emit, paragraph: -1}
	s.run()
}

// mdScanner walks a document line by line, carrying the little bit of state
// that spills across line boundaries.
type mdScanner struct {
	src  *source
	emit func(rawSpan)

	// Open code fence, tracked so its body is colorized as raw text.
	inFence     bool
	fenceChar   byte
	fenceLen    int
	fenceIndent int

	paragraph   int  // line number of the last paragraph line, or -1
	inTable     bool // the previous line was part of a table
	listContent int  // content column of the innermost open list item, or 0
}

func (s *mdScanner) run() {
	for ln, line := range s.src.lines {
		if s.inFence {
			// An unterminated fence just runs to the end of the document.
			if s.closesFence(line) {
				s.mark(ln, 0, len(line), editor.TtPunctuation)
				s.inFence = false
			} else {
				s.mark(ln, 0, len(line), editor.TtRawText)
			}
			continue
		}
		s.scanBlock(ln, line, 0)
	}
}

// scanBlock highlights the block construct that starts at or after col. col is
// what lets blocks nest: a block quote colorizes its own '>' and then hands the
// rest of the line back here, so "> ## Heading" is still a heading.
func (s *mdScanner) scanBlock(ln int, line string, col int) {
	indent := col
	for indent < len(line) && isSpaceByte(line[indent]) {
		indent++
	}
	if indent == len(line) {
		// A blank line closes the paragraph and the table above it.
		s.paragraph, s.inTable = -1, false
		return
	}

	// Only the outermost call owns the cross-line state; a nested call is still
	// looking at the same line.
	paragraph := s.paragraph
	if col == 0 {
		s.paragraph = -1
		if indent == 0 {
			s.listContent = 0
		}
		// An indented code block: four columns of indentation, following a blank
		// line, and not the continuation of a list item.
		if indent >= 4 && paragraph < 0 && s.listContent == 0 {
			s.mark(ln, indent, len(line), editor.TtRawText)
			return
		}
	}

	switch {
	case s.scanFenceOpen(ln, line, indent):
	case s.scanHeading(ln, line, indent):
	// A setext underline has to be tried before a thematic break, because '---'
	// under a paragraph line underlines it instead of breaking the page.
	case s.scanSetext(ln, line, indent, paragraph):
	case s.scanThematicBreak(ln, line, indent):
	case s.scanBlockQuote(ln, line, indent):
	case s.scanTableRow(ln, line, indent):
	case s.scanListItem(ln, line, indent):
	case s.scanLinkDefinition(ln, line, indent):
	default:
		s.scanInline(ln, line, indent, len(line))
		s.paragraph = ln
	}
}

// scanFenceOpen recognizes the ``` or ~~~ line that opens a code fence.
func (s *mdScanner) scanFenceOpen(ln int, line string, indent int) bool {
	c := line[indent]
	if c != '`' && c != '~' {
		return false
	}
	i := indent
	for i < len(line) && line[i] == c {
		i++
	}
	if i-indent < 3 {
		return false
	}

	s.mark(ln, indent, i, editor.TtPunctuation)
	// Whatever follows the fence is the info string naming the language.
	s.mark(ln, i, len(line), editor.TtKeyword)

	s.inFence, s.fenceChar, s.fenceLen, s.fenceIndent = true, c, i-indent, indent
	return true
}

// closesFence reports whether line ends the open fence: a run of at least as
// many of the fence's own characters, with nothing else on the line.
func (s *mdScanner) closesFence(line string) bool {
	i := 0
	for i < len(line) && line[i] == ' ' {
		i++
	}
	if i > s.fenceIndent+3 {
		return false
	}
	run := i
	for run < len(line) && line[run] == s.fenceChar {
		run++
	}
	if run-i < s.fenceLen {
		return false
	}
	return strings.TrimRight(line[run:], " \t") == ""
}

// scanHeading recognizes an ATX heading, "## Like this ##".
func (s *mdScanner) scanHeading(ln int, line string, indent int) bool {
	i := indent
	for i < len(line) && line[i] == '#' {
		i++
	}
	if level := i - indent; level == 0 || level > 6 {
		return false
	}
	// The marker has to be followed by a space, or be the whole line.
	if i < len(line) && !isSpaceByte(line[i]) {
		return false
	}
	s.mark(ln, indent, i, editor.TtPunctuation)

	end := len(line)
	for end > i && isSpaceByte(line[end-1]) {
		end--
	}
	// A closed heading ends with its own run of '#', which is punctuation too.
	closing := end
	for closing > i && line[closing-1] == '#' {
		closing--
	}
	if closing < end && (closing == i || isSpaceByte(line[closing-1])) {
		s.mark(ln, closing, end, editor.TtPunctuation)
		end = closing
	}

	// The heading text is colorized as a whole, then the inline markup inside it
	// overrides the narrower ranges it covers.
	s.mark(ln, i, end, editor.TtHeading)
	s.scanInline(ln, line, i, end)
	return true
}

// scanSetext recognizes the '===' or '---' line that underlines the paragraph
// line above it, turning it into a heading.
func (s *mdScanner) scanSetext(ln int, line string, indent, paragraph int) bool {
	if paragraph != ln-1 {
		return false
	}
	c := line[indent]
	if c != '=' && c != '-' {
		return false
	}
	i := indent
	for i < len(line) && line[i] == c {
		i++
	}
	if strings.TrimRight(line[i:], " \t") != "" {
		return false
	}

	s.mark(ln, indent, len(line), editor.TtPunctuation)
	// Promote the line above after the fact. Spans are collected before they are
	// flattened, so emitting one for an earlier line is fine, and the inline
	// spans already reported for it are narrower and still win.
	s.mark(paragraph, 0, len(s.src.line(paragraph)), editor.TtHeading)
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
	s.mark(ln, indent, len(line), editor.TtListMarker)
	return true
}

// scanBlockQuote colorizes a '>' marker and re-dispatches the rest of the line,
// so the blocks nested in a quote are highlighted as themselves.
func (s *mdScanner) scanBlockQuote(ln int, line string, indent int) bool {
	if line[indent] != '>' {
		return false
	}
	s.mark(ln, indent, indent+1, editor.TtPunctuation)
	s.mark(ln, indent+1, len(line), editor.TtQuote)
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
	// The row above the delimiter is the header, and reads as bold.
	header := isTableDelimiterRow(s.src.line(ln + 1))
	if !s.inTable && !delimiter && !header {
		return false
	}
	s.inTable = true

	if delimiter {
		s.mark(ln, indent, len(line), editor.TtPunctuation)
		return true
	}
	cellToken := editor.TtNothing
	if header {
		cellToken = editor.TtStrong
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
		s.mark(ln, i, i+1, editor.TtPunctuation)
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
	i := indent
	switch line[i] {
	case '-', '+', '*':
		i++
	default:
		for i < len(line) && line[i] >= '0' && line[i] <= '9' {
			i++
		}
		if i == indent || i-indent > 9 || i >= len(line) || (line[i] != '.' && line[i] != ')') {
			return false
		}
		i++
	}
	// The marker has to be followed by a space, or be the whole line. This is
	// also what keeps "*emphasis*" from looking like a bullet.
	if i < len(line) && !isSpaceByte(line[i]) {
		return false
	}
	s.mark(ln, indent, i, editor.TtListMarker)

	content := i
	for content < len(line) && isSpaceByte(line[content]) {
		content++
	}
	s.listContent = content

	if rest := line[content:]; len(rest) >= 3 && rest[0] == '[' && rest[2] == ']' &&
		(rest[1] == ' ' || rest[1] == 'x' || rest[1] == 'X') {
		s.mark(ln, content, content+3, editor.TtConstant)
		content += 3
	}
	s.scanInline(ln, line, content, len(line))
	return true
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

	s.mark(ln, indent, indent+1, editor.TtPunctuation)
	s.mark(ln, indent+1, label, editor.TtLink)
	s.mark(ln, label, label+2, editor.TtPunctuation)

	dest := label + 2
	for dest < len(line) && isSpaceByte(line[dest]) {
		dest++
	}
	destEnd := dest
	for destEnd < len(line) && !isSpaceByte(line[destEnd]) {
		destEnd++
	}
	s.mark(ln, dest, destEnd, editor.TtURL)

	title := destEnd
	for title < len(line) && isSpaceByte(line[title]) {
		title++
	}
	s.mark(ln, title, len(line), editor.TtStringLiteral)
	return true
}

// scanInline highlights the inline markup in line[from:to].
func (s *mdScanner) scanInline(ln int, line string, from, to int) {
	for i := from; i < to; {
		var next int
		switch line[i] {
		case '\\':
			// A backslash escape hides the punctuation character after it.
			if i+1 < to && isMarkdownPunct(line[i+1]) {
				s.mark(ln, i, i+2, editor.TtEscape)
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

// scanCodeSpan colorizes a `code span`, whose closing run of backticks has to
// be exactly as long as the one that opened it. Code spans are treated as
// line local, which is a simplification: CommonMark lets them wrap.
func (s *mdScanner) scanCodeSpan(ln int, line string, i, to int) int {
	open := i
	for i < to && line[i] == '`' {
		i++
	}
	n := i - open

	for j := i; j+n <= to; j++ {
		if line[j] != '`' {
			continue
		}
		run := 0
		for j+run < to && line[j+run] == '`' {
			run++
		}
		if run == n {
			s.mark(ln, open, i, editor.TtPunctuation)
			s.mark(ln, i, j, editor.TtRawText)
			s.mark(ln, j, j+n, editor.TtPunctuation)
			return j + n
		}
		j += run - 1
	}
	return open // no closing run on this line
}

// scanEmphasis colorizes *italic*, **bold** and ~~struck~~ runs. It follows the
// spirit of the CommonMark flanking rules rather than their letter: a delimiter
// opens when non-space text follows it and closes when non-space text precedes
// it, and '_' additionally has to sit on a word boundary so that snake_case
// identifiers survive.
func (s *mdScanner) scanEmphasis(ln int, line string, i, to int) int {
	c := line[i]
	open := i
	for i < to && line[i] == c {
		i++
	}
	n := i - open

	switch {
	case c == '~' && n != 2: // only ~~strikethrough~~ exists
		return open
	case n > 3:
		return open
	case i >= to || isSpaceByte(line[i]):
		return open // a closing or a lone run, not an opener
	case c == '_' && isWordByte(byteAt(line, open-1)):
		return open
	}

	token := editor.TtStrong
	switch {
	case c == '~':
		// The terminal writer only sets colors, not text attributes, so struck
		// text borrows the dimmed color comments use.
		token = editor.TtComment
	case n == 1:
		token = editor.TtEmphasis
	}

	for j := i; j < to; j++ {
		if line[j] != c {
			continue
		}
		run := 0
		for j+run < to && line[j+run] == c {
			run++
		}
		if run >= n && !isSpaceByte(line[j-1]) &&
			(c != '_' || !isWordByte(byteAt(line, j+run))) {
			s.mark(ln, open, i, editor.TtPunctuation)
			s.mark(ln, i, j, token)
			s.mark(ln, j, j+n, editor.TtPunctuation)
			s.scanInline(ln, line, i, j)
			return j + n
		}
		j += run - 1
	}
	return open
}

// scanLink colorizes [label](dest "Title") and ![image](dest), along with the
// reference forms [label][ref] and [ref].
func (s *mdScanner) scanLink(ln int, line string, i, to int) int {
	start := i
	if line[i] == '!' {
		if i+1 >= to || line[i+1] != '[' {
			return start
		}
		i++
	}

	label := i + 1
	end := -1
	for j, depth := i, 0; j < to && end < 0; j++ {
		switch line[j] {
		case '\\':
			j++
		case '[':
			depth++
		case ']':
			if depth--; depth == 0 {
				end = j
			}
		}
	}
	if end < 0 {
		return start
	}

	s.mark(ln, start, label, editor.TtPunctuation) // the '!' and the '['
	s.mark(ln, label, end, editor.TtLink)
	s.scanInline(ln, line, label, end)
	s.mark(ln, end, end+1, editor.TtPunctuation)
	i = end + 1

	if i >= to {
		return i
	}
	switch line[i] {
	case '(':
		return s.scanLinkTarget(ln, line, i, to, ')', editor.TtURL)
	case '[':
		return s.scanLinkTarget(ln, line, i, to, ']', editor.TtURL)
	}
	return i // a shortcut reference link, '[ref]' on its own
}

// scanLinkTarget colorizes the '(destination "Title")' or '[reference]' that
// follows a link label, up to the matching close byte.
func (s *mdScanner) scanLinkTarget(ln int, line string, i, to int, close byte, token editor.TokenType) int {
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
		return i
	}

	s.mark(ln, i, i+1, editor.TtPunctuation)
	destEnd := i + 1
	for destEnd < end && !isSpaceByte(line[destEnd]) {
		destEnd++
	}
	s.mark(ln, i+1, destEnd, token)
	// Anything after the destination is the optional quoted title.
	s.mark(ln, destEnd, end, editor.TtStringLiteral)
	s.mark(ln, end, end+1, editor.TtPunctuation)
	return end + 1
}

// scanAngle colorizes an <https://autolink> or an inline HTML tag.
func (s *mdScanner) scanAngle(ln int, line string, i, to int) int {
	end := strings.IndexByte(line[i:to], '>')
	if end <= 1 { // no closer, or an empty '<>'
		return i
	}
	end += i

	body := line[i+1 : end]
	token := editor.TtURL
	if !strings.Contains(body, "://") && !strings.Contains(body, "@") {
		token = editor.TtPunctuation // an HTML tag rather than an autolink
	}
	s.mark(ln, i, i+1, editor.TtPunctuation)
	s.mark(ln, i+1, end, token)
	s.mark(ln, end, end+1, editor.TtPunctuation)
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
	s.mark(ln, i, end, editor.TtURL)
	return end
}

func (s *mdScanner) mark(ln, start, end int, token editor.TokenType) {
	if start >= end {
		return
	}
	s.emit(lineSpan(ln, start, end, token))
}

func isSpaceByte(b byte) bool { return b == ' ' || b == '\t' }

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
