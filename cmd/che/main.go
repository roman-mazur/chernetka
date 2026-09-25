package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"rmazur.io/chernetka/internal/cheimg"
	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/debugflags"
	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/editor/extd2"
	"rmazur.io/chernetka/internal/editor/extlsp"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl"
	"rmazur.io/chernetka/internal/logger"
	"rmazur.io/chernetka/internal/remotectl"
	"rmazur.io/chernetka/internal/vt"
)

func main() {
	term, err := vt.SystemTerminal()
	if err != nil {
		log.Fatal("no terminal detected:", err)
	}
	defer term.Close()

	flag.Parse()
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

	debugEnv(logf)

	edit.Extend(new(extlsp.Integration))
	edit.Extend(new(extsyntaxhl.Integration))
	edit.Extend(&extd2.Integration{Viewer: viewer})

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
			log.Fatal("cannot get path info:", err)
		}
		if info.IsDir() {
			edit.OpenDir(path, &delegate)
			skipCtl = true
		} else {
			delegate.openFile(path)
		}
	}

	if !skipCtl {
		srv, err := remotectl.NewServer(remotectl.EditorEndpoint)
		if err == nil {
			defer srv.Close()
			go srv.Run(&delegate, logf)
		} else {
			logf("ctl error: %s", err)
		}
	}

	edit.Run(term)
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

func (ed *editDelegate) OpenFile(path string) {
	err := ed.sendOpenCommand(path)
	if err == nil {
		return
	}
	if errors.Is(err, os.ErrNotExist) {
		ed.logf("starting main editor")
		err := openMainEditor(ed.edit, path)
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
	return remotectl.SendCommand(remotectl.EditorEndpoint, &remotectl.CommandData{
		Action: "open",
		Args:   []string{path},
	})
}
