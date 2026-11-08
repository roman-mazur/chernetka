package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/tables"
)

// The files of a tiny grammar, shaped like the generated ones.
var (
	testHeader = `#include <stdint.h>
typedef struct TSLanguage {
  const uint16_t *parse_table;
  const uint32_t *map;
  const char *const *names;
} TSLanguage;
`
	testParser = `#include "tree_sitter/parser.h"

#define LARGE_STATE_COUNT 2

static const uint16_t ts_parse_table[LARGE_STATE_COUNT][3] = {
  [0] = {1, 2, 3},
  [1] = {4, 5, 0x102},
};

static const char * const ts_symbol_names[] = {
  "end",
};

static const uint32_t ts_small_parse_table_map[] = {
  7, 0x10000,
};

const TSLanguage *tree_sitter_lang(void) {
  static const TSLanguage language = {
    .parse_table = &ts_parse_table[0][0],
    .map = ts_small_parse_table_map,
    .names = ts_symbol_names,
  };
  return &language;
}
`
)

func TestFetch(t *testing.T) {
	cc := requireTools(t)

	repo := t.TempDir()
	writeFiles(t, repo, map[string]string{
		"src/parser.c":             testParser,
		"src/scanner.c":            "void tree_sitter_lang_scanner(void) {}\n",
		"src/tree_sitter/parser.h": testHeader,
		"src/grammar.json":         "not copied",
		"LICENSE":                  "MIT",
	})
	gitRun(t, repo, "init", "--quiet", "--initial-branch=trunk")
	gitRun(t, repo, "add", ".")
	gitRun(t, repo, "-c", "user.name=test", "-c", "user.email=test@example.com", "commit", "--quiet", "-m", "grammar")
	commit, err := git(repo, "rev-parse", "HEAD")
	if err != nil {
		t.Fatal(err)
	}

	dir := t.TempDir()
	writeFiles(t, dir, map[string]string{
		"doc.go":               "package lang",
		"src/parser.c":         "old parser",
		"src/dropped_upstream": "stale",
	})
	got, err := fetch(repo, "trunk", "", dir, "lang")
	if err != nil {
		t.Fatal(err)
	}
	if got != commit {
		t.Errorf("fetch returned commit %q, want %q", got, commit)
	}

	for name, want := range map[string]string{
		"doc.go":                   "package lang",
		"src/scanner.c":            "void tree_sitter_lang_scanner(void) {}\n",
		"src/tree_sitter/parser.h": testHeader,
		"LICENSE":                  "MIT",
	} {
		if data := readFile(t, dir, name); data != want {
			t.Errorf("%s = %q, want %q", name, data, want)
		}
	}
	for _, name := range []string{"src/grammar.json", "src/dropped_upstream"} {
		if _, err := os.Stat(filepath.Join(dir, name)); err == nil {
			t.Errorf("%s exists, want it gone", name)
		}
	}
	upstream := readFile(t, dir, upstreamFile)
	if !strings.Contains(upstream, "package lang\n") || !strings.Contains(upstream, strconv.Quote(repo+"@"+commit)) {
		t.Errorf("upstream.go doesn't record the package and the commit:\n%s", upstream)
	}

	// The tables are left to be filled, in the order of tableNames, and the
	// rest of the parser is kept.
	parser := readFile(t, dir, "src/parser.c")
	for _, want := range []string{
		"static uint16_t ts_parse_table[LARGE_STATE_COUNT][3];\n",
		"static uint32_t ts_small_parse_table_map[2];\n",
		"static const char * const ts_symbol_names[] = {\n  \"end\",\n};\n",
		"case 0: *size = sizeof(ts_parse_table);",
		"case 1: *size = sizeof(ts_small_parse_table_map);",
	} {
		if !strings.Contains(parser, want) {
			t.Errorf("parser.c doesn't contain %q:\n%s", want, parser)
		}
	}
	if strings.Contains(parser, "0x102") || strings.Contains(parser, "0x10000") {
		t.Errorf("parser.c still holds the table contents:\n%s", parser)
	}

	compressed, err := os.ReadFile(filepath.Join(dir, tablesFile))
	if err != nil {
		t.Fatal(err)
	}
	parseTable, smallMap := make([]byte, 12), make([]byte, 8)
	if err := tables.Inflate(compressed, [][]byte{parseTable, smallMap}); err != nil {
		t.Fatal(err)
	}
	if want := []byte{1, 0, 2, 0, 3, 0, 4, 0, 5, 0, 2, 1}; !bytes.Equal(parseTable, want) {
		t.Errorf("ts_parse_table = %v, want %v", parseTable, want)
	}
	if want := []byte{7, 0, 0, 0, 0, 0, 1, 0}; !bytes.Equal(smallMap, want) {
		t.Errorf("ts_small_parse_table_map = %v, want %v", smallMap, want)
	}

	grammar := readFile(t, dir, grammarFile)
	for _, want := range []string{"package lang\n", `#include "src/scanner.c"`, "make([][]byte, 2)", "C.tree_sitter_lang()"} {
		if !strings.Contains(grammar, want) {
			t.Errorf("grammar.go doesn't contain %q:\n%s", want, grammar)
		}
	}

	// The stripped parser compiles, and its arrays are as large as the tables.
	writeFiles(t, dir, map[string]string{"main.c": `#include <stdio.h>
#include "parser.c"
int main(void) {
  for (int i = 0; i < 3; i++) {
    size_t size;
    fetchgrammar_table(i, &size);
    printf("%zu ", size);
  }
  return tree_sitter_lang() == NULL;
}
`})
	exe := filepath.Join(dir, "sizes.exe")
	if out, err := exec.Command(cc, "-std=c11", "-I", filepath.Join(dir, "src"), "-o", exe, filepath.Join(dir, "main.c")).CombinedOutput(); err != nil {
		t.Fatalf("compiling the stripped parser: %v\n%s", err, out)
	}
	out, err := exec.Command(exe).Output()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(out), "12 8 0 "; got != want {
		t.Errorf("the table sizes are %q, want %q", got, want)
	}
}

// requireTools skips the test without git or a C compiler, and returns the compiler.
func requireTools(t *testing.T) string {
	t.Helper()
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git is not installed")
	}
	cc := os.Getenv("CC")
	if cc == "" {
		cc = "cc"
	}
	if _, err := exec.LookPath(cc); err != nil {
		t.Skip("no C compiler")
	}
	return cc
}

func readFile(t *testing.T, dir, name string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(dir, filepath.FromSlash(name)))
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

// TestFetchSubdir covers a repository holding several grammars, like
// tree-sitter-typescript, whose scanners share a header out of their src.
func TestFetchSubdir(t *testing.T) {
	requireTools(t)

	repo := t.TempDir()
	writeFiles(t, repo, map[string]string{
		"lang/src/parser.c":             testParser,
		"lang/src/scanner.c":            "#include \"../../common/scanner.h\"\nvoid tree_sitter_lang_scanner(void) {}\n",
		"lang/src/tree_sitter/parser.h": testHeader,
		"common/scanner.h":              "#include \"tree_sitter/parser.h\"\n",
		"common/unused.h":               "not included",
		"src/parser.c":                  "another grammar",
		"LICENSE.md":                    "MIT",
	})
	gitRun(t, repo, "init", "--quiet", "--initial-branch=trunk")
	gitRun(t, repo, "add", ".")
	gitRun(t, repo, "-c", "user.name=test", "-c", "user.email=test@example.com", "commit", "--quiet", "-m", "grammar")

	dir := t.TempDir()
	writeFiles(t, dir, map[string]string{"common/stale.h": "stale"})
	if _, err := fetch(repo, "trunk", "lang", dir, "lang"); err != nil {
		t.Fatal(err)
	}

	// The shared header keeps its place relative to the scanner, and the
	// header found through the include path isn't looked for next to it.
	if got := readFile(t, dir, "common/scanner.h"); got != "#include \"tree_sitter/parser.h\"\n" {
		t.Errorf("common/scanner.h = %q", got)
	}
	if got := readFile(t, dir, "LICENSE.md"); got != "MIT" {
		t.Errorf("LICENSE.md = %q, want the license of the repository", got)
	}
	for _, name := range []string{"common/unused.h", "common/stale.h", "src/parser.c", "lang/src/grammar.json"} {
		if _, err := os.Stat(filepath.Join(dir, name)); err == nil {
			t.Errorf("%s exists, want it gone", name)
		}
	}
	grammar := readFile(t, dir, grammarFile)
	for _, want := range []string{"-I${SRCDIR}/lang/src", `#include "lang/src/parser.c"`, `#include "lang/src/scanner.c"`} {
		if !strings.Contains(grammar, want) {
			t.Errorf("grammar.go doesn't contain %q:\n%s", want, grammar)
		}
	}
	if !strings.Contains(readFile(t, dir, "lang/src/parser.c"), "static uint16_t ts_parse_table[LARGE_STATE_COUNT][3];") {
		t.Error("the parse table is left in lang/src/parser.c")
	}
}

func TestFetchMissingFile(t *testing.T) {
	requireTools(t)

	repo := t.TempDir()
	writeFiles(t, repo, map[string]string{"src/parser.c": testParser})
	gitRun(t, repo, "init", "--quiet", "--initial-branch=trunk")
	gitRun(t, repo, "add", ".")
	gitRun(t, repo, "-c", "user.name=test", "-c", "user.email=test@example.com", "commit", "--quiet", "-m", "grammar")

	dir := t.TempDir()
	writeFiles(t, dir, map[string]string{"src/parser.c": "current"})
	if _, err := fetch(repo, "trunk", "", dir, "lang"); err == nil {
		t.Fatal("fetch succeeded without the headers and a license")
	}
	if data, _ := os.ReadFile(filepath.Join(dir, "src/parser.c")); string(data) != "current" {
		t.Errorf("the current copy was changed to %q", data)
	}
}

func writeFiles(t *testing.T, dir string, files map[string]string) {
	t.Helper()
	for name, data := range files {
		path := filepath.Join(dir, filepath.FromSlash(name))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(data), 0o644); err != nil {
			t.Fatal(err)
		}
	}
}

func gitRun(t *testing.T, dir string, args ...string) {
	t.Helper()
	if _, err := git(dir, args...); err != nil {
		t.Fatal(err)
	}
}
