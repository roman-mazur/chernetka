//go:build !darwin

package main

import (
	"rmazur.io/chernetka/internal/editor"
)

var errorTerminalNotSupported = errors.New("terminal not supported")

// openMainEditor does nothing
func openMainEditor(_ *editor.Editor, _ string) error {
	return errorTerminalNotSupported
}
