//go:build !darwin

package main

import "errors"

var errorTerminalNotSupported = errors.New("terminal not supported")

// openMainEditor reports that the terminal is not supported: panes are split in Ghostty on macOS only.
func openMainEditor(_, _ string) error {
	return errorTerminalNotSupported
}

// launchImageViewer reports that the terminal is not supported.
func launchImageViewer() error {
	return errorTerminalNotSupported
}

// launchCommandProxy reports that the terminal is not supported.
func launchCommandProxy(_, _ string) error {
	return errorTerminalNotSupported
}
