// Package lsp is a minimal LSP client used by the editor to obtain completion
// suggestions, definitions, formatting, imports, and diagnostics from a language
// server, like gopls.
//
// It exposes the small subset of LSP methods the editor needs: initialize,
// didOpen, didChange, completion, signature help, definition, formatting,
// organize imports, and shutdown. The diagnostics the server publishes are
// passed to a callback. The transport is JSON-RPC over the server's stdio.
package lsp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/pkg/fakenet"
	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

// Client owns a single LSP session over JSON-RPC. It is safe to call its
// methods from multiple goroutines because jsonrpc2.Conn serialises writes.
type Client struct {
	conn   jsonrpc2.Conn
	exited <-chan struct{} // closed once the server process is gone; nil for a supplied connection (tests)
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
//
// On Windows, every session runs its own gopls (see serveArgs).
func Start(ctx context.Context, rootDir string, opts Options) (*Client, error) {
	logf := opts.Logf
	if logf == nil {
		logf = func(string, ...any) {}
	}
	opts.Args = serveArgs("")

	tc, err := FindToolchain(ctx, rootDir)
	if err != nil {
		logf("cannot determine the Go toolchain, using gopls from PATH: %s", err)
		return StartWith(ctx, rootDir, opts)
	}
	logf("using %s selected by %s", tc.Version, tc.Source())

	opts.Env = tc.Env(os.Environ())
	opts.Args = serveArgs(tc.Version)
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

// serveArgs returns the gopls arguments making it forward to the daemon
// shared by the sessions with the same id.
//
// On Windows, gopls panics given an id, and without one, the sessions of any
// gopls binary and any toolchain share the daemon at a fixed address. So
// there gopls serves the session itself.
func serveArgs(id string) []string {
	if runtime.GOOS == "windows" {
		return []string{"serve"}
	}
	remote := "-remote=auto"
	if id != "" {
		remote += ";" + id
	}
	return []string{remote, "serve"}
}

// Options tune how the language server is launched and initialised.
type Options struct {
	Command string   // server binary; "gopls" from PATH when empty
	Args    []string // server arguments, ignored by Start
	Env     []string // server environment; the current one when nil

	// GoplsCache is where Start installs gopls built with the Go toolchains
	// that the gopls in PATH is too old for. If empty, the gopls in PATH is
	// used with any toolchain.
	GoplsCache string

	InitializationOptions map[string]any // gopls settings, see gopls/doc/settings.md
	SnippetSupport        bool

	// Logf receives errors the server reports, like a failed workspace load,
	// what the server writes to its stderr, and how Start chooses the gopls to
	// run. The server's stderr goes to os.Stderr when it's nil.
	Logf func(format string, args ...any)

	// Diagnostics receives the problems the server finds in a document,
	// replacing the ones it reported for the document before. It's called on
	// the connection goroutine, so it must not block.
	Diagnostics func(protocol.PublishDiagnosticsParams)
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
	var stderr *lineLogger
	if opts.Logf != nil {
		stderr = &lineLogger{logf: opts.Logf, prefix: filepath.Base(command) + " stderr: "}
		cmd.Stderr = stderr
	} else {
		cmd.Stderr = os.Stderr
	}
	// Don't wait for the processes the server leaves behind with its stderr.
	cmd.WaitDelay = time.Second

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("%s stdin pipe: %w", command, err)
	}
	// Unlike the pipe of StdoutPipe, this one isn't closed by Wait, which runs
	// while it's read.
	stdout, serverStdout, err := os.Pipe()
	if err != nil {
		return nil, fmt.Errorf("%s stdout pipe: %w", command, err)
	}
	cmd.Stdout = serverStdout
	err = cmd.Start()
	_ = serverStdout.Close() // The server has its own copy.
	if err != nil {
		_ = stdout.Close()
		return nil, fmt.Errorf("start %s: %w", command, err)
	}

	exited := make(chan struct{})
	var waitErr error
	go func() {
		waitErr = cmd.Wait()
		if stderr != nil {
			stderr.flush()
		}
		close(exited)
	}()

	c, err := newClient(ctx, newConn(stdout, stdin), exited, rootDir, opts)
	if err != nil {
		select {
		case <-exited:
			err = fmt.Errorf("%w: %s exited: %v", err, filepath.Base(command), waitErr)
		default:
			_ = cmd.Process.Kill()
		}
		_ = stdin.Close()
		_ = stdout.Close()
		return nil, err
	}
	c.exited = exited
	return c, nil
}

// lineLogger passes what a process writes to logf line by line.
type lineLogger struct {
	logf   func(format string, args ...any)
	prefix string
	buf    []byte // the incomplete last line
}

func (l *lineLogger) Write(p []byte) (int, error) {
	l.buf = append(l.buf, p...)
	for {
		i := bytes.IndexByte(l.buf, '\n')
		if i < 0 {
			break
		}
		l.log(l.buf[:i])
		l.buf = l.buf[i+1:]
	}
	return len(p), nil
}

// flush logs the last line if it doesn't end with a newline.
func (l *lineLogger) flush() {
	l.log(l.buf)
	l.buf = nil
}

func (l *lineLogger) log(line []byte) {
	if line = bytes.TrimRight(line, "\r"); len(line) > 0 {
		l.logf("%s%s", l.prefix, line)
	}
}

// newConn wraps a stdio pair (server's stdout to our reader; server's stdin to
// our writer) in a jsonrpc2 connection using LSP Content-Length framing.
func newConn(serverOut io.ReadCloser, serverIn io.WriteCloser) jsonrpc2.Conn {
	netConn := fakenet.NewConn("lsp", serverOut, serverIn)
	return jsonrpc2.NewConn(jsonrpc2.NewStream(netConn))
}

// newClient starts the connection's read loop and performs the LSP handshake,
// which is given up once serverExited is closed. It exists separately from
// Start so tests can supply a synthetic conn.
func newClient(ctx context.Context, conn jsonrpc2.Conn, serverExited <-chan struct{}, rootDir string, opts Options) (*Client, error) {
	conn.Go(ctx, jsonrpc2.ReplyHandler(serverHandler(opts.Logf, opts.Diagnostics)))

	// The connection doesn't fail the calls in flight when the server is gone:
	// they would wait for a reply forever.
	initCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		select {
		case <-serverExited:
			cancel()
		case <-initCtx.Done():
		}
	}()

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
				PublishDiagnostics: &protocol.PublishDiagnosticsClientCapabilities{},
			},
		},
	}
	if len(opts.InitializationOptions) > 0 {
		initParams.InitializationOptions = opts.InitializationOptions
	}
	var initResult protocol.InitializeResult
	if _, err := conn.Call(initCtx, protocol.MethodInitialize, initParams, &initResult); err != nil {
		return nil, fmt.Errorf("initialize: %w", err)
	}
	if err := conn.Notify(initCtx, protocol.MethodInitialized, &protocol.InitializedParams{}); err != nil {
		return nil, fmt.Errorf("initialized: %w", err)
	}
	return &Client{conn: conn}, nil
}

// serverHandler handles messages initiated by the server. Error messages are
// passed to logf, and diagnostics to diagnostics; otherwise we don't act on
// them yet — requests get a method-not-found error, notifications are dropped —
// but every request must be replied to (ReplyHandler enforces this).
func serverHandler(logf func(format string, args ...any), diagnostics func(protocol.PublishDiagnosticsParams)) jsonrpc2.Handler {
	return func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
		switch req.Method() {
		case protocol.MethodWindowShowMessage, protocol.MethodWindowLogMessage:
			var msg protocol.ShowMessageParams // same shape as LogMessageParams
			if logf != nil && json.Unmarshal(req.Params(), &msg) == nil && msg.Type == protocol.MessageTypeError {
				logf("server error: %s", strings.TrimSpace(msg.Message))
			}
		case protocol.MethodTextDocumentPublishDiagnostics:
			var params protocol.PublishDiagnosticsParams
			if diagnostics != nil && json.Unmarshal(req.Params(), &params) == nil {
				diagnostics(params)
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

// SignatureHelp returns the signatures of the call around the given
// zero-based position, with the parameter the position is at. It's nil
// outside of calls.
func (c *Client) SignatureHelp(ctx context.Context, fileURI uri.URI, line, character uint32) (*protocol.SignatureHelp, error) {
	var help *protocol.SignatureHelp
	if _, err := c.conn.Call(ctx, protocol.MethodTextDocumentSignatureHelp, &protocol.SignatureHelpParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{URI: fileURI},
			Position:     protocol.Position{Line: line, Character: character},
		},
	}, &help); err != nil {
		return nil, err
	}
	return help, nil
}

// Definition returns the locations where the symbol at the given zero-based
// position is defined. It's empty if the symbol is unknown.
func (c *Client) Definition(ctx context.Context, fileURI uri.URI, line, character uint32) ([]protocol.Location, error) {
	var raw json.RawMessage
	if _, err := c.conn.Call(ctx, protocol.MethodTextDocumentDefinition, &protocol.DefinitionParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: fileURI},
		Position:     protocol.Position{Line: line, Character: character},
	}, &raw); err != nil {
		return nil, err
	}
	return decodeLocations(raw)
}

// decodeLocations decodes a definition result: null, a location, or a list of
// them. Location links aren't expected: the client doesn't declare their support.
func decodeLocations(raw json.RawMessage) ([]protocol.Location, error) {
	raw = bytes.TrimSpace(raw)
	switch {
	case len(raw) == 0 || string(raw) == "null":
		return nil, nil
	case raw[0] == '[':
		var locs []protocol.Location
		err := json.Unmarshal(raw, &locs)
		return locs, err
	default:
		var loc protocol.Location
		if err := json.Unmarshal(raw, &loc); err != nil {
			return nil, err
		}
		return []protocol.Location{loc}, nil
	}
}

// Formatting returns the edits formatting the document with the given options
// (gopls ignores them: it runs gofmt). The edits are for the document version
// the server has.
func (c *Client) Formatting(ctx context.Context, fileURI uri.URI, opts protocol.FormattingOptions) ([]protocol.TextEdit, error) {
	var edits []protocol.TextEdit
	if _, err := c.conn.Call(ctx, protocol.MethodTextDocumentFormatting, &protocol.DocumentFormattingParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: fileURI},
		Options:      opts,
	}, &edits); err != nil {
		return nil, err
	}
	return edits, nil
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
	if c.exited != nil {
		<-c.exited
	}
	return nil
}
