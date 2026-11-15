#include "tree_sitter/parser.h"

#if defined(__GNUC__) || defined(__clang__)
#pragma GCC diagnostic ignored "-Wmissing-field-initializers"
#endif

#define LANGUAGE_VERSION 14
#define STATE_COUNT 495
#define LARGE_STATE_COUNT 2
#define SYMBOL_COUNT 143
#define ALIAS_COUNT 0
#define TOKEN_COUNT 74
#define EXTERNAL_TOKEN_COUNT 11
#define FIELD_COUNT 2
#define MAX_ALIAS_SEQUENCE_LENGTH 12
#define PRODUCTION_ID_COUNT 5

enum ts_symbol_identifiers {
  sym_Name = 1,
  anon_sym_LT_QMARK = 2,
  anon_sym_xml = 3,
  anon_sym_QMARK_GT = 4,
  anon_sym_standalone = 5,
  anon_sym_SQUOTE = 6,
  anon_sym_yes = 7,
  anon_sym_no = 8,
  anon_sym_DQUOTE = 9,
  anon_sym_LT_BANG = 10,
  anon_sym_DOCTYPE = 11,
  anon_sym_LBRACK = 12,
  anon_sym_RBRACK = 13,
  anon_sym_GT = 14,
  anon_sym_LT = 15,
  anon_sym_SLASH_GT = 16,
  anon_sym_LT_SLASH = 17,
  anon_sym_RBRACK_RBRACK_GT = 18,
  anon_sym_LT_BANG_LBRACK = 19,
  anon_sym_CDATA = 20,
  anon_sym_xml_DASHstylesheet = 21,
  anon_sym_xml_DASHmodel = 22,
  aux_sym_PseudoAttValue_token1 = 23,
  aux_sym_PseudoAttValue_token2 = 24,
  anon_sym_ELEMENT = 25,
  anon_sym_EMPTY = 26,
  anon_sym_ANY = 27,
  anon_sym_LPAREN = 28,
  anon_sym_POUNDPCDATA = 29,
  anon_sym_PIPE = 30,
  anon_sym_RPAREN = 31,
  anon_sym_STAR = 32,
  anon_sym_QMARK = 33,
  anon_sym_PLUS = 34,
  anon_sym_COMMA = 35,
  anon_sym_ATTLIST = 36,
  sym_TokenizedType = 37,
  anon_sym_NOTATION = 38,
  anon_sym_POUNDREQUIRED = 39,
  anon_sym_POUNDIMPLIED = 40,
  anon_sym_POUNDFIXED = 41,
  anon_sym_ENTITY = 42,
  anon_sym_PERCENT = 43,
  aux_sym_EntityValue_token1 = 44,
  aux_sym_EntityValue_token2 = 45,
  anon_sym_NDATA = 46,
  anon_sym_SEMI = 47,
  sym__S = 48,
  sym_Nmtoken = 49,
  anon_sym_AMP = 50,
  anon_sym_AMP_POUND = 51,
  aux_sym_CharRef_token1 = 52,
  anon_sym_AMP_POUNDx = 53,
  aux_sym_CharRef_token2 = 54,
  anon_sym_SYSTEM = 55,
  anon_sym_PUBLIC = 56,
  aux_sym_SystemLiteral_token1 = 57,
  aux_sym_SystemLiteral_token2 = 58,
  aux_sym_PubidLiteral_token1 = 59,
  aux_sym_PubidLiteral_token2 = 60,
  anon_sym_version = 61,
  sym_VersionNum = 62,
  anon_sym_encoding = 63,
  sym_EncName = 64,
  anon_sym_EQ = 65,
  sym_PITarget = 66,
  sym__pi_content = 67,
  sym_Comment = 68,
  sym_CharData = 69,
  sym_CData = 70,
  sym__start_tag_name = 71,
  sym__end_tag_name = 72,
  sym__erroneous_end_name = 73,
  sym_document = 74,
  sym_prolog = 75,
  sym__Misc = 76,
  sym_XMLDecl = 77,
  sym__SDDecl = 78,
  sym_doctypedecl = 79,
  aux_sym__intSubset = 80,
  sym_element = 81,
  sym_EmptyElemTag = 82,
  sym_Attribute = 83,
  sym_STag = 84,
  sym_ETag = 85,
  sym_content = 86,
  sym_CDSect = 87,
  sym_CDStart = 88,
  sym_StyleSheetPI = 89,
  sym_XmlModelPI = 90,
  sym_PseudoAtt = 91,
  sym_PseudoAttValue = 92,
  sym__markupdecl = 93,
  sym__DeclSep = 94,
  sym_elementdecl = 95,
  sym_contentspec = 96,
  sym_Mixed = 97,
  sym_children = 98,
  sym__cp = 99,
  sym__choice = 100,
  sym_AttlistDecl = 101,
  sym_AttDef = 102,
  sym__AttType = 103,
  sym_StringType = 104,
  sym__EnumeratedType = 105,
  sym_NotationType = 106,
  sym_Enumeration = 107,
  sym_DefaultDecl = 108,
  sym__EntityDecl = 109,
  sym_GEDecl = 110,
  sym_PEDecl = 111,
  sym_EntityValue = 112,
  sym_NDataDecl = 113,
  sym_NotationDecl = 114,
  sym_PEReference = 115,
  sym__Reference = 116,
  sym_EntityRef = 117,
  sym_CharRef = 118,
  sym_AttValue = 119,
  sym_ExternalID = 120,
  sym_PublicID = 121,
  sym_SystemLiteral = 122,
  sym_PubidLiteral = 123,
  sym__VersionInfo = 124,
  sym__EncodingDecl = 125,
  sym_PI = 126,
  sym__Eq = 127,
  aux_sym_document_repeat1 = 128,
  aux_sym_EmptyElemTag_repeat1 = 129,
  aux_sym_content_repeat1 = 130,
  aux_sym_StyleSheetPI_repeat1 = 131,
  aux_sym_PseudoAttValue_repeat1 = 132,
  aux_sym_PseudoAttValue_repeat2 = 133,
  aux_sym_Mixed_repeat1 = 134,
  aux_sym_Mixed_repeat2 = 135,
  aux_sym__choice_repeat1 = 136,
  aux_sym__choice_repeat2 = 137,
  aux_sym_AttlistDecl_repeat1 = 138,
  aux_sym_NotationType_repeat1 = 139,
  aux_sym_Enumeration_repeat1 = 140,
  aux_sym_EntityValue_repeat1 = 141,
  aux_sym_EntityValue_repeat2 = 142,
};

static const char * const ts_symbol_names[] = {
  [ts_builtin_sym_end] = "end",
  [sym_Name] = "Name",
  [anon_sym_LT_QMARK] = "<\?",
  [anon_sym_xml] = "xml",
  [anon_sym_QMARK_GT] = "\?>",
  [anon_sym_standalone] = "standalone",
  [anon_sym_SQUOTE] = "'",
  [anon_sym_yes] = "yes",
  [anon_sym_no] = "no",
  [anon_sym_DQUOTE] = "\"",
  [anon_sym_LT_BANG] = "<!",
  [anon_sym_DOCTYPE] = "DOCTYPE",
  [anon_sym_LBRACK] = "[",
  [anon_sym_RBRACK] = "]",
  [anon_sym_GT] = ">",
  [anon_sym_LT] = "<",
  [anon_sym_SLASH_GT] = "/>",
  [anon_sym_LT_SLASH] = "</",
  [anon_sym_RBRACK_RBRACK_GT] = "]]>",
  [anon_sym_LT_BANG_LBRACK] = "<![",
  [anon_sym_CDATA] = "CDATA",
  [anon_sym_xml_DASHstylesheet] = "xml-stylesheet",
  [anon_sym_xml_DASHmodel] = "xml-model",
  [aux_sym_PseudoAttValue_token1] = "PseudoAttValue_token1",
  [aux_sym_PseudoAttValue_token2] = "PseudoAttValue_token2",
  [anon_sym_ELEMENT] = "ELEMENT",
  [anon_sym_EMPTY] = "EMPTY",
  [anon_sym_ANY] = "ANY",
  [anon_sym_LPAREN] = "(",
  [anon_sym_POUNDPCDATA] = "#PCDATA",
  [anon_sym_PIPE] = "|",
  [anon_sym_RPAREN] = ")",
  [anon_sym_STAR] = "*",
  [anon_sym_QMARK] = "\?",
  [anon_sym_PLUS] = "+",
  [anon_sym_COMMA] = ",",
  [anon_sym_ATTLIST] = "ATTLIST",
  [sym_TokenizedType] = "TokenizedType",
  [anon_sym_NOTATION] = "NOTATION",
  [anon_sym_POUNDREQUIRED] = "#REQUIRED",
  [anon_sym_POUNDIMPLIED] = "#IMPLIED",
  [anon_sym_POUNDFIXED] = "#FIXED",
  [anon_sym_ENTITY] = "ENTITY",
  [anon_sym_PERCENT] = "%",
  [aux_sym_EntityValue_token1] = "EntityValue_token1",
  [aux_sym_EntityValue_token2] = "EntityValue_token2",
  [anon_sym_NDATA] = "NDATA",
  [anon_sym_SEMI] = ";",
  [sym__S] = "_S",
  [sym_Nmtoken] = "Nmtoken",
  [anon_sym_AMP] = "&",
  [anon_sym_AMP_POUND] = "&#",
  [aux_sym_CharRef_token1] = "CharRef_token1",
  [anon_sym_AMP_POUNDx] = "&#x",
  [aux_sym_CharRef_token2] = "CharRef_token2",
  [anon_sym_SYSTEM] = "SYSTEM",
  [anon_sym_PUBLIC] = "PUBLIC",
  [aux_sym_SystemLiteral_token1] = "URI",
  [aux_sym_SystemLiteral_token2] = "URI",
  [aux_sym_PubidLiteral_token1] = "PubidLiteral_token1",
  [aux_sym_PubidLiteral_token2] = "PubidLiteral_token2",
  [anon_sym_version] = "version",
  [sym_VersionNum] = "VersionNum",
  [anon_sym_encoding] = "encoding",
  [sym_EncName] = "EncName",
  [anon_sym_EQ] = "=",
  [sym_PITarget] = "PITarget",
  [sym__pi_content] = "_pi_content",
  [sym_Comment] = "Comment",
  [sym_CharData] = "CharData",
  [sym_CData] = "CData",
  [sym__start_tag_name] = "Name",
  [sym__end_tag_name] = "Name",
  [sym__erroneous_end_name] = "_erroneous_end_name",
  [sym_document] = "document",
  [sym_prolog] = "prolog",
  [sym__Misc] = "_Misc",
  [sym_XMLDecl] = "XMLDecl",
  [sym__SDDecl] = "_SDDecl",
  [sym_doctypedecl] = "doctypedecl",
  [aux_sym__intSubset] = "_intSubset",
  [sym_element] = "element",
  [sym_EmptyElemTag] = "EmptyElemTag",
  [sym_Attribute] = "Attribute",
  [sym_STag] = "STag",
  [sym_ETag] = "ETag",
  [sym_content] = "content",
  [sym_CDSect] = "CDSect",
  [sym_CDStart] = "CDStart",
  [sym_StyleSheetPI] = "StyleSheetPI",
  [sym_XmlModelPI] = "XmlModelPI",
  [sym_PseudoAtt] = "PseudoAtt",
  [sym_PseudoAttValue] = "PseudoAttValue",
  [sym__markupdecl] = "_markupdecl",
  [sym__DeclSep] = "_DeclSep",
  [sym_elementdecl] = "elementdecl",
  [sym_contentspec] = "contentspec",
  [sym_Mixed] = "Mixed",
  [sym_children] = "children",
  [sym__cp] = "_cp",
  [sym__choice] = "_choice",
  [sym_AttlistDecl] = "AttlistDecl",
  [sym_AttDef] = "AttDef",
  [sym__AttType] = "_AttType",
  [sym_StringType] = "StringType",
  [sym__EnumeratedType] = "_EnumeratedType",
  [sym_NotationType] = "NotationType",
  [sym_Enumeration] = "Enumeration",
  [sym_DefaultDecl] = "DefaultDecl",
  [sym__EntityDecl] = "_EntityDecl",
  [sym_GEDecl] = "GEDecl",
  [sym_PEDecl] = "PEDecl",
  [sym_EntityValue] = "EntityValue",
  [sym_NDataDecl] = "NDataDecl",
  [sym_NotationDecl] = "NotationDecl",
  [sym_PEReference] = "PEReference",
  [sym__Reference] = "_Reference",
  [sym_EntityRef] = "EntityRef",
  [sym_CharRef] = "CharRef",
  [sym_AttValue] = "AttValue",
  [sym_ExternalID] = "ExternalID",
  [sym_PublicID] = "PublicID",
  [sym_SystemLiteral] = "SystemLiteral",
  [sym_PubidLiteral] = "PubidLiteral",
  [sym__VersionInfo] = "_VersionInfo",
  [sym__EncodingDecl] = "_EncodingDecl",
  [sym_PI] = "PI",
  [sym__Eq] = "_Eq",
  [aux_sym_document_repeat1] = "document_repeat1",
  [aux_sym_EmptyElemTag_repeat1] = "EmptyElemTag_repeat1",
  [aux_sym_content_repeat1] = "content_repeat1",
  [aux_sym_StyleSheetPI_repeat1] = "StyleSheetPI_repeat1",
  [aux_sym_PseudoAttValue_repeat1] = "PseudoAttValue_repeat1",
  [aux_sym_PseudoAttValue_repeat2] = "PseudoAttValue_repeat2",
  [aux_sym_Mixed_repeat1] = "Mixed_repeat1",
  [aux_sym_Mixed_repeat2] = "Mixed_repeat2",
  [aux_sym__choice_repeat1] = "_choice_repeat1",
  [aux_sym__choice_repeat2] = "_choice_repeat2",
  [aux_sym_AttlistDecl_repeat1] = "AttlistDecl_repeat1",
  [aux_sym_NotationType_repeat1] = "NotationType_repeat1",
  [aux_sym_Enumeration_repeat1] = "Enumeration_repeat1",
  [aux_sym_EntityValue_repeat1] = "EntityValue_repeat1",
  [aux_sym_EntityValue_repeat2] = "EntityValue_repeat2",
};

static const TSSymbol ts_symbol_map[] = {
  [ts_builtin_sym_end] = ts_builtin_sym_end,
  [sym_Name] = sym_Name,
  [anon_sym_LT_QMARK] = anon_sym_LT_QMARK,
  [anon_sym_xml] = anon_sym_xml,
  [anon_sym_QMARK_GT] = anon_sym_QMARK_GT,
  [anon_sym_standalone] = anon_sym_standalone,
  [anon_sym_SQUOTE] = anon_sym_SQUOTE,
  [anon_sym_yes] = anon_sym_yes,
  [anon_sym_no] = anon_sym_no,
  [anon_sym_DQUOTE] = anon_sym_DQUOTE,
  [anon_sym_LT_BANG] = anon_sym_LT_BANG,
  [anon_sym_DOCTYPE] = anon_sym_DOCTYPE,
  [anon_sym_LBRACK] = anon_sym_LBRACK,
  [anon_sym_RBRACK] = anon_sym_RBRACK,
  [anon_sym_GT] = anon_sym_GT,
  [anon_sym_LT] = anon_sym_LT,
  [anon_sym_SLASH_GT] = anon_sym_SLASH_GT,
  [anon_sym_LT_SLASH] = anon_sym_LT_SLASH,
  [anon_sym_RBRACK_RBRACK_GT] = anon_sym_RBRACK_RBRACK_GT,
  [anon_sym_LT_BANG_LBRACK] = anon_sym_LT_BANG_LBRACK,
  [anon_sym_CDATA] = anon_sym_CDATA,
  [anon_sym_xml_DASHstylesheet] = anon_sym_xml_DASHstylesheet,
  [anon_sym_xml_DASHmodel] = anon_sym_xml_DASHmodel,
  [aux_sym_PseudoAttValue_token1] = aux_sym_PseudoAttValue_token1,
  [aux_sym_PseudoAttValue_token2] = aux_sym_PseudoAttValue_token2,
  [anon_sym_ELEMENT] = anon_sym_ELEMENT,
  [anon_sym_EMPTY] = anon_sym_EMPTY,
  [anon_sym_ANY] = anon_sym_ANY,
  [anon_sym_LPAREN] = anon_sym_LPAREN,
  [anon_sym_POUNDPCDATA] = anon_sym_POUNDPCDATA,
  [anon_sym_PIPE] = anon_sym_PIPE,
  [anon_sym_RPAREN] = anon_sym_RPAREN,
  [anon_sym_STAR] = anon_sym_STAR,
  [anon_sym_QMARK] = anon_sym_QMARK,
  [anon_sym_PLUS] = anon_sym_PLUS,
  [anon_sym_COMMA] = anon_sym_COMMA,
  [anon_sym_ATTLIST] = anon_sym_ATTLIST,
  [sym_TokenizedType] = sym_TokenizedType,
  [anon_sym_NOTATION] = anon_sym_NOTATION,
  [anon_sym_POUNDREQUIRED] = anon_sym_POUNDREQUIRED,
  [anon_sym_POUNDIMPLIED] = anon_sym_POUNDIMPLIED,
  [anon_sym_POUNDFIXED] = anon_sym_POUNDFIXED,
  [anon_sym_ENTITY] = anon_sym_ENTITY,
  [anon_sym_PERCENT] = anon_sym_PERCENT,
  [aux_sym_EntityValue_token1] = aux_sym_EntityValue_token1,
  [aux_sym_EntityValue_token2] = aux_sym_EntityValue_token2,
  [anon_sym_NDATA] = anon_sym_NDATA,
  [anon_sym_SEMI] = anon_sym_SEMI,
  [sym__S] = sym__S,
  [sym_Nmtoken] = sym_Nmtoken,
  [anon_sym_AMP] = anon_sym_AMP,
  [anon_sym_AMP_POUND] = anon_sym_AMP_POUND,
  [aux_sym_CharRef_token1] = aux_sym_CharRef_token1,
  [anon_sym_AMP_POUNDx] = anon_sym_AMP_POUNDx,
  [aux_sym_CharRef_token2] = aux_sym_CharRef_token2,
  [anon_sym_SYSTEM] = anon_sym_SYSTEM,
  [anon_sym_PUBLIC] = anon_sym_PUBLIC,
  [aux_sym_SystemLiteral_token1] = aux_sym_SystemLiteral_token1,
  [aux_sym_SystemLiteral_token2] = aux_sym_SystemLiteral_token1,
  [aux_sym_PubidLiteral_token1] = aux_sym_PubidLiteral_token1,
  [aux_sym_PubidLiteral_token2] = aux_sym_PubidLiteral_token2,
  [anon_sym_version] = anon_sym_version,
  [sym_VersionNum] = sym_VersionNum,
  [anon_sym_encoding] = anon_sym_encoding,
  [sym_EncName] = sym_EncName,
  [anon_sym_EQ] = anon_sym_EQ,
  [sym_PITarget] = sym_PITarget,
  [sym__pi_content] = sym__pi_content,
  [sym_Comment] = sym_Comment,
  [sym_CharData] = sym_CharData,
  [sym_CData] = sym_CData,
  [sym__start_tag_name] = sym_Name,
  [sym__end_tag_name] = sym_Name,
  [sym__erroneous_end_name] = sym__erroneous_end_name,
  [sym_document] = sym_document,
  [sym_prolog] = sym_prolog,
  [sym__Misc] = sym__Misc,
  [sym_XMLDecl] = sym_XMLDecl,
  [sym__SDDecl] = sym__SDDecl,
  [sym_doctypedecl] = sym_doctypedecl,
  [aux_sym__intSubset] = aux_sym__intSubset,
  [sym_element] = sym_element,
  [sym_EmptyElemTag] = sym_EmptyElemTag,
  [sym_Attribute] = sym_Attribute,
  [sym_STag] = sym_STag,
  [sym_ETag] = sym_ETag,
  [sym_content] = sym_content,
  [sym_CDSect] = sym_CDSect,
  [sym_CDStart] = sym_CDStart,
  [sym_StyleSheetPI] = sym_StyleSheetPI,
  [sym_XmlModelPI] = sym_XmlModelPI,
  [sym_PseudoAtt] = sym_PseudoAtt,
  [sym_PseudoAttValue] = sym_PseudoAttValue,
  [sym__markupdecl] = sym__markupdecl,
  [sym__DeclSep] = sym__DeclSep,
  [sym_elementdecl] = sym_elementdecl,
  [sym_contentspec] = sym_contentspec,
  [sym_Mixed] = sym_Mixed,
  [sym_children] = sym_children,
  [sym__cp] = sym__cp,
  [sym__choice] = sym__choice,
  [sym_AttlistDecl] = sym_AttlistDecl,
  [sym_AttDef] = sym_AttDef,
  [sym__AttType] = sym__AttType,
  [sym_StringType] = sym_StringType,
  [sym__EnumeratedType] = sym__EnumeratedType,
  [sym_NotationType] = sym_NotationType,
  [sym_Enumeration] = sym_Enumeration,
  [sym_DefaultDecl] = sym_DefaultDecl,
  [sym__EntityDecl] = sym__EntityDecl,
  [sym_GEDecl] = sym_GEDecl,
  [sym_PEDecl] = sym_PEDecl,
  [sym_EntityValue] = sym_EntityValue,
  [sym_NDataDecl] = sym_NDataDecl,
  [sym_NotationDecl] = sym_NotationDecl,
  [sym_PEReference] = sym_PEReference,
  [sym__Reference] = sym__Reference,
  [sym_EntityRef] = sym_EntityRef,
  [sym_CharRef] = sym_CharRef,
  [sym_AttValue] = sym_AttValue,
  [sym_ExternalID] = sym_ExternalID,
  [sym_PublicID] = sym_PublicID,
  [sym_SystemLiteral] = sym_SystemLiteral,
  [sym_PubidLiteral] = sym_PubidLiteral,
  [sym__VersionInfo] = sym__VersionInfo,
  [sym__EncodingDecl] = sym__EncodingDecl,
  [sym_PI] = sym_PI,
  [sym__Eq] = sym__Eq,
  [aux_sym_document_repeat1] = aux_sym_document_repeat1,
  [aux_sym_EmptyElemTag_repeat1] = aux_sym_EmptyElemTag_repeat1,
  [aux_sym_content_repeat1] = aux_sym_content_repeat1,
  [aux_sym_StyleSheetPI_repeat1] = aux_sym_StyleSheetPI_repeat1,
  [aux_sym_PseudoAttValue_repeat1] = aux_sym_PseudoAttValue_repeat1,
  [aux_sym_PseudoAttValue_repeat2] = aux_sym_PseudoAttValue_repeat2,
  [aux_sym_Mixed_repeat1] = aux_sym_Mixed_repeat1,
  [aux_sym_Mixed_repeat2] = aux_sym_Mixed_repeat2,
  [aux_sym__choice_repeat1] = aux_sym__choice_repeat1,
  [aux_sym__choice_repeat2] = aux_sym__choice_repeat2,
  [aux_sym_AttlistDecl_repeat1] = aux_sym_AttlistDecl_repeat1,
  [aux_sym_NotationType_repeat1] = aux_sym_NotationType_repeat1,
  [aux_sym_Enumeration_repeat1] = aux_sym_Enumeration_repeat1,
  [aux_sym_EntityValue_repeat1] = aux_sym_EntityValue_repeat1,
  [aux_sym_EntityValue_repeat2] = aux_sym_EntityValue_repeat2,
};

static const TSSymbolMetadata ts_symbol_metadata[] = {
  [ts_builtin_sym_end] = {
    .visible = false,
    .named = true,
  },
  [sym_Name] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_LT_QMARK] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_xml] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_QMARK_GT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_standalone] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_SQUOTE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_yes] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_no] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DQUOTE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LT_BANG] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DOCTYPE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LBRACK] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_RBRACK] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_GT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_SLASH_GT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LT_SLASH] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_RBRACK_RBRACK_GT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LT_BANG_LBRACK] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_CDATA] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_xml_DASHstylesheet] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_xml_DASHmodel] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_PseudoAttValue_token1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_PseudoAttValue_token2] = {
    .visible = false,
    .named = false,
  },
  [anon_sym_ELEMENT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_EMPTY] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_ANY] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LPAREN] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_POUNDPCDATA] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PIPE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_RPAREN] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_STAR] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_QMARK] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PLUS] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_COMMA] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_ATTLIST] = {
    .visible = true,
    .named = false,
  },
  [sym_TokenizedType] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_NOTATION] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_POUNDREQUIRED] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_POUNDIMPLIED] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_POUNDFIXED] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_ENTITY] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PERCENT] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_EntityValue_token1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_EntityValue_token2] = {
    .visible = false,
    .named = false,
  },
  [anon_sym_NDATA] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_SEMI] = {
    .visible = true,
    .named = false,
  },
  [sym__S] = {
    .visible = false,
    .named = true,
  },
  [sym_Nmtoken] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_AMP] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_AMP_POUND] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_CharRef_token1] = {
    .visible = false,
    .named = false,
  },
  [anon_sym_AMP_POUNDx] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_CharRef_token2] = {
    .visible = false,
    .named = false,
  },
  [anon_sym_SYSTEM] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PUBLIC] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_SystemLiteral_token1] = {
    .visible = true,
    .named = true,
  },
  [aux_sym_SystemLiteral_token2] = {
    .visible = true,
    .named = true,
  },
  [aux_sym_PubidLiteral_token1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_PubidLiteral_token2] = {
    .visible = false,
    .named = false,
  },
  [anon_sym_version] = {
    .visible = true,
    .named = false,
  },
  [sym_VersionNum] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_encoding] = {
    .visible = true,
    .named = false,
  },
  [sym_EncName] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_EQ] = {
    .visible = true,
    .named = false,
  },
  [sym_PITarget] = {
    .visible = true,
    .named = true,
  },
  [sym__pi_content] = {
    .visible = false,
    .named = true,
  },
  [sym_Comment] = {
    .visible = true,
    .named = true,
  },
  [sym_CharData] = {
    .visible = true,
    .named = true,
  },
  [sym_CData] = {
    .visible = true,
    .named = true,
  },
  [sym__start_tag_name] = {
    .visible = true,
    .named = true,
  },
  [sym__end_tag_name] = {
    .visible = true,
    .named = true,
  },
  [sym__erroneous_end_name] = {
    .visible = false,
    .named = true,
  },
  [sym_document] = {
    .visible = true,
    .named = true,
  },
  [sym_prolog] = {
    .visible = true,
    .named = true,
  },
  [sym__Misc] = {
    .visible = false,
    .named = true,
  },
  [sym_XMLDecl] = {
    .visible = true,
    .named = true,
  },
  [sym__SDDecl] = {
    .visible = false,
    .named = true,
  },
  [sym_doctypedecl] = {
    .visible = true,
    .named = true,
  },
  [aux_sym__intSubset] = {
    .visible = false,
    .named = false,
  },
  [sym_element] = {
    .visible = true,
    .named = true,
  },
  [sym_EmptyElemTag] = {
    .visible = true,
    .named = true,
  },
  [sym_Attribute] = {
    .visible = true,
    .named = true,
  },
  [sym_STag] = {
    .visible = true,
    .named = true,
  },
  [sym_ETag] = {
    .visible = true,
    .named = true,
  },
  [sym_content] = {
    .visible = true,
    .named = true,
  },
  [sym_CDSect] = {
    .visible = true,
    .named = true,
  },
  [sym_CDStart] = {
    .visible = true,
    .named = true,
  },
  [sym_StyleSheetPI] = {
    .visible = true,
    .named = true,
  },
  [sym_XmlModelPI] = {
    .visible = true,
    .named = true,
  },
  [sym_PseudoAtt] = {
    .visible = true,
    .named = true,
  },
  [sym_PseudoAttValue] = {
    .visible = true,
    .named = true,
  },
  [sym__markupdecl] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym__DeclSep] = {
    .visible = false,
    .named = true,
  },
  [sym_elementdecl] = {
    .visible = true,
    .named = true,
  },
  [sym_contentspec] = {
    .visible = true,
    .named = true,
  },
  [sym_Mixed] = {
    .visible = true,
    .named = true,
  },
  [sym_children] = {
    .visible = true,
    .named = true,
  },
  [sym__cp] = {
    .visible = false,
    .named = true,
  },
  [sym__choice] = {
    .visible = false,
    .named = true,
  },
  [sym_AttlistDecl] = {
    .visible = true,
    .named = true,
  },
  [sym_AttDef] = {
    .visible = true,
    .named = true,
  },
  [sym__AttType] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym_StringType] = {
    .visible = true,
    .named = true,
  },
  [sym__EnumeratedType] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym_NotationType] = {
    .visible = true,
    .named = true,
  },
  [sym_Enumeration] = {
    .visible = true,
    .named = true,
  },
  [sym_DefaultDecl] = {
    .visible = true,
    .named = true,
  },
  [sym__EntityDecl] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym_GEDecl] = {
    .visible = true,
    .named = true,
  },
  [sym_PEDecl] = {
    .visible = true,
    .named = true,
  },
  [sym_EntityValue] = {
    .visible = true,
    .named = true,
  },
  [sym_NDataDecl] = {
    .visible = true,
    .named = true,
  },
  [sym_NotationDecl] = {
    .visible = true,
    .named = true,
  },
  [sym_PEReference] = {
    .visible = true,
    .named = true,
  },
  [sym__Reference] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym_EntityRef] = {
    .visible = true,
    .named = true,
  },
  [sym_CharRef] = {
    .visible = true,
    .named = true,
  },
  [sym_AttValue] = {
    .visible = true,
    .named = true,
  },
  [sym_ExternalID] = {
    .visible = true,
    .named = true,
  },
  [sym_PublicID] = {
    .visible = true,
    .named = true,
  },
  [sym_SystemLiteral] = {
    .visible = true,
    .named = true,
  },
  [sym_PubidLiteral] = {
    .visible = true,
    .named = true,
  },
  [sym__VersionInfo] = {
    .visible = false,
    .named = true,
  },
  [sym__EncodingDecl] = {
    .visible = false,
    .named = true,
  },
  [sym_PI] = {
    .visible = true,
    .named = true,
  },
  [sym__Eq] = {
    .visible = false,
    .named = true,
  },
  [aux_sym_document_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_EmptyElemTag_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_content_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_StyleSheetPI_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_PseudoAttValue_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_PseudoAttValue_repeat2] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_Mixed_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_Mixed_repeat2] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__choice_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__choice_repeat2] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_AttlistDecl_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_NotationType_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_Enumeration_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_EntityValue_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_EntityValue_repeat2] = {
    .visible = false,
    .named = false,
  },
};

enum ts_field_identifiers {
  field_content = 1,
  field_root = 2,
};

static const char * const ts_field_names[] = {
  [0] = NULL,
  [field_content] = "content",
  [field_root] = "root",
};

static const TSFieldMapSlice ts_field_map_slices[PRODUCTION_ID_COUNT] = {
  [1] = {.index = 0, .length = 1},
  [2] = {.index = 1, .length = 1},
  [3] = {.index = 2, .length = 1},
  [4] = {.index = 3, .length = 1},
};

static const TSFieldMapEntry ts_field_map_entries[] = {
  [0] =
    {field_root, 0},
  [1] =
    {field_root, 1},
  [2] =
    {field_root, 2},
  [3] =
    {field_content, 1},
};

static const TSSymbol ts_alias_sequences[PRODUCTION_ID_COUNT][MAX_ALIAS_SEQUENCE_LENGTH] = {
  [0] = {0},
};

static const uint16_t ts_non_terminal_alias_map[] = {
  0,
};

// Filled from tables.zlib by grammar.go, see fetchgrammar.
static TSStateId ts_primary_state_ids[STATE_COUNT];

static TSCharacterRange aux_sym_PubidLiteral_token1_character_set_1[] = {
  {'\n', '\n'}, {'\r', '\r'}, {' ', '!'}, {'#', '%'}, {'\'', ';'}, {'=', '='}, {'?', 'Z'}, {'_', '_'},
  {'a', 'z'},
};

static TSCharacterRange aux_sym_PubidLiteral_token2_character_set_1[] = {
  {'\n', '\n'}, {'\r', '\r'}, {' ', '!'}, {'#', '%'}, {'(', ';'}, {'=', '='}, {'?', 'Z'}, {'_', '_'},
  {'a', 'z'},
};

static bool ts_lex(TSLexer *lexer, TSStateId state) {
  START_LEXER();
  eof = lexer->eof(lexer);
  switch (state) {
    case 0:
      if (eof) ADVANCE(42);
      ADVANCE_MAP(
        '"', 46,
        '#', 64,
        '%', 91,
        '&', 132,
        '\'', 45,
        '(', 74,
        ')', 77,
        '*', 78,
        '+', 80,
        ',', 81,
        '/', 62,
        '1', 61,
        ';', 94,
        '<', 53,
        '=', 143,
        '>', 52,
        '?', 79,
        'E', 66,
        'I', 63,
        'N', 65,
        '[', 49,
        ']', 51,
        '_', 72,
        '|', 76,
        '\t', 69,
        '\n', 69,
        '\r', 69,
        ' ', 69,
        '-', 71,
        '.', 71,
        ':', 71,
        0xb7, 71,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(70);
      if (('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(67);
      if (('G' <= lookahead && lookahead <= 'Z') ||
          ('g' <= lookahead && lookahead <= 'z')) ADVANCE(68);
      if (lookahead != 0) ADVANCE(60);
      END_STATE();
    case 1:
      ADVANCE_MAP(
        '"', 46,
        '#', 28,
        '%', 91,
        '\'', 45,
        '(', 74,
        ')', 77,
        '*', 78,
        '+', 80,
        ',', 81,
        '/', 7,
        '>', 52,
        '?', 79,
        '[', 49,
        ']', 36,
        '|', 76,
        '\t', 95,
        '\n', 95,
        '\r', 95,
        ' ', 95,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(134);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(127);
      END_STATE();
    case 2:
      if (lookahead == '"') ADVANCE(46);
      if (lookahead == '%') ADVANCE(91);
      if (lookahead == '&') ADVANCE(132);
      if (lookahead != 0 &&
          lookahead != '<') ADVANCE(92);
      END_STATE();
    case 3:
      if (lookahead == '"') ADVANCE(46);
      if (lookahead == '&') ADVANCE(132);
      if (lookahead != 0 &&
          lookahead != '<') ADVANCE(60);
      END_STATE();
    case 4:
      if (lookahead == '%') ADVANCE(91);
      if (lookahead == '&') ADVANCE(132);
      if (lookahead == '\'') ADVANCE(45);
      if (lookahead == '<') ADVANCE(55);
      if (lookahead != 0) ADVANCE(93);
      END_STATE();
    case 5:
      if (lookahead == '&') ADVANCE(132);
      if (lookahead == '\'') ADVANCE(45);
      if (lookahead != 0 &&
          lookahead != '<') ADVANCE(73);
      END_STATE();
    case 6:
      if (lookahead == '.') ADVANCE(38);
      END_STATE();
    case 7:
      if (lookahead == '>') ADVANCE(56);
      END_STATE();
    case 8:
      if (lookahead == '>') ADVANCE(58);
      END_STATE();
    case 9:
      if (lookahead == '>') ADVANCE(44);
      END_STATE();
    case 10:
      if (lookahead == '?') ADVANCE(9);
      if (lookahead == '\t' ||
          lookahead == '\n' ||
          lookahead == '\r' ||
          lookahead == ' ') ADVANCE(95);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(127);
      END_STATE();
    case 11:
      if (lookahead == 'A') ADVANCE(32);
      END_STATE();
    case 12:
      if (lookahead == 'A') ADVANCE(75);
      END_STATE();
    case 13:
      if (lookahead == 'C') ADVANCE(14);
      END_STATE();
    case 14:
      if (lookahead == 'D') ADVANCE(11);
      END_STATE();
    case 15:
      if (lookahead == 'D') ADVANCE(90);
      END_STATE();
    case 16:
      if (lookahead == 'D') ADVANCE(89);
      END_STATE();
    case 17:
      if (lookahead == 'D') ADVANCE(88);
      END_STATE();
    case 18:
      if (lookahead == 'E') ADVANCE(30);
      END_STATE();
    case 19:
      if (lookahead == 'E') ADVANCE(15);
      END_STATE();
    case 20:
      if (lookahead == 'E') ADVANCE(16);
      END_STATE();
    case 21:
      if (lookahead == 'E') ADVANCE(17);
      END_STATE();
    case 22:
      if (lookahead == 'F') ADVANCE(23);
      if (lookahead == 'I') ADVANCE(27);
      if (lookahead == 'R') ADVANCE(18);
      END_STATE();
    case 23:
      if (lookahead == 'I') ADVANCE(34);
      END_STATE();
    case 24:
      if (lookahead == 'I') ADVANCE(31);
      END_STATE();
    case 25:
      if (lookahead == 'I') ADVANCE(20);
      END_STATE();
    case 26:
      if (lookahead == 'L') ADVANCE(25);
      END_STATE();
    case 27:
      if (lookahead == 'M') ADVANCE(29);
      END_STATE();
    case 28:
      if (lookahead == 'P') ADVANCE(13);
      END_STATE();
    case 29:
      if (lookahead == 'P') ADVANCE(26);
      END_STATE();
    case 30:
      if (lookahead == 'Q') ADVANCE(33);
      END_STATE();
    case 31:
      if (lookahead == 'R') ADVANCE(21);
      END_STATE();
    case 32:
      if (lookahead == 'T') ADVANCE(12);
      END_STATE();
    case 33:
      if (lookahead == 'U') ADVANCE(24);
      END_STATE();
    case 34:
      if (lookahead == 'X') ADVANCE(19);
      END_STATE();
    case 35:
      if (lookahead == '[') ADVANCE(59);
      END_STATE();
    case 36:
      if (lookahead == ']') ADVANCE(8);
      END_STATE();
    case 37:
      if (lookahead == '\t' ||
          lookahead == '\n' ||
          lookahead == '\r' ||
          lookahead == ' ') ADVANCE(95);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= ':') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(131);
      END_STATE();
    case 38:
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(141);
      END_STATE();
    case 39:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(136);
      END_STATE();
    case 40:
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(142);
      END_STATE();
    case 41:
      if (eof) ADVANCE(42);
      ADVANCE_MAP(
        '"', 46,
        '#', 22,
        '%', 91,
        '\'', 45,
        '(', 74,
        ')', 77,
        '*', 78,
        '+', 80,
        ',', 81,
        '/', 7,
        '1', 6,
        ';', 94,
        '<', 54,
        '=', 143,
        '>', 52,
        '?', 79,
        'E', 113,
        'I', 96,
        'N', 111,
        '[', 49,
        ']', 50,
        '|', 76,
        '\t', 95,
        '\n', 95,
        '\r', 95,
        ' ', 95,
      );
      if (('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(127);
      END_STATE();
    case 42:
      ACCEPT_TOKEN(ts_builtin_sym_end);
      END_STATE();
    case 43:
      ACCEPT_TOKEN(anon_sym_LT_QMARK);
      END_STATE();
    case 44:
      ACCEPT_TOKEN(anon_sym_QMARK_GT);
      END_STATE();
    case 45:
      ACCEPT_TOKEN(anon_sym_SQUOTE);
      END_STATE();
    case 46:
      ACCEPT_TOKEN(anon_sym_DQUOTE);
      END_STATE();
    case 47:
      ACCEPT_TOKEN(anon_sym_LT_BANG);
      END_STATE();
    case 48:
      ACCEPT_TOKEN(anon_sym_LT_BANG);
      if (lookahead == '[') ADVANCE(59);
      END_STATE();
    case 49:
      ACCEPT_TOKEN(anon_sym_LBRACK);
      END_STATE();
    case 50:
      ACCEPT_TOKEN(anon_sym_RBRACK);
      END_STATE();
    case 51:
      ACCEPT_TOKEN(anon_sym_RBRACK);
      if (lookahead == ']') ADVANCE(8);
      END_STATE();
    case 52:
      ACCEPT_TOKEN(anon_sym_GT);
      END_STATE();
    case 53:
      ACCEPT_TOKEN(anon_sym_LT);
      if (lookahead == '!') ADVANCE(48);
      if (lookahead == '/') ADVANCE(57);
      if (lookahead == '?') ADVANCE(43);
      END_STATE();
    case 54:
      ACCEPT_TOKEN(anon_sym_LT);
      if (lookahead == '!') ADVANCE(47);
      if (lookahead == '?') ADVANCE(43);
      END_STATE();
    case 55:
      ACCEPT_TOKEN(anon_sym_LT);
      if (lookahead == '!') ADVANCE(35);
      if (lookahead == '/') ADVANCE(57);
      if (lookahead == '?') ADVANCE(43);
      END_STATE();
    case 56:
      ACCEPT_TOKEN(anon_sym_SLASH_GT);
      END_STATE();
    case 57:
      ACCEPT_TOKEN(anon_sym_LT_SLASH);
      END_STATE();
    case 58:
      ACCEPT_TOKEN(anon_sym_RBRACK_RBRACK_GT);
      END_STATE();
    case 59:
      ACCEPT_TOKEN(anon_sym_LT_BANG_LBRACK);
      END_STATE();
    case 60:
      ACCEPT_TOKEN(aux_sym_PseudoAttValue_token1);
      END_STATE();
    case 61:
      ACCEPT_TOKEN(aux_sym_PseudoAttValue_token1);
      if (lookahead == '.') ADVANCE(128);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(129);
      if (('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(130);
      if (lookahead == '-' ||
          lookahead == ':' ||
          ('G' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('g' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(131);
      END_STATE();
    case 62:
      ACCEPT_TOKEN(aux_sym_PseudoAttValue_token1);
      if (lookahead == '>') ADVANCE(56);
      END_STATE();
    case 63:
      ACCEPT_TOKEN(aux_sym_PseudoAttValue_token1);
      if (lookahead == 'D') ADVANCE(82);
      if (lookahead == ':' ||
          lookahead == 0xb7) ADVANCE(127);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(126);
      END_STATE();
    case 64:
      ACCEPT_TOKEN(aux_sym_PseudoAttValue_token1);
      if (lookahead == 'F') ADVANCE(23);
      if (lookahead == 'I') ADVANCE(27);
      if (lookahead == 'P') ADVANCE(13);
      if (lookahead == 'R') ADVANCE(18);
      END_STATE();
    case 65:
      ACCEPT_TOKEN(aux_sym_PseudoAttValue_token1);
      if (lookahead == 'M') ADVANCE(120);
      if (lookahead == ':' ||
          lookahead == 0xb7) ADVANCE(127);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(126);
      END_STATE();
    case 66:
      ACCEPT_TOKEN(aux_sym_PseudoAttValue_token1);
      if (lookahead == 'N') ADVANCE(119);
      if (lookahead == ':' ||
          lookahead == 0xb7) ADVANCE(127);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(125);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('G' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('g' <= lookahead && lookahead <= 'z')) ADVANCE(126);
      END_STATE();
    case 67:
      ACCEPT_TOKEN(aux_sym_PseudoAttValue_token1);
      if (lookahead == ':' ||
          lookahead == 0xb7) ADVANCE(127);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(125);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('G' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('g' <= lookahead && lookahead <= 'z')) ADVANCE(126);
      END_STATE();
    case 68:
      ACCEPT_TOKEN(aux_sym_PseudoAttValue_token1);
      if (lookahead == ':' ||
          lookahead == 0xb7) ADVANCE(127);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(126);
      END_STATE();
    case 69:
      ACCEPT_TOKEN(aux_sym_PseudoAttValue_token1);
      if (lookahead == '\t' ||
          lookahead == '\n' ||
          lookahead == '\r' ||
          lookahead == ' ') ADVANCE(95);
      END_STATE();
    case 70:
      ACCEPT_TOKEN(aux_sym_PseudoAttValue_token1);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(129);
      if (('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(130);
      if (lookahead == '-' ||
          lookahead == '.' ||
          lookahead == ':' ||
          ('G' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('g' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(131);
      END_STATE();
    case 71:
      ACCEPT_TOKEN(aux_sym_PseudoAttValue_token1);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= ':') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(131);
      END_STATE();
    case 72:
      ACCEPT_TOKEN(aux_sym_PseudoAttValue_token1);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= ':') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(127);
      END_STATE();
    case 73:
      ACCEPT_TOKEN(aux_sym_PseudoAttValue_token2);
      END_STATE();
    case 74:
      ACCEPT_TOKEN(anon_sym_LPAREN);
      END_STATE();
    case 75:
      ACCEPT_TOKEN(anon_sym_POUNDPCDATA);
      END_STATE();
    case 76:
      ACCEPT_TOKEN(anon_sym_PIPE);
      END_STATE();
    case 77:
      ACCEPT_TOKEN(anon_sym_RPAREN);
      END_STATE();
    case 78:
      ACCEPT_TOKEN(anon_sym_STAR);
      END_STATE();
    case 79:
      ACCEPT_TOKEN(anon_sym_QMARK);
      END_STATE();
    case 80:
      ACCEPT_TOKEN(anon_sym_PLUS);
      END_STATE();
    case 81:
      ACCEPT_TOKEN(anon_sym_COMMA);
      END_STATE();
    case 82:
      ACCEPT_TOKEN(sym_TokenizedType);
      if (lookahead == 'R') ADVANCE(97);
      if (lookahead == ':' ||
          lookahead == 0xb7) ADVANCE(127);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(126);
      END_STATE();
    case 83:
      ACCEPT_TOKEN(sym_TokenizedType);
      if (lookahead == 'R') ADVANCE(100);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= ':') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(127);
      END_STATE();
    case 84:
      ACCEPT_TOKEN(sym_TokenizedType);
      if (lookahead == 'S') ADVANCE(86);
      if (lookahead == ':' ||
          lookahead == 0xb7) ADVANCE(127);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(126);
      END_STATE();
    case 85:
      ACCEPT_TOKEN(sym_TokenizedType);
      if (lookahead == 'S') ADVANCE(87);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= ':') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(127);
      END_STATE();
    case 86:
      ACCEPT_TOKEN(sym_TokenizedType);
      if (lookahead == ':' ||
          lookahead == 0xb7) ADVANCE(127);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(126);
      END_STATE();
    case 87:
      ACCEPT_TOKEN(sym_TokenizedType);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= ':') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(127);
      END_STATE();
    case 88:
      ACCEPT_TOKEN(anon_sym_POUNDREQUIRED);
      END_STATE();
    case 89:
      ACCEPT_TOKEN(anon_sym_POUNDIMPLIED);
      END_STATE();
    case 90:
      ACCEPT_TOKEN(anon_sym_POUNDFIXED);
      END_STATE();
    case 91:
      ACCEPT_TOKEN(anon_sym_PERCENT);
      END_STATE();
    case 92:
      ACCEPT_TOKEN(aux_sym_EntityValue_token1);
      END_STATE();
    case 93:
      ACCEPT_TOKEN(aux_sym_EntityValue_token2);
      END_STATE();
    case 94:
      ACCEPT_TOKEN(anon_sym_SEMI);
      END_STATE();
    case 95:
      ACCEPT_TOKEN(sym__S);
      if (lookahead == '\t' ||
          lookahead == '\n' ||
          lookahead == '\r' ||
          lookahead == ' ') ADVANCE(95);
      END_STATE();
    case 96:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'D') ADVANCE(83);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= ':') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(127);
      END_STATE();
    case 97:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'E') ADVANCE(103);
      if (lookahead == ':' ||
          lookahead == 0xb7) ADVANCE(127);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(126);
      END_STATE();
    case 98:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'E') ADVANCE(112);
      if (lookahead == ':' ||
          lookahead == 0xb7) ADVANCE(127);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(126);
      END_STATE();
    case 99:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'E') ADVANCE(117);
      if (lookahead == ':' ||
          lookahead == 0xb7) ADVANCE(127);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(126);
      END_STATE();
    case 100:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'E') ADVANCE(104);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= ':') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(127);
      END_STATE();
    case 101:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'E') ADVANCE(118);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= ':') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(127);
      END_STATE();
    case 102:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'E') ADVANCE(114);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= ':') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(127);
      END_STATE();
    case 103:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'F') ADVANCE(84);
      if (lookahead == ':' ||
          lookahead == 0xb7) ADVANCE(127);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(126);
      END_STATE();
    case 104:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'F') ADVANCE(85);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= ':') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(127);
      END_STATE();
    case 105:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'I') ADVANCE(121);
      if (lookahead == ':' ||
          lookahead == 0xb7) ADVANCE(127);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(126);
      END_STATE();
    case 106:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'I') ADVANCE(99);
      if (lookahead == 'Y') ADVANCE(86);
      if (lookahead == ':' ||
          lookahead == 0xb7) ADVANCE(127);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(126);
      END_STATE();
    case 107:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'I') ADVANCE(124);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= ':') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(127);
      END_STATE();
    case 108:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'I') ADVANCE(101);
      if (lookahead == 'Y') ADVANCE(87);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= ':') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(127);
      END_STATE();
    case 109:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'K') ADVANCE(98);
      if (lookahead == ':' ||
          lookahead == 0xb7) ADVANCE(127);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(126);
      END_STATE();
    case 110:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'K') ADVANCE(102);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= ':') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(127);
      END_STATE();
    case 111:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'M') ADVANCE(123);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= ':') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(127);
      END_STATE();
    case 112:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'N') ADVANCE(84);
      if (lookahead == ':' ||
          lookahead == 0xb7) ADVANCE(127);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(126);
      END_STATE();
    case 113:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'N') ADVANCE(122);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= ':') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(127);
      END_STATE();
    case 114:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'N') ADVANCE(85);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= ':') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(127);
      END_STATE();
    case 115:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'O') ADVANCE(109);
      if (lookahead == ':' ||
          lookahead == 0xb7) ADVANCE(127);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(126);
      END_STATE();
    case 116:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'O') ADVANCE(110);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= ':') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(127);
      END_STATE();
    case 117:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'S') ADVANCE(86);
      if (lookahead == ':' ||
          lookahead == 0xb7) ADVANCE(127);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(126);
      END_STATE();
    case 118:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'S') ADVANCE(87);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= ':') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(127);
      END_STATE();
    case 119:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'T') ADVANCE(105);
      if (lookahead == ':' ||
          lookahead == 0xb7) ADVANCE(127);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(126);
      END_STATE();
    case 120:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'T') ADVANCE(115);
      if (lookahead == ':' ||
          lookahead == 0xb7) ADVANCE(127);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(126);
      END_STATE();
    case 121:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'T') ADVANCE(106);
      if (lookahead == ':' ||
          lookahead == 0xb7) ADVANCE(127);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(126);
      END_STATE();
    case 122:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'T') ADVANCE(107);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= ':') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(127);
      END_STATE();
    case 123:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'T') ADVANCE(116);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= ':') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(127);
      END_STATE();
    case 124:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == 'T') ADVANCE(108);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= ':') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(127);
      END_STATE();
    case 125:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == ':' ||
          lookahead == 0xb7) ADVANCE(127);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(125);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('G' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('g' <= lookahead && lookahead <= 'z')) ADVANCE(126);
      END_STATE();
    case 126:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == ':' ||
          lookahead == 0xb7) ADVANCE(127);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(126);
      END_STATE();
    case 127:
      ACCEPT_TOKEN(sym_Name);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= ':') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(127);
      END_STATE();
    case 128:
      ACCEPT_TOKEN(sym_Nmtoken);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(128);
      if (lookahead == '-' ||
          lookahead == '.' ||
          lookahead == ':' ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(131);
      END_STATE();
    case 129:
      ACCEPT_TOKEN(sym_Nmtoken);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(129);
      if (('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(130);
      if (lookahead == '-' ||
          lookahead == '.' ||
          lookahead == ':' ||
          ('G' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('g' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(131);
      END_STATE();
    case 130:
      ACCEPT_TOKEN(sym_Nmtoken);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(130);
      if (lookahead == '-' ||
          lookahead == '.' ||
          lookahead == ':' ||
          ('G' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('g' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(131);
      END_STATE();
    case 131:
      ACCEPT_TOKEN(sym_Nmtoken);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= ':') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z') ||
          lookahead == 0xb7) ADVANCE(131);
      END_STATE();
    case 132:
      ACCEPT_TOKEN(anon_sym_AMP);
      if (lookahead == '#') ADVANCE(133);
      END_STATE();
    case 133:
      ACCEPT_TOKEN(anon_sym_AMP_POUND);
      if (lookahead == 'x') ADVANCE(135);
      END_STATE();
    case 134:
      ACCEPT_TOKEN(aux_sym_CharRef_token1);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(134);
      END_STATE();
    case 135:
      ACCEPT_TOKEN(anon_sym_AMP_POUNDx);
      END_STATE();
    case 136:
      ACCEPT_TOKEN(aux_sym_CharRef_token2);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(136);
      END_STATE();
    case 137:
      ACCEPT_TOKEN(aux_sym_SystemLiteral_token1);
      if (lookahead != 0 &&
          lookahead != '"') ADVANCE(137);
      END_STATE();
    case 138:
      ACCEPT_TOKEN(aux_sym_SystemLiteral_token2);
      if (lookahead != 0 &&
          lookahead != '\'') ADVANCE(138);
      END_STATE();
    case 139:
      ACCEPT_TOKEN(aux_sym_PubidLiteral_token1);
      if (set_contains(aux_sym_PubidLiteral_token1_character_set_1, 9, lookahead)) ADVANCE(139);
      END_STATE();
    case 140:
      ACCEPT_TOKEN(aux_sym_PubidLiteral_token2);
      if (set_contains(aux_sym_PubidLiteral_token2_character_set_1, 9, lookahead)) ADVANCE(140);
      END_STATE();
    case 141:
      ACCEPT_TOKEN(sym_VersionNum);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(141);
      END_STATE();
    case 142:
      ACCEPT_TOKEN(sym_EncName);
      if (lookahead == '-' ||
          lookahead == '.' ||
          ('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(142);
      END_STATE();
    case 143:
      ACCEPT_TOKEN(anon_sym_EQ);
      END_STATE();
    default:
      return false;
  }
}

static bool ts_lex_keywords(TSLexer *lexer, TSStateId state) {
  START_LEXER();
  eof = lexer->eof(lexer);
  switch (state) {
    case 0:
      ADVANCE_MAP(
        'A', 1,
        'C', 2,
        'D', 3,
        'E', 4,
        'N', 5,
        'P', 6,
        'S', 7,
        'e', 8,
        'n', 9,
        's', 10,
        'v', 11,
        'x', 12,
        'y', 13,
      );
      END_STATE();
    case 1:
      if (lookahead == 'N') ADVANCE(14);
      if (lookahead == 'T') ADVANCE(15);
      END_STATE();
    case 2:
      if (lookahead == 'D') ADVANCE(16);
      END_STATE();
    case 3:
      if (lookahead == 'O') ADVANCE(17);
      END_STATE();
    case 4:
      if (lookahead == 'L') ADVANCE(18);
      if (lookahead == 'M') ADVANCE(19);
      if (lookahead == 'N') ADVANCE(20);
      END_STATE();
    case 5:
      if (lookahead == 'D') ADVANCE(21);
      if (lookahead == 'O') ADVANCE(22);
      END_STATE();
    case 6:
      if (lookahead == 'U') ADVANCE(23);
      END_STATE();
    case 7:
      if (lookahead == 'Y') ADVANCE(24);
      END_STATE();
    case 8:
      if (lookahead == 'n') ADVANCE(25);
      END_STATE();
    case 9:
      if (lookahead == 'o') ADVANCE(26);
      END_STATE();
    case 10:
      if (lookahead == 't') ADVANCE(27);
      END_STATE();
    case 11:
      if (lookahead == 'e') ADVANCE(28);
      END_STATE();
    case 12:
      if (lookahead == 'm') ADVANCE(29);
      END_STATE();
    case 13:
      if (lookahead == 'e') ADVANCE(30);
      END_STATE();
    case 14:
      if (lookahead == 'Y') ADVANCE(31);
      END_STATE();
    case 15:
      if (lookahead == 'T') ADVANCE(32);
      END_STATE();
    case 16:
      if (lookahead == 'A') ADVANCE(33);
      END_STATE();
    case 17:
      if (lookahead == 'C') ADVANCE(34);
      END_STATE();
    case 18:
      if (lookahead == 'E') ADVANCE(35);
      END_STATE();
    case 19:
      if (lookahead == 'P') ADVANCE(36);
      END_STATE();
    case 20:
      if (lookahead == 'T') ADVANCE(37);
      END_STATE();
    case 21:
      if (lookahead == 'A') ADVANCE(38);
      END_STATE();
    case 22:
      if (lookahead == 'T') ADVANCE(39);
      END_STATE();
    case 23:
      if (lookahead == 'B') ADVANCE(40);
      END_STATE();
    case 24:
      if (lookahead == 'S') ADVANCE(41);
      END_STATE();
    case 25:
      if (lookahead == 'c') ADVANCE(42);
      END_STATE();
    case 26:
      ACCEPT_TOKEN(anon_sym_no);
      END_STATE();
    case 27:
      if (lookahead == 'a') ADVANCE(43);
      END_STATE();
    case 28:
      if (lookahead == 'r') ADVANCE(44);
      END_STATE();
    case 29:
      if (lookahead == 'l') ADVANCE(45);
      END_STATE();
    case 30:
      if (lookahead == 's') ADVANCE(46);
      END_STATE();
    case 31:
      ACCEPT_TOKEN(anon_sym_ANY);
      END_STATE();
    case 32:
      if (lookahead == 'L') ADVANCE(47);
      END_STATE();
    case 33:
      if (lookahead == 'T') ADVANCE(48);
      END_STATE();
    case 34:
      if (lookahead == 'T') ADVANCE(49);
      END_STATE();
    case 35:
      if (lookahead == 'M') ADVANCE(50);
      END_STATE();
    case 36:
      if (lookahead == 'T') ADVANCE(51);
      END_STATE();
    case 37:
      if (lookahead == 'I') ADVANCE(52);
      END_STATE();
    case 38:
      if (lookahead == 'T') ADVANCE(53);
      END_STATE();
    case 39:
      if (lookahead == 'A') ADVANCE(54);
      END_STATE();
    case 40:
      if (lookahead == 'L') ADVANCE(55);
      END_STATE();
    case 41:
      if (lookahead == 'T') ADVANCE(56);
      END_STATE();
    case 42:
      if (lookahead == 'o') ADVANCE(57);
      END_STATE();
    case 43:
      if (lookahead == 'n') ADVANCE(58);
      END_STATE();
    case 44:
      if (lookahead == 's') ADVANCE(59);
      END_STATE();
    case 45:
      ACCEPT_TOKEN(anon_sym_xml);
      if (lookahead == '-') ADVANCE(60);
      END_STATE();
    case 46:
      ACCEPT_TOKEN(anon_sym_yes);
      END_STATE();
    case 47:
      if (lookahead == 'I') ADVANCE(61);
      END_STATE();
    case 48:
      if (lookahead == 'A') ADVANCE(62);
      END_STATE();
    case 49:
      if (lookahead == 'Y') ADVANCE(63);
      END_STATE();
    case 50:
      if (lookahead == 'E') ADVANCE(64);
      END_STATE();
    case 51:
      if (lookahead == 'Y') ADVANCE(65);
      END_STATE();
    case 52:
      if (lookahead == 'T') ADVANCE(66);
      END_STATE();
    case 53:
      if (lookahead == 'A') ADVANCE(67);
      END_STATE();
    case 54:
      if (lookahead == 'T') ADVANCE(68);
      END_STATE();
    case 55:
      if (lookahead == 'I') ADVANCE(69);
      END_STATE();
    case 56:
      if (lookahead == 'E') ADVANCE(70);
      END_STATE();
    case 57:
      if (lookahead == 'd') ADVANCE(71);
      END_STATE();
    case 58:
      if (lookahead == 'd') ADVANCE(72);
      END_STATE();
    case 59:
      if (lookahead == 'i') ADVANCE(73);
      END_STATE();
    case 60:
      if (lookahead == 'm') ADVANCE(74);
      if (lookahead == 's') ADVANCE(75);
      END_STATE();
    case 61:
      if (lookahead == 'S') ADVANCE(76);
      END_STATE();
    case 62:
      ACCEPT_TOKEN(anon_sym_CDATA);
      END_STATE();
    case 63:
      if (lookahead == 'P') ADVANCE(77);
      END_STATE();
    case 64:
      if (lookahead == 'N') ADVANCE(78);
      END_STATE();
    case 65:
      ACCEPT_TOKEN(anon_sym_EMPTY);
      END_STATE();
    case 66:
      if (lookahead == 'Y') ADVANCE(79);
      END_STATE();
    case 67:
      ACCEPT_TOKEN(anon_sym_NDATA);
      END_STATE();
    case 68:
      if (lookahead == 'I') ADVANCE(80);
      END_STATE();
    case 69:
      if (lookahead == 'C') ADVANCE(81);
      END_STATE();
    case 70:
      if (lookahead == 'M') ADVANCE(82);
      END_STATE();
    case 71:
      if (lookahead == 'i') ADVANCE(83);
      END_STATE();
    case 72:
      if (lookahead == 'a') ADVANCE(84);
      END_STATE();
    case 73:
      if (lookahead == 'o') ADVANCE(85);
      END_STATE();
    case 74:
      if (lookahead == 'o') ADVANCE(86);
      END_STATE();
    case 75:
      if (lookahead == 't') ADVANCE(87);
      END_STATE();
    case 76:
      if (lookahead == 'T') ADVANCE(88);
      END_STATE();
    case 77:
      if (lookahead == 'E') ADVANCE(89);
      END_STATE();
    case 78:
      if (lookahead == 'T') ADVANCE(90);
      END_STATE();
    case 79:
      ACCEPT_TOKEN(anon_sym_ENTITY);
      END_STATE();
    case 80:
      if (lookahead == 'O') ADVANCE(91);
      END_STATE();
    case 81:
      ACCEPT_TOKEN(anon_sym_PUBLIC);
      END_STATE();
    case 82:
      ACCEPT_TOKEN(anon_sym_SYSTEM);
      END_STATE();
    case 83:
      if (lookahead == 'n') ADVANCE(92);
      END_STATE();
    case 84:
      if (lookahead == 'l') ADVANCE(93);
      END_STATE();
    case 85:
      if (lookahead == 'n') ADVANCE(94);
      END_STATE();
    case 86:
      if (lookahead == 'd') ADVANCE(95);
      END_STATE();
    case 87:
      if (lookahead == 'y') ADVANCE(96);
      END_STATE();
    case 88:
      ACCEPT_TOKEN(anon_sym_ATTLIST);
      END_STATE();
    case 89:
      ACCEPT_TOKEN(anon_sym_DOCTYPE);
      END_STATE();
    case 90:
      ACCEPT_TOKEN(anon_sym_ELEMENT);
      END_STATE();
    case 91:
      if (lookahead == 'N') ADVANCE(97);
      END_STATE();
    case 92:
      if (lookahead == 'g') ADVANCE(98);
      END_STATE();
    case 93:
      if (lookahead == 'o') ADVANCE(99);
      END_STATE();
    case 94:
      ACCEPT_TOKEN(anon_sym_version);
      END_STATE();
    case 95:
      if (lookahead == 'e') ADVANCE(100);
      END_STATE();
    case 96:
      if (lookahead == 'l') ADVANCE(101);
      END_STATE();
    case 97:
      ACCEPT_TOKEN(anon_sym_NOTATION);
      END_STATE();
    case 98:
      ACCEPT_TOKEN(anon_sym_encoding);
      END_STATE();
    case 99:
      if (lookahead == 'n') ADVANCE(102);
      END_STATE();
    case 100:
      if (lookahead == 'l') ADVANCE(103);
      END_STATE();
    case 101:
      if (lookahead == 'e') ADVANCE(104);
      END_STATE();
    case 102:
      if (lookahead == 'e') ADVANCE(105);
      END_STATE();
    case 103:
      ACCEPT_TOKEN(anon_sym_xml_DASHmodel);
      END_STATE();
    case 104:
      if (lookahead == 's') ADVANCE(106);
      END_STATE();
    case 105:
      ACCEPT_TOKEN(anon_sym_standalone);
      END_STATE();
    case 106:
      if (lookahead == 'h') ADVANCE(107);
      END_STATE();
    case 107:
      if (lookahead == 'e') ADVANCE(108);
      END_STATE();
    case 108:
      if (lookahead == 'e') ADVANCE(109);
      END_STATE();
    case 109:
      if (lookahead == 't') ADVANCE(110);
      END_STATE();
    case 110:
      ACCEPT_TOKEN(anon_sym_xml_DASHstylesheet);
      END_STATE();
    default:
      return false;
  }
}

// Filled from tables.zlib by grammar.go, see fetchgrammar.
static TSLexMode ts_lex_modes[STATE_COUNT];

// Filled from tables.zlib by grammar.go, see fetchgrammar.
static uint16_t ts_parse_table[LARGE_STATE_COUNT][SYMBOL_COUNT];

// Filled from tables.zlib by grammar.go, see fetchgrammar.
static uint16_t ts_small_parse_table[5178];

// Filled from tables.zlib by grammar.go, see fetchgrammar.
static uint32_t ts_small_parse_table_map[493];

// Filled from tables.zlib by grammar.go, see fetchgrammar.
static TSParseActionEntry ts_parse_actions[1263];

enum ts_external_scanner_symbol_identifiers {
  ts_external_token_PITarget = 0,
  ts_external_token__pi_content = 1,
  ts_external_token_Comment = 2,
  ts_external_token_CharData = 3,
  ts_external_token_CData = 4,
  ts_external_token_xml_DASHmodel = 5,
  ts_external_token_xml_DASHstylesheet = 6,
  ts_external_token__start_tag_name = 7,
  ts_external_token__end_tag_name = 8,
  ts_external_token__erroneous_end_name = 9,
  ts_external_token_SLASH_GT = 10,
};

static const TSSymbol ts_external_scanner_symbol_map[EXTERNAL_TOKEN_COUNT] = {
  [ts_external_token_PITarget] = sym_PITarget,
  [ts_external_token__pi_content] = sym__pi_content,
  [ts_external_token_Comment] = sym_Comment,
  [ts_external_token_CharData] = sym_CharData,
  [ts_external_token_CData] = sym_CData,
  [ts_external_token_xml_DASHmodel] = anon_sym_xml_DASHmodel,
  [ts_external_token_xml_DASHstylesheet] = anon_sym_xml_DASHstylesheet,
  [ts_external_token__start_tag_name] = sym__start_tag_name,
  [ts_external_token__end_tag_name] = sym__end_tag_name,
  [ts_external_token__erroneous_end_name] = sym__erroneous_end_name,
  [ts_external_token_SLASH_GT] = anon_sym_SLASH_GT,
};

static const bool ts_external_scanner_states[11][EXTERNAL_TOKEN_COUNT] = {
  [1] = {
    [ts_external_token_PITarget] = true,
    [ts_external_token__pi_content] = true,
    [ts_external_token_Comment] = true,
    [ts_external_token_CharData] = true,
    [ts_external_token_CData] = true,
    [ts_external_token_xml_DASHmodel] = true,
    [ts_external_token_xml_DASHstylesheet] = true,
    [ts_external_token__start_tag_name] = true,
    [ts_external_token__end_tag_name] = true,
    [ts_external_token__erroneous_end_name] = true,
    [ts_external_token_SLASH_GT] = true,
  },
  [2] = {
    [ts_external_token_Comment] = true,
  },
  [3] = {
    [ts_external_token_Comment] = true,
    [ts_external_token_CharData] = true,
  },
  [4] = {
    [ts_external_token_SLASH_GT] = true,
  },
  [5] = {
    [ts_external_token_PITarget] = true,
    [ts_external_token_xml_DASHmodel] = true,
    [ts_external_token_xml_DASHstylesheet] = true,
  },
  [6] = {
    [ts_external_token_CData] = true,
  },
  [7] = {
    [ts_external_token__pi_content] = true,
  },
  [8] = {
    [ts_external_token_PITarget] = true,
  },
  [9] = {
    [ts_external_token__end_tag_name] = true,
  },
  [10] = {
    [ts_external_token__start_tag_name] = true,
  },
};

#ifdef __cplusplus
extern "C" {
#endif
void *tree_sitter_xml_external_scanner_create(void);
void tree_sitter_xml_external_scanner_destroy(void *);
bool tree_sitter_xml_external_scanner_scan(void *, TSLexer *, const bool *);
unsigned tree_sitter_xml_external_scanner_serialize(void *, char *);
void tree_sitter_xml_external_scanner_deserialize(void *, const char *, unsigned);

#ifdef TREE_SITTER_HIDE_SYMBOLS
#define TS_PUBLIC
#elif defined(_WIN32)
#define TS_PUBLIC __declspec(dllexport)
#else
#define TS_PUBLIC __attribute__((visibility("default")))
#endif

TS_PUBLIC const TSLanguage *tree_sitter_xml(void) {
  static const TSLanguage language = {
    .version = LANGUAGE_VERSION,
    .symbol_count = SYMBOL_COUNT,
    .alias_count = ALIAS_COUNT,
    .token_count = TOKEN_COUNT,
    .external_token_count = EXTERNAL_TOKEN_COUNT,
    .state_count = STATE_COUNT,
    .large_state_count = LARGE_STATE_COUNT,
    .production_id_count = PRODUCTION_ID_COUNT,
    .field_count = FIELD_COUNT,
    .max_alias_sequence_length = MAX_ALIAS_SEQUENCE_LENGTH,
    .parse_table = &ts_parse_table[0][0],
    .small_parse_table = ts_small_parse_table,
    .small_parse_table_map = ts_small_parse_table_map,
    .parse_actions = ts_parse_actions,
    .symbol_names = ts_symbol_names,
    .field_names = ts_field_names,
    .field_map_slices = ts_field_map_slices,
    .field_map_entries = ts_field_map_entries,
    .symbol_metadata = ts_symbol_metadata,
    .public_symbol_map = ts_symbol_map,
    .alias_map = ts_non_terminal_alias_map,
    .alias_sequences = &ts_alias_sequences[0][0],
    .lex_modes = ts_lex_modes,
    .lex_fn = ts_lex,
    .keyword_lex_fn = ts_lex_keywords,
    .keyword_capture_token = sym_Name,
    .external_scanner = {
      &ts_external_scanner_states[0][0],
      ts_external_scanner_symbol_map,
      tree_sitter_xml_external_scanner_create,
      tree_sitter_xml_external_scanner_destroy,
      tree_sitter_xml_external_scanner_scan,
      tree_sitter_xml_external_scanner_serialize,
      tree_sitter_xml_external_scanner_deserialize,
    },
    .primary_state_ids = ts_primary_state_ids,
  };
  return &language;
}
#ifdef __cplusplus
}
#endif

// Added by fetchgrammar: grammar.go fills the tables above from tables.zlib.

_Static_assert(__BYTE_ORDER__ == __ORDER_LITTLE_ENDIAN__, "the tables in tables.zlib are little endian");

static void *fetchgrammar_table(int i, size_t *size) {
  switch (i) {
    case 0: *size = sizeof(ts_parse_table); return (void *)ts_parse_table;
    case 1: *size = sizeof(ts_small_parse_table); return (void *)ts_small_parse_table;
    case 2: *size = sizeof(ts_small_parse_table_map); return (void *)ts_small_parse_table_map;
    case 3: *size = sizeof(ts_parse_actions); return (void *)ts_parse_actions;
    case 4: *size = sizeof(ts_lex_modes); return (void *)ts_lex_modes;
    case 5: *size = sizeof(ts_primary_state_ids); return (void *)ts_primary_state_ids;
  }
  *size = 0;
  return NULL;
}
