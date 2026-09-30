package extlsp

import (
	"context"
	"path/filepath"

	"rmazur.io/chernetka/internal"
	"rmazur.io/chernetka/internal/content/code"
	"rmazur.io/chernetka/internal/lsp"
)

// language describes how to get a language server for files of a syntax.
type language struct {
	syntax *code.Syntax // with the root finder

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
		syntax:           code.Go,
		start:            startGopls,
		rankedCompletion: true,
		goImports:        true,
	},
	{
		syntax: code.CUE,
		start:  startCUE,
	},
}

// id returns the LSP language identifier.
func (l *language) id() string { return l.syntax.Name }

// root returns the workspace root to start the server with for a file in dir.
func (l *language) root(dir string) string { return l.syntax.RootFinder(dir) }

func languageForPath(path string) *language {
	s := code.SyntaxForPath(path)
	for _, lang := range languages {
		if lang.syntax == s {
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
