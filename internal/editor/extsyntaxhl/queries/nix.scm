; Highlight query for Nix, adapted from tree-sitter-nix's queries/highlights.scm.
;
; Patterns are matched against the capture names in treesitter.go. When two
; patterns capture the exact same node the one listed first here wins, so the
; specific patterns come before the catch-all identifier rules below.

; Declarations: an attribute bound to a function.

(binding
  attrpath: (attrpath
    attr: (identifier) @function .)
  expression: (function_expression))

; Calls. Application is juxtaposition: in "f x" the function is f.

(apply_expression
  function: (variable_expression
    name: (identifier) @function.call))

(apply_expression
  function: (select_expression
    attrpath: (attrpath
      attr: (identifier) @function.call .)))

; Constants

((variable_expression
  name: (identifier) @constant.builtin)
  (#match? @constant.builtin "^(true|false|null)$"))

((variable_expression
  name: (identifier) @module)
  (#eq? @module "builtins"))

; Attributes

(binding
  attrpath: (attrpath
    attr: (identifier) @property))

(select_expression
  attrpath: (attrpath
    attr: (identifier) @property))

(has_attr_expression
  attrpath: (attrpath
    attr: (identifier) @property))

(inherited_attrs
  attr: (identifier) @property)

; Literals

[
  (string_expression)
  (indented_string_expression)
  (path_expression)
  (hpath_expression)
  (spath_expression)
  (uri_expression)
] @string

[
  (escape_sequence)
  (dollar_escape)
] @escape

; An interpolated expression is code, not string content. Its delimiters read
; like escapes, and the expression itself is reset to plain text before the
; patterns for its own tokens paint over it.
(interpolation
  "${" @escape
  expression: (_) @embedded
  "}" @escape)

[
  (integer_expression)
  (float_expression)
] @number

(comment) @comment

; Keywords

[
  "assert"
  "else"
  "if"
  "in"
  "inherit"
  "let"
  "or"
  "rec"
  "then"
  "with"
] @keyword

; Catch-all identifier rule, kept last so the patterns above take precedence.

(identifier) @variable
