package logger

import (
	"log"
	"os"
	"path/filepath"
)

// UserLogFile returns a logger that appends to a std log file of the editor.
// In case of an error a Discard function is returned.
func UserLogFile() (Func, error) {
	userDir, err := os.UserHomeDir()
	if err != nil {
		return Discard, err
	}
	logPath := filepath.Join(userDir, ".edit", "logs", "editor.log")
	_ = os.MkdirAll(filepath.Dir(logPath), 0755)
	f, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		return Discard, err
	}
	return log.New(f, "", log.LstdFlags|log.Lshortfile).Printf, nil
}
