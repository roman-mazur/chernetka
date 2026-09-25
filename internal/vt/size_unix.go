//go:build unix

package vt

import "golang.org/x/sys/unix"

func windowSize(fd int) (WindowSize, error) {
	ws, err := unix.IoctlGetWinsize(fd, unix.TIOCGWINSZ)
	if err != nil {
		return WindowSize{}, err
	}
	return WindowSize{
		Cols:   int(ws.Col),
		Rows:   int(ws.Row),
		XPixel: int(ws.Xpixel),
		YPixel: int(ws.Ypixel),
	}, nil
}
