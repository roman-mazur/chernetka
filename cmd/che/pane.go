package main

import (
	"sync"
	"time"

	"rmazur.io/chernetka/internal/logger"
)

// paneProcess delivers items to a helper process running in a terminal pane, launching a new one if none is running.
type paneProcess[T any] struct {
	logf   logger.Func
	send   func(it T) error
	launch func() error

	retryInterval time.Duration // how often to try sending the item to the launched process
	launchTimeout time.Duration // how long to wait for the launched process to start listening

	mu sync.Mutex // serializes deliveries so that the process is not launched twice
}

func (pp *paneProcess[T]) deliver(it T) {
	pp.mu.Lock()
	defer pp.mu.Unlock()

	err := pp.send(it)
	if err == nil {
		return
	}
	pp.logf("not available (%s), launching", err)
	if err := pp.launch(); err != nil {
		pp.logf("cannot launch: %s", err)
		return
	}

	deadline := time.Now().Add(pp.launchTimeout)
	for time.Now().Before(deadline) {
		time.Sleep(pp.retryInterval)
		if err = pp.send(it); err == nil {
			return
		}
	}
	pp.logf("did not start: %s", err)
}
