package main

import (
	"bytes"
	"context"
	"errors"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"rmazur.io/chernetka/internal/cheimg"
	"rmazur.io/chernetka/internal/remotectl"
	"rmazur.io/chernetka/internal/vt"
	"rmazur.io/chernetka/internal/vt/escape"
)

func TestFitImage(t *testing.T) {
	// 100x50 cells of 10x20 pixels: 1000x1000 pixels, 1000x960 available below 2 reserved rows.
	ws := vt.WindowSize{Cols: 100, Rows: 50, XPixel: 1000, YPixel: 1000}
	noPixels := vt.WindowSize{Cols: 100, Rows: 50}
	for _, tc := range []struct {
		name       string
		w, h       int
		ws         vt.WindowSize
		cols, rows int
	}{
		{name: "small image", w: 300, h: 200, ws: ws},
		{name: "exact fit", w: 1000, h: 960, ws: ws},
		{name: "too wide", w: 2000, h: 500, ws: ws, cols: 100},
		{name: "too tall", w: 500, h: 2000, ws: ws, rows: 48},
		{name: "too big, width limits", w: 4000, h: 2000, ws: ws, cols: 100},
		{name: "too big, height limits", w: 2000, h: 4000, ws: ws, rows: 48},
		{name: "unknown pixel size, fits", w: 100, h: 100, ws: noPixels},
		{name: "unknown pixel size, too big", w: 5000, h: 100, ws: noPixels, cols: 100},
		{name: "no space", w: 100, h: 100, ws: vt.WindowSize{Cols: 100, Rows: 2}},
		{name: "empty image", ws: ws},
	} {
		t.Run(tc.name, func(t *testing.T) {
			cols, rows := fitImage(tc.w, tc.h, tc.ws, 2)
			if cols != tc.cols || rows != tc.rows {
				t.Errorf("fitImage(%d, %d) = %d, %d; want %d, %d", tc.w, tc.h, cols, rows, tc.cols, tc.rows)
			}
		})
	}
}

func TestDetectGraphics(t *testing.T) {
	for _, tc := range []struct {
		name     string
		response []string
		want     bool
	}{
		{name: "supported", response: []string{"\x1b_Gi=31;OK\x1b\\", "\x1b[?62;4c"}, want: true},
		{name: "supported, one chunk", response: []string{"\x1b_Gi=31;OK\x1b\\\x1b[?62;4c"}, want: true},
		{name: "not supported", response: []string{"\x1b[?1;2c"}},
		{name: "no response"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			input := make(chan []byte, len(tc.response))
			for _, chunk := range tc.response {
				input <- []byte(chunk)
			}
			var out bytes.Buffer
			if got := detectGraphics(&out, input, 50*time.Millisecond); got != tc.want {
				t.Errorf("detectGraphics() = %t, want %t", got, tc.want)
			}
			if out.String() != escape.GraphicsQuery {
				t.Errorf("sent %q, want the graphics query", out.String())
			}
		})
	}
}

func TestReadInput(t *testing.T) {
	var got []byte
	for chunk := range readInput(strings.NewReader("hello")) {
		got = append(got, chunk...)
	}
	if string(got) != "hello" {
		t.Errorf("read %q", got)
	}
}

func TestIsQuitInput(t *testing.T) {
	for in, want := range map[string]bool{"q": true, "\x03": true, "x": false, "qq": false, "": false} {
		if got := isQuitInput([]byte(in)); got != want {
			t.Errorf("isQuitInput(%q) = %t", in, got)
		}
	}
}

// newTestViewer returns a viewer that "renders" a diagram into its source text or a tall image.
func newTestViewer(images bool) (*viewer, *bytes.Buffer, *[]cheimg.Item) {
	var (
		out      bytes.Buffer
		rendered []cheimg.Item
	)
	return &viewer{
		ctx:    context.Background(),
		out:    &out,
		size:   func() vt.WindowSize { return vt.WindowSize{Cols: 80, Rows: 24} },
		images: images,
		logf:   func(string, ...any) {},
		renderDiagram: func(_ context.Context, it cheimg.Item) frame {
			rendered = append(rendered, it)
			switch {
			case it.Source == "bad":
				return frame{err: errors.New("bad diagram\nline 2")}
			case images:
				return pngFrame(encodePNG(10, 500))
			default:
				return frame{text: it.Source}
			}
		},
	}, &out, &rendered
}

func encodePNG(w, h int) []byte {
	var buf bytes.Buffer
	if err := png.Encode(&buf, image.NewGray(image.Rect(0, 0, w, h))); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

// screen returns the text displayed after the last screen clear.
func screen(out *bytes.Buffer) string {
	got := out.String()
	return escape.Clean(got[max(0, strings.LastIndex(got, "\x1b[2J")):])
}

func TestViewer_ExecuteCommand(t *testing.T) {
	v, out, rendered := newTestViewer(false)

	v.ExecuteCommand(*cheimg.Item{Source: "a -> b\nb -> c", Path: "/notes/README.md"}.Command())
	if len(*rendered) != 1 {
		t.Fatalf("rendered %d diagrams", len(*rendered))
	}
	for _, want := range []string{"/notes/README.md", "a -> b\r\nb -> c"} {
		if !strings.Contains(screen(out), want) {
			t.Errorf("screen does not contain %q", want)
		}
	}
	if strings.Contains(out.String(), "\x1b_G") {
		t.Error("graphics sequences sent to a terminal without images support")
	}

	out.Reset()
	v.ExecuteCommand(remotectl.CommandData{Action: "unknown"})
	if len(*rendered) != 1 || out.Len() != 0 {
		t.Error("bad command was not ignored")
	}

	v.ExecuteCommand(*cheimg.Item{Source: "bad"}.Command())
	if got := screen(out); !strings.Contains(got, "bad diagram\r\nline 2") {
		t.Errorf("error is not displayed: %q", got)
	}
}

func TestViewer_Diagram(t *testing.T) {
	v, out, _ := newTestViewer(true)

	v.redraw()
	if got := escape.Clean(out.String()); !strings.Contains(got, "waiting for images and diagrams") {
		t.Errorf("initial screen: %q", got)
	}

	out.Reset()
	v.show(cheimg.Item{Path: "/tmp/a.d2"})
	got := out.String()
	if !strings.Contains(got, "\x1b_Ga=d,q=2\x1b\\") {
		t.Error("previous images are not deleted")
	}
	// The image is too tall for the window: it's scaled to the rows below the header.
	if !strings.Contains(got, "\x1b_Ga=T,f=100,q=2,r=22,m=0;") {
		t.Errorf("image is not displayed: %q", got)
	}

	out.Reset()
	v.redraw()
	if !strings.Contains(out.String(), "\x1b_Ga=T") {
		t.Error("image is not displayed on redraw")
	}
}

func TestViewer_PNG(t *testing.T) {
	dir := t.TempDir()
	imagePath := filepath.Join(dir, "small.png")
	if err := os.WriteFile(imagePath, encodePNG(40, 30), 0o600); err != nil {
		t.Fatal(err)
	}
	notImagePath := filepath.Join(dir, "text.png")
	if err := os.WriteFile(notImagePath, []byte("hello"), 0o600); err != nil {
		t.Fatal(err)
	}

	t.Run("displayed", func(t *testing.T) {
		v, out, rendered := newTestViewer(true)
		v.ExecuteCommand(*cheimg.Item{Path: imagePath}.Command())
		if len(*rendered) != 0 {
			t.Error("image is rendered as a diagram")
		}
		// The image fits the window: it's displayed in its original size.
		if !strings.Contains(out.String(), "\x1b_Ga=T,f=100,q=2,m=0;") {
			t.Errorf("image is not displayed: %q", out.String())
		}
		if !strings.Contains(screen(out), imagePath) {
			t.Error("image name is not displayed")
		}
	})

	for _, tc := range []struct {
		name    string
		images  bool
		path    string
		wantErr string
	}{
		{name: "no images support", path: imagePath, wantErr: errNoImages.Error()},
		{name: "not a PNG", images: true, path: notImagePath, wantErr: "not a PNG image"},
		{name: "missing file", images: true, path: filepath.Join(dir, "missing.png"), wantErr: "no such file"},
		{name: "unsupported", images: true, path: filepath.Join(dir, "photo.jpg"), wantErr: errUnsupported.Error()},
	} {
		t.Run(tc.name, func(t *testing.T) {
			v, out, _ := newTestViewer(tc.images)
			v.show(cheimg.Item{Path: tc.path})
			if strings.Contains(out.String(), "\x1b_Ga=T") {
				t.Error("image is displayed")
			}
			if got := screen(out); !strings.Contains(got, tc.wantErr) {
				t.Errorf("screen %q does not contain the error %q", got, tc.wantErr)
			}
		})
	}
}
