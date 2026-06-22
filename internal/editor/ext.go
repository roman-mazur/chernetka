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
type BufferExtData interface {
}

// CodeAssist can be optionally implemented by BufferExtData.
type CodeAssist interface {
	TextSuggestion() string
}

// SyntaxHighlighter can be optionally implemented by BufferExtData.
type SyntaxHighlighter interface {
	SyntaxSpans(lineNumber int, line string) []SyntaxSpan
}

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

//go:generate go run golang.org/x/tools/cmd/stringer -type=TokenType

const (
	TtNothing TokenType = iota
	TtKeyword
	TtIdentifier
	TtTypeRef
	TtImportRef
	TtDeclaration
	TtCall
	TtStringLiteral
	TtNumberLiteral
	TtComment
)
