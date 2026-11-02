# che falls behind the input

`che does not quit for 10s` after the typing: che is not hung, it's still handling the keys
typed before the quit. The monkey types thousands of keys in seconds, and some keys take
long to handle, so che falls behind. The goroutines dump catches the main goroutine wherever
it spends the time, under `Editor.handleKeys`.

## Saving on every buffer switch

`1791570040873849011.log`: the main goroutine waits in `extlsp.(*server).format`, under
`Editor.navigateHistory` → `GoTo` → `selectBuffer` → `saveTop` → `Save`.

Switching buffers saves the changed top buffer, and saving waits up to `formatTimeout` (1s)
for the language server to format it; the editor log has `formatting timed out`. The monkey
changes the buffer between most switches, and the navigation keys (Ctrl+Alt+Left/Right)
switch buffers often, so the monkey hits it since it presses them. The input `seed2` of the
closed fuzzing PR (#2, `.known-crashes/treesitter-parse-hang` on its branch) hit it once
too, under `:w` (`runExCommand` → `Save`).

Possible fixes: don't wait for the formatting when switching buffers (format on explicit
saves only, or apply the edits when they arrive), or save without formatting there.
