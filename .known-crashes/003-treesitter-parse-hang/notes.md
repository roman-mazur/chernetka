# Tree-sitter parse hang

The editor stops responding for 10s while editing a new `.cue` file. The stack dump shows
the editor loop inside a tree-sitter parse:

```
github.com/tree-sitter/go-tree-sitter.(*Parser).Parse(...)
rmazur.io/chernetka/internal/editor/extsyntaxhl.(*tsHighlighter).reparse(...)
	internal/editor/extsyntaxhl/treesitter.go:195
rmazur.io/chernetka/internal/editor/extsyntaxhl.(*document).ensureParsed(...)
	internal/editor/extsyntaxhl/syntaxhl.go:181
rmazur.io/chernetka/internal/editor/extsyntaxhl.(*Integration).AfterEdit(...)
	internal/editor/extsyntaxhl/syntaxhl.go:158
```

The parsed text is short (0x1b6 bytes), so it's likely the cue grammar looping
on some input rather than a slow parse.

## Architecture

The parsing runs on the editor loop after every edit, so a slow or stuck parser freezes the editor.
Consider parsing with a timeout or cancellation (tree-sitter supports a progress callback
in `ParseWithOptions`), or parsing off the editor loop.

## Next steps

Take the text of the buffer at the hang (e.g. log it before `reparse`) and reproduce with
the cue grammar alone.
