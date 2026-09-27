//go:build !darwin

package tabscope

// terminalTab does not know how to talk to the terminal on this platform.
func terminalTab() string { return "" }
