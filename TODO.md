- [x] Text selection: with kb and mouse
- [x] Experiment with markdown syntax highlight
- [x] Extend Go syntax highlight
- [x] Commit description syntax highlight
- [x] Rebase file highlight
- [x] Text selection: copy / paste / cut
- [x] Fix mouse click offset issue (with tabs and utf-8)
- [x] Horizontal scroll
- [x] Address race issues in the Editor/Run test
- [x] Navigate to line number
- [x] Action on md checkboxes
- [x] Run tests from normal mode (action on a test function definition)
- [x] Same for main function
- [x] ctrl+/ for comments
- [x] LSP diagnostics (line marker + status bar with problems count)
- [x] Unify how text content is distinguished by the buffer

- unify how external commands are run (consider a helper process that receives commands to execute): d2, go run, git diff
- undo/redo backed by tracking changes with content/changes.History, consider batch edit interface on a buffer
- project search (render with the tool buffer)
- project search inter-process (dir + main editor)
- find references (render with the tool buffer), call hierarchy, diagnostics using the tool buffer
- git diff coloring the line number (greenish for insert, blueish for change, greying for whitespace change, red underline for deletion)
- navigation history: go to definition, go back

- Windows and Linux support for spawning the main editor in a new pane
- Allow selecting the path in the status bar to copy
- code structure selection
- key bindings / commands registry

- Right col after return (understand tabs)

- multiple cursors
- SQL syntax highlight
