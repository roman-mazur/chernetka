; Highlight query for Rust, adapted from tree-sitter-rust's queries/highlights.scm.
;
; Patterns are matched against the capture names in treesitter.go. When two
; patterns capture the exact same node the one listed first here wins, so the
; specific patterns come before the catch-all identifier rule below.

(shebang) @directive

; Attributes, with their #, ! and brackets but not their arguments.

(attribute_item
  [
    "#"
    "["
    "]"
  ] @annotation)

(inner_attribute_item
  [
    "#"
    "!"
    "["
    "]"
  ] @annotation)

(attribute
  .
  (identifier) @annotation)

(attribute
  .
  (scoped_identifier
    path: (identifier) @annotation
    name: (identifier) @annotation))

; Declarations

(function_item
  name: (identifier) @function)

(function_signature_item
  name: (identifier) @function)

(macro_definition
  name: (identifier) @function)

; Calls, of functions, methods and macros. A capitalized function is an enum
; variant or a tuple struct, which reads as a type, like the variants without
; fields.

((call_expression
  function: (identifier) @type)
  (#match? @type "^[A-Z]"))

((call_expression
  function: (scoped_identifier
    name: (identifier) @type))
  (#match? @type "^[A-Z]"))

(call_expression
  function: (identifier) @function.call)

(call_expression
  function: (field_expression
    field: (field_identifier) @function.call))

(call_expression
  function: (scoped_identifier
    name: (identifier) @function.call))

(generic_function
  function: (identifier) @function.call)

(generic_function
  function: (scoped_identifier
    name: (identifier) @function.call))

(generic_function
  function: (field_expression
    field: (field_identifier) @function.call))

(macro_invocation
  macro: (identifier) @function.call
  "!" @function.call)

(macro_invocation
  macro: (scoped_identifier
    name: (identifier) @function.call)
  "!" @function.call)

; Constants. Upper case names are constants by convention.

(boolean_literal) @constant.builtin

((identifier) @constant
  (#match? @constant "^_*[A-Z][A-Z0-9_]+$"))

; Types, and the modules of a path. A capitalized name is more likely a type or
; an enum variant than a module or a variable.

((identifier) @type
  (#match? @type "^[A-Z]"))

(mod_item
  name: (identifier) @module)

(scoped_identifier
  path: (identifier) @module)

(scoped_identifier
  path: (scoped_identifier
    name: (identifier) @module))

(scoped_type_identifier
  path: (identifier) @module)

(scoped_type_identifier
  path: (scoped_identifier
    name: (identifier) @module))

(scoped_use_list
  path: (identifier) @module)

(scoped_use_list
  path: (scoped_identifier
    name: (identifier) @module))

[
  (type_identifier)
  (primitive_type)
  (lifetime)
] @type

(lifetime
  (identifier) @type)

; Fields

[
  (field_identifier)
  (shorthand_field_identifier)
] @property

; Literals

[
  (string_literal)
  (raw_string_literal)
  (char_literal)
] @string

(escape_sequence) @escape

[
  (integer_literal)
  (float_literal)
] @number

[
  (line_comment)
  (block_comment)
] @comment

; Keywords

[
  (crate)
  (super)
  (self)
  (mutable_specifier)
] @keyword

[
  "as"
  "async"
  "await"
  "break"
  "const"
  "continue"
  "default"
  "dyn"
  "else"
  "enum"
  "extern"
  "fn"
  "for"
  "gen"
  "if"
  "impl"
  "in"
  "let"
  "loop"
  "macro_rules!"
  "match"
  "mod"
  "move"
  "pub"
  "raw"
  "ref"
  "return"
  "static"
  "struct"
  "trait"
  "try"
  "type"
  "union"
  "unsafe"
  "use"
  "where"
  "while"
  "yield"
] @keyword

; Catch-all identifier rule, kept last so the patterns above take precedence.

(identifier) @variable
