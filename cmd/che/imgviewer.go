package main

import (
	"os"
	"path/filepath"
	"sync"
	"time"

	"rmazur.io/chernetka/internal/cheimg"
	"rmazur.io/chernetka/internal/logger"
)

// imageViewer shows images and diagrams in a che-img process, launching a new one if none is running.
type imageViewer struct {
	logf   logger.Func
	send   func(it cheimg.Item) error
	launch func() error

	retryInterval time.Duration // how often to try sending the item to the launched che-img
	launchTimeout time.Duration // how long to wait for the launched che-img to start listening

	mu sync.Mutex // serializes deliveries so that che-img is not launched twice
}

func newImageViewer(logf logger.Func) *imageViewer {
	return &imageViewer{
		logf:          logger.Prefix(logf, "imgviewer: "),
		send:          cheimg.Show,
		launch:        launchImageViewer,
		retryInterval: 100 * time.Millisecond,
		launchTimeout: 10 * time.Second,
	}
}

// Show implements extd2.Viewer.
func (v *imageViewer) Show(it cheimg.Item) {
	go v.deliver(it)
}

func (v *imageViewer) deliver(it cheimg.Item) {
	v.mu.Lock()
	defer v.mu.Unlock()

	err := v.send(it)
	if err == nil {
		return
	}
	v.logf("che-img is not available (%s), launching", err)
	if err := v.launch(); err != nil {
		v.logf("cannot launch che-img: %s", err)
		return
	}

	deadline := time.Now().Add(v.launchTimeout)
	for time.Now().Before(deadline) {
		time.Sleep(v.retryInterval)
		if err = v.send(it); err == nil {
			return
		}
	}
	v.logf("che-img did not start: %s", err)
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
