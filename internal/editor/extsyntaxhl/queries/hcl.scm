; Highlight query for HCL. The grammar comes with no highlight query of its
; own.
;
; Patterns are matched against the capture names in treesitter.go. When two
; patterns capture the exact same node the one listed first here wins, so the
; specific patterns come before the catch-all identifier rule below.

; Blocks. The type of a block, like resource, reads as a keyword, and its
; labels written as names as types. The quoted labels stay strings.

(block
  .
  (identifier) @keyword)

(block
  (identifier) @type)

; Attributes, and the keys of the objects written as names.

(attribute
  (identifier) @property)

(object_elem
  key: (expression
    (variable_expr
      (identifier) @property)))

(get_attr
  (identifier) @property)

(function_call
  (identifier) @function.call)

; The roots of the references in Terraform, like var in var.region, are
; namespaces rather than variables.

((variable_expr
  (identifier) @module)
  (#match? @module "^(var|local|module|data|each|count|path|terraform|self)$"))

; Literals

[
  (string_lit)
  (quoted_template)
  (heredoc_template)
] @string

(numeric_lit) @number

[
  (bool_lit)
  (null_lit)
] @constant.builtin

; An interpolation or a directive in a template is code, not string content.
; Its delimiters read like escapes, and the rest is reset to plain text
; before the patterns for its own tokens paint over it.

[
  (template_interpolation)
  (template_if_intro)
  (template_else_intro)
  (template_if_end)
  (template_for_start)
  (template_for_end)
] @embedded

[
  (template_interpolation_start)
  (template_interpolation_end)
  (template_directive_start)
  (template_directive_end)
  (strip_marker)
] @escape

(comment) @comment

; Keywords, of the for expressions and the template directives.

[
  "for"
  "in"
  "if"
  "else"
  "endif"
  "endfor"
] @keyword

; Catch-all identifier rule, kept last so the patterns above take precedence.

(identifier) @variable
