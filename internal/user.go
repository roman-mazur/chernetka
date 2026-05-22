package internal

import (
	"os"
	"path/filepath"
)

// UserDir returns the local user directory owned by the editor.
func UserDir() (string, error) {
	userDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	p := filepath.Join(userDir, ".edit")
	_ = os.MkdirAll(p, 0755)
	return p, nil
}
