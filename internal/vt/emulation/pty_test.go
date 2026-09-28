//go:build darwin || linux

package emulation

import (
	"io"
	"testing"

	"golang.org/x/sys/unix"
	"golang.org/x/term"
)

func TestOpenPTY(t *testing.T) {
	pty, tty, err := OpenPTY(80, 40)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = pty.Close()
		_ = tty.Close()
	})

	if !term.IsTerminal(int(tty.Fd())) {
		t.Error("tty is not a terminal")
	}
	size, err := unix.IoctlGetWinsize(int(tty.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		t.Fatal(err)
	}
	if size.Col != 80 || size.Row != 40 {
		t.Errorf("size = %dx%d, want 80x40", size.Col, size.Row)
	}

	// The terminal output reaches pty.
	if _, err := io.WriteString(tty, "out"); err != nil {
		t.Fatal(err)
	}
	buf := make([]byte, 3)
	if _, err := io.ReadFull(pty, buf); err != nil {
		t.Fatal(err)
	}
	if string(buf) != "out" {
		t.Errorf("output = %q, want %q", buf, "out")
	}

	// The input written to pty reaches the terminal.
	if _, err := term.MakeRaw(int(tty.Fd())); err != nil {
		t.Fatal(err)
	}
	if _, err := io.WriteString(pty, "in"); err != nil {
		t.Fatal(err)
	}
	buf = buf[:2]
	if _, err := io.ReadFull(tty, buf); err != nil {
		t.Fatal(err)
	}
	if string(buf) != "in" {
		t.Errorf("input = %q, want %q", buf, "in")
	}
}
