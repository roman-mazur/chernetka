package main

import (
	"flag"
	"log"
	"os"

	"rmazur.io/x/edit/internal/editor"
	"rmazur.io/x/edit/internal/logger"
	"rmazur.io/x/edit/internal/remotectl"
)

func main() {
	flag.Parse()
	var (
		ttyFile = os.Stdin
		edit    editor.Editor
		skipCtl bool
	)

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
			edit.OpenDir(path)
			skipCtl = true
		} else {
			(&editor.OpenFile{Path: path}).DoOnEditor(&edit)
		}
	}

	logf, _ := logger.UserLogFile()

	if !skipCtl {
		srv, err := remotectl.NewServer()
		if err == nil {
			defer srv.Close()
			go srv.Run(&remoteCmdExecutor{edit: &edit, logf: logf}, logf)
		} else {
			logf("ctl error: %s", err)
		}
	}

	edit.Run(ttyFile, logf)
}

type remoteCmdExecutor struct {
	edit *editor.Editor
	logf logger.Func
}

func (rce *remoteCmdExecutor) ExecuteCommand(cmd remotectl.CommandData) {
	var editorCommand editor.Command

	switch cmd.Action {
	case "open":
		editorCommand = &editor.OpenFile{Path: cmd.Args[0]}
	default:
		rce.logf("ignore remote cmd: %s", cmd.Action)
		return
	}

	rce.logf("remote cmd: %s", cmd.Action)
	rce.edit.Post(editorCommand)
}
