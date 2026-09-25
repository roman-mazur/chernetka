// Command che-img displays images and d2 diagrams (https://d2lang.com) in the terminal.
//
// Usage:
//
//	che-img [flags] [file.png|file.d2]
//
// che-img shows the given file (if any) and listens on a local socket for the next items to show.
// The che editor sends them there when it's asked to open a PNG image or a d2 diagram, when you
// press Enter on the first line of a .d2 file, or on the line opening a ```d2 code block in
// a Markdown file. If che-img is not running, che launches it in a new terminal pane.
//
// Terminals supporting the kitty graphics protocol (kitty, Ghostty, WezTerm, and others)
// display images, and diagrams are rendered as images too. Otherwise, only diagrams can be shown:
// they are drawn with text characters.
//
// Press q to quit.
package main

import (
	"context"
	"flag"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	d2log "oss.terrastruct.com/d2/lib/log"

	"rmazur.io/chernetka/cmd/che-img/d2view"
	"rmazur.io/chernetka/internal/cheimg"
	"rmazur.io/chernetka/internal/logger"
	"rmazur.io/chernetka/internal/remotectl"
	"rmazur.io/chernetka/internal/vt"
)

var (
	asciiFlag = flag.Bool("ascii", false, "draw diagrams with text characters even if the terminal can display images")
	themeFlag = flag.Int64("theme", d2view.DefaultOptions.ThemeID, "d2 theme ID used for diagram images")
)

func main() {
	flag.Parse()

	logf, _ := logger.UserLogFile()
	logf = logger.Prefix(logf, "che-img: ")

	srv, err := remotectl.NewServer(cheimg.Endpoint)
	if err != nil {
		log.Fatalf("cannot listen for images (is che-img already running?): %s", err)
	}
	defer srv.Close()

	term, err := setupTerminal()
	if err != nil {
		log.Fatalf("cannot get system terminal: %s", err)
	}
	defer logf.CloseAndLog(term, "terminal")

	input := readInput(term)

	rast := &d2view.Rasterizer{Options: d2view.DefaultOptions}
	rast.ThemeID = *themeFlag
	defer rast.Close()

	v := &viewer{
		// Route d2 logs to the log file: the terminal shows diagrams.
		ctx:           d2log.With(context.Background(), slog.New(slog.NewTextHandler(logger.Writer(logf), nil))),
		out:           term,
		size:          func() vt.WindowSize { return terminalSize(term) },
		logf:          logf,
		renderDiagram: renderASCII,
	}
	if detectGraphics(term, input, time.Second) {
		v.images = true
		if !*asciiFlag {
			v.renderDiagram = func(ctx context.Context, it cheimg.Item) frame { return renderImage(ctx, rast, it) }
		}
	}
	logf("started, images: %t", v.images)

	go srv.Run(v, logf)

	if flag.NArg() > 0 {
		path, _ := filepath.Abs(flag.Arg(0))
		go v.show(cheimg.Item{Path: path})
	} else {
		v.redraw()
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGHUP)
	winch := term.WindowSizeChanges()
	for {
		select {
		case in, ok := <-input:
			if !ok {
				input = nil // The terminal is closed: wait for a signal.
			} else if isQuitInput(in) {
				return
			}
		case <-winch:
			v.redraw()
		case <-signals:
			return
		}
	}
}

// terminalSize returns the terminal size falling back to the classic 80x24.
func terminalSize(term vt.Terminal) vt.WindowSize {
	size, err := term.Size()
	if err != nil || size.Cols == 0 || size.Rows == 0 {
		return vt.WindowSize{Cols: 80, Rows: 24}
	}
	return size
}

func renderASCII(ctx context.Context, it cheimg.Item) frame {
	text, err := d2view.ASCII(ctx, it)
	return frame{text: text, err: err}
}

func renderImage(ctx context.Context, rast *d2view.Rasterizer, it cheimg.Item) frame {
	data, err := rast.PNG(ctx, it)
	if err != nil {
		return frame{err: err}
	}
	return pngFrame(data)
}
