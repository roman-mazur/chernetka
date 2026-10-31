package exeproxy

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
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
	// Creates the file passed in the args, prints INTERRUPTED on SIGINT and exits.
	helperTrap = "exeproxy-helper-trap"
	helperExit = "exeproxy-helper-exit" // Exits with status 3.
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

		case helperExit:
			os.Exit(3)

		case helperTrap:
			interrupts := make(chan os.Signal, 1)
			signal.Notify(interrupts, os.Interrupt)
			if err := os.WriteFile(os.Args[2], nil, 0o600); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			<-interrupts
			fmt.Println("INTERRUPTED")
			os.Exit(0)

		}
	}
	os.Exit(m.Run())
}

type testClient struct {
	t        *testing.T
	home     string
	sender   *remotectl.Sender
	output   bytes.Buffer
	pWatcher <-chan (<-chan struct{})
}

func startProxy(t *testing.T, configure func(cp *CommandProxy)) *testClient {
	// Unix socket paths are limited in length, so the directory is not based on t.TempDir.
	home, err := os.MkdirTemp("", "proxytest")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(home) })
	// os.UserHomeDir reads USERPROFILE on Windows.
	t.Setenv("HOME", home)
	t.Setenv("USERPROFILE", home)

	tc := new(testClient)
	cp := &CommandProxy{
		Logf:     logger.Prefix(t.Logf, "test proxy:"),
		Stdout:   &tc.output,
		Stderr:   &tc.output,
		pWatcher: make(chan (<-chan struct{})),
	}
	if configure != nil {
		configure(cp)
	}

	tstEndpoint := Endpoint("tst")

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

	tc.sender = client
	tc.t = t
	tc.home = home
	tc.pWatcher = cp.pWatcher
	return tc
}

func (tc *testClient) send(cmd ...string) {
	tc.t.Helper()
	tc.t.Log("send", cmd)
	err := tc.sender.Send(&remotectl.CommandData{Action: cmd[0], Args: cmd[1:]})
	if err != nil {
		tc.t.Error("err", err)
	}
}

// waitFinished waits for n commands to finish and returns their output.
func (tc *testClient) waitFinished(n int) string {
	tc.t.Helper()
	timeout := time.After(10 * time.Second)
	for range n {
		select {
		case <-(<-tc.pWatcher):
		case <-timeout:
			tc.t.Fatal("timeout waiting for the processes to finish")
		}
	}
	out := tc.output.String()
	tc.output.Reset()
	return out
}

// waitFile waits for the file at path to appear.
func waitFile(t *testing.T, path string) {
	t.Helper()
	start := time.Now()
	for _, err := os.Stat(path); err != nil; _, err = os.Stat(path) {
		if time.Since(start) > 10*time.Second {
			t.Errorf("timeout waiting for file to appear: %s", path)
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func TestCommandProxy(t *testing.T) {
	tp := startProxy(t, nil)
	send := tp.send
	expectOut := func(s string, n int) {
		t.Helper()
		if out := tp.waitFinished(n); out != s {
			t.Errorf("output = %q, want %q", out, s)
		}
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
	spawned := filepath.Join(tp.home, "spawned")
	send("exe", os.Args[0], helperSpawn, spawned)
	waitFile(t, spawned)
	// Until the child sharing stdout is gone, the output is not finished.
	send("exe", os.Args[0], helperEcho, "NO CHILDREN")
	expectOut("NO CHILDREN\n", 2)

	if runtime.GOOS != "windows" {
		t.Log("check replace interrupts")
		trapped := filepath.Join(tp.home, "trapped")
		send("exe", os.Args[0], helperTrap, trapped)
		waitFile(t, trapped)
		send("exe", os.Args[0], helperEcho, "AFTER INTERRUPT")
		expectOut("INTERRUPTED\nAFTER INTERRUPT\n", 2)
	}

	t.Log("check shell")
	dir := t.TempDir()
	tp.sender.Send(&remotectl.CommandData{
		Action:  "exe",
		WorkDir: dir,
		Args:    []string{"bash", "-c", "echo SHELL > out.txt"},
	})
	expectOut("", 1)
	if out, err := os.ReadFile(filepath.Join(dir, "out.txt")); err != nil || strings.TrimSpace(string(out)) != "SHELL" {
		t.Errorf("the shell command output in its dir: %q, %v", out, err)
	}
}

func TestCommandProxy_Header(t *testing.T) {
	tp := startProxy(t, func(cp *CommandProxy) {
		cp.Header = true
		cp.Clear = true
		cp.Grace = 100 * time.Millisecond
	})

	tp.send("exe", os.Args[0], helperEcho, "it's OK")
	out := tp.waitFinished(1)
	t.Log(out)
	for _, want := range []string{"\x1b[2J", "it's OK\n", "✓ done in"} {
		if !strings.Contains(out, want) {
			t.Errorf("output %q misses %q", out, want)
		}
	}

	tp.send("exe", os.Args[0], helperBlock)
	tp.send("exe", os.Args[0], helperExit)
	out = tp.waitFinished(2)
	t.Log(out)
	for _, want := range []string{"✗ exit status 3 after"} {
		if !strings.Contains(out, want) {
			t.Errorf("output %q misses %q", out, want)
		}
	}

	tp.send("exe", filepath.Join(tp.home, "missing"))
	tp.send("exe", os.Args[0], helperEcho)
	t.Log(out)
	if out := tp.waitFinished(1); !strings.Contains(out, "cannot start") {
		t.Errorf("output %q misses the start error", out)
	}
}
