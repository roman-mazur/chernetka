; Highlight query for XML, adapted from tree-sitter-xml's
; queries/xml/highlights.scm.
;
; Patterns are matched against the capture names in extsyntaxhl's treesitter.go.
; Like in HTML, the tags read as keywords and their attributes as properties.
; The declaration and the processing instructions, like <?xml-stylesheet?>,
; are instructions to the processor and read as directives. The text is left
; plain.

[
  (XMLDecl)
  (PI)
  (StyleSheetPI)
  (XmlModelPI)
] @directive

; The document type declaration, which may hold a DTD.

(doctypedecl
  "DOCTYPE" @directive
  (Name) @type)

[
  "ELEMENT"
  "ATTLIST"
  "ENTITY"
  "NOTATION"
  "NDATA"
  "PUBLIC"
  "SYSTEM"
] @keyword

(elementdecl
  (Name) @keyword)

(AttlistDecl
  (Name) @keyword)

(AttDef
  (Name) @property)

(GEDecl
  (Name) @constant)

(PEDecl
  (Name) @constant)

(NotationDecl
  (Name) @constant)

[
  "#PCDATA"
  "EMPTY"
  "ANY"
  (StringType)
  (TokenizedType)
  "#REQUIRED"
  "#IMPLIED"
  "#FIXED"
] @constant

[
  (EntityValue)
  (PubidLiteral)
  (SystemLiteral)
] @string

(Enumeration
  (Nmtoken) @string)

; Elements

(STag
  (Name) @keyword)

(ETag
  (Name) @keyword)

(EmptyElemTag
  (Name) @keyword)

(Attribute
  (Name) @property)

(AttValue) @string

; A reference, like &amp; or &#x20;, reads like an escape, and so do the
; delimiters of a CDATA section, whose text is left plain.

[
  (EntityRef)
  (CharRef)
  (PEReference)
] @escape

(CDSect
  [
    (CDStart)
    "]]>"
  ] @escape)

(Comment) @comment
