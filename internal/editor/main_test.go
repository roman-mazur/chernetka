package editor

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(runTestsWithClipboard(m))
}

func runTestsWithClipboard(m *testing.M) int {
	// Restore current clipboard text after the test run.
	defer clipboard.Write(clipboard.Read())
	return m.Run()
}
