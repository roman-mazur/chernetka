; Highlight query for TypeScript, adapted from tree-sitter-javascript's and
; tree-sitter-typescript's queries/highlights.scm. The TSX grammar extends the
; TypeScript one with JSX, and the TSX query is jsx.scm followed by this file:
; everything here has to compile against both grammars. JavaScript is read
; with the TSX grammar.
;
; Patterns are matched against the capture names in treesitter.go. When two
; patterns capture the exact same node the one listed first here wins, so the
; specific patterns come before the catch-all identifier rule below.

; Directives: the shebang, and the directive prologues like "use strict" and
; React's "use client".

(hash_bang_line) @directive

((expression_statement
  (string) @directive)
  (#match? @directive "^[\"']use (strict|client|server)[\"']$"))

; Decorators, with their @ but not their arguments.

(decorator
  "@" @annotation)

(decorator
  (identifier) @annotation)

(decorator
  (call_expression
    function: (identifier) @annotation))

(decorator
  (call_expression
    function: (member_expression
      property: (property_identifier) @annotation)))

; Declarations, including the functions bound to a name.

[
  (function_declaration
    name: (identifier) @function)
  (function_expression
    name: (identifier) @function)
  (generator_function_declaration
    name: (identifier) @function)
  (function_signature
    name: (identifier) @function)
  (method_definition
    name: (property_identifier) @function)
  (method_signature
    name: (property_identifier) @function)
  (abstract_method_signature
    name: (property_identifier) @function)
]

(variable_declarator
  name: (identifier) @function
  value: [
    (function_expression)
    (arrow_function)
  ])

(pair
  key: (property_identifier) @function
  value: [
    (function_expression)
    (arrow_function)
  ])

(enum_body
  name: (property_identifier) @constant)

(enum_assignment
  name: (property_identifier) @constant)

; Calls

(call_expression
  function: (identifier) @function.call)

(call_expression
  function: (member_expression
    property: (property_identifier) @function.call))

; Imports

(import_statement
  source: (string) @module)

(export_statement
  source: (string) @module)

; Constants. Upper case names are constants by convention.

[
  (true)
  (false)
  (null)
  (undefined)
] @constant.builtin

([
  (identifier)
  (property_identifier)
  (shorthand_property_identifier)
  (shorthand_property_identifier_pattern)
] @constant
  (#match? @constant "^_*[A-Z][A-Z0-9_]+$"))

; Types. A capitalized identifier is more likely a class than a variable.

[
  (type_identifier)
  (predefined_type)
] @type

((identifier) @type
  (#match? @type "^[A-Z]"))

; Properties

[
  (property_identifier)
  (private_property_identifier)
] @property

; Literals

[
  (string)
  (template_string)
  (regex)
] @string

(escape_sequence) @escape

; A template substitution is code, not string content. Its delimiters read
; like escapes, and the rest is reset to plain text before the patterns for its
; own tokens paint over it.
(template_substitution) @embedded

(template_substitution
  [
    "${"
    "}"
  ] @escape)

(number) @number

[
  (comment)
  (html_comment)
] @comment

; Keywords

[
  (this)
  (super)
] @keyword

[
  "abstract"
  "accessor"
  "as"
  "assert"
  "asserts"
  "async"
  "await"
  "break"
  "case"
  "catch"
  "class"
  "const"
  "continue"
  "debugger"
  "declare"
  "default"
  "delete"
  "do"
  "else"
  "enum"
  "export"
  "extends"
  "finally"
  "for"
  "from"
  "function"
  "get"
  "global"
  "if"
  "implements"
  "import"
  "in"
  "infer"
  "instanceof"
  "interface"
  "is"
  "keyof"
  "let"
  "module"
  "namespace"
  "new"
  "of"
  "override"
  "private"
  "protected"
  "public"
  "readonly"
  "require"
  "return"
  "satisfies"
  "set"
  "static"
  "switch"
  "throw"
  "try"
  "type"
  "typeof"
  "using"
  "var"
  "void"
  "while"
  "with"
  "yield"
] @keyword

; Catch-all identifier rule, kept last so the patterns above take precedence.

(identifier) @variable
