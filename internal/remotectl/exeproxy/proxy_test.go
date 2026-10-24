package exeproxy

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"rmazur.io/chernetka/internal/logger"
	"rmazur.io/chernetka/internal/remotectl"
)

const (
	helperEcho  = "exeproxy-helper-echo"  // Prints the rest of the args and exits.
	helperBlock = "exeproxy-helper-block" // Prints its pid and runs until killed.
)

// TestMain lets the test binary act as a command executed by the proxy.
func TestMain(m *testing.M) {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case helperEcho:
			fmt.Println(strings.Join(os.Args[2:], " "))
			os.Exit(0)
		case helperBlock:
			time.Sleep(5 * time.Minute)
			os.Exit(1)
		}
	}
	os.Exit(m.Run())
}

func TestCommandProxy(t *testing.T) {
	// Unix socket paths are limited in length, so the directory is not based on t.TempDir.
	home, err := os.MkdirTemp("", "proxytest")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(home) })
	t.Setenv("HOME", home)

	var outputCapture bytes.Buffer
	cp := CommandProxy{
		Logf:     logger.Prefix(t.Logf, "test proxy:"),
		Stdout:   &outputCapture,
		Stderr:   &outputCapture,
		pWatcher: make(chan (<-chan struct{})),
	}

	tstEndpoint := remotectl.Endpoint("cmd-proxy-test")
	proxyFinishErr := make(chan error)
	go func() {
		proxyFinishErr <- cp.Run(tstEndpoint)
		close(proxyFinishErr)
	}()

	t.Cleanup(func() {
		_ = remotectl.SendCommand(tstEndpoint, &remotectl.CommandData{Action: "quit"})
		if err := <-proxyFinishErr; err != nil {
			t.Fatal("proxy finish error:", err)
		}
	})

	t.Log("waiting for the connection")
	var (
		waitTime  = 50 * time.Millisecond
		connError error
	)
	for range 10 {
		connError = remotectl.SendCommand(tstEndpoint, &remotectl.CommandData{Action: "ping"})
		if connError == nil {
			break
		}
		time.Sleep(waitTime)
		waitTime *= 2
	}
	if connError != nil {
		t.Fatal("proxy was not started")
	}

	send := func(cmd ...string) {
		t.Helper()
		t.Log("send", cmd)
		err := remotectl.SendCommand(
			tstEndpoint,
			&remotectl.CommandData{Action: cmd[0], Args: cmd[1:]},
		)
		if err != nil {
			t.Error("err", err)
		}
	}
	expectOut := func(s string, n int) {
		t.Helper()
		for range n {
			<-(<-cp.pWatcher)
		}
		if outputCapture.String() != s {
			t.Errorf("output = %q, want %q", outputCapture.String(), s)
		}
		outputCapture.Reset()
	}

	t.Log("check unknown command")
	send("unknown", "arg1")
	send("unknown2")

	t.Log("check normal cmd")
	send("exe", os.Args[0], helperEcho, "PASSED")
	expectOut("PASSED\n", 1)

	t.Log("check replace")
	send("exe", os.Args[0], helperBlock)
	send("exe", os.Args[0], helperEcho, "REPLACED")
	expectOut("REPLACED\n", 2)
}
