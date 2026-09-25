package d2view

import (
	"bytes"
	"context"
	"fmt"
	"regexp"
	"sync"

	resvg "github.com/kanrichan/resvg-go"
	"oss.terrastruct.com/d2/d2renderers/d2fonts"

	"rmazur.io/chernetka/internal/cheimg"
)

// Rasterizer converts diagrams to PNG images.
// d2 itself produces PNG images with a headless browser. Instead, the Rasterizer passes the SVG
// rendered by d2 to resvg, which is compiled to WASM and runs within the process.
type Rasterizer struct {
	Options

	once     sync.Once
	initErr  error
	rctx     *resvg.Context
	renderer *resvg.Renderer
}

// init prepares the resvg runtime. It takes a while, so it's done once on the first use.
func (r *Rasterizer) init() error {
	r.once.Do(func() {
		r.rctx, r.initErr = resvg.NewContext(context.Background())
		if r.initErr != nil {
			return
		}
		r.renderer, r.initErr = r.rctx.NewRenderer()
		if r.initErr != nil {
			return
		}
		for _, f := range embeddedFonts {
			if err := r.renderer.LoadFontData(d2fonts.FontFaces.Get(f)); err != nil {
				r.initErr = fmt.Errorf("load font %s %s: %w", f.Family, f.Style, err)
				return
			}
		}
		r.initErr = r.renderer.SetFontFamily(fontFamilySans)
	})
	return r.initErr
}

// PNG renders the diagram as a PNG image. It is not safe for concurrent use.
func (r *Rasterizer) PNG(ctx context.Context, d cheimg.Item) ([]byte, error) {
	svg, err := SVG(ctx, d, r.Options)
	if err != nil {
		return nil, err
	}
	if err := r.init(); err != nil {
		return nil, fmt.Errorf("rasterizer: %w", err)
	}
	return r.renderer.Render(replaceFonts(svg))
}

// Close releases the resvg runtime.
func (r *Rasterizer) Close() error {
	if r.renderer != nil {
		_ = r.renderer.Close()
	}
	if r.rctx != nil {
		return r.rctx.Close()
	}
	return nil
}

const (
	fontFamilySans = "Source Sans Pro"
	fontFamilyMono = "Source Code Pro"
)

var embeddedFonts = []d2fonts.Font{
	{Family: d2fonts.SourceSansPro, Style: d2fonts.FONT_STYLE_REGULAR},
	{Family: d2fonts.SourceSansPro, Style: d2fonts.FONT_STYLE_BOLD},
	{Family: d2fonts.SourceSansPro, Style: d2fonts.FONT_STYLE_SEMIBOLD},
	{Family: d2fonts.SourceSansPro, Style: d2fonts.FONT_STYLE_ITALIC},
	{Family: d2fonts.SourceCodePro, Style: d2fonts.FONT_STYLE_REGULAR},
	{Family: d2fonts.SourceCodePro, Style: d2fonts.FONT_STYLE_BOLD},
	{Family: d2fonts.SourceCodePro, Style: d2fonts.FONT_STYLE_SEMIBOLD},
	{Family: d2fonts.SourceCodePro, Style: d2fonts.FONT_STYLE_ITALIC},
}

// d2FontRE matches the font families d2 declares in the SVG styles, like "d2-123-font-mono-bold".
var d2FontRE = regexp.MustCompile(`"?d2-\d+-font-(mono-)?(regular|bold|semibold|italic)"?`)

// replaceFonts points the text styles to the fonts loaded to resvg.
// d2 embeds its fonts into the SVG as data URLs, which resvg does not support.
func replaceFonts(svg []byte) []byte {
	return d2FontRE.ReplaceAllFunc(svg, func(m []byte) []byte {
		sub := d2FontRE.FindSubmatch(m)
		family := fontFamilySans
		if len(sub[1]) > 0 {
			family = fontFamilyMono
		}
		var res bytes.Buffer
		fmt.Fprintf(&res, "%q", family)
		switch string(sub[2]) {
		case "bold":
			res.WriteString("; font-weight: bold")
		case "semibold":
			res.WriteString("; font-weight: 600")
		case "italic":
			res.WriteString("; font-style: italic")
		}
		return res.Bytes()
	})
}
