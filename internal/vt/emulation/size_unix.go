//go:build darwin || linux

package emulation

import (
	"os"

	"golang.org/x/sys/unix"
)

func setSize(pty *os.File, cols, rows int) error {
	return unix.IoctlSetWinsize(int(pty.Fd()), unix.TIOCSWINSZ, &unix.Winsize{Row: uint16(rows), Col: uint16(cols)})
}
