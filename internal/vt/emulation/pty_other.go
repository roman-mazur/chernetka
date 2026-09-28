//go:build !darwin && !linux

package emulation

import (
	"errors"
	"os"
)

func openPTY() (pty, tty *os.File, err error) { return nil, nil, errors.ErrUnsupported }

func setSize(*os.File, int, int) error { return errors.ErrUnsupported }
