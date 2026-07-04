//go:build !darwin

package main

// openMainEditor does nothing 
func openMainEditor(path string) error { return nil }
