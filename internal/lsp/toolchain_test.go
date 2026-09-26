package lsp

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestFindToolchain(t *testing.T) {
	t.Setenv("GOWORK", "") // discover go.work files as usual
	t.Setenv("GOFLAGS", "")

	root := t.TempDir()
	write := func(path, content string) {
		t.Helper()
		_ = os.MkdirAll(filepath.Dir(path), 0o755)
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	// Versions not newer than the running Go, so nothing is downloaded.
	write(filepath.Join(root, "mod", "go.mod"), "module example.com/mod\n\ngo 1.22\n")
	write(filepath.Join(root, "ws", "go.work"), "go 1.22\n\nuse ./a\n")
	write(filepath.Join(root, "ws", "a", "go.mod"), "module example.com/a\n\ngo 1.22\n")
	_ = os.MkdirAll(filepath.Join(root, "plain"), 0o755)

	cases := []struct {
		dir, gomod, gowork, source string
	}{
		{dir: "mod", gomod: "mod/go.mod", source: "mod/go.mod"},
		{dir: "ws/a", gomod: "ws/a/go.mod", gowork: "ws/go.work", source: "ws/go.work"},
		{dir: "plain", source: "the installed Go"},
	}
	for _, tc := range cases {
		t.Run(tc.dir, func(t *testing.T) {
			abs := func(p string) string {
				if p == "" || !strings.Contains(p, "/") {
					return p
				}
				return filepath.Join(root, p)
			}
			got, err := FindToolchain(context.Background(), filepath.Join(root, tc.dir))
			if err != nil {
				t.Fatal(err)
			}
			if !samePath(got.GOMOD, abs(tc.gomod)) || !samePath(got.GOWORK, abs(tc.gowork)) {
				t.Errorf("GOMOD = %q, GOWORK = %q; want %q, %q", got.GOMOD, got.GOWORK, abs(tc.gomod), abs(tc.gowork))
			}
			if src := got.Source(); !samePath(src, abs(tc.source)) {
				t.Errorf("Source() = %q, want %q", src, abs(tc.source))
			}
			if got.Version == "" || got.GOROOT == "" {
				t.Errorf("incomplete toolchain: %+v", got)
			}
		})
	}
}

// samePath compares paths resolving symlinks (like /var and /private/var on macOS).
func samePath(a, b string) bool {
	if a == b {
		return true
	}
	ra, err1 := filepath.EvalSymlinks(a)
	rb, err2 := filepath.EvalSymlinks(b)
	return err1 == nil && err2 == nil && ra == rb
}

func TestToolchain_Env(t *testing.T) {
	sep := string(filepath.ListSeparator)
	tc := Toolchain{Version: "go1.27.0", GOROOT: filepath.FromSlash("/sdk/go1.27.0")}
	got := tc.Env([]string{"HOME=/home/x", "PATH=/usr/bin" + sep + "/bin", "GOROOT=/usr/local/go"})
	want := []string{
		"HOME=/home/x",
		"PATH=" + filepath.FromSlash("/sdk/go1.27.0/bin") + sep + "/usr/bin" + sep + "/bin",
		"GOROOT=" + filepath.FromSlash("/sdk/go1.27.0"),
	}
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Errorf("Env =\n%s\nwant\n%s", strings.Join(got, "\n"), strings.Join(want, "\n"))
	}
}

func TestToolchain_Gopls(t *testing.T) {
	// Any Go binary works as a gopls stand-in: only its build info is read.
	exe, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	bin := t.TempDir()
	fakeGopls := filepath.Join(bin, goplsExe)
	copyFile(t, exe, fakeGopls)
	t.Setenv("PATH", bin)

	var logs []string
	logf := func(format string, args ...any) { logs = append(logs, format) }
	cache := t.TempDir()

	t.Run("gopls in PATH is recent enough", func(t *testing.T) {
		tc := Toolchain{Version: runtime.Version()}
		got, err := tc.Gopls(context.Background(), cache, logf)
		if err != nil || got != fakeGopls {
			t.Errorf("Gopls() = %q, %v; want %q", got, err, fakeGopls)
		}
	})

	t.Run("gopls in PATH is too old", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // don't actually build it
		tc := Toolchain{Version: "go1.999.0", GOROOT: t.TempDir()}
		if got, err := tc.Gopls(ctx, cache, logf); err == nil {
			t.Errorf("Gopls() = %q, want a build attempt failing", got)
		}
		if len(logs) == 0 || !strings.Contains(logs[len(logs)-1], "building gopls") {
			t.Errorf("logs = %q, want a build attempt", logs)
		}
	})

	t.Run("previously built gopls is reused", func(t *testing.T) {
		tc := Toolchain{Version: "go1.999.0"}
		// The stand-in has no module version: the latest gopls would be built.
		built := filepath.Join(cache, "gopls-latest-go1.999.0", goplsExe)
		copyFile(t, exe, built)
		got, err := tc.Gopls(context.Background(), cache, logf)
		if err != nil || got != built {
			t.Errorf("Gopls() = %q, %v; want %q", got, err, built)
		}
	})
}

func copyFile(t *testing.T, from, to string) {
	t.Helper()
	_ = os.MkdirAll(filepath.Dir(to), 0o755)
	in, err := os.Open(from)
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()
	out, err := os.OpenFile(to, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o755)
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()
	if _, err := io.Copy(out, in); err != nil {
		t.Fatal(err)
	}
}
