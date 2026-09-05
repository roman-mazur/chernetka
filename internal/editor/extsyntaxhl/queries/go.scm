; Highlight query for Go, adapted from tree-sitter-go's queries/highlights.scm.
;
; Patterns are matched against the capture names in treesitter.go. When two
; patterns capture the exact same node the one listed first here wins, so the
; specific patterns come before the catch-all identifier rules below.

; Declarations

(function_declaration
  name: (identifier) @function)

(method_declaration
  name: (field_identifier) @function)

(type_spec
  name: (type_identifier) @type)

; Calls

(call_expression
  function: (identifier) @function.call)

(call_expression
  function: (selector_expression
    field: (field_identifier) @function.call))

; Imports

(import_spec
  path: (_) @module)

(package_clause
  (package_identifier) @variable)

(package_identifier) @module

; A composite literal key names a struct field, so it reads like one.

(keyed_element
  key: (literal_element
    (identifier) @property))

; Types

(type_identifier) @type

; Constants

[
  (true)
  (false)
  (nil)
  (iota)
] @constant.builtin

; Literals

[
  (interpreted_string_literal)
  (raw_string_literal)
  (rune_literal)
] @string

(escape_sequence) @escape

[
  (int_literal)
  (float_literal)
  (imaginary_literal)
] @number

(comment) @comment

; Keywords. See https://go.dev/ref/spec#Keywords.

[
  "break"
  "case"
  "chan"
  "const"
  "continue"
  "default"
  "defer"
  "else"
  "fallthrough"
  "for"
  "func"
  "go"
  "goto"
  "if"
  "import"
  "interface"
  "map"
  "package"
  "range"
  "return"
  "select"
  "struct"
  "switch"
  "type"
  "var"
] @keyword

; Catch-all identifier rules, kept last so the patterns above take precedence.

(field_identifier) @property
(identifier) @variable
