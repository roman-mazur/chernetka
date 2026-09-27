// Package tabscope identifies the terminal tab a process runs in.
// Editor processes started in different tabs use different sockets, while the processes
// in the panes of the same tab talk to each other.
package tabscope

import (
	"os"
	"sync"
)

// EnvVar holds the tab ID for the processes started in the panes the editor creates,
// so that they don't need to ask the terminal.
const EnvVar = "CHE_TAB"

// ID returns the ID of the terminal tab the process runs in, or "" if it's unknown.
// The first call may take a while to talk to the terminal.
var ID = sync.OnceValue(func() string {
	if id := os.Getenv(EnvVar); id != "" {
		return id
	}
	id := terminalTab()
	if id != "" {
		// Processes replacing this one or started by it stay in the same tab.
		_ = os.Setenv(EnvVar, id)
	}
	return id
})
