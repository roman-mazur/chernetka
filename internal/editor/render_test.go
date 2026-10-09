package editor_test

import (
	"bufio"
	"io"
	"os"
	"testing"

	"rmazur.io/chernetka/internal/cheimg"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/editor/extcomment"
	"rmazur.io/chernetka/internal/editor/extd2"
	"rmazur.io/chernetka/internal/editor/extgo"
	"rmazur.io/chernetka/internal/editor/extmd"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl"
)

// BenchmarkEditor_Render measures a frame of a Go file without extensions, with
// the syntax highlighting, and with all the extensions of cmd/che but the LSP:
// it would start a language server, which reports to the editor loop.
func BenchmarkEditor_Render(b *testing.B) {
	for _, bc := range []struct {
		name string
		exts func() []editor.Extension
	}{
		{name: "plain", exts: func() []editor.Extension { return nil }},
		{name: "syntaxhl", exts: func() []editor.Extension {
			return []editor.Extension{new(extsyntaxhl.Integration)}
		}},
		{name: "all", exts: cheExtensions},
	} {
		b.Run(bc.name, func(b *testing.B) {
			f, err := os.Open("loop.go")
			if err != nil {
				b.Fatal(err)
			}
			b.Cleanup(func() { _ = f.Close() })

			h := editor.NewTestHarness()
			for _, ext := range bc.exts() {
				h.Extend(ext)
			}
			if err := h.OpenReader("loop.go", f); err != nil {
				b.Fatal(err)
			}

			out := bufio.NewWriter(io.Discard)
			for b.Loop() {
				h.RenderFrame(out)
			}
		})
	}
}

// cheExtensions returns new extensions of cmd/che but the LSP: it would start a language server,
// which reports to the editor loop. The diagrams are not shown, and the Go commands are not run.
func cheExtensions() []editor.Extension {
	return []editor.Extension{
		new(extsyntaxhl.Integration),
		&extd2.Integration{Viewer: nopViewer{}},
		new(extmd.Integration),
		new(extcomment.Integration),
		&extgo.Integration{Runner: nopRunner{}},
	}
}

type nopViewer struct{}

func (nopViewer) Show(cheimg.Item) {}

type nopRunner struct{}

func (nopRunner) Run(extgo.Command) {}
