#include "tree_sitter/parser.h"

#if defined(__GNUC__) || defined(__clang__)
#pragma GCC diagnostic ignored "-Wmissing-field-initializers"
#endif

#ifdef _MSC_VER
#pragma optimize("", off)
#elif defined(__clang__)
#pragma clang optimize off
#elif defined(__GNUC__)
#pragma GCC optimize ("O0")
#endif

#define LANGUAGE_VERSION 14
#define STATE_COUNT 10155
#define LARGE_STATE_COUNT 4529
#define SYMBOL_COUNT 357
#define ALIAS_COUNT 3
#define TOKEN_COUNT 153
#define EXTERNAL_TOKEN_COUNT 7
#define FIELD_COUNT 0
#define MAX_ALIAS_SEQUENCE_LENGTH 11
#define PRODUCTION_ID_COUNT 10

enum ts_symbol_identifiers {
  sym__alpha_identifier = 1,
  anon_sym_POUND_BANG = 2,
  aux_sym_shebang_line_token1 = 3,
  anon_sym_AT = 4,
  anon_sym_file = 5,
  anon_sym_COLON = 6,
  anon_sym_LBRACK = 7,
  anon_sym_RBRACK = 8,
  anon_sym_package = 9,
  anon_sym_import = 10,
  anon_sym_DOT = 11,
  sym_wildcard_import = 12,
  anon_sym_as = 13,
  anon_sym_typealias = 14,
  anon_sym_EQ = 15,
  anon_sym_class = 16,
  anon_sym_interface = 17,
  anon_sym_enum = 18,
  anon_sym_constructor = 19,
  anon_sym_LBRACE = 20,
  anon_sym_RBRACE = 21,
  anon_sym_LPAREN = 22,
  anon_sym_COMMA = 23,
  anon_sym_RPAREN = 24,
  anon_sym_val = 25,
  anon_sym_var = 26,
  anon_sym_by = 27,
  anon_sym_LT = 28,
  anon_sym_GT = 29,
  anon_sym_where = 30,
  anon_sym_init = 31,
  anon_sym_companion = 32,
  anon_sym_object = 33,
  anon_sym_fun = 34,
  anon_sym_SEMI = 35,
  anon_sym_get = 36,
  anon_sym_set = 37,
  anon_sym_this = 38,
  anon_sym_super = 39,
  anon_sym_dynamic = 40,
  anon_sym_AMP = 41,
  sym__quest = 42,
  anon_sym_STAR = 43,
  anon_sym_DASH_GT = 44,
  sym_label = 45,
  anon_sym_for = 46,
  anon_sym_in = 47,
  anon_sym_while = 48,
  anon_sym_do = 49,
  anon_sym_DOT_DOT = 50,
  anon_sym_QMARK_COLON = 51,
  anon_sym_AMP_AMP = 52,
  anon_sym_PIPE_PIPE = 53,
  anon_sym_null = 54,
  anon_sym_DOLLAR_LBRACE = 55,
  anon_sym_DOLLAR = 56,
  anon_sym_if = 57,
  anon_sym_else = 58,
  anon_sym_when = 59,
  anon_sym_try = 60,
  anon_sym_catch = 61,
  anon_sym_finally = 62,
  anon_sym_throw = 63,
  anon_sym_return = 64,
  anon_sym_continue = 65,
  anon_sym_break = 66,
  anon_sym_COLON_COLON = 67,
  anon_sym_PLUS_EQ = 68,
  anon_sym_DASH_EQ = 69,
  anon_sym_STAR_EQ = 70,
  anon_sym_SLASH_EQ = 71,
  anon_sym_PERCENT_EQ = 72,
  anon_sym_BANG_EQ = 73,
  anon_sym_BANG_EQ_EQ = 74,
  anon_sym_EQ_EQ = 75,
  anon_sym_EQ_EQ_EQ = 76,
  anon_sym_LT_EQ = 77,
  anon_sym_GT_EQ = 78,
  anon_sym_BANGin = 79,
  anon_sym_is = 80,
  anon_sym_BANGis = 81,
  anon_sym_PLUS = 82,
  anon_sym_DASH = 83,
  anon_sym_SLASH = 84,
  anon_sym_PERCENT = 85,
  anon_sym_as_QMARK = 86,
  anon_sym_PLUS_PLUS = 87,
  anon_sym_DASH_DASH = 88,
  anon_sym_BANG = 89,
  anon_sym_BANG_BANG = 90,
  anon_sym_suspend = 91,
  anon_sym_sealed = 92,
  anon_sym_annotation = 93,
  anon_sym_data = 94,
  anon_sym_inner = 95,
  anon_sym_value = 96,
  anon_sym_override = 97,
  anon_sym_lateinit = 98,
  anon_sym_public = 99,
  anon_sym_private = 100,
  anon_sym_internal = 101,
  anon_sym_protected = 102,
  anon_sym_out = 103,
  anon_sym_tailrec = 104,
  anon_sym_operator = 105,
  anon_sym_infix = 106,
  anon_sym_inline = 107,
  anon_sym_external = 108,
  sym_property_modifier = 109,
  anon_sym_abstract = 110,
  anon_sym_final = 111,
  anon_sym_open = 112,
  anon_sym_vararg = 113,
  anon_sym_noinline = 114,
  anon_sym_crossinline = 115,
  sym_reification_modifier = 116,
  anon_sym_expect = 117,
  anon_sym_actual = 118,
  anon_sym_field = 119,
  anon_sym_property = 120,
  anon_sym_receiver = 121,
  anon_sym_param = 122,
  anon_sym_setparam = 123,
  anon_sym_delegate = 124,
  sym_line_comment = 125,
  anon_sym_return_AT = 126,
  anon_sym_continue_AT = 127,
  anon_sym_break_AT = 128,
  anon_sym_this_AT = 129,
  anon_sym_super_AT = 130,
  anon_sym_AT2 = 131,
  sym_real_literal = 132,
  sym_integer_literal = 133,
  sym_hex_literal = 134,
  sym_bin_literal = 135,
  aux_sym_unsigned_literal_token1 = 136,
  anon_sym_L = 137,
  anon_sym_true = 138,
  anon_sym_false = 139,
  anon_sym_SQUOTE = 140,
  aux_sym_character_literal_token1 = 141,
  sym__backtick_identifier = 142,
  anon_sym_BSLASHu = 143,
  aux_sym__uni_character_literal_token1 = 144,
  sym__escaped_identifier = 145,
  sym__automatic_semicolon = 146,
  sym__import_list_delimiter = 147,
  sym_safe_nav = 148,
  sym_multiline_comment = 149,
  sym__string_start = 150,
  sym__string_end = 151,
  sym_string_content = 152,
  sym_source_file = 153,
  sym_shebang_line = 154,
  sym_file_annotation = 155,
  sym_package_header = 156,
  sym_import_list = 157,
  sym_import_header = 158,
  sym_import_alias = 159,
  sym_type_alias = 160,
  sym__declaration = 161,
  sym_class_declaration = 162,
  sym_primary_constructor = 163,
  sym_class_body = 164,
  sym__class_parameters = 165,
  sym_binding_pattern_kind = 166,
  sym_class_parameter = 167,
  sym__delegation_specifiers = 168,
  sym_delegation_specifier = 169,
  sym_constructor_invocation = 170,
  sym_explicit_delegation = 171,
  sym_type_parameters = 172,
  sym_type_parameter = 173,
  sym_type_constraints = 174,
  sym_type_constraint = 175,
  aux_sym__class_member_declarations = 176,
  sym__class_member_declaration = 177,
  sym_anonymous_initializer = 178,
  sym_companion_object = 179,
  sym_function_value_parameters = 180,
  sym__function_value_parameter = 181,
  sym__receiver_type = 182,
  sym_function_declaration = 183,
  sym_function_body = 184,
  sym_variable_declaration = 185,
  sym_property_declaration = 186,
  sym_property_delegate = 187,
  sym_getter = 188,
  sym_setter = 189,
  sym_parameter_with_optional_type = 190,
  sym_parameter = 191,
  sym_object_declaration = 192,
  sym_secondary_constructor = 193,
  sym_constructor_delegation_call = 194,
  sym_enum_class_body = 195,
  sym__enum_entries = 196,
  sym_enum_entry = 197,
  sym__type = 198,
  sym__type_reference = 199,
  sym_not_nullable_type = 200,
  sym_nullable_type = 201,
  sym_user_type = 202,
  sym__simple_user_type = 203,
  sym_type_projection = 204,
  sym_type_projection_modifiers = 205,
  sym__type_projection_modifier = 206,
  sym_function_type = 207,
  sym_function_type_parameters = 208,
  sym_parenthesized_type = 209,
  sym_parenthesized_user_type = 210,
  sym_statements = 211,
  sym__statement = 212,
  sym_control_structure_body = 213,
  sym__block = 214,
  sym__loop_statement = 215,
  sym_for_statement = 216,
  sym_while_statement = 217,
  sym_do_while_statement = 218,
  sym__semi = 219,
  sym_assignment = 220,
  sym__expression = 221,
  sym__unary_expression = 222,
  sym_postfix_expression = 223,
  sym_call_expression = 224,
  sym_indexing_expression = 225,
  sym_navigation_expression = 226,
  sym_prefix_expression = 227,
  sym_as_expression = 228,
  sym_spread_expression = 229,
  sym__binary_expression = 230,
  sym_multiplicative_expression = 231,
  sym_additive_expression = 232,
  sym_range_expression = 233,
  sym_infix_expression = 234,
  sym_elvis_expression = 235,
  sym_check_expression = 236,
  sym_comparison_expression = 237,
  sym_equality_expression = 238,
  sym_conjunction_expression = 239,
  sym_disjunction_expression = 240,
  sym_indexing_suffix = 241,
  sym_navigation_suffix = 242,
  sym_call_suffix = 243,
  sym_annotated_lambda = 244,
  sym_type_arguments = 245,
  sym_value_arguments = 246,
  sym_value_argument = 247,
  sym__primary_expression = 248,
  sym_parenthesized_expression = 249,
  sym_collection_literal = 250,
  sym__literal_constant = 251,
  sym_string_literal = 252,
  sym__interpolation = 253,
  sym_lambda_literal = 254,
  sym_multi_variable_declaration = 255,
  sym_lambda_parameters = 256,
  sym__lambda_parameter = 257,
  sym_anonymous_function = 258,
  sym__function_literal = 259,
  sym_object_literal = 260,
  sym_this_expression = 261,
  sym_super_expression = 262,
  sym_if_expression = 263,
  sym_when_subject = 264,
  sym_when_expression = 265,
  sym_when_entry = 266,
  sym_when_condition = 267,
  sym_range_test = 268,
  sym_type_test = 269,
  sym_try_expression = 270,
  sym_catch_block = 271,
  sym_finally_block = 272,
  sym_jump_expression = 273,
  sym_callable_reference = 274,
  sym__assignment_and_operator = 275,
  sym__equality_operator = 276,
  sym__comparison_operator = 277,
  sym__in_operator = 278,
  sym__is_operator = 279,
  sym__additive_operator = 280,
  sym__multiplicative_operator = 281,
  sym__as_operator = 282,
  sym__prefix_unary_operator = 283,
  sym__postfix_unary_operator = 284,
  sym__member_access_operator = 285,
  sym__postfix_unary_suffix = 286,
  sym__postfix_unary_expression = 287,
  sym_directly_assignable_expression = 288,
  sym_modifiers = 289,
  sym_parameter_modifiers = 290,
  sym__modifier = 291,
  sym_type_modifiers = 292,
  sym__type_modifier = 293,
  sym_class_modifier = 294,
  sym_member_modifier = 295,
  sym_visibility_modifier = 296,
  sym_variance_modifier = 297,
  sym_type_parameter_modifiers = 298,
  sym__type_parameter_modifier = 299,
  sym_function_modifier = 300,
  sym_inheritance_modifier = 301,
  sym_parameter_modifier = 302,
  sym_platform_modifier = 303,
  sym_annotation = 304,
  sym__single_annotation = 305,
  sym__multi_annotation = 306,
  sym_use_site_target = 307,
  sym__unescaped_annotation = 308,
  sym_simple_identifier = 309,
  sym_identifier = 310,
  sym__import_identifier = 311,
  sym__return_at = 312,
  sym__continue_at = 313,
  sym__break_at = 314,
  sym__this_at = 315,
  sym__super_at = 316,
  sym_unsigned_literal = 317,
  sym_long_literal = 318,
  sym_boolean_literal = 319,
  sym_character_literal = 320,
  sym_character_escape_seq = 321,
  sym__lexical_identifier = 322,
  sym__uni_character_literal = 323,
  aux_sym_source_file_repeat1 = 324,
  aux_sym_source_file_repeat2 = 325,
  aux_sym_source_file_repeat3 = 326,
  aux_sym_file_annotation_repeat1 = 327,
  aux_sym_import_list_repeat1 = 328,
  aux_sym__class_parameters_repeat1 = 329,
  aux_sym__delegation_specifiers_repeat1 = 330,
  aux_sym__annotated_delegation_specifier_repeat1 = 331,
  aux_sym_type_parameters_repeat1 = 332,
  aux_sym_type_constraints_repeat1 = 333,
  aux_sym_function_value_parameters_repeat1 = 334,
  aux_sym__enum_entries_repeat1 = 335,
  aux_sym_nullable_type_repeat1 = 336,
  aux_sym_user_type_repeat1 = 337,
  aux_sym_type_projection_modifiers_repeat1 = 338,
  aux_sym_function_type_parameters_repeat1 = 339,
  aux_sym_statements_repeat1 = 340,
  aux_sym__statement_repeat1 = 341,
  aux_sym_indexing_suffix_repeat1 = 342,
  aux_sym_type_arguments_repeat1 = 343,
  aux_sym_value_arguments_repeat1 = 344,
  aux_sym_string_literal_repeat1 = 345,
  aux_sym_multi_variable_declaration_repeat1 = 346,
  aux_sym_lambda_parameters_repeat1 = 347,
  aux_sym_when_expression_repeat1 = 348,
  aux_sym_when_entry_repeat1 = 349,
  aux_sym_try_expression_repeat1 = 350,
  aux_sym__postfix_unary_expression_repeat1 = 351,
  aux_sym_modifiers_repeat1 = 352,
  aux_sym_parameter_modifiers_repeat1 = 353,
  aux_sym_type_modifiers_repeat1 = 354,
  aux_sym_type_parameter_modifiers_repeat1 = 355,
  aux_sym_identifier_repeat1 = 356,
  alias_sym_interpolated_expression = 357,
  alias_sym_interpolated_identifier = 358,
  alias_sym_type_identifier = 359,
};

static const char * const ts_symbol_names[] = {
  [ts_builtin_sym_end] = "end",
  [sym__alpha_identifier] = "_alpha_identifier",
  [anon_sym_POUND_BANG] = "#!",
  [aux_sym_shebang_line_token1] = "shebang_line_token1",
  [anon_sym_AT] = "@",
  [anon_sym_file] = "file",
  [anon_sym_COLON] = ":",
  [anon_sym_LBRACK] = "[",
  [anon_sym_RBRACK] = "]",
  [anon_sym_package] = "package",
  [anon_sym_import] = "import",
  [anon_sym_DOT] = ".",
  [sym_wildcard_import] = "wildcard_import",
  [anon_sym_as] = "as",
  [anon_sym_typealias] = "typealias",
  [anon_sym_EQ] = "=",
  [anon_sym_class] = "class",
  [anon_sym_interface] = "interface",
  [anon_sym_enum] = "enum",
  [anon_sym_constructor] = "constructor",
  [anon_sym_LBRACE] = "{",
  [anon_sym_RBRACE] = "}",
  [anon_sym_LPAREN] = "(",
  [anon_sym_COMMA] = ",",
  [anon_sym_RPAREN] = ")",
  [anon_sym_val] = "val",
  [anon_sym_var] = "var",
  [anon_sym_by] = "by",
  [anon_sym_LT] = "<",
  [anon_sym_GT] = ">",
  [anon_sym_where] = "where",
  [anon_sym_init] = "init",
  [anon_sym_companion] = "companion",
  [anon_sym_object] = "object",
  [anon_sym_fun] = "fun",
  [anon_sym_SEMI] = ";",
  [anon_sym_get] = "get",
  [anon_sym_set] = "set",
  [anon_sym_this] = "this",
  [anon_sym_super] = "super",
  [anon_sym_dynamic] = "dynamic",
  [anon_sym_AMP] = "&",
  [sym__quest] = "_quest",
  [anon_sym_STAR] = "*",
  [anon_sym_DASH_GT] = "->",
  [sym_label] = "label",
  [anon_sym_for] = "for",
  [anon_sym_in] = "in",
  [anon_sym_while] = "while",
  [anon_sym_do] = "do",
  [anon_sym_DOT_DOT] = "..",
  [anon_sym_QMARK_COLON] = "\?:",
  [anon_sym_AMP_AMP] = "&&",
  [anon_sym_PIPE_PIPE] = "||",
  [anon_sym_null] = "null",
  [anon_sym_DOLLAR_LBRACE] = "${",
  [anon_sym_DOLLAR] = "$",
  [anon_sym_if] = "if",
  [anon_sym_else] = "else",
  [anon_sym_when] = "when",
  [anon_sym_try] = "try",
  [anon_sym_catch] = "catch",
  [anon_sym_finally] = "finally",
  [anon_sym_throw] = "throw",
  [anon_sym_return] = "return",
  [anon_sym_continue] = "continue",
  [anon_sym_break] = "break",
  [anon_sym_COLON_COLON] = "::",
  [anon_sym_PLUS_EQ] = "+=",
  [anon_sym_DASH_EQ] = "-=",
  [anon_sym_STAR_EQ] = "*=",
  [anon_sym_SLASH_EQ] = "/=",
  [anon_sym_PERCENT_EQ] = "%=",
  [anon_sym_BANG_EQ] = "!=",
  [anon_sym_BANG_EQ_EQ] = "!==",
  [anon_sym_EQ_EQ] = "==",
  [anon_sym_EQ_EQ_EQ] = "===",
  [anon_sym_LT_EQ] = "<=",
  [anon_sym_GT_EQ] = ">=",
  [anon_sym_BANGin] = "!in",
  [anon_sym_is] = "is",
  [anon_sym_BANGis] = "!is",
  [anon_sym_PLUS] = "+",
  [anon_sym_DASH] = "-",
  [anon_sym_SLASH] = "/",
  [anon_sym_PERCENT] = "%",
  [anon_sym_as_QMARK] = "as\?",
  [anon_sym_PLUS_PLUS] = "++",
  [anon_sym_DASH_DASH] = "--",
  [anon_sym_BANG] = "!",
  [anon_sym_BANG_BANG] = "!!",
  [anon_sym_suspend] = "suspend",
  [anon_sym_sealed] = "sealed",
  [anon_sym_annotation] = "annotation",
  [anon_sym_data] = "data",
  [anon_sym_inner] = "inner",
  [anon_sym_value] = "value",
  [anon_sym_override] = "override",
  [anon_sym_lateinit] = "lateinit",
  [anon_sym_public] = "public",
  [anon_sym_private] = "private",
  [anon_sym_internal] = "internal",
  [anon_sym_protected] = "protected",
  [anon_sym_out] = "out",
  [anon_sym_tailrec] = "tailrec",
  [anon_sym_operator] = "operator",
  [anon_sym_infix] = "infix",
  [anon_sym_inline] = "inline",
  [anon_sym_external] = "external",
  [sym_property_modifier] = "property_modifier",
  [anon_sym_abstract] = "abstract",
  [anon_sym_final] = "final",
  [anon_sym_open] = "open",
  [anon_sym_vararg] = "vararg",
  [anon_sym_noinline] = "noinline",
  [anon_sym_crossinline] = "crossinline",
  [sym_reification_modifier] = "reification_modifier",
  [anon_sym_expect] = "expect",
  [anon_sym_actual] = "actual",
  [anon_sym_field] = "field",
  [anon_sym_property] = "property",
  [anon_sym_receiver] = "receiver",
  [anon_sym_param] = "param",
  [anon_sym_setparam] = "setparam",
  [anon_sym_delegate] = "delegate",
  [sym_line_comment] = "line_comment",
  [anon_sym_return_AT] = "return@",
  [anon_sym_continue_AT] = "continue@",
  [anon_sym_break_AT] = "break@",
  [anon_sym_this_AT] = "this@",
  [anon_sym_super_AT] = "super@",
  [anon_sym_AT2] = "@",
  [sym_real_literal] = "real_literal",
  [sym_integer_literal] = "integer_literal",
  [sym_hex_literal] = "hex_literal",
  [sym_bin_literal] = "bin_literal",
  [aux_sym_unsigned_literal_token1] = "unsigned_literal_token1",
  [anon_sym_L] = "L",
  [anon_sym_true] = "true",
  [anon_sym_false] = "false",
  [anon_sym_SQUOTE] = "'",
  [aux_sym_character_literal_token1] = "character_literal_token1",
  [sym__backtick_identifier] = "_backtick_identifier",
  [anon_sym_BSLASHu] = "\\u",
  [aux_sym__uni_character_literal_token1] = "_uni_character_literal_token1",
  [sym__escaped_identifier] = "_escaped_identifier",
  [sym__automatic_semicolon] = "_automatic_semicolon",
  [sym__import_list_delimiter] = "_import_list_delimiter",
  [sym_safe_nav] = "\?.",
  [sym_multiline_comment] = "multiline_comment",
  [sym__string_start] = "_string_start",
  [sym__string_end] = "_string_end",
  [sym_string_content] = "string_content",
  [sym_source_file] = "source_file",
  [sym_shebang_line] = "shebang_line",
  [sym_file_annotation] = "file_annotation",
  [sym_package_header] = "package_header",
  [sym_import_list] = "import_list",
  [sym_import_header] = "import_header",
  [sym_import_alias] = "import_alias",
  [sym_type_alias] = "type_alias",
  [sym__declaration] = "_declaration",
  [sym_class_declaration] = "class_declaration",
  [sym_primary_constructor] = "primary_constructor",
  [sym_class_body] = "class_body",
  [sym__class_parameters] = "_class_parameters",
  [sym_binding_pattern_kind] = "binding_pattern_kind",
  [sym_class_parameter] = "class_parameter",
  [sym__delegation_specifiers] = "_delegation_specifiers",
  [sym_delegation_specifier] = "delegation_specifier",
  [sym_constructor_invocation] = "constructor_invocation",
  [sym_explicit_delegation] = "explicit_delegation",
  [sym_type_parameters] = "type_parameters",
  [sym_type_parameter] = "type_parameter",
  [sym_type_constraints] = "type_constraints",
  [sym_type_constraint] = "type_constraint",
  [aux_sym__class_member_declarations] = "_class_member_declarations",
  [sym__class_member_declaration] = "_class_member_declaration",
  [sym_anonymous_initializer] = "anonymous_initializer",
  [sym_companion_object] = "companion_object",
  [sym_function_value_parameters] = "function_value_parameters",
  [sym__function_value_parameter] = "_function_value_parameter",
  [sym__receiver_type] = "_receiver_type",
  [sym_function_declaration] = "function_declaration",
  [sym_function_body] = "function_body",
  [sym_variable_declaration] = "variable_declaration",
  [sym_property_declaration] = "property_declaration",
  [sym_property_delegate] = "property_delegate",
  [sym_getter] = "getter",
  [sym_setter] = "setter",
  [sym_parameter_with_optional_type] = "parameter_with_optional_type",
  [sym_parameter] = "parameter",
  [sym_object_declaration] = "object_declaration",
  [sym_secondary_constructor] = "secondary_constructor",
  [sym_constructor_delegation_call] = "constructor_delegation_call",
  [sym_enum_class_body] = "enum_class_body",
  [sym__enum_entries] = "_enum_entries",
  [sym_enum_entry] = "enum_entry",
  [sym__type] = "_type",
  [sym__type_reference] = "_type_reference",
  [sym_not_nullable_type] = "not_nullable_type",
  [sym_nullable_type] = "nullable_type",
  [sym_user_type] = "user_type",
  [sym__simple_user_type] = "_simple_user_type",
  [sym_type_projection] = "type_projection",
  [sym_type_projection_modifiers] = "type_projection_modifiers",
  [sym__type_projection_modifier] = "_type_projection_modifier",
  [sym_function_type] = "function_type",
  [sym_function_type_parameters] = "function_type_parameters",
  [sym_parenthesized_type] = "parenthesized_type",
  [sym_parenthesized_user_type] = "parenthesized_user_type",
  [sym_statements] = "statements",
  [sym__statement] = "_statement",
  [sym_control_structure_body] = "control_structure_body",
  [sym__block] = "_block",
  [sym__loop_statement] = "_loop_statement",
  [sym_for_statement] = "for_statement",
  [sym_while_statement] = "while_statement",
  [sym_do_while_statement] = "do_while_statement",
  [sym__semi] = "_semi",
  [sym_assignment] = "assignment",
  [sym__expression] = "_expression",
  [sym__unary_expression] = "_unary_expression",
  [sym_postfix_expression] = "postfix_expression",
  [sym_call_expression] = "call_expression",
  [sym_indexing_expression] = "indexing_expression",
  [sym_navigation_expression] = "navigation_expression",
  [sym_prefix_expression] = "prefix_expression",
  [sym_as_expression] = "as_expression",
  [sym_spread_expression] = "spread_expression",
  [sym__binary_expression] = "_binary_expression",
  [sym_multiplicative_expression] = "multiplicative_expression",
  [sym_additive_expression] = "additive_expression",
  [sym_range_expression] = "range_expression",
  [sym_infix_expression] = "infix_expression",
  [sym_elvis_expression] = "elvis_expression",
  [sym_check_expression] = "check_expression",
  [sym_comparison_expression] = "comparison_expression",
  [sym_equality_expression] = "equality_expression",
  [sym_conjunction_expression] = "conjunction_expression",
  [sym_disjunction_expression] = "disjunction_expression",
  [sym_indexing_suffix] = "indexing_suffix",
  [sym_navigation_suffix] = "navigation_suffix",
  [sym_call_suffix] = "call_suffix",
  [sym_annotated_lambda] = "annotated_lambda",
  [sym_type_arguments] = "type_arguments",
  [sym_value_arguments] = "value_arguments",
  [sym_value_argument] = "value_argument",
  [sym__primary_expression] = "_primary_expression",
  [sym_parenthesized_expression] = "parenthesized_expression",
  [sym_collection_literal] = "collection_literal",
  [sym__literal_constant] = "_literal_constant",
  [sym_string_literal] = "string_literal",
  [sym__interpolation] = "_interpolation",
  [sym_lambda_literal] = "lambda_literal",
  [sym_multi_variable_declaration] = "multi_variable_declaration",
  [sym_lambda_parameters] = "lambda_parameters",
  [sym__lambda_parameter] = "_lambda_parameter",
  [sym_anonymous_function] = "anonymous_function",
  [sym__function_literal] = "_function_literal",
  [sym_object_literal] = "object_literal",
  [sym_this_expression] = "this_expression",
  [sym_super_expression] = "super_expression",
  [sym_if_expression] = "if_expression",
  [sym_when_subject] = "when_subject",
  [sym_when_expression] = "when_expression",
  [sym_when_entry] = "when_entry",
  [sym_when_condition] = "when_condition",
  [sym_range_test] = "range_test",
  [sym_type_test] = "type_test",
  [sym_try_expression] = "try_expression",
  [sym_catch_block] = "catch_block",
  [sym_finally_block] = "finally_block",
  [sym_jump_expression] = "jump_expression",
  [sym_callable_reference] = "callable_reference",
  [sym__assignment_and_operator] = "_assignment_and_operator",
  [sym__equality_operator] = "_equality_operator",
  [sym__comparison_operator] = "_comparison_operator",
  [sym__in_operator] = "_in_operator",
  [sym__is_operator] = "_is_operator",
  [sym__additive_operator] = "_additive_operator",
  [sym__multiplicative_operator] = "_multiplicative_operator",
  [sym__as_operator] = "_as_operator",
  [sym__prefix_unary_operator] = "_prefix_unary_operator",
  [sym__postfix_unary_operator] = "_postfix_unary_operator",
  [sym__member_access_operator] = "_member_access_operator",
  [sym__postfix_unary_suffix] = "_postfix_unary_suffix",
  [sym__postfix_unary_expression] = "_postfix_unary_expression",
  [sym_directly_assignable_expression] = "directly_assignable_expression",
  [sym_modifiers] = "modifiers",
  [sym_parameter_modifiers] = "parameter_modifiers",
  [sym__modifier] = "_modifier",
  [sym_type_modifiers] = "type_modifiers",
  [sym__type_modifier] = "_type_modifier",
  [sym_class_modifier] = "class_modifier",
  [sym_member_modifier] = "member_modifier",
  [sym_visibility_modifier] = "visibility_modifier",
  [sym_variance_modifier] = "variance_modifier",
  [sym_type_parameter_modifiers] = "type_parameter_modifiers",
  [sym__type_parameter_modifier] = "_type_parameter_modifier",
  [sym_function_modifier] = "function_modifier",
  [sym_inheritance_modifier] = "inheritance_modifier",
  [sym_parameter_modifier] = "parameter_modifier",
  [sym_platform_modifier] = "platform_modifier",
  [sym_annotation] = "annotation",
  [sym__single_annotation] = "_single_annotation",
  [sym__multi_annotation] = "_multi_annotation",
  [sym_use_site_target] = "use_site_target",
  [sym__unescaped_annotation] = "_unescaped_annotation",
  [sym_simple_identifier] = "simple_identifier",
  [sym_identifier] = "identifier",
  [sym__import_identifier] = "_import_identifier",
  [sym__return_at] = "_return_at",
  [sym__continue_at] = "_continue_at",
  [sym__break_at] = "_break_at",
  [sym__this_at] = "_this_at",
  [sym__super_at] = "_super_at",
  [sym_unsigned_literal] = "unsigned_literal",
  [sym_long_literal] = "long_literal",
  [sym_boolean_literal] = "boolean_literal",
  [sym_character_literal] = "character_literal",
  [sym_character_escape_seq] = "character_escape_seq",
  [sym__lexical_identifier] = "_lexical_identifier",
  [sym__uni_character_literal] = "_uni_character_literal",
  [aux_sym_source_file_repeat1] = "source_file_repeat1",
  [aux_sym_source_file_repeat2] = "source_file_repeat2",
  [aux_sym_source_file_repeat3] = "source_file_repeat3",
  [aux_sym_file_annotation_repeat1] = "file_annotation_repeat1",
  [aux_sym_import_list_repeat1] = "import_list_repeat1",
  [aux_sym__class_parameters_repeat1] = "_class_parameters_repeat1",
  [aux_sym__delegation_specifiers_repeat1] = "_delegation_specifiers_repeat1",
  [aux_sym__annotated_delegation_specifier_repeat1] = "_annotated_delegation_specifier_repeat1",
  [aux_sym_type_parameters_repeat1] = "type_parameters_repeat1",
  [aux_sym_type_constraints_repeat1] = "type_constraints_repeat1",
  [aux_sym_function_value_parameters_repeat1] = "function_value_parameters_repeat1",
  [aux_sym__enum_entries_repeat1] = "_enum_entries_repeat1",
  [aux_sym_nullable_type_repeat1] = "nullable_type_repeat1",
  [aux_sym_user_type_repeat1] = "user_type_repeat1",
  [aux_sym_type_projection_modifiers_repeat1] = "type_projection_modifiers_repeat1",
  [aux_sym_function_type_parameters_repeat1] = "function_type_parameters_repeat1",
  [aux_sym_statements_repeat1] = "statements_repeat1",
  [aux_sym__statement_repeat1] = "_statement_repeat1",
  [aux_sym_indexing_suffix_repeat1] = "indexing_suffix_repeat1",
  [aux_sym_type_arguments_repeat1] = "type_arguments_repeat1",
  [aux_sym_value_arguments_repeat1] = "value_arguments_repeat1",
  [aux_sym_string_literal_repeat1] = "string_literal_repeat1",
  [aux_sym_multi_variable_declaration_repeat1] = "multi_variable_declaration_repeat1",
  [aux_sym_lambda_parameters_repeat1] = "lambda_parameters_repeat1",
  [aux_sym_when_expression_repeat1] = "when_expression_repeat1",
  [aux_sym_when_entry_repeat1] = "when_entry_repeat1",
  [aux_sym_try_expression_repeat1] = "try_expression_repeat1",
  [aux_sym__postfix_unary_expression_repeat1] = "_postfix_unary_expression_repeat1",
  [aux_sym_modifiers_repeat1] = "modifiers_repeat1",
  [aux_sym_parameter_modifiers_repeat1] = "parameter_modifiers_repeat1",
  [aux_sym_type_modifiers_repeat1] = "type_modifiers_repeat1",
  [aux_sym_type_parameter_modifiers_repeat1] = "type_parameter_modifiers_repeat1",
  [aux_sym_identifier_repeat1] = "identifier_repeat1",
  [alias_sym_interpolated_expression] = "interpolated_expression",
  [alias_sym_interpolated_identifier] = "interpolated_identifier",
  [alias_sym_type_identifier] = "type_identifier",
};

static const TSSymbol ts_symbol_map[] = {
  [ts_builtin_sym_end] = ts_builtin_sym_end,
  [sym__alpha_identifier] = sym__alpha_identifier,
  [anon_sym_POUND_BANG] = anon_sym_POUND_BANG,
  [aux_sym_shebang_line_token1] = aux_sym_shebang_line_token1,
  [anon_sym_AT] = anon_sym_AT,
  [anon_sym_file] = anon_sym_file,
  [anon_sym_COLON] = anon_sym_COLON,
  [anon_sym_LBRACK] = anon_sym_LBRACK,
  [anon_sym_RBRACK] = anon_sym_RBRACK,
  [anon_sym_package] = anon_sym_package,
  [anon_sym_import] = anon_sym_import,
  [anon_sym_DOT] = anon_sym_DOT,
  [sym_wildcard_import] = sym_wildcard_import,
  [anon_sym_as] = anon_sym_as,
  [anon_sym_typealias] = anon_sym_typealias,
  [anon_sym_EQ] = anon_sym_EQ,
  [anon_sym_class] = anon_sym_class,
  [anon_sym_interface] = anon_sym_interface,
  [anon_sym_enum] = anon_sym_enum,
  [anon_sym_constructor] = anon_sym_constructor,
  [anon_sym_LBRACE] = anon_sym_LBRACE,
  [anon_sym_RBRACE] = anon_sym_RBRACE,
  [anon_sym_LPAREN] = anon_sym_LPAREN,
  [anon_sym_COMMA] = anon_sym_COMMA,
  [anon_sym_RPAREN] = anon_sym_RPAREN,
  [anon_sym_val] = anon_sym_val,
  [anon_sym_var] = anon_sym_var,
  [anon_sym_by] = anon_sym_by,
  [anon_sym_LT] = anon_sym_LT,
  [anon_sym_GT] = anon_sym_GT,
  [anon_sym_where] = anon_sym_where,
  [anon_sym_init] = anon_sym_init,
  [anon_sym_companion] = anon_sym_companion,
  [anon_sym_object] = anon_sym_object,
  [anon_sym_fun] = anon_sym_fun,
  [anon_sym_SEMI] = anon_sym_SEMI,
  [anon_sym_get] = anon_sym_get,
  [anon_sym_set] = anon_sym_set,
  [anon_sym_this] = anon_sym_this,
  [anon_sym_super] = anon_sym_super,
  [anon_sym_dynamic] = anon_sym_dynamic,
  [anon_sym_AMP] = anon_sym_AMP,
  [sym__quest] = sym__quest,
  [anon_sym_STAR] = anon_sym_STAR,
  [anon_sym_DASH_GT] = anon_sym_DASH_GT,
  [sym_label] = sym_label,
  [anon_sym_for] = anon_sym_for,
  [anon_sym_in] = anon_sym_in,
  [anon_sym_while] = anon_sym_while,
  [anon_sym_do] = anon_sym_do,
  [anon_sym_DOT_DOT] = anon_sym_DOT_DOT,
  [anon_sym_QMARK_COLON] = anon_sym_QMARK_COLON,
  [anon_sym_AMP_AMP] = anon_sym_AMP_AMP,
  [anon_sym_PIPE_PIPE] = anon_sym_PIPE_PIPE,
  [anon_sym_null] = anon_sym_null,
  [anon_sym_DOLLAR_LBRACE] = anon_sym_DOLLAR_LBRACE,
  [anon_sym_DOLLAR] = anon_sym_DOLLAR,
  [anon_sym_if] = anon_sym_if,
  [anon_sym_else] = anon_sym_else,
  [anon_sym_when] = anon_sym_when,
  [anon_sym_try] = anon_sym_try,
  [anon_sym_catch] = anon_sym_catch,
  [anon_sym_finally] = anon_sym_finally,
  [anon_sym_throw] = anon_sym_throw,
  [anon_sym_return] = anon_sym_return,
  [anon_sym_continue] = anon_sym_continue,
  [anon_sym_break] = anon_sym_break,
  [anon_sym_COLON_COLON] = anon_sym_COLON_COLON,
  [anon_sym_PLUS_EQ] = anon_sym_PLUS_EQ,
  [anon_sym_DASH_EQ] = anon_sym_DASH_EQ,
  [anon_sym_STAR_EQ] = anon_sym_STAR_EQ,
  [anon_sym_SLASH_EQ] = anon_sym_SLASH_EQ,
  [anon_sym_PERCENT_EQ] = anon_sym_PERCENT_EQ,
  [anon_sym_BANG_EQ] = anon_sym_BANG_EQ,
  [anon_sym_BANG_EQ_EQ] = anon_sym_BANG_EQ_EQ,
  [anon_sym_EQ_EQ] = anon_sym_EQ_EQ,
  [anon_sym_EQ_EQ_EQ] = anon_sym_EQ_EQ_EQ,
  [anon_sym_LT_EQ] = anon_sym_LT_EQ,
  [anon_sym_GT_EQ] = anon_sym_GT_EQ,
  [anon_sym_BANGin] = anon_sym_BANGin,
  [anon_sym_is] = anon_sym_is,
  [anon_sym_BANGis] = anon_sym_BANGis,
  [anon_sym_PLUS] = anon_sym_PLUS,
  [anon_sym_DASH] = anon_sym_DASH,
  [anon_sym_SLASH] = anon_sym_SLASH,
  [anon_sym_PERCENT] = anon_sym_PERCENT,
  [anon_sym_as_QMARK] = anon_sym_as_QMARK,
  [anon_sym_PLUS_PLUS] = anon_sym_PLUS_PLUS,
  [anon_sym_DASH_DASH] = anon_sym_DASH_DASH,
  [anon_sym_BANG] = anon_sym_BANG,
  [anon_sym_BANG_BANG] = anon_sym_BANG_BANG,
  [anon_sym_suspend] = anon_sym_suspend,
  [anon_sym_sealed] = anon_sym_sealed,
  [anon_sym_annotation] = anon_sym_annotation,
  [anon_sym_data] = anon_sym_data,
  [anon_sym_inner] = anon_sym_inner,
  [anon_sym_value] = anon_sym_value,
  [anon_sym_override] = anon_sym_override,
  [anon_sym_lateinit] = anon_sym_lateinit,
  [anon_sym_public] = anon_sym_public,
  [anon_sym_private] = anon_sym_private,
  [anon_sym_internal] = anon_sym_internal,
  [anon_sym_protected] = anon_sym_protected,
  [anon_sym_out] = anon_sym_out,
  [anon_sym_tailrec] = anon_sym_tailrec,
  [anon_sym_operator] = anon_sym_operator,
  [anon_sym_infix] = anon_sym_infix,
  [anon_sym_inline] = anon_sym_inline,
  [anon_sym_external] = anon_sym_external,
  [sym_property_modifier] = sym_property_modifier,
  [anon_sym_abstract] = anon_sym_abstract,
  [anon_sym_final] = anon_sym_final,
  [anon_sym_open] = anon_sym_open,
  [anon_sym_vararg] = anon_sym_vararg,
  [anon_sym_noinline] = anon_sym_noinline,
  [anon_sym_crossinline] = anon_sym_crossinline,
  [sym_reification_modifier] = sym_reification_modifier,
  [anon_sym_expect] = anon_sym_expect,
  [anon_sym_actual] = anon_sym_actual,
  [anon_sym_field] = anon_sym_field,
  [anon_sym_property] = anon_sym_property,
  [anon_sym_receiver] = anon_sym_receiver,
  [anon_sym_param] = anon_sym_param,
  [anon_sym_setparam] = anon_sym_setparam,
  [anon_sym_delegate] = anon_sym_delegate,
  [sym_line_comment] = sym_line_comment,
  [anon_sym_return_AT] = anon_sym_return_AT,
  [anon_sym_continue_AT] = anon_sym_continue_AT,
  [anon_sym_break_AT] = anon_sym_break_AT,
  [anon_sym_this_AT] = anon_sym_this_AT,
  [anon_sym_super_AT] = anon_sym_super_AT,
  [anon_sym_AT2] = anon_sym_AT,
  [sym_real_literal] = sym_real_literal,
  [sym_integer_literal] = sym_integer_literal,
  [sym_hex_literal] = sym_hex_literal,
  [sym_bin_literal] = sym_bin_literal,
  [aux_sym_unsigned_literal_token1] = aux_sym_unsigned_literal_token1,
  [anon_sym_L] = anon_sym_L,
  [anon_sym_true] = anon_sym_true,
  [anon_sym_false] = anon_sym_false,
  [anon_sym_SQUOTE] = anon_sym_SQUOTE,
  [aux_sym_character_literal_token1] = aux_sym_character_literal_token1,
  [sym__backtick_identifier] = sym__backtick_identifier,
  [anon_sym_BSLASHu] = anon_sym_BSLASHu,
  [aux_sym__uni_character_literal_token1] = aux_sym__uni_character_literal_token1,
  [sym__escaped_identifier] = sym__escaped_identifier,
  [sym__automatic_semicolon] = sym__automatic_semicolon,
  [sym__import_list_delimiter] = sym__import_list_delimiter,
  [sym_safe_nav] = sym_safe_nav,
  [sym_multiline_comment] = sym_multiline_comment,
  [sym__string_start] = sym__string_start,
  [sym__string_end] = sym__string_end,
  [sym_string_content] = sym_string_content,
  [sym_source_file] = sym_source_file,
  [sym_shebang_line] = sym_shebang_line,
  [sym_file_annotation] = sym_file_annotation,
  [sym_package_header] = sym_package_header,
  [sym_import_list] = sym_import_list,
  [sym_import_header] = sym_import_header,
  [sym_import_alias] = sym_import_alias,
  [sym_type_alias] = sym_type_alias,
  [sym__declaration] = sym__declaration,
  [sym_class_declaration] = sym_class_declaration,
  [sym_primary_constructor] = sym_primary_constructor,
  [sym_class_body] = sym_class_body,
  [sym__class_parameters] = sym__class_parameters,
  [sym_binding_pattern_kind] = sym_binding_pattern_kind,
  [sym_class_parameter] = sym_class_parameter,
  [sym__delegation_specifiers] = sym__delegation_specifiers,
  [sym_delegation_specifier] = sym_delegation_specifier,
  [sym_constructor_invocation] = sym_constructor_invocation,
  [sym_explicit_delegation] = sym_explicit_delegation,
  [sym_type_parameters] = sym_type_parameters,
  [sym_type_parameter] = sym_type_parameter,
  [sym_type_constraints] = sym_type_constraints,
  [sym_type_constraint] = sym_type_constraint,
  [aux_sym__class_member_declarations] = aux_sym__class_member_declarations,
  [sym__class_member_declaration] = sym__class_member_declaration,
  [sym_anonymous_initializer] = sym_anonymous_initializer,
  [sym_companion_object] = sym_companion_object,
  [sym_function_value_parameters] = sym_function_value_parameters,
  [sym__function_value_parameter] = sym__function_value_parameter,
  [sym__receiver_type] = sym__receiver_type,
  [sym_function_declaration] = sym_function_declaration,
  [sym_function_body] = sym_function_body,
  [sym_variable_declaration] = sym_variable_declaration,
  [sym_property_declaration] = sym_property_declaration,
  [sym_property_delegate] = sym_property_delegate,
  [sym_getter] = sym_getter,
  [sym_setter] = sym_setter,
  [sym_parameter_with_optional_type] = sym_parameter_with_optional_type,
  [sym_parameter] = sym_parameter,
  [sym_object_declaration] = sym_object_declaration,
  [sym_secondary_constructor] = sym_secondary_constructor,
  [sym_constructor_delegation_call] = sym_constructor_delegation_call,
  [sym_enum_class_body] = sym_enum_class_body,
  [sym__enum_entries] = sym__enum_entries,
  [sym_enum_entry] = sym_enum_entry,
  [sym__type] = sym__type,
  [sym__type_reference] = sym__type_reference,
  [sym_not_nullable_type] = sym_not_nullable_type,
  [sym_nullable_type] = sym_nullable_type,
  [sym_user_type] = sym_user_type,
  [sym__simple_user_type] = sym__simple_user_type,
  [sym_type_projection] = sym_type_projection,
  [sym_type_projection_modifiers] = sym_type_projection_modifiers,
  [sym__type_projection_modifier] = sym__type_projection_modifier,
  [sym_function_type] = sym_function_type,
  [sym_function_type_parameters] = sym_function_type_parameters,
  [sym_parenthesized_type] = sym_parenthesized_type,
  [sym_parenthesized_user_type] = sym_parenthesized_user_type,
  [sym_statements] = sym_statements,
  [sym__statement] = sym__statement,
  [sym_control_structure_body] = sym_control_structure_body,
  [sym__block] = sym__block,
  [sym__loop_statement] = sym__loop_statement,
  [sym_for_statement] = sym_for_statement,
  [sym_while_statement] = sym_while_statement,
  [sym_do_while_statement] = sym_do_while_statement,
  [sym__semi] = sym__semi,
  [sym_assignment] = sym_assignment,
  [sym__expression] = sym__expression,
  [sym__unary_expression] = sym__unary_expression,
  [sym_postfix_expression] = sym_postfix_expression,
  [sym_call_expression] = sym_call_expression,
  [sym_indexing_expression] = sym_indexing_expression,
  [sym_navigation_expression] = sym_navigation_expression,
  [sym_prefix_expression] = sym_prefix_expression,
  [sym_as_expression] = sym_as_expression,
  [sym_spread_expression] = sym_spread_expression,
  [sym__binary_expression] = sym__binary_expression,
  [sym_multiplicative_expression] = sym_multiplicative_expression,
  [sym_additive_expression] = sym_additive_expression,
  [sym_range_expression] = sym_range_expression,
  [sym_infix_expression] = sym_infix_expression,
  [sym_elvis_expression] = sym_elvis_expression,
  [sym_check_expression] = sym_check_expression,
  [sym_comparison_expression] = sym_comparison_expression,
  [sym_equality_expression] = sym_equality_expression,
  [sym_conjunction_expression] = sym_conjunction_expression,
  [sym_disjunction_expression] = sym_disjunction_expression,
  [sym_indexing_suffix] = sym_indexing_suffix,
  [sym_navigation_suffix] = sym_navigation_suffix,
  [sym_call_suffix] = sym_call_suffix,
  [sym_annotated_lambda] = sym_annotated_lambda,
  [sym_type_arguments] = sym_type_arguments,
  [sym_value_arguments] = sym_value_arguments,
  [sym_value_argument] = sym_value_argument,
  [sym__primary_expression] = sym__primary_expression,
  [sym_parenthesized_expression] = sym_parenthesized_expression,
  [sym_collection_literal] = sym_collection_literal,
  [sym__literal_constant] = sym__literal_constant,
  [sym_string_literal] = sym_string_literal,
  [sym__interpolation] = sym__interpolation,
  [sym_lambda_literal] = sym_lambda_literal,
  [sym_multi_variable_declaration] = sym_multi_variable_declaration,
  [sym_lambda_parameters] = sym_lambda_parameters,
  [sym__lambda_parameter] = sym__lambda_parameter,
  [sym_anonymous_function] = sym_anonymous_function,
  [sym__function_literal] = sym__function_literal,
  [sym_object_literal] = sym_object_literal,
  [sym_this_expression] = sym_this_expression,
  [sym_super_expression] = sym_super_expression,
  [sym_if_expression] = sym_if_expression,
  [sym_when_subject] = sym_when_subject,
  [sym_when_expression] = sym_when_expression,
  [sym_when_entry] = sym_when_entry,
  [sym_when_condition] = sym_when_condition,
  [sym_range_test] = sym_range_test,
  [sym_type_test] = sym_type_test,
  [sym_try_expression] = sym_try_expression,
  [sym_catch_block] = sym_catch_block,
  [sym_finally_block] = sym_finally_block,
  [sym_jump_expression] = sym_jump_expression,
  [sym_callable_reference] = sym_callable_reference,
  [sym__assignment_and_operator] = sym__assignment_and_operator,
  [sym__equality_operator] = sym__equality_operator,
  [sym__comparison_operator] = sym__comparison_operator,
  [sym__in_operator] = sym__in_operator,
  [sym__is_operator] = sym__is_operator,
  [sym__additive_operator] = sym__additive_operator,
  [sym__multiplicative_operator] = sym__multiplicative_operator,
  [sym__as_operator] = sym__as_operator,
  [sym__prefix_unary_operator] = sym__prefix_unary_operator,
  [sym__postfix_unary_operator] = sym__postfix_unary_operator,
  [sym__member_access_operator] = sym__member_access_operator,
  [sym__postfix_unary_suffix] = sym__postfix_unary_suffix,
  [sym__postfix_unary_expression] = sym__postfix_unary_expression,
  [sym_directly_assignable_expression] = sym_directly_assignable_expression,
  [sym_modifiers] = sym_modifiers,
  [sym_parameter_modifiers] = sym_parameter_modifiers,
  [sym__modifier] = sym__modifier,
  [sym_type_modifiers] = sym_type_modifiers,
  [sym__type_modifier] = sym__type_modifier,
  [sym_class_modifier] = sym_class_modifier,
  [sym_member_modifier] = sym_member_modifier,
  [sym_visibility_modifier] = sym_visibility_modifier,
  [sym_variance_modifier] = sym_variance_modifier,
  [sym_type_parameter_modifiers] = sym_type_parameter_modifiers,
  [sym__type_parameter_modifier] = sym__type_parameter_modifier,
  [sym_function_modifier] = sym_function_modifier,
  [sym_inheritance_modifier] = sym_inheritance_modifier,
  [sym_parameter_modifier] = sym_parameter_modifier,
  [sym_platform_modifier] = sym_platform_modifier,
  [sym_annotation] = sym_annotation,
  [sym__single_annotation] = sym__single_annotation,
  [sym__multi_annotation] = sym__multi_annotation,
  [sym_use_site_target] = sym_use_site_target,
  [sym__unescaped_annotation] = sym__unescaped_annotation,
  [sym_simple_identifier] = sym_simple_identifier,
  [sym_identifier] = sym_identifier,
  [sym__import_identifier] = sym__import_identifier,
  [sym__return_at] = sym__return_at,
  [sym__continue_at] = sym__continue_at,
  [sym__break_at] = sym__break_at,
  [sym__this_at] = sym__this_at,
  [sym__super_at] = sym__super_at,
  [sym_unsigned_literal] = sym_unsigned_literal,
  [sym_long_literal] = sym_long_literal,
  [sym_boolean_literal] = sym_boolean_literal,
  [sym_character_literal] = sym_character_literal,
  [sym_character_escape_seq] = sym_character_escape_seq,
  [sym__lexical_identifier] = sym__lexical_identifier,
  [sym__uni_character_literal] = sym__uni_character_literal,
  [aux_sym_source_file_repeat1] = aux_sym_source_file_repeat1,
  [aux_sym_source_file_repeat2] = aux_sym_source_file_repeat2,
  [aux_sym_source_file_repeat3] = aux_sym_source_file_repeat3,
  [aux_sym_file_annotation_repeat1] = aux_sym_file_annotation_repeat1,
  [aux_sym_import_list_repeat1] = aux_sym_import_list_repeat1,
  [aux_sym__class_parameters_repeat1] = aux_sym__class_parameters_repeat1,
  [aux_sym__delegation_specifiers_repeat1] = aux_sym__delegation_specifiers_repeat1,
  [aux_sym__annotated_delegation_specifier_repeat1] = aux_sym__annotated_delegation_specifier_repeat1,
  [aux_sym_type_parameters_repeat1] = aux_sym_type_parameters_repeat1,
  [aux_sym_type_constraints_repeat1] = aux_sym_type_constraints_repeat1,
  [aux_sym_function_value_parameters_repeat1] = aux_sym_function_value_parameters_repeat1,
  [aux_sym__enum_entries_repeat1] = aux_sym__enum_entries_repeat1,
  [aux_sym_nullable_type_repeat1] = aux_sym_nullable_type_repeat1,
  [aux_sym_user_type_repeat1] = aux_sym_user_type_repeat1,
  [aux_sym_type_projection_modifiers_repeat1] = aux_sym_type_projection_modifiers_repeat1,
  [aux_sym_function_type_parameters_repeat1] = aux_sym_function_type_parameters_repeat1,
  [aux_sym_statements_repeat1] = aux_sym_statements_repeat1,
  [aux_sym__statement_repeat1] = aux_sym__statement_repeat1,
  [aux_sym_indexing_suffix_repeat1] = aux_sym_indexing_suffix_repeat1,
  [aux_sym_type_arguments_repeat1] = aux_sym_type_arguments_repeat1,
  [aux_sym_value_arguments_repeat1] = aux_sym_value_arguments_repeat1,
  [aux_sym_string_literal_repeat1] = aux_sym_string_literal_repeat1,
  [aux_sym_multi_variable_declaration_repeat1] = aux_sym_multi_variable_declaration_repeat1,
  [aux_sym_lambda_parameters_repeat1] = aux_sym_lambda_parameters_repeat1,
  [aux_sym_when_expression_repeat1] = aux_sym_when_expression_repeat1,
  [aux_sym_when_entry_repeat1] = aux_sym_when_entry_repeat1,
  [aux_sym_try_expression_repeat1] = aux_sym_try_expression_repeat1,
  [aux_sym__postfix_unary_expression_repeat1] = aux_sym__postfix_unary_expression_repeat1,
  [aux_sym_modifiers_repeat1] = aux_sym_modifiers_repeat1,
  [aux_sym_parameter_modifiers_repeat1] = aux_sym_parameter_modifiers_repeat1,
  [aux_sym_type_modifiers_repeat1] = aux_sym_type_modifiers_repeat1,
  [aux_sym_type_parameter_modifiers_repeat1] = aux_sym_type_parameter_modifiers_repeat1,
  [aux_sym_identifier_repeat1] = aux_sym_identifier_repeat1,
  [alias_sym_interpolated_expression] = alias_sym_interpolated_expression,
  [alias_sym_interpolated_identifier] = alias_sym_interpolated_identifier,
  [alias_sym_type_identifier] = alias_sym_type_identifier,
};

static const TSSymbolMetadata ts_symbol_metadata[] = {
  [ts_builtin_sym_end] = {
    .visible = false,
    .named = true,
  },
  [sym__alpha_identifier] = {
    .visible = false,
    .named = true,
  },
  [anon_sym_POUND_BANG] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_shebang_line_token1] = {
    .visible = false,
    .named = false,
  },
  [anon_sym_AT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_file] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_COLON] = {
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
  [anon_sym_package] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_import] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DOT] = {
    .visible = true,
    .named = false,
  },
  [sym_wildcard_import] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_as] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_typealias] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_class] = {
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
  [anon_sym_constructor] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LBRACE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_RBRACE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LPAREN] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_COMMA] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_RPAREN] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_val] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_var] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_by] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_GT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_where] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_init] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_companion] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_object] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_fun] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_SEMI] = {
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
  [anon_sym_this] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_super] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_dynamic] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_AMP] = {
    .visible = true,
    .named = false,
  },
  [sym__quest] = {
    .visible = false,
    .named = true,
  },
  [anon_sym_STAR] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DASH_GT] = {
    .visible = true,
    .named = false,
  },
  [sym_label] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_for] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_in] = {
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
  [anon_sym_DOT_DOT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_QMARK_COLON] = {
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
  [anon_sym_null] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DOLLAR_LBRACE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DOLLAR] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_if] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_else] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_when] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_try] = {
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
  [anon_sym_throw] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_return] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_continue] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_break] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_COLON_COLON] = {
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
  [anon_sym_BANG_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_BANG_EQ_EQ] = {
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
  [anon_sym_LT_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_GT_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_BANGin] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_is] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_BANGis] = {
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
  [anon_sym_as_QMARK] = {
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
  [anon_sym_BANG] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_BANG_BANG] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_suspend] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_sealed] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_annotation] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_data] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_inner] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_value] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_override] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_lateinit] = {
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
  [anon_sym_internal] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_protected] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_out] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_tailrec] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_operator] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_infix] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_inline] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_external] = {
    .visible = true,
    .named = false,
  },
  [sym_property_modifier] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_abstract] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_final] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_open] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_vararg] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_noinline] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_crossinline] = {
    .visible = true,
    .named = false,
  },
  [sym_reification_modifier] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_expect] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_actual] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_field] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_property] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_receiver] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_param] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_setparam] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_delegate] = {
    .visible = true,
    .named = false,
  },
  [sym_line_comment] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_return_AT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_continue_AT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_break_AT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_this_AT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_super_AT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_AT2] = {
    .visible = true,
    .named = false,
  },
  [sym_real_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_integer_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_hex_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_bin_literal] = {
    .visible = true,
    .named = true,
  },
  [aux_sym_unsigned_literal_token1] = {
    .visible = false,
    .named = false,
  },
  [anon_sym_L] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_true] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_false] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_SQUOTE] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_character_literal_token1] = {
    .visible = false,
    .named = false,
  },
  [sym__backtick_identifier] = {
    .visible = false,
    .named = true,
  },
  [anon_sym_BSLASHu] = {
    .visible = true,
    .named = false,
  },
  [aux_sym__uni_character_literal_token1] = {
    .visible = false,
    .named = false,
  },
  [sym__escaped_identifier] = {
    .visible = false,
    .named = true,
  },
  [sym__automatic_semicolon] = {
    .visible = false,
    .named = true,
  },
  [sym__import_list_delimiter] = {
    .visible = false,
    .named = true,
  },
  [sym_safe_nav] = {
    .visible = true,
    .named = false,
  },
  [sym_multiline_comment] = {
    .visible = true,
    .named = true,
  },
  [sym__string_start] = {
    .visible = false,
    .named = true,
  },
  [sym__string_end] = {
    .visible = false,
    .named = true,
  },
  [sym_string_content] = {
    .visible = true,
    .named = true,
  },
  [sym_source_file] = {
    .visible = true,
    .named = true,
  },
  [sym_shebang_line] = {
    .visible = true,
    .named = true,
  },
  [sym_file_annotation] = {
    .visible = true,
    .named = true,
  },
  [sym_package_header] = {
    .visible = true,
    .named = true,
  },
  [sym_import_list] = {
    .visible = true,
    .named = true,
  },
  [sym_import_header] = {
    .visible = true,
    .named = true,
  },
  [sym_import_alias] = {
    .visible = true,
    .named = true,
  },
  [sym_type_alias] = {
    .visible = true,
    .named = true,
  },
  [sym__declaration] = {
    .visible = false,
    .named = true,
  },
  [sym_class_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_primary_constructor] = {
    .visible = true,
    .named = true,
  },
  [sym_class_body] = {
    .visible = true,
    .named = true,
  },
  [sym__class_parameters] = {
    .visible = false,
    .named = true,
  },
  [sym_binding_pattern_kind] = {
    .visible = true,
    .named = true,
  },
  [sym_class_parameter] = {
    .visible = true,
    .named = true,
  },
  [sym__delegation_specifiers] = {
    .visible = false,
    .named = true,
  },
  [sym_delegation_specifier] = {
    .visible = true,
    .named = true,
  },
  [sym_constructor_invocation] = {
    .visible = true,
    .named = true,
  },
  [sym_explicit_delegation] = {
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
  [sym_type_constraints] = {
    .visible = true,
    .named = true,
  },
  [sym_type_constraint] = {
    .visible = true,
    .named = true,
  },
  [aux_sym__class_member_declarations] = {
    .visible = false,
    .named = false,
  },
  [sym__class_member_declaration] = {
    .visible = false,
    .named = true,
  },
  [sym_anonymous_initializer] = {
    .visible = true,
    .named = true,
  },
  [sym_companion_object] = {
    .visible = true,
    .named = true,
  },
  [sym_function_value_parameters] = {
    .visible = true,
    .named = true,
  },
  [sym__function_value_parameter] = {
    .visible = false,
    .named = true,
  },
  [sym__receiver_type] = {
    .visible = false,
    .named = true,
  },
  [sym_function_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_function_body] = {
    .visible = true,
    .named = true,
  },
  [sym_variable_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_property_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_property_delegate] = {
    .visible = true,
    .named = true,
  },
  [sym_getter] = {
    .visible = true,
    .named = true,
  },
  [sym_setter] = {
    .visible = true,
    .named = true,
  },
  [sym_parameter_with_optional_type] = {
    .visible = true,
    .named = true,
  },
  [sym_parameter] = {
    .visible = true,
    .named = true,
  },
  [sym_object_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_secondary_constructor] = {
    .visible = true,
    .named = true,
  },
  [sym_constructor_delegation_call] = {
    .visible = true,
    .named = true,
  },
  [sym_enum_class_body] = {
    .visible = true,
    .named = true,
  },
  [sym__enum_entries] = {
    .visible = false,
    .named = true,
  },
  [sym_enum_entry] = {
    .visible = true,
    .named = true,
  },
  [sym__type] = {
    .visible = false,
    .named = true,
  },
  [sym__type_reference] = {
    .visible = false,
    .named = true,
  },
  [sym_not_nullable_type] = {
    .visible = true,
    .named = true,
  },
  [sym_nullable_type] = {
    .visible = true,
    .named = true,
  },
  [sym_user_type] = {
    .visible = true,
    .named = true,
  },
  [sym__simple_user_type] = {
    .visible = false,
    .named = true,
  },
  [sym_type_projection] = {
    .visible = true,
    .named = true,
  },
  [sym_type_projection_modifiers] = {
    .visible = true,
    .named = true,
  },
  [sym__type_projection_modifier] = {
    .visible = false,
    .named = true,
  },
  [sym_function_type] = {
    .visible = true,
    .named = true,
  },
  [sym_function_type_parameters] = {
    .visible = true,
    .named = true,
  },
  [sym_parenthesized_type] = {
    .visible = true,
    .named = true,
  },
  [sym_parenthesized_user_type] = {
    .visible = true,
    .named = true,
  },
  [sym_statements] = {
    .visible = true,
    .named = true,
  },
  [sym__statement] = {
    .visible = false,
    .named = true,
  },
  [sym_control_structure_body] = {
    .visible = true,
    .named = true,
  },
  [sym__block] = {
    .visible = false,
    .named = true,
  },
  [sym__loop_statement] = {
    .visible = false,
    .named = true,
  },
  [sym_for_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_while_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_do_while_statement] = {
    .visible = true,
    .named = true,
  },
  [sym__semi] = {
    .visible = false,
    .named = true,
  },
  [sym_assignment] = {
    .visible = true,
    .named = true,
  },
  [sym__expression] = {
    .visible = false,
    .named = true,
  },
  [sym__unary_expression] = {
    .visible = false,
    .named = true,
  },
  [sym_postfix_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_call_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_indexing_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_navigation_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_prefix_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_as_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_spread_expression] = {
    .visible = true,
    .named = true,
  },
  [sym__binary_expression] = {
    .visible = false,
    .named = true,
  },
  [sym_multiplicative_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_additive_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_range_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_infix_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_elvis_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_check_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_comparison_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_equality_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_conjunction_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_disjunction_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_indexing_suffix] = {
    .visible = true,
    .named = true,
  },
  [sym_navigation_suffix] = {
    .visible = true,
    .named = true,
  },
  [sym_call_suffix] = {
    .visible = true,
    .named = true,
  },
  [sym_annotated_lambda] = {
    .visible = true,
    .named = true,
  },
  [sym_type_arguments] = {
    .visible = true,
    .named = true,
  },
  [sym_value_arguments] = {
    .visible = true,
    .named = true,
  },
  [sym_value_argument] = {
    .visible = true,
    .named = true,
  },
  [sym__primary_expression] = {
    .visible = false,
    .named = true,
  },
  [sym_parenthesized_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_collection_literal] = {
    .visible = true,
    .named = true,
  },
  [sym__literal_constant] = {
    .visible = false,
    .named = true,
  },
  [sym_string_literal] = {
    .visible = true,
    .named = true,
  },
  [sym__interpolation] = {
    .visible = false,
    .named = true,
  },
  [sym_lambda_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_multi_variable_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_lambda_parameters] = {
    .visible = true,
    .named = true,
  },
  [sym__lambda_parameter] = {
    .visible = false,
    .named = true,
  },
  [sym_anonymous_function] = {
    .visible = true,
    .named = true,
  },
  [sym__function_literal] = {
    .visible = false,
    .named = true,
  },
  [sym_object_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_this_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_super_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_if_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_when_subject] = {
    .visible = true,
    .named = true,
  },
  [sym_when_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_when_entry] = {
    .visible = true,
    .named = true,
  },
  [sym_when_condition] = {
    .visible = true,
    .named = true,
  },
  [sym_range_test] = {
    .visible = true,
    .named = true,
  },
  [sym_type_test] = {
    .visible = true,
    .named = true,
  },
  [sym_try_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_catch_block] = {
    .visible = true,
    .named = true,
  },
  [sym_finally_block] = {
    .visible = true,
    .named = true,
  },
  [sym_jump_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_callable_reference] = {
    .visible = true,
    .named = true,
  },
  [sym__assignment_and_operator] = {
    .visible = false,
    .named = true,
  },
  [sym__equality_operator] = {
    .visible = false,
    .named = true,
  },
  [sym__comparison_operator] = {
    .visible = false,
    .named = true,
  },
  [sym__in_operator] = {
    .visible = false,
    .named = true,
  },
  [sym__is_operator] = {
    .visible = false,
    .named = true,
  },
  [sym__additive_operator] = {
    .visible = false,
    .named = true,
  },
  [sym__multiplicative_operator] = {
    .visible = false,
    .named = true,
  },
  [sym__as_operator] = {
    .visible = false,
    .named = true,
  },
  [sym__prefix_unary_operator] = {
    .visible = false,
    .named = true,
  },
  [sym__postfix_unary_operator] = {
    .visible = false,
    .named = true,
  },
  [sym__member_access_operator] = {
    .visible = false,
    .named = true,
  },
  [sym__postfix_unary_suffix] = {
    .visible = false,
    .named = true,
  },
  [sym__postfix_unary_expression] = {
    .visible = false,
    .named = true,
  },
  [sym_directly_assignable_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_modifiers] = {
    .visible = true,
    .named = true,
  },
  [sym_parameter_modifiers] = {
    .visible = true,
    .named = true,
  },
  [sym__modifier] = {
    .visible = false,
    .named = true,
  },
  [sym_type_modifiers] = {
    .visible = true,
    .named = true,
  },
  [sym__type_modifier] = {
    .visible = false,
    .named = true,
  },
  [sym_class_modifier] = {
    .visible = true,
    .named = true,
  },
  [sym_member_modifier] = {
    .visible = true,
    .named = true,
  },
  [sym_visibility_modifier] = {
    .visible = true,
    .named = true,
  },
  [sym_variance_modifier] = {
    .visible = true,
    .named = true,
  },
  [sym_type_parameter_modifiers] = {
    .visible = true,
    .named = true,
  },
  [sym__type_parameter_modifier] = {
    .visible = false,
    .named = true,
  },
  [sym_function_modifier] = {
    .visible = true,
    .named = true,
  },
  [sym_inheritance_modifier] = {
    .visible = true,
    .named = true,
  },
  [sym_parameter_modifier] = {
    .visible = true,
    .named = true,
  },
  [sym_platform_modifier] = {
    .visible = true,
    .named = true,
  },
  [sym_annotation] = {
    .visible = true,
    .named = true,
  },
  [sym__single_annotation] = {
    .visible = false,
    .named = true,
  },
  [sym__multi_annotation] = {
    .visible = false,
    .named = true,
  },
  [sym_use_site_target] = {
    .visible = true,
    .named = true,
  },
  [sym__unescaped_annotation] = {
    .visible = false,
    .named = true,
  },
  [sym_simple_identifier] = {
    .visible = true,
    .named = true,
  },
  [sym_identifier] = {
    .visible = true,
    .named = true,
  },
  [sym__import_identifier] = {
    .visible = false,
    .named = true,
  },
  [sym__return_at] = {
    .visible = false,
    .named = true,
  },
  [sym__continue_at] = {
    .visible = false,
    .named = true,
  },
  [sym__break_at] = {
    .visible = false,
    .named = true,
  },
  [sym__this_at] = {
    .visible = false,
    .named = true,
  },
  [sym__super_at] = {
    .visible = false,
    .named = true,
  },
  [sym_unsigned_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_long_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_boolean_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_character_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_character_escape_seq] = {
    .visible = true,
    .named = true,
  },
  [sym__lexical_identifier] = {
    .visible = false,
    .named = true,
  },
  [sym__uni_character_literal] = {
    .visible = false,
    .named = true,
  },
  [aux_sym_source_file_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_source_file_repeat2] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_source_file_repeat3] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_file_annotation_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_import_list_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__class_parameters_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__delegation_specifiers_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__annotated_delegation_specifier_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_type_parameters_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_type_constraints_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_function_value_parameters_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__enum_entries_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_nullable_type_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_user_type_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_type_projection_modifiers_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_function_type_parameters_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_statements_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__statement_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_indexing_suffix_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_type_arguments_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_value_arguments_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_string_literal_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_multi_variable_declaration_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_lambda_parameters_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_when_expression_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_when_entry_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_try_expression_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__postfix_unary_expression_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_modifiers_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_parameter_modifiers_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_type_modifiers_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_type_parameter_modifiers_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_identifier_repeat1] = {
    .visible = false,
    .named = false,
  },
  [alias_sym_interpolated_expression] = {
    .visible = true,
    .named = true,
  },
  [alias_sym_interpolated_identifier] = {
    .visible = true,
    .named = true,
  },
  [alias_sym_type_identifier] = {
    .visible = true,
    .named = true,
  },
};

static const TSSymbol ts_alias_sequences[PRODUCTION_ID_COUNT][MAX_ALIAS_SEQUENCE_LENGTH] = {
  [0] = {0},
  [1] = {
    [0] = alias_sym_type_identifier,
  },
  [2] = {
    [1] = alias_sym_type_identifier,
  },
  [3] = {
    [1] = sym_label,
  },
  [4] = {
    [1] = alias_sym_interpolated_identifier,
  },
  [5] = {
    [1] = sym_identifier,
  },
  [6] = {
    [2] = alias_sym_type_identifier,
  },
  [7] = {
    [1] = alias_sym_interpolated_expression,
  },
  [8] = {
    [3] = alias_sym_type_identifier,
  },
  [9] = {
    [5] = alias_sym_type_identifier,
  },
};

static const uint16_t ts_non_terminal_alias_map[] = {
  sym__expression, 2,
    sym__expression,
    alias_sym_interpolated_expression,
  sym_simple_identifier, 3,
    sym_simple_identifier,
    alias_sym_interpolated_identifier,
    alias_sym_type_identifier,
  sym__import_identifier, 2,
    sym__import_identifier,
    sym_identifier,
  sym__lexical_identifier, 3,
    sym__lexical_identifier,
    alias_sym_type_identifier,
    sym_label,
  0,
};

// Filled from tables.zlib by grammar.go, see fetchgrammar.
static TSStateId ts_primary_state_ids[STATE_COUNT];

static TSCharacterRange sym__alpha_identifier_character_set_1[] = {
  {'A', 'Z'}, {'_', '_'}, {'a', 'z'}, {0xaa, 0xaa}, {0xb5, 0xb5}, {0xba, 0xba}, {0xc0, 0xd6}, {0xd8, 0xf6},
  {0xf8, 0x2c1}, {0x2c6, 0x2d1}, {0x2e0, 0x2e4}, {0x2ec, 0x2ec}, {0x2ee, 0x2ee}, {0x370, 0x374}, {0x376, 0x377}, {0x37a, 0x37d},
  {0x37f, 0x37f}, {0x386, 0x386}, {0x388, 0x38a}, {0x38c, 0x38c}, {0x38e, 0x3a1}, {0x3a3, 0x3f5}, {0x3f7, 0x481}, {0x48a, 0x52f},
  {0x531, 0x556}, {0x559, 0x559}, {0x560, 0x588}, {0x5d0, 0x5ea}, {0x5ef, 0x5f2}, {0x620, 0x64a}, {0x66e, 0x66f}, {0x671, 0x6d3},
  {0x6d5, 0x6d5}, {0x6e5, 0x6e6}, {0x6ee, 0x6ef}, {0x6fa, 0x6fc}, {0x6ff, 0x6ff}, {0x710, 0x710}, {0x712, 0x72f}, {0x74d, 0x7a5},
  {0x7b1, 0x7b1}, {0x7ca, 0x7ea}, {0x7f4, 0x7f5}, {0x7fa, 0x7fa}, {0x800, 0x815}, {0x81a, 0x81a}, {0x824, 0x824}, {0x828, 0x828},
  {0x840, 0x858}, {0x860, 0x86a}, {0x870, 0x887}, {0x889, 0x88e}, {0x8a0, 0x8c9}, {0x904, 0x939}, {0x93d, 0x93d}, {0x950, 0x950},
  {0x958, 0x961}, {0x971, 0x980}, {0x985, 0x98c}, {0x98f, 0x990}, {0x993, 0x9a8}, {0x9aa, 0x9b0}, {0x9b2, 0x9b2}, {0x9b6, 0x9b9},
  {0x9bd, 0x9bd}, {0x9ce, 0x9ce}, {0x9dc, 0x9dd}, {0x9df, 0x9e1}, {0x9f0, 0x9f1}, {0x9fc, 0x9fc}, {0xa05, 0xa0a}, {0xa0f, 0xa10},
  {0xa13, 0xa28}, {0xa2a, 0xa30}, {0xa32, 0xa33}, {0xa35, 0xa36}, {0xa38, 0xa39}, {0xa59, 0xa5c}, {0xa5e, 0xa5e}, {0xa72, 0xa74},
  {0xa85, 0xa8d}, {0xa8f, 0xa91}, {0xa93, 0xaa8}, {0xaaa, 0xab0}, {0xab2, 0xab3}, {0xab5, 0xab9}, {0xabd, 0xabd}, {0xad0, 0xad0},
  {0xae0, 0xae1}, {0xaf9, 0xaf9}, {0xb05, 0xb0c}, {0xb0f, 0xb10}, {0xb13, 0xb28}, {0xb2a, 0xb30}, {0xb32, 0xb33}, {0xb35, 0xb39},
  {0xb3d, 0xb3d}, {0xb5c, 0xb5d}, {0xb5f, 0xb61}, {0xb71, 0xb71}, {0xb83, 0xb83}, {0xb85, 0xb8a}, {0xb8e, 0xb90}, {0xb92, 0xb95},
  {0xb99, 0xb9a}, {0xb9c, 0xb9c}, {0xb9e, 0xb9f}, {0xba3, 0xba4}, {0xba8, 0xbaa}, {0xbae, 0xbb9}, {0xbd0, 0xbd0}, {0xc05, 0xc0c},
  {0xc0e, 0xc10}, {0xc12, 0xc28}, {0xc2a, 0xc39}, {0xc3d, 0xc3d}, {0xc58, 0xc5a}, {0xc5d, 0xc5d}, {0xc60, 0xc61}, {0xc80, 0xc80},
  {0xc85, 0xc8c}, {0xc8e, 0xc90}, {0xc92, 0xca8}, {0xcaa, 0xcb3}, {0xcb5, 0xcb9}, {0xcbd, 0xcbd}, {0xcdd, 0xcde}, {0xce0, 0xce1},
  {0xcf1, 0xcf2}, {0xd04, 0xd0c}, {0xd0e, 0xd10}, {0xd12, 0xd3a}, {0xd3d, 0xd3d}, {0xd4e, 0xd4e}, {0xd54, 0xd56}, {0xd5f, 0xd61},
  {0xd7a, 0xd7f}, {0xd85, 0xd96}, {0xd9a, 0xdb1}, {0xdb3, 0xdbb}, {0xdbd, 0xdbd}, {0xdc0, 0xdc6}, {0xe01, 0xe30}, {0xe32, 0xe33},
  {0xe40, 0xe46}, {0xe81, 0xe82}, {0xe84, 0xe84}, {0xe86, 0xe8a}, {0xe8c, 0xea3}, {0xea5, 0xea5}, {0xea7, 0xeb0}, {0xeb2, 0xeb3},
  {0xebd, 0xebd}, {0xec0, 0xec4}, {0xec6, 0xec6}, {0xedc, 0xedf}, {0xf00, 0xf00}, {0xf40, 0xf47}, {0xf49, 0xf6c}, {0xf88, 0xf8c},
  {0x1000, 0x102a}, {0x103f, 0x103f}, {0x1050, 0x1055}, {0x105a, 0x105d}, {0x1061, 0x1061}, {0x1065, 0x1066}, {0x106e, 0x1070}, {0x1075, 0x1081},
  {0x108e, 0x108e}, {0x10a0, 0x10c5}, {0x10c7, 0x10c7}, {0x10cd, 0x10cd}, {0x10d0, 0x10fa}, {0x10fc, 0x1248}, {0x124a, 0x124d}, {0x1250, 0x1256},
  {0x1258, 0x1258}, {0x125a, 0x125d}, {0x1260, 0x1288}, {0x128a, 0x128d}, {0x1290, 0x12b0}, {0x12b2, 0x12b5}, {0x12b8, 0x12be}, {0x12c0, 0x12c0},
  {0x12c2, 0x12c5}, {0x12c8, 0x12d6}, {0x12d8, 0x1310}, {0x1312, 0x1315}, {0x1318, 0x135a}, {0x1380, 0x138f}, {0x13a0, 0x13f5}, {0x13f8, 0x13fd},
  {0x1401, 0x166c}, {0x166f, 0x167f}, {0x1681, 0x169a}, {0x16a0, 0x16ea}, {0x16f1, 0x16f8}, {0x1700, 0x1711}, {0x171f, 0x1731}, {0x1740, 0x1751},
  {0x1760, 0x176c}, {0x176e, 0x1770}, {0x1780, 0x17b3}, {0x17d7, 0x17d7}, {0x17dc, 0x17dc}, {0x1820, 0x1878}, {0x1880, 0x1884}, {0x1887, 0x18a8},
  {0x18aa, 0x18aa}, {0x18b0, 0x18f5}, {0x1900, 0x191e}, {0x1950, 0x196d}, {0x1970, 0x1974}, {0x1980, 0x19ab}, {0x19b0, 0x19c9}, {0x1a00, 0x1a16},
  {0x1a20, 0x1a54}, {0x1aa7, 0x1aa7}, {0x1b05, 0x1b33}, {0x1b45, 0x1b4c}, {0x1b83, 0x1ba0}, {0x1bae, 0x1baf}, {0x1bba, 0x1be5}, {0x1c00, 0x1c23},
  {0x1c4d, 0x1c4f}, {0x1c5a, 0x1c7d}, {0x1c80, 0x1c88}, {0x1c90, 0x1cba}, {0x1cbd, 0x1cbf}, {0x1ce9, 0x1cec}, {0x1cee, 0x1cf3}, {0x1cf5, 0x1cf6},
  {0x1cfa, 0x1cfa}, {0x1d00, 0x1dbf}, {0x1e00, 0x1f15}, {0x1f18, 0x1f1d}, {0x1f20, 0x1f45}, {0x1f48, 0x1f4d}, {0x1f50, 0x1f57}, {0x1f59, 0x1f59},
  {0x1f5b, 0x1f5b}, {0x1f5d, 0x1f5d}, {0x1f5f, 0x1f7d}, {0x1f80, 0x1fb4}, {0x1fb6, 0x1fbc}, {0x1fbe, 0x1fbe}, {0x1fc2, 0x1fc4}, {0x1fc6, 0x1fcc},
  {0x1fd0, 0x1fd3}, {0x1fd6, 0x1fdb}, {0x1fe0, 0x1fec}, {0x1ff2, 0x1ff4}, {0x1ff6, 0x1ffc}, {0x2071, 0x2071}, {0x207f, 0x207f}, {0x2090, 0x209c},
  {0x2102, 0x2102}, {0x2107, 0x2107}, {0x210a, 0x2113}, {0x2115, 0x2115}, {0x2119, 0x211d}, {0x2124, 0x2124}, {0x2126, 0x2126}, {0x2128, 0x2128},
  {0x212a, 0x212d}, {0x212f, 0x2139}, {0x213c, 0x213f}, {0x2145, 0x2149}, {0x214e, 0x214e}, {0x2183, 0x2184}, {0x2c00, 0x2ce4}, {0x2ceb, 0x2cee},
  {0x2cf2, 0x2cf3}, {0x2d00, 0x2d25}, {0x2d27, 0x2d27}, {0x2d2d, 0x2d2d}, {0x2d30, 0x2d67}, {0x2d6f, 0x2d6f}, {0x2d80, 0x2d96}, {0x2da0, 0x2da6},
  {0x2da8, 0x2dae}, {0x2db0, 0x2db6}, {0x2db8, 0x2dbe}, {0x2dc0, 0x2dc6}, {0x2dc8, 0x2dce}, {0x2dd0, 0x2dd6}, {0x2dd8, 0x2dde}, {0x2e2f, 0x2e2f},
  {0x3005, 0x3006}, {0x3031, 0x3035}, {0x303b, 0x303c}, {0x3041, 0x3096}, {0x309d, 0x309f}, {0x30a1, 0x30fa}, {0x30fc, 0x30ff}, {0x3105, 0x312f},
  {0x3131, 0x318e}, {0x31a0, 0x31bf}, {0x31f0, 0x31ff}, {0x3400, 0x3400}, {0x4dbf, 0x4dbf}, {0x4e00, 0x4e00}, {0x9fff, 0xa48c}, {0xa4d0, 0xa4fd},
  {0xa500, 0xa60c}, {0xa610, 0xa61f}, {0xa62a, 0xa62b}, {0xa640, 0xa66e}, {0xa67f, 0xa69d}, {0xa6a0, 0xa6e5}, {0xa717, 0xa71f}, {0xa722, 0xa788},
  {0xa78b, 0xa7ca}, {0xa7d0, 0xa7d1}, {0xa7d3, 0xa7d3}, {0xa7d5, 0xa7d9}, {0xa7f2, 0xa801}, {0xa803, 0xa805}, {0xa807, 0xa80a}, {0xa80c, 0xa822},
  {0xa840, 0xa873}, {0xa882, 0xa8b3}, {0xa8f2, 0xa8f7}, {0xa8fb, 0xa8fb}, {0xa8fd, 0xa8fe}, {0xa90a, 0xa925}, {0xa930, 0xa946}, {0xa960, 0xa97c},
  {0xa984, 0xa9b2}, {0xa9cf, 0xa9cf}, {0xa9e0, 0xa9e4}, {0xa9e6, 0xa9ef}, {0xa9fa, 0xa9fe}, {0xaa00, 0xaa28}, {0xaa40, 0xaa42}, {0xaa44, 0xaa4b},
  {0xaa60, 0xaa76}, {0xaa7a, 0xaa7a}, {0xaa7e, 0xaaaf}, {0xaab1, 0xaab1}, {0xaab5, 0xaab6}, {0xaab9, 0xaabd}, {0xaac0, 0xaac0}, {0xaac2, 0xaac2},
  {0xaadb, 0xaadd}, {0xaae0, 0xaaea}, {0xaaf2, 0xaaf4}, {0xab01, 0xab06}, {0xab09, 0xab0e}, {0xab11, 0xab16}, {0xab20, 0xab26}, {0xab28, 0xab2e},
  {0xab30, 0xab5a}, {0xab5c, 0xab69}, {0xab70, 0xabe2}, {0xac00, 0xac00}, {0xd7a3, 0xd7a3}, {0xd7b0, 0xd7c6}, {0xd7cb, 0xd7fb}, {0xf900, 0xfa6d},
  {0xfa70, 0xfad9}, {0xfb00, 0xfb06}, {0xfb13, 0xfb17}, {0xfb1d, 0xfb1d}, {0xfb1f, 0xfb28}, {0xfb2a, 0xfb36}, {0xfb38, 0xfb3c}, {0xfb3e, 0xfb3e},
  {0xfb40, 0xfb41}, {0xfb43, 0xfb44}, {0xfb46, 0xfbb1}, {0xfbd3, 0xfd3d}, {0xfd50, 0xfd8f}, {0xfd92, 0xfdc7}, {0xfdf0, 0xfdfb}, {0xfe70, 0xfe74},
  {0xfe76, 0xfefc}, {0xff21, 0xff3a}, {0xff41, 0xff5a}, {0xff66, 0xffbe}, {0xffc2, 0xffc7}, {0xffca, 0xffcf}, {0xffd2, 0xffd7}, {0xffda, 0xffdc},
  {0x10000, 0x1000b}, {0x1000d, 0x10026}, {0x10028, 0x1003a}, {0x1003c, 0x1003d}, {0x1003f, 0x1004d}, {0x10050, 0x1005d}, {0x10080, 0x100fa}, {0x10280, 0x1029c},
  {0x102a0, 0x102d0}, {0x10300, 0x1031f}, {0x1032d, 0x10340}, {0x10342, 0x10349}, {0x10350, 0x10375}, {0x10380, 0x1039d}, {0x103a0, 0x103c3}, {0x103c8, 0x103cf},
  {0x10400, 0x1049d}, {0x104b0, 0x104d3}, {0x104d8, 0x104fb}, {0x10500, 0x10527}, {0x10530, 0x10563}, {0x10570, 0x1057a}, {0x1057c, 0x1058a}, {0x1058c, 0x10592},
  {0x10594, 0x10595}, {0x10597, 0x105a1}, {0x105a3, 0x105b1}, {0x105b3, 0x105b9}, {0x105bb, 0x105bc}, {0x10600, 0x10736}, {0x10740, 0x10755}, {0x10760, 0x10767},
  {0x10780, 0x10785}, {0x10787, 0x107b0}, {0x107b2, 0x107ba}, {0x10800, 0x10805}, {0x10808, 0x10808}, {0x1080a, 0x10835}, {0x10837, 0x10838}, {0x1083c, 0x1083c},
  {0x1083f, 0x10855}, {0x10860, 0x10876}, {0x10880, 0x1089e}, {0x108e0, 0x108f2}, {0x108f4, 0x108f5}, {0x10900, 0x10915}, {0x10920, 0x1092b},
};

static TSCharacterRange sym__alpha_identifier_character_set_2[] = {
  {'0', '9'}, {'A', 'Z'}, {'_', '_'}, {'a', 'z'}, {0xaa, 0xaa}, {0xb5, 0xb5}, {0xba, 0xba}, {0xc0, 0xd6},
  {0xd8, 0xf6}, {0xf8, 0x2c1}, {0x2c6, 0x2d1}, {0x2e0, 0x2e4}, {0x2ec, 0x2ec}, {0x2ee, 0x2ee}, {0x370, 0x374}, {0x376, 0x377},
  {0x37a, 0x37d}, {0x37f, 0x37f}, {0x386, 0x386}, {0x388, 0x38a}, {0x38c, 0x38c}, {0x38e, 0x3a1}, {0x3a3, 0x3f5}, {0x3f7, 0x481},
  {0x48a, 0x52f}, {0x531, 0x556}, {0x559, 0x559}, {0x560, 0x588}, {0x5d0, 0x5ea}, {0x5ef, 0x5f2}, {0x620, 0x64a}, {0x660, 0x669},
  {0x66e, 0x66f}, {0x671, 0x6d3}, {0x6d5, 0x6d5}, {0x6e5, 0x6e6}, {0x6ee, 0x6fc}, {0x6ff, 0x6ff}, {0x710, 0x710}, {0x712, 0x72f},
  {0x74d, 0x7a5}, {0x7b1, 0x7b1}, {0x7c0, 0x7ea}, {0x7f4, 0x7f5}, {0x7fa, 0x7fa}, {0x800, 0x815}, {0x81a, 0x81a}, {0x824, 0x824},
  {0x828, 0x828}, {0x840, 0x858}, {0x860, 0x86a}, {0x870, 0x887}, {0x889, 0x88e}, {0x8a0, 0x8c9}, {0x904, 0x939}, {0x93d, 0x93d},
  {0x950, 0x950}, {0x958, 0x961}, {0x966, 0x96f}, {0x971, 0x980}, {0x985, 0x98c}, {0x98f, 0x990}, {0x993, 0x9a8}, {0x9aa, 0x9b0},
  {0x9b2, 0x9b2}, {0x9b6, 0x9b9}, {0x9bd, 0x9bd}, {0x9ce, 0x9ce}, {0x9dc, 0x9dd}, {0x9df, 0x9e1}, {0x9e6, 0x9f1}, {0x9fc, 0x9fc},
  {0xa05, 0xa0a}, {0xa0f, 0xa10}, {0xa13, 0xa28}, {0xa2a, 0xa30}, {0xa32, 0xa33}, {0xa35, 0xa36}, {0xa38, 0xa39}, {0xa59, 0xa5c},
  {0xa5e, 0xa5e}, {0xa66, 0xa6f}, {0xa72, 0xa74}, {0xa85, 0xa8d}, {0xa8f, 0xa91}, {0xa93, 0xaa8}, {0xaaa, 0xab0}, {0xab2, 0xab3},
  {0xab5, 0xab9}, {0xabd, 0xabd}, {0xad0, 0xad0}, {0xae0, 0xae1}, {0xae6, 0xaef}, {0xaf9, 0xaf9}, {0xb05, 0xb0c}, {0xb0f, 0xb10},
  {0xb13, 0xb28}, {0xb2a, 0xb30}, {0xb32, 0xb33}, {0xb35, 0xb39}, {0xb3d, 0xb3d}, {0xb5c, 0xb5d}, {0xb5f, 0xb61}, {0xb66, 0xb6f},
  {0xb71, 0xb71}, {0xb83, 0xb83}, {0xb85, 0xb8a}, {0xb8e, 0xb90}, {0xb92, 0xb95}, {0xb99, 0xb9a}, {0xb9c, 0xb9c}, {0xb9e, 0xb9f},
  {0xba3, 0xba4}, {0xba8, 0xbaa}, {0xbae, 0xbb9}, {0xbd0, 0xbd0}, {0xbe6, 0xbef}, {0xc05, 0xc0c}, {0xc0e, 0xc10}, {0xc12, 0xc28},
  {0xc2a, 0xc39}, {0xc3d, 0xc3d}, {0xc58, 0xc5a}, {0xc5d, 0xc5d}, {0xc60, 0xc61}, {0xc66, 0xc6f}, {0xc80, 0xc80}, {0xc85, 0xc8c},
  {0xc8e, 0xc90}, {0xc92, 0xca8}, {0xcaa, 0xcb3}, {0xcb5, 0xcb9}, {0xcbd, 0xcbd}, {0xcdd, 0xcde}, {0xce0, 0xce1}, {0xce6, 0xcef},
  {0xcf1, 0xcf2}, {0xd04, 0xd0c}, {0xd0e, 0xd10}, {0xd12, 0xd3a}, {0xd3d, 0xd3d}, {0xd4e, 0xd4e}, {0xd54, 0xd56}, {0xd5f, 0xd61},
  {0xd66, 0xd6f}, {0xd7a, 0xd7f}, {0xd85, 0xd96}, {0xd9a, 0xdb1}, {0xdb3, 0xdbb}, {0xdbd, 0xdbd}, {0xdc0, 0xdc6}, {0xde6, 0xdef},
  {0xe01, 0xe30}, {0xe32, 0xe33}, {0xe40, 0xe46}, {0xe50, 0xe59}, {0xe81, 0xe82}, {0xe84, 0xe84}, {0xe86, 0xe8a}, {0xe8c, 0xea3},
  {0xea5, 0xea5}, {0xea7, 0xeb0}, {0xeb2, 0xeb3}, {0xebd, 0xebd}, {0xec0, 0xec4}, {0xec6, 0xec6}, {0xed0, 0xed9}, {0xedc, 0xedf},
  {0xf00, 0xf00}, {0xf20, 0xf29}, {0xf40, 0xf47}, {0xf49, 0xf6c}, {0xf88, 0xf8c}, {0x1000, 0x102a}, {0x103f, 0x1049}, {0x1050, 0x1055},
  {0x105a, 0x105d}, {0x1061, 0x1061}, {0x1065, 0x1066}, {0x106e, 0x1070}, {0x1075, 0x1081}, {0x108e, 0x108e}, {0x1090, 0x1099}, {0x10a0, 0x10c5},
  {0x10c7, 0x10c7}, {0x10cd, 0x10cd}, {0x10d0, 0x10fa}, {0x10fc, 0x1248}, {0x124a, 0x124d}, {0x1250, 0x1256}, {0x1258, 0x1258}, {0x125a, 0x125d},
  {0x1260, 0x1288}, {0x128a, 0x128d}, {0x1290, 0x12b0}, {0x12b2, 0x12b5}, {0x12b8, 0x12be}, {0x12c0, 0x12c0}, {0x12c2, 0x12c5}, {0x12c8, 0x12d6},
  {0x12d8, 0x1310}, {0x1312, 0x1315}, {0x1318, 0x135a}, {0x1380, 0x138f}, {0x13a0, 0x13f5}, {0x13f8, 0x13fd}, {0x1401, 0x166c}, {0x166f, 0x167f},
  {0x1681, 0x169a}, {0x16a0, 0x16ea}, {0x16f1, 0x16f8}, {0x1700, 0x1711}, {0x171f, 0x1731}, {0x1740, 0x1751}, {0x1760, 0x176c}, {0x176e, 0x1770},
  {0x1780, 0x17b3}, {0x17d7, 0x17d7}, {0x17dc, 0x17dc}, {0x17e0, 0x17e9}, {0x1810, 0x1819}, {0x1820, 0x1878}, {0x1880, 0x1884}, {0x1887, 0x18a8},
  {0x18aa, 0x18aa}, {0x18b0, 0x18f5}, {0x1900, 0x191e}, {0x1946, 0x196d}, {0x1970, 0x1974}, {0x1980, 0x19ab}, {0x19b0, 0x19c9}, {0x19d0, 0x19d9},
  {0x1a00, 0x1a16}, {0x1a20, 0x1a54}, {0x1a80, 0x1a89}, {0x1a90, 0x1a99}, {0x1aa7, 0x1aa7}, {0x1b05, 0x1b33}, {0x1b45, 0x1b4c}, {0x1b50, 0x1b59},
  {0x1b83, 0x1ba0}, {0x1bae, 0x1be5}, {0x1c00, 0x1c23}, {0x1c40, 0x1c49}, {0x1c4d, 0x1c7d}, {0x1c80, 0x1c88}, {0x1c90, 0x1cba}, {0x1cbd, 0x1cbf},
  {0x1ce9, 0x1cec}, {0x1cee, 0x1cf3}, {0x1cf5, 0x1cf6}, {0x1cfa, 0x1cfa}, {0x1d00, 0x1dbf}, {0x1e00, 0x1f15}, {0x1f18, 0x1f1d}, {0x1f20, 0x1f45},
  {0x1f48, 0x1f4d}, {0x1f50, 0x1f57}, {0x1f59, 0x1f59}, {0x1f5b, 0x1f5b}, {0x1f5d, 0x1f5d}, {0x1f5f, 0x1f7d}, {0x1f80, 0x1fb4}, {0x1fb6, 0x1fbc},
  {0x1fbe, 0x1fbe}, {0x1fc2, 0x1fc4}, {0x1fc6, 0x1fcc}, {0x1fd0, 0x1fd3}, {0x1fd6, 0x1fdb}, {0x1fe0, 0x1fec}, {0x1ff2, 0x1ff4}, {0x1ff6, 0x1ffc},
  {0x2071, 0x2071}, {0x207f, 0x207f}, {0x2090, 0x209c}, {0x2102, 0x2102}, {0x2107, 0x2107}, {0x210a, 0x2113}, {0x2115, 0x2115}, {0x2119, 0x211d},
  {0x2124, 0x2124}, {0x2126, 0x2126}, {0x2128, 0x2128}, {0x212a, 0x212d}, {0x212f, 0x2139}, {0x213c, 0x213f}, {0x2145, 0x2149}, {0x214e, 0x214e},
  {0x2183, 0x2184}, {0x2c00, 0x2ce4}, {0x2ceb, 0x2cee}, {0x2cf2, 0x2cf3}, {0x2d00, 0x2d25}, {0x2d27, 0x2d27}, {0x2d2d, 0x2d2d}, {0x2d30, 0x2d67},
  {0x2d6f, 0x2d6f}, {0x2d80, 0x2d96}, {0x2da0, 0x2da6}, {0x2da8, 0x2dae}, {0x2db0, 0x2db6}, {0x2db8, 0x2dbe}, {0x2dc0, 0x2dc6}, {0x2dc8, 0x2dce},
  {0x2dd0, 0x2dd6}, {0x2dd8, 0x2dde}, {0x2e2f, 0x2e2f}, {0x3005, 0x3006}, {0x3031, 0x3035}, {0x303b, 0x303c}, {0x3041, 0x3096}, {0x309d, 0x309f},
  {0x30a1, 0x30fa}, {0x30fc, 0x30ff}, {0x3105, 0x312f}, {0x3131, 0x318e}, {0x31a0, 0x31bf}, {0x31f0, 0x31ff}, {0x3400, 0x3400}, {0x4dbf, 0x4dbf},
  {0x4e00, 0x4e00}, {0x9fff, 0xa48c}, {0xa4d0, 0xa4fd}, {0xa500, 0xa60c}, {0xa610, 0xa62b}, {0xa640, 0xa66e}, {0xa67f, 0xa69d}, {0xa6a0, 0xa6e5},
  {0xa717, 0xa71f}, {0xa722, 0xa788}, {0xa78b, 0xa7ca}, {0xa7d0, 0xa7d1}, {0xa7d3, 0xa7d3}, {0xa7d5, 0xa7d9}, {0xa7f2, 0xa801}, {0xa803, 0xa805},
  {0xa807, 0xa80a}, {0xa80c, 0xa822}, {0xa840, 0xa873}, {0xa882, 0xa8b3}, {0xa8d0, 0xa8d9}, {0xa8f2, 0xa8f7}, {0xa8fb, 0xa8fb}, {0xa8fd, 0xa8fe},
  {0xa900, 0xa925}, {0xa930, 0xa946}, {0xa960, 0xa97c}, {0xa984, 0xa9b2}, {0xa9cf, 0xa9d9}, {0xa9e0, 0xa9e4}, {0xa9e6, 0xa9fe}, {0xaa00, 0xaa28},
  {0xaa40, 0xaa42}, {0xaa44, 0xaa4b}, {0xaa50, 0xaa59}, {0xaa60, 0xaa76}, {0xaa7a, 0xaa7a}, {0xaa7e, 0xaaaf}, {0xaab1, 0xaab1}, {0xaab5, 0xaab6},
  {0xaab9, 0xaabd}, {0xaac0, 0xaac0}, {0xaac2, 0xaac2}, {0xaadb, 0xaadd}, {0xaae0, 0xaaea}, {0xaaf2, 0xaaf4}, {0xab01, 0xab06}, {0xab09, 0xab0e},
  {0xab11, 0xab16}, {0xab20, 0xab26}, {0xab28, 0xab2e}, {0xab30, 0xab5a}, {0xab5c, 0xab69}, {0xab70, 0xabe2}, {0xabf0, 0xabf9}, {0xac00, 0xac00},
  {0xd7a3, 0xd7a3}, {0xd7b0, 0xd7c6}, {0xd7cb, 0xd7fb}, {0xf900, 0xfa6d}, {0xfa70, 0xfad9}, {0xfb00, 0xfb06}, {0xfb13, 0xfb17}, {0xfb1d, 0xfb1d},
  {0xfb1f, 0xfb28}, {0xfb2a, 0xfb36}, {0xfb38, 0xfb3c}, {0xfb3e, 0xfb3e}, {0xfb40, 0xfb41}, {0xfb43, 0xfb44}, {0xfb46, 0xfbb1}, {0xfbd3, 0xfd3d},
  {0xfd50, 0xfd8f}, {0xfd92, 0xfdc7}, {0xfdf0, 0xfdfb}, {0xfe70, 0xfe74}, {0xfe76, 0xfefc}, {0xff10, 0xff19}, {0xff21, 0xff3a}, {0xff41, 0xff5a},
  {0xff66, 0xffbe}, {0xffc2, 0xffc7}, {0xffca, 0xffcf}, {0xffd2, 0xffd7}, {0xffda, 0xffdc}, {0x10000, 0x1000b}, {0x1000d, 0x10026}, {0x10028, 0x1003a},
  {0x1003c, 0x1003d}, {0x1003f, 0x1004d}, {0x10050, 0x1005d}, {0x10080, 0x100fa}, {0x10280, 0x1029c}, {0x102a0, 0x102d0}, {0x10300, 0x1031f}, {0x1032d, 0x10340},
  {0x10342, 0x10349}, {0x10350, 0x10375}, {0x10380, 0x1039d}, {0x103a0, 0x103c3}, {0x103c8, 0x103cf}, {0x10400, 0x1049d}, {0x104a0, 0x104a9}, {0x104b0, 0x104d3},
  {0x104d8, 0x104fb}, {0x10500, 0x10527}, {0x10530, 0x10563}, {0x10570, 0x1057a}, {0x1057c, 0x1058a}, {0x1058c, 0x10592}, {0x10594, 0x10595}, {0x10597, 0x105a1},
  {0x105a3, 0x105b1}, {0x105b3, 0x105b9}, {0x105bb, 0x105bc}, {0x10600, 0x10736}, {0x10740, 0x10755}, {0x10760, 0x10767}, {0x10780, 0x10785}, {0x10787, 0x107b0},
  {0x107b2, 0x107ba}, {0x10800, 0x10805}, {0x10808, 0x10808}, {0x1080a, 0x10835}, {0x10837, 0x10838}, {0x1083c, 0x1083c}, {0x1083f, 0x10855}, {0x10860, 0x10876},
  {0x10880, 0x1089e}, {0x108e0, 0x108f2}, {0x108f4, 0x108f5}, {0x10900, 0x10915}, {0x10920, 0x1092b},
};

static bool ts_lex(TSLexer *lexer, TSStateId state) {
  START_LEXER();
  eof = lexer->eof(lexer);
  switch (state) {
    case 0:
      if (eof) ADVANCE(322);
      ADVANCE_MAP(
        '!', 399,
        '#', 8,
        '$', 372,
        '%', 394,
        '&', 360,
        '\'', 520,
        '(', 343,
        ')', 345,
        '*', 338,
        '+', 388,
        ',', 344,
        '-', 391,
        '.', 335,
        '/', 393,
        '0', 516,
        ':', 330,
        ';', 350,
        '<', 347,
        '=', 340,
        '>', 349,
        '?', 362,
        '@', 512,
        '[', 331,
        '\\', 304,
        ']', 332,
        '`', 319,
        'a', 728,
        'b', 835,
        'c', 820,
        'd', 714,
        'e', 874,
        'f', 777,
        'g', 754,
        'i', 799,
        'l', 717,
        'n', 825,
        'o', 826,
        'p', 831,
        'r', 765,
        's', 741,
        't', 712,
        'v', 709,
        '{', 341,
        '|', 310,
        '}', 342,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(320);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(517);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 1:
      ADVANCE_MAP(
        '!', 399,
        '%', 394,
        '&', 360,
        '\'', 520,
        '(', 343,
        ')', 345,
        '*', 364,
        '+', 388,
        ',', 344,
        '-', 391,
        '.', 335,
        '/', 393,
        '0', 516,
        ':', 330,
        ';', 350,
        '<', 347,
        '=', 340,
        '>', 349,
        '?', 362,
        '@', 328,
        '[', 331,
        ']', 332,
        '`', 319,
        'a', 554,
        'b', 666,
        'c', 649,
        'd', 541,
        'e', 705,
        'f', 605,
        'g', 581,
        'i', 627,
        'l', 543,
        'n', 655,
        'o', 656,
        'p', 662,
        'r', 592,
        's', 568,
        't', 539,
        'v', 536,
        '{', 341,
        '|', 310,
        '}', 342,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(1);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(517);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('_' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 2:
      ADVANCE_MAP(
        '!', 399,
        '%', 394,
        '&', 360,
        '\'', 520,
        '(', 343,
        ')', 345,
        '*', 364,
        '+', 388,
        ',', 344,
        '-', 391,
        '.', 335,
        '/', 393,
        '0', 516,
        ':', 330,
        ';', 350,
        '<', 347,
        '=', 340,
        '>', 349,
        '?', 362,
        '@', 328,
        '[', 331,
        ']', 332,
        '`', 319,
        'a', 557,
        'b', 666,
        'c', 653,
        'd', 541,
        'e', 707,
        'g', 581,
        'i', 646,
        'r', 592,
        's', 590,
        't', 598,
        'v', 550,
        '{', 341,
        '|', 310,
        '}', 342,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(2);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(517);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('_' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 3:
      ADVANCE_MAP(
        '!', 399,
        '%', 394,
        '&', 19,
        '\'', 520,
        '(', 343,
        '*', 364,
        '+', 388,
        ',', 344,
        '-', 390,
        '.', 335,
        '/', 393,
        '0', 516,
        ':', 35,
        ';', 350,
        '<', 347,
        '=', 340,
        '>', 349,
        '?', 36,
        '@', 512,
        '[', 331,
        '`', 319,
        'a', 554,
        'b', 666,
        'c', 649,
        'd', 541,
        'e', 705,
        'f', 605,
        'g', 581,
        'i', 627,
        'l', 543,
        'n', 655,
        'o', 656,
        'p', 662,
        'r', 592,
        's', 568,
        't', 539,
        'v', 536,
        '{', 341,
        '|', 310,
        '}', 342,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(5);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(517);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('_' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 4:
      ADVANCE_MAP(
        '!', 399,
        '%', 394,
        '&', 19,
        '\'', 520,
        '(', 343,
        '*', 364,
        '+', 388,
        ',', 344,
        '-', 390,
        '.', 335,
        '/', 393,
        '0', 516,
        ':', 35,
        ';', 350,
        '<', 347,
        '=', 340,
        '>', 349,
        '?', 36,
        '@', 512,
        '[', 331,
        '`', 319,
        'a', 557,
        'b', 666,
        'c', 653,
        'd', 541,
        'e', 707,
        'g', 581,
        'i', 646,
        'r', 592,
        's', 590,
        't', 598,
        'v', 550,
        '{', 341,
        '|', 310,
        '}', 342,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(6);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(517);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('_' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 5:
      ADVANCE_MAP(
        '!', 399,
        '%', 394,
        '&', 19,
        '\'', 520,
        '(', 343,
        '*', 364,
        '+', 388,
        ',', 344,
        '-', 390,
        '.', 335,
        '/', 393,
        '0', 516,
        ':', 35,
        ';', 350,
        '<', 347,
        '=', 340,
        '>', 349,
        '?', 36,
        '@', 328,
        '[', 331,
        '`', 319,
        'a', 554,
        'b', 666,
        'c', 649,
        'd', 541,
        'e', 705,
        'f', 605,
        'g', 581,
        'i', 627,
        'l', 543,
        'n', 655,
        'o', 656,
        'p', 662,
        'r', 592,
        's', 568,
        't', 539,
        'v', 536,
        '{', 341,
        '|', 310,
        '}', 342,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(5);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(517);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('_' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 6:
      ADVANCE_MAP(
        '!', 399,
        '%', 394,
        '&', 19,
        '\'', 520,
        '(', 343,
        '*', 364,
        '+', 388,
        ',', 344,
        '-', 390,
        '.', 335,
        '/', 393,
        '0', 516,
        ':', 35,
        ';', 350,
        '<', 347,
        '=', 340,
        '>', 349,
        '?', 36,
        '@', 328,
        '[', 331,
        '`', 319,
        'a', 557,
        'b', 666,
        'c', 653,
        'd', 541,
        'e', 707,
        'g', 581,
        'i', 646,
        'r', 592,
        's', 590,
        't', 598,
        'v', 550,
        '{', 341,
        '|', 310,
        '}', 342,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(6);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(517);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('_' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 7:
      if (lookahead == '!') ADVANCE(401);
      if (lookahead == '=') ADVANCE(379);
      if (lookahead == 'i') ADVANCE(249);
      END_STATE();
    case 8:
      if (lookahead == '!') ADVANCE(323);
      END_STATE();
    case 9:
      ADVANCE_MAP(
        '!', 398,
        '\'', 520,
        '(', 343,
        ')', 345,
        '*', 363,
        '+', 387,
        '-', 389,
        '.', 336,
        '/', 30,
        '0', 516,
        ':', 330,
        '<', 346,
        '@', 328,
        '[', 331,
        '`', 319,
        'a', 558,
        'b', 666,
        'c', 653,
        'd', 541,
        'e', 707,
        'g', 581,
        'i', 646,
        'r', 592,
        's', 590,
        't', 598,
        'v', 550,
        '{', 341,
        '}', 342,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(9);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(517);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('_' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 10:
      ADVANCE_MAP(
        '!', 398,
        '\'', 520,
        '(', 343,
        '*', 363,
        '+', 387,
        '-', 389,
        '.', 313,
        '/', 30,
        '0', 516,
        ':', 35,
        '@', 328,
        '[', 331,
        '`', 319,
        'a', 558,
        'b', 666,
        'c', 653,
        'd', 541,
        'e', 707,
        'g', 581,
        'i', 646,
        'r', 592,
        's', 589,
        't', 598,
        'v', 550,
        '{', 341,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(10);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(517);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('_' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 11:
      ADVANCE_MAP(
        '!', 400,
        '&', 359,
        '\'', 520,
        '(', 343,
        '*', 363,
        '+', 387,
        ',', 344,
        '-', 392,
        '.', 336,
        '/', 30,
        '0', 516,
        ':', 330,
        ';', 350,
        '<', 346,
        '=', 339,
        '?', 361,
        '@', 328,
        '[', 331,
        '`', 319,
        'a', 555,
        'b', 666,
        'c', 649,
        'd', 541,
        'e', 705,
        'f', 605,
        'g', 581,
        'i', 627,
        'l', 543,
        'n', 655,
        'o', 656,
        'p', 662,
        'r', 592,
        's', 568,
        't', 539,
        'v', 536,
        '{', 341,
        '}', 342,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(11);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(517);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('_' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 12:
      ADVANCE_MAP(
        '!', 400,
        '&', 359,
        '\'', 520,
        '(', 343,
        '*', 363,
        '+', 387,
        ',', 344,
        '-', 392,
        '.', 336,
        '/', 30,
        '0', 516,
        ':', 330,
        '<', 346,
        '=', 339,
        '?', 361,
        '@', 328,
        '[', 331,
        '`', 319,
        'a', 558,
        'b', 666,
        'c', 653,
        'd', 541,
        'e', 707,
        'g', 581,
        'i', 646,
        'r', 592,
        's', 590,
        't', 598,
        'v', 550,
        '{', 341,
        '}', 342,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(12);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(517);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('_' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 13:
      ADVANCE_MAP(
        '!', 7,
        '%', 394,
        '&', 360,
        '(', 343,
        ')', 345,
        '*', 364,
        '+', 388,
        ',', 344,
        '-', 391,
        '.', 334,
        '/', 393,
        ':', 330,
        ';', 350,
        '<', 347,
        '=', 340,
        '>', 349,
        '?', 362,
        '@', 328,
        '[', 331,
        ']', 332,
        '`', 319,
        'a', 554,
        'c', 654,
        'd', 541,
        'e', 705,
        'f', 605,
        'g', 581,
        'i', 627,
        'l', 543,
        'n', 655,
        'o', 656,
        'p', 662,
        's', 569,
        't', 540,
        'v', 536,
        '{', 341,
        '|', 310,
        '}', 342,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(13);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('_' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 14:
      ADVANCE_MAP(
        '!', 7,
        '%', 394,
        '&', 360,
        '(', 343,
        ')', 345,
        '*', 364,
        '+', 388,
        ',', 344,
        '-', 391,
        '.', 334,
        '/', 393,
        ':', 330,
        ';', 350,
        '<', 347,
        '=', 340,
        '>', 349,
        '?', 362,
        '@', 328,
        '[', 331,
        ']', 332,
        '`', 319,
        'a', 557,
        'd', 541,
        'e', 707,
        'g', 581,
        'i', 646,
        's', 591,
        'v', 550,
        '{', 341,
        '|', 310,
        '}', 342,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(14);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('_' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 15:
      ADVANCE_MAP(
        '!', 7,
        '%', 394,
        '&', 19,
        '(', 343,
        ')', 345,
        '*', 364,
        '+', 388,
        ',', 344,
        '-', 391,
        '.', 334,
        '/', 393,
        ':', 35,
        ';', 350,
        '<', 347,
        '=', 340,
        '>', 349,
        '?', 36,
        '@', 512,
        '[', 331,
        ']', 332,
        '`', 319,
        'a', 554,
        'c', 654,
        'd', 541,
        'e', 705,
        'f', 605,
        'g', 581,
        'i', 627,
        'l', 543,
        'n', 655,
        'o', 656,
        'p', 662,
        's', 569,
        't', 540,
        'v', 536,
        '{', 341,
        '|', 310,
        '}', 342,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(17);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('_' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 16:
      ADVANCE_MAP(
        '!', 7,
        '%', 394,
        '&', 19,
        '(', 343,
        ')', 345,
        '*', 364,
        '+', 388,
        ',', 344,
        '-', 391,
        '.', 334,
        '/', 393,
        ':', 35,
        ';', 350,
        '<', 347,
        '=', 340,
        '>', 349,
        '?', 36,
        '@', 512,
        '[', 331,
        ']', 332,
        '`', 319,
        'a', 557,
        'd', 541,
        'e', 707,
        'g', 581,
        'i', 646,
        's', 591,
        'v', 550,
        '{', 341,
        '|', 310,
        '}', 342,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(18);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('_' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 17:
      ADVANCE_MAP(
        '!', 7,
        '%', 394,
        '&', 19,
        '(', 343,
        ')', 345,
        '*', 364,
        '+', 388,
        ',', 344,
        '-', 391,
        '.', 334,
        '/', 393,
        ':', 35,
        ';', 350,
        '<', 347,
        '=', 340,
        '>', 349,
        '?', 36,
        '@', 328,
        '[', 331,
        ']', 332,
        '`', 319,
        'a', 554,
        'c', 654,
        'd', 541,
        'e', 705,
        'f', 605,
        'g', 581,
        'i', 627,
        'l', 543,
        'n', 655,
        'o', 656,
        'p', 662,
        's', 569,
        't', 540,
        'v', 536,
        '{', 341,
        '|', 310,
        '}', 342,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(17);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('_' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 18:
      ADVANCE_MAP(
        '!', 7,
        '%', 394,
        '&', 19,
        '(', 343,
        ')', 345,
        '*', 364,
        '+', 388,
        ',', 344,
        '-', 391,
        '.', 334,
        '/', 393,
        ':', 35,
        ';', 350,
        '<', 347,
        '=', 340,
        '>', 349,
        '?', 36,
        '@', 328,
        '[', 331,
        ']', 332,
        '`', 319,
        'a', 557,
        'd', 541,
        'e', 707,
        'g', 581,
        'i', 646,
        's', 591,
        'v', 550,
        '{', 341,
        '|', 310,
        '}', 342,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(18);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('_' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 19:
      if (lookahead == '&') ADVANCE(369);
      END_STATE();
    case 20:
      ADVANCE_MAP(
        '&', 359,
        '(', 343,
        ')', 345,
        '*', 337,
        ',', 344,
        '-', 38,
        '.', 333,
        '/', 30,
        ':', 329,
        ';', 350,
        '<', 346,
        '=', 339,
        '>', 348,
        '?', 361,
        '@', 328,
        '[', 331,
        ']', 332,
        '`', 319,
        'a', 731,
        'd', 714,
        'e', 876,
        'g', 754,
        'i', 817,
        's', 764,
        'v', 724,
        '{', 341,
        '}', 342,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(21);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 21:
      ADVANCE_MAP(
        '&', 359,
        '(', 343,
        ')', 345,
        ',', 344,
        '-', 38,
        '.', 333,
        '/', 30,
        ':', 329,
        ';', 350,
        '<', 346,
        '=', 339,
        '>', 348,
        '?', 361,
        '@', 328,
        '[', 331,
        ']', 332,
        '`', 319,
        'a', 731,
        'd', 714,
        'e', 876,
        'g', 754,
        'i', 817,
        's', 764,
        'v', 724,
        '{', 341,
        '}', 342,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(21);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 22:
      ADVANCE_MAP(
        '&', 359,
        '(', 343,
        ')', 345,
        ',', 344,
        '-', 38,
        '.', 333,
        '/', 30,
        ':', 329,
        ';', 350,
        '<', 346,
        '=', 339,
        '>', 348,
        '?', 361,
        '@', 328,
        '`', 319,
        '{', 341,
        '}', 342,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(22);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 23:
      ADVANCE_MAP(
        '&', 359,
        '(', 343,
        ')', 345,
        ',', 344,
        '-', 38,
        '.', 333,
        '/', 30,
        ':', 329,
        ';', 350,
        '<', 346,
        '=', 339,
        '?', 361,
        '@', 328,
        '`', 319,
        'a', 729,
        'c', 824,
        'd', 714,
        'e', 874,
        'f', 777,
        'g', 754,
        'i', 799,
        'l', 717,
        'n', 825,
        'o', 826,
        'p', 831,
        's', 742,
        't', 713,
        'v', 709,
        '{', 341,
        '}', 342,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(23);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 24:
      ADVANCE_MAP(
        '&', 359,
        '(', 343,
        ')', 345,
        ',', 344,
        '-', 38,
        '.', 333,
        '/', 30,
        '=', 339,
        '>', 348,
        '?', 361,
        '@', 328,
        'a', 555,
        'c', 654,
        'd', 541,
        'e', 705,
        'f', 605,
        'g', 581,
        'i', 627,
        'l', 543,
        'n', 655,
        'o', 656,
        'p', 662,
        's', 569,
        't', 540,
        'v', 536,
        '{', 341,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(24);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 25:
      ADVANCE_MAP(
        '&', 359,
        '(', 343,
        ',', 344,
        '.', 333,
        '/', 30,
        ':', 329,
        '<', 346,
        '=', 339,
        '?', 361,
        '@', 328,
        'a', 729,
        'c', 824,
        'd', 714,
        'e', 874,
        'f', 777,
        'i', 799,
        'l', 717,
        'n', 825,
        'o', 826,
        'p', 831,
        's', 762,
        't', 713,
        'v', 709,
        '{', 341,
        '}', 342,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(25);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 26:
      ADVANCE_MAP(
        '(', 343,
        ')', 345,
        '*', 363,
        '.', 333,
        '/', 30,
        ':', 329,
        '<', 346,
        '=', 339,
        '@', 328,
        '`', 319,
        'a', 731,
        'd', 714,
        'e', 876,
        'g', 754,
        'i', 817,
        's', 763,
        'v', 724,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(26);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 27:
      ADVANCE_MAP(
        '(', 343,
        ')', 345,
        ',', 344,
        '.', 333,
        '/', 30,
        ':', 329,
        '<', 346,
        '@', 328,
        '`', 319,
        'a', 731,
        'c', 836,
        'd', 714,
        'e', 876,
        'g', 754,
        'i', 817,
        'n', 825,
        's', 763,
        'v', 709,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(27);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 28:
      ADVANCE_MAP(
        '(', 343,
        ')', 345,
        ',', 344,
        '.', 333,
        '/', 30,
        ':', 329,
        '<', 346,
        '@', 328,
        '`', 319,
        'a', 731,
        'c', 836,
        'd', 714,
        'e', 876,
        'g', 754,
        'i', 817,
        'n', 825,
        's', 764,
        'v', 709,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(28);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 29:
      ADVANCE_MAP(
        '(', 343,
        '.', 333,
        '/', 30,
        ':', 329,
        '<', 346,
        '@', 328,
        'a', 54,
        'c', 126,
        'd', 42,
        'e', 168,
        'f', 92,
        'g', 76,
        'i', 111,
        'l', 44,
        'n', 131,
        'o', 132,
        'p', 135,
        's', 65,
        't', 41,
        'v', 39,
        '{', 341,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(29);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 30:
      if (lookahead == '/') ADVANCE(506);
      END_STATE();
    case 31:
      if (lookahead == '/') ADVANCE(506);
      if (lookahead == '=') ADVANCE(377);
      END_STATE();
    case 32:
      ADVANCE_MAP(
        '/', 30,
        ';', 350,
        '@', 328,
        'a', 192,
        'c', 266,
        'd', 179,
        'e', 308,
        'f', 230,
        'g', 203,
        'i', 250,
        'l', 180,
        'n', 265,
        'o', 271,
        'p', 274,
        's', 204,
        't', 181,
        'v', 177,
        '}', 342,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(32);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          lookahead == 'b') ADVANCE(318);
      END_STATE();
    case 33:
      if (lookahead == '/') ADVANCE(30);
      if (lookahead == 'g') ADVANCE(754);
      if (lookahead == 's') ADVANCE(764);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(33);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 34:
      if (lookahead == '/') ADVANCE(522);
      if (lookahead == '\\') ADVANCE(304);
      if (lookahead == '\n' ||
          lookahead == '\r') SKIP(34);
      if (('\t' <= lookahead && lookahead <= '\f') ||
          lookahead == ' ') ADVANCE(523);
      if (lookahead != 0 &&
          lookahead != '\'') ADVANCE(521);
      END_STATE();
    case 35:
      if (lookahead == ':') ADVANCE(373);
      END_STATE();
    case 36:
      if (lookahead == ':') ADVANCE(368);
      END_STATE();
    case 37:
      if (lookahead == '=') ADVANCE(378);
      END_STATE();
    case 38:
      if (lookahead == '>') ADVANCE(365);
      END_STATE();
    case 39:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(105);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 40:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(416);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 41:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(91);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 42:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(157);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 43:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(108);
      if (lookahead == 't') ADVANCE(357);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 44:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(160);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 45:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(101);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 46:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(102);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 47:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(137);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 48:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(103);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 49:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(104);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 50:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(158);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 51:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(159);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 52:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(161);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 53:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(59);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 54:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'b') ADVANCE(147);
      if (lookahead == 'c') ADVANCE(150);
      if (lookahead == 'n') ADVANCE(114);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 55:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'b') ADVANCE(106);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 56:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'c') ADVANCE(436);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 57:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'c') ADVANCE(452);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 58:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'c') ADVANCE(153);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 59:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'c') ADVANCE(154);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 60:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'c') ADVANCE(162);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 61:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'd') ADVANCE(408);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 62:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'd') ADVANCE(404);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 63:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'd') ADVANCE(448);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 64:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'd') ADVANCE(74);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 65:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(43);
      if (lookahead == 'u') ADVANCE(145);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 66:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(112);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 67:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(140);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 68:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(58);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 69:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(61);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 70:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(424);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 71:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(464);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 72:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(440);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 73:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(492);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 74:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(428);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 75:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(496);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 76:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(151);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 77:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(143);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 78:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(57);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 79:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(136);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 80:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(63);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 81:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(115);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 82:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(60);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 83:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(96);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 84:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(144);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 85:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'f') ADVANCE(88);
      if (lookahead == 'l') ADVANCE(95);
      if (lookahead == 'n') ADVANCE(79);
      if (lookahead == 't') ADVANCE(84);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 86:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'g') ADVANCE(488);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 87:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(167);
      if (lookahead == 'o') ADVANCE(163);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 88:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(169);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 89:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(56);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 90:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(64);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 91:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(107);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 92:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(116);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 93:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(117);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 94:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(130);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 95:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(121);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 96:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(118);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 97:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(155);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 98:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(122);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 99:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(123);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 100:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(125);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 101:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'l') ADVANCE(480);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 102:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'l') ADVANCE(504);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 103:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'l') ADVANCE(468);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 104:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'l') ADVANCE(444);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 105:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'l') ADVANCE(166);
      if (lookahead == 'r') ADVANCE(47);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 106:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'l') ADVANCE(89);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 107:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'l') ADVANCE(142);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 108:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'l') ADVANCE(69);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 109:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'l') ADVANCE(98);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 110:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'l') ADVANCE(99);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 111:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(85);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 112:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(484);
      if (lookahead == 'r') ADVANCE(50);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 113:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(412);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 114:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(128);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 115:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(62);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 116:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(45);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 117:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(109);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 118:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(97);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 119:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(48);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 120:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(49);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 121:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(71);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 122:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(73);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 123:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(75);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 124:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(148);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 125:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(110);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 126:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'o') ADVANCE(124);
      if (lookahead == 'r') ADVANCE(127);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 127:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'o') ADVANCE(146);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 128:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'o') ADVANCE(164);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 129:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'o') ADVANCE(138);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 130:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'o') ADVANCE(113);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 131:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'o') ADVANCE(93);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 132:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'p') ADVANCE(66);
      if (lookahead == 'v') ADVANCE(67);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 133:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'p') ADVANCE(68);
      if (lookahead == 't') ADVANCE(77);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 134:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'p') ADVANCE(81);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 135:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'r') ADVANCE(87);
      if (lookahead == 'u') ADVANCE(55);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 136:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'r') ADVANCE(420);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 137:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'r') ADVANCE(86);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 138:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'r') ADVANCE(456);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 139:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'r') ADVANCE(90);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 140:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'r') ADVANCE(139);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 141:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'r') ADVANCE(53);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 142:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'r') ADVANCE(78);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 143:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'r') ADVANCE(119);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 144:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'r') ADVANCE(120);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 145:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 's') ADVANCE(134);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 146:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 's') ADVANCE(149);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 147:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 's') ADVANCE(156);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 148:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 's') ADVANCE(152);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 149:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 's') ADVANCE(100);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 150:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(165);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 151:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(353);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 152:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(472);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 153:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(500);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 154:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(476);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 155:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(432);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 156:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(141);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 157:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(40);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 158:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(129);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 159:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(94);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 160:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(83);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 161:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(72);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 162:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(80);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 163:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(82);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 164:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(51);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 165:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'u') ADVANCE(46);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 166:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'u') ADVANCE(70);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 167:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'v') ADVANCE(52);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 168:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'x') ADVANCE(133);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 169:
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'x') ADVANCE(460);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 170:
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 171:
      if (lookahead == '_') ADVANCE(171);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(517);
      END_STATE();
    case 172:
      if (lookahead == '_') ADVANCE(172);
      if (lookahead == '0' ||
          lookahead == '1') ADVANCE(519);
      END_STATE();
    case 173:
      if (lookahead == '_') ADVANCE(173);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(518);
      END_STATE();
    case 174:
      if (lookahead == '_') ADVANCE(174);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(514);
      END_STATE();
    case 175:
      if (lookahead == '_') ADVANCE(175);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(515);
      END_STATE();
    case 176:
      if (lookahead == '`') ADVANCE(878);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r') ADVANCE(176);
      END_STATE();
    case 177:
      if (lookahead == 'a') ADVANCE(243);
      END_STATE();
    case 178:
      if (lookahead == 'a') ADVANCE(414);
      END_STATE();
    case 179:
      if (lookahead == 'a') ADVANCE(296);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('b' <= lookahead && lookahead <= 'f')) ADVANCE(317);
      END_STATE();
    case 180:
      if (lookahead == 'a') ADVANCE(298);
      END_STATE();
    case 181:
      if (lookahead == 'a') ADVANCE(229);
      END_STATE();
    case 182:
      if (lookahead == 'a') ADVANCE(246);
      if (lookahead == 't') ADVANCE(355);
      END_STATE();
    case 183:
      if (lookahead == 'a') ADVANCE(239);
      END_STATE();
    case 184:
      if (lookahead == 'a') ADVANCE(240);
      END_STATE();
    case 185:
      if (lookahead == 'a') ADVANCE(297);
      END_STATE();
    case 186:
      if (lookahead == 'a') ADVANCE(276);
      END_STATE();
    case 187:
      if (lookahead == 'a') ADVANCE(241);
      END_STATE();
    case 188:
      if (lookahead == 'a') ADVANCE(299);
      END_STATE();
    case 189:
      if (lookahead == 'a') ADVANCE(242);
      END_STATE();
    case 190:
      if (lookahead == 'a') ADVANCE(197);
      END_STATE();
    case 191:
      if (lookahead == 'a') ADVANCE(300);
      END_STATE();
    case 192:
      if (lookahead == 'b') ADVANCE(285);
      if (lookahead == 'c') ADVANCE(289);
      if (lookahead == 'n') ADVANCE(254);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(317);
      END_STATE();
    case 193:
      if (lookahead == 'b') ADVANCE(244);
      END_STATE();
    case 194:
      if (lookahead == 'c') ADVANCE(434);
      END_STATE();
    case 195:
      if (lookahead == 'c') ADVANCE(450);
      END_STATE();
    case 196:
      if (lookahead == 'c') ADVANCE(292);
      END_STATE();
    case 197:
      if (lookahead == 'c') ADVANCE(293);
      END_STATE();
    case 198:
      if (lookahead == 'c') ADVANCE(301);
      END_STATE();
    case 199:
      if (lookahead == 'd') ADVANCE(406);
      END_STATE();
    case 200:
      if (lookahead == 'd') ADVANCE(402);
      END_STATE();
    case 201:
      if (lookahead == 'd') ADVANCE(446);
      END_STATE();
    case 202:
      if (lookahead == 'd') ADVANCE(213);
      END_STATE();
    case 203:
      if (lookahead == 'e') ADVANCE(290);
      END_STATE();
    case 204:
      if (lookahead == 'e') ADVANCE(182);
      if (lookahead == 'u') ADVANCE(284);
      END_STATE();
    case 205:
      if (lookahead == 'e') ADVANCE(251);
      END_STATE();
    case 206:
      if (lookahead == 'e') ADVANCE(279);
      END_STATE();
    case 207:
      if (lookahead == 'e') ADVANCE(196);
      END_STATE();
    case 208:
      if (lookahead == 'e') ADVANCE(199);
      END_STATE();
    case 209:
      if (lookahead == 'e') ADVANCE(422);
      END_STATE();
    case 210:
      if (lookahead == 'e') ADVANCE(462);
      END_STATE();
    case 211:
      if (lookahead == 'e') ADVANCE(438);
      END_STATE();
    case 212:
      if (lookahead == 'e') ADVANCE(490);
      END_STATE();
    case 213:
      if (lookahead == 'e') ADVANCE(426);
      END_STATE();
    case 214:
      if (lookahead == 'e') ADVANCE(494);
      END_STATE();
    case 215:
      if (lookahead == 'e') ADVANCE(282);
      END_STATE();
    case 216:
      if (lookahead == 'e') ADVANCE(195);
      END_STATE();
    case 217:
      if (lookahead == 'e') ADVANCE(275);
      END_STATE();
    case 218:
      if (lookahead == 'e') ADVANCE(201);
      END_STATE();
    case 219:
      if (lookahead == 'e') ADVANCE(255);
      END_STATE();
    case 220:
      if (lookahead == 'e') ADVANCE(198);
      END_STATE();
    case 221:
      if (lookahead == 'e') ADVANCE(234);
      END_STATE();
    case 222:
      if (lookahead == 'e') ADVANCE(283);
      END_STATE();
    case 223:
      if (lookahead == 'f') ADVANCE(226);
      if (lookahead == 'l') ADVANCE(233);
      if (lookahead == 'n') ADVANCE(217);
      if (lookahead == 't') ADVANCE(222);
      END_STATE();
    case 224:
      if (lookahead == 'g') ADVANCE(486);
      END_STATE();
    case 225:
      if (lookahead == 'i') ADVANCE(307);
      if (lookahead == 'o') ADVANCE(302);
      END_STATE();
    case 226:
      if (lookahead == 'i') ADVANCE(309);
      END_STATE();
    case 227:
      if (lookahead == 'i') ADVANCE(194);
      END_STATE();
    case 228:
      if (lookahead == 'i') ADVANCE(202);
      END_STATE();
    case 229:
      if (lookahead == 'i') ADVANCE(245);
      END_STATE();
    case 230:
      if (lookahead == 'i') ADVANCE(256);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(317);
      END_STATE();
    case 231:
      if (lookahead == 'i') ADVANCE(257);
      END_STATE();
    case 232:
      if (lookahead == 'i') ADVANCE(270);
      END_STATE();
    case 233:
      if (lookahead == 'i') ADVANCE(261);
      END_STATE();
    case 234:
      if (lookahead == 'i') ADVANCE(258);
      END_STATE();
    case 235:
      if (lookahead == 'i') ADVANCE(294);
      END_STATE();
    case 236:
      if (lookahead == 'i') ADVANCE(262);
      END_STATE();
    case 237:
      if (lookahead == 'i') ADVANCE(263);
      END_STATE();
    case 238:
      if (lookahead == 'i') ADVANCE(264);
      END_STATE();
    case 239:
      if (lookahead == 'l') ADVANCE(478);
      END_STATE();
    case 240:
      if (lookahead == 'l') ADVANCE(502);
      END_STATE();
    case 241:
      if (lookahead == 'l') ADVANCE(466);
      END_STATE();
    case 242:
      if (lookahead == 'l') ADVANCE(442);
      END_STATE();
    case 243:
      if (lookahead == 'l') ADVANCE(306);
      if (lookahead == 'r') ADVANCE(186);
      END_STATE();
    case 244:
      if (lookahead == 'l') ADVANCE(227);
      END_STATE();
    case 245:
      if (lookahead == 'l') ADVANCE(281);
      END_STATE();
    case 246:
      if (lookahead == 'l') ADVANCE(208);
      END_STATE();
    case 247:
      if (lookahead == 'l') ADVANCE(236);
      END_STATE();
    case 248:
      if (lookahead == 'l') ADVANCE(237);
      END_STATE();
    case 249:
      if (lookahead == 'n') ADVANCE(385);
      if (lookahead == 's') ADVANCE(386);
      END_STATE();
    case 250:
      if (lookahead == 'n') ADVANCE(223);
      END_STATE();
    case 251:
      if (lookahead == 'n') ADVANCE(482);
      if (lookahead == 'r') ADVANCE(185);
      END_STATE();
    case 252:
      if (lookahead == 'n') ADVANCE(410);
      END_STATE();
    case 253:
      if (lookahead == 'n') ADVANCE(287);
      END_STATE();
    case 254:
      if (lookahead == 'n') ADVANCE(268);
      END_STATE();
    case 255:
      if (lookahead == 'n') ADVANCE(200);
      END_STATE();
    case 256:
      if (lookahead == 'n') ADVANCE(183);
      END_STATE();
    case 257:
      if (lookahead == 'n') ADVANCE(247);
      END_STATE();
    case 258:
      if (lookahead == 'n') ADVANCE(235);
      END_STATE();
    case 259:
      if (lookahead == 'n') ADVANCE(187);
      END_STATE();
    case 260:
      if (lookahead == 'n') ADVANCE(189);
      END_STATE();
    case 261:
      if (lookahead == 'n') ADVANCE(210);
      END_STATE();
    case 262:
      if (lookahead == 'n') ADVANCE(212);
      END_STATE();
    case 263:
      if (lookahead == 'n') ADVANCE(214);
      END_STATE();
    case 264:
      if (lookahead == 'n') ADVANCE(248);
      END_STATE();
    case 265:
      if (lookahead == 'o') ADVANCE(231);
      END_STATE();
    case 266:
      if (lookahead == 'o') ADVANCE(253);
      if (lookahead == 'r') ADVANCE(267);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(317);
      END_STATE();
    case 267:
      if (lookahead == 'o') ADVANCE(286);
      END_STATE();
    case 268:
      if (lookahead == 'o') ADVANCE(303);
      END_STATE();
    case 269:
      if (lookahead == 'o') ADVANCE(277);
      END_STATE();
    case 270:
      if (lookahead == 'o') ADVANCE(252);
      END_STATE();
    case 271:
      if (lookahead == 'p') ADVANCE(205);
      if (lookahead == 'v') ADVANCE(206);
      END_STATE();
    case 272:
      if (lookahead == 'p') ADVANCE(207);
      if (lookahead == 't') ADVANCE(215);
      END_STATE();
    case 273:
      if (lookahead == 'p') ADVANCE(219);
      END_STATE();
    case 274:
      if (lookahead == 'r') ADVANCE(225);
      if (lookahead == 'u') ADVANCE(193);
      END_STATE();
    case 275:
      if (lookahead == 'r') ADVANCE(418);
      END_STATE();
    case 276:
      if (lookahead == 'r') ADVANCE(224);
      END_STATE();
    case 277:
      if (lookahead == 'r') ADVANCE(454);
      END_STATE();
    case 278:
      if (lookahead == 'r') ADVANCE(228);
      END_STATE();
    case 279:
      if (lookahead == 'r') ADVANCE(278);
      END_STATE();
    case 280:
      if (lookahead == 'r') ADVANCE(190);
      END_STATE();
    case 281:
      if (lookahead == 'r') ADVANCE(216);
      END_STATE();
    case 282:
      if (lookahead == 'r') ADVANCE(259);
      END_STATE();
    case 283:
      if (lookahead == 'r') ADVANCE(260);
      END_STATE();
    case 284:
      if (lookahead == 's') ADVANCE(273);
      END_STATE();
    case 285:
      if (lookahead == 's') ADVANCE(295);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(316);
      END_STATE();
    case 286:
      if (lookahead == 's') ADVANCE(288);
      END_STATE();
    case 287:
      if (lookahead == 's') ADVANCE(291);
      END_STATE();
    case 288:
      if (lookahead == 's') ADVANCE(238);
      END_STATE();
    case 289:
      if (lookahead == 't') ADVANCE(305);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(316);
      END_STATE();
    case 290:
      if (lookahead == 't') ADVANCE(351);
      END_STATE();
    case 291:
      if (lookahead == 't') ADVANCE(470);
      END_STATE();
    case 292:
      if (lookahead == 't') ADVANCE(498);
      END_STATE();
    case 293:
      if (lookahead == 't') ADVANCE(474);
      END_STATE();
    case 294:
      if (lookahead == 't') ADVANCE(430);
      END_STATE();
    case 295:
      if (lookahead == 't') ADVANCE(280);
      END_STATE();
    case 296:
      if (lookahead == 't') ADVANCE(178);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(316);
      END_STATE();
    case 297:
      if (lookahead == 't') ADVANCE(269);
      END_STATE();
    case 298:
      if (lookahead == 't') ADVANCE(221);
      END_STATE();
    case 299:
      if (lookahead == 't') ADVANCE(232);
      END_STATE();
    case 300:
      if (lookahead == 't') ADVANCE(211);
      END_STATE();
    case 301:
      if (lookahead == 't') ADVANCE(218);
      END_STATE();
    case 302:
      if (lookahead == 't') ADVANCE(220);
      END_STATE();
    case 303:
      if (lookahead == 't') ADVANCE(188);
      END_STATE();
    case 304:
      ADVANCE_MAP(
        'u', 879,
        '"', 881,
        '$', 881,
        '\'', 881,
        '\\', 881,
        'b', 881,
        'n', 881,
        'r', 881,
        't', 881,
      );
      END_STATE();
    case 305:
      if (lookahead == 'u') ADVANCE(184);
      END_STATE();
    case 306:
      if (lookahead == 'u') ADVANCE(209);
      END_STATE();
    case 307:
      if (lookahead == 'v') ADVANCE(191);
      END_STATE();
    case 308:
      if (lookahead == 'x') ADVANCE(272);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(317);
      END_STATE();
    case 309:
      if (lookahead == 'x') ADVANCE(458);
      END_STATE();
    case 310:
      if (lookahead == '|') ADVANCE(370);
      END_STATE();
    case 311:
      if (lookahead == '+' ||
          lookahead == '-') ADVANCE(314);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(515);
      END_STATE();
    case 312:
      if (lookahead == '0' ||
          lookahead == '1') ADVANCE(519);
      END_STATE();
    case 313:
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(514);
      END_STATE();
    case 314:
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(515);
      END_STATE();
    case 315:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(518);
      END_STATE();
    case 316:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(880);
      END_STATE();
    case 317:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(316);
      END_STATE();
    case 318:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(317);
      END_STATE();
    case 319:
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '`') ADVANCE(176);
      END_STATE();
    case 320:
      if (eof) ADVANCE(322);
      ADVANCE_MAP(
        '!', 399,
        '#', 8,
        '$', 372,
        '%', 394,
        '&', 360,
        '\'', 520,
        '(', 343,
        ')', 345,
        '*', 364,
        '+', 388,
        ',', 344,
        '-', 391,
        '.', 335,
        '/', 393,
        '0', 516,
        ':', 330,
        ';', 350,
        '<', 347,
        '=', 340,
        '>', 349,
        '?', 362,
        '@', 328,
        '[', 331,
        '\\', 304,
        ']', 332,
        '`', 319,
        'a', 728,
        'b', 835,
        'c', 820,
        'd', 714,
        'e', 874,
        'f', 777,
        'g', 754,
        'i', 799,
        'l', 717,
        'n', 825,
        'o', 826,
        'p', 831,
        'r', 765,
        's', 741,
        't', 712,
        'v', 709,
        '{', 341,
        '|', 310,
        '}', 342,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(320);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(517);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 321:
      if (eof) ADVANCE(322);
      ADVANCE_MAP(
        '!', 398,
        '#', 8,
        '%', 37,
        '&', 359,
        '\'', 520,
        '(', 343,
        ')', 345,
        '*', 364,
        '+', 388,
        ',', 344,
        '-', 391,
        '.', 336,
        '/', 31,
        '0', 516,
        ':', 330,
        ';', 350,
        '<', 346,
        '=', 339,
        '>', 348,
        '?', 361,
        '@', 328,
        '[', 331,
        '`', 319,
        'a', 555,
        'b', 666,
        'c', 649,
        'd', 541,
        'e', 705,
        'f', 605,
        'g', 581,
        'i', 627,
        'l', 543,
        'n', 655,
        'o', 656,
        'p', 662,
        'r', 592,
        's', 568,
        't', 539,
        'v', 536,
        '{', 341,
        '}', 342,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(321);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(517);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          ('_' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_1, 431, lookahead)) ADVANCE(877);
      END_STATE();
    case 322:
      ACCEPT_TOKEN(ts_builtin_sym_end);
      END_STATE();
    case 323:
      ACCEPT_TOKEN(anon_sym_POUND_BANG);
      END_STATE();
    case 324:
      ACCEPT_TOKEN(aux_sym_shebang_line_token1);
      if (lookahead == '\r') ADVANCE(506);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(324);
      END_STATE();
    case 325:
      ACCEPT_TOKEN(aux_sym_shebang_line_token1);
      if (lookahead == '/') ADVANCE(326);
      if (lookahead == '\t' ||
          lookahead == 0x0b ||
          lookahead == '\f' ||
          lookahead == ' ') ADVANCE(325);
      if (lookahead != 0 &&
          (lookahead < '\t' || '\r' < lookahead)) ADVANCE(327);
      END_STATE();
    case 326:
      ACCEPT_TOKEN(aux_sym_shebang_line_token1);
      if (lookahead == '/') ADVANCE(324);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r') ADVANCE(327);
      END_STATE();
    case 327:
      ACCEPT_TOKEN(aux_sym_shebang_line_token1);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r') ADVANCE(327);
      END_STATE();
    case 328:
      ACCEPT_TOKEN(anon_sym_AT);
      END_STATE();
    case 329:
      ACCEPT_TOKEN(anon_sym_COLON);
      END_STATE();
    case 330:
      ACCEPT_TOKEN(anon_sym_COLON);
      if (lookahead == ':') ADVANCE(373);
      END_STATE();
    case 331:
      ACCEPT_TOKEN(anon_sym_LBRACK);
      END_STATE();
    case 332:
      ACCEPT_TOKEN(anon_sym_RBRACK);
      END_STATE();
    case 333:
      ACCEPT_TOKEN(anon_sym_DOT);
      END_STATE();
    case 334:
      ACCEPT_TOKEN(anon_sym_DOT);
      if (lookahead == '.') ADVANCE(367);
      END_STATE();
    case 335:
      ACCEPT_TOKEN(anon_sym_DOT);
      if (lookahead == '.') ADVANCE(367);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(514);
      END_STATE();
    case 336:
      ACCEPT_TOKEN(anon_sym_DOT);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(514);
      END_STATE();
    case 337:
      ACCEPT_TOKEN(sym_wildcard_import);
      END_STATE();
    case 338:
      ACCEPT_TOKEN(sym_wildcard_import);
      if (lookahead == '=') ADVANCE(376);
      END_STATE();
    case 339:
      ACCEPT_TOKEN(anon_sym_EQ);
      END_STATE();
    case 340:
      ACCEPT_TOKEN(anon_sym_EQ);
      if (lookahead == '=') ADVANCE(381);
      END_STATE();
    case 341:
      ACCEPT_TOKEN(anon_sym_LBRACE);
      END_STATE();
    case 342:
      ACCEPT_TOKEN(anon_sym_RBRACE);
      END_STATE();
    case 343:
      ACCEPT_TOKEN(anon_sym_LPAREN);
      END_STATE();
    case 344:
      ACCEPT_TOKEN(anon_sym_COMMA);
      END_STATE();
    case 345:
      ACCEPT_TOKEN(anon_sym_RPAREN);
      END_STATE();
    case 346:
      ACCEPT_TOKEN(anon_sym_LT);
      END_STATE();
    case 347:
      ACCEPT_TOKEN(anon_sym_LT);
      if (lookahead == '=') ADVANCE(383);
      END_STATE();
    case 348:
      ACCEPT_TOKEN(anon_sym_GT);
      END_STATE();
    case 349:
      ACCEPT_TOKEN(anon_sym_GT);
      if (lookahead == '=') ADVANCE(384);
      END_STATE();
    case 350:
      ACCEPT_TOKEN(anon_sym_SEMI);
      END_STATE();
    case 351:
      ACCEPT_TOKEN(anon_sym_get);
      END_STATE();
    case 352:
      ACCEPT_TOKEN(anon_sym_get);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 353:
      ACCEPT_TOKEN(anon_sym_get);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 354:
      ACCEPT_TOKEN(anon_sym_get);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 355:
      ACCEPT_TOKEN(anon_sym_set);
      END_STATE();
    case 356:
      ACCEPT_TOKEN(anon_sym_set);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 357:
      ACCEPT_TOKEN(anon_sym_set);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 358:
      ACCEPT_TOKEN(anon_sym_set);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 359:
      ACCEPT_TOKEN(anon_sym_AMP);
      END_STATE();
    case 360:
      ACCEPT_TOKEN(anon_sym_AMP);
      if (lookahead == '&') ADVANCE(369);
      END_STATE();
    case 361:
      ACCEPT_TOKEN(sym__quest);
      END_STATE();
    case 362:
      ACCEPT_TOKEN(sym__quest);
      if (lookahead == ':') ADVANCE(368);
      END_STATE();
    case 363:
      ACCEPT_TOKEN(anon_sym_STAR);
      END_STATE();
    case 364:
      ACCEPT_TOKEN(anon_sym_STAR);
      if (lookahead == '=') ADVANCE(376);
      END_STATE();
    case 365:
      ACCEPT_TOKEN(anon_sym_DASH_GT);
      END_STATE();
    case 366:
      ACCEPT_TOKEN(sym_label);
      END_STATE();
    case 367:
      ACCEPT_TOKEN(anon_sym_DOT_DOT);
      END_STATE();
    case 368:
      ACCEPT_TOKEN(anon_sym_QMARK_COLON);
      END_STATE();
    case 369:
      ACCEPT_TOKEN(anon_sym_AMP_AMP);
      END_STATE();
    case 370:
      ACCEPT_TOKEN(anon_sym_PIPE_PIPE);
      END_STATE();
    case 371:
      ACCEPT_TOKEN(anon_sym_DOLLAR_LBRACE);
      END_STATE();
    case 372:
      ACCEPT_TOKEN(anon_sym_DOLLAR);
      if (lookahead == '{') ADVANCE(371);
      END_STATE();
    case 373:
      ACCEPT_TOKEN(anon_sym_COLON_COLON);
      END_STATE();
    case 374:
      ACCEPT_TOKEN(anon_sym_PLUS_EQ);
      END_STATE();
    case 375:
      ACCEPT_TOKEN(anon_sym_DASH_EQ);
      END_STATE();
    case 376:
      ACCEPT_TOKEN(anon_sym_STAR_EQ);
      END_STATE();
    case 377:
      ACCEPT_TOKEN(anon_sym_SLASH_EQ);
      END_STATE();
    case 378:
      ACCEPT_TOKEN(anon_sym_PERCENT_EQ);
      END_STATE();
    case 379:
      ACCEPT_TOKEN(anon_sym_BANG_EQ);
      if (lookahead == '=') ADVANCE(380);
      END_STATE();
    case 380:
      ACCEPT_TOKEN(anon_sym_BANG_EQ_EQ);
      END_STATE();
    case 381:
      ACCEPT_TOKEN(anon_sym_EQ_EQ);
      if (lookahead == '=') ADVANCE(382);
      END_STATE();
    case 382:
      ACCEPT_TOKEN(anon_sym_EQ_EQ_EQ);
      END_STATE();
    case 383:
      ACCEPT_TOKEN(anon_sym_LT_EQ);
      END_STATE();
    case 384:
      ACCEPT_TOKEN(anon_sym_GT_EQ);
      END_STATE();
    case 385:
      ACCEPT_TOKEN(anon_sym_BANGin);
      END_STATE();
    case 386:
      ACCEPT_TOKEN(anon_sym_BANGis);
      END_STATE();
    case 387:
      ACCEPT_TOKEN(anon_sym_PLUS);
      if (lookahead == '+') ADVANCE(396);
      END_STATE();
    case 388:
      ACCEPT_TOKEN(anon_sym_PLUS);
      if (lookahead == '+') ADVANCE(396);
      if (lookahead == '=') ADVANCE(374);
      END_STATE();
    case 389:
      ACCEPT_TOKEN(anon_sym_DASH);
      if (lookahead == '-') ADVANCE(397);
      END_STATE();
    case 390:
      ACCEPT_TOKEN(anon_sym_DASH);
      if (lookahead == '-') ADVANCE(397);
      if (lookahead == '=') ADVANCE(375);
      END_STATE();
    case 391:
      ACCEPT_TOKEN(anon_sym_DASH);
      if (lookahead == '-') ADVANCE(397);
      if (lookahead == '=') ADVANCE(375);
      if (lookahead == '>') ADVANCE(365);
      END_STATE();
    case 392:
      ACCEPT_TOKEN(anon_sym_DASH);
      if (lookahead == '-') ADVANCE(397);
      if (lookahead == '>') ADVANCE(365);
      END_STATE();
    case 393:
      ACCEPT_TOKEN(anon_sym_SLASH);
      if (lookahead == '/') ADVANCE(506);
      if (lookahead == '=') ADVANCE(377);
      END_STATE();
    case 394:
      ACCEPT_TOKEN(anon_sym_PERCENT);
      if (lookahead == '=') ADVANCE(378);
      END_STATE();
    case 395:
      ACCEPT_TOKEN(anon_sym_as_QMARK);
      END_STATE();
    case 396:
      ACCEPT_TOKEN(anon_sym_PLUS_PLUS);
      END_STATE();
    case 397:
      ACCEPT_TOKEN(anon_sym_DASH_DASH);
      END_STATE();
    case 398:
      ACCEPT_TOKEN(anon_sym_BANG);
      END_STATE();
    case 399:
      ACCEPT_TOKEN(anon_sym_BANG);
      if (lookahead == '!') ADVANCE(401);
      if (lookahead == '=') ADVANCE(379);
      if (lookahead == 'i') ADVANCE(249);
      END_STATE();
    case 400:
      ACCEPT_TOKEN(anon_sym_BANG);
      if (lookahead == 'i') ADVANCE(249);
      END_STATE();
    case 401:
      ACCEPT_TOKEN(anon_sym_BANG_BANG);
      END_STATE();
    case 402:
      ACCEPT_TOKEN(anon_sym_suspend);
      END_STATE();
    case 403:
      ACCEPT_TOKEN(anon_sym_suspend);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 404:
      ACCEPT_TOKEN(anon_sym_suspend);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 405:
      ACCEPT_TOKEN(anon_sym_suspend);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 406:
      ACCEPT_TOKEN(anon_sym_sealed);
      END_STATE();
    case 407:
      ACCEPT_TOKEN(anon_sym_sealed);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 408:
      ACCEPT_TOKEN(anon_sym_sealed);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 409:
      ACCEPT_TOKEN(anon_sym_sealed);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 410:
      ACCEPT_TOKEN(anon_sym_annotation);
      END_STATE();
    case 411:
      ACCEPT_TOKEN(anon_sym_annotation);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 412:
      ACCEPT_TOKEN(anon_sym_annotation);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 413:
      ACCEPT_TOKEN(anon_sym_annotation);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 414:
      ACCEPT_TOKEN(anon_sym_data);
      END_STATE();
    case 415:
      ACCEPT_TOKEN(anon_sym_data);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 416:
      ACCEPT_TOKEN(anon_sym_data);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 417:
      ACCEPT_TOKEN(anon_sym_data);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 418:
      ACCEPT_TOKEN(anon_sym_inner);
      END_STATE();
    case 419:
      ACCEPT_TOKEN(anon_sym_inner);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 420:
      ACCEPT_TOKEN(anon_sym_inner);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 421:
      ACCEPT_TOKEN(anon_sym_inner);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 422:
      ACCEPT_TOKEN(anon_sym_value);
      END_STATE();
    case 423:
      ACCEPT_TOKEN(anon_sym_value);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 424:
      ACCEPT_TOKEN(anon_sym_value);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 425:
      ACCEPT_TOKEN(anon_sym_value);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 426:
      ACCEPT_TOKEN(anon_sym_override);
      END_STATE();
    case 427:
      ACCEPT_TOKEN(anon_sym_override);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 428:
      ACCEPT_TOKEN(anon_sym_override);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 429:
      ACCEPT_TOKEN(anon_sym_override);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 430:
      ACCEPT_TOKEN(anon_sym_lateinit);
      END_STATE();
    case 431:
      ACCEPT_TOKEN(anon_sym_lateinit);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 432:
      ACCEPT_TOKEN(anon_sym_lateinit);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 433:
      ACCEPT_TOKEN(anon_sym_lateinit);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 434:
      ACCEPT_TOKEN(anon_sym_public);
      END_STATE();
    case 435:
      ACCEPT_TOKEN(anon_sym_public);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 436:
      ACCEPT_TOKEN(anon_sym_public);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 437:
      ACCEPT_TOKEN(anon_sym_public);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 438:
      ACCEPT_TOKEN(anon_sym_private);
      END_STATE();
    case 439:
      ACCEPT_TOKEN(anon_sym_private);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 440:
      ACCEPT_TOKEN(anon_sym_private);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 441:
      ACCEPT_TOKEN(anon_sym_private);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 442:
      ACCEPT_TOKEN(anon_sym_internal);
      END_STATE();
    case 443:
      ACCEPT_TOKEN(anon_sym_internal);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 444:
      ACCEPT_TOKEN(anon_sym_internal);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 445:
      ACCEPT_TOKEN(anon_sym_internal);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 446:
      ACCEPT_TOKEN(anon_sym_protected);
      END_STATE();
    case 447:
      ACCEPT_TOKEN(anon_sym_protected);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 448:
      ACCEPT_TOKEN(anon_sym_protected);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 449:
      ACCEPT_TOKEN(anon_sym_protected);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 450:
      ACCEPT_TOKEN(anon_sym_tailrec);
      END_STATE();
    case 451:
      ACCEPT_TOKEN(anon_sym_tailrec);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 452:
      ACCEPT_TOKEN(anon_sym_tailrec);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 453:
      ACCEPT_TOKEN(anon_sym_tailrec);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 454:
      ACCEPT_TOKEN(anon_sym_operator);
      END_STATE();
    case 455:
      ACCEPT_TOKEN(anon_sym_operator);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 456:
      ACCEPT_TOKEN(anon_sym_operator);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 457:
      ACCEPT_TOKEN(anon_sym_operator);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 458:
      ACCEPT_TOKEN(anon_sym_infix);
      END_STATE();
    case 459:
      ACCEPT_TOKEN(anon_sym_infix);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 460:
      ACCEPT_TOKEN(anon_sym_infix);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 461:
      ACCEPT_TOKEN(anon_sym_infix);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 462:
      ACCEPT_TOKEN(anon_sym_inline);
      END_STATE();
    case 463:
      ACCEPT_TOKEN(anon_sym_inline);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 464:
      ACCEPT_TOKEN(anon_sym_inline);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 465:
      ACCEPT_TOKEN(anon_sym_inline);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 466:
      ACCEPT_TOKEN(anon_sym_external);
      END_STATE();
    case 467:
      ACCEPT_TOKEN(anon_sym_external);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 468:
      ACCEPT_TOKEN(anon_sym_external);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 469:
      ACCEPT_TOKEN(anon_sym_external);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 470:
      ACCEPT_TOKEN(sym_property_modifier);
      END_STATE();
    case 471:
      ACCEPT_TOKEN(sym_property_modifier);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 472:
      ACCEPT_TOKEN(sym_property_modifier);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 473:
      ACCEPT_TOKEN(sym_property_modifier);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 474:
      ACCEPT_TOKEN(anon_sym_abstract);
      END_STATE();
    case 475:
      ACCEPT_TOKEN(anon_sym_abstract);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 476:
      ACCEPT_TOKEN(anon_sym_abstract);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 477:
      ACCEPT_TOKEN(anon_sym_abstract);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 478:
      ACCEPT_TOKEN(anon_sym_final);
      END_STATE();
    case 479:
      ACCEPT_TOKEN(anon_sym_final);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 480:
      ACCEPT_TOKEN(anon_sym_final);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 481:
      ACCEPT_TOKEN(anon_sym_final);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 482:
      ACCEPT_TOKEN(anon_sym_open);
      END_STATE();
    case 483:
      ACCEPT_TOKEN(anon_sym_open);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 484:
      ACCEPT_TOKEN(anon_sym_open);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 485:
      ACCEPT_TOKEN(anon_sym_open);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 486:
      ACCEPT_TOKEN(anon_sym_vararg);
      END_STATE();
    case 487:
      ACCEPT_TOKEN(anon_sym_vararg);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 488:
      ACCEPT_TOKEN(anon_sym_vararg);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 489:
      ACCEPT_TOKEN(anon_sym_vararg);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 490:
      ACCEPT_TOKEN(anon_sym_noinline);
      END_STATE();
    case 491:
      ACCEPT_TOKEN(anon_sym_noinline);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 492:
      ACCEPT_TOKEN(anon_sym_noinline);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 493:
      ACCEPT_TOKEN(anon_sym_noinline);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 494:
      ACCEPT_TOKEN(anon_sym_crossinline);
      END_STATE();
    case 495:
      ACCEPT_TOKEN(anon_sym_crossinline);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 496:
      ACCEPT_TOKEN(anon_sym_crossinline);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 497:
      ACCEPT_TOKEN(anon_sym_crossinline);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 498:
      ACCEPT_TOKEN(anon_sym_expect);
      END_STATE();
    case 499:
      ACCEPT_TOKEN(anon_sym_expect);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 500:
      ACCEPT_TOKEN(anon_sym_expect);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 501:
      ACCEPT_TOKEN(anon_sym_expect);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 502:
      ACCEPT_TOKEN(anon_sym_actual);
      END_STATE();
    case 503:
      ACCEPT_TOKEN(anon_sym_actual);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 504:
      ACCEPT_TOKEN(anon_sym_actual);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(170);
      END_STATE();
    case 505:
      ACCEPT_TOKEN(anon_sym_actual);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 506:
      ACCEPT_TOKEN(sym_line_comment);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(506);
      END_STATE();
    case 507:
      ACCEPT_TOKEN(anon_sym_return_AT);
      END_STATE();
    case 508:
      ACCEPT_TOKEN(anon_sym_continue_AT);
      END_STATE();
    case 509:
      ACCEPT_TOKEN(anon_sym_break_AT);
      END_STATE();
    case 510:
      ACCEPT_TOKEN(anon_sym_this_AT);
      END_STATE();
    case 511:
      ACCEPT_TOKEN(anon_sym_super_AT);
      END_STATE();
    case 512:
      ACCEPT_TOKEN(anon_sym_AT2);
      END_STATE();
    case 513:
      ACCEPT_TOKEN(sym_real_literal);
      END_STATE();
    case 514:
      ACCEPT_TOKEN(sym_real_literal);
      if (lookahead == '_') ADVANCE(174);
      if (lookahead == 'E' ||
          lookahead == 'e') ADVANCE(311);
      if (lookahead == 'F' ||
          lookahead == 'f') ADVANCE(513);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(514);
      END_STATE();
    case 515:
      ACCEPT_TOKEN(sym_real_literal);
      if (lookahead == '_') ADVANCE(175);
      if (lookahead == 'F' ||
          lookahead == 'f') ADVANCE(513);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(515);
      END_STATE();
    case 516:
      ACCEPT_TOKEN(sym_integer_literal);
      ADVANCE_MAP(
        '.', 313,
        '_', 171,
        'B', 312,
        'b', 312,
        'E', 311,
        'e', 311,
        'F', 513,
        'f', 513,
        'X', 315,
        'x', 315,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(517);
      END_STATE();
    case 517:
      ACCEPT_TOKEN(sym_integer_literal);
      if (lookahead == '.') ADVANCE(313);
      if (lookahead == '_') ADVANCE(171);
      if (lookahead == 'E' ||
          lookahead == 'e') ADVANCE(311);
      if (lookahead == 'F' ||
          lookahead == 'f') ADVANCE(513);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(517);
      END_STATE();
    case 518:
      ACCEPT_TOKEN(sym_hex_literal);
      if (lookahead == '_') ADVANCE(173);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(518);
      END_STATE();
    case 519:
      ACCEPT_TOKEN(sym_bin_literal);
      if (lookahead == '_') ADVANCE(172);
      END_STATE();
    case 520:
      ACCEPT_TOKEN(anon_sym_SQUOTE);
      END_STATE();
    case 521:
      ACCEPT_TOKEN(aux_sym_character_literal_token1);
      END_STATE();
    case 522:
      ACCEPT_TOKEN(aux_sym_character_literal_token1);
      if (lookahead == '/') ADVANCE(506);
      END_STATE();
    case 523:
      ACCEPT_TOKEN(aux_sym_character_literal_token1);
      if (lookahead == '/') ADVANCE(522);
      if (lookahead == '\t' ||
          lookahead == 0x0b ||
          lookahead == '\f' ||
          lookahead == ' ') ADVANCE(523);
      if (lookahead != 0 &&
          (lookahead < '\t' || '\r' < lookahead) &&
          lookahead != '\'' &&
          lookahead != '\\') ADVANCE(521);
      END_STATE();
    case 524:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '?') ADVANCE(395);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 525:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '?') ADVANCE(395);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 526:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(510);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 527:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(510);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 528:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(509);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 529:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(509);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 530:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(511);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 531:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(511);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 532:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(507);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 533:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(507);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 534:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(508);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 535:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(508);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 536:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(620);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 537:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(615);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 538:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(415);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 539:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(604);
      if (lookahead == 'h') ADVANCE(603);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 540:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(604);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 541:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(691);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 542:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(623);
      if (lookahead == 't') ADVANCE(356);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 543:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(694);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 544:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(616);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 545:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(617);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 546:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(618);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 547:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(664);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 548:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(619);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 549:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(692);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 550:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(621);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 551:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(693);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 552:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(562);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 553:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'a') ADVANCE(695);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 554:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'b') ADVANCE(678);
      if (lookahead == 'c') ADVANCE(682);
      if (lookahead == 'n') ADVANCE(631);
      if (lookahead == 's') ADVANCE(524);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 555:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'b') ADVANCE(678);
      if (lookahead == 'c') ADVANCE(682);
      if (lookahead == 'n') ADVANCE(631);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 556:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'b') ADVANCE(622);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 557:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'c') ADVANCE(682);
      if (lookahead == 's') ADVANCE(524);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 558:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'c') ADVANCE(682);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 559:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'c') ADVANCE(435);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 560:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'c') ADVANCE(451);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 561:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'c') ADVANCE(686);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 562:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'c') ADVANCE(687);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 563:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'c') ADVANCE(696);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 564:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'd') ADVANCE(407);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 565:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'd') ADVANCE(403);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 566:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'd') ADVANCE(447);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 567:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'd') ADVANCE(577);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 568:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(542);
      if (lookahead == 'u') ADVANCE(659);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 569:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(542);
      if (lookahead == 'u') ADVANCE(675);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 570:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(629);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 571:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(561);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 572:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(564);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 573:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(423);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 574:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(463);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 575:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(439);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 576:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(491);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 577:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(427);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 578:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(495);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 579:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(668);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 580:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(537);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 581:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(683);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 582:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(560);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 583:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(673);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 584:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(566);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 585:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(663);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 586:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(534);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 587:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(667);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 588:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(632);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 589:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(684);
      if (lookahead == 'u') ADVANCE(659);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 590:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(684);
      if (lookahead == 'u') ADVANCE(660);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 591:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(684);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 592:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(689);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 593:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(563);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 594:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(610);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 595:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'e') ADVANCE(674);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 596:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'f') ADVANCE(600);
      if (lookahead == 'l') ADVANCE(608);
      if (lookahead == 'n') ADVANCE(585);
      if (lookahead == 't') ADVANCE(595);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 597:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'g') ADVANCE(487);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 598:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'h') ADVANCE(603);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 599:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(704);
      if (lookahead == 'o') ADVANCE(697);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 600:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(706);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 601:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(559);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 602:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(567);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 603:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(676);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 604:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(624);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 605:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(633);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 606:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(635);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 607:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(652);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 608:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(640);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 609:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(645);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 610:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(636);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 611:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(688);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 612:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(642);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 613:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(643);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 614:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'i') ADVANCE(647);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 615:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'k') ADVANCE(528);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 616:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'l') ADVANCE(479);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 617:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'l') ADVANCE(503);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 618:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'l') ADVANCE(467);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 619:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'l') ADVANCE(443);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 620:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'l') ADVANCE(701);
      if (lookahead == 'r') ADVANCE(547);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 621:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'l') ADVANCE(701);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 622:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'l') ADVANCE(601);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 623:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'l') ADVANCE(572);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 624:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'l') ADVANCE(672);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 625:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'l') ADVANCE(612);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 626:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'l') ADVANCE(613);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 627:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(596);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 628:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(679);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 629:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(483);
      if (lookahead == 'r') ADVANCE(549);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 630:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(411);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 631:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(650);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 632:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(565);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 633:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(544);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 634:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(532);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 635:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(625);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 636:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(611);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 637:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(585);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 638:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(546);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 639:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(548);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 640:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(574);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 641:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(698);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 642:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(576);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 643:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(578);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 644:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(680);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 645:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(702);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 646:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(637);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 647:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'n') ADVANCE(626);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 648:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'o') ADVANCE(677);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 649:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'o') ADVANCE(628);
      if (lookahead == 'r') ADVANCE(648);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 650:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'o') ADVANCE(699);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 651:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'o') ADVANCE(665);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 652:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'o') ADVANCE(630);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 653:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'o') ADVANCE(641);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 654:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'o') ADVANCE(644);
      if (lookahead == 'r') ADVANCE(648);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 655:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'o') ADVANCE(606);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 656:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'p') ADVANCE(570);
      if (lookahead == 'v') ADVANCE(579);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 657:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'p') ADVANCE(571);
      if (lookahead == 't') ADVANCE(583);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 658:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'p') ADVANCE(571);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 659:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'p') ADVANCE(587);
      if (lookahead == 's') ADVANCE(661);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 660:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'p') ADVANCE(587);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 661:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'p') ADVANCE(588);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 662:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'r') ADVANCE(599);
      if (lookahead == 'u') ADVANCE(556);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 663:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'r') ADVANCE(419);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 664:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'r') ADVANCE(597);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 665:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'r') ADVANCE(455);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 666:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'r') ADVANCE(580);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 667:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'r') ADVANCE(530);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 668:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'r') ADVANCE(669);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 669:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'r') ADVANCE(602);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 670:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'r') ADVANCE(552);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 671:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'r') ADVANCE(634);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 672:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'r') ADVANCE(582);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 673:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'r') ADVANCE(638);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 674:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'r') ADVANCE(639);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 675:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 's') ADVANCE(661);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 676:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 's') ADVANCE(526);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 677:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 's') ADVANCE(681);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 678:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 's') ADVANCE(690);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 679:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 's') ADVANCE(685);
      if (lookahead == 't') ADVANCE(609);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 680:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 's') ADVANCE(685);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 681:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 's') ADVANCE(614);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 682:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(700);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 683:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(352);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 684:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(356);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 685:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(471);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 686:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(499);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 687:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(475);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 688:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(431);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 689:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(703);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 690:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(670);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 691:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(538);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 692:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(651);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 693:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(607);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 694:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(594);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 695:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(575);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 696:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(584);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 697:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(593);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 698:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(609);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 699:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 't') ADVANCE(551);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 700:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'u') ADVANCE(545);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 701:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'u') ADVANCE(573);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 702:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'u') ADVANCE(586);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 703:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'u') ADVANCE(671);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 704:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'v') ADVANCE(553);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 705:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'x') ADVANCE(657);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 706:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'x') ADVANCE(459);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 707:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (lookahead == 'x') ADVANCE(658);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 708:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == '@') ADVANCE(366);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(708);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 709:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'a') ADVANCE(792);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 710:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'a') ADVANCE(787);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 711:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'a') ADVANCE(417);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 712:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'a') ADVANCE(776);
      if (lookahead == 'h') ADVANCE(775);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 713:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'a') ADVANCE(776);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 714:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'a') ADVANCE(861);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 715:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'a') ADVANCE(795);
      if (lookahead == 't') ADVANCE(358);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 716:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'a') ADVANCE(795);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 717:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'a') ADVANCE(864);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 718:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'a') ADVANCE(788);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 719:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'a') ADVANCE(789);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 720:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'a') ADVANCE(790);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 721:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'a') ADVANCE(833);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 722:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'a') ADVANCE(791);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 723:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'a') ADVANCE(862);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 724:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'a') ADVANCE(793);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 725:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'a') ADVANCE(863);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 726:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'a') ADVANCE(735);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 727:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'a') ADVANCE(865);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 728:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'b') ADVANCE(848);
      if (lookahead == 'c') ADVANCE(852);
      if (lookahead == 'n') ADVANCE(803);
      if (lookahead == 's') ADVANCE(525);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 729:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'b') ADVANCE(848);
      if (lookahead == 'c') ADVANCE(852);
      if (lookahead == 'n') ADVANCE(803);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 730:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'b') ADVANCE(794);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 731:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'c') ADVANCE(852);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 732:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'c') ADVANCE(437);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 733:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'c') ADVANCE(453);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 734:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'c') ADVANCE(856);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 735:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'c') ADVANCE(857);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 736:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'c') ADVANCE(866);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 737:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'd') ADVANCE(409);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 738:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'd') ADVANCE(405);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 739:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'd') ADVANCE(449);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 740:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'd') ADVANCE(750);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 741:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(715);
      if (lookahead == 'u') ADVANCE(829);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 742:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(715);
      if (lookahead == 'u') ADVANCE(845);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 743:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(801);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 744:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(734);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 745:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(737);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 746:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(425);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 747:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(465);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 748:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(441);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 749:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(493);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 750:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(429);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 751:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(497);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 752:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(838);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 753:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(710);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 754:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(853);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 755:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(733);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 756:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(843);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 757:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(739);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 758:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(832);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 759:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(535);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 760:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(837);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 761:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(804);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 762:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(716);
      if (lookahead == 'u') ADVANCE(845);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 763:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(854);
      if (lookahead == 'u') ADVANCE(845);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 764:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(854);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 765:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(859);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 766:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(736);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 767:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(782);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 768:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'e') ADVANCE(844);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 769:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'f') ADVANCE(772);
      if (lookahead == 'l') ADVANCE(780);
      if (lookahead == 'n') ADVANCE(758);
      if (lookahead == 't') ADVANCE(768);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 770:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'g') ADVANCE(489);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 771:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'i') ADVANCE(873);
      if (lookahead == 'o') ADVANCE(867);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 772:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'i') ADVANCE(875);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 773:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'i') ADVANCE(732);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 774:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'i') ADVANCE(740);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 775:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'i') ADVANCE(846);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 776:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'i') ADVANCE(796);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 777:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'i') ADVANCE(806);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 778:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'i') ADVANCE(807);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 779:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'i') ADVANCE(823);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 780:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'i') ADVANCE(812);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 781:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'i') ADVANCE(816);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 782:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'i') ADVANCE(808);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 783:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'i') ADVANCE(858);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 784:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'i') ADVANCE(813);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 785:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'i') ADVANCE(814);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 786:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'i') ADVANCE(818);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 787:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'k') ADVANCE(529);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 788:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'l') ADVANCE(481);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 789:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'l') ADVANCE(505);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 790:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'l') ADVANCE(469);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 791:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'l') ADVANCE(445);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 792:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'l') ADVANCE(870);
      if (lookahead == 'r') ADVANCE(721);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 793:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'l') ADVANCE(870);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 794:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'l') ADVANCE(773);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 795:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'l') ADVANCE(745);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 796:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'l') ADVANCE(842);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 797:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'l') ADVANCE(784);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 798:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'l') ADVANCE(785);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 799:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'n') ADVANCE(769);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 800:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'n') ADVANCE(849);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 801:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'n') ADVANCE(485);
      if (lookahead == 'r') ADVANCE(723);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 802:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'n') ADVANCE(413);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 803:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'n') ADVANCE(821);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 804:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'n') ADVANCE(738);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 805:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'n') ADVANCE(533);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 806:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'n') ADVANCE(718);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 807:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'n') ADVANCE(797);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 808:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'n') ADVANCE(783);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 809:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'n') ADVANCE(758);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 810:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'n') ADVANCE(720);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 811:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'n') ADVANCE(722);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 812:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'n') ADVANCE(747);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 813:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'n') ADVANCE(749);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 814:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'n') ADVANCE(751);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 815:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'n') ADVANCE(850);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 816:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'n') ADVANCE(871);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 817:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'n') ADVANCE(809);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 818:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'n') ADVANCE(798);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 819:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'o') ADVANCE(847);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 820:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'o') ADVANCE(800);
      if (lookahead == 'r') ADVANCE(819);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 821:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'o') ADVANCE(868);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 822:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'o') ADVANCE(834);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 823:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'o') ADVANCE(802);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 824:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'o') ADVANCE(815);
      if (lookahead == 'r') ADVANCE(819);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 825:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'o') ADVANCE(778);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 826:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'p') ADVANCE(743);
      if (lookahead == 'v') ADVANCE(752);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 827:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'p') ADVANCE(744);
      if (lookahead == 't') ADVANCE(756);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 828:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'p') ADVANCE(744);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 829:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'p') ADVANCE(760);
      if (lookahead == 's') ADVANCE(830);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 830:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'p') ADVANCE(761);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 831:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'r') ADVANCE(771);
      if (lookahead == 'u') ADVANCE(730);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 832:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'r') ADVANCE(421);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 833:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'r') ADVANCE(770);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 834:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'r') ADVANCE(457);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 835:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'r') ADVANCE(753);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 836:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'r') ADVANCE(819);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 837:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'r') ADVANCE(531);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 838:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'r') ADVANCE(839);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 839:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'r') ADVANCE(774);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 840:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'r') ADVANCE(726);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 841:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'r') ADVANCE(805);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 842:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'r') ADVANCE(755);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 843:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'r') ADVANCE(810);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 844:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'r') ADVANCE(811);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 845:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 's') ADVANCE(830);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 846:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 's') ADVANCE(527);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 847:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 's') ADVANCE(851);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 848:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 's') ADVANCE(860);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 849:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 's') ADVANCE(855);
      if (lookahead == 't') ADVANCE(781);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 850:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 's') ADVANCE(855);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 851:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 's') ADVANCE(786);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 852:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 't') ADVANCE(869);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 853:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 't') ADVANCE(354);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 854:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 't') ADVANCE(358);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 855:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 't') ADVANCE(473);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 856:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 't') ADVANCE(501);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 857:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 't') ADVANCE(477);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 858:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 't') ADVANCE(433);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 859:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 't') ADVANCE(872);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 860:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 't') ADVANCE(840);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 861:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 't') ADVANCE(711);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 862:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 't') ADVANCE(822);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 863:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 't') ADVANCE(779);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 864:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 't') ADVANCE(767);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 865:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 't') ADVANCE(748);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 866:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 't') ADVANCE(757);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 867:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 't') ADVANCE(766);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 868:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 't') ADVANCE(725);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 869:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'u') ADVANCE(719);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 870:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'u') ADVANCE(746);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 871:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'u') ADVANCE(759);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 872:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'u') ADVANCE(841);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 873:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'v') ADVANCE(727);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 874:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'x') ADVANCE(827);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 875:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'x') ADVANCE(461);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 876:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (lookahead == 'x') ADVANCE(828);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 877:
      ACCEPT_TOKEN(sym__alpha_identifier);
      if (set_contains(sym__alpha_identifier_character_set_2, 453, lookahead)) ADVANCE(877);
      END_STATE();
    case 878:
      ACCEPT_TOKEN(sym__backtick_identifier);
      END_STATE();
    case 879:
      ACCEPT_TOKEN(anon_sym_BSLASHu);
      END_STATE();
    case 880:
      ACCEPT_TOKEN(aux_sym__uni_character_literal_token1);
      END_STATE();
    case 881:
      ACCEPT_TOKEN(sym__escaped_identifier);
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
        'L', 1,
        'a', 2,
        'b', 3,
        'c', 4,
        'd', 5,
        'e', 6,
        'f', 7,
        'i', 8,
        'n', 9,
        'o', 10,
        'p', 11,
        'r', 12,
        's', 13,
        't', 14,
        'v', 15,
        'w', 16,
        'U', 17,
        'u', 17,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(0);
      END_STATE();
    case 1:
      ACCEPT_TOKEN(anon_sym_L);
      END_STATE();
    case 2:
      if (lookahead == 's') ADVANCE(18);
      END_STATE();
    case 3:
      if (lookahead == 'r') ADVANCE(19);
      if (lookahead == 'y') ADVANCE(20);
      END_STATE();
    case 4:
      if (lookahead == 'a') ADVANCE(21);
      if (lookahead == 'l') ADVANCE(22);
      if (lookahead == 'o') ADVANCE(23);
      END_STATE();
    case 5:
      if (lookahead == 'e') ADVANCE(24);
      if (lookahead == 'o') ADVANCE(25);
      if (lookahead == 'y') ADVANCE(26);
      END_STATE();
    case 6:
      if (lookahead == 'l') ADVANCE(27);
      if (lookahead == 'n') ADVANCE(28);
      END_STATE();
    case 7:
      if (lookahead == 'a') ADVANCE(29);
      if (lookahead == 'i') ADVANCE(30);
      if (lookahead == 'o') ADVANCE(31);
      if (lookahead == 'u') ADVANCE(32);
      END_STATE();
    case 8:
      if (lookahead == 'f') ADVANCE(33);
      if (lookahead == 'm') ADVANCE(34);
      if (lookahead == 'n') ADVANCE(35);
      if (lookahead == 's') ADVANCE(36);
      END_STATE();
    case 9:
      if (lookahead == 'u') ADVANCE(37);
      END_STATE();
    case 10:
      if (lookahead == 'b') ADVANCE(38);
      if (lookahead == 'u') ADVANCE(39);
      END_STATE();
    case 11:
      if (lookahead == 'a') ADVANCE(40);
      if (lookahead == 'r') ADVANCE(41);
      END_STATE();
    case 12:
      if (lookahead == 'e') ADVANCE(42);
      END_STATE();
    case 13:
      if (lookahead == 'e') ADVANCE(43);
      if (lookahead == 'u') ADVANCE(44);
      END_STATE();
    case 14:
      if (lookahead == 'h') ADVANCE(45);
      if (lookahead == 'r') ADVANCE(46);
      if (lookahead == 'y') ADVANCE(47);
      END_STATE();
    case 15:
      if (lookahead == 'a') ADVANCE(48);
      END_STATE();
    case 16:
      if (lookahead == 'h') ADVANCE(49);
      END_STATE();
    case 17:
      ACCEPT_TOKEN(aux_sym_unsigned_literal_token1);
      if (lookahead == 'L') ADVANCE(50);
      END_STATE();
    case 18:
      ACCEPT_TOKEN(anon_sym_as);
      END_STATE();
    case 19:
      if (lookahead == 'e') ADVANCE(51);
      END_STATE();
    case 20:
      ACCEPT_TOKEN(anon_sym_by);
      END_STATE();
    case 21:
      if (lookahead == 't') ADVANCE(52);
      END_STATE();
    case 22:
      if (lookahead == 'a') ADVANCE(53);
      END_STATE();
    case 23:
      if (lookahead == 'm') ADVANCE(54);
      if (lookahead == 'n') ADVANCE(55);
      END_STATE();
    case 24:
      if (lookahead == 'l') ADVANCE(56);
      END_STATE();
    case 25:
      ACCEPT_TOKEN(anon_sym_do);
      END_STATE();
    case 26:
      if (lookahead == 'n') ADVANCE(57);
      END_STATE();
    case 27:
      if (lookahead == 's') ADVANCE(58);
      END_STATE();
    case 28:
      if (lookahead == 'u') ADVANCE(59);
      END_STATE();
    case 29:
      if (lookahead == 'l') ADVANCE(60);
      END_STATE();
    case 30:
      if (lookahead == 'e') ADVANCE(61);
      if (lookahead == 'l') ADVANCE(62);
      if (lookahead == 'n') ADVANCE(63);
      END_STATE();
    case 31:
      if (lookahead == 'r') ADVANCE(64);
      END_STATE();
    case 32:
      if (lookahead == 'n') ADVANCE(65);
      END_STATE();
    case 33:
      ACCEPT_TOKEN(anon_sym_if);
      END_STATE();
    case 34:
      if (lookahead == 'p') ADVANCE(66);
      END_STATE();
    case 35:
      ACCEPT_TOKEN(anon_sym_in);
      if (lookahead == 'i') ADVANCE(67);
      if (lookahead == 't') ADVANCE(68);
      END_STATE();
    case 36:
      ACCEPT_TOKEN(anon_sym_is);
      END_STATE();
    case 37:
      if (lookahead == 'l') ADVANCE(69);
      END_STATE();
    case 38:
      if (lookahead == 'j') ADVANCE(70);
      END_STATE();
    case 39:
      if (lookahead == 't') ADVANCE(71);
      END_STATE();
    case 40:
      if (lookahead == 'c') ADVANCE(72);
      if (lookahead == 'r') ADVANCE(73);
      END_STATE();
    case 41:
      if (lookahead == 'o') ADVANCE(74);
      END_STATE();
    case 42:
      if (lookahead == 'c') ADVANCE(75);
      if (lookahead == 'i') ADVANCE(76);
      if (lookahead == 't') ADVANCE(77);
      END_STATE();
    case 43:
      if (lookahead == 't') ADVANCE(78);
      END_STATE();
    case 44:
      if (lookahead == 'p') ADVANCE(79);
      END_STATE();
    case 45:
      if (lookahead == 'i') ADVANCE(80);
      if (lookahead == 'r') ADVANCE(81);
      END_STATE();
    case 46:
      if (lookahead == 'u') ADVANCE(82);
      if (lookahead == 'y') ADVANCE(83);
      END_STATE();
    case 47:
      if (lookahead == 'p') ADVANCE(84);
      END_STATE();
    case 48:
      if (lookahead == 'l') ADVANCE(85);
      if (lookahead == 'r') ADVANCE(86);
      END_STATE();
    case 49:
      if (lookahead == 'e') ADVANCE(87);
      if (lookahead == 'i') ADVANCE(88);
      END_STATE();
    case 50:
      ACCEPT_TOKEN(aux_sym_unsigned_literal_token1);
      END_STATE();
    case 51:
      if (lookahead == 'a') ADVANCE(89);
      END_STATE();
    case 52:
      if (lookahead == 'c') ADVANCE(90);
      END_STATE();
    case 53:
      if (lookahead == 's') ADVANCE(91);
      END_STATE();
    case 54:
      if (lookahead == 'p') ADVANCE(92);
      END_STATE();
    case 55:
      if (lookahead == 's') ADVANCE(93);
      if (lookahead == 't') ADVANCE(94);
      END_STATE();
    case 56:
      if (lookahead == 'e') ADVANCE(95);
      END_STATE();
    case 57:
      if (lookahead == 'a') ADVANCE(96);
      END_STATE();
    case 58:
      if (lookahead == 'e') ADVANCE(97);
      END_STATE();
    case 59:
      if (lookahead == 'm') ADVANCE(98);
      END_STATE();
    case 60:
      if (lookahead == 's') ADVANCE(99);
      END_STATE();
    case 61:
      if (lookahead == 'l') ADVANCE(100);
      END_STATE();
    case 62:
      if (lookahead == 'e') ADVANCE(101);
      END_STATE();
    case 63:
      if (lookahead == 'a') ADVANCE(102);
      END_STATE();
    case 64:
      ACCEPT_TOKEN(anon_sym_for);
      END_STATE();
    case 65:
      ACCEPT_TOKEN(anon_sym_fun);
      END_STATE();
    case 66:
      if (lookahead == 'o') ADVANCE(103);
      END_STATE();
    case 67:
      if (lookahead == 't') ADVANCE(104);
      END_STATE();
    case 68:
      if (lookahead == 'e') ADVANCE(105);
      END_STATE();
    case 69:
      if (lookahead == 'l') ADVANCE(106);
      END_STATE();
    case 70:
      if (lookahead == 'e') ADVANCE(107);
      END_STATE();
    case 71:
      ACCEPT_TOKEN(anon_sym_out);
      END_STATE();
    case 72:
      if (lookahead == 'k') ADVANCE(108);
      END_STATE();
    case 73:
      if (lookahead == 'a') ADVANCE(109);
      END_STATE();
    case 74:
      if (lookahead == 'p') ADVANCE(110);
      END_STATE();
    case 75:
      if (lookahead == 'e') ADVANCE(111);
      END_STATE();
    case 76:
      if (lookahead == 'f') ADVANCE(112);
      END_STATE();
    case 77:
      if (lookahead == 'u') ADVANCE(113);
      END_STATE();
    case 78:
      if (lookahead == 'p') ADVANCE(114);
      END_STATE();
    case 79:
      if (lookahead == 'e') ADVANCE(115);
      END_STATE();
    case 80:
      if (lookahead == 's') ADVANCE(116);
      END_STATE();
    case 81:
      if (lookahead == 'o') ADVANCE(117);
      END_STATE();
    case 82:
      if (lookahead == 'e') ADVANCE(118);
      END_STATE();
    case 83:
      ACCEPT_TOKEN(anon_sym_try);
      END_STATE();
    case 84:
      if (lookahead == 'e') ADVANCE(119);
      END_STATE();
    case 85:
      ACCEPT_TOKEN(anon_sym_val);
      END_STATE();
    case 86:
      ACCEPT_TOKEN(anon_sym_var);
      END_STATE();
    case 87:
      if (lookahead == 'n') ADVANCE(120);
      if (lookahead == 'r') ADVANCE(121);
      END_STATE();
    case 88:
      if (lookahead == 'l') ADVANCE(122);
      END_STATE();
    case 89:
      if (lookahead == 'k') ADVANCE(123);
      END_STATE();
    case 90:
      if (lookahead == 'h') ADVANCE(124);
      END_STATE();
    case 91:
      if (lookahead == 's') ADVANCE(125);
      END_STATE();
    case 92:
      if (lookahead == 'a') ADVANCE(126);
      END_STATE();
    case 93:
      if (lookahead == 't') ADVANCE(127);
      END_STATE();
    case 94:
      if (lookahead == 'i') ADVANCE(128);
      END_STATE();
    case 95:
      if (lookahead == 'g') ADVANCE(129);
      END_STATE();
    case 96:
      if (lookahead == 'm') ADVANCE(130);
      END_STATE();
    case 97:
      ACCEPT_TOKEN(anon_sym_else);
      END_STATE();
    case 98:
      ACCEPT_TOKEN(anon_sym_enum);
      END_STATE();
    case 99:
      if (lookahead == 'e') ADVANCE(131);
      END_STATE();
    case 100:
      if (lookahead == 'd') ADVANCE(132);
      END_STATE();
    case 101:
      ACCEPT_TOKEN(anon_sym_file);
      END_STATE();
    case 102:
      if (lookahead == 'l') ADVANCE(133);
      END_STATE();
    case 103:
      if (lookahead == 'r') ADVANCE(134);
      END_STATE();
    case 104:
      ACCEPT_TOKEN(anon_sym_init);
      END_STATE();
    case 105:
      if (lookahead == 'r') ADVANCE(135);
      END_STATE();
    case 106:
      ACCEPT_TOKEN(anon_sym_null);
      END_STATE();
    case 107:
      if (lookahead == 'c') ADVANCE(136);
      END_STATE();
    case 108:
      if (lookahead == 'a') ADVANCE(137);
      END_STATE();
    case 109:
      if (lookahead == 'm') ADVANCE(138);
      END_STATE();
    case 110:
      if (lookahead == 'e') ADVANCE(139);
      END_STATE();
    case 111:
      if (lookahead == 'i') ADVANCE(140);
      END_STATE();
    case 112:
      if (lookahead == 'i') ADVANCE(141);
      END_STATE();
    case 113:
      if (lookahead == 'r') ADVANCE(142);
      END_STATE();
    case 114:
      if (lookahead == 'a') ADVANCE(143);
      END_STATE();
    case 115:
      if (lookahead == 'r') ADVANCE(144);
      END_STATE();
    case 116:
      ACCEPT_TOKEN(anon_sym_this);
      END_STATE();
    case 117:
      if (lookahead == 'w') ADVANCE(145);
      END_STATE();
    case 118:
      ACCEPT_TOKEN(anon_sym_true);
      END_STATE();
    case 119:
      if (lookahead == 'a') ADVANCE(146);
      END_STATE();
    case 120:
      ACCEPT_TOKEN(anon_sym_when);
      END_STATE();
    case 121:
      if (lookahead == 'e') ADVANCE(147);
      END_STATE();
    case 122:
      if (lookahead == 'e') ADVANCE(148);
      END_STATE();
    case 123:
      ACCEPT_TOKEN(anon_sym_break);
      END_STATE();
    case 124:
      ACCEPT_TOKEN(anon_sym_catch);
      END_STATE();
    case 125:
      ACCEPT_TOKEN(anon_sym_class);
      END_STATE();
    case 126:
      if (lookahead == 'n') ADVANCE(149);
      END_STATE();
    case 127:
      if (lookahead == 'r') ADVANCE(150);
      END_STATE();
    case 128:
      if (lookahead == 'n') ADVANCE(151);
      END_STATE();
    case 129:
      if (lookahead == 'a') ADVANCE(152);
      END_STATE();
    case 130:
      if (lookahead == 'i') ADVANCE(153);
      END_STATE();
    case 131:
      ACCEPT_TOKEN(anon_sym_false);
      END_STATE();
    case 132:
      ACCEPT_TOKEN(anon_sym_field);
      END_STATE();
    case 133:
      if (lookahead == 'l') ADVANCE(154);
      END_STATE();
    case 134:
      if (lookahead == 't') ADVANCE(155);
      END_STATE();
    case 135:
      if (lookahead == 'f') ADVANCE(156);
      END_STATE();
    case 136:
      if (lookahead == 't') ADVANCE(157);
      END_STATE();
    case 137:
      if (lookahead == 'g') ADVANCE(158);
      END_STATE();
    case 138:
      ACCEPT_TOKEN(anon_sym_param);
      END_STATE();
    case 139:
      if (lookahead == 'r') ADVANCE(159);
      END_STATE();
    case 140:
      if (lookahead == 'v') ADVANCE(160);
      END_STATE();
    case 141:
      if (lookahead == 'e') ADVANCE(161);
      END_STATE();
    case 142:
      if (lookahead == 'n') ADVANCE(162);
      END_STATE();
    case 143:
      if (lookahead == 'r') ADVANCE(163);
      END_STATE();
    case 144:
      ACCEPT_TOKEN(anon_sym_super);
      END_STATE();
    case 145:
      ACCEPT_TOKEN(anon_sym_throw);
      END_STATE();
    case 146:
      if (lookahead == 'l') ADVANCE(164);
      END_STATE();
    case 147:
      ACCEPT_TOKEN(anon_sym_where);
      END_STATE();
    case 148:
      ACCEPT_TOKEN(anon_sym_while);
      END_STATE();
    case 149:
      if (lookahead == 'i') ADVANCE(165);
      END_STATE();
    case 150:
      if (lookahead == 'u') ADVANCE(166);
      END_STATE();
    case 151:
      if (lookahead == 'u') ADVANCE(167);
      END_STATE();
    case 152:
      if (lookahead == 't') ADVANCE(168);
      END_STATE();
    case 153:
      if (lookahead == 'c') ADVANCE(169);
      END_STATE();
    case 154:
      if (lookahead == 'y') ADVANCE(170);
      END_STATE();
    case 155:
      ACCEPT_TOKEN(anon_sym_import);
      END_STATE();
    case 156:
      if (lookahead == 'a') ADVANCE(171);
      END_STATE();
    case 157:
      ACCEPT_TOKEN(anon_sym_object);
      END_STATE();
    case 158:
      if (lookahead == 'e') ADVANCE(172);
      END_STATE();
    case 159:
      if (lookahead == 't') ADVANCE(173);
      END_STATE();
    case 160:
      if (lookahead == 'e') ADVANCE(174);
      END_STATE();
    case 161:
      if (lookahead == 'd') ADVANCE(175);
      END_STATE();
    case 162:
      ACCEPT_TOKEN(anon_sym_return);
      END_STATE();
    case 163:
      if (lookahead == 'a') ADVANCE(176);
      END_STATE();
    case 164:
      if (lookahead == 'i') ADVANCE(177);
      END_STATE();
    case 165:
      if (lookahead == 'o') ADVANCE(178);
      END_STATE();
    case 166:
      if (lookahead == 'c') ADVANCE(179);
      END_STATE();
    case 167:
      if (lookahead == 'e') ADVANCE(180);
      END_STATE();
    case 168:
      if (lookahead == 'e') ADVANCE(181);
      END_STATE();
    case 169:
      ACCEPT_TOKEN(anon_sym_dynamic);
      END_STATE();
    case 170:
      ACCEPT_TOKEN(anon_sym_finally);
      END_STATE();
    case 171:
      if (lookahead == 'c') ADVANCE(182);
      END_STATE();
    case 172:
      ACCEPT_TOKEN(anon_sym_package);
      END_STATE();
    case 173:
      if (lookahead == 'y') ADVANCE(183);
      END_STATE();
    case 174:
      if (lookahead == 'r') ADVANCE(184);
      END_STATE();
    case 175:
      ACCEPT_TOKEN(sym_reification_modifier);
      END_STATE();
    case 176:
      if (lookahead == 'm') ADVANCE(185);
      END_STATE();
    case 177:
      if (lookahead == 'a') ADVANCE(186);
      END_STATE();
    case 178:
      if (lookahead == 'n') ADVANCE(187);
      END_STATE();
    case 179:
      if (lookahead == 't') ADVANCE(188);
      END_STATE();
    case 180:
      ACCEPT_TOKEN(anon_sym_continue);
      END_STATE();
    case 181:
      ACCEPT_TOKEN(anon_sym_delegate);
      END_STATE();
    case 182:
      if (lookahead == 'e') ADVANCE(189);
      END_STATE();
    case 183:
      ACCEPT_TOKEN(anon_sym_property);
      END_STATE();
    case 184:
      ACCEPT_TOKEN(anon_sym_receiver);
      END_STATE();
    case 185:
      ACCEPT_TOKEN(anon_sym_setparam);
      END_STATE();
    case 186:
      if (lookahead == 's') ADVANCE(190);
      END_STATE();
    case 187:
      ACCEPT_TOKEN(anon_sym_companion);
      END_STATE();
    case 188:
      if (lookahead == 'o') ADVANCE(191);
      END_STATE();
    case 189:
      ACCEPT_TOKEN(anon_sym_interface);
      END_STATE();
    case 190:
      ACCEPT_TOKEN(anon_sym_typealias);
      END_STATE();
    case 191:
      if (lookahead == 'r') ADVANCE(192);
      END_STATE();
    case 192:
      ACCEPT_TOKEN(anon_sym_constructor);
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
static uint16_t ts_small_parse_table[229504];

// Filled from tables.zlib by grammar.go, see fetchgrammar.
static uint32_t ts_small_parse_table_map[5626];

// Filled from tables.zlib by grammar.go, see fetchgrammar.
static TSParseActionEntry ts_parse_actions[13092];

enum ts_external_scanner_symbol_identifiers {
  ts_external_token__automatic_semicolon = 0,
  ts_external_token__import_list_delimiter = 1,
  ts_external_token_safe_nav = 2,
  ts_external_token_multiline_comment = 3,
  ts_external_token__string_start = 4,
  ts_external_token__string_end = 5,
  ts_external_token_string_content = 6,
};

static const TSSymbol ts_external_scanner_symbol_map[EXTERNAL_TOKEN_COUNT] = {
  [ts_external_token__automatic_semicolon] = sym__automatic_semicolon,
  [ts_external_token__import_list_delimiter] = sym__import_list_delimiter,
  [ts_external_token_safe_nav] = sym_safe_nav,
  [ts_external_token_multiline_comment] = sym_multiline_comment,
  [ts_external_token__string_start] = sym__string_start,
  [ts_external_token__string_end] = sym__string_end,
  [ts_external_token_string_content] = sym_string_content,
};

static const bool ts_external_scanner_states[12][EXTERNAL_TOKEN_COUNT] = {
  [1] = {
    [ts_external_token__automatic_semicolon] = true,
    [ts_external_token__import_list_delimiter] = true,
    [ts_external_token_safe_nav] = true,
    [ts_external_token_multiline_comment] = true,
    [ts_external_token__string_start] = true,
    [ts_external_token__string_end] = true,
    [ts_external_token_string_content] = true,
  },
  [2] = {
    [ts_external_token_multiline_comment] = true,
    [ts_external_token__string_start] = true,
  },
  [3] = {
    [ts_external_token_safe_nav] = true,
    [ts_external_token_multiline_comment] = true,
    [ts_external_token__string_start] = true,
  },
  [4] = {
    [ts_external_token__automatic_semicolon] = true,
    [ts_external_token_safe_nav] = true,
    [ts_external_token_multiline_comment] = true,
    [ts_external_token__string_start] = true,
  },
  [5] = {
    [ts_external_token__automatic_semicolon] = true,
    [ts_external_token_multiline_comment] = true,
    [ts_external_token__string_start] = true,
  },
  [6] = {
    [ts_external_token_safe_nav] = true,
    [ts_external_token_multiline_comment] = true,
  },
  [7] = {
    [ts_external_token__automatic_semicolon] = true,
    [ts_external_token_safe_nav] = true,
    [ts_external_token_multiline_comment] = true,
  },
  [8] = {
    [ts_external_token_multiline_comment] = true,
  },
  [9] = {
    [ts_external_token__automatic_semicolon] = true,
    [ts_external_token_multiline_comment] = true,
  },
  [10] = {
    [ts_external_token_multiline_comment] = true,
    [ts_external_token__string_end] = true,
    [ts_external_token_string_content] = true,
  },
  [11] = {
    [ts_external_token__import_list_delimiter] = true,
    [ts_external_token_multiline_comment] = true,
  },
};

#ifdef __cplusplus
extern "C" {
#endif
void *tree_sitter_kotlin_external_scanner_create(void);
void tree_sitter_kotlin_external_scanner_destroy(void *);
bool tree_sitter_kotlin_external_scanner_scan(void *, TSLexer *, const bool *);
unsigned tree_sitter_kotlin_external_scanner_serialize(void *, char *);
void tree_sitter_kotlin_external_scanner_deserialize(void *, const char *, unsigned);

#ifdef TREE_SITTER_HIDE_SYMBOLS
#define TS_PUBLIC
#elif defined(_WIN32)
#define TS_PUBLIC __declspec(dllexport)
#else
#define TS_PUBLIC __attribute__((visibility("default")))
#endif

TS_PUBLIC const TSLanguage *tree_sitter_kotlin(void) {
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
    .symbol_metadata = ts_symbol_metadata,
    .public_symbol_map = ts_symbol_map,
    .alias_map = ts_non_terminal_alias_map,
    .alias_sequences = &ts_alias_sequences[0][0],
    .lex_modes = ts_lex_modes,
    .lex_fn = ts_lex,
    .keyword_lex_fn = ts_lex_keywords,
    .keyword_capture_token = sym__alpha_identifier,
    .external_scanner = {
      &ts_external_scanner_states[0][0],
      ts_external_scanner_symbol_map,
      tree_sitter_kotlin_external_scanner_create,
      tree_sitter_kotlin_external_scanner_destroy,
      tree_sitter_kotlin_external_scanner_scan,
      tree_sitter_kotlin_external_scanner_serialize,
      tree_sitter_kotlin_external_scanner_deserialize,
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
