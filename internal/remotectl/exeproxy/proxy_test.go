package exeproxy

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"rmazur.io/chernetka/internal/logger"
	"rmazur.io/chernetka/internal/remotectl"
)

const (
	helperEcho  = "exeproxy-helper-echo"  // Prints the rest of the args and exits.
	helperBlock = "exeproxy-helper-block" // Runs until killed.
	// Starts helperBlock sharing its stdout, creates the file passed in the args and runs until killed.
	helperSpawn = "exeproxy-helper-spawn"
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
		case helperSpawn:
			child := exec.Command(os.Args[0], helperBlock)
			child.Stdout = os.Stdout
			if err := child.Start(); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			if err := os.WriteFile(os.Args[2], nil, 0o600); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
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

	t.Log("waiting for the connection")
	var (
		client    *remotectl.Sender
		waitTime  = 50 * time.Millisecond
		connError error
	)
	for range 10 {
		client, connError = remotectl.Client(tstEndpoint)
		if connError == nil {
			break
		}
		time.Sleep(waitTime)
		waitTime *= 2
	}
	if connError != nil {
		t.Fatal("proxy was not started")
	}
	t.Cleanup(func() { _ = client.Close() })

	t.Cleanup(func() {
		_ = client.Send(&remotectl.CommandData{Action: "quit"})
		if err := <-proxyFinishErr; err != nil {
			t.Fatal("proxy finish error:", err)
		}
	})

	send := func(cmd ...string) {
		t.Helper()
		t.Log("send", cmd)
		err := client.Send(&remotectl.CommandData{Action: cmd[0], Args: cmd[1:]})
		if err != nil {
			t.Error("err", err)
		}
	}
	expectOut := func(s string, n int) {
		t.Helper()
		timeout := time.After(10 * time.Second)
		for range n {
			select {
			case <-(<-cp.pWatcher):
			case <-timeout:
				t.Fatal("timeout waiting for the processes to finish")
			}
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

	t.Log("check replace kills children")
	spawned := filepath.Join(home, "spawned")
	send("exe", os.Args[0], helperSpawn, spawned)
	for _, err := os.Stat(spawned); err != nil; _, err = os.Stat(spawned) {
		time.Sleep(10 * time.Millisecond)
	}
	// Until the child sharing stdout is gone, the output is not finished.
	send("exe", os.Args[0], helperEcho, "NO CHILDREN")
	expectOut("NO CHILDREN\n", 2)
}
