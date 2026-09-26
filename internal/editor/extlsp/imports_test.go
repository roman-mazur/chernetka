package extlsp

import "testing"

func TestAddImports(t *testing.T) {
	strconvSpec := []importSpec{{path: "strconv"}}
	const body = "\nfunc main() {\n\tstrconv.Itoa(\n"
	cases := []struct {
		name  string
		src   string
		specs []importSpec
		want  string
	}{
		{
			name:  "no imports",
			src:   "package main\n" + body,
			specs: strconvSpec,
			want:  "package main\n\nimport \"strconv\"\n" + body,
		},
		{
			name:  "single import becomes a block",
			src:   "package main\n\nimport \"fmt\"\n" + body,
			specs: strconvSpec,
			want:  "package main\n\nimport (\n\t\"fmt\"\n\t\"strconv\"\n)\n" + body,
		},
		{
			name:  "sorted into the block",
			src:   "package main\n\nimport (\n\t\"fmt\"\n\t\"strings\"\n)\n" + body,
			specs: strconvSpec,
			want:  "package main\n\nimport (\n\t\"fmt\"\n\t\"strconv\"\n\t\"strings\"\n)\n" + body,
		},
		{
			name:  "appended to the end of the group",
			src:   "package main\n\nimport (\n\t\"errors\"\n\t\"fmt\"\n)\n" + body,
			specs: strconvSpec,
			want:  "package main\n\nimport (\n\t\"errors\"\n\t\"fmt\"\n\t\"strconv\"\n)\n" + body,
		},
		{
			name:  "into the standard library group",
			src:   "package main\n\nimport (\n\t\"fmt\"\n\n\t\"example.com/x\"\n)\n" + body,
			specs: strconvSpec,
			want:  "package main\n\nimport (\n\t\"fmt\"\n\t\"strconv\"\n\n\t\"example.com/x\"\n)\n" + body,
		},
		{
			name:  "into the third party group",
			src:   "package main\n\nimport (\n\t\"fmt\"\n\n\t\"example.com/x\"\n)\n" + body,
			specs: []importSpec{{path: "example.com/a"}},
			want:  "package main\n\nimport (\n\t\"fmt\"\n\n\t\"example.com/a\"\n\t\"example.com/x\"\n)\n" + body,
		},
		{
			name:  "a new group",
			src:   "package main\n\nimport (\n\t\"fmt\"\n)\n" + body,
			specs: []importSpec{{path: "example.com/a"}},
			want:  "package main\n\nimport (\n\t\"fmt\"\n\n\t\"example.com/a\"\n)\n" + body,
		},
		{
			name:  "several, renamed",
			src:   "package main\n\nimport (\n\t\"fmt\"\n)\n" + body,
			specs: []importSpec{{path: "os"}, {name: "crand", path: "crypto/rand"}},
			want:  "package main\n\nimport (\n\tcrand \"crypto/rand\"\n\t\"fmt\"\n\t\"os\"\n)\n" + body,
		},
		{
			name:  "keeps a doc comment with its import",
			src:   "package main\n\nimport (\n\t\"fmt\"\n\t// for Itoa\n\t\"strings\"\n)\n" + body,
			specs: strconvSpec,
			want:  "package main\n\nimport (\n\t\"fmt\"\n\t\"strconv\"\n\t// for Itoa\n\t\"strings\"\n)\n" + body,
		},
		{
			name:  "empty block",
			src:   "package main\n\nimport ()\n" + body,
			specs: strconvSpec,
			want:  "package main\n\nimport (\n\t\"strconv\"\n)\n" + body,
		},
		{
			name:  "unparsable header is left alone",
			src:   "package main\n\nimport (\n\t\"fmt\n" + body,
			specs: strconvSpec,
			want:  "package main\n\nimport (\n\t\"fmt\n" + body,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := addImports(tc.src, tc.specs); got != tc.want {
				t.Errorf("got:\n%s\nwant:\n%s", got, tc.want)
			}
		})
	}
}

func TestAddedImports(t *testing.T) {
	before := "package main\n\nimport (\n\t\"fmt\"\n\t\"io\"\n)\n\nfunc main() { fmt.Println(strconv.Itoa(1)) }\n"
	// What gopls organize imports makes of it: adds strconv, drops unused io.
	after := "package main\n\nimport (\n\t\"fmt\"\n\t\"strconv\"\n)\n\nfunc main() { fmt.Println(strconv.Itoa(1)) }\n"
	got := addedImports(before, after)
	if len(got) != 1 || got[0].path != "strconv" {
		t.Errorf("addedImports = %v, want strconv only", got)
	}
}

func TestIsImported(t *testing.T) {
	src := "package main\n\nimport (\n\t\"fmt\"\n\tyaml \"gopkg.in/yaml.v3\"\n\t\"github.com/x/y/v2\"\n\t\"gopkg.in/check.v1\"\n)\n\nfunc main() {\n\tbroken(\n"
	for name, want := range map[string]bool{
		"fmt": true, "yaml": true, "y": true, "check": true,
		"strconv": false, "v2": false, "main": false,
	} {
		if got := isImported(src, name); got != want {
			t.Errorf("isImported(%q) = %t, want %t", name, got, want)
		}
	}
}

func TestQualifiedBeforeCursor(t *testing.T) {
	cases := []struct {
		before, pkg string
	}{
		{"\tstrconv.Itoa(", "strconv"},
		{"x := strconv.Itoa ", "strconv"},
		{"f(os.Stdout,", "os"},
		{"\tstrconv.Itoa", ""}, // still typing
		{"\tstrconv.", ""},     // still typing
		{"\ta.b.C(", ""},       // a selector of a selector
		{"\tx.y[", "x"},        // can't tell a package from a variable: gopls decides
		{"\t1.5 ", ""},         // a number
		{"(", ""},
		{"", ""},
	}
	for _, tc := range cases {
		pkg, ok := qualifiedBeforeCursor(tc.before)
		if pkg != tc.pkg || ok != (tc.pkg != "") {
			t.Errorf("qualifiedBeforeCursor(%q) = %q, %t; want %q", tc.before, pkg, ok, tc.pkg)
		}
	}
}

func TestGuessPackageName(t *testing.T) {
	for p, want := range map[string]string{
		"fmt": "fmt", "math/rand": "rand", "github.com/x/y/v2": "y",
		"gopkg.in/yaml.v3": "yaml", "example.com/go-foo": "go-foo",
	} {
		if got := guessPackageName(p); got != want {
			t.Errorf("guessPackageName(%q) = %q, want %q", p, got, want)
		}
	}
}
