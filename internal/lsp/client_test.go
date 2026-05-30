package lsp

import (
	"context"
	"encoding/json"
	"io"
	"testing"
	"time"

	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/pkg/fakenet"
	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

// TestClient_RoundTrip stands up an in-memory JSON-RPC peer that plays the
// role of a language server, then drives the Client through its full lifecycle.
func TestClient_RoundTrip(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	clientConn, serverConn := pipeConns(t)

	// Bring up the fake server: log incoming methods and reply with canned data.
	var (
		initialized = make(chan struct{})
		opened      = make(chan string, 1)
		changed     = make(chan int32, 1)
	)
	serverConn.Go(ctx, jsonrpc2.ReplyHandler(func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
		switch req.Method() {
		case protocol.MethodInitialize:
			return reply(ctx, &protocol.InitializeResult{}, nil)
		case protocol.MethodInitialized:
			close(initialized)
			return reply(ctx, nil, nil)
		case protocol.MethodTextDocumentDidOpen:
			var p protocol.DidOpenTextDocumentParams
			_ = jsonrpc2DecodeParams(req, &p)
			opened <- p.TextDocument.Text
			return reply(ctx, nil, nil)
		case protocol.MethodTextDocumentDidChange:
			var p protocol.DidChangeTextDocumentParams
			_ = jsonrpc2DecodeParams(req, &p)
			changed <- p.TextDocument.Version
			return reply(ctx, nil, nil)
		case protocol.MethodTextDocumentCompletion:
			return reply(ctx, &protocol.CompletionList{
				Items: []protocol.CompletionItem{{Label: "Println", InsertText: "Println"}},
			}, nil)
		case protocol.MethodShutdown:
			return reply(ctx, nil, nil)
		case protocol.MethodExit:
			return reply(ctx, nil, nil)
		}
		return reply(ctx, nil, nil)
	}))

	c, err := newClient(ctx, clientConn, "/tmp")
	if err != nil {
		t.Fatalf("newClient: %v", err)
	}

	select {
	case <-initialized:
	case <-ctx.Done():
		t.Fatalf("server never saw initialized notification: %v", ctx.Err())
	}

	fileURI := uri.File("/tmp/x.go")
	if err := c.DidOpen(ctx, fileURI, "go", "package main\n", 1); err != nil {
		t.Fatalf("DidOpen: %v", err)
	}
	if got := <-opened; got != "package main\n" {
		t.Errorf("server received text %q, want %q", got, "package main\n")
	}

	if err := c.DidChange(ctx, fileURI, "package main\n\nfunc f() {}\n", 2); err != nil {
		t.Fatalf("DidChange: %v", err)
	}
	if got := <-changed; got != 2 {
		t.Errorf("server received version %d, want 2", got)
	}

	items, err := c.Completion(ctx, fileURI, 0, 0)
	if err != nil {
		t.Fatalf("Completion: %v", err)
	}
	if len(items) != 1 || items[0].Label != "Println" {
		t.Fatalf("got items=%v, want one Println", items)
	}

	if err := c.Shutdown(ctx); err != nil {
		t.Errorf("Shutdown: %v", err)
	}
}

// jsonrpc2DecodeParams unmarshals a request's params into v.
func jsonrpc2DecodeParams(req jsonrpc2.Request, v any) error {
	return json.Unmarshal(req.Params(), v)
}

// pipeConns wires two jsonrpc2.Conns together via in-memory pipes.
func pipeConns(t *testing.T) (clientConn, serverConn jsonrpc2.Conn) {
	t.Helper()
	cToSRead, cToSWrite := io.Pipe()
	sToCRead, sToCWrite := io.Pipe()
	clientConn = jsonrpc2.NewConn(jsonrpc2.NewStream(fakenet.NewConn("client", sToCRead, cToSWrite)))
	serverConn = jsonrpc2.NewConn(jsonrpc2.NewStream(fakenet.NewConn("server", cToSRead, sToCWrite)))
	t.Cleanup(func() {
		_ = clientConn.Close()
		_ = serverConn.Close()
	})
	return
}
