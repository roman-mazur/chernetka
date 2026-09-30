package code

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
)

// Syntax represents one of the syntaxes supported by the editor.
type Syntax struct {
	// Name identifies the syntax. For the syntaxes with a language server, it's
	// also the LSP language identifier.
	Name      string
	PathMatch SyntaxPathMatcher
	// LineCommentPrefix starts a line comment, like "//". Empty if the syntax has none.
	LineCommentPrefix string
	// RootFinder returns the workspace root for a file in dir, like the directory
	// with go.mod. Nil if the syntax has no notion of a workspace.
	RootFinder func(dir string) string
}

// SyntaxPathMatcher reports whether the file at the path is written in a syntax.
type SyntaxPathMatcher func(p string) bool

var (
	Go = &Syntax{
		Name:              "go",
		PathMatch:         suffixMatch(".go"),
		LineCommentPrefix: "//",
		RootFinder:        findRoot("go.mod"),
	}
	Nix = &Syntax{
		Name:              "nix",
		PathMatch:         suffixMatch(".nix"),
		LineCommentPrefix: "#",
	}
	CUE = &Syntax{
		Name:              "cue",
		PathMatch:         suffixMatch(".cue"),
		LineCommentPrefix: "//",
		RootFinder:        findRoot("cue.mod"),
	}
	JSON = &Syntax{
		Name: "json",
		PathMatch: compose(
			suffixMatch(".json", ".jsonl"),
			filenameMatch("flake.lock"), // Nix flakes pin their inputs in JSON.
		),
	}
	JSONC = &Syntax{
		Name:              "jsonc",
		PathMatch:         suffixMatch(".jsonc"),
		LineCommentPrefix: "//",
	}
	YAML = &Syntax{
		Name:              "yaml",
		PathMatch:         suffixMatch(".yaml", ".yml"),
		LineCommentPrefix: "#",
	}
	Shell = &Syntax{
		Name: "shell",
		PathMatch: compose(
			suffixMatch(".sh", ".bash", ".zsh"),
			filenameMatch(shellFileNames...),
		),
		LineCommentPrefix: "#",
	}
	Markdown = &Syntax{
		Name:      "markdown",
		PathMatch: suffixMatch(".markdown", ".md"),
	}
	D2 = &Syntax{
		Name:              "d2",
		PathMatch:         suffixMatch(".d2"),
		LineCommentPrefix: "#",
	}
	SQL = &Syntax{
		Name:              "sql",
		PathMatch:         suffixMatch(".sql"),
		LineCommentPrefix: "--",
	}
	GitCommitMsg = &Syntax{
		Name:              "git_commit_msg",
		PathMatch:         func(p string) bool { return strings.HasSuffix(p, ".git/COMMIT_EDITMSG") },
		LineCommentPrefix: "#",
	}
	GitRebaseTodo = &Syntax{
		Name:              "git_rebase",
		PathMatch:         func(p string) bool { return strings.HasSuffix(p, ".git/rebase-merge/git-rebase-todo") },
		LineCommentPrefix: "#",
	}

	// syntaxes are all the supported syntaxes, in the order they are matched.
	syntaxes = []*Syntax{Go, Nix, CUE, JSON, JSONC, YAML, Shell, Markdown, D2, SQL, GitCommitMsg, GitRebaseTodo}

	// shellFileNames are the shell scripts recognized by their name alone. The zsh
	// ones are parsed as bash, which covers most of what goes into them.
	shellFileNames = []string{
		".bashrc", ".bash_profile", ".bash_login", ".bash_logout", ".profile",
		".zshrc", ".zshenv", ".zprofile", ".zlogin", ".zlogout",
		".envrc",
	}
)

// SyntaxForPath returns the syntax of the file at the path, or nil if it's unknown.
func SyntaxForPath(p string) *Syntax {
	for _, s := range syntaxes {
		if s.PathMatch(p) {
			return s
		}
	}
	return nil
}

// suffixMatch matches the paths ending with one of the suffixes, ignoring the case.
func suffixMatch(s ...string) SyntaxPathMatcher {
	return func(p string) bool {
		p = strings.ToLower(p)
		for i := range s {
			if strings.HasSuffix(p, s[i]) {
				return true
			}
		}
		return false
	}
}

func filenameMatch(names ...string) SyntaxPathMatcher {
	return func(p string) bool { return slices.Contains(names, filepath.Base(p)) }
}

func compose(matchers ...SyntaxPathMatcher) SyntaxPathMatcher {
	return func(p string) bool {
		for i := range matchers {
			if matchers[i](p) {
				return true
			}
		}
		return false
	}
}

// findRoot returns a function walking up from dir looking for the file (or directory)
// marking the workspace root, like go.mod. The function returns the directory containing it,
// or dir if none is found before the filesystem root.
func findRoot(marker string) func(dir string) string {
	return func(dir string) string {
		for d := dir; ; {
			if _, err := os.Stat(filepath.Join(d, marker)); err == nil {
				return d
			}
			parent := filepath.Dir(d)
			if parent == d {
				return dir
			}
			d = parent
		}
	}
}
