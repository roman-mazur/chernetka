// Package ci keeps the configuration of the GitHub Actions workflows in CUE.
// The workflows in .github/workflows are generated from it:
//
//	go generate ./internal/ci
//
// The test workflow checks that they are up to date.
package ci

//go:generate go tool cue cmd gen
