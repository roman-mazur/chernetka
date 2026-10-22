// Package extindent implements an editor extension that keeps the indentation when Enter splits
// a line in the insert mode, and changes the indentation of the selected lines with Tab and Shift+Tab.
//
// The new line gets the indentation of the split line, tabs or spaces, so the cursor lands at
// the column the text continues at. An empty line takes it from the first non-empty line above.
// After an opening bracket the new line is indented one level deeper, and a closing bracket right
// after the cursor goes to a line of its own. The blanks around the split are dropped.
//
// With a selection, Tab indents the selected lines one level deeper instead of replacing them.
// Shift+Tab takes one level of indentation off the selected lines, or off the cursor line
// without a selection. Blank lines are left as they are.
//
// A level of indentation is a tab, or, in the lines indented with spaces, as many spaces as the
// smallest indentation in the file. An unindented line follows the first indented line of the file.
package extindent

import (
	"iter"
	"strings"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/editor/input"
)

// Integration implements an editor.Extension indenting the new lines of the text files.
type Integration struct{}

func (in *Integration) ID() string { return "indent" }

func (in *Integration) MakeBufferData(_ editor.Sender, buf *editor.Buffer) editor.BufferExtData {
	if !buf.HoldsTextFile() {
		return nil // Not a file, like a directory listing.
	}
	if _, ok := buf.Content.(content.Mutable); !ok {
		return nil
	}
	return &document{buf: buf}
}

func (in *Integration) AfterEdit(editor.Sender, *editor.Buffer) {}

// HandleInsertKey implements editor.InsertKeyHandler: it splits the line with Enter, and changes
// the indentation of the selected lines with Tab and Shift+Tab. The editor replaces a selection
// with Enter itself.
func (in *Integration) HandleInsertKey(buf *editor.Buffer, k input.Key) (handled bool) {
	d, ok := buf.ExtensionData(in.ID()).(*document)
	if !ok {
		return false
	}
	switch {
	case k.Special == input.Enter && !buf.HasSelection():
		d.splitLine()
		return true
	case k == input.Of(input.Tab) && buf.HasSelection():
		d.indentLines(1)
		return true
	case k.Special == input.Backtab:
		d.indentLines(-1)
		return true
	}
	return false
}

var _ editor.InsertKeyHandler = new(Integration)

// document is the per-buffer extension data.
type document struct {
	buf *editor.Buffer
}

// splitLine breaks the line at the cursor, indenting the new line like the reference line.
func (d *document) splitLine() {
	cx, cy := d.buf.Pos()
	lines := d.buf.Content.Lines()
	if cy >= len(lines) {
		return
	}
	line := lines[cy].String()
	cx = min(cx, len(line))
	before, rest := line[:cx], line[cx:]
	start := len(strings.TrimRight(before, blanks)) // The blanks around the cursor are dropped.
	end := len(line) - len(strings.TrimLeft(rest, blanks))

	ref, refLine := before, cy
	if line == "" {
		ref, refLine = "", -1
		for i := cy - 1; i >= 0; i-- {
			if l := lines[i].String(); l != "" {
				ref, refLine = l, i
				break
			}
		}
	}
	indent := leadingBlanks(ref)
	newIndent := indent
	if trimmed := strings.TrimRight(ref, blanks); trimmed != "" && strings.ContainsRune("{([", rune(trimmed[len(trimmed)-1])) {
		newIndent += indentUnit(lines, indent)
		if opening := trimmed[len(trimmed)-1]; refLine == cy && end < len(line) && line[end] == closing(opening) {
			// The closing bracket goes below the new line: {|} becomes {\n\t|\n}.
			// Replacing the bracket keeps the cursor before it.
			d.buf.ReplaceText(content.Span{
				Start: content.Position{Line: cy, Col: end},
				End:   content.Position{Line: cy, Col: end + 1},
			}, "\n"+indent+line[end:end+1])
		}
	}
	// The cursor is within the span, so it lands at the end of the new text.
	d.buf.ReplaceText(content.Span{
		Start: content.Position{Line: cy, Col: start},
		End:   content.Position{Line: cy, Col: end},
	}, "\n"+newIndent)
}

// indentLines indents the lines touched by the selection spans, or the cursor line without
// a selection, one level deeper with dir 1, or one level shallower with dir -1. The edits are
// undone at once.
func (d *document) indentLines(dir int) {
	_, cy := d.buf.Pos()
	d.indentSpans(selectedSpans(d.buf.Selection(), cy), dir)
}

// indentSpans changes the indentation of the lines touched by spans, sorted and apart as
// selectedSpans returns them, by one level of the first non-blank of these lines. A line touched by two spans
// is changed once.
func (d *document) indentSpans(spans []content.Span, dir int) {
	lines := d.buf.Content.Lines()
	for _, span := range spans {
		if span.End.Line >= len(lines) {
			return
		}
	}
	unit := ""
	for i := range forLines(spans) {
		if l := lines[i].String(); strings.TrimLeft(l, blanks) != "" {
			unit = indentUnit(lines, leadingBlanks(l))
			break
		}
	}
	if unit == "" {
		return // Only blank lines.
	}

	d.buf.MutateAsGroup(func(content.Mutable) {
		// The edits are at the line starts, so they don't move the other lines.
		for i := range forLines(spans) {
			line := d.buf.Content.Lines()[i].String()
			if strings.TrimLeft(line, blanks) == "" {
				continue
			}
			start := content.Position{Line: i}
			if dir > 0 {
				d.buf.ReplaceText(content.Span{Start: start, End: start}, unit)
				continue
			}
			n := 0
			switch {
			case line[0] == '\t':
				n = 1
			case line[0] == ' ':
				n = min(len(unit), len(line)-len(strings.TrimLeft(line, " ")))
			}
			if n > 0 {
				d.buf.ReplaceText(content.Span{Start: start, End: content.Position{Line: i, Col: n}}, "")
			}
		}
	})
}

// selectedSpans returns the selection spans as editor.Buffer.Selection gives them, or a span of
// the cursor line if no text is selected. A span ending at the start of a line only selects the
// line break of the previous one, so it's made to end at the start of that line, or at its own
// start if it starts there, and the next span may start on the same line.
func selectedSpans(sel []content.Span, cursorLine int) []content.Span {
	if len(sel) == 0 {
		at := content.Position{Line: cursorLine}
		return []content.Span{{Start: at, End: at}}
	}
	for i, span := range sel {
		if span.End.Col == 0 && span.End.Line > span.Start.Line {
			end := content.Position{Line: span.End.Line - 1}
			if end.Line == span.Start.Line {
				end.Col = span.Start.Col // Not before the start.
			}
			sel[i].End = end
		}
	}
	return sel
}

// forLines yields each line touched by the sorted spans once, in order.
func forLines(spans []content.Span) iter.Seq[int] {
	return func(yield func(line int) bool) {
		next := 0 // The first line not visited yet.
		for _, span := range spans {
			for l := max(span.Start.Line, next); l <= span.End.Line; l++ {
				if !yield(l) {
					return
				}
			}
			next = max(next, span.End.Line+1)
		}
	}
}

const blanks = " \t"

func leadingBlanks(s string) string { return s[:len(s)-len(strings.TrimLeft(s, blanks))] }

// indentUnit returns one level of indentation in the style of indent: a tab, or spaces if
// it's made of spaces. A level of spaces is the smallest indentation of the lines indented
// with spaces only. Without indent, it's in the style of the first indented line of the file,
// or a tab.
func indentUnit(lines []content.Line, indent string) string {
	spaces, minSpaces := indent != "" && indent[0] == ' ', 0
	for _, l := range lines {
		s := l.String()
		ind := leadingBlanks(s)
		if ind == "" || ind == s {
			continue // Not indented or blank.
		}
		if indent == "" {
			indent, spaces = ind, ind[0] == ' ' // The style of the first indented line.
		}
		if n := len(ind); strings.Trim(ind, " ") == "" && (minSpaces == 0 || n < minSpaces) {
			minSpaces = n
		}
	}
	if !spaces || minSpaces == 0 {
		return "\t"
	}
	return strings.Repeat(" ", minSpaces)
}

func closing(opening byte) byte {
	switch opening {
	case '{':
		return '}'
	case '(':
		return ')'
	case '[':
		return ']'
	}
	return 0
}
