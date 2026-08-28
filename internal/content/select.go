package content

import "bytes"

// Select returns the text delimited by the given spans, concatenated in order.
func Select(c Interface, spans []Span) string {
	if len(spans) == 0 {
		return ""
	}

	lines := c.Lines()
	var out bytes.Buffer
	for _, span := range spans {
		start, stop := span.Min(), span.Max()
		for i := start.Line; i <= stop.Line; i++ {
			line := lines[i].String()
			j, k := start.Col, stop.Col
			if i > start.Line {
				j = 0
				out.WriteByte('\n')
			}
			if i < stop.Line {
				k = len(line)
			}
			out.WriteString(line[j:k])
		}
	}
	return out.String()
}
