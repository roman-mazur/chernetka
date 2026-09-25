//go:build unix

package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

// replaceWithImageViewer replaces the current process with che-img showing the file at path.
// beforeExec is called once che-img is found, to restore the terminal state.
// It returns only if che-img cannot be found.
func replaceWithImageViewer(path string, beforeExec func()) error {
	bin, err := exec.LookPath(cheImgCommand())
	if err != nil {
		return err
	}
	beforeExec()
	err = syscall.Exec(bin, []string{"che-img", path}, os.Environ())
	// The terminal is restored already: the editor cannot continue.
	log.Fatalf("cannot start che-img: %s", err)
	return err
}
