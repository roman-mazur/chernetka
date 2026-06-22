package main

import (
	"errors"
	"flag"
	"log"
	"os"

	"rmazur.io/chernetka/internal/editor"
	"rmazur.io/chernetka/internal/editor/extlsp"
	"rmazur.io/chernetka/internal/editor/extsyntaxhl"
	"rmazur.io/chernetka/internal/logger"
	"rmazur.io/chernetka/internal/remotectl"
)

func main() {
	flag.Parse()
	var (
		ttyFile = os.Stdin
		edit    editor.Editor
		skipCtl bool
	)

	logf, _ := logger.UserLogFile()
	delegate := editDelegate{edit: &edit, logf: logf}

	edit.Extend(new(extlsp.Integration))
	edit.Extend(new(extsyntaxhl.Integration))

	if flag.NArg() < 1 {
		stat, err := os.Stdin.Stat()
		if err != nil {
			panic(err)
		}
		if pipeUsed := stat.Mode()&os.ModeCharDevice == 0; pipeUsed {
			_ = edit.OpenReader("", os.Stdin)
			ttyFile, err = os.Open("/dev/tty")
			if err != nil {
				panic(err)
			}
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
			(&editor.OpenFile{Path: path}).DoOnEditor(&edit)
		}
	}

	if !skipCtl {
		srv, err := remotectl.NewServer()
		if err == nil {
			defer srv.Close()
			go srv.Run(&delegate, logf)
		} else {
			logf("ctl error: %s", err)
		}
	}

	inOut := editor.InOut{
		Reader: ttyFile,
		Writer: os.Stdout,
	}
	edit.Run(&inOut, logf)
}

type editDelegate struct {
	edit *editor.Editor
	logf logger.Func
}

func (ed *editDelegate) ExecuteCommand(cmd remotectl.CommandData) {
	var editorCommand editor.Command

	switch cmd.Action {
	case "open":
		editorCommand = &editor.OpenFile{Path: cmd.Args[0]}
	default:
		ed.logf("ignore remote cmd: %s", cmd.Action)
		return
	}

	ed.logf("remote cmd: %s", cmd.Action)
	ed.edit.Post(editorCommand)
}

func (ed *editDelegate) OpenFile(path string) {
	err := ed.sendOpenCommand(path)
	if err == nil {
		return
	}
	if errors.Is(err, os.ErrNotExist) {
		ed.logf("starting main editor")
		err := openMainEditor(path)
		if err != nil {
			ed.logf("error with main editor: %s", err)
		}
	}
}

func (ed *editDelegate) sendOpenCommand(path string) error {
	return remotectl.SendCommand(&remotectl.CommandData{
		Action: "open",
		Args:   []string{path},
	})
}
