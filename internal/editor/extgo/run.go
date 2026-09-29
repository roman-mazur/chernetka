// Package extgo implements an editor extension that makes Go tests and main functions actionable.
//
// The line declaring a test function like "func TestName(t *testing.T)" in a _test.go file
// can be engaged to run the test with go test, and the line declaring the main function
// of a main package can be engaged to run the program with go run. The commands are passed
// to a Runner in the directory of the file, after the editor saves the unsaved changes.
package extgo

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/logger"
)

// Command is a shell command to run in a directory.
type Command struct {
	Dir  string // absolute path
	Line string // shell command line
}

// Runner runs the commands, usually in a new terminal pane.
type Runner interface {
	// Run starts the command. It's called by the editor and must not block.
	Run(cmd Command)
}

// Integration implements an editor.Extension that provides line actions for the Go tests and main functions.
type Integration struct {
	logger.LogEmbed
	Runner Runner
}

func (in *Integration) ID() string { return "go" }

func (in *Integration) MakeBufferData(buf *editor.Buffer) editor.BufferExtData {
	if in.Runner == nil || buf.Path == "" || filepath.Ext(buf.Path) != ".go" {
		return nil
	}
	if _, isDir := buf.Content.(*content.FsContent); isDir {
		return nil // Its path is only a display name.
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
	testFunc = regexp.MustCompile(`^func (Test\w*)\(\w+ \*testing\.T\)`)
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
		if m == nil || !isTestName(m[1]) {
			return nil
		}
		return run{gf: gf, line: fmt.Sprintf("go test -run '^%s$' .", m[1])}
	}
	if mainFunc.MatchString(text) && gf.pkg() == "main" {
		return run{gf: gf, line: "go run ."}
	}
	return nil
}

// isTestName reports whether go test treats the function name as a test:
// "Test" is not followed by a lower-case letter.
func isTestName(name string) bool {
	r, _ := utf8.DecodeRuneInString(strings.TrimPrefix(name, "Test"))
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
	line string
}

func (r run) Engage() {
	dir, err := filepath.Abs(filepath.Dir(r.gf.buf.Path))
	if err != nil {
		r.gf.Logf("cannot run %s: %s", r.line, err)
		return
	}
	r.gf.Logf("run %s in %s", r.line, dir)
	r.gf.Runner.Run(Command{Dir: dir, Line: r.line})
}
