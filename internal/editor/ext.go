package editor

import (
	"fmt"

	"rmazur.io/chernetka/internal/content/code"
)

// Extension represents an editor extension.
type Extension interface {
	ID() string
	MakeBufferData(buf *Buffer) BufferExtData

	AfterEdit(e *Editor, buf *Buffer)
	HandleInsertInput(buf *Buffer, prefs *RenderPrefs, b []byte) (handled bool)
}

// BufferExtData represents data associated with a Buffer and managed by an Extension.
type BufferExtData any

// CodeAssist can be optionally implemented by BufferExtData.
type CodeAssist interface {
	TextSuggestion() string
}

// SyntaxHighlighter can be optionally implemented by BufferExtData.
type SyntaxHighlighter interface {
	// SyntaxSpans returns the highlighted regions of the given line.
	// The result is sorted by Start, spans never overlap, and every span stays
	// within the bounds of line. Regions that need no highlight are omitted, so
	// an empty result means the whole line is rendered with the default color.
	SyntaxSpans(lineNumber int, line string) []SyntaxSpan
}

// SyntaxSpan is a half-open [Start, End) byte range of a single line that
// should be rendered with the color of its TokenType.
type SyntaxSpan struct {
	LineNumber int
	Start, End int
	code.TokenType
}

func (ss SyntaxSpan) String() string {
	return fmt.Sprintf("%d:%d:%d:%s", ss.LineNumber, ss.Start, ss.End, ss.TokenType)
}
