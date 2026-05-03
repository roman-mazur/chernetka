package content

import (
	"os"
	"path"
	"strconv"
	"strings"
	"testing"
)

func TestLoadFullText(t *testing.T) {
	entries, err := os.ReadDir("testdata")
	must(err)

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
				ft, err := LoadFullText(f)
				must(err)
				if ft.Len() != eln {
					t.Errorf("wrong number of lines for %s: got %d, want %d", name, ft.Len(), eln)
				}
				if lines := ft.Lines(); len(lines) != eln {
					t.Log(lines)
					t.Error("bad lines")
				}
			})
		}
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
