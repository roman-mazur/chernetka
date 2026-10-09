# Known crashes

The crashes and hangs found by the monkey fuzz test (`FuzzMonkey` in `cmd/che/monkey_test.go`)
that are not fixed yet. Each folder is one known problem: `notes.md` describes it, the fuzz inputs
reproducing it are kept in the `go test fuzz v1` format, and the `.log` files are the output
of the failed runs.

An input repeats the typed input only: the timing of the editor differs from run to run,
so an input may not reproduce a crash every time. The project files change over time too,
so an input editing a copy of a project file may pick another file later. The recent input,
the stack, and the editor log tail in the logs are the main evidence.

The failing inputs are saved by `go test -fuzz` in `cmd/che/testdata/fuzz/FuzzMonkey`.
To add a crash, move the input into the folder of the known problem, or into a new folder
if the stack is new, and put the output of its run next to it:

```bash
CHEMONKEY=1 go test -count=1 -run 'FuzzMonkey/<input>$' ./cmd/che > run.log 2>&1
```

To type a known input again, copy it to `cmd/che/testdata/fuzz/FuzzMonkey` and run the same command.

Remove the folder with the fix of the problem.
