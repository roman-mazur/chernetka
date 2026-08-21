//go:build windows

package main

func windowChangeSignal() <-chan struct{} {
	return nil // TODO: Implement window change event on Windows.
}
