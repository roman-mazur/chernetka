# SIMD packages of Go 1.27 for rendering

Can the `simd` and `simd/archsimd` packages, presented on the Go blog
([archsimd](https://go.dev/blog/archsimd), [simd-experiment](https://go.dev/blog/simd-experiment)),
make the editor render faster?

**No, not in a way that matters now.** Rendering is not bound by the kind of work SIMD speeds up.

## The packages

- `simd/archsimd`: the low-level layer with fixed-width hardware types (`Int8x16`, `Float32x8`, ...).
  Experimental on amd64 since Go 1.26; Go 1.27 adds arm64 and wasm.
  CPU features are checked at runtime, like `archsimd.X86.AVX2()`.
- `simd`: new in Go 1.27, portable and vector-size-agnostic (`Uint8s`, `Float32s`, mask types).
  It has `Load…`/`Load…Part`/`Store…`, element-wise arithmetic, comparisons producing masks,
  and is emulated in pure Go when the hardware has no support. `ToArch()` converts to `archsimd`.
- Both exist only with `GOEXPERIMENT=simd` and are not covered by the Go 1 compatibility promise.

The notes are based on the package docs of the go1.27.0 toolchain: the blog was not reachable from
the environment the research was done in.

## Where a frame spends time

`BenchmarkEditor_Render` on `loop.go` (Intel Xeon, AVX-512):

| Extensions | Time per frame |
|------------|----------------|
| none       | ~9 µs          |
| syntaxhl   | ~20–27 µs      |
| all of che | ~20–30 µs      |

A 60 fps frame is 16 ms, so the render takes 500–1000× less than that.

The CPU profile:

- ~30%: `io.WriteString`, `bufio.(*Writer).WriteString`, and interface dispatch of many small writes.
- ~27%: `escape.StyleTextSet` and `write8bitColor` building the SGR sequences.
- `TextStyle` struct comparisons, `slices.BinarySearchFunc` over the syntax spans, the map of diagnostics.
- The only byte loops SIMD fits are small: `runeToScreenCol` ~2%, `utf8.RuneCountInString` ~1%,
  `nlDigitsLen` ~1.5%. Splitting at tabs goes through `strings.Index`/`IndexByte`,
  which are already vectorized in the standard library.

## Experiment: the screen columns of a line

`runeToScreenCol(line, len(line), tabSize)` rewritten with the portable package, counting the UTF-8
lead bytes and the tabs with masks. It matches the scalar version on the tested lines.

```go
//go:build goexperiment.simd

func simdCols(line string, tabSize int) int {
	b := unsafe.Slice(unsafe.StringData(line), len(line))
	contMask := simd.BroadcastUint8s(0xC0)
	cont := simd.BroadcastUint8s(0x80)
	tab := simd.BroadcastUint8s('\t')
	var runes, tabs simd.Int8s // negative counters: a true mask lane is -1
	n := runes.Len()
	var buf [64]int8
	total, totalTabs, chunks := 0, 0, 0
	flush := func() {
		runes.Store(buf[:])
		for _, v := range buf[:n] {
			total -= int(v)
		}
		tabs.Store(buf[:])
		for _, v := range buf[:n] {
			totalTabs -= int(v)
		}
		runes, tabs, chunks = simd.Int8s{}, simd.Int8s{}, 0
	}
	for len(b) > 0 {
		v, k := simd.LoadUint8sPart(b)
		b = b[k:]
		lead := v.And(contMask).NotEqual(cont)
		if k < n { // the zero padding counts as lead bytes
			total -= n - k
		}
		runes = runes.Add(lead.ToInt8s())
		tabs = tabs.Add(v.Equal(tab).ToInt8s())
		if chunks++; chunks == 127 {
			flush()
		}
	}
	flush()
	return total + totalTabs*(tabSize-1)
}
```

| Line  | Scalar | SIMD                 |
|-------|--------|----------------------|
| 38 B  | 32 ns  | 57 ns (slower)       |
| 80 B  | 85 ns  | 58 ns                |
| 1 KB  | 970 ns | 87 ns (11× faster)   |

Code lines are short, so the gain is small or negative: at most ~1 µs for a frame.
Only lines longer than 1 KB, like minified JSON, gain clearly.

## Why not adopt now

- `go install rmazur.io/chernetka/cmd/che@latest` does not set `GOEXPERIMENT=simd`:
  a scalar fallback under `//go:build !goexperiment.simd` is needed, and most users would run only it.
- The API is experimental and may change with any Go release.
- The image viewer is not a better fit: base64 in `escape.ShowPNG` is small next to resvg
  rasterizing in WASM.

## What would make the frames faster

- Fewer and bigger writes: a concrete `*bufio.Writer` or a `[]byte` buffer instead of `io.Writer`.
- Cached escape sequences for every `TextStyle`.
- Redrawing only the changed lines, which also cuts what the terminal parses, likely the cost
  the user feels the most.

Revisit the columns counter if the experiment becomes stable and long lines are a problem.
