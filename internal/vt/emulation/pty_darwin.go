//go:build darwin

package emulation

import (
	"bytes"
	"os"
	"unsafe"

	"golang.org/x/sys/unix"
)

// openPTY opens a new pseudo-terminal returning its controlling side and the terminal device.
func openPTY() (pty, tty *os.File, err error) {
	pty, err = os.OpenFile("/dev/ptmx", os.O_RDWR|unix.O_NOCTTY, 0)
	if err != nil {
		return nil, nil, err
	}
	fd := pty.Fd()
	var name [128]byte
	for _, req := range []struct {
		code uint
		arg  uintptr
	}{
		{unix.TIOCPTYGRANT, 0},
		{unix.TIOCPTYUNLK, 0},
		{unix.TIOCPTYGNAME, uintptr(unsafe.Pointer(&name[0]))},
	} {
		if _, _, errno := unix.Syscall(unix.SYS_IOCTL, fd, uintptr(req.code), req.arg); errno != 0 {
			_ = pty.Close()
			return nil, nil, errno
		}
	}
	name0, _, _ := bytes.Cut(name[:], []byte{0})
	tty, err = os.OpenFile(string(name0), os.O_RDWR|unix.O_NOCTTY, 0)
	if err != nil {
		_ = pty.Close()
		return nil, nil, err
	}
	return pty, tty, nil
}
