package escape

import (
	"bytes"
	"encoding/base64"
	"errors"
	"io"
	"regexp"
	"strings"
	"testing"
	"testing/iotest"
)

func TestParseGraphicsQueryResponse(t *testing.T) {
	for _, tc := range []struct {
		name                string
		in                  string
		supported, complete bool
	}{
		{name: "nothing yet"},
		{name: "supported", in: "\x1b_Gi=31;OK\x1b\\\x1b[?62;22;52c", supported: true, complete: true},
		{name: "supported, waiting for attributes", in: "\x1b_Gi=31;OK\x1b\\", supported: true},
		{name: "not supported", in: "\x1b[?1;2c", complete: true},
		{name: "incomplete attributes", in: "\x1b[?1;2"},
		{name: "error", in: "\x1b_Gi=31;EINVAL:bad\x1b\\\x1b[?1;2c", complete: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			supported, done := ParseGraphicsQueryResponse([]byte(tc.in))
			if supported != tc.supported || done != tc.complete {
				t.Errorf("ParseGraphicsQueryResponse(%q) = %t, %t; want %t, %t",
					tc.in, supported, done, tc.supported, tc.complete)
			}
		})
	}
}

var graphicsRE = regexp.MustCompile(`\x1b_G([^;]*);([^\x1b]*)\x1b\\`)

// graphicsChunk is one escape sequence of the transmitted image.
type graphicsChunk struct {
	control, payload string
}

// parseGraphicsChunks checks that out consists of graphics escape sequences only and splits it.
func parseGraphicsChunks(t *testing.T, out string) []graphicsChunk {
	t.Helper()
	if rest := graphicsRE.ReplaceAllString(out, ""); rest != "" {
		t.Errorf("unexpected output besides graphics sequences: %q", rest)
	}
	var res []graphicsChunk
	for _, m := range graphicsRE.FindAllStringSubmatch(out, -1) {
		res = append(res, graphicsChunk{control: m[1], payload: m[2]})
	}
	return res
}

// checkContinuation verifies the chunks following the first one.
func checkContinuation(t *testing.T, chunks []graphicsChunk) {
	t.Helper()
	for i, chunk := range chunks {
		if len(chunk.payload) > graphicsChunkSize {
			t.Errorf("chunk %d is too big: %d", i, len(chunk.payload))
		}
		if len(chunk.payload)%4 != 0 {
			t.Errorf("chunk %d payload size %d is not a multiple of 4", i, len(chunk.payload))
		}
		if i == 0 {
			continue
		}
		wantControl := "m=1"
		if i == len(chunks)-1 {
			wantControl = "m=0"
		}
		if chunk.control != wantControl {
			t.Errorf("chunk %d control = %q, want %q", i, chunk.control, wantControl)
		}
	}
}

func decodePayload(t *testing.T, chunks []graphicsChunk) []byte {
	t.Helper()
	var payload strings.Builder
	for _, chunk := range chunks {
		payload.WriteString(chunk.payload)
	}
	decoded, err := base64.StdEncoding.DecodeString(payload.String())
	if err != nil {
		t.Fatal(err)
	}
	return decoded
}

func TestShowPNG(t *testing.T) {
	const rawChunkSize = graphicsChunkSize / 4 * 3 // image bytes encoded into one chunk

	readers := map[string]func([]byte) io.Reader{
		"whole":    func(b []byte) io.Reader { return bytes.NewReader(b) },
		"one byte": func(b []byte) io.Reader { return iotest.OneByteReader(bytes.NewReader(b)) },
		"halves":   func(b []byte) io.Reader { return iotest.HalfReader(bytes.NewReader(b)) },
	}

	for _, tc := range []struct {
		name        string
		size        int
		cols, rows  int
		wantControl string
		wantChunks  int
	}{
		{name: "small", size: 10, wantControl: "a=T,f=100,q=2,m=0", wantChunks: 1},
		{name: "empty", size: 0, wantControl: "a=T,f=100,q=2,m=0", wantChunks: 1},
		{name: "columns", size: 10, cols: 20, wantControl: "a=T,f=100,q=2,c=20,m=0", wantChunks: 1},
		{name: "rows", size: 10, rows: 5, wantControl: "a=T,f=100,q=2,r=5,m=0", wantChunks: 1},
		{name: "columns and rows", size: 10, cols: 1234, rows: 567, wantControl: "a=T,f=100,q=2,c=1234,r=567,m=0", wantChunks: 1},
		{name: "exactly one chunk", size: rawChunkSize, wantControl: "a=T,f=100,q=2,m=0", wantChunks: 1},
		{name: "one byte more", size: rawChunkSize + 1, wantControl: "a=T,f=100,q=2,m=1", wantChunks: 2},
		{name: "exactly two chunks", size: 2 * rawChunkSize, cols: 3, wantControl: "a=T,f=100,q=2,c=3,m=1", wantChunks: 2},
		{name: "chunked", size: 10000, wantControl: "a=T,f=100,q=2,m=1", wantChunks: 4},
	} {
		for readerName, newReader := range readers {
			t.Run(tc.name+"/"+readerName, func(t *testing.T) {
				data := make([]byte, tc.size)
				for i := range data {
					data[i] = byte(i * 7)
				}
				var out bytes.Buffer
				ShowPNG(&out, newReader(data), tc.cols, tc.rows)

				chunks := parseGraphicsChunks(t, out.String())
				if len(chunks) != tc.wantChunks {
					t.Fatalf("got %d chunks, want %d", len(chunks), tc.wantChunks)
				}
				if chunks[0].control != tc.wantControl {
					t.Errorf("first control data = %q, want %q", chunks[0].control, tc.wantControl)
				}
				checkContinuation(t, chunks)
				if !bytes.Equal(decodePayload(t, chunks), data) {
					t.Error("transmitted data does not match the image")
				}
			})
		}
	}
}

// TestShowPNG_ReadError checks that the escape sequence is terminated even if the image cannot be read fully,
// so the terminal does not treat the following output as the image data.
func TestShowPNG_ReadError(t *testing.T) {
	data := bytes.Repeat([]byte{0xCD}, 5000)
	in := io.MultiReader(bytes.NewReader(data), iotest.ErrReader(errors.New("broken")))

	var out bytes.Buffer
	ShowPNG(&out, in, 0, 0)

	chunks := parseGraphicsChunks(t, out.String())
	if len(chunks) != 2 {
		t.Fatalf("got %d chunks, want 2", len(chunks))
	}
	checkContinuation(t, chunks)
	if !bytes.Equal(decodePayload(t, chunks), data) {
		t.Error("transmitted data does not match the data read before the error")
	}
}
