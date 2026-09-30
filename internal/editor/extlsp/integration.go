package extlsp

import (
	"context"
	"errors"
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
	Definition(ctx context.Context, fileURI uri.URI, line, character uint32) ([]protocol.Location, error)
	OrganizeImports(ctx context.Context, fileURI uri.URI) ([]protocol.TextEdit, error)
	Formatting(ctx context.Context, fileURI uri.URI, opts protocol.FormattingOptions) ([]protocol.TextEdit, error)
	Shutdown(ctx context.Context) error
}

// lspStarter spawns and initialises an LSP client of the language with the
// given LSP identifier for the given workspace root.
type lspStarter func(ctx context.Context, languageID, rootDir string) (lspClient, error)

// Integration encapsulates the logic of integrating with LSP clients in an editor.Editor and editor.Buffer.
//
// Every edit of a buffer in a supported language (see languages) is pushed to
// the language server, and when the cursor is in the middle of typing an
// identifier (or right after a selector dot) the edit is followed by a
// completion request. The best candidates are shown as an inline suggestion
// that can be accepted with Tab.
//
// Saving a buffer formats it with the server first.
//
// Ctrl+click on a symbol goes to its definition found by the server.
type Integration struct {
	logger.LogEmbed

	Starter lspStarter

	// servers are the language servers started so far by language. Only used
	// on the editor loop.
	servers map[string]*server
}

func (le *Integration) ID() string { return "lsp" }

func (le *Integration) MakeBufferData(_ editor.Sender, buf *editor.Buffer) editor.BufferExtData {
	lang := languageForPath(buf.Path)
	if lang == nil {
		return nil
	}

	absPath, err := filepath.Abs(buf.Path)
	if err != nil {
		return nil
	}
	srv := le.serverFor(lang, filepath.Dir(absPath))

	bufData := BufferData{srv: srv, buf: buf}
	bufData.SetPath(absPath)
	bufData.version = 1
	// The server gets the document once it's ready.
	srv.queue.push(syncReq{
		bufData: &bufData,
		text:    buf.Text(),
		version: bufData.version,
	})
	return &bufData
}

// serverFor returns the server of the language, starting it for the workspace
// of dir if it's the first buffer of the language. The server keeps serving
// the workspace it's started for.
func (le *Integration) serverFor(lang *language, dir string) *server {
	if srv, ok := le.servers[lang.id()]; ok {
		return srv
	}
	if debugImpl {
		le.LogDebug = true // Before the loop starts using the logger.
	}
	ctx, cancel := context.WithCancel(context.Background())
	srv := &server{
		le:       le,
		lang:     lang,
		rootDir:  lang.root(dir),
		cancel:   cancel,
		loopDone: make(chan struct{}),
	}
	srv.queue.wake = make(chan struct{}, 1)
	if le.servers == nil {
		le.servers = make(map[string]*server)
	}
	le.servers[lang.id()] = srv
	go srv.syncLoop(ctx)
	return srv
}

// AfterEdit runs on the editor loop after every buffer mutation.
func (le *Integration) AfterEdit(loop editor.Sender, buf *editor.Buffer) {
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
		editor:  loop,
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
	if buf.Mode() == editor.ModeInsert && data.srv.lang.goImports {
		// A package used without an import, like "strconv.Itoa(" was typed.
		pkg, ok := qualifiedBeforeCursor(line[:cx])
		req.addImports = ok && !isImported(req.text, pkg)
	}
	data.srv.queue.push(req)
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

func (le *Integration) activeOn(buf *editor.Buffer) (*BufferData, bool) {
	data, ok := buf.ExtensionData(le.ID()).(*BufferData)
	if !ok || data.srv.failed.Load() {
		return nil, false
	}
	return data, data.docUri != ""
}

func (le *Integration) Close() error {
	var errs []error
	for id, srv := range le.servers {
		errs = append(errs, srv.close())
		delete(le.servers, id)
	}
	return errors.Join(errs...)
}

// server is the language server of one language together with the goroutine
// talking to it.
type server struct {
	le      *Integration
	lang    *language
	rootDir string // the workspace root to start the server with

	cancel   context.CancelFunc
	queue    syncQueue
	loopDone chan struct{}
	client   lspClient   // owned by the sync loop until it's done
	ready    atomic.Bool // the server has started
	failed   atomic.Bool // the server couldn't start
}

// syncLoop starts the language server, then sends it buffer changes and
// completion requests. Using a single goroutine guarantees the server sees a
// change before a completion request for the text it produced. Changes that
// pile up while the server is busy (or starting) are coalesced: only the
// latest text of each buffer is sent.
func (srv *server) syncLoop(ctx context.Context) {
	defer close(srv.loopDone)
	for {
		select {
		case <-ctx.Done():
			return
		case <-srv.queue.wake:
		}
		reqs := srv.queue.take()
		if srv.client == nil && !srv.start(ctx) {
			return // Without a server, there is nothing to do.
		}
		for _, req := range reqs {
			srv.sync(ctx, srv.client, req)
		}
	}
}

// start launches the language server. It may take a while if a gopls matching
// the workspace Go version needs to be built. A failure is final: it usually
// means the server isn't installed.
func (srv *server) start(ctx context.Context) bool {
	le := srv.le
	starter := le.Starter
	if starter == nil {
		starter = func(ctx context.Context, _, rootDir string) (lspClient, error) {
			return srv.lang.start(ctx, le, rootDir)
		}
	}
	le.Logf("starting the %s LSP server for %s", srv.lang.id(), srv.rootDir)
	started := time.Now()
	client, err := starter(ctx, srv.lang.id(), srv.rootDir)
	if err != nil {
		le.Logf("cannot start the %s LSP server: %s", srv.lang.id(), err)
		srv.failed.Store(true)
		return false
	}
	le.Logf("%s LSP server is ready in %s", srv.lang.id(), time.Since(started).Round(time.Millisecond))
	srv.client = client
	srv.ready.Store(true)
	return true
}

func (srv *server) sync(ctx context.Context, client lspClient, req syncReq) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	le := srv.le
	uri := req.bufData.docUri
	data := req.bufData
	synced := srv.syncText(ctx, client, req)
	if req.format != nil {
		// Reply before anything else: the editor is waiting.
		var edits []protocol.TextEdit
		if synced {
			edits = srv.formatting(ctx, client, req)
		}
		req.format.res <- edits
	}
	if !synced {
		return
	}
	if req.addImports {
		srv.addImports(ctx, client, req)
	}
	if req.definition != nil {
		srv.goToDefinition(ctx, client, req)
	}
	if req.completion == nil || srv.queue.hasPending(req.bufData) {
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
	if !srv.lang.rankedCompletion {
		items = rankByPrefix(items, at.line, at.cx)
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

// syncText sends the server the text of req.
func (srv *server) syncText(ctx context.Context, client lspClient, req syncReq) bool {
	data := req.bufData
	if data.serverOpen {
		return srv.sendChange(ctx, client, req) == nil
	}
	if err := client.DidOpen(ctx, data.docUri, srv.lang.id(), req.text, req.version); err != nil {
		srv.le.Debugf("didOpen failed: %s", err)
		return false
	}
	data.serverOpen, data.serverText, data.serverSynced = true, req.text, true
	return true
}

// formatTimeout is how long saving a buffer waits for the server to format it.
const formatTimeout = time.Second

// format formats the buffer with the server. It blocks the editor loop until
// the server replies, which usually takes milliseconds, but no longer than
// formatTimeout. The buffer is left as is if the server isn't ready, fails
// (as it does on syntax errors), or is too slow.
func (srv *server) format(buf *editor.Buffer, data *BufferData, prefs editor.RenderPrefs) {
	if !srv.ready.Load() {
		return // Saving doesn't wait for the server to start.
	}
	text := buf.Text()
	res := make(chan []protocol.TextEdit, 1)
	data.version++
	srv.queue.push(syncReq{
		buf:     buf,
		bufData: data,
		text:    text,
		version: data.version,
		format:  &formatReq{res: res, opts: protocol.FormattingOptions{TabSize: uint32(prefs.TabSize)}},
	})

	timer := time.NewTimer(formatTimeout)
	defer timer.Stop()
	var edits []protocol.TextEdit
	select {
	case edits = <-res:
	case <-timer.C:
		srv.le.Logf("formatting timed out")
		return
	}
	if len(edits) == 0 || buf.Text() != text {
		return
	}
	applyEdits(buf, edits)
}

// formatting returns the edits formatting the text of req.
func (srv *server) formatting(ctx context.Context, client lspClient, req syncReq) []protocol.TextEdit {
	started := time.Now()
	edits, err := client.Formatting(ctx, req.bufData.docUri, req.format.opts)
	if err != nil {
		srv.le.Debugf("formatting failed: %s", err)
		return nil
	}
	srv.le.Debugf("formatting: %d edits in %s", len(edits), time.Since(started).Round(time.Millisecond))
	return edits
}

// addImports adds the imports missing for the packages used in the text of
// req. The server resolves the packages with its "organize imports" action,
// but only the imports it adds are taken: it also removes unused imports,
// which are likely to be used soon while typing. The imports are added if the
// user didn't change them in the meantime.
func (srv *server) addImports(ctx context.Context, client lspClient, req syncReq) {
	le := srv.le
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

// goToDefinition asks the server where the symbol at the position of req is
// defined and moves the editor there.
func (srv *server) goToDefinition(ctx context.Context, client lspClient, req syncReq) {
	le := srv.le
	at := req.definition.at
	locs, err := client.Definition(ctx, req.bufData.docUri, at.Line, at.Character)
	if err != nil {
		le.Debugf("definition failed: %s", err)
		return
	}
	if len(locs) == 0 {
		le.Logf("no definition found at %d:%d", at.Line, at.Character)
		return
	}
	loc := locs[0]
	path := loc.URI.Filename()
	// The column is converted to bytes with the text the server has.
	text := req.text
	if loc.URI != req.bufData.docUri {
		data, err := os.ReadFile(path)
		if err != nil {
			le.Logf("cannot read the definition file: %s", err)
			return
		}
		text = string(data)
	}
	le.Debugf("definition at %s:%d:%d", path, loc.Range.Start.Line, loc.Range.Start.Character)
	req.editor.Send(&editor.GoTo{Path: path, Pos: textPosition(text, loc.Range.Start)})
}

func positionBefore(a, b protocol.Position) bool {
	return a.Line < b.Line || (a.Line == b.Line && a.Character < b.Character)
}

// sendChange sends the server only what changed since the last sync: usually
// a character.
func (srv *server) sendChange(ctx context.Context, client lspClient, req syncReq) error {
	data := req.bufData
	change := lsp.TextChange{Text: req.text}
	if data.serverSynced {
		change = diff(data.serverText, req.text)
	}
	if err := client.DidChange(ctx, data.docUri, req.version, change); err != nil {
		// Unknown state on the server: the next sync sends the whole text.
		data.serverSynced = false
		srv.le.Debugf("didChange failed: %s", err)
		return err
	}
	data.serverText, data.serverSynced = req.text, true
	return nil
}

// close stops the sync loop and shuts the server down.
func (srv *server) close() error {
	srv.cancel()
	<-srv.loopDone
	if srv.client == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err := srv.client.Shutdown(ctx)
	srv.client = nil
	return err
}

// syncReq is the state of a buffer to be sent to the server.
type syncReq struct {
	editor  editor.Sender // nil for a format request
	buf     *editor.Buffer
	bufData *BufferData

	text    string
	version int32

	completion *completionReq // nil if no completion is needed
	addImports bool           // add the imports missing after the change
	format     *formatReq     // nil if no formatting is needed
	definition *definitionReq // nil if no definition is looked for
}

type definitionReq struct {
	at protocol.Position // the symbol position
}

type completionReq struct {
	id int
	at anchor
}

type formatReq struct {
	res  chan<- []protocol.TextEdit // receives the formatting edits
	opts protocol.FormattingOptions
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
			if prev := q.pending[i]; prev.addImports && !req.addImports {
				req.addImports = true
				req.editor = prev.editor // A format request has no editor.
			}
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
