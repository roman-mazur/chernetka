; Highlight query for shell scripts, adapted from tree-sitter-bash's queries/highlights.scm.
;
; Patterns are matched against the capture names in treesitter.go. When two
; patterns capture the exact same node the one listed first here wins, so the
; specific patterns come before the general ones.

; Declarations and calls. A command's name reads as a call, including the
; builtins, as the grammar doesn't tell them apart.

(function_definition
  name: (word) @function)

(command_name) @function.call

; Options passed to a command, like -v or --force.
((command
  argument: (word) @constant)
  (#match? @constant "^-"))

; Variables

(special_variable_name) @constant.builtin

; Positional parameters are special too, though the grammar takes them for names.
((variable_name) @constant.builtin
  (#match? @constant.builtin "^[0-9]+$"))

(variable_name) @property

; Literals

[
  (string)
  (raw_string)
  (ansi_c_string)
  (translated_string)
  (heredoc_body)
  (heredoc_start)
  (heredoc_end)
  (regex)
] @string

[
  (number)
  (file_descriptor)
] @number

; An expansion or a substitution is code, not string content. Its delimiters
; read like escapes, and the rest is reset to plain text before the patterns
; for its own tokens paint over it.
[
  (simple_expansion)
  (expansion)
  (arithmetic_expansion)
  (command_substitution)
  (process_substitution)
] @embedded

(simple_expansion
  "$" @escape)

(expansion
  [
    "${"
    "}"
  ] @escape)

(arithmetic_expansion
  [
    "$(("
    "))"
  ] @escape)

(command_substitution
  [
    "$("
    "`"
    ")"
  ] @escape)

(process_substitution
  [
    "<("
    ">("
    ")"
  ] @escape)

(comment) @comment

; Keywords

(test_operator) @keyword

[
  "case"
  "declare"
  "do"
  "done"
  "elif"
  "else"
  "esac"
  "export"
  "fi"
  "for"
  "function"
  "if"
  "in"
  "local"
  "readonly"
  "select"
  "then"
  "typeset"
  "unset"
  "unsetenv"
  "until"
  "while"
] @keyword
