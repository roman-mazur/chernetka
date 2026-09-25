package main

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"image/color"
	"image/png"
	"io"
	"os"
	"strings"
	"sync"

	"rmazur.io/chernetka/internal/cheimg"
	"rmazur.io/chernetka/internal/editor/styles"
	"rmazur.io/chernetka/internal/logger"
	"rmazur.io/chernetka/internal/remotectl"
	"rmazur.io/chernetka/internal/vt"
	"rmazur.io/chernetka/internal/vt/escape"
)

// frame is a rendered item: an image, a text, or an error to display.
type frame struct {
	name string

	png        []byte
	imgW, imgH int // image size in pixels

	text string
	err  error
}

var (
	errNoImages    = errors.New("the terminal cannot display images")
	errUnsupported = errors.New("unsupported file type")
)

// pngFrame checks the image data and reads its size.
func pngFrame(data []byte) frame {
	cfg, err := png.DecodeConfig(bytes.NewReader(data))
	if err != nil {
		return frame{err: fmt.Errorf("not a PNG image: %w", err)}
	}
	return frame{png: data, imgW: cfg.Width, imgH: cfg.Height}
}

// viewer displays the items it's asked to show, one at a time.
type viewer struct {
	ctx    context.Context
	out    io.Writer
	size   func() vt.WindowSize
	images bool // whether the terminal supports images
	logf   logger.Func

	// renderDiagram draws a d2 diagram: as an image if the terminal supports them, or as text.
	renderDiagram func(ctx context.Context, it cheimg.Item) frame

	renderMu sync.Mutex // serializes rendering

	mu   sync.Mutex // guards the output and the last frame
	last *frame
}

// ExecuteCommand implements remotectl.Executor.
func (v *viewer) ExecuteCommand(cmd remotectl.CommandData) {
	it, err := cheimg.ParseCommand(cmd)
	if err != nil {
		v.logf("error parsing command: %s", err)
		return
	}
	v.show(it)
}

func (v *viewer) show(it cheimg.Item) {
	v.renderMu.Lock()
	defer v.renderMu.Unlock()

	v.logf("rendering %s", it.Name())
	v.mu.Lock()
	v.drawStatus(it.Name(), "rendering...")
	v.mu.Unlock()

	f := v.render(it)
	f.name = it.Name()
	if f.err != nil {
		v.logf("rendering %s failed: %s", f.name, f.err)
	}

	v.mu.Lock()
	defer v.mu.Unlock()
	v.last = &f
	v.draw()
}

func (v *viewer) render(it cheimg.Item) frame {
	switch it.Kind() {
	case cheimg.KindImage:
		if !v.images {
			return frame{err: errNoImages}
		}
		data, err := os.ReadFile(it.Path)
		if err != nil {
			return frame{err: err}
		}
		return pngFrame(data)
	case cheimg.KindDiagram:
		return v.renderDiagram(v.ctx, it)
	default:
		return frame{err: errUnsupported}
	}
}

// redraw displays the last frame again, for example, when the window size changes.
func (v *viewer) redraw() {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.draw()
}

var (
	headerStyle = styles.TextStyle{Bold: true}
	statusStyle = styles.TextStyle{TextColor: styles.DefaultColors.Suggestion}
	errorStyle  = styles.TextStyle{TextColor: color.RGBA{R: 0xF7, G: 0x54, B: 0x64, A: 0xff}}
)

// headerRows is the number of rows above the displayed item.
const headerRows = 2

func (v *viewer) drawStatus(name, status string) {
	out := bufio.NewWriter(v.out)
	defer out.Flush()
	escape.MoveTopLeft(out)
	escape.ClearLine(out)
	escape.StyleText(out, name, headerStyle)
	escape.StyleText(out, " "+status, statusStyle)
}

func (v *viewer) draw() {
	out := bufio.NewWriter(v.out)
	defer out.Flush()

	if v.images {
		escape.DeleteImages(out)
	}
	escape.ClearScreen(out)

	f := v.last
	if f == nil {
		escape.StyleText(out, "che-img", headerStyle)
		escape.StyleText(out, " waiting for images and diagrams... (q to quit)", statusStyle)
		return
	}

	escape.StyleText(out, f.name, headerStyle)
	_, _ = io.WriteString(out, strings.Repeat("\r\n", headerRows))

	switch {
	case f.err != nil:
		printLines(out, f.err.Error(), errorStyle)
	case f.png != nil:
		cols, rows := fitImage(f.imgW, f.imgH, v.size(), headerRows)
		escape.ShowPNG(out, bytes.NewReader(f.png), cols, rows)
	default:
		printLines(out, f.text, styles.TextStyle{})
	}
}

// printLines writes the text with line endings suitable for the terminal in raw mode.
func printLines(out io.Writer, text string, style styles.TextStyle) {
	for i, line := range strings.Split(text, "\n") {
		if i > 0 {
			_, _ = io.WriteString(out, "\r\n")
		}
		escape.StyleText(out, line, style)
	}
}

// fallbackCellSize is used when the terminal does not report its size in pixels.
var fallbackCellSize = [2]float64{10, 20}

// fitImage calculates the number of cells the image should occupy to fit the window below the reserved rows.
// Zero cols and rows mean the image fits in its original size.
// Otherwise, one of them is set to scale the image down preserving the aspect ratio.
func fitImage(imgW, imgH int, ws vt.WindowSize, reservedRows int) (cols, rows int) {
	availCols, availRows := ws.Cols, ws.Rows-reservedRows
	if imgW <= 0 || imgH <= 0 || availCols <= 0 || availRows <= 0 {
		return 0, 0
	}

	cellW, cellH, ok := ws.CellSize()
	if !ok {
		cellW, cellH = fallbackCellSize[0], fallbackCellSize[1]
	}

	scaleW := float64(availCols) * cellW / float64(imgW)
	scaleH := float64(availRows) * cellH / float64(imgH)
	if scaleW >= 1 && scaleH >= 1 {
		return 0, 0
	}
	if scaleW < scaleH {
		return availCols, 0
	}
	return 0, availRows
}
