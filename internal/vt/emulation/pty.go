// Package emulation provides pseudo-terminals to run the terminal apps in, like tests do
// to drive che without the user's terminal.
package emulation

import "os"

// OpenPTY opens a new pseudo-terminal of the given size in characters.
// The process running in the terminal is given tty as its standard streams, and it's
// driven by writing its input to pty and reading its output from pty.
// Reading from pty fails once all the tty descriptors are closed, so the caller closes
// its tty after starting the process.
func OpenPTY(cols, rows int) (pty, tty *os.File, err error) {
	pty, tty, err = openPTY()
	if err != nil {
		return nil, nil, err
	}
	if err := setSize(pty, cols, rows); err != nil {
		_ = pty.Close()
		_ = tty.Close()
		return nil, nil, err
	}
	return pty, tty, nil
}
