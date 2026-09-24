package remotectl

import (
	"errors"
	"net"
	"os"
	"slices"
	"testing"
	"time"
)

// useTempSocketDir points the endpoints to a new directory for the duration of the test.
// Unix socket paths are limited in length, so the directory is not based on t.TempDir.
func useTempSocketDir(t *testing.T) {
	t.Helper()
	dir, err := os.MkdirTemp(t.TempDir(), "rctl")
	if err != nil {
		t.Fatal("failed to create a tmp dir for sockets:", err)
	}

	prev := socketDir
	socketDir = func() (string, error) { return dir, nil }
	t.Cleanup(func() { socketDir = prev })
}

type chanExecutor chan CommandData

func (ce chanExecutor) ExecuteCommand(cmd CommandData) { ce <- cmd }

func startServer(t *testing.T, ep Endpoint) chanExecutor {
	t.Helper()
	srv, err := NewServer(ep)
	if err != nil {
		t.Fatal("NewServer:", err)
	}
	commands := make(chanExecutor, 1)
	done := make(chan struct{})
	go func() {
		defer close(done)
		srv.Run(commands, t.Logf)
	}()
	t.Cleanup(func() {
		_ = srv.Close()
		<-done
	})
	return commands
}

func receiveCommand(t *testing.T, commands chanExecutor) CommandData {
	t.Helper()
	select {
	case cmd := <-commands:
		return cmd
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for a command")
		return CommandData{}
	}
}

func TestSendCommand(t *testing.T) {
	useTempSocketDir(t)
	editorCommands := startServer(t, EditorEndpoint)
	otherCommands := startServer(t, "other")

	sent := CommandData{Action: "show", Args: []string{"multi\nline", ""}}
	if err := SendCommand("other", &sent); err != nil {
		t.Fatal("SendCommand:", err)
	}
	got := receiveCommand(t, otherCommands)
	if got.Action != sent.Action || !slices.Equal(got.Args, sent.Args) {
		t.Errorf("received %+v, want %+v", got, sent)
	}

	if err := SendCommand(EditorEndpoint, &CommandData{Action: "open"}); err != nil {
		t.Fatal("SendCommand:", err)
	}
	if got := receiveCommand(t, editorCommands); got.Action != "open" {
		t.Errorf("editor received %+v", got)
	}
}

func TestSendCommand_NoServer(t *testing.T) {
	useTempSocketDir(t)
	err := SendCommand("missing", &CommandData{Action: "open"})
	if !errors.Is(err, os.ErrNotExist) {
		t.Errorf("SendCommand without a server: %v, want os.ErrNotExist", err)
	}
}

func TestNewServer_StaleSocket(t *testing.T) {
	useTempSocketDir(t)
	const ep = Endpoint("stale")
	p, err := ep.path()
	if err != nil {
		t.Fatal(err)
	}

	// Simulate a crashed process: the socket file remains, nobody listens.
	l, err := net.Listen("unix", p)
	if err != nil {
		t.Fatal(err)
	}
	l.(*net.UnixListener).SetUnlinkOnClose(false)
	_ = l.Close()

	commands := startServer(t, ep)
	if err := SendCommand(ep, &CommandData{Action: "ping"}); err != nil {
		t.Fatal("SendCommand:", err)
	}
	receiveCommand(t, commands)
}

func TestNewServer_ActiveSocket(t *testing.T) {
	useTempSocketDir(t)
	startServer(t, "active")
	if srv, err := NewServer("active"); err == nil {
		_ = srv.Close()
		t.Error("second server started on the endpoint in use")
	}
}
