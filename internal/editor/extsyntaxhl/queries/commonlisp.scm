; Highlight query for Common Lisp. The grammar comes with no highlight query of
; its own: it parses the syntax of the reader, so the forms are told apart by
; their first symbol, ignoring the case as the reader does.
;
; Patterns are matched against the capture names in treesitter.go. When two
; patterns capture the exact same node the one listed first here wins, so the
; specific patterns come before the general ones.

; A reader conditional, like #+sbcl, is an instruction to the reader.

(include_reader_macro
  marker: _ @directive
  condition: _ @directive)

; Definitions

(defun_header
  keyword: (defun_keyword) @keyword
  function_name: (sym_lit) @function)

(defun_header
  function_name: (package_lit
    symbol: (sym_lit) @function))

; The grammar takes only the lower case words for definitions, so a definition
; like (DEFUN NAME (X) ...) is matched as a list.
((list_lit
  .
  (sym_lit) @keyword
  .
  (sym_lit) @function
  .
  (list_lit
    .
    (sym_lit) @variable))
  (#match? @keyword "(?i)^def(un|macro|generic|method)$"))

((list_lit
  .
  (sym_lit) @keyword
  .
  (sym_lit) @constant)
  (#match? @keyword "(?i)^def(var|parameter|constant)$"))

((list_lit
  .
  (sym_lit) @keyword
  .
  (sym_lit) @type)
  (#match? @keyword "(?i)^def(class|struct|type|condition)$"))

((list_lit
  .
  (sym_lit) @keyword
  .
  [
    (sym_lit)
    (kwd_lit)
    (str_lit)
  ] @module)
  (#match? @keyword "(?i)^(defpackage|in-package)$"))

; The special operators and the standard macros that shape the code read as
; keywords, like the loop clauses.

((list_lit
  .
  (sym_lit) @keyword)
  (#match? @keyword "(?i)^(block|catch|eval-when|flet|function|go|if|labels|let\\*?|load-time-value|locally|macrolet|multiple-value-call|multiple-value-prog1|progn|progv|quote|return-from|setq|symbol-macrolet|tagbody|the|throw|unwind-protect|and|assert|case|ccase|check-type|cond|ctypecase|decf|declaim|declare|defgeneric|defmacro|defmethod|defun|defsetf|define-condition|define-modify-macro|define-symbol-macro|destructuring-bind|do\\*?|do-all-symbols|do-external-symbols|do-symbols|dolist|dotimes|ecase|etypecase|handler-bind|handler-case|ignore-errors|in-package|incf|lambda|loop|multiple-value-bind|multiple-value-list|multiple-value-setq|or|pop|prog1|prog2|push|pushnew|restart-case|return|rotatef|setf|shiftf|typecase|unless|when|with-accessors|with-open-file|with-output-to-string|with-slots)$"))

(loop_macro
  "loop" @keyword)

[
  (for_clause_word)
  (accumulation_verb)
  "for"
  "and"
  "as"
  "with"
  "do"
  "while"
  "until"
  "repeat"
  "when"
  "if"
  "unless"
  "else"
  "then"
  "always"
  "never"
  "thereis"
  "into"
  "finally"
  "initially"
  "return"
] @keyword

; The lambda list keywords, like &optional and &key.

((sym_lit) @keyword
  (#match? @keyword "^&"))

; Bindings. The first element of a list is a name rather than an operator in
; a lambda list and in the binding lists of the forms that bind names. The
; captures starting with _ only select the forms and are not colored.

(defun_header
  lambda_list: (list_lit
    .
    (sym_lit) @variable))

(defun_header
  lambda_list: (list_lit
    (list_lit
      .
      (sym_lit) @variable)))

((list_lit
  .
  (sym_lit) @_form
  .
  (list_lit
    (list_lit
      .
      (sym_lit) @variable)))
  (#match? @_form "(?i)^(let\\*?|symbol-macrolet|do\\*?)$"))

((list_lit
  .
  (sym_lit) @_form
  .
  (list_lit
    (list_lit
      .
      (sym_lit) @function)))
  (#match? @_form "(?i)^(flet|labels|macrolet)$"))

((list_lit
  .
  (sym_lit) @_form
  .
  (list_lit
    .
    (sym_lit) @variable))
  (#match? @_form "(?i)^(lambda|dolist|dotimes|do-symbols|do-external-symbols|do-all-symbols|multiple-value-bind|destructuring-bind|with-.+)$"))

; Calls. The first element of a list is the operator, unless the list is quoted
; data: its first element is reset to plain text.

(quoting_lit
  value: (list_lit
    .
    (sym_lit) @variable))

(syn_quoting_lit
  value: (list_lit
    .
    (sym_lit) @variable))

(list_lit
  .
  (sym_lit) @function.call)

(list_lit
  .
  (package_lit
    symbol: (sym_lit) @function.call))

; A function named by #', like #'car.
(var_quoting_lit
  value: (sym_lit) @function.call)

; Packages

(package_lit
  package: _ @module)

; Constants: nil, t, the keywords, and by convention the +constants+ and the
; *special* variables.

[
  (nil_lit)
  (kwd_lit)
] @constant

((sym_lit) @constant
  (#match? @constant "(?i)^t$|^\\+.+\\+$|^\\*.+\\*$"))

; Literals

[
  (str_lit)
  (char_lit)
] @string

; A format directive, like ~A or ~%, reads like an escape in its string.
(format_specifier) @escape

[
  (num_lit)
  (complex_num_lit)
] @number

[
  (comment)
  (block_comment)
] @comment

; Catch-all symbol rule, kept last so the patterns above take precedence.

(sym_lit) @variable
