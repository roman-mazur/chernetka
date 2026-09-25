//go:build windows

package vt

func windowChangeSignal() <-chan struct{} {
	return nil // TODO: Implement window change event on Windows.
}
