# Tree-sitter parse hang

che hangs in `ts_parser_parse_with_options`, called from `tsHighlighter.reparse`
(`internal/editor/extsyntaxhl/treesitter.go`), after typing random input into a new file:
the editor loop never returns from the parse, and che does not quit.

Found with the seed inputs 0 (`new.cue`) and 2 (`new.go`), which reproduce it every run.
The document is about 4KB of random text at the time of the hang.
`run.log` is the output of both runs.
