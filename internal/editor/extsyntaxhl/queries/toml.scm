; Highlight query for TOML, adapted from tree-sitter-toml's queries/highlights.scm.
;
; Patterns are matched against the capture names in treesitter.go. When two
; patterns capture the exact same node the one listed first here wins, so the
; keys come before the strings they may be made of.

; Tables. The name of a table reads as a type, a key as a property.

(table
  [
    (bare_key)
    (quoted_key)
  ] @type)

(table
  (dotted_key
    [
      (bare_key)
      (quoted_key)
    ] @type))

(table_array_element
  [
    (bare_key)
    (quoted_key)
  ] @type)

(table_array_element
  (dotted_key
    [
      (bare_key)
      (quoted_key)
    ] @type))

[
  (bare_key)
  (quoted_key)
] @property

; Literals

(string) @string

(escape_sequence) @escape

[
  (integer)
  (float)
] @number

(boolean) @constant.builtin

[
  (offset_date_time)
  (local_date_time)
  (local_date)
  (local_time)
] @constant

(comment) @comment
