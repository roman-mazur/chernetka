// Package extindent implements an editor extension that keeps the indentation when Enter splits
// a line in the insert mode.
//
// The new line gets the indentation of the split line, tabs or spaces, so the cursor lands at
// the column the text continues at. An empty line takes it from the first non-empty line above.
// After an opening bracket the new line is indented one level deeper, and a closing bracket right
// after the cursor goes to a line of its own. The blanks around the split are dropped.
package extindent

import (
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

// HandleInsertKey implements editor.InsertKeyHandler: it splits the line with Enter.
// The editor replaces a selection with Enter itself.
func (in *Integration) HandleInsertKey(buf *editor.Buffer, k input.Key) (handled bool) {
	if k.Special != input.Enter {
		return false
	}
	d, ok := buf.ExtensionData(in.ID()).(*document)
	if !ok {
		return false
	}
	for _, s := range buf.Selection() {
		if s.Start != s.End {
			return false
		}
	}
	d.splitLine()
	return true
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
		newIndent += indentUnit(lines, refLine, indent)
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

const blanks = " \t"

func leadingBlanks(s string) string { return s[:len(s)-len(strings.TrimLeft(s, blanks))] }

// indentUnit returns one level of indentation in the style of indent of the line refLine:
// a tab, or spaces if the line is indented with spaces. The number of spaces is how much
// deeper the line is than the closest less indented line above, or the whole indent if
// there is none.
func indentUnit(lines []content.Line, refLine int, indent string) string {
	if indent == "" || indent[0] != ' ' {
		return "\t"
	}
	unit := len(indent)
	for i := refLine - 1; i >= 0; i-- {
		l := lines[i].String()
		if strings.TrimLeft(l, blanks) == "" {
			continue // Blank lines don't show the structure.
		}
		ind := leadingBlanks(l)
		if strings.Trim(ind, " ") != "" {
			continue // Indented with tabs.
		}
		if len(ind) < len(indent) {
			unit = len(indent) - len(ind)
			break
		}
	}
	return strings.Repeat(" ", unit)
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
