package extlsp

import (
	"context"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unicode"
	"unicode/utf8"

	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
	"rmazur.io/chernetka/internal"
	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/logger"
	"rmazur.io/chernetka/internal/lsp"
)

const debugImpl = true

// lspClient is the subset of an LSP backend the editor needs. *lsp.Client
// satisfies it; tests substitute a fake.
type lspClient interface {
	DidOpen(ctx context.Context, fileURI uri.URI, languageID, text string, version int32) error
	DidChange(ctx context.Context, fileURI uri.URI, version int32, changes ...lsp.TextChange) error
	Completion(ctx context.Context, fileURI uri.URI, line, character uint32) ([]protocol.CompletionItem, error)
	OrganizeImports(ctx context.Context, fileURI uri.URI) ([]protocol.TextEdit, error)
	Shutdown(ctx context.Context) error
}

// lspStarter spawns and initialises an LSP client for the given workspace root.
type lspStarter func(ctx context.Context, rootDir string) (lspClient, error)

func (le *Integration) startLSP(ctx context.Context, rootDir string) (lspClient, error) {
	var cache string
	if dir, err := internal.UserDir(); err == nil {
		cache = filepath.Join(dir, "gopls")
	}
	return lsp.Start(ctx, rootDir, lsp.Options{
		GoplsCache: cache,
		// Calls come with brackets and the cursor placed inside them.
		SnippetSupport: true,
		Logf:           le.Logf,
	})
}

// Integration encapsulates the logic of integrating with an LSP client in an editor.Editor and editor.Buffer.
//
// Every edit of a Go buffer is pushed to the language server, and when the
// cursor is in the middle of typing an identifier (or right after a selector
// dot) the edit is followed by a completion request. The best candidates are
// shown as an inline suggestion that can be accepted with Tab.
type Integration struct {
	logger.LogEmbed

	Starter lspStarter

	once     sync.Once
	cancel   context.CancelFunc
	queue    syncQueue
	loopDone chan struct{}
	client   lspClient   // owned by the sync loop until it's done
	failed   atomic.Bool // the server couldn't start
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
	le.startLoop()

	var bufData BufferData
	bufData.SetPath(absPath)
	bufData.version = 1
	// The server gets the document once it's ready.
	bufData.rootDir = findGoModRoot(filepath.Dir(absPath))
	le.queue.push(syncReq{
		bufData: &bufData,
		text:    buf.Text(),
		version: bufData.version,
	})
	return &bufData
}

// AfterEdit runs on the editor loop after every buffer mutation.
func (le *Integration) AfterEdit(e *editor.Editor, buf *editor.Buffer) {
	data, active := le.activeOn(buf)
	if !active {
		return
	}

	cx, cy := buf.Pos()
	var line string
	if lines := buf.Content.Lines(); cy < len(lines) {
		line = lines[cy].String()
	}
	cx = min(cx, len(line))

	// Keep showing what's still valid of the current suggestion while a fresh
	// completion is on its way.
	data.typeThrough(line, cx, cy)

	// Invalidate completions in flight: they were computed for an older text.
	data.reqCompletion++
	data.version++
	req := syncReq{
		editor:  e,
		buf:     buf,
		bufData: data,
		text:    buf.Text(), // snapshot: the server must see exactly this version
		version: data.version,
	}
	if buf.Mode() == editor.ModeInsert && shouldComplete(line[:cx]) && !identAt(line, cx) {
		req.completion = &completionReq{id: data.reqCompletion, at: anchor{line: line, cx: cx, cy: cy}}
	} else {
		data.ResetSuggestions()
	}
	if buf.Mode() == editor.ModeInsert {
		// A package used without an import, like "strconv.Itoa(" was typed.
		pkg, ok := qualifiedBeforeCursor(line[:cx])
		req.addImports = ok && !isImported(req.text, pkg)
	}
	le.queue.push(req)
}

// shouldComplete decides whether it's worth asking for a completion with the
// cursor after the given text: only while typing an identifier, and right
// after a selector dot. Everywhere else (after spaces, brackets, operators)
// the server has nothing to go on and would return an alphabetical list.
func shouldComplete(beforeCursor string) bool {
	if strings.HasSuffix(beforeCursor, ".") {
		// Not after a number literal like "1.".
		p := identTrailing(beforeCursor[:len(beforeCursor)-1])
		return p != "" && !startsWithDigit(p)
	}
	p := identTrailing(beforeCursor)
	return len(p) >= minPrefix && !startsWithDigit(p)
}

// minPrefix is how many identifier characters must be typed before completion
// kicks in.
const minPrefix = 1

// identAt reports whether an identifier character follows position i: the
// cursor is in the middle of a word, where an inline suggestion doesn't fit.
func identAt(line string, i int) bool {
	r, _ := utf8.DecodeRuneInString(line[i:])
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

func startsWithDigit(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return unicode.IsDigit(r)
}

// startLoop starts the goroutine talking to the language server.
func (le *Integration) startLoop() {
	le.once.Do(func() {
		if debugImpl {
			le.LogDebug = true // Before the loop starts using the logger.
		}
		ctx, cancel := context.WithCancel(context.Background())
		le.cancel = cancel
		le.queue.wake = make(chan struct{}, 1)
		le.loopDone = make(chan struct{})
		go le.syncLoop(ctx)
	})
}

// syncLoop starts the language server, then sends it buffer changes and
// completion requests. Using a single goroutine guarantees the server sees a
// change before a completion request for the text it produced. Changes that
// pile up while the server is busy (or starting) are coalesced: only the
// latest text of each buffer is sent.
func (le *Integration) syncLoop(ctx context.Context) {
	defer close(le.loopDone)
	for {
		select {
		case <-ctx.Done():
			return
		case <-le.queue.wake:
		}
		reqs := le.queue.take()
		if le.client == nil && !le.start(ctx, reqs[0].bufData.rootDir) {
			return // Without a server, there is nothing to do.
		}
		for _, req := range reqs {
			le.sync(ctx, le.client, req)
		}
	}
}

// start launches the language server. It may take a while if a gopls matching
// the workspace Go version needs to be built. A failure is final: it usually
// means gopls isn't installed.
func (le *Integration) start(ctx context.Context, rootDir string) bool {
	starter := le.Starter
	if starter == nil {
		starter = le.startLSP
	}
	le.Logf("starting an LSP server for %s", rootDir)
	started := time.Now()
	client, err := starter(ctx, rootDir)
	if err != nil {
		le.Logf("cannot start the LSP server: %s", err)
		le.failed.Store(true)
		return false
	}
	le.Logf("LSP server is ready in %s", time.Since(started).Round(time.Millisecond))
	le.client = client
	return true
}

func (le *Integration) sync(ctx context.Context, client lspClient, req syncReq) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	uri := req.bufData.docUri
	data := req.bufData
	if !data.serverOpen {
		if err := client.DidOpen(ctx, uri, "go", req.text, req.version); err != nil {
			le.Debugf("didOpen failed: %s", err)
			return
		}
		data.serverOpen, data.serverText, data.serverSynced = true, req.text, true
	} else if err := le.sendChange(ctx, client, req); err != nil {
		return
	}
	if req.addImports {
		le.addImports(ctx, client, req)
	}
	if req.completion == nil || le.queue.hasPending(req.bufData) {
		return // Nothing to ask, or a newer edit makes the answer useless.
	}

	at := req.completion.at
	started := time.Now()
	pos := lspPosition(at.line, content.Position{Line: at.cy, Col: at.cx})
	items, err := client.Completion(ctx, uri, pos.Line, pos.Character)
	if err != nil {
		le.Debugf("completion failed: %s", err)
		return
	}
	suggestions := extractSuggestions(items, at.line, at.cx)
	le.Debugf("completion at %d:%d: %d items, %d suggestions in %s",
		at.cy, at.cx, len(items), len(suggestions), time.Since(started).Round(time.Millisecond))

	id, buf := req.completion.id, req.buf
	req.editor.Send(editor.CommandFunc(func(e *editor.Editor) {
		if data.reqCompletion != id || buf.Mode() != editor.ModeInsert {
			return // A newer edit superseded this request, or the insert is over.
		}
		if cx, cy := buf.Pos(); cx != at.cx || cy != at.cy {
			return // The cursor moved away.
		}
		data.assign(suggestions, at)
		e.RequestRender()
	}))
}

// addImports adds the imports missing for the packages used in the text of
// req. The server resolves the packages with its "organize imports" action,
// but only the imports it adds are taken: it also removes unused imports,
// which are likely to be used soon while typing. The imports are added if the
// user didn't change them in the meantime.
func (le *Integration) addImports(ctx context.Context, client lspClient, req syncReq) {
	edits, err := client.OrganizeImports(ctx, req.bufData.docUri)
	if err != nil {
		le.Debugf("organize imports failed: %s", err)
		return
	}
	organized := req.text
	sort.Slice(edits, func(i, j int) bool { return positionBefore(edits[j].Range.Start, edits[i].Range.Start) })
	for _, te := range edits {
		organized = applyChange(organized, lsp.TextChange{Range: &te.Range, Text: te.NewText})
	}
	added := addedImports(req.text, organized)
	if len(added) == 0 {
		return
	}
	withImports := addImports(req.text, added)
	change := diff(req.text, withImports)
	header := req.text[:importsEnd(req.text)]

	req.editor.SendBufferCmd(editor.BufferCommandFunc(func(buf *editor.Buffer, _ editor.RenderPrefs) {
		if buf != req.buf {
			return // Another buffer is active now: it will be done next time.
		}
		if !strings.HasPrefix(buf.Text(), header) {
			return // The imports were edited in the meantime.
		}
		applyEdits(buf, []protocol.TextEdit{{Range: *change.Range, NewText: change.Text}})
		le.Logf("added imports %v", added)
	}))
}

func positionBefore(a, b protocol.Position) bool {
	return a.Line < b.Line || (a.Line == b.Line && a.Character < b.Character)
}

// sendChange sends the server only what changed since the last sync: usually
// a character.
func (le *Integration) sendChange(ctx context.Context, client lspClient, req syncReq) error {
	data := req.bufData
	change := lsp.TextChange{Text: req.text}
	if data.serverSynced {
		change = diff(data.serverText, req.text)
	}
	if err := client.DidChange(ctx, data.docUri, req.version, change); err != nil {
		// Unknown state on the server: the next sync sends the whole text.
		data.serverSynced = false
		le.Debugf("didChange failed: %s", err)
		return err
	}
	data.serverText, data.serverSynced = req.text, true
	return nil
}

func (le *Integration) activeOn(buf *editor.Buffer) (*BufferData, bool) {
	if le.failed.Load() {
		return nil, false
	}

	data, ok := buf.ExtensionData(le.ID()).(*BufferData)
	if !ok {
		return nil, false
	}
	return data, data.docUri != ""
}

func (le *Integration) Close() error {
	if le.cancel == nil {
		return nil // Never started.
	}
	le.cancel()
	<-le.loopDone
	le.cancel = nil
	if le.client == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err := le.client.Shutdown(ctx)
	le.client = nil
	return err
}

// syncReq is the state of a buffer to be sent to the server.
type syncReq struct {
	editor  *editor.Editor
	buf     *editor.Buffer
	bufData *BufferData

	text    string
	version int32

	completion *completionReq // nil if no completion is needed
	addImports bool           // add the imports missing after the change
}

type completionReq struct {
	id int
	at anchor
}

// syncQueue keeps the latest syncReq per buffer.
type syncQueue struct {
	mu      sync.Mutex
	pending []syncReq
	wake    chan struct{}
}

func (q *syncQueue) push(req syncReq) {
	q.mu.Lock()
	replaced := false
	for i := range q.pending {
		if q.pending[i].bufData == req.bufData {
			// The latest state wins, but the imports must still be checked.
			req.addImports = req.addImports || q.pending[i].addImports
			q.pending[i] = req
			replaced = true
		}
	}
	if !replaced {
		q.pending = append(q.pending, req)
	}
	q.mu.Unlock()

	select {
	case q.wake <- struct{}{}:
	default:
	}
}

func (q *syncQueue) take() []syncReq {
	q.mu.Lock()
	defer q.mu.Unlock()
	res := q.pending
	q.pending = nil
	return res
}

func (q *syncQueue) hasPending(data *BufferData) bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	for _, req := range q.pending {
		if req.bufData == data {
			return true
		}
	}
	return false
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
