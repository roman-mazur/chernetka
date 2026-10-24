package exeproxy_test

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"rmazur.io/chernetka/internal/remotectl"
	"rmazur.io/chernetka/internal/remotectl/exeproxy"
)

func ExampleCommandProxy() {
	// For socket files.
	tmpDir, err := os.MkdirTemp("", "proxyexample")
	if err != nil {
		panic(err)
	}
	os.Setenv("HOME", tmpDir)

	const endpoint = remotectl.Endpoint("example")

	// Server side.
	var cp exeproxy.CommandProxy
	done := make(chan struct{})
	go func() {
		cp.Run(endpoint)
		close(done)
	}()

	// Client side.
	fmt.Println("waiting for server")
	waitTime := 50 * time.Millisecond
	for range 10 {
		if err := remotectl.SendCommand(endpoint, &remotectl.CommandData{Action: "ping"}); err == nil {
			break
		}
		waitTime *= 2
	}
	fmt.Println("send echo")
	cmd := []string{"echo", "OK"}
	if runtime.GOOS == "windows" {
		cmd = []string{"cmd", "/c", "echo", "OK"}
	}
	remotectl.SendCommand(endpoint, &remotectl.CommandData{Action: "exe", Args: cmd})
	time.Sleep(time.Second / 2) // Give time for echo to finish.
	remotectl.SendCommand(endpoint, &remotectl.CommandData{Action: "quit"})

	<-done
	// Output:
	// waiting for server
	// send echo
	// OK
}
