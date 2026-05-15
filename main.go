package main

import (
	"flag"
	"log"
	"os"

	"rmazur.io/x/edit/internal/editor"
	"rmazur.io/x/edit/internal/logger"
)

func main() {
	flag.Parse()
	var (
		ttyFile = os.Stdin
		edit    editor.Editor
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
		} else {
			var f *os.File
			f, err = os.Open(path)
			if err != nil {
				log.Fatal("cannot open the specified file:", err)
			}
			defer f.Close()
			err = edit.OpenReader(path, f)
		}
		if err != nil {
			log.Fatal("cannot read the specified path:", err)
		}
	}

	logf, _ := logger.UserLogFile()
	edit.Run(ttyFile, logf)
}
