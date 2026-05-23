//go:build darwin

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// openMainEditor interacts with the terminal to launch the main editor process in a new pane.
func openMainEditor(path string) error {
	const appleScript = `
tell application "Ghostty"
    activate
    set currentTerm to focused terminal of selected tab of front window
    set newPane to split currentTerm direction right
    input text "%s \"%s\"" to newPane
    send key "enter" to newPane
end tell`

	cmd := exec.Command("osascript")
	fullScript := fmt.Sprintf(appleScript, os.Args[0], path)
	cmd.Stdin = strings.NewReader(fullScript)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("could not open editor: %w, output: %s, script: %s", err, string(out), fullScript)
	}
	return nil
}
