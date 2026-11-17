package lsp

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"go.lsp.dev/uri"
)

// TestRealGoplsRegistersWatchers starts a real gopls, directly and through
// its daemon, and checks that it asks to watch the workspace files once a
// document is open. It's skipped unless CHE_LSP_EVAL is set.
func TestRealGoplsRegistersWatchers(t *testing.T) {
	if os.Getenv("CHE_LSP_EVAL") == "" {
		t.Skip("set CHE_LSP_EVAL=1 to run tests against a real gopls")
	}
	for _, daemon := range []bool{false, true} {
		t.Run(fmt.Sprintf("daemon=%t", daemon), func(t *testing.T) {
			dir := t.TempDir()
			_ = os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module example.com/x\n\ngo 1.22\n"), 0o644)
			path := filepath.Join(dir, "main.go")
			src := "package main\n\nfunc main() {}\n"
			_ = os.WriteFile(path, []byte(src), 0o644)

			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()
			watched := make(chan []string, 4)
			opts := Options{
				Logf:       func(format string, args ...any) { t.Logf(format, args...) },
				WatchFiles: func(patterns []string) { watched <- patterns },
			}
			var (
				c   *Client
				err error
			)
			if daemon {
				c, err = Start(ctx, dir, opts)
			} else {
				c, err = StartWith(ctx, dir, opts)
			}
			if err != nil {
				t.Fatal(err)
			}
			defer func() { _ = c.Shutdown(context.Background()) }()
			if err := c.DidOpen(ctx, uri.File(path), "go", src, 1); err != nil {
				t.Fatal(err)
			}
			select {
			case patterns := <-watched:
				t.Logf("patterns: %q", patterns)
			case <-time.After(20 * time.Second):
				t.Fatal("gopls did not ask to watch the files")
			}
		})
	}
}
