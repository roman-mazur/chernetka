// Package content exposes an interface that represents the data visualized by a text editor.
// Implementations include maintaining all text lines in a memory or loading a part of a big file.
package content

// Interface defines the minimal set of methods needed to get the content visualized by the editor.
type Interface interface {
	// Lines returns the lines that should currently be displayed by the editor.
	Lines() []Line
	// Len returns the total number of lines.
	// It may be bigger than the size of a slice returned by Lines.
	Len() int
}

// Seekable extends the Interface with methods that allow the implementation to track where the user is and load
// more data on demand.
type Seekable interface {
	Interface

	UpdateUserCursor(lineNumber int) error
}

// Line represents a piece of content displayed by the editor. Usually a text line, but can also be an image.
type Line interface {
	MimeType() string
	String() string
	Len() int
}

type Mutable interface {
	Interface

	Insert(pos int, line Line)
	Update(pos int, line Line)
	Delete(pos int)
}
