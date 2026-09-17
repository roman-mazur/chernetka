package styles

import (
	"fmt"
	"image/color"
	"strconv"

	"rmazur.io/chernetka/internal/content/code"
)

var DefaultColors = ColorTheme{
	Suggestion:     color.Gray{Y: 100},
	LineSelected:   color.Gray{Y: 200},
	LineSelectedBg: color.Gray{Y: 70},
	TextSelected:   color.Gray{Y: 200},
	TextSelectedBg: parseColor("1010FF"),

	syntaxColors: map[code.TokenType]color.Color{
		code.TtKeyword:         parseColor("CF8E6D"),
		code.TtTypeRef:         parseColor("BCBEC4"),
		code.TtImportRef:       parseColor("57AAF7"),
		code.TtStringLiteral:   parseColor("6AAB73"),
		code.TtNumberLiteral:   parseColor("2AACB8"),
		code.TtFuncDeclaration: parseColor("56A8F5"),
		code.TtCall:            parseColor("56A8F5"),
		code.TtComment:         parseColor("7A7E85"),
		code.TtConstant:        parseColor("C77DBB"),
		code.TtField:           parseColor("C77DBB"),
		code.TtEscape:          parseColor("CF8E6D"),

		// Markup. Emphasis and strong text only differ by color: the terminal
		// writer sets foreground and background, not text attributes.
		code.TtPunctuation: parseColor("7A7E85"),
		code.TtHeading:     parseColor("56A8F5"),
		code.TtEmphasis:    parseColor("BCBEC4"),
		code.TtStrong:      parseColor("FFFFFF"),
		code.TtLink:        parseColor("C77DBB"),
		code.TtURL:         parseColor("548AF7"),
		code.TtListMarker:  parseColor("CF8E6D"),
		code.TtRawText:     parseColor("6AAB73"),
		code.TtQuote:       parseColor("9BA0A8"),
	},
}

func parseColor(s string) color.Color {
	if len(s) != 6 {
		panic(fmt.Errorf("expected 6 bytes, got %d", len(s)))
	}
	r, g, b := hex2i(s[:2]), hex2i(s[2:4]), hex2i(s[4:])
	return color.RGBA{R: r, G: g, B: b, A: 0xff}
}

func hex2i(s string) uint8 {
	res, err := strconv.ParseInt(s, 16, 32)
	if err != nil {
		panic(err)
	}
	return uint8(res)
}
