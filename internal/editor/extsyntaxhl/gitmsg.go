package extsyntaxhl

import (
	"regexp"
	"strings"

	"rmazur.io/chernetka/internal/content/code"
)

type gitMessage struct {
	subject *rawSpan
	parts   []rawSpan // footers and comments
}

var _ = highlighter(&gitMessage{}) // ensure interface implementation

func (gm *gitMessage) reparse(s *source) {
	gm.subject, gm.parts = nil, nil
	sLen := len(s.lines)
	if sLen == 0 {
		return
	}
	gm.subject = &rawSpan{
		EndCol:    len(s.line(0)),
		TokenType: code.TtHeading,
	}
	if len(s.lines) != 1 && s.line(1) != "" {
		gm.subject = nil // no distinct subject line
		return
	}
	if sLen <= 2 {
		return
	}
	gm.parseBody(s)
}

func (gm *gitMessage) parseBody(s *source) {
	_ = s.lines[2:] // bounds check

	afterEmptyLine, insideTags := true, false
	for i := 2; i < len(s.lines); i++ {
		switch {
		case s.lines[i] == "":
			afterEmptyLine = true

		case s.lines[i][0] == '#':
			afterEmptyLine = false
			gm.parts = append(gm.parts, rawSpan{
				StartLine: i,
				EndLine:   i,
				EndCol:    len(s.lines[i]),
				TokenType: code.TtComment,
			})

		case afterEmptyLine || insideTags:
			afterEmptyLine = false
			parts := strings.SplitN(s.lines[i], ":", 2)
			insideTags = len(parts) == 2 && !strings.ContainsAny(parts[0], " \t")
			if insideTags {
				gm.parts = append(gm.parts,
					rawSpan{
						StartLine: i,
						EndLine:   i,
						EndCol:    len(parts[0]) + 1,
						TokenType: code.TtField,
					},
					rawSpan{
						StartLine: i,
						StartCol:  len(parts[0]) + 2,
						EndLine:   i,
						EndCol:    len(s.lines[i]),
						TokenType: code.TtStringLiteral,
					},
				)
			}
		}
	}
}

func (gm *gitMessage) spans(_ *source, emit func(rawSpan)) {
	if gm.subject != nil {
		emit(*gm.subject)
	}
	for _, p := range gm.parts {
		emit(p)
	}
}

func (gm *gitMessage) Close() error { return nil }

type gitRebase struct {
	parts []rawSpan
}

func (gr *gitRebase) reparse(s *source) {
	for i, line := range s.lines {
		if strings.HasPrefix(line, "#") {
			gr.parts = append(gr.parts, rawSpan{
				StartLine: i,
				EndLine:   i,
				EndCol:    len(line),
				TokenType: code.TtComment,
			})
			continue
		}

		m := gitRebaseCmdRegexp.FindStringSubmatchIndex(line)
		if len(m) == 0 || !validateGitRebaseCmd(line[m[2]:m[3]]) {
			continue
		}
		gr.parts = append(gr.parts, gitRebaseTokenSpan(m, 1, code.TtKeyword, i))
		if gitRebaseHasMatch(m, 2) {
			gr.parts = append(gr.parts, gitRebaseTokenSpan(m, 2, code.TtNumberLiteral, i))
		}
		if gitRebaseHasMatch(m, 3) {
			gr.parts = append(gr.parts, gitRebaseTokenSpan(m, 3, code.TtComment, i))
		}
	}
}

func (gr *gitRebase) spans(_ *source, emit func(rawSpan)) {
	for _, p := range gr.parts {
		emit(p)
	}
}

func (gr *gitRebase) Close() error { return nil }

func gitRebaseTokenSpan(m []int, group int, token code.TokenType, line int) rawSpan {
	return rawSpan{
		StartLine: line,
		EndLine:   line,
		StartCol:  m[group*2],
		EndCol:    m[group*2+1],
		TokenType: token,
	}
}

func gitRebaseHasMatch(m []int, group int) bool {
	return len(m) >= (group+1)*2 && m[group*2] < m[group*2+1]
}

var gitRebaseCmdRegexp = regexp.MustCompile(`^(\w+)\s(.+?)(\s?#.*)?$`)

func validateGitRebaseCmd(name string) bool {
	for _, cmd := range rebaseCommands {
		if name == cmd.shortcut || name == cmd.name {
			return true
		}
	}
	return false
}

var rebaseCommands = []gitRebaseCmd{
	{"p", "pick"},
	{"r", "reword"},
	{"e", "edit"},
	{"s", "squash"},
	{"f", "fixup"},
	{"x", "exec"},
	{"b", "break"},
	{"d", "drop"},
	{"l", "label"},
	{"t", "reset"},
	{"m", "merge"},
	{"u", "update-ref"},
}

type gitRebaseCmd struct {
	shortcut string
	name     string
}
