; Highlight query for HTML, adapted from tree-sitter-html's queries/highlights.scm.
;
; Patterns are matched against the capture names in treesitter.go. Like the
; elements of JSX, the tags read as keywords and their attributes as
; properties. The scripts and the style sheets embedded in a page are left
; plain.

(doctype) @directive

[
  (tag_name)
  (erroneous_end_tag_name)
] @keyword

(attribute_name) @property

[
  (attribute_value)
  (quoted_attribute_value)
] @string

; A character reference, like &amp;, reads like an escape.
(entity) @escape

(comment) @comment
