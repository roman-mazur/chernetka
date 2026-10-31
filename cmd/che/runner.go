package main

import (
	"log"
	"os"
	"time"

	"golang.org/x/term"

	"rmazur.io/chernetka/internal/logger"
	"rmazur.io/chernetka/internal/remotectl"
	"rmazur.io/chernetka/internal/remotectl/exeproxy"
)

// paneRunner runs the commands of the Go line actions with a command proxy in a terminal pane below,
// launching one if none is running. A new command replaces the one still running.
type paneRunner struct {
	paneProcess[remotectl.CommandData]
}

func newPaneRunner(purpose string, logf logger.Func) *paneRunner {
	return &paneRunner{paneProcess[remotectl.CommandData]{
		logf: logger.Prefix(logf, "runner: "),
		send: func(cmd remotectl.CommandData) error {
			sender, err := remotectl.Client(exeproxy.Endpoint(purpose).InTab())
			if err != nil {
				return err
			}
			defer sender.Close()
			return sender.Send(&cmd)
		},
		launch:        func() error { return launchCommandProxy("down", purpose) },
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
		Clear:  term.IsTerminal(int(os.Stdout.Fd())),
	}
	if err := cp.Run(exeproxy.Endpoint(purpose).InTab()); err != nil {
		log.Fatal("cannot start the command proxy: ", err)
	}
}
