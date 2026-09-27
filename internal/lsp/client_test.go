package lsp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
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
		changed     = make(chan json.RawMessage, 1)
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
			changed <- req.Params()
			return reply(ctx, nil, nil)
		case protocol.MethodTextDocumentCompletion:
			return reply(ctx, &protocol.CompletionList{
				Items: []protocol.CompletionItem{{Label: "Println", InsertText: "Println"}},
			}, nil)
		case protocol.MethodTextDocumentDefinition:
			var p protocol.DefinitionParams
			if err := jsonrpc2DecodeParams(req, &p); err != nil || p.Position.Line != 2 || p.Position.Character != 5 {
				return reply(ctx, nil, fmt.Errorf("unexpected params %+v (%v)", p, err))
			}
			// A single location, not a list.
			return reply(ctx, &protocol.Location{
				URI:   uri.File("/tmp/y.go"),
				Range: protocol.Range{Start: protocol.Position{Line: 7, Character: 5}, End: protocol.Position{Line: 7, Character: 6}},
			}, nil)
		case protocol.MethodTextDocumentFormatting:
			var p protocol.DocumentFormattingParams
			if err := jsonrpc2DecodeParams(req, &p); err != nil || p.Options.TabSize != 8 {
				return reply(ctx, nil, fmt.Errorf("unexpected options %+v (%v)", p.Options, err))
			}
			return reply(ctx, []protocol.TextEdit{{
				Range:   protocol.Range{Start: protocol.Position{Line: 2, Character: 8}, End: protocol.Position{Line: 2, Character: 8}},
				NewText: " ",
			}}, nil)
		case protocol.MethodShutdown:
			return reply(ctx, nil, nil)
		case protocol.MethodExit:
			return reply(ctx, nil, nil)
		}
		return reply(ctx, nil, nil)
	}))

	c, err := newClient(ctx, clientConn, "/tmp", Options{})
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

	if err := c.DidChange(ctx, fileURI, 2, TextChange{Text: "package main\n\nfunc f() {}\n"}); err != nil {
		t.Fatalf("DidChange: %v", err)
	}
	raw := <-changed
	var p protocol.DidChangeTextDocumentParams
	_ = json.Unmarshal(raw, &p)
	if p.TextDocument.Version != 2 {
		t.Errorf("server received version %d, want 2", p.TextDocument.Version)
	}
	if len(p.ContentChanges) != 1 || p.ContentChanges[0].Text != "package main\n\nfunc f() {}\n" {
		t.Errorf("server received changes %+v", p.ContentChanges)
	}
	// A full content change must not have a range: with incremental sync, an
	// empty range means inserting the text at the start of the document.
	if strings.Contains(string(raw), `"range"`) {
		t.Errorf("full content change sent with a range: %s", raw)
	}

	items, err := c.Completion(ctx, fileURI, 0, 0)
	if err != nil {
		t.Fatalf("Completion: %v", err)
	}
	if len(items) != 1 || items[0].Label != "Println" {
		t.Fatalf("got items=%v, want one Println", items)
	}

	locs, err := c.Definition(ctx, fileURI, 2, 5)
	if err != nil {
		t.Fatalf("Definition: %v", err)
	}
	if len(locs) != 1 || locs[0].URI.Filename() != "/tmp/y.go" || locs[0].Range.Start.Line != 7 {
		t.Errorf("got locations=%+v, want one in /tmp/y.go on line 7", locs)
	}

	edits, err := c.Formatting(ctx, fileURI, protocol.FormattingOptions{TabSize: 8})
	if err != nil {
		t.Fatalf("Formatting: %v", err)
	}
	if len(edits) != 1 || edits[0].NewText != " " || edits[0].Range.Start.Line != 2 {
		t.Errorf("got edits=%+v, want one space inserted on line 2", edits)
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

// TestClient_LogsServerErrors checks that errors the server reports on its own
// (like a failed workspace load) reach the log instead of being dropped.
func TestClient_LogsServerErrors(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	clientConn, serverConn := pipeConns(t)
	serverConn.Go(ctx, jsonrpc2.ReplyHandler(func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
		if req.Method() == protocol.MethodInitialize {
			return reply(ctx, &protocol.InitializeResult{}, nil)
		}
		return reply(ctx, nil, nil)
	}))

	logged := make(chan string, 10)
	logf := func(format string, args ...any) { logged <- fmt.Sprintf(format, args...) }
	if _, err := newClient(ctx, clientConn, "/tmp", Options{Logf: logf}); err != nil {
		t.Fatal(err)
	}

	_ = serverConn.Notify(ctx, protocol.MethodWindowLogMessage, &protocol.LogMessageParams{Type: protocol.MessageTypeInfo, Message: "all good"})
	_ = serverConn.Notify(ctx, protocol.MethodWindowShowMessage, &protocol.ShowMessageParams{Type: protocol.MessageTypeError, Message: "load failed\n"})
	select {
	case got := <-logged:
		if got != "server error: load failed" {
			t.Errorf("logged %q", got)
		}
	case <-ctx.Done():
		t.Fatal("server error not logged")
	}
}

func TestDecodeLocations(t *testing.T) {
	loc := `{"uri":"file:///tmp/x.go","range":{"start":{"line":1,"character":2},"end":{"line":1,"character":3}}}`
	for _, tc := range []struct {
		raw  string
		want int
	}{
		{raw: "null", want: 0},
		{raw: "[]", want: 0},
		{raw: loc, want: 1},
		{raw: "[" + loc + "," + loc + "]", want: 2},
	} {
		locs, err := decodeLocations(json.RawMessage(tc.raw))
		if err != nil {
			t.Errorf("decodeLocations(%s): %s", tc.raw, err)
			continue
		}
		if len(locs) != tc.want {
			t.Errorf("decodeLocations(%s) = %+v, want %d locations", tc.raw, locs, tc.want)
		}
		for _, l := range locs {
			if l.URI.Filename() != "/tmp/x.go" || l.Range.Start.Character != 2 {
				t.Errorf("decodeLocations(%s): bad location %+v", tc.raw, l)
			}
		}
	}
}
