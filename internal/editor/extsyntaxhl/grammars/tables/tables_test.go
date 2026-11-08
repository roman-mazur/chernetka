package tables

import (
	"bytes"
	"compress/zlib"
	"testing"
)

func compressed(t *testing.T, data []byte) []byte {
	t.Helper()
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	if _, err := w.Write(data); err != nil {
		t.Fatal(err)
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	return b.Bytes()
}

func TestInflate(t *testing.T) {
	stream := compressed(t, []byte{1, 2, 3, 4, 5})
	a, b := make([]byte, 2), make([]byte, 3)
	if err := Inflate(stream, [][]byte{a, b}); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(a, []byte{1, 2}) || !bytes.Equal(b, []byte{3, 4, 5}) {
		t.Errorf("inflated %v %v, want [1 2] [3 4 5]", a, b)
	}
}

func TestInflateMismatch(t *testing.T) {
	stream := compressed(t, []byte{1, 2, 3, 4, 5})
	corrupt := bytes.Clone(stream)
	corrupt[len(corrupt)-1] ^= 0xff // the checksum

	for name, tc := range map[string]struct {
		stream []byte
		sizes  []int
	}{
		"short":    {stream, []int{2, 4}},
		"long":     {stream, []int{2, 2}},
		"corrupt":  {corrupt, []int{5}},
		"not zlib": {[]byte("tables"), []int{5}},
	} {
		dst := make([][]byte, len(tc.sizes))
		for i, n := range tc.sizes {
			dst[i] = make([]byte, n)
		}
		if err := Inflate(tc.stream, dst); err == nil {
			t.Errorf("%s: Inflate succeeded", name)
		}
	}
}
