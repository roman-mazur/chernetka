package main

import (
	"log"
	"time"

	"rmazur.io/chernetka/internal/logger"
	"rmazur.io/chernetka/internal/remotectl"
	"rmazur.io/chernetka/internal/remotectl/exeproxy"
)

// The purposes of the command proxies che launches: each one has its own pane.
const (
	goRunPurpose   = "go-run"   // the Go tests and programs of the line actions
	gitDiffPurpose = "git-diff" // the changes of a file shown with Ctrl+D
)

// paneRunner runs the external commands, like the Go tests or git diff, with the command proxy
// for its purpose in a terminal pane, launching one if none is running. The pane is split
// from the editor's one in the direction (see launchCommandProxy). A new command replaces the one
// still running in the pane.
type paneRunner struct {
	paneProcess[remotectl.CommandData]
}

func newPaneRunner(purpose, direction string, logf logger.Func) *paneRunner {
	return &paneRunner{paneProcess[remotectl.CommandData]{
		logf: logger.Prefix(logf, "runner "+purpose+": "),
		send: func(cmd remotectl.CommandData) error {
			sender, err := remotectl.Client(exeproxy.Endpoint(purpose).InTab())
			if err != nil {
				return err
			}
			defer sender.Close()
			return sender.Send(&cmd)
		},
		launch:        func() error { return launchCommandProxy(direction, purpose) },
		retryInterval: 100 * time.Millisecond,
		launchTimeout: 10 * time.Second,
	}}
}

func (pr *paneRunner) ExecuteCommand(cmd remotectl.CommandData) {
	go pr.deliver(cmd)
}

// runCommandProxy serves the command proxy for the purpose in the current terminal pane until it's killed.
// Each command clears the pane and is shown with its exit status.
func runCommandProxy(purpose string) {
	logf, _ := logger.UserLogFile()
	cp := exeproxy.CommandProxy{
		Logf:   logger.Prefix(logf, "exeproxy "+purpose+": "),
		Header: true,
		Clear:  true,
	}
	if err := cp.Run(exeproxy.Endpoint(purpose).InTab()); err != nil {
		log.Fatal("cannot start the command proxy: ", err)
	}
}
