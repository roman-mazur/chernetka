; Highlight query for makefiles, adapted from tree-sitter-make's queries/highlights.scm.
;
; Patterns are matched against the capture names in treesitter.go. When two
; patterns capture the exact same node the one listed first here wins, so the
; specific patterns come before the general ones. The recipes are left plain,
; but for the variables and the functions they use.

; Rules. A target is declared by its rule, like a function, but the special
; targets, like .PHONY, are instructions to make.

((targets
  (word) @keyword)
  (#match? @keyword "^\\.[A-Z_]+$"))

(targets
  (word) @function)

; The prefixes of a recipe line changing how it is run, like @ to keep it
; quiet.

(recipe_line
  [
    "@"
    "-"
    "+"
  ] @keyword)

; Variables. Like in shell scripts, the names read as properties, and the
; automatic variables, like $@, as builtin constants.

(automatic_variable) @constant.builtin

[
  "VPATH"
  ".RECIPEPREFIX"
] @constant.builtin

(variable_assignment
  name: (word) @property)

(define_directive
  name: (word) @property)

(variable_reference
  (word) @property)

; A reference or a function call is code, not text. Its delimiters read like
; escapes, and the rest is reset to plain text before the patterns for its
; own tokens paint over it.

[
  (variable_reference)
  (substitution_reference)
  (function_call)
  (shell_function)
] @embedded

(function_call
  function: _ @function.call)

(shell_function
  function: _ @function.call)

(variable_reference
  [
    "$"
    "$$"
    "("
    ")"
    "{"
    "}"
  ] @escape)

(substitution_reference
  [
    "$"
    "$$"
    "("
    ")"
    "{"
    "}"
  ] @escape)

(function_call
  [
    "$"
    "$$"
    "("
    ")"
    "{"
    "}"
  ] @escape)

(shell_function
  [
    "$"
    "$$"
    "("
    ")"
    "{"
    "}"
  ] @escape)

; Literals

(string) @string

(escape) @escape

(comment) @comment

; Directives

[
  "ifeq"
  "ifneq"
  "ifdef"
  "ifndef"
  "else"
  "endif"
  "define"
  "endef"
  "include"
  "sinclude"
  "-include"
  "export"
  "unexport"
  "override"
  "private"
  "undefine"
  "vpath"
] @keyword
