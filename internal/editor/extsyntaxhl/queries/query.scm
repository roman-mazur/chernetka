; Highlight query for tree-sitter queries, like this file, adapted from
; tree-sitter-query's queries/query/highlights.scm.
;
; Patterns are matched against the capture names in treesitter.go. When two
; patterns capture the exact same node the one listed first here wins.

; A node type reads as a type, and a capture names what it captures, like a
; constant.

(named_node
  name: (identifier) @type)

(missing_node
  name: (identifier) @type)

(capture) @constant

; A predicate, like #match? or #set!, is an instruction to the query engine.

(predicate
  [
    "#"
    "."
  ] @directive
  name: (identifier) @directive
  type: (predicate_type) @directive)

[
  (field_definition
    name: (identifier) @property)
  (negated_field
    "!" @property
    (identifier) @property)
]

; Anchors, quantifiers and wildcards change what a pattern matches.

[
  "."
  "_"
  "MISSING"
  (quantifier)
] @keyword

(string) @string

(escape_sequence) @escape

(comment) @comment
