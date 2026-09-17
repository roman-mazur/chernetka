package code

// TokenType represents a syntax token in the content.
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
