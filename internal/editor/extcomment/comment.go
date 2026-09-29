// Package extcomment implements an editor extension that comments out lines with Ctrl+/
// in the insert mode.
//
// The line comment prefix is chosen by the file type. Ctrl+/ comments out the line with
// the cursor or the lines of the selection, or uncomments them if all of them are comments.
package extcomment

import (
	"path/filepath"
	"slices"
	"strings"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/editor/input"
)

// Integration implements an editor.Extension toggling the line comments in the files
// of the known types.
type Integration struct{}

func (in *Integration) ID() string { return "comment" }

func (in *Integration) MakeBufferData(buf *editor.Buffer) editor.BufferExtData {
	lines := buf.Content.Lines()
	if len(lines) > 0 && lines[0].MimeType() != content.MimeTypeTextPlain {
		return nil // Not a file, like a directory listing.
	}
	if _, ok := buf.Content.(content.Mutable); !ok {
		return nil
	}
	if p := prefixForPath(buf.Path); p != "" {
		return &document{buf: buf, prefix: p}
	}
	return nil
}

func (in *Integration) AfterEdit(editor.Sender, *editor.Buffer) {}

// HandleInsertKey implements editor.InsertKeyHandler: it toggles the comments with Ctrl+/.
func (in *Integration) HandleInsertKey(buf *editor.Buffer, k input.Key) (handled bool) {
	if k != input.Ctrl('/') {
		return false
	}
	d, ok := buf.ExtensionData(in.ID()).(*document)
	if !ok {
		return false
	}
	d.toggle()
	return true
}

var _ editor.InsertKeyHandler = new(Integration)

// document is the per-buffer extension data.
type document struct {
	buf    *editor.Buffer
	prefix string // of the line comments, like "//"
}

// toggle comments out the lines with the cursor or the selection, or uncomments them
// if all of them are comments. The prefix is inserted at the smallest indentation of
// the lines, followed by a space. Blank lines are left as is.
func (d *document) toggle() {
	lines := d.buf.Content.Lines()
	first, last := d.lineRange()
	if first < 0 || last >= len(lines) {
		return
	}

	indent, commented, blank := -1, true, true
	for _, l := range lines[first : last+1] {
		line := l.String()
		text := strings.TrimLeft(line, " \t")
		if text == "" {
			continue
		}
		blank = false
		if n := len(line) - len(text); indent < 0 || n < indent {
			indent = n
		}
		commented = commented && strings.HasPrefix(text, d.prefix)
	}
	if blank {
		return
	}

	for i := first; i <= last; i++ {
		line := d.buf.Content.Lines()[i].String()
		text := strings.TrimLeft(line, " \t")
		if text == "" {
			continue
		}
		if !commented {
			at := content.Position{Line: i, Col: indent}
			d.buf.ReplaceText(content.Span{Start: at, End: at}, d.prefix+" ")
			continue
		}
		col := len(line) - len(text)
		n := len(d.prefix)
		if strings.HasPrefix(text[n:], " ") {
			n++
		}
		d.buf.ReplaceText(content.Span{
			Start: content.Position{Line: i, Col: col},
			End:   content.Position{Line: i, Col: col + n},
		}, "")
	}
}

// lineRange returns the lines touched by the selection, or the cursor line without it.
// A selection ending at the start of a line only selects the line break of the previous one.
func (d *document) lineRange() (first, last int) {
	_, cy := d.buf.Pos()
	first, last = cy, cy
	selected := false
	for _, span := range d.buf.Selection() {
		start, end := span.Min(), span.Max()
		if start == end {
			continue
		}
		if end.Col == 0 && end.Line > start.Line {
			end.Line--
		}
		if !selected {
			first, last, selected = start.Line, end.Line, true
		}
		first, last = min(first, start.Line), max(last, end.Line)
	}
	return first, last
}

// prefixes maps the file extensions to the prefix of their line comments.
var prefixes = map[string]string{
	".go": "//", ".cue": "//", ".jsonc": "//",
	".nix": "#", ".sh": "#", ".bash": "#", ".zsh": "#", ".yaml": "#", ".yml": "#", ".d2": "#",
	".sql": "--",
}

// shellFileNames are the shell scripts recognized by their name alone.
var shellFileNames = []string{
	".bashrc", ".bash_profile", ".bash_login", ".bash_logout", ".profile",
	".zshrc", ".zshenv", ".zprofile", ".zlogin", ".zlogout",
	".envrc",
}

// prefixForPath returns the prefix of the line comments in the file at path,
// or an empty string if it's unknown.
func prefixForPath(path string) string {
	if p, ok := prefixes[strings.ToLower(filepath.Ext(path))]; ok {
		return p
	}
	if slices.Contains(shellFileNames, filepath.Base(path)) {
		return "#"
	}
	return ""
}
