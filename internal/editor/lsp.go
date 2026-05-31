package editor

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"

	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
	"rmazur.io/x/edit/internal/lsp"
)

// lspClient is the subset of an LSP backend the editor needs. *lsp.Client
// satisfies it; tests substitute a fake.
type lspClient interface {
	DidOpen(ctx context.Context, fileURI uri.URI, languageID, text string, version int32) error
	DidChange(ctx context.Context, fileURI uri.URI, text string, version int32) error
	Completion(ctx context.Context, fileURI uri.URI, line, character uint32) ([]protocol.CompletionItem, error)
	Shutdown(ctx context.Context) error
}

// lspStarter spawns and initialises an LSP client for the given workspace root.
type lspStarter func(ctx context.Context, rootDir string) (lspClient, error)

func defaultLSPStarter(ctx context.Context, rootDir string) (lspClient, error) {
	return lsp.Start(ctx, rootDir)
}

// lspIntegration encapsulates the logic of integrating with an LSP client in an Editor and Buffer.
type lspIntegration struct {
	client     lspClient
	starter    lspStarter
	cancel     context.CancelFunc
	startError error
}

type lspBufferData struct {
	docUri      uri.URI
	suggestions []string

	version       int32
	reqCompletion int
}

// ensureLSP performs the one-time start of the language server. After the
// first call, subsequent calls return the cached start error (if any) without
// retrying — a failed start usually means gopls isn't installed.
func (le *lspIntegration) ensureLSP(fileDir string) error {
	if le.client != nil || le.startError != nil {
		return le.startError
	}
	starter := le.starter
	if starter == nil {
		starter = defaultLSPStarter
	}
	ctx, cancel := context.WithCancel(context.Background())
	client, err := starter(ctx, findGoModRoot(fileDir))
	if err != nil {
		cancel()
		le.startError = err
		return err
	}
	le.client = client
	le.cancel = cancel
	return nil
}

// maybeAttach starts the language server (lazily, on first .go file) and
// sends didOpen for buf. Errors are swallowed so the editor stays usable when
// the language server is unavailable.
func (le *lspIntegration) maybeAttach(buf *Buffer) {
	if !strings.HasSuffix(buf.path, ".go") {
		return
	}
	absPath, err := filepath.Abs(buf.path)
	if err != nil {
		return
	}
	if err := le.ensureLSP(filepath.Dir(absPath)); err != nil {
		return
	}
	buf.lsp = lspBufferData{
		docUri:  uri.File(absPath),
		version: 1,
	}
	_ = le.client.DidOpen(context.Background(), buf.lsp.docUri, "go", buf.text(), buf.lsp.version)
}

func (le *lspIntegration) activeOn(buf *Buffer) bool {
	return le.client != nil && buf.lsp.docUri != ""
}

func (le *lspIntegration) sendChange(buf *Buffer) {
	buf.lsp.version++
	_ = le.client.DidChange(context.Background(), buf.lsp.docUri, buf.text(), buf.lsp.version)
}

func (le *lspIntegration) Close() error {
	if le.client == nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := le.client.Shutdown(ctx)
	if le.cancel != nil {
		le.cancel()
		le.cancel = nil
	}
	le.client = nil
	return err
}

// afterEdit notifies the language server that buf changed and asks for a fresh
// completion at the cursor. The didChange notification is sent synchronously so
// versions reach gopls in order; only the blocking completion round-trip runs
// in a goroutine off a snapshot of buffer state, so the input loop never waits
// on gopls. The result is applied back on the main loop via Post and dropped if
// newer edits have since superseded it. It is a no-op for buffers the server
// doesn't track or when not in insert mode.
func (e *Editor) afterEdit(buf *Buffer) {
	if !e.lsp.activeOn(buf) {
		return
	}
	e.lsp.sendChange(buf) // TODO: consider bouncing and doing asynchronously.

	buf.lsp.reqCompletion++

	var line string
	if lines := buf.content.Lines(); buf.cy < len(lines) {
		line = lines[buf.cy].String()
	}
	cx := min(buf.cx, len(line))

	var (
		client  = e.lsp.client
		fileURI = buf.lsp.docUri
		req     = buf.lsp.reqCompletion
		ln      = uint32(buf.cy)
		ch      = utf16Len(line[:cx])
	)

	// TODO: replace with smarter logic when to ask for a suggestion.
	go func() {
		items, err := client.Completion(context.Background(), fileURI, ln, ch)
		if err != nil {
			return
		}
		suggestions := extractLspSuggestions(items, line, cx)
		e.Post(CommandFunc(func(e *Editor) {
			if buf.lsp.reqCompletion != req {
				return // A newer edit superseded this request.
			}
			buf.lsp.suggestions = suggestions
			e.layoutRequested = true
		}))
	}()
}

func utf16Len(s string) (n uint32) {
	for _, r := range s {
		n += uint32(utf16.RuneLen(r))
	}
	return n
}

func extractLspSuggestions(items []protocol.CompletionItem, line string, cx int) []string {
	if len(items) == 0 {
		return nil
	}
	prefix := identTrailing(line[:min(cx, len(line))])

	res := make([]string, 0, len(items))
	for _, item := range items {
		if item.Label == "" {
			continue
		}
		if prefix != "" && !strings.HasPrefix(item.Label, prefix) {
			// Avoid showing text that contradicts what was typed.
			continue
		}
		if len(prefix) == len(item.Label) {
			continue
		}
		res = append(res, item.Label[len(prefix):])
	}

	if len(res) == 0 {
		return nil
	}
	return res
}

// identTrailing returns the trailing run of identifier characters (letters,
// digits, underscore) of s — the partial word immediately before the cursor.
func identTrailing(s string) string {
	i := len(s)
	for i > 0 {
		r, sz := utf8.DecodeLastRuneInString(s[:i])
		if r != '_' && !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			break
		}
		i -= sz
	}
	return s[i:]
}

// findGoModRoot walks up from dir looking for a go.mod file. Returns the
// directory containing go.mod, or dir if none is found before the filesystem
// root.
func findGoModRoot(dir string) string {
	for d := dir; ; {
		if _, err := os.Stat(filepath.Join(d, "go.mod")); err == nil {
			return d
		}
		parent := filepath.Dir(d)
		if parent == d {
			return dir
		}
		d = parent
	}
}
