package extlsp

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// importSpec is an import of a Go file.
type importSpec struct {
	name string // empty unless the import is renamed
	path string
}

func (s importSpec) String() string {
	if s.name != "" {
		return s.name + " " + strconv.Quote(s.path)
	}
	return strconv.Quote(s.path)
}

// parseImports parses the package clause and imports of src. The rest of the
// file isn't looked at, so it may be broken, as it often is while typing.
func parseImports(src string) (*token.FileSet, *ast.File, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.ImportsOnly|parser.ParseComments)
	return fset, f, err
}

func importsOf(f *ast.File) []importSpec {
	res := make([]importSpec, 0, len(f.Imports))
	for _, is := range f.Imports {
		p, err := strconv.Unquote(is.Path.Value)
		if err != nil {
			continue
		}
		spec := importSpec{path: p}
		if is.Name != nil {
			spec.name = is.Name.Name
		}
		res = append(res, spec)
	}
	return res
}

// addedImports returns the imports of after that are not in before.
func addedImports(before, after string) []importSpec {
	_, fb, err := parseImports(before)
	if err != nil {
		return nil
	}
	_, fa, err := parseImports(after)
	if err != nil {
		return nil
	}
	had := make(map[importSpec]bool)
	for _, s := range importsOf(fb) {
		had[s] = true
	}
	var res []importSpec
	for _, s := range importsOf(fa) {
		if !had[s] {
			res = append(res, s)
		}
	}
	return res
}

// isImported reports whether a package may be referred to by name in src. It
// can't know the names of packages that differ from the last element of their
// import path: those are reported as not imported.
func isImported(src, name string) bool {
	_, f, err := parseImports(src)
	if err != nil {
		return false
	}
	for _, s := range importsOf(f) {
		if s.name == name || (s.name == "" && guessPackageName(s.path) == name) {
			return true
		}
	}
	return false
}

var majorVersion = regexp.MustCompile(`^v[0-9]+$`)

// guessPackageName guesses a package name from its import path:
// "math/rand" is "rand", "github.com/x/y/v2" is "y", "gopkg.in/yaml.v3" is "yaml".
func guessPackageName(importPath string) string {
	base := path.Base(importPath)
	if majorVersion.MatchString(base) {
		base = path.Base(path.Dir(importPath))
	}
	if i := strings.Index(base, ".v"); i > 0 && strings.HasPrefix(importPath, "gopkg.in/") {
		base = base[:i]
	}
	return base
}

// addImports inserts imports into src without touching the existing ones:
// into the import block (next to the imports of the same kind, standard
// library or not, in sorted order), turning a single import into a block, or
// after the package clause if there are no imports. It returns src unchanged
// if it can't be parsed.
func addImports(src string, specs []importSpec) string {
	fset, f, err := parseImports(src)
	if err != nil || len(specs) == 0 {
		return src
	}
	offset := func(p token.Pos) int { return fset.Position(p).Offset }

	var decls []*ast.GenDecl
	for _, d := range f.Decls {
		if gd, ok := d.(*ast.GenDecl); ok && gd.Tok == token.IMPORT {
			decls = append(decls, gd)
		}
	}

	sorted := append([]importSpec(nil), specs...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].path < sorted[j].path })

	switch {
	case len(decls) == 0:
		// After the package clause.
		at := lineEnd(src, offset(f.Name.End()))
		return src[:at] + "\n\n" + importDecl(sorted) + src[at:]

	case !decls[0].Lparen.IsValid() && len(decls) == 1:
		// import "fmt" → import ( "fmt" ... )
		all := append(importsOf(f), sorted...)
		sort.Slice(all, func(i, j int) bool { return all[i].path < all[j].path })
		d := decls[0]
		return src[:offset(d.Pos())] + importDecl(all) + src[offset(d.End()):]

	case !decls[0].Lparen.IsValid():
		// Several single imports: add another one after the last.
		at := lineEnd(src, offset(decls[len(decls)-1].End()))
		return src[:at] + "\n" + importDecl(sorted) + src[at:]
	}

	// Insert into the first block, from the end so offsets stay valid.
	block := decls[0]
	var inserts []insertion
	groups := importGroups(fset, block)
	for _, s := range sorted {
		inserts = append(inserts, insertIntoBlock(src, fset, block, groups, s))
	}
	sort.SliceStable(inserts, func(i, j int) bool { return inserts[i].at > inserts[j].at })
	for _, ins := range inserts {
		src = src[:ins.at] + ins.text + src[ins.at:]
	}
	return src
}

// insertion is text to insert at an offset.
type insertion struct {
	at   int
	text string
}

// importsEnd returns the offset where the imports of src end (or the package
// clause if there are none), or 0 if src can't be parsed.
func importsEnd(src string) int {
	fset, f, err := parseImports(src)
	if err != nil {
		return 0
	}
	end := f.Name.End()
	for _, d := range f.Decls {
		if gd, ok := d.(*ast.GenDecl); ok && gd.Tok == token.IMPORT {
			end = gd.End()
		}
	}
	return fset.Position(end).Offset
}

// importGroups splits the specs of an import block into groups separated by
// blank lines.
func importGroups(fset *token.FileSet, block *ast.GenDecl) [][]*ast.ImportSpec {
	var groups [][]*ast.ImportSpec
	lastLine := -1
	for _, s := range block.Specs {
		is := s.(*ast.ImportSpec)
		line := fset.Position(specStart(is)).Line
		if len(groups) == 0 || line > lastLine+1 {
			groups = append(groups, nil)
		}
		groups[len(groups)-1] = append(groups[len(groups)-1], is)
		lastLine = fset.Position(is.End()).Line
	}
	return groups
}

// insertIntoBlock finds where s goes in an import block.
func insertIntoBlock(src string, fset *token.FileSet, block *ast.GenDecl, groups [][]*ast.ImportSpec, s importSpec) insertion {
	offset := func(p token.Pos) int { return fset.Position(p).Offset }
	std := isStd(s.path)
	for _, g := range groups {
		if isStd(specPath(g[0])) != std {
			continue
		}
		indent := lineIndent(src, offset(g[0].Pos()))
		line := indent + s.String() + "\n"
		for _, is := range g {
			if specPath(is) > s.path {
				return insertion{lineStart(src, offset(specStart(is))), line}
			}
		}
		return insertion{lineEnd(src, offset(g[len(g)-1].End())) + 1, line}
	}
	// A new group at the end of the block.
	indent := "\t"
	if len(groups) > 0 {
		indent = lineIndent(src, offset(groups[0][0].Pos()))
	}
	text := indent + s.String() + "\n"
	if len(groups) > 0 {
		text = "\n" + text
	}
	rparen := offset(block.Rparen)
	at := lineStart(src, rparen)
	if strings.TrimSpace(src[at:rparen]) != "" {
		// ")" shares its line with the last import: start a line.
		at, text = rparen, "\n"+text
	}
	return insertion{at, text}
}

func importDecl(specs []importSpec) string {
	if len(specs) == 1 {
		return "import " + specs[0].String()
	}
	var sb strings.Builder
	sb.WriteString("import (\n")
	for _, s := range specs {
		sb.WriteString("\t" + s.String() + "\n")
	}
	sb.WriteString(")")
	return sb.String()
}

// specStart includes the doc comment of the import.
func specStart(is *ast.ImportSpec) token.Pos {
	if is.Doc != nil {
		return is.Doc.Pos()
	}
	return is.Pos()
}

func specPath(is *ast.ImportSpec) string {
	p, _ := strconv.Unquote(is.Path.Value)
	return p
}

// isStd reports whether an import path is from the standard library: its
// first element has no dot, unlike a domain.
func isStd(importPath string) bool {
	first, _, _ := strings.Cut(importPath, "/")
	return !strings.Contains(first, ".")
}

func lineStart(src string, offset int) int { return strings.LastIndexByte(src[:offset], '\n') + 1 }

func lineEnd(src string, offset int) int {
	if i := strings.IndexByte(src[offset:], '\n'); i >= 0 {
		return offset + i
	}
	return len(src)
}

func lineIndent(src string, offset int) string {
	start := lineStart(src, offset)
	end := start
	for end < len(src) && (src[end] == ' ' || src[end] == '\t') {
		end++
	}
	return src[start:end]
}

// qualifiedBeforeCursor reports the package name of a qualified identifier
// like "strconv.Itoa" that was just completed by typing a character after it,
// as in "strconv.Itoa(". Selectors of selectors ("a.b.C(") don't count.
func qualifiedBeforeCursor(before string) (pkg string, ok bool) {
	if before == "" || isIdentByte(before[len(before)-1]) || before[len(before)-1] == '.' {
		return "", false
	}
	s := before[:len(before)-1]
	sel := identTrailing(s)
	if sel == "" || startsWithDigit(sel) {
		return "", false
	}
	s = s[:len(s)-len(sel)]
	if !strings.HasSuffix(s, ".") {
		return "", false
	}
	s = s[:len(s)-1]
	pkg = identTrailing(s)
	if pkg == "" || startsWithDigit(pkg) {
		return "", false
	}
	if rest := s[:len(s)-len(pkg)]; strings.HasSuffix(rest, ".") {
		return "", false
	}
	return pkg, true
}

func isIdentByte(b byte) bool {
	return b == '_' || b >= 0x80 || ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9')
}
