; Highlight query for C, adapted from tree-sitter-c's queries/highlights.scm.
; The C++ grammar extends the C one, so the C++ query is cpp.scm followed by
; this file: everything here has to compile against both grammars.
;
; Patterns are matched against the capture names in treesitter.go. When two
; patterns capture the exact same node the one listed first here wins, so the
; specific patterns come before the catch-all identifier rules below.

; Attributes, like [[nodiscard]], with their brackets but not their arguments.

(attribute_declaration
  [
    "[["
    "]]"
  ] @annotation)

(attribute
  prefix: (identifier) @annotation)

(attribute
  name: (identifier) @annotation)

; Declarations

(function_declarator
  declarator: (identifier) @function)

(preproc_function_def
  name: (identifier) @function)

(preproc_def
  name: (identifier) @constant)

; Calls

(call_expression
  function: (identifier) @function.call)

(call_expression
  function: (field_expression
    field: (field_identifier) @function.call))

; Includes

(preproc_include
  path: (_) @module)

; Types

[
  (type_identifier)
  (primitive_type)
  (sized_type_specifier)
] @type

; Constants. Upper case names are macros or enumerators by convention.

[
  (true)
  (false)
  (null)
] @constant.builtin

((identifier) @constant
  (#match? @constant "^[A-Z][A-Z0-9_]*$"))

; Literals

[
  (string_literal)
  (system_lib_string)
  (char_literal)
] @string

(escape_sequence) @escape

(number_literal) @number

(comment) @comment

; Keywords. The storage classes and the type qualifiers are matched by their
; nodes rather than their words, as the grammars don't share all of the words.

[
  (storage_class_specifier)
  (type_qualifier)
] @keyword

[
  "break"
  "case"
  "continue"
  "default"
  "do"
  "else"
  "enum"
  "for"
  "goto"
  "if"
  "return"
  "sizeof"
  "struct"
  "switch"
  "typedef"
  "union"
  "while"
  "offsetof"
  "_Generic"
  "asm"
  "__asm__"
  "__attribute__"
  "defined"
] @keyword

[
  "#define"
  "#elif"
  "#elifdef"
  "#elifndef"
  "#else"
  "#endif"
  "#if"
  "#ifdef"
  "#ifndef"
  "#include"
  (preproc_directive)
] @keyword

; Catch-all identifier rules, kept last so the patterns above take precedence.

(field_identifier) @property
(identifier) @variable
