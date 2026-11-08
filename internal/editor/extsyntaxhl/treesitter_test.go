package extsyntaxhl

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
	"unicode"

	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/logger"
)

func TestGoHighlight(t *testing.T) {
	doc := hlDoc{
		{"// Command comment.", []string{"0:19:Comment"}},
		{"package main", []string{"0:7:Keyword", "8:12:Identifier"}},
		{"", nil},
		{"import \"fmt\"", []string{"0:6:Keyword", "7:12:ImportRef"}},
		{"", nil},
		{"type T struct{ N int }", []string{
			"0:4:Keyword", "5:6:TypeRef", "7:13:Keyword", "15:16:Field", "17:20:TypeRef",
		}},
		{"", nil},
		{"func (t *T) Do(s string) error {", []string{
			"0:4:Keyword", "6:7:Identifier", "9:10:TypeRef",
			"12:14:FuncDeclaration", "15:16:Identifier", "17:23:TypeRef", "25:30:TypeRef",
		}},
		{"\tconst x = 1.5", []string{"1:6:Keyword", "7:8:Identifier", "11:14:NumberLiteral"}},
		// The \t escape sequence sits inside the string literal and overrides it.
		{"\tfmt.Println(\"a\\tb\", t.N, x, nil)", []string{
			"1:4:Identifier", "5:12:Call", "13:15:StringLiteral", "15:17:Escape",
			"17:19:StringLiteral", "21:22:Identifier", "23:24:Field", "26:27:Identifier",
			"29:32:Constant",
		}},
		{"\treturn nil", []string{"1:7:Keyword", "8:11:Constant"}},
		{"}", nil},
	}

	buf, ext := openDoc(t, "main.go", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

// TestGoMultiLineLiteral covers a node that starts on one line and ends on
// another: it has to be cut into one span per line, each clipped to its own
// line, instead of being reported with the start line's number and the end
// line's column.
func TestGoMultiLineLiteral(t *testing.T) {
	doc := hlDoc{
		{"package main", []string{"0:7:Keyword", "8:12:Identifier"}},
		{"", nil},
		{"var raw = `first", []string{"0:3:Keyword", "4:7:Identifier", "10:16:StringLiteral"}},
		{"second line is longer", []string{"0:21:StringLiteral"}},
		{"third`", []string{"0:6:StringLiteral"}},
		{"", nil},
		{"/* a block", []string{"0:10:Comment"}},
		{"   comment */", []string{"0:13:Comment"}},
	}

	buf, ext := openDoc(t, "main.go", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestGoHighlightAfterEdit(t *testing.T) {
	doc := hlDoc{
		{"package main", []string{"0:7:Keyword", "8:12:Identifier"}},
		{"", nil},
		{"func main() {}", []string{"0:4:Keyword", "5:9:FuncDeclaration"}},
	}

	buf, ext := openDoc(t, "main.go", doc.text())
	hl := highlighterOf(t, buf, ext)
	doc.check(t, hl)

	// Inserting a line shifts everything below it down by one.
	insertLine(t, buf, ext, 1)
	shifted := hlDoc{doc[0], {"", nil}, doc[1], doc[2]}
	shifted.check(t, hl)
}

func TestNixHighlight(t *testing.T) {
	doc := hlDoc{
		{"# A comment.", []string{"0:12:Comment"}},
		{"{ lib, enable ? true, ... }@args:", []string{
			"2:5:Identifier", "7:13:Identifier", "16:20:Constant", "28:32:Identifier",
		}},
		{"let", []string{"0:3:Keyword"}},
		{"  greet = name: \"hi ${name}\\n\";", []string{
			"2:7:FuncDeclaration", "10:14:Identifier", "16:20:StringLiteral",
			"20:22:Escape", "22:26:Identifier", "26:29:Escape", "29:30:StringLiteral",
		}},
		{"in rec {", []string{"0:2:Keyword", "3:6:Keyword"}},
		{"  a.b = lib.mkIf enable [ ./src 1.5 ];", []string{
			"2:3:Field", "4:5:Field", "8:11:Identifier", "12:16:Call", "17:23:Identifier",
			"26:31:StringLiteral", "32:35:NumberLiteral",
		}},
		{"  inherit (builtins) toString;", []string{"2:9:Keyword", "11:19:ImportRef", "21:29:Field"}},
		{"  c = args.x or null;", []string{
			"2:3:Field", "6:10:Identifier", "11:12:Field", "13:15:Keyword", "16:20:Constant",
		}},
		{"}", nil},
	}

	buf, ext := openDoc(t, "default.nix", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestCUEHighlight(t *testing.T) {
	doc := hlDoc{
		{"package config", []string{"0:7:Keyword", "8:14:Identifier"}},
		{"", nil},
		{"import \"strings\"", []string{"0:6:Keyword", "7:16:ImportRef"}},
		{"", nil},
		{"// #Service is a definition.", []string{"0:28:Comment"}},
		{"#Service: {", []string{"0:8:TypeRef"}},
		{"\tname!:  string & strings.MinRunes(1)", []string{
			"1:5:Field", "9:15:TypeRef", "18:25:Identifier", "26:34:Call", "35:36:NumberLiteral",
		}},
		{"\tport?:  int | *8080 @go(Port)", []string{
			"1:5:Field", "9:12:TypeRef", "16:20:NumberLiteral", "21:30:Constant",
		}},
		{"}", nil},
		{"", nil},
		{"svc: #Service & {", []string{"0:3:Field", "5:13:TypeRef"}},
		// The interpolated expression is code, so the string doesn't color it.
		{"\turl: \"http://\\(name):\\(port + 1)\\n\"", []string{
			"1:4:Field", "6:14:StringLiteral", "14:16:Escape", "16:20:Identifier", "20:21:Escape",
			"21:22:StringLiteral", "22:24:Escape", "24:31:Identifier", "31:32:NumberLiteral",
			"32:35:Escape", "35:36:StringLiteral",
		}},
		{"\tif port > 80 {tls: true}", []string{
			"1:3:Keyword", "4:8:Identifier", "11:13:NumberLiteral", "15:18:Field", "20:24:Constant",
		}},
		{"\tn: len([for x in [1, 2.5] {x}])", []string{
			"1:2:Field", "4:7:Call", "9:12:Keyword", "13:14:Identifier", "15:17:Keyword",
			"19:20:NumberLiteral", "22:25:NumberLiteral", "28:29:Identifier",
		}},
		{"}", nil},
	}

	buf, ext := openDoc(t, "config.cue", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

// TestCUEErrorRecovery checks the input the core v0.25.0 never finished parsing
// with the CUE grammar.
func TestCUEErrorRecovery(t *testing.T) {
	const text = "{R(z&["
	done := make(chan struct{})
	go func() {
		defer close(done)
		openDoc(t, "recovery.cue", text)
	}()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("parsing does not finish")
	}
}

func TestJSONHighlight(t *testing.T) {
	doc := hlDoc{
		{"{", nil},
		{"  // A comment.", []string{"2:15:Comment"}},
		{"  \"a\\\"b\": \"x\\ty\",", []string{
			"2:4:Field", "4:6:Escape", "6:8:Field", "10:12:StringLiteral", "12:14:Escape", "14:16:StringLiteral",
		}},
		{"  \"n\": [1.5e3, true, false, null]", []string{
			"2:5:Field", "8:13:NumberLiteral", "15:19:Constant", "21:26:Constant", "28:32:Constant",
		}},
		{"}", nil},
	}

	buf, ext := openDoc(t, "package.json", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestYAMLHighlight(t *testing.T) {
	doc := hlDoc{
		{"---", []string{"0:3:Keyword"}},
		{"# A comment.", []string{"0:12:Comment"}},
		{"name: CI", []string{"0:4:Field", "6:8:StringLiteral"}},
		{"base: &base", []string{"0:4:Field", "6:11:Constant"}},
		{"  n: 30", []string{"2:3:Field", "5:7:NumberLiteral"}},
		{"  f: 2.5", []string{"2:3:Field", "5:8:NumberLiteral"}},
		{"  ok: true", []string{"2:4:Field", "6:10:Constant"}},
		{"  \"quoted key\": 'v'", []string{"2:14:Field", "16:19:StringLiteral"}},
		{"job:", []string{"0:3:Field"}},
		{"  <<: *base", []string{"2:4:Field", "6:11:Constant"}},
		{"  env: !!map {A: \"1\\n\"}", []string{
			"2:5:Field", "7:12:TypeRef", "14:15:Field", "17:19:StringLiteral", "19:21:Escape", "21:22:StringLiteral",
		}},
		{"  steps: [checkout, test]", []string{"2:7:Field", "10:18:StringLiteral", "20:24:StringLiteral"}},
	}

	buf, ext := openDoc(t, "ci.yaml", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestShellHighlight(t *testing.T) {
	doc := hlDoc{
		{"#!/bin/sh", []string{"0:9:Comment"}},
		{"export DIR=\"${1:-.}\"", []string{
			"0:6:Keyword", "7:10:Field", "11:12:StringLiteral", "12:14:Escape", "14:15:Constant",
			"15:18:Identifier", "18:19:Escape", "19:20:StringLiteral",
		}},
		{"run() { ls -la \"$DIR\" >&2; }", []string{
			"0:3:FuncDeclaration", "8:10:Call", "11:14:Constant",
			"15:16:StringLiteral", "16:17:Escape", "17:20:Field", "20:21:StringLiteral", "24:25:NumberLiteral",
		}},
		// A substitution in a string is code again, not string content.
		{"if [ -n \"x $(date)\" ]; then", []string{
			"0:2:Keyword", "5:7:Keyword", "8:11:StringLiteral", "11:13:Escape", "13:17:Call",
			"17:18:Escape", "18:19:StringLiteral", "23:27:Keyword",
		}},
		{"  echo 'a $b' $# 42", []string{
			"2:6:Call", "7:13:StringLiteral", "14:15:Escape", "15:16:Constant", "17:19:NumberLiteral",
		}},
		{"fi", []string{"0:2:Keyword"}},
	}

	buf, ext := openDoc(t, "run.sh", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestCHighlight(t *testing.T) {
	doc := hlDoc{
		{"#include <stdio.h>", []string{"0:8:Keyword", "9:18:ImportRef"}},
		{"#define MAX_LEN 16", []string{"0:7:Keyword", "8:15:Constant"}},
		{"", nil},
		{"typedef struct { int x; unsigned long y; } point_t;", []string{
			"0:7:Keyword", "8:14:Keyword", "17:20:TypeRef", "21:22:Field", "24:37:TypeRef", "38:39:Field", "43:50:TypeRef",
		}},
		{"", nil},
		{"static const char *name(point_t *p) {", []string{
			"0:6:Keyword", "7:12:Keyword", "13:17:TypeRef", "19:23:FuncDeclaration", "24:31:TypeRef", "33:34:Identifier",
		}},
		// Upper case names read as macros.
		{"\tif (p->x > MAX_LEN) return NULL;", []string{
			"1:3:Keyword", "5:6:Identifier", "8:9:Field", "12:19:Constant", "21:27:Keyword", "28:32:Constant",
		}},
		{"\tprintf(\"%d\\n\", p->x, 'a');", []string{
			"1:7:Call", "8:11:StringLiteral", "11:13:Escape", "13:14:StringLiteral", "16:17:Identifier", "19:20:Field",
			"22:25:StringLiteral",
		}},
		{"\tp->fn(sizeof(int), 0x1fu); // done", []string{
			"1:2:Identifier", "4:6:Call", "7:13:Keyword", "14:17:TypeRef", "20:25:NumberLiteral", "28:35:Comment",
		}},
		{"}", nil},
	}

	buf, ext := openDoc(t, "main.c", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestCPPHighlight(t *testing.T) {
	doc := hlDoc{
		{"namespace app {", []string{"0:9:Keyword", "10:13:ImportRef"}},
		{"template <typename T>", []string{"0:8:Keyword", "10:18:Keyword", "19:20:TypeRef"}},
		{"class Box final : public Base {", []string{
			"0:5:Keyword", "6:9:TypeRef", "10:15:Keyword", "18:24:Keyword", "25:29:TypeRef",
		}},
		{"\texplicit Box(T v) noexcept : v_(v) {}", []string{
			"1:9:Keyword", "10:13:FuncDeclaration", "14:15:TypeRef", "16:17:Identifier", "19:27:Keyword", "30:32:Field",
			"33:34:Identifier",
		}},
		{"\t~Box() override = default;", []string{"2:5:FuncDeclaration", "8:16:Keyword", "19:26:Keyword"}},
		{"\tauto get() const -> const T& { return v_; }", []string{
			"1:5:TypeRef", "6:9:FuncDeclaration", "12:17:Keyword", "21:26:Keyword", "27:28:TypeRef", "32:38:Keyword",
			"39:41:Identifier",
		}},
		{"};", nil},
		{"}", nil},
		{"void app::run() {", []string{"0:4:TypeRef", "5:8:ImportRef", "10:13:FuncDeclaration"}},
		{"\tauto b = std::make_unique<Box<int>>(1);", []string{
			"1:5:TypeRef", "6:7:Identifier", "10:13:ImportRef", "15:26:Call", "27:30:TypeRef", "31:34:TypeRef",
			"37:38:NumberLiteral",
		}},
		{"\tstd::cout << R\"(raw)\" << nullptr << this->x;", []string{
			"1:4:ImportRef", "6:10:Identifier", "14:22:StringLiteral", "26:33:Constant", "37:41:Keyword", "43:44:Field",
		}},
		{"}", nil},
	}

	buf, ext := openDoc(t, "box.cpp", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

func TestCSharpHighlight(t *testing.T) {
	doc := hlDoc{
		// A dotted namespace is one span, not a chain of identifiers.
		{"using System.Collections.Generic;", []string{"0:5:Keyword", "6:32:ImportRef"}},
		{"", nil},
		{"[Serializable]", []string{"1:13:TypeRef"}},
		{"public class Repo<T> : IRepo<T> where T : class", []string{
			"0:6:Keyword", "7:12:Keyword", "13:17:TypeRef", "18:19:TypeRef", "23:28:TypeRef", "29:30:TypeRef",
			"32:37:Keyword", "42:47:Keyword",
		}},
		{"{", nil},
		{"    public int Count { get; private set; }", []string{
			"4:10:Keyword", "11:14:TypeRef", "15:20:Field", "23:26:Keyword", "28:35:Keyword", "36:39:Keyword",
		}},
		{"    public async Task<T?> FindAsync(int id)", []string{
			"4:10:Keyword", "11:16:Keyword", "17:21:TypeRef", "22:23:TypeRef", "26:35:FuncDeclaration", "36:39:TypeRef",
		}},
		{"    {", nil},
		{"        var x = items.Find(i => i is null);", []string{
			"8:11:Keyword", "22:26:Call", "34:36:Keyword", "37:41:Constant",
		}},
		// An interpolated expression is code again, not string content.
		{"        Console.WriteLine($\"id={id} \\t{x}\" + 'c');", []string{
			"16:25:Call", "26:31:StringLiteral", "31:32:Escape", "32:34:Identifier", "34:35:Escape", "35:36:StringLiteral",
			"36:39:Escape", "39:40:Identifier", "40:41:Escape", "41:42:StringLiteral", "45:48:StringLiteral",
		}},
		{"        return default;", []string{"8:14:Keyword", "15:22:Keyword"}},
		{"    }", nil},
		{"    enum Color { Red, Green }", []string{"4:8:Keyword", "9:14:TypeRef", "17:20:Constant", "22:27:Constant"}},
		{"}", nil},
	}

	buf, ext := openDoc(t, "Repo.cs", doc.text())
	doc.check(t, highlighterOf(t, buf, ext))
}

// TestGrammarsPrepare checks that every tree-sitter grammar loads and that its
// highlight query compiles. A broken query leaves the language uncolored
// without any other sign of trouble.
func TestGrammarsPrepare(t *testing.T) {
	for _, lang := range languages {
		hl, ok := lang.newHighlighter().(*tsHighlighter)
		if !ok {
			continue
		}
		if err := hl.grammar.prepare(); err != nil {
			t.Errorf("%s: %s", lang.name(), err)
		}
	}
}

// TestLanguageForPath checks that only the recognized file types get a
// highlighter, and that an unknown one is left alone rather than mishandled.
func TestLanguageForPath(t *testing.T) {
	for path, want := range map[string]string{
		"main.go":         "go",
		"a/b/main.GO":     "go",
		"flake.nix":       "nix",
		"schema.CUE":      "cue",
		"package.json":    "json",
		"settings.jsonc":  "jsonc",
		"a/flake.lock":    "json",
		"Cargo.lock":      "",
		"ci.yaml":         "yaml",
		"ci.yml":          "yaml",
		"run.sh":          "shell",
		"a/b/run.BASH":    "shell",
		"init.zsh":        "shell",
		"/home/u/.zshrc":  "shell",
		".bashrc":         "shell",
		".envrc":          "shell",
		"bashrc":          "",
		"main.c":          "c",
		"lib.h":           "cpp",
		"lib.cpp":         "cpp",
		"a/b/Lib.HPP":     "cpp",
		"Program.cs":      "csharp",
		"notes.md":        "markdown",
		"notes.markdown":  "markdown",
		"":                "",
		"Makefile":        "",
		"query.sql":       "",
		"go":              "",
		"archive.go.bak":  "",
		"dir.go/file.txt": "",
	} {
		got := ""
		if lang := languageForPath(path); lang != nil {
			got = lang.name()
		}
		if got != want {
			t.Errorf("languageForPath(%q) = %q, want %q", path, got, want)
		}
	}
}

func TestNoHighlightForUnknownType(t *testing.T) {
	buf, ext := openDoc(t, "notes.txt", "some text\n")
	if data := buf.ExtensionData(ext.ID()); data != nil {
		t.Errorf("got extension data %#v for an unknown file type, want none", data)
	}
	// AfterEdit has to tolerate a buffer it never made data for.
	ext.AfterEdit(nil, buf, nil)
	// So does a buffer that was never passed through MakeBufferData at all.
	ext.AfterEdit(nil, new(editor.Buffer), nil)
}

// TestGoRealSources runs the highlighter over this package's own sources. It is
// a sanity check rather than an exact comparison: every line with real content
// should come back with some highlight, and no span may fall outside its line.
func TestGoRealSources(t *testing.T) {
	entries, err := os.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".go" {
			continue
		}
		t.Run(entry.Name(), func(t *testing.T) {
			src, err := os.ReadFile(entry.Name())
			if err != nil {
				t.Fatal(err)
			}
			buf, ext := openDoc(t, entry.Name(), string(src))
			hl := highlighterOf(t, buf, ext)

			check := func(stage string) {
				t.Helper()
				for i, line := range buf.Content.Lines() {
					text := line.String()
					spans := hl.SyntaxSpans(i, text)
					for _, s := range spans {
						if s.LineNumber != i || s.Start < 0 || s.End > len(text) || s.Start >= s.End {
							t.Fatalf("%s > line %d %q: span %s is out of range", stage, i, text, s)
						}
					}
					// Any line with a word on it is a keyword, an identifier, a
					// comment or a string, so it has to come back colorized.
					// Lines of pure punctuation such as "}}," legitimately do not.
					if len(spans) == 0 && strings.ContainsFunc(text, unicode.IsLetter) {
						t.Errorf("%s > line %d %q: no highlight", stage, i, text)
					}
				}
			}

			check("initial")

			before := buf.Content.Len()
			insertLine(t, buf, ext, before/2)
			if buf.Content.Len() != before+1 {
				t.Fatalf("content length %d after insert, want %d", buf.Content.Len(), before+1)
			}
			check("after insert")
		})
	}
}

// TestNoHighlightForDirectoryListing covers a directory whose own name ends in
// a recognized extension: the buffer lists file names, not code.
func TestNoHighlightForDirectoryListing(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "notes.md")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "a.txt"), []byte("x"), 0o600); err != nil {
		t.Fatal(err)
	}

	ext := new(Integration)
	ext.LogEmbed = logger.Embed(t.Logf)
	var edit editor.Editor
	edit.Extend(ext)
	edit.OpenDir(dir, noopOpener{})

	buf := edit.Top()
	if buf.Path != "notes.md" {
		t.Fatalf("buffer path is %q, want the directory name", buf.Path)
	}
	if data := buf.ExtensionData(ext.ID()); data != nil {
		t.Errorf("got extension data %#v for a directory listing, want none", data)
	}
}

// noopOpener satisfies content.OpenFile for a test that never opens anything.
type noopOpener struct{}

func (noopOpener) OpenFile(string) {}
