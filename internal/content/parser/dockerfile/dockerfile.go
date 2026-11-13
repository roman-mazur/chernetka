// Package dockerfile scans Dockerfiles with a hand written scanner following the
// parser of BuildKit, the builder behind docker build.
//
// A Dockerfile is a list of instructions, one per line, and a line ending with
// the escape character continues on the next one. The scanner reads the
// document the way BuildKit does: it takes the parser directives at its top,
// which may change the escape character, skips the comments and the empty lines
// between the continued lines, and reads the bodies of the heredocs following an
// instruction. Then it colorizes each instruction by its kind. It never fails:
// what BuildKit would reject is highlighted as far as it makes sense, or left
// unhighlighted.
//
// The commands an instruction runs, like the script of RUN, are not parsed as
// shell scripts: only their strings and their variables are highlighted.
package dockerfile

import (
	"encoding/json"
	"regexp"
	"sort"
	"strings"
	"unicode"

	"rmazur.io/chernetka/internal/content/code"
)

// Span is a range of a line, [Start, End) in bytes, recognized as the Token.
type Span struct {
	Line, Start, End int
	Token            code.TokenType
}

// Scan walks the document lines reporting the recognized ranges to emit. The spans
// of the nested constructs, like a variable in a string, follow the ones they are
// nested in.
func Scan(lines []string, emit func(Span)) {
	s := scanner{lines: lines, emit: emit, esc: defaultEscape, syntaxLine: -1}
	s.run()
}

const (
	defaultEscape = '\\'
	bom           = "\xef\xbb\xbf"

	// maxNesting is how deep the variables nested in the defaults of the others,
	// like ${A:-${B}}, are highlighted. It keeps the text scanned again for each
	// level within a limit.
	maxNesting = 4
)

// instructions are the instructions BuildKit knows.
var instructions = map[string]bool{
	"ADD": true, "ARG": true, "CMD": true, "COPY": true, "ENTRYPOINT": true,
	"ENV": true, "EXPOSE": true, "FROM": true, "HEALTHCHECK": true, "LABEL": true,
	"MAINTAINER": true, "ONBUILD": true, "RUN": true, "SHELL": true,
	"STOPSIGNAL": true, "USER": true, "VOLUME": true, "WORKDIR": true,
}

var (
	// directiveRegexp matches a parser directive after its # prefix, as BuildKit does.
	directiveRegexp = regexp.MustCompile(`^([a-zA-Z][a-zA-Z0-9]*)\s*=\s*(.+?)\s*$`)
	// heredocRegexp matches a word opening a heredoc, like <<EOF or 2<<-"EOF".
	heredocRegexp = regexp.MustCompile(`^(\d*)<<(-?)\s*([^<]*)$`)
)

// scanner walks a document instruction by instruction. An instruction is
// scanned as its logical line, the continued lines joined, which is mapped
// back to the document by its segments.
type scanner struct {
	lines []string
	emit  func(Span)

	esc        byte // the escape character, set by the escape directive
	syntaxLine int  // the syntax directive following a shebang, or -1

	t        string    // the logical line being scanned
	text     []byte    // the buffer the logical line is joined in
	segs     []segment // the pieces of the document t is joined from
	close    []int     // the index of the '}' closing each "${" of t, or -1
	stack    []int
	heredocs []heredoc // the heredocs opened by the instruction
}

// segment is a piece of the logical line: t[off:] starts at lines[line][col:].
type segment struct{ off, line, col int }

// word is the range t[start:end] of a word of the logical line.
type word struct{ start, end int }

// heredoc is a heredoc opened by an instruction, whose body follows it.
type heredoc struct {
	name   string // the word terminating the body
	chomp  bool   // whether the leading tabs of the terminator are ignored, for <<-
	expand bool   // whether the variables of the body are expanded: its word is not quoted
}

func (s *scanner) run() {
	ln := s.directives()
	for ln < len(s.lines) {
		if ln == s.syntaxLine {
			ln++
			continue
		}
		start, end := s.bounds(ln)
		i := end - len(strings.TrimLeftFunc(s.lines[ln][start:end], unicode.IsSpace))
		switch {
		case i == end:
			ln++
		case s.lines[ln][i] == '#':
			s.emitLine(ln, i, end, code.TtComment)
			ln++
		default:
			ln = s.instruction(ln)
		}
	}
}

// bounds returns the range of the line holding its content, without the byte
// order mark of the document and the carriage returns of a CRLF line ending.
func (s *scanner) bounds(ln int) (start, end int) {
	line := s.lines[ln]
	if ln == 0 && strings.HasPrefix(line, bom) {
		start = len(bom)
	}
	return start, start + len(strings.TrimRight(line[start:], "\r\n"))
}

func (s *scanner) emitLine(ln, start, end int, token code.TokenType) {
	if start < end {
		s.emit(Span{Line: ln, Start: start, End: end, Token: token})
	}
}

// directives highlights the parser directives at the top of the document and
// returns the number of the lines they take. Like BuildKit, it takes the
// lines starting with # as directives until the first line that is not a
// known directive: an instruction, an empty line, a comment, an unknown or a
// repeated directive, or an escape directive with an invalid character.
func (s *scanner) directives() int {
	seen := make(map[string]bool)
	ln := 0
	for ; ln < len(s.lines); ln++ {
		start, end := s.bounds(ln)
		line := s.lines[ln][start:end]
		trimmed := strings.TrimLeftFunc(line, unicode.IsSpace)
		rest, ok := strings.CutPrefix(trimmed, "#")
		if !ok {
			break
		}
		name, value, ok := parseDirective(rest)
		if !ok || seen[name] {
			break
		}
		if name == "escape" {
			if value != "`" && value != `\` {
				break
			}
			s.esc = value[0]
		}
		seen[name] = true
		s.emitLine(ln, end-len(trimmed), end, code.TtDirective)
	}
	if ln == 0 && len(s.lines) > 0 {
		return s.shebang()
	}
	return ln
}

// shebang highlights the shebang of a document starting with one, and returns
// the number of its lines. BuildKit takes no parser directives after it, but
// the syntax one: the frontend running the build is detected skipping the
// shebang. The other directives following it are comments.
func (s *scanner) shebang() int {
	start, end := s.bounds(0)
	if !strings.HasPrefix(s.lines[0][start:end], "#!") {
		return 0
	}
	s.emitLine(0, start, end, code.TtDirective)
	seen := make(map[string]bool)
	for ln := 1; ln < len(s.lines); ln++ {
		// Unlike the parser, the syntax detection takes no space before the #.
		start, end := s.bounds(ln)
		rest, ok := strings.CutPrefix(s.lines[ln][start:end], "#")
		if !ok {
			break
		}
		name, _, ok := parseDirective(rest)
		if !ok || seen[name] {
			break
		}
		seen[name] = true
		if name == "syntax" {
			s.emitLine(ln, start, end, code.TtDirective)
			s.syntaxLine = ln
		}
	}
	return 1
}

// parseDirective parses a line with its # prefix cut as a parser directive, and
// returns its lower case name and its value.
func parseDirective(text string) (name, value string, ok bool) {
	m := directiveRegexp.FindStringSubmatch(strings.TrimLeftFunc(text, unicode.IsSpace))
	if m == nil {
		return "", "", false
	}
	switch name = strings.ToLower(m[1]); name {
	case "syntax", "escape", "check":
		return name, m[2], true
	}
	return "", "", false
}

// instruction highlights the instruction starting on the line, with the lines
// continuing it and the bodies of its heredocs, and returns the line after them.
func (s *scanner) instruction(ln int) int {
	s.text, s.segs, s.heredocs = s.text[:0], s.segs[:0], s.heredocs[:0]
	first := true
	for ; ln < len(s.lines); ln++ {
		start, end := s.bounds(ln)
		line := s.lines[ln]
		trimmed := strings.TrimLeftFunc(line[start:end], unicode.IsSpace)
		col := start
		if first {
			col = end - len(trimmed)
		} else if trimmed == "" {
			continue // BuildKit warns about the empty continuation lines
		} else if trimmed[0] == '#' {
			s.emitLine(ln, end-len(trimmed), end, code.TtComment)
			continue
		}
		first = false
		stop, continued := s.continuation(line[col:end])
		s.segs = append(s.segs, segment{off: len(s.text), line: ln, col: col})
		s.text = append(s.text, line[col:col+stop]...)
		if !continued {
			ln++
			break
		}
	}
	s.t = string(s.text)
	s.match()
	s.instr(0, false)
	return s.heredocBodies(ln)
}

// continuation reports whether the line is continued on the next one, ending
// with the escape character that is not escaped itself, and returns the end
// of its content.
func (s *scanner) continuation(line string) (stop int, continued bool) {
	r := strings.TrimRight(line, " \t")
	if n := len(r); n > 0 && r[n-1] == s.esc && (n == 1 || r[n-2] != s.esc) {
		return n - 1, true
	}
	return len(line), false
}

// setLine makes the content of a single line the logical line to scan.
func (s *scanner) setLine(ln, start, end int) {
	s.t = s.lines[ln][start:end]
	s.segs = append(s.segs[:0], segment{line: ln, col: start})
	s.match()
}

// match finds the '}' closing each "${" of the logical line, once: searching
// for the closer of each opener is quadratic.
func (s *scanner) match() {
	s.close = s.close[:0]
	for range len(s.t) {
		s.close = append(s.close, -1)
	}
	stack := s.stack[:0]
	for i := 0; i < len(s.t); i++ {
		switch s.t[i] {
		case s.esc:
			i++
		case '$':
			if i+1 < len(s.t) && s.t[i+1] == '{' {
				stack = append(stack, i)
				i++
			}
		case '}':
			if n := len(stack); n > 0 {
				s.close[stack[n-1]] = i
				stack = stack[:n-1]
			}
		}
	}
	s.stack = stack
}

// mark emits the range of the logical line, split by the lines it spans.
func (s *scanner) mark(from, to int, token code.TokenType) {
	if from >= to {
		return
	}
	i := sort.Search(len(s.segs), func(i int) bool { return s.segs[i].off > from }) - 1
	for ; i < len(s.segs) && s.segs[i].off < to; i++ {
		seg := s.segs[i]
		end := len(s.t)
		if i+1 < len(s.segs) {
			end = s.segs[i+1].off
		}
		if a, b := max(from, seg.off), min(to, end); a < b {
			s.emit(Span{Line: seg.line, Start: seg.col + a - seg.off, End: seg.col + b - seg.off, Token: token})
		}
	}
}

// instr highlights the instruction starting at the index of the logical line.
// The instruction may be the trigger of an ONBUILD one.
func (s *scanner) instr(at int, trigger bool) {
	i := s.skipSpace(at)
	j := i
	for j < len(s.t) && !isSpace(s.t[j]) {
		j++
	}
	name := strings.ToUpper(s.t[i:j])
	if !instructions[name] {
		return
	}
	if trigger && (name == "ONBUILD" || name == "FROM" || name == "MAINTAINER") {
		return // not allowed as triggers
	}
	s.mark(i, j, code.TtKeyword)
	i = s.skipSpace(j)
	if name == "ONBUILD" {
		s.instr(i, true)
		return
	}

	i = s.flags(i)
	switch name {
	case "FROM":
		s.from(i)
	case "ARG":
		s.pairs(i, false)
	case "ENV", "LABEL":
		s.pairs(i, true)
	case "EXPOSE":
		s.expose(i)
	case "HEALTHCHECK":
		s.healthcheck(i)
	case "RUN", "CMD", "ENTRYPOINT", "SHELL":
		s.command(i)
	case "MAINTAINER":
		s.shell(i, len(s.t), false)
	default: // ADD, COPY, STOPSIGNAL, USER, VOLUME, WORKDIR
		if s.isJSON(i) {
			s.json(i, true)
		} else {
			s.shell(i, len(s.t), true)
		}
	}
	if name == "RUN" || name == "COPY" || name == "ADD" {
		s.heredocWords(i)
	}
}

// flags highlights the flags of an instruction, like --platform=linux/amd64,
// which precede its arguments, and returns the index after them.
func (s *scanner) flags(i int) int {
	for strings.HasPrefix(s.t[i:], "--") {
		end := s.wordEnd(i, s.esc)
		name := end
		if eq := strings.IndexByte(s.t[i:end], '='); eq >= 0 {
			name = i + eq
			s.shell(name+1, end, true)
		}
		s.mark(i, name, code.TtConstant)
		i = s.skipSpace(end)
	}
	return i
}

// from highlights the image of a stage, as an import, and the stage name
// declared after AS.
func (s *scanner) from(i int) {
	ws := s.words(i, s.esc)
	if len(ws) == 0 {
		return
	}
	s.mark(ws[0].start, ws[0].end, code.TtImportRef)
	s.shell(ws[0].start, ws[0].end, true)
	if len(ws) > 1 && strings.EqualFold(s.t[ws[1].start:ws[1].end], "AS") {
		s.mark(ws[1].start, ws[1].end, code.TtKeyword)
		if len(ws) > 2 {
			s.mark(ws[2].start, ws[2].end, code.TtFuncDeclaration)
		}
	}
}

// pairs highlights the names and the values of ARG, ENV and LABEL, like
// key=value. Like BuildKit, a word is split at its first '=', wherever it is.
// The legacy form of ENV and LABEL, the name and the value separated by a
// space, is taken when the first word has no '='.
func (s *scanner) pairs(i int, legacy bool) {
	ws := s.words(i, s.esc)
	if len(ws) == 0 {
		return
	}
	if legacy && !strings.Contains(s.t[ws[0].start:ws[0].end], "=") {
		s.mark(ws[0].start, ws[0].end, code.TtField)
		if len(ws) > 1 {
			s.shell(ws[1].start, len(s.t), true)
		}
		return
	}
	for _, w := range ws {
		eq := strings.IndexByte(s.t[w.start:w.end], '=')
		if eq < 0 {
			if !legacy {
				s.mark(w.start, w.end, code.TtField) // ARG with no default
			}
			continue
		}
		s.mark(w.start, w.start+eq, code.TtField)
		s.shell(w.start+eq+1, w.end, true)
	}
}

// expose highlights the ports of EXPOSE, like 80/tcp or 8000-8010.
func (s *scanner) expose(i int) {
	for _, w := range s.words(i, s.esc) {
		s.shell(w.start, w.end, true)
		n := s.digits(w.start, w.end)
		if n > w.start && n < w.end && s.t[n] == '-' {
			if m := s.digits(n+1, w.end); m > n+1 {
				n = m
			}
		}
		s.mark(w.start, n, code.TtNumberLiteral)
	}
}

func (s *scanner) digits(i, end int) int {
	for i < end && '0' <= s.t[i] && s.t[i] <= '9' {
		i++
	}
	return i
}

// healthcheck highlights the command of HEALTHCHECK, or its NONE.
func (s *scanner) healthcheck(i int) {
	end := s.wordEnd(i, s.esc)
	switch strings.ToUpper(s.t[i:end]) {
	case "CMD":
		s.mark(i, end, code.TtKeyword)
		s.command(s.skipSpace(end))
	case "NONE":
		s.mark(i, end, code.TtKeyword)
	}
}

// command highlights the command of RUN, CMD, ENTRYPOINT, SHELL or HEALTHCHECK.
// In the shell form the shell expands its variables, while in the exec form,
// a JSON array, nothing does.
func (s *scanner) command(i int) {
	if s.isJSON(i) {
		s.json(i, false)
	} else {
		s.shell(i, len(s.t), true)
	}
}

// heredocWords highlights the words opening the heredocs of RUN, COPY or ADD,
// like <<EOF, and records the heredocs whose bodies follow the instruction.
// Like BuildKit, the words are split with \ escaping, whatever the escape
// character, and a heredoc whose word is quoted is not expanded.
func (s *scanner) heredocWords(i int) {
	if !strings.Contains(s.t[i:], "<<") {
		return
	}
	for _, w := range s.words(i, '\\') {
		m := heredocRegexp.FindStringSubmatch(s.t[w.start:w.end])
		if m == nil || m[3] == "" {
			continue
		}
		name, quoted := unquote(m[3])
		if name == "" {
			continue
		}
		s.heredocs = append(s.heredocs, heredoc{name: name, chomp: m[2] == "-", expand: !quoted})
		s.mark(w.start, w.end, code.TtStringLiteral)
	}
}

// unquote removes the quotes and the escapes of a word, and reports whether it was quoted.
func unquote(w string) (string, bool) {
	var b strings.Builder
	quoted := false
	for i := 0; i < len(w); i++ {
		switch c := w[i]; c {
		case '\\':
			if i+1 < len(w) {
				i++
				b.WriteByte(w[i])
			}
		case '\'':
			quoted = true
			end := strings.IndexByte(w[i+1:], '\'')
			if end < 0 {
				end = len(w) - i - 1
			}
			b.WriteString(w[i+1 : i+1+end])
			i += end + 1
		case '"':
			quoted = true
			for i++; i < len(w) && w[i] != '"'; i++ {
				if w[i] == '\\' && i+1 < len(w) && strings.IndexByte("\"$\\", w[i+1]) >= 0 {
					i++
				}
				b.WriteByte(w[i])
			}
		default:
			b.WriteByte(c)
		}
	}
	return b.String(), quoted
}

// heredocBodies highlights the bodies of the heredocs opened by the instruction,
// one after another starting on the line, and returns the line after them.
func (s *scanner) heredocBodies(ln int) int {
	for _, h := range s.heredocs {
		for ; ln < len(s.lines); ln++ {
			start, end := s.bounds(ln)
			terminator := s.lines[ln][start:end]
			if h.chomp {
				terminator = strings.TrimLeft(terminator, "\t")
			}
			s.emitLine(ln, start, end, code.TtStringLiteral)
			if terminator == h.name {
				ln++
				break
			}
			if h.expand {
				s.setLine(ln, start, end)
				s.dollars(0, len(s.t), 0)
			}
		}
	}
	return ln
}

// isJSON reports whether the arguments starting at the index are in the exec
// form, a JSON array of strings. Otherwise, BuildKit takes them in the shell form.
func (s *scanner) isJSON(i int) bool {
	rest := strings.TrimSpace(s.t[i:])
	if !strings.HasPrefix(rest, "[") {
		return false
	}
	var args []string
	return json.Unmarshal([]byte(rest), &args) == nil
}

// json highlights the strings of a JSON array and their escapes, and their
// variables if the instruction expands them.
func (s *scanner) json(i int, expand bool) {
	for k := i; k < len(s.t); k++ {
		if s.t[k] != '"' {
			continue
		}
		end := k + 1
		for end < len(s.t) && s.t[end] != '"' {
			if s.t[end] == '\\' {
				end++
			}
			end++
		}
		end = min(end+1, len(s.t))
		s.mark(k, end, code.TtStringLiteral)
		for j := k + 1; j < end-1; {
			switch s.t[j] {
			case '\\':
				n := 2
				if j+1 < len(s.t) && s.t[j+1] == 'u' {
					n = 6
				}
				s.mark(j, min(j+n, end-1), code.TtEscape)
				j += n
			case '$':
				if expand {
					j = s.expansion(j, end-1, 0)
				} else {
					j++
				}
			default:
				j++
			}
		}
		k = end - 1
	}
}

// shell highlights the strings of t[from:to] in the shell form, and its
// variables if they are expanded. A single quoted string expands nothing.
func (s *scanner) shell(from, to int, expand bool) {
	for k := from; k < to; {
		switch c := s.t[k]; {
		case c == s.esc:
			k += 2
		case c == '\'' || c == '"':
			end := min(s.quoteEnd(k, s.esc), to)
			s.mark(k, end, code.TtStringLiteral)
			if c == '"' && expand {
				s.dollars(k+1, end, 0)
			}
			k = end
		case c == '$' && expand:
			k = s.expansion(k, to, 0)
		default:
			k++
		}
	}
}

// dollars highlights the variables of t[from:to].
func (s *scanner) dollars(from, to, depth int) {
	for k := from; k < to; {
		switch s.t[k] {
		case s.esc:
			k += 2
		case '$':
			k = s.expansion(k, to, depth)
		default:
			k++
		}
	}
}

// expansion highlights the variable at the '$' at the index, like $NAME,
// ${NAME} or ${NAME:-default}, and returns the index after it. The variables
// nested in a default are highlighted too.
func (s *scanner) expansion(k, to, depth int) int {
	j := k + 1
	if j < to && s.t[j] == '{' {
		name := s.nameEnd(j+1, to)
		s.mark(k, j+1, code.TtEscape)
		s.mark(j+1, name, code.TtField)
		end := s.close[k]
		if end < 0 || end >= to {
			return name // BuildKit fails on the missing '}'
		}
		if depth < maxNesting {
			s.dollars(name, end, depth+1)
		}
		s.mark(end, end+1, code.TtEscape)
		return end + 1
	}
	name := s.nameEnd(j, to)
	if name == j {
		return j // a lone $
	}
	s.mark(k, j, code.TtEscape)
	s.mark(j, name, code.TtField)
	return name
}

func (s *scanner) nameEnd(i, to int) int {
	for i < to {
		c := s.t[i]
		if c != '_' && !('a' <= c && c <= 'z') && !('A' <= c && c <= 'Z') && !('0' <= c && c <= '9') {
			break
		}
		i++
	}
	return i
}

// words splits the logical line from the index into words separated by spaces,
// keeping the quoted strings and the escaped characters in them.
func (s *scanner) words(i int, esc byte) []word {
	var ws []word
	for i = s.skipSpace(i); i < len(s.t); i = s.skipSpace(i) {
		end := s.wordEnd(i, esc)
		ws = append(ws, word{i, end})
		i = end
	}
	return ws
}

func (s *scanner) wordEnd(i int, esc byte) int {
	for i < len(s.t) && !isSpace(s.t[i]) {
		switch c := s.t[i]; {
		case c == esc:
			i += 2
		case c == '\'' || c == '"':
			i = s.quoteEnd(i, esc)
		default:
			i++
		}
	}
	return min(i, len(s.t))
}

// quoteEnd returns the index after the string quoted at the index. In double
// quotes, the escape character escapes the quote. An unterminated string runs
// to the end of the line.
func (s *scanner) quoteEnd(i int, esc byte) int {
	q := s.t[i]
	for i++; i < len(s.t); i++ {
		switch c := s.t[i]; {
		case c == q:
			return i + 1
		case c == esc && q == '"':
			i++
		}
	}
	return len(s.t)
}

func (s *scanner) skipSpace(i int) int {
	for i < len(s.t) && isSpace(s.t[i]) {
		i++
	}
	return i
}

func isSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\v' || c == '\f' || c == '\r' || c == '\n'
}
