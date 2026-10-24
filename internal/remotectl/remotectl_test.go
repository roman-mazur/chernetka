package remotectl

import (
	"errors"
	"fmt"
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
	dir, err := os.MkdirTemp("", "rctl")
	if err != nil {
		t.Fatal("failed to create a tmp dir for sockets:", err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(dir) })

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
	commands := make(chanExecutor, 10)
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

// connect creates a Client for the endpoint, closing it at the test end.
func connect(t *testing.T, ep Endpoint) *Sender {
	t.Helper()
	s, err := Client(ep)
	if err != nil {
		t.Fatal("Client:", err)
	}
	t.Cleanup(func() { _ = s.Close() })
	return s
}

func send(t *testing.T, s *Sender, cmd CommandData) {
	t.Helper()
	if err := s.Send(&cmd); err != nil {
		t.Fatal("Send:", err)
	}
}

func TestSender(t *testing.T) {
	useTempSocketDir(t)
	editorCommands := startServer(t, EditorEndpoint)
	otherCommands := startServer(t, "other")
	other := connect(t, "other")
	editor := connect(t, EditorEndpoint)

	sent := CommandData{Action: "show", Args: []string{"multi\nline", ""}}
	send(t, other, sent)
	got := receiveCommand(t, otherCommands)
	if got.Action != sent.Action || !slices.Equal(got.Args, sent.Args) {
		t.Errorf("received %+v, want %+v", got, sent)
	}

	send(t, editor, CommandData{Action: "open"})
	if got := receiveCommand(t, editorCommands); got.Action != "open" {
		t.Errorf("editor received %+v", got)
	}
}

func TestSender_Order(t *testing.T) {
	useTempSocketDir(t)
	commands := startServer(t, "order")
	s := connect(t, "order")

	actions := []string{"first", "second", "third", "fourth"}
	for _, a := range actions {
		send(t, s, CommandData{Action: a})
	}
	for _, a := range actions {
		if got := receiveCommand(t, commands); got.Action != a {
			t.Errorf("received %q, want %q", got.Action, a)
		}
	}
}

func TestServer_IgnoresPing(t *testing.T) {
	useTempSocketDir(t)
	commands := startServer(t, "ping")
	s := connect(t, "ping")

	send(t, s, CommandData{Action: "ping"})
	send(t, s, CommandData{Action: "open"})
	if got := receiveCommand(t, commands); got.Action != "open" {
		t.Errorf("received %+v, want only open", got)
	}
	select {
	case cmd := <-commands:
		t.Errorf("unexpected command %+v", cmd)
	default:
	}
}

func TestSender_Concurrent(t *testing.T) {
	useTempSocketDir(t)
	commands := startServer(t, "concurrent")
	s := connect(t, "concurrent")

	const n = 5
	errs := make(chan error, n)
	for i := range n {
		go func() { errs <- s.Send(&CommandData{Action: fmt.Sprint(i)}) }()
	}
	var received []string
	for range n {
		if err := <-errs; err != nil {
			t.Error("Send:", err)
		}
		received = append(received, receiveCommand(t, commands).Action)
	}
	slices.Sort(received)
	if want := []string{"0", "1", "2", "3", "4"}; !slices.Equal(received, want) {
		t.Errorf("received %q, want %q", received, want)
	}
}

func TestClient_NoServer(t *testing.T) {
	useTempSocketDir(t)
	if s, err := Client("missing"); !errors.Is(err, os.ErrNotExist) {
		if err == nil {
			_ = s.Close()
		}
		t.Errorf("Client without a server: %v, want os.ErrNotExist", err)
	}
}

func TestServer_CloseWithConnectedClient(t *testing.T) {
	useTempSocketDir(t)
	srv, err := NewServer("closing")
	if err != nil {
		t.Fatal("NewServer:", err)
	}
	commands := make(chanExecutor, 10)
	done := make(chan struct{})
	go func() {
		defer close(done)
		srv.Run(commands, t.Logf)
	}()
	s := connect(t, "closing")

	// The client keeps the connection open, it must not prevent the server from stopping.
	_ = srv.Close()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("Run did not return after Close")
	}
	if err := s.Send(&CommandData{Action: "open"}); err == nil {
		t.Error("Send succeeded after the server was closed")
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
	send(t, connect(t, ep), CommandData{Action: "open"})
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

func TestEndpoint_Scoped(t *testing.T) {
	for _, tc := range []struct {
		scope string
		want  Endpoint
	}{
		{"", "mainctl"},
		{"tab-7f98af016230", "mainctl-tab-7f98af016230"},
		{"../a b", "mainctl-___a_b"},
	} {
		if got := EditorEndpoint.Scoped(tc.scope); got != tc.want {
			t.Errorf("Scoped(%q) = %q, want %q", tc.scope, got, tc.want)
		}
	}
}

func TestSender_Scoped(t *testing.T) {
	useTempSocketDir(t)
	tab1 := startServer(t, EditorEndpoint.Scoped("tab1"))
	tab2 := startServer(t, EditorEndpoint.Scoped("tab2"))

	s := connect(t, EditorEndpoint.Scoped("tab2"))
	send(t, s, CommandData{Action: "open"})
	receiveCommand(t, tab2)
	select {
	case cmd := <-tab1:
		t.Errorf("the other tab received %+v", cmd)
	default:
	}
	if s, err := Client(EditorEndpoint.Scoped("tab3")); !errors.Is(err, os.ErrNotExist) {
		if err == nil {
			_ = s.Close()
		}
		t.Errorf("Client for a tab without a server: %v, want os.ErrNotExist", err)
	}
}
