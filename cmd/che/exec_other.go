//go:build !unix

package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

// replaceWithImageViewer runs che-img showing the file at path and exits with its status:
// the process cannot be replaced on this platform.
// beforeExec is called once che-img is found, to restore the terminal state.
// It returns only if che-img cannot be found.
func replaceWithImageViewer(path string, beforeExec func()) error {
	bin, err := exec.LookPath(cheImgCommand())
	if err != nil {
		return err
	}
	beforeExec()
	cmd := exec.Command(bin, path)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	err = cmd.Run()
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		os.Exit(exitErr.ExitCode())
	}
	if err != nil {
		// The terminal is restored already: the editor cannot continue.
		log.Fatalf("cannot start che-img: %s", err)
	}
	os.Exit(0)
	return nil
}
