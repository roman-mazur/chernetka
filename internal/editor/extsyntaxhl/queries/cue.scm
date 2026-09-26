; Highlight query for CUE, adapted from tree-sitter-cue's queries/highlights.scm.
;
; Patterns are matched against the capture names in treesitter.go. When two
; patterns capture the exact same node the one listed first here wins, so the
; specific patterns come before the catch-all identifier rules below.

; Imports

(import_spec
  path: (_) @module)

(package_clause
  (package_identifier) @variable)

; Definitions (#Name, _#Name) are CUE's types, both where they are declared and
; where they are referenced.

((identifier) @type
  (#match? @type "^_?#"))

; Calls

(call_expression
  function: (identifier) @function.call)

(call_expression
  function: (selector_expression
    (_)
    (identifier) @function.call))

(call_expression
  function: (builtin_function) @function.call)

; Fields

(label
  (identifier) @property)

(label
  (required
    (identifier) @property))

(label
  (optional
    (identifier) @property))

(selector_expression
  (_)
  (identifier) @property)

; Types

(primitive_type) @type

; Constants

[
  (boolean)
  (null)
  (top)
  (bottom)
] @constant.builtin

; An attribute, like @go(Name), is metadata for tools rather than a part of the
; value, so its names are not colorized as identifiers.
(attribute) @constant

(attribute
  (identifier) @constant)

; Literals

(string) @string

[
  (escape_char)
  (escape_unicode)
  (escape_byte)
] @escape

; An interpolated expression is code, not string content. Its delimiters read
; like escapes, and the expression itself is reset to plain text before the
; patterns for its own tokens paint over it.
(interpolation
  "\\(" @escape
  (_) @embedded
  ")" @escape)

[
  (number)
  (float)
  (si_unit)
] @number

(comment) @comment

; Keywords

[
  "else"
  "for"
  "if"
  "import"
  "in"
  "let"
  "package"
  "try"
] @keyword

; Catch-all identifier rule, kept last so the patterns above take precedence.

(identifier) @variable
