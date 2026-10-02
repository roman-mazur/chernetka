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

// The states of parsing a commit message.
const (
	gmBeforeSubject = iota // only empty lines so far
	gmAfterSubject         // expecting the empty line separating the body
	gmBody
	gmNoBody // the subject is not separated from the rest, which is not parsed
)

// reparse finds the subject, the trailers, and the comments of the message. Like git,
// it treats the lines starting with '#' as comments wherever they are, and ignores
// them together with the leading empty lines when looking for the subject and the body.
func (gm *gitMessage) reparse(s *source) error {
	gm.subject, gm.parts = nil, nil
	state := gmBeforeSubject
	afterEmptyLine, insideTags := false, false
	for i, line := range s.lines {
		if strings.HasPrefix(line, "#") {
			gm.parts = append(gm.parts, rawSpan{
				StartLine: i,
				EndLine:   i,
				EndCol:    len(line),
				TokenType: code.TtComment,
			})
			continue
		}

		switch state {
		case gmBeforeSubject:
			if line != "" {
				gm.subject = &rawSpan{StartLine: i, EndLine: i, EndCol: len(line), TokenType: code.TtHeading}
				state = gmAfterSubject
			}

		case gmAfterSubject:
			if line != "" {
				gm.subject = nil // no distinct subject line
				state = gmNoBody
			} else {
				afterEmptyLine = true
				state = gmBody
			}

		case gmBody:
			switch {
			case line == "":
				afterEmptyLine = true

			case afterEmptyLine || insideTags:
				afterEmptyLine = false
				parts := strings.SplitN(line, ":", 2)
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
							EndCol:    len(line),
							TokenType: code.TtStringLiteral,
						},
					)
				}
			}
		}
	}
	return nil
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

func (gr *gitRebase) reparse(s *source) error {
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
	return nil
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
