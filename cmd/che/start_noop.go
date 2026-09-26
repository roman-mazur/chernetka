//go:build !darwin

package main

import "errors"

var errorTerminalNotSupported = errors.New("terminal not supported")

// openMainEditor does nothing
func openMainEditor(_, _ string) error {
	return errorTerminalNotSupported
}

// launchImageViewer does nothing
func launchImageViewer() error {
	return errorTerminalNotSupported
}

// launchDiffViewer does nothing
func launchDiffViewer(_ string) error {
	return errorTerminalNotSupported
}
