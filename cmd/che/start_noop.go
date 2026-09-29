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

// runInNewPane does nothing
func runInNewPane(_, _ string) error {
	return errorTerminalNotSupported
}
