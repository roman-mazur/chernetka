// Command fetchgrammar copies the C sources of a tree-sitter grammar from its
// git repository into the current directory and makes a Go package of them.
// It's meant to be run by go generate in the package of the grammar:
//
//	//go:generate go run ../fetchgrammar https://github.com/owner/tree-sitter-lang v1.2.3 [subdir]
//
// The reference, a tag, a branch or a commit, is fetched into a temporary
// directory, and the parser, the external scanner, the tree-sitter headers and
// the license replace the ones in the current directory. The commit they come
// from is recorded in upstream.go.
//
// The sources are under src, or under subdir/src for a repository holding
// several grammars, like tree-sitter-typescript. They keep their paths in the
// repository, and the files they include by a relative path, like a scanner
// header shared by the grammars, are copied along.
//
// The parse tables make up most of a grammar, and they compress well: they are
// mostly zeros. So they are taken out of parser.c, leaving arrays to be filled,
// and stored in tables.zlib, which grammar.go embeds and inflates into the
// arrays when the language is first asked for. The table bytes are read out by
// compiling and running a program with the original parser.c, so a C compiler
// is needed, the one in $CC or cc.
package main

import (
	"bytes"
	"compress/zlib"
	"errors"
	"fmt"
	"go/format"
	"io/fs"
	"log"
	"maps"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("fetchgrammar: ")
	if len(os.Args) != 3 && len(os.Args) != 4 {
		log.Fatal("usage: fetchgrammar <repository URL> <tag, branch or commit> [subdir]")
	}
	subdir := ""
	if len(os.Args) == 4 {
		subdir = os.Args[3]
	}
	pkg := os.Getenv("GOPACKAGE")
	if pkg == "" {
		log.Fatal("GOPACKAGE is not set: run it with go generate")
	}
	commit, err := fetch(os.Args[1], os.Args[2], subdir, ".", pkg)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("copied %s at %s", os.Args[1], commit)
}

// Outputs are the files fetch writes into the package directory, besides the
// sources under src.
const (
	tablesFile   = "tables.zlib"
	grammarFile  = "grammar.go"
	upstreamFile = "upstream.go"
)

// tableNames are the parse tables moved to tablesFile, in the order they are
// stored there. They hold numbers only, no pointers, so their bytes are the same
// on every little endian target. A table missing from a grammar is skipped.
var tableNames = []string{
	"ts_parse_table",
	"ts_small_parse_table",
	"ts_small_parse_table_map",
	"ts_parse_actions",
	"ts_lex_modes",
	"ts_primary_state_ids",
}

// fetch copies the grammar in subdir, which is empty for the repository root,
// at the reference of the repository into dir, which holds the Go package pkg,
// and returns the commit it's copied from.
func fetch(repo, ref, subdir, dir, pkg string) (string, error) {
	tmp, err := os.MkdirTemp("", "fetchgrammar-")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(tmp)

	// Unlike clone --branch, a fetch takes a commit as well as a tag or a branch.
	// The files are checked out as they are stored, with LF line endings, even
	// where git is set to convert them, like on Windows: parser.c is edited by
	// patterns matching its lines.
	clone := filepath.Join(tmp, "repo")
	for _, args := range [][]string{
		{"init", "--quiet", clone},
		{"-C", clone, "config", "core.autocrlf", "false"},
		{"-C", clone, "config", "core.eol", "lf"},
		{"-C", clone, "fetch", "--quiet", "--depth=1", repo, ref},
		{"-C", clone, "checkout", "--quiet", "FETCH_HEAD"},
	} {
		if _, err := git(tmp, args...); err != nil {
			return "", err
		}
	}
	commit, err := git(clone, "rev-parse", "HEAD")
	if err != nil {
		return "", err
	}

	// Collect everything before touching dir: a repository lacking a file must
	// leave the current copy intact.
	src := path.Join(subdir, "src") // in the repository and in dir
	files := make(map[string][]byte)
	for _, pattern := range []string{src + "/*.c", src + "/*.h", src + "/tree_sitter/*.h", "LICENSE*", "COPYING*"} {
		paths, err := filepath.Glob(filepath.Join(clone, filepath.FromSlash(pattern)))
		if err != nil {
			return "", err
		}
		for _, file := range paths {
			data, err := os.ReadFile(file)
			if err != nil {
				return "", err
			}
			rel, err := filepath.Rel(clone, file)
			if err != nil {
				return "", err
			}
			files[filepath.ToSlash(rel)] = data
		}
	}
	for _, name := range []string{src + "/parser.c", src + "/tree_sitter/parser.h"} {
		if files[name] == nil {
			return "", fmt.Errorf("no %s in %s", name, repo)
		}
	}
	// The license comes along with the code, whatever its file is named, like
	// LICENSE or LICENSE.md.
	if !slices.ContainsFunc(slices.Collect(maps.Keys(files)), isLicense) {
		return "", fmt.Errorf("no license in %s", repo)
	}
	if err := addIncluded(clone, files); err != nil {
		return "", err
	}
	_, hasScanner := files[src+"/scanner.c"]

	parser := string(files[src+"/parser.c"])
	langFunc := languageFunc.FindStringSubmatch(parser)
	if langFunc == nil {
		return "", fmt.Errorf("no language function in the parser.c of %s", repo)
	}
	tables, err := dumpTables(filepath.Join(clone, filepath.FromSlash(src)), parser, hasScanner)
	if err != nil {
		return "", err
	}
	stripped, err := stripTables(parser, tables)
	if err != nil {
		return "", err
	}
	compressed, err := compress(tables)
	if err != nil {
		return "", err
	}
	files[src+"/parser.c"] = []byte(stripped)
	files[tablesFile] = compressed
	if files[grammarFile], err = format.Source(grammarSource(pkg, src, langFunc[1], len(tables), hasScanner)); err != nil {
		return "", fmt.Errorf("formatting %s: %w", grammarFile, err)
	}
	files[upstreamFile] = upstreamSource(pkg, repo, ref, commit)

	// Files dropped upstream must not linger in the directories copied to.
	for name := range files {
		if top, _, nested := strings.Cut(name, "/"); nested {
			if err := os.RemoveAll(filepath.Join(dir, top)); err != nil {
				return "", err
			}
		}
	}
	for name, data := range files {
		dst := filepath.Join(dir, filepath.FromSlash(name))
		if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
			return "", err
		}
		if err := os.WriteFile(dst, data, fs.FileMode(0o644)); err != nil {
			return "", err
		}
	}
	return commit, nil
}

func isLicense(name string) bool {
	return !strings.Contains(name, "/") && (strings.HasPrefix(name, "LICENSE") || strings.HasPrefix(name, "COPYING"))
}

// includeLine finds the includes of the files in quotes, which are looked up
// relative to the including file first.
var includeLine = regexp.MustCompile(`(?m)^\s*#\s*include\s+"([^"]+)"`)

// addIncluded adds the files of the repository in clone that the C sources in
// files include by a relative path, and the files those include in turn. An
// include that isn't a file of the repository, like one found through the
// include path, is left to the compiler.
func addIncluded(clone string, files map[string][]byte) error {
	pending := make([]string, 0, len(files))
	for name := range files {
		pending = append(pending, name)
	}
	for len(pending) > 0 {
		name := pending[len(pending)-1]
		pending = pending[:len(pending)-1]
		if ext := path.Ext(name); ext != ".c" && ext != ".h" {
			continue
		}
		for _, m := range includeLine.FindAllStringSubmatch(string(files[name]), -1) {
			included := path.Join(path.Dir(name), m[1])
			if _, ok := files[included]; ok || strings.HasPrefix(included, "../") {
				continue
			}
			data, err := os.ReadFile(filepath.Join(clone, filepath.FromSlash(included)))
			if errors.Is(err, fs.ErrNotExist) {
				continue
			} else if err != nil {
				return err
			}
			files[included] = data
			pending = append(pending, included)
		}
	}
	return nil
}

// languageFunc finds the function returning the language in parser.c.
var languageFunc = regexp.MustCompile(`(?m)^(?:TS_PUBLIC |extern )?const TSLanguage \*(tree_sitter_\w+)\(void\)`)

// table is a parse table read out of a compiled parser.c.
type table struct {
	name  string
	data  []byte
	count int // of the elements in the first dimension
}

// dumpTables compiles a program printing the parse tables of the parser.c in
// src, given as parser, and returns them in the order of tableNames.
func dumpTables(src, parser string, hasScanner bool) ([]table, error) {
	var names []string
	for _, name := range tableNames {
		if findTable(parser, name) != nil {
			names = append(names, name)
		}
	}
	if len(names) == 0 {
		return nil, fmt.Errorf("no parse tables in parser.c")
	}

	var prog strings.Builder
	prog.WriteString("#include <stdio.h>\n#include \"parser.c\"\n\n")
	prog.WriteString("int main(int argc, char **argv) {\n\tFILE *f = fopen(argv[1], \"wb\");\n\tif (f == NULL) return 1;\n")
	for _, n := range names {
		fmt.Fprintf(&prog, "\tif (fwrite(%[1]s, 1, sizeof(%[1]s), f) != sizeof(%[1]s)) return 1;\n", n)
		fmt.Fprintf(&prog, "\tprintf(\"%%zu %%zu\\n\", sizeof(%[1]s), sizeof(%[1]s) / sizeof(%[1]s[0]));\n", n)
	}
	prog.WriteString("\treturn fclose(f) != 0;\n}\n")

	tmp, err := os.MkdirTemp("", "fetchgrammar-dump-")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tmp)
	progPath := filepath.Join(tmp, "dump.c")
	if err := os.WriteFile(progPath, []byte(prog.String()), 0o644); err != nil {
		return nil, err
	}

	cc := os.Getenv("CC")
	if cc == "" {
		cc = "cc"
	}
	exe := filepath.Join(tmp, "dump.exe") // gcc adds .exe on Windows otherwise
	// The scanner is linked in as the language refers to its functions.
	args := []string{"-std=c11", "-w", "-I", src, "-o", exe, progPath}
	if hasScanner {
		args = append(args, filepath.Join(src, "scanner.c"))
	}
	if out, err := exec.Command(cc, args...).CombinedOutput(); err != nil {
		return nil, fmt.Errorf("compiling the table dump: %w: %s", err, bytes.TrimSpace(out))
	}
	blobPath := filepath.Join(tmp, "tables.bin")
	sizes, err := exec.Command(exe, blobPath).Output()
	if err != nil {
		return nil, fmt.Errorf("dumping the tables: %w", err)
	}
	blob, err := os.ReadFile(blobPath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(sizes)), "\n")
	if len(lines) != len(names) {
		return nil, fmt.Errorf("the table dump printed %d sizes for %d tables", len(lines), len(names))
	}
	tables := make([]table, len(names))
	for i, line := range lines {
		var size, count int
		if _, err := fmt.Sscan(line, &size, &count); err != nil {
			return nil, fmt.Errorf("the table dump printed %q: %w", line, err)
		}
		if size > len(blob) {
			return nil, fmt.Errorf("the table dump wrote %d bytes, short of %s", len(blob), names[i])
		}
		tables[i] = table{name: names[i], data: blob[:size], count: count}
		blob = blob[size:]
	}
	if len(blob) != 0 {
		return nil, fmt.Errorf("the table dump wrote %d bytes more than its tables", len(blob))
	}
	return tables, nil
}

// tableDecl matches the definition of a table in parser.c, which is generated
// in a fixed shape, like:
//
//	static const uint16_t ts_parse_table[LARGE_STATE_COUNT][SYMBOL_COUNT] = {
//	  ...
//	};
var tableDecl = regexp.MustCompile(`(?m)^static const (\w+) (\w+)\[([^\]\n]*)\]((?:\[[^\]\n]*\])*) = \{\n`)

// tableLoc is where a table is defined in parser.c.
type tableLoc struct {
	start, end int // of the whole definition, up to the line after it
	typ        string
	dim, rest  string // the first dimension, and the following ones
}

func findTable(parser, name string) *tableLoc {
	for _, m := range tableDecl.FindAllStringSubmatchIndex(parser, -1) {
		if parser[m[4]:m[5]] != name {
			continue
		}
		closing := strings.Index(parser[m[1]:], "\n};\n")
		if closing < 0 {
			return nil
		}
		return &tableLoc{
			start: m[0],
			end:   m[1] + closing + len("\n};\n"),
			typ:   parser[m[2]:m[3]],
			dim:   parser[m[6]:m[7]],
			rest:  parser[m[8]:m[9]],
		}
	}
	return nil
}

// stripTables replaces the definitions of the tables in parser with arrays to
// be filled from tablesFile, and appends the function that grammar.go finds
// them by.
func stripTables(parser string, tables []table) (string, error) {
	for _, t := range tables {
		loc := findTable(parser, t.name)
		if loc == nil {
			return "", fmt.Errorf("no definition of %s in parser.c", t.name)
		}
		dim := loc.dim
		if dim == "" {
			dim = strconv.Itoa(t.count)
		}
		decl := fmt.Sprintf("// Filled from %s by grammar.go, see fetchgrammar.\nstatic %s %s[%s]%s;\n", tablesFile, loc.typ, t.name, dim, loc.rest)
		parser = parser[:loc.start] + decl + parser[loc.end:]
	}

	var b strings.Builder
	b.WriteString(parser)
	fmt.Fprintf(&b, "\n// Added by fetchgrammar: grammar.go fills the tables above from %s.\n\n", tablesFile)
	fmt.Fprintf(&b, "_Static_assert(__BYTE_ORDER__ == __ORDER_LITTLE_ENDIAN__, \"the tables in %s are little endian\");\n\n", tablesFile)
	b.WriteString("static void *fetchgrammar_table(int i, size_t *size) {\n  switch (i) {\n")
	for i, t := range tables {
		fmt.Fprintf(&b, "    case %d: *size = sizeof(%s); return (void *)%[2]s;\n", i, t.name)
	}
	b.WriteString("  }\n  *size = 0;\n  return NULL;\n}\n")
	return b.String(), nil
}

func compress(tables []table) ([]byte, error) {
	var b bytes.Buffer
	w, err := zlib.NewWriterLevel(&b, zlib.BestCompression)
	if err != nil {
		return nil, err
	}
	for _, t := range tables {
		if _, err := w.Write(t.data); err != nil {
			return nil, err
		}
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func grammarSource(pkg, src, langFunc string, tableCount int, hasScanner bool) []byte {
	var b bytes.Buffer
	fmt.Fprintf(&b, "// Code generated by fetchgrammar; DO NOT EDIT.\n\n")
	fmt.Fprintf(&b, "package %s\n\n", pkg)
	fmt.Fprintf(&b, "// #cgo CFLAGS: -std=c11 -fPIC -I${SRCDIR}/%s\n", src)
	fmt.Fprintf(&b, "// #include \"%s/parser.c\"\n", src)
	if hasScanner {
		fmt.Fprintf(&b, "// #include \"%s/scanner.c\"\n", src)
	}
	b.WriteString("import \"C\"\n\n")
	b.WriteString("import (\n\t_ \"embed\"\n\t\"sync\"\n\t\"unsafe\"\n\n\t\"rmazur.io/chernetka/internal/editor/extsyntaxhl/grammars/tables\"\n)\n\n")
	fmt.Fprintf(&b, "//go:embed %s\nvar compressedTables []byte\n\n", tablesFile)
	fmt.Fprintf(&b, "// Language returns the tree-sitter language, to be passed to treesitter.NewLanguage.\n")
	fmt.Fprintf(&b, "// Its parse tables are inflated on the first call.\n")
	fmt.Fprintf(&b, "var Language = sync.OnceValues(func() (unsafe.Pointer, error) {\n")
	fmt.Fprintf(&b, "\tdst := make([][]byte, %d)\n", tableCount)
	fmt.Fprintf(&b, "\tfor i := range dst {\n")
	fmt.Fprintf(&b, "\t\tvar size C.size_t\n")
	fmt.Fprintf(&b, "\t\tdst[i] = unsafe.Slice((*byte)(C.fetchgrammar_table(C.int(i), &size)), size)\n")
	fmt.Fprintf(&b, "\t}\n")
	fmt.Fprintf(&b, "\tif err := tables.Inflate(compressedTables, dst); err != nil {\n")
	fmt.Fprintf(&b, "\t\treturn nil, err\n")
	fmt.Fprintf(&b, "\t}\n")
	fmt.Fprintf(&b, "\treturn unsafe.Pointer(C.%s()), nil\n", langFunc)
	fmt.Fprintf(&b, "})\n")
	return b.Bytes()
}

func upstreamSource(pkg, repo, ref, commit string) []byte {
	var b bytes.Buffer
	fmt.Fprintf(&b, "// Code generated by fetchgrammar; DO NOT EDIT.\n\n")
	fmt.Fprintf(&b, "package %s\n\n", pkg)
	fmt.Fprintf(&b, "// Upstream is the repository and the commit the sources under src are copied\n")
	fmt.Fprintf(&b, "// from, fetched as %s.\n", ref)
	fmt.Fprintf(&b, "const Upstream = %q\n", repo+"@"+commit)
	return b.Bytes()
}

func git(dir string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, bytes.TrimSpace(stderr.Bytes()))
	}
	return strings.TrimSpace(string(out)), nil
}
