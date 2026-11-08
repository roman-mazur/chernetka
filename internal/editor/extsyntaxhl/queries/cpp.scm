; Highlight query for C++, adapted from tree-sitter-cpp's queries/highlights.scm.
; The C++ grammar extends the C one, and like upstream this query only covers
; what C++ adds: it's followed by c.scm, which it takes precedence over.

; Declarations

(function_declarator
  declarator: (qualified_identifier
    name: (identifier) @function))

(function_declarator
  declarator: (qualified_identifier
    name: (qualified_identifier
      name: (identifier) @function)))

; A method declared in its class.
(function_declarator
  declarator: (field_identifier) @function)

(function_declarator
  declarator: (destructor_name
    (identifier) @function))

(function_declarator
  declarator: (qualified_identifier
    name: (destructor_name
      (identifier) @function)))

(function_declarator
  declarator: (operator_name) @function)

; Calls

(call_expression
  function: (qualified_identifier
    name: (identifier) @function.call))

(call_expression
  function: (template_function
    name: (identifier) @function.call))

(call_expression
  function: (qualified_identifier
    name: (template_function
      name: (identifier) @function.call)))

(call_expression
  function: (field_expression
    field: (template_method
      name: (field_identifier) @function.call)))

; Namespaces. A capitalized scope is more likely a class than a namespace.

((namespace_identifier) @type
  (#match? @type "^[A-Z]"))

(namespace_identifier) @module

(using_declaration
  "namespace"
  [
    (identifier) @module
    (qualified_identifier
      name: (identifier) @module)
  ])

; Types

(auto) @type

; Literals

(raw_string_literal) @string

; Keywords

(this) @keyword

[
  "catch"
  "class"
  "co_await"
  "co_return"
  "co_yield"
  "concept"
  "decltype"
  "delete"
  "explicit"
  "final"
  "friend"
  "namespace"
  "new"
  "noexcept"
  "operator"
  "override"
  "private"
  "protected"
  "public"
  "requires"
  "static_assert"
  "template"
  "throw"
  "try"
  "typename"
  "using"
  "virtual"
] @keyword
