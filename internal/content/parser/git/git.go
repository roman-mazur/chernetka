// Package git scans the files git opens in an editor: the commit messages and
// the todo lists of the interactive rebases.
//
// Both are line oriented, so they are scanned line by line with hand written
// scanners. Like git, the scanners take the lines starting with '#' as
// comments wherever they are, after spaces too in a rebase todo list.
package git

import (
	"iter"
	"strings"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/content/code"
)

// token returns the token of the range [start, end) of the line in bytes.
func token(line, start, end int, tt code.TokenType) code.Token {
	return code.Token{
		Start:     content.Position{Line: line, Col: start},
		End:       content.Position{Line: line, Col: end},
		TokenType: tt,
	}
}

// The states of scanning a commit message.
const (
	beforeSubject = iota // only empty lines so far
	afterSubject         // expecting the empty line separating the body
	inBody
	noBody // the subject is not separated from the rest, which is not scanned
)

// ScanCommitMessage returns the tokens of a commit message: its subject, its
// trailers, like Signed-off-by, and its comments. Like git, it ignores the
// comments together with the leading empty lines when looking for the subject
// and the body. So, the subject is yielded once the next other line shows it's
// one, after the comments following it.
func ScanCommitMessage(lines []string) iter.Seq[code.Token] {
	return func(yield func(code.Token) bool) {
		var subject code.Token // the subject waiting for the line after it
		state := beforeSubject
		afterEmptyLine, inTrailers := false, false
		for i, line := range lines {
			if strings.HasPrefix(line, "#") {
				if !yield(token(i, 0, len(line), code.TtComment)) {
					return
				}
				continue
			}

			switch state {
			case beforeSubject:
				if line != "" {
					subject = token(i, 0, len(line), code.TtHeading)
					state = afterSubject
				}

			case afterSubject:
				if line != "" {
					state = noBody // no distinct subject line
					continue
				}
				if !yield(subject) {
					return
				}
				afterEmptyLine = true
				state = inBody

			case inBody:
				switch {
				case line == "":
					afterEmptyLine = true

				case afterEmptyLine || inTrailers:
					afterEmptyLine = false
					key, value, found := strings.Cut(line, ":")
					inTrailers = found && !strings.ContainsAny(key, " \t")
					if !inTrailers {
						continue
					}
					if !yield(token(i, 0, len(key)+1, code.TtField)) {
						return
					}
					if start := len(line) - len(strings.TrimLeft(value, " \t")); start < len(line) &&
						!yield(token(i, start, len(line), code.TtStringLiteral)) {
						return
					}
				}
			}
		}
		if state == afterSubject {
			yield(subject) // the message is just a subject
		}
	}
}

// ScanRebaseTodo returns the tokens of a rebase todo list: its commands, like
// pick or its shortcut p, the commits they take, and the comments. Like git, it
// ignores the spaces before a command or a comment. The subject following a
// commit is left plain, and so are the other arguments, like the label of reset
// or the shell command of exec.
func ScanRebaseTodo(lines []string) iter.Seq[code.Token] {
	return func(yield func(code.Token) bool) {
		for i, line := range lines {
			if !scanRebaseLine(yield, i, line) {
				return
			}
		}
	}
}

// scanRebaseLine yields the tokens of the line i of a rebase todo list. It
// returns false if yield does.
func scanRebaseLine(yield func(code.Token) bool, i int, line string) bool {
	start := len(line) - len(strings.TrimLeft(line, " \t"))
	if strings.HasPrefix(line[start:], "#") {
		return yield(token(i, start, len(line), code.TtComment))
	}

	end := wordEnd(line, start)
	args, ok := rebaseCommands[line[start:end]]
	if !ok {
		return true
	}
	if !yield(token(i, start, end, code.TtKeyword)) {
		return false
	}

	// A comment may follow the arguments, like the subject of merge.
	comment := len(line)
	for j := end; j+1 < len(line); j++ {
		if isSpace(line[j]) && line[j+1] == '#' {
			comment = j + 1
			break
		}
	}

	// The commit, maybe after the -C or -c flag, which merge requires and
	// fixup takes with no value of its own.
	if args != argsOther {
		ws, we := nextWord(line, end, comment)
		if line[ws:we] == "-C" || line[ws:we] == "-c" {
			if !yield(token(i, ws, we, code.TtConstant)) {
				return false
			}
			ws, we = nextWord(line, we, comment)
		} else if args == argsMerge {
			ws = we // a label with no commit
		}
		if ws < we && !yield(token(i, ws, we, code.TtNumberLiteral)) {
			return false
		}
	}
	return comment == len(line) || yield(token(i, comment, len(line), code.TtComment))
}

// nextWord returns the range of the word after the index in line[:end].
func nextWord(line string, i, end int) (start, stop int) {
	for i < end && isSpace(line[i]) {
		i++
	}
	return i, min(wordEnd(line, i), end)
}

func wordEnd(line string, i int) int {
	for i < len(line) && !isSpace(line[i]) {
		i++
	}
	return i
}

func isSpace(c byte) bool { return c == ' ' || c == '\t' }

// rebaseArgs is what a command of a rebase todo list takes.
type rebaseArgs int

const (
	argsOther  rebaseArgs = iota // nothing, a label, a ref, or a shell command
	argsCommit                   // a commit and its subject
	argsMerge                    // maybe a commit after -C or -c, and a label
)

// rebaseCommands are the commands of a rebase todo list, with their shortcuts.
var rebaseCommands = map[string]rebaseArgs{
	"pick": argsCommit, "p": argsCommit,
	"reword": argsCommit, "r": argsCommit,
	"edit": argsCommit, "e": argsCommit,
	"squash": argsCommit, "s": argsCommit,
	"fixup": argsCommit, "f": argsCommit,
	"drop": argsCommit, "d": argsCommit,
	"merge": argsMerge, "m": argsMerge,
	"exec": argsOther, "x": argsOther,
	"break": argsOther, "b": argsOther,
	"label": argsOther, "l": argsOther,
	"reset": argsOther, "t": argsOther,
	"update-ref": argsOther, "u": argsOther,
}
