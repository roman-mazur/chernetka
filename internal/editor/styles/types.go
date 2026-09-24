package styles

import (
	"image/color"

	"rmazur.io/chernetka/internal/content/code"
)

type TextStyle struct {
	Bold   bool
	Italic bool

	TextColor color.Color
	BgColor   color.Color
}

func ResolveTokenStyle(t code.TokenType) TextStyle {
	s := TextStyle{
		TextColor: DefaultColors.ColorForTokenType(t),
	}
	switch t {
	case code.TtKeyword, code.TtStrong, code.TtHeading:
		s.Bold = true
	case code.TtConstant, code.TtEmphasis:
		s.Italic = true
	default:
	}
	return s
}

type ColorTheme struct {
	Suggestion     color.Color
	LineSelected   color.Color
	LineSelectedBg color.Color
	TextSelected   color.Color
	TextSelectedBg color.Color
	LineAction     color.Color // marker of the lines that can be engaged

	syntaxColors map[code.TokenType]color.Color
}

func (ct *ColorTheme) ColorForTokenType(t code.TokenType) color.Color {
	if ct.syntaxColors == nil {
		return nil
	}
	return ct.syntaxColors[t]
}
