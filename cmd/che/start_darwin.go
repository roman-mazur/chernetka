//go:build darwin

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"rmazur.io/chernetka/internal/editor"
)

// openMainEditor interacts with the terminal to launch the main editor process in a new pane.
func openMainEditor(_ *editor.Editor, path string) error {
	return runInNewPane("right", fmt.Sprintf("%s %q", os.Args[0], path))
}

// launchImageViewer interacts with the terminal to launch che-img in a new pane below.
func launchImageViewer() error {
	return runInNewPane("down", fmt.Sprintf("%q", cheImgCommand()))
}

// runInNewPane splits the focused terminal pane in the given direction and runs the command there.
func runInNewPane(direction, command string) error {
	const appleScript = `
tell application "Ghostty"
    activate
    set currentTerm to focused terminal of selected tab of front window
    set newPane to split currentTerm direction %s
    input text "%s" to newPane
    send key "enter" to newPane
end tell`

	cmd := exec.Command("osascript")
	fullScript := fmt.Sprintf(appleScript, direction, escapeAppleScript(command))
	cmd.Stdin = strings.NewReader(fullScript)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("could not run in a new pane: %w, output: %s, script: %s", err, string(out), fullScript)
	}
	return nil
}

func escapeAppleScript(s string) string {
	return strings.NewReplacer(`\`, `\\`, `"`, `\"`).Replace(s)
}
