package code

import "rmazur.io/chernetka/internal/content"

// Token is a range of the content recognized as a token of the TokenType, like
// the ranges the syntax scanners report to highlight.
type Token struct {
	content.Span
	TokenType
}
