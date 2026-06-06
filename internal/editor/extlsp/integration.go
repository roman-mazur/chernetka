package extlsp

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
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/lsp"
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

// Integration encapsulates the logic of integrating with an LSP client in an editor.Editor and editor.Buffer.
type Integration struct {
	Starter lspStarter

	client lspClient
	cancel context.CancelFunc

	startError error
}

func (le *Integration) ID() string { return "lsp" }

func (le *Integration) MakeBufferData(buf *editor.Buffer) editor.BufferExtData {
	if !strings.HasSuffix(buf.Path, ".go") {
		return nil
	}

	absPath, err := filepath.Abs(buf.Path)
	if err != nil {
		return nil
	}
	if err := le.ensureLSP(filepath.Dir(absPath)); err != nil {
		return nil
	}

	var bufData BufferData
	bufData.SetPath(absPath)
	bufData.version = 1
	_ = le.client.DidOpen(context.Background(), bufData.docUri, "go", buf.Text(), bufData.version)
	return &bufData
}

func (le *Integration) AfterEdit(e *editor.Editor, buf *editor.Buffer) {
	data, active := le.activeOn(buf)
	if !active {
		return
	}
	le.sendChange(buf, data) // TODO: consider bouncing and doing asynchronously.

	cx, cy := buf.Pos()
	var line string
	if lines := buf.Content.Lines(); cy < len(lines) {
		line = lines[cy].String()
	}
	cx = min(cx, len(line))
	le.askForNewSuggestion(line, cx, cy, e, data)
}

// ensureLSP performs the one-time start of the language server. After the
// first call, subsequent calls return the cached start error (if any) without
// retrying — a failed start usually means gopls isn't installed.
func (le *Integration) ensureLSP(fileDir string) error {
	if le.client != nil || le.startError != nil {
		return le.startError
	}
	starter := le.Starter
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

func (le *Integration) activeOn(buf *editor.Buffer) (*BufferData, bool) {
	if le.client == nil {
		return nil, false
	}

	data := buf.ExtensionData(le.ID())
	if data == nil {
		return nil, false
	}
	return data.(*BufferData), data.(*BufferData).docUri != ""
}

func (le *Integration) sendChange(buf *editor.Buffer, bufData *BufferData) {
	bufData.version++
	_ = le.client.DidChange(context.Background(), bufData.docUri, buf.Text(), bufData.version)
}

func (le *Integration) askForNewSuggestion(line string, cx, cy int, e *editor.Editor, bufData *BufferData) {
	bufData.reqCompletion++

	var (
		client  = le.client
		fileURI = bufData.docUri
		req     = bufData.reqCompletion
		ln      = uint32(cy)
		ch      = utf16Len(line[:cx])
	)

	// TODO: replace with smarter logic when to ask for a suggestion.
	go func() {
		items, err := client.Completion(context.Background(), fileURI, ln, ch)
		if err != nil {
			return
		}
		suggestions := extractLspSuggestions(items, line, cx)
		e.Post(editor.CommandFunc(func(e *editor.Editor) {
			if bufData.reqCompletion != req {
				return // A newer edit superseded this request.
			}
			bufData.Assign(suggestions)
			e.RequestLayout()
		}))
	}()
}

func (le *Integration) Close() error {
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

func utf16Len(s string) (n uint32) {
	for _, r := range s {
		n += uint32(utf16.RuneLen(r))
	}
	return n
}
