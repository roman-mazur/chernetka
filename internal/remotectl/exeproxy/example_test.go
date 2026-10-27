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
	// os.UserHomeDir reads USERPROFILE on Windows.
	os.Setenv("HOME", tmpDir)
	os.Setenv("USERPROFILE", tmpDir)

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
	var client *remotectl.Sender
	waitTime := 50 * time.Millisecond
	for range 10 {
		client, err = remotectl.Client(endpoint)
		if err == nil {
			break
		}
		time.Sleep(waitTime)
		waitTime *= 2
	}
	fmt.Println("send echo")
	cmd := []string{"echo", "OK"}
	if runtime.GOOS == "windows" {
		cmd = []string{"cmd", "/c", "echo", "OK"}
	}
	client.Send(&remotectl.CommandData{Action: "exe", Args: cmd})
	time.Sleep(time.Second / 2) // Give time for echo to finish.
	client.Send(&remotectl.CommandData{Action: "quit"})

	<-done
	// Output:
	// waiting for server
	// send echo
	// OK
}
