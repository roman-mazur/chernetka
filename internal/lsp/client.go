// Package lsp is a minimal LSP client used by the editor to obtain completion
// suggestions from a language server (gopls).
//
// It exposes the small subset of LSP methods the editor needs: initialize,
// didOpen, didChange, completion, and shutdown. The transport is JSON-RPC over
// the server's stdio.
package lsp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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

// Start launches gopls for the workspace in rootDir (typically the directory
// containing go.mod) and performs the LSP initialize handshake.
//
// gopls runs with the Go toolchain the workspace requires (see FindToolchain),
// and is built with it if the gopls in PATH is too old (see Toolchain.Gopls).
// If the toolchain can't be determined, gopls from PATH is used as is.
//
// gopls is invoked with -remote=auto, which makes it a thin forwarder to a
// shared daemon: the first editor session spawns the daemon and subsequent
// sessions reuse it, so a single gopls process serves the whole machine. The
// daemon is shared only by the sessions using the same gopls binary and Go
// toolchain: it inherits the environment of the session that spawns it.
//
// Server settings are per session, so they apply even with a shared daemon.
func Start(ctx context.Context, rootDir string, opts Options) (*Client, error) {
	logf := opts.Logf
	if logf == nil {
		logf = func(string, ...any) {}
	}
	opts.Args = []string{"-remote=auto", "serve"}

	tc, err := FindToolchain(ctx, rootDir)
	if err != nil {
		logf("cannot determine the Go toolchain, using gopls from PATH: %s", err)
		return StartWith(ctx, rootDir, opts)
	}
	logf("using %s selected by %s", tc.Version, tc.Source())

	opts.Env = tc.Env(os.Environ())
	opts.Args = []string{"-remote=auto;" + tc.Version, "serve"}
	if opts.GoplsCache != "" {
		gopls, err := tc.Gopls(ctx, opts.GoplsCache, logf)
		if err != nil {
			logf("using gopls from PATH: %s", err)
		} else {
			opts.Command = gopls
		}
	}
	return StartWith(ctx, rootDir, opts)
}

// Options tune how the language server is launched and initialised.
type Options struct {
	Command string   // gopls binary; "gopls" from PATH when empty
	Args    []string // gopls arguments, ignored by Start
	Env     []string // gopls environment; the current one when nil

	// GoplsCache is where Start installs gopls built with the Go toolchains
	// that the gopls in PATH is too old for. If empty, the gopls in PATH is
	// used with any toolchain.
	GoplsCache string

	InitializationOptions map[string]any // gopls settings, see gopls/doc/settings.md
	SnippetSupport        bool

	// Logf receives errors the server reports, like a failed workspace load,
	// and how Start chooses the gopls to run.
	Logf func(format string, args ...any)
}

// StartWith is like Start but runs opts.Command with opts.Args ("serve" when
// empty) and opts.Env as is.
func StartWith(ctx context.Context, rootDir string, opts Options) (*Client, error) {
	args := opts.Args
	if len(args) == 0 {
		args = []string{"serve"}
	}
	command := opts.Command
	if command == "" {
		command = "gopls"
	}
	cmd := exec.Command(command, args...)
	cmd.Env = opts.Env
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

	c, err := newClient(ctx, newConn(stdout, stdin), rootDir, opts)
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
func newClient(ctx context.Context, conn jsonrpc2.Conn, rootDir string, opts Options) (*Client, error) {
	conn.Go(ctx, jsonrpc2.ReplyHandler(serverHandler(opts.Logf)))

	rootURI := uri.File(rootDir)
	initParams := &protocol.InitializeParams{
		ProcessID: int32(os.Getpid()),
		WorkspaceFolders: []protocol.WorkspaceFolder{{
			URI:  string(rootURI),
			Name: filepath.Base(rootDir),
		}},
		Capabilities: protocol.ClientCapabilities{
			TextDocument: &protocol.TextDocumentClientCapabilities{
				Completion: &protocol.CompletionTextDocumentClientCapabilities{
					CompletionItem: &protocol.CompletionTextDocumentClientCapabilitiesItem{
						SnippetSupport: opts.SnippetSupport,
					},
				},
			},
		},
	}
	if len(opts.InitializationOptions) > 0 {
		initParams.InitializationOptions = opts.InitializationOptions
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

// serverHandler handles messages initiated by the server. Error messages are
// passed to logf; otherwise we don't act on them yet — requests get a
// method-not-found error, notifications are dropped — but every request must
// be replied to (ReplyHandler enforces this).
func serverHandler(logf func(format string, args ...any)) jsonrpc2.Handler {
	return func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
		switch req.Method() {
		case protocol.MethodWindowShowMessage, protocol.MethodWindowLogMessage:
			var msg protocol.ShowMessageParams // same shape as LogMessageParams
			if logf != nil && json.Unmarshal(req.Params(), &msg) == nil && msg.Type == protocol.MessageTypeError {
				logf("server error: %s", strings.TrimSpace(msg.Message))
			}
		}
		if _, isCall := req.(*jsonrpc2.Call); isCall {
			return reply(ctx, nil, fmt.Errorf("%q: %w", req.Method(), jsonrpc2.ErrMethodNotFound))
		}
		return reply(ctx, nil, nil)
	}
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

// TextChange is a change of a document content: Text replaces Range, or the
// whole document if Range is nil.
type TextChange struct {
	Range *protocol.Range `json:"range,omitempty"`
	Text  string          `json:"text"`
}

// DidChange sends changes of a document content, which are applied in order.
func (c *Client) DidChange(ctx context.Context, fileURI uri.URI, version int32, changes ...TextChange) error {
	return c.conn.Notify(ctx, protocol.MethodTextDocumentDidChange, &didChangeParams{
		TextDocument: protocol.VersionedTextDocumentIdentifier{
			TextDocumentIdentifier: protocol.TextDocumentIdentifier{URI: fileURI},
			Version:                version,
		},
		ContentChanges: changes,
	})
}

// didChangeParams mirrors protocol.DidChangeTextDocumentParams with changes
// that can omit the range. protocol.TextDocumentContentChangeEvent always
// encodes it, and a server that negotiated incremental sync (like gopls)
// treats an empty range as an insertion at the start of the document instead
// of a replacement of the whole content.
type didChangeParams struct {
	TextDocument   protocol.VersionedTextDocumentIdentifier `json:"textDocument"`
	ContentChanges []TextChange                             `json:"contentChanges"`
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

// OrganizeImports returns the edits of the "organize imports" source action:
// missing imports are added, unused ones are removed, and the import block is
// sorted. The edits are for the document version the server has.
func (c *Client) OrganizeImports(ctx context.Context, fileURI uri.URI) ([]protocol.TextEdit, error) {
	var actions []json.RawMessage // commands or code actions
	_, err := c.conn.Call(ctx, protocol.MethodTextDocumentCodeAction, &protocol.CodeActionParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: fileURI},
		Context: protocol.CodeActionContext{
			Diagnostics: []protocol.Diagnostic{},
			Only:        []protocol.CodeActionKind{protocol.SourceOrganizeImports},
		},
	}, &actions)
	if err != nil {
		return nil, err
	}
	for _, raw := range actions {
		var action protocol.CodeAction
		if json.Unmarshal(raw, &action) != nil || action.Kind != protocol.SourceOrganizeImports || action.Edit == nil {
			continue
		}
		var edits []protocol.TextEdit
		for _, dc := range action.Edit.DocumentChanges {
			if dc.TextDocument.URI == fileURI {
				edits = append(edits, dc.Edits...)
			}
		}
		edits = append(edits, action.Edit.Changes[fileURI]...)
		return edits, nil
	}
	return nil, nil
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
