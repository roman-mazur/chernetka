; Highlight query for YAML, adapted from tree-sitter-yaml's queries/highlights.scm.
;
; Patterns are matched against the capture names in treesitter.go. When two
; patterns capture the exact same node the one listed first here wins, so the
; mapping keys come before the scalars they are made of.

; Keys

(block_mapping_pair
  key: (flow_node
    [
      (double_quote_scalar)
      (single_quote_scalar)
      (plain_scalar)
    ] @property))

(flow_pair
  key: (flow_node
    [
      (double_quote_scalar)
      (single_quote_scalar)
      (plain_scalar)
    ] @property))

; Scalars

[
  (double_quote_scalar)
  (single_quote_scalar)
  (block_scalar)
  (string_scalar)
] @string

(escape_sequence) @escape

[
  (integer_scalar)
  (float_scalar)
] @number

[
  (boolean_scalar)
  (null_scalar)
] @constant.builtin

(timestamp_scalar) @constant

; Anchors, aliases and tags annotate the nodes rather than being values.

[
  (anchor)
  (alias)
] @constant

(tag) @type

; Structure

[
  (yaml_directive)
  (tag_directive)
  (reserved_directive)
  "---"
  "..."
] @keyword

(comment) @comment
