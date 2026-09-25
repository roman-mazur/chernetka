package main

import (
	"errors"
	"sync"
	"testing"
	"time"

	"rmazur.io/chernetka/internal/cheimg"
)

// fakeCheImg accepts items once it's launched and has started.
type fakeCheImg struct {
	mu        sync.Mutex
	launches  int
	launchErr error
	startIn   int  // number of failed sends after the launch
	launched  bool // launched without an error
	running   bool
	shown     []cheimg.Item
}

var errNotRunning = errors.New("not running")

func (f *fakeCheImg) send(d cheimg.Item) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if !f.running {
		if !f.launched || f.startIn > 0 {
			f.startIn--
			return errNotRunning
		}
		f.running = true
	}
	f.shown = append(f.shown, d)
	return nil
}

func (f *fakeCheImg) launch() error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.launches++
	f.launched = f.launchErr == nil
	return f.launchErr
}

func (f *fakeCheImg) viewer(t *testing.T) *imageViewer {
	return &imageViewer{
		logf:          t.Logf,
		send:          f.send,
		launch:        f.launch,
		retryInterval: time.Millisecond,
		launchTimeout: 100 * time.Millisecond,
	}
}

func TestImageViewer(t *testing.T) {
	a, b := cheimg.Item{Path: "a.d2"}, cheimg.Item{Source: "b"}

	for _, tc := range []struct {
		name         string
		cheImg       fakeCheImg
		wantLaunches int
		wantShown    int
	}{
		{name: "running", cheImg: fakeCheImg{running: true}, wantShown: 2},
		{name: "launched", wantLaunches: 1, wantShown: 2},
		{name: "launched, slow start", cheImg: fakeCheImg{startIn: 5}, wantLaunches: 1, wantShown: 2},
		{name: "launch error", cheImg: fakeCheImg{launchErr: errors.New("no terminal")}, wantLaunches: 2},
		// The first item times out, the second finds che-img running.
		{name: "start timeout", cheImg: fakeCheImg{startIn: 1000}, wantLaunches: 2},
	} {
		t.Run(tc.name, func(t *testing.T) {
			v := tc.cheImg.viewer(t)
			v.deliver(a)
			v.deliver(b)
			if tc.cheImg.launches != tc.wantLaunches {
				t.Errorf("launched %d times, want %d", tc.cheImg.launches, tc.wantLaunches)
			}
			if len(tc.cheImg.shown) != tc.wantShown {
				t.Errorf("shown %v, want %d items", tc.cheImg.shown, tc.wantShown)
			}
		})
	}
}

func TestImageViewer_ShowDoesNotBlock(t *testing.T) {
	blocked := make(chan struct{})
	shown := make(chan cheimg.Item, 1)
	v := &imageViewer{
		logf: t.Logf,
		send: func(d cheimg.Item) error {
			<-blocked
			shown <- d
			return nil
		},
	}
	v.Show(cheimg.Item{Source: "a"})
	close(blocked)
	select {
	case <-shown:
	case <-time.After(time.Second):
		t.Fatal("item not delivered")
	}
}
