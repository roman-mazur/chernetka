#include "tree_sitter/parser.h"

#if defined(__GNUC__) || defined(__clang__)
#pragma GCC diagnostic ignored "-Wmissing-field-initializers"
#endif

#define LANGUAGE_VERSION 14
#define STATE_COUNT 5986
#define LARGE_STATE_COUNT 1167
#define SYMBOL_COUNT 393
#define ALIAS_COUNT 7
#define TOKEN_COUNT 172
#define EXTERNAL_TOKEN_COUNT 10
#define FIELD_COUNT 43
#define MAX_ALIAS_SEQUENCE_LENGTH 10
#define PRODUCTION_ID_COUNT 351

enum ts_symbol_identifiers {
  sym_identifier = 1,
  sym_hash_bang_line = 2,
  anon_sym_export = 3,
  anon_sym_STAR = 4,
  anon_sym_default = 5,
  anon_sym_type = 6,
  anon_sym_EQ = 7,
  anon_sym_as = 8,
  anon_sym_namespace = 9,
  anon_sym_LBRACE = 10,
  anon_sym_COMMA = 11,
  anon_sym_RBRACE = 12,
  anon_sym_typeof = 13,
  anon_sym_import = 14,
  anon_sym_from = 15,
  anon_sym_with = 16,
  anon_sym_assert = 17,
  anon_sym_var = 18,
  anon_sym_let = 19,
  anon_sym_const = 20,
  anon_sym_BANG = 21,
  anon_sym_else = 22,
  anon_sym_if = 23,
  anon_sym_switch = 24,
  anon_sym_for = 25,
  anon_sym_LPAREN = 26,
  anon_sym_SEMI = 27,
  anon_sym_RPAREN = 28,
  anon_sym_await = 29,
  anon_sym_in = 30,
  anon_sym_of = 31,
  anon_sym_while = 32,
  anon_sym_do = 33,
  anon_sym_try = 34,
  anon_sym_break = 35,
  anon_sym_continue = 36,
  anon_sym_debugger = 37,
  anon_sym_return = 38,
  anon_sym_throw = 39,
  anon_sym_COLON = 40,
  anon_sym_case = 41,
  anon_sym_catch = 42,
  anon_sym_finally = 43,
  anon_sym_yield = 44,
  anon_sym_LBRACK = 45,
  anon_sym_RBRACK = 46,
  sym_html_character_reference = 47,
  anon_sym_GT = 48,
  sym_jsx_identifier = 49,
  anon_sym_DOT = 50,
  anon_sym_LT_SLASH = 51,
  anon_sym_SLASH_GT = 52,
  anon_sym_DQUOTE = 53,
  anon_sym_SQUOTE = 54,
  sym_unescaped_double_jsx_string_fragment = 55,
  sym_unescaped_single_jsx_string_fragment = 56,
  anon_sym_class = 57,
  anon_sym_async = 58,
  anon_sym_function = 59,
  anon_sym_EQ_GT = 60,
  anon_sym_QMARK_DOT = 61,
  anon_sym_new = 62,
  anon_sym_using = 63,
  anon_sym_PLUS_EQ = 64,
  anon_sym_DASH_EQ = 65,
  anon_sym_STAR_EQ = 66,
  anon_sym_SLASH_EQ = 67,
  anon_sym_PERCENT_EQ = 68,
  anon_sym_CARET_EQ = 69,
  anon_sym_AMP_EQ = 70,
  anon_sym_PIPE_EQ = 71,
  anon_sym_GT_GT_EQ = 72,
  anon_sym_GT_GT_GT_EQ = 73,
  anon_sym_LT_LT_EQ = 74,
  anon_sym_STAR_STAR_EQ = 75,
  anon_sym_AMP_AMP_EQ = 76,
  anon_sym_PIPE_PIPE_EQ = 77,
  anon_sym_QMARK_QMARK_EQ = 78,
  anon_sym_DOT_DOT_DOT = 79,
  anon_sym_AMP_AMP = 80,
  anon_sym_PIPE_PIPE = 81,
  anon_sym_GT_GT = 82,
  anon_sym_GT_GT_GT = 83,
  anon_sym_LT_LT = 84,
  anon_sym_AMP = 85,
  anon_sym_CARET = 86,
  anon_sym_PIPE = 87,
  anon_sym_PLUS = 88,
  anon_sym_DASH = 89,
  anon_sym_SLASH = 90,
  anon_sym_PERCENT = 91,
  anon_sym_STAR_STAR = 92,
  anon_sym_LT = 93,
  anon_sym_LT_EQ = 94,
  anon_sym_EQ_EQ = 95,
  anon_sym_EQ_EQ_EQ = 96,
  anon_sym_BANG_EQ = 97,
  anon_sym_BANG_EQ_EQ = 98,
  anon_sym_GT_EQ = 99,
  anon_sym_QMARK_QMARK = 100,
  anon_sym_instanceof = 101,
  anon_sym_TILDE = 102,
  anon_sym_void = 103,
  anon_sym_delete = 104,
  anon_sym_PLUS_PLUS = 105,
  anon_sym_DASH_DASH = 106,
  sym_unescaped_double_string_fragment = 107,
  sym_unescaped_single_string_fragment = 108,
  sym_escape_sequence = 109,
  sym_comment = 110,
  anon_sym_BQUOTE = 111,
  anon_sym_DOLLAR_LBRACE = 112,
  anon_sym_SLASH2 = 113,
  sym_regex_pattern = 114,
  sym_regex_flags = 115,
  sym_number = 116,
  sym_private_property_identifier = 117,
  anon_sym_target = 118,
  anon_sym_meta = 119,
  sym_this = 120,
  sym_super = 121,
  sym_true = 122,
  sym_false = 123,
  sym_null = 124,
  sym_undefined = 125,
  anon_sym_AT = 126,
  anon_sym_static = 127,
  anon_sym_readonly = 128,
  anon_sym_get = 129,
  anon_sym_set = 130,
  anon_sym_QMARK = 131,
  anon_sym_declare = 132,
  anon_sym_public = 133,
  anon_sym_private = 134,
  anon_sym_protected = 135,
  anon_sym_override = 136,
  anon_sym_module = 137,
  anon_sym_any = 138,
  anon_sym_number = 139,
  anon_sym_boolean = 140,
  anon_sym_string = 141,
  anon_sym_symbol = 142,
  anon_sym_object = 143,
  anon_sym_abstract = 144,
  anon_sym_accessor = 145,
  anon_sym_satisfies = 146,
  anon_sym_require = 147,
  anon_sym_extends = 148,
  anon_sym_implements = 149,
  anon_sym_global = 150,
  anon_sym_interface = 151,
  anon_sym_enum = 152,
  anon_sym_DASH_QMARK_COLON = 153,
  anon_sym_PLUS_QMARK_COLON = 154,
  anon_sym_QMARK_COLON = 155,
  anon_sym_asserts = 156,
  anon_sym_infer = 157,
  anon_sym_is = 158,
  anon_sym_keyof = 159,
  anon_sym_unique = 160,
  anon_sym_unknown = 161,
  anon_sym_never = 162,
  anon_sym_LBRACE_PIPE = 163,
  anon_sym_PIPE_RBRACE = 164,
  sym__automatic_semicolon = 165,
  sym__template_chars = 166,
  sym__ternary_qmark = 167,
  sym_html_comment = 168,
  sym_jsx_text = 169,
  sym__function_signature_automatic_semicolon = 170,
  sym___error_recovery = 171,
  sym_program = 172,
  sym_export_statement = 173,
  sym_namespace_export = 174,
  sym_export_clause = 175,
  sym_export_specifier = 176,
  sym__module_export_name = 177,
  sym_declaration = 178,
  sym_import = 179,
  sym_import_statement = 180,
  sym_import_clause = 181,
  sym__from_clause = 182,
  sym_namespace_import = 183,
  sym_named_imports = 184,
  sym_import_specifier = 185,
  sym_import_attribute = 186,
  sym_statement = 187,
  sym_expression_statement = 188,
  sym_variable_declaration = 189,
  sym_lexical_declaration = 190,
  sym_variable_declarator = 191,
  sym_statement_block = 192,
  sym_else_clause = 193,
  sym_if_statement = 194,
  sym_switch_statement = 195,
  sym_for_statement = 196,
  sym_for_in_statement = 197,
  sym__for_header = 198,
  sym_while_statement = 199,
  sym_do_statement = 200,
  sym_try_statement = 201,
  sym_with_statement = 202,
  sym_break_statement = 203,
  sym_continue_statement = 204,
  sym_debugger_statement = 205,
  sym_return_statement = 206,
  sym_throw_statement = 207,
  sym_empty_statement = 208,
  sym_labeled_statement = 209,
  sym_switch_body = 210,
  sym_switch_case = 211,
  sym_switch_default = 212,
  sym_catch_clause = 213,
  sym_finally_clause = 214,
  sym_parenthesized_expression = 215,
  sym_expression = 216,
  sym_primary_expression = 217,
  sym_yield_expression = 218,
  sym_object = 219,
  sym_object_pattern = 220,
  sym_assignment_pattern = 221,
  sym_object_assignment_pattern = 222,
  sym_array = 223,
  sym_array_pattern = 224,
  sym_jsx_element = 225,
  sym_jsx_expression = 226,
  sym_jsx_opening_element = 227,
  sym_nested_identifier = 228,
  sym_jsx_namespace_name = 229,
  sym_jsx_closing_element = 230,
  sym_jsx_self_closing_element = 231,
  sym_jsx_attribute = 232,
  sym__jsx_string = 233,
  sym_class = 234,
  sym_class_declaration = 235,
  sym_class_heritage = 236,
  sym_function_expression = 237,
  sym_function_declaration = 238,
  sym_generator_function = 239,
  sym_generator_function_declaration = 240,
  sym_arrow_function = 241,
  sym__call_signature = 242,
  sym__formal_parameter = 243,
  sym_optional_chain = 244,
  sym_call_expression = 245,
  sym_new_expression = 246,
  sym_await_expression = 247,
  sym_member_expression = 248,
  sym_subscript_expression = 249,
  sym_assignment_expression = 250,
  sym__augmented_assignment_lhs = 251,
  sym_augmented_assignment_expression = 252,
  sym__initializer = 253,
  sym__destructuring_pattern = 254,
  sym_spread_element = 255,
  sym_ternary_expression = 256,
  sym_binary_expression = 257,
  sym_unary_expression = 258,
  sym_update_expression = 259,
  sym_sequence_expression = 260,
  sym_string = 261,
  sym_template_string = 262,
  sym_template_substitution = 263,
  sym_regex = 264,
  sym_meta_property = 265,
  sym_arguments = 266,
  sym_decorator = 267,
  sym_decorator_member_expression = 268,
  sym_decorator_call_expression = 269,
  sym_class_body = 270,
  sym_formal_parameters = 271,
  sym_class_static_block = 272,
  sym_pattern = 273,
  sym_rest_pattern = 274,
  sym_method_definition = 275,
  sym_pair = 276,
  sym_pair_pattern = 277,
  sym__property_name = 278,
  sym_computed_property_name = 279,
  sym_public_field_definition = 280,
  sym__import_identifier = 281,
  sym_non_null_expression = 282,
  sym_method_signature = 283,
  sym_abstract_method_signature = 284,
  sym_function_signature = 285,
  sym_decorator_parenthesized_expression = 286,
  sym_as_expression = 287,
  sym_satisfies_expression = 288,
  sym_instantiation_expression = 289,
  sym_import_require_clause = 290,
  sym_extends_clause = 291,
  sym__extends_clause_single = 292,
  sym_implements_clause = 293,
  sym_ambient_declaration = 294,
  sym_abstract_class_declaration = 295,
  sym_module = 296,
  sym_internal_module = 297,
  sym__module = 298,
  sym_import_alias = 299,
  sym_nested_type_identifier = 300,
  sym_interface_declaration = 301,
  sym_extends_type_clause = 302,
  sym_enum_declaration = 303,
  sym_enum_body = 304,
  sym_enum_assignment = 305,
  sym_type_alias_declaration = 306,
  sym_accessibility_modifier = 307,
  sym_override_modifier = 308,
  sym_required_parameter = 309,
  sym_optional_parameter = 310,
  sym__parameter_name = 311,
  sym_omitting_type_annotation = 312,
  sym_adding_type_annotation = 313,
  sym_opting_type_annotation = 314,
  sym_type_annotation = 315,
  sym__type_query_member_expression_in_type_annotation = 316,
  sym__type_query_call_expression_in_type_annotation = 317,
  sym_asserts = 318,
  sym_asserts_annotation = 319,
  sym_type = 320,
  sym_tuple_parameter = 321,
  sym_optional_tuple_parameter = 322,
  sym_optional_type = 323,
  sym_rest_type = 324,
  sym__tuple_type_member = 325,
  sym_constructor_type = 326,
  sym_primary_type = 327,
  sym_template_type = 328,
  sym_template_literal_type = 329,
  sym_infer_type = 330,
  sym_conditional_type = 331,
  sym_generic_type = 332,
  sym_type_predicate = 333,
  sym_type_predicate_annotation = 334,
  sym__type_query_member_expression = 335,
  sym__type_query_subscript_expression = 336,
  sym__type_query_call_expression = 337,
  sym__type_query_instantiation_expression = 338,
  sym_type_query = 339,
  sym_index_type_query = 340,
  sym_lookup_type = 341,
  sym_mapped_type_clause = 342,
  sym_literal_type = 343,
  sym__number = 344,
  sym_existential_type = 345,
  sym_flow_maybe_type = 346,
  sym_parenthesized_type = 347,
  sym_predefined_type = 348,
  sym_type_arguments = 349,
  sym_object_type = 350,
  sym_call_signature = 351,
  sym_property_signature = 352,
  sym_type_parameters = 353,
  sym_type_parameter = 354,
  sym_default_type = 355,
  sym_constraint = 356,
  sym_construct_signature = 357,
  sym_index_signature = 358,
  sym_array_type = 359,
  sym_tuple_type = 360,
  sym_readonly_type = 361,
  sym_union_type = 362,
  sym_intersection_type = 363,
  sym_function_type = 364,
  aux_sym_program_repeat1 = 365,
  aux_sym_export_statement_repeat1 = 366,
  aux_sym_export_clause_repeat1 = 367,
  aux_sym_named_imports_repeat1 = 368,
  aux_sym_variable_declaration_repeat1 = 369,
  aux_sym_switch_body_repeat1 = 370,
  aux_sym_object_repeat1 = 371,
  aux_sym_object_pattern_repeat1 = 372,
  aux_sym_array_repeat1 = 373,
  aux_sym_array_pattern_repeat1 = 374,
  aux_sym_jsx_element_repeat1 = 375,
  aux_sym__jsx_string_repeat1 = 376,
  aux_sym__jsx_string_repeat2 = 377,
  aux_sym_sequence_expression_repeat1 = 378,
  aux_sym_string_repeat1 = 379,
  aux_sym_string_repeat2 = 380,
  aux_sym_template_string_repeat1 = 381,
  aux_sym_class_body_repeat1 = 382,
  aux_sym_formal_parameters_repeat1 = 383,
  aux_sym__jsx_start_opening_element_repeat1 = 384,
  aux_sym_extends_clause_repeat1 = 385,
  aux_sym_implements_clause_repeat1 = 386,
  aux_sym_extends_type_clause_repeat1 = 387,
  aux_sym_enum_body_repeat1 = 388,
  aux_sym_template_literal_type_repeat1 = 389,
  aux_sym_object_type_repeat1 = 390,
  aux_sym_type_parameters_repeat1 = 391,
  aux_sym_tuple_type_repeat1 = 392,
  alias_sym_interface_body = 393,
  alias_sym_property_identifier = 394,
  alias_sym_shorthand_property_identifier = 395,
  alias_sym_shorthand_property_identifier_pattern = 396,
  alias_sym_statement_identifier = 397,
  alias_sym_this_type = 398,
  alias_sym_type_identifier = 399,
};

static const char * const ts_symbol_names[] = {
  [ts_builtin_sym_end] = "end",
  [sym_identifier] = "identifier",
  [sym_hash_bang_line] = "hash_bang_line",
  [anon_sym_export] = "export",
  [anon_sym_STAR] = "*",
  [anon_sym_default] = "default",
  [anon_sym_type] = "type",
  [anon_sym_EQ] = "=",
  [anon_sym_as] = "as",
  [anon_sym_namespace] = "namespace",
  [anon_sym_LBRACE] = "{",
  [anon_sym_COMMA] = ",",
  [anon_sym_RBRACE] = "}",
  [anon_sym_typeof] = "typeof",
  [anon_sym_import] = "import",
  [anon_sym_from] = "from",
  [anon_sym_with] = "with",
  [anon_sym_assert] = "assert",
  [anon_sym_var] = "var",
  [anon_sym_let] = "let",
  [anon_sym_const] = "const",
  [anon_sym_BANG] = "!",
  [anon_sym_else] = "else",
  [anon_sym_if] = "if",
  [anon_sym_switch] = "switch",
  [anon_sym_for] = "for",
  [anon_sym_LPAREN] = "(",
  [anon_sym_SEMI] = ";",
  [anon_sym_RPAREN] = ")",
  [anon_sym_await] = "await",
  [anon_sym_in] = "in",
  [anon_sym_of] = "of",
  [anon_sym_while] = "while",
  [anon_sym_do] = "do",
  [anon_sym_try] = "try",
  [anon_sym_break] = "break",
  [anon_sym_continue] = "continue",
  [anon_sym_debugger] = "debugger",
  [anon_sym_return] = "return",
  [anon_sym_throw] = "throw",
  [anon_sym_COLON] = ":",
  [anon_sym_case] = "case",
  [anon_sym_catch] = "catch",
  [anon_sym_finally] = "finally",
  [anon_sym_yield] = "yield",
  [anon_sym_LBRACK] = "[",
  [anon_sym_RBRACK] = "]",
  [sym_html_character_reference] = "html_character_reference",
  [anon_sym_GT] = ">",
  [sym_jsx_identifier] = "identifier",
  [anon_sym_DOT] = ".",
  [anon_sym_LT_SLASH] = "</",
  [anon_sym_SLASH_GT] = "/>",
  [anon_sym_DQUOTE] = "\"",
  [anon_sym_SQUOTE] = "'",
  [sym_unescaped_double_jsx_string_fragment] = "string_fragment",
  [sym_unescaped_single_jsx_string_fragment] = "string_fragment",
  [anon_sym_class] = "class",
  [anon_sym_async] = "async",
  [anon_sym_function] = "function",
  [anon_sym_EQ_GT] = "=>",
  [anon_sym_QMARK_DOT] = "\?.",
  [anon_sym_new] = "new",
  [anon_sym_using] = "using",
  [anon_sym_PLUS_EQ] = "+=",
  [anon_sym_DASH_EQ] = "-=",
  [anon_sym_STAR_EQ] = "*=",
  [anon_sym_SLASH_EQ] = "/=",
  [anon_sym_PERCENT_EQ] = "%=",
  [anon_sym_CARET_EQ] = "^=",
  [anon_sym_AMP_EQ] = "&=",
  [anon_sym_PIPE_EQ] = "|=",
  [anon_sym_GT_GT_EQ] = ">>=",
  [anon_sym_GT_GT_GT_EQ] = ">>>=",
  [anon_sym_LT_LT_EQ] = "<<=",
  [anon_sym_STAR_STAR_EQ] = "**=",
  [anon_sym_AMP_AMP_EQ] = "&&=",
  [anon_sym_PIPE_PIPE_EQ] = "||=",
  [anon_sym_QMARK_QMARK_EQ] = "\?\?=",
  [anon_sym_DOT_DOT_DOT] = "...",
  [anon_sym_AMP_AMP] = "&&",
  [anon_sym_PIPE_PIPE] = "||",
  [anon_sym_GT_GT] = ">>",
  [anon_sym_GT_GT_GT] = ">>>",
  [anon_sym_LT_LT] = "<<",
  [anon_sym_AMP] = "&",
  [anon_sym_CARET] = "^",
  [anon_sym_PIPE] = "|",
  [anon_sym_PLUS] = "+",
  [anon_sym_DASH] = "-",
  [anon_sym_SLASH] = "/",
  [anon_sym_PERCENT] = "%",
  [anon_sym_STAR_STAR] = "**",
  [anon_sym_LT] = "<",
  [anon_sym_LT_EQ] = "<=",
  [anon_sym_EQ_EQ] = "==",
  [anon_sym_EQ_EQ_EQ] = "===",
  [anon_sym_BANG_EQ] = "!=",
  [anon_sym_BANG_EQ_EQ] = "!==",
  [anon_sym_GT_EQ] = ">=",
  [anon_sym_QMARK_QMARK] = "\?\?",
  [anon_sym_instanceof] = "instanceof",
  [anon_sym_TILDE] = "~",
  [anon_sym_void] = "void",
  [anon_sym_delete] = "delete",
  [anon_sym_PLUS_PLUS] = "++",
  [anon_sym_DASH_DASH] = "--",
  [sym_unescaped_double_string_fragment] = "string_fragment",
  [sym_unescaped_single_string_fragment] = "string_fragment",
  [sym_escape_sequence] = "escape_sequence",
  [sym_comment] = "comment",
  [anon_sym_BQUOTE] = "`",
  [anon_sym_DOLLAR_LBRACE] = "${",
  [anon_sym_SLASH2] = "/",
  [sym_regex_pattern] = "regex_pattern",
  [sym_regex_flags] = "regex_flags",
  [sym_number] = "number",
  [sym_private_property_identifier] = "private_property_identifier",
  [anon_sym_target] = "target",
  [anon_sym_meta] = "meta",
  [sym_this] = "this",
  [sym_super] = "super",
  [sym_true] = "true",
  [sym_false] = "false",
  [sym_null] = "null",
  [sym_undefined] = "undefined",
  [anon_sym_AT] = "@",
  [anon_sym_static] = "static",
  [anon_sym_readonly] = "readonly",
  [anon_sym_get] = "get",
  [anon_sym_set] = "set",
  [anon_sym_QMARK] = "\?",
  [anon_sym_declare] = "declare",
  [anon_sym_public] = "public",
  [anon_sym_private] = "private",
  [anon_sym_protected] = "protected",
  [anon_sym_override] = "override",
  [anon_sym_module] = "module",
  [anon_sym_any] = "any",
  [anon_sym_number] = "number",
  [anon_sym_boolean] = "boolean",
  [anon_sym_string] = "string",
  [anon_sym_symbol] = "symbol",
  [anon_sym_object] = "object",
  [anon_sym_abstract] = "abstract",
  [anon_sym_accessor] = "accessor",
  [anon_sym_satisfies] = "satisfies",
  [anon_sym_require] = "require",
  [anon_sym_extends] = "extends",
  [anon_sym_implements] = "implements",
  [anon_sym_global] = "global",
  [anon_sym_interface] = "interface",
  [anon_sym_enum] = "enum",
  [anon_sym_DASH_QMARK_COLON] = "-\?:",
  [anon_sym_PLUS_QMARK_COLON] = "+\?:",
  [anon_sym_QMARK_COLON] = "\?:",
  [anon_sym_asserts] = "asserts",
  [anon_sym_infer] = "infer",
  [anon_sym_is] = "is",
  [anon_sym_keyof] = "keyof",
  [anon_sym_unique] = "unique symbol",
  [anon_sym_unknown] = "unknown",
  [anon_sym_never] = "never",
  [anon_sym_LBRACE_PIPE] = "{|",
  [anon_sym_PIPE_RBRACE] = "|}",
  [sym__automatic_semicolon] = "_automatic_semicolon",
  [sym__template_chars] = "string_fragment",
  [sym__ternary_qmark] = "\?",
  [sym_html_comment] = "html_comment",
  [sym_jsx_text] = "jsx_text",
  [sym__function_signature_automatic_semicolon] = "_function_signature_automatic_semicolon",
  [sym___error_recovery] = "__error_recovery",
  [sym_program] = "program",
  [sym_export_statement] = "export_statement",
  [sym_namespace_export] = "namespace_export",
  [sym_export_clause] = "export_clause",
  [sym_export_specifier] = "export_specifier",
  [sym__module_export_name] = "_module_export_name",
  [sym_declaration] = "declaration",
  [sym_import] = "import",
  [sym_import_statement] = "import_statement",
  [sym_import_clause] = "import_clause",
  [sym__from_clause] = "_from_clause",
  [sym_namespace_import] = "namespace_import",
  [sym_named_imports] = "named_imports",
  [sym_import_specifier] = "import_specifier",
  [sym_import_attribute] = "import_attribute",
  [sym_statement] = "statement",
  [sym_expression_statement] = "expression_statement",
  [sym_variable_declaration] = "variable_declaration",
  [sym_lexical_declaration] = "lexical_declaration",
  [sym_variable_declarator] = "variable_declarator",
  [sym_statement_block] = "statement_block",
  [sym_else_clause] = "else_clause",
  [sym_if_statement] = "if_statement",
  [sym_switch_statement] = "switch_statement",
  [sym_for_statement] = "for_statement",
  [sym_for_in_statement] = "for_in_statement",
  [sym__for_header] = "_for_header",
  [sym_while_statement] = "while_statement",
  [sym_do_statement] = "do_statement",
  [sym_try_statement] = "try_statement",
  [sym_with_statement] = "with_statement",
  [sym_break_statement] = "break_statement",
  [sym_continue_statement] = "continue_statement",
  [sym_debugger_statement] = "debugger_statement",
  [sym_return_statement] = "return_statement",
  [sym_throw_statement] = "throw_statement",
  [sym_empty_statement] = "empty_statement",
  [sym_labeled_statement] = "labeled_statement",
  [sym_switch_body] = "switch_body",
  [sym_switch_case] = "switch_case",
  [sym_switch_default] = "switch_default",
  [sym_catch_clause] = "catch_clause",
  [sym_finally_clause] = "finally_clause",
  [sym_parenthesized_expression] = "parenthesized_expression",
  [sym_expression] = "expression",
  [sym_primary_expression] = "primary_expression",
  [sym_yield_expression] = "yield_expression",
  [sym_object] = "object",
  [sym_object_pattern] = "object_pattern",
  [sym_assignment_pattern] = "assignment_pattern",
  [sym_object_assignment_pattern] = "object_assignment_pattern",
  [sym_array] = "array",
  [sym_array_pattern] = "array_pattern",
  [sym_jsx_element] = "jsx_element",
  [sym_jsx_expression] = "jsx_expression",
  [sym_jsx_opening_element] = "jsx_opening_element",
  [sym_nested_identifier] = "nested_identifier",
  [sym_jsx_namespace_name] = "jsx_namespace_name",
  [sym_jsx_closing_element] = "jsx_closing_element",
  [sym_jsx_self_closing_element] = "jsx_self_closing_element",
  [sym_jsx_attribute] = "jsx_attribute",
  [sym__jsx_string] = "string",
  [sym_class] = "class",
  [sym_class_declaration] = "class_declaration",
  [sym_class_heritage] = "class_heritage",
  [sym_function_expression] = "function_expression",
  [sym_function_declaration] = "function_declaration",
  [sym_generator_function] = "generator_function",
  [sym_generator_function_declaration] = "generator_function_declaration",
  [sym_arrow_function] = "arrow_function",
  [sym__call_signature] = "_call_signature",
  [sym__formal_parameter] = "_formal_parameter",
  [sym_optional_chain] = "optional_chain",
  [sym_call_expression] = "call_expression",
  [sym_new_expression] = "new_expression",
  [sym_await_expression] = "await_expression",
  [sym_member_expression] = "member_expression",
  [sym_subscript_expression] = "subscript_expression",
  [sym_assignment_expression] = "assignment_expression",
  [sym__augmented_assignment_lhs] = "_augmented_assignment_lhs",
  [sym_augmented_assignment_expression] = "augmented_assignment_expression",
  [sym__initializer] = "_initializer",
  [sym__destructuring_pattern] = "_destructuring_pattern",
  [sym_spread_element] = "spread_element",
  [sym_ternary_expression] = "ternary_expression",
  [sym_binary_expression] = "binary_expression",
  [sym_unary_expression] = "unary_expression",
  [sym_update_expression] = "update_expression",
  [sym_sequence_expression] = "sequence_expression",
  [sym_string] = "string",
  [sym_template_string] = "template_string",
  [sym_template_substitution] = "template_substitution",
  [sym_regex] = "regex",
  [sym_meta_property] = "meta_property",
  [sym_arguments] = "arguments",
  [sym_decorator] = "decorator",
  [sym_decorator_member_expression] = "member_expression",
  [sym_decorator_call_expression] = "call_expression",
  [sym_class_body] = "class_body",
  [sym_formal_parameters] = "formal_parameters",
  [sym_class_static_block] = "class_static_block",
  [sym_pattern] = "pattern",
  [sym_rest_pattern] = "rest_pattern",
  [sym_method_definition] = "method_definition",
  [sym_pair] = "pair",
  [sym_pair_pattern] = "pair_pattern",
  [sym__property_name] = "_property_name",
  [sym_computed_property_name] = "computed_property_name",
  [sym_public_field_definition] = "public_field_definition",
  [sym__import_identifier] = "_import_identifier",
  [sym_non_null_expression] = "non_null_expression",
  [sym_method_signature] = "method_signature",
  [sym_abstract_method_signature] = "abstract_method_signature",
  [sym_function_signature] = "function_signature",
  [sym_decorator_parenthesized_expression] = "parenthesized_expression",
  [sym_as_expression] = "as_expression",
  [sym_satisfies_expression] = "satisfies_expression",
  [sym_instantiation_expression] = "instantiation_expression",
  [sym_import_require_clause] = "import_require_clause",
  [sym_extends_clause] = "extends_clause",
  [sym__extends_clause_single] = "_extends_clause_single",
  [sym_implements_clause] = "implements_clause",
  [sym_ambient_declaration] = "ambient_declaration",
  [sym_abstract_class_declaration] = "abstract_class_declaration",
  [sym_module] = "module",
  [sym_internal_module] = "internal_module",
  [sym__module] = "_module",
  [sym_import_alias] = "import_alias",
  [sym_nested_type_identifier] = "nested_type_identifier",
  [sym_interface_declaration] = "interface_declaration",
  [sym_extends_type_clause] = "extends_type_clause",
  [sym_enum_declaration] = "enum_declaration",
  [sym_enum_body] = "enum_body",
  [sym_enum_assignment] = "enum_assignment",
  [sym_type_alias_declaration] = "type_alias_declaration",
  [sym_accessibility_modifier] = "accessibility_modifier",
  [sym_override_modifier] = "override_modifier",
  [sym_required_parameter] = "required_parameter",
  [sym_optional_parameter] = "optional_parameter",
  [sym__parameter_name] = "_parameter_name",
  [sym_omitting_type_annotation] = "omitting_type_annotation",
  [sym_adding_type_annotation] = "adding_type_annotation",
  [sym_opting_type_annotation] = "opting_type_annotation",
  [sym_type_annotation] = "type_annotation",
  [sym__type_query_member_expression_in_type_annotation] = "member_expression",
  [sym__type_query_call_expression_in_type_annotation] = "call_expression",
  [sym_asserts] = "asserts",
  [sym_asserts_annotation] = "asserts_annotation",
  [sym_type] = "type",
  [sym_tuple_parameter] = "required_parameter",
  [sym_optional_tuple_parameter] = "optional_parameter",
  [sym_optional_type] = "optional_type",
  [sym_rest_type] = "rest_type",
  [sym__tuple_type_member] = "_tuple_type_member",
  [sym_constructor_type] = "constructor_type",
  [sym_primary_type] = "primary_type",
  [sym_template_type] = "template_type",
  [sym_template_literal_type] = "template_literal_type",
  [sym_infer_type] = "infer_type",
  [sym_conditional_type] = "conditional_type",
  [sym_generic_type] = "generic_type",
  [sym_type_predicate] = "type_predicate",
  [sym_type_predicate_annotation] = "type_predicate_annotation",
  [sym__type_query_member_expression] = "member_expression",
  [sym__type_query_subscript_expression] = "subscript_expression",
  [sym__type_query_call_expression] = "call_expression",
  [sym__type_query_instantiation_expression] = "instantiation_expression",
  [sym_type_query] = "type_query",
  [sym_index_type_query] = "index_type_query",
  [sym_lookup_type] = "lookup_type",
  [sym_mapped_type_clause] = "mapped_type_clause",
  [sym_literal_type] = "literal_type",
  [sym__number] = "unary_expression",
  [sym_existential_type] = "existential_type",
  [sym_flow_maybe_type] = "flow_maybe_type",
  [sym_parenthesized_type] = "parenthesized_type",
  [sym_predefined_type] = "predefined_type",
  [sym_type_arguments] = "type_arguments",
  [sym_object_type] = "object_type",
  [sym_call_signature] = "call_signature",
  [sym_property_signature] = "property_signature",
  [sym_type_parameters] = "type_parameters",
  [sym_type_parameter] = "type_parameter",
  [sym_default_type] = "default_type",
  [sym_constraint] = "constraint",
  [sym_construct_signature] = "construct_signature",
  [sym_index_signature] = "index_signature",
  [sym_array_type] = "array_type",
  [sym_tuple_type] = "tuple_type",
  [sym_readonly_type] = "readonly_type",
  [sym_union_type] = "union_type",
  [sym_intersection_type] = "intersection_type",
  [sym_function_type] = "function_type",
  [aux_sym_program_repeat1] = "program_repeat1",
  [aux_sym_export_statement_repeat1] = "export_statement_repeat1",
  [aux_sym_export_clause_repeat1] = "export_clause_repeat1",
  [aux_sym_named_imports_repeat1] = "named_imports_repeat1",
  [aux_sym_variable_declaration_repeat1] = "variable_declaration_repeat1",
  [aux_sym_switch_body_repeat1] = "switch_body_repeat1",
  [aux_sym_object_repeat1] = "object_repeat1",
  [aux_sym_object_pattern_repeat1] = "object_pattern_repeat1",
  [aux_sym_array_repeat1] = "array_repeat1",
  [aux_sym_array_pattern_repeat1] = "array_pattern_repeat1",
  [aux_sym_jsx_element_repeat1] = "jsx_element_repeat1",
  [aux_sym__jsx_string_repeat1] = "_jsx_string_repeat1",
  [aux_sym__jsx_string_repeat2] = "_jsx_string_repeat2",
  [aux_sym_sequence_expression_repeat1] = "sequence_expression_repeat1",
  [aux_sym_string_repeat1] = "string_repeat1",
  [aux_sym_string_repeat2] = "string_repeat2",
  [aux_sym_template_string_repeat1] = "template_string_repeat1",
  [aux_sym_class_body_repeat1] = "class_body_repeat1",
  [aux_sym_formal_parameters_repeat1] = "formal_parameters_repeat1",
  [aux_sym__jsx_start_opening_element_repeat1] = "_jsx_start_opening_element_repeat1",
  [aux_sym_extends_clause_repeat1] = "extends_clause_repeat1",
  [aux_sym_implements_clause_repeat1] = "implements_clause_repeat1",
  [aux_sym_extends_type_clause_repeat1] = "extends_type_clause_repeat1",
  [aux_sym_enum_body_repeat1] = "enum_body_repeat1",
  [aux_sym_template_literal_type_repeat1] = "template_literal_type_repeat1",
  [aux_sym_object_type_repeat1] = "object_type_repeat1",
  [aux_sym_type_parameters_repeat1] = "type_parameters_repeat1",
  [aux_sym_tuple_type_repeat1] = "tuple_type_repeat1",
  [alias_sym_interface_body] = "interface_body",
  [alias_sym_property_identifier] = "property_identifier",
  [alias_sym_shorthand_property_identifier] = "shorthand_property_identifier",
  [alias_sym_shorthand_property_identifier_pattern] = "shorthand_property_identifier_pattern",
  [alias_sym_statement_identifier] = "statement_identifier",
  [alias_sym_this_type] = "this_type",
  [alias_sym_type_identifier] = "type_identifier",
};

static const TSSymbol ts_symbol_map[] = {
  [ts_builtin_sym_end] = ts_builtin_sym_end,
  [sym_identifier] = sym_identifier,
  [sym_hash_bang_line] = sym_hash_bang_line,
  [anon_sym_export] = anon_sym_export,
  [anon_sym_STAR] = anon_sym_STAR,
  [anon_sym_default] = anon_sym_default,
  [anon_sym_type] = anon_sym_type,
  [anon_sym_EQ] = anon_sym_EQ,
  [anon_sym_as] = anon_sym_as,
  [anon_sym_namespace] = anon_sym_namespace,
  [anon_sym_LBRACE] = anon_sym_LBRACE,
  [anon_sym_COMMA] = anon_sym_COMMA,
  [anon_sym_RBRACE] = anon_sym_RBRACE,
  [anon_sym_typeof] = anon_sym_typeof,
  [anon_sym_import] = anon_sym_import,
  [anon_sym_from] = anon_sym_from,
  [anon_sym_with] = anon_sym_with,
  [anon_sym_assert] = anon_sym_assert,
  [anon_sym_var] = anon_sym_var,
  [anon_sym_let] = anon_sym_let,
  [anon_sym_const] = anon_sym_const,
  [anon_sym_BANG] = anon_sym_BANG,
  [anon_sym_else] = anon_sym_else,
  [anon_sym_if] = anon_sym_if,
  [anon_sym_switch] = anon_sym_switch,
  [anon_sym_for] = anon_sym_for,
  [anon_sym_LPAREN] = anon_sym_LPAREN,
  [anon_sym_SEMI] = anon_sym_SEMI,
  [anon_sym_RPAREN] = anon_sym_RPAREN,
  [anon_sym_await] = anon_sym_await,
  [anon_sym_in] = anon_sym_in,
  [anon_sym_of] = anon_sym_of,
  [anon_sym_while] = anon_sym_while,
  [anon_sym_do] = anon_sym_do,
  [anon_sym_try] = anon_sym_try,
  [anon_sym_break] = anon_sym_break,
  [anon_sym_continue] = anon_sym_continue,
  [anon_sym_debugger] = anon_sym_debugger,
  [anon_sym_return] = anon_sym_return,
  [anon_sym_throw] = anon_sym_throw,
  [anon_sym_COLON] = anon_sym_COLON,
  [anon_sym_case] = anon_sym_case,
  [anon_sym_catch] = anon_sym_catch,
  [anon_sym_finally] = anon_sym_finally,
  [anon_sym_yield] = anon_sym_yield,
  [anon_sym_LBRACK] = anon_sym_LBRACK,
  [anon_sym_RBRACK] = anon_sym_RBRACK,
  [sym_html_character_reference] = sym_html_character_reference,
  [anon_sym_GT] = anon_sym_GT,
  [sym_jsx_identifier] = sym_identifier,
  [anon_sym_DOT] = anon_sym_DOT,
  [anon_sym_LT_SLASH] = anon_sym_LT_SLASH,
  [anon_sym_SLASH_GT] = anon_sym_SLASH_GT,
  [anon_sym_DQUOTE] = anon_sym_DQUOTE,
  [anon_sym_SQUOTE] = anon_sym_SQUOTE,
  [sym_unescaped_double_jsx_string_fragment] = sym__template_chars,
  [sym_unescaped_single_jsx_string_fragment] = sym__template_chars,
  [anon_sym_class] = anon_sym_class,
  [anon_sym_async] = anon_sym_async,
  [anon_sym_function] = anon_sym_function,
  [anon_sym_EQ_GT] = anon_sym_EQ_GT,
  [anon_sym_QMARK_DOT] = anon_sym_QMARK_DOT,
  [anon_sym_new] = anon_sym_new,
  [anon_sym_using] = anon_sym_using,
  [anon_sym_PLUS_EQ] = anon_sym_PLUS_EQ,
  [anon_sym_DASH_EQ] = anon_sym_DASH_EQ,
  [anon_sym_STAR_EQ] = anon_sym_STAR_EQ,
  [anon_sym_SLASH_EQ] = anon_sym_SLASH_EQ,
  [anon_sym_PERCENT_EQ] = anon_sym_PERCENT_EQ,
  [anon_sym_CARET_EQ] = anon_sym_CARET_EQ,
  [anon_sym_AMP_EQ] = anon_sym_AMP_EQ,
  [anon_sym_PIPE_EQ] = anon_sym_PIPE_EQ,
  [anon_sym_GT_GT_EQ] = anon_sym_GT_GT_EQ,
  [anon_sym_GT_GT_GT_EQ] = anon_sym_GT_GT_GT_EQ,
  [anon_sym_LT_LT_EQ] = anon_sym_LT_LT_EQ,
  [anon_sym_STAR_STAR_EQ] = anon_sym_STAR_STAR_EQ,
  [anon_sym_AMP_AMP_EQ] = anon_sym_AMP_AMP_EQ,
  [anon_sym_PIPE_PIPE_EQ] = anon_sym_PIPE_PIPE_EQ,
  [anon_sym_QMARK_QMARK_EQ] = anon_sym_QMARK_QMARK_EQ,
  [anon_sym_DOT_DOT_DOT] = anon_sym_DOT_DOT_DOT,
  [anon_sym_AMP_AMP] = anon_sym_AMP_AMP,
  [anon_sym_PIPE_PIPE] = anon_sym_PIPE_PIPE,
  [anon_sym_GT_GT] = anon_sym_GT_GT,
  [anon_sym_GT_GT_GT] = anon_sym_GT_GT_GT,
  [anon_sym_LT_LT] = anon_sym_LT_LT,
  [anon_sym_AMP] = anon_sym_AMP,
  [anon_sym_CARET] = anon_sym_CARET,
  [anon_sym_PIPE] = anon_sym_PIPE,
  [anon_sym_PLUS] = anon_sym_PLUS,
  [anon_sym_DASH] = anon_sym_DASH,
  [anon_sym_SLASH] = anon_sym_SLASH,
  [anon_sym_PERCENT] = anon_sym_PERCENT,
  [anon_sym_STAR_STAR] = anon_sym_STAR_STAR,
  [anon_sym_LT] = anon_sym_LT,
  [anon_sym_LT_EQ] = anon_sym_LT_EQ,
  [anon_sym_EQ_EQ] = anon_sym_EQ_EQ,
  [anon_sym_EQ_EQ_EQ] = anon_sym_EQ_EQ_EQ,
  [anon_sym_BANG_EQ] = anon_sym_BANG_EQ,
  [anon_sym_BANG_EQ_EQ] = anon_sym_BANG_EQ_EQ,
  [anon_sym_GT_EQ] = anon_sym_GT_EQ,
  [anon_sym_QMARK_QMARK] = anon_sym_QMARK_QMARK,
  [anon_sym_instanceof] = anon_sym_instanceof,
  [anon_sym_TILDE] = anon_sym_TILDE,
  [anon_sym_void] = anon_sym_void,
  [anon_sym_delete] = anon_sym_delete,
  [anon_sym_PLUS_PLUS] = anon_sym_PLUS_PLUS,
  [anon_sym_DASH_DASH] = anon_sym_DASH_DASH,
  [sym_unescaped_double_string_fragment] = sym__template_chars,
  [sym_unescaped_single_string_fragment] = sym__template_chars,
  [sym_escape_sequence] = sym_escape_sequence,
  [sym_comment] = sym_comment,
  [anon_sym_BQUOTE] = anon_sym_BQUOTE,
  [anon_sym_DOLLAR_LBRACE] = anon_sym_DOLLAR_LBRACE,
  [anon_sym_SLASH2] = anon_sym_SLASH,
  [sym_regex_pattern] = sym_regex_pattern,
  [sym_regex_flags] = sym_regex_flags,
  [sym_number] = sym_number,
  [sym_private_property_identifier] = sym_private_property_identifier,
  [anon_sym_target] = anon_sym_target,
  [anon_sym_meta] = anon_sym_meta,
  [sym_this] = sym_this,
  [sym_super] = sym_super,
  [sym_true] = sym_true,
  [sym_false] = sym_false,
  [sym_null] = sym_null,
  [sym_undefined] = sym_undefined,
  [anon_sym_AT] = anon_sym_AT,
  [anon_sym_static] = anon_sym_static,
  [anon_sym_readonly] = anon_sym_readonly,
  [anon_sym_get] = anon_sym_get,
  [anon_sym_set] = anon_sym_set,
  [anon_sym_QMARK] = anon_sym_QMARK,
  [anon_sym_declare] = anon_sym_declare,
  [anon_sym_public] = anon_sym_public,
  [anon_sym_private] = anon_sym_private,
  [anon_sym_protected] = anon_sym_protected,
  [anon_sym_override] = anon_sym_override,
  [anon_sym_module] = anon_sym_module,
  [anon_sym_any] = anon_sym_any,
  [anon_sym_number] = anon_sym_number,
  [anon_sym_boolean] = anon_sym_boolean,
  [anon_sym_string] = anon_sym_string,
  [anon_sym_symbol] = anon_sym_symbol,
  [anon_sym_object] = anon_sym_object,
  [anon_sym_abstract] = anon_sym_abstract,
  [anon_sym_accessor] = anon_sym_accessor,
  [anon_sym_satisfies] = anon_sym_satisfies,
  [anon_sym_require] = anon_sym_require,
  [anon_sym_extends] = anon_sym_extends,
  [anon_sym_implements] = anon_sym_implements,
  [anon_sym_global] = anon_sym_global,
  [anon_sym_interface] = anon_sym_interface,
  [anon_sym_enum] = anon_sym_enum,
  [anon_sym_DASH_QMARK_COLON] = anon_sym_DASH_QMARK_COLON,
  [anon_sym_PLUS_QMARK_COLON] = anon_sym_PLUS_QMARK_COLON,
  [anon_sym_QMARK_COLON] = anon_sym_QMARK_COLON,
  [anon_sym_asserts] = anon_sym_asserts,
  [anon_sym_infer] = anon_sym_infer,
  [anon_sym_is] = anon_sym_is,
  [anon_sym_keyof] = anon_sym_keyof,
  [anon_sym_unique] = anon_sym_unique,
  [anon_sym_unknown] = anon_sym_unknown,
  [anon_sym_never] = anon_sym_never,
  [anon_sym_LBRACE_PIPE] = anon_sym_LBRACE_PIPE,
  [anon_sym_PIPE_RBRACE] = anon_sym_PIPE_RBRACE,
  [sym__automatic_semicolon] = sym__automatic_semicolon,
  [sym__template_chars] = sym__template_chars,
  [sym__ternary_qmark] = anon_sym_QMARK,
  [sym_html_comment] = sym_html_comment,
  [sym_jsx_text] = sym_jsx_text,
  [sym__function_signature_automatic_semicolon] = sym__function_signature_automatic_semicolon,
  [sym___error_recovery] = sym___error_recovery,
  [sym_program] = sym_program,
  [sym_export_statement] = sym_export_statement,
  [sym_namespace_export] = sym_namespace_export,
  [sym_export_clause] = sym_export_clause,
  [sym_export_specifier] = sym_export_specifier,
  [sym__module_export_name] = sym__module_export_name,
  [sym_declaration] = sym_declaration,
  [sym_import] = sym_import,
  [sym_import_statement] = sym_import_statement,
  [sym_import_clause] = sym_import_clause,
  [sym__from_clause] = sym__from_clause,
  [sym_namespace_import] = sym_namespace_import,
  [sym_named_imports] = sym_named_imports,
  [sym_import_specifier] = sym_import_specifier,
  [sym_import_attribute] = sym_import_attribute,
  [sym_statement] = sym_statement,
  [sym_expression_statement] = sym_expression_statement,
  [sym_variable_declaration] = sym_variable_declaration,
  [sym_lexical_declaration] = sym_lexical_declaration,
  [sym_variable_declarator] = sym_variable_declarator,
  [sym_statement_block] = sym_statement_block,
  [sym_else_clause] = sym_else_clause,
  [sym_if_statement] = sym_if_statement,
  [sym_switch_statement] = sym_switch_statement,
  [sym_for_statement] = sym_for_statement,
  [sym_for_in_statement] = sym_for_in_statement,
  [sym__for_header] = sym__for_header,
  [sym_while_statement] = sym_while_statement,
  [sym_do_statement] = sym_do_statement,
  [sym_try_statement] = sym_try_statement,
  [sym_with_statement] = sym_with_statement,
  [sym_break_statement] = sym_break_statement,
  [sym_continue_statement] = sym_continue_statement,
  [sym_debugger_statement] = sym_debugger_statement,
  [sym_return_statement] = sym_return_statement,
  [sym_throw_statement] = sym_throw_statement,
  [sym_empty_statement] = sym_empty_statement,
  [sym_labeled_statement] = sym_labeled_statement,
  [sym_switch_body] = sym_switch_body,
  [sym_switch_case] = sym_switch_case,
  [sym_switch_default] = sym_switch_default,
  [sym_catch_clause] = sym_catch_clause,
  [sym_finally_clause] = sym_finally_clause,
  [sym_parenthesized_expression] = sym_parenthesized_expression,
  [sym_expression] = sym_expression,
  [sym_primary_expression] = sym_primary_expression,
  [sym_yield_expression] = sym_yield_expression,
  [sym_object] = sym_object,
  [sym_object_pattern] = sym_object_pattern,
  [sym_assignment_pattern] = sym_assignment_pattern,
  [sym_object_assignment_pattern] = sym_object_assignment_pattern,
  [sym_array] = sym_array,
  [sym_array_pattern] = sym_array_pattern,
  [sym_jsx_element] = sym_jsx_element,
  [sym_jsx_expression] = sym_jsx_expression,
  [sym_jsx_opening_element] = sym_jsx_opening_element,
  [sym_nested_identifier] = sym_nested_identifier,
  [sym_jsx_namespace_name] = sym_jsx_namespace_name,
  [sym_jsx_closing_element] = sym_jsx_closing_element,
  [sym_jsx_self_closing_element] = sym_jsx_self_closing_element,
  [sym_jsx_attribute] = sym_jsx_attribute,
  [sym__jsx_string] = sym_string,
  [sym_class] = sym_class,
  [sym_class_declaration] = sym_class_declaration,
  [sym_class_heritage] = sym_class_heritage,
  [sym_function_expression] = sym_function_expression,
  [sym_function_declaration] = sym_function_declaration,
  [sym_generator_function] = sym_generator_function,
  [sym_generator_function_declaration] = sym_generator_function_declaration,
  [sym_arrow_function] = sym_arrow_function,
  [sym__call_signature] = sym__call_signature,
  [sym__formal_parameter] = sym__formal_parameter,
  [sym_optional_chain] = sym_optional_chain,
  [sym_call_expression] = sym_call_expression,
  [sym_new_expression] = sym_new_expression,
  [sym_await_expression] = sym_await_expression,
  [sym_member_expression] = sym_member_expression,
  [sym_subscript_expression] = sym_subscript_expression,
  [sym_assignment_expression] = sym_assignment_expression,
  [sym__augmented_assignment_lhs] = sym__augmented_assignment_lhs,
  [sym_augmented_assignment_expression] = sym_augmented_assignment_expression,
  [sym__initializer] = sym__initializer,
  [sym__destructuring_pattern] = sym__destructuring_pattern,
  [sym_spread_element] = sym_spread_element,
  [sym_ternary_expression] = sym_ternary_expression,
  [sym_binary_expression] = sym_binary_expression,
  [sym_unary_expression] = sym_unary_expression,
  [sym_update_expression] = sym_update_expression,
  [sym_sequence_expression] = sym_sequence_expression,
  [sym_string] = sym_string,
  [sym_template_string] = sym_template_string,
  [sym_template_substitution] = sym_template_substitution,
  [sym_regex] = sym_regex,
  [sym_meta_property] = sym_meta_property,
  [sym_arguments] = sym_arguments,
  [sym_decorator] = sym_decorator,
  [sym_decorator_member_expression] = sym_member_expression,
  [sym_decorator_call_expression] = sym_call_expression,
  [sym_class_body] = sym_class_body,
  [sym_formal_parameters] = sym_formal_parameters,
  [sym_class_static_block] = sym_class_static_block,
  [sym_pattern] = sym_pattern,
  [sym_rest_pattern] = sym_rest_pattern,
  [sym_method_definition] = sym_method_definition,
  [sym_pair] = sym_pair,
  [sym_pair_pattern] = sym_pair_pattern,
  [sym__property_name] = sym__property_name,
  [sym_computed_property_name] = sym_computed_property_name,
  [sym_public_field_definition] = sym_public_field_definition,
  [sym__import_identifier] = sym__import_identifier,
  [sym_non_null_expression] = sym_non_null_expression,
  [sym_method_signature] = sym_method_signature,
  [sym_abstract_method_signature] = sym_abstract_method_signature,
  [sym_function_signature] = sym_function_signature,
  [sym_decorator_parenthesized_expression] = sym_parenthesized_expression,
  [sym_as_expression] = sym_as_expression,
  [sym_satisfies_expression] = sym_satisfies_expression,
  [sym_instantiation_expression] = sym_instantiation_expression,
  [sym_import_require_clause] = sym_import_require_clause,
  [sym_extends_clause] = sym_extends_clause,
  [sym__extends_clause_single] = sym__extends_clause_single,
  [sym_implements_clause] = sym_implements_clause,
  [sym_ambient_declaration] = sym_ambient_declaration,
  [sym_abstract_class_declaration] = sym_abstract_class_declaration,
  [sym_module] = sym_module,
  [sym_internal_module] = sym_internal_module,
  [sym__module] = sym__module,
  [sym_import_alias] = sym_import_alias,
  [sym_nested_type_identifier] = sym_nested_type_identifier,
  [sym_interface_declaration] = sym_interface_declaration,
  [sym_extends_type_clause] = sym_extends_type_clause,
  [sym_enum_declaration] = sym_enum_declaration,
  [sym_enum_body] = sym_enum_body,
  [sym_enum_assignment] = sym_enum_assignment,
  [sym_type_alias_declaration] = sym_type_alias_declaration,
  [sym_accessibility_modifier] = sym_accessibility_modifier,
  [sym_override_modifier] = sym_override_modifier,
  [sym_required_parameter] = sym_required_parameter,
  [sym_optional_parameter] = sym_optional_parameter,
  [sym__parameter_name] = sym__parameter_name,
  [sym_omitting_type_annotation] = sym_omitting_type_annotation,
  [sym_adding_type_annotation] = sym_adding_type_annotation,
  [sym_opting_type_annotation] = sym_opting_type_annotation,
  [sym_type_annotation] = sym_type_annotation,
  [sym__type_query_member_expression_in_type_annotation] = sym_member_expression,
  [sym__type_query_call_expression_in_type_annotation] = sym_call_expression,
  [sym_asserts] = sym_asserts,
  [sym_asserts_annotation] = sym_asserts_annotation,
  [sym_type] = sym_type,
  [sym_tuple_parameter] = sym_required_parameter,
  [sym_optional_tuple_parameter] = sym_optional_parameter,
  [sym_optional_type] = sym_optional_type,
  [sym_rest_type] = sym_rest_type,
  [sym__tuple_type_member] = sym__tuple_type_member,
  [sym_constructor_type] = sym_constructor_type,
  [sym_primary_type] = sym_primary_type,
  [sym_template_type] = sym_template_type,
  [sym_template_literal_type] = sym_template_literal_type,
  [sym_infer_type] = sym_infer_type,
  [sym_conditional_type] = sym_conditional_type,
  [sym_generic_type] = sym_generic_type,
  [sym_type_predicate] = sym_type_predicate,
  [sym_type_predicate_annotation] = sym_type_predicate_annotation,
  [sym__type_query_member_expression] = sym_member_expression,
  [sym__type_query_subscript_expression] = sym_subscript_expression,
  [sym__type_query_call_expression] = sym_call_expression,
  [sym__type_query_instantiation_expression] = sym_instantiation_expression,
  [sym_type_query] = sym_type_query,
  [sym_index_type_query] = sym_index_type_query,
  [sym_lookup_type] = sym_lookup_type,
  [sym_mapped_type_clause] = sym_mapped_type_clause,
  [sym_literal_type] = sym_literal_type,
  [sym__number] = sym_unary_expression,
  [sym_existential_type] = sym_existential_type,
  [sym_flow_maybe_type] = sym_flow_maybe_type,
  [sym_parenthesized_type] = sym_parenthesized_type,
  [sym_predefined_type] = sym_predefined_type,
  [sym_type_arguments] = sym_type_arguments,
  [sym_object_type] = sym_object_type,
  [sym_call_signature] = sym_call_signature,
  [sym_property_signature] = sym_property_signature,
  [sym_type_parameters] = sym_type_parameters,
  [sym_type_parameter] = sym_type_parameter,
  [sym_default_type] = sym_default_type,
  [sym_constraint] = sym_constraint,
  [sym_construct_signature] = sym_construct_signature,
  [sym_index_signature] = sym_index_signature,
  [sym_array_type] = sym_array_type,
  [sym_tuple_type] = sym_tuple_type,
  [sym_readonly_type] = sym_readonly_type,
  [sym_union_type] = sym_union_type,
  [sym_intersection_type] = sym_intersection_type,
  [sym_function_type] = sym_function_type,
  [aux_sym_program_repeat1] = aux_sym_program_repeat1,
  [aux_sym_export_statement_repeat1] = aux_sym_export_statement_repeat1,
  [aux_sym_export_clause_repeat1] = aux_sym_export_clause_repeat1,
  [aux_sym_named_imports_repeat1] = aux_sym_named_imports_repeat1,
  [aux_sym_variable_declaration_repeat1] = aux_sym_variable_declaration_repeat1,
  [aux_sym_switch_body_repeat1] = aux_sym_switch_body_repeat1,
  [aux_sym_object_repeat1] = aux_sym_object_repeat1,
  [aux_sym_object_pattern_repeat1] = aux_sym_object_pattern_repeat1,
  [aux_sym_array_repeat1] = aux_sym_array_repeat1,
  [aux_sym_array_pattern_repeat1] = aux_sym_array_pattern_repeat1,
  [aux_sym_jsx_element_repeat1] = aux_sym_jsx_element_repeat1,
  [aux_sym__jsx_string_repeat1] = aux_sym__jsx_string_repeat1,
  [aux_sym__jsx_string_repeat2] = aux_sym__jsx_string_repeat2,
  [aux_sym_sequence_expression_repeat1] = aux_sym_sequence_expression_repeat1,
  [aux_sym_string_repeat1] = aux_sym_string_repeat1,
  [aux_sym_string_repeat2] = aux_sym_string_repeat2,
  [aux_sym_template_string_repeat1] = aux_sym_template_string_repeat1,
  [aux_sym_class_body_repeat1] = aux_sym_class_body_repeat1,
  [aux_sym_formal_parameters_repeat1] = aux_sym_formal_parameters_repeat1,
  [aux_sym__jsx_start_opening_element_repeat1] = aux_sym__jsx_start_opening_element_repeat1,
  [aux_sym_extends_clause_repeat1] = aux_sym_extends_clause_repeat1,
  [aux_sym_implements_clause_repeat1] = aux_sym_implements_clause_repeat1,
  [aux_sym_extends_type_clause_repeat1] = aux_sym_extends_type_clause_repeat1,
  [aux_sym_enum_body_repeat1] = aux_sym_enum_body_repeat1,
  [aux_sym_template_literal_type_repeat1] = aux_sym_template_literal_type_repeat1,
  [aux_sym_object_type_repeat1] = aux_sym_object_type_repeat1,
  [aux_sym_type_parameters_repeat1] = aux_sym_type_parameters_repeat1,
  [aux_sym_tuple_type_repeat1] = aux_sym_tuple_type_repeat1,
  [alias_sym_interface_body] = alias_sym_interface_body,
  [alias_sym_property_identifier] = alias_sym_property_identifier,
  [alias_sym_shorthand_property_identifier] = alias_sym_shorthand_property_identifier,
  [alias_sym_shorthand_property_identifier_pattern] = alias_sym_shorthand_property_identifier_pattern,
  [alias_sym_statement_identifier] = alias_sym_statement_identifier,
  [alias_sym_this_type] = alias_sym_this_type,
  [alias_sym_type_identifier] = alias_sym_type_identifier,
};

static const TSSymbolMetadata ts_symbol_metadata[] = {
  [ts_builtin_sym_end] = {
    .visible = false,
    .named = true,
  },
  [sym_identifier] = {
    .visible = true,
    .named = true,
  },
  [sym_hash_bang_line] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_export] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_STAR] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_default] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_type] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_as] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_namespace] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LBRACE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_COMMA] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_RBRACE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_typeof] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_import] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_from] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_with] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_assert] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_var] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_let] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_const] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_BANG] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_else] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_if] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_switch] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_for] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LPAREN] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_SEMI] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_RPAREN] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_await] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_in] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_of] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_while] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_do] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_try] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_break] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_continue] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_debugger] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_return] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_throw] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_COLON] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_case] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_catch] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_finally] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_yield] = {
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
  [sym_html_character_reference] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_GT] = {
    .visible = true,
    .named = false,
  },
  [sym_jsx_identifier] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_DOT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LT_SLASH] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_SLASH_GT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DQUOTE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_SQUOTE] = {
    .visible = true,
    .named = false,
  },
  [sym_unescaped_double_jsx_string_fragment] = {
    .visible = true,
    .named = true,
  },
  [sym_unescaped_single_jsx_string_fragment] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_class] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_async] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_function] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_EQ_GT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_QMARK_DOT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_new] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_using] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PLUS_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DASH_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_STAR_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_SLASH_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PERCENT_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_CARET_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_AMP_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PIPE_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_GT_GT_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_GT_GT_GT_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LT_LT_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_STAR_STAR_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_AMP_AMP_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PIPE_PIPE_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_QMARK_QMARK_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DOT_DOT_DOT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_AMP_AMP] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PIPE_PIPE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_GT_GT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_GT_GT_GT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LT_LT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_AMP] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_CARET] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PIPE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PLUS] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DASH] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_SLASH] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PERCENT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_STAR_STAR] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LT_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_EQ_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_EQ_EQ_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_BANG_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_BANG_EQ_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_GT_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_QMARK_QMARK] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_instanceof] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_TILDE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_void] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_delete] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PLUS_PLUS] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DASH_DASH] = {
    .visible = true,
    .named = false,
  },
  [sym_unescaped_double_string_fragment] = {
    .visible = true,
    .named = true,
  },
  [sym_unescaped_single_string_fragment] = {
    .visible = true,
    .named = true,
  },
  [sym_escape_sequence] = {
    .visible = true,
    .named = true,
  },
  [sym_comment] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_BQUOTE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DOLLAR_LBRACE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_SLASH2] = {
    .visible = true,
    .named = false,
  },
  [sym_regex_pattern] = {
    .visible = true,
    .named = true,
  },
  [sym_regex_flags] = {
    .visible = true,
    .named = true,
  },
  [sym_number] = {
    .visible = true,
    .named = true,
  },
  [sym_private_property_identifier] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_target] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_meta] = {
    .visible = true,
    .named = false,
  },
  [sym_this] = {
    .visible = true,
    .named = true,
  },
  [sym_super] = {
    .visible = true,
    .named = true,
  },
  [sym_true] = {
    .visible = true,
    .named = true,
  },
  [sym_false] = {
    .visible = true,
    .named = true,
  },
  [sym_null] = {
    .visible = true,
    .named = true,
  },
  [sym_undefined] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_AT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_static] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_readonly] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_get] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_set] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_QMARK] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_declare] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_public] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_private] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_protected] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_override] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_module] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_any] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_number] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_boolean] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_string] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_symbol] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_object] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_abstract] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_accessor] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_satisfies] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_require] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_extends] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_implements] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_global] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_interface] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_enum] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DASH_QMARK_COLON] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PLUS_QMARK_COLON] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_QMARK_COLON] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_asserts] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_infer] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_is] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_keyof] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_unique] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_unknown] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_never] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LBRACE_PIPE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PIPE_RBRACE] = {
    .visible = true,
    .named = false,
  },
  [sym__automatic_semicolon] = {
    .visible = false,
    .named = true,
  },
  [sym__template_chars] = {
    .visible = true,
    .named = true,
  },
  [sym__ternary_qmark] = {
    .visible = true,
    .named = false,
  },
  [sym_html_comment] = {
    .visible = true,
    .named = true,
  },
  [sym_jsx_text] = {
    .visible = true,
    .named = true,
  },
  [sym__function_signature_automatic_semicolon] = {
    .visible = false,
    .named = true,
  },
  [sym___error_recovery] = {
    .visible = false,
    .named = true,
  },
  [sym_program] = {
    .visible = true,
    .named = true,
  },
  [sym_export_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_namespace_export] = {
    .visible = true,
    .named = true,
  },
  [sym_export_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_export_specifier] = {
    .visible = true,
    .named = true,
  },
  [sym__module_export_name] = {
    .visible = false,
    .named = true,
  },
  [sym_declaration] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym_import] = {
    .visible = true,
    .named = true,
  },
  [sym_import_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_import_clause] = {
    .visible = true,
    .named = true,
  },
  [sym__from_clause] = {
    .visible = false,
    .named = true,
  },
  [sym_namespace_import] = {
    .visible = true,
    .named = true,
  },
  [sym_named_imports] = {
    .visible = true,
    .named = true,
  },
  [sym_import_specifier] = {
    .visible = true,
    .named = true,
  },
  [sym_import_attribute] = {
    .visible = true,
    .named = true,
  },
  [sym_statement] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym_expression_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_variable_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_lexical_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_variable_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_statement_block] = {
    .visible = true,
    .named = true,
  },
  [sym_else_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_if_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_switch_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_for_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_for_in_statement] = {
    .visible = true,
    .named = true,
  },
  [sym__for_header] = {
    .visible = false,
    .named = true,
  },
  [sym_while_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_do_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_try_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_with_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_break_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_continue_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_debugger_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_return_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_throw_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_empty_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_labeled_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_switch_body] = {
    .visible = true,
    .named = true,
  },
  [sym_switch_case] = {
    .visible = true,
    .named = true,
  },
  [sym_switch_default] = {
    .visible = true,
    .named = true,
  },
  [sym_catch_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_finally_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_parenthesized_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_expression] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym_primary_expression] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym_yield_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_object] = {
    .visible = true,
    .named = true,
  },
  [sym_object_pattern] = {
    .visible = true,
    .named = true,
  },
  [sym_assignment_pattern] = {
    .visible = true,
    .named = true,
  },
  [sym_object_assignment_pattern] = {
    .visible = true,
    .named = true,
  },
  [sym_array] = {
    .visible = true,
    .named = true,
  },
  [sym_array_pattern] = {
    .visible = true,
    .named = true,
  },
  [sym_jsx_element] = {
    .visible = true,
    .named = true,
  },
  [sym_jsx_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_jsx_opening_element] = {
    .visible = true,
    .named = true,
  },
  [sym_nested_identifier] = {
    .visible = true,
    .named = true,
  },
  [sym_jsx_namespace_name] = {
    .visible = true,
    .named = true,
  },
  [sym_jsx_closing_element] = {
    .visible = true,
    .named = true,
  },
  [sym_jsx_self_closing_element] = {
    .visible = true,
    .named = true,
  },
  [sym_jsx_attribute] = {
    .visible = true,
    .named = true,
  },
  [sym__jsx_string] = {
    .visible = true,
    .named = true,
  },
  [sym_class] = {
    .visible = true,
    .named = true,
  },
  [sym_class_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_class_heritage] = {
    .visible = true,
    .named = true,
  },
  [sym_function_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_function_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_generator_function] = {
    .visible = true,
    .named = true,
  },
  [sym_generator_function_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_arrow_function] = {
    .visible = true,
    .named = true,
  },
  [sym__call_signature] = {
    .visible = false,
    .named = true,
  },
  [sym__formal_parameter] = {
    .visible = false,
    .named = true,
  },
  [sym_optional_chain] = {
    .visible = true,
    .named = true,
  },
  [sym_call_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_new_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_await_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_member_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_subscript_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_assignment_expression] = {
    .visible = true,
    .named = true,
  },
  [sym__augmented_assignment_lhs] = {
    .visible = false,
    .named = true,
  },
  [sym_augmented_assignment_expression] = {
    .visible = true,
    .named = true,
  },
  [sym__initializer] = {
    .visible = false,
    .named = true,
  },
  [sym__destructuring_pattern] = {
    .visible = false,
    .named = true,
  },
  [sym_spread_element] = {
    .visible = true,
    .named = true,
  },
  [sym_ternary_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_binary_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_unary_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_update_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_sequence_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_string] = {
    .visible = true,
    .named = true,
  },
  [sym_template_string] = {
    .visible = true,
    .named = true,
  },
  [sym_template_substitution] = {
    .visible = true,
    .named = true,
  },
  [sym_regex] = {
    .visible = true,
    .named = true,
  },
  [sym_meta_property] = {
    .visible = true,
    .named = true,
  },
  [sym_arguments] = {
    .visible = true,
    .named = true,
  },
  [sym_decorator] = {
    .visible = true,
    .named = true,
  },
  [sym_decorator_member_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_decorator_call_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_class_body] = {
    .visible = true,
    .named = true,
  },
  [sym_formal_parameters] = {
    .visible = true,
    .named = true,
  },
  [sym_class_static_block] = {
    .visible = true,
    .named = true,
  },
  [sym_pattern] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym_rest_pattern] = {
    .visible = true,
    .named = true,
  },
  [sym_method_definition] = {
    .visible = true,
    .named = true,
  },
  [sym_pair] = {
    .visible = true,
    .named = true,
  },
  [sym_pair_pattern] = {
    .visible = true,
    .named = true,
  },
  [sym__property_name] = {
    .visible = false,
    .named = true,
  },
  [sym_computed_property_name] = {
    .visible = true,
    .named = true,
  },
  [sym_public_field_definition] = {
    .visible = true,
    .named = true,
  },
  [sym__import_identifier] = {
    .visible = false,
    .named = true,
  },
  [sym_non_null_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_method_signature] = {
    .visible = true,
    .named = true,
  },
  [sym_abstract_method_signature] = {
    .visible = true,
    .named = true,
  },
  [sym_function_signature] = {
    .visible = true,
    .named = true,
  },
  [sym_decorator_parenthesized_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_as_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_satisfies_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_instantiation_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_import_require_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_extends_clause] = {
    .visible = true,
    .named = true,
  },
  [sym__extends_clause_single] = {
    .visible = false,
    .named = true,
  },
  [sym_implements_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_ambient_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_abstract_class_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_module] = {
    .visible = true,
    .named = true,
  },
  [sym_internal_module] = {
    .visible = true,
    .named = true,
  },
  [sym__module] = {
    .visible = false,
    .named = true,
  },
  [sym_import_alias] = {
    .visible = true,
    .named = true,
  },
  [sym_nested_type_identifier] = {
    .visible = true,
    .named = true,
  },
  [sym_interface_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_extends_type_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_enum_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_enum_body] = {
    .visible = true,
    .named = true,
  },
  [sym_enum_assignment] = {
    .visible = true,
    .named = true,
  },
  [sym_type_alias_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_accessibility_modifier] = {
    .visible = true,
    .named = true,
  },
  [sym_override_modifier] = {
    .visible = true,
    .named = true,
  },
  [sym_required_parameter] = {
    .visible = true,
    .named = true,
  },
  [sym_optional_parameter] = {
    .visible = true,
    .named = true,
  },
  [sym__parameter_name] = {
    .visible = false,
    .named = true,
  },
  [sym_omitting_type_annotation] = {
    .visible = true,
    .named = true,
  },
  [sym_adding_type_annotation] = {
    .visible = true,
    .named = true,
  },
  [sym_opting_type_annotation] = {
    .visible = true,
    .named = true,
  },
  [sym_type_annotation] = {
    .visible = true,
    .named = true,
  },
  [sym__type_query_member_expression_in_type_annotation] = {
    .visible = true,
    .named = true,
  },
  [sym__type_query_call_expression_in_type_annotation] = {
    .visible = true,
    .named = true,
  },
  [sym_asserts] = {
    .visible = true,
    .named = true,
  },
  [sym_asserts_annotation] = {
    .visible = true,
    .named = true,
  },
  [sym_type] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym_tuple_parameter] = {
    .visible = true,
    .named = true,
  },
  [sym_optional_tuple_parameter] = {
    .visible = true,
    .named = true,
  },
  [sym_optional_type] = {
    .visible = true,
    .named = true,
  },
  [sym_rest_type] = {
    .visible = true,
    .named = true,
  },
  [sym__tuple_type_member] = {
    .visible = false,
    .named = true,
  },
  [sym_constructor_type] = {
    .visible = true,
    .named = true,
  },
  [sym_primary_type] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym_template_type] = {
    .visible = true,
    .named = true,
  },
  [sym_template_literal_type] = {
    .visible = true,
    .named = true,
  },
  [sym_infer_type] = {
    .visible = true,
    .named = true,
  },
  [sym_conditional_type] = {
    .visible = true,
    .named = true,
  },
  [sym_generic_type] = {
    .visible = true,
    .named = true,
  },
  [sym_type_predicate] = {
    .visible = true,
    .named = true,
  },
  [sym_type_predicate_annotation] = {
    .visible = true,
    .named = true,
  },
  [sym__type_query_member_expression] = {
    .visible = true,
    .named = true,
  },
  [sym__type_query_subscript_expression] = {
    .visible = true,
    .named = true,
  },
  [sym__type_query_call_expression] = {
    .visible = true,
    .named = true,
  },
  [sym__type_query_instantiation_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_type_query] = {
    .visible = true,
    .named = true,
  },
  [sym_index_type_query] = {
    .visible = true,
    .named = true,
  },
  [sym_lookup_type] = {
    .visible = true,
    .named = true,
  },
  [sym_mapped_type_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_literal_type] = {
    .visible = true,
    .named = true,
  },
  [sym__number] = {
    .visible = true,
    .named = true,
  },
  [sym_existential_type] = {
    .visible = true,
    .named = true,
  },
  [sym_flow_maybe_type] = {
    .visible = true,
    .named = true,
  },
  [sym_parenthesized_type] = {
    .visible = true,
    .named = true,
  },
  [sym_predefined_type] = {
    .visible = true,
    .named = true,
  },
  [sym_type_arguments] = {
    .visible = true,
    .named = true,
  },
  [sym_object_type] = {
    .visible = true,
    .named = true,
  },
  [sym_call_signature] = {
    .visible = true,
    .named = true,
  },
  [sym_property_signature] = {
    .visible = true,
    .named = true,
  },
  [sym_type_parameters] = {
    .visible = true,
    .named = true,
  },
  [sym_type_parameter] = {
    .visible = true,
    .named = true,
  },
  [sym_default_type] = {
    .visible = true,
    .named = true,
  },
  [sym_constraint] = {
    .visible = true,
    .named = true,
  },
  [sym_construct_signature] = {
    .visible = true,
    .named = true,
  },
  [sym_index_signature] = {
    .visible = true,
    .named = true,
  },
  [sym_array_type] = {
    .visible = true,
    .named = true,
  },
  [sym_tuple_type] = {
    .visible = true,
    .named = true,
  },
  [sym_readonly_type] = {
    .visible = true,
    .named = true,
  },
  [sym_union_type] = {
    .visible = true,
    .named = true,
  },
  [sym_intersection_type] = {
    .visible = true,
    .named = true,
  },
  [sym_function_type] = {
    .visible = true,
    .named = true,
  },
  [aux_sym_program_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_export_statement_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_export_clause_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_named_imports_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_variable_declaration_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_switch_body_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_object_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_object_pattern_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_array_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_array_pattern_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_jsx_element_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__jsx_string_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__jsx_string_repeat2] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_sequence_expression_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_string_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_string_repeat2] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_template_string_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_class_body_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_formal_parameters_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__jsx_start_opening_element_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_extends_clause_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_implements_clause_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_extends_type_clause_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_enum_body_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_template_literal_type_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_object_type_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_type_parameters_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_tuple_type_repeat1] = {
    .visible = false,
    .named = false,
  },
  [alias_sym_interface_body] = {
    .visible = true,
    .named = true,
  },
  [alias_sym_property_identifier] = {
    .visible = true,
    .named = true,
  },
  [alias_sym_shorthand_property_identifier] = {
    .visible = true,
    .named = true,
  },
  [alias_sym_shorthand_property_identifier_pattern] = {
    .visible = true,
    .named = true,
  },
  [alias_sym_statement_identifier] = {
    .visible = true,
    .named = true,
  },
  [alias_sym_this_type] = {
    .visible = true,
    .named = true,
  },
  [alias_sym_type_identifier] = {
    .visible = true,
    .named = true,
  },
};

enum ts_field_identifiers {
  field_alias = 1,
  field_alternative = 2,
  field_argument = 3,
  field_arguments = 4,
  field_attribute = 5,
  field_body = 6,
  field_close_tag = 7,
  field_condition = 8,
  field_consequence = 9,
  field_constraint = 10,
  field_constructor = 11,
  field_declaration = 12,
  field_decorator = 13,
  field_finalizer = 14,
  field_flags = 15,
  field_function = 16,
  field_handler = 17,
  field_increment = 18,
  field_index = 19,
  field_index_type = 20,
  field_initializer = 21,
  field_key = 22,
  field_kind = 23,
  field_label = 24,
  field_left = 25,
  field_module = 26,
  field_name = 27,
  field_object = 28,
  field_open_tag = 29,
  field_operator = 30,
  field_optional_chain = 31,
  field_parameter = 32,
  field_parameters = 33,
  field_pattern = 34,
  field_property = 35,
  field_return_type = 36,
  field_right = 37,
  field_sign = 38,
  field_source = 39,
  field_type = 40,
  field_type_arguments = 41,
  field_type_parameters = 42,
  field_value = 43,
};

static const char * const ts_field_names[] = {
  [0] = NULL,
  [field_alias] = "alias",
  [field_alternative] = "alternative",
  [field_argument] = "argument",
  [field_arguments] = "arguments",
  [field_attribute] = "attribute",
  [field_body] = "body",
  [field_close_tag] = "close_tag",
  [field_condition] = "condition",
  [field_consequence] = "consequence",
  [field_constraint] = "constraint",
  [field_constructor] = "constructor",
  [field_declaration] = "declaration",
  [field_decorator] = "decorator",
  [field_finalizer] = "finalizer",
  [field_flags] = "flags",
  [field_function] = "function",
  [field_handler] = "handler",
  [field_increment] = "increment",
  [field_index] = "index",
  [field_index_type] = "index_type",
  [field_initializer] = "initializer",
  [field_key] = "key",
  [field_kind] = "kind",
  [field_label] = "label",
  [field_left] = "left",
  [field_module] = "module",
  [field_name] = "name",
  [field_object] = "object",
  [field_open_tag] = "open_tag",
  [field_operator] = "operator",
  [field_optional_chain] = "optional_chain",
  [field_parameter] = "parameter",
  [field_parameters] = "parameters",
  [field_pattern] = "pattern",
  [field_property] = "property",
  [field_return_type] = "return_type",
  [field_right] = "right",
  [field_sign] = "sign",
  [field_source] = "source",
  [field_type] = "type",
  [field_type_arguments] = "type_arguments",
  [field_type_parameters] = "type_parameters",
  [field_value] = "value",
};

static const TSFieldMapSlice ts_field_map_slices[PRODUCTION_ID_COUNT] = {
  [2] = {.index = 0, .length = 1},
  [3] = {.index = 1, .length = 1},
  [4] = {.index = 2, .length = 1},
  [5] = {.index = 3, .length = 1},
  [6] = {.index = 4, .length = 2},
  [8] = {.index = 6, .length = 2},
  [9] = {.index = 8, .length = 1},
  [10] = {.index = 9, .length = 2},
  [11] = {.index = 11, .length = 1},
  [12] = {.index = 12, .length = 1},
  [13] = {.index = 3, .length = 1},
  [14] = {.index = 13, .length = 2},
  [15] = {.index = 15, .length = 2},
  [16] = {.index = 17, .length = 1},
  [17] = {.index = 18, .length = 2},
  [18] = {.index = 20, .length = 2},
  [19] = {.index = 22, .length = 2},
  [20] = {.index = 24, .length = 2},
  [21] = {.index = 26, .length = 1},
  [22] = {.index = 27, .length = 2},
  [23] = {.index = 29, .length = 2},
  [24] = {.index = 31, .length = 2},
  [25] = {.index = 33, .length = 2},
  [29] = {.index = 35, .length = 1},
  [30] = {.index = 36, .length = 2},
  [31] = {.index = 38, .length = 2},
  [32] = {.index = 40, .length = 2},
  [33] = {.index = 42, .length = 1},
  [34] = {.index = 43, .length = 2},
  [35] = {.index = 45, .length = 2},
  [36] = {.index = 47, .length = 6},
  [37] = {.index = 53, .length = 1},
  [38] = {.index = 54, .length = 3},
  [39] = {.index = 57, .length = 3},
  [40] = {.index = 60, .length = 2},
  [41] = {.index = 62, .length = 2},
  [42] = {.index = 64, .length = 2},
  [43] = {.index = 66, .length = 2},
  [44] = {.index = 68, .length = 1},
  [45] = {.index = 69, .length = 2},
  [46] = {.index = 71, .length = 1},
  [47] = {.index = 72, .length = 2},
  [50] = {.index = 74, .length = 2},
  [51] = {.index = 76, .length = 2},
  [52] = {.index = 78, .length = 2},
  [53] = {.index = 80, .length = 1},
  [54] = {.index = 81, .length = 2},
  [55] = {.index = 83, .length = 4},
  [56] = {.index = 87, .length = 2},
  [57] = {.index = 89, .length = 2},
  [58] = {.index = 91, .length = 1},
  [59] = {.index = 91, .length = 1},
  [60] = {.index = 92, .length = 1},
  [61] = {.index = 93, .length = 2},
  [62] = {.index = 95, .length = 2},
  [63] = {.index = 91, .length = 1},
  [64] = {.index = 27, .length = 2},
  [65] = {.index = 31, .length = 2},
  [66] = {.index = 97, .length = 3},
  [67] = {.index = 69, .length = 2},
  [68] = {.index = 69, .length = 2},
  [69] = {.index = 100, .length = 2},
  [70] = {.index = 100, .length = 2},
  [71] = {.index = 102, .length = 3},
  [72] = {.index = 102, .length = 3},
  [73] = {.index = 105, .length = 3},
  [74] = {.index = 108, .length = 2},
  [75] = {.index = 110, .length = 2},
  [76] = {.index = 112, .length = 4},
  [77] = {.index = 116, .length = 3},
  [78] = {.index = 119, .length = 2},
  [79] = {.index = 121, .length = 2},
  [80] = {.index = 123, .length = 1},
  [81] = {.index = 124, .length = 1},
  [82] = {.index = 100, .length = 2},
  [83] = {.index = 27, .length = 2},
  [84] = {.index = 125, .length = 2},
  [85] = {.index = 127, .length = 5},
  [86] = {.index = 132, .length = 1},
  [87] = {.index = 133, .length = 1},
  [88] = {.index = 134, .length = 2},
  [89] = {.index = 136, .length = 3},
  [90] = {.index = 139, .length = 2},
  [91] = {.index = 141, .length = 3},
  [92] = {.index = 144, .length = 6},
  [93] = {.index = 150, .length = 1},
  [94] = {.index = 151, .length = 1},
  [95] = {.index = 152, .length = 3},
  [96] = {.index = 155, .length = 3},
  [97] = {.index = 158, .length = 4},
  [98] = {.index = 162, .length = 2},
  [99] = {.index = 164, .length = 2},
  [100] = {.index = 166, .length = 3},
  [101] = {.index = 169, .length = 4},
  [102] = {.index = 173, .length = 1},
  [103] = {.index = 174, .length = 2},
  [104] = {.index = 176, .length = 1},
  [105] = {.index = 177, .length = 2},
  [106] = {.index = 179, .length = 3},
  [107] = {.index = 182, .length = 2},
  [108] = {.index = 184, .length = 4},
  [109] = {.index = 188, .length = 3},
  [110] = {.index = 191, .length = 2},
  [111] = {.index = 193, .length = 2},
  [112] = {.index = 195, .length = 2},
  [113] = {.index = 197, .length = 2},
  [114] = {.index = 199, .length = 2},
  [117] = {.index = 199, .length = 2},
  [118] = {.index = 201, .length = 4},
  [119] = {.index = 205, .length = 2},
  [120] = {.index = 207, .length = 2},
  [121] = {.index = 209, .length = 4},
  [122] = {.index = 207, .length = 2},
  [123] = {.index = 213, .length = 4},
  [124] = {.index = 217, .length = 4},
  [125] = {.index = 221, .length = 5},
  [126] = {.index = 226, .length = 3},
  [127] = {.index = 229, .length = 2},
  [128] = {.index = 229, .length = 2},
  [129] = {.index = 231, .length = 2},
  [130] = {.index = 233, .length = 2},
  [131] = {.index = 235, .length = 2},
  [132] = {.index = 237, .length = 2},
  [133] = {.index = 239, .length = 2},
  [134] = {.index = 241, .length = 2},
  [135] = {.index = 243, .length = 3},
  [136] = {.index = 241, .length = 2},
  [137] = {.index = 237, .length = 2},
  [138] = {.index = 139, .length = 2},
  [139] = {.index = 246, .length = 1},
  [140] = {.index = 246, .length = 1},
  [141] = {.index = 177, .length = 2},
  [142] = {.index = 179, .length = 3},
  [143] = {.index = 247, .length = 2},
  [144] = {.index = 249, .length = 3},
  [145] = {.index = 252, .length = 2},
  [146] = {.index = 254, .length = 3},
  [147] = {.index = 257, .length = 2},
  [148] = {.index = 259, .length = 3},
  [149] = {.index = 262, .length = 1},
  [150] = {.index = 263, .length = 2},
  [151] = {.index = 265, .length = 2},
  [152] = {.index = 267, .length = 5},
  [153] = {.index = 263, .length = 2},
  [154] = {.index = 272, .length = 1},
  [155] = {.index = 273, .length = 4},
  [156] = {.index = 277, .length = 2},
  [157] = {.index = 279, .length = 1},
  [158] = {.index = 280, .length = 2},
  [159] = {.index = 282, .length = 2},
  [160] = {.index = 284, .length = 2},
  [161] = {.index = 286, .length = 4},
  [162] = {.index = 290, .length = 2},
  [163] = {.index = 292, .length = 3},
  [164] = {.index = 295, .length = 3},
  [165] = {.index = 298, .length = 3},
  [166] = {.index = 301, .length = 4},
  [167] = {.index = 305, .length = 1},
  [168] = {.index = 306, .length = 2},
  [169] = {.index = 308, .length = 4},
  [170] = {.index = 312, .length = 4},
  [171] = {.index = 316, .length = 4},
  [172] = {.index = 320, .length = 2},
  [173] = {.index = 322, .length = 2},
  [174] = {.index = 324, .length = 4},
  [175] = {.index = 324, .length = 4},
  [176] = {.index = 328, .length = 4},
  [177] = {.index = 328, .length = 4},
  [178] = {.index = 332, .length = 4},
  [179] = {.index = 336, .length = 4},
  [180] = {.index = 340, .length = 5},
  [181] = {.index = 345, .length = 3},
  [182] = {.index = 348, .length = 3},
  [183] = {.index = 348, .length = 3},
  [184] = {.index = 351, .length = 2},
  [185] = {.index = 353, .length = 3},
  [186] = {.index = 356, .length = 2},
  [187] = {.index = 356, .length = 2},
  [188] = {.index = 298, .length = 3},
  [189] = {.index = 358, .length = 3},
  [190] = {.index = 361, .length = 3},
  [191] = {.index = 134, .length = 2},
  [192] = {.index = 364, .length = 2},
  [193] = {.index = 366, .length = 3},
  [194] = {.index = 369, .length = 4},
  [195] = {.index = 373, .length = 3},
  [196] = {.index = 376, .length = 3},
  [197] = {.index = 379, .length = 2},
  [198] = {.index = 381, .length = 3},
  [199] = {.index = 384, .length = 5},
  [200] = {.index = 379, .length = 2},
  [201] = {.index = 389, .length = 3},
  [202] = {.index = 389, .length = 3},
  [203] = {.index = 392, .length = 3},
  [204] = {.index = 395, .length = 2},
  [205] = {.index = 397, .length = 4},
  [206] = {.index = 134, .length = 2},
  [207] = {.index = 401, .length = 1},
  [208] = {.index = 402, .length = 2},
  [209] = {.index = 404, .length = 2},
  [210] = {.index = 406, .length = 2},
  [211] = {.index = 408, .length = 2},
  [212] = {.index = 410, .length = 3},
  [213] = {.index = 413, .length = 1},
  [214] = {.index = 414, .length = 3},
  [215] = {.index = 417, .length = 2},
  [216] = {.index = 419, .length = 3},
  [217] = {.index = 422, .length = 3},
  [218] = {.index = 425, .length = 3},
  [219] = {.index = 428, .length = 3},
  [220] = {.index = 431, .length = 4},
  [221] = {.index = 435, .length = 2},
  [222] = {.index = 437, .length = 2},
  [223] = {.index = 439, .length = 1},
  [224] = {.index = 440, .length = 4},
  [225] = {.index = 440, .length = 4},
  [226] = {.index = 444, .length = 2},
  [227] = {.index = 446, .length = 3},
  [228] = {.index = 449, .length = 5},
  [229] = {.index = 454, .length = 3},
  [230] = {.index = 457, .length = 2},
  [231] = {.index = 459, .length = 2},
  [232] = {.index = 461, .length = 2},
  [233] = {.index = 463, .length = 1},
  [234] = {.index = 464, .length = 4},
  [235] = {.index = 468, .length = 3},
  [236] = {.index = 471, .length = 4},
  [237] = {.index = 475, .length = 5},
  [238] = {.index = 480, .length = 1},
  [239] = {.index = 481, .length = 2},
  [240] = {.index = 483, .length = 4},
  [241] = {.index = 487, .length = 4},
  [242] = {.index = 491, .length = 4},
  [243] = {.index = 495, .length = 3},
  [244] = {.index = 498, .length = 2},
  [245] = {.index = 500, .length = 4},
  [246] = {.index = 504, .length = 4},
  [247] = {.index = 508, .length = 2},
  [248] = {.index = 510, .length = 2},
  [249] = {.index = 512, .length = 3},
  [250] = {.index = 515, .length = 3},
  [251] = {.index = 518, .length = 2},
  [252] = {.index = 520, .length = 2},
  [253] = {.index = 522, .length = 1},
  [254] = {.index = 523, .length = 1},
  [255] = {.index = 524, .length = 3},
  [256] = {.index = 527, .length = 3},
  [257] = {.index = 530, .length = 3},
  [258] = {.index = 533, .length = 3},
  [259] = {.index = 536, .length = 4},
  [260] = {.index = 540, .length = 2},
  [261] = {.index = 542, .length = 4},
  [262] = {.index = 546, .length = 3},
  [263] = {.index = 549, .length = 2},
  [264] = {.index = 551, .length = 2},
  [265] = {.index = 553, .length = 4},
  [266] = {.index = 557, .length = 4},
  [267] = {.index = 561, .length = 4},
  [268] = {.index = 565, .length = 3},
  [269] = {.index = 568, .length = 2},
  [271] = {.index = 570, .length = 4},
  [272] = {.index = 574, .length = 5},
  [273] = {.index = 579, .length = 5},
  [274] = {.index = 584, .length = 5},
  [275] = {.index = 589, .length = 4},
  [276] = {.index = 593, .length = 5},
  [277] = {.index = 598, .length = 4},
  [278] = {.index = 602, .length = 4},
  [279] = {.index = 606, .length = 3},
  [280] = {.index = 609, .length = 3},
  [281] = {.index = 612, .length = 3},
  [282] = {.index = 609, .length = 3},
  [283] = {.index = 615, .length = 2},
  [284] = {.index = 617, .length = 4},
  [285] = {.index = 621, .length = 4},
  [286] = {.index = 625, .length = 3},
  [287] = {.index = 628, .length = 2},
  [288] = {.index = 630, .length = 2},
  [289] = {.index = 632, .length = 3},
  [290] = {.index = 635, .length = 2},
  [291] = {.index = 637, .length = 2},
  [292] = {.index = 639, .length = 1},
  [293] = {.index = 640, .length = 3},
  [294] = {.index = 643, .length = 3},
  [295] = {.index = 646, .length = 4},
  [296] = {.index = 650, .length = 4},
  [297] = {.index = 654, .length = 3},
  [298] = {.index = 657, .length = 3},
  [299] = {.index = 660, .length = 2},
  [300] = {.index = 662, .length = 3},
  [301] = {.index = 665, .length = 2},
  [302] = {.index = 667, .length = 4},
  [303] = {.index = 671, .length = 4},
  [304] = {.index = 675, .length = 4},
  [305] = {.index = 679, .length = 3},
  [306] = {.index = 682, .length = 5},
  [307] = {.index = 687, .length = 5},
  [308] = {.index = 692, .length = 5},
  [309] = {.index = 697, .length = 4},
  [310] = {.index = 701, .length = 4},
  [311] = {.index = 705, .length = 3},
  [312] = {.index = 708, .length = 3},
  [313] = {.index = 708, .length = 3},
  [314] = {.index = 711, .length = 2},
  [315] = {.index = 713, .length = 2},
  [316] = {.index = 715, .length = 3},
  [317] = {.index = 718, .length = 2},
  [318] = {.index = 720, .length = 2},
  [319] = {.index = 722, .length = 4},
  [320] = {.index = 726, .length = 3},
  [321] = {.index = 729, .length = 3},
  [322] = {.index = 732, .length = 4},
  [323] = {.index = 736, .length = 3},
  [324] = {.index = 739, .length = 3},
  [325] = {.index = 742, .length = 2},
  [326] = {.index = 744, .length = 4},
  [327] = {.index = 748, .length = 5},
  [328] = {.index = 753, .length = 5},
  [329] = {.index = 758, .length = 4},
  [330] = {.index = 758, .length = 4},
  [331] = {.index = 762, .length = 4},
  [332] = {.index = 766, .length = 3},
  [333] = {.index = 769, .length = 2},
  [334] = {.index = 771, .length = 2},
  [335] = {.index = 773, .length = 3},
  [336] = {.index = 776, .length = 4},
  [337] = {.index = 780, .length = 4},
  [338] = {.index = 784, .length = 3},
  [339] = {.index = 787, .length = 3},
  [340] = {.index = 790, .length = 4},
  [341] = {.index = 794, .length = 3},
  [342] = {.index = 797, .length = 3},
  [343] = {.index = 800, .length = 5},
  [344] = {.index = 805, .length = 3},
  [345] = {.index = 808, .length = 4},
  [346] = {.index = 812, .length = 4},
  [347] = {.index = 816, .length = 3},
  [348] = {.index = 819, .length = 3},
  [349] = {.index = 822, .length = 4},
  [350] = {.index = 826, .length = 4},
};

static const TSFieldMapEntry ts_field_map_entries[] = {
  [0] =
    {field_decorator, 0},
  [1] =
    {field_parameters, 0},
  [2] =
    {field_declaration, 1},
  [3] =
    {field_name, 0},
  [4] =
    {field_body, 1, .inherited = true},
    {field_name, 1, .inherited = true},
  [6] =
    {field_argument, 1},
    {field_operator, 0},
  [8] =
    {field_pattern, 0},
  [9] =
    {field_decorator, 0, .inherited = true},
    {field_pattern, 0, .inherited = true},
  [11] =
    {field_body, 1},
  [12] =
    {field_constructor, 1},
  [13] =
    {field_arguments, 1},
    {field_function, 0},
  [15] =
    {field_argument, 0},
    {field_operator, 1},
  [17] =
    {field_type_arguments, 1},
  [18] =
    {field_close_tag, 1},
    {field_open_tag, 0},
  [20] =
    {field_parameters, 0},
    {field_return_type, 1},
  [22] =
    {field_parameters, 1},
    {field_type_parameters, 0},
  [24] =
    {field_decorator, 0, .inherited = true},
    {field_decorator, 1, .inherited = true},
  [26] =
    {field_declaration, 2},
  [27] =
    {field_left, 0},
    {field_right, 2},
  [29] =
    {field_body, 2},
    {field_label, 0},
  [31] =
    {field_body, 2},
    {field_parameter, 0},
  [33] =
    {field_body, 1},
    {field_name, 0},
  [35] =
    {field_source, 1},
  [36] =
    {field_body, 2},
    {field_object, 1},
  [38] =
    {field_name, 0},
    {field_value, 1, .inherited = true},
  [40] =
    {field_name, 0},
    {field_type, 1},
  [42] =
    {field_kind, 0},
  [43] =
    {field_condition, 1},
    {field_consequence, 2},
  [45] =
    {field_body, 2},
    {field_value, 1},
  [47] =
    {field_body, 2},
    {field_kind, 1, .inherited = true},
    {field_left, 1, .inherited = true},
    {field_operator, 1, .inherited = true},
    {field_right, 1, .inherited = true},
    {field_value, 1, .inherited = true},
  [53] =
    {field_pattern, 1},
  [54] =
    {field_decorator, 0, .inherited = true},
    {field_pattern, 0, .inherited = true},
    {field_value, 1, .inherited = true},
  [57] =
    {field_decorator, 0, .inherited = true},
    {field_pattern, 0, .inherited = true},
    {field_type, 1},
  [60] =
    {field_decorator, 0, .inherited = true},
    {field_pattern, 1},
  [62] =
    {field_body, 2},
    {field_condition, 1},
  [64] =
    {field_body, 1},
    {field_handler, 2},
  [66] =
    {field_body, 1},
    {field_finalizer, 2},
  [68] =
    {field_label, 1},
  [69] =
    {field_body, 2},
    {field_name, 1},
  [71] =
    {field_value, 0},
  [72] =
    {field_type_arguments, 1, .inherited = true},
    {field_value, 1, .inherited = true},
  [74] =
    {field_object, 0, .inherited = true},
    {field_property, 0, .inherited = true},
  [76] =
    {field_arguments, 0, .inherited = true},
    {field_function, 0, .inherited = true},
  [78] =
    {field_argument, 0, .inherited = true},
    {field_operator, 0, .inherited = true},
  [80] =
    {field_body, 2},
  [81] =
    {field_body, 2},
    {field_type_parameters, 1},
  [83] =
    {field_body, 2},
    {field_parameters, 1, .inherited = true},
    {field_return_type, 1, .inherited = true},
    {field_type_parameters, 1, .inherited = true},
  [87] =
    {field_arguments, 2},
    {field_constructor, 1},
  [89] =
    {field_constructor, 1},
    {field_type_arguments, 2},
  [91] =
    {field_name, 1},
  [92] =
    {field_attribute, 0},
  [93] =
    {field_name, 0},
    {field_value, 1},
  [95] =
    {field_constraint, 1},
    {field_name, 0},
  [97] =
    {field_left, 0},
    {field_operator, 1},
    {field_right, 2},
  [100] =
    {field_object, 0},
    {field_property, 2},
  [102] =
    {field_object, 0},
    {field_optional_chain, 1},
    {field_property, 2},
  [105] =
    {field_arguments, 2},
    {field_function, 0},
    {field_type_arguments, 1},
  [108] =
    {field_arguments, 2},
    {field_function, 0},
  [110] =
    {field_close_tag, 2},
    {field_open_tag, 0},
  [112] =
    {field_body, 2},
    {field_parameters, 0, .inherited = true},
    {field_return_type, 0, .inherited = true},
    {field_type_parameters, 0, .inherited = true},
  [116] =
    {field_parameters, 1},
    {field_return_type, 2},
    {field_type_parameters, 0},
  [119] =
    {field_declaration, 2},
    {field_decorator, 0, .inherited = true},
  [121] =
    {field_body, 2},
    {field_decorator, 0, .inherited = true},
  [123] =
    {field_source, 2, .inherited = true},
  [124] =
    {field_value, 2},
  [125] =
    {field_key, 0},
    {field_value, 2},
  [127] =
    {field_body, 2},
    {field_name, 0},
    {field_parameters, 1, .inherited = true},
    {field_return_type, 1, .inherited = true},
    {field_type_parameters, 1, .inherited = true},
  [132] =
    {field_source, 2},
  [133] =
    {field_value, 1},
  [134] =
    {field_name, 0},
    {field_type, 2},
  [136] =
    {field_name, 0},
    {field_type, 1},
    {field_value, 2, .inherited = true},
  [139] =
    {field_body, 3},
    {field_name, 2},
  [141] =
    {field_alternative, 3},
    {field_condition, 1},
    {field_consequence, 2},
  [144] =
    {field_body, 3},
    {field_kind, 2, .inherited = true},
    {field_left, 2, .inherited = true},
    {field_operator, 2, .inherited = true},
    {field_right, 2, .inherited = true},
    {field_value, 2, .inherited = true},
  [150] =
    {field_type, 2},
  [151] =
    {field_pattern, 2},
  [152] =
    {field_decorator, 0, .inherited = true},
    {field_pattern, 0, .inherited = true},
    {field_value, 2, .inherited = true},
  [155] =
    {field_decorator, 0, .inherited = true},
    {field_pattern, 0, .inherited = true},
    {field_type, 2},
  [158] =
    {field_decorator, 0, .inherited = true},
    {field_pattern, 0, .inherited = true},
    {field_type, 1},
    {field_value, 2, .inherited = true},
  [162] =
    {field_decorator, 0, .inherited = true},
    {field_pattern, 2},
  [164] =
    {field_body, 1},
    {field_condition, 3},
  [166] =
    {field_body, 1},
    {field_finalizer, 3},
    {field_handler, 2},
  [169] =
    {field_name, 0},
    {field_parameters, 1, .inherited = true},
    {field_return_type, 1, .inherited = true},
    {field_type_parameters, 1, .inherited = true},
  [173] =
    {field_decorator, 0, .inherited = true},
  [174] =
    {field_decorator, 0, .inherited = true},
    {field_name, 1},
  [176] =
    {field_decorator, 1, .inherited = true},
  [177] =
    {field_body, 3},
    {field_name, 1},
  [179] =
    {field_body, 3},
    {field_name, 1},
    {field_type_parameters, 2},
  [182] =
    {field_type_arguments, 1},
    {field_value, 0},
  [184] =
    {field_type_arguments, 1, .inherited = true},
    {field_type_arguments, 2, .inherited = true},
    {field_value, 1, .inherited = true},
    {field_value, 2, .inherited = true},
  [188] =
    {field_parameters, 0, .inherited = true},
    {field_return_type, 0, .inherited = true},
    {field_type_parameters, 0, .inherited = true},
  [191] =
    {field_object, 1, .inherited = true},
    {field_property, 1, .inherited = true},
  [193] =
    {field_index, 1, .inherited = true},
    {field_object, 1, .inherited = true},
  [195] =
    {field_arguments, 1, .inherited = true},
    {field_function, 1, .inherited = true},
  [197] =
    {field_function, 1, .inherited = true},
    {field_type_arguments, 1, .inherited = true},
  [199] =
    {field_name, 0},
    {field_type_arguments, 1},
  [201] =
    {field_arguments, 1},
    {field_function, 0},
    {field_object, 0, .inherited = true},
    {field_property, 0, .inherited = true},
  [205] =
    {field_body, 3},
    {field_type_parameters, 1},
  [207] =
    {field_body, 3},
    {field_parameter, 1},
  [209] =
    {field_body, 3},
    {field_parameters, 2, .inherited = true},
    {field_return_type, 2, .inherited = true},
    {field_type_parameters, 2, .inherited = true},
  [213] =
    {field_body, 3},
    {field_parameters, 1, .inherited = true},
    {field_return_type, 1, .inherited = true},
    {field_type_parameters, 1, .inherited = true},
  [217] =
    {field_name, 1},
    {field_parameters, 2, .inherited = true},
    {field_return_type, 2, .inherited = true},
    {field_type_parameters, 2, .inherited = true},
  [221] =
    {field_body, 3},
    {field_name, 1},
    {field_parameters, 2, .inherited = true},
    {field_return_type, 2, .inherited = true},
    {field_type_parameters, 2, .inherited = true},
  [226] =
    {field_arguments, 3},
    {field_constructor, 1},
    {field_type_arguments, 2},
  [229] =
    {field_left, 1},
    {field_right, 3},
  [231] =
    {field_flags, 3},
    {field_pattern, 1},
  [233] =
    {field_name, 1},
    {field_value, 2},
  [235] =
    {field_constraint, 2},
    {field_name, 1},
  [237] =
    {field_attribute, 2, .inherited = true},
    {field_name, 1},
  [239] =
    {field_attribute, 0, .inherited = true},
    {field_attribute, 1, .inherited = true},
  [241] =
    {field_name, 1},
    {field_type_arguments, 2},
  [243] =
    {field_constraint, 1},
    {field_name, 0},
    {field_value, 2},
  [246] =
    {field_type, 1},
  [247] =
    {field_index, 2},
    {field_object, 0},
  [249] =
    {field_arguments, 3},
    {field_function, 0},
    {field_type_arguments, 2},
  [252] =
    {field_declaration, 3},
    {field_decorator, 0, .inherited = true},
  [254] =
    {field_body, 3},
    {field_decorator, 0, .inherited = true},
    {field_name, 2},
  [257] =
    {field_body, 3},
    {field_decorator, 0, .inherited = true},
  [259] =
    {field_body, 3},
    {field_decorator, 0, .inherited = true},
    {field_type_parameters, 2},
  [262] =
    {field_source, 3, .inherited = true},
  [263] =
    {field_alias, 2},
    {field_name, 0},
  [265] =
    {field_name, 1},
    {field_value, 3},
  [267] =
    {field_body, 3},
    {field_name, 0},
    {field_parameters, 2, .inherited = true},
    {field_return_type, 2, .inherited = true},
    {field_type_parameters, 2, .inherited = true},
  [272] =
    {field_pattern, 3},
  [273] =
    {field_decorator, 0, .inherited = true},
    {field_pattern, 0, .inherited = true},
    {field_type, 2},
    {field_value, 3, .inherited = true},
  [277] =
    {field_decorator, 0, .inherited = true},
    {field_pattern, 3},
  [279] =
    {field_name, 2},
  [280] =
    {field_name, 1},
    {field_value, 2, .inherited = true},
  [282] =
    {field_name, 1},
    {field_type, 2},
  [284] =
    {field_name, 0},
    {field_value, 2, .inherited = true},
  [286] =
    {field_name, 0},
    {field_parameters, 2, .inherited = true},
    {field_return_type, 2, .inherited = true},
    {field_type_parameters, 2, .inherited = true},
  [290] =
    {field_decorator, 0, .inherited = true},
    {field_name, 2},
  [292] =
    {field_decorator, 0, .inherited = true},
    {field_name, 1},
    {field_value, 2, .inherited = true},
  [295] =
    {field_decorator, 0, .inherited = true},
    {field_name, 1},
    {field_type, 2},
  [298] =
    {field_body, 4},
    {field_name, 1},
    {field_type_parameters, 2},
  [301] =
    {field_type_arguments, 0, .inherited = true},
    {field_type_arguments, 1, .inherited = true},
    {field_value, 0, .inherited = true},
    {field_value, 1, .inherited = true},
  [305] =
    {field_parameters, 1},
  [306] =
    {field_function, 0},
    {field_type_arguments, 1},
  [308] =
    {field_function, 0},
    {field_object, 0, .inherited = true},
    {field_property, 0, .inherited = true},
    {field_type_arguments, 1},
  [312] =
    {field_arguments, 1},
    {field_function, 0},
    {field_index, 0, .inherited = true},
    {field_object, 0, .inherited = true},
  [316] =
    {field_function, 0},
    {field_index, 0, .inherited = true},
    {field_object, 0, .inherited = true},
    {field_type_arguments, 1},
  [320] =
    {field_module, 0},
    {field_name, 2},
  [322] =
    {field_parameters, 0},
    {field_return_type, 2},
  [324] =
    {field_object, 0},
    {field_object, 0, .inherited = true},
    {field_property, 0, .inherited = true},
    {field_property, 2},
  [328] =
    {field_arguments, 0, .inherited = true},
    {field_function, 0, .inherited = true},
    {field_object, 0},
    {field_property, 2},
  [332] =
    {field_body, 4},
    {field_parameters, 3, .inherited = true},
    {field_return_type, 3, .inherited = true},
    {field_type_parameters, 3, .inherited = true},
  [336] =
    {field_name, 2},
    {field_parameters, 3, .inherited = true},
    {field_return_type, 3, .inherited = true},
    {field_type_parameters, 3, .inherited = true},
  [340] =
    {field_body, 4},
    {field_name, 2},
    {field_parameters, 3, .inherited = true},
    {field_return_type, 3, .inherited = true},
    {field_type_parameters, 3, .inherited = true},
  [345] =
    {field_constraint, 2},
    {field_name, 1},
    {field_value, 3},
  [348] =
    {field_attribute, 3, .inherited = true},
    {field_name, 1},
    {field_type_arguments, 2},
  [351] =
    {field_body, 4},
    {field_name, 2},
  [353] =
    {field_body, 4},
    {field_name, 2},
    {field_type_parameters, 3},
  [356] =
    {field_type, 1},
    {field_type, 2, .inherited = true},
  [358] =
    {field_alternative, 4},
    {field_condition, 0},
    {field_consequence, 2},
  [361] =
    {field_index, 3},
    {field_object, 0},
    {field_optional_chain, 1},
  [364] =
    {field_decorator, 0, .inherited = true},
    {field_value, 3},
  [366] =
    {field_body, 4},
    {field_decorator, 0, .inherited = true},
    {field_name, 2},
  [369] =
    {field_body, 4},
    {field_decorator, 0, .inherited = true},
    {field_name, 2},
    {field_type_parameters, 3},
  [373] =
    {field_body, 4},
    {field_decorator, 0, .inherited = true},
    {field_type_parameters, 2},
  [376] =
    {field_body, 4},
    {field_decorator, 0, .inherited = true},
    {field_name, 3},
  [379] =
    {field_alias, 3},
    {field_name, 1},
  [381] =
    {field_name, 1},
    {field_type_parameters, 2},
    {field_value, 4},
  [384] =
    {field_body, 4},
    {field_name, 1},
    {field_parameters, 3, .inherited = true},
    {field_return_type, 3, .inherited = true},
    {field_type_parameters, 3, .inherited = true},
  [389] =
    {field_left, 1},
    {field_operator, 2},
    {field_right, 3},
  [392] =
    {field_body, 5},
    {field_condition, 3},
    {field_initializer, 2},
  [395] =
    {field_decorator, 0, .inherited = true},
    {field_pattern, 4},
  [397] =
    {field_name, 1},
    {field_parameters, 3, .inherited = true},
    {field_return_type, 3, .inherited = true},
    {field_type_parameters, 3, .inherited = true},
  [401] =
    {field_type, 3},
  [402] =
    {field_name, 2},
    {field_value, 3, .inherited = true},
  [404] =
    {field_name, 2},
    {field_type, 3},
  [406] =
    {field_name, 1},
    {field_value, 3, .inherited = true},
  [408] =
    {field_name, 1},
    {field_type, 3},
  [410] =
    {field_name, 1},
    {field_type, 2},
    {field_value, 3, .inherited = true},
  [413] =
    {field_name, 3},
  [414] =
    {field_name, 0},
    {field_type, 2},
    {field_value, 3, .inherited = true},
  [417] =
    {field_decorator, 0, .inherited = true},
    {field_name, 3},
  [419] =
    {field_decorator, 0, .inherited = true},
    {field_name, 2},
    {field_value, 3, .inherited = true},
  [422] =
    {field_decorator, 0, .inherited = true},
    {field_name, 2},
    {field_type, 3},
  [425] =
    {field_decorator, 0, .inherited = true},
    {field_name, 1},
    {field_value, 3, .inherited = true},
  [428] =
    {field_decorator, 0, .inherited = true},
    {field_name, 1},
    {field_type, 3},
  [431] =
    {field_decorator, 0, .inherited = true},
    {field_name, 1},
    {field_type, 2},
    {field_value, 3, .inherited = true},
  [435] =
    {field_parameters, 1},
    {field_type, 2},
  [437] =
    {field_parameters, 2},
    {field_type_parameters, 1},
  [439] =
    {field_parameters, 2},
  [440] =
    {field_index, 0, .inherited = true},
    {field_object, 0},
    {field_object, 0, .inherited = true},
    {field_property, 2},
  [444] =
    {field_parameters, 1},
    {field_type, 3},
  [446] =
    {field_parameters, 1},
    {field_return_type, 3},
    {field_type_parameters, 0},
  [449] =
    {field_body, 5},
    {field_name, 3},
    {field_parameters, 4, .inherited = true},
    {field_return_type, 4, .inherited = true},
    {field_type_parameters, 4, .inherited = true},
  [454] =
    {field_body, 5},
    {field_name, 2},
    {field_type_parameters, 3},
  [457] =
    {field_type, 0, .inherited = true},
    {field_type, 1, .inherited = true},
  [459] =
    {field_name, 1},
    {field_name, 2, .inherited = true},
  [461] =
    {field_name, 0, .inherited = true},
    {field_name, 1, .inherited = true},
  [463] =
    {field_name, 2, .inherited = true},
  [464] =
    {field_body, 5},
    {field_decorator, 0, .inherited = true},
    {field_name, 2},
    {field_type_parameters, 3},
  [468] =
    {field_body, 5},
    {field_decorator, 0, .inherited = true},
    {field_name, 3},
  [471] =
    {field_body, 5},
    {field_decorator, 0, .inherited = true},
    {field_name, 3},
    {field_type_parameters, 4},
  [475] =
    {field_body, 5},
    {field_name, 2},
    {field_parameters, 4, .inherited = true},
    {field_return_type, 4, .inherited = true},
    {field_type_parameters, 4, .inherited = true},
  [480] =
    {field_source, 4},
  [481] =
    {field_body, 3},
    {field_value, 1},
  [483] =
    {field_kind, 1},
    {field_left, 2},
    {field_operator, 3},
    {field_right, 4},
  [487] =
    {field_body, 6},
    {field_condition, 3},
    {field_increment, 4},
    {field_initializer, 2},
  [491] =
    {field_body, 6},
    {field_condition, 3},
    {field_condition, 4},
    {field_initializer, 2},
  [495] =
    {field_body, 6},
    {field_condition, 4},
    {field_initializer, 2},
  [498] =
    {field_body, 4},
    {field_parameter, 2},
  [500] =
    {field_name, 2},
    {field_parameters, 4, .inherited = true},
    {field_return_type, 4, .inherited = true},
    {field_type_parameters, 4, .inherited = true},
  [504] =
    {field_name, 3},
    {field_parameters, 4, .inherited = true},
    {field_return_type, 4, .inherited = true},
    {field_type_parameters, 4, .inherited = true},
  [508] =
    {field_name, 2},
    {field_value, 4, .inherited = true},
  [510] =
    {field_name, 2},
    {field_type, 4},
  [512] =
    {field_name, 2},
    {field_type, 3},
    {field_value, 4, .inherited = true},
  [515] =
    {field_name, 1},
    {field_type, 3},
    {field_value, 4, .inherited = true},
  [518] =
    {field_name, 3},
    {field_value, 4, .inherited = true},
  [520] =
    {field_name, 3},
    {field_type, 4},
  [522] =
    {field_type, 4},
  [523] =
    {field_name, 4},
  [524] =
    {field_decorator, 0, .inherited = true},
    {field_name, 3},
    {field_value, 4, .inherited = true},
  [527] =
    {field_decorator, 0, .inherited = true},
    {field_name, 3},
    {field_type, 4},
  [530] =
    {field_decorator, 0, .inherited = true},
    {field_name, 2},
    {field_value, 4, .inherited = true},
  [533] =
    {field_decorator, 0, .inherited = true},
    {field_name, 2},
    {field_type, 4},
  [536] =
    {field_decorator, 0, .inherited = true},
    {field_name, 2},
    {field_type, 3},
    {field_value, 4, .inherited = true},
  [540] =
    {field_decorator, 0, .inherited = true},
    {field_name, 4},
  [542] =
    {field_decorator, 0, .inherited = true},
    {field_name, 1},
    {field_type, 3},
    {field_value, 4, .inherited = true},
  [546] =
    {field_parameters, 2},
    {field_type, 3},
    {field_type_parameters, 1},
  [549] =
    {field_parameters, 2},
    {field_type, 3},
  [551] =
    {field_parameters, 3},
    {field_type_parameters, 2},
  [553] =
    {field_index, 2},
    {field_object, 0},
    {field_object, 0, .inherited = true},
    {field_property, 0, .inherited = true},
  [557] =
    {field_index, 0, .inherited = true},
    {field_index, 2},
    {field_object, 0},
    {field_object, 0, .inherited = true},
  [561] =
    {field_arguments, 0, .inherited = true},
    {field_function, 0, .inherited = true},
    {field_index, 2},
    {field_object, 0},
  [565] =
    {field_parameters, 2},
    {field_type, 4},
    {field_type_parameters, 1},
  [568] =
    {field_parameters, 2},
    {field_type, 4},
  [570] =
    {field_body, 6},
    {field_decorator, 0, .inherited = true},
    {field_name, 3},
    {field_type_parameters, 4},
  [574] =
    {field_body, 6},
    {field_name, 3},
    {field_parameters, 5, .inherited = true},
    {field_return_type, 5, .inherited = true},
    {field_type_parameters, 5, .inherited = true},
  [579] =
    {field_body, 6},
    {field_name, 4},
    {field_parameters, 5, .inherited = true},
    {field_return_type, 5, .inherited = true},
    {field_type_parameters, 5, .inherited = true},
  [584] =
    {field_kind, 1},
    {field_left, 2},
    {field_operator, 4},
    {field_right, 5},
    {field_value, 3, .inherited = true},
  [589] =
    {field_kind, 1},
    {field_left, 2},
    {field_operator, 4},
    {field_right, 5},
  [593] =
    {field_body, 7},
    {field_condition, 3},
    {field_condition, 4},
    {field_increment, 5},
    {field_initializer, 2},
  [598] =
    {field_body, 7},
    {field_condition, 4},
    {field_increment, 5},
    {field_initializer, 2},
  [602] =
    {field_body, 7},
    {field_condition, 4},
    {field_condition, 5},
    {field_initializer, 2},
  [606] =
    {field_body, 5},
    {field_parameter, 2},
    {field_type, 3},
  [609] =
    {field_index_type, 3},
    {field_name, 1},
    {field_type, 5},
  [612] =
    {field_alias, 4},
    {field_name, 0},
    {field_type, 2},
  [615] =
    {field_sign, 0},
    {field_type, 5},
  [617] =
    {field_name, 3},
    {field_parameters, 5, .inherited = true},
    {field_return_type, 5, .inherited = true},
    {field_type_parameters, 5, .inherited = true},
  [621] =
    {field_name, 4},
    {field_parameters, 5, .inherited = true},
    {field_return_type, 5, .inherited = true},
    {field_type_parameters, 5, .inherited = true},
  [625] =
    {field_name, 2},
    {field_type, 4},
    {field_value, 5, .inherited = true},
  [628] =
    {field_name, 3},
    {field_value, 5, .inherited = true},
  [630] =
    {field_name, 3},
    {field_type, 5},
  [632] =
    {field_name, 3},
    {field_type, 4},
    {field_value, 5, .inherited = true},
  [635] =
    {field_name, 4},
    {field_value, 5, .inherited = true},
  [637] =
    {field_name, 4},
    {field_type, 5},
  [639] =
    {field_name, 5},
  [640] =
    {field_decorator, 0, .inherited = true},
    {field_name, 3},
    {field_value, 5, .inherited = true},
  [643] =
    {field_decorator, 0, .inherited = true},
    {field_name, 3},
    {field_type, 5},
  [646] =
    {field_decorator, 0, .inherited = true},
    {field_name, 3},
    {field_type, 4},
    {field_value, 5, .inherited = true},
  [650] =
    {field_decorator, 0, .inherited = true},
    {field_name, 2},
    {field_type, 4},
    {field_value, 5, .inherited = true},
  [654] =
    {field_decorator, 0, .inherited = true},
    {field_name, 4},
    {field_value, 5, .inherited = true},
  [657] =
    {field_decorator, 0, .inherited = true},
    {field_name, 4},
    {field_type, 5},
  [660] =
    {field_decorator, 0, .inherited = true},
    {field_name, 5},
  [662] =
    {field_parameters, 3},
    {field_type, 4},
    {field_type_parameters, 2},
  [665] =
    {field_index, 3},
    {field_object, 0},
  [667] =
    {field_index, 3},
    {field_object, 0},
    {field_object, 0, .inherited = true},
    {field_property, 0, .inherited = true},
  [671] =
    {field_index, 0, .inherited = true},
    {field_index, 3},
    {field_object, 0},
    {field_object, 0, .inherited = true},
  [675] =
    {field_arguments, 0, .inherited = true},
    {field_function, 0, .inherited = true},
    {field_index, 3},
    {field_object, 0},
  [679] =
    {field_parameters, 3},
    {field_type, 5},
    {field_type_parameters, 2},
  [682] =
    {field_body, 7},
    {field_name, 4},
    {field_parameters, 6, .inherited = true},
    {field_return_type, 6, .inherited = true},
    {field_type_parameters, 6, .inherited = true},
  [687] =
    {field_body, 7},
    {field_name, 5},
    {field_parameters, 6, .inherited = true},
    {field_return_type, 6, .inherited = true},
    {field_type_parameters, 6, .inherited = true},
  [692] =
    {field_body, 8},
    {field_condition, 4},
    {field_condition, 5},
    {field_increment, 6},
    {field_initializer, 2},
  [697] =
    {field_name, 4},
    {field_parameters, 6, .inherited = true},
    {field_return_type, 6, .inherited = true},
    {field_type_parameters, 6, .inherited = true},
  [701] =
    {field_name, 5},
    {field_parameters, 6, .inherited = true},
    {field_return_type, 6, .inherited = true},
    {field_type_parameters, 6, .inherited = true},
  [705] =
    {field_name, 3},
    {field_type, 5},
    {field_value, 6, .inherited = true},
  [708] =
    {field_index_type, 4},
    {field_name, 2},
    {field_type, 6},
  [711] =
    {field_name, 4},
    {field_value, 6, .inherited = true},
  [713] =
    {field_name, 4},
    {field_type, 6},
  [715] =
    {field_name, 4},
    {field_type, 5},
    {field_value, 6, .inherited = true},
  [718] =
    {field_name, 5},
    {field_value, 6, .inherited = true},
  [720] =
    {field_name, 5},
    {field_type, 6},
  [722] =
    {field_decorator, 0, .inherited = true},
    {field_name, 3},
    {field_type, 5},
    {field_value, 6, .inherited = true},
  [726] =
    {field_decorator, 0, .inherited = true},
    {field_name, 4},
    {field_value, 6, .inherited = true},
  [729] =
    {field_decorator, 0, .inherited = true},
    {field_name, 4},
    {field_type, 6},
  [732] =
    {field_decorator, 0, .inherited = true},
    {field_name, 4},
    {field_type, 5},
    {field_value, 6, .inherited = true},
  [736] =
    {field_decorator, 0, .inherited = true},
    {field_name, 5},
    {field_value, 6, .inherited = true},
  [739] =
    {field_decorator, 0, .inherited = true},
    {field_name, 5},
    {field_type, 6},
  [742] =
    {field_decorator, 0, .inherited = true},
    {field_name, 6},
  [744] =
    {field_alternative, 6},
    {field_consequence, 4},
    {field_left, 0},
    {field_right, 2},
  [748] =
    {field_body, 8},
    {field_name, 5},
    {field_parameters, 7, .inherited = true},
    {field_return_type, 7, .inherited = true},
    {field_type_parameters, 7, .inherited = true},
  [753] =
    {field_body, 8},
    {field_name, 6},
    {field_parameters, 7, .inherited = true},
    {field_return_type, 7, .inherited = true},
    {field_type_parameters, 7, .inherited = true},
  [758] =
    {field_index_type, 5},
    {field_name, 3},
    {field_sign, 0},
    {field_type, 7},
  [762] =
    {field_name, 5},
    {field_parameters, 7, .inherited = true},
    {field_return_type, 7, .inherited = true},
    {field_type_parameters, 7, .inherited = true},
  [766] =
    {field_name, 4},
    {field_type, 6},
    {field_value, 7, .inherited = true},
  [769] =
    {field_name, 5},
    {field_value, 7, .inherited = true},
  [771] =
    {field_name, 5},
    {field_type, 7},
  [773] =
    {field_name, 5},
    {field_type, 6},
    {field_value, 7, .inherited = true},
  [776] =
    {field_name, 6},
    {field_parameters, 7, .inherited = true},
    {field_return_type, 7, .inherited = true},
    {field_type_parameters, 7, .inherited = true},
  [780] =
    {field_decorator, 0, .inherited = true},
    {field_name, 4},
    {field_type, 6},
    {field_value, 7, .inherited = true},
  [784] =
    {field_decorator, 0, .inherited = true},
    {field_name, 5},
    {field_value, 7, .inherited = true},
  [787] =
    {field_decorator, 0, .inherited = true},
    {field_name, 5},
    {field_type, 7},
  [790] =
    {field_decorator, 0, .inherited = true},
    {field_name, 5},
    {field_type, 6},
    {field_value, 7, .inherited = true},
  [794] =
    {field_decorator, 0, .inherited = true},
    {field_name, 6},
    {field_value, 7, .inherited = true},
  [797] =
    {field_decorator, 0, .inherited = true},
    {field_name, 6},
    {field_type, 7},
  [800] =
    {field_body, 9},
    {field_name, 6},
    {field_parameters, 8, .inherited = true},
    {field_return_type, 8, .inherited = true},
    {field_type_parameters, 8, .inherited = true},
  [805] =
    {field_name, 5},
    {field_type, 7},
    {field_value, 8, .inherited = true},
  [808] =
    {field_name, 6},
    {field_parameters, 8, .inherited = true},
    {field_return_type, 8, .inherited = true},
    {field_type_parameters, 8, .inherited = true},
  [812] =
    {field_decorator, 0, .inherited = true},
    {field_name, 5},
    {field_type, 7},
    {field_value, 8, .inherited = true},
  [816] =
    {field_decorator, 0, .inherited = true},
    {field_name, 6},
    {field_value, 8, .inherited = true},
  [819] =
    {field_decorator, 0, .inherited = true},
    {field_name, 6},
    {field_type, 8},
  [822] =
    {field_decorator, 0, .inherited = true},
    {field_name, 6},
    {field_type, 7},
    {field_value, 8, .inherited = true},
  [826] =
    {field_decorator, 0, .inherited = true},
    {field_name, 6},
    {field_type, 8},
    {field_value, 9, .inherited = true},
};

static const TSSymbol ts_alias_sequences[PRODUCTION_ID_COUNT][MAX_ALIAS_SEQUENCE_LENGTH] = {
  [0] = {0},
  [1] = {
    [0] = sym_identifier,
  },
  [7] = {
    [0] = alias_sym_property_identifier,
  },
  [13] = {
    [0] = alias_sym_type_identifier,
  },
  [22] = {
    [0] = sym_identifier,
  },
  [23] = {
    [0] = alias_sym_statement_identifier,
  },
  [24] = {
    [0] = sym_identifier,
  },
  [26] = {
    [1] = alias_sym_shorthand_property_identifier,
  },
  [27] = {
    [1] = alias_sym_shorthand_property_identifier_pattern,
  },
  [28] = {
    [1] = sym_identifier,
  },
  [44] = {
    [1] = alias_sym_statement_identifier,
  },
  [45] = {
    [1] = alias_sym_type_identifier,
  },
  [48] = {
    [0] = alias_sym_type_identifier,
  },
  [49] = {
    [0] = alias_sym_this_type,
  },
  [58] = {
    [1] = alias_sym_type_identifier,
  },
  [61] = {
    [0] = alias_sym_type_identifier,
  },
  [62] = {
    [0] = alias_sym_type_identifier,
  },
  [63] = {
    [1] = sym_member_expression,
  },
  [67] = {
    [1] = alias_sym_type_identifier,
    [2] = alias_sym_interface_body,
  },
  [69] = {
    [2] = alias_sym_property_identifier,
  },
  [71] = {
    [2] = alias_sym_property_identifier,
  },
  [82] = {
    [0] = sym_member_expression,
    [2] = alias_sym_property_identifier,
  },
  [83] = {
    [0] = alias_sym_shorthand_property_identifier_pattern,
  },
  [105] = {
    [1] = alias_sym_type_identifier,
  },
  [106] = {
    [1] = alias_sym_type_identifier,
  },
  [114] = {
    [0] = alias_sym_type_identifier,
  },
  [115] = {
    [1] = alias_sym_type_identifier,
  },
  [116] = {
    [1] = anon_sym_unique,
  },
  [120] = {
    [1] = sym_identifier,
  },
  [127] = {
    [1] = sym_identifier,
  },
  [130] = {
    [1] = alias_sym_type_identifier,
  },
  [131] = {
    [1] = alias_sym_type_identifier,
  },
  [135] = {
    [0] = alias_sym_type_identifier,
  },
  [136] = {
    [1] = sym_member_expression,
  },
  [137] = {
    [1] = sym_member_expression,
  },
  [138] = {
    [2] = alias_sym_type_identifier,
  },
  [139] = {
    [1] = alias_sym_type_identifier,
  },
  [141] = {
    [1] = alias_sym_type_identifier,
    [3] = alias_sym_interface_body,
  },
  [142] = {
    [1] = alias_sym_type_identifier,
    [3] = alias_sym_interface_body,
  },
  [146] = {
    [2] = alias_sym_type_identifier,
  },
  [151] = {
    [1] = alias_sym_type_identifier,
  },
  [153] = {
    [0] = sym_identifier,
  },
  [165] = {
    [1] = alias_sym_type_identifier,
  },
  [172] = {
    [2] = alias_sym_type_identifier,
  },
  [174] = {
    [2] = alias_sym_property_identifier,
  },
  [176] = {
    [2] = alias_sym_property_identifier,
  },
  [181] = {
    [1] = alias_sym_type_identifier,
  },
  [183] = {
    [1] = sym_member_expression,
  },
  [184] = {
    [2] = alias_sym_type_identifier,
  },
  [185] = {
    [2] = alias_sym_type_identifier,
  },
  [186] = {
    [1] = alias_sym_type_identifier,
  },
  [188] = {
    [1] = alias_sym_type_identifier,
    [4] = alias_sym_interface_body,
  },
  [191] = {
    [0] = sym_identifier,
  },
  [193] = {
    [2] = alias_sym_type_identifier,
  },
  [194] = {
    [2] = alias_sym_type_identifier,
  },
  [196] = {
    [3] = alias_sym_type_identifier,
  },
  [198] = {
    [1] = alias_sym_type_identifier,
  },
  [200] = {
    [1] = sym_identifier,
  },
  [201] = {
    [1] = sym_identifier,
  },
  [206] = {
    [0] = alias_sym_type_identifier,
  },
  [224] = {
    [2] = alias_sym_property_identifier,
  },
  [229] = {
    [2] = alias_sym_type_identifier,
  },
  [234] = {
    [2] = alias_sym_type_identifier,
  },
  [235] = {
    [3] = alias_sym_type_identifier,
  },
  [236] = {
    [3] = alias_sym_type_identifier,
  },
  [270] = {
    [3] = alias_sym_property_identifier,
  },
  [271] = {
    [3] = alias_sym_type_identifier,
  },
  [280] = {
    [1] = sym_identifier,
  },
  [281] = {
    [0] = alias_sym_type_identifier,
  },
  [312] = {
    [2] = sym_identifier,
  },
  [329] = {
    [3] = sym_identifier,
  },
};

static const uint16_t ts_non_terminal_alias_map[] = {
  sym_nested_identifier, 2,
    sym_nested_identifier,
    sym_member_expression,
  sym_predefined_type, 2,
    sym_predefined_type,
    sym_identifier,
  sym_object_type, 2,
    sym_object_type,
    alias_sym_interface_body,
  0,
};

// Filled from tables.zlib by grammar.go, see fetchgrammar.
static TSStateId ts_primary_state_ids[STATE_COUNT];

static TSCharacterRange extras_character_set_1[] = {
  {'\t', '\r'}, {' ', ' '}, {0xa0, 0xa0}, {0x1680, 0x1680}, {0x2000, 0x200b}, {0x2028, 0x2029}, {0x202f, 0x202f}, {0x205f, 0x2060},
  {0x3000, 0x3000}, {0xfeff, 0xfeff},
};

static TSCharacterRange sym_identifier_character_set_1[] = {
  {'$', '$'}, {'A', 'Z'}, {'\\', '\\'}, {'_', '_'}, {'a', 'z'}, {0x7f, 0x9f}, {0xa1, 0x167f}, {0x1681, 0x1fff},
  {0x200c, 0x2027}, {0x202a, 0x202e}, {0x2030, 0x205e}, {0x2061, 0x2fff}, {0x3001, 0xfefe}, {0xff00, 0x10ffff},
};

static TSCharacterRange sym_identifier_character_set_2[] = {
  {'$', '$'}, {'0', '9'}, {'A', 'Z'}, {'\\', '\\'}, {'_', '_'}, {'a', 'z'}, {0x7f, 0x9f}, {0xa1, 0x167f},
  {0x1681, 0x1fff}, {0x200c, 0x2027}, {0x202a, 0x202e}, {0x2030, 0x205e}, {0x2061, 0x2fff}, {0x3001, 0xfefe}, {0xff00, 0x10ffff},
};

static bool ts_lex(TSLexer *lexer, TSStateId state) {
  START_LEXER();
  eof = lexer->eof(lexer);
  switch (state) {
    case 0:
      if (eof) ADVANCE(129);
      ADVANCE_MAP(
        '!', 143,
        '"', 160,
        '#', 6,
        '$', 281,
        '%', 224,
        '&', 203,
        '\'', 161,
        '(', 144,
        ')', 146,
        '*', 132,
        '+', 216,
        ',', 140,
        '-', 220,
        '.', 156,
        '/', 266,
        '0', 271,
        ':', 147,
        ';', 145,
        '<', 229,
        '=', 136,
        '>', 152,
        '?', 287,
        '@', 284,
        '[', 148,
        '\\', 82,
        ']', 149,
        '^', 206,
        '`', 264,
        '{', 139,
        '|', 209,
        '}', 141,
        '~', 240,
      );
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(272);
      if (set_contains(extras_character_set_1, 10, lookahead)) SKIP(126);
      if (lookahead > '@') ADVANCE(282);
      END_STATE();
    case 1:
      if (lookahead == '\n') SKIP(34);
      if (lookahead == '/') ADVANCE(24);
      if (lookahead == '[') ADVANCE(81);
      if (lookahead == '\\') ADVANCE(125);
      if (set_contains(extras_character_set_1, 10, lookahead)) ADVANCE(267);
      if (lookahead != 0) ADVANCE(268);
      END_STATE();
    case 2:
      ADVANCE_MAP(
        '!', 143,
        '"', 160,
        '#', 80,
        '%', 224,
        '&', 203,
        '\'', 161,
        '(', 144,
        ')', 146,
        '*', 132,
        '+', 215,
        ',', 140,
        '-', 219,
        '.', 156,
        '/', 222,
        '0', 271,
        ':', 147,
        ';', 145,
        '<', 230,
        '=', 136,
        '>', 152,
        '?', 287,
        '@', 284,
        '[', 148,
        '\\', 84,
        ']', 149,
        '^', 206,
        '`', 264,
        '{', 139,
        '|', 209,
        '}', 141,
        '~', 240,
      );
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(272);
      if (set_contains(extras_character_set_1, 10, lookahead)) SKIP(2);
      if (lookahead > '#') ADVANCE(282);
      END_STATE();
    case 3:
      ADVANCE_MAP(
        '!', 143,
        '"', 160,
        '#', 80,
        '%', 224,
        '&', 203,
        '\'', 161,
        '(', 144,
        ')', 146,
        '*', 132,
        '+', 215,
        ',', 140,
        '-', 219,
        '.', 156,
        '/', 222,
        '0', 271,
        ':', 147,
        ';', 145,
        '<', 230,
        '=', 136,
        '>', 152,
        '?', 287,
        '@', 284,
        '[', 148,
        '\\', 84,
        ']', 149,
        '^', 206,
        '`', 264,
        '{', 138,
        '|', 208,
        '}', 141,
        '~', 240,
      );
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(272);
      if (set_contains(extras_character_set_1, 10, lookahead)) SKIP(3);
      if (lookahead > '#') ADVANCE(282);
      END_STATE();
    case 4:
      ADVANCE_MAP(
        '!', 143,
        '%', 223,
        '&', 204,
        '(', 144,
        ')', 146,
        '*', 133,
        '+', 214,
        ',', 140,
        '-', 218,
        '.', 155,
        '/', 221,
        ':', 147,
        ';', 145,
        '<', 231,
        '=', 77,
        '>', 153,
        '?', 31,
        '[', 148,
        '\\', 84,
        ']', 149,
        '^', 205,
        '`', 264,
        '{', 138,
        '|', 210,
        '}', 141,
      );
      if (('a' <= lookahead && lookahead <= 'z')) ADVANCE(269);
      if (set_contains(extras_character_set_1, 10, lookahead)) SKIP(5);
      if (lookahead > '#' &&
          (lookahead < '%' || '@' < lookahead) &&
          (lookahead < '`' || '~' < lookahead)) ADVANCE(282);
      END_STATE();
    case 5:
      ADVANCE_MAP(
        '!', 143,
        '%', 223,
        '&', 204,
        '(', 144,
        ')', 146,
        '*', 133,
        '+', 214,
        ',', 140,
        '-', 218,
        '.', 155,
        '/', 221,
        ':', 147,
        ';', 145,
        '<', 231,
        '=', 77,
        '>', 153,
        '?', 31,
        '[', 148,
        '\\', 84,
        ']', 149,
        '^', 205,
        '`', 264,
        '{', 138,
        '|', 210,
        '}', 141,
      );
      if (set_contains(extras_character_set_1, 10, lookahead)) SKIP(5);
      if (lookahead > '#' &&
          (lookahead < '%' || '@' < lookahead) &&
          (lookahead < '{' || '~' < lookahead)) ADVANCE(282);
      END_STATE();
    case 6:
      if (lookahead == '!') ADVANCE(130);
      if (lookahead == '\\') ADVANCE(83);
      if (set_contains(sym_identifier_character_set_1, 14, lookahead)) ADVANCE(283);
      END_STATE();
    case 7:
      ADVANCE_MAP(
        '!', 142,
        '"', 160,
        '#', 80,
        '&', 202,
        '\'', 161,
        '(', 144,
        ')', 146,
        '*', 131,
        '+', 214,
        ',', 140,
        '-', 218,
        '.', 156,
        '/', 221,
        '0', 271,
        ':', 147,
        ';', 145,
        '<', 227,
        '=', 137,
        '>', 151,
        '?', 286,
        '@', 284,
        '[', 148,
        '\\', 84,
        ']', 149,
        '`', 264,
        '{', 138,
        '|', 212,
        '}', 141,
        '~', 240,
      );
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(272);
      if (set_contains(extras_character_set_1, 10, lookahead)) SKIP(7);
      if (lookahead > '#' &&
          (lookahead < '%' || '@' < lookahead) &&
          (lookahead < '[' || '^' < lookahead)) ADVANCE(282);
      END_STATE();
    case 8:
      ADVANCE_MAP(
        '!', 142,
        '"', 160,
        '#', 80,
        '&', 202,
        '\'', 161,
        '(', 144,
        ')', 146,
        '*', 131,
        '+', 214,
        ',', 140,
        '-', 218,
        '.', 156,
        '/', 221,
        '0', 271,
        '<', 227,
        '?', 285,
        '@', 284,
        '[', 148,
        '\\', 84,
        ']', 149,
        '`', 264,
        '{', 139,
        '|', 207,
        '~', 240,
      );
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(272);
      if (set_contains(extras_character_set_1, 10, lookahead)) SKIP(8);
      if (lookahead > '#' &&
          (lookahead < '%' || '@' < lookahead) &&
          (lookahead < '[' || '^' < lookahead) &&
          (lookahead < '{' || '~' < lookahead)) ADVANCE(282);
      END_STATE();
    case 9:
      ADVANCE_MAP(
        '"', 160,
        '#', 80,
        '&', 202,
        '\'', 161,
        '(', 144,
        '*', 131,
        '+', 213,
        ',', 140,
        '-', 217,
        '.', 30,
        '/', 24,
        '0', 271,
        ';', 145,
        '<', 227,
        '>', 151,
        '?', 285,
        '@', 284,
        '[', 148,
        '\\', 84,
        ']', 149,
        '`', 264,
        '{', 139,
        '|', 212,
        '}', 141,
      );
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(272);
      if (set_contains(extras_character_set_1, 10, lookahead)) SKIP(9);
      if (lookahead > '#' &&
          (lookahead < '%' || '@' < lookahead) &&
          (lookahead < '[' || '^' < lookahead) &&
          (lookahead < '{' || '~' < lookahead)) ADVANCE(282);
      END_STATE();
    case 10:
      ADVANCE_MAP(
        '"', 160,
        '#', 80,
        '\'', 161,
        '(', 144,
        '*', 131,
        '+', 213,
        ',', 140,
        '-', 217,
        '.', 30,
        '/', 24,
        '0', 271,
        ';', 145,
        '<', 227,
        '@', 284,
        '[', 148,
        '\\', 84,
        '{', 138,
        '|', 94,
        '}', 141,
      );
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(272);
      if (set_contains(extras_character_set_1, 10, lookahead)) SKIP(10);
      if (lookahead > '#' &&
          (lookahead < '%' || '@' < lookahead) &&
          (lookahead < '[' || '^' < lookahead) &&
          lookahead != '`' &&
          (lookahead < '{' || '~' < lookahead)) ADVANCE(282);
      END_STATE();
    case 11:
      ADVANCE_MAP(
        '"', 160,
        '&', 202,
        '\'', 161,
        '(', 144,
        '*', 131,
        '+', 213,
        '-', 217,
        '.', 101,
        '/', 25,
        '0', 271,
        '<', 227,
        '>', 151,
        '?', 285,
        '[', 148,
        '\\', 84,
        '`', 264,
        '{', 139,
        '|', 207,
      );
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(272);
      if (set_contains(extras_character_set_1, 10, lookahead)) SKIP(11);
      if (lookahead == '$' ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          ('_' <= lookahead && lookahead <= 'z')) ADVANCE(280);
      if (lookahead > '~') ADVANCE(282);
      END_STATE();
    case 12:
      if (lookahead == '"') ADVANCE(160);
      if (lookahead == '&') ADVANCE(16);
      if (lookahead == '/') ADVANCE(163);
      if (set_contains(extras_character_set_1, 10, lookahead)) ADVANCE(162);
      if (lookahead != 0) ADVANCE(164);
      END_STATE();
    case 13:
      if (lookahead == '"') ADVANCE(160);
      if (lookahead == '/') ADVANCE(24);
      if (set_contains(extras_character_set_1, 10, lookahead)) SKIP(13);
      END_STATE();
    case 14:
      if (lookahead == '"') ADVANCE(160);
      if (lookahead == '/') ADVANCE(243);
      if (lookahead == '\\') ADVANCE(85);
      if (lookahead == '\n' ||
          lookahead == '\r') SKIP(13);
      if (set_contains(extras_character_set_1, 10, lookahead)) ADVANCE(246);
      if (lookahead != 0) ADVANCE(248);
      END_STATE();
    case 15:
      if (lookahead == '#') ADVANCE(97);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(76);
      END_STATE();
    case 16:
      if (lookahead == '#') ADVANCE(97);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(76);
      if (lookahead != 0) ADVANCE(164);
      END_STATE();
    case 17:
      if (lookahead == '#') ADVANCE(97);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(76);
      if (lookahead != 0) ADVANCE(170);
      END_STATE();
    case 18:
      if (lookahead == '$') ADVANCE(86);
      if (lookahead == '&') ADVANCE(15);
      if (lookahead == '/') ADVANCE(24);
      if (lookahead == '<') ADVANCE(228);
      if (lookahead == '\\') ADVANCE(85);
      if (lookahead == '`') ADVANCE(264);
      if (lookahead == '{') ADVANCE(138);
      if (set_contains(extras_character_set_1, 10, lookahead)) SKIP(19);
      END_STATE();
    case 19:
      if (lookahead == '$') ADVANCE(86);
      if (lookahead == '&') ADVANCE(15);
      if (lookahead == '/') ADVANCE(24);
      if (lookahead == '<') ADVANCE(228);
      if (lookahead == '`') ADVANCE(264);
      if (lookahead == '{') ADVANCE(138);
      if (set_contains(extras_character_set_1, 10, lookahead)) SKIP(19);
      END_STATE();
    case 20:
      ADVANCE_MAP(
        '&', 202,
        '(', 144,
        '+', 78,
        ',', 140,
        '-', 79,
        '.', 155,
        '/', 25,
        ':', 147,
        '<', 227,
        '=', 134,
        '>', 151,
        '?', 37,
        '[', 148,
        '\\', 84,
        '{', 138,
        '|', 207,
      );
      if (set_contains(extras_character_set_1, 10, lookahead)) SKIP(20);
      if (lookahead == '$' ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(280);
      if (lookahead > '~') ADVANCE(282);
      END_STATE();
    case 21:
      if (lookahead == '&') ADVANCE(17);
      if (lookahead == '\'') ADVANCE(161);
      if (lookahead == '/') ADVANCE(169);
      if (set_contains(extras_character_set_1, 10, lookahead)) ADVANCE(168);
      if (lookahead != 0) ADVANCE(170);
      END_STATE();
    case 22:
      if (lookahead == '\'') ADVANCE(161);
      if (lookahead == '/') ADVANCE(24);
      if (set_contains(extras_character_set_1, 10, lookahead)) SKIP(22);
      END_STATE();
    case 23:
      if (lookahead == '\'') ADVANCE(161);
      if (lookahead == '/') ADVANCE(249);
      if (lookahead == '\\') ADVANCE(85);
      if (lookahead == '\n' ||
          lookahead == '\r') SKIP(22);
      if (set_contains(extras_character_set_1, 10, lookahead)) ADVANCE(252);
      if (lookahead != 0) ADVANCE(254);
      END_STATE();
    case 24:
      if (lookahead == '*') ADVANCE(27);
      if (lookahead == '/') ADVANCE(263);
      END_STATE();
    case 25:
      if (lookahead == '*') ADVANCE(27);
      if (lookahead == '/') ADVANCE(263);
      if (lookahead == '>') ADVANCE(159);
      END_STATE();
    case 26:
      if (lookahead == '*') ADVANCE(26);
      if (lookahead == '/') ADVANCE(260);
      if (lookahead != 0) ADVANCE(27);
      END_STATE();
    case 27:
      if (lookahead == '*') ADVANCE(26);
      if (lookahead != 0) ADVANCE(27);
      END_STATE();
    case 28:
      if (lookahead == '*') ADVANCE(165);
      if (lookahead == '#' ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(27);
      if (lookahead != 0) ADVANCE(166);
      END_STATE();
    case 29:
      if (lookahead == '*') ADVANCE(171);
      if (lookahead == '#' ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(27);
      if (lookahead != 0) ADVANCE(172);
      END_STATE();
    case 30:
      if (lookahead == '.') ADVANCE(32);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(277);
      END_STATE();
    case 31:
      if (lookahead == '.') ADVANCE(175);
      if (lookahead == '?') ADVANCE(238);
      END_STATE();
    case 32:
      if (lookahead == '.') ADVANCE(191);
      END_STATE();
    case 33:
      if (lookahead == '/') ADVANCE(266);
      if (set_contains(extras_character_set_1, 10, lookahead)) SKIP(34);
      END_STATE();
    case 34:
      if (lookahead == '/') ADVANCE(24);
      if (set_contains(extras_character_set_1, 10, lookahead)) SKIP(34);
      END_STATE();
    case 35:
      if (lookahead == ':') ADVANCE(290);
      END_STATE();
    case 36:
      if (lookahead == ':') ADVANCE(289);
      END_STATE();
    case 37:
      if (lookahead == ':') ADVANCE(291);
      END_STATE();
    case 38:
      if (lookahead == ';') ADVANCE(150);
      END_STATE();
    case 39:
      if (lookahead == ';') ADVANCE(150);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(38);
      END_STATE();
    case 40:
      if (lookahead == ';') ADVANCE(150);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(39);
      END_STATE();
    case 41:
      if (lookahead == ';') ADVANCE(150);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(40);
      END_STATE();
    case 42:
      if (lookahead == ';') ADVANCE(150);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(41);
      END_STATE();
    case 43:
      if (lookahead == ';') ADVANCE(150);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(38);
      END_STATE();
    case 44:
      if (lookahead == ';') ADVANCE(150);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(43);
      END_STATE();
    case 45:
      if (lookahead == ';') ADVANCE(150);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(44);
      END_STATE();
    case 46:
      if (lookahead == ';') ADVANCE(150);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(45);
      END_STATE();
    case 47:
      if (lookahead == ';') ADVANCE(150);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(46);
      END_STATE();
    case 48:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(38);
      END_STATE();
    case 49:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(48);
      END_STATE();
    case 50:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(49);
      END_STATE();
    case 51:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(50);
      END_STATE();
    case 52:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(51);
      END_STATE();
    case 53:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(52);
      END_STATE();
    case 54:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(53);
      END_STATE();
    case 55:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(54);
      END_STATE();
    case 56:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(55);
      END_STATE();
    case 57:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(56);
      END_STATE();
    case 58:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(57);
      END_STATE();
    case 59:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(58);
      END_STATE();
    case 60:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(59);
      END_STATE();
    case 61:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(60);
      END_STATE();
    case 62:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(61);
      END_STATE();
    case 63:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(62);
      END_STATE();
    case 64:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(63);
      END_STATE();
    case 65:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(64);
      END_STATE();
    case 66:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(65);
      END_STATE();
    case 67:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(66);
      END_STATE();
    case 68:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(67);
      END_STATE();
    case 69:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(68);
      END_STATE();
    case 70:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(69);
      END_STATE();
    case 71:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(70);
      END_STATE();
    case 72:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(71);
      END_STATE();
    case 73:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(72);
      END_STATE();
    case 74:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(73);
      END_STATE();
    case 75:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(74);
      END_STATE();
    case 76:
      if (lookahead == ';') ADVANCE(150);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(75);
      END_STATE();
    case 77:
      if (lookahead == '=') ADVANCE(233);
      END_STATE();
    case 78:
      if (lookahead == '?') ADVANCE(35);
      END_STATE();
    case 79:
      if (lookahead == '?') ADVANCE(36);
      END_STATE();
    case 80:
      if (lookahead == '\\') ADVANCE(83);
      if (set_contains(sym_identifier_character_set_1, 14, lookahead)) ADVANCE(283);
      END_STATE();
    case 81:
      if (lookahead == '\\') ADVANCE(124);
      if (lookahead == ']') ADVANCE(268);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(81);
      END_STATE();
    case 82:
      if (lookahead == 'u') ADVANCE(87);
      if (lookahead == 'x') ADVANCE(114);
      if (lookahead == '\r' ||
          lookahead == '?') ADVANCE(257);
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(259);
      if (lookahead != 0) ADVANCE(255);
      END_STATE();
    case 83:
      if (lookahead == 'u') ADVANCE(88);
      END_STATE();
    case 84:
      if (lookahead == 'u') ADVANCE(89);
      END_STATE();
    case 85:
      if (lookahead == 'u') ADVANCE(90);
      if (lookahead == 'x') ADVANCE(114);
      if (lookahead == '\r' ||
          lookahead == '?') ADVANCE(257);
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(259);
      if (lookahead != 0) ADVANCE(255);
      END_STATE();
    case 86:
      if (lookahead == '{') ADVANCE(265);
      END_STATE();
    case 87:
      if (lookahead == '{') ADVANCE(108);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(119);
      END_STATE();
    case 88:
      if (lookahead == '{') ADVANCE(112);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(120);
      END_STATE();
    case 89:
      if (lookahead == '{') ADVANCE(113);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(121);
      END_STATE();
    case 90:
      if (lookahead == '{') ADVANCE(115);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(111);
      END_STATE();
    case 91:
      if (lookahead == '}') ADVANCE(282);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(91);
      END_STATE();
    case 92:
      if (lookahead == '}') ADVANCE(283);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(92);
      END_STATE();
    case 93:
      if (lookahead == '}') ADVANCE(255);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(93);
      END_STATE();
    case 94:
      if (lookahead == '}') ADVANCE(293);
      END_STATE();
    case 95:
      if (lookahead == '}') ADVANCE(256);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(95);
      END_STATE();
    case 96:
      if (lookahead == '+' ||
          lookahead == '-') ADVANCE(103);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(278);
      END_STATE();
    case 97:
      if (lookahead == 'X' ||
          lookahead == 'x') ADVANCE(110);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(42);
      END_STATE();
    case 98:
      if (lookahead == '0' ||
          lookahead == '1') ADVANCE(274);
      END_STATE();
    case 99:
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(275);
      END_STATE();
    case 100:
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(272);
      END_STATE();
    case 101:
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(277);
      END_STATE();
    case 102:
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(273);
      END_STATE();
    case 103:
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(278);
      END_STATE();
    case 104:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(282);
      END_STATE();
    case 105:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(283);
      END_STATE();
    case 106:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(255);
      END_STATE();
    case 107:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(276);
      END_STATE();
    case 108:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(95);
      END_STATE();
    case 109:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(256);
      END_STATE();
    case 110:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(47);
      END_STATE();
    case 111:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(114);
      END_STATE();
    case 112:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(92);
      END_STATE();
    case 113:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(91);
      END_STATE();
    case 114:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(106);
      END_STATE();
    case 115:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(93);
      END_STATE();
    case 116:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(109);
      END_STATE();
    case 117:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(105);
      END_STATE();
    case 118:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(104);
      END_STATE();
    case 119:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(116);
      END_STATE();
    case 120:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(117);
      END_STATE();
    case 121:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(118);
      END_STATE();
    case 122:
      if (lookahead != 0 &&
          lookahead != '#' &&
          (lookahead < 'A' || 'Z' < lookahead) &&
          (lookahead < 'a' || 'z' < lookahead)) ADVANCE(164);
      END_STATE();
    case 123:
      if (lookahead != 0 &&
          lookahead != '#' &&
          (lookahead < 'A' || 'Z' < lookahead) &&
          (lookahead < 'a' || 'z' < lookahead)) ADVANCE(170);
      END_STATE();
    case 124:
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(81);
      END_STATE();
    case 125:
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(268);
      END_STATE();
    case 126:
      if (eof) ADVANCE(129);
      ADVANCE_MAP(
        '!', 143,
        '"', 160,
        '#', 6,
        '$', 281,
        '%', 224,
        '&', 203,
        '\'', 161,
        '(', 144,
        ')', 146,
        '*', 132,
        '+', 216,
        ',', 140,
        '-', 220,
        '.', 156,
        '/', 222,
        '0', 271,
        ':', 147,
        ';', 145,
        '<', 229,
        '=', 136,
        '>', 152,
        '?', 287,
        '@', 284,
        '[', 148,
        '\\', 84,
        ']', 149,
        '^', 206,
        '`', 264,
        '{', 139,
        '|', 209,
        '}', 141,
        '~', 240,
      );
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(272);
      if (set_contains(extras_character_set_1, 10, lookahead)) SKIP(126);
      if (lookahead > '@') ADVANCE(282);
      END_STATE();
    case 127:
      if (eof) ADVANCE(129);
      ADVANCE_MAP(
        '!', 143,
        '"', 160,
        '#', 80,
        '%', 223,
        '&', 204,
        '\'', 161,
        '(', 144,
        ')', 146,
        '*', 133,
        '+', 214,
        ',', 140,
        '-', 218,
        '.', 157,
        '/', 221,
        '0', 271,
        ':', 147,
        ';', 145,
        '<', 231,
        '=', 135,
        '>', 153,
        '?', 288,
        '@', 284,
        '[', 148,
        '\\', 84,
        ']', 149,
        '^', 205,
        '`', 264,
        '{', 138,
        '|', 211,
        '}', 141,
        '~', 240,
      );
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(272);
      if (set_contains(extras_character_set_1, 10, lookahead)) SKIP(127);
      if (lookahead > '#') ADVANCE(282);
      END_STATE();
    case 128:
      if (eof) ADVANCE(129);
      ADVANCE_MAP(
        '!', 142,
        '"', 160,
        '#', 6,
        '&', 202,
        '\'', 161,
        '(', 144,
        ')', 146,
        '*', 131,
        '+', 214,
        ',', 140,
        '-', 218,
        '.', 156,
        '/', 221,
        '0', 271,
        ':', 147,
        ';', 145,
        '<', 227,
        '=', 137,
        '>', 151,
        '?', 285,
        '@', 284,
        '[', 148,
        '\\', 84,
        ']', 149,
        '`', 264,
        '{', 138,
        '|', 212,
        '}', 141,
        '~', 240,
      );
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(272);
      if (set_contains(extras_character_set_1, 10, lookahead)) SKIP(128);
      if (lookahead > '#' &&
          (lookahead < '%' || '@' < lookahead) &&
          (lookahead < '[' || '^' < lookahead)) ADVANCE(282);
      END_STATE();
    case 129:
      ACCEPT_TOKEN(ts_builtin_sym_end);
      END_STATE();
    case 130:
      ACCEPT_TOKEN(sym_hash_bang_line);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(130);
      END_STATE();
    case 131:
      ACCEPT_TOKEN(anon_sym_STAR);
      END_STATE();
    case 132:
      ACCEPT_TOKEN(anon_sym_STAR);
      if (lookahead == '*') ADVANCE(226);
      if (lookahead == '=') ADVANCE(178);
      END_STATE();
    case 133:
      ACCEPT_TOKEN(anon_sym_STAR);
      if (lookahead == '*') ADVANCE(225);
      END_STATE();
    case 134:
      ACCEPT_TOKEN(anon_sym_EQ);
      END_STATE();
    case 135:
      ACCEPT_TOKEN(anon_sym_EQ);
      if (lookahead == '=') ADVANCE(233);
      END_STATE();
    case 136:
      ACCEPT_TOKEN(anon_sym_EQ);
      if (lookahead == '=') ADVANCE(233);
      if (lookahead == '>') ADVANCE(174);
      END_STATE();
    case 137:
      ACCEPT_TOKEN(anon_sym_EQ);
      if (lookahead == '>') ADVANCE(174);
      END_STATE();
    case 138:
      ACCEPT_TOKEN(anon_sym_LBRACE);
      END_STATE();
    case 139:
      ACCEPT_TOKEN(anon_sym_LBRACE);
      if (lookahead == '|') ADVANCE(292);
      END_STATE();
    case 140:
      ACCEPT_TOKEN(anon_sym_COMMA);
      END_STATE();
    case 141:
      ACCEPT_TOKEN(anon_sym_RBRACE);
      END_STATE();
    case 142:
      ACCEPT_TOKEN(anon_sym_BANG);
      END_STATE();
    case 143:
      ACCEPT_TOKEN(anon_sym_BANG);
      if (lookahead == '=') ADVANCE(235);
      END_STATE();
    case 144:
      ACCEPT_TOKEN(anon_sym_LPAREN);
      END_STATE();
    case 145:
      ACCEPT_TOKEN(anon_sym_SEMI);
      END_STATE();
    case 146:
      ACCEPT_TOKEN(anon_sym_RPAREN);
      END_STATE();
    case 147:
      ACCEPT_TOKEN(anon_sym_COLON);
      END_STATE();
    case 148:
      ACCEPT_TOKEN(anon_sym_LBRACK);
      END_STATE();
    case 149:
      ACCEPT_TOKEN(anon_sym_RBRACK);
      END_STATE();
    case 150:
      ACCEPT_TOKEN(sym_html_character_reference);
      END_STATE();
    case 151:
      ACCEPT_TOKEN(anon_sym_GT);
      END_STATE();
    case 152:
      ACCEPT_TOKEN(anon_sym_GT);
      if (lookahead == '=') ADVANCE(237);
      if (lookahead == '>') ADVANCE(196);
      END_STATE();
    case 153:
      ACCEPT_TOKEN(anon_sym_GT);
      if (lookahead == '=') ADVANCE(237);
      if (lookahead == '>') ADVANCE(197);
      END_STATE();
    case 154:
      ACCEPT_TOKEN(sym_jsx_identifier);
      if (lookahead == '$' ||
          lookahead == '-' ||
          ('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(154);
      END_STATE();
    case 155:
      ACCEPT_TOKEN(anon_sym_DOT);
      END_STATE();
    case 156:
      ACCEPT_TOKEN(anon_sym_DOT);
      if (lookahead == '.') ADVANCE(32);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(277);
      END_STATE();
    case 157:
      ACCEPT_TOKEN(anon_sym_DOT);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(277);
      END_STATE();
    case 158:
      ACCEPT_TOKEN(anon_sym_LT_SLASH);
      END_STATE();
    case 159:
      ACCEPT_TOKEN(anon_sym_SLASH_GT);
      END_STATE();
    case 160:
      ACCEPT_TOKEN(anon_sym_DQUOTE);
      END_STATE();
    case 161:
      ACCEPT_TOKEN(anon_sym_SQUOTE);
      END_STATE();
    case 162:
      ACCEPT_TOKEN(sym_unescaped_double_jsx_string_fragment);
      if (lookahead == '&') ADVANCE(16);
      if (lookahead == '/') ADVANCE(163);
      if (set_contains(extras_character_set_1, 10, lookahead)) ADVANCE(162);
      if (lookahead != 0 &&
          lookahead != '"') ADVANCE(164);
      END_STATE();
    case 163:
      ACCEPT_TOKEN(sym_unescaped_double_jsx_string_fragment);
      if (lookahead == '&') ADVANCE(122);
      if (lookahead == '*') ADVANCE(166);
      if (lookahead == '/') ADVANCE(167);
      if (lookahead != 0 &&
          lookahead != '"') ADVANCE(164);
      END_STATE();
    case 164:
      ACCEPT_TOKEN(sym_unescaped_double_jsx_string_fragment);
      if (lookahead == '&') ADVANCE(122);
      if (lookahead != 0 &&
          lookahead != '"') ADVANCE(164);
      END_STATE();
    case 165:
      ACCEPT_TOKEN(sym_unescaped_double_jsx_string_fragment);
      if (lookahead == '&') ADVANCE(28);
      if (lookahead == '*') ADVANCE(165);
      if (lookahead == '/') ADVANCE(164);
      if (lookahead != 0 &&
          lookahead != '"') ADVANCE(166);
      END_STATE();
    case 166:
      ACCEPT_TOKEN(sym_unescaped_double_jsx_string_fragment);
      if (lookahead == '&') ADVANCE(28);
      if (lookahead == '*') ADVANCE(165);
      if (lookahead != 0 &&
          lookahead != '"') ADVANCE(166);
      END_STATE();
    case 167:
      ACCEPT_TOKEN(sym_unescaped_double_jsx_string_fragment);
      if (lookahead == '&') ADVANCE(261);
      if (lookahead == '\n' ||
          lookahead == '\r' ||
          lookahead == 0x2028 ||
          lookahead == 0x2029) ADVANCE(164);
      if (lookahead != 0 &&
          lookahead != '"') ADVANCE(167);
      END_STATE();
    case 168:
      ACCEPT_TOKEN(sym_unescaped_single_jsx_string_fragment);
      if (lookahead == '&') ADVANCE(17);
      if (lookahead == '/') ADVANCE(169);
      if (set_contains(extras_character_set_1, 10, lookahead)) ADVANCE(168);
      if (lookahead != 0 &&
          lookahead != '&' &&
          lookahead != '\'') ADVANCE(170);
      END_STATE();
    case 169:
      ACCEPT_TOKEN(sym_unescaped_single_jsx_string_fragment);
      if (lookahead == '&') ADVANCE(123);
      if (lookahead == '*') ADVANCE(172);
      if (lookahead == '/') ADVANCE(173);
      if (lookahead != 0 &&
          lookahead != '&' &&
          lookahead != '\'') ADVANCE(170);
      END_STATE();
    case 170:
      ACCEPT_TOKEN(sym_unescaped_single_jsx_string_fragment);
      if (lookahead == '&') ADVANCE(123);
      if (lookahead != 0 &&
          lookahead != '&' &&
          lookahead != '\'') ADVANCE(170);
      END_STATE();
    case 171:
      ACCEPT_TOKEN(sym_unescaped_single_jsx_string_fragment);
      if (lookahead == '&') ADVANCE(29);
      if (lookahead == '*') ADVANCE(171);
      if (lookahead == '/') ADVANCE(170);
      if (lookahead != 0 &&
          lookahead != '&' &&
          lookahead != '\'') ADVANCE(172);
      END_STATE();
    case 172:
      ACCEPT_TOKEN(sym_unescaped_single_jsx_string_fragment);
      if (lookahead == '&') ADVANCE(29);
      if (lookahead == '*') ADVANCE(171);
      if (lookahead != 0 &&
          lookahead != '&' &&
          lookahead != '\'') ADVANCE(172);
      END_STATE();
    case 173:
      ACCEPT_TOKEN(sym_unescaped_single_jsx_string_fragment);
      if (lookahead == '&') ADVANCE(262);
      if (lookahead == '\n' ||
          lookahead == '\r' ||
          lookahead == 0x2028 ||
          lookahead == 0x2029) ADVANCE(170);
      if (lookahead != 0 &&
          lookahead != '&' &&
          lookahead != '\'') ADVANCE(173);
      END_STATE();
    case 174:
      ACCEPT_TOKEN(anon_sym_EQ_GT);
      END_STATE();
    case 175:
      ACCEPT_TOKEN(anon_sym_QMARK_DOT);
      END_STATE();
    case 176:
      ACCEPT_TOKEN(anon_sym_PLUS_EQ);
      END_STATE();
    case 177:
      ACCEPT_TOKEN(anon_sym_DASH_EQ);
      END_STATE();
    case 178:
      ACCEPT_TOKEN(anon_sym_STAR_EQ);
      END_STATE();
    case 179:
      ACCEPT_TOKEN(anon_sym_SLASH_EQ);
      END_STATE();
    case 180:
      ACCEPT_TOKEN(anon_sym_PERCENT_EQ);
      END_STATE();
    case 181:
      ACCEPT_TOKEN(anon_sym_CARET_EQ);
      END_STATE();
    case 182:
      ACCEPT_TOKEN(anon_sym_AMP_EQ);
      END_STATE();
    case 183:
      ACCEPT_TOKEN(anon_sym_PIPE_EQ);
      END_STATE();
    case 184:
      ACCEPT_TOKEN(anon_sym_GT_GT_EQ);
      END_STATE();
    case 185:
      ACCEPT_TOKEN(anon_sym_GT_GT_GT_EQ);
      END_STATE();
    case 186:
      ACCEPT_TOKEN(anon_sym_LT_LT_EQ);
      END_STATE();
    case 187:
      ACCEPT_TOKEN(anon_sym_STAR_STAR_EQ);
      END_STATE();
    case 188:
      ACCEPT_TOKEN(anon_sym_AMP_AMP_EQ);
      END_STATE();
    case 189:
      ACCEPT_TOKEN(anon_sym_PIPE_PIPE_EQ);
      END_STATE();
    case 190:
      ACCEPT_TOKEN(anon_sym_QMARK_QMARK_EQ);
      END_STATE();
    case 191:
      ACCEPT_TOKEN(anon_sym_DOT_DOT_DOT);
      END_STATE();
    case 192:
      ACCEPT_TOKEN(anon_sym_AMP_AMP);
      END_STATE();
    case 193:
      ACCEPT_TOKEN(anon_sym_AMP_AMP);
      if (lookahead == '=') ADVANCE(188);
      END_STATE();
    case 194:
      ACCEPT_TOKEN(anon_sym_PIPE_PIPE);
      END_STATE();
    case 195:
      ACCEPT_TOKEN(anon_sym_PIPE_PIPE);
      if (lookahead == '=') ADVANCE(189);
      END_STATE();
    case 196:
      ACCEPT_TOKEN(anon_sym_GT_GT);
      if (lookahead == '=') ADVANCE(184);
      if (lookahead == '>') ADVANCE(199);
      END_STATE();
    case 197:
      ACCEPT_TOKEN(anon_sym_GT_GT);
      if (lookahead == '>') ADVANCE(198);
      END_STATE();
    case 198:
      ACCEPT_TOKEN(anon_sym_GT_GT_GT);
      END_STATE();
    case 199:
      ACCEPT_TOKEN(anon_sym_GT_GT_GT);
      if (lookahead == '=') ADVANCE(185);
      END_STATE();
    case 200:
      ACCEPT_TOKEN(anon_sym_LT_LT);
      END_STATE();
    case 201:
      ACCEPT_TOKEN(anon_sym_LT_LT);
      if (lookahead == '=') ADVANCE(186);
      END_STATE();
    case 202:
      ACCEPT_TOKEN(anon_sym_AMP);
      END_STATE();
    case 203:
      ACCEPT_TOKEN(anon_sym_AMP);
      if (lookahead == '&') ADVANCE(193);
      if (lookahead == '=') ADVANCE(182);
      END_STATE();
    case 204:
      ACCEPT_TOKEN(anon_sym_AMP);
      if (lookahead == '&') ADVANCE(192);
      END_STATE();
    case 205:
      ACCEPT_TOKEN(anon_sym_CARET);
      END_STATE();
    case 206:
      ACCEPT_TOKEN(anon_sym_CARET);
      if (lookahead == '=') ADVANCE(181);
      END_STATE();
    case 207:
      ACCEPT_TOKEN(anon_sym_PIPE);
      END_STATE();
    case 208:
      ACCEPT_TOKEN(anon_sym_PIPE);
      if (lookahead == '=') ADVANCE(183);
      if (lookahead == '|') ADVANCE(195);
      END_STATE();
    case 209:
      ACCEPT_TOKEN(anon_sym_PIPE);
      if (lookahead == '=') ADVANCE(183);
      if (lookahead == '|') ADVANCE(195);
      if (lookahead == '}') ADVANCE(293);
      END_STATE();
    case 210:
      ACCEPT_TOKEN(anon_sym_PIPE);
      if (lookahead == '|') ADVANCE(194);
      END_STATE();
    case 211:
      ACCEPT_TOKEN(anon_sym_PIPE);
      if (lookahead == '|') ADVANCE(194);
      if (lookahead == '}') ADVANCE(293);
      END_STATE();
    case 212:
      ACCEPT_TOKEN(anon_sym_PIPE);
      if (lookahead == '}') ADVANCE(293);
      END_STATE();
    case 213:
      ACCEPT_TOKEN(anon_sym_PLUS);
      END_STATE();
    case 214:
      ACCEPT_TOKEN(anon_sym_PLUS);
      if (lookahead == '+') ADVANCE(241);
      END_STATE();
    case 215:
      ACCEPT_TOKEN(anon_sym_PLUS);
      if (lookahead == '+') ADVANCE(241);
      if (lookahead == '=') ADVANCE(176);
      END_STATE();
    case 216:
      ACCEPT_TOKEN(anon_sym_PLUS);
      if (lookahead == '+') ADVANCE(241);
      if (lookahead == '=') ADVANCE(176);
      if (lookahead == '?') ADVANCE(35);
      END_STATE();
    case 217:
      ACCEPT_TOKEN(anon_sym_DASH);
      END_STATE();
    case 218:
      ACCEPT_TOKEN(anon_sym_DASH);
      if (lookahead == '-') ADVANCE(242);
      END_STATE();
    case 219:
      ACCEPT_TOKEN(anon_sym_DASH);
      if (lookahead == '-') ADVANCE(242);
      if (lookahead == '=') ADVANCE(177);
      END_STATE();
    case 220:
      ACCEPT_TOKEN(anon_sym_DASH);
      if (lookahead == '-') ADVANCE(242);
      if (lookahead == '=') ADVANCE(177);
      if (lookahead == '?') ADVANCE(36);
      END_STATE();
    case 221:
      ACCEPT_TOKEN(anon_sym_SLASH);
      if (lookahead == '*') ADVANCE(27);
      if (lookahead == '/') ADVANCE(263);
      END_STATE();
    case 222:
      ACCEPT_TOKEN(anon_sym_SLASH);
      if (lookahead == '*') ADVANCE(27);
      if (lookahead == '/') ADVANCE(263);
      if (lookahead == '=') ADVANCE(179);
      END_STATE();
    case 223:
      ACCEPT_TOKEN(anon_sym_PERCENT);
      END_STATE();
    case 224:
      ACCEPT_TOKEN(anon_sym_PERCENT);
      if (lookahead == '=') ADVANCE(180);
      END_STATE();
    case 225:
      ACCEPT_TOKEN(anon_sym_STAR_STAR);
      END_STATE();
    case 226:
      ACCEPT_TOKEN(anon_sym_STAR_STAR);
      if (lookahead == '=') ADVANCE(187);
      END_STATE();
    case 227:
      ACCEPT_TOKEN(anon_sym_LT);
      END_STATE();
    case 228:
      ACCEPT_TOKEN(anon_sym_LT);
      if (lookahead == '/') ADVANCE(158);
      END_STATE();
    case 229:
      ACCEPT_TOKEN(anon_sym_LT);
      if (lookahead == '/') ADVANCE(158);
      if (lookahead == '<') ADVANCE(201);
      if (lookahead == '=') ADVANCE(232);
      END_STATE();
    case 230:
      ACCEPT_TOKEN(anon_sym_LT);
      if (lookahead == '<') ADVANCE(201);
      if (lookahead == '=') ADVANCE(232);
      END_STATE();
    case 231:
      ACCEPT_TOKEN(anon_sym_LT);
      if (lookahead == '<') ADVANCE(200);
      if (lookahead == '=') ADVANCE(232);
      END_STATE();
    case 232:
      ACCEPT_TOKEN(anon_sym_LT_EQ);
      END_STATE();
    case 233:
      ACCEPT_TOKEN(anon_sym_EQ_EQ);
      if (lookahead == '=') ADVANCE(234);
      END_STATE();
    case 234:
      ACCEPT_TOKEN(anon_sym_EQ_EQ_EQ);
      END_STATE();
    case 235:
      ACCEPT_TOKEN(anon_sym_BANG_EQ);
      if (lookahead == '=') ADVANCE(236);
      END_STATE();
    case 236:
      ACCEPT_TOKEN(anon_sym_BANG_EQ_EQ);
      END_STATE();
    case 237:
      ACCEPT_TOKEN(anon_sym_GT_EQ);
      END_STATE();
    case 238:
      ACCEPT_TOKEN(anon_sym_QMARK_QMARK);
      END_STATE();
    case 239:
      ACCEPT_TOKEN(anon_sym_QMARK_QMARK);
      if (lookahead == '=') ADVANCE(190);
      END_STATE();
    case 240:
      ACCEPT_TOKEN(anon_sym_TILDE);
      END_STATE();
    case 241:
      ACCEPT_TOKEN(anon_sym_PLUS_PLUS);
      END_STATE();
    case 242:
      ACCEPT_TOKEN(anon_sym_DASH_DASH);
      END_STATE();
    case 243:
      ACCEPT_TOKEN(sym_unescaped_double_string_fragment);
      if (lookahead == '*') ADVANCE(245);
      if (lookahead == '/') ADVANCE(247);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '"' &&
          lookahead != '\\') ADVANCE(248);
      END_STATE();
    case 244:
      ACCEPT_TOKEN(sym_unescaped_double_string_fragment);
      if (lookahead == '*') ADVANCE(244);
      if (lookahead == '/') ADVANCE(248);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '"' &&
          lookahead != '\\') ADVANCE(245);
      END_STATE();
    case 245:
      ACCEPT_TOKEN(sym_unescaped_double_string_fragment);
      if (lookahead == '*') ADVANCE(244);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '"' &&
          lookahead != '\\') ADVANCE(245);
      END_STATE();
    case 246:
      ACCEPT_TOKEN(sym_unescaped_double_string_fragment);
      if (lookahead == '/') ADVANCE(243);
      if ((set_contains(extras_character_set_1, 10, lookahead)) &&
          lookahead != '\n' &&
          lookahead != '\r') ADVANCE(246);
      if (lookahead != 0 &&
          (lookahead < '\t' || '\r' < lookahead) &&
          lookahead != '"' &&
          lookahead != '\\') ADVANCE(248);
      END_STATE();
    case 247:
      ACCEPT_TOKEN(sym_unescaped_double_string_fragment);
      if (lookahead == 0x2028 ||
          lookahead == 0x2029) ADVANCE(248);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '"' &&
          lookahead != '\\') ADVANCE(247);
      END_STATE();
    case 248:
      ACCEPT_TOKEN(sym_unescaped_double_string_fragment);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '"' &&
          lookahead != '\\') ADVANCE(248);
      END_STATE();
    case 249:
      ACCEPT_TOKEN(sym_unescaped_single_string_fragment);
      if (lookahead == '*') ADVANCE(251);
      if (lookahead == '/') ADVANCE(253);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '\'' &&
          lookahead != '\\') ADVANCE(254);
      END_STATE();
    case 250:
      ACCEPT_TOKEN(sym_unescaped_single_string_fragment);
      if (lookahead == '*') ADVANCE(250);
      if (lookahead == '/') ADVANCE(254);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '\'' &&
          lookahead != '\\') ADVANCE(251);
      END_STATE();
    case 251:
      ACCEPT_TOKEN(sym_unescaped_single_string_fragment);
      if (lookahead == '*') ADVANCE(250);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '\'' &&
          lookahead != '\\') ADVANCE(251);
      END_STATE();
    case 252:
      ACCEPT_TOKEN(sym_unescaped_single_string_fragment);
      if (lookahead == '/') ADVANCE(249);
      if ((set_contains(extras_character_set_1, 10, lookahead)) &&
          lookahead != '\n' &&
          lookahead != '\r') ADVANCE(252);
      if (lookahead != 0 &&
          (lookahead < '\t' || '\r' < lookahead) &&
          lookahead != '\'' &&
          lookahead != '\\') ADVANCE(254);
      END_STATE();
    case 253:
      ACCEPT_TOKEN(sym_unescaped_single_string_fragment);
      if (lookahead == 0x2028 ||
          lookahead == 0x2029) ADVANCE(254);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '\'' &&
          lookahead != '\\') ADVANCE(253);
      END_STATE();
    case 254:
      ACCEPT_TOKEN(sym_unescaped_single_string_fragment);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '\'' &&
          lookahead != '\\') ADVANCE(254);
      END_STATE();
    case 255:
      ACCEPT_TOKEN(sym_escape_sequence);
      END_STATE();
    case 256:
      ACCEPT_TOKEN(sym_escape_sequence);
      if (lookahead == '\\') ADVANCE(84);
      if (set_contains(sym_identifier_character_set_2, 15, lookahead)) ADVANCE(282);
      END_STATE();
    case 257:
      ACCEPT_TOKEN(sym_escape_sequence);
      if (lookahead == '\n' ||
          lookahead == 0x2028 ||
          lookahead == 0x2029) ADVANCE(255);
      END_STATE();
    case 258:
      ACCEPT_TOKEN(sym_escape_sequence);
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(255);
      END_STATE();
    case 259:
      ACCEPT_TOKEN(sym_escape_sequence);
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(258);
      END_STATE();
    case 260:
      ACCEPT_TOKEN(sym_comment);
      END_STATE();
    case 261:
      ACCEPT_TOKEN(sym_comment);
      if (lookahead == '\n' ||
          lookahead == '\r' ||
          lookahead == 0x2028 ||
          lookahead == 0x2029) ADVANCE(164);
      if (lookahead == '#' ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(263);
      if (lookahead != 0) ADVANCE(167);
      END_STATE();
    case 262:
      ACCEPT_TOKEN(sym_comment);
      if (lookahead == '\n' ||
          lookahead == '\r' ||
          lookahead == 0x2028 ||
          lookahead == 0x2029) ADVANCE(170);
      if (lookahead == '#' ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(263);
      if (lookahead != 0) ADVANCE(173);
      END_STATE();
    case 263:
      ACCEPT_TOKEN(sym_comment);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != 0x2028 &&
          lookahead != 0x2029) ADVANCE(263);
      END_STATE();
    case 264:
      ACCEPT_TOKEN(anon_sym_BQUOTE);
      END_STATE();
    case 265:
      ACCEPT_TOKEN(anon_sym_DOLLAR_LBRACE);
      END_STATE();
    case 266:
      ACCEPT_TOKEN(anon_sym_SLASH2);
      END_STATE();
    case 267:
      ACCEPT_TOKEN(sym_regex_pattern);
      if (lookahead == '\n') SKIP(34);
      if (lookahead == '/') ADVANCE(24);
      if (lookahead == '[') ADVANCE(81);
      if (lookahead == '\\') ADVANCE(125);
      if (set_contains(extras_character_set_1, 10, lookahead)) ADVANCE(267);
      if (lookahead != 0) ADVANCE(268);
      END_STATE();
    case 268:
      ACCEPT_TOKEN(sym_regex_pattern);
      if (lookahead == '[') ADVANCE(81);
      if (lookahead == '\\') ADVANCE(125);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '/') ADVANCE(268);
      END_STATE();
    case 269:
      ACCEPT_TOKEN(sym_regex_flags);
      if (lookahead == '\\') ADVANCE(84);
      if (('a' <= lookahead && lookahead <= 'z')) ADVANCE(269);
      if (set_contains(sym_identifier_character_set_2, 15, lookahead)) ADVANCE(282);
      END_STATE();
    case 270:
      ACCEPT_TOKEN(sym_number);
      END_STATE();
    case 271:
      ACCEPT_TOKEN(sym_number);
      ADVANCE_MAP(
        '.', 279,
        '0', 273,
        '_', 102,
        'n', 270,
        'B', 98,
        'b', 98,
        'E', 96,
        'e', 96,
        'O', 99,
        'o', 99,
        'X', 107,
        'x', 107,
      );
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(272);
      END_STATE();
    case 272:
      ACCEPT_TOKEN(sym_number);
      if (lookahead == '.') ADVANCE(279);
      if (lookahead == '_') ADVANCE(100);
      if (lookahead == 'n') ADVANCE(270);
      if (lookahead == 'E' ||
          lookahead == 'e') ADVANCE(96);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(272);
      END_STATE();
    case 273:
      ACCEPT_TOKEN(sym_number);
      if (lookahead == '_') ADVANCE(102);
      if (lookahead == 'n') ADVANCE(270);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(273);
      END_STATE();
    case 274:
      ACCEPT_TOKEN(sym_number);
      if (lookahead == '_') ADVANCE(98);
      if (lookahead == 'n') ADVANCE(270);
      if (lookahead == '0' ||
          lookahead == '1') ADVANCE(274);
      END_STATE();
    case 275:
      ACCEPT_TOKEN(sym_number);
      if (lookahead == '_') ADVANCE(99);
      if (lookahead == 'n') ADVANCE(270);
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(275);
      END_STATE();
    case 276:
      ACCEPT_TOKEN(sym_number);
      if (lookahead == '_') ADVANCE(107);
      if (lookahead == 'n') ADVANCE(270);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(276);
      END_STATE();
    case 277:
      ACCEPT_TOKEN(sym_number);
      if (lookahead == '_') ADVANCE(101);
      if (lookahead == 'E' ||
          lookahead == 'e') ADVANCE(96);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(277);
      END_STATE();
    case 278:
      ACCEPT_TOKEN(sym_number);
      if (lookahead == '_') ADVANCE(103);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(278);
      END_STATE();
    case 279:
      ACCEPT_TOKEN(sym_number);
      if (lookahead == 'E' ||
          lookahead == 'e') ADVANCE(96);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(277);
      END_STATE();
    case 280:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '-') ADVANCE(154);
      if (lookahead == '\\') ADVANCE(84);
      if (lookahead == '$' ||
          ('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(280);
      if (set_contains(sym_identifier_character_set_2, 15, lookahead)) ADVANCE(282);
      END_STATE();
    case 281:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(84);
      if (lookahead == '{') ADVANCE(265);
      if (set_contains(sym_identifier_character_set_2, 15, lookahead)) ADVANCE(282);
      END_STATE();
    case 282:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(84);
      if (set_contains(sym_identifier_character_set_2, 15, lookahead)) ADVANCE(282);
      END_STATE();
    case 283:
      ACCEPT_TOKEN(sym_private_property_identifier);
      if (lookahead == '\\') ADVANCE(83);
      if (set_contains(sym_identifier_character_set_2, 15, lookahead)) ADVANCE(283);
      END_STATE();
    case 284:
      ACCEPT_TOKEN(anon_sym_AT);
      END_STATE();
    case 285:
      ACCEPT_TOKEN(anon_sym_QMARK);
      END_STATE();
    case 286:
      ACCEPT_TOKEN(anon_sym_QMARK);
      if (lookahead == '.') ADVANCE(175);
      END_STATE();
    case 287:
      ACCEPT_TOKEN(anon_sym_QMARK);
      if (lookahead == '.') ADVANCE(175);
      if (lookahead == '?') ADVANCE(239);
      END_STATE();
    case 288:
      ACCEPT_TOKEN(anon_sym_QMARK);
      if (lookahead == '.') ADVANCE(175);
      if (lookahead == '?') ADVANCE(238);
      END_STATE();
    case 289:
      ACCEPT_TOKEN(anon_sym_DASH_QMARK_COLON);
      END_STATE();
    case 290:
      ACCEPT_TOKEN(anon_sym_PLUS_QMARK_COLON);
      END_STATE();
    case 291:
      ACCEPT_TOKEN(anon_sym_QMARK_COLON);
      END_STATE();
    case 292:
      ACCEPT_TOKEN(anon_sym_LBRACE_PIPE);
      END_STATE();
    case 293:
      ACCEPT_TOKEN(anon_sym_PIPE_RBRACE);
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
        'a', 1,
        'b', 2,
        'c', 3,
        'd', 4,
        'e', 5,
        'f', 6,
        'g', 7,
        'i', 8,
        'k', 9,
        'l', 10,
        'm', 11,
        'n', 12,
        'o', 13,
        'p', 14,
        'r', 15,
        's', 16,
        't', 17,
        'u', 18,
        'v', 19,
        'w', 20,
        'y', 21,
      );
      if (set_contains(extras_character_set_1, 10, lookahead)) SKIP(0);
      END_STATE();
    case 1:
      if (lookahead == 'b') ADVANCE(22);
      if (lookahead == 'c') ADVANCE(23);
      if (lookahead == 'n') ADVANCE(24);
      if (lookahead == 's') ADVANCE(25);
      if (lookahead == 'w') ADVANCE(26);
      END_STATE();
    case 2:
      if (lookahead == 'o') ADVANCE(27);
      if (lookahead == 'r') ADVANCE(28);
      END_STATE();
    case 3:
      if (lookahead == 'a') ADVANCE(29);
      if (lookahead == 'l') ADVANCE(30);
      if (lookahead == 'o') ADVANCE(31);
      END_STATE();
    case 4:
      if (lookahead == 'e') ADVANCE(32);
      if (lookahead == 'o') ADVANCE(33);
      END_STATE();
    case 5:
      if (lookahead == 'l') ADVANCE(34);
      if (lookahead == 'n') ADVANCE(35);
      if (lookahead == 'x') ADVANCE(36);
      END_STATE();
    case 6:
      if (lookahead == 'a') ADVANCE(37);
      if (lookahead == 'i') ADVANCE(38);
      if (lookahead == 'o') ADVANCE(39);
      if (lookahead == 'r') ADVANCE(40);
      if (lookahead == 'u') ADVANCE(41);
      END_STATE();
    case 7:
      if (lookahead == 'e') ADVANCE(42);
      if (lookahead == 'l') ADVANCE(43);
      END_STATE();
    case 8:
      if (lookahead == 'f') ADVANCE(44);
      if (lookahead == 'm') ADVANCE(45);
      if (lookahead == 'n') ADVANCE(46);
      if (lookahead == 's') ADVANCE(47);
      END_STATE();
    case 9:
      if (lookahead == 'e') ADVANCE(48);
      END_STATE();
    case 10:
      if (lookahead == 'e') ADVANCE(49);
      END_STATE();
    case 11:
      if (lookahead == 'e') ADVANCE(50);
      if (lookahead == 'o') ADVANCE(51);
      END_STATE();
    case 12:
      if (lookahead == 'a') ADVANCE(52);
      if (lookahead == 'e') ADVANCE(53);
      if (lookahead == 'u') ADVANCE(54);
      END_STATE();
    case 13:
      if (lookahead == 'b') ADVANCE(55);
      if (lookahead == 'f') ADVANCE(56);
      if (lookahead == 'v') ADVANCE(57);
      END_STATE();
    case 14:
      if (lookahead == 'r') ADVANCE(58);
      if (lookahead == 'u') ADVANCE(59);
      END_STATE();
    case 15:
      if (lookahead == 'e') ADVANCE(60);
      END_STATE();
    case 16:
      if (lookahead == 'a') ADVANCE(61);
      if (lookahead == 'e') ADVANCE(62);
      if (lookahead == 't') ADVANCE(63);
      if (lookahead == 'u') ADVANCE(64);
      if (lookahead == 'w') ADVANCE(65);
      if (lookahead == 'y') ADVANCE(66);
      END_STATE();
    case 17:
      if (lookahead == 'a') ADVANCE(67);
      if (lookahead == 'h') ADVANCE(68);
      if (lookahead == 'r') ADVANCE(69);
      if (lookahead == 'y') ADVANCE(70);
      END_STATE();
    case 18:
      if (lookahead == 'n') ADVANCE(71);
      if (lookahead == 's') ADVANCE(72);
      END_STATE();
    case 19:
      if (lookahead == 'a') ADVANCE(73);
      if (lookahead == 'o') ADVANCE(74);
      END_STATE();
    case 20:
      if (lookahead == 'h') ADVANCE(75);
      if (lookahead == 'i') ADVANCE(76);
      END_STATE();
    case 21:
      if (lookahead == 'i') ADVANCE(77);
      END_STATE();
    case 22:
      if (lookahead == 's') ADVANCE(78);
      END_STATE();
    case 23:
      if (lookahead == 'c') ADVANCE(79);
      END_STATE();
    case 24:
      if (lookahead == 'y') ADVANCE(80);
      END_STATE();
    case 25:
      ACCEPT_TOKEN(anon_sym_as);
      if (lookahead == 's') ADVANCE(81);
      if (lookahead == 'y') ADVANCE(82);
      END_STATE();
    case 26:
      if (lookahead == 'a') ADVANCE(83);
      END_STATE();
    case 27:
      if (lookahead == 'o') ADVANCE(84);
      END_STATE();
    case 28:
      if (lookahead == 'e') ADVANCE(85);
      END_STATE();
    case 29:
      if (lookahead == 's') ADVANCE(86);
      if (lookahead == 't') ADVANCE(87);
      END_STATE();
    case 30:
      if (lookahead == 'a') ADVANCE(88);
      END_STATE();
    case 31:
      if (lookahead == 'n') ADVANCE(89);
      END_STATE();
    case 32:
      if (lookahead == 'b') ADVANCE(90);
      if (lookahead == 'c') ADVANCE(91);
      if (lookahead == 'f') ADVANCE(92);
      if (lookahead == 'l') ADVANCE(93);
      END_STATE();
    case 33:
      ACCEPT_TOKEN(anon_sym_do);
      END_STATE();
    case 34:
      if (lookahead == 's') ADVANCE(94);
      END_STATE();
    case 35:
      if (lookahead == 'u') ADVANCE(95);
      END_STATE();
    case 36:
      if (lookahead == 'p') ADVANCE(96);
      if (lookahead == 't') ADVANCE(97);
      END_STATE();
    case 37:
      if (lookahead == 'l') ADVANCE(98);
      END_STATE();
    case 38:
      if (lookahead == 'n') ADVANCE(99);
      END_STATE();
    case 39:
      if (lookahead == 'r') ADVANCE(100);
      END_STATE();
    case 40:
      if (lookahead == 'o') ADVANCE(101);
      END_STATE();
    case 41:
      if (lookahead == 'n') ADVANCE(102);
      END_STATE();
    case 42:
      if (lookahead == 't') ADVANCE(103);
      END_STATE();
    case 43:
      if (lookahead == 'o') ADVANCE(104);
      END_STATE();
    case 44:
      ACCEPT_TOKEN(anon_sym_if);
      END_STATE();
    case 45:
      if (lookahead == 'p') ADVANCE(105);
      END_STATE();
    case 46:
      ACCEPT_TOKEN(anon_sym_in);
      if (lookahead == 'f') ADVANCE(106);
      if (lookahead == 's') ADVANCE(107);
      if (lookahead == 't') ADVANCE(108);
      END_STATE();
    case 47:
      ACCEPT_TOKEN(anon_sym_is);
      END_STATE();
    case 48:
      if (lookahead == 'y') ADVANCE(109);
      END_STATE();
    case 49:
      if (lookahead == 't') ADVANCE(110);
      END_STATE();
    case 50:
      if (lookahead == 't') ADVANCE(111);
      END_STATE();
    case 51:
      if (lookahead == 'd') ADVANCE(112);
      END_STATE();
    case 52:
      if (lookahead == 'm') ADVANCE(113);
      END_STATE();
    case 53:
      if (lookahead == 'v') ADVANCE(114);
      if (lookahead == 'w') ADVANCE(115);
      END_STATE();
    case 54:
      if (lookahead == 'l') ADVANCE(116);
      if (lookahead == 'm') ADVANCE(117);
      END_STATE();
    case 55:
      if (lookahead == 'j') ADVANCE(118);
      END_STATE();
    case 56:
      ACCEPT_TOKEN(anon_sym_of);
      END_STATE();
    case 57:
      if (lookahead == 'e') ADVANCE(119);
      END_STATE();
    case 58:
      if (lookahead == 'i') ADVANCE(120);
      if (lookahead == 'o') ADVANCE(121);
      END_STATE();
    case 59:
      if (lookahead == 'b') ADVANCE(122);
      END_STATE();
    case 60:
      if (lookahead == 'a') ADVANCE(123);
      if (lookahead == 'q') ADVANCE(124);
      if (lookahead == 't') ADVANCE(125);
      END_STATE();
    case 61:
      if (lookahead == 't') ADVANCE(126);
      END_STATE();
    case 62:
      if (lookahead == 't') ADVANCE(127);
      END_STATE();
    case 63:
      if (lookahead == 'a') ADVANCE(128);
      if (lookahead == 'r') ADVANCE(129);
      END_STATE();
    case 64:
      if (lookahead == 'p') ADVANCE(130);
      END_STATE();
    case 65:
      if (lookahead == 'i') ADVANCE(131);
      END_STATE();
    case 66:
      if (lookahead == 'm') ADVANCE(132);
      END_STATE();
    case 67:
      if (lookahead == 'r') ADVANCE(133);
      END_STATE();
    case 68:
      if (lookahead == 'i') ADVANCE(134);
      if (lookahead == 'r') ADVANCE(135);
      END_STATE();
    case 69:
      if (lookahead == 'u') ADVANCE(136);
      if (lookahead == 'y') ADVANCE(137);
      END_STATE();
    case 70:
      if (lookahead == 'p') ADVANCE(138);
      END_STATE();
    case 71:
      if (lookahead == 'd') ADVANCE(139);
      if (lookahead == 'i') ADVANCE(140);
      if (lookahead == 'k') ADVANCE(141);
      END_STATE();
    case 72:
      if (lookahead == 'i') ADVANCE(142);
      END_STATE();
    case 73:
      if (lookahead == 'r') ADVANCE(143);
      END_STATE();
    case 74:
      if (lookahead == 'i') ADVANCE(144);
      END_STATE();
    case 75:
      if (lookahead == 'i') ADVANCE(145);
      END_STATE();
    case 76:
      if (lookahead == 't') ADVANCE(146);
      END_STATE();
    case 77:
      if (lookahead == 'e') ADVANCE(147);
      END_STATE();
    case 78:
      if (lookahead == 't') ADVANCE(148);
      END_STATE();
    case 79:
      if (lookahead == 'e') ADVANCE(149);
      END_STATE();
    case 80:
      ACCEPT_TOKEN(anon_sym_any);
      END_STATE();
    case 81:
      if (lookahead == 'e') ADVANCE(150);
      END_STATE();
    case 82:
      if (lookahead == 'n') ADVANCE(151);
      END_STATE();
    case 83:
      if (lookahead == 'i') ADVANCE(152);
      END_STATE();
    case 84:
      if (lookahead == 'l') ADVANCE(153);
      END_STATE();
    case 85:
      if (lookahead == 'a') ADVANCE(154);
      END_STATE();
    case 86:
      if (lookahead == 'e') ADVANCE(155);
      END_STATE();
    case 87:
      if (lookahead == 'c') ADVANCE(156);
      END_STATE();
    case 88:
      if (lookahead == 's') ADVANCE(157);
      END_STATE();
    case 89:
      if (lookahead == 's') ADVANCE(158);
      if (lookahead == 't') ADVANCE(159);
      END_STATE();
    case 90:
      if (lookahead == 'u') ADVANCE(160);
      END_STATE();
    case 91:
      if (lookahead == 'l') ADVANCE(161);
      END_STATE();
    case 92:
      if (lookahead == 'a') ADVANCE(162);
      END_STATE();
    case 93:
      if (lookahead == 'e') ADVANCE(163);
      END_STATE();
    case 94:
      if (lookahead == 'e') ADVANCE(164);
      END_STATE();
    case 95:
      if (lookahead == 'm') ADVANCE(165);
      END_STATE();
    case 96:
      if (lookahead == 'o') ADVANCE(166);
      END_STATE();
    case 97:
      if (lookahead == 'e') ADVANCE(167);
      END_STATE();
    case 98:
      if (lookahead == 's') ADVANCE(168);
      END_STATE();
    case 99:
      if (lookahead == 'a') ADVANCE(169);
      END_STATE();
    case 100:
      ACCEPT_TOKEN(anon_sym_for);
      END_STATE();
    case 101:
      if (lookahead == 'm') ADVANCE(170);
      END_STATE();
    case 102:
      if (lookahead == 'c') ADVANCE(171);
      END_STATE();
    case 103:
      ACCEPT_TOKEN(anon_sym_get);
      END_STATE();
    case 104:
      if (lookahead == 'b') ADVANCE(172);
      END_STATE();
    case 105:
      if (lookahead == 'l') ADVANCE(173);
      if (lookahead == 'o') ADVANCE(174);
      END_STATE();
    case 106:
      if (lookahead == 'e') ADVANCE(175);
      END_STATE();
    case 107:
      if (lookahead == 't') ADVANCE(176);
      END_STATE();
    case 108:
      if (lookahead == 'e') ADVANCE(177);
      END_STATE();
    case 109:
      if (lookahead == 'o') ADVANCE(178);
      END_STATE();
    case 110:
      ACCEPT_TOKEN(anon_sym_let);
      END_STATE();
    case 111:
      if (lookahead == 'a') ADVANCE(179);
      END_STATE();
    case 112:
      if (lookahead == 'u') ADVANCE(180);
      END_STATE();
    case 113:
      if (lookahead == 'e') ADVANCE(181);
      END_STATE();
    case 114:
      if (lookahead == 'e') ADVANCE(182);
      END_STATE();
    case 115:
      ACCEPT_TOKEN(anon_sym_new);
      END_STATE();
    case 116:
      if (lookahead == 'l') ADVANCE(183);
      END_STATE();
    case 117:
      if (lookahead == 'b') ADVANCE(184);
      END_STATE();
    case 118:
      if (lookahead == 'e') ADVANCE(185);
      END_STATE();
    case 119:
      if (lookahead == 'r') ADVANCE(186);
      END_STATE();
    case 120:
      if (lookahead == 'v') ADVANCE(187);
      END_STATE();
    case 121:
      if (lookahead == 't') ADVANCE(188);
      END_STATE();
    case 122:
      if (lookahead == 'l') ADVANCE(189);
      END_STATE();
    case 123:
      if (lookahead == 'd') ADVANCE(190);
      END_STATE();
    case 124:
      if (lookahead == 'u') ADVANCE(191);
      END_STATE();
    case 125:
      if (lookahead == 'u') ADVANCE(192);
      END_STATE();
    case 126:
      if (lookahead == 'i') ADVANCE(193);
      END_STATE();
    case 127:
      ACCEPT_TOKEN(anon_sym_set);
      END_STATE();
    case 128:
      if (lookahead == 't') ADVANCE(194);
      END_STATE();
    case 129:
      if (lookahead == 'i') ADVANCE(195);
      END_STATE();
    case 130:
      if (lookahead == 'e') ADVANCE(196);
      END_STATE();
    case 131:
      if (lookahead == 't') ADVANCE(197);
      END_STATE();
    case 132:
      if (lookahead == 'b') ADVANCE(198);
      END_STATE();
    case 133:
      if (lookahead == 'g') ADVANCE(199);
      END_STATE();
    case 134:
      if (lookahead == 's') ADVANCE(200);
      END_STATE();
    case 135:
      if (lookahead == 'o') ADVANCE(201);
      END_STATE();
    case 136:
      if (lookahead == 'e') ADVANCE(202);
      END_STATE();
    case 137:
      ACCEPT_TOKEN(anon_sym_try);
      END_STATE();
    case 138:
      if (lookahead == 'e') ADVANCE(203);
      END_STATE();
    case 139:
      if (lookahead == 'e') ADVANCE(204);
      END_STATE();
    case 140:
      if (lookahead == 'q') ADVANCE(205);
      END_STATE();
    case 141:
      if (lookahead == 'n') ADVANCE(206);
      END_STATE();
    case 142:
      if (lookahead == 'n') ADVANCE(207);
      END_STATE();
    case 143:
      ACCEPT_TOKEN(anon_sym_var);
      END_STATE();
    case 144:
      if (lookahead == 'd') ADVANCE(208);
      END_STATE();
    case 145:
      if (lookahead == 'l') ADVANCE(209);
      END_STATE();
    case 146:
      if (lookahead == 'h') ADVANCE(210);
      END_STATE();
    case 147:
      if (lookahead == 'l') ADVANCE(211);
      END_STATE();
    case 148:
      if (lookahead == 'r') ADVANCE(212);
      END_STATE();
    case 149:
      if (lookahead == 's') ADVANCE(213);
      END_STATE();
    case 150:
      if (lookahead == 'r') ADVANCE(214);
      END_STATE();
    case 151:
      if (lookahead == 'c') ADVANCE(215);
      END_STATE();
    case 152:
      if (lookahead == 't') ADVANCE(216);
      END_STATE();
    case 153:
      if (lookahead == 'e') ADVANCE(217);
      END_STATE();
    case 154:
      if (lookahead == 'k') ADVANCE(218);
      END_STATE();
    case 155:
      ACCEPT_TOKEN(anon_sym_case);
      END_STATE();
    case 156:
      if (lookahead == 'h') ADVANCE(219);
      END_STATE();
    case 157:
      if (lookahead == 's') ADVANCE(220);
      END_STATE();
    case 158:
      if (lookahead == 't') ADVANCE(221);
      END_STATE();
    case 159:
      if (lookahead == 'i') ADVANCE(222);
      END_STATE();
    case 160:
      if (lookahead == 'g') ADVANCE(223);
      END_STATE();
    case 161:
      if (lookahead == 'a') ADVANCE(224);
      END_STATE();
    case 162:
      if (lookahead == 'u') ADVANCE(225);
      END_STATE();
    case 163:
      if (lookahead == 't') ADVANCE(226);
      END_STATE();
    case 164:
      ACCEPT_TOKEN(anon_sym_else);
      END_STATE();
    case 165:
      ACCEPT_TOKEN(anon_sym_enum);
      END_STATE();
    case 166:
      if (lookahead == 'r') ADVANCE(227);
      END_STATE();
    case 167:
      if (lookahead == 'n') ADVANCE(228);
      END_STATE();
    case 168:
      if (lookahead == 'e') ADVANCE(229);
      END_STATE();
    case 169:
      if (lookahead == 'l') ADVANCE(230);
      END_STATE();
    case 170:
      ACCEPT_TOKEN(anon_sym_from);
      END_STATE();
    case 171:
      if (lookahead == 't') ADVANCE(231);
      END_STATE();
    case 172:
      if (lookahead == 'a') ADVANCE(232);
      END_STATE();
    case 173:
      if (lookahead == 'e') ADVANCE(233);
      END_STATE();
    case 174:
      if (lookahead == 'r') ADVANCE(234);
      END_STATE();
    case 175:
      if (lookahead == 'r') ADVANCE(235);
      END_STATE();
    case 176:
      if (lookahead == 'a') ADVANCE(236);
      END_STATE();
    case 177:
      if (lookahead == 'r') ADVANCE(237);
      END_STATE();
    case 178:
      if (lookahead == 'f') ADVANCE(238);
      END_STATE();
    case 179:
      ACCEPT_TOKEN(anon_sym_meta);
      END_STATE();
    case 180:
      if (lookahead == 'l') ADVANCE(239);
      END_STATE();
    case 181:
      if (lookahead == 's') ADVANCE(240);
      END_STATE();
    case 182:
      if (lookahead == 'r') ADVANCE(241);
      END_STATE();
    case 183:
      ACCEPT_TOKEN(sym_null);
      END_STATE();
    case 184:
      if (lookahead == 'e') ADVANCE(242);
      END_STATE();
    case 185:
      if (lookahead == 'c') ADVANCE(243);
      END_STATE();
    case 186:
      if (lookahead == 'r') ADVANCE(244);
      END_STATE();
    case 187:
      if (lookahead == 'a') ADVANCE(245);
      END_STATE();
    case 188:
      if (lookahead == 'e') ADVANCE(246);
      END_STATE();
    case 189:
      if (lookahead == 'i') ADVANCE(247);
      END_STATE();
    case 190:
      if (lookahead == 'o') ADVANCE(248);
      END_STATE();
    case 191:
      if (lookahead == 'i') ADVANCE(249);
      END_STATE();
    case 192:
      if (lookahead == 'r') ADVANCE(250);
      END_STATE();
    case 193:
      if (lookahead == 's') ADVANCE(251);
      END_STATE();
    case 194:
      if (lookahead == 'i') ADVANCE(252);
      END_STATE();
    case 195:
      if (lookahead == 'n') ADVANCE(253);
      END_STATE();
    case 196:
      if (lookahead == 'r') ADVANCE(254);
      END_STATE();
    case 197:
      if (lookahead == 'c') ADVANCE(255);
      END_STATE();
    case 198:
      if (lookahead == 'o') ADVANCE(256);
      END_STATE();
    case 199:
      if (lookahead == 'e') ADVANCE(257);
      END_STATE();
    case 200:
      ACCEPT_TOKEN(sym_this);
      END_STATE();
    case 201:
      if (lookahead == 'w') ADVANCE(258);
      END_STATE();
    case 202:
      ACCEPT_TOKEN(sym_true);
      END_STATE();
    case 203:
      ACCEPT_TOKEN(anon_sym_type);
      if (lookahead == 'o') ADVANCE(259);
      END_STATE();
    case 204:
      if (lookahead == 'f') ADVANCE(260);
      END_STATE();
    case 205:
      if (lookahead == 'u') ADVANCE(261);
      END_STATE();
    case 206:
      if (lookahead == 'o') ADVANCE(262);
      END_STATE();
    case 207:
      if (lookahead == 'g') ADVANCE(263);
      END_STATE();
    case 208:
      ACCEPT_TOKEN(anon_sym_void);
      END_STATE();
    case 209:
      if (lookahead == 'e') ADVANCE(264);
      END_STATE();
    case 210:
      ACCEPT_TOKEN(anon_sym_with);
      END_STATE();
    case 211:
      if (lookahead == 'd') ADVANCE(265);
      END_STATE();
    case 212:
      if (lookahead == 'a') ADVANCE(266);
      END_STATE();
    case 213:
      if (lookahead == 's') ADVANCE(267);
      END_STATE();
    case 214:
      if (lookahead == 't') ADVANCE(268);
      END_STATE();
    case 215:
      ACCEPT_TOKEN(anon_sym_async);
      END_STATE();
    case 216:
      ACCEPT_TOKEN(anon_sym_await);
      END_STATE();
    case 217:
      if (lookahead == 'a') ADVANCE(269);
      END_STATE();
    case 218:
      ACCEPT_TOKEN(anon_sym_break);
      END_STATE();
    case 219:
      ACCEPT_TOKEN(anon_sym_catch);
      END_STATE();
    case 220:
      ACCEPT_TOKEN(anon_sym_class);
      END_STATE();
    case 221:
      ACCEPT_TOKEN(anon_sym_const);
      END_STATE();
    case 222:
      if (lookahead == 'n') ADVANCE(270);
      END_STATE();
    case 223:
      if (lookahead == 'g') ADVANCE(271);
      END_STATE();
    case 224:
      if (lookahead == 'r') ADVANCE(272);
      END_STATE();
    case 225:
      if (lookahead == 'l') ADVANCE(273);
      END_STATE();
    case 226:
      if (lookahead == 'e') ADVANCE(274);
      END_STATE();
    case 227:
      if (lookahead == 't') ADVANCE(275);
      END_STATE();
    case 228:
      if (lookahead == 'd') ADVANCE(276);
      END_STATE();
    case 229:
      ACCEPT_TOKEN(sym_false);
      END_STATE();
    case 230:
      if (lookahead == 'l') ADVANCE(277);
      END_STATE();
    case 231:
      if (lookahead == 'i') ADVANCE(278);
      END_STATE();
    case 232:
      if (lookahead == 'l') ADVANCE(279);
      END_STATE();
    case 233:
      if (lookahead == 'm') ADVANCE(280);
      END_STATE();
    case 234:
      if (lookahead == 't') ADVANCE(281);
      END_STATE();
    case 235:
      ACCEPT_TOKEN(anon_sym_infer);
      END_STATE();
    case 236:
      if (lookahead == 'n') ADVANCE(282);
      END_STATE();
    case 237:
      if (lookahead == 'f') ADVANCE(283);
      END_STATE();
    case 238:
      ACCEPT_TOKEN(anon_sym_keyof);
      END_STATE();
    case 239:
      if (lookahead == 'e') ADVANCE(284);
      END_STATE();
    case 240:
      if (lookahead == 'p') ADVANCE(285);
      END_STATE();
    case 241:
      ACCEPT_TOKEN(anon_sym_never);
      END_STATE();
    case 242:
      if (lookahead == 'r') ADVANCE(286);
      END_STATE();
    case 243:
      if (lookahead == 't') ADVANCE(287);
      END_STATE();
    case 244:
      if (lookahead == 'i') ADVANCE(288);
      END_STATE();
    case 245:
      if (lookahead == 't') ADVANCE(289);
      END_STATE();
    case 246:
      if (lookahead == 'c') ADVANCE(290);
      END_STATE();
    case 247:
      if (lookahead == 'c') ADVANCE(291);
      END_STATE();
    case 248:
      if (lookahead == 'n') ADVANCE(292);
      END_STATE();
    case 249:
      if (lookahead == 'r') ADVANCE(293);
      END_STATE();
    case 250:
      if (lookahead == 'n') ADVANCE(294);
      END_STATE();
    case 251:
      if (lookahead == 'f') ADVANCE(295);
      END_STATE();
    case 252:
      if (lookahead == 'c') ADVANCE(296);
      END_STATE();
    case 253:
      if (lookahead == 'g') ADVANCE(297);
      END_STATE();
    case 254:
      ACCEPT_TOKEN(sym_super);
      END_STATE();
    case 255:
      if (lookahead == 'h') ADVANCE(298);
      END_STATE();
    case 256:
      if (lookahead == 'l') ADVANCE(299);
      END_STATE();
    case 257:
      if (lookahead == 't') ADVANCE(300);
      END_STATE();
    case 258:
      ACCEPT_TOKEN(anon_sym_throw);
      END_STATE();
    case 259:
      if (lookahead == 'f') ADVANCE(301);
      END_STATE();
    case 260:
      if (lookahead == 'i') ADVANCE(302);
      END_STATE();
    case 261:
      if (lookahead == 'e') ADVANCE(303);
      END_STATE();
    case 262:
      if (lookahead == 'w') ADVANCE(304);
      END_STATE();
    case 263:
      ACCEPT_TOKEN(anon_sym_using);
      END_STATE();
    case 264:
      ACCEPT_TOKEN(anon_sym_while);
      END_STATE();
    case 265:
      ACCEPT_TOKEN(anon_sym_yield);
      END_STATE();
    case 266:
      if (lookahead == 'c') ADVANCE(305);
      END_STATE();
    case 267:
      if (lookahead == 'o') ADVANCE(306);
      END_STATE();
    case 268:
      ACCEPT_TOKEN(anon_sym_assert);
      if (lookahead == 's') ADVANCE(307);
      END_STATE();
    case 269:
      if (lookahead == 'n') ADVANCE(308);
      END_STATE();
    case 270:
      if (lookahead == 'u') ADVANCE(309);
      END_STATE();
    case 271:
      if (lookahead == 'e') ADVANCE(310);
      END_STATE();
    case 272:
      if (lookahead == 'e') ADVANCE(311);
      END_STATE();
    case 273:
      if (lookahead == 't') ADVANCE(312);
      END_STATE();
    case 274:
      ACCEPT_TOKEN(anon_sym_delete);
      END_STATE();
    case 275:
      ACCEPT_TOKEN(anon_sym_export);
      END_STATE();
    case 276:
      if (lookahead == 's') ADVANCE(313);
      END_STATE();
    case 277:
      if (lookahead == 'y') ADVANCE(314);
      END_STATE();
    case 278:
      if (lookahead == 'o') ADVANCE(315);
      END_STATE();
    case 279:
      ACCEPT_TOKEN(anon_sym_global);
      END_STATE();
    case 280:
      if (lookahead == 'e') ADVANCE(316);
      END_STATE();
    case 281:
      ACCEPT_TOKEN(anon_sym_import);
      END_STATE();
    case 282:
      if (lookahead == 'c') ADVANCE(317);
      END_STATE();
    case 283:
      if (lookahead == 'a') ADVANCE(318);
      END_STATE();
    case 284:
      ACCEPT_TOKEN(anon_sym_module);
      END_STATE();
    case 285:
      if (lookahead == 'a') ADVANCE(319);
      END_STATE();
    case 286:
      ACCEPT_TOKEN(anon_sym_number);
      END_STATE();
    case 287:
      ACCEPT_TOKEN(anon_sym_object);
      END_STATE();
    case 288:
      if (lookahead == 'd') ADVANCE(320);
      END_STATE();
    case 289:
      if (lookahead == 'e') ADVANCE(321);
      END_STATE();
    case 290:
      if (lookahead == 't') ADVANCE(322);
      END_STATE();
    case 291:
      ACCEPT_TOKEN(anon_sym_public);
      END_STATE();
    case 292:
      if (lookahead == 'l') ADVANCE(323);
      END_STATE();
    case 293:
      if (lookahead == 'e') ADVANCE(324);
      END_STATE();
    case 294:
      ACCEPT_TOKEN(anon_sym_return);
      END_STATE();
    case 295:
      if (lookahead == 'i') ADVANCE(325);
      END_STATE();
    case 296:
      ACCEPT_TOKEN(anon_sym_static);
      END_STATE();
    case 297:
      ACCEPT_TOKEN(anon_sym_string);
      END_STATE();
    case 298:
      ACCEPT_TOKEN(anon_sym_switch);
      END_STATE();
    case 299:
      ACCEPT_TOKEN(anon_sym_symbol);
      END_STATE();
    case 300:
      ACCEPT_TOKEN(anon_sym_target);
      END_STATE();
    case 301:
      ACCEPT_TOKEN(anon_sym_typeof);
      END_STATE();
    case 302:
      if (lookahead == 'n') ADVANCE(326);
      END_STATE();
    case 303:
      ACCEPT_TOKEN(anon_sym_unique);
      END_STATE();
    case 304:
      if (lookahead == 'n') ADVANCE(327);
      END_STATE();
    case 305:
      if (lookahead == 't') ADVANCE(328);
      END_STATE();
    case 306:
      if (lookahead == 'r') ADVANCE(329);
      END_STATE();
    case 307:
      ACCEPT_TOKEN(anon_sym_asserts);
      END_STATE();
    case 308:
      ACCEPT_TOKEN(anon_sym_boolean);
      END_STATE();
    case 309:
      if (lookahead == 'e') ADVANCE(330);
      END_STATE();
    case 310:
      if (lookahead == 'r') ADVANCE(331);
      END_STATE();
    case 311:
      ACCEPT_TOKEN(anon_sym_declare);
      END_STATE();
    case 312:
      ACCEPT_TOKEN(anon_sym_default);
      END_STATE();
    case 313:
      ACCEPT_TOKEN(anon_sym_extends);
      END_STATE();
    case 314:
      ACCEPT_TOKEN(anon_sym_finally);
      END_STATE();
    case 315:
      if (lookahead == 'n') ADVANCE(332);
      END_STATE();
    case 316:
      if (lookahead == 'n') ADVANCE(333);
      END_STATE();
    case 317:
      if (lookahead == 'e') ADVANCE(334);
      END_STATE();
    case 318:
      if (lookahead == 'c') ADVANCE(335);
      END_STATE();
    case 319:
      if (lookahead == 'c') ADVANCE(336);
      END_STATE();
    case 320:
      if (lookahead == 'e') ADVANCE(337);
      END_STATE();
    case 321:
      ACCEPT_TOKEN(anon_sym_private);
      END_STATE();
    case 322:
      if (lookahead == 'e') ADVANCE(338);
      END_STATE();
    case 323:
      if (lookahead == 'y') ADVANCE(339);
      END_STATE();
    case 324:
      ACCEPT_TOKEN(anon_sym_require);
      END_STATE();
    case 325:
      if (lookahead == 'e') ADVANCE(340);
      END_STATE();
    case 326:
      if (lookahead == 'e') ADVANCE(341);
      END_STATE();
    case 327:
      ACCEPT_TOKEN(anon_sym_unknown);
      END_STATE();
    case 328:
      ACCEPT_TOKEN(anon_sym_abstract);
      END_STATE();
    case 329:
      ACCEPT_TOKEN(anon_sym_accessor);
      END_STATE();
    case 330:
      ACCEPT_TOKEN(anon_sym_continue);
      END_STATE();
    case 331:
      ACCEPT_TOKEN(anon_sym_debugger);
      END_STATE();
    case 332:
      ACCEPT_TOKEN(anon_sym_function);
      END_STATE();
    case 333:
      if (lookahead == 't') ADVANCE(342);
      END_STATE();
    case 334:
      if (lookahead == 'o') ADVANCE(343);
      END_STATE();
    case 335:
      if (lookahead == 'e') ADVANCE(344);
      END_STATE();
    case 336:
      if (lookahead == 'e') ADVANCE(345);
      END_STATE();
    case 337:
      ACCEPT_TOKEN(anon_sym_override);
      END_STATE();
    case 338:
      if (lookahead == 'd') ADVANCE(346);
      END_STATE();
    case 339:
      ACCEPT_TOKEN(anon_sym_readonly);
      END_STATE();
    case 340:
      if (lookahead == 's') ADVANCE(347);
      END_STATE();
    case 341:
      if (lookahead == 'd') ADVANCE(348);
      END_STATE();
    case 342:
      if (lookahead == 's') ADVANCE(349);
      END_STATE();
    case 343:
      if (lookahead == 'f') ADVANCE(350);
      END_STATE();
    case 344:
      ACCEPT_TOKEN(anon_sym_interface);
      END_STATE();
    case 345:
      ACCEPT_TOKEN(anon_sym_namespace);
      END_STATE();
    case 346:
      ACCEPT_TOKEN(anon_sym_protected);
      END_STATE();
    case 347:
      ACCEPT_TOKEN(anon_sym_satisfies);
      END_STATE();
    case 348:
      ACCEPT_TOKEN(sym_undefined);
      END_STATE();
    case 349:
      ACCEPT_TOKEN(anon_sym_implements);
      END_STATE();
    case 350:
      ACCEPT_TOKEN(anon_sym_instanceof);
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
static uint16_t ts_small_parse_table[167371];

// Filled from tables.zlib by grammar.go, see fetchgrammar.
static uint32_t ts_small_parse_table_map[4819];

// Filled from tables.zlib by grammar.go, see fetchgrammar.
static TSParseActionEntry ts_parse_actions[9910];

enum ts_external_scanner_symbol_identifiers {
  ts_external_token__automatic_semicolon = 0,
  ts_external_token__template_chars = 1,
  ts_external_token__ternary_qmark = 2,
  ts_external_token_html_comment = 3,
  ts_external_token_PIPE_PIPE = 4,
  ts_external_token_escape_sequence = 5,
  ts_external_token_regex_pattern = 6,
  ts_external_token_jsx_text = 7,
  ts_external_token__function_signature_automatic_semicolon = 8,
  ts_external_token___error_recovery = 9,
};

static const TSSymbol ts_external_scanner_symbol_map[EXTERNAL_TOKEN_COUNT] = {
  [ts_external_token__automatic_semicolon] = sym__automatic_semicolon,
  [ts_external_token__template_chars] = sym__template_chars,
  [ts_external_token__ternary_qmark] = sym__ternary_qmark,
  [ts_external_token_html_comment] = sym_html_comment,
  [ts_external_token_PIPE_PIPE] = anon_sym_PIPE_PIPE,
  [ts_external_token_escape_sequence] = sym_escape_sequence,
  [ts_external_token_regex_pattern] = sym_regex_pattern,
  [ts_external_token_jsx_text] = sym_jsx_text,
  [ts_external_token__function_signature_automatic_semicolon] = sym__function_signature_automatic_semicolon,
  [ts_external_token___error_recovery] = sym___error_recovery,
};

static const bool ts_external_scanner_states[12][EXTERNAL_TOKEN_COUNT] = {
  [1] = {
    [ts_external_token__automatic_semicolon] = true,
    [ts_external_token__template_chars] = true,
    [ts_external_token__ternary_qmark] = true,
    [ts_external_token_html_comment] = true,
    [ts_external_token_PIPE_PIPE] = true,
    [ts_external_token_escape_sequence] = true,
    [ts_external_token_jsx_text] = true,
    [ts_external_token__function_signature_automatic_semicolon] = true,
    [ts_external_token___error_recovery] = true,
  },
  [2] = {
    [ts_external_token_html_comment] = true,
  },
  [3] = {
    [ts_external_token__ternary_qmark] = true,
    [ts_external_token_html_comment] = true,
    [ts_external_token_PIPE_PIPE] = true,
  },
  [4] = {
    [ts_external_token__automatic_semicolon] = true,
    [ts_external_token__ternary_qmark] = true,
    [ts_external_token_html_comment] = true,
    [ts_external_token_PIPE_PIPE] = true,
  },
  [5] = {
    [ts_external_token__automatic_semicolon] = true,
    [ts_external_token_html_comment] = true,
  },
  [6] = {
    [ts_external_token__automatic_semicolon] = true,
    [ts_external_token_html_comment] = true,
    [ts_external_token__function_signature_automatic_semicolon] = true,
  },
  [7] = {
    [ts_external_token_html_comment] = true,
    [ts_external_token_jsx_text] = true,
  },
  [8] = {
    [ts_external_token__template_chars] = true,
    [ts_external_token_html_comment] = true,
    [ts_external_token_escape_sequence] = true,
  },
  [9] = {
    [ts_external_token__template_chars] = true,
    [ts_external_token_html_comment] = true,
  },
  [10] = {
    [ts_external_token_html_comment] = true,
    [ts_external_token_escape_sequence] = true,
  },
  [11] = {
    [ts_external_token_html_comment] = true,
    [ts_external_token_regex_pattern] = true,
  },
};

#ifdef __cplusplus
extern "C" {
#endif
void *tree_sitter_tsx_external_scanner_create(void);
void tree_sitter_tsx_external_scanner_destroy(void *);
bool tree_sitter_tsx_external_scanner_scan(void *, TSLexer *, const bool *);
unsigned tree_sitter_tsx_external_scanner_serialize(void *, char *);
void tree_sitter_tsx_external_scanner_deserialize(void *, const char *, unsigned);

#ifdef TREE_SITTER_HIDE_SYMBOLS
#define TS_PUBLIC
#elif defined(_WIN32)
#define TS_PUBLIC __declspec(dllexport)
#else
#define TS_PUBLIC __attribute__((visibility("default")))
#endif

TS_PUBLIC const TSLanguage *tree_sitter_tsx(void) {
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
    .keyword_capture_token = sym_identifier,
    .external_scanner = {
      &ts_external_scanner_states[0][0],
      ts_external_scanner_symbol_map,
      tree_sitter_tsx_external_scanner_create,
      tree_sitter_tsx_external_scanner_destroy,
      tree_sitter_tsx_external_scanner_scan,
      tree_sitter_tsx_external_scanner_serialize,
      tree_sitter_tsx_external_scanner_deserialize,
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
