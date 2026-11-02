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
	"rmazur.io/chernetka/internal/editor/extindent"
	"rmazur.io/chernetka/internal/editor/extmd"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl"
	"rmazur.io/chernetka/internal/remotectl"
)

// BenchmarkEditor_Render measures a frame of a Go file without extensions, with
// the syntax highlighting, and with all the extensions of cmd/che but the LSP:
// it would start a language server, which reports to the editor loop.
func BenchmarkEditor_Render(b *testing.B) {
	for _, bc := range renderTestConfigs() {
		b.Run(bc.name, func(b *testing.B) {
			h, out := prepareRenderTest(b, &bc)
			for b.Loop() {
				h.RenderFrame(out)
			}
		})
	}
}

func TestEditor_Render_Allocations(t *testing.T) {
	for _, tc := range renderTestConfigs() {
		t.Run(tc.name, func(t *testing.T) {
			h, out := prepareRenderTest(t, &tc)
			n := testing.AllocsPerRun(100, func() {
				h.RenderFrame(out)
			})
			t.Log("avg allocations", n, "per 100 calls")
			if n > 1 {
				t.Errorf("allocations number is too big: %f", n)
			}
		})
	}
}

func prepareRenderTest(t testing.TB, cfg *renderTestConfig) (*editor.TestHarness, *bufio.Writer) {
	t.Helper()
	f, err := os.Open("loop.go")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = f.Close() })

	h := editor.NewTestHarness(t)
	for _, ext := range cfg.exts() {
		h.Extend(ext)
	}
	if err := h.OpenReader("loop.go", f); err != nil {
		t.Fatal(err)
	}
	out := bufio.NewWriter(io.Discard)
	return h, out
}

func renderTestConfigs() []renderTestConfig {
	return []renderTestConfig{
		{name: "plain", exts: func() []editor.Extension { return nil }},
		{name: "syntaxhl", exts: func() []editor.Extension {
			return []editor.Extension{new(extsyntaxhl.Integration)}
		}},
		{name: "all", exts: cheExtensions},
	}
}

type renderTestConfig struct {
	name string
	exts func() []editor.Extension
}

// cheExtensions returns new extensions of cmd/che but the LSP: it would start a language server,
// which reports to the editor loop. The diagrams are not shown, and the Go commands are not run.
func cheExtensions() []editor.Extension {
	return []editor.Extension{
		new(extsyntaxhl.Integration),
		&extd2.Integration{Viewer: nopViewer{}},
		new(extmd.Integration),
		new(extcomment.Integration),
		new(extindent.Integration),
		&extgo.Integration{Runner: nopRunner{}},
	}
}

type nopViewer struct{}

func (nopViewer) Show(cheimg.Item) {}

type nopRunner struct{}

func (nopRunner) ExecuteCommand(remotectl.CommandData) {}
