// Package extgo implements an editor extension that makes Go tests and main functions actionable.
//
// The line declaring a test function like "func TestName(t *testing.T)" in a _test.go file
// can be engaged to run the test with go test, a benchmark like "func BenchmarkName(b *testing.B)"
// is run with go test -bench, and the line declaring the main function
// of a main package can be engaged to run the program with go run. The commands are passed
// to a Runner in the directory of the file, after the editor saves the unsaved changes.
package extgo

import (
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/content/code"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/logger"
	"rmazur.io/chernetka/internal/remotectl"
)

// Integration implements an editor.Extension that provides line actions for the Go tests and main functions.
type Integration struct {
	logger.LogEmbed
	Runner remotectl.Executor
}

func (in *Integration) ID() string { return "go" }

func (in *Integration) MakeBufferData(_ editor.Sender, buf *editor.Buffer) editor.BufferExtData {
	if in.Runner == nil || code.SyntaxForPath(buf.Path) != code.Go {
		return nil
	}
	if !buf.HoldsTextFile() {
		return nil // Like a directory listing, whose path is only a display name.
	}
	return &goFile{Integration: in, buf: buf, test: strings.HasSuffix(buf.Path, "_test.go")}
}

func (in *Integration) AfterEdit(editor.Sender, *editor.Buffer) {}

// goFile is the per-buffer extension data.
type goFile struct {
	*Integration
	buf  *editor.Buffer
	test bool // a _test.go file
}

var (
	testFunc = regexp.MustCompile(`^func ((Test|Benchmark)(\w*))\(\w+ \*testing\.[TB]\)`)
	mainFunc = regexp.MustCompile(`^func main\(\)`)
	pkgName  = regexp.MustCompile(`^package (\w+)`)
)

func (gf *goFile) LineAction(lineNumber int) content.LineAction {
	lines := gf.buf.Content.Lines()
	if lineNumber < 0 || lineNumber >= len(lines) {
		return nil
	}
	text := lines[lineNumber].String()
	if !strings.HasPrefix(text, "func ") {
		return nil // Most lines are not.
	}
	if gf.test {
		m := testFunc.FindStringSubmatch(text)
		if m == nil || !isTestName(m[3]) {
			return nil
		}
		args := []string{"go", "test", "-v", "-run"}
		if m[2] == "Benchmark" {
			args = append(args, "^$", "-benchmem", "-bench")
		}
		return run{gf: gf, args: append(args, "^"+m[1]+"$", ".")}
	}
	if mainFunc.MatchString(text) && gf.pkg() == "main" {
		return run{gf: gf, args: []string{"go", "run", "."}}
	}
	return nil
}

// isTestName reports whether go test treats the function name as a test:
// "Test" or "Benchmark" is not followed by a lower-case letter.
func isTestName(name string) bool {
	r, _ := utf8.DecodeRuneInString(name)
	return !unicode.IsLower(r)
}

// pkg returns the name of the package from the first package clause of the file.
func (gf *goFile) pkg() string {
	for _, line := range gf.buf.Content.Lines() {
		if m := pkgName.FindStringSubmatch(line.String()); m != nil {
			return m[1]
		}
	}
	return ""
}

// run passes the command to the Runner in the directory of the file.
type run struct {
	gf   *goFile
	args []string
}

func (r run) Engage() {
	dir, err := filepath.Abs(filepath.Dir(r.gf.buf.Path))
	if err != nil {
		r.gf.Logf("cannot run %s: %s", r.args, err)
		return
	}
	r.gf.Logf("run %s in %s", r.args, dir)
	r.gf.Runner.ExecuteCommand(remotectl.CommandData{
		Action:  "exe",
		WorkDir: dir,
		Args:    r.args,
	})
}
