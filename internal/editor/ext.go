package editor

// Extension represents an editor extension.
type Extension interface {
	ID() string
	MakeBufferData(buf *Buffer) BufferExtData

	AfterEdit(e *Editor, buf *Buffer)
	HandleInsertInput(buf *Buffer, prefs *RenderPrefs, b []byte) (handled, changed bool)
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
	SyntaxSpans(lineNumber int) []SyntaxSpan
}

type SyntaxSpan struct {
	LineNumber int
	Start, End int
	TokenType
}

type TokenType int

const (
	TokenTypeKeyword TokenType = iota
	TokenTypeIdentifier
	TokenTypeTypeRef
	TokenTypeImportRef
	TokenTypeDeclaration
	TokenTypeCall
	TokenTypeStringLiteral
	TokenTypeNumberLiteral
)
