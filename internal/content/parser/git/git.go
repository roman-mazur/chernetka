// Package git scans the files git opens in an editor: the commit messages and
// the todo lists of the interactive rebases.
//
// Both are line oriented, so they are scanned line by line with hand written
// scanners. Like git, the scanners take the lines starting with '#' as
// comments wherever they are.
package git

import (
	"regexp"
	"strings"

	"rmazur.io/chernetka/internal/content/code"
)

// Span is a range of a line, [Start, End) in bytes, recognized as the Token.
type Span struct {
	Line, Start, End int
	Token            code.TokenType
}

// The states of scanning a commit message.
const (
	beforeSubject = iota // only empty lines so far
	afterSubject         // expecting the empty line separating the body
	inBody
	noBody // the subject is not separated from the rest, which is not scanned
)

// ScanCommitMessage walks the lines of a commit message reporting its subject,
// its trailers, like Signed-off-by, and its comments to emit. The subject is
// reported first. Like git, it ignores the comments together with the leading
// empty lines when looking for the subject and the body.
func ScanCommitMessage(lines []string, emit func(Span)) {
	var (
		subject *Span
		parts   []Span // the trailers and the comments
	)
	state := beforeSubject
	afterEmptyLine, inTrailers := false, false
	for i, line := range lines {
		if strings.HasPrefix(line, "#") {
			parts = append(parts, Span{Line: i, End: len(line), Token: code.TtComment})
			continue
		}

		switch state {
		case beforeSubject:
			if line != "" {
				subject = &Span{Line: i, End: len(line), Token: code.TtHeading}
				state = afterSubject
			}

		case afterSubject:
			if line != "" {
				subject = nil // no distinct subject line
				state = noBody
			} else {
				afterEmptyLine = true
				state = inBody
			}

		case inBody:
			switch {
			case line == "":
				afterEmptyLine = true

			case afterEmptyLine || inTrailers:
				afterEmptyLine = false
				key, _, found := strings.Cut(line, ":")
				inTrailers = found && !strings.ContainsAny(key, " \t")
				if inTrailers {
					parts = append(parts,
						Span{Line: i, End: len(key) + 1, Token: code.TtField},
						Span{Line: i, Start: len(key) + 2, End: len(line), Token: code.TtStringLiteral},
					)
				}
			}
		}
	}

	if subject != nil {
		emit(*subject)
	}
	for _, p := range parts {
		if p.Start < p.End {
			emit(p)
		}
	}
}

// ScanRebaseTodo walks the lines of a rebase todo list reporting its commands,
// like pick or its shortcut p, their arguments, and the comments to emit.
func ScanRebaseTodo(lines []string, emit func(Span)) {
	for i, line := range lines {
		if strings.HasPrefix(line, "#") {
			emit(Span{Line: i, End: len(line), Token: code.TtComment})
			continue
		}

		m := rebaseCommandRegexp.FindStringSubmatchIndex(line)
		if m == nil || !isRebaseCommand(line[m[2]:m[3]]) {
			continue
		}
		for group, token := range []code.TokenType{1: code.TtKeyword, 2: code.TtNumberLiteral, 3: code.TtComment} {
			if group > 0 && m[group*2] < m[group*2+1] {
				emit(Span{Line: i, Start: m[group*2], End: m[group*2+1], Token: token})
			}
		}
	}
}

// rebaseCommandRegexp matches a command line, like "pick 1234 change # note": the
// command, its arguments, and the comment after them.
var rebaseCommandRegexp = regexp.MustCompile(`^(\w+)\s(.+?)(\s?#.*)?$`)

func isRebaseCommand(name string) bool {
	for _, cmd := range rebaseCommands {
		if name == cmd.shortcut || name == cmd.name {
			return true
		}
	}
	return false
}

// rebaseCommands are the commands of a rebase todo list, with their shortcuts.
var rebaseCommands = []struct{ shortcut, name string }{
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
