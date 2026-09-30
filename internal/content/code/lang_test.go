package code

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSyntaxForPath(t *testing.T) {
	for path, want := range map[string]string{
		"main.go":                                "go",
		"a/b/MAIN.GO":                            "go",
		"flake.nix":                              "nix",
		"schema.CUE":                             "cue",
		"package.json":                           "json",
		"a/flake.lock":                           "json",
		"settings.jsonc":                         "jsonc",
		"ci.yml":                                 "yaml",
		"a/b/run.BASH":                           "shell",
		"/home/u/.zshrc":                         "shell",
		"bashrc":                                 "",
		"notes.markdown":                         "markdown",
		"arch.d2":                                "d2",
		"DB.SQL":                                 "sql",
		"repo/.git/COMMIT_EDITMSG":               "git_commit_msg",
		"repo/.git/rebase-merge/git-rebase-todo": "git_rebase",
		"":                                       "",
		"Makefile":                               "",
		"go":                                     "",
		"archive.go.bak":                         "",
		"dir.go/file.txt":                        "",
	} {
		got := ""
		if s := SyntaxForPath(path); s != nil {
			got = s.Name
		}
		if got != want {
			t.Errorf("SyntaxForPath(%q) = %q, want %q", path, got, want)
		}
	}
}

func TestSyntaxNamesAreUnique(t *testing.T) {
	seen := make(map[string]bool)
	for _, s := range syntaxes {
		if seen[s.Name] {
			t.Errorf("syntax %q is registered twice", s.Name)
		}
		seen[s.Name] = true
	}
}

func TestFindRoot(t *testing.T) {
	dir := t.TempDir()
	nested := filepath.Join(dir, "a", "b")
	if err := os.MkdirAll(filepath.Join(nested), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(dir, "a", "cue.mod"), 0o755); err != nil {
		t.Fatal(err)
	}
	if got, want := findRoot("cue.mod")(nested), filepath.Join(dir, "a"); got != want {
		t.Errorf("findRoot(cue.mod)(%q) = %q, want %q", nested, got, want)
	}
	if got := findRoot("no.such.marker")(nested); got != nested {
		t.Errorf("findRoot without a marker = %q, want the directory itself", got)
	}
}
