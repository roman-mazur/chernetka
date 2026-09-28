//go:build linux

package emulation

import (
	"os"
	"strconv"

	"golang.org/x/sys/unix"
)

// openPTY opens a new pseudo-terminal returning its controlling side and the terminal device.
func openPTY() (pty, tty *os.File, err error) {
	pty, err = os.OpenFile("/dev/ptmx", os.O_RDWR|unix.O_NOCTTY, 0)
	if err != nil {
		return nil, nil, err
	}
	fd := int(pty.Fd())
	if err := unix.IoctlSetPointerInt(fd, unix.TIOCSPTLCK, 0); err != nil {
		_ = pty.Close()
		return nil, nil, err
	}
	n, err := unix.IoctlGetInt(fd, unix.TIOCGPTN)
	if err != nil {
		_ = pty.Close()
		return nil, nil, err
	}
	tty, err = os.OpenFile("/dev/pts/"+strconv.Itoa(n), os.O_RDWR|unix.O_NOCTTY, 0)
	if err != nil {
		_ = pty.Close()
		return nil, nil, err
	}
	return pty, tty, nil
}
