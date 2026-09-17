package styles

import (
	"strconv"
	"testing"

	"rmazur.io/chernetka/internal/content/code"
)

// TestSyntaxColorCoverage guards against adding a token type and forgetting to
// give it a color, which would leave it rendered in the default foreground and
// so indistinguishable from unhighlighted text.
func TestSyntaxColorCoverage(t *testing.T) {
	// TtNothing marks the absence of a token. Identifiers are deliberately left
	// in the default color: in a typical source file most words are identifiers.
	exempt := map[code.TokenType]bool{code.TtNothing: true, code.TtIdentifier: true}

	for token := code.TtNothing; token <= code.TtQuote; token++ {
		if token.String() == "TokenType("+strconv.Itoa(int(token))+")" {
			t.Errorf("%d has no name; re-run go generate", int(token))
			continue
		}
		if got := DefaultColors.ColorForTokenType(token); (got == nil) != exempt[token] {
			t.Errorf("ColorForTokenType(%s) = %v, exempt = %v", token, got, exempt[token])
		}
	}
}
