package d2view

import (
	"oss.terrastruct.com/d2/d2renderers/d2svg"
	"oss.terrastruct.com/d2/d2themes/d2themescatalog"
)

// Options configure image rendering.
type Options struct {
	ThemeID int64   // d2 theme, see `d2 themes`
	Scale   float64 // image size multiplier, 1 when not set
	Pad     int64   // space around the diagram in pixels before scaling
}

// DefaultOptions suit a terminal with a dark background.
var DefaultOptions = Options{
	ThemeID: d2themescatalog.DarkMauve.ID,
	Scale:   2,
	Pad:     20,
}

func (o Options) renderOpts() *d2svg.RenderOpts {
	res := &d2svg.RenderOpts{
		ThemeID: &o.ThemeID,
		Pad:     &o.Pad,
	}
	if o.Scale > 0 {
		res.Scale = &o.Scale
	}
	return res
}
