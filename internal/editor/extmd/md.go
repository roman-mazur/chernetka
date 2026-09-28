// Package extmd implements an editor extension that makes Markdown task list items actionable.
//
// A line like "- [ ] task" can be engaged to check the item turning it into "- [x] task", and back.
package extmd

import (
	"path/filepath"
	"strings"
	"unicode"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor"
)

// Integration implements an editor.Extension that provides line actions for the Markdown task list items.
type Integration struct{}

func (in *Integration) ID() string { return "md" }

func (in *Integration) MakeBufferData(buf *editor.Buffer) editor.BufferExtData {
	switch strings.ToLower(filepath.Ext(buf.Path)) {
	case ".md", ".markdown":
	default:
		return nil
	}
	lines := buf.Content.Lines()
	if len(lines) > 0 && lines[0].MimeType() != content.MimeTypeTextPlain {
		return nil // Not a file, like a directory listing.
	}
	if _, ok := buf.Content.(content.Mutable); !ok {
		return nil
	}
	return &markdown{buf: buf}
}

func (in *Integration) AfterEdit(editor.Sender, *editor.Buffer) {}

type markdown struct {
	buf *editor.Buffer
}

func (md *markdown) LineAction(lineNumber int) content.LineAction {
	lines := md.buf.Content.Lines()
	if lineNumber < 0 || lineNumber >= len(lines) {
		return nil
	}
	if checkboxAt(lines[lineNumber].String()) < 0 {
		return nil
	}
	return toggle{md: md, line: lineNumber}
}

// toggle checks or unchecks the task list item on the line.
type toggle struct {
	md   *markdown
	line int
}

func (t toggle) Engage() {
	lines := t.md.buf.Content.Lines()
	if t.line >= len(lines) {
		return
	}
	text := lines[t.line].String()
	i := checkboxAt(text)
	if i < 0 {
		return
	}
	mark := "x"
	if text[i] != ' ' {
		mark = " "
	}
	t.md.buf.Mutate().Update(t.line, content.TextLine(text[:i]+mark+text[i+1:]))
}

// SingleShot tells the editor to not re-run the toggle as the last engaged action.
func (toggle) SingleShot() {}

// checkboxAt returns the index of the mark within the checkbox of a task list item line, or -1 if it's not one.
// The item starts with a bullet (-, *, +) or a number followed by . or ), then a space and [ ], [x], or [X].
func checkboxAt(line string) int {
	rest := strings.TrimLeftFunc(line, unicode.IsSpace)
	switch {
	case rest == "":
		return -1
	case strings.ContainsRune("-*+", rune(rest[0])):
		rest = rest[1:]
	default:
		digits := strings.TrimLeftFunc(rest, unicode.IsDigit)
		if len(digits) == len(rest) || digits == "" || (digits[0] != '.' && digits[0] != ')') {
			return -1
		}
		rest = digits[1:]
	}
	box := strings.TrimLeft(rest, " ")
	if len(box) == len(rest) || len(box) < 3 || box[0] != '[' || box[2] != ']' || !strings.ContainsRune(" xX", rune(box[1])) {
		return -1
	}
	if len(box) > 3 && box[3] != ' ' && box[3] != '\t' {
		return -1
	}
	return len(line) - len(box) + 1
}
