// Package debugflags provides API to discover the user settings that affect chernetka behaviour
// mainly for the debugging purposes. Available flags are listed below:
// - logdebug - enables debug logging of the editor and extensions
// - loginputs - enables dumping the input bytes received by the editor
package debugflags

import (
	"os"
	"strings"
)

func IsEnabled(name string) bool { return flags[name] == "1" }

func IsDisabled(name string) bool { return flags[name] == "0" }

func Value(name string) string { return flags[name] }

func init() {
	config := os.Getenv("CHEDEBUG")
	for pair := range strings.SplitSeq(config, ",") {
		key, val, _ := strings.Cut(pair, "=")
		flags[key] = val
	}
}

var flags = make(map[string]string)
