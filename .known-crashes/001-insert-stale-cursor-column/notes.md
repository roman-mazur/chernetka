# Insert at a stale cursor column

Typing or pasting in the insert mode panics because the cursor column is past the end of its line:

```
panic: runtime error: slice bounds out of range [:20] with length 0

rmazur.io/chernetka/internal/editor.insertContent(...)
	internal/editor/insert.go:112
rmazur.io/chernetka/internal/editor.insertInput(...)
	internal/editor/insert.go:100
```

The paste variant panics in `content.InsertText` (`internal/content/mutate.go:32`) called by `PasteText`.

Seen with Go, JSON, and new files, with and without the language server.

## Findings

- The cursor is clamped on every render (`Buffer.clampCursor`), but not between the commands
  handled before the next render. Any command leaving `Buffer.c` outside the content is only
  caught if a render comes first.
- Replaying the recent input of the logs, both through `editor.TestHarness` and against the built
  `che`, doesn't reproduce the panic: it depends on the timing or on the earlier input, which
  isn't in the logs (only the last 50 inputs are).
- Modified vertical arrows (`\x1b[1;10B`, `\x1b[1;3B`, `\x1b[1;9A`) precede several of the
  panics, but tried alone they keep the cursor valid.
- Many places assign `Buffer.c` directly (`grep -n 'c\.Col *=\|\.c = ' internal/editor/*.go`),
  and only some clamp it. Rather than fixing them one by one, consider a single place keeping
  the cursor valid after any content or cursor change, e.g. clamping after every command
  on the editor loop, or making the cursor private to a type that validates it.

## Next steps

Find the command breaking the cursor: check the invariant after every key and command on the
editor loop (`Editor.handleKey`, the command loop of `Editor.Run`), panicking with the name of
the key or command, and run the monkey test with it.
