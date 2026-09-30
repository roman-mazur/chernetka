package extlsp

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"testing"
	"time"

	"go.lsp.dev/protocol"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/logger"
)

func TestLanguageForPath(t *testing.T) {
	for path, want := range map[string]string{
		"main.go":         "go",
		"a/b/MAIN.GO":     "go",
		"schema.cue":      "cue",
		"notes.md":        "",
		"go":              "",
		"dir.cue/file.md": "",
	} {
		got := ""
		if lang := languageForPath(path); lang != nil {
			got = lang.id()
		}
		if got != want {
			t.Errorf("languageForPath(%q) = %q, want %q", path, got, want)
		}
	}
}

// cueItems are completion items the way cue lsp returns them: every candidate
// in alphabetical order, fields inserting a colon, and definitions replacing
// the typed name together with its '#'.
func cueItems(line string, cx int) []protocol.CompletionItem {
	start := strings.LastIndexAny(line[:cx], " \t:") + 1
	edit := func(text string) *protocol.TextEdit {
		return &protocol.TextEdit{
			Range: protocol.Range{
				Start: protocol.Position{Character: uint32(start)},
				End:   protocol.Position{Character: uint32(cx)},
			},
			NewText: text,
		}
	}
	return []protocol.CompletionItem{
		{Label: "#Service", TextEdit: edit("#Service")},
		{Label: "list", TextEdit: edit("list: ")},
		{Label: "name", TextEdit: edit("name: ")},
		{Label: "names", TextEdit: edit("names: ")},
		{Label: "port", TextEdit: edit("port: ")},
		{Label: "strings", TextEdit: edit("strings")},
		{Label: "svc", TextEdit: edit("svc")},
	}
}

func TestRankByPrefix(t *testing.T) {
	for _, tc := range []struct {
		line string
		want []string // suggested texts
	}{
		{"svc: #Se", s("rvice")},
		{"\tpo", s("rt: ")},
		// The common prefix goes first, the candidates can be cycled.
		{"\tna", s("me", "me: ", "mes: ")},
		// A complete name isn't extended with a longer one.
		{"\tname", nil},
		{"\tzz", nil},
	} {
		cx := len(tc.line)
		items := rankByPrefix(cueItems(tc.line, cx), tc.line, cx)
		var got []string
		for _, sug := range extractSuggestions(items, tc.line, cx) {
			got = append(got, sug.text)
		}
		if !slices.Equal(got, tc.want) {
			t.Errorf("%q: suggested %q, want %q", tc.line, got, tc.want)
		}
	}
}

func TestIntegration_CUEBuffer(t *testing.T) {
	fake := &fakeLSP{}
	fake.items = cueItems("svc: #Se", len("svc: #Se"))
	var le Integration
	var (
		mu      sync.Mutex
		started []string
	)
	le.Starter = func(_ context.Context, languageID, _ string) (lspClient, error) {
		mu.Lock()
		defer mu.Unlock()
		started = append(started, languageID)
		return fake, nil
	}
	h, _, data := newBuffer(t, &le, fake, "config.cue", "")
	fake.waitVersion(t, 1)
	h.Run(t)
	h.Post(t, editor.CommandFunc(func(*editor.Editor) { h.SetMode(editor.ModeInsert) }))

	h.SendInputSequence(t, "svc: #Se")
	waitSuggestion(t, h, data, "rvice")
	// Would add a Go import in a Go buffer.
	h.SendInputSequence(t, " & strings.Min(")
	fake.waitVersion(t, int32(1+len("svc: #Se & strings.Min(")))

	fake.mu.Lock()
	defer fake.mu.Unlock()
	if fake.languageID != "cue" {
		t.Errorf("document opened as %q, want cue", fake.languageID)
	}
	if fake.organized != 0 {
		t.Errorf("%d organize imports requests for a CUE buffer, want none", fake.organized)
	}
	mu.Lock()
	defer mu.Unlock()
	if !slices.Equal(started, []string{"cue"}) {
		t.Errorf("started servers %q, want just cue", started)
	}
}

// TestIntegration_ServerPerLanguage checks that buffers of different languages
// get their own servers, and that a server failing to start doesn't disable
// the others.
func TestIntegration_ServerPerLanguage(t *testing.T) {
	goServer := &fakeLSP{}
	var le Integration
	le.Starter = func(_ context.Context, languageID, _ string) (lspClient, error) {
		if languageID == "go" {
			return goServer, nil
		}
		return nil, exec.ErrNotFound // Like cue not installed.
	}
	h := editor.NewTestHarness()
	h.LogEmbed = logger.Embed(logger.Prefix(t.Logf, "editor: "))
	h.Extend(&le)
	t.Cleanup(func() { _ = le.Close() })
	for _, path := range []string{"a.go", "b.cue", "c.go"} {
		if err := h.OpenReader(path, strings.NewReader("")); err != nil {
			t.Fatal(err)
		}
	}
	if n := len(le.servers); n != 2 {
		t.Fatalf("%d servers, want 2", n)
	}
	goServer.waitVersion(t, 1)
	cue := le.servers["cue"]
	deadline := time.Now().Add(2 * time.Second)
	for !cue.failed.Load() {
		if time.Now().After(deadline) {
			t.Fatal("the cue server didn't fail")
		}
		time.Sleep(5 * time.Millisecond)
	}
	if le.servers["go"].failed.Load() {
		t.Error("the go server failed with the cue one")
	}
}

// TestRealCUE drives the editor with the production integration and a real
// cue lsp: typing a definition name completes it, and saving formats the file.
func TestRealCUE(t *testing.T) {
	skipUnlessEval(t)
	if _, err := exec.LookPath("cue"); err != nil {
		t.Skip("cue is not installed")
	}
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, "cue.mod"), 0o755); err != nil {
		t.Fatal(err)
	}
	module := "module: \"example.com/x@v0\"\nlanguage: version: \"v0.12.0\"\n"
	if err := os.WriteFile(filepath.Join(dir, "cue.mod", "module.cue"), []byte(module), 0o644); err != nil {
		t.Fatal(err)
	}
	// Unformatted: saving fixes the indentation of the field.
	src := "package x\n\n#Service: {\n  name:   string\n}\n\nsvc: \n"
	path := filepath.Join(dir, "x.cue")
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}

	var le Integration
	h := editor.NewTestHarness()
	h.LogEmbed = logger.Embed(logger.Prefix(t.Logf, "editor: "))
	h.Extend(&le)
	if err := h.OpenReader(path, strings.NewReader(src)); err != nil {
		t.Fatal(err)
	}
	defer le.Close()
	buf := h.Top()
	data := buf.ExtensionData(le.ID()).(*BufferData)
	h.Run(t)
	deadline := time.Now().Add(10 * time.Second)
	for !data.srv.ready.Load() {
		if data.srv.failed.Load() || time.Now().After(deadline) {
			t.Fatal("cue lsp didn't start")
		}
		time.Sleep(10 * time.Millisecond)
	}

	h.Post(t, editor.CommandFunc(func(*editor.Editor) {
		for range 6 {
			editor.RelMove{Dy: 1}.DoOnBuffer(buf, editor.RenderPrefs{})
		}
		h.MoveCursorToLineEnd()
		h.SetMode(editor.ModeInsert)
	}))
	h.SendInputSequence(t, "#Se")
	waitSuggestion(t, h, data, "rvice")

	h.SendInput(t, []byte{'\t'}) // Accept.
	h.SendInput(t, []byte{0x1b}) // Leave the insert mode.
	h.Post(t, editor.CommandFunc(func(e *editor.Editor) {
		(&editor.Save{DstPath: path}).DoOnBuffer(e.Top(), editor.RenderPrefs{TabSize: 4})
	}))
	want := "package x\n\n#Service: {\n\tname: string\n}\n\nsvc: #Service\n"
	if got := onLoop(t, h, buf.Text); got != want {
		t.Errorf("text after save =\n%s\nwant\n%s", got, want)
	}
}
