//go:build darwin

package tabscope

import (
	"os"
	"os/exec"
	"strings"
)

// terminalTab asks Ghostty for its selected tab. Ghostty does not tell a process which terminal
// surface it runs in, but the tab of the focused window is the one where the user has just
// typed the command (or where the editor has just split a pane) in all but rare cases.
func terminalTab() string {
	if os.Getenv("TERM_PROGRAM") != "ghostty" {
		return ""
	}
	out, err := exec.Command("osascript", "-e", `tell application "Ghostty" to get id of selected tab of front window`).Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}
