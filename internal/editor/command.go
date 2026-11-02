package editor

import (
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"rmazur.io/chernetka/internal/content"
	"rmazur.io/chernetka/internal/content/symbol"
	"rmazur.io/chernetka/internal/editor/input"
)

const (
	cmdPrefixNewFile   = "n "
	cmdPrefixQuickOpen = "e "
)

// exPrompt runs the typed command on Enter, see runExCommand.
// Typing ":e " turns it into the file picker, and ":n " into the new file prompt.
type exPrompt struct {
	c *cmdLine
}

func newExPrompt(c *cmdLine) prompt { return &exPrompt{c: c} }

func (p *exPrompt) prefix() string { return ":" }

func (p *exPrompt) input(*Editor, input.Key) bool { return false }

func (p *exPrompt) changed(e *Editor) {
	switch {
	case strings.HasPrefix(p.c.text, cmdPrefixQuickOpen):
		p.c.prompt = newQuickOpen(p.c)
		p.c.prompt.changed(e)
	case strings.HasPrefix(p.c.text, cmdPrefixNewFile):
		p.c.text += selectedDirPrefix(e, p.c.buf)
		p.c.prompt = &newFilePrompt{c: p.c}
		p.c.prompt.changed(e)
	}
}

func (p *exPrompt) submit(e *Editor) (quit bool) { return runExCommand(p.c.buf, p.c.text, &e.rPrefs) }

func (p *exPrompt) cancel(*Editor) {}

func (p *exPrompt) height(int) int { return 1 }

func (p *exPrompt) render(out io.Writer, w int) { renderCmdline(out, w, p.prefix()+p.c.text, "") }

func runExCommand(buf *Buffer, cmd string, prefs *RenderPrefs) (quit bool) {
	if n, err := strconv.Atoi(cmd); err == nil {
		GoToLine(n).DoOnBuffer(buf, *prefs)
		return false
	}
	for len(cmd) > 0 {
		key := cmd[0:1]
		cmd = cmd[1:]

		switch key {
		case "q":
			return true

		case "w":
			dstPath := buf.Path
			if len(cmd) > 2 {
				cmd, dstPath, _ = strings.Cut(cmd, " ")
			}
			(&Save{DstPath: dstPath}).DoOnBuffer(buf, *prefs)

		case "p":
			switch cmd {
			case "bcopy":
				ClipboardCopy.DoOnBuffer(buf, *prefs)
			case "bcut":
				ClipboardCut.DoOnBuffer(buf, *prefs)
			case "bpaste":
				ClipboardPaste.DoOnBuffer(buf, *prefs)
			}
		}
	}
	return false
}

// newFilePrompt asks for the path of a new file, relative to the project root.
// The path starts with the directory of the buffer, see selectedDirPrefix.
// The buffer of the file is started on Enter, and the file is written when it's saved.
type newFilePrompt struct {
	c *cmdLine

	path string // the absolute path of the file to create or open, empty if the text does not name a file
	info string // what is shown on the right of the command line
}

func (p *newFilePrompt) prefix() string { return ":" }

func (p *newFilePrompt) input(*Editor, input.Key) bool { return false }

// changed resolves the typed path. The prompt turns back into the ex command if the prefix is deleted.
func (p *newFilePrompt) changed(e *Editor) {
	inputPath, active := strings.CutPrefix(p.c.text, cmdPrefixNewFile)
	if !active {
		p.c.prompt = newExPrompt(p.c)
		return
	}
	p.path, p.info = filePathFromSlash(e, inputPath), ""
	if p.path == "" {
		return
	}

	if info, err := os.Stat(p.path); err == nil {
		if info.IsDir() {
			p.path = ""
			p.info = "(dir)"
		} else {
			p.info = "(exists)"
		}
	}
}

func (p *newFilePrompt) submit(e *Editor) (quit bool) {
	if p.path != "" {
		e.open(p.path)
	}
	return false
}

func (p *newFilePrompt) cancel(*Editor) {}

func (p *newFilePrompt) height(int) int { return 1 }

func (p *newFilePrompt) render(out io.Writer, w int) {
	cmd := p.prefix() + p.c.text
	const margins = 3 // A gap after the command, and a space with the cursor at the end.
	info := ""
	avail := w - symbol.DisplayWidth(cmd, symbol.NoTabExpansion) - margins
	if avail >= symbol.DisplayWidth(p.info, symbol.NoTabExpansion) {
		info = p.info
	}
	renderCmdline(out, w, cmd, info)
}

// filePathFromSlash returns the path of the file typed as the slash-separated p,
// resolving it against the project root. It's empty if p does not name a file.
func filePathFromSlash(e *Editor, p string) string {
	if p == "" || strings.HasSuffix(p, "/") {
		return ""
	}
	res := filepath.FromSlash(p)
	if !filepath.IsAbs(res) {
		res = filepath.Join(e.root(), res)
	}
	return res
}

// selectedDirPrefix returns the slash-separated directory of the buffer relative to the project
// root, with a trailing slash: the directory of the file, or the directory selected in the tree.
// It's empty for the project root and for a buffer without a file.
func selectedDirPrefix(e *Editor, buf *Buffer) string {
	dir := e.root()
	if fc, isDir := buf.Content.(*content.FsContent); isDir {
		// The tree paths are relative to its root, not to the working directory.
		dir = filepath.Join(fc.RootPath(), filepath.FromSlash(fc.DirAt(buf.c().Line)))
	} else if buf.Path != "" {
		dir = filepath.Dir(buf.Path)
	}
	dir = displayPath(e.root(), dir)
	if dir == "." {
		return ""
	}
	dir = filepath.ToSlash(dir)
	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}
	return dir
}
