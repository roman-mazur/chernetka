# che falls behind the input

`che does not quit for 10s` after the typing: che is not hung, it's still handling the keys
typed before the quit. The monkey types thousands of keys in seconds, and some keys take
long to handle, so che falls behind. The goroutines dump catches the main goroutine wherever
it spends the time, under `Editor.handleKeys`.

Both logs edit a copy of a Go file. A seed rarely repeats the failure, but the inputs
`seed0` and `seed2` of the closed fuzzing PR (#2, `.known-crashes/treesitter-parse-hang`
on its branch) fail every run on a new CUE or Go file. With `monkeyHangTimeout` raised
to 60s, `seed2` passes: che quits about 13s after the typing ends.

## Full reparse after every key

`1791570979259612320.log`: the main goroutine is in `ts_parser_parse_with_options`, under
`extsyntaxhl.(*Integration).AfterEdit` → `document.ensureParsed` → `tsHighlighter.reparse`.

`handleKey` notifies the extensions after every key, and the syntax highlighting reparses
the whole text in `AfterEdit` (incremental parsing is a TODO in `reparse`). Every C allocation
of tree-sitter is a cgo callback (`gotreesitter/allocator.go`) too. One failing input of a few
KB of Go (`seed2` above) took 1,122 parses, 12.3s in total, up to 30ms each: none of them
hangs, and the same text parses in 17ms out of che.

Possible fix: `AfterEdit` only marks the document as outdated, and it's parsed when the spans
or the code blocks are requested, once for the keys handled together. The spans are built
lazily already.

## Saving on every buffer switch

`1791570040873849011.log`: the main goroutine waits in `extlsp.(*server).format`, under
`Editor.navigateHistory` → `GoTo` → `selectBuffer` → `saveTop` → `Save`.

Switching buffers saves the changed top buffer, and saving waits up to `formatTimeout` (1s)
for the language server to format it; the editor log has `formatting timed out`. The monkey
changes the buffer between most switches, and the navigation keys (Ctrl+Alt+Left/Right)
switch buffers often, so the monkey hits it since it presses them.

Possible fixes: don't wait for the formatting when switching buffers (format on explicit
saves only, or apply the edits when they arrive), or save without formatting there.
