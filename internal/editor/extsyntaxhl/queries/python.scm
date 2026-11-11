; Highlight query for Python, adapted from tree-sitter-python's queries/highlights.scm.
;
; Patterns are matched against the capture names in treesitter.go. When two
; patterns capture the exact same node the one listed first here wins, so the
; specific patterns come before the catch-all identifier rule below.

; The shebang, an instruction to the loader, is parsed as the first comment.
((module
  .
  (comment) @directive)
  (#match? @directive "^#!"))

; Decorators, with their @ and the name of the decorator but not its arguments.

(decorator
  "@" @annotation)

(decorator
  (identifier) @annotation)

(decorator
  (attribute
    (identifier) @annotation))

(decorator
  (attribute
    (attribute
      (identifier) @annotation)))

(decorator
  (call
    function: (identifier) @annotation))

(decorator
  (call
    function: (attribute
      (identifier) @annotation)))

(decorator
  (call
    function: (attribute
      (attribute
        (identifier) @annotation))))

; Definitions

(function_definition
  name: (identifier) @function)

(class_definition
  name: (identifier) @type)

; Imports. The alias of a module is a module too, but the names imported from
; a module may be anything.

(import_statement
  name: (dotted_name
    (identifier) @module))

(import_statement
  name: (aliased_import
    name: (dotted_name
      (identifier) @module)
    alias: (identifier) @module))

(import_from_statement
  module_name: (dotted_name
    (identifier) @module))

(import_from_statement
  module_name: (relative_import
    (dotted_name
      (identifier) @module)))

(future_import_statement
  "__future__" @module)

; Calls. A capitalized function is most likely a class being instantiated,
; which reads as a type.

((call
  function: (identifier) @type)
  (#match? @type "^_*[A-Z]"))

((call
  function: (attribute
    attribute: (identifier) @type))
  (#match? @type "^_*[A-Z]"))

(call
  function: (identifier) @function.call)

(call
  function: (attribute
    attribute: (identifier) @function.call))

; Constants. Upper case names are constants by convention.

[
  (none)
  (true)
  (false)
  (ellipsis)
] @constant.builtin

((identifier) @constant
  (#match? @constant "^_*[A-Z][A-Z0-9_]+$"))

; Types: the annotations, and the capitalized names, which are more likely
; classes than variables.

(type
  (identifier) @type)

(type
  (attribute
    attribute: (identifier) @type))

(generic_type
  (identifier) @type)

(member_type
  (identifier) @type)

(splat_type
  (identifier) @type)

((identifier) @type
  (#match? @type "^_*[A-Z]"))

; Fields

(attribute
  attribute: (identifier) @property)

; Literals

(string) @string

(escape_sequence) @escape

; An interpolated expression is code, not string content. Its braces read
; like escapes, and the expression is reset to plain text before the
; patterns for its own tokens paint over it. The format specification after
; the colon is a part of the string again.

(interpolation) @embedded

(interpolation
  [
    "{"
    "}"
  ] @escape)

(format_expression
  [
    "{"
    "}"
  ] @escape)

(format_specifier) @string

(type_conversion) @escape

(escape_interpolation) @escape

[
  (integer)
  (float)
] @number

(comment) @comment

; Keywords, with the operators written as words. The soft keywords, like match
; and type, are tokens only where they act as keywords.

[
  "and"
  "as"
  "assert"
  "async"
  "await"
  "break"
  "case"
  "class"
  "continue"
  "def"
  "del"
  "elif"
  "else"
  "except"
  "exec"
  "finally"
  "for"
  "from"
  "global"
  "if"
  "import"
  "in"
  "is"
  "is not"
  "lambda"
  "match"
  "nonlocal"
  "not"
  "not in"
  "or"
  "pass"
  "print"
  "raise"
  "return"
  "try"
  "type"
  "while"
  "with"
  "yield"
] @keyword

; Catch-all identifier rule, kept last so the patterns above take precedence.

(identifier) @variable
