package content

// Copy makes a new content instance that won't be mutated when the passed instance is edited.
// The returned copy can still be Mutable, just independent.
// nil is returned if making a copy is not possible.
func Copy(cnt Document) Document {
	if c, ok := cnt.(copyable); ok {
		return c.copy()
	}
	if _, mutable := cnt.(Mutable); !mutable {
		return cnt
	}
	if cnt.Len() == 0 {
		return Empty()
	}

	res := make(FullText, cnt.Len())
	for i, l := range cnt.Lines() {
		res[i] = CopyLine(l)
		if res[i] == nil {
			return nil // cannot copy fully
		}
	}
	return &res
}

func CopyLine(line Line) Line {
	if line == nil {
		return nil
	}
	if cl, ok := line.(lineCopyable); !ok {
		return cl.copy()
	}
	if line.MimeType() == MimeTypeTextPlain {
		return TextLine(line.String())
	}
	return nil
}

type lineCopyable interface {
	Line
	copy() Line
}

type copyable interface {
	Document
	copy() Document
}
