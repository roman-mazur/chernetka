// Package d2view renders d2 diagrams using the d2 module: as PNG images or as text art.
package d2view

import (
	"context"
	"os"
	"path/filepath"

	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2layouts/d2dagrelayout"
	"oss.terrastruct.com/d2/d2layouts/d2elklayout"
	"oss.terrastruct.com/d2/d2lib"
	"oss.terrastruct.com/d2/d2renderers/d2ascii"
	"oss.terrastruct.com/d2/d2renderers/d2ascii/charset"
	"oss.terrastruct.com/d2/d2renderers/d2svg"
	"oss.terrastruct.com/d2/d2target"
	"oss.terrastruct.com/d2/lib/textmeasure"

	"rmazur.io/chernetka/internal/cheimg"
)

// load returns the d2 source of the diagram, reading the file if needed.
func load(d cheimg.Item) (string, error) {
	if d.Source != "" || d.Path == "" {
		return d.Source, nil
	}
	data, err := os.ReadFile(d.Path)
	return string(data), err
}

func compile(ctx context.Context, d cheimg.Item, layout d2graph.LayoutGraph, renderOpts *d2svg.RenderOpts) (*d2target.Diagram, error) {
	src, err := load(d)
	if err != nil {
		return nil, err
	}
	ruler, err := textmeasure.NewRuler()
	if err != nil {
		return nil, err
	}
	opts := &d2lib.CompileOptions{
		Ruler:          ruler,
		LayoutResolver: func(string) (d2graph.LayoutGraph, error) { return layout, nil },
	}
	if d.Path != "" {
		// Resolve imports relative to the diagram file.
		opts.FS = os.DirFS(filepath.Dir(d.Path))
		opts.InputPath = filepath.Base(d.Path)
	}
	diagram, _, err := d2lib.Compile(ctx, src, opts, renderOpts)
	return diagram, err
}

// ASCII draws the diagram with box drawing characters.
func ASCII(ctx context.Context, d cheimg.Item) (string, error) {
	// ELK produces orthogonal routes that fit the character grid better than dagre does.
	// The d2 CLI makes the same choice for its text output.
	diagram, err := compile(ctx, d, d2elklayout.DefaultLayout, nil)
	if err != nil {
		return "", err
	}
	out, err := d2ascii.NewASCIIartist().Render(ctx, diagram, &d2ascii.RenderOpts{Charset: charset.Unicode})
	return string(out), err
}

// SVG renders the diagram as an SVG image.
func SVG(ctx context.Context, d cheimg.Item, opts Options) ([]byte, error) {
	renderOpts := opts.renderOpts()
	diagram, err := compile(ctx, d, d2dagrelayout.DefaultLayout, renderOpts)
	if err != nil {
		return nil, err
	}
	return d2svg.Render(diagram, renderOpts)
}
