; Highlight query for Kotlin, adapted from tree-sitter-kotlin's queries/highlights.scm.
;
; Patterns are matched against the capture names in treesitter.go. When two
; patterns capture the exact same node the one listed first here wins, so the
; specific patterns come before the catch-all identifier rule below.

; Annotations, with their @ and use-site target but not their arguments.

(annotation
  "@" @annotation)

(annotation
  (use_site_target) @annotation)

(annotation
  (user_type
    (type_identifier) @annotation))

(annotation
  (constructor_invocation
    (user_type
      (type_identifier) @annotation)))

(file_annotation
  [
    "@"
    "file"
    ":"
  ] @annotation)

(file_annotation
  (user_type
    (type_identifier) @annotation))

(file_annotation
  (constructor_invocation
    (user_type
      (type_identifier) @annotation)))

; Declarations. A function's own name is its only direct simple_identifier:
; the receiver, the parameters and the body are nested in nodes of their own.

(function_declaration
  (simple_identifier) @function)

(class_parameter
  (simple_identifier) @property)

(class_body
  (property_declaration
    (variable_declaration
      (simple_identifier) @property)))

(enum_entry
  (simple_identifier) @constant)

; Calls

(call_expression
  .
  (simple_identifier) @function.call)

(call_expression
  (navigation_expression
    (navigation_suffix
      (simple_identifier) @function.call)
    .))

; Packages. A dotted name is a flat list of simple identifiers.

(package_header
  (identifier
    (simple_identifier) @module))

(import_header
  (identifier
    (simple_identifier) @module))

; Types. A capitalized name a member is taken from is more likely a class or
; an object than a variable.

((navigation_expression
  .
  (simple_identifier) @type)
  (#match? @type "^[A-Z]"))

(type_identifier) @type

; Constants. Upper case names are constants by convention.

[
  (boolean_literal)
  "null"
] @constant.builtin

((simple_identifier) @constant
  (#match? @constant "^_*[A-Z][A-Z0-9_]+$"))

; Members, after the constants: Color.RED reads as a constant.

(navigation_suffix
  (simple_identifier) @property)

; Literals

[
  (string_literal)
  (character_literal)
] @string

(character_escape_seq) @escape

; A string template is code, not string content. Its delimiters read like
; escapes, and the rest is reset to plain text before the patterns for its
; own tokens paint over it.
(string_literal
  "$" @escape
  (interpolated_identifier) @embedded)

(string_literal
  "${" @escape
  (interpolated_expression) @embedded
  "}" @escape)

[
  (integer_literal)
  (long_literal)
  (hex_literal)
  (bin_literal)
  (unsigned_literal)
  (real_literal)
] @number

(shebang_line) @directive

[
  (line_comment)
  (multiline_comment)
] @comment

; Keywords

[
  (class_modifier)
  (member_modifier)
  (function_modifier)
  (property_modifier)
  (platform_modifier)
  (variance_modifier)
  (parameter_modifier)
  (visibility_modifier)
  (reification_modifier)
  (inheritance_modifier)
  (this_expression)
  (super_expression)
] @keyword

[
  "as"
  "as?"
  "break"
  "break@"
  "by"
  "catch"
  "class"
  "companion"
  "constructor"
  "continue"
  "continue@"
  "do"
  "else"
  "enum"
  "finally"
  "for"
  "fun"
  "get"
  "if"
  "import"
  "in"
  "!in"
  "init"
  "interface"
  "is"
  "!is"
  "object"
  "package"
  "return"
  "return@"
  "set"
  "throw"
  "try"
  "typealias"
  "val"
  "var"
  "when"
  "where"
  "while"
] @keyword

; Catch-all identifier rule, kept last so the patterns above take precedence.

(simple_identifier) @variable
