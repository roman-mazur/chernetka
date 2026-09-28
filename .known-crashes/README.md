# Known crashes

The crashes and hangs found by the monkey test (`cmd/che/monkey_test.go`) that are not fixed yet.
Each folder is one known problem: `notes.md` describes it, and the `<seed>.log` files are
the output of the failed runs, named by their `CHEMONKEY_SEED`.

The seed repeats the typed input only: the timing of the editor differs from run to run,
so a seed rarely reproduces a crash. The recent input, the stack, and the editor log tail
in the logs are the main evidence.

To add a crash, run the test and put its output into the folder of the known problem,
or into a new folder if the stack is new:

```bash
CHEMONKEY=1 go test -count=1 -run 'TestMonkey$' ./cmd/che > run.log 2>&1
```

Remove the folder with the fix of the problem.

| Problem                                                              | Seen | Kind  | Where                                         |
|----------------------------------------------------------------------|------|-------|-----------------------------------------------|
| [001 Insert at a stale cursor column](001-insert-stale-cursor-column) | 5    | panic | `editor.insertContent`, `content.InsertText`  |
| [002 Unexpected exit](002-unexpected-exit)                            | 2    | exit  | unknown input quits che                       |
| [003 Tree-sitter parse hang](003-treesitter-parse-hang)               | 1    | hang  | `extsyntaxhl.(*tsHighlighter).reparse` (cue)  |
