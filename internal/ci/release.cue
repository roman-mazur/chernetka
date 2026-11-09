package ci

import (
	"list"

	"cue.dev/x/githubactions"
)

// #target is a platform the release ships a binary for.
//
// The binary needs cgo (tree-sitter is C code), so each target builds on a runner of its OS, with
// its native C toolchain. Cross-compiling to macOS needs the Apple SDK, which is licensed for
// Apple hardware only, so macOS needs its own runner, and the other targets follow suit. Apple's
// clang targets x86_64 from an arm64 runner, since Go passes it -arch x86_64.
#target: {
	goos:   string
	goarch: "amd64"
	runner: string
	ext:    string | *""

	asset: "che-\(goos)-\(goarch)\(ext)"
}

releaseTargets: [...#target] & [
	{goos: "linux", runner: "ubuntu-latest"},
	{goos: "windows", runner: "windows-latest", ext: ".exe"},
	{goos: "darwin", runner: "macos-latest"},
]

workflows: release: githubactions.#Workflow & {
	name: "Release Binaries"
	on: {
		workflow_dispatch: {}
		release: types: ["created"]
	}

	defaults: run: shell: "bash"

	for t in releaseTargets {
		jobs: "build-\(t.goos)-\(t.goarch)": {
			"runs-on": t.runner
			permissions: {
				"id-token": "write"
				contents:   "write"
			}
			steps: list.Concat([#goSetup, [
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
					}
					run: """
						mkdir -p ./out
						go build -ldflags='-s -w' -trimpath -tags=netgo -o ./out/\(t.asset) ./cmd/che
						"""
				},
				{
					// Manual runs only check that the binaries build: there is no release to
					// upload them to.
					name: "Upload release asset"
					if:   "github.event_name == 'release'"
					env: {
						GH_TOKEN: "${{ github.token }}"
						GH_REPO:  "${{ github.repository }}"
						TAG:      "${{ github.event.release.tag_name }}"
					}
					run: "gh release upload \"$TAG\" ./out/\(t.asset) --clobber"
				},
			]])
		}
	}
}
