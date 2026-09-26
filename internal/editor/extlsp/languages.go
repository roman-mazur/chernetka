package extlsp

import (
	"context"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"rmazur.io/chernetka/internal"
	"rmazur.io/chernetka/internal/lsp"
)

// language describes how to get a language server for files of a language.
type language struct {
	id         string   // the LSP language identifier
	extensions []string // the file name extensions, lowercase

	// root returns the workspace root to start the server with for a file in dir.
	root func(dir string) string
	// start launches the server for the workspace in rootDir.
	start func(ctx context.Context, le *Integration, rootDir string) (lspClient, error)

	// rankedCompletion is set if the server sorts completion items by
	// relevance. Items of other servers are ranked by the typed prefix.
	rankedCompletion bool
	// goImports is set to add the Go imports missing for qualified names typed.
	goImports bool
}

var languages = []*language{
	{
		id:               "go",
		extensions:       []string{".go"},
		root:             func(dir string) string { return findRoot(dir, "go.mod") },
		start:            startGopls,
		rankedCompletion: true,
		goImports:        true,
	},
	{
		id:         "cue",
		extensions: []string{".cue"},
		root:       func(dir string) string { return findRoot(dir, "cue.mod") },
		start:      startCUE,
	},
}

func languageForPath(path string) *language {
	ext := strings.ToLower(filepath.Ext(path))
	for _, lang := range languages {
		if slices.Contains(lang.extensions, ext) {
			return lang
		}
	}
	return nil
}

// startGopls starts gopls with the Go toolchain the workspace requires.
func startGopls(ctx context.Context, le *Integration, rootDir string) (lspClient, error) {
	var cache string
	if dir, err := internal.UserDir(); err == nil {
		cache = filepath.Join(dir, "gopls")
	}
	return lsp.Start(ctx, rootDir, lsp.Options{
		GoplsCache: cache,
		// Calls come with brackets and the cursor placed inside them.
		SnippetSupport: true,
		Logf:           le.Logf,
	})
}

// startCUE starts the language server built into the cue command.
func startCUE(ctx context.Context, le *Integration, rootDir string) (lspClient, error) {
	return lsp.StartWith(ctx, rootDir, lsp.Options{
		Command:        "cue",
		Args:           []string{"lsp"},
		SnippetSupport: true,
		Logf:           le.Logf,
	})
}

// findRoot walks up from dir looking for the file (or directory) marking the
// workspace root, like go.mod. Returns the directory containing it, or dir if
// none is found before the filesystem root.
func findRoot(dir, marker string) string {
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
