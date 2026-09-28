# Unexpected exit

che exits with status 0 in the middle of the test (`che exited: <nil>`), without a panic.
The editor log ends with `session done`, so the editor loop quit normally.

The test bans the inputs quitting on purpose (`q`), and types `:q` only at the end.
Either some other input quits the editor (then the test should ban it, or the editor
shouldn't quit on it), or the editor quits by mistake.

Seen with `new.go` and `new.txt`.

## Next steps

Log the key or command that sets `Editor.quitRequested`, or pops the last buffer, and run the
monkey test until it exits again.
