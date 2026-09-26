; Highlight query for JSON, adapted from tree-sitter-json's queries/highlights.scm.
;
; Patterns are matched against the capture names in treesitter.go. When two
; patterns capture the exact same node the one listed first here wins, so an
; object key is a property rather than a string.

(pair
  key: (string) @property)

(string) @string

(escape_sequence) @escape

(number) @number

[
  (true)
  (false)
  (null)
] @constant.builtin

; Not in the JSON spec, but accepted by the grammar, as JSONC files have them.
(comment) @comment
