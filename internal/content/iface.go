// Package content exposes an interface that represents the data visualized by a text editor.
// Implementations include maintaining all text lines in a memory or loading a part of a big file.
package content

import "io"

// Document defines the minimal set of methods needed to get the content visualized by the editor.
type Document interface {
	// Lines returns the lines that should currently be displayed by the editor.
	Lines() []Line
	// Len returns the total number of lines.
	// It may be bigger than the size of a slice returned by Lines.
	Len() int
}

// Line represents a piece of content displayed by the editor. Usually a text line, but can also be an image.
type Line interface {
	MimeType() string
	String() string
	Len() int
}

// Mutable is a Document that can be modified.
type Mutable interface {
	Document

	// Insert puts the lines at pos, shifting the lines starting there.
	Insert(pos int, lines ...Line)
	Update(pos int, line Line)
	Delete(startPos, endPos int)
}

// UndoRedo extends a Mutable document with under/redo operations.
type UndoRedo interface {
	Mutable

	HasUndo() bool // Whether there is an action in the history that can be undone.
	HasRedo() bool
	Undo() // Revert the last mutation.
	Redo() // Cancel the last revert of a mutation.

	SplitChangeGroup() // Ensure that the next change is not grouped with the previous one.
}

// Print is useful to dump a part of the document for check what's inside.
func Print(out io.Writer, data Document, i, j int) {
	lines := data.Lines()
	_ = lines[i:j] // boundary check
	for x := i; x < j; x++ {
		out.Write([]byte(lines[x].String()))
		out.Write([]byte("\r\n"))
	}
}

// IsText returns true if the document holds text. It false for a file system document, for example.
func IsText(d Document) bool {
	switch d.(type) {
	case *FsContent, *ErrorContent:
		return false
	}
	lines := d.Lines()
	return len(lines) > 0 && lines[0].MimeType() == MimeTypeTextPlain
}
