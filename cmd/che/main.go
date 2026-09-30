package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"rmazur.io/chernetka/internal/cheimg"
	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/debugflags"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/editor/extcomment"
	"rmazur.io/chernetka/internal/editor/extd2"
	"rmazur.io/chernetka/internal/editor/extgo"
	"rmazur.io/chernetka/internal/editor/extlsp"
	"rmazur.io/chernetka/internal/editor/extmd"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl"
	"rmazur.io/chernetka/internal/logger"
	"rmazur.io/chernetka/internal/remotectl"
	"rmazur.io/chernetka/internal/vt"
	"rmazur.io/chernetka/internal/vt/tabscope"
)

func main() {
	term, err := vt.SystemTerminal()
	if err != nil {
		log.Fatal("no terminal detected:", err)
	}
	defer term.Close()

	rootFlag := flag.String("root", "", "project `dir` to search files in (defaults to the opened directory or the working one)")
	flag.Parse()
	go tabscope.ID() // Resolve the terminal tab early: talking to the terminal is slow.
	var (
		edit    editor.Editor
		skipCtl bool
	)

	logf, _ := logger.UserLogFile()
	edit.LogEmbed = logger.Embed(logf)
	edit.LogDebug = debugflags.IsEnabled("logdebug")
	viewer := newImageViewer(logf)
	delegate := editDelegate{
		edit:   &edit,
		logf:   logf,
		viewer: viewer,
		replaceWithViewer: func(path string) error {
			return replaceWithImageViewer(path, func() { _ = term.Close() })
		},
	}

	edit.OpenPath = delegate.openFile
	edit.ShowDiff = delegate.showDiff
	debugEnv(logf)

	edit.Extend(new(extlsp.Integration))
	edit.Extend(new(extsyntaxhl.Integration))
	edit.Extend(&extd2.Integration{Viewer: viewer})
	edit.Extend(new(extmd.Integration))
	edit.Extend(new(extcomment.Integration))
	edit.Extend(&extgo.Integration{Runner: paneRunner{logf: logf}})

	if flag.NArg() < 1 {
		stat, err := os.Stdin.Stat()
		if err != nil {
			panic(err)
		}
		if pipeUsed := stat.Mode()&os.ModeCharDevice == 0; pipeUsed {
			_ = edit.OpenReader("", os.Stdin)
		} else {
			edit.New()
		}
	} else {
		path := flag.Arg(0)

		info, err := os.Stat(path)
		if err != nil {
			_ = term.Close() // log.Fatal skips deferred calls.
			log.Fatal("cannot get path info:", err)
		}
		if info.IsDir() {
			delegate.root = path
			edit.OpenDir(path, &delegate)
			skipCtl = true
		} else {
			delegate.openFile(path)
		}
	}
	if *rootFlag != "" {
		delegate.root = *rootFlag
	}
	if abs, err := filepath.Abs(delegate.root); err == nil {
		delegate.root = abs
	}
	edit.Root = delegate.root

	if !skipCtl {
		defer startCtlServer(&delegate, logf)()
	}

	edit.Run(term)
}

// startCtlServer receives the files to open from the other che processes in the same terminal tab.
// Resolving the tab may take a while, so the server starts in the background.
// The returned function stops the server.
func startCtlServer(e remotectl.Executor, logf logger.Func) (stop func()) {
	var (
		mu      sync.Mutex
		srv     *remotectl.Server
		stopped bool
	)
	go func() {
		s, err := remotectl.NewServer(remotectl.EditorEndpoint.InTab())
		if err != nil {
			logf("ctl error: %s", err)
			return
		}
		mu.Lock()
		if stopped {
			mu.Unlock()
			_ = s.Close()
			return
		}
		srv = s
		mu.Unlock()
		s.Run(e, logf)
	}()
	return func() {
		mu.Lock()
		defer mu.Unlock()
		stopped = true
		if srv != nil {
			_ = srv.Close()
		}
	}
}

const doDebugEnv = false

func debugEnv(logf logger.Func) {
	if !doDebugEnv {
		return
	}
	env := os.Environ()
	envKeys := make([]string, len(env))
	for i := range env {
		envKeys[i], _, _ = strings.Cut(env[i], "=")
	}
	logf("env: %v", envKeys)
	for _, key := range []string{"GHOSTTY_SHELL_FEATURES", "TERM_PROGRAM"} {
		logf("%s = %q", key, os.Getenv(key))
	}
}

type editDelegate struct {
	edit   *editor.Editor
	logf   logger.Func
	viewer extd2.Viewer
	root   string // the project directory, paths passed to OpenFile are relative to it

	// replaceWithViewer turns che into che-img showing the image at path.
	// It returns only if che-img cannot be found.
	replaceWithViewer func(path string) error
}

// openFile opens the file at path in the editor. Images are shown with che-img instead,
// and diagrams are shown with che-img in addition to being opened for editing.
// If nothing is open in the editor yet, che turns into che-img to show an image.
// It must be called on the editor goroutine or before the editor runs.
func (ed *editDelegate) openFile(path string) {
	imgKind := cheimg.KindOf(path)
	if imgKind == cheimg.KindImage && ed.edit.Top() == nil {
		ed.logf("replacing the editor with che-img for %s", path)
		err := fmt.Errorf("cannot start che-img: %w", ed.replaceWithViewer(path))
		ed.logf("%s", err)
		ed.edit.OpenBuffer(&editor.Buffer{Path: path, Content: &content.ErrorContent{Error: err}})
		return
	}

	if imgKind != cheimg.KindUnsupported {
		if abs, err := filepath.Abs(path); err == nil {
			// The viewer runs in another process that may have a different working directory.
			path = abs
		}
		ed.viewer.Show(cheimg.Item{Path: path})
	}
	if imgKind != cheimg.KindImage {
		(&editor.OpenFile{Path: path}).DoOnEditor(ed.edit)
	}
}

// showDiff shows the git changes of the file at path in a new terminal pane.
func (ed *editDelegate) showDiff(path string) {
	go launchCommandViaTerminal("right", gitDiffCommand(path), ed.logf)
}

// gitDiffCommand returns a shell command showing the changes of the file at path
// since the last commit, both staged and not.
func gitDiffCommand(path string) string {
	return fmt.Sprintf("git -C %s diff HEAD -- %s", shellQuote(filepath.Dir(path)), shellQuote(path))
}

// paneRunner runs the commands of the Go line actions in a new terminal pane below.
type paneRunner struct {
	logf logger.Func
}

func (pr paneRunner) Run(cmd extgo.Command) {
	go launchCommandViaTerminal("down", fmt.Sprintf("cd %s && %s", shellQuote(cmd.Dir), cmd.Line), pr.logf)
}

// shellQuote quotes s as a single word for a POSIX shell: nothing is expanded inside
// the single quotes, and a single quote is closed, escaped, and reopened.
func shellQuote(s string) string {
	return "'" + strings.ReplaceAll(s, "'", `'\''`) + "'"
}

func launchCommandViaTerminal(paneDirection string, cmdLine string, logf logger.Func) {
	if err := runInNewPane(paneDirection, cmdLine); err != nil {
		logf("cannot run %s: %s", cmdLine, err)
	}
}

func (ed *editDelegate) ExecuteCommand(cmd remotectl.CommandData) {
	var editorCommand editor.Command

	switch cmd.Action {
	case "open":
		path := cmd.Args[0]
		editorCommand = editor.CommandFunc(func(*editor.Editor) { ed.openFile(path) })
	default:
		ed.logf("ignore remote cmd: %s", cmd.Action)
		return
	}

	ed.logf("remote cmd: %s", cmd.Action)
	ed.edit.Send(editorCommand)
}

// OpenFile opens the file from the project directory in the main editor of the terminal tab,
// starting one if necessary.
func (ed *editDelegate) OpenFile(path string) {
	if !filepath.IsAbs(path) {
		// The main editor may run in another directory.
		path = filepath.Join(ed.root, path)
	}
	err := ed.sendOpenCommand(path)
	if err == nil {
		return
	}
	if errors.Is(err, os.ErrNotExist) {
		ed.logf("starting main editor")
		err := openMainEditor(ed.root, path)
		if err != nil {
			ed.logf("error with main editor: %s", err)
			ed.openFile(path)
		}
	} else {
		ed.logf("cannot send a command: %s", err)
		ed.openFile(path)
	}
}

func (ed *editDelegate) sendOpenCommand(path string) error {
	return remotectl.SendCommand(remotectl.EditorEndpoint.InTab(), &remotectl.CommandData{
		Action: "open",
		Args:   []string{path},
	})
}
