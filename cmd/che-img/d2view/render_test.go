package d2view

import (
	"bytes"
	"context"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"rmazur.io/chernetka/internal/cheimg"
)

func TestASCII(t *testing.T) {
	out, err := ASCII(context.Background(), cheimg.Item{Source: "alpha -> beta: link"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("\n" + out)
	for _, want := range []string{"alpha", "beta", "link", "┌"} {
		if !strings.Contains(out, want) {
			t.Errorf("ASCII output does not contain %q", want)
		}
	}
}

func TestASCII_FileWithImport(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "shared.d2"), "beta: Imported Beta")
	path := filepath.Join(dir, "main.d2")
	writeFile(t, path, "...@shared\nalpha -> beta")

	for _, d := range []cheimg.Item{
		{Path: path},
		// Unsaved source of the file.
		{Path: path, Source: "...@shared\ngamma -> beta"},
	} {
		out, err := ASCII(context.Background(), d)
		if err != nil {
			t.Fatalf("ASCII(%+v): %s", d, err)
		}
		if !strings.Contains(out, "Imported Beta") {
			t.Errorf("ASCII(%+v) does not contain the imported label:\n%s", d, out)
		}
	}
}

func TestASCII_Errors(t *testing.T) {
	for _, d := range []cheimg.Item{
		{Path: filepath.Join(t.TempDir(), "missing.d2")},
		{Source: "a -> {"},
	} {
		if _, err := ASCII(context.Background(), d); err == nil {
			t.Errorf("ASCII(%+v) did not fail", d)
		}
	}
}

func TestRasterizer_PNG(t *testing.T) {
	var r Rasterizer
	r.Options = DefaultOptions
	defer r.Close()

	data, err := r.PNG(context.Background(), cheimg.Item{Source: "alpha -> beta: link"})
	if err != nil {
		t.Fatal(err)
	}
	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		t.Fatal("result is not a PNG:", err)
	}
	if size := img.Bounds().Size(); size.X < 100 || size.Y < 200 {
		t.Errorf("image is too small: %s", size)
	}
	if out := os.Getenv("D2VIEW_OUT"); out != "" {
		writeFile(t, out, string(data)) // for inspection by eye
	}

	if _, err := r.PNG(context.Background(), cheimg.Item{Source: "a -> {"}); err == nil {
		t.Error("PNG of an invalid diagram did not fail")
	}
}

func TestReplaceFonts(t *testing.T) {
	in := `.text{font-family:"d2-1-font-regular"}.text-bold{font-family:d2-1-font-bold}` +
		`.text-mono-italic{font-family:"d2-1-font-mono-italic"}.text-semi{font-family:d2-1-font-semibold}`
	want := `.text{font-family:"Source Sans Pro"}.text-bold{font-family:"Source Sans Pro"; font-weight: bold}` +
		`.text-mono-italic{font-family:"Source Code Pro"; font-style: italic}.text-semi{font-family:"Source Sans Pro"; font-weight: 600}`
	if got := string(replaceFonts([]byte(in))); got != want {
		t.Errorf("replaceFonts:\n got %s\nwant %s", got, want)
	}
}

func writeFile(t *testing.T, path, data string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(data), 0644); err != nil {
		t.Fatal(err)
	}
}
