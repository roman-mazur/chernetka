package editor

import "fmt"

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
	TokenType
}

func (ss SyntaxSpan) String() string {
	return fmt.Sprintf("%d:%d:%d:%s", ss.LineNumber, ss.Start, ss.End, ss.TokenType)
}

// TokenType represents a syntax token recognized by SyntaxHighlighter.
type TokenType int

//go:generate go run golang.org/x/tools/cmd/stringer -type=TokenType -trimprefix=Tt

const (
	TtNothing TokenType = iota
	TtKeyword
	TtIdentifier
	TtTypeRef
	TtImportRef
	TtFuncDeclaration
	TtCall
	TtStringLiteral
	TtNumberLiteral
	TtComment

	// Programming language tokens.

	TtConstant // language-defined constant, e.g. Go's true, false, nil, iota
	TtField    // struct field or property reference
	TtEscape   // escape sequence inside a string literal

	// Markup tokens.

	TtPunctuation // markup syntax characters, e.g. '#', '**', '`', '>'
	TtHeading     // heading text
	TtEmphasis    // italic text
	TtStrong      // bold text
	TtLink        // link or image label
	TtURL         // link destination or autolink
	TtListMarker  // list bullet, ordered list marker, or thematic break
	TtRawText     // inline code and fenced code block content
	TtQuote       // block quote text
)
