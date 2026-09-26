package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestFetch(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git is not installed")
	}

	repo := t.TempDir()
	writeFiles(t, repo, map[string]string{
		"src/parser.c":             "parser",
		"src/scanner.c":            "scanner",
		"src/tree_sitter/parser.h": "header",
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
		"binding.go":           "package lang",
		"src/parser.c":         "old parser",
		"src/dropped_upstream": "stale",
	})
	got, err := fetch(repo, "trunk", dir, "lang")
	if err != nil {
		t.Fatal(err)
	}
	if got != commit {
		t.Errorf("fetch returned commit %q, want %q", got, commit)
	}

	for name, want := range map[string]string{
		"binding.go":               "package lang",
		"src/parser.c":             "parser",
		"src/scanner.c":            "scanner",
		"src/tree_sitter/parser.h": "header",
		"LICENSE":                  "MIT",
	} {
		data, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			t.Error(err)
		} else if string(data) != want {
			t.Errorf("%s = %q, want %q", name, data, want)
		}
	}
	for _, name := range []string{"src/grammar.json", "src/dropped_upstream"} {
		if _, err := os.Stat(filepath.Join(dir, name)); err == nil {
			t.Errorf("%s exists, want it gone", name)
		}
	}
	upstream, err := os.ReadFile(filepath.Join(dir, "upstream.go"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(upstream), "package lang\n") || !strings.Contains(string(upstream), repo+"@"+commit) {
		t.Errorf("upstream.go doesn't record the package and the commit:\n%s", upstream)
	}
}

func TestFetchMissingFile(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git is not installed")
	}

	repo := t.TempDir()
	writeFiles(t, repo, map[string]string{"src/parser.c": "parser"})
	gitRun(t, repo, "init", "--quiet", "--initial-branch=trunk")
	gitRun(t, repo, "add", ".")
	gitRun(t, repo, "-c", "user.name=test", "-c", "user.email=test@example.com", "commit", "--quiet", "-m", "grammar")

	dir := t.TempDir()
	writeFiles(t, dir, map[string]string{"src/parser.c": "current"})
	if _, err := fetch(repo, "trunk", dir, "lang"); err == nil {
		t.Fatal("fetch succeeded without a scanner and a license")
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
