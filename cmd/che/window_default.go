//go:build unix

package main

import (
	"os"
	"os/signal"
	"syscall"
)

func windowChangeSignal() <-chan struct{} {
	c := make(chan struct{})
	go func() {
		defer close(c)
		s := make(chan os.Signal, 1)
		signal.Notify(s, syscall.SIGWINCH)
		for range s {
			c <- struct{}{}
		}
	}()
	return c
}
