// Package lsp is a minimal LSP client used by the editor to obtain completion
// suggestions from a language server (gopls).
//
// It exposes the small subset of LSP methods the editor needs: initialize,
// didOpen, didChange, completion, and shutdown. The transport is JSON-RPC over
// the server's stdio.
package lsp

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"

	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/pkg/fakenet"
	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

// Client owns a single LSP session over JSON-RPC. It is safe to call its
// methods from multiple goroutines because jsonrpc2.Conn serialises writes.
type Client struct {
	conn jsonrpc2.Conn
	proc *os.Process // nil if the connection was supplied externally (tests).
}

// Start launches gopls and performs the LSP initialize handshake.
// rootDir is the workspace folder (typically the directory containing go.mod).
//
// gopls is invoked with -remote=auto, which makes it a thin forwarder to a
// shared daemon: the first editor session spawns the daemon and subsequent
// sessions reuse it, so a single gopls process serves the whole machine.
func Start(ctx context.Context, rootDir string) (*Client, error) {
	cmd := exec.Command("gopls", "-remote=auto", "serve")
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("gopls stdin pipe: %w", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("gopls stdout pipe: %w", err)
	}
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("start gopls: %w", err)
	}

	c, err := newClient(ctx, newConn(stdout, stdin), rootDir)
	if err != nil {
		_ = stdin.Close()
		_ = cmd.Process.Kill()
		return nil, err
	}
	c.proc = cmd.Process
	return c, nil
}

// newConn wraps a stdio pair (server's stdout to our reader; server's stdin to
// our writer) in a jsonrpc2 connection using LSP Content-Length framing.
func newConn(serverOut io.ReadCloser, serverIn io.WriteCloser) jsonrpc2.Conn {
	netConn := fakenet.NewConn("lsp", serverOut, serverIn)
	return jsonrpc2.NewConn(jsonrpc2.NewStream(netConn))
}

// newClient starts the connection's read loop and performs the LSP handshake.
// It exists separately from Start so tests can supply a synthetic conn.
func newClient(ctx context.Context, conn jsonrpc2.Conn, rootDir string) (*Client, error) {
	conn.Go(ctx, jsonrpc2.ReplyHandler(serverHandler))

	rootURI := uri.File(rootDir)
	initParams := &protocol.InitializeParams{
		ProcessID: int32(os.Getpid()),
		RootURI:   rootURI,
	}
	var initResult protocol.InitializeResult
	if _, err := conn.Call(ctx, protocol.MethodInitialize, initParams, &initResult); err != nil {
		return nil, fmt.Errorf("initialize: %w", err)
	}
	if err := conn.Notify(ctx, protocol.MethodInitialized, &protocol.InitializedParams{}); err != nil {
		return nil, fmt.Errorf("initialized: %w", err)
	}
	return &Client{conn: conn}, nil
}

// serverHandler handles messages initiated by the server. We don't act on any
// of them yet — requests get a method-not-found error, notifications are
// dropped — but every request must be replied to (ReplyHandler enforces this).
func serverHandler(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
	if _, isCall := req.(*jsonrpc2.Call); isCall {
		return reply(ctx, nil, fmt.Errorf("%q: %w", req.Method(), jsonrpc2.ErrMethodNotFound))
	}
	return reply(ctx, nil, nil)
}

// DidOpen tells the server about a newly opened document.
func (c *Client) DidOpen(ctx context.Context, fileURI uri.URI, languageID, text string, version int32) error {
	return c.conn.Notify(ctx, protocol.MethodTextDocumentDidOpen, &protocol.DidOpenTextDocumentParams{
		TextDocument: protocol.TextDocumentItem{
			URI:        fileURI,
			LanguageID: protocol.LanguageIdentifier(languageID),
			Version:    version,
			Text:       text,
		},
	})
}

// DidChange notifies the server about a full-document content change.
func (c *Client) DidChange(ctx context.Context, fileURI uri.URI, text string, version int32) error {
	return c.conn.Notify(ctx, protocol.MethodTextDocumentDidChange, &protocol.DidChangeTextDocumentParams{
		TextDocument: protocol.VersionedTextDocumentIdentifier{
			TextDocumentIdentifier: protocol.TextDocumentIdentifier{URI: fileURI},
			Version:                version,
		},
		ContentChanges: []protocol.TextDocumentContentChangeEvent{{Text: text}},
	})
}

// Completion requests completion items at the given zero-based position.
func (c *Client) Completion(ctx context.Context, fileURI uri.URI, line, character uint32) ([]protocol.CompletionItem, error) {
	var list protocol.CompletionList
	if _, err := c.conn.Call(ctx, protocol.MethodTextDocumentCompletion, &protocol.CompletionParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: fileURI},
		Position:     protocol.Position{Line: line, Character: character},
	}, &list); err != nil {
		return nil, err
	}
	return list.Items, nil
}

// Shutdown performs a graceful LSP shutdown and terminates the server process.
func (c *Client) Shutdown(ctx context.Context) error {
	_, _ = c.conn.Call(ctx, protocol.MethodShutdown, nil, nil)
	_ = c.conn.Notify(ctx, protocol.MethodExit, nil)
	_ = c.conn.Close()
	<-c.conn.Done()
	if c.proc != nil {
		_, _ = c.proc.Wait()
	}
	return nil
}
