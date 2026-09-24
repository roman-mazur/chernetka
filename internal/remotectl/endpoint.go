package remotectl

import (
	"path/filepath"

	"rmazur.io/chernetka/internal"
)

// Endpoint names a local socket a Server listens on.
type Endpoint string

// EditorEndpoint is used by the main editor process to receive files to open.
const EditorEndpoint Endpoint = "ctl"

// socketDir resolves the directory to put the sockets in. Overridden in tests.
var socketDir = internal.UserDir

func (ep Endpoint) path() (string, error) {
	dir, err := socketDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, string(ep)+".socket"), nil
}
