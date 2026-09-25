//go:build !unix

package vt

import "golang.org/x/term"

// windowSize reports the size in cells only: the pixel size is not available.
func windowSize(fd int) (WindowSize, error) {
	w, h, err := term.GetSize(fd)
	return WindowSize{Cols: w, Rows: h}, err
}
