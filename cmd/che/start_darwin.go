//go:build darwin

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"rmazur.io/chernetka/internal/vt/tabscope"
)

// openMainEditor interacts with the terminal to launch the main editor process in a new pane.
// The main editor uses root as its project directory.
func openMainEditor(root, path string) error {
	return runInNewPane("right", fmt.Sprintf("%s -root %q %q", os.Args[0], root, path))
}

// launchImageViewer interacts with the terminal to launch che-img in a new pane below.
func launchImageViewer() error {
	return runInNewPane("down", fmt.Sprintf("%q", cheImgCommand()))
}

// runInNewPane splits the focused terminal pane of the editor's tab in the given direction
// and runs the command there. The new pane's shell knows the tab ID (see tabscope).
func runInNewPane(direction, command string) error {
	const appleScript = `
tell application "Ghostty"
    activate
    set targetTab to selected tab of front window
    set tabID to "%s"
    if tabID is not "" then
        repeat with w in windows
            repeat with t in tabs of w
                if id of t is tabID then set targetTab to t
            end repeat
        end repeat
    end if
    set currentTerm to focused terminal of targetTab
    set cfg to new surface configuration
    if tabID is not "" then set environment variables of cfg to {"%s=" & tabID}
    set newPane to split currentTerm direction %s with configuration cfg
    input text "%s" to newPane
    send key "enter" to newPane
end tell`

	cmd := exec.Command("osascript")
	fullScript := fmt.Sprintf(appleScript,
		escapeAppleScript(tabscope.ID()), tabscope.EnvVar, direction, escapeAppleScript(command))
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
