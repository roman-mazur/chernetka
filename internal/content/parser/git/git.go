// Package git scans the files git opens in an editor: the commit messages and
// the todo lists of the interactive rebases.
//
// Both are line oriented, so they are scanned line by line with hand written
// scanners. Like git, the scanners take the lines starting with '#' as
// comments wherever they are, after spaces too in a rebase todo list.
package git

import (
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
				key, value, found := strings.Cut(line, ":")
				inTrailers = found && !strings.ContainsAny(key, " \t")
				if inTrailers {
					parts = append(parts,
						Span{Line: i, End: len(key) + 1, Token: code.TtField},
						Span{Line: i, Start: len(line) - len(strings.TrimLeft(value, " \t")), End: len(line), Token: code.TtStringLiteral},
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
// like pick or its shortcut p, the commits they take, and the comments to emit.
// Like git, it ignores the spaces before a command or a comment. The subject
// following a commit is left plain, and so are the other arguments, like the
// label of reset or the shell command of exec.
func ScanRebaseTodo(lines []string, emit func(Span)) {
	for i, line := range lines {
		start := len(line) - len(strings.TrimLeft(line, " \t"))
		if strings.HasPrefix(line[start:], "#") {
			emit(Span{Line: i, Start: start, End: len(line), Token: code.TtComment})
			continue
		}

		end := wordEnd(line, start)
		args, ok := rebaseCommands[line[start:end]]
		if !ok {
			continue
		}
		emit(Span{Line: i, Start: start, End: end, Token: code.TtKeyword})

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
				emit(Span{Line: i, Start: ws, End: we, Token: code.TtConstant})
				ws, we = nextWord(line, we, comment)
			} else if args == argsMerge {
				ws = we // a label with no commit
			}
			if ws < we {
				emit(Span{Line: i, Start: ws, End: we, Token: code.TtNumberLiteral})
			}
		}
		if comment < len(line) {
			emit(Span{Line: i, Start: comment, End: len(line), Token: code.TtComment})
		}
	}
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
