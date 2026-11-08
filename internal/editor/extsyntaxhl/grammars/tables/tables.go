// Package tables inflates the parse tables of the grammars into their arrays.
//
// fetchgrammar takes the parse tables out of a grammar's parser.c and stores
// them compressed, one after another, as a zlib stream. The arrays left in
// their place are zeroed until Inflate fills them.
package tables

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
)

// Inflate decompresses the zlib stream into the tables, in order. The stream
// must fill them exactly: anything short, longer or corrupt is an error, as
// the parser would read garbage. zlib checks the stream with its checksum.
func Inflate(compressed []byte, tables [][]byte) error {
	r, err := zlib.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return fmt.Errorf("parse tables: %w", err)
	}
	defer r.Close()
	for i, t := range tables {
		if _, err := io.ReadFull(r, t); err != nil {
			return fmt.Errorf("parse table %d of %d bytes: %w", i, len(t), err)
		}
	}
	// The checksum is verified at the end of the stream, so reading up to it
	// also tells a corrupt stream apart.
	switch n, err := io.Copy(io.Discard, r); {
	case err != nil:
		return fmt.Errorf("parse tables: %w", err)
	case n > 0:
		return fmt.Errorf("parse tables: the stream is %d bytes longer than the tables", n)
	}
	return nil
}
