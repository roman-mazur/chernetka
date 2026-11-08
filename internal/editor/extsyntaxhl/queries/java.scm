; Highlight query for Java, adapted from tree-sitter-java's queries/highlights.scm.
;
; Patterns are matched against the capture names in treesitter.go. When two
; patterns capture the exact same node the one listed first here wins, so the
; specific patterns come before the general ones.
;
; Like the C# query, this one has no catch-all rule for identifiers: it would
; paint over the dotted names captured as a whole, like the package of an
; import, and plain identifiers are not colored anyway.

; Annotations, with their @ but not their arguments.

(annotation
  "@" @annotation
  name: (_) @annotation)

(marker_annotation
  "@" @annotation
  name: (_) @annotation)

; Declarations

(method_declaration
  name: (identifier) @function)

(constructor_declaration
  name: (identifier) @function)

(compact_constructor_declaration
  name: (identifier) @function)

(enum_constant
  name: (identifier) @constant)

; Calls

(method_invocation
  name: (identifier) @function.call)

; Packages

(package_declaration
  [
    (identifier)
    (scoped_identifier)
  ] @module)

(import_declaration
  [
    (identifier)
    (scoped_identifier)
  ] @module)

; Types. A capitalized name a member is taken from is more likely a class than
; a variable.

(class_declaration
  name: (identifier) @type)

(interface_declaration
  name: (identifier) @type)

(enum_declaration
  name: (identifier) @type)

(record_declaration
  name: (identifier) @type)

(annotation_type_declaration
  name: (identifier) @type)

((field_access
  object: (identifier) @type)
  (#match? @type "^[A-Z]"))

((method_invocation
  object: (identifier) @type)
  (#match? @type "^[A-Z]"))

((method_reference
  .
  (identifier) @type)
  (#match? @type "^[A-Z]"))

[
  (type_identifier)
  (boolean_type)
  (integral_type)
  (floating_point_type)
  (void_type)
] @type

; Constants. Upper case names are constants by convention.

[
  (true)
  (false)
  (null_literal)
] @constant.builtin

((identifier) @constant
  (#match? @constant "^_*[A-Z][A-Z0-9_]+$"))

; Members, after the constants: Color.RED reads as a constant.

(field_access
  field: (identifier) @property)

; Literals

[
  (string_literal)
  (character_literal)
] @string

(escape_sequence) @escape

; An embedded expression of a string template is code, not string content.
(string_interpolation) @embedded

[
  (hex_integer_literal)
  (decimal_integer_literal)
  (octal_integer_literal)
  (binary_integer_literal)
  (decimal_floating_point_literal)
  (hex_floating_point_literal)
] @number

[
  (line_comment)
  (block_comment)
] @comment

; Keywords

[
  (this)
  (super)
] @keyword

[
  "@interface"
  "abstract"
  "assert"
  "break"
  "case"
  "catch"
  "class"
  "continue"
  "default"
  "do"
  "else"
  "enum"
  "exports"
  "extends"
  "final"
  "finally"
  "for"
  "if"
  "implements"
  "import"
  "instanceof"
  "interface"
  "module"
  "native"
  "new"
  "non-sealed"
  "open"
  "opens"
  "package"
  "permits"
  "private"
  "protected"
  "provides"
  "public"
  "record"
  "requires"
  "return"
  "sealed"
  "static"
  "strictfp"
  "switch"
  "synchronized"
  "throw"
  "throws"
  "to"
  "transient"
  "transitive"
  "try"
  "uses"
  "volatile"
  "when"
  "while"
  "with"
  "yield"
] @keyword
