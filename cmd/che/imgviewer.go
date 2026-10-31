package main

import (
	"os"
	"path/filepath"
	"time"

	"rmazur.io/chernetka/internal/cheimg"
	"rmazur.io/chernetka/internal/logger"
)

// imageViewer shows images and diagrams in a che-img process, launching a new one if none is running.
type imageViewer struct {
	paneProcess[cheimg.Item]
}

func newImageViewer(logf logger.Func) *imageViewer {
	return &imageViewer{paneProcess[cheimg.Item]{
		logf:          logger.Prefix(logf, "imgviewer: "),
		send:          cheimg.Show,
		launch:        launchImageViewer,
		retryInterval: 100 * time.Millisecond,
		launchTimeout: 10 * time.Second,
	}}
}

// Show implements extd2.Viewer.
func (v *imageViewer) Show(it cheimg.Item) {
	go v.deliver(it)
}

// cheImgCommand returns the che-img command, preferring the one installed next to the editor.
func cheImgCommand() string {
	if exe, err := os.Executable(); err == nil {
		p := filepath.Join(filepath.Dir(exe), "che-img")
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}
	return "che-img"
}
