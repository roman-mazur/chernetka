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
				for _, line := range ft.Lines() {
					if strings.HasSuffix(line.String(), "\n") {
						t.Fatal("\\n detected")
					}
				}
			})
		}
	}
}

func TestFullText_Mutate(t *testing.T) {
	data := FullText{
		TextLine("a"),
		TextLine("b"),
		TextLine("c"),
	}
	data.Update(1, TextLine("d"))
	data.Insert(0, TextLine("0"))
	data.Delete(1)
	data.Insert(3, TextLine("e"))

	var buf strings.Builder
	Print(&buf, &data, 0, data.Len())
	if res := buf.String(); res != "0\r\nd\r\nc\r\ne\r\n" {
		t.Log(res)
		t.Error("mutations failed")
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
