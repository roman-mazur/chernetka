package content

import (
	"bytes"
	"testing"
)

func TestCopy(t *testing.T) {
	for _, tc := range []struct {
		src      Document
		identity bool
	}{
		{src: Empty()},
		{src: &FullText{TextLine("test")}},
		{src: nil, identity: true},
	} {
		var srcOut bytes.Buffer
		if err := SaveToWriter(tc.src, &srcOut); err != nil {
			t.Fatal(err)
		}
		res := Copy(tc.src)
		var resOut bytes.Buffer
		if err := SaveToWriter(res, &resOut); err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(srcOut.Bytes(), resOut.Bytes()) {
			t.Errorf("copied content does not match original: %T", tc.src)
		}
		if tc.identity && res != tc.src {
			t.Errorf("content identity is not maintained: %T", tc.src)
		}
	}
}
