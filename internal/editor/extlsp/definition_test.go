package extlsp

import (
	"os"
	"path/filepath"
	"testing"

	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor"
)

func TestIntegration_FindDefinition(t *testing.T) {
	dir := t.TempDir()
	t.Chdir(dir)
	other := filepath.Join(dir, "other.go")
	if err := os.WriteFile(other, []byte("package main\n\n// 日本\nfunc /* 語 */ foo() {}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	at := func(path string, line, char uint32) []protocol.Location {
		p := protocol.Position{Line: line, Character: char}
		return []protocol.Location{{URI: uri.File(path), Range: protocol.Range{Start: p, End: p}}}
	}

	for _, tc := range []struct {
		name        string
		definitions []protocol.Location
		wantPath    string
		wantPos     content.Position
	}{
		{
			name:        "same file",
			definitions: at(filepath.Join(dir, "main.go"), 2, 6),
			wantPath:    filepath.Join(dir, "main.go"),
			wantPos:     content.Position{Line: 2, Col: 6},
		},
		{
			// The UTF-16 column is converted with the text of the other file.
			name:        "other file",
			definitions: at(other, 3, 13),
			wantPath:    other,
			wantPos:     content.Position{Line: 3, Col: len("func /* 語 */ ")},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			fake := &fakeLSP{definitions: tc.definitions}
			var le Integration
			// The buffer is not saved: the server must see its text.
			h, _, data := newBuffer(t, &le, fake, "main.go", "package main\n\nfunc bar() {}\n\nvar x = bar()\n")

			data.FindDefinition(h.Editor, content.Position{Line: 4, Col: len("var x = b")})
			cmd := <-h.Commands()
			goTo, ok := cmd.(*editor.GoTo)
			if !ok {
				t.Fatalf("posted %T, want *editor.GoTo", cmd)
			}
			if goTo.Path != tc.wantPath || goTo.Pos != tc.wantPos {
				t.Errorf("go to %s at %v, want %s at %v", goTo.Path, goTo.Pos, tc.wantPath, tc.wantPos)
			}

			fake.mu.Lock()
			defer fake.mu.Unlock()
			if want := "package main\n\nfunc bar() {}\n\nvar x = b"; len(fake.definedAt) != 1 || fake.definedAt[0] != want {
				t.Errorf("definition requested after %q, want after %q", fake.definedAt, want)
			}
		})
	}
}

func TestTextPosition(t *testing.T) {
	text := "a\n日本語x\n"
	for _, tc := range []struct {
		p    protocol.Position
		want content.Position
	}{
		{protocol.Position{Line: 0, Character: 1}, content.Position{Line: 0, Col: 1}},
		{protocol.Position{Line: 1, Character: 3}, content.Position{Line: 1, Col: len("日本語")}},
		{protocol.Position{Line: 5, Character: 3}, content.Position{Line: 5}},
	} {
		if got := textPosition(text, tc.p); got != tc.want {
			t.Errorf("textPosition(%v) = %v, want %v", tc.p, got, tc.want)
		}
	}
}
