package ci

import (
	"list"
	"strings"

	"cue.dev/x/githubactions"
)

// #target is a platform the release ships a binary for.
//
// The binary needs cgo (tree-sitter is C code), so a C toolchain for the target must be at hand.
// Linux and Windows build on Linux: Windows with MinGW-w64, which links against the system DLLs
// only. macOS builds on a macOS runner: cross-compiling to it needs the Apple SDK, which is
// licensed for Apple hardware only. Apple's clang targets x86_64 from an arm64 runner, since Go
// passes it -arch x86_64.
#target: {
	goos:   string
	goarch: "amd64"
	runner: string
	// aptPackages are the packages installing the C toolchain the build uses.
	aptPackages: *[] | [...string]
	// cc is the C compiler; empty for the default one.
	cc:  string | *""
	ext: string | *""

	asset: "che-\(goos)-\(goarch)\(ext)"
}

releaseTargets: [...#target] & [
	{goos: "linux", runner: "ubuntu-latest"},
	{
		goos:   "windows"
		runner: "ubuntu-latest"
		aptPackages: ["gcc-mingw-w64-x86-64"]
		cc:  "x86_64-w64-mingw32-gcc"
		ext: ".exe"
	},
	{goos: "darwin", runner: "macos-latest"},
]

workflows: release: githubactions.#Workflow & {
	name: "Release Binaries"
	on: {
		workflow_dispatch: {}
		release: types: ["created"]
	}

	for t in releaseTargets {
		jobs: "build-\(t.goos)-\(t.goarch)": {
			"runs-on": t.runner
			permissions: {
				"id-token": "write"
				contents:   "write"
			}
			steps: list.Concat([#goSetup, [
				if len(t.aptPackages) > 0 {
					name: "Install the C toolchain"
					run:  """
						sudo apt-get update
						sudo apt-get install -y \(strings.Join(t.aptPackages, " "))
						"""
				},
				{
					name: "Print go env"
					run: """
						go version
						go env
						"""
				},
				{
					name: "Download modules"
					run:  "go mod download"
				},
				{
					name: "Build"
					env: {
						GOOS:        t.goos
						GOARCH:      t.goarch
						CGO_ENABLED: "1"
						if t.cc != "" {
							CC: t.cc
						}
					}
					run: """
						mkdir -p ./out
						go build -ldflags='-s -w' -trimpath -tags=netgo -o ./out/\(t.asset) ./cmd/che
						"""
				},
				{
					name: "Upload release asset"
					uses: "actions/upload-release-asset@v1"
					with: {
						upload_url:         "${{ github.event.release.upload_url }}"
						asset_path:         "./out/\(t.asset)"
						asset_name:         t.asset
						asset_content_type: "application/octet-stream"
					}
					env: GITHUB_TOKEN: "${{ secrets.GITHUB_TOKEN }}"
				},
			]])
		}
	}
}
