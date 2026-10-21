package content

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"
	"testing/iotest"
)

func TestLoadFullText(t *testing.T) {
	entries, err := os.ReadDir("testdata")
	must(t, err)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if strings.Contains(name, "fulltext") {
			parts := strings.Split(name, "-")
			last := parts[len(parts)-1]
			expectedLines := strings.TrimSuffix(last, path.Ext(last))
			eln, err := strconv.Atoi(expectedLines)
			if err != nil {
				t.Fatalf("cannot get expected lines number for %s: %s", name, err)
			}

			t.Run(name, func(t *testing.T) {
				f, err := os.Open("testdata/" + name)
				t.Cleanup(func() {
					_ = f.Close()
				})
				ft, _, err := LoadFullText(f)
				must(t, err)
				if ft.Len() != eln {
					t.Errorf("wrong number of lines for %s: got %d, want %d", name, ft.Len(), eln)
				}
				if lines := ft.Lines(); len(lines) != eln {
					t.Log(lines)
					t.Error("bad lines")
				}
				for _, line := range ft.Lines() {
					if strings.HasSuffix(line.String(), "\n") {
						t.Fatal("\\n detected")
					}
				}
			})
		}
	}
}

func TestLoadFullText_Fingerprint(t *testing.T) {
	sha := func(data string) string {
		sum := sha256.Sum256([]byte(data))
		return hex.EncodeToString(sum[:])
	}
	load := func(t *testing.T, data string) string {
		t.Helper()
		_, fingerprint, err := LoadFullText(strings.NewReader(data))
		must(t, err)
		return fingerprint
	}

	for _, data := range []string{
		"",
		"one",
		"one\ntwo\n",
		"one\r\ntwo",
		strings.Repeat("a long line ", 1000) + "\n" + strings.Repeat("x\n", 10000), // Over the bufio buffer.
	} {
		if got, want := load(t, data), sha(data); got != want {
			t.Errorf("fingerprint of %.20q is %s, want the hash of the read data %s", data, got, want)
		}
	}

	t.Run("same text", func(t *testing.T) {
		if load(t, "one\ntwo") != load(t, "one\ntwo") {
			t.Error("the same text has different fingerprints")
		}
	})

	// The files differ, even when only in the line endings that are not part of the lines.
	for _, pair := range [][2]string{
		{"one\ntwo", "one\r\ntwo"},
		{"one\ntwo", "one\ntwo\n"},
		{"one", "two"},
	} {
		t.Run(fmt.Sprintf("%q and %q", pair[0], pair[1]), func(t *testing.T) {
			if load(t, pair[0]) == load(t, pair[1]) {
				t.Error("the different files have the same fingerprint")
			}
		})
	}

	t.Run("read error", func(t *testing.T) {
		readErr := errors.New("read error")
		doc, fingerprint, err := LoadFullText(io.MultiReader(strings.NewReader("one\n"), iotest.ErrReader(readErr)))
		if !errors.Is(err, readErr) || doc != nil || fingerprint != "" {
			t.Errorf("got %v, %q, %v; want no document and fingerprint with the read error", doc, fingerprint, err)
		}
	})
}

func TestFullText_Mutate(t *testing.T) {
	data := FullText{
		TextLine("a"),
		TextLine("b"),
		TextLine("c"),
	}
	data.Update(1, TextLine("d"))
	data.Insert(0, TextLine("0"))
	data.Delete(1, 2)
	data.Insert(3, TextLine("e"))

	var buf strings.Builder
	Print(&buf, &data, 0, data.Len())
	if res := buf.String(); res != "0\r\nd\r\nc\r\ne\r\n" {
		t.Log(res)
		t.Error("mutations failed")
	}
}

func TestEmpty(t *testing.T) {
	c := Empty()
	if c.Len() != 1 {
		t.Error("wrong length", c.Len())
	}
	t.Logf("content: %q", c.Lines()[0].String())
}

func must(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
