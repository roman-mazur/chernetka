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
#define STATE_COUNT 8637
#define LARGE_STATE_COUNT 2322
#define SYMBOL_COUNT 538
#define ALIAS_COUNT 5
#define TOKEN_COUNT 219
#define EXTERNAL_TOKEN_COUNT 2
#define FIELD_COUNT 50
#define MAX_ALIAS_SEQUENCE_LENGTH 9
#define PRODUCTION_ID_COUNT 220

enum ts_symbol_identifiers {
  sym_identifier = 1,
  aux_sym_preproc_include_token1 = 2,
  aux_sym_preproc_include_token2 = 3,
  aux_sym_preproc_def_token1 = 4,
  anon_sym_LPAREN = 5,
  anon_sym_DOT_DOT_DOT = 6,
  anon_sym_COMMA = 7,
  anon_sym_RPAREN = 8,
  aux_sym_preproc_if_token1 = 9,
  anon_sym_LF = 10,
  aux_sym_preproc_if_token2 = 11,
  aux_sym_preproc_ifdef_token1 = 12,
  aux_sym_preproc_ifdef_token2 = 13,
  aux_sym_preproc_else_token1 = 14,
  aux_sym_preproc_elif_token1 = 15,
  aux_sym_preproc_elifdef_token1 = 16,
  aux_sym_preproc_elifdef_token2 = 17,
  sym_preproc_arg = 18,
  sym_preproc_directive = 19,
  anon_sym_LPAREN2 = 20,
  anon_sym_defined = 21,
  anon_sym_BANG = 22,
  anon_sym_TILDE = 23,
  anon_sym_DASH = 24,
  anon_sym_PLUS = 25,
  anon_sym_STAR = 26,
  anon_sym_SLASH = 27,
  anon_sym_PERCENT = 28,
  anon_sym_PIPE_PIPE = 29,
  anon_sym_AMP_AMP = 30,
  anon_sym_PIPE = 31,
  anon_sym_CARET = 32,
  anon_sym_AMP = 33,
  anon_sym_EQ_EQ = 34,
  anon_sym_BANG_EQ = 35,
  anon_sym_GT = 36,
  anon_sym_GT_EQ = 37,
  anon_sym_LT_EQ = 38,
  anon_sym_LT = 39,
  anon_sym_LT_LT = 40,
  anon_sym_GT_GT = 41,
  anon_sym_SEMI = 42,
  anon_sym___extension__ = 43,
  anon_sym_typedef = 44,
  anon_sym_virtual = 45,
  anon_sym_extern = 46,
  anon_sym___attribute__ = 47,
  anon_sym___attribute = 48,
  anon_sym_COLON_COLON = 49,
  anon_sym_LBRACK_LBRACK = 50,
  anon_sym_RBRACK_RBRACK = 51,
  anon_sym___declspec = 52,
  anon_sym___based = 53,
  anon_sym___cdecl = 54,
  anon_sym___clrcall = 55,
  anon_sym___stdcall = 56,
  anon_sym___fastcall = 57,
  anon_sym___thiscall = 58,
  anon_sym___vectorcall = 59,
  sym_ms_restrict_modifier = 60,
  sym_ms_unsigned_ptr_modifier = 61,
  sym_ms_signed_ptr_modifier = 62,
  anon_sym__unaligned = 63,
  anon_sym___unaligned = 64,
  anon_sym_LBRACE = 65,
  anon_sym_RBRACE = 66,
  anon_sym_signed = 67,
  anon_sym_unsigned = 68,
  anon_sym_long = 69,
  anon_sym_short = 70,
  anon_sym_LBRACK = 71,
  anon_sym_static = 72,
  anon_sym_RBRACK = 73,
  anon_sym_EQ = 74,
  anon_sym_register = 75,
  anon_sym_inline = 76,
  anon_sym___inline = 77,
  anon_sym___inline__ = 78,
  anon_sym___forceinline = 79,
  anon_sym_thread_local = 80,
  anon_sym___thread = 81,
  anon_sym_const = 82,
  anon_sym_constexpr = 83,
  anon_sym_volatile = 84,
  anon_sym_restrict = 85,
  anon_sym___restrict__ = 86,
  anon_sym__Atomic = 87,
  anon_sym__Noreturn = 88,
  anon_sym_noreturn = 89,
  anon_sym__Nonnull = 90,
  anon_sym_mutable = 91,
  anon_sym_constinit = 92,
  anon_sym_consteval = 93,
  anon_sym_alignas = 94,
  anon_sym__Alignas = 95,
  sym_primitive_type = 96,
  anon_sym_enum = 97,
  anon_sym_class = 98,
  anon_sym_struct = 99,
  anon_sym_union = 100,
  anon_sym_COLON = 101,
  anon_sym_if = 102,
  anon_sym_else = 103,
  anon_sym_switch = 104,
  anon_sym_case = 105,
  anon_sym_default = 106,
  anon_sym_while = 107,
  anon_sym_do = 108,
  anon_sym_for = 109,
  anon_sym_return = 110,
  anon_sym_break = 111,
  anon_sym_continue = 112,
  anon_sym_goto = 113,
  anon_sym___try = 114,
  anon_sym___except = 115,
  anon_sym___finally = 116,
  anon_sym___leave = 117,
  anon_sym_QMARK = 118,
  anon_sym_STAR_EQ = 119,
  anon_sym_SLASH_EQ = 120,
  anon_sym_PERCENT_EQ = 121,
  anon_sym_PLUS_EQ = 122,
  anon_sym_DASH_EQ = 123,
  anon_sym_LT_LT_EQ = 124,
  anon_sym_GT_GT_EQ = 125,
  anon_sym_AMP_EQ = 126,
  anon_sym_CARET_EQ = 127,
  anon_sym_PIPE_EQ = 128,
  anon_sym_and_eq = 129,
  anon_sym_or_eq = 130,
  anon_sym_xor_eq = 131,
  anon_sym_not = 132,
  anon_sym_compl = 133,
  anon_sym_LT_EQ_GT = 134,
  anon_sym_or = 135,
  anon_sym_and = 136,
  anon_sym_bitor = 137,
  anon_sym_xor = 138,
  anon_sym_bitand = 139,
  anon_sym_not_eq = 140,
  anon_sym_DASH_DASH = 141,
  anon_sym_PLUS_PLUS = 142,
  anon_sym_sizeof = 143,
  anon_sym___alignof__ = 144,
  anon_sym___alignof = 145,
  anon_sym__alignof = 146,
  anon_sym_alignof = 147,
  anon_sym__Alignof = 148,
  anon_sym_offsetof = 149,
  anon_sym__Generic = 150,
  anon_sym_asm = 151,
  anon_sym___asm__ = 152,
  anon_sym___asm = 153,
  anon_sym___volatile__ = 154,
  anon_sym_DOT = 155,
  anon_sym_DOT_STAR = 156,
  anon_sym_DASH_GT = 157,
  sym_number_literal = 158,
  anon_sym_L_SQUOTE = 159,
  anon_sym_u_SQUOTE = 160,
  anon_sym_U_SQUOTE = 161,
  anon_sym_u8_SQUOTE = 162,
  anon_sym_SQUOTE = 163,
  aux_sym_char_literal_token1 = 164,
  anon_sym_L_DQUOTE = 165,
  anon_sym_u_DQUOTE = 166,
  anon_sym_U_DQUOTE = 167,
  anon_sym_u8_DQUOTE = 168,
  anon_sym_DQUOTE = 169,
  aux_sym_string_literal_token1 = 170,
  sym_escape_sequence = 171,
  sym_system_lib_string = 172,
  sym_true = 173,
  sym_false = 174,
  anon_sym_NULL = 175,
  anon_sym_nullptr = 176,
  sym_comment = 177,
  sym_auto = 178,
  anon_sym_decltype = 179,
  anon_sym_final = 180,
  anon_sym_override = 181,
  anon_sym_explicit = 182,
  anon_sym_typename = 183,
  anon_sym_template = 184,
  anon_sym_GT2 = 185,
  anon_sym_operator = 186,
  anon_sym_try = 187,
  anon_sym_delete = 188,
  aux_sym_pure_virtual_clause_token1 = 189,
  anon_sym_friend = 190,
  anon_sym_public = 191,
  anon_sym_private = 192,
  anon_sym_protected = 193,
  anon_sym_noexcept = 194,
  anon_sym_throw = 195,
  anon_sym_namespace = 196,
  anon_sym_using = 197,
  anon_sym_static_assert = 198,
  anon_sym_concept = 199,
  anon_sym_co_return = 200,
  anon_sym_co_yield = 201,
  anon_sym_catch = 202,
  anon_sym_R_DQUOTE = 203,
  anon_sym_LR_DQUOTE = 204,
  anon_sym_uR_DQUOTE = 205,
  anon_sym_UR_DQUOTE = 206,
  anon_sym_u8R_DQUOTE = 207,
  anon_sym_co_await = 208,
  anon_sym_new = 209,
  anon_sym_requires = 210,
  anon_sym_DASH_GT_STAR = 211,
  anon_sym_LPAREN_RPAREN = 212,
  anon_sym_LBRACK_RBRACK = 213,
  anon_sym_DQUOTE_DQUOTE = 214,
  sym_this = 215,
  sym_literal_suffix = 216,
  sym_raw_string_delimiter = 217,
  sym_raw_string_content = 218,
  sym_translation_unit = 219,
  sym__top_level_item = 220,
  sym__block_item = 221,
  sym_preproc_include = 222,
  sym_preproc_def = 223,
  sym_preproc_function_def = 224,
  sym_preproc_params = 225,
  sym_preproc_call = 226,
  sym_preproc_if = 227,
  sym_preproc_ifdef = 228,
  sym_preproc_else = 229,
  sym_preproc_elif = 230,
  sym_preproc_elifdef = 231,
  sym_preproc_if_in_field_declaration_list = 232,
  sym_preproc_ifdef_in_field_declaration_list = 233,
  sym_preproc_else_in_field_declaration_list = 234,
  sym_preproc_elif_in_field_declaration_list = 235,
  sym_preproc_elifdef_in_field_declaration_list = 236,
  sym_preproc_if_in_enumerator_list = 237,
  sym_preproc_ifdef_in_enumerator_list = 238,
  sym_preproc_else_in_enumerator_list = 239,
  sym_preproc_elif_in_enumerator_list = 240,
  sym_preproc_elifdef_in_enumerator_list = 241,
  sym_preproc_if_in_enumerator_list_no_comma = 242,
  sym_preproc_ifdef_in_enumerator_list_no_comma = 243,
  sym_preproc_else_in_enumerator_list_no_comma = 244,
  sym_preproc_elif_in_enumerator_list_no_comma = 245,
  sym_preproc_elifdef_in_enumerator_list_no_comma = 246,
  sym__preproc_expression = 247,
  sym_preproc_parenthesized_expression = 248,
  sym_preproc_defined = 249,
  sym_preproc_unary_expression = 250,
  sym_preproc_call_expression = 251,
  sym_preproc_argument_list = 252,
  sym_preproc_binary_expression = 253,
  sym_function_definition = 254,
  sym_declaration = 255,
  sym_type_definition = 256,
  sym__type_definition_type = 257,
  sym__type_definition_declarators = 258,
  sym__declaration_modifiers = 259,
  sym__declaration_specifiers = 260,
  sym_linkage_specification = 261,
  sym_attribute_specifier = 262,
  sym_attribute = 263,
  sym_attribute_declaration = 264,
  sym_ms_declspec_modifier = 265,
  sym_ms_based_modifier = 266,
  sym_ms_call_modifier = 267,
  sym_ms_unaligned_ptr_modifier = 268,
  sym_ms_pointer_modifier = 269,
  sym_declaration_list = 270,
  sym__declarator = 271,
  sym__field_declarator = 272,
  sym__type_declarator = 273,
  sym__abstract_declarator = 274,
  sym_parenthesized_declarator = 275,
  sym_parenthesized_field_declarator = 276,
  sym_parenthesized_type_declarator = 277,
  sym_abstract_parenthesized_declarator = 278,
  sym_attributed_declarator = 279,
  sym_attributed_field_declarator = 280,
  sym_attributed_type_declarator = 281,
  sym_pointer_declarator = 282,
  sym_pointer_field_declarator = 283,
  sym_pointer_type_declarator = 284,
  sym_abstract_pointer_declarator = 285,
  sym_function_declarator = 286,
  sym_function_field_declarator = 287,
  sym_function_type_declarator = 288,
  sym_abstract_function_declarator = 289,
  sym_array_declarator = 290,
  sym_array_field_declarator = 291,
  sym_array_type_declarator = 292,
  sym_abstract_array_declarator = 293,
  sym_init_declarator = 294,
  sym_compound_statement = 295,
  sym_storage_class_specifier = 296,
  sym_type_qualifier = 297,
  sym_alignas_qualifier = 298,
  sym_type_specifier = 299,
  sym_sized_type_specifier = 300,
  sym_enum_specifier = 301,
  sym_enumerator_list = 302,
  sym_struct_specifier = 303,
  sym_union_specifier = 304,
  sym_field_declaration_list = 305,
  sym__field_declaration_list_item = 306,
  sym_field_declaration = 307,
  sym_bitfield_clause = 308,
  sym_enumerator = 309,
  sym_parameter_list = 310,
  sym_parameter_declaration = 311,
  sym_attributed_statement = 312,
  sym_statement = 313,
  sym__top_level_statement = 314,
  sym_labeled_statement = 315,
  sym__top_level_expression_statement = 316,
  sym_expression_statement = 317,
  sym_if_statement = 318,
  sym_else_clause = 319,
  sym_switch_statement = 320,
  sym_case_statement = 321,
  sym_while_statement = 322,
  sym_do_statement = 323,
  sym_for_statement = 324,
  sym__for_statement_body = 325,
  sym_return_statement = 326,
  sym_break_statement = 327,
  sym_continue_statement = 328,
  sym_goto_statement = 329,
  sym_seh_try_statement = 330,
  sym_seh_except_clause = 331,
  sym_seh_finally_clause = 332,
  sym_seh_leave_statement = 333,
  sym_expression = 334,
  sym__string = 335,
  sym_comma_expression = 336,
  sym_conditional_expression = 337,
  sym_assignment_expression = 338,
  sym_pointer_expression = 339,
  sym_unary_expression = 340,
  sym_binary_expression = 341,
  sym_update_expression = 342,
  sym_cast_expression = 343,
  sym_type_descriptor = 344,
  sym_sizeof_expression = 345,
  sym_alignof_expression = 346,
  sym_offsetof_expression = 347,
  sym_generic_expression = 348,
  sym_subscript_expression = 349,
  sym_call_expression = 350,
  sym_gnu_asm_expression = 351,
  sym_gnu_asm_qualifier = 352,
  sym_gnu_asm_output_operand_list = 353,
  sym_gnu_asm_output_operand = 354,
  sym_gnu_asm_input_operand_list = 355,
  sym_gnu_asm_input_operand = 356,
  sym_gnu_asm_clobber_list = 357,
  sym_gnu_asm_goto_list = 358,
  sym_extension_expression = 359,
  sym_argument_list = 360,
  sym_field_expression = 361,
  sym_compound_literal_expression = 362,
  sym_parenthesized_expression = 363,
  sym_initializer_list = 364,
  sym_initializer_pair = 365,
  sym_subscript_designator = 366,
  sym_subscript_range_designator = 367,
  sym_field_designator = 368,
  sym_char_literal = 369,
  sym_concatenated_string = 370,
  sym_string_literal = 371,
  sym_null = 372,
  sym__empty_declaration = 373,
  sym_placeholder_type_specifier = 374,
  sym_decltype_auto = 375,
  sym_decltype = 376,
  sym__class_declaration = 377,
  sym__class_declaration_item = 378,
  sym_class_specifier = 379,
  sym__class_name = 380,
  sym_virtual_specifier = 381,
  sym_explicit_function_specifier = 382,
  sym_base_class_clause = 383,
  sym__enum_base_clause = 384,
  sym_dependent_type = 385,
  sym_template_declaration = 386,
  sym_template_instantiation = 387,
  sym_template_parameter_list = 388,
  sym_type_parameter_declaration = 389,
  sym_variadic_type_parameter_declaration = 390,
  sym_optional_type_parameter_declaration = 391,
  sym_template_template_parameter_declaration = 392,
  sym_optional_parameter_declaration = 393,
  sym_variadic_parameter_declaration = 394,
  sym_variadic_declarator = 395,
  sym_variadic_reference_declarator = 396,
  sym_operator_cast = 397,
  sym_field_initializer_list = 398,
  sym_field_initializer = 399,
  sym_inline_method_definition = 400,
  sym__constructor_specifiers = 401,
  sym_operator_cast_definition = 402,
  sym_operator_cast_declaration = 403,
  sym_constructor_try_statement = 404,
  sym_constructor_or_destructor_definition = 405,
  sym_constructor_or_destructor_declaration = 406,
  sym_default_method_clause = 407,
  sym_delete_method_clause = 408,
  sym_pure_virtual_clause = 409,
  sym_friend_declaration = 410,
  sym_access_specifier = 411,
  sym_reference_declarator = 412,
  sym_reference_field_declarator = 413,
  sym_reference_type_declarator = 414,
  sym_abstract_reference_declarator = 415,
  sym_structured_binding_declarator = 416,
  sym_ref_qualifier = 417,
  sym__function_declarator_seq = 418,
  sym__function_attributes_start = 419,
  sym__function_exception_specification = 420,
  sym__function_attributes_end = 421,
  sym__function_postfix = 422,
  sym_trailing_return_type = 423,
  sym_noexcept = 424,
  sym_throw_specifier = 425,
  sym_template_type = 426,
  sym_template_method = 427,
  sym_template_function = 428,
  sym_template_argument_list = 429,
  sym_namespace_definition = 430,
  sym_namespace_alias_definition = 431,
  sym__namespace_specifier = 432,
  sym_nested_namespace_specifier = 433,
  sym_using_declaration = 434,
  sym_alias_declaration = 435,
  sym_static_assert_declaration = 436,
  sym_concept_definition = 437,
  sym_for_range_loop = 438,
  sym__for_range_loop_body = 439,
  sym_init_statement = 440,
  sym_condition_clause = 441,
  sym_condition_declaration = 442,
  sym_co_return_statement = 443,
  sym_co_yield_statement = 444,
  sym_throw_statement = 445,
  sym_try_statement = 446,
  sym_catch_clause = 447,
  sym_raw_string_literal = 448,
  sym_subscript_argument_list = 449,
  sym_co_await_expression = 450,
  sym_new_expression = 451,
  sym_new_declarator = 452,
  sym_delete_expression = 453,
  sym_type_requirement = 454,
  sym_compound_requirement = 455,
  sym__requirement = 456,
  sym_requirement_seq = 457,
  sym_constraint_conjunction = 458,
  sym_constraint_disjunction = 459,
  sym__requirement_clause_constraint = 460,
  sym_requires_clause = 461,
  sym_requires_parameter_list = 462,
  sym_requires_expression = 463,
  sym_lambda_expression = 464,
  sym_lambda_capture_specifier = 465,
  sym_lambda_default_capture = 466,
  sym__lambda_capture_identifier = 467,
  sym_lambda_capture_initializer = 468,
  sym__lambda_capture = 469,
  sym__fold_operator = 470,
  sym__binary_fold_operator = 471,
  sym__unary_left_fold = 472,
  sym__unary_right_fold = 473,
  sym__binary_fold = 474,
  sym_fold_expression = 475,
  sym_parameter_pack_expansion = 476,
  sym_type_parameter_pack_expansion = 477,
  sym_identifier_parameter_pack_expansion = 478,
  sym_destructor_name = 479,
  sym_dependent_identifier = 480,
  sym_dependent_field_identifier = 481,
  sym_dependent_type_identifier = 482,
  sym__scope_resolution = 483,
  sym_qualified_field_identifier = 484,
  sym_qualified_identifier = 485,
  sym_qualified_type_identifier = 486,
  sym_qualified_operator_cast_identifier = 487,
  sym__assignment_expression_lhs = 488,
  sym_operator_name = 489,
  sym_user_defined_literal = 490,
  aux_sym_translation_unit_repeat1 = 491,
  aux_sym_preproc_params_repeat1 = 492,
  aux_sym_preproc_if_repeat1 = 493,
  aux_sym_preproc_if_in_field_declaration_list_repeat1 = 494,
  aux_sym_preproc_if_in_enumerator_list_repeat1 = 495,
  aux_sym_preproc_if_in_enumerator_list_no_comma_repeat1 = 496,
  aux_sym_preproc_argument_list_repeat1 = 497,
  aux_sym_declaration_repeat1 = 498,
  aux_sym_type_definition_repeat1 = 499,
  aux_sym__type_definition_type_repeat1 = 500,
  aux_sym__type_definition_declarators_repeat1 = 501,
  aux_sym__declaration_specifiers_repeat1 = 502,
  aux_sym_attribute_declaration_repeat1 = 503,
  aux_sym_attributed_declarator_repeat1 = 504,
  aux_sym_pointer_declarator_repeat1 = 505,
  aux_sym_array_declarator_repeat1 = 506,
  aux_sym_sized_type_specifier_repeat1 = 507,
  aux_sym_enumerator_list_repeat1 = 508,
  aux_sym_field_declaration_repeat1 = 509,
  aux_sym_parameter_list_repeat1 = 510,
  aux_sym_case_statement_repeat1 = 511,
  aux_sym_generic_expression_repeat1 = 512,
  aux_sym_gnu_asm_expression_repeat1 = 513,
  aux_sym_gnu_asm_output_operand_list_repeat1 = 514,
  aux_sym_gnu_asm_input_operand_list_repeat1 = 515,
  aux_sym_gnu_asm_clobber_list_repeat1 = 516,
  aux_sym_gnu_asm_goto_list_repeat1 = 517,
  aux_sym_argument_list_repeat1 = 518,
  aux_sym_initializer_list_repeat1 = 519,
  aux_sym_initializer_pair_repeat1 = 520,
  aux_sym_char_literal_repeat1 = 521,
  aux_sym_concatenated_string_repeat1 = 522,
  aux_sym_string_literal_repeat1 = 523,
  aux_sym__class_declaration_repeat1 = 524,
  aux_sym_base_class_clause_repeat1 = 525,
  aux_sym_template_parameter_list_repeat1 = 526,
  aux_sym_field_initializer_list_repeat1 = 527,
  aux_sym_operator_cast_definition_repeat1 = 528,
  aux_sym_constructor_try_statement_repeat1 = 529,
  aux_sym_structured_binding_declarator_repeat1 = 530,
  aux_sym__function_postfix_repeat1 = 531,
  aux_sym_throw_specifier_repeat1 = 532,
  aux_sym_template_argument_list_repeat1 = 533,
  aux_sym_subscript_argument_list_repeat1 = 534,
  aux_sym_requirement_seq_repeat1 = 535,
  aux_sym_requires_parameter_list_repeat1 = 536,
  aux_sym_lambda_capture_specifier_repeat1 = 537,
  alias_sym_field_identifier = 538,
  alias_sym_namespace_identifier = 539,
  alias_sym_simple_requirement = 540,
  alias_sym_statement_identifier = 541,
  alias_sym_type_identifier = 542,
};

static const char * const ts_symbol_names[] = {
  [ts_builtin_sym_end] = "end",
  [sym_identifier] = "identifier",
  [aux_sym_preproc_include_token1] = "#include",
  [aux_sym_preproc_include_token2] = "preproc_include_token2",
  [aux_sym_preproc_def_token1] = "#define",
  [anon_sym_LPAREN] = "(",
  [anon_sym_DOT_DOT_DOT] = "...",
  [anon_sym_COMMA] = ",",
  [anon_sym_RPAREN] = ")",
  [aux_sym_preproc_if_token1] = "#if",
  [anon_sym_LF] = "\n",
  [aux_sym_preproc_if_token2] = "#endif",
  [aux_sym_preproc_ifdef_token1] = "#ifdef",
  [aux_sym_preproc_ifdef_token2] = "#ifndef",
  [aux_sym_preproc_else_token1] = "#else",
  [aux_sym_preproc_elif_token1] = "#elif",
  [aux_sym_preproc_elifdef_token1] = "#elifdef",
  [aux_sym_preproc_elifdef_token2] = "#elifndef",
  [sym_preproc_arg] = "preproc_arg",
  [sym_preproc_directive] = "preproc_directive",
  [anon_sym_LPAREN2] = "(",
  [anon_sym_defined] = "defined",
  [anon_sym_BANG] = "!",
  [anon_sym_TILDE] = "~",
  [anon_sym_DASH] = "-",
  [anon_sym_PLUS] = "+",
  [anon_sym_STAR] = "*",
  [anon_sym_SLASH] = "/",
  [anon_sym_PERCENT] = "%",
  [anon_sym_PIPE_PIPE] = "||",
  [anon_sym_AMP_AMP] = "&&",
  [anon_sym_PIPE] = "|",
  [anon_sym_CARET] = "^",
  [anon_sym_AMP] = "&",
  [anon_sym_EQ_EQ] = "==",
  [anon_sym_BANG_EQ] = "!=",
  [anon_sym_GT] = ">",
  [anon_sym_GT_EQ] = ">=",
  [anon_sym_LT_EQ] = "<=",
  [anon_sym_LT] = "<",
  [anon_sym_LT_LT] = "<<",
  [anon_sym_GT_GT] = ">>",
  [anon_sym_SEMI] = ";",
  [anon_sym___extension__] = "__extension__",
  [anon_sym_typedef] = "typedef",
  [anon_sym_virtual] = "virtual",
  [anon_sym_extern] = "extern",
  [anon_sym___attribute__] = "__attribute__",
  [anon_sym___attribute] = "__attribute",
  [anon_sym_COLON_COLON] = "::",
  [anon_sym_LBRACK_LBRACK] = "[[",
  [anon_sym_RBRACK_RBRACK] = "]]",
  [anon_sym___declspec] = "__declspec",
  [anon_sym___based] = "__based",
  [anon_sym___cdecl] = "__cdecl",
  [anon_sym___clrcall] = "__clrcall",
  [anon_sym___stdcall] = "__stdcall",
  [anon_sym___fastcall] = "__fastcall",
  [anon_sym___thiscall] = "__thiscall",
  [anon_sym___vectorcall] = "__vectorcall",
  [sym_ms_restrict_modifier] = "ms_restrict_modifier",
  [sym_ms_unsigned_ptr_modifier] = "ms_unsigned_ptr_modifier",
  [sym_ms_signed_ptr_modifier] = "ms_signed_ptr_modifier",
  [anon_sym__unaligned] = "_unaligned",
  [anon_sym___unaligned] = "__unaligned",
  [anon_sym_LBRACE] = "{",
  [anon_sym_RBRACE] = "}",
  [anon_sym_signed] = "signed",
  [anon_sym_unsigned] = "unsigned",
  [anon_sym_long] = "long",
  [anon_sym_short] = "short",
  [anon_sym_LBRACK] = "[",
  [anon_sym_static] = "static",
  [anon_sym_RBRACK] = "]",
  [anon_sym_EQ] = "=",
  [anon_sym_register] = "register",
  [anon_sym_inline] = "inline",
  [anon_sym___inline] = "__inline",
  [anon_sym___inline__] = "__inline__",
  [anon_sym___forceinline] = "__forceinline",
  [anon_sym_thread_local] = "thread_local",
  [anon_sym___thread] = "__thread",
  [anon_sym_const] = "const",
  [anon_sym_constexpr] = "constexpr",
  [anon_sym_volatile] = "volatile",
  [anon_sym_restrict] = "restrict",
  [anon_sym___restrict__] = "__restrict__",
  [anon_sym__Atomic] = "_Atomic",
  [anon_sym__Noreturn] = "_Noreturn",
  [anon_sym_noreturn] = "noreturn",
  [anon_sym__Nonnull] = "_Nonnull",
  [anon_sym_mutable] = "mutable",
  [anon_sym_constinit] = "constinit",
  [anon_sym_consteval] = "consteval",
  [anon_sym_alignas] = "alignas",
  [anon_sym__Alignas] = "_Alignas",
  [sym_primitive_type] = "primitive_type",
  [anon_sym_enum] = "enum",
  [anon_sym_class] = "class",
  [anon_sym_struct] = "struct",
  [anon_sym_union] = "union",
  [anon_sym_COLON] = ":",
  [anon_sym_if] = "if",
  [anon_sym_else] = "else",
  [anon_sym_switch] = "switch",
  [anon_sym_case] = "case",
  [anon_sym_default] = "default",
  [anon_sym_while] = "while",
  [anon_sym_do] = "do",
  [anon_sym_for] = "for",
  [anon_sym_return] = "return",
  [anon_sym_break] = "break",
  [anon_sym_continue] = "continue",
  [anon_sym_goto] = "goto",
  [anon_sym___try] = "__try",
  [anon_sym___except] = "__except",
  [anon_sym___finally] = "__finally",
  [anon_sym___leave] = "__leave",
  [anon_sym_QMARK] = "\?",
  [anon_sym_STAR_EQ] = "*=",
  [anon_sym_SLASH_EQ] = "/=",
  [anon_sym_PERCENT_EQ] = "%=",
  [anon_sym_PLUS_EQ] = "+=",
  [anon_sym_DASH_EQ] = "-=",
  [anon_sym_LT_LT_EQ] = "<<=",
  [anon_sym_GT_GT_EQ] = ">>=",
  [anon_sym_AMP_EQ] = "&=",
  [anon_sym_CARET_EQ] = "^=",
  [anon_sym_PIPE_EQ] = "|=",
  [anon_sym_and_eq] = "and_eq",
  [anon_sym_or_eq] = "or_eq",
  [anon_sym_xor_eq] = "xor_eq",
  [anon_sym_not] = "not",
  [anon_sym_compl] = "compl",
  [anon_sym_LT_EQ_GT] = "<=>",
  [anon_sym_or] = "or",
  [anon_sym_and] = "and",
  [anon_sym_bitor] = "bitor",
  [anon_sym_xor] = "xor",
  [anon_sym_bitand] = "bitand",
  [anon_sym_not_eq] = "not_eq",
  [anon_sym_DASH_DASH] = "--",
  [anon_sym_PLUS_PLUS] = "++",
  [anon_sym_sizeof] = "sizeof",
  [anon_sym___alignof__] = "__alignof__",
  [anon_sym___alignof] = "__alignof",
  [anon_sym__alignof] = "_alignof",
  [anon_sym_alignof] = "alignof",
  [anon_sym__Alignof] = "_Alignof",
  [anon_sym_offsetof] = "offsetof",
  [anon_sym__Generic] = "_Generic",
  [anon_sym_asm] = "asm",
  [anon_sym___asm__] = "__asm__",
  [anon_sym___asm] = "__asm",
  [anon_sym___volatile__] = "__volatile__",
  [anon_sym_DOT] = ".",
  [anon_sym_DOT_STAR] = ".*",
  [anon_sym_DASH_GT] = "->",
  [sym_number_literal] = "number_literal",
  [anon_sym_L_SQUOTE] = "L'",
  [anon_sym_u_SQUOTE] = "u'",
  [anon_sym_U_SQUOTE] = "U'",
  [anon_sym_u8_SQUOTE] = "u8'",
  [anon_sym_SQUOTE] = "'",
  [aux_sym_char_literal_token1] = "character",
  [anon_sym_L_DQUOTE] = "L\"",
  [anon_sym_u_DQUOTE] = "u\"",
  [anon_sym_U_DQUOTE] = "U\"",
  [anon_sym_u8_DQUOTE] = "u8\"",
  [anon_sym_DQUOTE] = "\"",
  [aux_sym_string_literal_token1] = "string_content",
  [sym_escape_sequence] = "escape_sequence",
  [sym_system_lib_string] = "system_lib_string",
  [sym_true] = "true",
  [sym_false] = "false",
  [anon_sym_NULL] = "NULL",
  [anon_sym_nullptr] = "nullptr",
  [sym_comment] = "comment",
  [sym_auto] = "auto",
  [anon_sym_decltype] = "decltype",
  [anon_sym_final] = "final",
  [anon_sym_override] = "override",
  [anon_sym_explicit] = "explicit",
  [anon_sym_typename] = "typename",
  [anon_sym_template] = "template",
  [anon_sym_GT2] = ">",
  [anon_sym_operator] = "operator",
  [anon_sym_try] = "try",
  [anon_sym_delete] = "delete",
  [aux_sym_pure_virtual_clause_token1] = "pure_virtual_clause_token1",
  [anon_sym_friend] = "friend",
  [anon_sym_public] = "public",
  [anon_sym_private] = "private",
  [anon_sym_protected] = "protected",
  [anon_sym_noexcept] = "noexcept",
  [anon_sym_throw] = "throw",
  [anon_sym_namespace] = "namespace",
  [anon_sym_using] = "using",
  [anon_sym_static_assert] = "static_assert",
  [anon_sym_concept] = "concept",
  [anon_sym_co_return] = "co_return",
  [anon_sym_co_yield] = "co_yield",
  [anon_sym_catch] = "catch",
  [anon_sym_R_DQUOTE] = "R\"",
  [anon_sym_LR_DQUOTE] = "LR\"",
  [anon_sym_uR_DQUOTE] = "uR\"",
  [anon_sym_UR_DQUOTE] = "UR\"",
  [anon_sym_u8R_DQUOTE] = "u8R\"",
  [anon_sym_co_await] = "co_await",
  [anon_sym_new] = "new",
  [anon_sym_requires] = "requires",
  [anon_sym_DASH_GT_STAR] = "->*",
  [anon_sym_LPAREN_RPAREN] = "()",
  [anon_sym_LBRACK_RBRACK] = "[]",
  [anon_sym_DQUOTE_DQUOTE] = "\"\"",
  [sym_this] = "this",
  [sym_literal_suffix] = "literal_suffix",
  [sym_raw_string_delimiter] = "raw_string_delimiter",
  [sym_raw_string_content] = "raw_string_content",
  [sym_translation_unit] = "translation_unit",
  [sym__top_level_item] = "_top_level_item",
  [sym__block_item] = "_block_item",
  [sym_preproc_include] = "preproc_include",
  [sym_preproc_def] = "preproc_def",
  [sym_preproc_function_def] = "preproc_function_def",
  [sym_preproc_params] = "preproc_params",
  [sym_preproc_call] = "preproc_call",
  [sym_preproc_if] = "preproc_if",
  [sym_preproc_ifdef] = "preproc_ifdef",
  [sym_preproc_else] = "preproc_else",
  [sym_preproc_elif] = "preproc_elif",
  [sym_preproc_elifdef] = "preproc_elifdef",
  [sym_preproc_if_in_field_declaration_list] = "preproc_if",
  [sym_preproc_ifdef_in_field_declaration_list] = "preproc_ifdef",
  [sym_preproc_else_in_field_declaration_list] = "preproc_else",
  [sym_preproc_elif_in_field_declaration_list] = "preproc_elif",
  [sym_preproc_elifdef_in_field_declaration_list] = "preproc_elifdef",
  [sym_preproc_if_in_enumerator_list] = "preproc_if",
  [sym_preproc_ifdef_in_enumerator_list] = "preproc_ifdef",
  [sym_preproc_else_in_enumerator_list] = "preproc_else",
  [sym_preproc_elif_in_enumerator_list] = "preproc_elif",
  [sym_preproc_elifdef_in_enumerator_list] = "preproc_elifdef",
  [sym_preproc_if_in_enumerator_list_no_comma] = "preproc_if",
  [sym_preproc_ifdef_in_enumerator_list_no_comma] = "preproc_ifdef",
  [sym_preproc_else_in_enumerator_list_no_comma] = "preproc_else",
  [sym_preproc_elif_in_enumerator_list_no_comma] = "preproc_elif",
  [sym_preproc_elifdef_in_enumerator_list_no_comma] = "preproc_elifdef",
  [sym__preproc_expression] = "_preproc_expression",
  [sym_preproc_parenthesized_expression] = "parenthesized_expression",
  [sym_preproc_defined] = "preproc_defined",
  [sym_preproc_unary_expression] = "unary_expression",
  [sym_preproc_call_expression] = "call_expression",
  [sym_preproc_argument_list] = "argument_list",
  [sym_preproc_binary_expression] = "binary_expression",
  [sym_function_definition] = "function_definition",
  [sym_declaration] = "declaration",
  [sym_type_definition] = "type_definition",
  [sym__type_definition_type] = "_type_definition_type",
  [sym__type_definition_declarators] = "_type_definition_declarators",
  [sym__declaration_modifiers] = "_declaration_modifiers",
  [sym__declaration_specifiers] = "_declaration_specifiers",
  [sym_linkage_specification] = "linkage_specification",
  [sym_attribute_specifier] = "attribute_specifier",
  [sym_attribute] = "attribute",
  [sym_attribute_declaration] = "attribute_declaration",
  [sym_ms_declspec_modifier] = "ms_declspec_modifier",
  [sym_ms_based_modifier] = "ms_based_modifier",
  [sym_ms_call_modifier] = "ms_call_modifier",
  [sym_ms_unaligned_ptr_modifier] = "ms_unaligned_ptr_modifier",
  [sym_ms_pointer_modifier] = "ms_pointer_modifier",
  [sym_declaration_list] = "declaration_list",
  [sym__declarator] = "_declarator",
  [sym__field_declarator] = "_field_declarator",
  [sym__type_declarator] = "_type_declarator",
  [sym__abstract_declarator] = "_abstract_declarator",
  [sym_parenthesized_declarator] = "parenthesized_declarator",
  [sym_parenthesized_field_declarator] = "parenthesized_declarator",
  [sym_parenthesized_type_declarator] = "parenthesized_declarator",
  [sym_abstract_parenthesized_declarator] = "abstract_parenthesized_declarator",
  [sym_attributed_declarator] = "attributed_declarator",
  [sym_attributed_field_declarator] = "attributed_declarator",
  [sym_attributed_type_declarator] = "attributed_declarator",
  [sym_pointer_declarator] = "pointer_declarator",
  [sym_pointer_field_declarator] = "pointer_declarator",
  [sym_pointer_type_declarator] = "pointer_type_declarator",
  [sym_abstract_pointer_declarator] = "abstract_pointer_declarator",
  [sym_function_declarator] = "function_declarator",
  [sym_function_field_declarator] = "function_declarator",
  [sym_function_type_declarator] = "function_declarator",
  [sym_abstract_function_declarator] = "abstract_function_declarator",
  [sym_array_declarator] = "array_declarator",
  [sym_array_field_declarator] = "array_declarator",
  [sym_array_type_declarator] = "array_declarator",
  [sym_abstract_array_declarator] = "abstract_array_declarator",
  [sym_init_declarator] = "init_declarator",
  [sym_compound_statement] = "compound_statement",
  [sym_storage_class_specifier] = "storage_class_specifier",
  [sym_type_qualifier] = "type_qualifier",
  [sym_alignas_qualifier] = "alignas_qualifier",
  [sym_type_specifier] = "type_specifier",
  [sym_sized_type_specifier] = "sized_type_specifier",
  [sym_enum_specifier] = "enum_specifier",
  [sym_enumerator_list] = "enumerator_list",
  [sym_struct_specifier] = "struct_specifier",
  [sym_union_specifier] = "union_specifier",
  [sym_field_declaration_list] = "field_declaration_list",
  [sym__field_declaration_list_item] = "_field_declaration_list_item",
  [sym_field_declaration] = "field_declaration",
  [sym_bitfield_clause] = "bitfield_clause",
  [sym_enumerator] = "enumerator",
  [sym_parameter_list] = "parameter_list",
  [sym_parameter_declaration] = "parameter_declaration",
  [sym_attributed_statement] = "attributed_statement",
  [sym_statement] = "statement",
  [sym__top_level_statement] = "_top_level_statement",
  [sym_labeled_statement] = "labeled_statement",
  [sym__top_level_expression_statement] = "expression_statement",
  [sym_expression_statement] = "expression_statement",
  [sym_if_statement] = "if_statement",
  [sym_else_clause] = "else_clause",
  [sym_switch_statement] = "switch_statement",
  [sym_case_statement] = "case_statement",
  [sym_while_statement] = "while_statement",
  [sym_do_statement] = "do_statement",
  [sym_for_statement] = "for_statement",
  [sym__for_statement_body] = "_for_statement_body",
  [sym_return_statement] = "return_statement",
  [sym_break_statement] = "break_statement",
  [sym_continue_statement] = "continue_statement",
  [sym_goto_statement] = "goto_statement",
  [sym_seh_try_statement] = "seh_try_statement",
  [sym_seh_except_clause] = "seh_except_clause",
  [sym_seh_finally_clause] = "seh_finally_clause",
  [sym_seh_leave_statement] = "seh_leave_statement",
  [sym_expression] = "expression",
  [sym__string] = "_string",
  [sym_comma_expression] = "comma_expression",
  [sym_conditional_expression] = "conditional_expression",
  [sym_assignment_expression] = "assignment_expression",
  [sym_pointer_expression] = "pointer_expression",
  [sym_unary_expression] = "unary_expression",
  [sym_binary_expression] = "binary_expression",
  [sym_update_expression] = "update_expression",
  [sym_cast_expression] = "cast_expression",
  [sym_type_descriptor] = "type_descriptor",
  [sym_sizeof_expression] = "sizeof_expression",
  [sym_alignof_expression] = "alignof_expression",
  [sym_offsetof_expression] = "offsetof_expression",
  [sym_generic_expression] = "generic_expression",
  [sym_subscript_expression] = "subscript_expression",
  [sym_call_expression] = "call_expression",
  [sym_gnu_asm_expression] = "gnu_asm_expression",
  [sym_gnu_asm_qualifier] = "gnu_asm_qualifier",
  [sym_gnu_asm_output_operand_list] = "gnu_asm_output_operand_list",
  [sym_gnu_asm_output_operand] = "gnu_asm_output_operand",
  [sym_gnu_asm_input_operand_list] = "gnu_asm_input_operand_list",
  [sym_gnu_asm_input_operand] = "gnu_asm_input_operand",
  [sym_gnu_asm_clobber_list] = "gnu_asm_clobber_list",
  [sym_gnu_asm_goto_list] = "gnu_asm_goto_list",
  [sym_extension_expression] = "extension_expression",
  [sym_argument_list] = "argument_list",
  [sym_field_expression] = "field_expression",
  [sym_compound_literal_expression] = "compound_literal_expression",
  [sym_parenthesized_expression] = "parenthesized_expression",
  [sym_initializer_list] = "initializer_list",
  [sym_initializer_pair] = "initializer_pair",
  [sym_subscript_designator] = "subscript_designator",
  [sym_subscript_range_designator] = "subscript_range_designator",
  [sym_field_designator] = "field_designator",
  [sym_char_literal] = "char_literal",
  [sym_concatenated_string] = "concatenated_string",
  [sym_string_literal] = "string_literal",
  [sym_null] = "null",
  [sym__empty_declaration] = "_empty_declaration",
  [sym_placeholder_type_specifier] = "placeholder_type_specifier",
  [sym_decltype_auto] = "decltype",
  [sym_decltype] = "decltype",
  [sym__class_declaration] = "_class_declaration",
  [sym__class_declaration_item] = "_class_declaration_item",
  [sym_class_specifier] = "class_specifier",
  [sym__class_name] = "_class_name",
  [sym_virtual_specifier] = "virtual_specifier",
  [sym_explicit_function_specifier] = "explicit_function_specifier",
  [sym_base_class_clause] = "base_class_clause",
  [sym__enum_base_clause] = "_enum_base_clause",
  [sym_dependent_type] = "dependent_type",
  [sym_template_declaration] = "template_declaration",
  [sym_template_instantiation] = "template_instantiation",
  [sym_template_parameter_list] = "template_parameter_list",
  [sym_type_parameter_declaration] = "type_parameter_declaration",
  [sym_variadic_type_parameter_declaration] = "variadic_type_parameter_declaration",
  [sym_optional_type_parameter_declaration] = "optional_type_parameter_declaration",
  [sym_template_template_parameter_declaration] = "template_template_parameter_declaration",
  [sym_optional_parameter_declaration] = "optional_parameter_declaration",
  [sym_variadic_parameter_declaration] = "variadic_parameter_declaration",
  [sym_variadic_declarator] = "variadic_declarator",
  [sym_variadic_reference_declarator] = "reference_declarator",
  [sym_operator_cast] = "operator_cast",
  [sym_field_initializer_list] = "field_initializer_list",
  [sym_field_initializer] = "field_initializer",
  [sym_inline_method_definition] = "function_definition",
  [sym__constructor_specifiers] = "_constructor_specifiers",
  [sym_operator_cast_definition] = "function_definition",
  [sym_operator_cast_declaration] = "declaration",
  [sym_constructor_try_statement] = "try_statement",
  [sym_constructor_or_destructor_definition] = "function_definition",
  [sym_constructor_or_destructor_declaration] = "declaration",
  [sym_default_method_clause] = "default_method_clause",
  [sym_delete_method_clause] = "delete_method_clause",
  [sym_pure_virtual_clause] = "pure_virtual_clause",
  [sym_friend_declaration] = "friend_declaration",
  [sym_access_specifier] = "access_specifier",
  [sym_reference_declarator] = "reference_declarator",
  [sym_reference_field_declarator] = "reference_declarator",
  [sym_reference_type_declarator] = "reference_declarator",
  [sym_abstract_reference_declarator] = "abstract_reference_declarator",
  [sym_structured_binding_declarator] = "structured_binding_declarator",
  [sym_ref_qualifier] = "ref_qualifier",
  [sym__function_declarator_seq] = "_function_declarator_seq",
  [sym__function_attributes_start] = "_function_attributes_start",
  [sym__function_exception_specification] = "_function_exception_specification",
  [sym__function_attributes_end] = "_function_attributes_end",
  [sym__function_postfix] = "_function_postfix",
  [sym_trailing_return_type] = "trailing_return_type",
  [sym_noexcept] = "noexcept",
  [sym_throw_specifier] = "throw_specifier",
  [sym_template_type] = "template_type",
  [sym_template_method] = "template_method",
  [sym_template_function] = "template_function",
  [sym_template_argument_list] = "template_argument_list",
  [sym_namespace_definition] = "namespace_definition",
  [sym_namespace_alias_definition] = "namespace_alias_definition",
  [sym__namespace_specifier] = "_namespace_specifier",
  [sym_nested_namespace_specifier] = "nested_namespace_specifier",
  [sym_using_declaration] = "using_declaration",
  [sym_alias_declaration] = "alias_declaration",
  [sym_static_assert_declaration] = "static_assert_declaration",
  [sym_concept_definition] = "concept_definition",
  [sym_for_range_loop] = "for_range_loop",
  [sym__for_range_loop_body] = "_for_range_loop_body",
  [sym_init_statement] = "init_statement",
  [sym_condition_clause] = "condition_clause",
  [sym_condition_declaration] = "declaration",
  [sym_co_return_statement] = "co_return_statement",
  [sym_co_yield_statement] = "co_yield_statement",
  [sym_throw_statement] = "throw_statement",
  [sym_try_statement] = "try_statement",
  [sym_catch_clause] = "catch_clause",
  [sym_raw_string_literal] = "raw_string_literal",
  [sym_subscript_argument_list] = "subscript_argument_list",
  [sym_co_await_expression] = "co_await_expression",
  [sym_new_expression] = "new_expression",
  [sym_new_declarator] = "new_declarator",
  [sym_delete_expression] = "delete_expression",
  [sym_type_requirement] = "type_requirement",
  [sym_compound_requirement] = "compound_requirement",
  [sym__requirement] = "_requirement",
  [sym_requirement_seq] = "requirement_seq",
  [sym_constraint_conjunction] = "constraint_conjunction",
  [sym_constraint_disjunction] = "constraint_disjunction",
  [sym__requirement_clause_constraint] = "_requirement_clause_constraint",
  [sym_requires_clause] = "requires_clause",
  [sym_requires_parameter_list] = "parameter_list",
  [sym_requires_expression] = "requires_expression",
  [sym_lambda_expression] = "lambda_expression",
  [sym_lambda_capture_specifier] = "lambda_capture_specifier",
  [sym_lambda_default_capture] = "lambda_default_capture",
  [sym__lambda_capture_identifier] = "_lambda_capture_identifier",
  [sym_lambda_capture_initializer] = "lambda_capture_initializer",
  [sym__lambda_capture] = "_lambda_capture",
  [sym__fold_operator] = "_fold_operator",
  [sym__binary_fold_operator] = "_binary_fold_operator",
  [sym__unary_left_fold] = "_unary_left_fold",
  [sym__unary_right_fold] = "_unary_right_fold",
  [sym__binary_fold] = "_binary_fold",
  [sym_fold_expression] = "fold_expression",
  [sym_parameter_pack_expansion] = "parameter_pack_expansion",
  [sym_type_parameter_pack_expansion] = "parameter_pack_expansion",
  [sym_identifier_parameter_pack_expansion] = "parameter_pack_expansion",
  [sym_destructor_name] = "destructor_name",
  [sym_dependent_identifier] = "dependent_name",
  [sym_dependent_field_identifier] = "dependent_name",
  [sym_dependent_type_identifier] = "dependent_name",
  [sym__scope_resolution] = "_scope_resolution",
  [sym_qualified_field_identifier] = "qualified_identifier",
  [sym_qualified_identifier] = "qualified_identifier",
  [sym_qualified_type_identifier] = "qualified_identifier",
  [sym_qualified_operator_cast_identifier] = "qualified_identifier",
  [sym__assignment_expression_lhs] = "assignment_expression",
  [sym_operator_name] = "operator_name",
  [sym_user_defined_literal] = "user_defined_literal",
  [aux_sym_translation_unit_repeat1] = "translation_unit_repeat1",
  [aux_sym_preproc_params_repeat1] = "preproc_params_repeat1",
  [aux_sym_preproc_if_repeat1] = "preproc_if_repeat1",
  [aux_sym_preproc_if_in_field_declaration_list_repeat1] = "preproc_if_in_field_declaration_list_repeat1",
  [aux_sym_preproc_if_in_enumerator_list_repeat1] = "preproc_if_in_enumerator_list_repeat1",
  [aux_sym_preproc_if_in_enumerator_list_no_comma_repeat1] = "preproc_if_in_enumerator_list_no_comma_repeat1",
  [aux_sym_preproc_argument_list_repeat1] = "preproc_argument_list_repeat1",
  [aux_sym_declaration_repeat1] = "declaration_repeat1",
  [aux_sym_type_definition_repeat1] = "type_definition_repeat1",
  [aux_sym__type_definition_type_repeat1] = "_type_definition_type_repeat1",
  [aux_sym__type_definition_declarators_repeat1] = "_type_definition_declarators_repeat1",
  [aux_sym__declaration_specifiers_repeat1] = "_declaration_specifiers_repeat1",
  [aux_sym_attribute_declaration_repeat1] = "attribute_declaration_repeat1",
  [aux_sym_attributed_declarator_repeat1] = "attributed_declarator_repeat1",
  [aux_sym_pointer_declarator_repeat1] = "pointer_declarator_repeat1",
  [aux_sym_array_declarator_repeat1] = "array_declarator_repeat1",
  [aux_sym_sized_type_specifier_repeat1] = "sized_type_specifier_repeat1",
  [aux_sym_enumerator_list_repeat1] = "enumerator_list_repeat1",
  [aux_sym_field_declaration_repeat1] = "field_declaration_repeat1",
  [aux_sym_parameter_list_repeat1] = "parameter_list_repeat1",
  [aux_sym_case_statement_repeat1] = "case_statement_repeat1",
  [aux_sym_generic_expression_repeat1] = "generic_expression_repeat1",
  [aux_sym_gnu_asm_expression_repeat1] = "gnu_asm_expression_repeat1",
  [aux_sym_gnu_asm_output_operand_list_repeat1] = "gnu_asm_output_operand_list_repeat1",
  [aux_sym_gnu_asm_input_operand_list_repeat1] = "gnu_asm_input_operand_list_repeat1",
  [aux_sym_gnu_asm_clobber_list_repeat1] = "gnu_asm_clobber_list_repeat1",
  [aux_sym_gnu_asm_goto_list_repeat1] = "gnu_asm_goto_list_repeat1",
  [aux_sym_argument_list_repeat1] = "argument_list_repeat1",
  [aux_sym_initializer_list_repeat1] = "initializer_list_repeat1",
  [aux_sym_initializer_pair_repeat1] = "initializer_pair_repeat1",
  [aux_sym_char_literal_repeat1] = "char_literal_repeat1",
  [aux_sym_concatenated_string_repeat1] = "concatenated_string_repeat1",
  [aux_sym_string_literal_repeat1] = "string_literal_repeat1",
  [aux_sym__class_declaration_repeat1] = "_class_declaration_repeat1",
  [aux_sym_base_class_clause_repeat1] = "base_class_clause_repeat1",
  [aux_sym_template_parameter_list_repeat1] = "template_parameter_list_repeat1",
  [aux_sym_field_initializer_list_repeat1] = "field_initializer_list_repeat1",
  [aux_sym_operator_cast_definition_repeat1] = "operator_cast_definition_repeat1",
  [aux_sym_constructor_try_statement_repeat1] = "constructor_try_statement_repeat1",
  [aux_sym_structured_binding_declarator_repeat1] = "structured_binding_declarator_repeat1",
  [aux_sym__function_postfix_repeat1] = "_function_postfix_repeat1",
  [aux_sym_throw_specifier_repeat1] = "throw_specifier_repeat1",
  [aux_sym_template_argument_list_repeat1] = "template_argument_list_repeat1",
  [aux_sym_subscript_argument_list_repeat1] = "subscript_argument_list_repeat1",
  [aux_sym_requirement_seq_repeat1] = "requirement_seq_repeat1",
  [aux_sym_requires_parameter_list_repeat1] = "requires_parameter_list_repeat1",
  [aux_sym_lambda_capture_specifier_repeat1] = "lambda_capture_specifier_repeat1",
  [alias_sym_field_identifier] = "field_identifier",
  [alias_sym_namespace_identifier] = "namespace_identifier",
  [alias_sym_simple_requirement] = "simple_requirement",
  [alias_sym_statement_identifier] = "statement_identifier",
  [alias_sym_type_identifier] = "type_identifier",
};

static const TSSymbol ts_symbol_map[] = {
  [ts_builtin_sym_end] = ts_builtin_sym_end,
  [sym_identifier] = sym_identifier,
  [aux_sym_preproc_include_token1] = aux_sym_preproc_include_token1,
  [aux_sym_preproc_include_token2] = aux_sym_preproc_include_token2,
  [aux_sym_preproc_def_token1] = aux_sym_preproc_def_token1,
  [anon_sym_LPAREN] = anon_sym_LPAREN,
  [anon_sym_DOT_DOT_DOT] = anon_sym_DOT_DOT_DOT,
  [anon_sym_COMMA] = anon_sym_COMMA,
  [anon_sym_RPAREN] = anon_sym_RPAREN,
  [aux_sym_preproc_if_token1] = aux_sym_preproc_if_token1,
  [anon_sym_LF] = anon_sym_LF,
  [aux_sym_preproc_if_token2] = aux_sym_preproc_if_token2,
  [aux_sym_preproc_ifdef_token1] = aux_sym_preproc_ifdef_token1,
  [aux_sym_preproc_ifdef_token2] = aux_sym_preproc_ifdef_token2,
  [aux_sym_preproc_else_token1] = aux_sym_preproc_else_token1,
  [aux_sym_preproc_elif_token1] = aux_sym_preproc_elif_token1,
  [aux_sym_preproc_elifdef_token1] = aux_sym_preproc_elifdef_token1,
  [aux_sym_preproc_elifdef_token2] = aux_sym_preproc_elifdef_token2,
  [sym_preproc_arg] = sym_preproc_arg,
  [sym_preproc_directive] = sym_preproc_directive,
  [anon_sym_LPAREN2] = anon_sym_LPAREN,
  [anon_sym_defined] = anon_sym_defined,
  [anon_sym_BANG] = anon_sym_BANG,
  [anon_sym_TILDE] = anon_sym_TILDE,
  [anon_sym_DASH] = anon_sym_DASH,
  [anon_sym_PLUS] = anon_sym_PLUS,
  [anon_sym_STAR] = anon_sym_STAR,
  [anon_sym_SLASH] = anon_sym_SLASH,
  [anon_sym_PERCENT] = anon_sym_PERCENT,
  [anon_sym_PIPE_PIPE] = anon_sym_PIPE_PIPE,
  [anon_sym_AMP_AMP] = anon_sym_AMP_AMP,
  [anon_sym_PIPE] = anon_sym_PIPE,
  [anon_sym_CARET] = anon_sym_CARET,
  [anon_sym_AMP] = anon_sym_AMP,
  [anon_sym_EQ_EQ] = anon_sym_EQ_EQ,
  [anon_sym_BANG_EQ] = anon_sym_BANG_EQ,
  [anon_sym_GT] = anon_sym_GT,
  [anon_sym_GT_EQ] = anon_sym_GT_EQ,
  [anon_sym_LT_EQ] = anon_sym_LT_EQ,
  [anon_sym_LT] = anon_sym_LT,
  [anon_sym_LT_LT] = anon_sym_LT_LT,
  [anon_sym_GT_GT] = anon_sym_GT_GT,
  [anon_sym_SEMI] = anon_sym_SEMI,
  [anon_sym___extension__] = anon_sym___extension__,
  [anon_sym_typedef] = anon_sym_typedef,
  [anon_sym_virtual] = anon_sym_virtual,
  [anon_sym_extern] = anon_sym_extern,
  [anon_sym___attribute__] = anon_sym___attribute__,
  [anon_sym___attribute] = anon_sym___attribute,
  [anon_sym_COLON_COLON] = anon_sym_COLON_COLON,
  [anon_sym_LBRACK_LBRACK] = anon_sym_LBRACK_LBRACK,
  [anon_sym_RBRACK_RBRACK] = anon_sym_RBRACK_RBRACK,
  [anon_sym___declspec] = anon_sym___declspec,
  [anon_sym___based] = anon_sym___based,
  [anon_sym___cdecl] = anon_sym___cdecl,
  [anon_sym___clrcall] = anon_sym___clrcall,
  [anon_sym___stdcall] = anon_sym___stdcall,
  [anon_sym___fastcall] = anon_sym___fastcall,
  [anon_sym___thiscall] = anon_sym___thiscall,
  [anon_sym___vectorcall] = anon_sym___vectorcall,
  [sym_ms_restrict_modifier] = sym_ms_restrict_modifier,
  [sym_ms_unsigned_ptr_modifier] = sym_ms_unsigned_ptr_modifier,
  [sym_ms_signed_ptr_modifier] = sym_ms_signed_ptr_modifier,
  [anon_sym__unaligned] = anon_sym__unaligned,
  [anon_sym___unaligned] = anon_sym___unaligned,
  [anon_sym_LBRACE] = anon_sym_LBRACE,
  [anon_sym_RBRACE] = anon_sym_RBRACE,
  [anon_sym_signed] = anon_sym_signed,
  [anon_sym_unsigned] = anon_sym_unsigned,
  [anon_sym_long] = anon_sym_long,
  [anon_sym_short] = anon_sym_short,
  [anon_sym_LBRACK] = anon_sym_LBRACK,
  [anon_sym_static] = anon_sym_static,
  [anon_sym_RBRACK] = anon_sym_RBRACK,
  [anon_sym_EQ] = anon_sym_EQ,
  [anon_sym_register] = anon_sym_register,
  [anon_sym_inline] = anon_sym_inline,
  [anon_sym___inline] = anon_sym___inline,
  [anon_sym___inline__] = anon_sym___inline__,
  [anon_sym___forceinline] = anon_sym___forceinline,
  [anon_sym_thread_local] = anon_sym_thread_local,
  [anon_sym___thread] = anon_sym___thread,
  [anon_sym_const] = anon_sym_const,
  [anon_sym_constexpr] = anon_sym_constexpr,
  [anon_sym_volatile] = anon_sym_volatile,
  [anon_sym_restrict] = anon_sym_restrict,
  [anon_sym___restrict__] = anon_sym___restrict__,
  [anon_sym__Atomic] = anon_sym__Atomic,
  [anon_sym__Noreturn] = anon_sym__Noreturn,
  [anon_sym_noreturn] = anon_sym_noreturn,
  [anon_sym__Nonnull] = anon_sym__Nonnull,
  [anon_sym_mutable] = anon_sym_mutable,
  [anon_sym_constinit] = anon_sym_constinit,
  [anon_sym_consteval] = anon_sym_consteval,
  [anon_sym_alignas] = anon_sym_alignas,
  [anon_sym__Alignas] = anon_sym__Alignas,
  [sym_primitive_type] = sym_primitive_type,
  [anon_sym_enum] = anon_sym_enum,
  [anon_sym_class] = anon_sym_class,
  [anon_sym_struct] = anon_sym_struct,
  [anon_sym_union] = anon_sym_union,
  [anon_sym_COLON] = anon_sym_COLON,
  [anon_sym_if] = anon_sym_if,
  [anon_sym_else] = anon_sym_else,
  [anon_sym_switch] = anon_sym_switch,
  [anon_sym_case] = anon_sym_case,
  [anon_sym_default] = anon_sym_default,
  [anon_sym_while] = anon_sym_while,
  [anon_sym_do] = anon_sym_do,
  [anon_sym_for] = anon_sym_for,
  [anon_sym_return] = anon_sym_return,
  [anon_sym_break] = anon_sym_break,
  [anon_sym_continue] = anon_sym_continue,
  [anon_sym_goto] = anon_sym_goto,
  [anon_sym___try] = anon_sym___try,
  [anon_sym___except] = anon_sym___except,
  [anon_sym___finally] = anon_sym___finally,
  [anon_sym___leave] = anon_sym___leave,
  [anon_sym_QMARK] = anon_sym_QMARK,
  [anon_sym_STAR_EQ] = anon_sym_STAR_EQ,
  [anon_sym_SLASH_EQ] = anon_sym_SLASH_EQ,
  [anon_sym_PERCENT_EQ] = anon_sym_PERCENT_EQ,
  [anon_sym_PLUS_EQ] = anon_sym_PLUS_EQ,
  [anon_sym_DASH_EQ] = anon_sym_DASH_EQ,
  [anon_sym_LT_LT_EQ] = anon_sym_LT_LT_EQ,
  [anon_sym_GT_GT_EQ] = anon_sym_GT_GT_EQ,
  [anon_sym_AMP_EQ] = anon_sym_AMP_EQ,
  [anon_sym_CARET_EQ] = anon_sym_CARET_EQ,
  [anon_sym_PIPE_EQ] = anon_sym_PIPE_EQ,
  [anon_sym_and_eq] = anon_sym_and_eq,
  [anon_sym_or_eq] = anon_sym_or_eq,
  [anon_sym_xor_eq] = anon_sym_xor_eq,
  [anon_sym_not] = anon_sym_not,
  [anon_sym_compl] = anon_sym_compl,
  [anon_sym_LT_EQ_GT] = anon_sym_LT_EQ_GT,
  [anon_sym_or] = anon_sym_or,
  [anon_sym_and] = anon_sym_and,
  [anon_sym_bitor] = anon_sym_bitor,
  [anon_sym_xor] = anon_sym_xor,
  [anon_sym_bitand] = anon_sym_bitand,
  [anon_sym_not_eq] = anon_sym_not_eq,
  [anon_sym_DASH_DASH] = anon_sym_DASH_DASH,
  [anon_sym_PLUS_PLUS] = anon_sym_PLUS_PLUS,
  [anon_sym_sizeof] = anon_sym_sizeof,
  [anon_sym___alignof__] = anon_sym___alignof__,
  [anon_sym___alignof] = anon_sym___alignof,
  [anon_sym__alignof] = anon_sym__alignof,
  [anon_sym_alignof] = anon_sym_alignof,
  [anon_sym__Alignof] = anon_sym__Alignof,
  [anon_sym_offsetof] = anon_sym_offsetof,
  [anon_sym__Generic] = anon_sym__Generic,
  [anon_sym_asm] = anon_sym_asm,
  [anon_sym___asm__] = anon_sym___asm__,
  [anon_sym___asm] = anon_sym___asm,
  [anon_sym___volatile__] = anon_sym___volatile__,
  [anon_sym_DOT] = anon_sym_DOT,
  [anon_sym_DOT_STAR] = anon_sym_DOT_STAR,
  [anon_sym_DASH_GT] = anon_sym_DASH_GT,
  [sym_number_literal] = sym_number_literal,
  [anon_sym_L_SQUOTE] = anon_sym_L_SQUOTE,
  [anon_sym_u_SQUOTE] = anon_sym_u_SQUOTE,
  [anon_sym_U_SQUOTE] = anon_sym_U_SQUOTE,
  [anon_sym_u8_SQUOTE] = anon_sym_u8_SQUOTE,
  [anon_sym_SQUOTE] = anon_sym_SQUOTE,
  [aux_sym_char_literal_token1] = aux_sym_char_literal_token1,
  [anon_sym_L_DQUOTE] = anon_sym_L_DQUOTE,
  [anon_sym_u_DQUOTE] = anon_sym_u_DQUOTE,
  [anon_sym_U_DQUOTE] = anon_sym_U_DQUOTE,
  [anon_sym_u8_DQUOTE] = anon_sym_u8_DQUOTE,
  [anon_sym_DQUOTE] = anon_sym_DQUOTE,
  [aux_sym_string_literal_token1] = aux_sym_string_literal_token1,
  [sym_escape_sequence] = sym_escape_sequence,
  [sym_system_lib_string] = sym_system_lib_string,
  [sym_true] = sym_true,
  [sym_false] = sym_false,
  [anon_sym_NULL] = anon_sym_NULL,
  [anon_sym_nullptr] = anon_sym_nullptr,
  [sym_comment] = sym_comment,
  [sym_auto] = sym_auto,
  [anon_sym_decltype] = anon_sym_decltype,
  [anon_sym_final] = anon_sym_final,
  [anon_sym_override] = anon_sym_override,
  [anon_sym_explicit] = anon_sym_explicit,
  [anon_sym_typename] = anon_sym_typename,
  [anon_sym_template] = anon_sym_template,
  [anon_sym_GT2] = anon_sym_GT,
  [anon_sym_operator] = anon_sym_operator,
  [anon_sym_try] = anon_sym_try,
  [anon_sym_delete] = anon_sym_delete,
  [aux_sym_pure_virtual_clause_token1] = aux_sym_pure_virtual_clause_token1,
  [anon_sym_friend] = anon_sym_friend,
  [anon_sym_public] = anon_sym_public,
  [anon_sym_private] = anon_sym_private,
  [anon_sym_protected] = anon_sym_protected,
  [anon_sym_noexcept] = anon_sym_noexcept,
  [anon_sym_throw] = anon_sym_throw,
  [anon_sym_namespace] = anon_sym_namespace,
  [anon_sym_using] = anon_sym_using,
  [anon_sym_static_assert] = anon_sym_static_assert,
  [anon_sym_concept] = anon_sym_concept,
  [anon_sym_co_return] = anon_sym_co_return,
  [anon_sym_co_yield] = anon_sym_co_yield,
  [anon_sym_catch] = anon_sym_catch,
  [anon_sym_R_DQUOTE] = anon_sym_R_DQUOTE,
  [anon_sym_LR_DQUOTE] = anon_sym_LR_DQUOTE,
  [anon_sym_uR_DQUOTE] = anon_sym_uR_DQUOTE,
  [anon_sym_UR_DQUOTE] = anon_sym_UR_DQUOTE,
  [anon_sym_u8R_DQUOTE] = anon_sym_u8R_DQUOTE,
  [anon_sym_co_await] = anon_sym_co_await,
  [anon_sym_new] = anon_sym_new,
  [anon_sym_requires] = anon_sym_requires,
  [anon_sym_DASH_GT_STAR] = anon_sym_DASH_GT_STAR,
  [anon_sym_LPAREN_RPAREN] = anon_sym_LPAREN_RPAREN,
  [anon_sym_LBRACK_RBRACK] = anon_sym_LBRACK_RBRACK,
  [anon_sym_DQUOTE_DQUOTE] = anon_sym_DQUOTE_DQUOTE,
  [sym_this] = sym_this,
  [sym_literal_suffix] = sym_literal_suffix,
  [sym_raw_string_delimiter] = sym_raw_string_delimiter,
  [sym_raw_string_content] = sym_raw_string_content,
  [sym_translation_unit] = sym_translation_unit,
  [sym__top_level_item] = sym__top_level_item,
  [sym__block_item] = sym__block_item,
  [sym_preproc_include] = sym_preproc_include,
  [sym_preproc_def] = sym_preproc_def,
  [sym_preproc_function_def] = sym_preproc_function_def,
  [sym_preproc_params] = sym_preproc_params,
  [sym_preproc_call] = sym_preproc_call,
  [sym_preproc_if] = sym_preproc_if,
  [sym_preproc_ifdef] = sym_preproc_ifdef,
  [sym_preproc_else] = sym_preproc_else,
  [sym_preproc_elif] = sym_preproc_elif,
  [sym_preproc_elifdef] = sym_preproc_elifdef,
  [sym_preproc_if_in_field_declaration_list] = sym_preproc_if,
  [sym_preproc_ifdef_in_field_declaration_list] = sym_preproc_ifdef,
  [sym_preproc_else_in_field_declaration_list] = sym_preproc_else,
  [sym_preproc_elif_in_field_declaration_list] = sym_preproc_elif,
  [sym_preproc_elifdef_in_field_declaration_list] = sym_preproc_elifdef,
  [sym_preproc_if_in_enumerator_list] = sym_preproc_if,
  [sym_preproc_ifdef_in_enumerator_list] = sym_preproc_ifdef,
  [sym_preproc_else_in_enumerator_list] = sym_preproc_else,
  [sym_preproc_elif_in_enumerator_list] = sym_preproc_elif,
  [sym_preproc_elifdef_in_enumerator_list] = sym_preproc_elifdef,
  [sym_preproc_if_in_enumerator_list_no_comma] = sym_preproc_if,
  [sym_preproc_ifdef_in_enumerator_list_no_comma] = sym_preproc_ifdef,
  [sym_preproc_else_in_enumerator_list_no_comma] = sym_preproc_else,
  [sym_preproc_elif_in_enumerator_list_no_comma] = sym_preproc_elif,
  [sym_preproc_elifdef_in_enumerator_list_no_comma] = sym_preproc_elifdef,
  [sym__preproc_expression] = sym__preproc_expression,
  [sym_preproc_parenthesized_expression] = sym_parenthesized_expression,
  [sym_preproc_defined] = sym_preproc_defined,
  [sym_preproc_unary_expression] = sym_unary_expression,
  [sym_preproc_call_expression] = sym_call_expression,
  [sym_preproc_argument_list] = sym_argument_list,
  [sym_preproc_binary_expression] = sym_binary_expression,
  [sym_function_definition] = sym_function_definition,
  [sym_declaration] = sym_declaration,
  [sym_type_definition] = sym_type_definition,
  [sym__type_definition_type] = sym__type_definition_type,
  [sym__type_definition_declarators] = sym__type_definition_declarators,
  [sym__declaration_modifiers] = sym__declaration_modifiers,
  [sym__declaration_specifiers] = sym__declaration_specifiers,
  [sym_linkage_specification] = sym_linkage_specification,
  [sym_attribute_specifier] = sym_attribute_specifier,
  [sym_attribute] = sym_attribute,
  [sym_attribute_declaration] = sym_attribute_declaration,
  [sym_ms_declspec_modifier] = sym_ms_declspec_modifier,
  [sym_ms_based_modifier] = sym_ms_based_modifier,
  [sym_ms_call_modifier] = sym_ms_call_modifier,
  [sym_ms_unaligned_ptr_modifier] = sym_ms_unaligned_ptr_modifier,
  [sym_ms_pointer_modifier] = sym_ms_pointer_modifier,
  [sym_declaration_list] = sym_declaration_list,
  [sym__declarator] = sym__declarator,
  [sym__field_declarator] = sym__field_declarator,
  [sym__type_declarator] = sym__type_declarator,
  [sym__abstract_declarator] = sym__abstract_declarator,
  [sym_parenthesized_declarator] = sym_parenthesized_declarator,
  [sym_parenthesized_field_declarator] = sym_parenthesized_declarator,
  [sym_parenthesized_type_declarator] = sym_parenthesized_declarator,
  [sym_abstract_parenthesized_declarator] = sym_abstract_parenthesized_declarator,
  [sym_attributed_declarator] = sym_attributed_declarator,
  [sym_attributed_field_declarator] = sym_attributed_declarator,
  [sym_attributed_type_declarator] = sym_attributed_declarator,
  [sym_pointer_declarator] = sym_pointer_declarator,
  [sym_pointer_field_declarator] = sym_pointer_declarator,
  [sym_pointer_type_declarator] = sym_pointer_type_declarator,
  [sym_abstract_pointer_declarator] = sym_abstract_pointer_declarator,
  [sym_function_declarator] = sym_function_declarator,
  [sym_function_field_declarator] = sym_function_declarator,
  [sym_function_type_declarator] = sym_function_declarator,
  [sym_abstract_function_declarator] = sym_abstract_function_declarator,
  [sym_array_declarator] = sym_array_declarator,
  [sym_array_field_declarator] = sym_array_declarator,
  [sym_array_type_declarator] = sym_array_declarator,
  [sym_abstract_array_declarator] = sym_abstract_array_declarator,
  [sym_init_declarator] = sym_init_declarator,
  [sym_compound_statement] = sym_compound_statement,
  [sym_storage_class_specifier] = sym_storage_class_specifier,
  [sym_type_qualifier] = sym_type_qualifier,
  [sym_alignas_qualifier] = sym_alignas_qualifier,
  [sym_type_specifier] = sym_type_specifier,
  [sym_sized_type_specifier] = sym_sized_type_specifier,
  [sym_enum_specifier] = sym_enum_specifier,
  [sym_enumerator_list] = sym_enumerator_list,
  [sym_struct_specifier] = sym_struct_specifier,
  [sym_union_specifier] = sym_union_specifier,
  [sym_field_declaration_list] = sym_field_declaration_list,
  [sym__field_declaration_list_item] = sym__field_declaration_list_item,
  [sym_field_declaration] = sym_field_declaration,
  [sym_bitfield_clause] = sym_bitfield_clause,
  [sym_enumerator] = sym_enumerator,
  [sym_parameter_list] = sym_parameter_list,
  [sym_parameter_declaration] = sym_parameter_declaration,
  [sym_attributed_statement] = sym_attributed_statement,
  [sym_statement] = sym_statement,
  [sym__top_level_statement] = sym__top_level_statement,
  [sym_labeled_statement] = sym_labeled_statement,
  [sym__top_level_expression_statement] = sym_expression_statement,
  [sym_expression_statement] = sym_expression_statement,
  [sym_if_statement] = sym_if_statement,
  [sym_else_clause] = sym_else_clause,
  [sym_switch_statement] = sym_switch_statement,
  [sym_case_statement] = sym_case_statement,
  [sym_while_statement] = sym_while_statement,
  [sym_do_statement] = sym_do_statement,
  [sym_for_statement] = sym_for_statement,
  [sym__for_statement_body] = sym__for_statement_body,
  [sym_return_statement] = sym_return_statement,
  [sym_break_statement] = sym_break_statement,
  [sym_continue_statement] = sym_continue_statement,
  [sym_goto_statement] = sym_goto_statement,
  [sym_seh_try_statement] = sym_seh_try_statement,
  [sym_seh_except_clause] = sym_seh_except_clause,
  [sym_seh_finally_clause] = sym_seh_finally_clause,
  [sym_seh_leave_statement] = sym_seh_leave_statement,
  [sym_expression] = sym_expression,
  [sym__string] = sym__string,
  [sym_comma_expression] = sym_comma_expression,
  [sym_conditional_expression] = sym_conditional_expression,
  [sym_assignment_expression] = sym_assignment_expression,
  [sym_pointer_expression] = sym_pointer_expression,
  [sym_unary_expression] = sym_unary_expression,
  [sym_binary_expression] = sym_binary_expression,
  [sym_update_expression] = sym_update_expression,
  [sym_cast_expression] = sym_cast_expression,
  [sym_type_descriptor] = sym_type_descriptor,
  [sym_sizeof_expression] = sym_sizeof_expression,
  [sym_alignof_expression] = sym_alignof_expression,
  [sym_offsetof_expression] = sym_offsetof_expression,
  [sym_generic_expression] = sym_generic_expression,
  [sym_subscript_expression] = sym_subscript_expression,
  [sym_call_expression] = sym_call_expression,
  [sym_gnu_asm_expression] = sym_gnu_asm_expression,
  [sym_gnu_asm_qualifier] = sym_gnu_asm_qualifier,
  [sym_gnu_asm_output_operand_list] = sym_gnu_asm_output_operand_list,
  [sym_gnu_asm_output_operand] = sym_gnu_asm_output_operand,
  [sym_gnu_asm_input_operand_list] = sym_gnu_asm_input_operand_list,
  [sym_gnu_asm_input_operand] = sym_gnu_asm_input_operand,
  [sym_gnu_asm_clobber_list] = sym_gnu_asm_clobber_list,
  [sym_gnu_asm_goto_list] = sym_gnu_asm_goto_list,
  [sym_extension_expression] = sym_extension_expression,
  [sym_argument_list] = sym_argument_list,
  [sym_field_expression] = sym_field_expression,
  [sym_compound_literal_expression] = sym_compound_literal_expression,
  [sym_parenthesized_expression] = sym_parenthesized_expression,
  [sym_initializer_list] = sym_initializer_list,
  [sym_initializer_pair] = sym_initializer_pair,
  [sym_subscript_designator] = sym_subscript_designator,
  [sym_subscript_range_designator] = sym_subscript_range_designator,
  [sym_field_designator] = sym_field_designator,
  [sym_char_literal] = sym_char_literal,
  [sym_concatenated_string] = sym_concatenated_string,
  [sym_string_literal] = sym_string_literal,
  [sym_null] = sym_null,
  [sym__empty_declaration] = sym__empty_declaration,
  [sym_placeholder_type_specifier] = sym_placeholder_type_specifier,
  [sym_decltype_auto] = sym_decltype,
  [sym_decltype] = sym_decltype,
  [sym__class_declaration] = sym__class_declaration,
  [sym__class_declaration_item] = sym__class_declaration_item,
  [sym_class_specifier] = sym_class_specifier,
  [sym__class_name] = sym__class_name,
  [sym_virtual_specifier] = sym_virtual_specifier,
  [sym_explicit_function_specifier] = sym_explicit_function_specifier,
  [sym_base_class_clause] = sym_base_class_clause,
  [sym__enum_base_clause] = sym__enum_base_clause,
  [sym_dependent_type] = sym_dependent_type,
  [sym_template_declaration] = sym_template_declaration,
  [sym_template_instantiation] = sym_template_instantiation,
  [sym_template_parameter_list] = sym_template_parameter_list,
  [sym_type_parameter_declaration] = sym_type_parameter_declaration,
  [sym_variadic_type_parameter_declaration] = sym_variadic_type_parameter_declaration,
  [sym_optional_type_parameter_declaration] = sym_optional_type_parameter_declaration,
  [sym_template_template_parameter_declaration] = sym_template_template_parameter_declaration,
  [sym_optional_parameter_declaration] = sym_optional_parameter_declaration,
  [sym_variadic_parameter_declaration] = sym_variadic_parameter_declaration,
  [sym_variadic_declarator] = sym_variadic_declarator,
  [sym_variadic_reference_declarator] = sym_reference_declarator,
  [sym_operator_cast] = sym_operator_cast,
  [sym_field_initializer_list] = sym_field_initializer_list,
  [sym_field_initializer] = sym_field_initializer,
  [sym_inline_method_definition] = sym_function_definition,
  [sym__constructor_specifiers] = sym__constructor_specifiers,
  [sym_operator_cast_definition] = sym_function_definition,
  [sym_operator_cast_declaration] = sym_declaration,
  [sym_constructor_try_statement] = sym_try_statement,
  [sym_constructor_or_destructor_definition] = sym_function_definition,
  [sym_constructor_or_destructor_declaration] = sym_declaration,
  [sym_default_method_clause] = sym_default_method_clause,
  [sym_delete_method_clause] = sym_delete_method_clause,
  [sym_pure_virtual_clause] = sym_pure_virtual_clause,
  [sym_friend_declaration] = sym_friend_declaration,
  [sym_access_specifier] = sym_access_specifier,
  [sym_reference_declarator] = sym_reference_declarator,
  [sym_reference_field_declarator] = sym_reference_declarator,
  [sym_reference_type_declarator] = sym_reference_declarator,
  [sym_abstract_reference_declarator] = sym_abstract_reference_declarator,
  [sym_structured_binding_declarator] = sym_structured_binding_declarator,
  [sym_ref_qualifier] = sym_ref_qualifier,
  [sym__function_declarator_seq] = sym__function_declarator_seq,
  [sym__function_attributes_start] = sym__function_attributes_start,
  [sym__function_exception_specification] = sym__function_exception_specification,
  [sym__function_attributes_end] = sym__function_attributes_end,
  [sym__function_postfix] = sym__function_postfix,
  [sym_trailing_return_type] = sym_trailing_return_type,
  [sym_noexcept] = sym_noexcept,
  [sym_throw_specifier] = sym_throw_specifier,
  [sym_template_type] = sym_template_type,
  [sym_template_method] = sym_template_method,
  [sym_template_function] = sym_template_function,
  [sym_template_argument_list] = sym_template_argument_list,
  [sym_namespace_definition] = sym_namespace_definition,
  [sym_namespace_alias_definition] = sym_namespace_alias_definition,
  [sym__namespace_specifier] = sym__namespace_specifier,
  [sym_nested_namespace_specifier] = sym_nested_namespace_specifier,
  [sym_using_declaration] = sym_using_declaration,
  [sym_alias_declaration] = sym_alias_declaration,
  [sym_static_assert_declaration] = sym_static_assert_declaration,
  [sym_concept_definition] = sym_concept_definition,
  [sym_for_range_loop] = sym_for_range_loop,
  [sym__for_range_loop_body] = sym__for_range_loop_body,
  [sym_init_statement] = sym_init_statement,
  [sym_condition_clause] = sym_condition_clause,
  [sym_condition_declaration] = sym_declaration,
  [sym_co_return_statement] = sym_co_return_statement,
  [sym_co_yield_statement] = sym_co_yield_statement,
  [sym_throw_statement] = sym_throw_statement,
  [sym_try_statement] = sym_try_statement,
  [sym_catch_clause] = sym_catch_clause,
  [sym_raw_string_literal] = sym_raw_string_literal,
  [sym_subscript_argument_list] = sym_subscript_argument_list,
  [sym_co_await_expression] = sym_co_await_expression,
  [sym_new_expression] = sym_new_expression,
  [sym_new_declarator] = sym_new_declarator,
  [sym_delete_expression] = sym_delete_expression,
  [sym_type_requirement] = sym_type_requirement,
  [sym_compound_requirement] = sym_compound_requirement,
  [sym__requirement] = sym__requirement,
  [sym_requirement_seq] = sym_requirement_seq,
  [sym_constraint_conjunction] = sym_constraint_conjunction,
  [sym_constraint_disjunction] = sym_constraint_disjunction,
  [sym__requirement_clause_constraint] = sym__requirement_clause_constraint,
  [sym_requires_clause] = sym_requires_clause,
  [sym_requires_parameter_list] = sym_parameter_list,
  [sym_requires_expression] = sym_requires_expression,
  [sym_lambda_expression] = sym_lambda_expression,
  [sym_lambda_capture_specifier] = sym_lambda_capture_specifier,
  [sym_lambda_default_capture] = sym_lambda_default_capture,
  [sym__lambda_capture_identifier] = sym__lambda_capture_identifier,
  [sym_lambda_capture_initializer] = sym_lambda_capture_initializer,
  [sym__lambda_capture] = sym__lambda_capture,
  [sym__fold_operator] = sym__fold_operator,
  [sym__binary_fold_operator] = sym__binary_fold_operator,
  [sym__unary_left_fold] = sym__unary_left_fold,
  [sym__unary_right_fold] = sym__unary_right_fold,
  [sym__binary_fold] = sym__binary_fold,
  [sym_fold_expression] = sym_fold_expression,
  [sym_parameter_pack_expansion] = sym_parameter_pack_expansion,
  [sym_type_parameter_pack_expansion] = sym_parameter_pack_expansion,
  [sym_identifier_parameter_pack_expansion] = sym_parameter_pack_expansion,
  [sym_destructor_name] = sym_destructor_name,
  [sym_dependent_identifier] = sym_dependent_identifier,
  [sym_dependent_field_identifier] = sym_dependent_identifier,
  [sym_dependent_type_identifier] = sym_dependent_identifier,
  [sym__scope_resolution] = sym__scope_resolution,
  [sym_qualified_field_identifier] = sym_qualified_identifier,
  [sym_qualified_identifier] = sym_qualified_identifier,
  [sym_qualified_type_identifier] = sym_qualified_identifier,
  [sym_qualified_operator_cast_identifier] = sym_qualified_identifier,
  [sym__assignment_expression_lhs] = sym_assignment_expression,
  [sym_operator_name] = sym_operator_name,
  [sym_user_defined_literal] = sym_user_defined_literal,
  [aux_sym_translation_unit_repeat1] = aux_sym_translation_unit_repeat1,
  [aux_sym_preproc_params_repeat1] = aux_sym_preproc_params_repeat1,
  [aux_sym_preproc_if_repeat1] = aux_sym_preproc_if_repeat1,
  [aux_sym_preproc_if_in_field_declaration_list_repeat1] = aux_sym_preproc_if_in_field_declaration_list_repeat1,
  [aux_sym_preproc_if_in_enumerator_list_repeat1] = aux_sym_preproc_if_in_enumerator_list_repeat1,
  [aux_sym_preproc_if_in_enumerator_list_no_comma_repeat1] = aux_sym_preproc_if_in_enumerator_list_no_comma_repeat1,
  [aux_sym_preproc_argument_list_repeat1] = aux_sym_preproc_argument_list_repeat1,
  [aux_sym_declaration_repeat1] = aux_sym_declaration_repeat1,
  [aux_sym_type_definition_repeat1] = aux_sym_type_definition_repeat1,
  [aux_sym__type_definition_type_repeat1] = aux_sym__type_definition_type_repeat1,
  [aux_sym__type_definition_declarators_repeat1] = aux_sym__type_definition_declarators_repeat1,
  [aux_sym__declaration_specifiers_repeat1] = aux_sym__declaration_specifiers_repeat1,
  [aux_sym_attribute_declaration_repeat1] = aux_sym_attribute_declaration_repeat1,
  [aux_sym_attributed_declarator_repeat1] = aux_sym_attributed_declarator_repeat1,
  [aux_sym_pointer_declarator_repeat1] = aux_sym_pointer_declarator_repeat1,
  [aux_sym_array_declarator_repeat1] = aux_sym_array_declarator_repeat1,
  [aux_sym_sized_type_specifier_repeat1] = aux_sym_sized_type_specifier_repeat1,
  [aux_sym_enumerator_list_repeat1] = aux_sym_enumerator_list_repeat1,
  [aux_sym_field_declaration_repeat1] = aux_sym_field_declaration_repeat1,
  [aux_sym_parameter_list_repeat1] = aux_sym_parameter_list_repeat1,
  [aux_sym_case_statement_repeat1] = aux_sym_case_statement_repeat1,
  [aux_sym_generic_expression_repeat1] = aux_sym_generic_expression_repeat1,
  [aux_sym_gnu_asm_expression_repeat1] = aux_sym_gnu_asm_expression_repeat1,
  [aux_sym_gnu_asm_output_operand_list_repeat1] = aux_sym_gnu_asm_output_operand_list_repeat1,
  [aux_sym_gnu_asm_input_operand_list_repeat1] = aux_sym_gnu_asm_input_operand_list_repeat1,
  [aux_sym_gnu_asm_clobber_list_repeat1] = aux_sym_gnu_asm_clobber_list_repeat1,
  [aux_sym_gnu_asm_goto_list_repeat1] = aux_sym_gnu_asm_goto_list_repeat1,
  [aux_sym_argument_list_repeat1] = aux_sym_argument_list_repeat1,
  [aux_sym_initializer_list_repeat1] = aux_sym_initializer_list_repeat1,
  [aux_sym_initializer_pair_repeat1] = aux_sym_initializer_pair_repeat1,
  [aux_sym_char_literal_repeat1] = aux_sym_char_literal_repeat1,
  [aux_sym_concatenated_string_repeat1] = aux_sym_concatenated_string_repeat1,
  [aux_sym_string_literal_repeat1] = aux_sym_string_literal_repeat1,
  [aux_sym__class_declaration_repeat1] = aux_sym__class_declaration_repeat1,
  [aux_sym_base_class_clause_repeat1] = aux_sym_base_class_clause_repeat1,
  [aux_sym_template_parameter_list_repeat1] = aux_sym_template_parameter_list_repeat1,
  [aux_sym_field_initializer_list_repeat1] = aux_sym_field_initializer_list_repeat1,
  [aux_sym_operator_cast_definition_repeat1] = aux_sym_operator_cast_definition_repeat1,
  [aux_sym_constructor_try_statement_repeat1] = aux_sym_constructor_try_statement_repeat1,
  [aux_sym_structured_binding_declarator_repeat1] = aux_sym_structured_binding_declarator_repeat1,
  [aux_sym__function_postfix_repeat1] = aux_sym__function_postfix_repeat1,
  [aux_sym_throw_specifier_repeat1] = aux_sym_throw_specifier_repeat1,
  [aux_sym_template_argument_list_repeat1] = aux_sym_template_argument_list_repeat1,
  [aux_sym_subscript_argument_list_repeat1] = aux_sym_subscript_argument_list_repeat1,
  [aux_sym_requirement_seq_repeat1] = aux_sym_requirement_seq_repeat1,
  [aux_sym_requires_parameter_list_repeat1] = aux_sym_requires_parameter_list_repeat1,
  [aux_sym_lambda_capture_specifier_repeat1] = aux_sym_lambda_capture_specifier_repeat1,
  [alias_sym_field_identifier] = alias_sym_field_identifier,
  [alias_sym_namespace_identifier] = alias_sym_namespace_identifier,
  [alias_sym_simple_requirement] = alias_sym_simple_requirement,
  [alias_sym_statement_identifier] = alias_sym_statement_identifier,
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
  [aux_sym_preproc_include_token1] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_preproc_include_token2] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_preproc_def_token1] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LPAREN] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DOT_DOT_DOT] = {
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
  [aux_sym_preproc_if_token1] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LF] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_preproc_if_token2] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_preproc_ifdef_token1] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_preproc_ifdef_token2] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_preproc_else_token1] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_preproc_elif_token1] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_preproc_elifdef_token1] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_preproc_elifdef_token2] = {
    .visible = true,
    .named = false,
  },
  [sym_preproc_arg] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_directive] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_LPAREN2] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_defined] = {
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
  [anon_sym_DASH] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PLUS] = {
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
  [anon_sym_PERCENT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PIPE_PIPE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_AMP_AMP] = {
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
  [anon_sym_AMP] = {
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
  [anon_sym_GT] = {
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
  [anon_sym_LT] = {
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
  [anon_sym_SEMI] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___extension__] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_typedef] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_virtual] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_extern] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___attribute__] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___attribute] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_COLON_COLON] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LBRACK_LBRACK] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_RBRACK_RBRACK] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___declspec] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___based] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___cdecl] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___clrcall] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___stdcall] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___fastcall] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___thiscall] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___vectorcall] = {
    .visible = true,
    .named = false,
  },
  [sym_ms_restrict_modifier] = {
    .visible = true,
    .named = true,
  },
  [sym_ms_unsigned_ptr_modifier] = {
    .visible = true,
    .named = true,
  },
  [sym_ms_signed_ptr_modifier] = {
    .visible = true,
    .named = true,
  },
  [anon_sym__unaligned] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___unaligned] = {
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
  [anon_sym_signed] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_unsigned] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_long] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_short] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LBRACK] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_static] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_RBRACK] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_register] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_inline] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___inline] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___inline__] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___forceinline] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_thread_local] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___thread] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_const] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_constexpr] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_volatile] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_restrict] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___restrict__] = {
    .visible = true,
    .named = false,
  },
  [anon_sym__Atomic] = {
    .visible = true,
    .named = false,
  },
  [anon_sym__Noreturn] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_noreturn] = {
    .visible = true,
    .named = false,
  },
  [anon_sym__Nonnull] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_mutable] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_constinit] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_consteval] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_alignas] = {
    .visible = true,
    .named = false,
  },
  [anon_sym__Alignas] = {
    .visible = true,
    .named = false,
  },
  [sym_primitive_type] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_enum] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_class] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_struct] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_union] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_COLON] = {
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
  [anon_sym_switch] = {
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
  [anon_sym_while] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_do] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_for] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_return] = {
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
  [anon_sym_goto] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___try] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___except] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___finally] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___leave] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_QMARK] = {
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
  [anon_sym_PLUS_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DASH_EQ] = {
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
  [anon_sym_AMP_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_CARET_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PIPE_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_and_eq] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_or_eq] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_xor_eq] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_not] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_compl] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LT_EQ_GT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_or] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_and] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_bitor] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_xor] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_bitand] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_not_eq] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DASH_DASH] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PLUS_PLUS] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_sizeof] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___alignof__] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___alignof] = {
    .visible = true,
    .named = false,
  },
  [anon_sym__alignof] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_alignof] = {
    .visible = true,
    .named = false,
  },
  [anon_sym__Alignof] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_offsetof] = {
    .visible = true,
    .named = false,
  },
  [anon_sym__Generic] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_asm] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___asm__] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___asm] = {
    .visible = true,
    .named = false,
  },
  [anon_sym___volatile__] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DOT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DOT_STAR] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DASH_GT] = {
    .visible = true,
    .named = false,
  },
  [sym_number_literal] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_L_SQUOTE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_u_SQUOTE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_U_SQUOTE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_u8_SQUOTE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_SQUOTE] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_char_literal_token1] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_L_DQUOTE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_u_DQUOTE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_U_DQUOTE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_u8_DQUOTE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DQUOTE] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_string_literal_token1] = {
    .visible = true,
    .named = true,
  },
  [sym_escape_sequence] = {
    .visible = true,
    .named = true,
  },
  [sym_system_lib_string] = {
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
  [anon_sym_NULL] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_nullptr] = {
    .visible = true,
    .named = false,
  },
  [sym_comment] = {
    .visible = true,
    .named = true,
  },
  [sym_auto] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_decltype] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_final] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_override] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_explicit] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_typename] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_template] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_GT2] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_operator] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_try] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_delete] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_pure_virtual_clause_token1] = {
    .visible = false,
    .named = false,
  },
  [anon_sym_friend] = {
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
  [anon_sym_noexcept] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_throw] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_namespace] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_using] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_static_assert] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_concept] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_co_return] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_co_yield] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_catch] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_R_DQUOTE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LR_DQUOTE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_uR_DQUOTE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_UR_DQUOTE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_u8R_DQUOTE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_co_await] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_new] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_requires] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DASH_GT_STAR] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LPAREN_RPAREN] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LBRACK_RBRACK] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DQUOTE_DQUOTE] = {
    .visible = true,
    .named = false,
  },
  [sym_this] = {
    .visible = true,
    .named = true,
  },
  [sym_literal_suffix] = {
    .visible = true,
    .named = true,
  },
  [sym_raw_string_delimiter] = {
    .visible = true,
    .named = true,
  },
  [sym_raw_string_content] = {
    .visible = true,
    .named = true,
  },
  [sym_translation_unit] = {
    .visible = true,
    .named = true,
  },
  [sym__top_level_item] = {
    .visible = false,
    .named = true,
  },
  [sym__block_item] = {
    .visible = false,
    .named = true,
  },
  [sym_preproc_include] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_def] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_function_def] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_params] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_call] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_if] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_ifdef] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_else] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_elif] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_elifdef] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_if_in_field_declaration_list] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_ifdef_in_field_declaration_list] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_else_in_field_declaration_list] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_elif_in_field_declaration_list] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_elifdef_in_field_declaration_list] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_if_in_enumerator_list] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_ifdef_in_enumerator_list] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_else_in_enumerator_list] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_elif_in_enumerator_list] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_elifdef_in_enumerator_list] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_if_in_enumerator_list_no_comma] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_ifdef_in_enumerator_list_no_comma] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_else_in_enumerator_list_no_comma] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_elif_in_enumerator_list_no_comma] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_elifdef_in_enumerator_list_no_comma] = {
    .visible = true,
    .named = true,
  },
  [sym__preproc_expression] = {
    .visible = false,
    .named = true,
  },
  [sym_preproc_parenthesized_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_defined] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_unary_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_call_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_argument_list] = {
    .visible = true,
    .named = true,
  },
  [sym_preproc_binary_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_function_definition] = {
    .visible = true,
    .named = true,
  },
  [sym_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_type_definition] = {
    .visible = true,
    .named = true,
  },
  [sym__type_definition_type] = {
    .visible = false,
    .named = true,
  },
  [sym__type_definition_declarators] = {
    .visible = false,
    .named = true,
  },
  [sym__declaration_modifiers] = {
    .visible = false,
    .named = true,
  },
  [sym__declaration_specifiers] = {
    .visible = false,
    .named = true,
  },
  [sym_linkage_specification] = {
    .visible = true,
    .named = true,
  },
  [sym_attribute_specifier] = {
    .visible = true,
    .named = true,
  },
  [sym_attribute] = {
    .visible = true,
    .named = true,
  },
  [sym_attribute_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_ms_declspec_modifier] = {
    .visible = true,
    .named = true,
  },
  [sym_ms_based_modifier] = {
    .visible = true,
    .named = true,
  },
  [sym_ms_call_modifier] = {
    .visible = true,
    .named = true,
  },
  [sym_ms_unaligned_ptr_modifier] = {
    .visible = true,
    .named = true,
  },
  [sym_ms_pointer_modifier] = {
    .visible = true,
    .named = true,
  },
  [sym_declaration_list] = {
    .visible = true,
    .named = true,
  },
  [sym__declarator] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym__field_declarator] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym__type_declarator] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym__abstract_declarator] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym_parenthesized_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_parenthesized_field_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_parenthesized_type_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_abstract_parenthesized_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_attributed_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_attributed_field_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_attributed_type_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_pointer_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_pointer_field_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_pointer_type_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_abstract_pointer_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_function_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_function_field_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_function_type_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_abstract_function_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_array_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_array_field_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_array_type_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_abstract_array_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_init_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_compound_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_storage_class_specifier] = {
    .visible = true,
    .named = true,
  },
  [sym_type_qualifier] = {
    .visible = true,
    .named = true,
  },
  [sym_alignas_qualifier] = {
    .visible = true,
    .named = true,
  },
  [sym_type_specifier] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym_sized_type_specifier] = {
    .visible = true,
    .named = true,
  },
  [sym_enum_specifier] = {
    .visible = true,
    .named = true,
  },
  [sym_enumerator_list] = {
    .visible = true,
    .named = true,
  },
  [sym_struct_specifier] = {
    .visible = true,
    .named = true,
  },
  [sym_union_specifier] = {
    .visible = true,
    .named = true,
  },
  [sym_field_declaration_list] = {
    .visible = true,
    .named = true,
  },
  [sym__field_declaration_list_item] = {
    .visible = false,
    .named = true,
  },
  [sym_field_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_bitfield_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_enumerator] = {
    .visible = true,
    .named = true,
  },
  [sym_parameter_list] = {
    .visible = true,
    .named = true,
  },
  [sym_parameter_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_attributed_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_statement] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym__top_level_statement] = {
    .visible = false,
    .named = true,
  },
  [sym_labeled_statement] = {
    .visible = true,
    .named = true,
  },
  [sym__top_level_expression_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_expression_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_if_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_else_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_switch_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_case_statement] = {
    .visible = true,
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
  [sym_for_statement] = {
    .visible = true,
    .named = true,
  },
  [sym__for_statement_body] = {
    .visible = false,
    .named = true,
  },
  [sym_return_statement] = {
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
  [sym_goto_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_seh_try_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_seh_except_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_seh_finally_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_seh_leave_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_expression] = {
    .visible = false,
    .named = true,
    .supertype = true,
  },
  [sym__string] = {
    .visible = false,
    .named = true,
  },
  [sym_comma_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_conditional_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_assignment_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_pointer_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_unary_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_binary_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_update_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_cast_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_type_descriptor] = {
    .visible = true,
    .named = true,
  },
  [sym_sizeof_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_alignof_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_offsetof_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_generic_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_subscript_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_call_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_gnu_asm_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_gnu_asm_qualifier] = {
    .visible = true,
    .named = true,
  },
  [sym_gnu_asm_output_operand_list] = {
    .visible = true,
    .named = true,
  },
  [sym_gnu_asm_output_operand] = {
    .visible = true,
    .named = true,
  },
  [sym_gnu_asm_input_operand_list] = {
    .visible = true,
    .named = true,
  },
  [sym_gnu_asm_input_operand] = {
    .visible = true,
    .named = true,
  },
  [sym_gnu_asm_clobber_list] = {
    .visible = true,
    .named = true,
  },
  [sym_gnu_asm_goto_list] = {
    .visible = true,
    .named = true,
  },
  [sym_extension_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_argument_list] = {
    .visible = true,
    .named = true,
  },
  [sym_field_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_compound_literal_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_parenthesized_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_initializer_list] = {
    .visible = true,
    .named = true,
  },
  [sym_initializer_pair] = {
    .visible = true,
    .named = true,
  },
  [sym_subscript_designator] = {
    .visible = true,
    .named = true,
  },
  [sym_subscript_range_designator] = {
    .visible = true,
    .named = true,
  },
  [sym_field_designator] = {
    .visible = true,
    .named = true,
  },
  [sym_char_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_concatenated_string] = {
    .visible = true,
    .named = true,
  },
  [sym_string_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_null] = {
    .visible = true,
    .named = true,
  },
  [sym__empty_declaration] = {
    .visible = false,
    .named = true,
  },
  [sym_placeholder_type_specifier] = {
    .visible = true,
    .named = true,
  },
  [sym_decltype_auto] = {
    .visible = true,
    .named = true,
  },
  [sym_decltype] = {
    .visible = true,
    .named = true,
  },
  [sym__class_declaration] = {
    .visible = false,
    .named = true,
  },
  [sym__class_declaration_item] = {
    .visible = false,
    .named = true,
  },
  [sym_class_specifier] = {
    .visible = true,
    .named = true,
  },
  [sym__class_name] = {
    .visible = false,
    .named = true,
  },
  [sym_virtual_specifier] = {
    .visible = true,
    .named = true,
  },
  [sym_explicit_function_specifier] = {
    .visible = true,
    .named = true,
  },
  [sym_base_class_clause] = {
    .visible = true,
    .named = true,
  },
  [sym__enum_base_clause] = {
    .visible = false,
    .named = true,
  },
  [sym_dependent_type] = {
    .visible = true,
    .named = true,
  },
  [sym_template_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_template_instantiation] = {
    .visible = true,
    .named = true,
  },
  [sym_template_parameter_list] = {
    .visible = true,
    .named = true,
  },
  [sym_type_parameter_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_variadic_type_parameter_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_optional_type_parameter_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_template_template_parameter_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_optional_parameter_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_variadic_parameter_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_variadic_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_variadic_reference_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_operator_cast] = {
    .visible = true,
    .named = true,
  },
  [sym_field_initializer_list] = {
    .visible = true,
    .named = true,
  },
  [sym_field_initializer] = {
    .visible = true,
    .named = true,
  },
  [sym_inline_method_definition] = {
    .visible = true,
    .named = true,
  },
  [sym__constructor_specifiers] = {
    .visible = false,
    .named = true,
  },
  [sym_operator_cast_definition] = {
    .visible = true,
    .named = true,
  },
  [sym_operator_cast_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_constructor_try_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_constructor_or_destructor_definition] = {
    .visible = true,
    .named = true,
  },
  [sym_constructor_or_destructor_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_default_method_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_delete_method_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_pure_virtual_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_friend_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_access_specifier] = {
    .visible = true,
    .named = true,
  },
  [sym_reference_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_reference_field_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_reference_type_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_abstract_reference_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_structured_binding_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_ref_qualifier] = {
    .visible = true,
    .named = true,
  },
  [sym__function_declarator_seq] = {
    .visible = false,
    .named = true,
  },
  [sym__function_attributes_start] = {
    .visible = false,
    .named = true,
  },
  [sym__function_exception_specification] = {
    .visible = false,
    .named = true,
  },
  [sym__function_attributes_end] = {
    .visible = false,
    .named = true,
  },
  [sym__function_postfix] = {
    .visible = false,
    .named = true,
  },
  [sym_trailing_return_type] = {
    .visible = true,
    .named = true,
  },
  [sym_noexcept] = {
    .visible = true,
    .named = true,
  },
  [sym_throw_specifier] = {
    .visible = true,
    .named = true,
  },
  [sym_template_type] = {
    .visible = true,
    .named = true,
  },
  [sym_template_method] = {
    .visible = true,
    .named = true,
  },
  [sym_template_function] = {
    .visible = true,
    .named = true,
  },
  [sym_template_argument_list] = {
    .visible = true,
    .named = true,
  },
  [sym_namespace_definition] = {
    .visible = true,
    .named = true,
  },
  [sym_namespace_alias_definition] = {
    .visible = true,
    .named = true,
  },
  [sym__namespace_specifier] = {
    .visible = false,
    .named = true,
  },
  [sym_nested_namespace_specifier] = {
    .visible = true,
    .named = true,
  },
  [sym_using_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_alias_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_static_assert_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_concept_definition] = {
    .visible = true,
    .named = true,
  },
  [sym_for_range_loop] = {
    .visible = true,
    .named = true,
  },
  [sym__for_range_loop_body] = {
    .visible = false,
    .named = true,
  },
  [sym_init_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_condition_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_condition_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_co_return_statement] = {
    .visible = true,
    .named = true,
  },
  [sym_co_yield_statement] = {
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
  [sym_raw_string_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_subscript_argument_list] = {
    .visible = true,
    .named = true,
  },
  [sym_co_await_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_new_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_new_declarator] = {
    .visible = true,
    .named = true,
  },
  [sym_delete_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_type_requirement] = {
    .visible = true,
    .named = true,
  },
  [sym_compound_requirement] = {
    .visible = true,
    .named = true,
  },
  [sym__requirement] = {
    .visible = false,
    .named = true,
  },
  [sym_requirement_seq] = {
    .visible = true,
    .named = true,
  },
  [sym_constraint_conjunction] = {
    .visible = true,
    .named = true,
  },
  [sym_constraint_disjunction] = {
    .visible = true,
    .named = true,
  },
  [sym__requirement_clause_constraint] = {
    .visible = false,
    .named = true,
  },
  [sym_requires_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_requires_parameter_list] = {
    .visible = true,
    .named = true,
  },
  [sym_requires_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_lambda_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_lambda_capture_specifier] = {
    .visible = true,
    .named = true,
  },
  [sym_lambda_default_capture] = {
    .visible = true,
    .named = true,
  },
  [sym__lambda_capture_identifier] = {
    .visible = false,
    .named = true,
  },
  [sym_lambda_capture_initializer] = {
    .visible = true,
    .named = true,
  },
  [sym__lambda_capture] = {
    .visible = false,
    .named = true,
  },
  [sym__fold_operator] = {
    .visible = false,
    .named = true,
  },
  [sym__binary_fold_operator] = {
    .visible = false,
    .named = true,
  },
  [sym__unary_left_fold] = {
    .visible = false,
    .named = true,
  },
  [sym__unary_right_fold] = {
    .visible = false,
    .named = true,
  },
  [sym__binary_fold] = {
    .visible = false,
    .named = true,
  },
  [sym_fold_expression] = {
    .visible = true,
    .named = true,
  },
  [sym_parameter_pack_expansion] = {
    .visible = true,
    .named = true,
  },
  [sym_type_parameter_pack_expansion] = {
    .visible = true,
    .named = true,
  },
  [sym_identifier_parameter_pack_expansion] = {
    .visible = true,
    .named = true,
  },
  [sym_destructor_name] = {
    .visible = true,
    .named = true,
  },
  [sym_dependent_identifier] = {
    .visible = true,
    .named = true,
  },
  [sym_dependent_field_identifier] = {
    .visible = true,
    .named = true,
  },
  [sym_dependent_type_identifier] = {
    .visible = true,
    .named = true,
  },
  [sym__scope_resolution] = {
    .visible = false,
    .named = true,
  },
  [sym_qualified_field_identifier] = {
    .visible = true,
    .named = true,
  },
  [sym_qualified_identifier] = {
    .visible = true,
    .named = true,
  },
  [sym_qualified_type_identifier] = {
    .visible = true,
    .named = true,
  },
  [sym_qualified_operator_cast_identifier] = {
    .visible = true,
    .named = true,
  },
  [sym__assignment_expression_lhs] = {
    .visible = true,
    .named = true,
  },
  [sym_operator_name] = {
    .visible = true,
    .named = true,
  },
  [sym_user_defined_literal] = {
    .visible = true,
    .named = true,
  },
  [aux_sym_translation_unit_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_preproc_params_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_preproc_if_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_preproc_if_in_field_declaration_list_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_preproc_if_in_enumerator_list_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_preproc_if_in_enumerator_list_no_comma_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_preproc_argument_list_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_declaration_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_type_definition_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__type_definition_type_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__type_definition_declarators_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__declaration_specifiers_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_attribute_declaration_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_attributed_declarator_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_pointer_declarator_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_array_declarator_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_sized_type_specifier_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_enumerator_list_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_field_declaration_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_parameter_list_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_case_statement_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_generic_expression_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_gnu_asm_expression_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_gnu_asm_output_operand_list_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_gnu_asm_input_operand_list_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_gnu_asm_clobber_list_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_gnu_asm_goto_list_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_argument_list_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_initializer_list_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_initializer_pair_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_char_literal_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_concatenated_string_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_string_literal_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__class_declaration_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_base_class_clause_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_template_parameter_list_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_field_initializer_list_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_operator_cast_definition_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_constructor_try_statement_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_structured_binding_declarator_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__function_postfix_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_throw_specifier_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_template_argument_list_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_subscript_argument_list_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_requirement_seq_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_requires_parameter_list_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_lambda_capture_specifier_repeat1] = {
    .visible = false,
    .named = false,
  },
  [alias_sym_field_identifier] = {
    .visible = true,
    .named = true,
  },
  [alias_sym_namespace_identifier] = {
    .visible = true,
    .named = true,
  },
  [alias_sym_simple_requirement] = {
    .visible = true,
    .named = true,
  },
  [alias_sym_statement_identifier] = {
    .visible = true,
    .named = true,
  },
  [alias_sym_type_identifier] = {
    .visible = true,
    .named = true,
  },
};

enum ts_field_identifiers {
  field_alternative = 1,
  field_argument = 2,
  field_arguments = 3,
  field_assembly_code = 4,
  field_base = 5,
  field_body = 6,
  field_captures = 7,
  field_clobbers = 8,
  field_condition = 9,
  field_consequence = 10,
  field_constraint = 11,
  field_declarator = 12,
  field_default_type = 13,
  field_default_value = 14,
  field_delimiter = 15,
  field_designator = 16,
  field_directive = 17,
  field_end = 18,
  field_field = 19,
  field_filter = 20,
  field_function = 21,
  field_goto_labels = 22,
  field_indices = 23,
  field_initializer = 24,
  field_input_operands = 25,
  field_label = 26,
  field_left = 27,
  field_length = 28,
  field_member = 29,
  field_message = 30,
  field_name = 31,
  field_operand = 32,
  field_operator = 33,
  field_output_operands = 34,
  field_parameters = 35,
  field_path = 36,
  field_pattern = 37,
  field_placement = 38,
  field_prefix = 39,
  field_register = 40,
  field_requirements = 41,
  field_right = 42,
  field_scope = 43,
  field_size = 44,
  field_start = 45,
  field_symbol = 46,
  field_template_parameters = 47,
  field_type = 48,
  field_update = 49,
  field_value = 50,
};

static const char * const ts_field_names[] = {
  [0] = NULL,
  [field_alternative] = "alternative",
  [field_argument] = "argument",
  [field_arguments] = "arguments",
  [field_assembly_code] = "assembly_code",
  [field_base] = "base",
  [field_body] = "body",
  [field_captures] = "captures",
  [field_clobbers] = "clobbers",
  [field_condition] = "condition",
  [field_consequence] = "consequence",
  [field_constraint] = "constraint",
  [field_declarator] = "declarator",
  [field_default_type] = "default_type",
  [field_default_value] = "default_value",
  [field_delimiter] = "delimiter",
  [field_designator] = "designator",
  [field_directive] = "directive",
  [field_end] = "end",
  [field_field] = "field",
  [field_filter] = "filter",
  [field_function] = "function",
  [field_goto_labels] = "goto_labels",
  [field_indices] = "indices",
  [field_initializer] = "initializer",
  [field_input_operands] = "input_operands",
  [field_label] = "label",
  [field_left] = "left",
  [field_length] = "length",
  [field_member] = "member",
  [field_message] = "message",
  [field_name] = "name",
  [field_operand] = "operand",
  [field_operator] = "operator",
  [field_output_operands] = "output_operands",
  [field_parameters] = "parameters",
  [field_path] = "path",
  [field_pattern] = "pattern",
  [field_placement] = "placement",
  [field_prefix] = "prefix",
  [field_register] = "register",
  [field_requirements] = "requirements",
  [field_right] = "right",
  [field_scope] = "scope",
  [field_size] = "size",
  [field_start] = "start",
  [field_symbol] = "symbol",
  [field_template_parameters] = "template_parameters",
  [field_type] = "type",
  [field_update] = "update",
  [field_value] = "value",
};

static const TSFieldMapSlice ts_field_map_slices[PRODUCTION_ID_COUNT] = {
  [2] = {.index = 0, .length = 1},
  [3] = {.index = 1, .length = 1},
  [4] = {.index = 2, .length = 2},
  [5] = {.index = 4, .length = 1},
  [6] = {.index = 5, .length = 1},
  [7] = {.index = 6, .length = 2},
  [8] = {.index = 8, .length = 2},
  [9] = {.index = 10, .length = 1},
  [10] = {.index = 11, .length = 1},
  [11] = {.index = 12, .length = 1},
  [12] = {.index = 13, .length = 2},
  [13] = {.index = 15, .length = 2},
  [14] = {.index = 17, .length = 1},
  [15] = {.index = 18, .length = 1},
  [16] = {.index = 19, .length = 2},
  [17] = {.index = 19, .length = 2},
  [18] = {.index = 0, .length = 1},
  [20] = {.index = 21, .length = 1},
  [21] = {.index = 22, .length = 1},
  [22] = {.index = 23, .length = 1},
  [23] = {.index = 24, .length = 1},
  [24] = {.index = 25, .length = 2},
  [25] = {.index = 27, .length = 2},
  [26] = {.index = 29, .length = 1},
  [27] = {.index = 30, .length = 1},
  [28] = {.index = 31, .length = 1},
  [29] = {.index = 32, .length = 2},
  [30] = {.index = 34, .length = 2},
  [31] = {.index = 18, .length = 1},
  [32] = {.index = 36, .length = 2},
  [33] = {.index = 38, .length = 1},
  [34] = {.index = 39, .length = 2},
  [35] = {.index = 39, .length = 2},
  [36] = {.index = 21, .length = 1},
  [37] = {.index = 41, .length = 1},
  [38] = {.index = 42, .length = 2},
  [39] = {.index = 44, .length = 2},
  [40] = {.index = 46, .length = 3},
  [41] = {.index = 49, .length = 1},
  [44] = {.index = 50, .length = 2},
  [45] = {.index = 52, .length = 1},
  [46] = {.index = 53, .length = 1},
  [47] = {.index = 54, .length = 1},
  [48] = {.index = 55, .length = 2},
  [49] = {.index = 57, .length = 2},
  [50] = {.index = 59, .length = 2},
  [51] = {.index = 61, .length = 2},
  [52] = {.index = 63, .length = 2},
  [53] = {.index = 65, .length = 1},
  [54] = {.index = 66, .length = 3},
  [55] = {.index = 69, .length = 1},
  [56] = {.index = 70, .length = 1},
  [57] = {.index = 71, .length = 1},
  [58] = {.index = 72, .length = 2},
  [60] = {.index = 55, .length = 2},
  [61] = {.index = 74, .length = 2},
  [62] = {.index = 76, .length = 2},
  [63] = {.index = 78, .length = 2},
  [65] = {.index = 80, .length = 2},
  [66] = {.index = 82, .length = 2},
  [67] = {.index = 84, .length = 3},
  [68] = {.index = 87, .length = 2},
  [69] = {.index = 89, .length = 2},
  [70] = {.index = 91, .length = 3},
  [71] = {.index = 91, .length = 3},
  [72] = {.index = 94, .length = 3},
  [73] = {.index = 97, .length = 3},
  [74] = {.index = 100, .length = 3},
  [75] = {.index = 52, .length = 1},
  [76] = {.index = 103, .length = 2},
  [77] = {.index = 105, .length = 2},
  [78] = {.index = 107, .length = 2},
  [79] = {.index = 109, .length = 1},
  [80] = {.index = 110, .length = 2},
  [81] = {.index = 112, .length = 2},
  [82] = {.index = 114, .length = 2},
  [83] = {.index = 116, .length = 3},
  [84] = {.index = 119, .length = 2},
  [85] = {.index = 121, .length = 1},
  [86] = {.index = 122, .length = 2},
  [87] = {.index = 124, .length = 2},
  [88] = {.index = 126, .length = 2},
  [89] = {.index = 128, .length = 2},
  [90] = {.index = 130, .length = 2},
  [91] = {.index = 132, .length = 2},
  [92] = {.index = 134, .length = 2},
  [93] = {.index = 136, .length = 2},
  [94] = {.index = 138, .length = 1},
  [95] = {.index = 136, .length = 2},
  [97] = {.index = 139, .length = 2},
  [98] = {.index = 141, .length = 1},
  [99] = {.index = 141, .length = 1},
  [100] = {.index = 142, .length = 3},
  [102] = {.index = 145, .length = 2},
  [103] = {.index = 147, .length = 2},
  [104] = {.index = 149, .length = 2},
  [105] = {.index = 151, .length = 3},
  [106] = {.index = 154, .length = 1},
  [107] = {.index = 155, .length = 1},
  [109] = {.index = 156, .length = 3},
  [110] = {.index = 159, .length = 3},
  [111] = {.index = 162, .length = 3},
  [112] = {.index = 165, .length = 3},
  [113] = {.index = 168, .length = 2},
  [114] = {.index = 170, .length = 3},
  [115] = {.index = 173, .length = 3},
  [116] = {.index = 176, .length = 2},
  [117] = {.index = 178, .length = 3},
  [118] = {.index = 181, .length = 2},
  [119] = {.index = 19, .length = 2},
  [120] = {.index = 39, .length = 2},
  [121] = {.index = 183, .length = 2},
  [122] = {.index = 185, .length = 2},
  [123] = {.index = 187, .length = 1},
  [124] = {.index = 188, .length = 4},
  [125] = {.index = 192, .length = 4},
  [126] = {.index = 196, .length = 2},
  [127] = {.index = 198, .length = 3},
  [128] = {.index = 201, .length = 2},
  [129] = {.index = 203, .length = 2},
  [130] = {.index = 205, .length = 1},
  [131] = {.index = 206, .length = 2},
  [132] = {.index = 208, .length = 2},
  [133] = {.index = 210, .length = 3},
  [134] = {.index = 213, .length = 3},
  [135] = {.index = 216, .length = 3},
  [136] = {.index = 219, .length = 2},
  [137] = {.index = 221, .length = 2},
  [138] = {.index = 221, .length = 2},
  [139] = {.index = 223, .length = 2},
  [140] = {.index = 223, .length = 2},
  [141] = {.index = 225, .length = 2},
  [142] = {.index = 227, .length = 3},
  [143] = {.index = 230, .length = 2},
  [144] = {.index = 232, .length = 2},
  [145] = {.index = 234, .length = 3},
  [146] = {.index = 237, .length = 2},
  [147] = {.index = 239, .length = 3},
  [148] = {.index = 242, .length = 2},
  [149] = {.index = 244, .length = 1},
  [150] = {.index = 245, .length = 2},
  [151] = {.index = 247, .length = 2},
  [152] = {.index = 249, .length = 4},
  [153] = {.index = 253, .length = 5},
  [154] = {.index = 258, .length = 1},
  [155] = {.index = 259, .length = 1},
  [156] = {.index = 260, .length = 2},
  [157] = {.index = 262, .length = 1},
  [159] = {.index = 263, .length = 1},
  [160] = {.index = 264, .length = 2},
  [161] = {.index = 266, .length = 2},
  [162] = {.index = 11, .length = 1},
  [163] = {.index = 11, .length = 1},
  [164] = {.index = 268, .length = 2},
  [165] = {.index = 270, .length = 1},
  [166] = {.index = 271, .length = 1},
  [167] = {.index = 272, .length = 4},
  [168] = {.index = 276, .length = 2},
  [169] = {.index = 278, .length = 4},
  [170] = {.index = 282, .length = 1},
  [171] = {.index = 283, .length = 3},
  [172] = {.index = 286, .length = 2},
  [173] = {.index = 288, .length = 3},
  [174] = {.index = 291, .length = 1},
  [175] = {.index = 292, .length = 5},
  [176] = {.index = 297, .length = 2},
  [177] = {.index = 299, .length = 2},
  [178] = {.index = 301, .length = 4},
  [179] = {.index = 305, .length = 2},
  [180] = {.index = 307, .length = 2},
  [181] = {.index = 309, .length = 3},
  [182] = {.index = 312, .length = 4},
  [183] = {.index = 316, .length = 4},
  [184] = {.index = 320, .length = 3},
  [185] = {.index = 323, .length = 2},
  [186] = {.index = 325, .length = 3},
  [187] = {.index = 328, .length = 3},
  [188] = {.index = 331, .length = 2},
  [189] = {.index = 333, .length = 2},
  [190] = {.index = 335, .length = 2},
  [191] = {.index = 337, .length = 2},
  [192] = {.index = 339, .length = 3},
  [193] = {.index = 342, .length = 2},
  [194] = {.index = 344, .length = 2},
  [195] = {.index = 346, .length = 3},
  [196] = {.index = 349, .length = 2},
  [197] = {.index = 351, .length = 2},
  [198] = {.index = 353, .length = 2},
  [199] = {.index = 355, .length = 4},
  [200] = {.index = 359, .length = 5},
  [201] = {.index = 364, .length = 3},
  [202] = {.index = 367, .length = 4},
  [203] = {.index = 371, .length = 2},
  [204] = {.index = 373, .length = 1},
  [205] = {.index = 374, .length = 4},
  [206] = {.index = 378, .length = 3},
  [207] = {.index = 381, .length = 2},
  [208] = {.index = 383, .length = 1},
  [209] = {.index = 384, .length = 5},
  [210] = {.index = 389, .length = 2},
  [211] = {.index = 391, .length = 2},
  [212] = {.index = 65, .length = 1},
  [213] = {.index = 393, .length = 5},
  [214] = {.index = 398, .length = 4},
  [215] = {.index = 402, .length = 2},
  [216] = {.index = 404, .length = 2},
  [217] = {.index = 406, .length = 5},
  [218] = {.index = 411, .length = 2},
  [219] = {.index = 413, .length = 3},
};

static const TSFieldMapEntry ts_field_map_entries[] = {
  [0] =
    {field_type, 0},
  [1] =
    {field_directive, 0},
  [2] =
    {field_argument, 1},
    {field_operator, 0},
  [4] =
    {field_declarator, 1},
  [5] =
    {field_name, 0},
  [6] =
    {field_arguments, 1},
    {field_function, 0},
  [8] =
    {field_type, 0},
    {field_value, 1},
  [10] =
    {field_body, 1},
  [11] =
    {field_name, 1},
  [12] =
    {field_body, 0},
  [13] =
    {field_body, 1, .inherited = true},
    {field_name, 1, .inherited = true},
  [15] =
    {field_body, 0, .inherited = true},
    {field_name, 0, .inherited = true},
  [17] =
    {field_value, 1},
  [18] =
    {field_scope, 0},
  [19] =
    {field_arguments, 1},
    {field_name, 0},
  [21] =
    {field_type, 1},
  [22] =
    {field_requirements, 1},
  [23] =
    {field_constraint, 1},
  [24] =
    {field_parameters, 0},
  [25] =
    {field_declarator, 0},
    {field_parameters, 1, .inherited = true},
  [27] =
    {field_body, 1},
    {field_declarator, 0},
  [29] =
    {field_declarator, 0},
  [30] =
    {field_constraint, 0},
  [31] =
    {field_pattern, 0},
  [32] =
    {field_argument, 0},
    {field_operator, 1},
  [34] =
    {field_argument, 0},
    {field_indices, 1},
  [36] =
    {field_body, 1},
    {field_captures, 0},
  [38] =
    {field_parameters, 0, .inherited = true},
  [39] =
    {field_name, 1},
    {field_scope, 0, .inherited = true},
  [41] =
    {field_path, 1},
  [42] =
    {field_argument, 1},
    {field_directive, 0},
  [44] =
    {field_declarator, 1},
    {field_type, 0},
  [46] =
    {field_left, 1, .inherited = true},
    {field_operator, 1, .inherited = true},
    {field_right, 1, .inherited = true},
  [49] =
    {field_declarator, 2},
  [50] =
    {field_body, 2},
    {field_value, 1},
  [52] =
    {field_type, 2},
  [53] =
    {field_body, 2},
  [54] =
    {field_name, 2},
  [55] =
    {field_body, 2},
    {field_name, 1},
  [57] =
    {field_base, 2, .inherited = true},
    {field_name, 1},
  [59] =
    {field_body, 1},
    {field_name, 0},
  [61] =
    {field_condition, 1},
    {field_consequence, 2},
  [63] =
    {field_body, 2},
    {field_condition, 1},
  [65] =
    {field_label, 1},
  [66] =
    {field_left, 0},
    {field_operator, 1},
    {field_right, 2},
  [69] =
    {field_label, 0},
  [70] =
    {field_type, 0, .inherited = true},
  [71] =
    {field_parameters, 1},
  [72] =
    {field_declarator, 2},
    {field_type, 1, .inherited = true},
  [74] =
    {field_arguments, 2},
    {field_type, 1},
  [76] =
    {field_declarator, 2},
    {field_type, 1},
  [78] =
    {field_placement, 1},
    {field_type, 2},
  [80] =
    {field_parameters, 1},
    {field_requirements, 2},
  [82] =
    {field_declarator, 1},
    {field_type, 0, .inherited = true},
  [84] =
    {field_body, 2},
    {field_declarator, 1},
    {field_type, 0, .inherited = true},
  [87] =
    {field_declarator, 0},
    {field_value, 1},
  [89] =
    {field_body, 2},
    {field_declarator, 0},
  [91] =
    {field_argument, 0},
    {field_field, 2},
    {field_operator, 1},
  [94] =
    {field_body, 2},
    {field_captures, 0},
    {field_declarator, 1},
  [97] =
    {field_body, 2},
    {field_captures, 0},
    {field_template_parameters, 1},
  [100] =
    {field_name, 1},
    {field_name, 2},
    {field_scope, 0, .inherited = true},
  [103] =
    {field_body, 2},
    {field_declarator, 1},
  [105] =
    {field_name, 1},
    {field_value, 2},
  [107] =
    {field_name, 1},
    {field_parameters, 2},
  [109] =
    {field_condition, 1},
  [110] =
    {field_alternative, 2},
    {field_name, 1},
  [112] =
    {field_declarator, 2},
    {field_type, 0},
  [114] =
    {field_left, 0},
    {field_right, 2},
  [116] =
    {field_left, 0},
    {field_operator, 1, .inherited = true},
    {field_right, 2},
  [119] =
    {field_type, 1},
    {field_value, 3},
  [121] =
    {field_declarator, 3},
  [122] =
    {field_declarator, 2, .inherited = true},
    {field_type, 1, .inherited = true},
  [124] =
    {field_declarator, 0},
    {field_parameters, 1},
  [126] =
    {field_declarator, 0},
    {field_declarator, 1, .inherited = true},
  [128] =
    {field_arguments, 3},
    {field_type, 2},
  [130] =
    {field_declarator, 3},
    {field_type, 2},
  [132] =
    {field_placement, 2},
    {field_type, 3},
  [134] =
    {field_name, 2},
    {field_prefix, 0},
  [136] =
    {field_body, 3},
    {field_name, 2},
  [138] =
    {field_body, 3},
  [139] =
    {field_base, 3, .inherited = true},
    {field_name, 2},
  [141] =
    {field_base, 1},
  [142] =
    {field_base, 2, .inherited = true},
    {field_body, 3},
    {field_name, 1},
  [145] =
    {field_body, 2, .inherited = true},
    {field_name, 2, .inherited = true},
  [147] =
    {field_body, 2},
    {field_name, 0},
  [149] =
    {field_condition, 2},
    {field_consequence, 3},
  [151] =
    {field_alternative, 3},
    {field_condition, 1},
    {field_consequence, 2},
  [154] =
    {field_initializer, 0},
  [155] =
    {field_assembly_code, 2},
  [156] =
    {field_arguments, 3},
    {field_declarator, 2},
    {field_type, 1},
  [159] =
    {field_arguments, 3},
    {field_placement, 1},
    {field_type, 2},
  [162] =
    {field_declarator, 3},
    {field_placement, 1},
    {field_type, 2},
  [165] =
    {field_body, 3},
    {field_declarator, 2},
    {field_type, 0, .inherited = true},
  [168] =
    {field_declarator, 0},
    {field_value, 2},
  [170] =
    {field_declarator, 1},
    {field_declarator, 2},
    {field_type, 0, .inherited = true},
  [173] =
    {field_declarator, 1},
    {field_declarator, 2, .inherited = true},
    {field_type, 0, .inherited = true},
  [176] =
    {field_declarator, 0, .inherited = true},
    {field_declarator, 1, .inherited = true},
  [178] =
    {field_body, 3},
    {field_declarator, 2},
    {field_type, 1, .inherited = true},
  [181] =
    {field_declarator, 0},
    {field_size, 2},
  [183] =
    {field_alternative, 3},
    {field_condition, 0},
  [185] =
    {field_declarator, 0},
    {field_default_value, 2},
  [187] =
    {field_size, 1},
  [188] =
    {field_body, 3},
    {field_captures, 0},
    {field_declarator, 2},
    {field_template_parameters, 1},
  [192] =
    {field_body, 3},
    {field_captures, 0},
    {field_constraint, 2},
    {field_template_parameters, 1},
  [196] =
    {field_body, 3},
    {field_declarator, 1},
  [198] =
    {field_name, 1},
    {field_parameters, 2},
    {field_value, 3},
  [201] =
    {field_alternative, 3},
    {field_condition, 1},
  [203] =
    {field_alternative, 3},
    {field_name, 1},
  [205] =
    {field_operator, 0},
  [206] =
    {field_declarator, 3},
    {field_type, 1},
  [208] =
    {field_declarator, 3, .inherited = true},
    {field_type, 2, .inherited = true},
  [210] =
    {field_arguments, 4},
    {field_declarator, 3},
    {field_type, 2},
  [213] =
    {field_arguments, 4},
    {field_placement, 2},
    {field_type, 3},
  [216] =
    {field_declarator, 4},
    {field_placement, 2},
    {field_type, 3},
  [219] =
    {field_left, 1},
    {field_right, 3},
  [221] =
    {field_body, 4},
    {field_name, 3},
  [223] =
    {field_designator, 0},
    {field_value, 2},
  [225] =
    {field_name, 0},
    {field_value, 2},
  [227] =
    {field_base, 3, .inherited = true},
    {field_body, 4},
    {field_name, 2},
  [230] =
    {field_body, 3},
    {field_name, 0},
  [232] =
    {field_body, 3, .inherited = true},
    {field_name, 3, .inherited = true},
  [234] =
    {field_declarator, 1},
    {field_type, 0, .inherited = true},
    {field_value, 2},
  [237] =
    {field_initializer, 1},
    {field_value, 2},
  [239] =
    {field_alternative, 4},
    {field_condition, 2},
    {field_consequence, 3},
  [242] =
    {field_body, 1},
    {field_condition, 3},
  [244] =
    {field_update, 2},
  [245] =
    {field_initializer, 0},
    {field_update, 2},
  [247] =
    {field_condition, 1},
    {field_initializer, 0},
  [249] =
    {field_body, 4},
    {field_condition, 2, .inherited = true},
    {field_initializer, 2, .inherited = true},
    {field_update, 2, .inherited = true},
  [253] =
    {field_body, 4},
    {field_declarator, 2, .inherited = true},
    {field_initializer, 2, .inherited = true},
    {field_right, 2, .inherited = true},
    {field_type, 2, .inherited = true},
  [258] =
    {field_value, 3},
  [259] =
    {field_operand, 1},
  [260] =
    {field_assembly_code, 2},
    {field_output_operands, 3},
  [262] =
    {field_assembly_code, 3},
  [263] =
    {field_default_type, 2},
  [264] =
    {field_default_value, 2},
    {field_type, 0, .inherited = true},
  [266] =
    {field_body, 2},
    {field_parameters, 1},
  [268] =
    {field_name, 1},
    {field_type, 3},
  [270] =
    {field_condition, 2},
  [271] =
    {field_length, 1},
  [272] =
    {field_arguments, 4},
    {field_declarator, 3},
    {field_placement, 1},
    {field_type, 2},
  [276] =
    {field_declarator, 1},
    {field_declarator, 2},
  [278] =
    {field_declarator, 1},
    {field_declarator, 2},
    {field_declarator, 3, .inherited = true},
    {field_type, 0, .inherited = true},
  [282] =
    {field_declarator, 4},
  [283] =
    {field_body, 4},
    {field_declarator, 3},
    {field_type, 1, .inherited = true},
  [286] =
    {field_declarator, 0},
    {field_size, 3},
  [288] =
    {field_alternative, 4},
    {field_condition, 0},
    {field_consequence, 2},
  [291] =
    {field_size, 2},
  [292] =
    {field_body, 4},
    {field_captures, 0},
    {field_constraint, 2},
    {field_declarator, 3},
    {field_template_parameters, 1},
  [297] =
    {field_declarator, 1},
    {field_default_value, 3},
  [299] =
    {field_alternative, 4},
    {field_condition, 1},
  [301] =
    {field_arguments, 5},
    {field_declarator, 4},
    {field_placement, 2},
    {field_type, 3},
  [305] =
    {field_body, 2},
    {field_filter, 1},
  [307] =
    {field_left, 2},
    {field_right, 4},
  [309] =
    {field_declarator, 1},
    {field_default_value, 2},
    {field_type, 0, .inherited = true},
  [312] =
    {field_declarator, 1},
    {field_declarator, 2, .inherited = true},
    {field_default_value, 2, .inherited = true},
    {field_type, 0, .inherited = true},
  [316] =
    {field_declarator, 0, .inherited = true},
    {field_declarator, 1, .inherited = true},
    {field_default_value, 0, .inherited = true},
    {field_default_value, 1, .inherited = true},
  [320] =
    {field_declarator, 1},
    {field_type, 0, .inherited = true},
    {field_value, 3},
  [323] =
    {field_condition, 1},
    {field_update, 3},
  [325] =
    {field_condition, 1},
    {field_initializer, 0},
    {field_update, 3},
  [328] =
    {field_declarator, 1},
    {field_right, 3},
    {field_type, 0, .inherited = true},
  [331] =
    {field_initializer, 0},
    {field_update, 3},
  [333] =
    {field_condition, 2},
    {field_initializer, 0},
  [335] =
    {field_member, 4},
    {field_type, 2},
  [337] =
    {field_operand, 1},
    {field_operand, 2, .inherited = true},
  [339] =
    {field_assembly_code, 2},
    {field_input_operands, 4},
    {field_output_operands, 3},
  [342] =
    {field_assembly_code, 3},
    {field_output_operands, 4},
  [344] =
    {field_default_type, 3},
    {field_name, 1},
  [346] =
    {field_declarator, 1},
    {field_default_value, 3},
    {field_type, 0, .inherited = true},
  [349] =
    {field_name, 1},
    {field_type, 4},
  [351] =
    {field_end, 3},
    {field_start, 1},
  [353] =
    {field_declarator, 1},
    {field_default_value, 2},
  [355] =
    {field_declarator, 1},
    {field_declarator, 3, .inherited = true},
    {field_default_value, 3, .inherited = true},
    {field_type, 0, .inherited = true},
  [359] =
    {field_declarator, 1},
    {field_declarator, 3, .inherited = true},
    {field_default_value, 2},
    {field_default_value, 3, .inherited = true},
    {field_type, 0, .inherited = true},
  [364] =
    {field_condition, 2},
    {field_initializer, 0},
    {field_update, 4},
  [367] =
    {field_declarator, 2},
    {field_initializer, 0},
    {field_right, 4},
    {field_type, 1, .inherited = true},
  [371] =
    {field_operand, 0, .inherited = true},
    {field_operand, 1, .inherited = true},
  [373] =
    {field_register, 1},
  [374] =
    {field_assembly_code, 2},
    {field_clobbers, 5},
    {field_input_operands, 4},
    {field_output_operands, 3},
  [378] =
    {field_assembly_code, 3},
    {field_input_operands, 5},
    {field_output_operands, 4},
  [381] =
    {field_condition, 2},
    {field_message, 4},
  [383] =
    {field_delimiter, 1},
  [384] =
    {field_declarator, 1},
    {field_declarator, 4, .inherited = true},
    {field_default_value, 3},
    {field_default_value, 4, .inherited = true},
    {field_type, 0, .inherited = true},
  [389] =
    {field_constraint, 0},
    {field_value, 2},
  [391] =
    {field_register, 1},
    {field_register, 2, .inherited = true},
  [393] =
    {field_assembly_code, 2},
    {field_clobbers, 5},
    {field_goto_labels, 6},
    {field_input_operands, 4},
    {field_output_operands, 3},
  [398] =
    {field_assembly_code, 3},
    {field_clobbers, 6},
    {field_input_operands, 5},
    {field_output_operands, 4},
  [402] =
    {field_register, 0, .inherited = true},
    {field_register, 1, .inherited = true},
  [404] =
    {field_label, 1},
    {field_label, 2, .inherited = true},
  [406] =
    {field_assembly_code, 3},
    {field_clobbers, 6},
    {field_goto_labels, 7},
    {field_input_operands, 5},
    {field_output_operands, 4},
  [411] =
    {field_label, 0, .inherited = true},
    {field_label, 1, .inherited = true},
  [413] =
    {field_constraint, 3},
    {field_symbol, 1},
    {field_value, 5},
};

static const TSSymbol ts_alias_sequences[PRODUCTION_ID_COUNT][MAX_ALIAS_SEQUENCE_LENGTH] = {
  [0] = {0},
  [1] = {
    [0] = alias_sym_type_identifier,
  },
  [15] = {
    [0] = alias_sym_namespace_identifier,
  },
  [16] = {
    [0] = alias_sym_type_identifier,
  },
  [18] = {
    [0] = alias_sym_type_identifier,
  },
  [19] = {
    [0] = alias_sym_namespace_identifier,
  },
  [35] = {
    [1] = alias_sym_type_identifier,
  },
  [36] = {
    [1] = alias_sym_type_identifier,
  },
  [42] = {
    [0] = sym_primitive_type,
  },
  [43] = {
    [0] = sym_pointer_declarator,
  },
  [53] = {
    [1] = alias_sym_statement_identifier,
  },
  [55] = {
    [0] = alias_sym_statement_identifier,
  },
  [59] = {
    [1] = alias_sym_namespace_identifier,
  },
  [60] = {
    [1] = alias_sym_namespace_identifier,
  },
  [64] = {
    [0] = alias_sym_simple_requirement,
  },
  [70] = {
    [2] = alias_sym_field_identifier,
  },
  [75] = {
    [2] = alias_sym_type_identifier,
  },
  [93] = {
    [2] = alias_sym_namespace_identifier,
  },
  [96] = {
    [1] = alias_sym_field_identifier,
  },
  [99] = {
    [1] = alias_sym_type_identifier,
  },
  [101] = {
    [0] = alias_sym_field_identifier,
  },
  [108] = {
    [1] = alias_sym_type_identifier,
  },
  [119] = {
    [0] = alias_sym_field_identifier,
  },
  [120] = {
    [1] = alias_sym_field_identifier,
  },
  [137] = {
    [3] = alias_sym_namespace_identifier,
  },
  [139] = {
    [0] = alias_sym_field_identifier,
  },
  [158] = {
    [2] = alias_sym_type_identifier,
  },
  [162] = {
    [1] = alias_sym_namespace_identifier,
    [3] = alias_sym_namespace_identifier,
  },
  [163] = {
    [1] = alias_sym_namespace_identifier,
  },
  [164] = {
    [1] = alias_sym_type_identifier,
  },
  [190] = {
    [4] = alias_sym_field_identifier,
  },
  [194] = {
    [1] = alias_sym_type_identifier,
  },
  [196] = {
    [1] = alias_sym_type_identifier,
  },
};

static const uint16_t ts_non_terminal_alias_map[] = {
  sym_pointer_type_declarator, 2,
    sym_pointer_type_declarator,
    sym_pointer_declarator,
  sym_expression_statement, 2,
    sym_expression_statement,
    alias_sym_simple_requirement,
  0,
};

// Filled from tables.zlib by grammar.go, see fetchgrammar.
static TSStateId ts_primary_state_ids[STATE_COUNT];

static TSCharacterRange sym_identifier_character_set_1[] = {
  {'$', '$'}, {'A', 'Z'}, {'\\', '\\'}, {'_', '_'}, {'a', 'z'}, {0xaa, 0xaa}, {0xb5, 0xb5}, {0xba, 0xba},
  {0xc0, 0xd6}, {0xd8, 0xf6}, {0xf8, 0x2c1}, {0x2c6, 0x2d1}, {0x2e0, 0x2e4}, {0x2ec, 0x2ec}, {0x2ee, 0x2ee}, {0x370, 0x374},
  {0x376, 0x377}, {0x37b, 0x37d}, {0x37f, 0x37f}, {0x386, 0x386}, {0x388, 0x38a}, {0x38c, 0x38c}, {0x38e, 0x3a1}, {0x3a3, 0x3f5},
  {0x3f7, 0x481}, {0x48a, 0x52f}, {0x531, 0x556}, {0x559, 0x559}, {0x560, 0x588}, {0x5d0, 0x5ea}, {0x5ef, 0x5f2}, {0x620, 0x64a},
  {0x66e, 0x66f}, {0x671, 0x6d3}, {0x6d5, 0x6d5}, {0x6e5, 0x6e6}, {0x6ee, 0x6ef}, {0x6fa, 0x6fc}, {0x6ff, 0x6ff}, {0x710, 0x710},
  {0x712, 0x72f}, {0x74d, 0x7a5}, {0x7b1, 0x7b1}, {0x7ca, 0x7ea}, {0x7f4, 0x7f5}, {0x7fa, 0x7fa}, {0x800, 0x815}, {0x81a, 0x81a},
  {0x824, 0x824}, {0x828, 0x828}, {0x840, 0x858}, {0x860, 0x86a}, {0x870, 0x887}, {0x889, 0x88e}, {0x8a0, 0x8c9}, {0x904, 0x939},
  {0x93d, 0x93d}, {0x950, 0x950}, {0x958, 0x961}, {0x971, 0x980}, {0x985, 0x98c}, {0x98f, 0x990}, {0x993, 0x9a8}, {0x9aa, 0x9b0},
  {0x9b2, 0x9b2}, {0x9b6, 0x9b9}, {0x9bd, 0x9bd}, {0x9ce, 0x9ce}, {0x9dc, 0x9dd}, {0x9df, 0x9e1}, {0x9f0, 0x9f1}, {0x9fc, 0x9fc},
  {0xa05, 0xa0a}, {0xa0f, 0xa10}, {0xa13, 0xa28}, {0xa2a, 0xa30}, {0xa32, 0xa33}, {0xa35, 0xa36}, {0xa38, 0xa39}, {0xa59, 0xa5c},
  {0xa5e, 0xa5e}, {0xa72, 0xa74}, {0xa85, 0xa8d}, {0xa8f, 0xa91}, {0xa93, 0xaa8}, {0xaaa, 0xab0}, {0xab2, 0xab3}, {0xab5, 0xab9},
  {0xabd, 0xabd}, {0xad0, 0xad0}, {0xae0, 0xae1}, {0xaf9, 0xaf9}, {0xb05, 0xb0c}, {0xb0f, 0xb10}, {0xb13, 0xb28}, {0xb2a, 0xb30},
  {0xb32, 0xb33}, {0xb35, 0xb39}, {0xb3d, 0xb3d}, {0xb5c, 0xb5d}, {0xb5f, 0xb61}, {0xb71, 0xb71}, {0xb83, 0xb83}, {0xb85, 0xb8a},
  {0xb8e, 0xb90}, {0xb92, 0xb95}, {0xb99, 0xb9a}, {0xb9c, 0xb9c}, {0xb9e, 0xb9f}, {0xba3, 0xba4}, {0xba8, 0xbaa}, {0xbae, 0xbb9},
  {0xbd0, 0xbd0}, {0xc05, 0xc0c}, {0xc0e, 0xc10}, {0xc12, 0xc28}, {0xc2a, 0xc39}, {0xc3d, 0xc3d}, {0xc58, 0xc5a}, {0xc5d, 0xc5d},
  {0xc60, 0xc61}, {0xc80, 0xc80}, {0xc85, 0xc8c}, {0xc8e, 0xc90}, {0xc92, 0xca8}, {0xcaa, 0xcb3}, {0xcb5, 0xcb9}, {0xcbd, 0xcbd},
  {0xcdd, 0xcde}, {0xce0, 0xce1}, {0xcf1, 0xcf2}, {0xd04, 0xd0c}, {0xd0e, 0xd10}, {0xd12, 0xd3a}, {0xd3d, 0xd3d}, {0xd4e, 0xd4e},
  {0xd54, 0xd56}, {0xd5f, 0xd61}, {0xd7a, 0xd7f}, {0xd85, 0xd96}, {0xd9a, 0xdb1}, {0xdb3, 0xdbb}, {0xdbd, 0xdbd}, {0xdc0, 0xdc6},
  {0xe01, 0xe30}, {0xe32, 0xe32}, {0xe40, 0xe46}, {0xe81, 0xe82}, {0xe84, 0xe84}, {0xe86, 0xe8a}, {0xe8c, 0xea3}, {0xea5, 0xea5},
  {0xea7, 0xeb0}, {0xeb2, 0xeb2}, {0xebd, 0xebd}, {0xec0, 0xec4}, {0xec6, 0xec6}, {0xedc, 0xedf}, {0xf00, 0xf00}, {0xf40, 0xf47},
  {0xf49, 0xf6c}, {0xf88, 0xf8c}, {0x1000, 0x102a}, {0x103f, 0x103f}, {0x1050, 0x1055}, {0x105a, 0x105d}, {0x1061, 0x1061}, {0x1065, 0x1066},
  {0x106e, 0x1070}, {0x1075, 0x1081}, {0x108e, 0x108e}, {0x10a0, 0x10c5}, {0x10c7, 0x10c7}, {0x10cd, 0x10cd}, {0x10d0, 0x10fa}, {0x10fc, 0x1248},
  {0x124a, 0x124d}, {0x1250, 0x1256}, {0x1258, 0x1258}, {0x125a, 0x125d}, {0x1260, 0x1288}, {0x128a, 0x128d}, {0x1290, 0x12b0}, {0x12b2, 0x12b5},
  {0x12b8, 0x12be}, {0x12c0, 0x12c0}, {0x12c2, 0x12c5}, {0x12c8, 0x12d6}, {0x12d8, 0x1310}, {0x1312, 0x1315}, {0x1318, 0x135a}, {0x1380, 0x138f},
  {0x13a0, 0x13f5}, {0x13f8, 0x13fd}, {0x1401, 0x166c}, {0x166f, 0x167f}, {0x1681, 0x169a}, {0x16a0, 0x16ea}, {0x16ee, 0x16f8}, {0x1700, 0x1711},
  {0x171f, 0x1731}, {0x1740, 0x1751}, {0x1760, 0x176c}, {0x176e, 0x1770}, {0x1780, 0x17b3}, {0x17d7, 0x17d7}, {0x17dc, 0x17dc}, {0x1820, 0x1878},
  {0x1880, 0x18a8}, {0x18aa, 0x18aa}, {0x18b0, 0x18f5}, {0x1900, 0x191e}, {0x1950, 0x196d}, {0x1970, 0x1974}, {0x1980, 0x19ab}, {0x19b0, 0x19c9},
  {0x1a00, 0x1a16}, {0x1a20, 0x1a54}, {0x1aa7, 0x1aa7}, {0x1b05, 0x1b33}, {0x1b45, 0x1b4c}, {0x1b83, 0x1ba0}, {0x1bae, 0x1baf}, {0x1bba, 0x1be5},
  {0x1c00, 0x1c23}, {0x1c4d, 0x1c4f}, {0x1c5a, 0x1c7d}, {0x1c80, 0x1c88}, {0x1c90, 0x1cba}, {0x1cbd, 0x1cbf}, {0x1ce9, 0x1cec}, {0x1cee, 0x1cf3},
  {0x1cf5, 0x1cf6}, {0x1cfa, 0x1cfa}, {0x1d00, 0x1dbf}, {0x1e00, 0x1f15}, {0x1f18, 0x1f1d}, {0x1f20, 0x1f45}, {0x1f48, 0x1f4d}, {0x1f50, 0x1f57},
  {0x1f59, 0x1f59}, {0x1f5b, 0x1f5b}, {0x1f5d, 0x1f5d}, {0x1f5f, 0x1f7d}, {0x1f80, 0x1fb4}, {0x1fb6, 0x1fbc}, {0x1fbe, 0x1fbe}, {0x1fc2, 0x1fc4},
  {0x1fc6, 0x1fcc}, {0x1fd0, 0x1fd3}, {0x1fd6, 0x1fdb}, {0x1fe0, 0x1fec}, {0x1ff2, 0x1ff4}, {0x1ff6, 0x1ffc}, {0x2071, 0x2071}, {0x207f, 0x207f},
  {0x2090, 0x209c}, {0x2102, 0x2102}, {0x2107, 0x2107}, {0x210a, 0x2113}, {0x2115, 0x2115}, {0x2118, 0x211d}, {0x2124, 0x2124}, {0x2126, 0x2126},
  {0x2128, 0x2128}, {0x212a, 0x2139}, {0x213c, 0x213f}, {0x2145, 0x2149}, {0x214e, 0x214e}, {0x2160, 0x2188}, {0x2c00, 0x2ce4}, {0x2ceb, 0x2cee},
  {0x2cf2, 0x2cf3}, {0x2d00, 0x2d25}, {0x2d27, 0x2d27}, {0x2d2d, 0x2d2d}, {0x2d30, 0x2d67}, {0x2d6f, 0x2d6f}, {0x2d80, 0x2d96}, {0x2da0, 0x2da6},
  {0x2da8, 0x2dae}, {0x2db0, 0x2db6}, {0x2db8, 0x2dbe}, {0x2dc0, 0x2dc6}, {0x2dc8, 0x2dce}, {0x2dd0, 0x2dd6}, {0x2dd8, 0x2dde}, {0x3005, 0x3007},
  {0x3021, 0x3029}, {0x3031, 0x3035}, {0x3038, 0x303c}, {0x3041, 0x3096}, {0x309d, 0x309f}, {0x30a1, 0x30fa}, {0x30fc, 0x30ff}, {0x3105, 0x312f},
  {0x3131, 0x318e}, {0x31a0, 0x31bf}, {0x31f0, 0x31ff}, {0x3400, 0x4dbf}, {0x4e00, 0xa48c}, {0xa4d0, 0xa4fd}, {0xa500, 0xa60c}, {0xa610, 0xa61f},
  {0xa62a, 0xa62b}, {0xa640, 0xa66e}, {0xa67f, 0xa69d}, {0xa6a0, 0xa6ef}, {0xa717, 0xa71f}, {0xa722, 0xa788}, {0xa78b, 0xa7ca}, {0xa7d0, 0xa7d1},
  {0xa7d3, 0xa7d3}, {0xa7d5, 0xa7d9}, {0xa7f2, 0xa801}, {0xa803, 0xa805}, {0xa807, 0xa80a}, {0xa80c, 0xa822}, {0xa840, 0xa873}, {0xa882, 0xa8b3},
  {0xa8f2, 0xa8f7}, {0xa8fb, 0xa8fb}, {0xa8fd, 0xa8fe}, {0xa90a, 0xa925}, {0xa930, 0xa946}, {0xa960, 0xa97c}, {0xa984, 0xa9b2}, {0xa9cf, 0xa9cf},
  {0xa9e0, 0xa9e4}, {0xa9e6, 0xa9ef}, {0xa9fa, 0xa9fe}, {0xaa00, 0xaa28}, {0xaa40, 0xaa42}, {0xaa44, 0xaa4b}, {0xaa60, 0xaa76}, {0xaa7a, 0xaa7a},
  {0xaa7e, 0xaaaf}, {0xaab1, 0xaab1}, {0xaab5, 0xaab6}, {0xaab9, 0xaabd}, {0xaac0, 0xaac0}, {0xaac2, 0xaac2}, {0xaadb, 0xaadd}, {0xaae0, 0xaaea},
  {0xaaf2, 0xaaf4}, {0xab01, 0xab06}, {0xab09, 0xab0e}, {0xab11, 0xab16}, {0xab20, 0xab26}, {0xab28, 0xab2e}, {0xab30, 0xab5a}, {0xab5c, 0xab69},
  {0xab70, 0xabe2}, {0xac00, 0xd7a3}, {0xd7b0, 0xd7c6}, {0xd7cb, 0xd7fb}, {0xf900, 0xfa6d}, {0xfa70, 0xfad9}, {0xfb00, 0xfb06}, {0xfb13, 0xfb17},
  {0xfb1d, 0xfb1d}, {0xfb1f, 0xfb28}, {0xfb2a, 0xfb36}, {0xfb38, 0xfb3c}, {0xfb3e, 0xfb3e}, {0xfb40, 0xfb41}, {0xfb43, 0xfb44}, {0xfb46, 0xfbb1},
  {0xfbd3, 0xfc5d}, {0xfc64, 0xfd3d}, {0xfd50, 0xfd8f}, {0xfd92, 0xfdc7}, {0xfdf0, 0xfdf9}, {0xfe71, 0xfe71}, {0xfe73, 0xfe73}, {0xfe77, 0xfe77},
  {0xfe79, 0xfe79}, {0xfe7b, 0xfe7b}, {0xfe7d, 0xfe7d}, {0xfe7f, 0xfefc}, {0xff21, 0xff3a}, {0xff41, 0xff5a}, {0xff66, 0xff9d}, {0xffa0, 0xffbe},
  {0xffc2, 0xffc7}, {0xffca, 0xffcf}, {0xffd2, 0xffd7}, {0xffda, 0xffdc}, {0x10000, 0x1000b}, {0x1000d, 0x10026}, {0x10028, 0x1003a}, {0x1003c, 0x1003d},
  {0x1003f, 0x1004d}, {0x10050, 0x1005d}, {0x10080, 0x100fa}, {0x10140, 0x10174}, {0x10280, 0x1029c}, {0x102a0, 0x102d0}, {0x10300, 0x1031f}, {0x1032d, 0x1034a},
  {0x10350, 0x10375}, {0x10380, 0x1039d}, {0x103a0, 0x103c3}, {0x103c8, 0x103cf}, {0x103d1, 0x103d5}, {0x10400, 0x1049d}, {0x104b0, 0x104d3}, {0x104d8, 0x104fb},
  {0x10500, 0x10527}, {0x10530, 0x10563}, {0x10570, 0x1057a}, {0x1057c, 0x1058a}, {0x1058c, 0x10592}, {0x10594, 0x10595}, {0x10597, 0x105a1}, {0x105a3, 0x105b1},
  {0x105b3, 0x105b9}, {0x105bb, 0x105bc}, {0x10600, 0x10736}, {0x10740, 0x10755}, {0x10760, 0x10767}, {0x10780, 0x10785}, {0x10787, 0x107b0}, {0x107b2, 0x107ba},
  {0x10800, 0x10805}, {0x10808, 0x10808}, {0x1080a, 0x10835}, {0x10837, 0x10838}, {0x1083c, 0x1083c}, {0x1083f, 0x10855}, {0x10860, 0x10876}, {0x10880, 0x1089e},
  {0x108e0, 0x108f2}, {0x108f4, 0x108f5}, {0x10900, 0x10915}, {0x10920, 0x10939}, {0x10980, 0x109b7}, {0x109be, 0x109bf}, {0x10a00, 0x10a00}, {0x10a10, 0x10a13},
  {0x10a15, 0x10a17}, {0x10a19, 0x10a35}, {0x10a60, 0x10a7c}, {0x10a80, 0x10a9c}, {0x10ac0, 0x10ac7}, {0x10ac9, 0x10ae4}, {0x10b00, 0x10b35}, {0x10b40, 0x10b55},
  {0x10b60, 0x10b72}, {0x10b80, 0x10b91}, {0x10c00, 0x10c48}, {0x10c80, 0x10cb2}, {0x10cc0, 0x10cf2}, {0x10d00, 0x10d23}, {0x10e80, 0x10ea9}, {0x10eb0, 0x10eb1},
  {0x10f00, 0x10f1c}, {0x10f27, 0x10f27}, {0x10f30, 0x10f45}, {0x10f70, 0x10f81}, {0x10fb0, 0x10fc4}, {0x10fe0, 0x10ff6}, {0x11003, 0x11037}, {0x11071, 0x11072},
  {0x11075, 0x11075}, {0x11083, 0x110af}, {0x110d0, 0x110e8}, {0x11103, 0x11126}, {0x11144, 0x11144}, {0x11147, 0x11147}, {0x11150, 0x11172}, {0x11176, 0x11176},
  {0x11183, 0x111b2}, {0x111c1, 0x111c4}, {0x111da, 0x111da}, {0x111dc, 0x111dc}, {0x11200, 0x11211}, {0x11213, 0x1122b}, {0x1123f, 0x11240}, {0x11280, 0x11286},
  {0x11288, 0x11288}, {0x1128a, 0x1128d}, {0x1128f, 0x1129d}, {0x1129f, 0x112a8}, {0x112b0, 0x112de}, {0x11305, 0x1130c}, {0x1130f, 0x11310}, {0x11313, 0x11328},
  {0x1132a, 0x11330}, {0x11332, 0x11333}, {0x11335, 0x11339}, {0x1133d, 0x1133d}, {0x11350, 0x11350}, {0x1135d, 0x11361}, {0x11400, 0x11434}, {0x11447, 0x1144a},
  {0x1145f, 0x11461}, {0x11480, 0x114af}, {0x114c4, 0x114c5}, {0x114c7, 0x114c7}, {0x11580, 0x115ae}, {0x115d8, 0x115db}, {0x11600, 0x1162f}, {0x11644, 0x11644},
  {0x11680, 0x116aa}, {0x116b8, 0x116b8}, {0x11700, 0x1171a}, {0x11740, 0x11746}, {0x11800, 0x1182b}, {0x118a0, 0x118df}, {0x118ff, 0x11906}, {0x11909, 0x11909},
  {0x1190c, 0x11913}, {0x11915, 0x11916}, {0x11918, 0x1192f}, {0x1193f, 0x1193f}, {0x11941, 0x11941}, {0x119a0, 0x119a7}, {0x119aa, 0x119d0}, {0x119e1, 0x119e1},
  {0x119e3, 0x119e3}, {0x11a00, 0x11a00}, {0x11a0b, 0x11a32}, {0x11a3a, 0x11a3a}, {0x11a50, 0x11a50}, {0x11a5c, 0x11a89}, {0x11a9d, 0x11a9d}, {0x11ab0, 0x11af8},
  {0x11c00, 0x11c08}, {0x11c0a, 0x11c2e}, {0x11c40, 0x11c40}, {0x11c72, 0x11c8f}, {0x11d00, 0x11d06}, {0x11d08, 0x11d09}, {0x11d0b, 0x11d30}, {0x11d46, 0x11d46},
  {0x11d60, 0x11d65}, {0x11d67, 0x11d68}, {0x11d6a, 0x11d89}, {0x11d98, 0x11d98}, {0x11ee0, 0x11ef2}, {0x11f02, 0x11f02}, {0x11f04, 0x11f10}, {0x11f12, 0x11f33},
  {0x11fb0, 0x11fb0}, {0x12000, 0x12399}, {0x12400, 0x1246e}, {0x12480, 0x12543}, {0x12f90, 0x12ff0}, {0x13000, 0x1342f}, {0x13441, 0x13446}, {0x14400, 0x14646},
  {0x16800, 0x16a38}, {0x16a40, 0x16a5e}, {0x16a70, 0x16abe}, {0x16ad0, 0x16aed}, {0x16b00, 0x16b2f}, {0x16b40, 0x16b43}, {0x16b63, 0x16b77}, {0x16b7d, 0x16b8f},
  {0x16e40, 0x16e7f}, {0x16f00, 0x16f4a}, {0x16f50, 0x16f50}, {0x16f93, 0x16f9f}, {0x16fe0, 0x16fe1}, {0x16fe3, 0x16fe3}, {0x17000, 0x187f7}, {0x18800, 0x18cd5},
  {0x18d00, 0x18d08}, {0x1aff0, 0x1aff3}, {0x1aff5, 0x1affb}, {0x1affd, 0x1affe}, {0x1b000, 0x1b122}, {0x1b132, 0x1b132}, {0x1b150, 0x1b152}, {0x1b155, 0x1b155},
  {0x1b164, 0x1b167}, {0x1b170, 0x1b2fb}, {0x1bc00, 0x1bc6a}, {0x1bc70, 0x1bc7c}, {0x1bc80, 0x1bc88}, {0x1bc90, 0x1bc99}, {0x1d400, 0x1d454}, {0x1d456, 0x1d49c},
  {0x1d49e, 0x1d49f}, {0x1d4a2, 0x1d4a2}, {0x1d4a5, 0x1d4a6}, {0x1d4a9, 0x1d4ac}, {0x1d4ae, 0x1d4b9}, {0x1d4bb, 0x1d4bb}, {0x1d4bd, 0x1d4c3}, {0x1d4c5, 0x1d505},
  {0x1d507, 0x1d50a}, {0x1d50d, 0x1d514}, {0x1d516, 0x1d51c}, {0x1d51e, 0x1d539}, {0x1d53b, 0x1d53e}, {0x1d540, 0x1d544}, {0x1d546, 0x1d546}, {0x1d54a, 0x1d550},
  {0x1d552, 0x1d6a5}, {0x1d6a8, 0x1d6c0}, {0x1d6c2, 0x1d6da}, {0x1d6dc, 0x1d6fa}, {0x1d6fc, 0x1d714}, {0x1d716, 0x1d734}, {0x1d736, 0x1d74e}, {0x1d750, 0x1d76e},
  {0x1d770, 0x1d788}, {0x1d78a, 0x1d7a8}, {0x1d7aa, 0x1d7c2}, {0x1d7c4, 0x1d7cb}, {0x1df00, 0x1df1e}, {0x1df25, 0x1df2a}, {0x1e030, 0x1e06d}, {0x1e100, 0x1e12c},
  {0x1e137, 0x1e13d}, {0x1e14e, 0x1e14e}, {0x1e290, 0x1e2ad}, {0x1e2c0, 0x1e2eb}, {0x1e4d0, 0x1e4eb}, {0x1e7e0, 0x1e7e6}, {0x1e7e8, 0x1e7eb}, {0x1e7ed, 0x1e7ee},
  {0x1e7f0, 0x1e7fe}, {0x1e800, 0x1e8c4}, {0x1e900, 0x1e943}, {0x1e94b, 0x1e94b}, {0x1ee00, 0x1ee03}, {0x1ee05, 0x1ee1f}, {0x1ee21, 0x1ee22}, {0x1ee24, 0x1ee24},
  {0x1ee27, 0x1ee27}, {0x1ee29, 0x1ee32}, {0x1ee34, 0x1ee37}, {0x1ee39, 0x1ee39}, {0x1ee3b, 0x1ee3b}, {0x1ee42, 0x1ee42}, {0x1ee47, 0x1ee47}, {0x1ee49, 0x1ee49},
  {0x1ee4b, 0x1ee4b}, {0x1ee4d, 0x1ee4f}, {0x1ee51, 0x1ee52}, {0x1ee54, 0x1ee54}, {0x1ee57, 0x1ee57}, {0x1ee59, 0x1ee59}, {0x1ee5b, 0x1ee5b}, {0x1ee5d, 0x1ee5d},
  {0x1ee5f, 0x1ee5f}, {0x1ee61, 0x1ee62}, {0x1ee64, 0x1ee64}, {0x1ee67, 0x1ee6a}, {0x1ee6c, 0x1ee72}, {0x1ee74, 0x1ee77}, {0x1ee79, 0x1ee7c}, {0x1ee7e, 0x1ee7e},
  {0x1ee80, 0x1ee89}, {0x1ee8b, 0x1ee9b}, {0x1eea1, 0x1eea3}, {0x1eea5, 0x1eea9}, {0x1eeab, 0x1eebb}, {0x20000, 0x2a6df}, {0x2a700, 0x2b739}, {0x2b740, 0x2b81d},
  {0x2b820, 0x2cea1}, {0x2ceb0, 0x2ebe0}, {0x2ebf0, 0x2ee5d}, {0x2f800, 0x2fa1d}, {0x30000, 0x3134a}, {0x31350, 0x323af},
};

static TSCharacterRange sym_identifier_character_set_2[] = {
  {'$', '$'}, {'0', '9'}, {'A', 'Z'}, {'\\', '\\'}, {'_', '_'}, {'a', 'z'}, {0xaa, 0xaa}, {0xb5, 0xb5},
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
      if (eof) ADVANCE(337);
      ADVANCE_MAP(
        '!', 404,
        '"', 522,
        '#', 283,
        '%', 428,
        '&', 438,
        '\'', 513,
        '(', 341,
        ')', 344,
        '*', 424,
        '+', 418,
        ',', 343,
        '-', 407,
        '.', 488,
        '/', 426,
        '0', 495,
        ':', 472,
        ';', 456,
        '<', 448,
        '=', 468,
        '>', 638,
        '?', 473,
        'F', 565,
        'L', 539,
        'R', 542,
        'T', 569,
        'U', 543,
        '[', 464,
        '\\', 2,
        ']', 466,
        '^', 435,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 574,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        't', 616,
        'u', 546,
        'v', 608,
        '{', 460,
        '|', 432,
        '}', 461,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(335);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(498);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 1:
      if (lookahead == '\n') SKIP(161);
      END_STATE();
    case 2:
      if (lookahead == '\n') SKIP(161);
      if (lookahead == '\r') SKIP(1);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 3:
      if (lookahead == '\n') SKIP(170);
      END_STATE();
    case 4:
      if (lookahead == '\n') SKIP(170);
      if (lookahead == '\r') SKIP(3);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 5:
      if (lookahead == '\n') SKIP(169);
      END_STATE();
    case 6:
      if (lookahead == '\n') SKIP(169);
      if (lookahead == '\r') SKIP(5);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 7:
      if (lookahead == '\n') SKIP(172);
      END_STATE();
    case 8:
      if (lookahead == '\n') SKIP(172);
      if (lookahead == '\r') SKIP(7);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 9:
      if (lookahead == '\n') SKIP(171);
      END_STATE();
    case 10:
      if (lookahead == '\n') SKIP(171);
      if (lookahead == '\r') SKIP(9);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 11:
      if (lookahead == '\n') SKIP(173);
      END_STATE();
    case 12:
      if (lookahead == '\n') SKIP(173);
      if (lookahead == '\r') SKIP(11);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 13:
      if (lookahead == '\n') SKIP(174);
      END_STATE();
    case 14:
      if (lookahead == '\n') SKIP(174);
      if (lookahead == '\r') SKIP(13);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 15:
      if (lookahead == '\n') SKIP(164);
      END_STATE();
    case 16:
      if (lookahead == '\n') SKIP(164);
      if (lookahead == '\r') SKIP(15);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 17:
      if (lookahead == '\n') SKIP(165);
      END_STATE();
    case 18:
      if (lookahead == '\n') SKIP(165);
      if (lookahead == '\r') SKIP(17);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 19:
      if (lookahead == '\n') SKIP(236);
      END_STATE();
    case 20:
      if (lookahead == '\n') SKIP(236);
      if (lookahead == '\r') SKIP(19);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 21:
      if (lookahead == '\n') SKIP(192);
      END_STATE();
    case 22:
      if (lookahead == '\n') SKIP(192);
      if (lookahead == '\r') SKIP(21);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 23:
      if (lookahead == '\n') SKIP(237);
      END_STATE();
    case 24:
      if (lookahead == '\n') SKIP(237);
      if (lookahead == '\r') SKIP(23);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 25:
      if (lookahead == '\n') SKIP(167);
      END_STATE();
    case 26:
      if (lookahead == '\n') SKIP(167);
      if (lookahead == '\r') SKIP(25);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 27:
      if (lookahead == '\n') SKIP(185);
      END_STATE();
    case 28:
      if (lookahead == '\n') SKIP(185);
      if (lookahead == '\r') SKIP(27);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 29:
      if (lookahead == '\n') SKIP(181);
      END_STATE();
    case 30:
      if (lookahead == '\n') SKIP(181);
      if (lookahead == '\r') SKIP(29);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 31:
      if (lookahead == '\n') SKIP(193);
      END_STATE();
    case 32:
      if (lookahead == '\n') SKIP(193);
      if (lookahead == '\r') SKIP(31);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 33:
      if (lookahead == '\n') SKIP(191);
      END_STATE();
    case 34:
      if (lookahead == '\n') SKIP(191);
      if (lookahead == '\r') SKIP(33);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 35:
      if (lookahead == '\n') SKIP(222);
      END_STATE();
    case 36:
      if (lookahead == '\n') SKIP(222);
      if (lookahead == '\r') SKIP(35);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 37:
      if (lookahead == '\n') SKIP(182);
      END_STATE();
    case 38:
      if (lookahead == '\n') SKIP(182);
      if (lookahead == '\r') SKIP(37);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 39:
      if (lookahead == '\n') SKIP(196);
      END_STATE();
    case 40:
      if (lookahead == '\n') SKIP(196);
      if (lookahead == '\r') SKIP(39);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 41:
      if (lookahead == '\n') SKIP(186);
      END_STATE();
    case 42:
      if (lookahead == '\n') SKIP(186);
      if (lookahead == '\r') SKIP(41);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 43:
      if (lookahead == '\n') SKIP(176);
      END_STATE();
    case 44:
      if (lookahead == '\n') SKIP(176);
      if (lookahead == '\r') SKIP(43);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 45:
      if (lookahead == '\n') SKIP(194);
      END_STATE();
    case 46:
      if (lookahead == '\n') SKIP(194);
      if (lookahead == '\r') SKIP(45);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 47:
      if (lookahead == '\n') SKIP(209);
      END_STATE();
    case 48:
      if (lookahead == '\n') SKIP(209);
      if (lookahead == '\r') SKIP(47);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 49:
      if (lookahead == '\n') SKIP(217);
      END_STATE();
    case 50:
      if (lookahead == '\n') SKIP(217);
      if (lookahead == '\r') SKIP(49);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 51:
      if (lookahead == '\n') SKIP(206);
      END_STATE();
    case 52:
      if (lookahead == '\n') SKIP(206);
      if (lookahead == '\r') SKIP(51);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 53:
      if (lookahead == '\n') SKIP(213);
      END_STATE();
    case 54:
      if (lookahead == '\n') SKIP(213);
      if (lookahead == '\r') SKIP(53);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 55:
      if (lookahead == '\n') SKIP(177);
      END_STATE();
    case 56:
      if (lookahead == '\n') SKIP(177);
      if (lookahead == '\r') SKIP(55);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 57:
      if (lookahead == '\n') SKIP(223);
      END_STATE();
    case 58:
      if (lookahead == '\n') SKIP(223);
      if (lookahead == '\r') SKIP(57);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 59:
      if (lookahead == '\n') SKIP(201);
      END_STATE();
    case 60:
      if (lookahead == '\n') SKIP(201);
      if (lookahead == '\r') SKIP(59);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 61:
      if (lookahead == '\n') SKIP(235);
      END_STATE();
    case 62:
      if (lookahead == '\n') SKIP(235);
      if (lookahead == '\r') SKIP(61);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 63:
      if (lookahead == '\n') SKIP(238);
      END_STATE();
    case 64:
      if (lookahead == '\n') SKIP(238);
      if (lookahead == '\r') SKIP(63);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 65:
      if (lookahead == '\n') SKIP(183);
      END_STATE();
    case 66:
      if (lookahead == '\n') SKIP(183);
      if (lookahead == '\r') SKIP(65);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 67:
      if (lookahead == '\n') SKIP(187);
      END_STATE();
    case 68:
      if (lookahead == '\n') SKIP(187);
      if (lookahead == '\r') SKIP(67);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 69:
      if (lookahead == '\n') SKIP(226);
      END_STATE();
    case 70:
      if (lookahead == '\n') SKIP(226);
      if (lookahead == '\r') SKIP(69);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 71:
      if (lookahead == '\n') SKIP(200);
      END_STATE();
    case 72:
      if (lookahead == '\n') SKIP(200);
      if (lookahead == '\r') SKIP(71);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 73:
      if (lookahead == '\n') SKIP(230);
      END_STATE();
    case 74:
      if (lookahead == '\n') SKIP(230);
      if (lookahead == '\r') SKIP(73);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 75:
      if (lookahead == '\n') SKIP(240);
      END_STATE();
    case 76:
      if (lookahead == '\n') SKIP(240);
      if (lookahead == '\r') SKIP(75);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 77:
      if (lookahead == '\n') SKIP(239);
      END_STATE();
    case 78:
      if (lookahead == '\n') SKIP(239);
      if (lookahead == '\r') SKIP(77);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 79:
      if (lookahead == '\n') SKIP(168);
      END_STATE();
    case 80:
      if (lookahead == '\n') SKIP(168);
      if (lookahead == '\r') SKIP(79);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 81:
      if (lookahead == '\n') SKIP(216);
      END_STATE();
    case 82:
      if (lookahead == '\n') SKIP(216);
      if (lookahead == '\r') SKIP(81);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 83:
      if (lookahead == '\n') SKIP(208);
      END_STATE();
    case 84:
      if (lookahead == '\n') SKIP(208);
      if (lookahead == '\r') SKIP(83);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 85:
      if (lookahead == '\n') SKIP(215);
      END_STATE();
    case 86:
      if (lookahead == '\n') SKIP(215);
      if (lookahead == '\r') SKIP(85);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 87:
      if (lookahead == '\n') SKIP(242);
      END_STATE();
    case 88:
      if (lookahead == '\n') SKIP(242);
      if (lookahead == '\r') SKIP(87);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 89:
      if (lookahead == '\n') SKIP(221);
      END_STATE();
    case 90:
      if (lookahead == '\n') SKIP(221);
      if (lookahead == '\r') SKIP(89);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 91:
      if (lookahead == '\n') SKIP(241);
      END_STATE();
    case 92:
      if (lookahead == '\n') SKIP(241);
      if (lookahead == '\r') SKIP(91);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 93:
      if (lookahead == '\n') SKIP(248);
      END_STATE();
    case 94:
      if (lookahead == '\n') SKIP(248);
      if (lookahead == '\r') SKIP(93);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 95:
      if (lookahead == '\n') SKIP(175);
      END_STATE();
    case 96:
      if (lookahead == '\n') SKIP(175);
      if (lookahead == '\r') SKIP(95);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 97:
      if (lookahead == '\n') SKIP(99);
      END_STATE();
    case 98:
      if (lookahead == '\n') SKIP(99);
      if (lookahead == '\r') SKIP(97);
      END_STATE();
    case 99:
      ADVANCE_MAP(
        '\n', 346,
        '!', 270,
        '%', 427,
        '&', 437,
        '(', 402,
        '*', 423,
        '+', 416,
        '-', 406,
        '/', 425,
        '<', 451,
        '=', 271,
        '>', 442,
      );
      if (lookahead == '\\') SKIP(98);
      if (lookahead == '^') ADVANCE(434);
      if (lookahead == '|') ADVANCE(433);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(99);
      END_STATE();
    case 100:
      if (lookahead == '\n') SKIP(243);
      END_STATE();
    case 101:
      if (lookahead == '\n') SKIP(243);
      if (lookahead == '\r') SKIP(100);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 102:
      if (lookahead == '\n') SKIP(247);
      END_STATE();
    case 103:
      if (lookahead == '\n') SKIP(247);
      if (lookahead == '\r') SKIP(102);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 104:
      if (lookahead == '\n') SKIP(231);
      END_STATE();
    case 105:
      if (lookahead == '\n') SKIP(231);
      if (lookahead == '\r') SKIP(104);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 106:
      if (lookahead == '\n') SKIP(232);
      END_STATE();
    case 107:
      if (lookahead == '\n') SKIP(232);
      if (lookahead == '\r') SKIP(106);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 108:
      if (lookahead == '\n') SKIP(233);
      if (lookahead == '"') ADVANCE(522);
      if (lookahead == '/') ADVANCE(523);
      if (lookahead == '\\') ADVANCE(109);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') ADVANCE(526);
      if (lookahead != 0) ADVANCE(527);
      END_STATE();
    case 109:
      if (lookahead == '\n') ADVANCE(529);
      if (lookahead == '\r') ADVANCE(528);
      if (lookahead == 'U') ADVANCE(333);
      if (lookahead == 'u') ADVANCE(325);
      if (lookahead == 'x') ADVANCE(319);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(531);
      if (lookahead != 0) ADVANCE(528);
      END_STATE();
    case 110:
      if (lookahead == '\n') SKIP(244);
      if (lookahead == '\'') ADVANCE(513);
      if (lookahead == '/') ADVANCE(516);
      if (lookahead == '\\') ADVANCE(515);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') ADVANCE(517);
      if (lookahead != 0) ADVANCE(514);
      END_STATE();
    case 111:
      if (lookahead == '\n') ADVANCE(339);
      if (lookahead == '\r') ADVANCE(115);
      if (lookahead == '(') ADVANCE(341);
      if (lookahead == '/') ADVANCE(367);
      if (lookahead == '\\') ADVANCE(362);
      if (('\t' <= lookahead && lookahead <= '\f') ||
          lookahead == ' ') SKIP(261);
      if (lookahead != 0) ADVANCE(369);
      END_STATE();
    case 112:
      if (lookahead == '\n') ADVANCE(339);
      if (lookahead == '\r') ADVANCE(115);
      if (lookahead == '/') ADVANCE(367);
      if (lookahead == '\\') ADVANCE(362);
      if (('\t' <= lookahead && lookahead <= '\f') ||
          lookahead == ' ') SKIP(261);
      if (lookahead != 0) ADVANCE(369);
      END_STATE();
    case 113:
      if (lookahead == '\n') ADVANCE(339);
      if (lookahead == '\r') ADVANCE(114);
      if (lookahead == '(') ADVANCE(402);
      if (lookahead == '/') ADVANCE(252);
      if (lookahead == '\\') SKIP(117);
      if (('\t' <= lookahead && lookahead <= '\f') ||
          lookahead == ' ') SKIP(249);
      END_STATE();
    case 114:
      if (lookahead == '\n') ADVANCE(339);
      if (lookahead == '(') ADVANCE(402);
      if (lookahead == '/') ADVANCE(252);
      if (lookahead == '\\') SKIP(117);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(249);
      END_STATE();
    case 115:
      if (lookahead == '\n') ADVANCE(339);
      if (lookahead == '/') ADVANCE(367);
      if (lookahead == '\\') ADVANCE(362);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(261);
      if (lookahead != 0) ADVANCE(369);
      END_STATE();
    case 116:
      if (lookahead == '\n') SKIP(249);
      END_STATE();
    case 117:
      if (lookahead == '\n') SKIP(249);
      if (lookahead == '\r') SKIP(116);
      END_STATE();
    case 118:
      if (lookahead == '\n') SKIP(162);
      END_STATE();
    case 119:
      if (lookahead == '\n') SKIP(162);
      if (lookahead == '\r') SKIP(118);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 120:
      if (lookahead == '\n') SKIP(166);
      END_STATE();
    case 121:
      if (lookahead == '\n') SKIP(166);
      if (lookahead == '\r') SKIP(120);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 122:
      if (lookahead == '\n') SKIP(202);
      END_STATE();
    case 123:
      if (lookahead == '\n') SKIP(202);
      if (lookahead == '\r') SKIP(122);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 124:
      if (lookahead == '\n') SKIP(197);
      END_STATE();
    case 125:
      if (lookahead == '\n') SKIP(197);
      if (lookahead == '\r') SKIP(124);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 126:
      if (lookahead == '\n') SKIP(227);
      END_STATE();
    case 127:
      if (lookahead == '\n') SKIP(227);
      if (lookahead == '\r') SKIP(126);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 128:
      if (lookahead == '\n') SKIP(207);
      END_STATE();
    case 129:
      if (lookahead == '\n') SKIP(207);
      if (lookahead == '\r') SKIP(128);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 130:
      if (lookahead == '\n') SKIP(179);
      END_STATE();
    case 131:
      if (lookahead == '\n') SKIP(179);
      if (lookahead == '\r') SKIP(130);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 132:
      if (lookahead == '\n') SKIP(224);
      END_STATE();
    case 133:
      if (lookahead == '\n') SKIP(224);
      if (lookahead == '\r') SKIP(132);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 134:
      if (lookahead == '\n') SKIP(189);
      END_STATE();
    case 135:
      if (lookahead == '\n') SKIP(189);
      if (lookahead == '\r') SKIP(134);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 136:
      if (lookahead == '\n') SKIP(214);
      END_STATE();
    case 137:
      if (lookahead == '\n') SKIP(214);
      if (lookahead == '\r') SKIP(136);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 138:
      if (lookahead == '\n') SKIP(199);
      END_STATE();
    case 139:
      if (lookahead == '\n') SKIP(199);
      if (lookahead == '\r') SKIP(138);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 140:
      if (lookahead == '\n') SKIP(219);
      END_STATE();
    case 141:
      if (lookahead == '\n') SKIP(219);
      if (lookahead == '\r') SKIP(140);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 142:
      if (lookahead == '\n') SKIP(211);
      END_STATE();
    case 143:
      if (lookahead == '\n') SKIP(211);
      if (lookahead == '\r') SKIP(142);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 144:
      if (lookahead == '\n') SKIP(163);
      END_STATE();
    case 145:
      if (lookahead == '\n') SKIP(163);
      if (lookahead == '\r') SKIP(144);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 146:
      if (lookahead == '\n') SKIP(203);
      END_STATE();
    case 147:
      if (lookahead == '\n') SKIP(203);
      if (lookahead == '\r') SKIP(146);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 148:
      if (lookahead == '\n') SKIP(225);
      END_STATE();
    case 149:
      if (lookahead == '\n') SKIP(225);
      if (lookahead == '\r') SKIP(148);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 150:
      if (lookahead == '\n') SKIP(205);
      END_STATE();
    case 151:
      if (lookahead == '\n') SKIP(205);
      if (lookahead == '\r') SKIP(150);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 152:
      if (lookahead == '\n') SKIP(220);
      END_STATE();
    case 153:
      if (lookahead == '\n') SKIP(220);
      if (lookahead == '\r') SKIP(152);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 154:
      if (lookahead == '\n') SKIP(212);
      END_STATE();
    case 155:
      if (lookahead == '\n') SKIP(212);
      if (lookahead == '\r') SKIP(154);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 156:
      if (lookahead == '\n') SKIP(195);
      END_STATE();
    case 157:
      if (lookahead == '\n') SKIP(195);
      if (lookahead == '\r') SKIP(156);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 158:
      if (lookahead == '\n') SKIP(229);
      END_STATE();
    case 159:
      if (lookahead == '\n') SKIP(229);
      if (lookahead == '\r') SKIP(158);
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 160:
      if (lookahead == '\r') ADVANCE(637);
      if (lookahead == '\\') ADVANCE(631);
      if (lookahead != 0) ADVANCE(636);
      END_STATE();
    case 161:
      ADVANCE_MAP(
        '!', 404,
        '"', 522,
        '#', 283,
        '%', 428,
        '&', 438,
        '\'', 513,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 418,
        ',', 343,
        '-', 407,
        '.', 488,
        '/', 426,
        '0', 495,
        ':', 472,
        ';', 456,
        '<', 448,
        '=', 468,
        '>', 638,
        '?', 473,
        'F', 565,
        'L', 539,
        'R', 542,
        'T', 569,
        'U', 543,
        '[', 464,
        '\\', 2,
        ']', 466,
        '^', 435,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 574,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        't', 616,
        'u', 546,
        'v', 608,
        '{', 460,
        '|', 432,
        '}', 461,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(161);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(498);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 162:
      ADVANCE_MAP(
        '!', 404,
        '"', 522,
        '#', 292,
        '%', 428,
        '&', 438,
        '\'', 513,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 418,
        ',', 343,
        '-', 408,
        '.', 488,
        '/', 426,
        '0', 495,
        ':', 472,
        ';', 456,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        'F', 565,
        'L', 539,
        'R', 542,
        'T', 569,
        'U', 543,
        '[', 462,
        '\\', 119,
        ']', 466,
        '^', 435,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 574,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        't', 616,
        'u', 546,
        'v', 608,
        '{', 460,
        '|', 432,
        '}', 461,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(162);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(498);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 163:
      ADVANCE_MAP(
        '!', 404,
        '"', 522,
        '#', 292,
        '%', 427,
        '&', 437,
        '\'', 513,
        '(', 402,
        ')', 344,
        '*', 423,
        '+', 419,
        ',', 343,
        '-', 409,
        '.', 488,
        '/', 425,
        '0', 495,
        ':', 472,
        ';', 456,
        '<', 450,
        '=', 271,
        '>', 442,
        '?', 473,
        'F', 565,
        'L', 539,
        'R', 542,
        'T', 569,
        'U', 543,
        '[', 462,
        '\\', 145,
        ']', 466,
        '^', 434,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 574,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        't', 616,
        'u', 546,
        'v', 608,
        '{', 460,
        '|', 433,
        '}', 461,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(163);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(498);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 164:
      ADVANCE_MAP(
        '!', 404,
        '"', 522,
        '%', 428,
        '&', 438,
        '\'', 513,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 418,
        ',', 343,
        '-', 407,
        '.', 488,
        '/', 426,
        '0', 495,
        ':', 268,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        'F', 565,
        'L', 539,
        'R', 542,
        'T', 569,
        'U', 543,
        '[', 462,
        '\\', 16,
        '^', 435,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 574,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        't', 616,
        'u', 546,
        'v', 608,
        '{', 460,
        '|', 432,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(164);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(498);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 165:
      ADVANCE_MAP(
        '!', 404,
        '"', 522,
        '%', 428,
        '&', 438,
        '\'', 513,
        '(', 402,
        '*', 424,
        '+', 418,
        ',', 343,
        '-', 408,
        '.', 488,
        '/', 426,
        '0', 495,
        ':', 268,
        '<', 448,
        '=', 468,
        '>', 638,
        '?', 473,
        'F', 565,
        'L', 539,
        'R', 542,
        'T', 569,
        'U', 543,
        '[', 462,
        '\\', 18,
        '^', 435,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 574,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        't', 616,
        'u', 546,
        'v', 608,
        '{', 460,
        '|', 432,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(165);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(498);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 166:
      ADVANCE_MAP(
        '!', 404,
        '"', 522,
        '%', 427,
        '&', 437,
        '\'', 513,
        '(', 402,
        '*', 423,
        '+', 419,
        ',', 343,
        '-', 409,
        '.', 488,
        '/', 425,
        '0', 495,
        ':', 268,
        '<', 450,
        '=', 271,
        '>', 638,
        '?', 473,
        'F', 565,
        'L', 539,
        'R', 542,
        'T', 569,
        'U', 543,
        '[', 462,
        '\\', 121,
        '^', 434,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 574,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        't', 616,
        'u', 546,
        'v', 608,
        '{', 460,
        '|', 433,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(166);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(498);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 167:
      ADVANCE_MAP(
        '!', 404,
        '"', 234,
        '%', 428,
        '&', 438,
        '(', 250,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 411,
        '/', 426,
        ':', 268,
        '<', 448,
        '=', 468,
        '>', 443,
        '[', 280,
        '\\', 26,
        '^', 435,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 603,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        'u', 596,
        'v', 608,
        '|', 432,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(167);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 168:
      ADVANCE_MAP(
        '!', 404,
        '"', 234,
        '%', 428,
        '&', 438,
        '(', 250,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 411,
        '/', 426,
        '<', 448,
        '=', 468,
        '>', 443,
        '[', 281,
        '\\', 80,
        '^', 435,
        '|', 432,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(168);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 169:
      ADVANCE_MAP(
        '!', 403,
        '"', 522,
        '#', 283,
        '&', 437,
        '\'', 513,
        '(', 402,
        '*', 423,
        '+', 419,
        ',', 343,
        '-', 410,
        '.', 312,
        '/', 252,
        '0', 495,
        ':', 268,
        ';', 456,
        '<', 269,
        '>', 441,
        'F', 565,
        'L', 539,
        'R', 542,
        'T', 569,
        'U', 543,
        '[', 463,
        '\\', 6,
        ']', 282,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 574,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        't', 616,
        'u', 546,
        'v', 608,
        '{', 460,
        '|', 431,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(169);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(498);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 170:
      ADVANCE_MAP(
        '!', 403,
        '"', 522,
        '#', 287,
        '%', 427,
        '&', 437,
        '\'', 513,
        '(', 402,
        ')', 344,
        '*', 423,
        '+', 419,
        ',', 343,
        '-', 410,
        '.', 489,
        '/', 425,
        '0', 495,
        ':', 472,
        ';', 456,
        '<', 447,
        '=', 467,
        '>', 638,
        'F', 565,
        'L', 539,
        'R', 542,
        'T', 569,
        'U', 543,
        '[', 463,
        '\\', 4,
        ']', 466,
        '^', 434,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 574,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        't', 616,
        'u', 546,
        'v', 608,
        '{', 460,
        '|', 306,
        '}', 461,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(170);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(498);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 171:
      ADVANCE_MAP(
        '!', 403,
        '"', 522,
        '#', 291,
        '&', 436,
        '\'', 513,
        '(', 402,
        ')', 344,
        '*', 423,
        '+', 419,
        '-', 410,
        '.', 258,
        '/', 252,
        '0', 495,
        ':', 268,
        ';', 456,
        '>', 274,
        'F', 565,
        'L', 539,
        'R', 542,
        'T', 569,
        'U', 543,
        '[', 463,
        '\\', 10,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 574,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        't', 616,
        'u', 546,
        'v', 608,
        '{', 460,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(171);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(498);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 172:
      ADVANCE_MAP(
        '!', 403,
        '"', 522,
        '#', 285,
        '&', 437,
        '\'', 513,
        '(', 402,
        '*', 423,
        '+', 419,
        ',', 343,
        '-', 410,
        '.', 312,
        '/', 252,
        '0', 495,
        ':', 268,
        ';', 456,
        '>', 272,
        'F', 565,
        'L', 539,
        'R', 542,
        'T', 569,
        'U', 543,
        '[', 463,
        '\\', 8,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 574,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        't', 616,
        'u', 546,
        'v', 608,
        '{', 460,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(172);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(498);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 173:
      ADVANCE_MAP(
        '!', 403,
        '"', 522,
        '&', 437,
        '\'', 513,
        '(', 402,
        ')', 344,
        '*', 423,
        '+', 419,
        ',', 343,
        '-', 410,
        '.', 489,
        '/', 252,
        '0', 495,
        ':', 268,
        ';', 456,
        '<', 447,
        '=', 467,
        '>', 638,
        'F', 565,
        'L', 539,
        'R', 542,
        'T', 569,
        'U', 543,
        '[', 462,
        '\\', 12,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 574,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        't', 616,
        'u', 546,
        'v', 608,
        '{', 460,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(173);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(498);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 174:
      ADVANCE_MAP(
        '!', 403,
        '"', 522,
        '&', 436,
        '\'', 513,
        '(', 402,
        ')', 344,
        '*', 423,
        '+', 419,
        ',', 343,
        '-', 410,
        '.', 489,
        '/', 252,
        '0', 495,
        ':', 472,
        ';', 456,
        '<', 447,
        '=', 467,
        '>', 638,
        'F', 565,
        'L', 539,
        'R', 542,
        'T', 569,
        'U', 543,
        '[', 462,
        '\\', 14,
        ']', 466,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 574,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        't', 616,
        'u', 546,
        'v', 608,
        '{', 460,
        '}', 461,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(174);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(498);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 175:
      ADVANCE_MAP(
        '!', 403,
        '\'', 513,
        '(', 402,
        ')', 344,
        '+', 421,
        '-', 414,
        '.', 312,
        '/', 252,
        '0', 495,
        'L', 557,
        'U', 558,
        '\\', 96,
        'u', 559,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(175);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(498);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 176:
      ADVANCE_MAP(
        '!', 270,
        '"', 522,
        '#', 292,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 412,
        '.', 487,
        '/', 426,
        ':', 472,
        ';', 456,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        'L', 540,
        'R', 542,
        'U', 544,
        '[', 462,
        '\\', 44,
        ']', 466,
        '^', 435,
        'u', 547,
        '{', 460,
        '|', 432,
        '}', 461,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(176);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 177:
      ADVANCE_MAP(
        '!', 270,
        '"', 522,
        '#', 292,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 412,
        '.', 487,
        '/', 426,
        ':', 471,
        ';', 456,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        'L', 540,
        'R', 542,
        'U', 544,
        '[', 462,
        '\\', 56,
        ']', 466,
        '^', 435,
        'u', 547,
        '|', 432,
        '}', 461,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(177);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 178:
      ADVANCE_MAP(
        '!', 270,
        '"', 522,
        '#', 292,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 412,
        '.', 487,
        '/', 426,
        ':', 471,
        ';', 456,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        'L', 649,
        'R', 650,
        'U', 651,
        '[', 462,
        '\\', 56,
        ']', 466,
        '^', 435,
        'u', 652,
        '|', 432,
        '}', 461,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(177);
      if (('A' <= lookahead && lookahead <= '_') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(658);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 179:
      ADVANCE_MAP(
        '!', 270,
        '"', 522,
        '#', 292,
        '%', 427,
        '&', 437,
        '(', 402,
        ')', 344,
        '*', 423,
        '+', 417,
        ',', 343,
        '-', 413,
        '.', 487,
        '/', 425,
        ':', 471,
        ';', 456,
        '<', 450,
        '=', 271,
        '>', 442,
        '?', 473,
        'L', 540,
        'R', 542,
        'U', 544,
        '[', 462,
        '\\', 131,
        ']', 466,
        '^', 434,
        'u', 547,
        '|', 433,
        '}', 461,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(179);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 180:
      ADVANCE_MAP(
        '!', 270,
        '"', 522,
        '#', 292,
        '%', 427,
        '&', 437,
        '(', 402,
        ')', 344,
        '*', 423,
        '+', 417,
        ',', 343,
        '-', 413,
        '.', 487,
        '/', 425,
        ':', 471,
        ';', 456,
        '<', 450,
        '=', 271,
        '>', 442,
        '?', 473,
        'L', 649,
        'R', 650,
        'U', 651,
        '[', 462,
        '\\', 131,
        ']', 466,
        '^', 434,
        'u', 652,
        '|', 433,
        '}', 461,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(179);
      if (('A' <= lookahead && lookahead <= '_') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(658);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 181:
      ADVANCE_MAP(
        '!', 270,
        '"', 522,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 411,
        '.', 487,
        '/', 426,
        ':', 268,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        'L', 540,
        'R', 542,
        'U', 544,
        '[', 463,
        '\\', 30,
        '^', 435,
        'u', 547,
        '{', 460,
        '|', 432,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(181);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 182:
      ADVANCE_MAP(
        '!', 270,
        '"', 522,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 411,
        '.', 487,
        '/', 426,
        ':', 268,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        'L', 540,
        'R', 542,
        'U', 544,
        '[', 462,
        '\\', 38,
        '^', 435,
        'u', 547,
        '{', 460,
        '|', 432,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(182);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 183:
      ADVANCE_MAP(
        '!', 270,
        '"', 522,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 411,
        '.', 487,
        '/', 426,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        'L', 540,
        'R', 542,
        'U', 544,
        '[', 462,
        '\\', 66,
        '^', 435,
        'u', 547,
        '|', 432,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(183);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 184:
      ADVANCE_MAP(
        '!', 270,
        '"', 522,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 411,
        '.', 487,
        '/', 426,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        'L', 649,
        'R', 650,
        'U', 651,
        '[', 462,
        '\\', 66,
        '^', 435,
        'u', 652,
        '|', 432,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(183);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(658);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 185:
      ADVANCE_MAP(
        '!', 270,
        '"', 522,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 412,
        '.', 487,
        '/', 426,
        ':', 472,
        ';', 456,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        'L', 540,
        'R', 542,
        'U', 544,
        '[', 463,
        '\\', 28,
        '^', 435,
        'u', 547,
        '{', 460,
        '|', 432,
        '}', 461,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(185);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 186:
      ADVANCE_MAP(
        '!', 270,
        '"', 522,
        '%', 428,
        '&', 438,
        '(', 402,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 412,
        '.', 487,
        '/', 426,
        ':', 268,
        '<', 448,
        '=', 468,
        '>', 638,
        '?', 473,
        'L', 540,
        'R', 542,
        'U', 544,
        '[', 462,
        '\\', 42,
        '^', 435,
        'u', 547,
        '{', 460,
        '|', 432,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(186);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 187:
      ADVANCE_MAP(
        '!', 270,
        '"', 522,
        '%', 428,
        '&', 438,
        '(', 402,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 412,
        '.', 487,
        '/', 426,
        '<', 448,
        '=', 468,
        '>', 638,
        '?', 473,
        'L', 540,
        'R', 542,
        'U', 544,
        '[', 462,
        '\\', 68,
        '^', 435,
        'u', 547,
        '|', 432,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(187);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 188:
      ADVANCE_MAP(
        '!', 270,
        '"', 522,
        '%', 428,
        '&', 438,
        '(', 402,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 412,
        '.', 487,
        '/', 426,
        '<', 448,
        '=', 468,
        '>', 638,
        '?', 473,
        'L', 649,
        'R', 650,
        'U', 651,
        '[', 462,
        '\\', 68,
        '^', 435,
        'u', 652,
        '|', 432,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(187);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(658);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 189:
      ADVANCE_MAP(
        '!', 270,
        '"', 522,
        '%', 427,
        '&', 437,
        '(', 402,
        '*', 423,
        '+', 417,
        ',', 343,
        '-', 413,
        '.', 487,
        '/', 425,
        '<', 450,
        '=', 271,
        '>', 638,
        '?', 473,
        'L', 540,
        'R', 542,
        'U', 544,
        '[', 462,
        '\\', 135,
        '^', 434,
        'u', 547,
        '|', 433,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(189);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 190:
      ADVANCE_MAP(
        '!', 270,
        '"', 522,
        '%', 427,
        '&', 437,
        '(', 402,
        '*', 423,
        '+', 417,
        ',', 343,
        '-', 413,
        '.', 487,
        '/', 425,
        '<', 450,
        '=', 271,
        '>', 638,
        '?', 473,
        'L', 649,
        'R', 650,
        'U', 651,
        '[', 462,
        '\\', 135,
        '^', 434,
        'u', 652,
        '|', 433,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(189);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(658);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 191:
      ADVANCE_MAP(
        '!', 270,
        '#', 303,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 411,
        '.', 487,
        '/', 426,
        ':', 268,
        ';', 456,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        '[', 463,
        '\\', 34,
        '^', 435,
        '{', 460,
        '|', 432,
        '}', 461,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(191);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 192:
      ADVANCE_MAP(
        '!', 270,
        '#', 288,
        '%', 427,
        '&', 437,
        '(', 402,
        ')', 344,
        '*', 423,
        '+', 416,
        ',', 343,
        '-', 406,
        '/', 425,
        ':', 268,
        ';', 456,
        '<', 451,
        '=', 271,
        '>', 442,
        '[', 463,
        '\\', 22,
        '^', 434,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 603,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        'u', 596,
        'v', 608,
        '|', 433,
        '}', 461,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(192);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 193:
      ADVANCE_MAP(
        '!', 270,
        '#', 292,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 412,
        '.', 487,
        '/', 426,
        '0', 639,
        ':', 472,
        ';', 456,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        '[', 463,
        '\\', 32,
        ']', 466,
        '^', 435,
        '{', 460,
        '|', 432,
        '}', 461,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(193);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 194:
      ADVANCE_MAP(
        '!', 270,
        '#', 292,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 412,
        '.', 487,
        '/', 426,
        ':', 472,
        ';', 456,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        '[', 462,
        '\\', 46,
        ']', 466,
        '^', 435,
        '{', 460,
        '|', 432,
        '}', 461,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(194);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 195:
      ADVANCE_MAP(
        '!', 270,
        '#', 292,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 412,
        '.', 487,
        '/', 426,
        ':', 471,
        ';', 456,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        '[', 463,
        '\\', 157,
        ']', 466,
        '^', 435,
        '{', 460,
        '|', 432,
        '}', 461,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(195);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 196:
      ADVANCE_MAP(
        '!', 270,
        '#', 292,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 412,
        '.', 487,
        '/', 426,
        ':', 471,
        ';', 456,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        '[', 462,
        '\\', 40,
        ']', 466,
        '^', 435,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 603,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        'u', 596,
        'v', 608,
        '{', 460,
        '|', 432,
        '}', 461,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(196);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 197:
      ADVANCE_MAP(
        '!', 270,
        '#', 292,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 412,
        '.', 487,
        '/', 426,
        ':', 471,
        ';', 456,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        '[', 462,
        '\\', 125,
        ']', 466,
        '^', 435,
        '{', 460,
        '|', 432,
        '}', 461,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(197);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 198:
      ADVANCE_MAP(
        '!', 270,
        '#', 292,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 412,
        '.', 487,
        '/', 426,
        ':', 471,
        ';', 456,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        '[', 462,
        '\\', 139,
        ']', 466,
        '^', 435,
        '|', 432,
        '}', 461,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(199);
      if (('A' <= lookahead && lookahead <= '_') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(658);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 199:
      ADVANCE_MAP(
        '!', 270,
        '#', 292,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 412,
        '.', 487,
        '/', 426,
        ':', 471,
        ';', 456,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        '[', 462,
        '\\', 139,
        ']', 466,
        '^', 435,
        '|', 432,
        '}', 461,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(199);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 200:
      ADVANCE_MAP(
        '!', 270,
        '#', 292,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 412,
        '.', 487,
        '/', 426,
        ':', 471,
        ';', 456,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        '[', 465,
        '\\', 72,
        ']', 466,
        '^', 435,
        '|', 432,
        '}', 461,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(200);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 201:
      ADVANCE_MAP(
        '!', 270,
        '#', 292,
        '%', 427,
        '&', 437,
        '(', 402,
        ')', 344,
        '*', 423,
        '+', 417,
        ',', 343,
        '-', 413,
        '.', 487,
        '/', 425,
        ':', 472,
        ';', 456,
        '<', 450,
        '=', 271,
        '>', 442,
        '?', 473,
        '[', 462,
        '\\', 60,
        ']', 466,
        '^', 434,
        '{', 460,
        '|', 433,
        '}', 461,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(201);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 202:
      ADVANCE_MAP(
        '!', 270,
        '#', 292,
        '%', 427,
        '&', 437,
        '(', 402,
        ')', 344,
        '*', 423,
        '+', 417,
        ',', 343,
        '-', 413,
        '.', 487,
        '/', 425,
        ':', 471,
        ';', 456,
        '<', 450,
        '=', 271,
        '>', 442,
        '?', 473,
        '[', 462,
        '\\', 123,
        ']', 466,
        '^', 434,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 603,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        'u', 596,
        'v', 608,
        '{', 460,
        '|', 433,
        '}', 461,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(202);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 203:
      ADVANCE_MAP(
        '!', 270,
        '#', 292,
        '%', 427,
        '&', 437,
        '(', 402,
        ')', 344,
        '*', 423,
        '+', 417,
        ',', 343,
        '-', 413,
        '.', 487,
        '/', 425,
        ':', 471,
        ';', 456,
        '<', 450,
        '=', 271,
        '>', 442,
        '?', 473,
        '[', 462,
        '\\', 147,
        ']', 466,
        '^', 434,
        '{', 460,
        '|', 433,
        '}', 461,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(203);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 204:
      ADVANCE_MAP(
        '!', 270,
        '#', 292,
        '%', 427,
        '&', 437,
        '(', 402,
        ')', 344,
        '*', 423,
        '+', 417,
        ',', 343,
        '-', 413,
        '.', 487,
        '/', 425,
        ':', 471,
        ';', 456,
        '<', 450,
        '=', 271,
        '>', 442,
        '?', 473,
        '[', 462,
        '\\', 151,
        ']', 466,
        '^', 434,
        '|', 433,
        '}', 461,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(205);
      if (('A' <= lookahead && lookahead <= '_') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(658);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 205:
      ADVANCE_MAP(
        '!', 270,
        '#', 292,
        '%', 427,
        '&', 437,
        '(', 402,
        ')', 344,
        '*', 423,
        '+', 417,
        ',', 343,
        '-', 413,
        '.', 487,
        '/', 425,
        ':', 471,
        ';', 456,
        '<', 450,
        '=', 271,
        '>', 442,
        '?', 473,
        '[', 462,
        '\\', 151,
        ']', 466,
        '^', 434,
        '|', 433,
        '}', 461,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(205);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 206:
      ADVANCE_MAP(
        '!', 270,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 411,
        '.', 487,
        '/', 426,
        ':', 472,
        ';', 456,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        '[', 462,
        '\\', 52,
        '^', 435,
        '{', 460,
        '|', 432,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(206);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 207:
      ADVANCE_MAP(
        '!', 270,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 411,
        '.', 487,
        '/', 426,
        ':', 471,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        '[', 462,
        '\\', 129,
        '^', 435,
        '{', 460,
        '|', 432,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(207);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 208:
      ADVANCE_MAP(
        '!', 270,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 411,
        '.', 487,
        '/', 426,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        '[', 464,
        '\\', 84,
        '^', 435,
        '|', 432,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(208);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 209:
      ADVANCE_MAP(
        '!', 270,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 411,
        '.', 487,
        '/', 426,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        '[', 462,
        '\\', 48,
        '^', 435,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 603,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        'u', 596,
        'v', 608,
        '{', 460,
        '|', 432,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(209);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 210:
      ADVANCE_MAP(
        '!', 270,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 411,
        '.', 487,
        '/', 426,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        '[', 462,
        '\\', 143,
        '^', 435,
        '|', 432,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(211);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(658);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 211:
      ADVANCE_MAP(
        '!', 270,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 411,
        '.', 487,
        '/', 426,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        '[', 462,
        '\\', 143,
        '^', 435,
        '|', 432,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(211);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 212:
      ADVANCE_MAP(
        '!', 270,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 411,
        '.', 487,
        '/', 426,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        '[', 465,
        '\\', 155,
        '^', 435,
        '|', 432,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(212);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 213:
      ADVANCE_MAP(
        '!', 270,
        '%', 428,
        '&', 438,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 412,
        '.', 487,
        '/', 426,
        ':', 472,
        '<', 448,
        '=', 468,
        '>', 638,
        '?', 473,
        '[', 462,
        '\\', 54,
        '^', 435,
        '{', 460,
        '|', 432,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(213);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 214:
      ADVANCE_MAP(
        '!', 270,
        '%', 428,
        '&', 438,
        '(', 402,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 412,
        '.', 487,
        '/', 426,
        ':', 471,
        '<', 448,
        '=', 468,
        '>', 638,
        '?', 473,
        '[', 462,
        '\\', 137,
        '^', 435,
        '{', 460,
        '|', 432,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(214);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 215:
      ADVANCE_MAP(
        '!', 270,
        '%', 428,
        '&', 438,
        '(', 402,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 412,
        '.', 487,
        '/', 426,
        ';', 456,
        '<', 448,
        '=', 468,
        '>', 443,
        '?', 473,
        '[', 464,
        '\\', 86,
        '^', 435,
        '|', 432,
        '}', 461,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(215);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 216:
      ADVANCE_MAP(
        '!', 270,
        '%', 428,
        '&', 438,
        '(', 402,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 412,
        '.', 487,
        '/', 426,
        '<', 448,
        '=', 468,
        '>', 638,
        '?', 473,
        '[', 463,
        '\\', 82,
        '^', 435,
        '|', 432,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(216);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 217:
      ADVANCE_MAP(
        '!', 270,
        '%', 428,
        '&', 438,
        '(', 402,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 412,
        '.', 487,
        '/', 426,
        '<', 448,
        '=', 468,
        '>', 638,
        '?', 473,
        '[', 462,
        '\\', 50,
        '^', 435,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 603,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        'u', 596,
        'v', 608,
        '{', 460,
        '|', 432,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(217);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 218:
      ADVANCE_MAP(
        '!', 270,
        '%', 428,
        '&', 438,
        '(', 402,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 412,
        '.', 487,
        '/', 426,
        '<', 448,
        '=', 468,
        '>', 638,
        '?', 473,
        '[', 462,
        '\\', 141,
        '^', 435,
        '|', 432,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(219);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(658);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 219:
      ADVANCE_MAP(
        '!', 270,
        '%', 428,
        '&', 438,
        '(', 402,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 412,
        '.', 487,
        '/', 426,
        '<', 448,
        '=', 468,
        '>', 638,
        '?', 473,
        '[', 462,
        '\\', 141,
        '^', 435,
        '|', 432,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(219);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 220:
      ADVANCE_MAP(
        '!', 270,
        '%', 428,
        '&', 438,
        '(', 402,
        '*', 424,
        '+', 420,
        ',', 343,
        '-', 412,
        '.', 487,
        '/', 426,
        '<', 448,
        '=', 468,
        '>', 638,
        '?', 473,
        '[', 465,
        '\\', 153,
        '^', 435,
        '|', 432,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(220);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 221:
      ADVANCE_MAP(
        '!', 270,
        '%', 428,
        '&', 438,
        '*', 424,
        '+', 422,
        ',', 343,
        '-', 415,
        '.', 251,
        '/', 426,
        '<', 449,
        '=', 468,
        '>', 443,
        '\\', 90,
        '^', 435,
        '|', 432,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(221);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 222:
      ADVANCE_MAP(
        '!', 270,
        '%', 427,
        '&', 437,
        '(', 402,
        ')', 344,
        '*', 423,
        '+', 417,
        ',', 343,
        '-', 413,
        '.', 487,
        '/', 425,
        ':', 472,
        ';', 456,
        '<', 450,
        '=', 468,
        '>', 442,
        '?', 473,
        '[', 463,
        '\\', 36,
        ']', 466,
        '^', 434,
        '{', 460,
        '|', 433,
        '}', 461,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(222);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 223:
      ADVANCE_MAP(
        '!', 270,
        '%', 427,
        '&', 437,
        '(', 402,
        ')', 344,
        '*', 423,
        '+', 417,
        ',', 343,
        '-', 413,
        '.', 487,
        '/', 425,
        ':', 472,
        ';', 456,
        '<', 450,
        '=', 271,
        '>', 442,
        '?', 473,
        '[', 462,
        '\\', 58,
        ']', 466,
        '^', 434,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 603,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        'u', 596,
        'v', 608,
        '{', 460,
        '|', 433,
        '}', 461,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(223);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 224:
      ADVANCE_MAP(
        '!', 270,
        '%', 427,
        '&', 437,
        '(', 402,
        ')', 344,
        '*', 423,
        '+', 417,
        ',', 343,
        '-', 413,
        '.', 487,
        '/', 425,
        ':', 268,
        ';', 456,
        '<', 450,
        '=', 271,
        '>', 442,
        '?', 473,
        '[', 462,
        '\\', 133,
        ']', 282,
        '^', 434,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 603,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        'u', 596,
        'v', 608,
        '{', 460,
        '|', 433,
        '}', 461,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(224);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 225:
      ADVANCE_MAP(
        '!', 270,
        '%', 427,
        '&', 437,
        '(', 402,
        '*', 423,
        '+', 417,
        ',', 343,
        '-', 413,
        '.', 487,
        '/', 425,
        ':', 472,
        '<', 450,
        '=', 271,
        '>', 638,
        '?', 473,
        '[', 462,
        '\\', 149,
        '^', 434,
        '{', 460,
        '|', 433,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(225);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 226:
      ADVANCE_MAP(
        '!', 270,
        '%', 427,
        '&', 437,
        '(', 402,
        '*', 423,
        '+', 417,
        ',', 343,
        '-', 413,
        '.', 487,
        '/', 425,
        ':', 471,
        '<', 450,
        '=', 271,
        '>', 638,
        '?', 473,
        '[', 462,
        '\\', 70,
        '^', 434,
        '{', 460,
        '|', 433,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(226);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 227:
      ADVANCE_MAP(
        '!', 270,
        '%', 427,
        '&', 437,
        '(', 402,
        '*', 423,
        '+', 417,
        ',', 343,
        '-', 413,
        '.', 487,
        '/', 425,
        '<', 450,
        '=', 271,
        '>', 638,
        '?', 473,
        '[', 462,
        '\\', 127,
        '^', 434,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 603,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        'u', 596,
        'v', 608,
        '{', 460,
        '|', 433,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(227);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 228:
      ADVANCE_MAP(
        '!', 270,
        '%', 427,
        '&', 437,
        '(', 402,
        '*', 423,
        '+', 417,
        ',', 343,
        '-', 413,
        '.', 487,
        '/', 425,
        '<', 450,
        '=', 271,
        '>', 638,
        '?', 473,
        '[', 462,
        '\\', 159,
        '^', 434,
        '|', 433,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(229);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(658);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 229:
      ADVANCE_MAP(
        '!', 270,
        '%', 427,
        '&', 437,
        '(', 402,
        '*', 423,
        '+', 417,
        ',', 343,
        '-', 413,
        '.', 487,
        '/', 425,
        '<', 450,
        '=', 271,
        '>', 638,
        '?', 473,
        '[', 462,
        '\\', 159,
        '^', 434,
        '|', 433,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(229);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 230:
      ADVANCE_MAP(
        '"', 522,
        '&', 437,
        '(', 402,
        '*', 423,
        '/', 252,
        ':', 268,
        'L', 541,
        'U', 545,
        '[', 463,
        '\\', 74,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 603,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        'u', 548,
        'v', 608,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(230);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 231:
      ADVANCE_MAP(
        '"', 522,
        ')', 344,
        ',', 343,
        '/', 252,
        ':', 471,
        'L', 540,
        'R', 542,
        'U', 544,
        '\\', 105,
        'u', 547,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(231);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 232:
      if (lookahead == '"') ADVANCE(522);
      if (lookahead == '/') ADVANCE(252);
      if (lookahead == '<') ADVANCE(276);
      if (lookahead == 'L') ADVANCE(541);
      if (lookahead == 'U') ADVANCE(545);
      if (lookahead == '\\') ADVANCE(107);
      if (lookahead == 'u') ADVANCE(549);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(232);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 233:
      if (lookahead == '"') ADVANCE(522);
      if (lookahead == '/') ADVANCE(252);
      if (lookahead == '\\') ADVANCE(109);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(233);
      END_STATE();
    case 234:
      if (lookahead == '"') ADVANCE(648);
      END_STATE();
    case 235:
      ADVANCE_MAP(
        '#', 292,
        '&', 437,
        '(', 402,
        ')', 344,
        '*', 423,
        ',', 343,
        '-', 275,
        '.', 257,
        '/', 252,
        ':', 472,
        ';', 456,
        '<', 447,
        '=', 467,
        '>', 638,
        '[', 462,
        '\\', 62,
        '{', 460,
        '|', 306,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(235);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 236:
      ADVANCE_MAP(
        '#', 284,
        '&', 437,
        '(', 402,
        ')', 344,
        '*', 423,
        '+', 416,
        ',', 343,
        '-', 275,
        '.', 257,
        '/', 252,
        ':', 472,
        ';', 456,
        '<', 447,
        '=', 467,
        '>', 638,
        '[', 463,
        '\\', 20,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 603,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        'u', 596,
        'v', 608,
        '{', 460,
        '|', 306,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(236);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 237:
      ADVANCE_MAP(
        '#', 286,
        '&', 437,
        '(', 402,
        '*', 423,
        ',', 343,
        '/', 252,
        ':', 268,
        ';', 456,
        '[', 463,
        '\\', 24,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 603,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        'u', 596,
        'v', 608,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(237);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 238:
      ADVANCE_MAP(
        '&', 437,
        '(', 402,
        ')', 344,
        '*', 423,
        ',', 343,
        '-', 275,
        '.', 257,
        '/', 252,
        ':', 472,
        ';', 456,
        '<', 447,
        '=', 467,
        '>', 638,
        '[', 463,
        '\\', 64,
        ']', 466,
        '{', 460,
        '|', 306,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(238);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 239:
      ADVANCE_MAP(
        '&', 437,
        '(', 402,
        ')', 344,
        '*', 423,
        ',', 343,
        '-', 275,
        '.', 257,
        '/', 252,
        ':', 471,
        ';', 456,
        '<', 447,
        '=', 467,
        '>', 638,
        '[', 463,
        '\\', 78,
        ']', 466,
        '{', 460,
        '|', 306,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(239);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 240:
      ADVANCE_MAP(
        '&', 437,
        '(', 402,
        ')', 344,
        '*', 423,
        ',', 343,
        '-', 275,
        '/', 252,
        ':', 471,
        ';', 456,
        '<', 447,
        '=', 467,
        '>', 638,
        '[', 463,
        '\\', 76,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 603,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        'u', 596,
        'v', 608,
        '{', 460,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(240);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 241:
      ADVANCE_MAP(
        '&', 437,
        '(', 402,
        ')', 344,
        '*', 423,
        ',', 343,
        '.', 257,
        '/', 252,
        ':', 471,
        ';', 456,
        '=', 467,
        '>', 638,
        '[', 462,
        '\\', 92,
        '{', 460,
        '|', 306,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(241);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 242:
      ADVANCE_MAP(
        '&', 437,
        '(', 402,
        ')', 344,
        '*', 423,
        ',', 343,
        '/', 252,
        ';', 456,
        '=', 467,
        '>', 638,
        '[', 462,
        '\\', 88,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 603,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        'u', 596,
        'v', 608,
        '{', 460,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(242);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 243:
      ADVANCE_MAP(
        '&', 436,
        '*', 423,
        '.', 257,
        '/', 252,
        ':', 268,
        '=', 467,
        '\\', 101,
        ']', 466,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(243);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 244:
      if (lookahead == '\'') ADVANCE(513);
      if (lookahead == '/') ADVANCE(252);
      if (lookahead == '\\') ADVANCE(109);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(244);
      END_STATE();
    case 245:
      if (lookahead == '\'') ADVANCE(317);
      if (lookahead == 'P' ||
          lookahead == 'p') ADVANCE(307);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(245);
      END_STATE();
    case 246:
      if (lookahead == '\'') ADVANCE(313);
      if (lookahead == '.') ADVANCE(502);
      if (lookahead == 'E' ||
          lookahead == 'e') ADVANCE(307);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(246);
      END_STATE();
    case 247:
      ADVANCE_MAP(
        '(', 402,
        ')', 344,
        ',', 343,
        '/', 252,
        ':', 471,
        ';', 456,
        '<', 447,
        '=', 467,
        '>', 638,
        '[', 464,
        '\\', 103,
        ']', 466,
        '{', 460,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(247);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 248:
      ADVANCE_MAP(
        '(', 402,
        '/', 252,
        ':', 268,
        'F', 565,
        'T', 569,
        '[', 462,
        '\\', 94,
        'f', 575,
        't', 616,
        '{', 460,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(248);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 249:
      if (lookahead == '(') ADVANCE(402);
      if (lookahead == '/') ADVANCE(252);
      if (lookahead == '\\') SKIP(117);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(249);
      END_STATE();
    case 250:
      if (lookahead == ')') ADVANCE(646);
      END_STATE();
    case 251:
      if (lookahead == '*') ADVANCE(490);
      END_STATE();
    case 252:
      if (lookahead == '*') ADVANCE(255);
      if (lookahead == '/') ADVANCE(636);
      END_STATE();
    case 253:
      if (lookahead == '*') ADVANCE(645);
      END_STATE();
    case 254:
      if (lookahead == '*') ADVANCE(254);
      if (lookahead == '/') ADVANCE(629);
      if (lookahead != 0) ADVANCE(255);
      END_STATE();
    case 255:
      if (lookahead == '*') ADVANCE(254);
      if (lookahead != 0) ADVANCE(255);
      END_STATE();
    case 256:
      if (lookahead == '*') ADVANCE(254);
      if (lookahead != 0) ADVANCE(360);
      END_STATE();
    case 257:
      if (lookahead == '.') ADVANCE(259);
      END_STATE();
    case 258:
      if (lookahead == '.') ADVANCE(259);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(494);
      END_STATE();
    case 259:
      if (lookahead == '.') ADVANCE(342);
      END_STATE();
    case 260:
      if (lookahead == '.') ADVANCE(317);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(500);
      END_STATE();
    case 261:
      if (lookahead == '/') ADVANCE(367);
      if (lookahead == '\\') ADVANCE(362);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(261);
      if (lookahead != 0) ADVANCE(369);
      END_STATE();
    case 262:
      if (lookahead == '1') ADVANCE(266);
      END_STATE();
    case 263:
      if (lookahead == '2') ADVANCE(493);
      END_STATE();
    case 264:
      if (lookahead == '2') ADVANCE(267);
      if (lookahead == '6') ADVANCE(493);
      END_STATE();
    case 265:
      if (lookahead == '4') ADVANCE(493);
      END_STATE();
    case 266:
      if (lookahead == '6') ADVANCE(493);
      END_STATE();
    case 267:
      if (lookahead == '8') ADVANCE(493);
      END_STATE();
    case 268:
      if (lookahead == ':') ADVANCE(457);
      END_STATE();
    case 269:
      if (lookahead == '<') ADVANCE(452);
      if (lookahead == '=') ADVANCE(445);
      END_STATE();
    case 270:
      if (lookahead == '=') ADVANCE(440);
      END_STATE();
    case 271:
      if (lookahead == '=') ADVANCE(439);
      END_STATE();
    case 272:
      if (lookahead == '=') ADVANCE(444);
      if (lookahead == '>') ADVANCE(454);
      END_STATE();
    case 273:
      if (lookahead == '=') ADVANCE(480);
      END_STATE();
    case 274:
      if (lookahead == '>') ADVANCE(273);
      END_STATE();
    case 275:
      if (lookahead == '>') ADVANCE(491);
      END_STATE();
    case 276:
      if (lookahead == '>') ADVANCE(535);
      if (lookahead == '\\') ADVANCE(277);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(276);
      END_STATE();
    case 277:
      if (lookahead == '>') ADVANCE(536);
      if (lookahead == '\\') ADVANCE(277);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(276);
      END_STATE();
    case 278:
      if (lookahead == 'F') ADVANCE(262);
      END_STATE();
    case 279:
      if (lookahead == 'U') ADVANCE(332);
      if (lookahead == 'u') ADVANCE(324);
      END_STATE();
    case 280:
      if (lookahead == '[') ADVANCE(458);
      if (lookahead == ']') ADVANCE(647);
      END_STATE();
    case 281:
      if (lookahead == ']') ADVANCE(647);
      END_STATE();
    case 282:
      if (lookahead == ']') ADVANCE(459);
      END_STATE();
    case 283:
      if (lookahead == 'd') ADVANCE(376);
      if (lookahead == 'e') ADVANCE(396);
      if (lookahead == 'i') ADVANCE(384);
      if (lookahead == '\t' ||
          lookahead == ' ') ADVANCE(283);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 284:
      if (lookahead == 'd') ADVANCE(376);
      if (lookahead == 'e') ADVANCE(396);
      if (lookahead == 'i') ADVANCE(385);
      if (lookahead == '\t' ||
          lookahead == ' ') ADVANCE(284);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 285:
      if (lookahead == 'd') ADVANCE(376);
      if (lookahead == 'e') ADVANCE(398);
      if (lookahead == 'i') ADVANCE(384);
      if (lookahead == '\t' ||
          lookahead == ' ') ADVANCE(285);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 286:
      if (lookahead == 'd') ADVANCE(376);
      if (lookahead == 'e') ADVANCE(398);
      if (lookahead == 'i') ADVANCE(385);
      if (lookahead == '\t' ||
          lookahead == ' ') ADVANCE(286);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 287:
      if (lookahead == 'd') ADVANCE(376);
      if (lookahead == 'i') ADVANCE(384);
      if (lookahead == '\t' ||
          lookahead == ' ') ADVANCE(287);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 288:
      if (lookahead == 'd') ADVANCE(376);
      if (lookahead == 'i') ADVANCE(385);
      if (lookahead == '\t' ||
          lookahead == ' ') ADVANCE(288);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 289:
      if (lookahead == 'd') ADVANCE(301);
      END_STATE();
    case 290:
      if (lookahead == 'd') ADVANCE(295);
      END_STATE();
    case 291:
      if (lookahead == 'e') ADVANCE(305);
      if (lookahead == '\t' ||
          lookahead == ' ') ADVANCE(291);
      END_STATE();
    case 292:
      if (lookahead == 'e') ADVANCE(304);
      if (lookahead == '\t' ||
          lookahead == ' ') ADVANCE(292);
      END_STATE();
    case 293:
      if (lookahead == 'e') ADVANCE(351);
      END_STATE();
    case 294:
      if (lookahead == 'e') ADVANCE(299);
      END_STATE();
    case 295:
      if (lookahead == 'e') ADVANCE(300);
      END_STATE();
    case 296:
      if (lookahead == 'f') ADVANCE(262);
      END_STATE();
    case 297:
      if (lookahead == 'f') ADVANCE(347);
      END_STATE();
    case 298:
      if (lookahead == 'f') ADVANCE(353);
      END_STATE();
    case 299:
      if (lookahead == 'f') ADVANCE(355);
      END_STATE();
    case 300:
      if (lookahead == 'f') ADVANCE(357);
      END_STATE();
    case 301:
      if (lookahead == 'i') ADVANCE(297);
      END_STATE();
    case 302:
      if (lookahead == 'i') ADVANCE(298);
      if (lookahead == 's') ADVANCE(293);
      END_STATE();
    case 303:
      if (lookahead == 'i') ADVANCE(385);
      if (lookahead == '\t' ||
          lookahead == ' ') ADVANCE(303);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 304:
      if (lookahead == 'l') ADVANCE(302);
      if (lookahead == 'n') ADVANCE(289);
      END_STATE();
    case 305:
      if (lookahead == 'n') ADVANCE(289);
      END_STATE();
    case 306:
      if (lookahead == '|') ADVANCE(429);
      END_STATE();
    case 307:
      if (lookahead == '+' ||
          lookahead == '-') ADVANCE(314);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(499);
      END_STATE();
    case 308:
      if (lookahead == 'P' ||
          lookahead == 'p') ADVANCE(307);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(245);
      END_STATE();
    case 309:
      if (lookahead == '0' ||
          lookahead == '1') ADVANCE(497);
      END_STATE();
    case 310:
      if (lookahead == '8' ||
          lookahead == '9') ADVANCE(246);
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(496);
      END_STATE();
    case 311:
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(498);
      END_STATE();
    case 312:
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(494);
      END_STATE();
    case 313:
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(246);
      END_STATE();
    case 314:
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(499);
      END_STATE();
    case 315:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(628);
      END_STATE();
    case 316:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(500);
      END_STATE();
    case 317:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(245);
      END_STATE();
    case 318:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(528);
      END_STATE();
    case 319:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(534);
      END_STATE();
    case 320:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(315);
      END_STATE();
    case 321:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(318);
      END_STATE();
    case 322:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(320);
      END_STATE();
    case 323:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(321);
      END_STATE();
    case 324:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(322);
      END_STATE();
    case 325:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(323);
      END_STATE();
    case 326:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(324);
      END_STATE();
    case 327:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(325);
      END_STATE();
    case 328:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(326);
      END_STATE();
    case 329:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(327);
      END_STATE();
    case 330:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(328);
      END_STATE();
    case 331:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(329);
      END_STATE();
    case 332:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(330);
      END_STATE();
    case 333:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(331);
      END_STATE();
    case 334:
      if (lookahead != 0 &&
          lookahead != '*') ADVANCE(369);
      END_STATE();
    case 335:
      if (eof) ADVANCE(337);
      ADVANCE_MAP(
        '!', 404,
        '"', 522,
        '#', 283,
        '%', 428,
        '&', 438,
        '\'', 513,
        '(', 402,
        ')', 344,
        '*', 424,
        '+', 418,
        ',', 343,
        '-', 407,
        '.', 488,
        '/', 426,
        '0', 495,
        ':', 472,
        ';', 456,
        '<', 448,
        '=', 468,
        '>', 638,
        '?', 473,
        'F', 565,
        'L', 539,
        'R', 542,
        'T', 569,
        'U', 543,
        '[', 464,
        '\\', 2,
        ']', 466,
        '^', 435,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 574,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        't', 616,
        'u', 546,
        'v', 608,
        '{', 460,
        '|', 432,
        '}', 461,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(335);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(498);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 336:
      if (eof) ADVANCE(337);
      ADVANCE_MAP(
        '!', 403,
        '"', 522,
        '#', 287,
        '%', 427,
        '&', 437,
        '\'', 513,
        '(', 402,
        ')', 344,
        '*', 423,
        '+', 419,
        ',', 343,
        '-', 410,
        '.', 489,
        '/', 425,
        '0', 495,
        ':', 472,
        ';', 456,
        '<', 447,
        '=', 467,
        '>', 638,
        'F', 565,
        'L', 539,
        'R', 542,
        'T', 569,
        'U', 543,
        '[', 463,
        '\\', 4,
        ']', 466,
        '^', 434,
        'b', 611,
        'c', 590,
        'd', 607,
        'f', 574,
        'i', 604,
        'm', 576,
        'n', 624,
        'p', 621,
        's', 591,
        't', 616,
        'u', 546,
        'v', 608,
        '{', 460,
        '|', 306,
        '}', 461,
        '~', 405,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(336);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(498);
      if (set_contains(sym_identifier_character_set_1, 670, lookahead)) ADVANCE(628);
      END_STATE();
    case 337:
      ACCEPT_TOKEN(ts_builtin_sym_end);
      END_STATE();
    case 338:
      ACCEPT_TOKEN(aux_sym_preproc_include_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 339:
      ACCEPT_TOKEN(aux_sym_preproc_include_token2);
      END_STATE();
    case 340:
      ACCEPT_TOKEN(aux_sym_preproc_def_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 341:
      ACCEPT_TOKEN(anon_sym_LPAREN);
      END_STATE();
    case 342:
      ACCEPT_TOKEN(anon_sym_DOT_DOT_DOT);
      END_STATE();
    case 343:
      ACCEPT_TOKEN(anon_sym_COMMA);
      END_STATE();
    case 344:
      ACCEPT_TOKEN(anon_sym_RPAREN);
      END_STATE();
    case 345:
      ACCEPT_TOKEN(aux_sym_preproc_if_token1);
      if (lookahead == 'd') ADVANCE(380);
      if (lookahead == 'n') ADVANCE(374);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 346:
      ACCEPT_TOKEN(anon_sym_LF);
      if (lookahead == '\n') ADVANCE(346);
      END_STATE();
    case 347:
      ACCEPT_TOKEN(aux_sym_preproc_if_token2);
      END_STATE();
    case 348:
      ACCEPT_TOKEN(aux_sym_preproc_if_token2);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 349:
      ACCEPT_TOKEN(aux_sym_preproc_ifdef_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 350:
      ACCEPT_TOKEN(aux_sym_preproc_ifdef_token2);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 351:
      ACCEPT_TOKEN(aux_sym_preproc_else_token1);
      END_STATE();
    case 352:
      ACCEPT_TOKEN(aux_sym_preproc_else_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 353:
      ACCEPT_TOKEN(aux_sym_preproc_elif_token1);
      if (lookahead == 'd') ADVANCE(294);
      if (lookahead == 'n') ADVANCE(290);
      END_STATE();
    case 354:
      ACCEPT_TOKEN(aux_sym_preproc_elif_token1);
      if (lookahead == 'd') ADVANCE(382);
      if (lookahead == 'n') ADVANCE(375);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 355:
      ACCEPT_TOKEN(aux_sym_preproc_elifdef_token1);
      END_STATE();
    case 356:
      ACCEPT_TOKEN(aux_sym_preproc_elifdef_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 357:
      ACCEPT_TOKEN(aux_sym_preproc_elifdef_token2);
      END_STATE();
    case 358:
      ACCEPT_TOKEN(aux_sym_preproc_elifdef_token2);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 359:
      ACCEPT_TOKEN(sym_preproc_arg);
      if (lookahead == '\n') ADVANCE(255);
      if (lookahead == '*') ADVANCE(359);
      if (lookahead == '/') ADVANCE(629);
      if (lookahead == '\\') ADVANCE(365);
      if (lookahead != 0) ADVANCE(360);
      END_STATE();
    case 360:
      ACCEPT_TOKEN(sym_preproc_arg);
      if (lookahead == '\n') ADVANCE(255);
      if (lookahead == '*') ADVANCE(359);
      if (lookahead == '/') ADVANCE(256);
      if (lookahead == '\\') ADVANCE(365);
      if (lookahead != 0) ADVANCE(360);
      END_STATE();
    case 361:
      ACCEPT_TOKEN(sym_preproc_arg);
      if (lookahead == '\n') ADVANCE(636);
      if (lookahead == '\r') ADVANCE(630);
      if (lookahead == '/') ADVANCE(633);
      if (lookahead == '\\') ADVANCE(632);
      if (lookahead != 0) ADVANCE(634);
      END_STATE();
    case 362:
      ACCEPT_TOKEN(sym_preproc_arg);
      if (lookahead == '\n') SKIP(261);
      if (lookahead == '\r') ADVANCE(363);
      if (lookahead == '/') ADVANCE(334);
      if (lookahead == '\\') ADVANCE(364);
      if (lookahead != 0) ADVANCE(369);
      END_STATE();
    case 363:
      ACCEPT_TOKEN(sym_preproc_arg);
      if (lookahead == '\n') SKIP(261);
      if (lookahead == '/') ADVANCE(334);
      if (lookahead == '\\') ADVANCE(364);
      if (lookahead != 0) ADVANCE(369);
      END_STATE();
    case 364:
      ACCEPT_TOKEN(sym_preproc_arg);
      if (lookahead == '\r') ADVANCE(370);
      if (lookahead == '/') ADVANCE(334);
      if (lookahead == '\\') ADVANCE(364);
      if (lookahead != 0) ADVANCE(369);
      END_STATE();
    case 365:
      ACCEPT_TOKEN(sym_preproc_arg);
      if (lookahead == '\r') ADVANCE(368);
      if (lookahead == '*') ADVANCE(359);
      if (lookahead == '/') ADVANCE(256);
      if (lookahead == '\\') ADVANCE(365);
      if (lookahead != 0) ADVANCE(360);
      END_STATE();
    case 366:
      ACCEPT_TOKEN(sym_preproc_arg);
      if (lookahead == '\r') ADVANCE(635);
      if (lookahead == '/') ADVANCE(633);
      if (lookahead == '\\') ADVANCE(632);
      if (lookahead != 0) ADVANCE(634);
      END_STATE();
    case 367:
      ACCEPT_TOKEN(sym_preproc_arg);
      if (lookahead == '*') ADVANCE(360);
      if (lookahead == '/') ADVANCE(633);
      if (lookahead == '\\') ADVANCE(364);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(369);
      END_STATE();
    case 368:
      ACCEPT_TOKEN(sym_preproc_arg);
      if (lookahead == '*') ADVANCE(359);
      if (lookahead == '/') ADVANCE(256);
      if (lookahead == '\\') ADVANCE(365);
      if (lookahead != 0) ADVANCE(360);
      END_STATE();
    case 369:
      ACCEPT_TOKEN(sym_preproc_arg);
      if (lookahead == '/') ADVANCE(334);
      if (lookahead == '\\') ADVANCE(364);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(369);
      END_STATE();
    case 370:
      ACCEPT_TOKEN(sym_preproc_arg);
      if (lookahead == '/') ADVANCE(334);
      if (lookahead == '\\') ADVANCE(364);
      if (lookahead != 0) ADVANCE(369);
      END_STATE();
    case 371:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'c') ADVANCE(397);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 372:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'd') ADVANCE(395);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 373:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'd') ADVANCE(379);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 374:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'd') ADVANCE(381);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 375:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'd') ADVANCE(383);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 376:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'e') ADVANCE(386);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 377:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'e') ADVANCE(352);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 378:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'e') ADVANCE(340);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 379:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'e') ADVANCE(338);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 380:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'e') ADVANCE(389);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 381:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'e') ADVANCE(390);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 382:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'e') ADVANCE(391);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 383:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'e') ADVANCE(392);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 384:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'f') ADVANCE(345);
      if (lookahead == 'n') ADVANCE(371);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 385:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'f') ADVANCE(345);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 386:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'f') ADVANCE(393);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 387:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'f') ADVANCE(354);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 388:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'f') ADVANCE(348);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 389:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'f') ADVANCE(349);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 390:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'f') ADVANCE(350);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 391:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'f') ADVANCE(356);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 392:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'f') ADVANCE(358);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 393:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'i') ADVANCE(399);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 394:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'i') ADVANCE(387);
      if (lookahead == 's') ADVANCE(377);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 395:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'i') ADVANCE(388);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 396:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'l') ADVANCE(394);
      if (lookahead == 'n') ADVANCE(372);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 397:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'l') ADVANCE(400);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 398:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'n') ADVANCE(372);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 399:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'n') ADVANCE(378);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 400:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (lookahead == 'u') ADVANCE(373);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 401:
      ACCEPT_TOKEN(sym_preproc_directive);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(401);
      END_STATE();
    case 402:
      ACCEPT_TOKEN(anon_sym_LPAREN2);
      END_STATE();
    case 403:
      ACCEPT_TOKEN(anon_sym_BANG);
      END_STATE();
    case 404:
      ACCEPT_TOKEN(anon_sym_BANG);
      if (lookahead == '=') ADVANCE(440);
      END_STATE();
    case 405:
      ACCEPT_TOKEN(anon_sym_TILDE);
      END_STATE();
    case 406:
      ACCEPT_TOKEN(anon_sym_DASH);
      END_STATE();
    case 407:
      ACCEPT_TOKEN(anon_sym_DASH);
      if (lookahead == '-') ADVANCE(485);
      if (lookahead == '.') ADVANCE(312);
      if (lookahead == '0') ADVANCE(495);
      if (lookahead == '=') ADVANCE(478);
      if (lookahead == '>') ADVANCE(492);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(498);
      END_STATE();
    case 408:
      ACCEPT_TOKEN(anon_sym_DASH);
      if (lookahead == '-') ADVANCE(485);
      if (lookahead == '.') ADVANCE(312);
      if (lookahead == '0') ADVANCE(495);
      if (lookahead == '=') ADVANCE(478);
      if (lookahead == '>') ADVANCE(491);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(498);
      END_STATE();
    case 409:
      ACCEPT_TOKEN(anon_sym_DASH);
      if (lookahead == '-') ADVANCE(485);
      if (lookahead == '.') ADVANCE(312);
      if (lookahead == '0') ADVANCE(495);
      if (lookahead == '>') ADVANCE(491);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(498);
      END_STATE();
    case 410:
      ACCEPT_TOKEN(anon_sym_DASH);
      if (lookahead == '-') ADVANCE(485);
      if (lookahead == '.') ADVANCE(312);
      if (lookahead == '0') ADVANCE(495);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(498);
      END_STATE();
    case 411:
      ACCEPT_TOKEN(anon_sym_DASH);
      if (lookahead == '-') ADVANCE(485);
      if (lookahead == '=') ADVANCE(478);
      if (lookahead == '>') ADVANCE(492);
      END_STATE();
    case 412:
      ACCEPT_TOKEN(anon_sym_DASH);
      if (lookahead == '-') ADVANCE(485);
      if (lookahead == '=') ADVANCE(478);
      if (lookahead == '>') ADVANCE(491);
      END_STATE();
    case 413:
      ACCEPT_TOKEN(anon_sym_DASH);
      if (lookahead == '-') ADVANCE(485);
      if (lookahead == '>') ADVANCE(491);
      END_STATE();
    case 414:
      ACCEPT_TOKEN(anon_sym_DASH);
      if (lookahead == '.') ADVANCE(312);
      if (lookahead == '0') ADVANCE(495);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(498);
      END_STATE();
    case 415:
      ACCEPT_TOKEN(anon_sym_DASH);
      if (lookahead == '=') ADVANCE(478);
      if (lookahead == '>') ADVANCE(253);
      END_STATE();
    case 416:
      ACCEPT_TOKEN(anon_sym_PLUS);
      END_STATE();
    case 417:
      ACCEPT_TOKEN(anon_sym_PLUS);
      if (lookahead == '+') ADVANCE(486);
      END_STATE();
    case 418:
      ACCEPT_TOKEN(anon_sym_PLUS);
      if (lookahead == '+') ADVANCE(486);
      if (lookahead == '.') ADVANCE(312);
      if (lookahead == '0') ADVANCE(495);
      if (lookahead == '=') ADVANCE(477);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(498);
      END_STATE();
    case 419:
      ACCEPT_TOKEN(anon_sym_PLUS);
      if (lookahead == '+') ADVANCE(486);
      if (lookahead == '.') ADVANCE(312);
      if (lookahead == '0') ADVANCE(495);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(498);
      END_STATE();
    case 420:
      ACCEPT_TOKEN(anon_sym_PLUS);
      if (lookahead == '+') ADVANCE(486);
      if (lookahead == '=') ADVANCE(477);
      END_STATE();
    case 421:
      ACCEPT_TOKEN(anon_sym_PLUS);
      if (lookahead == '.') ADVANCE(312);
      if (lookahead == '0') ADVANCE(495);
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(498);
      END_STATE();
    case 422:
      ACCEPT_TOKEN(anon_sym_PLUS);
      if (lookahead == '=') ADVANCE(477);
      END_STATE();
    case 423:
      ACCEPT_TOKEN(anon_sym_STAR);
      END_STATE();
    case 424:
      ACCEPT_TOKEN(anon_sym_STAR);
      if (lookahead == '=') ADVANCE(474);
      END_STATE();
    case 425:
      ACCEPT_TOKEN(anon_sym_SLASH);
      if (lookahead == '*') ADVANCE(255);
      if (lookahead == '/') ADVANCE(636);
      END_STATE();
    case 426:
      ACCEPT_TOKEN(anon_sym_SLASH);
      if (lookahead == '*') ADVANCE(255);
      if (lookahead == '/') ADVANCE(636);
      if (lookahead == '=') ADVANCE(475);
      END_STATE();
    case 427:
      ACCEPT_TOKEN(anon_sym_PERCENT);
      END_STATE();
    case 428:
      ACCEPT_TOKEN(anon_sym_PERCENT);
      if (lookahead == '=') ADVANCE(476);
      END_STATE();
    case 429:
      ACCEPT_TOKEN(anon_sym_PIPE_PIPE);
      END_STATE();
    case 430:
      ACCEPT_TOKEN(anon_sym_AMP_AMP);
      END_STATE();
    case 431:
      ACCEPT_TOKEN(anon_sym_PIPE);
      END_STATE();
    case 432:
      ACCEPT_TOKEN(anon_sym_PIPE);
      if (lookahead == '=') ADVANCE(483);
      if (lookahead == '|') ADVANCE(429);
      END_STATE();
    case 433:
      ACCEPT_TOKEN(anon_sym_PIPE);
      if (lookahead == '|') ADVANCE(429);
      END_STATE();
    case 434:
      ACCEPT_TOKEN(anon_sym_CARET);
      END_STATE();
    case 435:
      ACCEPT_TOKEN(anon_sym_CARET);
      if (lookahead == '=') ADVANCE(482);
      END_STATE();
    case 436:
      ACCEPT_TOKEN(anon_sym_AMP);
      END_STATE();
    case 437:
      ACCEPT_TOKEN(anon_sym_AMP);
      if (lookahead == '&') ADVANCE(430);
      END_STATE();
    case 438:
      ACCEPT_TOKEN(anon_sym_AMP);
      if (lookahead == '&') ADVANCE(430);
      if (lookahead == '=') ADVANCE(481);
      END_STATE();
    case 439:
      ACCEPT_TOKEN(anon_sym_EQ_EQ);
      END_STATE();
    case 440:
      ACCEPT_TOKEN(anon_sym_BANG_EQ);
      END_STATE();
    case 441:
      ACCEPT_TOKEN(anon_sym_GT);
      END_STATE();
    case 442:
      ACCEPT_TOKEN(anon_sym_GT);
      if (lookahead == '=') ADVANCE(444);
      if (lookahead == '>') ADVANCE(454);
      END_STATE();
    case 443:
      ACCEPT_TOKEN(anon_sym_GT);
      if (lookahead == '=') ADVANCE(444);
      if (lookahead == '>') ADVANCE(455);
      END_STATE();
    case 444:
      ACCEPT_TOKEN(anon_sym_GT_EQ);
      END_STATE();
    case 445:
      ACCEPT_TOKEN(anon_sym_LT_EQ);
      END_STATE();
    case 446:
      ACCEPT_TOKEN(anon_sym_LT_EQ);
      if (lookahead == '>') ADVANCE(484);
      END_STATE();
    case 447:
      ACCEPT_TOKEN(anon_sym_LT);
      END_STATE();
    case 448:
      ACCEPT_TOKEN(anon_sym_LT);
      if (lookahead == '<') ADVANCE(453);
      if (lookahead == '=') ADVANCE(446);
      END_STATE();
    case 449:
      ACCEPT_TOKEN(anon_sym_LT);
      if (lookahead == '<') ADVANCE(453);
      if (lookahead == '=') ADVANCE(445);
      END_STATE();
    case 450:
      ACCEPT_TOKEN(anon_sym_LT);
      if (lookahead == '<') ADVANCE(452);
      if (lookahead == '=') ADVANCE(446);
      END_STATE();
    case 451:
      ACCEPT_TOKEN(anon_sym_LT);
      if (lookahead == '<') ADVANCE(452);
      if (lookahead == '=') ADVANCE(445);
      END_STATE();
    case 452:
      ACCEPT_TOKEN(anon_sym_LT_LT);
      END_STATE();
    case 453:
      ACCEPT_TOKEN(anon_sym_LT_LT);
      if (lookahead == '=') ADVANCE(479);
      END_STATE();
    case 454:
      ACCEPT_TOKEN(anon_sym_GT_GT);
      END_STATE();
    case 455:
      ACCEPT_TOKEN(anon_sym_GT_GT);
      if (lookahead == '=') ADVANCE(480);
      END_STATE();
    case 456:
      ACCEPT_TOKEN(anon_sym_SEMI);
      END_STATE();
    case 457:
      ACCEPT_TOKEN(anon_sym_COLON_COLON);
      END_STATE();
    case 458:
      ACCEPT_TOKEN(anon_sym_LBRACK_LBRACK);
      END_STATE();
    case 459:
      ACCEPT_TOKEN(anon_sym_RBRACK_RBRACK);
      END_STATE();
    case 460:
      ACCEPT_TOKEN(anon_sym_LBRACE);
      END_STATE();
    case 461:
      ACCEPT_TOKEN(anon_sym_RBRACE);
      END_STATE();
    case 462:
      ACCEPT_TOKEN(anon_sym_LBRACK);
      END_STATE();
    case 463:
      ACCEPT_TOKEN(anon_sym_LBRACK);
      if (lookahead == '[') ADVANCE(458);
      END_STATE();
    case 464:
      ACCEPT_TOKEN(anon_sym_LBRACK);
      if (lookahead == '[') ADVANCE(458);
      if (lookahead == ']') ADVANCE(647);
      END_STATE();
    case 465:
      ACCEPT_TOKEN(anon_sym_LBRACK);
      if (lookahead == ']') ADVANCE(647);
      END_STATE();
    case 466:
      ACCEPT_TOKEN(anon_sym_RBRACK);
      END_STATE();
    case 467:
      ACCEPT_TOKEN(anon_sym_EQ);
      END_STATE();
    case 468:
      ACCEPT_TOKEN(anon_sym_EQ);
      if (lookahead == '=') ADVANCE(439);
      END_STATE();
    case 469:
      ACCEPT_TOKEN(sym_primitive_type);
      if (lookahead == '1') ADVANCE(564);
      if (lookahead == '3') ADVANCE(562);
      if (lookahead == '6') ADVANCE(563);
      if (lookahead == '8') ADVANCE(573);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'p') ADVANCE(622);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 470:
      ACCEPT_TOKEN(sym_primitive_type);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 471:
      ACCEPT_TOKEN(anon_sym_COLON);
      END_STATE();
    case 472:
      ACCEPT_TOKEN(anon_sym_COLON);
      if (lookahead == ':') ADVANCE(457);
      END_STATE();
    case 473:
      ACCEPT_TOKEN(anon_sym_QMARK);
      END_STATE();
    case 474:
      ACCEPT_TOKEN(anon_sym_STAR_EQ);
      END_STATE();
    case 475:
      ACCEPT_TOKEN(anon_sym_SLASH_EQ);
      END_STATE();
    case 476:
      ACCEPT_TOKEN(anon_sym_PERCENT_EQ);
      END_STATE();
    case 477:
      ACCEPT_TOKEN(anon_sym_PLUS_EQ);
      END_STATE();
    case 478:
      ACCEPT_TOKEN(anon_sym_DASH_EQ);
      END_STATE();
    case 479:
      ACCEPT_TOKEN(anon_sym_LT_LT_EQ);
      END_STATE();
    case 480:
      ACCEPT_TOKEN(anon_sym_GT_GT_EQ);
      END_STATE();
    case 481:
      ACCEPT_TOKEN(anon_sym_AMP_EQ);
      END_STATE();
    case 482:
      ACCEPT_TOKEN(anon_sym_CARET_EQ);
      END_STATE();
    case 483:
      ACCEPT_TOKEN(anon_sym_PIPE_EQ);
      END_STATE();
    case 484:
      ACCEPT_TOKEN(anon_sym_LT_EQ_GT);
      END_STATE();
    case 485:
      ACCEPT_TOKEN(anon_sym_DASH_DASH);
      END_STATE();
    case 486:
      ACCEPT_TOKEN(anon_sym_PLUS_PLUS);
      END_STATE();
    case 487:
      ACCEPT_TOKEN(anon_sym_DOT);
      if (lookahead == '*') ADVANCE(490);
      if (lookahead == '.') ADVANCE(259);
      END_STATE();
    case 488:
      ACCEPT_TOKEN(anon_sym_DOT);
      if (lookahead == '*') ADVANCE(490);
      if (lookahead == '.') ADVANCE(259);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(494);
      END_STATE();
    case 489:
      ACCEPT_TOKEN(anon_sym_DOT);
      if (lookahead == '.') ADVANCE(259);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(494);
      END_STATE();
    case 490:
      ACCEPT_TOKEN(anon_sym_DOT_STAR);
      END_STATE();
    case 491:
      ACCEPT_TOKEN(anon_sym_DASH_GT);
      END_STATE();
    case 492:
      ACCEPT_TOKEN(anon_sym_DASH_GT);
      if (lookahead == '*') ADVANCE(645);
      END_STATE();
    case 493:
      ACCEPT_TOKEN(sym_number_literal);
      END_STATE();
    case 494:
      ACCEPT_TOKEN(sym_number_literal);
      ADVANCE_MAP(
        '\'', 312,
        'B', 278,
        'b', 296,
        'E', 307,
        'e', 307,
        'F', 501,
        'f', 501,
        'L', 493,
        'l', 493,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(494);
      END_STATE();
    case 495:
      ACCEPT_TOKEN(sym_number_literal);
      ADVANCE_MAP(
        '\'', 310,
        '.', 502,
        'L', 503,
        'l', 506,
        'B', 309,
        'b', 309,
        'E', 307,
        'e', 307,
        'U', 505,
        'u', 505,
        'X', 260,
        'x', 260,
        'Z', 508,
        'z', 508,
        '8', 246,
        '9', 246,
      );
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(496);
      END_STATE();
    case 496:
      ACCEPT_TOKEN(sym_number_literal);
      ADVANCE_MAP(
        '\'', 310,
        '.', 502,
        'L', 503,
        'l', 506,
        'E', 307,
        'e', 307,
        'U', 505,
        'u', 505,
        'Z', 508,
        'z', 508,
        '8', 246,
        '9', 246,
      );
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(496);
      END_STATE();
    case 497:
      ACCEPT_TOKEN(sym_number_literal);
      ADVANCE_MAP(
        '\'', 309,
        'L', 503,
        'l', 506,
        'U', 505,
        'u', 505,
        'Z', 508,
        'z', 508,
        '0', 497,
        '1', 497,
      );
      END_STATE();
    case 498:
      ACCEPT_TOKEN(sym_number_literal);
      ADVANCE_MAP(
        '\'', 311,
        '.', 502,
        'L', 503,
        'l', 506,
        'E', 307,
        'e', 307,
        'U', 505,
        'u', 505,
        'Z', 508,
        'z', 508,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(498);
      END_STATE();
    case 499:
      ACCEPT_TOKEN(sym_number_literal);
      if (lookahead == '\'') ADVANCE(314);
      if (lookahead == 'B') ADVANCE(278);
      if (lookahead == 'b') ADVANCE(296);
      if (lookahead == 'F' ||
          lookahead == 'f') ADVANCE(501);
      if (lookahead == 'L' ||
          lookahead == 'l') ADVANCE(493);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(499);
      END_STATE();
    case 500:
      ACCEPT_TOKEN(sym_number_literal);
      ADVANCE_MAP(
        '\'', 316,
        '.', 308,
        'L', 503,
        'l', 506,
        'P', 307,
        'p', 307,
        'U', 505,
        'u', 505,
        'Z', 508,
        'z', 508,
      );
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(500);
      END_STATE();
    case 501:
      ACCEPT_TOKEN(sym_number_literal);
      if (lookahead == '1') ADVANCE(264);
      if (lookahead == '3') ADVANCE(263);
      if (lookahead == '6') ADVANCE(265);
      END_STATE();
    case 502:
      ACCEPT_TOKEN(sym_number_literal);
      ADVANCE_MAP(
        'B', 278,
        'b', 296,
        'E', 307,
        'e', 307,
        'F', 501,
        'f', 501,
        'L', 493,
        'l', 493,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(494);
      END_STATE();
    case 503:
      ACCEPT_TOKEN(sym_number_literal);
      if (lookahead == 'L') ADVANCE(508);
      if (lookahead == 'U' ||
          lookahead == 'u') ADVANCE(493);
      END_STATE();
    case 504:
      ACCEPT_TOKEN(sym_number_literal);
      if (lookahead == 'L') ADVANCE(493);
      END_STATE();
    case 505:
      ACCEPT_TOKEN(sym_number_literal);
      if (lookahead == 'L') ADVANCE(504);
      if (lookahead == 'l') ADVANCE(507);
      if (lookahead == 'Z' ||
          lookahead == 'z') ADVANCE(493);
      END_STATE();
    case 506:
      ACCEPT_TOKEN(sym_number_literal);
      if (lookahead == 'l') ADVANCE(508);
      if (lookahead == 'U' ||
          lookahead == 'u') ADVANCE(493);
      END_STATE();
    case 507:
      ACCEPT_TOKEN(sym_number_literal);
      if (lookahead == 'l') ADVANCE(493);
      END_STATE();
    case 508:
      ACCEPT_TOKEN(sym_number_literal);
      if (lookahead == 'U' ||
          lookahead == 'u') ADVANCE(493);
      END_STATE();
    case 509:
      ACCEPT_TOKEN(anon_sym_L_SQUOTE);
      END_STATE();
    case 510:
      ACCEPT_TOKEN(anon_sym_u_SQUOTE);
      END_STATE();
    case 511:
      ACCEPT_TOKEN(anon_sym_U_SQUOTE);
      END_STATE();
    case 512:
      ACCEPT_TOKEN(anon_sym_u8_SQUOTE);
      END_STATE();
    case 513:
      ACCEPT_TOKEN(anon_sym_SQUOTE);
      END_STATE();
    case 514:
      ACCEPT_TOKEN(aux_sym_char_literal_token1);
      END_STATE();
    case 515:
      ACCEPT_TOKEN(aux_sym_char_literal_token1);
      if (lookahead == '\n') ADVANCE(529);
      if (lookahead == '\r') ADVANCE(528);
      if (lookahead == 'U') ADVANCE(333);
      if (lookahead == 'u') ADVANCE(325);
      if (lookahead == 'x') ADVANCE(319);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(531);
      if (lookahead != 0) ADVANCE(528);
      END_STATE();
    case 516:
      ACCEPT_TOKEN(aux_sym_char_literal_token1);
      if (lookahead == '*') ADVANCE(255);
      if (lookahead == '/') ADVANCE(636);
      END_STATE();
    case 517:
      ACCEPT_TOKEN(aux_sym_char_literal_token1);
      if (lookahead == '\\') ADVANCE(109);
      END_STATE();
    case 518:
      ACCEPT_TOKEN(anon_sym_L_DQUOTE);
      END_STATE();
    case 519:
      ACCEPT_TOKEN(anon_sym_u_DQUOTE);
      END_STATE();
    case 520:
      ACCEPT_TOKEN(anon_sym_U_DQUOTE);
      END_STATE();
    case 521:
      ACCEPT_TOKEN(anon_sym_u8_DQUOTE);
      END_STATE();
    case 522:
      ACCEPT_TOKEN(anon_sym_DQUOTE);
      END_STATE();
    case 523:
      ACCEPT_TOKEN(aux_sym_string_literal_token1);
      if (lookahead == '*') ADVANCE(525);
      if (lookahead == '/') ADVANCE(527);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '"' &&
          lookahead != '\\') ADVANCE(527);
      END_STATE();
    case 524:
      ACCEPT_TOKEN(aux_sym_string_literal_token1);
      if (lookahead == '*') ADVANCE(524);
      if (lookahead == '/') ADVANCE(527);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '"' &&
          lookahead != '\\') ADVANCE(525);
      END_STATE();
    case 525:
      ACCEPT_TOKEN(aux_sym_string_literal_token1);
      if (lookahead == '*') ADVANCE(524);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '"' &&
          lookahead != '\\') ADVANCE(525);
      END_STATE();
    case 526:
      ACCEPT_TOKEN(aux_sym_string_literal_token1);
      if (lookahead == '/') ADVANCE(523);
      if (lookahead == '\t' ||
          (0x0b <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') ADVANCE(526);
      if (lookahead != 0 &&
          (lookahead < '\t' || '\r' < lookahead) &&
          lookahead != '"' &&
          lookahead != '\\') ADVANCE(527);
      END_STATE();
    case 527:
      ACCEPT_TOKEN(aux_sym_string_literal_token1);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '"' &&
          lookahead != '\\') ADVANCE(527);
      END_STATE();
    case 528:
      ACCEPT_TOKEN(sym_escape_sequence);
      END_STATE();
    case 529:
      ACCEPT_TOKEN(sym_escape_sequence);
      if (lookahead == '\\') ADVANCE(109);
      END_STATE();
    case 530:
      ACCEPT_TOKEN(sym_escape_sequence);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(528);
      END_STATE();
    case 531:
      ACCEPT_TOKEN(sym_escape_sequence);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(530);
      END_STATE();
    case 532:
      ACCEPT_TOKEN(sym_escape_sequence);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(528);
      END_STATE();
    case 533:
      ACCEPT_TOKEN(sym_escape_sequence);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(532);
      END_STATE();
    case 534:
      ACCEPT_TOKEN(sym_escape_sequence);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(533);
      END_STATE();
    case 535:
      ACCEPT_TOKEN(sym_system_lib_string);
      END_STATE();
    case 536:
      ACCEPT_TOKEN(sym_system_lib_string);
      if (lookahead == '>') ADVANCE(535);
      if (lookahead == '\\') ADVANCE(277);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(276);
      END_STATE();
    case 537:
      ACCEPT_TOKEN(sym_true);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 538:
      ACCEPT_TOKEN(sym_false);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 539:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '"') ADVANCE(518);
      if (lookahead == '\'') ADVANCE(509);
      if (lookahead == 'R') ADVANCE(550);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 540:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '"') ADVANCE(518);
      if (lookahead == 'R') ADVANCE(550);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 541:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '"') ADVANCE(518);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 542:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '"') ADVANCE(640);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 543:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '"') ADVANCE(520);
      if (lookahead == '\'') ADVANCE(511);
      if (lookahead == 'R') ADVANCE(551);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 544:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '"') ADVANCE(520);
      if (lookahead == 'R') ADVANCE(551);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 545:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '"') ADVANCE(520);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 546:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '"') ADVANCE(519);
      if (lookahead == '\'') ADVANCE(510);
      if (lookahead == '8') ADVANCE(552);
      if (lookahead == 'R') ADVANCE(555);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'i') ADVANCE(606);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 547:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '"') ADVANCE(519);
      if (lookahead == '8') ADVANCE(553);
      if (lookahead == 'R') ADVANCE(555);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 548:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '"') ADVANCE(519);
      if (lookahead == '8') ADVANCE(554);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'i') ADVANCE(606);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 549:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '"') ADVANCE(519);
      if (lookahead == '8') ADVANCE(554);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 550:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '"') ADVANCE(641);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 551:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '"') ADVANCE(643);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 552:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '"') ADVANCE(521);
      if (lookahead == '\'') ADVANCE(512);
      if (lookahead == 'R') ADVANCE(556);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 553:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '"') ADVANCE(521);
      if (lookahead == 'R') ADVANCE(556);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 554:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '"') ADVANCE(521);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 555:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '"') ADVANCE(642);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 556:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '"') ADVANCE(644);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 557:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\'') ADVANCE(509);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 558:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\'') ADVANCE(511);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 559:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\'') ADVANCE(510);
      if (lookahead == '8') ADVANCE(560);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 560:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\'') ADVANCE(512);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 561:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '1') ADVANCE(564);
      if (lookahead == '3') ADVANCE(562);
      if (lookahead == '6') ADVANCE(563);
      if (lookahead == '8') ADVANCE(573);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'p') ADVANCE(622);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 562:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '2') ADVANCE(573);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 563:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '4') ADVANCE(573);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 564:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '6') ADVANCE(573);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 565:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == 'A') ADVANCE(568);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 566:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == 'E') ADVANCE(537);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 567:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == 'E') ADVANCE(538);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 568:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == 'L') ADVANCE(570);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 569:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == 'R') ADVANCE(571);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 570:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == 'S') ADVANCE(567);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 571:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == 'U') ADVANCE(566);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 572:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == '_') ADVANCE(579);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 573:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == '_') ADVANCE(619);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 574:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'a') ADVANCE(597);
      if (lookahead == 'l') ADVANCE(609);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 575:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'a') ADVANCE(597);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 576:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'a') ADVANCE(626);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 577:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'a') ADVANCE(613);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 578:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'a') ADVANCE(619);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 579:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'a') ADVANCE(601);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 580:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'b') ADVANCE(602);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 581:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'd') ADVANCE(470);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 582:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'd') ADVANCE(593);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 583:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'e') ADVANCE(537);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 584:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'e') ADVANCE(470);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 585:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'e') ADVANCE(538);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 586:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'e') ADVANCE(573);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 587:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'f') ADVANCE(573);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 588:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'f') ADVANCE(587);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 589:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'g') ADVANCE(605);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 590:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'h') ADVANCE(577);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 591:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'i') ADVANCE(627);
      if (lookahead == 's') ADVANCE(592);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 592:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'i') ADVANCE(627);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 593:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'i') ADVANCE(588);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 594:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'i') ADVANCE(589);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 595:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'i') ADVANCE(581);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 596:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'i') ADVANCE(606);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 597:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'l') ADVANCE(617);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 598:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'l') ADVANCE(470);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 599:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'l') ADVANCE(612);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 600:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'l') ADVANCE(599);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 601:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'l') ADVANCE(594);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 602:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'l') ADVANCE(584);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 603:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'l') ADVANCE(609);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 604:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'n') ADVANCE(618);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 605:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'n') ADVANCE(573);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 606:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'n') ADVANCE(620);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 607:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'o') ADVANCE(623);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 608:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'o') ADVANCE(595);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 609:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'o') ADVANCE(578);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 610:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'o') ADVANCE(598);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 611:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'o') ADVANCE(610);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 612:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'p') ADVANCE(622);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 613:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'r') ADVANCE(469);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 614:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'r') ADVANCE(582);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 615:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'r') ADVANCE(573);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 616:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'r') ADVANCE(625);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 617:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 's') ADVANCE(585);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 618:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 't') ADVANCE(469);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 619:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 't') ADVANCE(470);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 620:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 't') ADVANCE(561);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 621:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 't') ADVANCE(614);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 622:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 't') ADVANCE(615);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 623:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'u') ADVANCE(580);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 624:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'u') ADVANCE(600);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 625:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'u') ADVANCE(583);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 626:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'x') ADVANCE(572);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 627:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (lookahead == 'z') ADVANCE(586);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 628:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == '\\') ADVANCE(279);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 629:
      ACCEPT_TOKEN(sym_comment);
      END_STATE();
    case 630:
      ACCEPT_TOKEN(sym_comment);
      if (lookahead == '\n') ADVANCE(636);
      if (lookahead == '/') ADVANCE(633);
      if (lookahead == '\\') ADVANCE(366);
      if (lookahead != 0) ADVANCE(634);
      END_STATE();
    case 631:
      ACCEPT_TOKEN(sym_comment);
      if (lookahead == '\r') ADVANCE(637);
      if (lookahead == '\\') ADVANCE(631);
      if (lookahead != 0) ADVANCE(636);
      END_STATE();
    case 632:
      ACCEPT_TOKEN(sym_comment);
      if (lookahead == '\r') ADVANCE(635);
      if (lookahead == '/') ADVANCE(633);
      if (lookahead == '\\') ADVANCE(632);
      if (lookahead != 0) ADVANCE(634);
      END_STATE();
    case 633:
      ACCEPT_TOKEN(sym_comment);
      if (lookahead == '*') ADVANCE(636);
      if (lookahead == '\\') ADVANCE(361);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(634);
      END_STATE();
    case 634:
      ACCEPT_TOKEN(sym_comment);
      if (lookahead == '/') ADVANCE(633);
      if (lookahead == '\\') ADVANCE(366);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(634);
      END_STATE();
    case 635:
      ACCEPT_TOKEN(sym_comment);
      if (lookahead == '/') ADVANCE(633);
      if (lookahead == '\\') ADVANCE(366);
      if (lookahead != 0) ADVANCE(634);
      END_STATE();
    case 636:
      ACCEPT_TOKEN(sym_comment);
      if (lookahead == '\\') ADVANCE(160);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(636);
      END_STATE();
    case 637:
      ACCEPT_TOKEN(sym_comment);
      if (lookahead == '\\') ADVANCE(160);
      if (lookahead != 0) ADVANCE(636);
      END_STATE();
    case 638:
      ACCEPT_TOKEN(anon_sym_GT2);
      END_STATE();
    case 639:
      ACCEPT_TOKEN(aux_sym_pure_virtual_clause_token1);
      END_STATE();
    case 640:
      ACCEPT_TOKEN(anon_sym_R_DQUOTE);
      END_STATE();
    case 641:
      ACCEPT_TOKEN(anon_sym_LR_DQUOTE);
      END_STATE();
    case 642:
      ACCEPT_TOKEN(anon_sym_uR_DQUOTE);
      END_STATE();
    case 643:
      ACCEPT_TOKEN(anon_sym_UR_DQUOTE);
      END_STATE();
    case 644:
      ACCEPT_TOKEN(anon_sym_u8R_DQUOTE);
      END_STATE();
    case 645:
      ACCEPT_TOKEN(anon_sym_DASH_GT_STAR);
      END_STATE();
    case 646:
      ACCEPT_TOKEN(anon_sym_LPAREN_RPAREN);
      END_STATE();
    case 647:
      ACCEPT_TOKEN(anon_sym_LBRACK_RBRACK);
      END_STATE();
    case 648:
      ACCEPT_TOKEN(anon_sym_DQUOTE_DQUOTE);
      END_STATE();
    case 649:
      ACCEPT_TOKEN(sym_literal_suffix);
      if (lookahead == '"') ADVANCE(518);
      if (lookahead == 'R') ADVANCE(653);
      if (lookahead == '\\') ADVANCE(279);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(658);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 650:
      ACCEPT_TOKEN(sym_literal_suffix);
      if (lookahead == '"') ADVANCE(640);
      if (lookahead == '\\') ADVANCE(279);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(658);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 651:
      ACCEPT_TOKEN(sym_literal_suffix);
      if (lookahead == '"') ADVANCE(520);
      if (lookahead == 'R') ADVANCE(654);
      if (lookahead == '\\') ADVANCE(279);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(658);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 652:
      ACCEPT_TOKEN(sym_literal_suffix);
      if (lookahead == '"') ADVANCE(519);
      if (lookahead == '8') ADVANCE(655);
      if (lookahead == 'R') ADVANCE(656);
      if (lookahead == '\\') ADVANCE(279);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(658);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 653:
      ACCEPT_TOKEN(sym_literal_suffix);
      if (lookahead == '"') ADVANCE(641);
      if (lookahead == '\\') ADVANCE(279);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(658);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 654:
      ACCEPT_TOKEN(sym_literal_suffix);
      if (lookahead == '"') ADVANCE(643);
      if (lookahead == '\\') ADVANCE(279);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(658);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 655:
      ACCEPT_TOKEN(sym_literal_suffix);
      if (lookahead == '"') ADVANCE(521);
      if (lookahead == 'R') ADVANCE(657);
      if (lookahead == '\\') ADVANCE(279);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(658);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 656:
      ACCEPT_TOKEN(sym_literal_suffix);
      if (lookahead == '"') ADVANCE(642);
      if (lookahead == '\\') ADVANCE(279);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(658);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 657:
      ACCEPT_TOKEN(sym_literal_suffix);
      if (lookahead == '"') ADVANCE(644);
      if (lookahead == '\\') ADVANCE(279);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(658);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
      END_STATE();
    case 658:
      ACCEPT_TOKEN(sym_literal_suffix);
      if (lookahead == '\\') ADVANCE(279);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(658);
      if (set_contains(sym_identifier_character_set_2, 778, lookahead)) ADVANCE(628);
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
      if (lookahead == 'N') ADVANCE(1);
      if (lookahead == '\\') SKIP(2);
      if (lookahead == '_') ADVANCE(3);
      if (lookahead == 'a') ADVANCE(4);
      if (lookahead == 'b') ADVANCE(5);
      if (lookahead == 'c') ADVANCE(6);
      if (lookahead == 'd') ADVANCE(7);
      if (lookahead == 'e') ADVANCE(8);
      if (lookahead == 'f') ADVANCE(9);
      if (lookahead == 'g') ADVANCE(10);
      if (lookahead == 'i') ADVANCE(11);
      if (lookahead == 'l') ADVANCE(12);
      if (lookahead == 'm') ADVANCE(13);
      if (lookahead == 'n') ADVANCE(14);
      if (lookahead == 'o') ADVANCE(15);
      if (lookahead == 'p') ADVANCE(16);
      if (lookahead == 'r') ADVANCE(17);
      if (lookahead == 's') ADVANCE(18);
      if (lookahead == 't') ADVANCE(19);
      if (lookahead == 'u') ADVANCE(20);
      if (lookahead == 'v') ADVANCE(21);
      if (lookahead == 'w') ADVANCE(22);
      if (lookahead == 'x') ADVANCE(23);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(0);
      END_STATE();
    case 1:
      if (lookahead == 'U') ADVANCE(24);
      END_STATE();
    case 2:
      if (lookahead == '\n') SKIP(0);
      if (lookahead == '\r') SKIP(25);
      END_STATE();
    case 3:
      if (lookahead == 'A') ADVANCE(26);
      if (lookahead == 'G') ADVANCE(27);
      if (lookahead == 'N') ADVANCE(28);
      if (lookahead == '_') ADVANCE(29);
      if (lookahead == 'a') ADVANCE(30);
      if (lookahead == 'u') ADVANCE(31);
      END_STATE();
    case 4:
      if (lookahead == 'l') ADVANCE(32);
      if (lookahead == 'n') ADVANCE(33);
      if (lookahead == 's') ADVANCE(34);
      if (lookahead == 'u') ADVANCE(35);
      END_STATE();
    case 5:
      if (lookahead == 'i') ADVANCE(36);
      if (lookahead == 'r') ADVANCE(37);
      END_STATE();
    case 6:
      if (lookahead == 'a') ADVANCE(38);
      if (lookahead == 'l') ADVANCE(39);
      if (lookahead == 'o') ADVANCE(40);
      END_STATE();
    case 7:
      if (lookahead == 'e') ADVANCE(41);
      if (lookahead == 'o') ADVANCE(42);
      END_STATE();
    case 8:
      if (lookahead == 'l') ADVANCE(43);
      if (lookahead == 'n') ADVANCE(44);
      if (lookahead == 'x') ADVANCE(45);
      END_STATE();
    case 9:
      if (lookahead == 'i') ADVANCE(46);
      if (lookahead == 'o') ADVANCE(47);
      if (lookahead == 'r') ADVANCE(48);
      END_STATE();
    case 10:
      if (lookahead == 'o') ADVANCE(49);
      END_STATE();
    case 11:
      if (lookahead == 'f') ADVANCE(50);
      if (lookahead == 'n') ADVANCE(51);
      END_STATE();
    case 12:
      if (lookahead == 'o') ADVANCE(52);
      END_STATE();
    case 13:
      if (lookahead == 'u') ADVANCE(53);
      END_STATE();
    case 14:
      if (lookahead == 'a') ADVANCE(54);
      if (lookahead == 'e') ADVANCE(55);
      if (lookahead == 'o') ADVANCE(56);
      if (lookahead == 'u') ADVANCE(57);
      END_STATE();
    case 15:
      if (lookahead == 'f') ADVANCE(58);
      if (lookahead == 'p') ADVANCE(59);
      if (lookahead == 'r') ADVANCE(60);
      if (lookahead == 'v') ADVANCE(61);
      END_STATE();
    case 16:
      if (lookahead == 'r') ADVANCE(62);
      if (lookahead == 'u') ADVANCE(63);
      END_STATE();
    case 17:
      if (lookahead == 'e') ADVANCE(64);
      END_STATE();
    case 18:
      if (lookahead == 'h') ADVANCE(65);
      if (lookahead == 'i') ADVANCE(66);
      if (lookahead == 't') ADVANCE(67);
      if (lookahead == 'w') ADVANCE(68);
      END_STATE();
    case 19:
      if (lookahead == 'e') ADVANCE(69);
      if (lookahead == 'h') ADVANCE(70);
      if (lookahead == 'r') ADVANCE(71);
      if (lookahead == 'y') ADVANCE(72);
      END_STATE();
    case 20:
      if (lookahead == 'n') ADVANCE(73);
      if (lookahead == 's') ADVANCE(74);
      END_STATE();
    case 21:
      if (lookahead == 'i') ADVANCE(75);
      if (lookahead == 'o') ADVANCE(76);
      END_STATE();
    case 22:
      if (lookahead == 'h') ADVANCE(77);
      END_STATE();
    case 23:
      if (lookahead == 'o') ADVANCE(78);
      END_STATE();
    case 24:
      if (lookahead == 'L') ADVANCE(79);
      END_STATE();
    case 25:
      if (lookahead == '\n') SKIP(0);
      END_STATE();
    case 26:
      if (lookahead == 'l') ADVANCE(80);
      if (lookahead == 't') ADVANCE(81);
      END_STATE();
    case 27:
      if (lookahead == 'e') ADVANCE(82);
      END_STATE();
    case 28:
      if (lookahead == 'o') ADVANCE(83);
      END_STATE();
    case 29:
      ADVANCE_MAP(
        'a', 84,
        'b', 85,
        'c', 86,
        'd', 87,
        'e', 88,
        'f', 89,
        'i', 90,
        'l', 91,
        'r', 92,
        's', 93,
        't', 94,
        'u', 95,
        'v', 96,
      );
      END_STATE();
    case 30:
      if (lookahead == 'l') ADVANCE(97);
      END_STATE();
    case 31:
      if (lookahead == 'n') ADVANCE(98);
      END_STATE();
    case 32:
      if (lookahead == 'i') ADVANCE(99);
      END_STATE();
    case 33:
      if (lookahead == 'd') ADVANCE(100);
      END_STATE();
    case 34:
      if (lookahead == 'm') ADVANCE(101);
      END_STATE();
    case 35:
      if (lookahead == 't') ADVANCE(102);
      END_STATE();
    case 36:
      if (lookahead == 't') ADVANCE(103);
      END_STATE();
    case 37:
      if (lookahead == 'e') ADVANCE(104);
      END_STATE();
    case 38:
      if (lookahead == 's') ADVANCE(105);
      if (lookahead == 't') ADVANCE(106);
      END_STATE();
    case 39:
      if (lookahead == 'a') ADVANCE(107);
      END_STATE();
    case 40:
      if (lookahead == '_') ADVANCE(108);
      if (lookahead == 'm') ADVANCE(109);
      if (lookahead == 'n') ADVANCE(110);
      END_STATE();
    case 41:
      if (lookahead == 'c') ADVANCE(111);
      if (lookahead == 'f') ADVANCE(112);
      if (lookahead == 'l') ADVANCE(113);
      END_STATE();
    case 42:
      ACCEPT_TOKEN(anon_sym_do);
      END_STATE();
    case 43:
      if (lookahead == 's') ADVANCE(114);
      END_STATE();
    case 44:
      if (lookahead == 'u') ADVANCE(115);
      END_STATE();
    case 45:
      if (lookahead == 'p') ADVANCE(116);
      if (lookahead == 't') ADVANCE(117);
      END_STATE();
    case 46:
      if (lookahead == 'n') ADVANCE(118);
      END_STATE();
    case 47:
      if (lookahead == 'r') ADVANCE(119);
      END_STATE();
    case 48:
      if (lookahead == 'i') ADVANCE(120);
      END_STATE();
    case 49:
      if (lookahead == 't') ADVANCE(121);
      END_STATE();
    case 50:
      ACCEPT_TOKEN(anon_sym_if);
      END_STATE();
    case 51:
      if (lookahead == 'l') ADVANCE(122);
      END_STATE();
    case 52:
      if (lookahead == 'n') ADVANCE(123);
      END_STATE();
    case 53:
      if (lookahead == 't') ADVANCE(124);
      END_STATE();
    case 54:
      if (lookahead == 'm') ADVANCE(125);
      END_STATE();
    case 55:
      if (lookahead == 'w') ADVANCE(126);
      END_STATE();
    case 56:
      if (lookahead == 'e') ADVANCE(127);
      if (lookahead == 'r') ADVANCE(128);
      if (lookahead == 't') ADVANCE(129);
      END_STATE();
    case 57:
      if (lookahead == 'l') ADVANCE(130);
      END_STATE();
    case 58:
      if (lookahead == 'f') ADVANCE(131);
      END_STATE();
    case 59:
      if (lookahead == 'e') ADVANCE(132);
      END_STATE();
    case 60:
      ACCEPT_TOKEN(anon_sym_or);
      if (lookahead == '_') ADVANCE(133);
      END_STATE();
    case 61:
      if (lookahead == 'e') ADVANCE(134);
      END_STATE();
    case 62:
      if (lookahead == 'i') ADVANCE(135);
      if (lookahead == 'o') ADVANCE(136);
      END_STATE();
    case 63:
      if (lookahead == 'b') ADVANCE(137);
      END_STATE();
    case 64:
      if (lookahead == 'g') ADVANCE(138);
      if (lookahead == 'q') ADVANCE(139);
      if (lookahead == 's') ADVANCE(140);
      if (lookahead == 't') ADVANCE(141);
      END_STATE();
    case 65:
      if (lookahead == 'o') ADVANCE(142);
      END_STATE();
    case 66:
      if (lookahead == 'g') ADVANCE(143);
      if (lookahead == 'z') ADVANCE(144);
      END_STATE();
    case 67:
      if (lookahead == 'a') ADVANCE(145);
      if (lookahead == 'r') ADVANCE(146);
      END_STATE();
    case 68:
      if (lookahead == 'i') ADVANCE(147);
      END_STATE();
    case 69:
      if (lookahead == 'm') ADVANCE(148);
      END_STATE();
    case 70:
      if (lookahead == 'i') ADVANCE(149);
      if (lookahead == 'r') ADVANCE(150);
      END_STATE();
    case 71:
      if (lookahead == 'y') ADVANCE(151);
      END_STATE();
    case 72:
      if (lookahead == 'p') ADVANCE(152);
      END_STATE();
    case 73:
      if (lookahead == 'i') ADVANCE(153);
      if (lookahead == 's') ADVANCE(154);
      END_STATE();
    case 74:
      if (lookahead == 'i') ADVANCE(155);
      END_STATE();
    case 75:
      if (lookahead == 'r') ADVANCE(156);
      END_STATE();
    case 76:
      if (lookahead == 'l') ADVANCE(157);
      END_STATE();
    case 77:
      if (lookahead == 'i') ADVANCE(158);
      END_STATE();
    case 78:
      if (lookahead == 'r') ADVANCE(159);
      END_STATE();
    case 79:
      if (lookahead == 'L') ADVANCE(160);
      END_STATE();
    case 80:
      if (lookahead == 'i') ADVANCE(161);
      END_STATE();
    case 81:
      if (lookahead == 'o') ADVANCE(162);
      END_STATE();
    case 82:
      if (lookahead == 'n') ADVANCE(163);
      END_STATE();
    case 83:
      if (lookahead == 'n') ADVANCE(164);
      if (lookahead == 'r') ADVANCE(165);
      END_STATE();
    case 84:
      if (lookahead == 'l') ADVANCE(166);
      if (lookahead == 's') ADVANCE(167);
      if (lookahead == 't') ADVANCE(168);
      END_STATE();
    case 85:
      if (lookahead == 'a') ADVANCE(169);
      END_STATE();
    case 86:
      if (lookahead == 'd') ADVANCE(170);
      if (lookahead == 'l') ADVANCE(171);
      END_STATE();
    case 87:
      if (lookahead == 'e') ADVANCE(172);
      END_STATE();
    case 88:
      if (lookahead == 'x') ADVANCE(173);
      END_STATE();
    case 89:
      if (lookahead == 'a') ADVANCE(174);
      if (lookahead == 'i') ADVANCE(175);
      if (lookahead == 'o') ADVANCE(176);
      END_STATE();
    case 90:
      if (lookahead == 'n') ADVANCE(177);
      END_STATE();
    case 91:
      if (lookahead == 'e') ADVANCE(178);
      END_STATE();
    case 92:
      if (lookahead == 'e') ADVANCE(179);
      END_STATE();
    case 93:
      if (lookahead == 'p') ADVANCE(180);
      if (lookahead == 't') ADVANCE(181);
      END_STATE();
    case 94:
      if (lookahead == 'h') ADVANCE(182);
      if (lookahead == 'r') ADVANCE(183);
      END_STATE();
    case 95:
      if (lookahead == 'n') ADVANCE(184);
      if (lookahead == 'p') ADVANCE(185);
      END_STATE();
    case 96:
      if (lookahead == 'e') ADVANCE(186);
      if (lookahead == 'o') ADVANCE(187);
      END_STATE();
    case 97:
      if (lookahead == 'i') ADVANCE(188);
      END_STATE();
    case 98:
      if (lookahead == 'a') ADVANCE(189);
      END_STATE();
    case 99:
      if (lookahead == 'g') ADVANCE(190);
      END_STATE();
    case 100:
      ACCEPT_TOKEN(anon_sym_and);
      if (lookahead == '_') ADVANCE(191);
      END_STATE();
    case 101:
      ACCEPT_TOKEN(anon_sym_asm);
      END_STATE();
    case 102:
      if (lookahead == 'o') ADVANCE(192);
      END_STATE();
    case 103:
      if (lookahead == 'a') ADVANCE(193);
      if (lookahead == 'o') ADVANCE(194);
      END_STATE();
    case 104:
      if (lookahead == 'a') ADVANCE(195);
      END_STATE();
    case 105:
      if (lookahead == 'e') ADVANCE(196);
      END_STATE();
    case 106:
      if (lookahead == 'c') ADVANCE(197);
      END_STATE();
    case 107:
      if (lookahead == 's') ADVANCE(198);
      END_STATE();
    case 108:
      if (lookahead == 'a') ADVANCE(199);
      if (lookahead == 'r') ADVANCE(200);
      if (lookahead == 'y') ADVANCE(201);
      END_STATE();
    case 109:
      if (lookahead == 'p') ADVANCE(202);
      END_STATE();
    case 110:
      if (lookahead == 'c') ADVANCE(203);
      if (lookahead == 's') ADVANCE(204);
      if (lookahead == 't') ADVANCE(205);
      END_STATE();
    case 111:
      if (lookahead == 'l') ADVANCE(206);
      END_STATE();
    case 112:
      if (lookahead == 'a') ADVANCE(207);
      if (lookahead == 'i') ADVANCE(208);
      END_STATE();
    case 113:
      if (lookahead == 'e') ADVANCE(209);
      END_STATE();
    case 114:
      if (lookahead == 'e') ADVANCE(210);
      END_STATE();
    case 115:
      if (lookahead == 'm') ADVANCE(211);
      END_STATE();
    case 116:
      if (lookahead == 'l') ADVANCE(212);
      END_STATE();
    case 117:
      if (lookahead == 'e') ADVANCE(213);
      END_STATE();
    case 118:
      if (lookahead == 'a') ADVANCE(214);
      END_STATE();
    case 119:
      ACCEPT_TOKEN(anon_sym_for);
      END_STATE();
    case 120:
      if (lookahead == 'e') ADVANCE(215);
      END_STATE();
    case 121:
      if (lookahead == 'o') ADVANCE(216);
      END_STATE();
    case 122:
      if (lookahead == 'i') ADVANCE(217);
      END_STATE();
    case 123:
      if (lookahead == 'g') ADVANCE(218);
      END_STATE();
    case 124:
      if (lookahead == 'a') ADVANCE(219);
      END_STATE();
    case 125:
      if (lookahead == 'e') ADVANCE(220);
      END_STATE();
    case 126:
      ACCEPT_TOKEN(anon_sym_new);
      END_STATE();
    case 127:
      if (lookahead == 'x') ADVANCE(221);
      END_STATE();
    case 128:
      if (lookahead == 'e') ADVANCE(222);
      END_STATE();
    case 129:
      ACCEPT_TOKEN(anon_sym_not);
      if (lookahead == '_') ADVANCE(223);
      END_STATE();
    case 130:
      if (lookahead == 'l') ADVANCE(224);
      END_STATE();
    case 131:
      if (lookahead == 's') ADVANCE(225);
      END_STATE();
    case 132:
      if (lookahead == 'r') ADVANCE(226);
      END_STATE();
    case 133:
      if (lookahead == 'e') ADVANCE(227);
      END_STATE();
    case 134:
      if (lookahead == 'r') ADVANCE(228);
      END_STATE();
    case 135:
      if (lookahead == 'v') ADVANCE(229);
      END_STATE();
    case 136:
      if (lookahead == 't') ADVANCE(230);
      END_STATE();
    case 137:
      if (lookahead == 'l') ADVANCE(231);
      END_STATE();
    case 138:
      if (lookahead == 'i') ADVANCE(232);
      END_STATE();
    case 139:
      if (lookahead == 'u') ADVANCE(233);
      END_STATE();
    case 140:
      if (lookahead == 't') ADVANCE(234);
      END_STATE();
    case 141:
      if (lookahead == 'u') ADVANCE(235);
      END_STATE();
    case 142:
      if (lookahead == 'r') ADVANCE(236);
      END_STATE();
    case 143:
      if (lookahead == 'n') ADVANCE(237);
      END_STATE();
    case 144:
      if (lookahead == 'e') ADVANCE(238);
      END_STATE();
    case 145:
      if (lookahead == 't') ADVANCE(239);
      END_STATE();
    case 146:
      if (lookahead == 'u') ADVANCE(240);
      END_STATE();
    case 147:
      if (lookahead == 't') ADVANCE(241);
      END_STATE();
    case 148:
      if (lookahead == 'p') ADVANCE(242);
      END_STATE();
    case 149:
      if (lookahead == 's') ADVANCE(243);
      END_STATE();
    case 150:
      if (lookahead == 'e') ADVANCE(244);
      if (lookahead == 'o') ADVANCE(245);
      END_STATE();
    case 151:
      ACCEPT_TOKEN(anon_sym_try);
      END_STATE();
    case 152:
      if (lookahead == 'e') ADVANCE(246);
      END_STATE();
    case 153:
      if (lookahead == 'o') ADVANCE(247);
      END_STATE();
    case 154:
      if (lookahead == 'i') ADVANCE(248);
      END_STATE();
    case 155:
      if (lookahead == 'n') ADVANCE(249);
      END_STATE();
    case 156:
      if (lookahead == 't') ADVANCE(250);
      END_STATE();
    case 157:
      if (lookahead == 'a') ADVANCE(251);
      END_STATE();
    case 158:
      if (lookahead == 'l') ADVANCE(252);
      END_STATE();
    case 159:
      ACCEPT_TOKEN(anon_sym_xor);
      if (lookahead == '_') ADVANCE(253);
      END_STATE();
    case 160:
      ACCEPT_TOKEN(anon_sym_NULL);
      END_STATE();
    case 161:
      if (lookahead == 'g') ADVANCE(254);
      END_STATE();
    case 162:
      if (lookahead == 'm') ADVANCE(255);
      END_STATE();
    case 163:
      if (lookahead == 'e') ADVANCE(256);
      END_STATE();
    case 164:
      if (lookahead == 'n') ADVANCE(257);
      END_STATE();
    case 165:
      if (lookahead == 'e') ADVANCE(258);
      END_STATE();
    case 166:
      if (lookahead == 'i') ADVANCE(259);
      END_STATE();
    case 167:
      if (lookahead == 'm') ADVANCE(260);
      END_STATE();
    case 168:
      if (lookahead == 't') ADVANCE(261);
      END_STATE();
    case 169:
      if (lookahead == 's') ADVANCE(262);
      END_STATE();
    case 170:
      if (lookahead == 'e') ADVANCE(263);
      END_STATE();
    case 171:
      if (lookahead == 'r') ADVANCE(264);
      END_STATE();
    case 172:
      if (lookahead == 'c') ADVANCE(265);
      END_STATE();
    case 173:
      if (lookahead == 'c') ADVANCE(266);
      if (lookahead == 't') ADVANCE(267);
      END_STATE();
    case 174:
      if (lookahead == 's') ADVANCE(268);
      END_STATE();
    case 175:
      if (lookahead == 'n') ADVANCE(269);
      END_STATE();
    case 176:
      if (lookahead == 'r') ADVANCE(270);
      END_STATE();
    case 177:
      if (lookahead == 'l') ADVANCE(271);
      END_STATE();
    case 178:
      if (lookahead == 'a') ADVANCE(272);
      END_STATE();
    case 179:
      if (lookahead == 's') ADVANCE(273);
      END_STATE();
    case 180:
      if (lookahead == 't') ADVANCE(274);
      END_STATE();
    case 181:
      if (lookahead == 'd') ADVANCE(275);
      END_STATE();
    case 182:
      if (lookahead == 'i') ADVANCE(276);
      if (lookahead == 'r') ADVANCE(277);
      END_STATE();
    case 183:
      if (lookahead == 'y') ADVANCE(278);
      END_STATE();
    case 184:
      if (lookahead == 'a') ADVANCE(279);
      END_STATE();
    case 185:
      if (lookahead == 't') ADVANCE(280);
      END_STATE();
    case 186:
      if (lookahead == 'c') ADVANCE(281);
      END_STATE();
    case 187:
      if (lookahead == 'l') ADVANCE(282);
      END_STATE();
    case 188:
      if (lookahead == 'g') ADVANCE(283);
      END_STATE();
    case 189:
      if (lookahead == 'l') ADVANCE(284);
      END_STATE();
    case 190:
      if (lookahead == 'n') ADVANCE(285);
      END_STATE();
    case 191:
      if (lookahead == 'e') ADVANCE(286);
      END_STATE();
    case 192:
      ACCEPT_TOKEN(sym_auto);
      END_STATE();
    case 193:
      if (lookahead == 'n') ADVANCE(287);
      END_STATE();
    case 194:
      if (lookahead == 'r') ADVANCE(288);
      END_STATE();
    case 195:
      if (lookahead == 'k') ADVANCE(289);
      END_STATE();
    case 196:
      ACCEPT_TOKEN(anon_sym_case);
      END_STATE();
    case 197:
      if (lookahead == 'h') ADVANCE(290);
      END_STATE();
    case 198:
      if (lookahead == 's') ADVANCE(291);
      END_STATE();
    case 199:
      if (lookahead == 'w') ADVANCE(292);
      END_STATE();
    case 200:
      if (lookahead == 'e') ADVANCE(293);
      END_STATE();
    case 201:
      if (lookahead == 'i') ADVANCE(294);
      END_STATE();
    case 202:
      if (lookahead == 'l') ADVANCE(295);
      END_STATE();
    case 203:
      if (lookahead == 'e') ADVANCE(296);
      END_STATE();
    case 204:
      if (lookahead == 't') ADVANCE(297);
      END_STATE();
    case 205:
      if (lookahead == 'i') ADVANCE(298);
      END_STATE();
    case 206:
      if (lookahead == 't') ADVANCE(299);
      END_STATE();
    case 207:
      if (lookahead == 'u') ADVANCE(300);
      END_STATE();
    case 208:
      if (lookahead == 'n') ADVANCE(301);
      END_STATE();
    case 209:
      if (lookahead == 't') ADVANCE(302);
      END_STATE();
    case 210:
      ACCEPT_TOKEN(anon_sym_else);
      END_STATE();
    case 211:
      ACCEPT_TOKEN(anon_sym_enum);
      END_STATE();
    case 212:
      if (lookahead == 'i') ADVANCE(303);
      END_STATE();
    case 213:
      if (lookahead == 'r') ADVANCE(304);
      END_STATE();
    case 214:
      if (lookahead == 'l') ADVANCE(305);
      END_STATE();
    case 215:
      if (lookahead == 'n') ADVANCE(306);
      END_STATE();
    case 216:
      ACCEPT_TOKEN(anon_sym_goto);
      END_STATE();
    case 217:
      if (lookahead == 'n') ADVANCE(307);
      END_STATE();
    case 218:
      ACCEPT_TOKEN(anon_sym_long);
      END_STATE();
    case 219:
      if (lookahead == 'b') ADVANCE(308);
      END_STATE();
    case 220:
      if (lookahead == 's') ADVANCE(309);
      END_STATE();
    case 221:
      if (lookahead == 'c') ADVANCE(310);
      END_STATE();
    case 222:
      if (lookahead == 't') ADVANCE(311);
      END_STATE();
    case 223:
      if (lookahead == 'e') ADVANCE(312);
      END_STATE();
    case 224:
      if (lookahead == 'p') ADVANCE(313);
      END_STATE();
    case 225:
      if (lookahead == 'e') ADVANCE(314);
      END_STATE();
    case 226:
      if (lookahead == 'a') ADVANCE(315);
      END_STATE();
    case 227:
      if (lookahead == 'q') ADVANCE(316);
      END_STATE();
    case 228:
      if (lookahead == 'r') ADVANCE(317);
      END_STATE();
    case 229:
      if (lookahead == 'a') ADVANCE(318);
      END_STATE();
    case 230:
      if (lookahead == 'e') ADVANCE(319);
      END_STATE();
    case 231:
      if (lookahead == 'i') ADVANCE(320);
      END_STATE();
    case 232:
      if (lookahead == 's') ADVANCE(321);
      END_STATE();
    case 233:
      if (lookahead == 'i') ADVANCE(322);
      END_STATE();
    case 234:
      if (lookahead == 'r') ADVANCE(323);
      END_STATE();
    case 235:
      if (lookahead == 'r') ADVANCE(324);
      END_STATE();
    case 236:
      if (lookahead == 't') ADVANCE(325);
      END_STATE();
    case 237:
      if (lookahead == 'e') ADVANCE(326);
      END_STATE();
    case 238:
      if (lookahead == 'o') ADVANCE(327);
      END_STATE();
    case 239:
      if (lookahead == 'i') ADVANCE(328);
      END_STATE();
    case 240:
      if (lookahead == 'c') ADVANCE(329);
      END_STATE();
    case 241:
      if (lookahead == 'c') ADVANCE(330);
      END_STATE();
    case 242:
      if (lookahead == 'l') ADVANCE(331);
      END_STATE();
    case 243:
      ACCEPT_TOKEN(sym_this);
      END_STATE();
    case 244:
      if (lookahead == 'a') ADVANCE(332);
      END_STATE();
    case 245:
      if (lookahead == 'w') ADVANCE(333);
      END_STATE();
    case 246:
      if (lookahead == 'd') ADVANCE(334);
      if (lookahead == 'n') ADVANCE(335);
      END_STATE();
    case 247:
      if (lookahead == 'n') ADVANCE(336);
      END_STATE();
    case 248:
      if (lookahead == 'g') ADVANCE(337);
      END_STATE();
    case 249:
      if (lookahead == 'g') ADVANCE(338);
      END_STATE();
    case 250:
      if (lookahead == 'u') ADVANCE(339);
      END_STATE();
    case 251:
      if (lookahead == 't') ADVANCE(340);
      END_STATE();
    case 252:
      if (lookahead == 'e') ADVANCE(341);
      END_STATE();
    case 253:
      if (lookahead == 'e') ADVANCE(342);
      END_STATE();
    case 254:
      if (lookahead == 'n') ADVANCE(343);
      END_STATE();
    case 255:
      if (lookahead == 'i') ADVANCE(344);
      END_STATE();
    case 256:
      if (lookahead == 'r') ADVANCE(345);
      END_STATE();
    case 257:
      if (lookahead == 'u') ADVANCE(346);
      END_STATE();
    case 258:
      if (lookahead == 't') ADVANCE(347);
      END_STATE();
    case 259:
      if (lookahead == 'g') ADVANCE(348);
      END_STATE();
    case 260:
      ACCEPT_TOKEN(anon_sym___asm);
      if (lookahead == '_') ADVANCE(349);
      END_STATE();
    case 261:
      if (lookahead == 'r') ADVANCE(350);
      END_STATE();
    case 262:
      if (lookahead == 'e') ADVANCE(351);
      END_STATE();
    case 263:
      if (lookahead == 'c') ADVANCE(352);
      END_STATE();
    case 264:
      if (lookahead == 'c') ADVANCE(353);
      END_STATE();
    case 265:
      if (lookahead == 'l') ADVANCE(354);
      END_STATE();
    case 266:
      if (lookahead == 'e') ADVANCE(355);
      END_STATE();
    case 267:
      if (lookahead == 'e') ADVANCE(356);
      END_STATE();
    case 268:
      if (lookahead == 't') ADVANCE(357);
      END_STATE();
    case 269:
      if (lookahead == 'a') ADVANCE(358);
      END_STATE();
    case 270:
      if (lookahead == 'c') ADVANCE(359);
      END_STATE();
    case 271:
      if (lookahead == 'i') ADVANCE(360);
      END_STATE();
    case 272:
      if (lookahead == 'v') ADVANCE(361);
      END_STATE();
    case 273:
      if (lookahead == 't') ADVANCE(362);
      END_STATE();
    case 274:
      if (lookahead == 'r') ADVANCE(363);
      END_STATE();
    case 275:
      if (lookahead == 'c') ADVANCE(364);
      END_STATE();
    case 276:
      if (lookahead == 's') ADVANCE(365);
      END_STATE();
    case 277:
      if (lookahead == 'e') ADVANCE(366);
      END_STATE();
    case 278:
      ACCEPT_TOKEN(anon_sym___try);
      END_STATE();
    case 279:
      if (lookahead == 'l') ADVANCE(367);
      END_STATE();
    case 280:
      if (lookahead == 'r') ADVANCE(368);
      END_STATE();
    case 281:
      if (lookahead == 't') ADVANCE(369);
      END_STATE();
    case 282:
      if (lookahead == 'a') ADVANCE(370);
      END_STATE();
    case 283:
      if (lookahead == 'n') ADVANCE(371);
      END_STATE();
    case 284:
      if (lookahead == 'i') ADVANCE(372);
      END_STATE();
    case 285:
      if (lookahead == 'a') ADVANCE(373);
      if (lookahead == 'o') ADVANCE(374);
      END_STATE();
    case 286:
      if (lookahead == 'q') ADVANCE(375);
      END_STATE();
    case 287:
      if (lookahead == 'd') ADVANCE(376);
      END_STATE();
    case 288:
      ACCEPT_TOKEN(anon_sym_bitor);
      END_STATE();
    case 289:
      ACCEPT_TOKEN(anon_sym_break);
      END_STATE();
    case 290:
      ACCEPT_TOKEN(anon_sym_catch);
      END_STATE();
    case 291:
      ACCEPT_TOKEN(anon_sym_class);
      END_STATE();
    case 292:
      if (lookahead == 'a') ADVANCE(377);
      END_STATE();
    case 293:
      if (lookahead == 't') ADVANCE(378);
      END_STATE();
    case 294:
      if (lookahead == 'e') ADVANCE(379);
      END_STATE();
    case 295:
      ACCEPT_TOKEN(anon_sym_compl);
      END_STATE();
    case 296:
      if (lookahead == 'p') ADVANCE(380);
      END_STATE();
    case 297:
      ACCEPT_TOKEN(anon_sym_const);
      if (lookahead == 'e') ADVANCE(381);
      if (lookahead == 'i') ADVANCE(382);
      END_STATE();
    case 298:
      if (lookahead == 'n') ADVANCE(383);
      END_STATE();
    case 299:
      if (lookahead == 'y') ADVANCE(384);
      END_STATE();
    case 300:
      if (lookahead == 'l') ADVANCE(385);
      END_STATE();
    case 301:
      if (lookahead == 'e') ADVANCE(386);
      END_STATE();
    case 302:
      if (lookahead == 'e') ADVANCE(387);
      END_STATE();
    case 303:
      if (lookahead == 'c') ADVANCE(388);
      END_STATE();
    case 304:
      if (lookahead == 'n') ADVANCE(389);
      END_STATE();
    case 305:
      ACCEPT_TOKEN(anon_sym_final);
      END_STATE();
    case 306:
      if (lookahead == 'd') ADVANCE(390);
      END_STATE();
    case 307:
      if (lookahead == 'e') ADVANCE(391);
      END_STATE();
    case 308:
      if (lookahead == 'l') ADVANCE(392);
      END_STATE();
    case 309:
      if (lookahead == 'p') ADVANCE(393);
      END_STATE();
    case 310:
      if (lookahead == 'e') ADVANCE(394);
      END_STATE();
    case 311:
      if (lookahead == 'u') ADVANCE(395);
      END_STATE();
    case 312:
      if (lookahead == 'q') ADVANCE(396);
      END_STATE();
    case 313:
      if (lookahead == 't') ADVANCE(397);
      END_STATE();
    case 314:
      if (lookahead == 't') ADVANCE(398);
      END_STATE();
    case 315:
      if (lookahead == 't') ADVANCE(399);
      END_STATE();
    case 316:
      ACCEPT_TOKEN(anon_sym_or_eq);
      END_STATE();
    case 317:
      if (lookahead == 'i') ADVANCE(400);
      END_STATE();
    case 318:
      if (lookahead == 't') ADVANCE(401);
      END_STATE();
    case 319:
      if (lookahead == 'c') ADVANCE(402);
      END_STATE();
    case 320:
      if (lookahead == 'c') ADVANCE(403);
      END_STATE();
    case 321:
      if (lookahead == 't') ADVANCE(404);
      END_STATE();
    case 322:
      if (lookahead == 'r') ADVANCE(405);
      END_STATE();
    case 323:
      if (lookahead == 'i') ADVANCE(406);
      END_STATE();
    case 324:
      if (lookahead == 'n') ADVANCE(407);
      END_STATE();
    case 325:
      ACCEPT_TOKEN(anon_sym_short);
      END_STATE();
    case 326:
      if (lookahead == 'd') ADVANCE(408);
      END_STATE();
    case 327:
      if (lookahead == 'f') ADVANCE(409);
      END_STATE();
    case 328:
      if (lookahead == 'c') ADVANCE(410);
      END_STATE();
    case 329:
      if (lookahead == 't') ADVANCE(411);
      END_STATE();
    case 330:
      if (lookahead == 'h') ADVANCE(412);
      END_STATE();
    case 331:
      if (lookahead == 'a') ADVANCE(413);
      END_STATE();
    case 332:
      if (lookahead == 'd') ADVANCE(414);
      END_STATE();
    case 333:
      ACCEPT_TOKEN(anon_sym_throw);
      END_STATE();
    case 334:
      if (lookahead == 'e') ADVANCE(415);
      END_STATE();
    case 335:
      if (lookahead == 'a') ADVANCE(416);
      END_STATE();
    case 336:
      ACCEPT_TOKEN(anon_sym_union);
      END_STATE();
    case 337:
      if (lookahead == 'n') ADVANCE(417);
      END_STATE();
    case 338:
      ACCEPT_TOKEN(anon_sym_using);
      END_STATE();
    case 339:
      if (lookahead == 'a') ADVANCE(418);
      END_STATE();
    case 340:
      if (lookahead == 'i') ADVANCE(419);
      END_STATE();
    case 341:
      ACCEPT_TOKEN(anon_sym_while);
      END_STATE();
    case 342:
      if (lookahead == 'q') ADVANCE(420);
      END_STATE();
    case 343:
      if (lookahead == 'a') ADVANCE(421);
      if (lookahead == 'o') ADVANCE(422);
      END_STATE();
    case 344:
      if (lookahead == 'c') ADVANCE(423);
      END_STATE();
    case 345:
      if (lookahead == 'i') ADVANCE(424);
      END_STATE();
    case 346:
      if (lookahead == 'l') ADVANCE(425);
      END_STATE();
    case 347:
      if (lookahead == 'u') ADVANCE(426);
      END_STATE();
    case 348:
      if (lookahead == 'n') ADVANCE(427);
      END_STATE();
    case 349:
      if (lookahead == '_') ADVANCE(428);
      END_STATE();
    case 350:
      if (lookahead == 'i') ADVANCE(429);
      END_STATE();
    case 351:
      if (lookahead == 'd') ADVANCE(430);
      END_STATE();
    case 352:
      if (lookahead == 'l') ADVANCE(431);
      END_STATE();
    case 353:
      if (lookahead == 'a') ADVANCE(432);
      END_STATE();
    case 354:
      if (lookahead == 's') ADVANCE(433);
      END_STATE();
    case 355:
      if (lookahead == 'p') ADVANCE(434);
      END_STATE();
    case 356:
      if (lookahead == 'n') ADVANCE(435);
      END_STATE();
    case 357:
      if (lookahead == 'c') ADVANCE(436);
      END_STATE();
    case 358:
      if (lookahead == 'l') ADVANCE(437);
      END_STATE();
    case 359:
      if (lookahead == 'e') ADVANCE(438);
      END_STATE();
    case 360:
      if (lookahead == 'n') ADVANCE(439);
      END_STATE();
    case 361:
      if (lookahead == 'e') ADVANCE(440);
      END_STATE();
    case 362:
      if (lookahead == 'r') ADVANCE(441);
      END_STATE();
    case 363:
      ACCEPT_TOKEN(sym_ms_signed_ptr_modifier);
      END_STATE();
    case 364:
      if (lookahead == 'a') ADVANCE(442);
      END_STATE();
    case 365:
      if (lookahead == 'c') ADVANCE(443);
      END_STATE();
    case 366:
      if (lookahead == 'a') ADVANCE(444);
      END_STATE();
    case 367:
      if (lookahead == 'i') ADVANCE(445);
      END_STATE();
    case 368:
      ACCEPT_TOKEN(sym_ms_unsigned_ptr_modifier);
      END_STATE();
    case 369:
      if (lookahead == 'o') ADVANCE(446);
      END_STATE();
    case 370:
      if (lookahead == 't') ADVANCE(447);
      END_STATE();
    case 371:
      if (lookahead == 'o') ADVANCE(448);
      END_STATE();
    case 372:
      if (lookahead == 'g') ADVANCE(449);
      END_STATE();
    case 373:
      if (lookahead == 's') ADVANCE(450);
      END_STATE();
    case 374:
      if (lookahead == 'f') ADVANCE(451);
      END_STATE();
    case 375:
      ACCEPT_TOKEN(anon_sym_and_eq);
      END_STATE();
    case 376:
      ACCEPT_TOKEN(anon_sym_bitand);
      END_STATE();
    case 377:
      if (lookahead == 'i') ADVANCE(452);
      END_STATE();
    case 378:
      if (lookahead == 'u') ADVANCE(453);
      END_STATE();
    case 379:
      if (lookahead == 'l') ADVANCE(454);
      END_STATE();
    case 380:
      if (lookahead == 't') ADVANCE(455);
      END_STATE();
    case 381:
      if (lookahead == 'v') ADVANCE(456);
      if (lookahead == 'x') ADVANCE(457);
      END_STATE();
    case 382:
      if (lookahead == 'n') ADVANCE(458);
      END_STATE();
    case 383:
      if (lookahead == 'u') ADVANCE(459);
      END_STATE();
    case 384:
      if (lookahead == 'p') ADVANCE(460);
      END_STATE();
    case 385:
      if (lookahead == 't') ADVANCE(461);
      END_STATE();
    case 386:
      if (lookahead == 'd') ADVANCE(462);
      END_STATE();
    case 387:
      ACCEPT_TOKEN(anon_sym_delete);
      END_STATE();
    case 388:
      if (lookahead == 'i') ADVANCE(463);
      END_STATE();
    case 389:
      ACCEPT_TOKEN(anon_sym_extern);
      END_STATE();
    case 390:
      ACCEPT_TOKEN(anon_sym_friend);
      END_STATE();
    case 391:
      ACCEPT_TOKEN(anon_sym_inline);
      END_STATE();
    case 392:
      if (lookahead == 'e') ADVANCE(464);
      END_STATE();
    case 393:
      if (lookahead == 'a') ADVANCE(465);
      END_STATE();
    case 394:
      if (lookahead == 'p') ADVANCE(466);
      END_STATE();
    case 395:
      if (lookahead == 'r') ADVANCE(467);
      END_STATE();
    case 396:
      ACCEPT_TOKEN(anon_sym_not_eq);
      END_STATE();
    case 397:
      if (lookahead == 'r') ADVANCE(468);
      END_STATE();
    case 398:
      if (lookahead == 'o') ADVANCE(469);
      END_STATE();
    case 399:
      if (lookahead == 'o') ADVANCE(470);
      END_STATE();
    case 400:
      if (lookahead == 'd') ADVANCE(471);
      END_STATE();
    case 401:
      if (lookahead == 'e') ADVANCE(472);
      END_STATE();
    case 402:
      if (lookahead == 't') ADVANCE(473);
      END_STATE();
    case 403:
      ACCEPT_TOKEN(anon_sym_public);
      END_STATE();
    case 404:
      if (lookahead == 'e') ADVANCE(474);
      END_STATE();
    case 405:
      if (lookahead == 'e') ADVANCE(475);
      END_STATE();
    case 406:
      if (lookahead == 'c') ADVANCE(476);
      END_STATE();
    case 407:
      ACCEPT_TOKEN(anon_sym_return);
      END_STATE();
    case 408:
      ACCEPT_TOKEN(anon_sym_signed);
      END_STATE();
    case 409:
      ACCEPT_TOKEN(anon_sym_sizeof);
      END_STATE();
    case 410:
      ACCEPT_TOKEN(anon_sym_static);
      if (lookahead == '_') ADVANCE(477);
      END_STATE();
    case 411:
      ACCEPT_TOKEN(anon_sym_struct);
      END_STATE();
    case 412:
      ACCEPT_TOKEN(anon_sym_switch);
      END_STATE();
    case 413:
      if (lookahead == 't') ADVANCE(478);
      END_STATE();
    case 414:
      if (lookahead == '_') ADVANCE(479);
      END_STATE();
    case 415:
      if (lookahead == 'f') ADVANCE(480);
      END_STATE();
    case 416:
      if (lookahead == 'm') ADVANCE(481);
      END_STATE();
    case 417:
      if (lookahead == 'e') ADVANCE(482);
      END_STATE();
    case 418:
      if (lookahead == 'l') ADVANCE(483);
      END_STATE();
    case 419:
      if (lookahead == 'l') ADVANCE(484);
      END_STATE();
    case 420:
      ACCEPT_TOKEN(anon_sym_xor_eq);
      END_STATE();
    case 421:
      if (lookahead == 's') ADVANCE(485);
      END_STATE();
    case 422:
      if (lookahead == 'f') ADVANCE(486);
      END_STATE();
    case 423:
      ACCEPT_TOKEN(anon_sym__Atomic);
      END_STATE();
    case 424:
      if (lookahead == 'c') ADVANCE(487);
      END_STATE();
    case 425:
      if (lookahead == 'l') ADVANCE(488);
      END_STATE();
    case 426:
      if (lookahead == 'r') ADVANCE(489);
      END_STATE();
    case 427:
      if (lookahead == 'o') ADVANCE(490);
      END_STATE();
    case 428:
      ACCEPT_TOKEN(anon_sym___asm__);
      END_STATE();
    case 429:
      if (lookahead == 'b') ADVANCE(491);
      END_STATE();
    case 430:
      ACCEPT_TOKEN(anon_sym___based);
      END_STATE();
    case 431:
      ACCEPT_TOKEN(anon_sym___cdecl);
      END_STATE();
    case 432:
      if (lookahead == 'l') ADVANCE(492);
      END_STATE();
    case 433:
      if (lookahead == 'p') ADVANCE(493);
      END_STATE();
    case 434:
      if (lookahead == 't') ADVANCE(494);
      END_STATE();
    case 435:
      if (lookahead == 's') ADVANCE(495);
      END_STATE();
    case 436:
      if (lookahead == 'a') ADVANCE(496);
      END_STATE();
    case 437:
      if (lookahead == 'l') ADVANCE(497);
      END_STATE();
    case 438:
      if (lookahead == 'i') ADVANCE(498);
      END_STATE();
    case 439:
      if (lookahead == 'e') ADVANCE(499);
      END_STATE();
    case 440:
      ACCEPT_TOKEN(anon_sym___leave);
      END_STATE();
    case 441:
      if (lookahead == 'i') ADVANCE(500);
      END_STATE();
    case 442:
      if (lookahead == 'l') ADVANCE(501);
      END_STATE();
    case 443:
      if (lookahead == 'a') ADVANCE(502);
      END_STATE();
    case 444:
      if (lookahead == 'd') ADVANCE(503);
      END_STATE();
    case 445:
      if (lookahead == 'g') ADVANCE(504);
      END_STATE();
    case 446:
      if (lookahead == 'r') ADVANCE(505);
      END_STATE();
    case 447:
      if (lookahead == 'i') ADVANCE(506);
      END_STATE();
    case 448:
      if (lookahead == 'f') ADVANCE(507);
      END_STATE();
    case 449:
      if (lookahead == 'n') ADVANCE(508);
      END_STATE();
    case 450:
      ACCEPT_TOKEN(anon_sym_alignas);
      END_STATE();
    case 451:
      ACCEPT_TOKEN(anon_sym_alignof);
      END_STATE();
    case 452:
      if (lookahead == 't') ADVANCE(509);
      END_STATE();
    case 453:
      if (lookahead == 'r') ADVANCE(510);
      END_STATE();
    case 454:
      if (lookahead == 'd') ADVANCE(511);
      END_STATE();
    case 455:
      ACCEPT_TOKEN(anon_sym_concept);
      END_STATE();
    case 456:
      if (lookahead == 'a') ADVANCE(512);
      END_STATE();
    case 457:
      if (lookahead == 'p') ADVANCE(513);
      END_STATE();
    case 458:
      if (lookahead == 'i') ADVANCE(514);
      END_STATE();
    case 459:
      if (lookahead == 'e') ADVANCE(515);
      END_STATE();
    case 460:
      if (lookahead == 'e') ADVANCE(516);
      END_STATE();
    case 461:
      ACCEPT_TOKEN(anon_sym_default);
      END_STATE();
    case 462:
      ACCEPT_TOKEN(anon_sym_defined);
      END_STATE();
    case 463:
      if (lookahead == 't') ADVANCE(517);
      END_STATE();
    case 464:
      ACCEPT_TOKEN(anon_sym_mutable);
      END_STATE();
    case 465:
      if (lookahead == 'c') ADVANCE(518);
      END_STATE();
    case 466:
      if (lookahead == 't') ADVANCE(519);
      END_STATE();
    case 467:
      if (lookahead == 'n') ADVANCE(520);
      END_STATE();
    case 468:
      ACCEPT_TOKEN(anon_sym_nullptr);
      END_STATE();
    case 469:
      if (lookahead == 'f') ADVANCE(521);
      END_STATE();
    case 470:
      if (lookahead == 'r') ADVANCE(522);
      END_STATE();
    case 471:
      if (lookahead == 'e') ADVANCE(523);
      END_STATE();
    case 472:
      ACCEPT_TOKEN(anon_sym_private);
      END_STATE();
    case 473:
      if (lookahead == 'e') ADVANCE(524);
      END_STATE();
    case 474:
      if (lookahead == 'r') ADVANCE(525);
      END_STATE();
    case 475:
      if (lookahead == 's') ADVANCE(526);
      END_STATE();
    case 476:
      if (lookahead == 't') ADVANCE(527);
      END_STATE();
    case 477:
      if (lookahead == 'a') ADVANCE(528);
      END_STATE();
    case 478:
      if (lookahead == 'e') ADVANCE(529);
      END_STATE();
    case 479:
      if (lookahead == 'l') ADVANCE(530);
      END_STATE();
    case 480:
      ACCEPT_TOKEN(anon_sym_typedef);
      END_STATE();
    case 481:
      if (lookahead == 'e') ADVANCE(531);
      END_STATE();
    case 482:
      if (lookahead == 'd') ADVANCE(532);
      END_STATE();
    case 483:
      ACCEPT_TOKEN(anon_sym_virtual);
      END_STATE();
    case 484:
      if (lookahead == 'e') ADVANCE(533);
      END_STATE();
    case 485:
      ACCEPT_TOKEN(anon_sym__Alignas);
      END_STATE();
    case 486:
      ACCEPT_TOKEN(anon_sym__Alignof);
      END_STATE();
    case 487:
      ACCEPT_TOKEN(anon_sym__Generic);
      END_STATE();
    case 488:
      ACCEPT_TOKEN(anon_sym__Nonnull);
      END_STATE();
    case 489:
      if (lookahead == 'n') ADVANCE(534);
      END_STATE();
    case 490:
      if (lookahead == 'f') ADVANCE(535);
      END_STATE();
    case 491:
      if (lookahead == 'u') ADVANCE(536);
      END_STATE();
    case 492:
      if (lookahead == 'l') ADVANCE(537);
      END_STATE();
    case 493:
      if (lookahead == 'e') ADVANCE(538);
      END_STATE();
    case 494:
      ACCEPT_TOKEN(anon_sym___except);
      END_STATE();
    case 495:
      if (lookahead == 'i') ADVANCE(539);
      END_STATE();
    case 496:
      if (lookahead == 'l') ADVANCE(540);
      END_STATE();
    case 497:
      if (lookahead == 'y') ADVANCE(541);
      END_STATE();
    case 498:
      if (lookahead == 'n') ADVANCE(542);
      END_STATE();
    case 499:
      ACCEPT_TOKEN(anon_sym___inline);
      if (lookahead == '_') ADVANCE(543);
      END_STATE();
    case 500:
      if (lookahead == 'c') ADVANCE(544);
      END_STATE();
    case 501:
      if (lookahead == 'l') ADVANCE(545);
      END_STATE();
    case 502:
      if (lookahead == 'l') ADVANCE(546);
      END_STATE();
    case 503:
      ACCEPT_TOKEN(anon_sym___thread);
      END_STATE();
    case 504:
      if (lookahead == 'n') ADVANCE(547);
      END_STATE();
    case 505:
      if (lookahead == 'c') ADVANCE(548);
      END_STATE();
    case 506:
      if (lookahead == 'l') ADVANCE(549);
      END_STATE();
    case 507:
      ACCEPT_TOKEN(anon_sym__alignof);
      END_STATE();
    case 508:
      if (lookahead == 'e') ADVANCE(550);
      END_STATE();
    case 509:
      ACCEPT_TOKEN(anon_sym_co_await);
      END_STATE();
    case 510:
      if (lookahead == 'n') ADVANCE(551);
      END_STATE();
    case 511:
      ACCEPT_TOKEN(anon_sym_co_yield);
      END_STATE();
    case 512:
      if (lookahead == 'l') ADVANCE(552);
      END_STATE();
    case 513:
      if (lookahead == 'r') ADVANCE(553);
      END_STATE();
    case 514:
      if (lookahead == 't') ADVANCE(554);
      END_STATE();
    case 515:
      ACCEPT_TOKEN(anon_sym_continue);
      END_STATE();
    case 516:
      ACCEPT_TOKEN(anon_sym_decltype);
      END_STATE();
    case 517:
      ACCEPT_TOKEN(anon_sym_explicit);
      END_STATE();
    case 518:
      if (lookahead == 'e') ADVANCE(555);
      END_STATE();
    case 519:
      ACCEPT_TOKEN(anon_sym_noexcept);
      END_STATE();
    case 520:
      ACCEPT_TOKEN(anon_sym_noreturn);
      END_STATE();
    case 521:
      ACCEPT_TOKEN(anon_sym_offsetof);
      END_STATE();
    case 522:
      ACCEPT_TOKEN(anon_sym_operator);
      END_STATE();
    case 523:
      ACCEPT_TOKEN(anon_sym_override);
      END_STATE();
    case 524:
      if (lookahead == 'd') ADVANCE(556);
      END_STATE();
    case 525:
      ACCEPT_TOKEN(anon_sym_register);
      END_STATE();
    case 526:
      ACCEPT_TOKEN(anon_sym_requires);
      END_STATE();
    case 527:
      ACCEPT_TOKEN(anon_sym_restrict);
      END_STATE();
    case 528:
      if (lookahead == 's') ADVANCE(557);
      END_STATE();
    case 529:
      ACCEPT_TOKEN(anon_sym_template);
      END_STATE();
    case 530:
      if (lookahead == 'o') ADVANCE(558);
      END_STATE();
    case 531:
      ACCEPT_TOKEN(anon_sym_typename);
      END_STATE();
    case 532:
      ACCEPT_TOKEN(anon_sym_unsigned);
      END_STATE();
    case 533:
      ACCEPT_TOKEN(anon_sym_volatile);
      END_STATE();
    case 534:
      ACCEPT_TOKEN(anon_sym__Noreturn);
      END_STATE();
    case 535:
      ACCEPT_TOKEN(anon_sym___alignof);
      if (lookahead == '_') ADVANCE(559);
      END_STATE();
    case 536:
      if (lookahead == 't') ADVANCE(560);
      END_STATE();
    case 537:
      ACCEPT_TOKEN(anon_sym___clrcall);
      END_STATE();
    case 538:
      if (lookahead == 'c') ADVANCE(561);
      END_STATE();
    case 539:
      if (lookahead == 'o') ADVANCE(562);
      END_STATE();
    case 540:
      if (lookahead == 'l') ADVANCE(563);
      END_STATE();
    case 541:
      ACCEPT_TOKEN(anon_sym___finally);
      END_STATE();
    case 542:
      if (lookahead == 'l') ADVANCE(564);
      END_STATE();
    case 543:
      if (lookahead == '_') ADVANCE(565);
      END_STATE();
    case 544:
      if (lookahead == 't') ADVANCE(566);
      END_STATE();
    case 545:
      ACCEPT_TOKEN(anon_sym___stdcall);
      END_STATE();
    case 546:
      if (lookahead == 'l') ADVANCE(567);
      END_STATE();
    case 547:
      if (lookahead == 'e') ADVANCE(568);
      END_STATE();
    case 548:
      if (lookahead == 'a') ADVANCE(569);
      END_STATE();
    case 549:
      if (lookahead == 'e') ADVANCE(570);
      END_STATE();
    case 550:
      if (lookahead == 'd') ADVANCE(571);
      END_STATE();
    case 551:
      ACCEPT_TOKEN(anon_sym_co_return);
      END_STATE();
    case 552:
      ACCEPT_TOKEN(anon_sym_consteval);
      END_STATE();
    case 553:
      ACCEPT_TOKEN(anon_sym_constexpr);
      END_STATE();
    case 554:
      ACCEPT_TOKEN(anon_sym_constinit);
      END_STATE();
    case 555:
      ACCEPT_TOKEN(anon_sym_namespace);
      END_STATE();
    case 556:
      ACCEPT_TOKEN(anon_sym_protected);
      END_STATE();
    case 557:
      if (lookahead == 's') ADVANCE(572);
      END_STATE();
    case 558:
      if (lookahead == 'c') ADVANCE(573);
      END_STATE();
    case 559:
      if (lookahead == '_') ADVANCE(574);
      END_STATE();
    case 560:
      if (lookahead == 'e') ADVANCE(575);
      END_STATE();
    case 561:
      ACCEPT_TOKEN(anon_sym___declspec);
      END_STATE();
    case 562:
      if (lookahead == 'n') ADVANCE(576);
      END_STATE();
    case 563:
      ACCEPT_TOKEN(anon_sym___fastcall);
      END_STATE();
    case 564:
      if (lookahead == 'i') ADVANCE(577);
      END_STATE();
    case 565:
      ACCEPT_TOKEN(anon_sym___inline__);
      END_STATE();
    case 566:
      ACCEPT_TOKEN(sym_ms_restrict_modifier);
      if (lookahead == '_') ADVANCE(578);
      END_STATE();
    case 567:
      ACCEPT_TOKEN(anon_sym___thiscall);
      END_STATE();
    case 568:
      if (lookahead == 'd') ADVANCE(579);
      END_STATE();
    case 569:
      if (lookahead == 'l') ADVANCE(580);
      END_STATE();
    case 570:
      if (lookahead == '_') ADVANCE(581);
      END_STATE();
    case 571:
      ACCEPT_TOKEN(anon_sym__unaligned);
      END_STATE();
    case 572:
      if (lookahead == 'e') ADVANCE(582);
      END_STATE();
    case 573:
      if (lookahead == 'a') ADVANCE(583);
      END_STATE();
    case 574:
      ACCEPT_TOKEN(anon_sym___alignof__);
      END_STATE();
    case 575:
      ACCEPT_TOKEN(anon_sym___attribute);
      if (lookahead == '_') ADVANCE(584);
      END_STATE();
    case 576:
      if (lookahead == '_') ADVANCE(585);
      END_STATE();
    case 577:
      if (lookahead == 'n') ADVANCE(586);
      END_STATE();
    case 578:
      if (lookahead == '_') ADVANCE(587);
      END_STATE();
    case 579:
      ACCEPT_TOKEN(anon_sym___unaligned);
      END_STATE();
    case 580:
      if (lookahead == 'l') ADVANCE(588);
      END_STATE();
    case 581:
      if (lookahead == '_') ADVANCE(589);
      END_STATE();
    case 582:
      if (lookahead == 'r') ADVANCE(590);
      END_STATE();
    case 583:
      if (lookahead == 'l') ADVANCE(591);
      END_STATE();
    case 584:
      if (lookahead == '_') ADVANCE(592);
      END_STATE();
    case 585:
      if (lookahead == '_') ADVANCE(593);
      END_STATE();
    case 586:
      if (lookahead == 'e') ADVANCE(594);
      END_STATE();
    case 587:
      ACCEPT_TOKEN(anon_sym___restrict__);
      END_STATE();
    case 588:
      ACCEPT_TOKEN(anon_sym___vectorcall);
      END_STATE();
    case 589:
      ACCEPT_TOKEN(anon_sym___volatile__);
      END_STATE();
    case 590:
      if (lookahead == 't') ADVANCE(595);
      END_STATE();
    case 591:
      ACCEPT_TOKEN(anon_sym_thread_local);
      END_STATE();
    case 592:
      ACCEPT_TOKEN(anon_sym___attribute__);
      END_STATE();
    case 593:
      ACCEPT_TOKEN(anon_sym___extension__);
      END_STATE();
    case 594:
      ACCEPT_TOKEN(anon_sym___forceinline);
      END_STATE();
    case 595:
      ACCEPT_TOKEN(anon_sym_static_assert);
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
static uint16_t ts_small_parse_table[286030];

// Filled from tables.zlib by grammar.go, see fetchgrammar.
static uint32_t ts_small_parse_table_map[6315];

// Filled from tables.zlib by grammar.go, see fetchgrammar.
static TSParseActionEntry ts_parse_actions[14617];

enum ts_external_scanner_symbol_identifiers {
  ts_external_token_raw_string_delimiter = 0,
  ts_external_token_raw_string_content = 1,
};

static const TSSymbol ts_external_scanner_symbol_map[EXTERNAL_TOKEN_COUNT] = {
  [ts_external_token_raw_string_delimiter] = sym_raw_string_delimiter,
  [ts_external_token_raw_string_content] = sym_raw_string_content,
};

static const bool ts_external_scanner_states[4][EXTERNAL_TOKEN_COUNT] = {
  [1] = {
    [ts_external_token_raw_string_delimiter] = true,
    [ts_external_token_raw_string_content] = true,
  },
  [2] = {
    [ts_external_token_raw_string_delimiter] = true,
  },
  [3] = {
    [ts_external_token_raw_string_content] = true,
  },
};

#ifdef __cplusplus
extern "C" {
#endif
void *tree_sitter_cpp_external_scanner_create(void);
void tree_sitter_cpp_external_scanner_destroy(void *);
bool tree_sitter_cpp_external_scanner_scan(void *, TSLexer *, const bool *);
unsigned tree_sitter_cpp_external_scanner_serialize(void *, char *);
void tree_sitter_cpp_external_scanner_deserialize(void *, const char *, unsigned);

#ifdef TREE_SITTER_HIDE_SYMBOLS
#define TS_PUBLIC
#elif defined(_WIN32)
#define TS_PUBLIC __declspec(dllexport)
#else
#define TS_PUBLIC __attribute__((visibility("default")))
#endif

TS_PUBLIC const TSLanguage *tree_sitter_cpp(void) {
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
      tree_sitter_cpp_external_scanner_create,
      tree_sitter_cpp_external_scanner_destroy,
      tree_sitter_cpp_external_scanner_scan,
      tree_sitter_cpp_external_scanner_serialize,
      tree_sitter_cpp_external_scanner_deserialize,
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
