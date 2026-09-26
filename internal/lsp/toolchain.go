package lsp

import (
	"bytes"
	"context"
	"debug/buildinfo"
	"encoding/json"
	"errors"
	"fmt"
	"go/version"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var goplsExe = map[bool]string{false: "gopls", true: "gopls.exe"}[runtime.GOOS == "windows"]

// Toolchain is the Go toolchain the go command selects for a directory: the
// one required by the go.work of the workspace the directory belongs to, or by
// its go.mod otherwise, honoring their go and toolchain lines and GOTOOLCHAIN
// (see https://go.dev/doc/toolchain). It may be newer than the installed Go,
// in which case the go command downloads it into the module cache.
type Toolchain struct {
	Version string // like "go1.27.0"
	GOROOT  string

	// The files that select the version: GOWORK takes precedence over GOMOD.
	// Empty if the directory is outside a workspace or a module.
	GOWORK, GOMOD string
}

// FindToolchain asks the go command which toolchain it uses in dir.
func FindToolchain(ctx context.Context, dir string) (Toolchain, error) {
	cmd := exec.CommandContext(ctx, "go", "env", "-json", "GOVERSION", "GOROOT", "GOWORK", "GOMOD")
	cmd.Dir = dir
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		return Toolchain{}, fmt.Errorf("go env in %s: %w: %s", dir, err, strings.TrimSpace(stderr.String()))
	}
	var env struct{ GOVERSION, GOROOT, GOWORK, GOMOD string }
	if err := json.Unmarshal(out, &env); err != nil {
		return Toolchain{}, fmt.Errorf("go env output: %w", err)
	}
	tc := Toolchain{Version: env.GOVERSION, GOROOT: env.GOROOT, GOWORK: env.GOWORK, GOMOD: env.GOMOD}
	if tc.GOWORK == "off" {
		tc.GOWORK = ""
	}
	if tc.GOMOD == os.DevNull {
		tc.GOMOD = ""
	}
	if !version.IsValid(tc.Version) || tc.GOROOT == "" {
		return Toolchain{}, fmt.Errorf("unexpected go env output: %s", out)
	}
	return tc, nil
}

// Source describes which file selects the toolchain.
func (tc Toolchain) Source() string {
	switch {
	case tc.GOWORK != "":
		return tc.GOWORK
	case tc.GOMOD != "":
		return tc.GOMOD
	default:
		return "the installed Go"
	}
}

// Env returns env changed so that the go command found in PATH is the
// toolchain's one. That's what the go command itself does for the programs
// it runs with "go run" or "go test".
func (tc Toolchain) Env(env []string) []string {
	bin := filepath.Join(tc.GOROOT, "bin")
	res := make([]string, 0, len(env)+2)
	path := bin
	for _, kv := range env {
		switch k, v, _ := strings.Cut(kv, "="); k {
		case "PATH":
			path = bin + string(filepath.ListSeparator) + v
		case "GOROOT":
			// Dropped: a GOROOT of another Go must not be used with this one.
		default:
			res = append(res, kv)
		}
	}
	return append(res, "PATH="+path, "GOROOT="+tc.GOROOT)
}

// Gopls returns the path of a gopls built with a Go at least as new as the
// toolchain. gopls type checks with the go/types it was built with, which
// rejects packages requiring a newer Go version, so a gopls built with an
// older Go can't be used for a module that requires a newer one.
//
// The gopls found in PATH is used if it's recent enough. Otherwise the same
// gopls version (the latest one if there is no gopls in PATH) is built with
// the toolchain and installed into cacheDir, where it's reused later.
// Building takes a while: it may need to download gopls first.
func (tc Toolchain) Gopls(ctx context.Context, cacheDir string, logf func(format string, args ...any)) (string, error) {
	goplsVersion := "latest"
	if path, err := exec.LookPath("gopls"); err == nil {
		info, err := buildinfo.ReadFile(path)
		if err == nil && version.Compare(version.Lang(info.GoVersion), version.Lang(tc.Version)) >= 0 {
			return path, nil
		}
		if err == nil && info.Main.Version != "" && info.Main.Version != "(devel)" {
			goplsVersion = info.Main.Version
		}
		if err == nil {
			logf("%s is built with %s, but %s requires %s", path, info.GoVersion, tc.Source(), tc.Version)
		}
	}

	dir := filepath.Join(cacheDir, "gopls-"+goplsVersion+"-"+tc.Version)
	bin := filepath.Join(dir, goplsExe)
	if _, err := os.Stat(bin); err == nil {
		return bin, nil
	}

	logf("building gopls@%s with %s into %s", goplsVersion, tc.Version, dir)
	if err := os.MkdirAll(cacheDir, 0o755); err != nil {
		return "", err
	}
	// Install into a temporary directory first: another editor may be doing
	// the same concurrently.
	tmp, err := os.MkdirTemp(cacheDir, "build-")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(tmp)

	cmd := exec.CommandContext(ctx, filepath.Join(tc.GOROOT, "bin", "go"), "install", "golang.org/x/tools/gopls@"+goplsVersion)
	cmd.Dir = tmp // Outside of any module or workspace.
	cmd.Env = append(tc.Env(os.Environ()), "GOBIN="+tmp, "GOTOOLCHAIN=local", "GOWORK=off")
	if out, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("go install gopls@%s: %w: %s", goplsVersion, err, bytes.TrimSpace(out))
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	if err := os.Rename(filepath.Join(tmp, goplsExe), bin); err != nil && !errors.Is(err, os.ErrExist) {
		if _, statErr := os.Stat(bin); statErr != nil {
			return "", err
		}
	}
	return bin, nil
}
