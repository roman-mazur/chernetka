; Highlight query for C#, adapted from tree-sitter-c-sharp's queries/highlights.scm.
;
; Patterns are matched against the capture names in treesitter.go. When two
; patterns capture the exact same node the one listed first here wins, so the
; specific patterns come before the general ones.
;
; Unlike the other queries, this one has no catch-all rule for identifiers:
; it would paint over the dotted names captured as a whole, like the
; namespace of a using directive, and plain identifiers are not colored anyway.

; Attributes, with their brackets and target but not their arguments.

(attribute_list
  [
    "["
    "]"
  ] @annotation)

(attribute_target_specifier) @annotation

(attribute
  name: (_) @annotation)

; Declarations

(method_declaration
  name: (identifier) @function)

(local_function_statement
  name: (identifier) @function)

(constructor_declaration
  name: (identifier) @function)

(destructor_declaration
  name: (identifier) @function)

(property_declaration
  name: (identifier) @property)

(enum_member_declaration
  name: (identifier) @constant)

; Calls

(invocation_expression
  function: (identifier) @function.call)

(invocation_expression
  function: (generic_name
    (identifier) @function.call))

(invocation_expression
  function: (member_access_expression
    name: (identifier) @function.call))

(invocation_expression
  function: (member_access_expression
    name: (generic_name
      (identifier) @function.call)))

; Namespaces

(using_directive
  [
    (identifier)
    (qualified_name)
  ] @module)

(namespace_declaration
  name: (_) @module)

(file_scoped_namespace_declaration
  name: (_) @module)

; Types

(class_declaration
  name: (identifier) @type)

(struct_declaration
  name: (identifier) @type)

(interface_declaration
  name: (identifier) @type)

(enum_declaration
  name: (identifier) @type)

(record_declaration
  name: (identifier) @type)

(delegate_declaration
  name: (identifier) @type)

(type_parameter
  (identifier) @type)

(generic_name
  (identifier) @type)

(type_argument_list
  (identifier) @type)

(base_list
  (identifier) @type)

(_
  type: (identifier) @type)

(as_expression
  right: (identifier) @type)

(is_expression
  right: (identifier) @type)

(type_parameter_constraint
  (type
    (identifier) @type))

(predefined_type) @type

(member_access_expression
  name: (identifier) @property)

; Constants

[
  (boolean_literal)
  (null_literal)
] @constant.builtin

; Literals

[
  (string_literal)
  (verbatim_string_literal)
  (raw_string_literal)
  (character_literal)
  (interpolated_string_expression)
] @string

(escape_sequence) @escape

; An interpolated expression is code, not string content. Its braces read
; like escapes, and the expression is reset to plain text before the
; patterns for its own tokens paint over it.
(interpolation) @embedded

(interpolation_brace) @escape

[
  (integer_literal)
  (real_literal)
] @number

[
  (comment)
  (shebang_directive)
] @comment

; Keywords

[
  (modifier)
  (implicit_type)
] @keyword

[
  "add"
  "alias"
  "and"
  "as"
  "ascending"
  "await"
  "base"
  "break"
  "by"
  "case"
  "catch"
  "checked"
  "class"
  "continue"
  "default"
  "delegate"
  "descending"
  "do"
  "else"
  "enum"
  "equals"
  "event"
  "explicit"
  "extern"
  "finally"
  "fixed"
  "for"
  "foreach"
  "from"
  "get"
  "global"
  "goto"
  "group"
  "if"
  "implicit"
  "in"
  "init"
  "interface"
  "into"
  "is"
  "join"
  "let"
  "lock"
  "namespace"
  "new"
  "not"
  "notnull"
  "on"
  "operator"
  "or"
  "orderby"
  "out"
  "params"
  "record"
  "ref"
  "remove"
  "return"
  "scoped"
  "select"
  "set"
  "sizeof"
  "stackalloc"
  "struct"
  "switch"
  "this"
  "throw"
  "try"
  "typeof"
  "unchecked"
  "unmanaged"
  "using"
  "when"
  "where"
  "while"
  "with"
  "yield"
] @keyword

[
  "#define"
  "#elif"
  "#else"
  "#endif"
  "#endregion"
  "#error"
  "#if"
  "#line"
  "#nullable"
  "#pragma"
  "#region"
  "#undef"
  "#warning"
] @keyword
