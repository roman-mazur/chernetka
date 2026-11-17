package extlsp

import (
	"context"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"time"

	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
	"rmazur.io/watch/dirwatch"
)

// The server reads the files it's not sent the text of from the disk, but it
// doesn't watch them: it relies on the editor to tell about their changes
// (workspace/didChangeWatchedFiles). Without that, the files edited elsewhere,
// like in another editor pane, or changed by a git checkout, stay as they were
// for the server, which then reports stale problems, like an undefined symbol
// that was just added in another file. The server registers the glob patterns
// of the files it's interested in once it loads the workspace, and the
// fileWatcher watches the directories of the patterns for it.

// watchDelay groups the events of a change, like a git checkout touching many files.
const watchDelay = 100 * time.Millisecond

// newFileWatcher starts the directory watcher. Tests may replace it.
var newFileWatcher = dirwatch.New

// fileWatcher watches the files the server asks for and tells it about their changes.
type fileWatcher struct {
	mu       sync.Mutex
	ctx      context.Context
	srv      *server
	client   lspClient
	w        *dirwatch.Watcher
	stops    []func()
	patterns []string // the latest patterns the server asked for, applied once started
	started  bool
	stopped  bool
}

// watch replaces the watched patterns with the given ones. It's called on the
// connection goroutine: the watches are set up in the background.
func (fw *fileWatcher) watch(patterns []string) {
	fw.mu.Lock()
	defer fw.mu.Unlock()
	fw.patterns = patterns
	if fw.srv != nil {
		fw.srv.le.Debugf("the server asks to watch %q", patterns)
	}
	if fw.started {
		go fw.apply(patterns)
	}
}

// start lets the watcher send the changes to the client.
func (fw *fileWatcher) start(ctx context.Context, srv *server, client lspClient) {
	fw.mu.Lock()
	defer fw.mu.Unlock()
	fw.ctx, fw.srv, fw.client, fw.started = ctx, srv, client, true
	if fw.patterns != nil {
		go fw.apply(fw.patterns)
	}
}

// apply watches the directories of the patterns instead of the ones watched before.
func (fw *fileWatcher) apply(patterns []string) {
	fw.mu.Lock()
	defer fw.mu.Unlock()
	if fw.stopped || !slices.Equal(patterns, fw.patterns) {
		return // Stopped, or newer patterns are on their way.
	}
	for _, stop := range fw.stops {
		stop()
	}
	fw.stops = nil
	if fw.w == nil {
		w, err := newFileWatcher(func(err error) { fw.srv.le.Logf("watch: %s", err) })
		if err != nil {
			fw.srv.le.Logf("cannot watch the files for the server: %s", err)
			return
		}
		fw.w = w
	}
	for _, p := range watchedDirs(patterns, fw.srv.rootDir) {
		dir, accept := p.dir, p.accept
		stop, err := fw.w.WatchTree(dir, watchDelay, func(paths []string) { fw.changed(paths, accept) })
		if err != nil {
			fw.srv.le.Logf("cannot watch %s: %s", dir, err)
			continue
		}
		fw.srv.le.Debugf("watching %s for the server", dir)
		fw.stops = append(fw.stops, stop)
	}
}

// changed tells the server about the changed files the server is interested in.
// It's called on the watcher goroutine.
func (fw *fileWatcher) changed(paths []string, accept func(path string) bool) {
	var events []protocol.FileEvent
	for _, p := range paths {
		if !accept(p) {
			continue
		}
		ev := protocol.FileEvent{URI: uri.File(p), Type: protocol.FileChangeTypeChanged}
		if info, err := os.Stat(p); err != nil {
			ev.Type = protocol.FileChangeTypeDeleted
		} else if info.IsDir() {
			continue
		}
		events = append(events, ev)
	}
	if len(events) == 0 {
		return
	}
	fw.mu.Lock()
	ctx, client, stopped := fw.ctx, fw.client, fw.stopped
	fw.mu.Unlock()
	if stopped {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if err := client.DidChangeWatchedFiles(ctx, events); err != nil {
		fw.srv.le.Debugf("didChangeWatchedFiles failed: %s", err)
	}
}

func (fw *fileWatcher) stop() {
	fw.mu.Lock()
	defer fw.mu.Unlock()
	fw.stopped = true
	for _, stop := range fw.stops {
		stop()
	}
	fw.stops = nil
	if fw.w != nil {
		_ = fw.w.Close()
		fw.w = nil
	}
}

// watchedDir is a directory to watch with the filter of the files to report.
type watchedDir struct {
	dir    string
	accept func(path string) bool
}

// watchedDirs returns the directories to watch for the glob patterns, the relative
// ones taken from the root. A directory inside another one is left out: the whole
// tree is watched. The patterns like "**/*.{go,mod}" filter the files by their
// extensions; the files of the other patterns are all reported.
func watchedDirs(patterns []string, root string) []watchedDir {
	type filter struct {
		exts []string // empty for any file
	}
	dirs := make(map[string]*filter)
	for _, p := range patterns {
		dir, glob := splitPattern(p)
		if !filepath.IsAbs(dir) {
			dir = filepath.Join(root, dir)
		}
		dir = filepath.Clean(dir)
		f, ok := dirs[dir]
		if !ok {
			f = &filter{}
			dirs[dir] = f
		}
		if exts := globExtensions(glob); len(exts) > 0 && (len(f.exts) > 0 || !ok) {
			f.exts = append(f.exts, exts...)
		} else {
			f.exts = nil // Any file.
		}
	}
	paths := make([]string, 0, len(dirs))
	for d := range dirs {
		paths = append(paths, d)
	}
	slices.Sort(paths)
	var res []watchedDir
	for _, d := range paths {
		if n := len(res); n > 0 && strings.HasPrefix(d, res[n-1].dir+string(filepath.Separator)) {
			continue // Inside a watched directory.
		}
		exts := dirs[d].exts
		accept := func(string) bool { return true }
		if len(exts) > 0 {
			accept = func(path string) bool { return slices.Contains(exts, strings.TrimPrefix(filepath.Ext(path), ".")) }
		}
		res = append(res, watchedDir{dir: d, accept: accept})
	}
	return res
}

// splitPattern splits the glob pattern into the directory before its first
// wildcard and the rest, slash-separated.
func splitPattern(p string) (dir, glob string) {
	i := strings.IndexAny(p, "*?[{")
	if i < 0 {
		return filepath.FromSlash(filepath.Dir(p)), filepath.Base(p)
	}
	j := strings.LastIndex(p[:i], "/")
	if j < 0 {
		return ".", p
	}
	return filepath.FromSlash(p[:max(j, 1)]), p[j+1:]
}

// globExtensions returns the extensions a pattern like "**/*.{go,mod}" or "*.go"
// matches, or nil if it's not such a pattern.
func globExtensions(glob string) []string {
	base := glob[strings.LastIndex(glob, "/")+1:]
	rest, ok := strings.CutPrefix(base, "*.")
	if !ok || strings.ContainsAny(rest, "*?[") {
		return nil
	}
	if strings.HasPrefix(rest, "{") && strings.HasSuffix(rest, "}") {
		return strings.Split(rest[1:len(rest)-1], ",")
	}
	return []string{rest}
}
