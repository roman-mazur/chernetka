; Highlight query for CSS, adapted from tree-sitter-css's queries/highlights.scm.
;
; Patterns are matched against the capture names in treesitter.go. When two
; patterns capture the exact same node the one listed first here wins, so the
; specific patterns come before the general ones.

; Selectors. Like in HTML, the tags read as keywords. The classes read as
; types and the ids as constants, with their . and #, and the pseudo-classes
; and pseudo-elements as annotations of the selector before them.

(pseudo_class_selector
  ":" @annotation
  (class_name) @annotation)

(pseudo_element_selector
  "::" @annotation
  (tag_name) @annotation)

[
  (tag_name)
  (nesting_selector)
  (universal_selector)
] @keyword

(class_selector
  "." @type
  (class_name) @type)

(id_selector
  "#" @constant
  (id_name) @constant)

(attribute_selector
  (attribute_name) @property)

(attribute_selector
  (plain_value) @string)

(namespace_name) @module

; Declarations. A custom property, like --gap, reads as a variable, where it
; is declared and where it is used.

((property_name) @variable
  (#match? @variable "^--"))

((plain_value) @variable
  (#match? @variable "^--"))

[
  (property_name)
  (feature_name)
] @property

(function_name) @function.call

(keyframes_name) @type

; At-rules and the words of their queries.

[
  (at_keyword)
  "@charset"
  "@import"
  "@keyframes"
  "@media"
  "@namespace"
  "@scope"
  "@supports"
  (to)
  (from)
  (important)
  "and"
  "not"
  "of"
  "only"
  "or"
  "selector"
  "to"
] @keyword

; Literals

(string_value) @string

(escape_sequence) @escape

(color_value) @constant

[
  (integer_value)
  (float_value)
] @number

[
  (comment)
  (js_comment)
] @comment
