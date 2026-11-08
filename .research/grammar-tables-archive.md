# One archive for the grammar parse tables

fetchgrammar stores the parse tables of each grammar in its own `tables.zlib`, inflated when the
language is first asked for. Could `cmd/che` get smaller by putting all grammars into one archive,
for example with FastCDC chunks shared between grammars, while still loading each grammar on demand?

**No.** The grammars have almost no bytes in common, so combining them saves about 1%. The codec matters
much more than combining: LZMA2 roughly halves the payload, and it keeps the per-grammar streams.
None of this is worth doing now.

## Where the bytes are

Measured with 16 grammars (bash … yaml) at 8274843, Go 1.27, darwin/amd64:

- `che` is 15.7 MB, 10.9 MB stripped.
- The `tables.zlib` files total **1.60 MB** and inflate to **19.6 MB**. The inflated arrays are in BSS,
  so they take no space in the file.
- The grammar code left in `parser.c` costs about as much: `ts_lex` + `ts_lex_keywords` are ~0.74 MB.
  It's code, so no archive helps with it.

LZMA2 bytes by table, summed over the grammars:

| Table                      | Raw     | zlib   | LZMA2  |
|----------------------------|---------|--------|--------|
| `ts_parse_table`           | 14.0 MB | 686 KB | 292 KB |
| `ts_small_parse_table`     | 4.1 MB  | 458 KB | 317 KB |
| `ts_parse_actions`         | 933 KB  | 167 KB | 114 KB |
| `ts_primary_state_ids`     | 125 KB  | 78 KB  | 65 KB  |
| `ts_small_parse_table_map` | 183 KB  | 78 KB  | 53 KB  |
| `ts_lex_modes`             | 300 KB  | 24 KB  | 20 KB  |

## Sharing between grammars

| Approach                                            | Result                                            |
|-----------------------------------------------------|---------------------------------------------------|
| LZMA2, solid archive vs one stream per grammar      | 852 KB vs 858 KB: −0.7%                           |
| zstd -22 `--long`, solid vs one stream per grammar  | 1010 KB vs 1026 KB: −1.5%                         |
| FastCDC, average chunk of 256 B, 1 KB, 4 KB, 16 KB  | 0 bytes shared between grammars                   |
| FastCDC, average chunk of 64 B                      | 10.9 KB shared in total, 2.2 KB of it tsx+typescript |
| FastCDC on byte planes (see below), 64 B            | 64 KB shared, ~0.3% of raw                        |
| tsx compressed against typescript (zstd `--patch-from`) | 86.5 KB vs 89.0 KB alone: −3%                 |
| cpp compressed against c                            | −0.06%                                            |
| Trained zstd dictionary, 16–110 KB                  | Dictionary + frames come out larger than no dictionary, even when trained on the same data |

The tables are written in each grammar's own state numbers and symbol IDs. Any difference between
grammars renumbers everything: tsx is typescript plus JSX, yet their tables hardly match byte for byte.
They're alike in structure, not in bytes, so content-defined chunks find nothing to share. The
redundancy is within each grammar, in runs of zeros and repeated rows, and LZ compression already
captures it.

A solid archive would also break loading on demand: reading a grammar means inflating everything
stored before it. One embedded file of independent streams with an offset index would keep lazy
loading, but it's packaging only and doesn't change the size.

## Codecs, one stream per grammar

The decoder cost is the growth of a stripped `che` when the decoder is linked in. It was measured with
`go build -overlay`, which leaves the repository unchanged. The load time is for the largest grammar,
csharp (5.2 MB raw).

| Codec                                        | Payload  | Decoder        | Net       | csharp load |
|----------------------------------------------|----------|----------------|-----------|-------------|
| Go `compress/zlib` BestCompression (current) | 1601 KB  | 0              | —         | 8 ms        |
| `gzip -9`                                    | 1525 KB  | 0              | −76 KB    | 8 ms        |
| libdeflate `-12` (zopfli-class encoder)      | 1413 KB  | 0              | −189 KB   | 8 ms        |
| zstd -22, klauspost/compress                 | 1026 KB  | +213 KB        | −363 KB   | 6 ms        |
| brotli -11, andybalholm/brotli               | 970 KB   | +2.2 MB¹       | worse     | —           |
| LZMA2 `lc=0,lp=1,pb=1`, ulikunitz/xz/lzma    | 848 KB   | +62 KB         | **−690 KB** | 27 ms     |

¹ Measured in a minimal binary: the package's encoder tables can't be removed by the linker.

Go's BestCompression is ~5% worse than `gzip -9`. LZMA2's `lp=1,pb=1` match the 2-byte alignment of the
tables, and LZMA2 decodes slower but still fast enough for a first load.

## Transforms

- Delta-coding the increasing uint32 offsets of `ts_small_parse_table_map`: LZMA2 53 → 17 KB,
  zlib 78 → 21 KB.
- Storing `i − ts_primary_state_ids[i]`: LZMA2 65 → 58 KB, zlib 78 → 69 KB.
- Splitting values into byte planes, and storing `ts_parse_table` column by column: both made
  LZMA2 and zstd worse, and gzip ~7% better.

The two deltas together save ~43 KB with LZMA2, or ~66 KB with zlib, for whichever codec is used.

## If it's ever worth doing

From the most saved to the least work:

1. Raw LZMA2 per grammar plus the deltas: ~1.6 MB → ~0.8 MB. The 805 KB total with the deltas is an
   estimate, not measured end to end. The decoder could be the LZMA SDK's `LzmaDec.c` through the cgo
   that's already there, instead of ulikunitz/xz. That should cost ~10–15 KB, not measured. fetchgrammar
   would need an LZMA2 encoder.
2. Keep zlib and encode with a zopfli-class encoder plus the deltas: ~−250 KB, with no change to the decoder.

A shared archive, a chunk store or a dictionary is not worth doing: there's nothing for them to share.
