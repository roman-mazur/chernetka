#include "tree_sitter/parser.h"

#if defined(__GNUC__) || defined(__clang__)
#pragma GCC diagnostic ignored "-Wmissing-field-initializers"
#endif

#define LANGUAGE_VERSION 14
#define STATE_COUNT 1385
#define LARGE_STATE_COUNT 406
#define SYMBOL_COUNT 320
#define ALIAS_COUNT 1
#define TOKEN_COUNT 138
#define EXTERNAL_TOKEN_COUNT 0
#define FIELD_COUNT 40
#define MAX_ALIAS_SEQUENCE_LENGTH 11
#define PRODUCTION_ID_COUNT 208

enum ts_symbol_identifiers {
  sym_identifier = 1,
  sym_decimal_integer_literal = 2,
  sym_hex_integer_literal = 3,
  sym_octal_integer_literal = 4,
  sym_binary_integer_literal = 5,
  sym_decimal_floating_point_literal = 6,
  sym_hex_floating_point_literal = 7,
  sym_true = 8,
  sym_false = 9,
  sym_character_literal = 10,
  anon_sym_DQUOTE = 11,
  anon_sym_DQUOTE_DQUOTE_DQUOTE = 12,
  sym_string_fragment = 13,
  aux_sym__multiline_string_fragment_token1 = 14,
  aux_sym__multiline_string_fragment_token2 = 15,
  anon_sym_BSLASH_LBRACE = 16,
  anon_sym_RBRACE = 17,
  aux_sym__escape_sequence_token1 = 18,
  sym_escape_sequence = 19,
  sym_null_literal = 20,
  anon_sym_LPAREN = 21,
  anon_sym_RPAREN = 22,
  anon_sym_AMP = 23,
  anon_sym_EQ = 24,
  anon_sym_PLUS_EQ = 25,
  anon_sym_DASH_EQ = 26,
  anon_sym_STAR_EQ = 27,
  anon_sym_SLASH_EQ = 28,
  anon_sym_AMP_EQ = 29,
  anon_sym_PIPE_EQ = 30,
  anon_sym_CARET_EQ = 31,
  anon_sym_PERCENT_EQ = 32,
  anon_sym_LT_LT_EQ = 33,
  anon_sym_GT_GT_EQ = 34,
  anon_sym_GT_GT_GT_EQ = 35,
  anon_sym_GT = 36,
  anon_sym_LT = 37,
  anon_sym_GT_EQ = 38,
  anon_sym_LT_EQ = 39,
  anon_sym_EQ_EQ = 40,
  anon_sym_BANG_EQ = 41,
  anon_sym_AMP_AMP = 42,
  anon_sym_PIPE_PIPE = 43,
  anon_sym_PLUS = 44,
  anon_sym_DASH = 45,
  anon_sym_STAR = 46,
  anon_sym_SLASH = 47,
  anon_sym_PIPE = 48,
  anon_sym_CARET = 49,
  anon_sym_PERCENT = 50,
  anon_sym_LT_LT = 51,
  anon_sym_GT_GT = 52,
  anon_sym_GT_GT_GT = 53,
  anon_sym_instanceof = 54,
  anon_sym_final = 55,
  anon_sym_DASH_GT = 56,
  anon_sym_COMMA = 57,
  anon_sym_QMARK = 58,
  anon_sym_COLON = 59,
  anon_sym_BANG = 60,
  anon_sym_TILDE = 61,
  anon_sym_PLUS_PLUS = 62,
  anon_sym_DASH_DASH = 63,
  anon_sym_new = 64,
  anon_sym_LBRACK = 65,
  anon_sym_RBRACK = 66,
  anon_sym_DOT = 67,
  anon_sym_class = 68,
  anon_sym_COLON_COLON = 69,
  anon_sym_extends = 70,
  anon_sym_switch = 71,
  anon_sym_LBRACE = 72,
  anon_sym_case = 73,
  anon_sym_default = 74,
  sym_underscore_pattern = 75,
  anon_sym_when = 76,
  anon_sym_SEMI = 77,
  anon_sym_assert = 78,
  anon_sym_do = 79,
  anon_sym_while = 80,
  anon_sym_break = 81,
  anon_sym_continue = 82,
  anon_sym_return = 83,
  anon_sym_yield = 84,
  anon_sym_synchronized = 85,
  anon_sym_throw = 86,
  anon_sym_try = 87,
  anon_sym_catch = 88,
  anon_sym_finally = 89,
  anon_sym_if = 90,
  anon_sym_else = 91,
  anon_sym_for = 92,
  anon_sym_AT = 93,
  anon_sym_open = 94,
  anon_sym_module = 95,
  anon_sym_requires = 96,
  anon_sym_transitive = 97,
  anon_sym_static = 98,
  anon_sym_exports = 99,
  anon_sym_to = 100,
  anon_sym_opens = 101,
  anon_sym_uses = 102,
  anon_sym_provides = 103,
  anon_sym_with = 104,
  anon_sym_package = 105,
  anon_sym_import = 106,
  anon_sym_enum = 107,
  anon_sym_public = 108,
  anon_sym_protected = 109,
  anon_sym_private = 110,
  anon_sym_abstract = 111,
  anon_sym_strictfp = 112,
  anon_sym_native = 113,
  anon_sym_transient = 114,
  anon_sym_volatile = 115,
  anon_sym_sealed = 116,
  anon_sym_non_DASHsealed = 117,
  anon_sym_implements = 118,
  anon_sym_permits = 119,
  anon_sym_record = 120,
  anon_sym_ATinterface = 121,
  anon_sym_interface = 122,
  anon_sym_byte = 123,
  anon_sym_short = 124,
  anon_sym_int = 125,
  anon_sym_long = 126,
  anon_sym_char = 127,
  anon_sym_float = 128,
  anon_sym_double = 129,
  sym_boolean_type = 130,
  sym_void_type = 131,
  anon_sym_DOT_DOT_DOT = 132,
  anon_sym_throws = 133,
  sym_this = 134,
  sym_super = 135,
  sym_line_comment = 136,
  sym_block_comment = 137,
  sym_program = 138,
  sym__toplevel_statement = 139,
  sym__literal = 140,
  sym_string_literal = 141,
  sym__string_literal = 142,
  sym__multiline_string_literal = 143,
  sym__multiline_string_fragment = 144,
  sym_string_interpolation = 145,
  sym__escape_sequence = 146,
  sym_expression = 147,
  sym_cast_expression = 148,
  sym_assignment_expression = 149,
  sym_binary_expression = 150,
  sym_instanceof_expression = 151,
  sym_lambda_expression = 152,
  sym_inferred_parameters = 153,
  sym_ternary_expression = 154,
  sym_unary_expression = 155,
  sym_update_expression = 156,
  sym_primary_expression = 157,
  sym_array_creation_expression = 158,
  sym_dimensions_expr = 159,
  sym_parenthesized_expression = 160,
  sym_class_literal = 161,
  sym_object_creation_expression = 162,
  sym__unqualified_object_creation_expression = 163,
  sym_field_access = 164,
  sym_template_expression = 165,
  sym_array_access = 166,
  sym_method_invocation = 167,
  sym_argument_list = 168,
  sym_method_reference = 169,
  sym_type_arguments = 170,
  sym_wildcard = 171,
  sym__wildcard_bounds = 172,
  sym_dimensions = 173,
  sym_switch_expression = 174,
  sym_switch_block = 175,
  sym_switch_block_statement_group = 176,
  sym_switch_rule = 177,
  sym_switch_label = 178,
  sym_pattern = 179,
  sym_type_pattern = 180,
  sym_record_pattern = 181,
  sym_record_pattern_body = 182,
  sym_record_pattern_component = 183,
  sym_guard = 184,
  sym_statement = 185,
  sym_block = 186,
  sym_expression_statement = 187,
  sym_labeled_statement = 188,
  sym_assert_statement = 189,
  sym_do_statement = 190,
  sym_break_statement = 191,
  sym_continue_statement = 192,
  sym_return_statement = 193,
  sym_yield_statement = 194,
  sym_synchronized_statement = 195,
  sym_throw_statement = 196,
  sym_try_statement = 197,
  sym_catch_clause = 198,
  sym_catch_formal_parameter = 199,
  sym_catch_type = 200,
  sym_finally_clause = 201,
  sym_try_with_resources_statement = 202,
  sym_resource_specification = 203,
  sym_resource = 204,
  sym_if_statement = 205,
  sym_while_statement = 206,
  sym_for_statement = 207,
  sym_enhanced_for_statement = 208,
  sym__annotation = 209,
  sym_marker_annotation = 210,
  sym_annotation = 211,
  sym_annotation_argument_list = 212,
  sym_element_value_pair = 213,
  sym__element_value = 214,
  sym_element_value_array_initializer = 215,
  sym_declaration = 216,
  sym_module_declaration = 217,
  sym_module_body = 218,
  sym_module_directive = 219,
  sym_requires_module_directive = 220,
  sym_requires_modifier = 221,
  sym_exports_module_directive = 222,
  sym_opens_module_directive = 223,
  sym_uses_module_directive = 224,
  sym_provides_module_directive = 225,
  sym_package_declaration = 226,
  sym_import_declaration = 227,
  sym_asterisk = 228,
  sym_enum_declaration = 229,
  sym_enum_body = 230,
  sym_enum_body_declarations = 231,
  sym_enum_constant = 232,
  sym_class_declaration = 233,
  sym_modifiers = 234,
  sym_type_parameters = 235,
  sym_type_parameter = 236,
  sym_type_bound = 237,
  sym_superclass = 238,
  sym_super_interfaces = 239,
  sym_type_list = 240,
  sym_permits = 241,
  sym_class_body = 242,
  sym_static_initializer = 243,
  sym_constructor_declaration = 244,
  sym__constructor_declarator = 245,
  sym_constructor_body = 246,
  sym_explicit_constructor_invocation = 247,
  sym_scoped_identifier = 248,
  sym_field_declaration = 249,
  sym_record_declaration = 250,
  sym_annotation_type_declaration = 251,
  sym_annotation_type_body = 252,
  sym_annotation_type_element_declaration = 253,
  sym__default_value = 254,
  sym_interface_declaration = 255,
  sym_extends_interfaces = 256,
  sym_interface_body = 257,
  sym_constant_declaration = 258,
  sym__variable_declarator_list = 259,
  sym_variable_declarator = 260,
  sym__variable_declarator_id = 261,
  sym_array_initializer = 262,
  sym__type = 263,
  sym__unannotated_type = 264,
  sym_annotated_type = 265,
  sym_scoped_type_identifier = 266,
  sym_generic_type = 267,
  sym_array_type = 268,
  sym_integral_type = 269,
  sym_floating_point_type = 270,
  sym__method_header = 271,
  sym__method_declarator = 272,
  sym_formal_parameters = 273,
  sym_formal_parameter = 274,
  sym_receiver_parameter = 275,
  sym_spread_parameter = 276,
  sym_throws = 277,
  sym_local_variable_declaration = 278,
  sym_method_declaration = 279,
  sym_compact_constructor_declaration = 280,
  sym__reserved_identifier = 281,
  aux_sym_program_repeat1 = 282,
  aux_sym__string_literal_repeat1 = 283,
  aux_sym__multiline_string_literal_repeat1 = 284,
  aux_sym_cast_expression_repeat1 = 285,
  aux_sym_inferred_parameters_repeat1 = 286,
  aux_sym_array_creation_expression_repeat1 = 287,
  aux_sym_array_creation_expression_repeat2 = 288,
  aux_sym_argument_list_repeat1 = 289,
  aux_sym_type_arguments_repeat1 = 290,
  aux_sym_dimensions_repeat1 = 291,
  aux_sym_switch_block_repeat1 = 292,
  aux_sym_switch_block_repeat2 = 293,
  aux_sym_switch_block_statement_group_repeat1 = 294,
  aux_sym_switch_block_statement_group_repeat2 = 295,
  aux_sym_record_pattern_body_repeat1 = 296,
  aux_sym_try_statement_repeat1 = 297,
  aux_sym_catch_type_repeat1 = 298,
  aux_sym_resource_specification_repeat1 = 299,
  aux_sym_for_statement_repeat1 = 300,
  aux_sym_for_statement_repeat2 = 301,
  aux_sym_annotation_argument_list_repeat1 = 302,
  aux_sym_element_value_array_initializer_repeat1 = 303,
  aux_sym_module_body_repeat1 = 304,
  aux_sym_requires_module_directive_repeat1 = 305,
  aux_sym_exports_module_directive_repeat1 = 306,
  aux_sym_provides_module_directive_repeat1 = 307,
  aux_sym_enum_body_repeat1 = 308,
  aux_sym_enum_body_declarations_repeat1 = 309,
  aux_sym_modifiers_repeat1 = 310,
  aux_sym_type_parameters_repeat1 = 311,
  aux_sym_type_bound_repeat1 = 312,
  aux_sym_type_list_repeat1 = 313,
  aux_sym_annotation_type_body_repeat1 = 314,
  aux_sym_interface_body_repeat1 = 315,
  aux_sym__variable_declarator_list_repeat1 = 316,
  aux_sym_array_initializer_repeat1 = 317,
  aux_sym_formal_parameters_repeat1 = 318,
  aux_sym_receiver_parameter_repeat1 = 319,
  alias_sym_type_identifier = 320,
};

static const char * const ts_symbol_names[] = {
  [ts_builtin_sym_end] = "end",
  [sym_identifier] = "identifier",
  [sym_decimal_integer_literal] = "decimal_integer_literal",
  [sym_hex_integer_literal] = "hex_integer_literal",
  [sym_octal_integer_literal] = "octal_integer_literal",
  [sym_binary_integer_literal] = "binary_integer_literal",
  [sym_decimal_floating_point_literal] = "decimal_floating_point_literal",
  [sym_hex_floating_point_literal] = "hex_floating_point_literal",
  [sym_true] = "true",
  [sym_false] = "false",
  [sym_character_literal] = "character_literal",
  [anon_sym_DQUOTE] = "\"",
  [anon_sym_DQUOTE_DQUOTE_DQUOTE] = "\"\"\"",
  [sym_string_fragment] = "string_fragment",
  [aux_sym__multiline_string_fragment_token1] = "_multiline_string_fragment_token1",
  [aux_sym__multiline_string_fragment_token2] = "_multiline_string_fragment_token2",
  [anon_sym_BSLASH_LBRACE] = "\\{",
  [anon_sym_RBRACE] = "}",
  [aux_sym__escape_sequence_token1] = "_escape_sequence_token1",
  [sym_escape_sequence] = "escape_sequence",
  [sym_null_literal] = "null_literal",
  [anon_sym_LPAREN] = "(",
  [anon_sym_RPAREN] = ")",
  [anon_sym_AMP] = "&",
  [anon_sym_EQ] = "=",
  [anon_sym_PLUS_EQ] = "+=",
  [anon_sym_DASH_EQ] = "-=",
  [anon_sym_STAR_EQ] = "*=",
  [anon_sym_SLASH_EQ] = "/=",
  [anon_sym_AMP_EQ] = "&=",
  [anon_sym_PIPE_EQ] = "|=",
  [anon_sym_CARET_EQ] = "^=",
  [anon_sym_PERCENT_EQ] = "%=",
  [anon_sym_LT_LT_EQ] = "<<=",
  [anon_sym_GT_GT_EQ] = ">>=",
  [anon_sym_GT_GT_GT_EQ] = ">>>=",
  [anon_sym_GT] = ">",
  [anon_sym_LT] = "<",
  [anon_sym_GT_EQ] = ">=",
  [anon_sym_LT_EQ] = "<=",
  [anon_sym_EQ_EQ] = "==",
  [anon_sym_BANG_EQ] = "!=",
  [anon_sym_AMP_AMP] = "&&",
  [anon_sym_PIPE_PIPE] = "||",
  [anon_sym_PLUS] = "+",
  [anon_sym_DASH] = "-",
  [anon_sym_STAR] = "*",
  [anon_sym_SLASH] = "/",
  [anon_sym_PIPE] = "|",
  [anon_sym_CARET] = "^",
  [anon_sym_PERCENT] = "%",
  [anon_sym_LT_LT] = "<<",
  [anon_sym_GT_GT] = ">>",
  [anon_sym_GT_GT_GT] = ">>>",
  [anon_sym_instanceof] = "instanceof",
  [anon_sym_final] = "final",
  [anon_sym_DASH_GT] = "->",
  [anon_sym_COMMA] = ",",
  [anon_sym_QMARK] = "\?",
  [anon_sym_COLON] = ":",
  [anon_sym_BANG] = "!",
  [anon_sym_TILDE] = "~",
  [anon_sym_PLUS_PLUS] = "++",
  [anon_sym_DASH_DASH] = "--",
  [anon_sym_new] = "new",
  [anon_sym_LBRACK] = "[",
  [anon_sym_RBRACK] = "]",
  [anon_sym_DOT] = ".",
  [anon_sym_class] = "class",
  [anon_sym_COLON_COLON] = "::",
  [anon_sym_extends] = "extends",
  [anon_sym_switch] = "switch",
  [anon_sym_LBRACE] = "{",
  [anon_sym_case] = "case",
  [anon_sym_default] = "default",
  [sym_underscore_pattern] = "underscore_pattern",
  [anon_sym_when] = "when",
  [anon_sym_SEMI] = ";",
  [anon_sym_assert] = "assert",
  [anon_sym_do] = "do",
  [anon_sym_while] = "while",
  [anon_sym_break] = "break",
  [anon_sym_continue] = "continue",
  [anon_sym_return] = "return",
  [anon_sym_yield] = "yield",
  [anon_sym_synchronized] = "synchronized",
  [anon_sym_throw] = "throw",
  [anon_sym_try] = "try",
  [anon_sym_catch] = "catch",
  [anon_sym_finally] = "finally",
  [anon_sym_if] = "if",
  [anon_sym_else] = "else",
  [anon_sym_for] = "for",
  [anon_sym_AT] = "@",
  [anon_sym_open] = "open",
  [anon_sym_module] = "module",
  [anon_sym_requires] = "requires",
  [anon_sym_transitive] = "transitive",
  [anon_sym_static] = "static",
  [anon_sym_exports] = "exports",
  [anon_sym_to] = "to",
  [anon_sym_opens] = "opens",
  [anon_sym_uses] = "uses",
  [anon_sym_provides] = "provides",
  [anon_sym_with] = "with",
  [anon_sym_package] = "package",
  [anon_sym_import] = "import",
  [anon_sym_enum] = "enum",
  [anon_sym_public] = "public",
  [anon_sym_protected] = "protected",
  [anon_sym_private] = "private",
  [anon_sym_abstract] = "abstract",
  [anon_sym_strictfp] = "strictfp",
  [anon_sym_native] = "native",
  [anon_sym_transient] = "transient",
  [anon_sym_volatile] = "volatile",
  [anon_sym_sealed] = "sealed",
  [anon_sym_non_DASHsealed] = "non-sealed",
  [anon_sym_implements] = "implements",
  [anon_sym_permits] = "permits",
  [anon_sym_record] = "record",
  [anon_sym_ATinterface] = "@interface",
  [anon_sym_interface] = "interface",
  [anon_sym_byte] = "byte",
  [anon_sym_short] = "short",
  [anon_sym_int] = "int",
  [anon_sym_long] = "long",
  [anon_sym_char] = "char",
  [anon_sym_float] = "float",
  [anon_sym_double] = "double",
  [sym_boolean_type] = "boolean_type",
  [sym_void_type] = "void_type",
  [anon_sym_DOT_DOT_DOT] = "...",
  [anon_sym_throws] = "throws",
  [sym_this] = "this",
  [sym_super] = "super",
  [sym_line_comment] = "line_comment",
  [sym_block_comment] = "block_comment",
  [sym_program] = "program",
  [sym__toplevel_statement] = "_toplevel_statement",
  [sym__literal] = "_literal",
  [sym_string_literal] = "string_literal",
  [sym__string_literal] = "_string_literal",
  [sym__multiline_string_literal] = "_multiline_string_literal",
  [sym__multiline_string_fragment] = "multiline_string_fragment",
  [sym_string_interpolation] = "string_interpolation",
  [sym__escape_sequence] = "_escape_sequence",
  [sym_expression] = "expression",
  [sym_cast_expression] = "cast_expression",
  [sym_assignment_expression] = "assignment_expression",
  [sym_binary_expression] = "binary_expression",
  [sym_instanceof_expression] = "instanceof_expression",
  [sym_lambda_expression] = "lambda_expression",
  [sym_inferred_parameters] = "inferred_parameters",
  [sym_ternary_expression] = "ternary_expression",
  [sym_unary_expression] = "unary_expression",
  [sym_update_expression] = "update_expression",
  [sym_primary_expression] = "primary_expression",
  [sym_array_creation_expression] = "array_creation_expression",
  [sym_dimensions_expr] = "dimensions_expr",
  [sym_parenthesized_expression] = "parenthesized_expression",
  [sym_class_literal] = "class_literal",
  [sym_object_creation_expression] = "object_creation_expression",
  [sym__unqualified_object_creation_expression] = "_unqualified_object_creation_expression",
  [sym_field_access] = "field_access",
  [sym_template_expression] = "template_expression",
  [sym_array_access] = "array_access",
  [sym_method_invocation] = "method_invocation",
  [sym_argument_list] = "argument_list",
  [sym_method_reference] = "method_reference",
  [sym_type_arguments] = "type_arguments",
  [sym_wildcard] = "wildcard",
  [sym__wildcard_bounds] = "_wildcard_bounds",
  [sym_dimensions] = "dimensions",
  [sym_switch_expression] = "switch_expression",
  [sym_switch_block] = "switch_block",
  [sym_switch_block_statement_group] = "switch_block_statement_group",
  [sym_switch_rule] = "switch_rule",
  [sym_switch_label] = "switch_label",
  [sym_pattern] = "pattern",
  [sym_type_pattern] = "type_pattern",
  [sym_record_pattern] = "record_pattern",
  [sym_record_pattern_body] = "record_pattern_body",
  [sym_record_pattern_component] = "record_pattern_component",
  [sym_guard] = "guard",
  [sym_statement] = "statement",
  [sym_block] = "block",
  [sym_expression_statement] = "expression_statement",
  [sym_labeled_statement] = "labeled_statement",
  [sym_assert_statement] = "assert_statement",
  [sym_do_statement] = "do_statement",
  [sym_break_statement] = "break_statement",
  [sym_continue_statement] = "continue_statement",
  [sym_return_statement] = "return_statement",
  [sym_yield_statement] = "yield_statement",
  [sym_synchronized_statement] = "synchronized_statement",
  [sym_throw_statement] = "throw_statement",
  [sym_try_statement] = "try_statement",
  [sym_catch_clause] = "catch_clause",
  [sym_catch_formal_parameter] = "catch_formal_parameter",
  [sym_catch_type] = "catch_type",
  [sym_finally_clause] = "finally_clause",
  [sym_try_with_resources_statement] = "try_with_resources_statement",
  [sym_resource_specification] = "resource_specification",
  [sym_resource] = "resource",
  [sym_if_statement] = "if_statement",
  [sym_while_statement] = "while_statement",
  [sym_for_statement] = "for_statement",
  [sym_enhanced_for_statement] = "enhanced_for_statement",
  [sym__annotation] = "_annotation",
  [sym_marker_annotation] = "marker_annotation",
  [sym_annotation] = "annotation",
  [sym_annotation_argument_list] = "annotation_argument_list",
  [sym_element_value_pair] = "element_value_pair",
  [sym__element_value] = "_element_value",
  [sym_element_value_array_initializer] = "element_value_array_initializer",
  [sym_declaration] = "declaration",
  [sym_module_declaration] = "module_declaration",
  [sym_module_body] = "module_body",
  [sym_module_directive] = "module_directive",
  [sym_requires_module_directive] = "requires_module_directive",
  [sym_requires_modifier] = "requires_modifier",
  [sym_exports_module_directive] = "exports_module_directive",
  [sym_opens_module_directive] = "opens_module_directive",
  [sym_uses_module_directive] = "uses_module_directive",
  [sym_provides_module_directive] = "provides_module_directive",
  [sym_package_declaration] = "package_declaration",
  [sym_import_declaration] = "import_declaration",
  [sym_asterisk] = "asterisk",
  [sym_enum_declaration] = "enum_declaration",
  [sym_enum_body] = "enum_body",
  [sym_enum_body_declarations] = "enum_body_declarations",
  [sym_enum_constant] = "enum_constant",
  [sym_class_declaration] = "class_declaration",
  [sym_modifiers] = "modifiers",
  [sym_type_parameters] = "type_parameters",
  [sym_type_parameter] = "type_parameter",
  [sym_type_bound] = "type_bound",
  [sym_superclass] = "superclass",
  [sym_super_interfaces] = "super_interfaces",
  [sym_type_list] = "type_list",
  [sym_permits] = "permits",
  [sym_class_body] = "class_body",
  [sym_static_initializer] = "static_initializer",
  [sym_constructor_declaration] = "constructor_declaration",
  [sym__constructor_declarator] = "_constructor_declarator",
  [sym_constructor_body] = "constructor_body",
  [sym_explicit_constructor_invocation] = "explicit_constructor_invocation",
  [sym_scoped_identifier] = "scoped_identifier",
  [sym_field_declaration] = "field_declaration",
  [sym_record_declaration] = "record_declaration",
  [sym_annotation_type_declaration] = "annotation_type_declaration",
  [sym_annotation_type_body] = "annotation_type_body",
  [sym_annotation_type_element_declaration] = "annotation_type_element_declaration",
  [sym__default_value] = "_default_value",
  [sym_interface_declaration] = "interface_declaration",
  [sym_extends_interfaces] = "extends_interfaces",
  [sym_interface_body] = "interface_body",
  [sym_constant_declaration] = "constant_declaration",
  [sym__variable_declarator_list] = "_variable_declarator_list",
  [sym_variable_declarator] = "variable_declarator",
  [sym__variable_declarator_id] = "_variable_declarator_id",
  [sym_array_initializer] = "array_initializer",
  [sym__type] = "_type",
  [sym__unannotated_type] = "_unannotated_type",
  [sym_annotated_type] = "annotated_type",
  [sym_scoped_type_identifier] = "scoped_type_identifier",
  [sym_generic_type] = "generic_type",
  [sym_array_type] = "array_type",
  [sym_integral_type] = "integral_type",
  [sym_floating_point_type] = "floating_point_type",
  [sym__method_header] = "_method_header",
  [sym__method_declarator] = "_method_declarator",
  [sym_formal_parameters] = "formal_parameters",
  [sym_formal_parameter] = "formal_parameter",
  [sym_receiver_parameter] = "receiver_parameter",
  [sym_spread_parameter] = "spread_parameter",
  [sym_throws] = "throws",
  [sym_local_variable_declaration] = "local_variable_declaration",
  [sym_method_declaration] = "method_declaration",
  [sym_compact_constructor_declaration] = "compact_constructor_declaration",
  [sym__reserved_identifier] = "_reserved_identifier",
  [aux_sym_program_repeat1] = "program_repeat1",
  [aux_sym__string_literal_repeat1] = "_string_literal_repeat1",
  [aux_sym__multiline_string_literal_repeat1] = "_multiline_string_literal_repeat1",
  [aux_sym_cast_expression_repeat1] = "cast_expression_repeat1",
  [aux_sym_inferred_parameters_repeat1] = "inferred_parameters_repeat1",
  [aux_sym_array_creation_expression_repeat1] = "array_creation_expression_repeat1",
  [aux_sym_array_creation_expression_repeat2] = "array_creation_expression_repeat2",
  [aux_sym_argument_list_repeat1] = "argument_list_repeat1",
  [aux_sym_type_arguments_repeat1] = "type_arguments_repeat1",
  [aux_sym_dimensions_repeat1] = "dimensions_repeat1",
  [aux_sym_switch_block_repeat1] = "switch_block_repeat1",
  [aux_sym_switch_block_repeat2] = "switch_block_repeat2",
  [aux_sym_switch_block_statement_group_repeat1] = "switch_block_statement_group_repeat1",
  [aux_sym_switch_block_statement_group_repeat2] = "switch_block_statement_group_repeat2",
  [aux_sym_record_pattern_body_repeat1] = "record_pattern_body_repeat1",
  [aux_sym_try_statement_repeat1] = "try_statement_repeat1",
  [aux_sym_catch_type_repeat1] = "catch_type_repeat1",
  [aux_sym_resource_specification_repeat1] = "resource_specification_repeat1",
  [aux_sym_for_statement_repeat1] = "for_statement_repeat1",
  [aux_sym_for_statement_repeat2] = "for_statement_repeat2",
  [aux_sym_annotation_argument_list_repeat1] = "annotation_argument_list_repeat1",
  [aux_sym_element_value_array_initializer_repeat1] = "element_value_array_initializer_repeat1",
  [aux_sym_module_body_repeat1] = "module_body_repeat1",
  [aux_sym_requires_module_directive_repeat1] = "requires_module_directive_repeat1",
  [aux_sym_exports_module_directive_repeat1] = "exports_module_directive_repeat1",
  [aux_sym_provides_module_directive_repeat1] = "provides_module_directive_repeat1",
  [aux_sym_enum_body_repeat1] = "enum_body_repeat1",
  [aux_sym_enum_body_declarations_repeat1] = "enum_body_declarations_repeat1",
  [aux_sym_modifiers_repeat1] = "modifiers_repeat1",
  [aux_sym_type_parameters_repeat1] = "type_parameters_repeat1",
  [aux_sym_type_bound_repeat1] = "type_bound_repeat1",
  [aux_sym_type_list_repeat1] = "type_list_repeat1",
  [aux_sym_annotation_type_body_repeat1] = "annotation_type_body_repeat1",
  [aux_sym_interface_body_repeat1] = "interface_body_repeat1",
  [aux_sym__variable_declarator_list_repeat1] = "_variable_declarator_list_repeat1",
  [aux_sym_array_initializer_repeat1] = "array_initializer_repeat1",
  [aux_sym_formal_parameters_repeat1] = "formal_parameters_repeat1",
  [aux_sym_receiver_parameter_repeat1] = "receiver_parameter_repeat1",
  [alias_sym_type_identifier] = "type_identifier",
};

static const TSSymbol ts_symbol_map[] = {
  [ts_builtin_sym_end] = ts_builtin_sym_end,
  [sym_identifier] = sym_identifier,
  [sym_decimal_integer_literal] = sym_decimal_integer_literal,
  [sym_hex_integer_literal] = sym_hex_integer_literal,
  [sym_octal_integer_literal] = sym_octal_integer_literal,
  [sym_binary_integer_literal] = sym_binary_integer_literal,
  [sym_decimal_floating_point_literal] = sym_decimal_floating_point_literal,
  [sym_hex_floating_point_literal] = sym_hex_floating_point_literal,
  [sym_true] = sym_true,
  [sym_false] = sym_false,
  [sym_character_literal] = sym_character_literal,
  [anon_sym_DQUOTE] = anon_sym_DQUOTE,
  [anon_sym_DQUOTE_DQUOTE_DQUOTE] = anon_sym_DQUOTE_DQUOTE_DQUOTE,
  [sym_string_fragment] = sym_string_fragment,
  [aux_sym__multiline_string_fragment_token1] = aux_sym__multiline_string_fragment_token1,
  [aux_sym__multiline_string_fragment_token2] = aux_sym__multiline_string_fragment_token2,
  [anon_sym_BSLASH_LBRACE] = anon_sym_BSLASH_LBRACE,
  [anon_sym_RBRACE] = anon_sym_RBRACE,
  [aux_sym__escape_sequence_token1] = aux_sym__escape_sequence_token1,
  [sym_escape_sequence] = sym_escape_sequence,
  [sym_null_literal] = sym_null_literal,
  [anon_sym_LPAREN] = anon_sym_LPAREN,
  [anon_sym_RPAREN] = anon_sym_RPAREN,
  [anon_sym_AMP] = anon_sym_AMP,
  [anon_sym_EQ] = anon_sym_EQ,
  [anon_sym_PLUS_EQ] = anon_sym_PLUS_EQ,
  [anon_sym_DASH_EQ] = anon_sym_DASH_EQ,
  [anon_sym_STAR_EQ] = anon_sym_STAR_EQ,
  [anon_sym_SLASH_EQ] = anon_sym_SLASH_EQ,
  [anon_sym_AMP_EQ] = anon_sym_AMP_EQ,
  [anon_sym_PIPE_EQ] = anon_sym_PIPE_EQ,
  [anon_sym_CARET_EQ] = anon_sym_CARET_EQ,
  [anon_sym_PERCENT_EQ] = anon_sym_PERCENT_EQ,
  [anon_sym_LT_LT_EQ] = anon_sym_LT_LT_EQ,
  [anon_sym_GT_GT_EQ] = anon_sym_GT_GT_EQ,
  [anon_sym_GT_GT_GT_EQ] = anon_sym_GT_GT_GT_EQ,
  [anon_sym_GT] = anon_sym_GT,
  [anon_sym_LT] = anon_sym_LT,
  [anon_sym_GT_EQ] = anon_sym_GT_EQ,
  [anon_sym_LT_EQ] = anon_sym_LT_EQ,
  [anon_sym_EQ_EQ] = anon_sym_EQ_EQ,
  [anon_sym_BANG_EQ] = anon_sym_BANG_EQ,
  [anon_sym_AMP_AMP] = anon_sym_AMP_AMP,
  [anon_sym_PIPE_PIPE] = anon_sym_PIPE_PIPE,
  [anon_sym_PLUS] = anon_sym_PLUS,
  [anon_sym_DASH] = anon_sym_DASH,
  [anon_sym_STAR] = anon_sym_STAR,
  [anon_sym_SLASH] = anon_sym_SLASH,
  [anon_sym_PIPE] = anon_sym_PIPE,
  [anon_sym_CARET] = anon_sym_CARET,
  [anon_sym_PERCENT] = anon_sym_PERCENT,
  [anon_sym_LT_LT] = anon_sym_LT_LT,
  [anon_sym_GT_GT] = anon_sym_GT_GT,
  [anon_sym_GT_GT_GT] = anon_sym_GT_GT_GT,
  [anon_sym_instanceof] = anon_sym_instanceof,
  [anon_sym_final] = anon_sym_final,
  [anon_sym_DASH_GT] = anon_sym_DASH_GT,
  [anon_sym_COMMA] = anon_sym_COMMA,
  [anon_sym_QMARK] = anon_sym_QMARK,
  [anon_sym_COLON] = anon_sym_COLON,
  [anon_sym_BANG] = anon_sym_BANG,
  [anon_sym_TILDE] = anon_sym_TILDE,
  [anon_sym_PLUS_PLUS] = anon_sym_PLUS_PLUS,
  [anon_sym_DASH_DASH] = anon_sym_DASH_DASH,
  [anon_sym_new] = anon_sym_new,
  [anon_sym_LBRACK] = anon_sym_LBRACK,
  [anon_sym_RBRACK] = anon_sym_RBRACK,
  [anon_sym_DOT] = anon_sym_DOT,
  [anon_sym_class] = anon_sym_class,
  [anon_sym_COLON_COLON] = anon_sym_COLON_COLON,
  [anon_sym_extends] = anon_sym_extends,
  [anon_sym_switch] = anon_sym_switch,
  [anon_sym_LBRACE] = anon_sym_LBRACE,
  [anon_sym_case] = anon_sym_case,
  [anon_sym_default] = anon_sym_default,
  [sym_underscore_pattern] = sym_underscore_pattern,
  [anon_sym_when] = anon_sym_when,
  [anon_sym_SEMI] = anon_sym_SEMI,
  [anon_sym_assert] = anon_sym_assert,
  [anon_sym_do] = anon_sym_do,
  [anon_sym_while] = anon_sym_while,
  [anon_sym_break] = anon_sym_break,
  [anon_sym_continue] = anon_sym_continue,
  [anon_sym_return] = anon_sym_return,
  [anon_sym_yield] = anon_sym_yield,
  [anon_sym_synchronized] = anon_sym_synchronized,
  [anon_sym_throw] = anon_sym_throw,
  [anon_sym_try] = anon_sym_try,
  [anon_sym_catch] = anon_sym_catch,
  [anon_sym_finally] = anon_sym_finally,
  [anon_sym_if] = anon_sym_if,
  [anon_sym_else] = anon_sym_else,
  [anon_sym_for] = anon_sym_for,
  [anon_sym_AT] = anon_sym_AT,
  [anon_sym_open] = anon_sym_open,
  [anon_sym_module] = anon_sym_module,
  [anon_sym_requires] = anon_sym_requires,
  [anon_sym_transitive] = anon_sym_transitive,
  [anon_sym_static] = anon_sym_static,
  [anon_sym_exports] = anon_sym_exports,
  [anon_sym_to] = anon_sym_to,
  [anon_sym_opens] = anon_sym_opens,
  [anon_sym_uses] = anon_sym_uses,
  [anon_sym_provides] = anon_sym_provides,
  [anon_sym_with] = anon_sym_with,
  [anon_sym_package] = anon_sym_package,
  [anon_sym_import] = anon_sym_import,
  [anon_sym_enum] = anon_sym_enum,
  [anon_sym_public] = anon_sym_public,
  [anon_sym_protected] = anon_sym_protected,
  [anon_sym_private] = anon_sym_private,
  [anon_sym_abstract] = anon_sym_abstract,
  [anon_sym_strictfp] = anon_sym_strictfp,
  [anon_sym_native] = anon_sym_native,
  [anon_sym_transient] = anon_sym_transient,
  [anon_sym_volatile] = anon_sym_volatile,
  [anon_sym_sealed] = anon_sym_sealed,
  [anon_sym_non_DASHsealed] = anon_sym_non_DASHsealed,
  [anon_sym_implements] = anon_sym_implements,
  [anon_sym_permits] = anon_sym_permits,
  [anon_sym_record] = anon_sym_record,
  [anon_sym_ATinterface] = anon_sym_ATinterface,
  [anon_sym_interface] = anon_sym_interface,
  [anon_sym_byte] = anon_sym_byte,
  [anon_sym_short] = anon_sym_short,
  [anon_sym_int] = anon_sym_int,
  [anon_sym_long] = anon_sym_long,
  [anon_sym_char] = anon_sym_char,
  [anon_sym_float] = anon_sym_float,
  [anon_sym_double] = anon_sym_double,
  [sym_boolean_type] = sym_boolean_type,
  [sym_void_type] = sym_void_type,
  [anon_sym_DOT_DOT_DOT] = anon_sym_DOT_DOT_DOT,
  [anon_sym_throws] = anon_sym_throws,
  [sym_this] = sym_this,
  [sym_super] = sym_super,
  [sym_line_comment] = sym_line_comment,
  [sym_block_comment] = sym_block_comment,
  [sym_program] = sym_program,
  [sym__toplevel_statement] = sym__toplevel_statement,
  [sym__literal] = sym__literal,
  [sym_string_literal] = sym_string_literal,
  [sym__string_literal] = sym__string_literal,
  [sym__multiline_string_literal] = sym__multiline_string_literal,
  [sym__multiline_string_fragment] = sym__multiline_string_fragment,
  [sym_string_interpolation] = sym_string_interpolation,
  [sym__escape_sequence] = sym__escape_sequence,
  [sym_expression] = sym_expression,
  [sym_cast_expression] = sym_cast_expression,
  [sym_assignment_expression] = sym_assignment_expression,
  [sym_binary_expression] = sym_binary_expression,
  [sym_instanceof_expression] = sym_instanceof_expression,
  [sym_lambda_expression] = sym_lambda_expression,
  [sym_inferred_parameters] = sym_inferred_parameters,
  [sym_ternary_expression] = sym_ternary_expression,
  [sym_unary_expression] = sym_unary_expression,
  [sym_update_expression] = sym_update_expression,
  [sym_primary_expression] = sym_primary_expression,
  [sym_array_creation_expression] = sym_array_creation_expression,
  [sym_dimensions_expr] = sym_dimensions_expr,
  [sym_parenthesized_expression] = sym_parenthesized_expression,
  [sym_class_literal] = sym_class_literal,
  [sym_object_creation_expression] = sym_object_creation_expression,
  [sym__unqualified_object_creation_expression] = sym__unqualified_object_creation_expression,
  [sym_field_access] = sym_field_access,
  [sym_template_expression] = sym_template_expression,
  [sym_array_access] = sym_array_access,
  [sym_method_invocation] = sym_method_invocation,
  [sym_argument_list] = sym_argument_list,
  [sym_method_reference] = sym_method_reference,
  [sym_type_arguments] = sym_type_arguments,
  [sym_wildcard] = sym_wildcard,
  [sym__wildcard_bounds] = sym__wildcard_bounds,
  [sym_dimensions] = sym_dimensions,
  [sym_switch_expression] = sym_switch_expression,
  [sym_switch_block] = sym_switch_block,
  [sym_switch_block_statement_group] = sym_switch_block_statement_group,
  [sym_switch_rule] = sym_switch_rule,
  [sym_switch_label] = sym_switch_label,
  [sym_pattern] = sym_pattern,
  [sym_type_pattern] = sym_type_pattern,
  [sym_record_pattern] = sym_record_pattern,
  [sym_record_pattern_body] = sym_record_pattern_body,
  [sym_record_pattern_component] = sym_record_pattern_component,
  [sym_guard] = sym_guard,
  [sym_statement] = sym_statement,
  [sym_block] = sym_block,
  [sym_expression_statement] = sym_expression_statement,
  [sym_labeled_statement] = sym_labeled_statement,
  [sym_assert_statement] = sym_assert_statement,
  [sym_do_statement] = sym_do_statement,
  [sym_break_statement] = sym_break_statement,
  [sym_continue_statement] = sym_continue_statement,
  [sym_return_statement] = sym_return_statement,
  [sym_yield_statement] = sym_yield_statement,
  [sym_synchronized_statement] = sym_synchronized_statement,
  [sym_throw_statement] = sym_throw_statement,
  [sym_try_statement] = sym_try_statement,
  [sym_catch_clause] = sym_catch_clause,
  [sym_catch_formal_parameter] = sym_catch_formal_parameter,
  [sym_catch_type] = sym_catch_type,
  [sym_finally_clause] = sym_finally_clause,
  [sym_try_with_resources_statement] = sym_try_with_resources_statement,
  [sym_resource_specification] = sym_resource_specification,
  [sym_resource] = sym_resource,
  [sym_if_statement] = sym_if_statement,
  [sym_while_statement] = sym_while_statement,
  [sym_for_statement] = sym_for_statement,
  [sym_enhanced_for_statement] = sym_enhanced_for_statement,
  [sym__annotation] = sym__annotation,
  [sym_marker_annotation] = sym_marker_annotation,
  [sym_annotation] = sym_annotation,
  [sym_annotation_argument_list] = sym_annotation_argument_list,
  [sym_element_value_pair] = sym_element_value_pair,
  [sym__element_value] = sym__element_value,
  [sym_element_value_array_initializer] = sym_element_value_array_initializer,
  [sym_declaration] = sym_declaration,
  [sym_module_declaration] = sym_module_declaration,
  [sym_module_body] = sym_module_body,
  [sym_module_directive] = sym_module_directive,
  [sym_requires_module_directive] = sym_requires_module_directive,
  [sym_requires_modifier] = sym_requires_modifier,
  [sym_exports_module_directive] = sym_exports_module_directive,
  [sym_opens_module_directive] = sym_opens_module_directive,
  [sym_uses_module_directive] = sym_uses_module_directive,
  [sym_provides_module_directive] = sym_provides_module_directive,
  [sym_package_declaration] = sym_package_declaration,
  [sym_import_declaration] = sym_import_declaration,
  [sym_asterisk] = sym_asterisk,
  [sym_enum_declaration] = sym_enum_declaration,
  [sym_enum_body] = sym_enum_body,
  [sym_enum_body_declarations] = sym_enum_body_declarations,
  [sym_enum_constant] = sym_enum_constant,
  [sym_class_declaration] = sym_class_declaration,
  [sym_modifiers] = sym_modifiers,
  [sym_type_parameters] = sym_type_parameters,
  [sym_type_parameter] = sym_type_parameter,
  [sym_type_bound] = sym_type_bound,
  [sym_superclass] = sym_superclass,
  [sym_super_interfaces] = sym_super_interfaces,
  [sym_type_list] = sym_type_list,
  [sym_permits] = sym_permits,
  [sym_class_body] = sym_class_body,
  [sym_static_initializer] = sym_static_initializer,
  [sym_constructor_declaration] = sym_constructor_declaration,
  [sym__constructor_declarator] = sym__constructor_declarator,
  [sym_constructor_body] = sym_constructor_body,
  [sym_explicit_constructor_invocation] = sym_explicit_constructor_invocation,
  [sym_scoped_identifier] = sym_scoped_identifier,
  [sym_field_declaration] = sym_field_declaration,
  [sym_record_declaration] = sym_record_declaration,
  [sym_annotation_type_declaration] = sym_annotation_type_declaration,
  [sym_annotation_type_body] = sym_annotation_type_body,
  [sym_annotation_type_element_declaration] = sym_annotation_type_element_declaration,
  [sym__default_value] = sym__default_value,
  [sym_interface_declaration] = sym_interface_declaration,
  [sym_extends_interfaces] = sym_extends_interfaces,
  [sym_interface_body] = sym_interface_body,
  [sym_constant_declaration] = sym_constant_declaration,
  [sym__variable_declarator_list] = sym__variable_declarator_list,
  [sym_variable_declarator] = sym_variable_declarator,
  [sym__variable_declarator_id] = sym__variable_declarator_id,
  [sym_array_initializer] = sym_array_initializer,
  [sym__type] = sym__type,
  [sym__unannotated_type] = sym__unannotated_type,
  [sym_annotated_type] = sym_annotated_type,
  [sym_scoped_type_identifier] = sym_scoped_type_identifier,
  [sym_generic_type] = sym_generic_type,
  [sym_array_type] = sym_array_type,
  [sym_integral_type] = sym_integral_type,
  [sym_floating_point_type] = sym_floating_point_type,
  [sym__method_header] = sym__method_header,
  [sym__method_declarator] = sym__method_declarator,
  [sym_formal_parameters] = sym_formal_parameters,
  [sym_formal_parameter] = sym_formal_parameter,
  [sym_receiver_parameter] = sym_receiver_parameter,
  [sym_spread_parameter] = sym_spread_parameter,
  [sym_throws] = sym_throws,
  [sym_local_variable_declaration] = sym_local_variable_declaration,
  [sym_method_declaration] = sym_method_declaration,
  [sym_compact_constructor_declaration] = sym_compact_constructor_declaration,
  [sym__reserved_identifier] = sym__reserved_identifier,
  [aux_sym_program_repeat1] = aux_sym_program_repeat1,
  [aux_sym__string_literal_repeat1] = aux_sym__string_literal_repeat1,
  [aux_sym__multiline_string_literal_repeat1] = aux_sym__multiline_string_literal_repeat1,
  [aux_sym_cast_expression_repeat1] = aux_sym_cast_expression_repeat1,
  [aux_sym_inferred_parameters_repeat1] = aux_sym_inferred_parameters_repeat1,
  [aux_sym_array_creation_expression_repeat1] = aux_sym_array_creation_expression_repeat1,
  [aux_sym_array_creation_expression_repeat2] = aux_sym_array_creation_expression_repeat2,
  [aux_sym_argument_list_repeat1] = aux_sym_argument_list_repeat1,
  [aux_sym_type_arguments_repeat1] = aux_sym_type_arguments_repeat1,
  [aux_sym_dimensions_repeat1] = aux_sym_dimensions_repeat1,
  [aux_sym_switch_block_repeat1] = aux_sym_switch_block_repeat1,
  [aux_sym_switch_block_repeat2] = aux_sym_switch_block_repeat2,
  [aux_sym_switch_block_statement_group_repeat1] = aux_sym_switch_block_statement_group_repeat1,
  [aux_sym_switch_block_statement_group_repeat2] = aux_sym_switch_block_statement_group_repeat2,
  [aux_sym_record_pattern_body_repeat1] = aux_sym_record_pattern_body_repeat1,
  [aux_sym_try_statement_repeat1] = aux_sym_try_statement_repeat1,
  [aux_sym_catch_type_repeat1] = aux_sym_catch_type_repeat1,
  [aux_sym_resource_specification_repeat1] = aux_sym_resource_specification_repeat1,
  [aux_sym_for_statement_repeat1] = aux_sym_for_statement_repeat1,
  [aux_sym_for_statement_repeat2] = aux_sym_for_statement_repeat2,
  [aux_sym_annotation_argument_list_repeat1] = aux_sym_annotation_argument_list_repeat1,
  [aux_sym_element_value_array_initializer_repeat1] = aux_sym_element_value_array_initializer_repeat1,
  [aux_sym_module_body_repeat1] = aux_sym_module_body_repeat1,
  [aux_sym_requires_module_directive_repeat1] = aux_sym_requires_module_directive_repeat1,
  [aux_sym_exports_module_directive_repeat1] = aux_sym_exports_module_directive_repeat1,
  [aux_sym_provides_module_directive_repeat1] = aux_sym_provides_module_directive_repeat1,
  [aux_sym_enum_body_repeat1] = aux_sym_enum_body_repeat1,
  [aux_sym_enum_body_declarations_repeat1] = aux_sym_enum_body_declarations_repeat1,
  [aux_sym_modifiers_repeat1] = aux_sym_modifiers_repeat1,
  [aux_sym_type_parameters_repeat1] = aux_sym_type_parameters_repeat1,
  [aux_sym_type_bound_repeat1] = aux_sym_type_bound_repeat1,
  [aux_sym_type_list_repeat1] = aux_sym_type_list_repeat1,
  [aux_sym_annotation_type_body_repeat1] = aux_sym_annotation_type_body_repeat1,
  [aux_sym_interface_body_repeat1] = aux_sym_interface_body_repeat1,
  [aux_sym__variable_declarator_list_repeat1] = aux_sym__variable_declarator_list_repeat1,
  [aux_sym_array_initializer_repeat1] = aux_sym_array_initializer_repeat1,
  [aux_sym_formal_parameters_repeat1] = aux_sym_formal_parameters_repeat1,
  [aux_sym_receiver_parameter_repeat1] = aux_sym_receiver_parameter_repeat1,
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
  [sym_decimal_integer_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_hex_integer_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_octal_integer_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_binary_integer_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_decimal_floating_point_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_hex_floating_point_literal] = {
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
  [sym_character_literal] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_DQUOTE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DQUOTE_DQUOTE_DQUOTE] = {
    .visible = true,
    .named = false,
  },
  [sym_string_fragment] = {
    .visible = true,
    .named = true,
  },
  [aux_sym__multiline_string_fragment_token1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__multiline_string_fragment_token2] = {
    .visible = false,
    .named = false,
  },
  [anon_sym_BSLASH_LBRACE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_RBRACE] = {
    .visible = true,
    .named = false,
  },
  [aux_sym__escape_sequence_token1] = {
    .visible = false,
    .named = false,
  },
  [sym_escape_sequence] = {
    .visible = true,
    .named = true,
  },
  [sym_null_literal] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_LPAREN] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_RPAREN] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_AMP] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_EQ] = {
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
  [anon_sym_AMP_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PIPE_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_CARET_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PERCENT_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LT_LT_EQ] = {
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
  [anon_sym_GT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_GT_EQ] = {
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
  [anon_sym_BANG_EQ] = {
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
  [anon_sym_PLUS] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DASH] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_STAR] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_SLASH] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PIPE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_CARET] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PERCENT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LT_LT] = {
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
  [anon_sym_instanceof] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_final] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DASH_GT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_COMMA] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_QMARK] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_COLON] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_BANG] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_TILDE] = {
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
  [anon_sym_new] = {
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
  [anon_sym_DOT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_class] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_COLON_COLON] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_extends] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_switch] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LBRACE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_case] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_default] = {
    .visible = true,
    .named = false,
  },
  [sym_underscore_pattern] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_when] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_SEMI] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_assert] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_do] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_while] = {
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
  [anon_sym_return] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_yield] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_synchronized] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_throw] = {
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
  [anon_sym_if] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_else] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_for] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_AT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_open] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_module] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_requires] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_transitive] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_static] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_exports] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_to] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_opens] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_uses] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_provides] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_with] = {
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
  [anon_sym_enum] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_public] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_protected] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_private] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_abstract] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_strictfp] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_native] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_transient] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_volatile] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_sealed] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_non_DASHsealed] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_implements] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_permits] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_record] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_ATinterface] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_interface] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_byte] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_short] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_int] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_long] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_char] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_float] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_double] = {
    .visible = true,
    .named = false,
  },
  [sym_boolean_type] = {
    .visible = true,
    .named = true,
  },
  [sym_void_type] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_DOT_DOT_DOT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_throws] = {
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
  [sym_line_comment] = {
    .visible = true,
    .named = true,
  },
  [sym_block_comment] = {
    .visible = true,
    .named = true,
  },
  [sym_program] = {
    .visible = true,
    .named = true,
  },
  [sym__toplevel_statement] = {
    .visible = false,
    .named = true,
  },
  [sym__literal] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym_string_literal] = {
    .visible = true,
    .named = true,
  },
  [sym__string_literal] = {
    .visible = false,
    .named = true,
  },
  [sym__multiline_string_literal] = {
    .visible = false,
    .named = true,
  },
  [sym__multiline_string_fragment] = {
    .visible = true,
    .named = true,
  },
  [sym_string_interpolation] = {
    .visible = true,
    .named = true,
  },
  [sym__escape_sequence] = {
    .visible = false,
    .named = true,
  },
  [sym_expression] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym_cast_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_assignment_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_binary_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_instanceof_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_lambda_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_inferred_parameters] = {
    .visible = true,
    .named = true,
  },
  [sym_ternary_expression] = {
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
  [sym_primary_expression] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym_array_creation_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_dimensions_expr] = {
    .visible = true,
    .named = true,
  },
  [sym_parenthesized_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_class_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_object_creation_expression] = {
    .visible = true,
    .named = true,
  },
  [sym__unqualified_object_creation_expression] = {
    .visible = false,
    .named = true,
  },
  [sym_field_access] = {
    .visible = true,
    .named = true,
  },
  [sym_template_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_array_access] = {
    .visible = true,
    .named = true,
  },
  [sym_method_invocation] = {
    .visible = true,
    .named = true,
  },
  [sym_argument_list] = {
    .visible = true,
    .named = true,
  },
  [sym_method_reference] = {
    .visible = true,
    .named = true,
  },
  [sym_type_arguments] = {
    .visible = true,
    .named = true,
  },
  [sym_wildcard] = {
    .visible = true,
    .named = true,
  },
  [sym__wildcard_bounds] = {
    .visible = false,
    .named = true,
  },
  [sym_dimensions] = {
    .visible = true,
    .named = true,
  },
  [sym_switch_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_switch_block] = {
    .visible = true,
    .named = true,
  },
  [sym_switch_block_statement_group] = {
    .visible = true,
    .named = true,
  },
  [sym_switch_rule] = {
    .visible = true,
    .named = true,
  },
  [sym_switch_label] = {
    .visible = true,
    .named = true,
  },
  [sym_pattern] = {
    .visible = true,
    .named = true,
  },
  [sym_type_pattern] = {
    .visible = true,
    .named = true,
  },
  [sym_record_pattern] = {
    .visible = true,
    .named = true,
  },
  [sym_record_pattern_body] = {
    .visible = true,
    .named = true,
  },
  [sym_record_pattern_component] = {
    .visible = true,
    .named = true,
  },
  [sym_guard] = {
    .visible = true,
    .named = true,
  },
  [sym_statement] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym_block] = {
    .visible = true,
    .named = true,
  },
  [sym_expression_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_labeled_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_assert_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_do_statement] = {
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
  [sym_return_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_yield_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_synchronized_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_throw_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_try_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_catch_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_catch_formal_parameter] = {
    .visible = true,
    .named = true,
  },
  [sym_catch_type] = {
    .visible = true,
    .named = true,
  },
  [sym_finally_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_try_with_resources_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_resource_specification] = {
    .visible = true,
    .named = true,
  },
  [sym_resource] = {
    .visible = true,
    .named = true,
  },
  [sym_if_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_while_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_for_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_enhanced_for_statement] = {
    .visible = true,
    .named = true,
  },
  [sym__annotation] = {
    .visible = false,
    .named = true,
  },
  [sym_marker_annotation] = {
    .visible = true,
    .named = true,
  },
  [sym_annotation] = {
    .visible = true,
    .named = true,
  },
  [sym_annotation_argument_list] = {
    .visible = true,
    .named = true,
  },
  [sym_element_value_pair] = {
    .visible = true,
    .named = true,
  },
  [sym__element_value] = {
    .visible = false,
    .named = true,
  },
  [sym_element_value_array_initializer] = {
    .visible = true,
    .named = true,
  },
  [sym_declaration] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym_module_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_module_body] = {
    .visible = true,
    .named = true,
  },
  [sym_module_directive] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym_requires_module_directive] = {
    .visible = true,
    .named = true,
  },
  [sym_requires_modifier] = {
    .visible = true,
    .named = true,
  },
  [sym_exports_module_directive] = {
    .visible = true,
    .named = true,
  },
  [sym_opens_module_directive] = {
    .visible = true,
    .named = true,
  },
  [sym_uses_module_directive] = {
    .visible = true,
    .named = true,
  },
  [sym_provides_module_directive] = {
    .visible = true,
    .named = true,
  },
  [sym_package_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_import_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_asterisk] = {
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
  [sym_enum_body_declarations] = {
    .visible = true,
    .named = true,
  },
  [sym_enum_constant] = {
    .visible = true,
    .named = true,
  },
  [sym_class_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_modifiers] = {
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
  [sym_type_bound] = {
    .visible = true,
    .named = true,
  },
  [sym_superclass] = {
    .visible = true,
    .named = true,
  },
  [sym_super_interfaces] = {
    .visible = true,
    .named = true,
  },
  [sym_type_list] = {
    .visible = true,
    .named = true,
  },
  [sym_permits] = {
    .visible = true,
    .named = true,
  },
  [sym_class_body] = {
    .visible = true,
    .named = true,
  },
  [sym_static_initializer] = {
    .visible = true,
    .named = true,
  },
  [sym_constructor_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym__constructor_declarator] = {
    .visible = false,
    .named = true,
  },
  [sym_constructor_body] = {
    .visible = true,
    .named = true,
  },
  [sym_explicit_constructor_invocation] = {
    .visible = true,
    .named = true,
  },
  [sym_scoped_identifier] = {
    .visible = true,
    .named = true,
  },
  [sym_field_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_record_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_annotation_type_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_annotation_type_body] = {
    .visible = true,
    .named = true,
  },
  [sym_annotation_type_element_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym__default_value] = {
    .visible = false,
    .named = true,
  },
  [sym_interface_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_extends_interfaces] = {
    .visible = true,
    .named = true,
  },
  [sym_interface_body] = {
    .visible = true,
    .named = true,
  },
  [sym_constant_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym__variable_declarator_list] = {
    .visible = false,
    .named = true,
  },
  [sym_variable_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym__variable_declarator_id] = {
    .visible = false,
    .named = true,
  },
  [sym_array_initializer] = {
    .visible = true,
    .named = true,
  },
  [sym__type] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym__unannotated_type] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym_annotated_type] = {
    .visible = true,
    .named = true,
  },
  [sym_scoped_type_identifier] = {
    .visible = true,
    .named = true,
  },
  [sym_generic_type] = {
    .visible = true,
    .named = true,
  },
  [sym_array_type] = {
    .visible = true,
    .named = true,
  },
  [sym_integral_type] = {
    .visible = true,
    .named = true,
  },
  [sym_floating_point_type] = {
    .visible = true,
    .named = true,
  },
  [sym__method_header] = {
    .visible = false,
    .named = true,
  },
  [sym__method_declarator] = {
    .visible = false,
    .named = true,
  },
  [sym_formal_parameters] = {
    .visible = true,
    .named = true,
  },
  [sym_formal_parameter] = {
    .visible = true,
    .named = true,
  },
  [sym_receiver_parameter] = {
    .visible = true,
    .named = true,
  },
  [sym_spread_parameter] = {
    .visible = true,
    .named = true,
  },
  [sym_throws] = {
    .visible = true,
    .named = true,
  },
  [sym_local_variable_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_method_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_compact_constructor_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym__reserved_identifier] = {
    .visible = false,
    .named = true,
  },
  [aux_sym_program_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__string_literal_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__multiline_string_literal_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_cast_expression_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_inferred_parameters_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_array_creation_expression_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_array_creation_expression_repeat2] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_argument_list_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_type_arguments_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_dimensions_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_switch_block_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_switch_block_repeat2] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_switch_block_statement_group_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_switch_block_statement_group_repeat2] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_record_pattern_body_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_try_statement_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_catch_type_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_resource_specification_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_for_statement_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_for_statement_repeat2] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_annotation_argument_list_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_element_value_array_initializer_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_module_body_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_requires_module_directive_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_exports_module_directive_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_provides_module_directive_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_enum_body_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_enum_body_declarations_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_modifiers_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_type_parameters_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_type_bound_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_type_list_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_annotation_type_body_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_interface_body_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__variable_declarator_list_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_array_initializer_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_formal_parameters_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_receiver_parameter_repeat1] = {
    .visible = false,
    .named = false,
  },
  [alias_sym_type_identifier] = {
    .visible = true,
    .named = true,
  },
};

enum ts_field_identifiers {
  field_alternative = 1,
  field_arguments = 2,
  field_array = 3,
  field_body = 4,
  field_condition = 5,
  field_consequence = 6,
  field_constructor = 7,
  field_declarator = 8,
  field_dimensions = 9,
  field_element = 10,
  field_field = 11,
  field_index = 12,
  field_init = 13,
  field_interfaces = 14,
  field_key = 15,
  field_left = 16,
  field_modifiers = 17,
  field_module = 18,
  field_modules = 19,
  field_name = 20,
  field_object = 21,
  field_operand = 22,
  field_operator = 23,
  field_package = 24,
  field_parameters = 25,
  field_pattern = 26,
  field_permits = 27,
  field_provided = 28,
  field_provider = 29,
  field_resources = 30,
  field_right = 31,
  field_scope = 32,
  field_superclass = 33,
  field_template_argument = 34,
  field_template_processor = 35,
  field_type = 36,
  field_type_arguments = 37,
  field_type_parameters = 38,
  field_update = 39,
  field_value = 40,
};

static const char * const ts_field_names[] = {
  [0] = NULL,
  [field_alternative] = "alternative",
  [field_arguments] = "arguments",
  [field_array] = "array",
  [field_body] = "body",
  [field_condition] = "condition",
  [field_consequence] = "consequence",
  [field_constructor] = "constructor",
  [field_declarator] = "declarator",
  [field_dimensions] = "dimensions",
  [field_element] = "element",
  [field_field] = "field",
  [field_index] = "index",
  [field_init] = "init",
  [field_interfaces] = "interfaces",
  [field_key] = "key",
  [field_left] = "left",
  [field_modifiers] = "modifiers",
  [field_module] = "module",
  [field_modules] = "modules",
  [field_name] = "name",
  [field_object] = "object",
  [field_operand] = "operand",
  [field_operator] = "operator",
  [field_package] = "package",
  [field_parameters] = "parameters",
  [field_pattern] = "pattern",
  [field_permits] = "permits",
  [field_provided] = "provided",
  [field_provider] = "provider",
  [field_resources] = "resources",
  [field_right] = "right",
  [field_scope] = "scope",
  [field_superclass] = "superclass",
  [field_template_argument] = "template_argument",
  [field_template_processor] = "template_processor",
  [field_type] = "type",
  [field_type_arguments] = "type_arguments",
  [field_type_parameters] = "type_parameters",
  [field_update] = "update",
  [field_value] = "value",
};

static const TSFieldMapSlice ts_field_map_slices[PRODUCTION_ID_COUNT] = {
  [3] = {.index = 0, .length = 3},
  [4] = {.index = 3, .length = 2},
  [5] = {.index = 5, .length = 1},
  [6] = {.index = 6, .length = 2},
  [7] = {.index = 8, .length = 1},
  [8] = {.index = 9, .length = 2},
  [9] = {.index = 11, .length = 1},
  [10] = {.index = 12, .length = 2},
  [11] = {.index = 14, .length = 4},
  [12] = {.index = 18, .length = 5},
  [13] = {.index = 23, .length = 6},
  [14] = {.index = 29, .length = 3},
  [16] = {.index = 32, .length = 2},
  [17] = {.index = 34, .length = 2},
  [18] = {.index = 32, .length = 2},
  [19] = {.index = 34, .length = 2},
  [20] = {.index = 36, .length = 2},
  [21] = {.index = 38, .length = 2},
  [22] = {.index = 40, .length = 1},
  [23] = {.index = 41, .length = 1},
  [24] = {.index = 42, .length = 2},
  [25] = {.index = 44, .length = 2},
  [26] = {.index = 46, .length = 2},
  [27] = {.index = 48, .length = 2},
  [28] = {.index = 50, .length = 3},
  [29] = {.index = 53, .length = 2},
  [31] = {.index = 55, .length = 2},
  [32] = {.index = 57, .length = 2},
  [33] = {.index = 59, .length = 2},
  [34] = {.index = 61, .length = 3},
  [35] = {.index = 64, .length = 5},
  [36] = {.index = 69, .length = 6},
  [37] = {.index = 75, .length = 5},
  [38] = {.index = 80, .length = 2},
  [39] = {.index = 82, .length = 2},
  [40] = {.index = 84, .length = 2},
  [41] = {.index = 86, .length = 2},
  [43] = {.index = 88, .length = 3},
  [44] = {.index = 91, .length = 2},
  [45] = {.index = 93, .length = 1},
  [46] = {.index = 94, .length = 2},
  [47] = {.index = 96, .length = 3},
  [48] = {.index = 99, .length = 3},
  [49] = {.index = 96, .length = 3},
  [50] = {.index = 99, .length = 3},
  [51] = {.index = 102, .length = 3},
  [52] = {.index = 102, .length = 3},
  [53] = {.index = 105, .length = 2},
  [54] = {.index = 107, .length = 2},
  [55] = {.index = 105, .length = 2},
  [56] = {.index = 107, .length = 2},
  [57] = {.index = 109, .length = 3},
  [58] = {.index = 112, .length = 3},
  [59] = {.index = 115, .length = 3},
  [60] = {.index = 118, .length = 3},
  [61] = {.index = 121, .length = 2},
  [62] = {.index = 123, .length = 2},
  [63] = {.index = 125, .length = 3},
  [64] = {.index = 128, .length = 2},
  [65] = {.index = 130, .length = 3},
  [67] = {.index = 133, .length = 2},
  [68] = {.index = 135, .length = 2},
  [69] = {.index = 137, .length = 3},
  [70] = {.index = 140, .length = 2},
  [71] = {.index = 142, .length = 2},
  [72] = {.index = 144, .length = 5},
  [73] = {.index = 149, .length = 3},
  [74] = {.index = 152, .length = 1},
  [75] = {.index = 153, .length = 2},
  [76] = {.index = 155, .length = 3},
  [78] = {.index = 158, .length = 3},
  [79] = {.index = 161, .length = 3},
  [80] = {.index = 161, .length = 3},
  [81] = {.index = 164, .length = 3},
  [82] = {.index = 167, .length = 3},
  [83] = {.index = 164, .length = 3},
  [84] = {.index = 167, .length = 3},
  [85] = {.index = 170, .length = 3},
  [86] = {.index = 170, .length = 3},
  [87] = {.index = 173, .length = 2},
  [88] = {.index = 175, .length = 4},
  [89] = {.index = 179, .length = 4},
  [90] = {.index = 183, .length = 4},
  [91] = {.index = 187, .length = 4},
  [92] = {.index = 191, .length = 4},
  [93] = {.index = 195, .length = 4},
  [94] = {.index = 199, .length = 4},
  [95] = {.index = 203, .length = 2},
  [96] = {.index = 205, .length = 3},
  [97] = {.index = 208, .length = 1},
  [98] = {.index = 209, .length = 2},
  [99] = {.index = 211, .length = 1},
  [100] = {.index = 212, .length = 4},
  [101] = {.index = 216, .length = 4},
  [102] = {.index = 220, .length = 3},
  [103] = {.index = 223, .length = 3},
  [104] = {.index = 226, .length = 2},
  [105] = {.index = 228, .length = 4},
  [106] = {.index = 232, .length = 3},
  [107] = {.index = 235, .length = 3},
  [108] = {.index = 238, .length = 3},
  [109] = {.index = 241, .length = 3},
  [110] = {.index = 244, .length = 3},
  [111] = {.index = 247, .length = 3},
  [112] = {.index = 250, .length = 3},
  [113] = {.index = 253, .length = 2},
  [114] = {.index = 255, .length = 2},
  [115] = {.index = 257, .length = 3},
  [116] = {.index = 257, .length = 3},
  [117] = {.index = 260, .length = 4},
  [118] = {.index = 264, .length = 3},
  [119] = {.index = 267, .length = 4},
  [120] = {.index = 271, .length = 5},
  [121] = {.index = 276, .length = 5},
  [122] = {.index = 281, .length = 5},
  [123] = {.index = 286, .length = 5},
  [124] = {.index = 291, .length = 4},
  [125] = {.index = 295, .length = 2},
  [126] = {.index = 297, .length = 1},
  [127] = {.index = 298, .length = 2},
  [128] = {.index = 300, .length = 2},
  [129] = {.index = 302, .length = 1},
  [130] = {.index = 303, .length = 2},
  [131] = {.index = 305, .length = 1},
  [132] = {.index = 306, .length = 3},
  [133] = {.index = 309, .length = 5},
  [134] = {.index = 314, .length = 4},
  [135] = {.index = 318, .length = 3},
  [136] = {.index = 321, .length = 4},
  [137] = {.index = 325, .length = 4},
  [138] = {.index = 329, .length = 4},
  [139] = {.index = 333, .length = 4},
  [140] = {.index = 337, .length = 4},
  [141] = {.index = 341, .length = 4},
  [142] = {.index = 345, .length = 4},
  [143] = {.index = 349, .length = 4},
  [144] = {.index = 353, .length = 3},
  [145] = {.index = 356, .length = 3},
  [146] = {.index = 359, .length = 4},
  [147] = {.index = 363, .length = 6},
  [148] = {.index = 369, .length = 4},
  [149] = {.index = 373, .length = 1},
  [150] = {.index = 374, .length = 2},
  [151] = {.index = 376, .length = 2},
  [152] = {.index = 378, .length = 1},
  [153] = {.index = 379, .length = 2},
  [154] = {.index = 381, .length = 2},
  [155] = {.index = 383, .length = 2},
  [156] = {.index = 385, .length = 3},
  [157] = {.index = 388, .length = 3},
  [158] = {.index = 391, .length = 2},
  [159] = {.index = 393, .length = 3},
  [160] = {.index = 396, .length = 4},
  [161] = {.index = 400, .length = 5},
  [162] = {.index = 405, .length = 5},
  [163] = {.index = 410, .length = 5},
  [164] = {.index = 415, .length = 5},
  [165] = {.index = 420, .length = 5},
  [166] = {.index = 425, .length = 4},
  [167] = {.index = 429, .length = 2},
  [168] = {.index = 431, .length = 3},
  [169] = {.index = 434, .length = 3},
  [170] = {.index = 437, .length = 3},
  [171] = {.index = 440, .length = 3},
  [172] = {.index = 443, .length = 3},
  [173] = {.index = 446, .length = 5},
  [174] = {.index = 451, .length = 4},
  [175] = {.index = 455, .length = 4},
  [176] = {.index = 459, .length = 2},
  [177] = {.index = 461, .length = 1},
  [178] = {.index = 462, .length = 2},
  [179] = {.index = 464, .length = 6},
  [180] = {.index = 470, .length = 3},
  [181] = {.index = 473, .length = 4},
  [182] = {.index = 477, .length = 4},
  [183] = {.index = 481, .length = 4},
  [184] = {.index = 485, .length = 4},
  [185] = {.index = 489, .length = 4},
  [186] = {.index = 493, .length = 5},
  [187] = {.index = 498, .length = 5},
  [188] = {.index = 503, .length = 1},
  [189] = {.index = 504, .length = 3},
  [190] = {.index = 507, .length = 2},
  [191] = {.index = 509, .length = 1},
  [192] = {.index = 510, .length = 2},
  [193] = {.index = 512, .length = 2},
  [194] = {.index = 514, .length = 2},
  [195] = {.index = 516, .length = 1},
  [196] = {.index = 517, .length = 3},
  [197] = {.index = 520, .length = 3},
  [198] = {.index = 523, .length = 3},
  [199] = {.index = 526, .length = 5},
  [200] = {.index = 531, .length = 5},
  [201] = {.index = 536, .length = 5},
  [202] = {.index = 541, .length = 3},
  [203] = {.index = 544, .length = 3},
  [204] = {.index = 547, .length = 4},
  [205] = {.index = 551, .length = 4},
  [206] = {.index = 555, .length = 6},
  [207] = {.index = 561, .length = 4},
};

static const TSFieldMapEntry ts_field_map_entries[] = {
  [0] =
    {field_arguments, 0, .inherited = true},
    {field_type, 0, .inherited = true},
    {field_type_arguments, 0, .inherited = true},
  [3] =
    {field_operand, 1},
    {field_operator, 0},
  [5] =
    {field_name, 1},
  [6] =
    {field_arguments, 1},
    {field_name, 0},
  [8] =
    {field_name, 0},
  [9] =
    {field_dimensions, 1},
    {field_element, 0},
  [11] =
    {field_declarator, 0},
  [12] =
    {field_dimensions, 0, .inherited = true},
    {field_name, 0, .inherited = true},
  [14] =
    {field_dimensions, 1, .inherited = true},
    {field_name, 1, .inherited = true},
    {field_parameters, 1, .inherited = true},
    {field_type, 0},
  [18] =
    {field_dimensions, 0, .inherited = true},
    {field_name, 0, .inherited = true},
    {field_parameters, 0, .inherited = true},
    {field_type, 0, .inherited = true},
    {field_type_parameters, 0, .inherited = true},
  [23] =
    {field_body, 1},
    {field_dimensions, 0, .inherited = true},
    {field_name, 0, .inherited = true},
    {field_parameters, 0, .inherited = true},
    {field_type, 0, .inherited = true},
    {field_type_parameters, 0, .inherited = true},
  [29] =
    {field_dimensions, 1, .inherited = true},
    {field_name, 1, .inherited = true},
    {field_type, 0},
  [32] =
    {field_arguments, 2},
    {field_type, 1},
  [34] =
    {field_dimensions, 2},
    {field_type, 1},
  [36] =
    {field_body, 2},
    {field_name, 1},
  [38] =
    {field_body, 2},
    {field_condition, 1},
  [40] =
    {field_body, 2},
  [41] =
    {field_body, 1},
  [42] =
    {field_body, 2},
    {field_resources, 1},
  [44] =
    {field_condition, 1},
    {field_consequence, 2},
  [46] =
    {field_arguments, 2},
    {field_name, 1},
  [48] =
    {field_field, 2},
    {field_object, 0},
  [50] =
    {field_left, 0},
    {field_operator, 1},
    {field_right, 2},
  [53] =
    {field_body, 2},
    {field_parameters, 0},
  [55] =
    {field_left, 0},
    {field_pattern, 2},
  [57] =
    {field_left, 0},
    {field_right, 2},
  [59] =
    {field_template_argument, 2},
    {field_template_processor, 0},
  [61] =
    {field_arguments, 2, .inherited = true},
    {field_type, 2, .inherited = true},
    {field_type_arguments, 2, .inherited = true},
  [64] =
    {field_dimensions, 1, .inherited = true},
    {field_name, 1, .inherited = true},
    {field_parameters, 1, .inherited = true},
    {field_type, 1, .inherited = true},
    {field_type_parameters, 1, .inherited = true},
  [69] =
    {field_body, 2},
    {field_dimensions, 1, .inherited = true},
    {field_name, 1, .inherited = true},
    {field_parameters, 1, .inherited = true},
    {field_type, 1, .inherited = true},
    {field_type_parameters, 1, .inherited = true},
  [75] =
    {field_dimensions, 2, .inherited = true},
    {field_name, 2, .inherited = true},
    {field_parameters, 2, .inherited = true},
    {field_type, 1},
    {field_type_parameters, 0},
  [80] =
    {field_dimensions, 1},
    {field_name, 0},
  [82] =
    {field_name, 0},
    {field_parameters, 1},
  [84] =
    {field_declarator, 1, .inherited = true},
    {field_type, 0},
  [86] =
    {field_declarator, 0},
    {field_declarator, 1, .inherited = true},
  [88] =
    {field_dimensions, 2, .inherited = true},
    {field_name, 2, .inherited = true},
    {field_type, 1},
  [91] =
    {field_type, 1},
    {field_value, 3},
  [93] =
    {field_type, 1},
  [94] =
    {field_type, 0, .inherited = true},
    {field_type, 1, .inherited = true},
  [96] =
    {field_dimensions, 2},
    {field_type, 1},
    {field_value, 3},
  [99] =
    {field_dimensions, 2},
    {field_dimensions, 3},
    {field_type, 1},
  [102] =
    {field_arguments, 3},
    {field_type, 2},
    {field_type_arguments, 1},
  [105] =
    {field_arguments, 3},
    {field_type, 2},
  [107] =
    {field_dimensions, 3},
    {field_type, 2},
  [109] =
    {field_body, 3},
    {field_name, 1},
    {field_type_parameters, 2},
  [112] =
    {field_body, 3},
    {field_name, 1},
    {field_superclass, 2},
  [115] =
    {field_body, 3},
    {field_interfaces, 2},
    {field_name, 1},
  [118] =
    {field_body, 3},
    {field_name, 1},
    {field_permits, 2},
  [121] =
    {field_name, 2},
    {field_scope, 0},
  [123] =
    {field_body, 3},
    {field_name, 2},
  [125] =
    {field_body, 3},
    {field_name, 1},
    {field_parameters, 2},
  [128] =
    {field_body, 3},
    {field_name, 1},
  [130] =
    {field_arguments, 3},
    {field_name, 2},
    {field_object, 0},
  [133] =
    {field_left, 0},
    {field_pattern, 3},
  [135] =
    {field_left, 0},
    {field_right, 3},
  [137] =
    {field_left, 0},
    {field_name, 3},
    {field_right, 2},
  [140] =
    {field_array, 0},
    {field_index, 2},
  [142] =
    {field_declarator, 2, .inherited = true},
    {field_type, 1},
  [144] =
    {field_dimensions, 3, .inherited = true},
    {field_name, 3, .inherited = true},
    {field_parameters, 3, .inherited = true},
    {field_type, 2},
    {field_type_parameters, 0},
  [149] =
    {field_dimensions, 2},
    {field_name, 0},
    {field_parameters, 1},
  [152] =
    {field_declarator, 1},
  [153] =
    {field_declarator, 0, .inherited = true},
    {field_declarator, 1, .inherited = true},
  [155] =
    {field_dimensions, 0, .inherited = true},
    {field_name, 0, .inherited = true},
    {field_value, 2},
  [158] =
    {field_type, 1},
    {field_type, 2, .inherited = true},
    {field_value, 4},
  [161] =
    {field_arguments, 4},
    {field_type, 3},
    {field_type_arguments, 1},
  [164] =
    {field_dimensions, 3},
    {field_type, 2},
    {field_value, 4},
  [167] =
    {field_dimensions, 3},
    {field_dimensions, 4},
    {field_type, 2},
  [170] =
    {field_arguments, 4},
    {field_type, 3},
    {field_type_arguments, 2},
  [173] =
    {field_body, 1},
    {field_name, 0},
  [175] =
    {field_body, 1},
    {field_name, 0, .inherited = true},
    {field_parameters, 0, .inherited = true},
    {field_type_parameters, 0, .inherited = true},
  [179] =
    {field_body, 4},
    {field_name, 1},
    {field_superclass, 3},
    {field_type_parameters, 2},
  [183] =
    {field_body, 4},
    {field_interfaces, 3},
    {field_name, 1},
    {field_type_parameters, 2},
  [187] =
    {field_body, 4},
    {field_name, 1},
    {field_permits, 3},
    {field_type_parameters, 2},
  [191] =
    {field_body, 4},
    {field_interfaces, 3},
    {field_name, 1},
    {field_superclass, 2},
  [195] =
    {field_body, 4},
    {field_name, 1},
    {field_permits, 3},
    {field_superclass, 2},
  [199] =
    {field_body, 4},
    {field_interfaces, 2},
    {field_name, 1},
    {field_permits, 3},
  [203] =
    {field_body, 1},
    {field_condition, 3},
  [205] =
    {field_alternative, 4},
    {field_condition, 1},
    {field_consequence, 2},
  [208] =
    {field_init, 1},
  [209] =
    {field_init, 0, .inherited = true},
    {field_init, 1, .inherited = true},
  [211] =
    {field_modifiers, 0},
  [212] =
    {field_body, 4},
    {field_name, 1},
    {field_parameters, 3},
    {field_type_parameters, 2},
  [216] =
    {field_body, 4},
    {field_interfaces, 3},
    {field_name, 1},
    {field_parameters, 2},
  [220] =
    {field_body, 4},
    {field_name, 1},
    {field_type_parameters, 2},
  [223] =
    {field_body, 4},
    {field_name, 1},
    {field_permits, 3},
  [226] =
    {field_field, 4},
    {field_object, 0},
  [228] =
    {field_arguments, 4},
    {field_name, 3},
    {field_object, 0},
    {field_type_arguments, 2},
  [232] =
    {field_left, 0},
    {field_name, 4},
    {field_right, 3},
  [235] =
    {field_alternative, 4},
    {field_condition, 0},
    {field_consequence, 2},
  [238] =
    {field_body, 4},
    {field_name, 2},
    {field_type_parameters, 3},
  [241] =
    {field_body, 4},
    {field_name, 2},
    {field_superclass, 3},
  [244] =
    {field_body, 4},
    {field_interfaces, 3},
    {field_name, 2},
  [247] =
    {field_body, 4},
    {field_name, 2},
    {field_permits, 3},
  [250] =
    {field_body, 4},
    {field_name, 2},
    {field_parameters, 3},
  [253] =
    {field_body, 4},
    {field_name, 2},
  [255] =
    {field_body, 4},
    {field_name, 3},
  [257] =
    {field_arguments, 5},
    {field_type, 4},
    {field_type_arguments, 2},
  [260] =
    {field_body, 2},
    {field_name, 1, .inherited = true},
    {field_parameters, 1, .inherited = true},
    {field_type_parameters, 1, .inherited = true},
  [264] =
    {field_name, 1},
    {field_parameters, 2},
    {field_type_parameters, 0},
  [267] =
    {field_body, 2},
    {field_name, 0, .inherited = true},
    {field_parameters, 0, .inherited = true},
    {field_type_parameters, 0, .inherited = true},
  [271] =
    {field_body, 5},
    {field_interfaces, 4},
    {field_name, 1},
    {field_superclass, 3},
    {field_type_parameters, 2},
  [276] =
    {field_body, 5},
    {field_name, 1},
    {field_permits, 4},
    {field_superclass, 3},
    {field_type_parameters, 2},
  [281] =
    {field_body, 5},
    {field_interfaces, 3},
    {field_name, 1},
    {field_permits, 4},
    {field_type_parameters, 2},
  [286] =
    {field_body, 5},
    {field_interfaces, 3},
    {field_name, 1},
    {field_permits, 4},
    {field_superclass, 2},
  [291] =
    {field_dimensions, 1, .inherited = true},
    {field_name, 1, .inherited = true},
    {field_type, 0},
    {field_value, 3},
  [295] =
    {field_dimensions, 1, .inherited = true},
    {field_name, 1, .inherited = true},
  [297] =
    {field_body, 5},
  [298] =
    {field_body, 5},
    {field_init, 2},
  [300] =
    {field_key, 0},
    {field_value, 2},
  [302] =
    {field_module, 1},
  [303] =
    {field_modifiers, 0, .inherited = true},
    {field_modifiers, 1, .inherited = true},
  [305] =
    {field_package, 1},
  [306] =
    {field_arguments, 1},
    {field_body, 2},
    {field_name, 0},
  [309] =
    {field_body, 5},
    {field_interfaces, 4},
    {field_name, 1},
    {field_parameters, 3},
    {field_type_parameters, 2},
  [314] =
    {field_body, 5},
    {field_name, 1},
    {field_permits, 4},
    {field_type_parameters, 2},
  [318] =
    {field_arguments, 5},
    {field_name, 4},
    {field_object, 0},
  [321] =
    {field_body, 5},
    {field_name, 2},
    {field_superclass, 4},
    {field_type_parameters, 3},
  [325] =
    {field_body, 5},
    {field_interfaces, 4},
    {field_name, 2},
    {field_type_parameters, 3},
  [329] =
    {field_body, 5},
    {field_name, 2},
    {field_permits, 4},
    {field_type_parameters, 3},
  [333] =
    {field_body, 5},
    {field_interfaces, 4},
    {field_name, 2},
    {field_superclass, 3},
  [337] =
    {field_body, 5},
    {field_name, 2},
    {field_permits, 4},
    {field_superclass, 3},
  [341] =
    {field_body, 5},
    {field_interfaces, 3},
    {field_name, 2},
    {field_permits, 4},
  [345] =
    {field_body, 5},
    {field_name, 2},
    {field_parameters, 4},
    {field_type_parameters, 3},
  [349] =
    {field_body, 5},
    {field_interfaces, 4},
    {field_name, 2},
    {field_parameters, 3},
  [353] =
    {field_body, 5},
    {field_name, 2},
    {field_type_parameters, 3},
  [356] =
    {field_body, 5},
    {field_name, 2},
    {field_permits, 4},
  [359] =
    {field_body, 3},
    {field_name, 1, .inherited = true},
    {field_parameters, 1, .inherited = true},
    {field_type_parameters, 1, .inherited = true},
  [363] =
    {field_body, 6},
    {field_interfaces, 4},
    {field_name, 1},
    {field_permits, 5},
    {field_superclass, 3},
    {field_type_parameters, 2},
  [369] =
    {field_dimensions, 2, .inherited = true},
    {field_name, 2, .inherited = true},
    {field_type, 1},
    {field_value, 4},
  [373] =
    {field_body, 4},
  [374] =
    {field_dimensions, 2, .inherited = true},
    {field_name, 2, .inherited = true},
  [376] =
    {field_body, 6},
    {field_update, 4},
  [378] =
    {field_update, 1},
  [379] =
    {field_update, 0, .inherited = true},
    {field_update, 1, .inherited = true},
  [381] =
    {field_body, 6},
    {field_condition, 3},
  [383] =
    {field_body, 6},
    {field_init, 2},
  [385] =
    {field_body, 6},
    {field_init, 2},
    {field_update, 4},
  [388] =
    {field_body, 6},
    {field_condition, 3},
    {field_init, 2},
  [391] =
    {field_modifiers, 1, .inherited = true},
    {field_module, 2},
  [393] =
    {field_arguments, 2},
    {field_body, 3},
    {field_name, 1},
  [396] =
    {field_arguments, 6},
    {field_name, 5},
    {field_object, 0},
    {field_type_arguments, 4},
  [400] =
    {field_body, 6},
    {field_interfaces, 5},
    {field_name, 2},
    {field_superclass, 4},
    {field_type_parameters, 3},
  [405] =
    {field_body, 6},
    {field_name, 2},
    {field_permits, 5},
    {field_superclass, 4},
    {field_type_parameters, 3},
  [410] =
    {field_body, 6},
    {field_interfaces, 4},
    {field_name, 2},
    {field_permits, 5},
    {field_type_parameters, 3},
  [415] =
    {field_body, 6},
    {field_interfaces, 4},
    {field_name, 2},
    {field_permits, 5},
    {field_superclass, 3},
  [420] =
    {field_body, 6},
    {field_interfaces, 5},
    {field_name, 2},
    {field_parameters, 4},
    {field_type_parameters, 3},
  [425] =
    {field_body, 6},
    {field_name, 2},
    {field_permits, 5},
    {field_type_parameters, 3},
  [429] =
    {field_arguments, 1},
    {field_constructor, 0},
  [431] =
    {field_body, 7},
    {field_update, 4},
    {field_update, 5, .inherited = true},
  [434] =
    {field_body, 7},
    {field_condition, 3},
    {field_update, 5},
  [437] =
    {field_body, 7},
    {field_init, 2},
    {field_update, 5},
  [440] =
    {field_body, 7},
    {field_condition, 4},
    {field_init, 2},
  [443] =
    {field_body, 7},
    {field_init, 2},
    {field_init, 3, .inherited = true},
  [446] =
    {field_body, 7},
    {field_dimensions, 3, .inherited = true},
    {field_name, 3, .inherited = true},
    {field_type, 2},
    {field_value, 5},
  [451] =
    {field_body, 7},
    {field_init, 2},
    {field_update, 4},
    {field_update, 5, .inherited = true},
  [455] =
    {field_body, 7},
    {field_condition, 3},
    {field_init, 2},
    {field_update, 5},
  [459] =
    {field_modules, 3},
    {field_package, 1},
  [461] =
    {field_provided, 1},
  [462] =
    {field_name, 1},
    {field_type, 0},
  [464] =
    {field_body, 7},
    {field_interfaces, 5},
    {field_name, 2},
    {field_permits, 6},
    {field_superclass, 4},
    {field_type_parameters, 3},
  [470] =
    {field_arguments, 2},
    {field_constructor, 1},
    {field_type_arguments, 0},
  [473] =
    {field_body, 8},
    {field_condition, 3},
    {field_update, 5},
    {field_update, 6, .inherited = true},
  [477] =
    {field_body, 8},
    {field_init, 2},
    {field_update, 5},
    {field_update, 6, .inherited = true},
  [481] =
    {field_body, 8},
    {field_condition, 4},
    {field_init, 2},
    {field_update, 6},
  [485] =
    {field_body, 8},
    {field_init, 2},
    {field_init, 3, .inherited = true},
    {field_update, 6},
  [489] =
    {field_body, 8},
    {field_condition, 5},
    {field_init, 2},
    {field_init, 3, .inherited = true},
  [493] =
    {field_body, 8},
    {field_dimensions, 4, .inherited = true},
    {field_name, 4, .inherited = true},
    {field_type, 3},
    {field_value, 6},
  [498] =
    {field_body, 8},
    {field_condition, 3},
    {field_init, 2},
    {field_update, 5},
    {field_update, 6, .inherited = true},
  [503] =
    {field_modules, 1},
  [504] =
    {field_modules, 3},
    {field_modules, 4, .inherited = true},
    {field_package, 1},
  [507] =
    {field_modules, 0, .inherited = true},
    {field_modules, 1, .inherited = true},
  [509] =
    {field_provider, 1},
  [510] =
    {field_provided, 1},
    {field_provider, 4, .inherited = true},
  [512] =
    {field_provider, 0, .inherited = true},
    {field_provider, 1, .inherited = true},
  [514] =
    {field_name, 2},
    {field_type, 1},
  [516] =
    {field_value, 1},
  [517] =
    {field_dimensions, 4},
    {field_name, 1},
    {field_type, 0},
  [520] =
    {field_name, 1},
    {field_type, 0},
    {field_value, 4, .inherited = true},
  [523] =
    {field_arguments, 3},
    {field_constructor, 2},
    {field_object, 0},
  [526] =
    {field_body, 9},
    {field_condition, 4},
    {field_init, 2},
    {field_update, 6},
    {field_update, 7, .inherited = true},
  [531] =
    {field_body, 9},
    {field_init, 2},
    {field_init, 3, .inherited = true},
    {field_update, 6},
    {field_update, 7, .inherited = true},
  [536] =
    {field_body, 9},
    {field_condition, 5},
    {field_init, 2},
    {field_init, 3, .inherited = true},
    {field_update, 7},
  [541] =
    {field_dimensions, 5},
    {field_name, 2},
    {field_type, 1},
  [544] =
    {field_name, 2},
    {field_type, 1},
    {field_value, 5, .inherited = true},
  [547] =
    {field_dimensions, 4},
    {field_name, 1},
    {field_type, 0},
    {field_value, 5, .inherited = true},
  [551] =
    {field_arguments, 4},
    {field_constructor, 3},
    {field_object, 0},
    {field_type_arguments, 2},
  [555] =
    {field_body, 10},
    {field_condition, 5},
    {field_init, 2},
    {field_init, 3, .inherited = true},
    {field_update, 7},
    {field_update, 8, .inherited = true},
  [561] =
    {field_dimensions, 5},
    {field_name, 2},
    {field_type, 1},
    {field_value, 6, .inherited = true},
};

static const TSSymbol ts_alias_sequences[PRODUCTION_ID_COUNT][MAX_ALIAS_SEQUENCE_LENGTH] = {
  [0] = {0},
  [1] = {
    [0] = sym_identifier,
  },
  [2] = {
    [0] = alias_sym_type_identifier,
  },
  [15] = {
    [1] = alias_sym_type_identifier,
  },
  [18] = {
    [1] = alias_sym_type_identifier,
  },
  [19] = {
    [1] = alias_sym_type_identifier,
  },
  [30] = {
    [0] = alias_sym_type_identifier,
    [2] = alias_sym_type_identifier,
  },
  [42] = {
    [2] = alias_sym_type_identifier,
  },
  [49] = {
    [1] = alias_sym_type_identifier,
  },
  [50] = {
    [1] = alias_sym_type_identifier,
  },
  [52] = {
    [2] = alias_sym_type_identifier,
  },
  [55] = {
    [2] = alias_sym_type_identifier,
  },
  [56] = {
    [2] = alias_sym_type_identifier,
  },
  [66] = {
    [0] = alias_sym_type_identifier,
    [3] = alias_sym_type_identifier,
  },
  [77] = {
    [3] = alias_sym_type_identifier,
  },
  [80] = {
    [3] = alias_sym_type_identifier,
  },
  [83] = {
    [2] = alias_sym_type_identifier,
  },
  [84] = {
    [2] = alias_sym_type_identifier,
  },
  [86] = {
    [3] = alias_sym_type_identifier,
  },
  [116] = {
    [4] = alias_sym_type_identifier,
  },
};

static const uint16_t ts_non_terminal_alias_map[] = {
  0,
};

// Filled from tables.zlib by grammar.go, see fetchgrammar.
static TSStateId ts_primary_state_ids[STATE_COUNT];

static TSCharacterRange sym_identifier_character_set_1[] = {
  {'$', '$'}, {'A', 'Z'}, {'_', '_'}, {'a', 'z'}, {0xaa, 0xaa}, {0xb5, 0xb5}, {0xba, 0xba}, {0xc0, 0xd6},
  {0xd8, 0xf6}, {0xf8, 0x2c1}, {0x2c6, 0x2d1}, {0x2e0, 0x2e4}, {0x2ec, 0x2ec}, {0x2ee, 0x2ee}, {0x370, 0x374}, {0x376, 0x377},
  {0x37b, 0x37d}, {0x37f, 0x37f}, {0x386, 0x386}, {0x388, 0x38a}, {0x38c, 0x38c}, {0x38e, 0x3a1}, {0x3a3, 0x3f5}, {0x3f7, 0x481},
  {0x48a, 0x52f}, {0x531, 0x556}, {0x559, 0x559}, {0x560, 0x588}, {0x5d0, 0x5ea}, {0x5ef, 0x5f2}, {0x620, 0x64a}, {0x66e, 0x66f},
  {0x671, 0x6d3}, {0x6d5, 0x6d5}, {0x6e5, 0x6e6}, {0x6ee, 0x6ef}, {0x6fa, 0x6fc}, {0x6ff, 0x6ff}, {0x710, 0x710}, {0x712, 0x72f},
  {0x74d, 0x7a5}, {0x7b1, 0x7b1}, {0x7ca, 0x7ea}, {0x7f4, 0x7f5}, {0x7fa, 0x7fa}, {0x800, 0x815}, {0x81a, 0x81a}, {0x824, 0x824},
  {0x828, 0x828}, {0x840, 0x858}, {0x860, 0x86a}, {0x870, 0x887}, {0x889, 0x88e}, {0x8a0, 0x8c9}, {0x904, 0x939}, {0x93d, 0x93d},
  {0x950, 0x950}, {0x958, 0x961}, {0x971, 0x980}, {0x985, 0x98c}, {0x98f, 0x990}, {0x993, 0x9a8}, {0x9aa, 0x9b0}, {0x9b2, 0x9b2},
  {0x9b6, 0x9b9}, {0x9bd, 0x9bd}, {0x9ce, 0x9ce}, {0x9dc, 0x9dd}, {0x9df, 0x9e1}, {0x9f0, 0x9f1}, {0x9fc, 0x9fc}, {0xa05, 0xa0a},
  {0xa0f, 0xa10}, {0xa13, 0xa28}, {0xa2a, 0xa30}, {0xa32, 0xa33}, {0xa35, 0xa36}, {0xa38, 0xa39}, {0xa59, 0xa5c}, {0xa5e, 0xa5e},
  {0xa72, 0xa74}, {0xa85, 0xa8d}, {0xa8f, 0xa91}, {0xa93, 0xaa8}, {0xaaa, 0xab0}, {0xab2, 0xab3}, {0xab5, 0xab9}, {0xabd, 0xabd},
  {0xad0, 0xad0}, {0xae0, 0xae1}, {0xaf9, 0xaf9}, {0xb05, 0xb0c}, {0xb0f, 0xb10}, {0xb13, 0xb28}, {0xb2a, 0xb30}, {0xb32, 0xb33},
  {0xb35, 0xb39}, {0xb3d, 0xb3d}, {0xb5c, 0xb5d}, {0xb5f, 0xb61}, {0xb71, 0xb71}, {0xb83, 0xb83}, {0xb85, 0xb8a}, {0xb8e, 0xb90},
  {0xb92, 0xb95}, {0xb99, 0xb9a}, {0xb9c, 0xb9c}, {0xb9e, 0xb9f}, {0xba3, 0xba4}, {0xba8, 0xbaa}, {0xbae, 0xbb9}, {0xbd0, 0xbd0},
  {0xc05, 0xc0c}, {0xc0e, 0xc10}, {0xc12, 0xc28}, {0xc2a, 0xc39}, {0xc3d, 0xc3d}, {0xc58, 0xc5a}, {0xc5d, 0xc5d}, {0xc60, 0xc61},
  {0xc80, 0xc80}, {0xc85, 0xc8c}, {0xc8e, 0xc90}, {0xc92, 0xca8}, {0xcaa, 0xcb3}, {0xcb5, 0xcb9}, {0xcbd, 0xcbd}, {0xcdd, 0xcde},
  {0xce0, 0xce1}, {0xcf1, 0xcf2}, {0xd04, 0xd0c}, {0xd0e, 0xd10}, {0xd12, 0xd3a}, {0xd3d, 0xd3d}, {0xd4e, 0xd4e}, {0xd54, 0xd56},
  {0xd5f, 0xd61}, {0xd7a, 0xd7f}, {0xd85, 0xd96}, {0xd9a, 0xdb1}, {0xdb3, 0xdbb}, {0xdbd, 0xdbd}, {0xdc0, 0xdc6}, {0xe01, 0xe30},
  {0xe32, 0xe32}, {0xe40, 0xe46}, {0xe81, 0xe82}, {0xe84, 0xe84}, {0xe86, 0xe8a}, {0xe8c, 0xea3}, {0xea5, 0xea5}, {0xea7, 0xeb0},
  {0xeb2, 0xeb2}, {0xebd, 0xebd}, {0xec0, 0xec4}, {0xec6, 0xec6}, {0xedc, 0xedf}, {0xf00, 0xf00}, {0xf40, 0xf47}, {0xf49, 0xf6c},
  {0xf88, 0xf8c}, {0x1000, 0x102a}, {0x103f, 0x103f}, {0x1050, 0x1055}, {0x105a, 0x105d}, {0x1061, 0x1061}, {0x1065, 0x1066}, {0x106e, 0x1070},
  {0x1075, 0x1081}, {0x108e, 0x108e}, {0x10a0, 0x10c5}, {0x10c7, 0x10c7}, {0x10cd, 0x10cd}, {0x10d0, 0x10fa}, {0x10fc, 0x1248}, {0x124a, 0x124d},
  {0x1250, 0x1256}, {0x1258, 0x1258}, {0x125a, 0x125d}, {0x1260, 0x1288}, {0x128a, 0x128d}, {0x1290, 0x12b0}, {0x12b2, 0x12b5}, {0x12b8, 0x12be},
  {0x12c0, 0x12c0}, {0x12c2, 0x12c5}, {0x12c8, 0x12d6}, {0x12d8, 0x1310}, {0x1312, 0x1315}, {0x1318, 0x135a}, {0x1380, 0x138f}, {0x13a0, 0x13f5},
  {0x13f8, 0x13fd}, {0x1401, 0x166c}, {0x166f, 0x167f}, {0x1681, 0x169a}, {0x16a0, 0x16ea}, {0x16ee, 0x16f8}, {0x1700, 0x1711}, {0x171f, 0x1731},
  {0x1740, 0x1751}, {0x1760, 0x176c}, {0x176e, 0x1770}, {0x1780, 0x17b3}, {0x17d7, 0x17d7}, {0x17dc, 0x17dc}, {0x1820, 0x1878}, {0x1880, 0x18a8},
  {0x18aa, 0x18aa}, {0x18b0, 0x18f5}, {0x1900, 0x191e}, {0x1950, 0x196d}, {0x1970, 0x1974}, {0x1980, 0x19ab}, {0x19b0, 0x19c9}, {0x1a00, 0x1a16},
  {0x1a20, 0x1a54}, {0x1aa7, 0x1aa7}, {0x1b05, 0x1b33}, {0x1b45, 0x1b4c}, {0x1b83, 0x1ba0}, {0x1bae, 0x1baf}, {0x1bba, 0x1be5}, {0x1c00, 0x1c23},
  {0x1c4d, 0x1c4f}, {0x1c5a, 0x1c7d}, {0x1c80, 0x1c88}, {0x1c90, 0x1cba}, {0x1cbd, 0x1cbf}, {0x1ce9, 0x1cec}, {0x1cee, 0x1cf3}, {0x1cf5, 0x1cf6},
  {0x1cfa, 0x1cfa}, {0x1d00, 0x1dbf}, {0x1e00, 0x1f15}, {0x1f18, 0x1f1d}, {0x1f20, 0x1f45}, {0x1f48, 0x1f4d}, {0x1f50, 0x1f57}, {0x1f59, 0x1f59},
  {0x1f5b, 0x1f5b}, {0x1f5d, 0x1f5d}, {0x1f5f, 0x1f7d}, {0x1f80, 0x1fb4}, {0x1fb6, 0x1fbc}, {0x1fbe, 0x1fbe}, {0x1fc2, 0x1fc4}, {0x1fc6, 0x1fcc},
  {0x1fd0, 0x1fd3}, {0x1fd6, 0x1fdb}, {0x1fe0, 0x1fec}, {0x1ff2, 0x1ff4}, {0x1ff6, 0x1ffc}, {0x2071, 0x2071}, {0x207f, 0x207f}, {0x2090, 0x209c},
  {0x2102, 0x2102}, {0x2107, 0x2107}, {0x210a, 0x2113}, {0x2115, 0x2115}, {0x2118, 0x211d}, {0x2124, 0x2124}, {0x2126, 0x2126}, {0x2128, 0x2128},
  {0x212a, 0x2139}, {0x213c, 0x213f}, {0x2145, 0x2149}, {0x214e, 0x214e}, {0x2160, 0x2188}, {0x2c00, 0x2ce4}, {0x2ceb, 0x2cee}, {0x2cf2, 0x2cf3},
  {0x2d00, 0x2d25}, {0x2d27, 0x2d27}, {0x2d2d, 0x2d2d}, {0x2d30, 0x2d67}, {0x2d6f, 0x2d6f}, {0x2d80, 0x2d96}, {0x2da0, 0x2da6}, {0x2da8, 0x2dae},
  {0x2db0, 0x2db6}, {0x2db8, 0x2dbe}, {0x2dc0, 0x2dc6}, {0x2dc8, 0x2dce}, {0x2dd0, 0x2dd6}, {0x2dd8, 0x2dde}, {0x3005, 0x3007}, {0x3021, 0x3029},
  {0x3031, 0x3035}, {0x3038, 0x303c}, {0x3041, 0x3096}, {0x309d, 0x309f}, {0x30a1, 0x30fa}, {0x30fc, 0x30ff}, {0x3105, 0x312f}, {0x3131, 0x318e},
  {0x31a0, 0x31bf}, {0x31f0, 0x31ff}, {0x3400, 0x4dbf}, {0x4e00, 0xa48c}, {0xa4d0, 0xa4fd}, {0xa500, 0xa60c}, {0xa610, 0xa61f}, {0xa62a, 0xa62b},
  {0xa640, 0xa66e}, {0xa67f, 0xa69d}, {0xa6a0, 0xa6ef}, {0xa717, 0xa71f}, {0xa722, 0xa788}, {0xa78b, 0xa7ca}, {0xa7d0, 0xa7d1}, {0xa7d3, 0xa7d3},
  {0xa7d5, 0xa7d9}, {0xa7f2, 0xa801}, {0xa803, 0xa805}, {0xa807, 0xa80a}, {0xa80c, 0xa822}, {0xa840, 0xa873}, {0xa882, 0xa8b3}, {0xa8f2, 0xa8f7},
  {0xa8fb, 0xa8fb}, {0xa8fd, 0xa8fe}, {0xa90a, 0xa925}, {0xa930, 0xa946}, {0xa960, 0xa97c}, {0xa984, 0xa9b2}, {0xa9cf, 0xa9cf}, {0xa9e0, 0xa9e4},
  {0xa9e6, 0xa9ef}, {0xa9fa, 0xa9fe}, {0xaa00, 0xaa28}, {0xaa40, 0xaa42}, {0xaa44, 0xaa4b}, {0xaa60, 0xaa76}, {0xaa7a, 0xaa7a}, {0xaa7e, 0xaaaf},
  {0xaab1, 0xaab1}, {0xaab5, 0xaab6}, {0xaab9, 0xaabd}, {0xaac0, 0xaac0}, {0xaac2, 0xaac2}, {0xaadb, 0xaadd}, {0xaae0, 0xaaea}, {0xaaf2, 0xaaf4},
  {0xab01, 0xab06}, {0xab09, 0xab0e}, {0xab11, 0xab16}, {0xab20, 0xab26}, {0xab28, 0xab2e}, {0xab30, 0xab5a}, {0xab5c, 0xab69}, {0xab70, 0xabe2},
  {0xac00, 0xd7a3}, {0xd7b0, 0xd7c6}, {0xd7cb, 0xd7fb}, {0xf900, 0xfa6d}, {0xfa70, 0xfad9}, {0xfb00, 0xfb06}, {0xfb13, 0xfb17}, {0xfb1d, 0xfb1d},
  {0xfb1f, 0xfb28}, {0xfb2a, 0xfb36}, {0xfb38, 0xfb3c}, {0xfb3e, 0xfb3e}, {0xfb40, 0xfb41}, {0xfb43, 0xfb44}, {0xfb46, 0xfbb1}, {0xfbd3, 0xfc5d},
  {0xfc64, 0xfd3d}, {0xfd50, 0xfd8f}, {0xfd92, 0xfdc7}, {0xfdf0, 0xfdf9}, {0xfe71, 0xfe71}, {0xfe73, 0xfe73}, {0xfe77, 0xfe77}, {0xfe79, 0xfe79},
  {0xfe7b, 0xfe7b}, {0xfe7d, 0xfe7d}, {0xfe7f, 0xfefc}, {0xff21, 0xff3a}, {0xff41, 0xff5a}, {0xff66, 0xff9d}, {0xffa0, 0xffbe}, {0xffc2, 0xffc7},
  {0xffca, 0xffcf}, {0xffd2, 0xffd7}, {0xffda, 0xffdc}, {0x10000, 0x1000b}, {0x1000d, 0x10026}, {0x10028, 0x1003a}, {0x1003c, 0x1003d}, {0x1003f, 0x1004d},
  {0x10050, 0x1005d}, {0x10080, 0x100fa}, {0x10140, 0x10174}, {0x10280, 0x1029c}, {0x102a0, 0x102d0}, {0x10300, 0x1031f}, {0x1032d, 0x1034a}, {0x10350, 0x10375},
  {0x10380, 0x1039d}, {0x103a0, 0x103c3}, {0x103c8, 0x103cf}, {0x103d1, 0x103d5}, {0x10400, 0x1049d}, {0x104b0, 0x104d3}, {0x104d8, 0x104fb}, {0x10500, 0x10527},
  {0x10530, 0x10563}, {0x10570, 0x1057a}, {0x1057c, 0x1058a}, {0x1058c, 0x10592}, {0x10594, 0x10595}, {0x10597, 0x105a1}, {0x105a3, 0x105b1}, {0x105b3, 0x105b9},
  {0x105bb, 0x105bc}, {0x10600, 0x10736}, {0x10740, 0x10755}, {0x10760, 0x10767}, {0x10780, 0x10785}, {0x10787, 0x107b0}, {0x107b2, 0x107ba}, {0x10800, 0x10805},
  {0x10808, 0x10808}, {0x1080a, 0x10835}, {0x10837, 0x10838}, {0x1083c, 0x1083c}, {0x1083f, 0x10855}, {0x10860, 0x10876}, {0x10880, 0x1089e}, {0x108e0, 0x108f2},
  {0x108f4, 0x108f5}, {0x10900, 0x10915}, {0x10920, 0x10939}, {0x10980, 0x109b7}, {0x109be, 0x109bf}, {0x10a00, 0x10a00}, {0x10a10, 0x10a13}, {0x10a15, 0x10a17},
  {0x10a19, 0x10a35}, {0x10a60, 0x10a7c}, {0x10a80, 0x10a9c}, {0x10ac0, 0x10ac7}, {0x10ac9, 0x10ae4}, {0x10b00, 0x10b35}, {0x10b40, 0x10b55}, {0x10b60, 0x10b72},
  {0x10b80, 0x10b91}, {0x10c00, 0x10c48}, {0x10c80, 0x10cb2}, {0x10cc0, 0x10cf2}, {0x10d00, 0x10d23}, {0x10e80, 0x10ea9}, {0x10eb0, 0x10eb1}, {0x10f00, 0x10f1c},
  {0x10f27, 0x10f27}, {0x10f30, 0x10f45}, {0x10f70, 0x10f81}, {0x10fb0, 0x10fc4}, {0x10fe0, 0x10ff6}, {0x11003, 0x11037}, {0x11071, 0x11072}, {0x11075, 0x11075},
  {0x11083, 0x110af}, {0x110d0, 0x110e8}, {0x11103, 0x11126}, {0x11144, 0x11144}, {0x11147, 0x11147}, {0x11150, 0x11172}, {0x11176, 0x11176}, {0x11183, 0x111b2},
  {0x111c1, 0x111c4}, {0x111da, 0x111da}, {0x111dc, 0x111dc}, {0x11200, 0x11211}, {0x11213, 0x1122b}, {0x1123f, 0x11240}, {0x11280, 0x11286}, {0x11288, 0x11288},
  {0x1128a, 0x1128d}, {0x1128f, 0x1129d}, {0x1129f, 0x112a8}, {0x112b0, 0x112de}, {0x11305, 0x1130c}, {0x1130f, 0x11310}, {0x11313, 0x11328}, {0x1132a, 0x11330},
  {0x11332, 0x11333}, {0x11335, 0x11339}, {0x1133d, 0x1133d}, {0x11350, 0x11350}, {0x1135d, 0x11361}, {0x11400, 0x11434}, {0x11447, 0x1144a}, {0x1145f, 0x11461},
  {0x11480, 0x114af}, {0x114c4, 0x114c5}, {0x114c7, 0x114c7}, {0x11580, 0x115ae}, {0x115d8, 0x115db}, {0x11600, 0x1162f}, {0x11644, 0x11644}, {0x11680, 0x116aa},
  {0x116b8, 0x116b8}, {0x11700, 0x1171a}, {0x11740, 0x11746}, {0x11800, 0x1182b}, {0x118a0, 0x118df}, {0x118ff, 0x11906}, {0x11909, 0x11909}, {0x1190c, 0x11913},
  {0x11915, 0x11916}, {0x11918, 0x1192f}, {0x1193f, 0x1193f}, {0x11941, 0x11941}, {0x119a0, 0x119a7}, {0x119aa, 0x119d0}, {0x119e1, 0x119e1}, {0x119e3, 0x119e3},
  {0x11a00, 0x11a00}, {0x11a0b, 0x11a32}, {0x11a3a, 0x11a3a}, {0x11a50, 0x11a50}, {0x11a5c, 0x11a89}, {0x11a9d, 0x11a9d}, {0x11ab0, 0x11af8}, {0x11c00, 0x11c08},
  {0x11c0a, 0x11c2e}, {0x11c40, 0x11c40}, {0x11c72, 0x11c8f}, {0x11d00, 0x11d06}, {0x11d08, 0x11d09}, {0x11d0b, 0x11d30}, {0x11d46, 0x11d46}, {0x11d60, 0x11d65},
  {0x11d67, 0x11d68}, {0x11d6a, 0x11d89}, {0x11d98, 0x11d98}, {0x11ee0, 0x11ef2}, {0x11f02, 0x11f02}, {0x11f04, 0x11f10}, {0x11f12, 0x11f33}, {0x11fb0, 0x11fb0},
  {0x12000, 0x12399}, {0x12400, 0x1246e}, {0x12480, 0x12543}, {0x12f90, 0x12ff0}, {0x13000, 0x1342f}, {0x13441, 0x13446}, {0x14400, 0x14646}, {0x16800, 0x16a38},
  {0x16a40, 0x16a5e}, {0x16a70, 0x16abe}, {0x16ad0, 0x16aed}, {0x16b00, 0x16b2f}, {0x16b40, 0x16b43}, {0x16b63, 0x16b77}, {0x16b7d, 0x16b8f}, {0x16e40, 0x16e7f},
  {0x16f00, 0x16f4a}, {0x16f50, 0x16f50}, {0x16f93, 0x16f9f}, {0x16fe0, 0x16fe1}, {0x16fe3, 0x16fe3}, {0x17000, 0x187f7}, {0x18800, 0x18cd5}, {0x18d00, 0x18d08},
  {0x1aff0, 0x1aff3}, {0x1aff5, 0x1affb}, {0x1affd, 0x1affe}, {0x1b000, 0x1b122}, {0x1b132, 0x1b132}, {0x1b150, 0x1b152}, {0x1b155, 0x1b155}, {0x1b164, 0x1b167},
  {0x1b170, 0x1b2fb}, {0x1bc00, 0x1bc6a}, {0x1bc70, 0x1bc7c}, {0x1bc80, 0x1bc88}, {0x1bc90, 0x1bc99}, {0x1d400, 0x1d454}, {0x1d456, 0x1d49c}, {0x1d49e, 0x1d49f},
  {0x1d4a2, 0x1d4a2}, {0x1d4a5, 0x1d4a6}, {0x1d4a9, 0x1d4ac}, {0x1d4ae, 0x1d4b9}, {0x1d4bb, 0x1d4bb}, {0x1d4bd, 0x1d4c3}, {0x1d4c5, 0x1d505}, {0x1d507, 0x1d50a},
  {0x1d50d, 0x1d514}, {0x1d516, 0x1d51c}, {0x1d51e, 0x1d539}, {0x1d53b, 0x1d53e}, {0x1d540, 0x1d544}, {0x1d546, 0x1d546}, {0x1d54a, 0x1d550}, {0x1d552, 0x1d6a5},
  {0x1d6a8, 0x1d6c0}, {0x1d6c2, 0x1d6da}, {0x1d6dc, 0x1d6fa}, {0x1d6fc, 0x1d714}, {0x1d716, 0x1d734}, {0x1d736, 0x1d74e}, {0x1d750, 0x1d76e}, {0x1d770, 0x1d788},
  {0x1d78a, 0x1d7a8}, {0x1d7aa, 0x1d7c2}, {0x1d7c4, 0x1d7cb}, {0x1df00, 0x1df1e}, {0x1df25, 0x1df2a}, {0x1e030, 0x1e06d}, {0x1e100, 0x1e12c}, {0x1e137, 0x1e13d},
  {0x1e14e, 0x1e14e}, {0x1e290, 0x1e2ad}, {0x1e2c0, 0x1e2eb}, {0x1e4d0, 0x1e4eb}, {0x1e7e0, 0x1e7e6}, {0x1e7e8, 0x1e7eb}, {0x1e7ed, 0x1e7ee}, {0x1e7f0, 0x1e7fe},
  {0x1e800, 0x1e8c4}, {0x1e900, 0x1e943}, {0x1e94b, 0x1e94b}, {0x1ee00, 0x1ee03}, {0x1ee05, 0x1ee1f}, {0x1ee21, 0x1ee22}, {0x1ee24, 0x1ee24}, {0x1ee27, 0x1ee27},
  {0x1ee29, 0x1ee32}, {0x1ee34, 0x1ee37}, {0x1ee39, 0x1ee39}, {0x1ee3b, 0x1ee3b}, {0x1ee42, 0x1ee42}, {0x1ee47, 0x1ee47}, {0x1ee49, 0x1ee49}, {0x1ee4b, 0x1ee4b},
  {0x1ee4d, 0x1ee4f}, {0x1ee51, 0x1ee52}, {0x1ee54, 0x1ee54}, {0x1ee57, 0x1ee57}, {0x1ee59, 0x1ee59}, {0x1ee5b, 0x1ee5b}, {0x1ee5d, 0x1ee5d}, {0x1ee5f, 0x1ee5f},
  {0x1ee61, 0x1ee62}, {0x1ee64, 0x1ee64}, {0x1ee67, 0x1ee6a}, {0x1ee6c, 0x1ee72}, {0x1ee74, 0x1ee77}, {0x1ee79, 0x1ee7c}, {0x1ee7e, 0x1ee7e}, {0x1ee80, 0x1ee89},
  {0x1ee8b, 0x1ee9b}, {0x1eea1, 0x1eea3}, {0x1eea5, 0x1eea9}, {0x1eeab, 0x1eebb}, {0x20000, 0x2a6df}, {0x2a700, 0x2b739}, {0x2b740, 0x2b81d}, {0x2b820, 0x2cea1},
  {0x2ceb0, 0x2ebe0}, {0x2ebf0, 0x2ee5d}, {0x2f800, 0x2fa1d}, {0x30000, 0x3134a}, {0x31350, 0x323af},
};

static TSCharacterRange sym_identifier_character_set_2[] = {
  {'$', '$'}, {'0', '9'}, {'A', 'Z'}, {'_', '_'}, {'a', 'z'}, {0xa2, 0xa2}, {0xaa, 0xaa}, {0xb5, 0xb5},
  {0xb7, 0xb7}, {0xba, 0xba}, {0xc0, 0xd6}, {0xd8, 0xf6}, {0xf8, 0x2c1}, {0x2c6, 0x2d1}, {0x2e0, 0x2e4}, {0x2ec, 0x2ec},
  {0x2ee, 0x2ee}, {0x300, 0x374}, {0x376, 0x377}, {0x37b, 0x37d}, {0x37f, 0x37f}, {0x386, 0x38a}, {0x38c, 0x38c}, {0x38e, 0x3a1},
  {0x3a3, 0x3f5}, {0x3f7, 0x481}, {0x483, 0x487}, {0x48a, 0x52f}, {0x531, 0x556}, {0x559, 0x559}, {0x560, 0x588}, {0x591, 0x5bd},
  {0x5bf, 0x5bf}, {0x5c1, 0x5c2}, {0x5c4, 0x5c5}, {0x5c7, 0x5c7}, {0x5d0, 0x5ea}, {0x5ef, 0x5f2}, {0x610, 0x61a}, {0x620, 0x669},
  {0x66e, 0x6d3}, {0x6d5, 0x6dc}, {0x6df, 0x6e8}, {0x6ea, 0x6fc}, {0x6ff, 0x6ff}, {0x710, 0x74a}, {0x74d, 0x7b1}, {0x7c0, 0x7f5},
  {0x7fa, 0x7fa}, {0x7fd, 0x7fd}, {0x800, 0x82d}, {0x840, 0x85b}, {0x860, 0x86a}, {0x870, 0x887}, {0x889, 0x88e}, {0x898, 0x8e1},
  {0x8e3, 0x963}, {0x966, 0x96f}, {0x971, 0x983}, {0x985, 0x98c}, {0x98f, 0x990}, {0x993, 0x9a8}, {0x9aa, 0x9b0}, {0x9b2, 0x9b2},
  {0x9b6, 0x9b9}, {0x9bc, 0x9c4}, {0x9c7, 0x9c8}, {0x9cb, 0x9ce}, {0x9d7, 0x9d7}, {0x9dc, 0x9dd}, {0x9df, 0x9e3}, {0x9e6, 0x9f1},
  {0x9fc, 0x9fc}, {0x9fe, 0x9fe}, {0xa01, 0xa03}, {0xa05, 0xa0a}, {0xa0f, 0xa10}, {0xa13, 0xa28}, {0xa2a, 0xa30}, {0xa32, 0xa33},
  {0xa35, 0xa36}, {0xa38, 0xa39}, {0xa3c, 0xa3c}, {0xa3e, 0xa42}, {0xa47, 0xa48}, {0xa4b, 0xa4d}, {0xa51, 0xa51}, {0xa59, 0xa5c},
  {0xa5e, 0xa5e}, {0xa66, 0xa75}, {0xa81, 0xa83}, {0xa85, 0xa8d}, {0xa8f, 0xa91}, {0xa93, 0xaa8}, {0xaaa, 0xab0}, {0xab2, 0xab3},
  {0xab5, 0xab9}, {0xabc, 0xac5}, {0xac7, 0xac9}, {0xacb, 0xacd}, {0xad0, 0xad0}, {0xae0, 0xae3}, {0xae6, 0xaef}, {0xaf9, 0xaff},
  {0xb01, 0xb03}, {0xb05, 0xb0c}, {0xb0f, 0xb10}, {0xb13, 0xb28}, {0xb2a, 0xb30}, {0xb32, 0xb33}, {0xb35, 0xb39}, {0xb3c, 0xb44},
  {0xb47, 0xb48}, {0xb4b, 0xb4d}, {0xb55, 0xb57}, {0xb5c, 0xb5d}, {0xb5f, 0xb63}, {0xb66, 0xb6f}, {0xb71, 0xb71}, {0xb82, 0xb83},
  {0xb85, 0xb8a}, {0xb8e, 0xb90}, {0xb92, 0xb95}, {0xb99, 0xb9a}, {0xb9c, 0xb9c}, {0xb9e, 0xb9f}, {0xba3, 0xba4}, {0xba8, 0xbaa},
  {0xbae, 0xbb9}, {0xbbe, 0xbc2}, {0xbc6, 0xbc8}, {0xbca, 0xbcd}, {0xbd0, 0xbd0}, {0xbd7, 0xbd7}, {0xbe6, 0xbef}, {0xc00, 0xc0c},
  {0xc0e, 0xc10}, {0xc12, 0xc28}, {0xc2a, 0xc39}, {0xc3c, 0xc44}, {0xc46, 0xc48}, {0xc4a, 0xc4d}, {0xc55, 0xc56}, {0xc58, 0xc5a},
  {0xc5d, 0xc5d}, {0xc60, 0xc63}, {0xc66, 0xc6f}, {0xc80, 0xc83}, {0xc85, 0xc8c}, {0xc8e, 0xc90}, {0xc92, 0xca8}, {0xcaa, 0xcb3},
  {0xcb5, 0xcb9}, {0xcbc, 0xcc4}, {0xcc6, 0xcc8}, {0xcca, 0xccd}, {0xcd5, 0xcd6}, {0xcdd, 0xcde}, {0xce0, 0xce3}, {0xce6, 0xcef},
  {0xcf1, 0xcf3}, {0xd00, 0xd0c}, {0xd0e, 0xd10}, {0xd12, 0xd44}, {0xd46, 0xd48}, {0xd4a, 0xd4e}, {0xd54, 0xd57}, {0xd5f, 0xd63},
  {0xd66, 0xd6f}, {0xd7a, 0xd7f}, {0xd81, 0xd83}, {0xd85, 0xd96}, {0xd9a, 0xdb1}, {0xdb3, 0xdbb}, {0xdbd, 0xdbd}, {0xdc0, 0xdc6},
  {0xdca, 0xdca}, {0xdcf, 0xdd4}, {0xdd6, 0xdd6}, {0xdd8, 0xddf}, {0xde6, 0xdef}, {0xdf2, 0xdf3}, {0xe01, 0xe3a}, {0xe40, 0xe4e},
  {0xe50, 0xe59}, {0xe81, 0xe82}, {0xe84, 0xe84}, {0xe86, 0xe8a}, {0xe8c, 0xea3}, {0xea5, 0xea5}, {0xea7, 0xebd}, {0xec0, 0xec4},
  {0xec6, 0xec6}, {0xec8, 0xece}, {0xed0, 0xed9}, {0xedc, 0xedf}, {0xf00, 0xf00}, {0xf18, 0xf19}, {0xf20, 0xf29}, {0xf35, 0xf35},
  {0xf37, 0xf37}, {0xf39, 0xf39}, {0xf3e, 0xf47}, {0xf49, 0xf6c}, {0xf71, 0xf84}, {0xf86, 0xf97}, {0xf99, 0xfbc}, {0xfc6, 0xfc6},
  {0x1000, 0x1049}, {0x1050, 0x109d}, {0x10a0, 0x10c5}, {0x10c7, 0x10c7}, {0x10cd, 0x10cd}, {0x10d0, 0x10fa}, {0x10fc, 0x1248}, {0x124a, 0x124d},
  {0x1250, 0x1256}, {0x1258, 0x1258}, {0x125a, 0x125d}, {0x1260, 0x1288}, {0x128a, 0x128d}, {0x1290, 0x12b0}, {0x12b2, 0x12b5}, {0x12b8, 0x12be},
  {0x12c0, 0x12c0}, {0x12c2, 0x12c5}, {0x12c8, 0x12d6}, {0x12d8, 0x1310}, {0x1312, 0x1315}, {0x1318, 0x135a}, {0x135d, 0x135f}, {0x1369, 0x1371},
  {0x1380, 0x138f}, {0x13a0, 0x13f5}, {0x13f8, 0x13fd}, {0x1401, 0x166c}, {0x166f, 0x167f}, {0x1681, 0x169a}, {0x16a0, 0x16ea}, {0x16ee, 0x16f8},
  {0x1700, 0x1715}, {0x171f, 0x1734}, {0x1740, 0x1753}, {0x1760, 0x176c}, {0x176e, 0x1770}, {0x1772, 0x1773}, {0x1780, 0x17d3}, {0x17d7, 0x17d7},
  {0x17dc, 0x17dd}, {0x17e0, 0x17e9}, {0x180b, 0x180d}, {0x180f, 0x1819}, {0x1820, 0x1878}, {0x1880, 0x18aa}, {0x18b0, 0x18f5}, {0x1900, 0x191e},
  {0x1920, 0x192b}, {0x1930, 0x193b}, {0x1946, 0x196d}, {0x1970, 0x1974}, {0x1980, 0x19ab}, {0x19b0, 0x19c9}, {0x19d0, 0x19da}, {0x1a00, 0x1a1b},
  {0x1a20, 0x1a5e}, {0x1a60, 0x1a7c}, {0x1a7f, 0x1a89}, {0x1a90, 0x1a99}, {0x1aa7, 0x1aa7}, {0x1ab0, 0x1abd}, {0x1abf, 0x1ace}, {0x1b00, 0x1b4c},
  {0x1b50, 0x1b59}, {0x1b6b, 0x1b73}, {0x1b80, 0x1bf3}, {0x1c00, 0x1c37}, {0x1c40, 0x1c49}, {0x1c4d, 0x1c7d}, {0x1c80, 0x1c88}, {0x1c90, 0x1cba},
  {0x1cbd, 0x1cbf}, {0x1cd0, 0x1cd2}, {0x1cd4, 0x1cfa}, {0x1d00, 0x1f15}, {0x1f18, 0x1f1d}, {0x1f20, 0x1f45}, {0x1f48, 0x1f4d}, {0x1f50, 0x1f57},
  {0x1f59, 0x1f59}, {0x1f5b, 0x1f5b}, {0x1f5d, 0x1f5d}, {0x1f5f, 0x1f7d}, {0x1f80, 0x1fb4}, {0x1fb6, 0x1fbc}, {0x1fbe, 0x1fbe}, {0x1fc2, 0x1fc4},
  {0x1fc6, 0x1fcc}, {0x1fd0, 0x1fd3}, {0x1fd6, 0x1fdb}, {0x1fe0, 0x1fec}, {0x1ff2, 0x1ff4}, {0x1ff6, 0x1ffc}, {0x200c, 0x200d}, {0x203f, 0x2040},
  {0x2054, 0x2054}, {0x2071, 0x2071}, {0x207f, 0x207f}, {0x2090, 0x209c}, {0x20d0, 0x20dc}, {0x20e1, 0x20e1}, {0x20e5, 0x20f0}, {0x2102, 0x2102},
  {0x2107, 0x2107}, {0x210a, 0x2113}, {0x2115, 0x2115}, {0x2118, 0x211d}, {0x2124, 0x2124}, {0x2126, 0x2126}, {0x2128, 0x2128}, {0x212a, 0x2139},
  {0x213c, 0x213f}, {0x2145, 0x2149}, {0x214e, 0x214e}, {0x2160, 0x2188}, {0x2c00, 0x2ce4}, {0x2ceb, 0x2cf3}, {0x2d00, 0x2d25}, {0x2d27, 0x2d27},
  {0x2d2d, 0x2d2d}, {0x2d30, 0x2d67}, {0x2d6f, 0x2d6f}, {0x2d7f, 0x2d96}, {0x2da0, 0x2da6}, {0x2da8, 0x2dae}, {0x2db0, 0x2db6}, {0x2db8, 0x2dbe},
  {0x2dc0, 0x2dc6}, {0x2dc8, 0x2dce}, {0x2dd0, 0x2dd6}, {0x2dd8, 0x2dde}, {0x2de0, 0x2dff}, {0x3005, 0x3007}, {0x3021, 0x302f}, {0x3031, 0x3035},
  {0x3038, 0x303c}, {0x3041, 0x3096}, {0x3099, 0x309a}, {0x309d, 0x309f}, {0x30a1, 0x30ff}, {0x3105, 0x312f}, {0x3131, 0x318e}, {0x31a0, 0x31bf},
  {0x31f0, 0x31ff}, {0x3400, 0x4dbf}, {0x4e00, 0xa48c}, {0xa4d0, 0xa4fd}, {0xa500, 0xa60c}, {0xa610, 0xa62b}, {0xa640, 0xa66f}, {0xa674, 0xa67d},
  {0xa67f, 0xa6f1}, {0xa717, 0xa71f}, {0xa722, 0xa788}, {0xa78b, 0xa7ca}, {0xa7d0, 0xa7d1}, {0xa7d3, 0xa7d3}, {0xa7d5, 0xa7d9}, {0xa7f2, 0xa827},
  {0xa82c, 0xa82c}, {0xa840, 0xa873}, {0xa880, 0xa8c5}, {0xa8d0, 0xa8d9}, {0xa8e0, 0xa8f7}, {0xa8fb, 0xa8fb}, {0xa8fd, 0xa92d}, {0xa930, 0xa953},
  {0xa960, 0xa97c}, {0xa980, 0xa9c0}, {0xa9cf, 0xa9d9}, {0xa9e0, 0xa9fe}, {0xaa00, 0xaa36}, {0xaa40, 0xaa4d}, {0xaa50, 0xaa59}, {0xaa60, 0xaa76},
  {0xaa7a, 0xaac2}, {0xaadb, 0xaadd}, {0xaae0, 0xaaef}, {0xaaf2, 0xaaf6}, {0xab01, 0xab06}, {0xab09, 0xab0e}, {0xab11, 0xab16}, {0xab20, 0xab26},
  {0xab28, 0xab2e}, {0xab30, 0xab5a}, {0xab5c, 0xab69}, {0xab70, 0xabea}, {0xabec, 0xabed}, {0xabf0, 0xabf9}, {0xac00, 0xd7a3}, {0xd7b0, 0xd7c6},
  {0xd7cb, 0xd7fb}, {0xf900, 0xfa6d}, {0xfa70, 0xfad9}, {0xfb00, 0xfb06}, {0xfb13, 0xfb17}, {0xfb1d, 0xfb28}, {0xfb2a, 0xfb36}, {0xfb38, 0xfb3c},
  {0xfb3e, 0xfb3e}, {0xfb40, 0xfb41}, {0xfb43, 0xfb44}, {0xfb46, 0xfbb1}, {0xfbd3, 0xfc5d}, {0xfc64, 0xfd3d}, {0xfd50, 0xfd8f}, {0xfd92, 0xfdc7},
  {0xfdf0, 0xfdf9}, {0xfe00, 0xfe0f}, {0xfe20, 0xfe2f}, {0xfe33, 0xfe34}, {0xfe4d, 0xfe4f}, {0xfe71, 0xfe71}, {0xfe73, 0xfe73}, {0xfe77, 0xfe77},
  {0xfe79, 0xfe79}, {0xfe7b, 0xfe7b}, {0xfe7d, 0xfe7d}, {0xfe7f, 0xfefc}, {0xff10, 0xff19}, {0xff21, 0xff3a}, {0xff3f, 0xff3f}, {0xff41, 0xff5a},
  {0xff65, 0xffbe}, {0xffc2, 0xffc7}, {0xffca, 0xffcf}, {0xffd2, 0xffd7}, {0xffda, 0xffdc}, {0x10000, 0x1000b}, {0x1000d, 0x10026}, {0x10028, 0x1003a},
  {0x1003c, 0x1003d}, {0x1003f, 0x1004d}, {0x10050, 0x1005d}, {0x10080, 0x100fa}, {0x10140, 0x10174}, {0x101fd, 0x101fd}, {0x10280, 0x1029c}, {0x102a0, 0x102d0},
  {0x102e0, 0x102e0}, {0x10300, 0x1031f}, {0x1032d, 0x1034a}, {0x10350, 0x1037a}, {0x10380, 0x1039d}, {0x103a0, 0x103c3}, {0x103c8, 0x103cf}, {0x103d1, 0x103d5},
  {0x10400, 0x1049d}, {0x104a0, 0x104a9}, {0x104b0, 0x104d3}, {0x104d8, 0x104fb}, {0x10500, 0x10527}, {0x10530, 0x10563}, {0x10570, 0x1057a}, {0x1057c, 0x1058a},
  {0x1058c, 0x10592}, {0x10594, 0x10595}, {0x10597, 0x105a1}, {0x105a3, 0x105b1}, {0x105b3, 0x105b9}, {0x105bb, 0x105bc}, {0x10600, 0x10736}, {0x10740, 0x10755},
  {0x10760, 0x10767}, {0x10780, 0x10785}, {0x10787, 0x107b0}, {0x107b2, 0x107ba}, {0x10800, 0x10805}, {0x10808, 0x10808}, {0x1080a, 0x10835}, {0x10837, 0x10838},
  {0x1083c, 0x1083c}, {0x1083f, 0x10855}, {0x10860, 0x10876}, {0x10880, 0x1089e}, {0x108e0, 0x108f2}, {0x108f4, 0x108f5}, {0x10900, 0x10915}, {0x10920, 0x10939},
  {0x10980, 0x109b7}, {0x109be, 0x109bf}, {0x10a00, 0x10a03}, {0x10a05, 0x10a06}, {0x10a0c, 0x10a13}, {0x10a15, 0x10a17}, {0x10a19, 0x10a35}, {0x10a38, 0x10a3a},
  {0x10a3f, 0x10a3f}, {0x10a60, 0x10a7c}, {0x10a80, 0x10a9c}, {0x10ac0, 0x10ac7}, {0x10ac9, 0x10ae6}, {0x10b00, 0x10b35}, {0x10b40, 0x10b55}, {0x10b60, 0x10b72},
  {0x10b80, 0x10b91}, {0x10c00, 0x10c48}, {0x10c80, 0x10cb2}, {0x10cc0, 0x10cf2}, {0x10d00, 0x10d27}, {0x10d30, 0x10d39}, {0x10e80, 0x10ea9}, {0x10eab, 0x10eac},
  {0x10eb0, 0x10eb1}, {0x10efd, 0x10f1c}, {0x10f27, 0x10f27}, {0x10f30, 0x10f50}, {0x10f70, 0x10f85}, {0x10fb0, 0x10fc4}, {0x10fe0, 0x10ff6}, {0x11000, 0x11046},
  {0x11066, 0x11075}, {0x1107f, 0x110ba}, {0x110c2, 0x110c2}, {0x110d0, 0x110e8}, {0x110f0, 0x110f9}, {0x11100, 0x11134}, {0x11136, 0x1113f}, {0x11144, 0x11147},
  {0x11150, 0x11173}, {0x11176, 0x11176}, {0x11180, 0x111c4}, {0x111c9, 0x111cc}, {0x111ce, 0x111da}, {0x111dc, 0x111dc}, {0x11200, 0x11211}, {0x11213, 0x11237},
  {0x1123e, 0x11241}, {0x11280, 0x11286}, {0x11288, 0x11288}, {0x1128a, 0x1128d}, {0x1128f, 0x1129d}, {0x1129f, 0x112a8}, {0x112b0, 0x112ea}, {0x112f0, 0x112f9},
  {0x11300, 0x11303}, {0x11305, 0x1130c}, {0x1130f, 0x11310}, {0x11313, 0x11328}, {0x1132a, 0x11330}, {0x11332, 0x11333}, {0x11335, 0x11339}, {0x1133b, 0x11344},
  {0x11347, 0x11348}, {0x1134b, 0x1134d}, {0x11350, 0x11350}, {0x11357, 0x11357}, {0x1135d, 0x11363}, {0x11366, 0x1136c}, {0x11370, 0x11374}, {0x11400, 0x1144a},
  {0x11450, 0x11459}, {0x1145e, 0x11461}, {0x11480, 0x114c5}, {0x114c7, 0x114c7}, {0x114d0, 0x114d9}, {0x11580, 0x115b5}, {0x115b8, 0x115c0}, {0x115d8, 0x115dd},
  {0x11600, 0x11640}, {0x11644, 0x11644}, {0x11650, 0x11659}, {0x11680, 0x116b8}, {0x116c0, 0x116c9}, {0x11700, 0x1171a}, {0x1171d, 0x1172b}, {0x11730, 0x11739},
  {0x11740, 0x11746}, {0x11800, 0x1183a}, {0x118a0, 0x118e9}, {0x118ff, 0x11906}, {0x11909, 0x11909}, {0x1190c, 0x11913}, {0x11915, 0x11916}, {0x11918, 0x11935},
  {0x11937, 0x11938}, {0x1193b, 0x11943}, {0x11950, 0x11959}, {0x119a0, 0x119a7}, {0x119aa, 0x119d7}, {0x119da, 0x119e1}, {0x119e3, 0x119e4}, {0x11a00, 0x11a3e},
  {0x11a47, 0x11a47}, {0x11a50, 0x11a99}, {0x11a9d, 0x11a9d}, {0x11ab0, 0x11af8}, {0x11c00, 0x11c08}, {0x11c0a, 0x11c36}, {0x11c38, 0x11c40}, {0x11c50, 0x11c59},
  {0x11c72, 0x11c8f}, {0x11c92, 0x11ca7}, {0x11ca9, 0x11cb6}, {0x11d00, 0x11d06}, {0x11d08, 0x11d09}, {0x11d0b, 0x11d36}, {0x11d3a, 0x11d3a}, {0x11d3c, 0x11d3d},
  {0x11d3f, 0x11d47}, {0x11d50, 0x11d59}, {0x11d60, 0x11d65}, {0x11d67, 0x11d68}, {0x11d6a, 0x11d8e}, {0x11d90, 0x11d91}, {0x11d93, 0x11d98}, {0x11da0, 0x11da9},
  {0x11ee0, 0x11ef6}, {0x11f00, 0x11f10}, {0x11f12, 0x11f3a}, {0x11f3e, 0x11f42}, {0x11f50, 0x11f59}, {0x11fb0, 0x11fb0}, {0x12000, 0x12399}, {0x12400, 0x1246e},
  {0x12480, 0x12543}, {0x12f90, 0x12ff0}, {0x13000, 0x1342f}, {0x13440, 0x13455}, {0x14400, 0x14646}, {0x16800, 0x16a38}, {0x16a40, 0x16a5e}, {0x16a60, 0x16a69},
  {0x16a70, 0x16abe}, {0x16ac0, 0x16ac9}, {0x16ad0, 0x16aed}, {0x16af0, 0x16af4}, {0x16b00, 0x16b36}, {0x16b40, 0x16b43}, {0x16b50, 0x16b59}, {0x16b63, 0x16b77},
  {0x16b7d, 0x16b8f}, {0x16e40, 0x16e7f}, {0x16f00, 0x16f4a}, {0x16f4f, 0x16f87}, {0x16f8f, 0x16f9f}, {0x16fe0, 0x16fe1}, {0x16fe3, 0x16fe4}, {0x16ff0, 0x16ff1},
  {0x17000, 0x187f7}, {0x18800, 0x18cd5}, {0x18d00, 0x18d08}, {0x1aff0, 0x1aff3}, {0x1aff5, 0x1affb}, {0x1affd, 0x1affe}, {0x1b000, 0x1b122}, {0x1b132, 0x1b132},
  {0x1b150, 0x1b152}, {0x1b155, 0x1b155}, {0x1b164, 0x1b167}, {0x1b170, 0x1b2fb}, {0x1bc00, 0x1bc6a}, {0x1bc70, 0x1bc7c}, {0x1bc80, 0x1bc88}, {0x1bc90, 0x1bc99},
  {0x1bc9d, 0x1bc9e}, {0x1cf00, 0x1cf2d}, {0x1cf30, 0x1cf46}, {0x1d165, 0x1d169}, {0x1d16d, 0x1d172}, {0x1d17b, 0x1d182}, {0x1d185, 0x1d18b}, {0x1d1aa, 0x1d1ad},
  {0x1d242, 0x1d244}, {0x1d400, 0x1d454}, {0x1d456, 0x1d49c}, {0x1d49e, 0x1d49f}, {0x1d4a2, 0x1d4a2}, {0x1d4a5, 0x1d4a6}, {0x1d4a9, 0x1d4ac}, {0x1d4ae, 0x1d4b9},
  {0x1d4bb, 0x1d4bb}, {0x1d4bd, 0x1d4c3}, {0x1d4c5, 0x1d505}, {0x1d507, 0x1d50a}, {0x1d50d, 0x1d514}, {0x1d516, 0x1d51c}, {0x1d51e, 0x1d539}, {0x1d53b, 0x1d53e},
  {0x1d540, 0x1d544}, {0x1d546, 0x1d546}, {0x1d54a, 0x1d550}, {0x1d552, 0x1d6a5}, {0x1d6a8, 0x1d6c0}, {0x1d6c2, 0x1d6da}, {0x1d6dc, 0x1d6fa}, {0x1d6fc, 0x1d714},
  {0x1d716, 0x1d734}, {0x1d736, 0x1d74e}, {0x1d750, 0x1d76e}, {0x1d770, 0x1d788}, {0x1d78a, 0x1d7a8}, {0x1d7aa, 0x1d7c2}, {0x1d7c4, 0x1d7cb}, {0x1d7ce, 0x1d7ff},
  {0x1da00, 0x1da36}, {0x1da3b, 0x1da6c}, {0x1da75, 0x1da75}, {0x1da84, 0x1da84}, {0x1da9b, 0x1da9f}, {0x1daa1, 0x1daaf}, {0x1df00, 0x1df1e}, {0x1df25, 0x1df2a},
  {0x1e000, 0x1e006}, {0x1e008, 0x1e018}, {0x1e01b, 0x1e021}, {0x1e023, 0x1e024}, {0x1e026, 0x1e02a}, {0x1e030, 0x1e06d}, {0x1e08f, 0x1e08f}, {0x1e100, 0x1e12c},
  {0x1e130, 0x1e13d}, {0x1e140, 0x1e149}, {0x1e14e, 0x1e14e}, {0x1e290, 0x1e2ae}, {0x1e2c0, 0x1e2f9}, {0x1e4d0, 0x1e4f9}, {0x1e7e0, 0x1e7e6}, {0x1e7e8, 0x1e7eb},
  {0x1e7ed, 0x1e7ee}, {0x1e7f0, 0x1e7fe}, {0x1e800, 0x1e8c4}, {0x1e8d0, 0x1e8d6}, {0x1e900, 0x1e94b}, {0x1e950, 0x1e959}, {0x1ee00, 0x1ee03}, {0x1ee05, 0x1ee1f},
  {0x1ee21, 0x1ee22}, {0x1ee24, 0x1ee24}, {0x1ee27, 0x1ee27}, {0x1ee29, 0x1ee32}, {0x1ee34, 0x1ee37}, {0x1ee39, 0x1ee39}, {0x1ee3b, 0x1ee3b}, {0x1ee42, 0x1ee42},
  {0x1ee47, 0x1ee47}, {0x1ee49, 0x1ee49}, {0x1ee4b, 0x1ee4b}, {0x1ee4d, 0x1ee4f}, {0x1ee51, 0x1ee52}, {0x1ee54, 0x1ee54}, {0x1ee57, 0x1ee57}, {0x1ee59, 0x1ee59},
  {0x1ee5b, 0x1ee5b}, {0x1ee5d, 0x1ee5d}, {0x1ee5f, 0x1ee5f}, {0x1ee61, 0x1ee62}, {0x1ee64, 0x1ee64}, {0x1ee67, 0x1ee6a}, {0x1ee6c, 0x1ee72}, {0x1ee74, 0x1ee77},
  {0x1ee79, 0x1ee7c}, {0x1ee7e, 0x1ee7e}, {0x1ee80, 0x1ee89}, {0x1ee8b, 0x1ee9b}, {0x1eea1, 0x1eea3}, {0x1eea5, 0x1eea9}, {0x1eeab, 0x1eebb}, {0x1fbf0, 0x1fbf9},
  {0x20000, 0x2a6df}, {0x2a700, 0x2b739}, {0x2b740, 0x2b81d}, {0x2b820, 0x2cea1}, {0x2ceb0, 0x2ebe0}, {0x2ebf0, 0x2ee5d}, {0x2f800, 0x2fa1d}, {0x30000, 0x3134a},
  {0x31350, 0x323af}, {0xe0100, 0xe01ef},
};

static bool ts_lex(TSLexer *lexer, TSStateId state) {
  START_LEXER();
  eof = lexer->eof(lexer);
  switch (state) {
    case 0:
      if (eof) ADVANCE(65);
      ADVANCE_MAP(
        '!', 170,
        '"', 90,
        '%', 159,
        '&', 119,
        '\'', 23,
        '(', 115,
        ')', 116,
        '*', 150,
        '+', 146,
        ',', 167,
        '-', 147,
        '.', 178,
        '/', 152,
        '0', 67,
        ':', 169,
        ';', 182,
        '<', 137,
        '=', 121,
        '>', 134,
        '?', 168,
        '@', 184,
        '[', 174,
        '\\', 42,
        ']', 175,
        '^', 157,
        'n', 190,
        '{', 181,
        '|', 154,
        '}', 107,
        '~', 171,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(63);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(69);
      if (set_contains(sym_identifier_character_set_1, 669, lookahead)) ADVANCE(191);
      END_STATE();
    case 1:
      ADVANCE_MAP(
        '!', 170,
        '"', 90,
        '%', 159,
        '&', 119,
        '\'', 23,
        '(', 115,
        ')', 116,
        '*', 150,
        '+', 146,
        ',', 167,
        '-', 147,
        '.', 178,
        '/', 152,
        '0', 67,
        ':', 169,
        ';', 182,
        '<', 137,
        '=', 121,
        '>', 134,
        '?', 168,
        '@', 183,
        '[', 174,
        ']', 175,
        '^', 157,
        '{', 181,
        '|', 154,
        '}', 107,
        '~', 171,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(1);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(69);
      if (set_contains(sym_identifier_character_set_1, 669, lookahead)) ADVANCE(191);
      END_STATE();
    case 2:
      ADVANCE_MAP(
        '!', 170,
        '"', 90,
        '%', 159,
        '&', 119,
        '\'', 23,
        '(', 115,
        ')', 116,
        '*', 150,
        '+', 146,
        ',', 167,
        '-', 147,
        '.', 179,
        '/', 152,
        '0', 67,
        ':', 20,
        ';', 182,
        '<', 137,
        '=', 121,
        '>', 134,
        '?', 168,
        '@', 183,
        '[', 174,
        '^', 157,
        'n', 190,
        '|', 154,
        '}', 107,
        '~', 171,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(2);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(69);
      if (set_contains(sym_identifier_character_set_1, 669, lookahead)) ADVANCE(191);
      END_STATE();
    case 3:
      ADVANCE_MAP(
        '!', 21,
        '%', 158,
        '&', 118,
        '(', 115,
        ')', 116,
        '*', 149,
        '+', 145,
        ',', 167,
        '-', 148,
        '.', 176,
        '/', 151,
        ':', 169,
        ';', 182,
        '<', 138,
        '=', 121,
        '>', 135,
        '?', 168,
        '@', 183,
        '[', 174,
        ']', 175,
        '^', 156,
        'n', 190,
        '|', 155,
        '}', 107,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(3);
      if (set_contains(sym_identifier_character_set_1, 669, lookahead)) ADVANCE(191);
      END_STATE();
    case 4:
      ADVANCE_MAP(
        '!', 21,
        '%', 158,
        '&', 118,
        '(', 115,
        ')', 116,
        '*', 149,
        '+', 145,
        ',', 167,
        '-', 148,
        '.', 176,
        '/', 151,
        ':', 169,
        ';', 182,
        '<', 138,
        '=', 22,
        '>', 135,
        '?', 168,
        '@', 183,
        '[', 174,
        ']', 175,
        '^', 156,
        '{', 181,
        '|', 155,
        '}', 107,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(4);
      if (set_contains(sym_identifier_character_set_1, 669, lookahead)) ADVANCE(191);
      END_STATE();
    case 5:
      if (lookahead == '"') ADVANCE(91);
      END_STATE();
    case 6:
      if (lookahead == '"') ADVANCE(104);
      if (lookahead == '/') ADVANCE(99);
      if (lookahead == '\\') ADVANCE(42);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') ADVANCE(102);
      if (lookahead != 0) ADVANCE(103);
      END_STATE();
    case 7:
      if (lookahead == '"') ADVANCE(105);
      END_STATE();
    case 8:
      if (lookahead == '"') ADVANCE(89);
      if (lookahead == '/') ADVANCE(93);
      if (lookahead == '\\') ADVANCE(43);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') ADVANCE(96);
      if (lookahead != 0) ADVANCE(97);
      END_STATE();
    case 9:
      ADVANCE_MAP(
        '&', 117,
        '(', 115,
        ')', 116,
        ',', 167,
        '.', 177,
        '/', 11,
        ':', 169,
        ';', 182,
        '<', 136,
        '=', 120,
        '>', 133,
        '?', 168,
        '@', 183,
        '[', 174,
        '{', 181,
        '|', 153,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(9);
      if (set_contains(sym_identifier_character_set_1, 669, lookahead)) ADVANCE(191);
      END_STATE();
    case 10:
      if (lookahead == '\'') ADVANCE(88);
      if (lookahead == '\\') ADVANCE(62);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(10);
      END_STATE();
    case 11:
      if (lookahead == '*') ADVANCE(13);
      if (lookahead == '/') ADVANCE(192);
      END_STATE();
    case 12:
      if (lookahead == '*') ADVANCE(12);
      if (lookahead == '/') ADVANCE(193);
      if (lookahead != 0) ADVANCE(13);
      END_STATE();
    case 13:
      if (lookahead == '*') ADVANCE(12);
      if (lookahead != 0) ADVANCE(13);
      END_STATE();
    case 14:
      if (lookahead == '.') ADVANCE(81);
      if (lookahead == '_') ADVANCE(52);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(14);
      END_STATE();
    case 15:
      if (lookahead == '.') ADVANCE(187);
      END_STATE();
    case 16:
      if (lookahead == '.') ADVANCE(58);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(72);
      END_STATE();
    case 17:
      if (lookahead == '/') ADVANCE(11);
      if (lookahead == '<') ADVANCE(136);
      if (lookahead == '@') ADVANCE(36);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(17);
      if (set_contains(sym_identifier_character_set_1, 669, lookahead)) ADVANCE(191);
      END_STATE();
    case 18:
      if (lookahead == '0') ADVANCE(87);
      if (lookahead == '+' ||
          lookahead == '-') ADVANCE(19);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(85);
      END_STATE();
    case 19:
      if (lookahead == '0') ADVANCE(87);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(85);
      END_STATE();
    case 20:
      if (lookahead == ':') ADVANCE(180);
      END_STATE();
    case 21:
      if (lookahead == '=') ADVANCE(142);
      END_STATE();
    case 22:
      if (lookahead == '=') ADVANCE(141);
      END_STATE();
    case 23:
      if (lookahead == '\\') ADVANCE(62);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\'') ADVANCE(10);
      END_STATE();
    case 24:
      if (lookahead == '_') ADVANCE(24);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(70);
      END_STATE();
    case 25:
      if (lookahead == '_') ADVANCE(24);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(68);
      END_STATE();
    case 26:
      if (lookahead == '_') ADVANCE(26);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(84);
      END_STATE();
    case 27:
      if (lookahead == 'a') ADVANCE(37);
      END_STATE();
    case 28:
      if (lookahead == 'a') ADVANCE(29);
      END_STATE();
    case 29:
      if (lookahead == 'c') ADVANCE(34);
      END_STATE();
    case 30:
      if (lookahead == 'd') ADVANCE(185);
      END_STATE();
    case 31:
      if (lookahead == 'e') ADVANCE(39);
      END_STATE();
    case 32:
      if (lookahead == 'e') ADVANCE(27);
      END_STATE();
    case 33:
      if (lookahead == 'e') ADVANCE(30);
      END_STATE();
    case 34:
      if (lookahead == 'e') ADVANCE(186);
      END_STATE();
    case 35:
      if (lookahead == 'f') ADVANCE(28);
      END_STATE();
    case 36:
      if (lookahead == 'i') ADVANCE(38);
      END_STATE();
    case 37:
      if (lookahead == 'l') ADVANCE(33);
      END_STATE();
    case 38:
      if (lookahead == 'n') ADVANCE(41);
      END_STATE();
    case 39:
      if (lookahead == 'r') ADVANCE(35);
      END_STATE();
    case 40:
      if (lookahead == 's') ADVANCE(32);
      END_STATE();
    case 41:
      if (lookahead == 't') ADVANCE(31);
      END_STATE();
    case 42:
      if (lookahead == 'u') ADVANCE(109);
      if (lookahead == 'x') ADVANCE(111);
      if (lookahead == '{') ADVANCE(106);
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(110);
      if (lookahead == '"' ||
          lookahead == '\'' ||
          lookahead == '\\' ||
          lookahead == 'b' ||
          lookahead == 'f' ||
          lookahead == 'n' ||
          ('r' <= lookahead && lookahead <= 't')) ADVANCE(112);
      if (lookahead != 0) ADVANCE(108);
      END_STATE();
    case 43:
      if (lookahead == 'u') ADVANCE(45);
      if (lookahead == 'x') ADVANCE(60);
      if (lookahead == '{') ADVANCE(106);
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(114);
      if (lookahead != 0) ADVANCE(112);
      END_STATE();
    case 44:
      if (lookahead == '{') ADVANCE(106);
      END_STATE();
    case 45:
      if (lookahead == '{') ADVANCE(59);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(61);
      END_STATE();
    case 46:
      if (lookahead == '}') ADVANCE(112);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(46);
      END_STATE();
    case 47:
      if (lookahead == '+' ||
          lookahead == '-') ADVANCE(53);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(80);
      END_STATE();
    case 48:
      if (lookahead == '0' ||
          lookahead == '1') ADVANCE(77);
      END_STATE();
    case 49:
      if (lookahead == '8' ||
          lookahead == '9') ADVANCE(14);
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(74);
      END_STATE();
    case 50:
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(75);
      END_STATE();
    case 51:
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(79);
      END_STATE();
    case 52:
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(14);
      END_STATE();
    case 53:
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(80);
      END_STATE();
    case 54:
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(68);
      END_STATE();
    case 55:
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(84);
      END_STATE();
    case 56:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(112);
      END_STATE();
    case 57:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(72);
      END_STATE();
    case 58:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(83);
      END_STATE();
    case 59:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(46);
      END_STATE();
    case 60:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(56);
      END_STATE();
    case 61:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(60);
      END_STATE();
    case 62:
      if (lookahead != 0) ADVANCE(10);
      END_STATE();
    case 63:
      if (eof) ADVANCE(65);
      ADVANCE_MAP(
        '!', 170,
        '"', 90,
        '%', 159,
        '&', 119,
        '\'', 23,
        '(', 115,
        ')', 116,
        '*', 150,
        '+', 146,
        ',', 167,
        '-', 147,
        '.', 178,
        '/', 152,
        '0', 67,
        ':', 169,
        ';', 182,
        '<', 137,
        '=', 121,
        '>', 134,
        '?', 168,
        '@', 184,
        '[', 174,
        '\\', 44,
        ']', 175,
        '^', 157,
        'n', 190,
        '{', 181,
        '|', 154,
        '}', 107,
        '~', 171,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(63);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(69);
      if (set_contains(sym_identifier_character_set_1, 669, lookahead)) ADVANCE(191);
      END_STATE();
    case 64:
      if (eof) ADVANCE(65);
      ADVANCE_MAP(
        '!', 170,
        '"', 90,
        '%', 158,
        '&', 118,
        '\'', 23,
        '(', 115,
        ')', 116,
        '*', 149,
        '+', 145,
        ',', 167,
        '-', 148,
        '.', 179,
        '/', 151,
        '0', 67,
        ':', 169,
        ';', 182,
        '<', 138,
        '=', 22,
        '>', 135,
        '?', 168,
        '@', 184,
        '[', 174,
        ']', 175,
        '^', 156,
        'n', 190,
        '{', 181,
        '|', 155,
        '}', 107,
        '~', 171,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(64);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(69);
      if (set_contains(sym_identifier_character_set_1, 669, lookahead)) ADVANCE(191);
      END_STATE();
    case 65:
      ACCEPT_TOKEN(ts_builtin_sym_end);
      END_STATE();
    case 66:
      ACCEPT_TOKEN(sym_decimal_integer_literal);
      END_STATE();
    case 67:
      ACCEPT_TOKEN(sym_decimal_integer_literal);
      ADVANCE_MAP(
        '.', 81,
        '_', 52,
        'B', 48,
        'b', 48,
        'E', 47,
        'e', 47,
        'L', 66,
        'l', 66,
        'O', 50,
        'o', 50,
        'X', 16,
        'x', 16,
        '8', 14,
        '9', 14,
        'D', 78,
        'F', 78,
        'd', 78,
        'f', 78,
      );
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(74);
      END_STATE();
    case 68:
      ACCEPT_TOKEN(sym_decimal_integer_literal);
      ADVANCE_MAP(
        '.', 81,
        '_', 25,
        'E', 47,
        'e', 47,
        'L', 66,
        'l', 66,
        'D', 78,
        'F', 78,
        'd', 78,
        'f', 78,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(68);
      END_STATE();
    case 69:
      ACCEPT_TOKEN(sym_decimal_integer_literal);
      ADVANCE_MAP(
        '.', 81,
        '_', 54,
        'E', 47,
        'e', 47,
        'L', 66,
        'l', 66,
        'D', 78,
        'F', 78,
        'd', 78,
        'f', 78,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(68);
      END_STATE();
    case 70:
      ACCEPT_TOKEN(sym_decimal_integer_literal);
      ADVANCE_MAP(
        '_', 24,
        'E', 47,
        'e', 47,
        'L', 66,
        'l', 66,
        'D', 78,
        'F', 78,
        'd', 78,
        'f', 78,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(70);
      END_STATE();
    case 71:
      ACCEPT_TOKEN(sym_hex_integer_literal);
      END_STATE();
    case 72:
      ACCEPT_TOKEN(sym_hex_integer_literal);
      if (lookahead == '.') ADVANCE(86);
      if (lookahead == '_') ADVANCE(57);
      if (lookahead == 'L' ||
          lookahead == 'l') ADVANCE(71);
      if (lookahead == 'P' ||
          lookahead == 'p') ADVANCE(18);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(72);
      END_STATE();
    case 73:
      ACCEPT_TOKEN(sym_octal_integer_literal);
      END_STATE();
    case 74:
      ACCEPT_TOKEN(sym_octal_integer_literal);
      if (lookahead == '.') ADVANCE(81);
      if (lookahead == '_') ADVANCE(49);
      if (lookahead == 'L' ||
          lookahead == 'l') ADVANCE(73);
      if (lookahead == '8' ||
          lookahead == '9') ADVANCE(14);
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(74);
      END_STATE();
    case 75:
      ACCEPT_TOKEN(sym_octal_integer_literal);
      if (lookahead == '_') ADVANCE(50);
      if (lookahead == 'L' ||
          lookahead == 'l') ADVANCE(73);
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(75);
      END_STATE();
    case 76:
      ACCEPT_TOKEN(sym_binary_integer_literal);
      END_STATE();
    case 77:
      ACCEPT_TOKEN(sym_binary_integer_literal);
      if (lookahead == '_') ADVANCE(48);
      if (lookahead == 'L' ||
          lookahead == 'l') ADVANCE(76);
      if (lookahead == '0' ||
          lookahead == '1') ADVANCE(77);
      END_STATE();
    case 78:
      ACCEPT_TOKEN(sym_decimal_floating_point_literal);
      END_STATE();
    case 79:
      ACCEPT_TOKEN(sym_decimal_floating_point_literal);
      if (lookahead == '_') ADVANCE(51);
      if (lookahead == 'E' ||
          lookahead == 'e') ADVANCE(47);
      if (('D' <= lookahead && lookahead <= 'F') ||
          ('d' <= lookahead && lookahead <= 'f')) ADVANCE(78);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(79);
      END_STATE();
    case 80:
      ACCEPT_TOKEN(sym_decimal_floating_point_literal);
      if (lookahead == '_') ADVANCE(53);
      if (lookahead == 'D' ||
          lookahead == 'F' ||
          lookahead == 'd' ||
          lookahead == 'f') ADVANCE(78);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(80);
      END_STATE();
    case 81:
      ACCEPT_TOKEN(sym_decimal_floating_point_literal);
      if (lookahead == 'E' ||
          lookahead == 'e') ADVANCE(47);
      if (('D' <= lookahead && lookahead <= 'F') ||
          ('d' <= lookahead && lookahead <= 'f')) ADVANCE(78);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(79);
      END_STATE();
    case 82:
      ACCEPT_TOKEN(sym_hex_floating_point_literal);
      END_STATE();
    case 83:
      ACCEPT_TOKEN(sym_hex_floating_point_literal);
      if (lookahead == '_') ADVANCE(58);
      if (lookahead == 'P' ||
          lookahead == 'p') ADVANCE(18);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(83);
      END_STATE();
    case 84:
      ACCEPT_TOKEN(sym_hex_floating_point_literal);
      if (lookahead == '_') ADVANCE(26);
      if (lookahead == 'D' ||
          lookahead == 'F' ||
          lookahead == 'd' ||
          lookahead == 'f') ADVANCE(82);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(84);
      END_STATE();
    case 85:
      ACCEPT_TOKEN(sym_hex_floating_point_literal);
      if (lookahead == '_') ADVANCE(55);
      if (lookahead == 'D' ||
          lookahead == 'F' ||
          lookahead == 'd' ||
          lookahead == 'f') ADVANCE(82);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(84);
      END_STATE();
    case 86:
      ACCEPT_TOKEN(sym_hex_floating_point_literal);
      if (lookahead == 'P' ||
          lookahead == 'p') ADVANCE(18);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(83);
      END_STATE();
    case 87:
      ACCEPT_TOKEN(sym_hex_floating_point_literal);
      if (lookahead == 'D' ||
          lookahead == 'F' ||
          lookahead == 'd' ||
          lookahead == 'f') ADVANCE(82);
      END_STATE();
    case 88:
      ACCEPT_TOKEN(sym_character_literal);
      END_STATE();
    case 89:
      ACCEPT_TOKEN(anon_sym_DQUOTE);
      END_STATE();
    case 90:
      ACCEPT_TOKEN(anon_sym_DQUOTE);
      if (lookahead == '"') ADVANCE(5);
      END_STATE();
    case 91:
      ACCEPT_TOKEN(anon_sym_DQUOTE_DQUOTE_DQUOTE);
      END_STATE();
    case 92:
      ACCEPT_TOKEN(sym_string_fragment);
      if (lookahead == '\n') ADVANCE(97);
      if (lookahead != 0 &&
          lookahead != '"' &&
          lookahead != '\\') ADVANCE(92);
      END_STATE();
    case 93:
      ACCEPT_TOKEN(sym_string_fragment);
      if (lookahead == '*') ADVANCE(95);
      if (lookahead == '/') ADVANCE(92);
      if (lookahead != 0 &&
          lookahead != '"' &&
          lookahead != '\\') ADVANCE(97);
      END_STATE();
    case 94:
      ACCEPT_TOKEN(sym_string_fragment);
      if (lookahead == '*') ADVANCE(94);
      if (lookahead == '/') ADVANCE(97);
      if (lookahead != 0 &&
          lookahead != '"' &&
          lookahead != '\\') ADVANCE(95);
      END_STATE();
    case 95:
      ACCEPT_TOKEN(sym_string_fragment);
      if (lookahead == '*') ADVANCE(94);
      if (lookahead != 0 &&
          lookahead != '"' &&
          lookahead != '\\') ADVANCE(95);
      END_STATE();
    case 96:
      ACCEPT_TOKEN(sym_string_fragment);
      if (lookahead == '/') ADVANCE(93);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') ADVANCE(96);
      if (lookahead != 0 &&
          lookahead != '"' &&
          lookahead != '\\') ADVANCE(97);
      END_STATE();
    case 97:
      ACCEPT_TOKEN(sym_string_fragment);
      if (lookahead != 0 &&
          lookahead != '"' &&
          lookahead != '\\') ADVANCE(97);
      END_STATE();
    case 98:
      ACCEPT_TOKEN(aux_sym__multiline_string_fragment_token1);
      if (lookahead == '\n') ADVANCE(103);
      if (lookahead == '"' ||
          lookahead == '\\') ADVANCE(192);
      if (lookahead != 0) ADVANCE(98);
      END_STATE();
    case 99:
      ACCEPT_TOKEN(aux_sym__multiline_string_fragment_token1);
      if (lookahead == '*') ADVANCE(101);
      if (lookahead == '/') ADVANCE(98);
      if (lookahead != 0 &&
          lookahead != '"' &&
          lookahead != '\\') ADVANCE(103);
      END_STATE();
    case 100:
      ACCEPT_TOKEN(aux_sym__multiline_string_fragment_token1);
      if (lookahead == '*') ADVANCE(100);
      if (lookahead == '/') ADVANCE(103);
      if (lookahead == '"' ||
          lookahead == '\\') ADVANCE(13);
      if (lookahead != 0) ADVANCE(101);
      END_STATE();
    case 101:
      ACCEPT_TOKEN(aux_sym__multiline_string_fragment_token1);
      if (lookahead == '*') ADVANCE(100);
      if (lookahead == '"' ||
          lookahead == '\\') ADVANCE(13);
      if (lookahead != 0) ADVANCE(101);
      END_STATE();
    case 102:
      ACCEPT_TOKEN(aux_sym__multiline_string_fragment_token1);
      if (lookahead == '/') ADVANCE(99);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') ADVANCE(102);
      if (lookahead != 0 &&
          lookahead != '"' &&
          lookahead != '\\') ADVANCE(103);
      END_STATE();
    case 103:
      ACCEPT_TOKEN(aux_sym__multiline_string_fragment_token1);
      if (lookahead != 0 &&
          lookahead != '"' &&
          lookahead != '\\') ADVANCE(103);
      END_STATE();
    case 104:
      ACCEPT_TOKEN(aux_sym__multiline_string_fragment_token2);
      if (lookahead == '"') ADVANCE(5);
      if (lookahead == '\\') ADVANCE(7);
      if (lookahead != 0) ADVANCE(105);
      END_STATE();
    case 105:
      ACCEPT_TOKEN(aux_sym__multiline_string_fragment_token2);
      if (lookahead == '\\') ADVANCE(7);
      if (lookahead != 0 &&
          lookahead != '"') ADVANCE(105);
      END_STATE();
    case 106:
      ACCEPT_TOKEN(anon_sym_BSLASH_LBRACE);
      END_STATE();
    case 107:
      ACCEPT_TOKEN(anon_sym_RBRACE);
      END_STATE();
    case 108:
      ACCEPT_TOKEN(aux_sym__escape_sequence_token1);
      END_STATE();
    case 109:
      ACCEPT_TOKEN(aux_sym__escape_sequence_token1);
      if (lookahead == '{') ADVANCE(59);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(61);
      END_STATE();
    case 110:
      ACCEPT_TOKEN(aux_sym__escape_sequence_token1);
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(113);
      END_STATE();
    case 111:
      ACCEPT_TOKEN(aux_sym__escape_sequence_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(56);
      END_STATE();
    case 112:
      ACCEPT_TOKEN(sym_escape_sequence);
      END_STATE();
    case 113:
      ACCEPT_TOKEN(sym_escape_sequence);
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(112);
      END_STATE();
    case 114:
      ACCEPT_TOKEN(sym_escape_sequence);
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(113);
      END_STATE();
    case 115:
      ACCEPT_TOKEN(anon_sym_LPAREN);
      END_STATE();
    case 116:
      ACCEPT_TOKEN(anon_sym_RPAREN);
      END_STATE();
    case 117:
      ACCEPT_TOKEN(anon_sym_AMP);
      END_STATE();
    case 118:
      ACCEPT_TOKEN(anon_sym_AMP);
      if (lookahead == '&') ADVANCE(143);
      END_STATE();
    case 119:
      ACCEPT_TOKEN(anon_sym_AMP);
      if (lookahead == '&') ADVANCE(143);
      if (lookahead == '=') ADVANCE(126);
      END_STATE();
    case 120:
      ACCEPT_TOKEN(anon_sym_EQ);
      END_STATE();
    case 121:
      ACCEPT_TOKEN(anon_sym_EQ);
      if (lookahead == '=') ADVANCE(141);
      END_STATE();
    case 122:
      ACCEPT_TOKEN(anon_sym_PLUS_EQ);
      END_STATE();
    case 123:
      ACCEPT_TOKEN(anon_sym_DASH_EQ);
      END_STATE();
    case 124:
      ACCEPT_TOKEN(anon_sym_STAR_EQ);
      END_STATE();
    case 125:
      ACCEPT_TOKEN(anon_sym_SLASH_EQ);
      END_STATE();
    case 126:
      ACCEPT_TOKEN(anon_sym_AMP_EQ);
      END_STATE();
    case 127:
      ACCEPT_TOKEN(anon_sym_PIPE_EQ);
      END_STATE();
    case 128:
      ACCEPT_TOKEN(anon_sym_CARET_EQ);
      END_STATE();
    case 129:
      ACCEPT_TOKEN(anon_sym_PERCENT_EQ);
      END_STATE();
    case 130:
      ACCEPT_TOKEN(anon_sym_LT_LT_EQ);
      END_STATE();
    case 131:
      ACCEPT_TOKEN(anon_sym_GT_GT_EQ);
      END_STATE();
    case 132:
      ACCEPT_TOKEN(anon_sym_GT_GT_GT_EQ);
      END_STATE();
    case 133:
      ACCEPT_TOKEN(anon_sym_GT);
      END_STATE();
    case 134:
      ACCEPT_TOKEN(anon_sym_GT);
      if (lookahead == '=') ADVANCE(139);
      if (lookahead == '>') ADVANCE(162);
      END_STATE();
    case 135:
      ACCEPT_TOKEN(anon_sym_GT);
      if (lookahead == '=') ADVANCE(139);
      if (lookahead == '>') ADVANCE(163);
      END_STATE();
    case 136:
      ACCEPT_TOKEN(anon_sym_LT);
      END_STATE();
    case 137:
      ACCEPT_TOKEN(anon_sym_LT);
      if (lookahead == '<') ADVANCE(161);
      if (lookahead == '=') ADVANCE(140);
      END_STATE();
    case 138:
      ACCEPT_TOKEN(anon_sym_LT);
      if (lookahead == '<') ADVANCE(160);
      if (lookahead == '=') ADVANCE(140);
      END_STATE();
    case 139:
      ACCEPT_TOKEN(anon_sym_GT_EQ);
      END_STATE();
    case 140:
      ACCEPT_TOKEN(anon_sym_LT_EQ);
      END_STATE();
    case 141:
      ACCEPT_TOKEN(anon_sym_EQ_EQ);
      END_STATE();
    case 142:
      ACCEPT_TOKEN(anon_sym_BANG_EQ);
      END_STATE();
    case 143:
      ACCEPT_TOKEN(anon_sym_AMP_AMP);
      END_STATE();
    case 144:
      ACCEPT_TOKEN(anon_sym_PIPE_PIPE);
      END_STATE();
    case 145:
      ACCEPT_TOKEN(anon_sym_PLUS);
      if (lookahead == '+') ADVANCE(172);
      END_STATE();
    case 146:
      ACCEPT_TOKEN(anon_sym_PLUS);
      if (lookahead == '+') ADVANCE(172);
      if (lookahead == '=') ADVANCE(122);
      END_STATE();
    case 147:
      ACCEPT_TOKEN(anon_sym_DASH);
      if (lookahead == '-') ADVANCE(173);
      if (lookahead == '=') ADVANCE(123);
      if (lookahead == '>') ADVANCE(166);
      END_STATE();
    case 148:
      ACCEPT_TOKEN(anon_sym_DASH);
      if (lookahead == '-') ADVANCE(173);
      if (lookahead == '>') ADVANCE(166);
      END_STATE();
    case 149:
      ACCEPT_TOKEN(anon_sym_STAR);
      END_STATE();
    case 150:
      ACCEPT_TOKEN(anon_sym_STAR);
      if (lookahead == '=') ADVANCE(124);
      END_STATE();
    case 151:
      ACCEPT_TOKEN(anon_sym_SLASH);
      if (lookahead == '*') ADVANCE(13);
      if (lookahead == '/') ADVANCE(192);
      END_STATE();
    case 152:
      ACCEPT_TOKEN(anon_sym_SLASH);
      if (lookahead == '*') ADVANCE(13);
      if (lookahead == '/') ADVANCE(192);
      if (lookahead == '=') ADVANCE(125);
      END_STATE();
    case 153:
      ACCEPT_TOKEN(anon_sym_PIPE);
      END_STATE();
    case 154:
      ACCEPT_TOKEN(anon_sym_PIPE);
      if (lookahead == '=') ADVANCE(127);
      if (lookahead == '|') ADVANCE(144);
      END_STATE();
    case 155:
      ACCEPT_TOKEN(anon_sym_PIPE);
      if (lookahead == '|') ADVANCE(144);
      END_STATE();
    case 156:
      ACCEPT_TOKEN(anon_sym_CARET);
      END_STATE();
    case 157:
      ACCEPT_TOKEN(anon_sym_CARET);
      if (lookahead == '=') ADVANCE(128);
      END_STATE();
    case 158:
      ACCEPT_TOKEN(anon_sym_PERCENT);
      END_STATE();
    case 159:
      ACCEPT_TOKEN(anon_sym_PERCENT);
      if (lookahead == '=') ADVANCE(129);
      END_STATE();
    case 160:
      ACCEPT_TOKEN(anon_sym_LT_LT);
      END_STATE();
    case 161:
      ACCEPT_TOKEN(anon_sym_LT_LT);
      if (lookahead == '=') ADVANCE(130);
      END_STATE();
    case 162:
      ACCEPT_TOKEN(anon_sym_GT_GT);
      if (lookahead == '=') ADVANCE(131);
      if (lookahead == '>') ADVANCE(165);
      END_STATE();
    case 163:
      ACCEPT_TOKEN(anon_sym_GT_GT);
      if (lookahead == '>') ADVANCE(164);
      END_STATE();
    case 164:
      ACCEPT_TOKEN(anon_sym_GT_GT_GT);
      END_STATE();
    case 165:
      ACCEPT_TOKEN(anon_sym_GT_GT_GT);
      if (lookahead == '=') ADVANCE(132);
      END_STATE();
    case 166:
      ACCEPT_TOKEN(anon_sym_DASH_GT);
      END_STATE();
    case 167:
      ACCEPT_TOKEN(anon_sym_COMMA);
      END_STATE();
    case 168:
      ACCEPT_TOKEN(anon_sym_QMARK);
      END_STATE();
    case 169:
      ACCEPT_TOKEN(anon_sym_COLON);
      if (lookahead == ':') ADVANCE(180);
      END_STATE();
    case 170:
      ACCEPT_TOKEN(anon_sym_BANG);
      if (lookahead == '=') ADVANCE(142);
      END_STATE();
    case 171:
      ACCEPT_TOKEN(anon_sym_TILDE);
      END_STATE();
    case 172:
      ACCEPT_TOKEN(anon_sym_PLUS_PLUS);
      END_STATE();
    case 173:
      ACCEPT_TOKEN(anon_sym_DASH_DASH);
      END_STATE();
    case 174:
      ACCEPT_TOKEN(anon_sym_LBRACK);
      END_STATE();
    case 175:
      ACCEPT_TOKEN(anon_sym_RBRACK);
      END_STATE();
    case 176:
      ACCEPT_TOKEN(anon_sym_DOT);
      END_STATE();
    case 177:
      ACCEPT_TOKEN(anon_sym_DOT);
      if (lookahead == '.') ADVANCE(15);
      END_STATE();
    case 178:
      ACCEPT_TOKEN(anon_sym_DOT);
      if (lookahead == '.') ADVANCE(15);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(79);
      END_STATE();
    case 179:
      ACCEPT_TOKEN(anon_sym_DOT);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(79);
      END_STATE();
    case 180:
      ACCEPT_TOKEN(anon_sym_COLON_COLON);
      END_STATE();
    case 181:
      ACCEPT_TOKEN(anon_sym_LBRACE);
      END_STATE();
    case 182:
      ACCEPT_TOKEN(anon_sym_SEMI);
      END_STATE();
    case 183:
      ACCEPT_TOKEN(anon_sym_AT);
      END_STATE();
    case 184:
      ACCEPT_TOKEN(anon_sym_AT);
      if (lookahead == 'i') ADVANCE(38);
      END_STATE();
    case 185:
      ACCEPT_TOKEN(anon_sym_non_DASHsealed);
      END_STATE();
    case 186:
      ACCEPT_TOKEN(anon_sym_ATinterface);
      END_STATE();
    case 187:
      ACCEPT_TOKEN(anon_sym_DOT_DOT_DOT);
      END_STATE();
    case 188:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '-') ADVANCE(40);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(191);
      END_STATE();
    case 189:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == 'n') ADVANCE(188);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(191);
      END_STATE();
    case 190:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == 'o') ADVANCE(189);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(191);
      END_STATE();
    case 191:
      ACCEPT_TOKEN(sym_identifier);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(191);
      END_STATE();
    case 192:
      ACCEPT_TOKEN(sym_line_comment);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(192);
      END_STATE();
    case 193:
      ACCEPT_TOKEN(sym_block_comment);
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
        '_', 1,
        'a', 2,
        'b', 3,
        'c', 4,
        'd', 5,
        'e', 6,
        'f', 7,
        'i', 8,
        'l', 9,
        'm', 10,
        'n', 11,
        'o', 12,
        'p', 13,
        'r', 14,
        's', 15,
        't', 16,
        'u', 17,
        'v', 18,
        'w', 19,
        'y', 20,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(0);
      END_STATE();
    case 1:
      ACCEPT_TOKEN(sym_underscore_pattern);
      END_STATE();
    case 2:
      if (lookahead == 'b') ADVANCE(21);
      if (lookahead == 's') ADVANCE(22);
      END_STATE();
    case 3:
      if (lookahead == 'o') ADVANCE(23);
      if (lookahead == 'r') ADVANCE(24);
      if (lookahead == 'y') ADVANCE(25);
      END_STATE();
    case 4:
      if (lookahead == 'a') ADVANCE(26);
      if (lookahead == 'h') ADVANCE(27);
      if (lookahead == 'l') ADVANCE(28);
      if (lookahead == 'o') ADVANCE(29);
      END_STATE();
    case 5:
      if (lookahead == 'e') ADVANCE(30);
      if (lookahead == 'o') ADVANCE(31);
      END_STATE();
    case 6:
      if (lookahead == 'l') ADVANCE(32);
      if (lookahead == 'n') ADVANCE(33);
      if (lookahead == 'x') ADVANCE(34);
      END_STATE();
    case 7:
      if (lookahead == 'a') ADVANCE(35);
      if (lookahead == 'i') ADVANCE(36);
      if (lookahead == 'l') ADVANCE(37);
      if (lookahead == 'o') ADVANCE(38);
      END_STATE();
    case 8:
      if (lookahead == 'f') ADVANCE(39);
      if (lookahead == 'm') ADVANCE(40);
      if (lookahead == 'n') ADVANCE(41);
      END_STATE();
    case 9:
      if (lookahead == 'o') ADVANCE(42);
      END_STATE();
    case 10:
      if (lookahead == 'o') ADVANCE(43);
      END_STATE();
    case 11:
      if (lookahead == 'a') ADVANCE(44);
      if (lookahead == 'e') ADVANCE(45);
      if (lookahead == 'u') ADVANCE(46);
      END_STATE();
    case 12:
      if (lookahead == 'p') ADVANCE(47);
      END_STATE();
    case 13:
      if (lookahead == 'a') ADVANCE(48);
      if (lookahead == 'e') ADVANCE(49);
      if (lookahead == 'r') ADVANCE(50);
      if (lookahead == 'u') ADVANCE(51);
      END_STATE();
    case 14:
      if (lookahead == 'e') ADVANCE(52);
      END_STATE();
    case 15:
      if (lookahead == 'e') ADVANCE(53);
      if (lookahead == 'h') ADVANCE(54);
      if (lookahead == 't') ADVANCE(55);
      if (lookahead == 'u') ADVANCE(56);
      if (lookahead == 'w') ADVANCE(57);
      if (lookahead == 'y') ADVANCE(58);
      END_STATE();
    case 16:
      if (lookahead == 'h') ADVANCE(59);
      if (lookahead == 'o') ADVANCE(60);
      if (lookahead == 'r') ADVANCE(61);
      END_STATE();
    case 17:
      if (lookahead == 's') ADVANCE(62);
      END_STATE();
    case 18:
      if (lookahead == 'o') ADVANCE(63);
      END_STATE();
    case 19:
      if (lookahead == 'h') ADVANCE(64);
      if (lookahead == 'i') ADVANCE(65);
      END_STATE();
    case 20:
      if (lookahead == 'i') ADVANCE(66);
      END_STATE();
    case 21:
      if (lookahead == 's') ADVANCE(67);
      END_STATE();
    case 22:
      if (lookahead == 's') ADVANCE(68);
      END_STATE();
    case 23:
      if (lookahead == 'o') ADVANCE(69);
      END_STATE();
    case 24:
      if (lookahead == 'e') ADVANCE(70);
      END_STATE();
    case 25:
      if (lookahead == 't') ADVANCE(71);
      END_STATE();
    case 26:
      if (lookahead == 's') ADVANCE(72);
      if (lookahead == 't') ADVANCE(73);
      END_STATE();
    case 27:
      if (lookahead == 'a') ADVANCE(74);
      END_STATE();
    case 28:
      if (lookahead == 'a') ADVANCE(75);
      END_STATE();
    case 29:
      if (lookahead == 'n') ADVANCE(76);
      END_STATE();
    case 30:
      if (lookahead == 'f') ADVANCE(77);
      END_STATE();
    case 31:
      ACCEPT_TOKEN(anon_sym_do);
      if (lookahead == 'u') ADVANCE(78);
      END_STATE();
    case 32:
      if (lookahead == 's') ADVANCE(79);
      END_STATE();
    case 33:
      if (lookahead == 'u') ADVANCE(80);
      END_STATE();
    case 34:
      if (lookahead == 'p') ADVANCE(81);
      if (lookahead == 't') ADVANCE(82);
      END_STATE();
    case 35:
      if (lookahead == 'l') ADVANCE(83);
      END_STATE();
    case 36:
      if (lookahead == 'n') ADVANCE(84);
      END_STATE();
    case 37:
      if (lookahead == 'o') ADVANCE(85);
      END_STATE();
    case 38:
      if (lookahead == 'r') ADVANCE(86);
      END_STATE();
    case 39:
      ACCEPT_TOKEN(anon_sym_if);
      END_STATE();
    case 40:
      if (lookahead == 'p') ADVANCE(87);
      END_STATE();
    case 41:
      if (lookahead == 's') ADVANCE(88);
      if (lookahead == 't') ADVANCE(89);
      END_STATE();
    case 42:
      if (lookahead == 'n') ADVANCE(90);
      END_STATE();
    case 43:
      if (lookahead == 'd') ADVANCE(91);
      END_STATE();
    case 44:
      if (lookahead == 't') ADVANCE(92);
      END_STATE();
    case 45:
      if (lookahead == 'w') ADVANCE(93);
      END_STATE();
    case 46:
      if (lookahead == 'l') ADVANCE(94);
      END_STATE();
    case 47:
      if (lookahead == 'e') ADVANCE(95);
      END_STATE();
    case 48:
      if (lookahead == 'c') ADVANCE(96);
      END_STATE();
    case 49:
      if (lookahead == 'r') ADVANCE(97);
      END_STATE();
    case 50:
      if (lookahead == 'i') ADVANCE(98);
      if (lookahead == 'o') ADVANCE(99);
      END_STATE();
    case 51:
      if (lookahead == 'b') ADVANCE(100);
      END_STATE();
    case 52:
      if (lookahead == 'c') ADVANCE(101);
      if (lookahead == 'q') ADVANCE(102);
      if (lookahead == 't') ADVANCE(103);
      END_STATE();
    case 53:
      if (lookahead == 'a') ADVANCE(104);
      END_STATE();
    case 54:
      if (lookahead == 'o') ADVANCE(105);
      END_STATE();
    case 55:
      if (lookahead == 'a') ADVANCE(106);
      if (lookahead == 'r') ADVANCE(107);
      END_STATE();
    case 56:
      if (lookahead == 'p') ADVANCE(108);
      END_STATE();
    case 57:
      if (lookahead == 'i') ADVANCE(109);
      END_STATE();
    case 58:
      if (lookahead == 'n') ADVANCE(110);
      END_STATE();
    case 59:
      if (lookahead == 'i') ADVANCE(111);
      if (lookahead == 'r') ADVANCE(112);
      END_STATE();
    case 60:
      ACCEPT_TOKEN(anon_sym_to);
      END_STATE();
    case 61:
      if (lookahead == 'a') ADVANCE(113);
      if (lookahead == 'u') ADVANCE(114);
      if (lookahead == 'y') ADVANCE(115);
      END_STATE();
    case 62:
      if (lookahead == 'e') ADVANCE(116);
      END_STATE();
    case 63:
      if (lookahead == 'i') ADVANCE(117);
      if (lookahead == 'l') ADVANCE(118);
      END_STATE();
    case 64:
      if (lookahead == 'e') ADVANCE(119);
      if (lookahead == 'i') ADVANCE(120);
      END_STATE();
    case 65:
      if (lookahead == 't') ADVANCE(121);
      END_STATE();
    case 66:
      if (lookahead == 'e') ADVANCE(122);
      END_STATE();
    case 67:
      if (lookahead == 't') ADVANCE(123);
      END_STATE();
    case 68:
      if (lookahead == 'e') ADVANCE(124);
      END_STATE();
    case 69:
      if (lookahead == 'l') ADVANCE(125);
      END_STATE();
    case 70:
      if (lookahead == 'a') ADVANCE(126);
      END_STATE();
    case 71:
      if (lookahead == 'e') ADVANCE(127);
      END_STATE();
    case 72:
      if (lookahead == 'e') ADVANCE(128);
      END_STATE();
    case 73:
      if (lookahead == 'c') ADVANCE(129);
      END_STATE();
    case 74:
      if (lookahead == 'r') ADVANCE(130);
      END_STATE();
    case 75:
      if (lookahead == 's') ADVANCE(131);
      END_STATE();
    case 76:
      if (lookahead == 't') ADVANCE(132);
      END_STATE();
    case 77:
      if (lookahead == 'a') ADVANCE(133);
      END_STATE();
    case 78:
      if (lookahead == 'b') ADVANCE(134);
      END_STATE();
    case 79:
      if (lookahead == 'e') ADVANCE(135);
      END_STATE();
    case 80:
      if (lookahead == 'm') ADVANCE(136);
      END_STATE();
    case 81:
      if (lookahead == 'o') ADVANCE(137);
      END_STATE();
    case 82:
      if (lookahead == 'e') ADVANCE(138);
      END_STATE();
    case 83:
      if (lookahead == 's') ADVANCE(139);
      END_STATE();
    case 84:
      if (lookahead == 'a') ADVANCE(140);
      END_STATE();
    case 85:
      if (lookahead == 'a') ADVANCE(141);
      END_STATE();
    case 86:
      ACCEPT_TOKEN(anon_sym_for);
      END_STATE();
    case 87:
      if (lookahead == 'l') ADVANCE(142);
      if (lookahead == 'o') ADVANCE(143);
      END_STATE();
    case 88:
      if (lookahead == 't') ADVANCE(144);
      END_STATE();
    case 89:
      ACCEPT_TOKEN(anon_sym_int);
      if (lookahead == 'e') ADVANCE(145);
      END_STATE();
    case 90:
      if (lookahead == 'g') ADVANCE(146);
      END_STATE();
    case 91:
      if (lookahead == 'u') ADVANCE(147);
      END_STATE();
    case 92:
      if (lookahead == 'i') ADVANCE(148);
      END_STATE();
    case 93:
      ACCEPT_TOKEN(anon_sym_new);
      END_STATE();
    case 94:
      if (lookahead == 'l') ADVANCE(149);
      END_STATE();
    case 95:
      if (lookahead == 'n') ADVANCE(150);
      END_STATE();
    case 96:
      if (lookahead == 'k') ADVANCE(151);
      END_STATE();
    case 97:
      if (lookahead == 'm') ADVANCE(152);
      END_STATE();
    case 98:
      if (lookahead == 'v') ADVANCE(153);
      END_STATE();
    case 99:
      if (lookahead == 't') ADVANCE(154);
      if (lookahead == 'v') ADVANCE(155);
      END_STATE();
    case 100:
      if (lookahead == 'l') ADVANCE(156);
      END_STATE();
    case 101:
      if (lookahead == 'o') ADVANCE(157);
      END_STATE();
    case 102:
      if (lookahead == 'u') ADVANCE(158);
      END_STATE();
    case 103:
      if (lookahead == 'u') ADVANCE(159);
      END_STATE();
    case 104:
      if (lookahead == 'l') ADVANCE(160);
      END_STATE();
    case 105:
      if (lookahead == 'r') ADVANCE(161);
      END_STATE();
    case 106:
      if (lookahead == 't') ADVANCE(162);
      END_STATE();
    case 107:
      if (lookahead == 'i') ADVANCE(163);
      END_STATE();
    case 108:
      if (lookahead == 'e') ADVANCE(164);
      END_STATE();
    case 109:
      if (lookahead == 't') ADVANCE(165);
      END_STATE();
    case 110:
      if (lookahead == 'c') ADVANCE(166);
      END_STATE();
    case 111:
      if (lookahead == 's') ADVANCE(167);
      END_STATE();
    case 112:
      if (lookahead == 'o') ADVANCE(168);
      END_STATE();
    case 113:
      if (lookahead == 'n') ADVANCE(169);
      END_STATE();
    case 114:
      if (lookahead == 'e') ADVANCE(170);
      END_STATE();
    case 115:
      ACCEPT_TOKEN(anon_sym_try);
      END_STATE();
    case 116:
      if (lookahead == 's') ADVANCE(171);
      END_STATE();
    case 117:
      if (lookahead == 'd') ADVANCE(172);
      END_STATE();
    case 118:
      if (lookahead == 'a') ADVANCE(173);
      END_STATE();
    case 119:
      if (lookahead == 'n') ADVANCE(174);
      END_STATE();
    case 120:
      if (lookahead == 'l') ADVANCE(175);
      END_STATE();
    case 121:
      if (lookahead == 'h') ADVANCE(176);
      END_STATE();
    case 122:
      if (lookahead == 'l') ADVANCE(177);
      END_STATE();
    case 123:
      if (lookahead == 'r') ADVANCE(178);
      END_STATE();
    case 124:
      if (lookahead == 'r') ADVANCE(179);
      END_STATE();
    case 125:
      if (lookahead == 'e') ADVANCE(180);
      END_STATE();
    case 126:
      if (lookahead == 'k') ADVANCE(181);
      END_STATE();
    case 127:
      ACCEPT_TOKEN(anon_sym_byte);
      END_STATE();
    case 128:
      ACCEPT_TOKEN(anon_sym_case);
      END_STATE();
    case 129:
      if (lookahead == 'h') ADVANCE(182);
      END_STATE();
    case 130:
      ACCEPT_TOKEN(anon_sym_char);
      END_STATE();
    case 131:
      if (lookahead == 's') ADVANCE(183);
      END_STATE();
    case 132:
      if (lookahead == 'i') ADVANCE(184);
      END_STATE();
    case 133:
      if (lookahead == 'u') ADVANCE(185);
      END_STATE();
    case 134:
      if (lookahead == 'l') ADVANCE(186);
      END_STATE();
    case 135:
      ACCEPT_TOKEN(anon_sym_else);
      END_STATE();
    case 136:
      ACCEPT_TOKEN(anon_sym_enum);
      END_STATE();
    case 137:
      if (lookahead == 'r') ADVANCE(187);
      END_STATE();
    case 138:
      if (lookahead == 'n') ADVANCE(188);
      END_STATE();
    case 139:
      if (lookahead == 'e') ADVANCE(189);
      END_STATE();
    case 140:
      if (lookahead == 'l') ADVANCE(190);
      END_STATE();
    case 141:
      if (lookahead == 't') ADVANCE(191);
      END_STATE();
    case 142:
      if (lookahead == 'e') ADVANCE(192);
      END_STATE();
    case 143:
      if (lookahead == 'r') ADVANCE(193);
      END_STATE();
    case 144:
      if (lookahead == 'a') ADVANCE(194);
      END_STATE();
    case 145:
      if (lookahead == 'r') ADVANCE(195);
      END_STATE();
    case 146:
      ACCEPT_TOKEN(anon_sym_long);
      END_STATE();
    case 147:
      if (lookahead == 'l') ADVANCE(196);
      END_STATE();
    case 148:
      if (lookahead == 'v') ADVANCE(197);
      END_STATE();
    case 149:
      ACCEPT_TOKEN(sym_null_literal);
      END_STATE();
    case 150:
      ACCEPT_TOKEN(anon_sym_open);
      if (lookahead == 's') ADVANCE(198);
      END_STATE();
    case 151:
      if (lookahead == 'a') ADVANCE(199);
      END_STATE();
    case 152:
      if (lookahead == 'i') ADVANCE(200);
      END_STATE();
    case 153:
      if (lookahead == 'a') ADVANCE(201);
      END_STATE();
    case 154:
      if (lookahead == 'e') ADVANCE(202);
      END_STATE();
    case 155:
      if (lookahead == 'i') ADVANCE(203);
      END_STATE();
    case 156:
      if (lookahead == 'i') ADVANCE(204);
      END_STATE();
    case 157:
      if (lookahead == 'r') ADVANCE(205);
      END_STATE();
    case 158:
      if (lookahead == 'i') ADVANCE(206);
      END_STATE();
    case 159:
      if (lookahead == 'r') ADVANCE(207);
      END_STATE();
    case 160:
      if (lookahead == 'e') ADVANCE(208);
      END_STATE();
    case 161:
      if (lookahead == 't') ADVANCE(209);
      END_STATE();
    case 162:
      if (lookahead == 'i') ADVANCE(210);
      END_STATE();
    case 163:
      if (lookahead == 'c') ADVANCE(211);
      END_STATE();
    case 164:
      if (lookahead == 'r') ADVANCE(212);
      END_STATE();
    case 165:
      if (lookahead == 'c') ADVANCE(213);
      END_STATE();
    case 166:
      if (lookahead == 'h') ADVANCE(214);
      END_STATE();
    case 167:
      ACCEPT_TOKEN(sym_this);
      END_STATE();
    case 168:
      if (lookahead == 'w') ADVANCE(215);
      END_STATE();
    case 169:
      if (lookahead == 's') ADVANCE(216);
      END_STATE();
    case 170:
      ACCEPT_TOKEN(sym_true);
      END_STATE();
    case 171:
      ACCEPT_TOKEN(anon_sym_uses);
      END_STATE();
    case 172:
      ACCEPT_TOKEN(sym_void_type);
      END_STATE();
    case 173:
      if (lookahead == 't') ADVANCE(217);
      END_STATE();
    case 174:
      ACCEPT_TOKEN(anon_sym_when);
      END_STATE();
    case 175:
      if (lookahead == 'e') ADVANCE(218);
      END_STATE();
    case 176:
      ACCEPT_TOKEN(anon_sym_with);
      END_STATE();
    case 177:
      if (lookahead == 'd') ADVANCE(219);
      END_STATE();
    case 178:
      if (lookahead == 'a') ADVANCE(220);
      END_STATE();
    case 179:
      if (lookahead == 't') ADVANCE(221);
      END_STATE();
    case 180:
      if (lookahead == 'a') ADVANCE(222);
      END_STATE();
    case 181:
      ACCEPT_TOKEN(anon_sym_break);
      END_STATE();
    case 182:
      ACCEPT_TOKEN(anon_sym_catch);
      END_STATE();
    case 183:
      ACCEPT_TOKEN(anon_sym_class);
      END_STATE();
    case 184:
      if (lookahead == 'n') ADVANCE(223);
      END_STATE();
    case 185:
      if (lookahead == 'l') ADVANCE(224);
      END_STATE();
    case 186:
      if (lookahead == 'e') ADVANCE(225);
      END_STATE();
    case 187:
      if (lookahead == 't') ADVANCE(226);
      END_STATE();
    case 188:
      if (lookahead == 'd') ADVANCE(227);
      END_STATE();
    case 189:
      ACCEPT_TOKEN(sym_false);
      END_STATE();
    case 190:
      ACCEPT_TOKEN(anon_sym_final);
      if (lookahead == 'l') ADVANCE(228);
      END_STATE();
    case 191:
      ACCEPT_TOKEN(anon_sym_float);
      END_STATE();
    case 192:
      if (lookahead == 'm') ADVANCE(229);
      END_STATE();
    case 193:
      if (lookahead == 't') ADVANCE(230);
      END_STATE();
    case 194:
      if (lookahead == 'n') ADVANCE(231);
      END_STATE();
    case 195:
      if (lookahead == 'f') ADVANCE(232);
      END_STATE();
    case 196:
      if (lookahead == 'e') ADVANCE(233);
      END_STATE();
    case 197:
      if (lookahead == 'e') ADVANCE(234);
      END_STATE();
    case 198:
      ACCEPT_TOKEN(anon_sym_opens);
      END_STATE();
    case 199:
      if (lookahead == 'g') ADVANCE(235);
      END_STATE();
    case 200:
      if (lookahead == 't') ADVANCE(236);
      END_STATE();
    case 201:
      if (lookahead == 't') ADVANCE(237);
      END_STATE();
    case 202:
      if (lookahead == 'c') ADVANCE(238);
      END_STATE();
    case 203:
      if (lookahead == 'd') ADVANCE(239);
      END_STATE();
    case 204:
      if (lookahead == 'c') ADVANCE(240);
      END_STATE();
    case 205:
      if (lookahead == 'd') ADVANCE(241);
      END_STATE();
    case 206:
      if (lookahead == 'r') ADVANCE(242);
      END_STATE();
    case 207:
      if (lookahead == 'n') ADVANCE(243);
      END_STATE();
    case 208:
      if (lookahead == 'd') ADVANCE(244);
      END_STATE();
    case 209:
      ACCEPT_TOKEN(anon_sym_short);
      END_STATE();
    case 210:
      if (lookahead == 'c') ADVANCE(245);
      END_STATE();
    case 211:
      if (lookahead == 't') ADVANCE(246);
      END_STATE();
    case 212:
      ACCEPT_TOKEN(sym_super);
      END_STATE();
    case 213:
      if (lookahead == 'h') ADVANCE(247);
      END_STATE();
    case 214:
      if (lookahead == 'r') ADVANCE(248);
      END_STATE();
    case 215:
      ACCEPT_TOKEN(anon_sym_throw);
      if (lookahead == 's') ADVANCE(249);
      END_STATE();
    case 216:
      if (lookahead == 'i') ADVANCE(250);
      END_STATE();
    case 217:
      if (lookahead == 'i') ADVANCE(251);
      END_STATE();
    case 218:
      ACCEPT_TOKEN(anon_sym_while);
      END_STATE();
    case 219:
      ACCEPT_TOKEN(anon_sym_yield);
      END_STATE();
    case 220:
      if (lookahead == 'c') ADVANCE(252);
      END_STATE();
    case 221:
      ACCEPT_TOKEN(anon_sym_assert);
      END_STATE();
    case 222:
      if (lookahead == 'n') ADVANCE(253);
      END_STATE();
    case 223:
      if (lookahead == 'u') ADVANCE(254);
      END_STATE();
    case 224:
      if (lookahead == 't') ADVANCE(255);
      END_STATE();
    case 225:
      ACCEPT_TOKEN(anon_sym_double);
      END_STATE();
    case 226:
      if (lookahead == 's') ADVANCE(256);
      END_STATE();
    case 227:
      if (lookahead == 's') ADVANCE(257);
      END_STATE();
    case 228:
      if (lookahead == 'y') ADVANCE(258);
      END_STATE();
    case 229:
      if (lookahead == 'e') ADVANCE(259);
      END_STATE();
    case 230:
      ACCEPT_TOKEN(anon_sym_import);
      END_STATE();
    case 231:
      if (lookahead == 'c') ADVANCE(260);
      END_STATE();
    case 232:
      if (lookahead == 'a') ADVANCE(261);
      END_STATE();
    case 233:
      ACCEPT_TOKEN(anon_sym_module);
      END_STATE();
    case 234:
      ACCEPT_TOKEN(anon_sym_native);
      END_STATE();
    case 235:
      if (lookahead == 'e') ADVANCE(262);
      END_STATE();
    case 236:
      if (lookahead == 's') ADVANCE(263);
      END_STATE();
    case 237:
      if (lookahead == 'e') ADVANCE(264);
      END_STATE();
    case 238:
      if (lookahead == 't') ADVANCE(265);
      END_STATE();
    case 239:
      if (lookahead == 'e') ADVANCE(266);
      END_STATE();
    case 240:
      ACCEPT_TOKEN(anon_sym_public);
      END_STATE();
    case 241:
      ACCEPT_TOKEN(anon_sym_record);
      END_STATE();
    case 242:
      if (lookahead == 'e') ADVANCE(267);
      END_STATE();
    case 243:
      ACCEPT_TOKEN(anon_sym_return);
      END_STATE();
    case 244:
      ACCEPT_TOKEN(anon_sym_sealed);
      END_STATE();
    case 245:
      ACCEPT_TOKEN(anon_sym_static);
      END_STATE();
    case 246:
      if (lookahead == 'f') ADVANCE(268);
      END_STATE();
    case 247:
      ACCEPT_TOKEN(anon_sym_switch);
      END_STATE();
    case 248:
      if (lookahead == 'o') ADVANCE(269);
      END_STATE();
    case 249:
      ACCEPT_TOKEN(anon_sym_throws);
      END_STATE();
    case 250:
      if (lookahead == 'e') ADVANCE(270);
      if (lookahead == 't') ADVANCE(271);
      END_STATE();
    case 251:
      if (lookahead == 'l') ADVANCE(272);
      END_STATE();
    case 252:
      if (lookahead == 't') ADVANCE(273);
      END_STATE();
    case 253:
      ACCEPT_TOKEN(sym_boolean_type);
      END_STATE();
    case 254:
      if (lookahead == 'e') ADVANCE(274);
      END_STATE();
    case 255:
      ACCEPT_TOKEN(anon_sym_default);
      END_STATE();
    case 256:
      ACCEPT_TOKEN(anon_sym_exports);
      END_STATE();
    case 257:
      ACCEPT_TOKEN(anon_sym_extends);
      END_STATE();
    case 258:
      ACCEPT_TOKEN(anon_sym_finally);
      END_STATE();
    case 259:
      if (lookahead == 'n') ADVANCE(275);
      END_STATE();
    case 260:
      if (lookahead == 'e') ADVANCE(276);
      END_STATE();
    case 261:
      if (lookahead == 'c') ADVANCE(277);
      END_STATE();
    case 262:
      ACCEPT_TOKEN(anon_sym_package);
      END_STATE();
    case 263:
      ACCEPT_TOKEN(anon_sym_permits);
      END_STATE();
    case 264:
      ACCEPT_TOKEN(anon_sym_private);
      END_STATE();
    case 265:
      if (lookahead == 'e') ADVANCE(278);
      END_STATE();
    case 266:
      if (lookahead == 's') ADVANCE(279);
      END_STATE();
    case 267:
      if (lookahead == 's') ADVANCE(280);
      END_STATE();
    case 268:
      if (lookahead == 'p') ADVANCE(281);
      END_STATE();
    case 269:
      if (lookahead == 'n') ADVANCE(282);
      END_STATE();
    case 270:
      if (lookahead == 'n') ADVANCE(283);
      END_STATE();
    case 271:
      if (lookahead == 'i') ADVANCE(284);
      END_STATE();
    case 272:
      if (lookahead == 'e') ADVANCE(285);
      END_STATE();
    case 273:
      ACCEPT_TOKEN(anon_sym_abstract);
      END_STATE();
    case 274:
      ACCEPT_TOKEN(anon_sym_continue);
      END_STATE();
    case 275:
      if (lookahead == 't') ADVANCE(286);
      END_STATE();
    case 276:
      if (lookahead == 'o') ADVANCE(287);
      END_STATE();
    case 277:
      if (lookahead == 'e') ADVANCE(288);
      END_STATE();
    case 278:
      if (lookahead == 'd') ADVANCE(289);
      END_STATE();
    case 279:
      ACCEPT_TOKEN(anon_sym_provides);
      END_STATE();
    case 280:
      ACCEPT_TOKEN(anon_sym_requires);
      END_STATE();
    case 281:
      ACCEPT_TOKEN(anon_sym_strictfp);
      END_STATE();
    case 282:
      if (lookahead == 'i') ADVANCE(290);
      END_STATE();
    case 283:
      if (lookahead == 't') ADVANCE(291);
      END_STATE();
    case 284:
      if (lookahead == 'v') ADVANCE(292);
      END_STATE();
    case 285:
      ACCEPT_TOKEN(anon_sym_volatile);
      END_STATE();
    case 286:
      if (lookahead == 's') ADVANCE(293);
      END_STATE();
    case 287:
      if (lookahead == 'f') ADVANCE(294);
      END_STATE();
    case 288:
      ACCEPT_TOKEN(anon_sym_interface);
      END_STATE();
    case 289:
      ACCEPT_TOKEN(anon_sym_protected);
      END_STATE();
    case 290:
      if (lookahead == 'z') ADVANCE(295);
      END_STATE();
    case 291:
      ACCEPT_TOKEN(anon_sym_transient);
      END_STATE();
    case 292:
      if (lookahead == 'e') ADVANCE(296);
      END_STATE();
    case 293:
      ACCEPT_TOKEN(anon_sym_implements);
      END_STATE();
    case 294:
      ACCEPT_TOKEN(anon_sym_instanceof);
      END_STATE();
    case 295:
      if (lookahead == 'e') ADVANCE(297);
      END_STATE();
    case 296:
      ACCEPT_TOKEN(anon_sym_transitive);
      END_STATE();
    case 297:
      if (lookahead == 'd') ADVANCE(298);
      END_STATE();
    case 298:
      ACCEPT_TOKEN(anon_sym_synchronized);
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
static uint16_t ts_small_parse_table[31700];

// Filled from tables.zlib by grammar.go, see fetchgrammar.
static uint32_t ts_small_parse_table_map[979];

// Filled from tables.zlib by grammar.go, see fetchgrammar.
static TSParseActionEntry ts_parse_actions[3432];

#ifdef __cplusplus
extern "C" {
#endif
#ifdef TREE_SITTER_HIDE_SYMBOLS
#define TS_PUBLIC
#elif defined(_WIN32)
#define TS_PUBLIC __declspec(dllexport)
#else
#define TS_PUBLIC __attribute__((visibility("default")))
#endif

TS_PUBLIC const TSLanguage *tree_sitter_java(void) {
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
