; Highlight query for JSX, adapted from tree-sitter-javascript's
; queries/highlights-jsx.scm. Like upstream, it only covers what JSX adds: the
; TSX query is this file followed by the TypeScript query, which it takes
; precedence over.

; A lower case tag is an element of the platform, like <div>, and reads like a
; keyword. A capitalized one is a component, which reads as a type by the rules
; of the TypeScript query.

(jsx_opening_element
  name: (identifier) @keyword
  (#match? @keyword "^[a-z][^.]*$"))

(jsx_closing_element
  name: (identifier) @keyword
  (#match? @keyword "^[a-z][^.]*$"))

(jsx_self_closing_element
  name: (identifier) @keyword
  (#match? @keyword "^[a-z][^.]*$"))

(jsx_attribute
  (property_identifier) @property)

(html_character_reference) @escape
