; Highlight query for Protocol Buffers, adapted from tree-sitter-proto's
; queries/highlights.scm.
;
; Patterns are matched against the capture names in extsyntaxhl's treesitter.go.
; The messages, the enums and the services read as types, the methods as
; functions, and the fields and the options as properties.

(package
  (full_ident
    (identifier) @module))

; Types

[
  (message_name)
  (enum_name)
  (service_name)
  (message_or_enum_type)
  (key_type)
  (type)
] @type

(extend
  (full_ident
    (identifier) @type))

(extension_name
  type: (full_ident
    (identifier) @type))

(rpc_name) @function

; Fields and options

[
  (field
    (identifier) @property)
  (map_field
    (identifier) @property)
  (oneof
    (identifier) @property)
  (oneof_field
    (identifier) @property)
  (option
    (identifier) @property)
  (option
    (full_ident
      (identifier) @property))
  (field_option
    (identifier) @property)
  (field_option
    (full_ident
      (identifier) @property))
  (enum_value_option
    (identifier) @property)
  (enum_value_option
    (full_ident
      (identifier) @property))
  (block_lit
    (identifier) @property)
  (extension_name
    name: (full_ident
      (identifier) @property))
]

(enum_field
  (identifier) @constant)

(constant
  (full_ident
    (identifier) @constant))

; Literals

[
  (string)
  (reserved_identifier)
] @string

(escape_sequence) @escape

[
  (int_lit)
  (float_lit)
] @number

[
  (true)
  (false)
] @constant

(comment) @comment

; Keywords

[
  "syntax"
  "edition"
  "package"
  "import"
  "option"
  "message"
  "enum"
  "service"
  "rpc"
  "returns"
  "stream"
  "oneof"
  "map"
  "extend"
  "extensions"
  "group"
  "reserved"
  "to"
  "max"
  "optional"
  "repeated"
  "required"
  "weak"
  "public"
  "export"
  "local"
] @keyword
