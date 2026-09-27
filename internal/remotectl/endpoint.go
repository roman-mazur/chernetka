package remotectl

import (
	"path/filepath"
	"strings"

	"rmazur.io/chernetka/internal"
	"rmazur.io/chernetka/internal/vt/tabscope"
)

// Endpoint names a local socket a Server listens on.
type Endpoint string

// EditorEndpoint is used by the main editor process to receive files to open.
const EditorEndpoint Endpoint = "ctl"

// socketDir resolves the directory to put the sockets in. Overridden in tests.
var socketDir = internal.UserDir

// Scoped returns a separate endpoint for the scope, for example a terminal tab.
// The empty scope returns the endpoint unchanged.
func (ep Endpoint) Scoped(scope string) Endpoint {
	if scope == "" {
		return ep
	}
	// The scope becomes a part of a file name.
	scope = strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' || r == '-' || r == '_' {
			return r
		}
		return '_'
	}, scope)
	return ep + "-" + Endpoint(scope)
}

// InTab returns the endpoint scoped to the terminal tab the process runs in (see tabscope.ID).
func (ep Endpoint) InTab() Endpoint {
	return ep.Scoped(tabscope.ID())
}

func (ep Endpoint) path() (string, error) {
	dir, err := socketDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, string(ep)+".socket"), nil
}
