//go:build unix

package vt

import (
	"errors"
	"os"

	"golang.org/x/term"
)

// openTTY opens the controlling terminal device.
func openTTY() (in, out *os.File, closeFn func(), err error) {
	f, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		return nil, nil, nil, err
	}
	if !term.IsTerminal(int(f.Fd())) {
		_ = f.Close()
		return nil, nil, nil, errors.New("/dev/tty is not a terminal")
	}
	return f, f, func() { _ = f.Close() }, nil
}
