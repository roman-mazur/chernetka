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
#define STATE_COUNT 4165
#define LARGE_STATE_COUNT 405
#define SYMBOL_COUNT 182
#define ALIAS_COUNT 0
#define TOKEN_COUNT 116
#define EXTERNAL_TOKEN_COUNT 0
#define FIELD_COUNT 20
#define MAX_ALIAS_SEQUENCE_LENGTH 14
#define PRODUCTION_ID_COUNT 177

enum ts_symbol_identifiers {
  sym__ws = 1,
  sym_comment = 2,
  anon_sym_POUND_ = 3,
  anon_sym_POUND = 4,
  aux_sym__form_token1 = 5,
  anon_sym_DOT = 6,
  aux_sym_num_lit_token1 = 7,
  aux_sym_num_lit_token2 = 8,
  anon_sym_COLON = 9,
  anon_sym_COLON_COLON = 10,
  anon_sym_DQUOTE = 11,
  aux_sym_str_lit_token1 = 12,
  aux_sym_str_lit_token2 = 13,
  anon_sym_TILDE = 14,
  aux_sym_char_lit_token1 = 15,
  sym_nil_lit = 16,
  aux_sym_sym_lit_token1 = 17,
  anon_sym_CARET = 18,
  anon_sym_POUND_CARET = 19,
  anon_sym_LPAREN = 20,
  anon_sym_RPAREN = 21,
  anon_sym_LBRACE = 22,
  anon_sym_RBRACE = 23,
  anon_sym_POUND0A = 24,
  anon_sym_POUND0a = 25,
  anon_sym_POUND_QMARK = 26,
  anon_sym_POUND_QMARK_AT = 27,
  anon_sym_POUND_SQUOTE = 28,
  anon_sym_SQUOTE = 29,
  anon_sym_BQUOTE = 30,
  anon_sym_COMMA_AT = 31,
  anon_sym_COMMA = 32,
  sym_block_comment = 33,
  sym_fancy_literal = 34,
  aux_sym__format_token_token1 = 35,
  anon_sym_v = 36,
  anon_sym_V = 37,
  anon_sym_AT = 38,
  anon_sym_AT_COLON = 39,
  anon_sym_COLON_AT = 40,
  anon_sym_PERCENT = 41,
  anon_sym_AMP = 42,
  anon_sym_PIPE = 43,
  aux_sym_format_directive_type_token1 = 44,
  aux_sym_format_directive_type_token2 = 45,
  anon_sym_LF = 46,
  anon_sym_CR = 47,
  aux_sym_format_directive_type_token3 = 48,
  aux_sym_format_directive_type_token4 = 49,
  aux_sym_format_directive_type_token5 = 50,
  aux_sym_format_directive_type_token6 = 51,
  anon_sym__ = 52,
  aux_sym_format_directive_type_token7 = 53,
  aux_sym_format_directive_type_token8 = 54,
  aux_sym_format_directive_type_token9 = 55,
  aux_sym_format_directive_type_token10 = 56,
  anon_sym_SEMI = 57,
  anon_sym_STAR = 58,
  anon_sym_SLASH = 59,
  anon_sym_QMARK = 60,
  anon_sym_Newline = 61,
  aux_sym_format_directive_type_token11 = 62,
  anon_sym_cl = 63,
  anon_sym_in = 64,
  anon_sym_across = 65,
  anon_sym_being = 66,
  anon_sym_using = 67,
  aux_sym_for_clause_word_token1 = 68,
  anon_sym_below = 69,
  anon_sym_above = 70,
  anon_sym_from = 71,
  anon_sym_to = 72,
  anon_sym_upto = 73,
  anon_sym_upfrom = 74,
  anon_sym_downto = 75,
  anon_sym_downfrom = 76,
  anon_sym_on = 77,
  anon_sym_by = 78,
  anon_sym_then = 79,
  anon_sym_EQ = 80,
  aux_sym_accumulation_verb_token1 = 81,
  anon_sym_for = 82,
  anon_sym_and = 83,
  anon_sym_as = 84,
  anon_sym_with = 85,
  anon_sym_do = 86,
  anon_sym_while = 87,
  anon_sym_until = 88,
  anon_sym_repeat = 89,
  anon_sym_when = 90,
  anon_sym_if = 91,
  anon_sym_unless = 92,
  anon_sym_always = 93,
  anon_sym_thereis = 94,
  anon_sym_never = 95,
  anon_sym_else = 96,
  anon_sym_into = 97,
  anon_sym_finally = 98,
  anon_sym_return = 99,
  anon_sym_initially = 100,
  anon_sym_loop = 101,
  anon_sym_defun = 102,
  anon_sym_defmacro = 103,
  anon_sym_defgeneric = 104,
  anon_sym_defmethod = 105,
  anon_sym_lambda = 106,
  anon_sym_POUNDP = 107,
  anon_sym_POUNDp = 108,
  aux_sym_path_lit_token1 = 109,
  aux_sym__sym_lit_without_slash_token1 = 110,
  sym_self_referential_reader_macro = 111,
  anon_sym_POUND_PLUS = 112,
  anon_sym_POUND_DASH = 113,
  anon_sym_POUNDC = 114,
  anon_sym_POUNDc = 115,
  sym_source = 116,
  sym__gap = 117,
  sym_dis_expr = 118,
  sym__form = 119,
  sym_num_lit = 120,
  sym_kwd_lit = 121,
  sym_str_lit = 122,
  sym_char_lit = 123,
  sym_sym_lit = 124,
  sym__metadata_lit = 125,
  sym_meta_lit = 126,
  sym_old_meta_lit = 127,
  sym_list_lit = 128,
  sym__bare_list_lit = 129,
  sym_map_lit = 130,
  sym__bare_map_lit = 131,
  sym_vec_lit = 132,
  sym_set_lit = 133,
  sym__bare_set_lit = 134,
  sym_read_cond_lit = 135,
  sym_splicing_read_cond_lit = 136,
  sym_var_quoting_lit = 137,
  sym_quoting_lit = 138,
  sym_syn_quoting_lit = 139,
  sym_unquote_splicing_lit = 140,
  sym_unquoting_lit = 141,
  sym_defun = 142,
  sym__format_token = 143,
  sym_format_prefix_parameters = 144,
  sym_format_modifiers = 145,
  sym_format_directive_type = 146,
  sym_format_specifier = 147,
  sym_for_clause_word = 148,
  sym__for_part = 149,
  sym_accumulation_verb = 150,
  sym_for_clause = 151,
  sym_with_clause = 152,
  sym_do_clause = 153,
  sym_while_clause = 154,
  sym_repeat_clause = 155,
  sym_condition_clause = 156,
  sym_accumulation_clause = 157,
  sym_termination_clause = 158,
  sym_loop_clause = 159,
  sym_loop_macro = 160,
  sym_defun_keyword = 161,
  sym_defun_header = 162,
  sym_array_dimension = 163,
  sym_path_lit = 164,
  sym_package_lit = 165,
  sym__package_lit_without_slash = 166,
  sym__sym_lit_without_slash = 167,
  sym_kwd_symbol = 168,
  sym_include_reader_macro = 169,
  sym_complex_num_lit = 170,
  aux_sym_source_repeat1 = 171,
  aux_sym_dis_expr_repeat1 = 172,
  aux_sym_str_lit_repeat1 = 173,
  aux_sym_list_lit_repeat1 = 174,
  aux_sym__bare_list_lit_repeat1 = 175,
  aux_sym_read_cond_lit_repeat1 = 176,
  aux_sym_format_modifiers_repeat1 = 177,
  aux_sym_for_clause_repeat1 = 178,
  aux_sym_do_clause_repeat1 = 179,
  aux_sym_loop_macro_repeat1 = 180,
  aux_sym__sym_lit_without_slash_repeat1 = 181,
};

static const char * const ts_symbol_names[] = {
  [ts_builtin_sym_end] = "end",
  [sym__ws] = "_ws",
  [sym_comment] = "comment",
  [anon_sym_POUND_] = "#_",
  [anon_sym_POUND] = "#",
  [aux_sym__form_token1] = "_form_token1",
  [anon_sym_DOT] = ".",
  [aux_sym_num_lit_token1] = "num_lit_token1",
  [aux_sym_num_lit_token2] = "num_lit_token2",
  [anon_sym_COLON] = ":",
  [anon_sym_COLON_COLON] = "::",
  [anon_sym_DQUOTE] = "\"",
  [aux_sym_str_lit_token1] = "str_lit_token1",
  [aux_sym_str_lit_token2] = "str_lit_token2",
  [anon_sym_TILDE] = "~",
  [aux_sym_char_lit_token1] = "char_lit_token1",
  [sym_nil_lit] = "nil_lit",
  [aux_sym_sym_lit_token1] = "sym_lit_token1",
  [anon_sym_CARET] = "^",
  [anon_sym_POUND_CARET] = "#^",
  [anon_sym_LPAREN] = "(",
  [anon_sym_RPAREN] = ")",
  [anon_sym_LBRACE] = "{",
  [anon_sym_RBRACE] = "}",
  [anon_sym_POUND0A] = "#0A",
  [anon_sym_POUND0a] = "#0a",
  [anon_sym_POUND_QMARK] = "#\?",
  [anon_sym_POUND_QMARK_AT] = "#\?@",
  [anon_sym_POUND_SQUOTE] = "#'",
  [anon_sym_SQUOTE] = "'",
  [anon_sym_BQUOTE] = "`",
  [anon_sym_COMMA_AT] = ",@",
  [anon_sym_COMMA] = ",",
  [sym_block_comment] = "block_comment",
  [sym_fancy_literal] = "fancy_literal",
  [aux_sym__format_token_token1] = "char_lit",
  [anon_sym_v] = "v",
  [anon_sym_V] = "V",
  [anon_sym_AT] = "@",
  [anon_sym_AT_COLON] = "@:",
  [anon_sym_COLON_AT] = ":@",
  [anon_sym_PERCENT] = "%",
  [anon_sym_AMP] = "&",
  [anon_sym_PIPE] = "|",
  [aux_sym_format_directive_type_token1] = "format_directive_type_token1",
  [aux_sym_format_directive_type_token2] = "format_directive_type_token2",
  [anon_sym_LF] = "\n",
  [anon_sym_CR] = "\r",
  [aux_sym_format_directive_type_token3] = "format_directive_type_token3",
  [aux_sym_format_directive_type_token4] = "format_directive_type_token4",
  [aux_sym_format_directive_type_token5] = "format_directive_type_token5",
  [aux_sym_format_directive_type_token6] = "format_directive_type_token6",
  [anon_sym__] = "_",
  [aux_sym_format_directive_type_token7] = "format_directive_type_token7",
  [aux_sym_format_directive_type_token8] = "format_directive_type_token8",
  [aux_sym_format_directive_type_token9] = "format_directive_type_token9",
  [aux_sym_format_directive_type_token10] = "format_directive_type_token10",
  [anon_sym_SEMI] = ";",
  [anon_sym_STAR] = "*",
  [anon_sym_SLASH] = "/",
  [anon_sym_QMARK] = "\?",
  [anon_sym_Newline] = "Newline",
  [aux_sym_format_directive_type_token11] = "format_directive_type_token11",
  [anon_sym_cl] = "cl",
  [anon_sym_in] = "in",
  [anon_sym_across] = "across",
  [anon_sym_being] = "being",
  [anon_sym_using] = "using",
  [aux_sym_for_clause_word_token1] = "for_clause_word_token1",
  [anon_sym_below] = "below",
  [anon_sym_above] = "above",
  [anon_sym_from] = "from",
  [anon_sym_to] = "to",
  [anon_sym_upto] = "upto",
  [anon_sym_upfrom] = "upfrom",
  [anon_sym_downto] = "downto",
  [anon_sym_downfrom] = "downfrom",
  [anon_sym_on] = "on",
  [anon_sym_by] = "by",
  [anon_sym_then] = "then",
  [anon_sym_EQ] = "=",
  [aux_sym_accumulation_verb_token1] = "accumulation_verb_token1",
  [anon_sym_for] = "for",
  [anon_sym_and] = "and",
  [anon_sym_as] = "as",
  [anon_sym_with] = "with",
  [anon_sym_do] = "do",
  [anon_sym_while] = "while",
  [anon_sym_until] = "until",
  [anon_sym_repeat] = "repeat",
  [anon_sym_when] = "when",
  [anon_sym_if] = "if",
  [anon_sym_unless] = "unless",
  [anon_sym_always] = "always",
  [anon_sym_thereis] = "thereis",
  [anon_sym_never] = "never",
  [anon_sym_else] = "else",
  [anon_sym_into] = "into",
  [anon_sym_finally] = "finally",
  [anon_sym_return] = "return",
  [anon_sym_initially] = "initially",
  [anon_sym_loop] = "loop",
  [anon_sym_defun] = "defun",
  [anon_sym_defmacro] = "defmacro",
  [anon_sym_defgeneric] = "defgeneric",
  [anon_sym_defmethod] = "defmethod",
  [anon_sym_lambda] = "defun_keyword",
  [anon_sym_POUNDP] = "#P",
  [anon_sym_POUNDp] = "#p",
  [aux_sym_path_lit_token1] = "str_lit",
  [aux_sym__sym_lit_without_slash_token1] = "_sym_lit_without_slash_token1",
  [sym_self_referential_reader_macro] = "self_referential_reader_macro",
  [anon_sym_POUND_PLUS] = "#+",
  [anon_sym_POUND_DASH] = "#-",
  [anon_sym_POUNDC] = "#C",
  [anon_sym_POUNDc] = "#c",
  [sym_source] = "source",
  [sym__gap] = "_gap",
  [sym_dis_expr] = "dis_expr",
  [sym__form] = "_form",
  [sym_num_lit] = "num_lit",
  [sym_kwd_lit] = "kwd_lit",
  [sym_str_lit] = "str_lit",
  [sym_char_lit] = "char_lit",
  [sym_sym_lit] = "sym_lit",
  [sym__metadata_lit] = "_metadata_lit",
  [sym_meta_lit] = "meta_lit",
  [sym_old_meta_lit] = "old_meta_lit",
  [sym_list_lit] = "list_lit",
  [sym__bare_list_lit] = "_bare_list_lit",
  [sym_map_lit] = "map_lit",
  [sym__bare_map_lit] = "_bare_map_lit",
  [sym_vec_lit] = "vec_lit",
  [sym_set_lit] = "set_lit",
  [sym__bare_set_lit] = "_bare_set_lit",
  [sym_read_cond_lit] = "read_cond_lit",
  [sym_splicing_read_cond_lit] = "splicing_read_cond_lit",
  [sym_var_quoting_lit] = "var_quoting_lit",
  [sym_quoting_lit] = "quoting_lit",
  [sym_syn_quoting_lit] = "syn_quoting_lit",
  [sym_unquote_splicing_lit] = "unquote_splicing_lit",
  [sym_unquoting_lit] = "unquoting_lit",
  [sym_defun] = "defun",
  [sym__format_token] = "_format_token",
  [sym_format_prefix_parameters] = "format_prefix_parameters",
  [sym_format_modifiers] = "format_modifiers",
  [sym_format_directive_type] = "format_directive_type",
  [sym_format_specifier] = "format_specifier",
  [sym_for_clause_word] = "for_clause_word",
  [sym__for_part] = "_for_part",
  [sym_accumulation_verb] = "accumulation_verb",
  [sym_for_clause] = "for_clause",
  [sym_with_clause] = "with_clause",
  [sym_do_clause] = "do_clause",
  [sym_while_clause] = "while_clause",
  [sym_repeat_clause] = "repeat_clause",
  [sym_condition_clause] = "condition_clause",
  [sym_accumulation_clause] = "accumulation_clause",
  [sym_termination_clause] = "termination_clause",
  [sym_loop_clause] = "loop_clause",
  [sym_loop_macro] = "loop_macro",
  [sym_defun_keyword] = "defun_keyword",
  [sym_defun_header] = "defun_header",
  [sym_array_dimension] = "array_dimension",
  [sym_path_lit] = "path_lit",
  [sym_package_lit] = "package_lit",
  [sym__package_lit_without_slash] = "package_lit",
  [sym__sym_lit_without_slash] = "_sym_lit_without_slash",
  [sym_kwd_symbol] = "kwd_symbol",
  [sym_include_reader_macro] = "include_reader_macro",
  [sym_complex_num_lit] = "complex_num_lit",
  [aux_sym_source_repeat1] = "source_repeat1",
  [aux_sym_dis_expr_repeat1] = "dis_expr_repeat1",
  [aux_sym_str_lit_repeat1] = "str_lit_repeat1",
  [aux_sym_list_lit_repeat1] = "list_lit_repeat1",
  [aux_sym__bare_list_lit_repeat1] = "_bare_list_lit_repeat1",
  [aux_sym_read_cond_lit_repeat1] = "read_cond_lit_repeat1",
  [aux_sym_format_modifiers_repeat1] = "format_modifiers_repeat1",
  [aux_sym_for_clause_repeat1] = "for_clause_repeat1",
  [aux_sym_do_clause_repeat1] = "do_clause_repeat1",
  [aux_sym_loop_macro_repeat1] = "loop_macro_repeat1",
  [aux_sym__sym_lit_without_slash_repeat1] = "_sym_lit_without_slash_repeat1",
};

static const TSSymbol ts_symbol_map[] = {
  [ts_builtin_sym_end] = ts_builtin_sym_end,
  [sym__ws] = sym__ws,
  [sym_comment] = sym_comment,
  [anon_sym_POUND_] = anon_sym_POUND_,
  [anon_sym_POUND] = anon_sym_POUND,
  [aux_sym__form_token1] = aux_sym__form_token1,
  [anon_sym_DOT] = anon_sym_DOT,
  [aux_sym_num_lit_token1] = aux_sym_num_lit_token1,
  [aux_sym_num_lit_token2] = aux_sym_num_lit_token2,
  [anon_sym_COLON] = anon_sym_COLON,
  [anon_sym_COLON_COLON] = anon_sym_COLON_COLON,
  [anon_sym_DQUOTE] = anon_sym_DQUOTE,
  [aux_sym_str_lit_token1] = aux_sym_str_lit_token1,
  [aux_sym_str_lit_token2] = aux_sym_str_lit_token2,
  [anon_sym_TILDE] = anon_sym_TILDE,
  [aux_sym_char_lit_token1] = aux_sym_char_lit_token1,
  [sym_nil_lit] = sym_nil_lit,
  [aux_sym_sym_lit_token1] = aux_sym_sym_lit_token1,
  [anon_sym_CARET] = anon_sym_CARET,
  [anon_sym_POUND_CARET] = anon_sym_POUND_CARET,
  [anon_sym_LPAREN] = anon_sym_LPAREN,
  [anon_sym_RPAREN] = anon_sym_RPAREN,
  [anon_sym_LBRACE] = anon_sym_LBRACE,
  [anon_sym_RBRACE] = anon_sym_RBRACE,
  [anon_sym_POUND0A] = anon_sym_POUND0A,
  [anon_sym_POUND0a] = anon_sym_POUND0a,
  [anon_sym_POUND_QMARK] = anon_sym_POUND_QMARK,
  [anon_sym_POUND_QMARK_AT] = anon_sym_POUND_QMARK_AT,
  [anon_sym_POUND_SQUOTE] = anon_sym_POUND_SQUOTE,
  [anon_sym_SQUOTE] = anon_sym_SQUOTE,
  [anon_sym_BQUOTE] = anon_sym_BQUOTE,
  [anon_sym_COMMA_AT] = anon_sym_COMMA_AT,
  [anon_sym_COMMA] = anon_sym_COMMA,
  [sym_block_comment] = sym_block_comment,
  [sym_fancy_literal] = sym_fancy_literal,
  [aux_sym__format_token_token1] = sym_char_lit,
  [anon_sym_v] = anon_sym_v,
  [anon_sym_V] = anon_sym_V,
  [anon_sym_AT] = anon_sym_AT,
  [anon_sym_AT_COLON] = anon_sym_AT_COLON,
  [anon_sym_COLON_AT] = anon_sym_COLON_AT,
  [anon_sym_PERCENT] = anon_sym_PERCENT,
  [anon_sym_AMP] = anon_sym_AMP,
  [anon_sym_PIPE] = anon_sym_PIPE,
  [aux_sym_format_directive_type_token1] = aux_sym_format_directive_type_token1,
  [aux_sym_format_directive_type_token2] = aux_sym_format_directive_type_token2,
  [anon_sym_LF] = anon_sym_LF,
  [anon_sym_CR] = anon_sym_CR,
  [aux_sym_format_directive_type_token3] = aux_sym_format_directive_type_token3,
  [aux_sym_format_directive_type_token4] = aux_sym_format_directive_type_token4,
  [aux_sym_format_directive_type_token5] = aux_sym_format_directive_type_token5,
  [aux_sym_format_directive_type_token6] = aux_sym_format_directive_type_token6,
  [anon_sym__] = anon_sym__,
  [aux_sym_format_directive_type_token7] = aux_sym_format_directive_type_token7,
  [aux_sym_format_directive_type_token8] = aux_sym_format_directive_type_token8,
  [aux_sym_format_directive_type_token9] = aux_sym_format_directive_type_token9,
  [aux_sym_format_directive_type_token10] = aux_sym_format_directive_type_token10,
  [anon_sym_SEMI] = anon_sym_SEMI,
  [anon_sym_STAR] = anon_sym_STAR,
  [anon_sym_SLASH] = anon_sym_SLASH,
  [anon_sym_QMARK] = anon_sym_QMARK,
  [anon_sym_Newline] = anon_sym_Newline,
  [aux_sym_format_directive_type_token11] = aux_sym_format_directive_type_token11,
  [anon_sym_cl] = anon_sym_cl,
  [anon_sym_in] = anon_sym_in,
  [anon_sym_across] = anon_sym_across,
  [anon_sym_being] = anon_sym_being,
  [anon_sym_using] = anon_sym_using,
  [aux_sym_for_clause_word_token1] = aux_sym_for_clause_word_token1,
  [anon_sym_below] = anon_sym_below,
  [anon_sym_above] = anon_sym_above,
  [anon_sym_from] = anon_sym_from,
  [anon_sym_to] = anon_sym_to,
  [anon_sym_upto] = anon_sym_upto,
  [anon_sym_upfrom] = anon_sym_upfrom,
  [anon_sym_downto] = anon_sym_downto,
  [anon_sym_downfrom] = anon_sym_downfrom,
  [anon_sym_on] = anon_sym_on,
  [anon_sym_by] = anon_sym_by,
  [anon_sym_then] = anon_sym_then,
  [anon_sym_EQ] = anon_sym_EQ,
  [aux_sym_accumulation_verb_token1] = aux_sym_accumulation_verb_token1,
  [anon_sym_for] = anon_sym_for,
  [anon_sym_and] = anon_sym_and,
  [anon_sym_as] = anon_sym_as,
  [anon_sym_with] = anon_sym_with,
  [anon_sym_do] = anon_sym_do,
  [anon_sym_while] = anon_sym_while,
  [anon_sym_until] = anon_sym_until,
  [anon_sym_repeat] = anon_sym_repeat,
  [anon_sym_when] = anon_sym_when,
  [anon_sym_if] = anon_sym_if,
  [anon_sym_unless] = anon_sym_unless,
  [anon_sym_always] = anon_sym_always,
  [anon_sym_thereis] = anon_sym_thereis,
  [anon_sym_never] = anon_sym_never,
  [anon_sym_else] = anon_sym_else,
  [anon_sym_into] = anon_sym_into,
  [anon_sym_finally] = anon_sym_finally,
  [anon_sym_return] = anon_sym_return,
  [anon_sym_initially] = anon_sym_initially,
  [anon_sym_loop] = anon_sym_loop,
  [anon_sym_defun] = anon_sym_defun,
  [anon_sym_defmacro] = anon_sym_defmacro,
  [anon_sym_defgeneric] = anon_sym_defgeneric,
  [anon_sym_defmethod] = anon_sym_defmethod,
  [anon_sym_lambda] = sym_defun_keyword,
  [anon_sym_POUNDP] = anon_sym_POUNDP,
  [anon_sym_POUNDp] = anon_sym_POUNDp,
  [aux_sym_path_lit_token1] = sym_str_lit,
  [aux_sym__sym_lit_without_slash_token1] = aux_sym__sym_lit_without_slash_token1,
  [sym_self_referential_reader_macro] = sym_self_referential_reader_macro,
  [anon_sym_POUND_PLUS] = anon_sym_POUND_PLUS,
  [anon_sym_POUND_DASH] = anon_sym_POUND_DASH,
  [anon_sym_POUNDC] = anon_sym_POUNDC,
  [anon_sym_POUNDc] = anon_sym_POUNDc,
  [sym_source] = sym_source,
  [sym__gap] = sym__gap,
  [sym_dis_expr] = sym_dis_expr,
  [sym__form] = sym__form,
  [sym_num_lit] = sym_num_lit,
  [sym_kwd_lit] = sym_kwd_lit,
  [sym_str_lit] = sym_str_lit,
  [sym_char_lit] = sym_char_lit,
  [sym_sym_lit] = sym_sym_lit,
  [sym__metadata_lit] = sym__metadata_lit,
  [sym_meta_lit] = sym_meta_lit,
  [sym_old_meta_lit] = sym_old_meta_lit,
  [sym_list_lit] = sym_list_lit,
  [sym__bare_list_lit] = sym__bare_list_lit,
  [sym_map_lit] = sym_map_lit,
  [sym__bare_map_lit] = sym__bare_map_lit,
  [sym_vec_lit] = sym_vec_lit,
  [sym_set_lit] = sym_set_lit,
  [sym__bare_set_lit] = sym__bare_set_lit,
  [sym_read_cond_lit] = sym_read_cond_lit,
  [sym_splicing_read_cond_lit] = sym_splicing_read_cond_lit,
  [sym_var_quoting_lit] = sym_var_quoting_lit,
  [sym_quoting_lit] = sym_quoting_lit,
  [sym_syn_quoting_lit] = sym_syn_quoting_lit,
  [sym_unquote_splicing_lit] = sym_unquote_splicing_lit,
  [sym_unquoting_lit] = sym_unquoting_lit,
  [sym_defun] = sym_defun,
  [sym__format_token] = sym__format_token,
  [sym_format_prefix_parameters] = sym_format_prefix_parameters,
  [sym_format_modifiers] = sym_format_modifiers,
  [sym_format_directive_type] = sym_format_directive_type,
  [sym_format_specifier] = sym_format_specifier,
  [sym_for_clause_word] = sym_for_clause_word,
  [sym__for_part] = sym__for_part,
  [sym_accumulation_verb] = sym_accumulation_verb,
  [sym_for_clause] = sym_for_clause,
  [sym_with_clause] = sym_with_clause,
  [sym_do_clause] = sym_do_clause,
  [sym_while_clause] = sym_while_clause,
  [sym_repeat_clause] = sym_repeat_clause,
  [sym_condition_clause] = sym_condition_clause,
  [sym_accumulation_clause] = sym_accumulation_clause,
  [sym_termination_clause] = sym_termination_clause,
  [sym_loop_clause] = sym_loop_clause,
  [sym_loop_macro] = sym_loop_macro,
  [sym_defun_keyword] = sym_defun_keyword,
  [sym_defun_header] = sym_defun_header,
  [sym_array_dimension] = sym_array_dimension,
  [sym_path_lit] = sym_path_lit,
  [sym_package_lit] = sym_package_lit,
  [sym__package_lit_without_slash] = sym_package_lit,
  [sym__sym_lit_without_slash] = sym__sym_lit_without_slash,
  [sym_kwd_symbol] = sym_kwd_symbol,
  [sym_include_reader_macro] = sym_include_reader_macro,
  [sym_complex_num_lit] = sym_complex_num_lit,
  [aux_sym_source_repeat1] = aux_sym_source_repeat1,
  [aux_sym_dis_expr_repeat1] = aux_sym_dis_expr_repeat1,
  [aux_sym_str_lit_repeat1] = aux_sym_str_lit_repeat1,
  [aux_sym_list_lit_repeat1] = aux_sym_list_lit_repeat1,
  [aux_sym__bare_list_lit_repeat1] = aux_sym__bare_list_lit_repeat1,
  [aux_sym_read_cond_lit_repeat1] = aux_sym_read_cond_lit_repeat1,
  [aux_sym_format_modifiers_repeat1] = aux_sym_format_modifiers_repeat1,
  [aux_sym_for_clause_repeat1] = aux_sym_for_clause_repeat1,
  [aux_sym_do_clause_repeat1] = aux_sym_do_clause_repeat1,
  [aux_sym_loop_macro_repeat1] = aux_sym_loop_macro_repeat1,
  [aux_sym__sym_lit_without_slash_repeat1] = aux_sym__sym_lit_without_slash_repeat1,
};

static const TSSymbolMetadata ts_symbol_metadata[] = {
  [ts_builtin_sym_end] = {
    .visible = false,
    .named = true,
  },
  [sym__ws] = {
    .visible = false,
    .named = true,
  },
  [sym_comment] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_POUND_] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_POUND] = {
    .visible = true,
    .named = false,
  },
  [aux_sym__form_token1] = {
    .visible = false,
    .named = false,
  },
  [anon_sym_DOT] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_num_lit_token1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_num_lit_token2] = {
    .visible = false,
    .named = false,
  },
  [anon_sym_COLON] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_COLON_COLON] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DQUOTE] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_str_lit_token1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_str_lit_token2] = {
    .visible = false,
    .named = false,
  },
  [anon_sym_TILDE] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_char_lit_token1] = {
    .visible = false,
    .named = false,
  },
  [sym_nil_lit] = {
    .visible = true,
    .named = true,
  },
  [aux_sym_sym_lit_token1] = {
    .visible = false,
    .named = false,
  },
  [anon_sym_CARET] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_POUND_CARET] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LPAREN] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_RPAREN] = {
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
  [anon_sym_POUND0A] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_POUND0a] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_POUND_QMARK] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_POUND_QMARK_AT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_POUND_SQUOTE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_SQUOTE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_BQUOTE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_COMMA_AT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_COMMA] = {
    .visible = true,
    .named = false,
  },
  [sym_block_comment] = {
    .visible = true,
    .named = true,
  },
  [sym_fancy_literal] = {
    .visible = true,
    .named = true,
  },
  [aux_sym__format_token_token1] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_v] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_V] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_AT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_AT_COLON] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_COLON_AT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PERCENT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_AMP] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PIPE] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_format_directive_type_token1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_format_directive_type_token2] = {
    .visible = false,
    .named = false,
  },
  [anon_sym_LF] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_CR] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_format_directive_type_token3] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_format_directive_type_token4] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_format_directive_type_token5] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_format_directive_type_token6] = {
    .visible = false,
    .named = false,
  },
  [anon_sym__] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_format_directive_type_token7] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_format_directive_type_token8] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_format_directive_type_token9] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_format_directive_type_token10] = {
    .visible = false,
    .named = false,
  },
  [anon_sym_SEMI] = {
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
  [anon_sym_QMARK] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_Newline] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_format_directive_type_token11] = {
    .visible = false,
    .named = false,
  },
  [anon_sym_cl] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_in] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_across] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_being] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_using] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_for_clause_word_token1] = {
    .visible = false,
    .named = false,
  },
  [anon_sym_below] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_above] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_from] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_to] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_upto] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_upfrom] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_downto] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_downfrom] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_on] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_by] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_then] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_EQ] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_accumulation_verb_token1] = {
    .visible = false,
    .named = false,
  },
  [anon_sym_for] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_and] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_as] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_with] = {
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
  [anon_sym_until] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_repeat] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_when] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_if] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_unless] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_always] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_thereis] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_never] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_else] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_into] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_finally] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_return] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_initially] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_loop] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_defun] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_defmacro] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_defgeneric] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_defmethod] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_lambda] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_POUNDP] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_POUNDp] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_path_lit_token1] = {
    .visible = true,
    .named = true,
  },
  [aux_sym__sym_lit_without_slash_token1] = {
    .visible = false,
    .named = false,
  },
  [sym_self_referential_reader_macro] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_POUND_PLUS] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_POUND_DASH] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_POUNDC] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_POUNDc] = {
    .visible = true,
    .named = false,
  },
  [sym_source] = {
    .visible = true,
    .named = true,
  },
  [sym__gap] = {
    .visible = false,
    .named = true,
  },
  [sym_dis_expr] = {
    .visible = true,
    .named = true,
  },
  [sym__form] = {
    .visible = false,
    .named = true,
  },
  [sym_num_lit] = {
    .visible = true,
    .named = true,
  },
  [sym_kwd_lit] = {
    .visible = true,
    .named = true,
  },
  [sym_str_lit] = {
    .visible = true,
    .named = true,
  },
  [sym_char_lit] = {
    .visible = true,
    .named = true,
  },
  [sym_sym_lit] = {
    .visible = true,
    .named = true,
  },
  [sym__metadata_lit] = {
    .visible = false,
    .named = true,
  },
  [sym_meta_lit] = {
    .visible = true,
    .named = true,
  },
  [sym_old_meta_lit] = {
    .visible = true,
    .named = true,
  },
  [sym_list_lit] = {
    .visible = true,
    .named = true,
  },
  [sym__bare_list_lit] = {
    .visible = false,
    .named = true,
  },
  [sym_map_lit] = {
    .visible = true,
    .named = true,
  },
  [sym__bare_map_lit] = {
    .visible = false,
    .named = true,
  },
  [sym_vec_lit] = {
    .visible = true,
    .named = true,
  },
  [sym_set_lit] = {
    .visible = true,
    .named = true,
  },
  [sym__bare_set_lit] = {
    .visible = false,
    .named = true,
  },
  [sym_read_cond_lit] = {
    .visible = true,
    .named = true,
  },
  [sym_splicing_read_cond_lit] = {
    .visible = true,
    .named = true,
  },
  [sym_var_quoting_lit] = {
    .visible = true,
    .named = true,
  },
  [sym_quoting_lit] = {
    .visible = true,
    .named = true,
  },
  [sym_syn_quoting_lit] = {
    .visible = true,
    .named = true,
  },
  [sym_unquote_splicing_lit] = {
    .visible = true,
    .named = true,
  },
  [sym_unquoting_lit] = {
    .visible = true,
    .named = true,
  },
  [sym_defun] = {
    .visible = true,
    .named = true,
  },
  [sym__format_token] = {
    .visible = false,
    .named = true,
  },
  [sym_format_prefix_parameters] = {
    .visible = true,
    .named = true,
  },
  [sym_format_modifiers] = {
    .visible = true,
    .named = true,
  },
  [sym_format_directive_type] = {
    .visible = true,
    .named = true,
  },
  [sym_format_specifier] = {
    .visible = true,
    .named = true,
  },
  [sym_for_clause_word] = {
    .visible = true,
    .named = true,
  },
  [sym__for_part] = {
    .visible = false,
    .named = true,
  },
  [sym_accumulation_verb] = {
    .visible = true,
    .named = true,
  },
  [sym_for_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_with_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_do_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_while_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_repeat_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_condition_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_accumulation_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_termination_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_loop_clause] = {
    .visible = true,
    .named = true,
  },
  [sym_loop_macro] = {
    .visible = true,
    .named = true,
  },
  [sym_defun_keyword] = {
    .visible = true,
    .named = true,
  },
  [sym_defun_header] = {
    .visible = true,
    .named = true,
  },
  [sym_array_dimension] = {
    .visible = true,
    .named = true,
  },
  [sym_path_lit] = {
    .visible = true,
    .named = true,
  },
  [sym_package_lit] = {
    .visible = true,
    .named = true,
  },
  [sym__package_lit_without_slash] = {
    .visible = true,
    .named = true,
  },
  [sym__sym_lit_without_slash] = {
    .visible = false,
    .named = true,
  },
  [sym_kwd_symbol] = {
    .visible = true,
    .named = true,
  },
  [sym_include_reader_macro] = {
    .visible = true,
    .named = true,
  },
  [sym_complex_num_lit] = {
    .visible = true,
    .named = true,
  },
  [aux_sym_source_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_dis_expr_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_str_lit_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_list_lit_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__bare_list_lit_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_read_cond_lit_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_format_modifiers_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_for_clause_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_do_clause_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_loop_macro_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__sym_lit_without_slash_repeat1] = {
    .visible = false,
    .named = false,
  },
};

enum ts_field_identifiers {
  field_close = 1,
  field_condition = 2,
  field_function_name = 3,
  field_imaginary = 4,
  field_keyword = 5,
  field_lambda_list = 6,
  field_marker = 7,
  field_meta = 8,
  field_numberOfArgs = 9,
  field_old_meta = 10,
  field_open = 11,
  field_package = 12,
  field_real = 13,
  field_repetitions = 14,
  field_specifier = 15,
  field_symbol = 16,
  field_target = 17,
  field_type = 18,
  field_value = 19,
  field_variable = 20,
};

static const char * const ts_field_names[] = {
  [0] = NULL,
  [field_close] = "close",
  [field_condition] = "condition",
  [field_function_name] = "function_name",
  [field_imaginary] = "imaginary",
  [field_keyword] = "keyword",
  [field_lambda_list] = "lambda_list",
  [field_marker] = "marker",
  [field_meta] = "meta",
  [field_numberOfArgs] = "numberOfArgs",
  [field_old_meta] = "old_meta",
  [field_open] = "open",
  [field_package] = "package",
  [field_real] = "real",
  [field_repetitions] = "repetitions",
  [field_specifier] = "specifier",
  [field_symbol] = "symbol",
  [field_target] = "target",
  [field_type] = "type",
  [field_value] = "value",
  [field_variable] = "variable",
};

static const TSFieldMapSlice ts_field_map_slices[PRODUCTION_ID_COUNT] = {
  [1] = {.index = 0, .length = 1},
  [2] = {.index = 1, .length = 2},
  [3] = {.index = 3, .length = 1},
  [4] = {.index = 4, .length = 1},
  [5] = {.index = 5, .length = 3},
  [6] = {.index = 8, .length = 4},
  [7] = {.index = 12, .length = 3},
  [8] = {.index = 15, .length = 1},
  [9] = {.index = 15, .length = 1},
  [10] = {.index = 16, .length = 2},
  [11] = {.index = 18, .length = 2},
  [12] = {.index = 20, .length = 2},
  [13] = {.index = 22, .length = 4},
  [14] = {.index = 26, .length = 2},
  [15] = {.index = 28, .length = 5},
  [16] = {.index = 33, .length = 6},
  [17] = {.index = 39, .length = 4},
  [18] = {.index = 43, .length = 3},
  [19] = {.index = 46, .length = 1},
  [20] = {.index = 15, .length = 1},
  [21] = {.index = 47, .length = 3},
  [23] = {.index = 50, .length = 2},
  [24] = {.index = 52, .length = 2},
  [25] = {.index = 54, .length = 2},
  [26] = {.index = 56, .length = 1},
  [27] = {.index = 57, .length = 4},
  [28] = {.index = 61, .length = 4},
  [29] = {.index = 65, .length = 4},
  [30] = {.index = 69, .length = 2},
  [31] = {.index = 71, .length = 5},
  [32] = {.index = 76, .length = 6},
  [33] = {.index = 82, .length = 5},
  [34] = {.index = 46, .length = 1},
  [35] = {.index = 87, .length = 5},
  [37] = {.index = 92, .length = 1},
  [38] = {.index = 93, .length = 1},
  [39] = {.index = 94, .length = 1},
  [40] = {.index = 95, .length = 2},
  [41] = {.index = 97, .length = 2},
  [42] = {.index = 99, .length = 4},
  [43] = {.index = 103, .length = 4},
  [44] = {.index = 107, .length = 5},
  [45] = {.index = 112, .length = 5},
  [46] = {.index = 117, .length = 6},
  [47] = {.index = 123, .length = 5},
  [48] = {.index = 128, .length = 7},
  [49] = {.index = 135, .length = 2},
  [50] = {.index = 137, .length = 2},
  [51] = {.index = 139, .length = 1},
  [52] = {.index = 140, .length = 3},
  [53] = {.index = 143, .length = 2},
  [54] = {.index = 145, .length = 3},
  [55] = {.index = 148, .length = 4},
  [56] = {.index = 152, .length = 5},
  [57] = {.index = 157, .length = 4},
  [58] = {.index = 161, .length = 4},
  [59] = {.index = 165, .length = 5},
  [60] = {.index = 170, .length = 3},
  [61] = {.index = 173, .length = 7},
  [62] = {.index = 180, .length = 7},
  [63] = {.index = 187, .length = 2},
  [64] = {.index = 189, .length = 3},
  [65] = {.index = 192, .length = 2},
  [66] = {.index = 194, .length = 3},
  [67] = {.index = 197, .length = 1},
  [68] = {.index = 198, .length = 5},
  [69] = {.index = 203, .length = 2},
  [70] = {.index = 205, .length = 4},
  [71] = {.index = 209, .length = 3},
  [72] = {.index = 212, .length = 4},
  [73] = {.index = 216, .length = 5},
  [74] = {.index = 221, .length = 6},
  [75] = {.index = 227, .length = 5},
  [76] = {.index = 232, .length = 4},
  [77] = {.index = 236, .length = 3},
  [78] = {.index = 239, .length = 3},
  [79] = {.index = 242, .length = 7},
  [80] = {.index = 249, .length = 5},
  [81] = {.index = 254, .length = 5},
  [82] = {.index = 259, .length = 3},
  [83] = {.index = 262, .length = 2},
  [84] = {.index = 264, .length = 4},
  [85] = {.index = 268, .length = 3},
  [86] = {.index = 271, .length = 2},
  [87] = {.index = 273, .length = 3},
  [88] = {.index = 276, .length = 1},
  [89] = {.index = 277, .length = 6},
  [90] = {.index = 283, .length = 2},
  [91] = {.index = 285, .length = 4},
  [92] = {.index = 289, .length = 4},
  [93] = {.index = 293, .length = 2},
  [94] = {.index = 295, .length = 6},
  [95] = {.index = 301, .length = 5},
  [96] = {.index = 306, .length = 6},
  [97] = {.index = 312, .length = 3},
  [98] = {.index = 315, .length = 3},
  [99] = {.index = 318, .length = 5},
  [100] = {.index = 323, .length = 5},
  [101] = {.index = 328, .length = 6},
  [102] = {.index = 334, .length = 5},
  [103] = {.index = 339, .length = 2},
  [104] = {.index = 341, .length = 4},
  [105] = {.index = 345, .length = 4},
  [106] = {.index = 349, .length = 2},
  [107] = {.index = 351, .length = 4},
  [108] = {.index = 355, .length = 3},
  [109] = {.index = 358, .length = 3},
  [110] = {.index = 361, .length = 2},
  [111] = {.index = 363, .length = 3},
  [112] = {.index = 366, .length = 2},
  [113] = {.index = 368, .length = 4},
  [114] = {.index = 372, .length = 4},
  [115] = {.index = 376, .length = 6},
  [116] = {.index = 382, .length = 3},
  [117] = {.index = 385, .length = 5},
  [118] = {.index = 390, .length = 5},
  [119] = {.index = 395, .length = 6},
  [120] = {.index = 401, .length = 2},
  [121] = {.index = 403, .length = 4},
  [122] = {.index = 407, .length = 4},
  [123] = {.index = 411, .length = 2},
  [124] = {.index = 413, .length = 4},
  [125] = {.index = 417, .length = 4},
  [126] = {.index = 421, .length = 5},
  [127] = {.index = 426, .length = 2},
  [128] = {.index = 428, .length = 4},
  [129] = {.index = 432, .length = 3},
  [130] = {.index = 435, .length = 2},
  [131] = {.index = 437, .length = 4},
  [132] = {.index = 441, .length = 4},
  [133] = {.index = 445, .length = 5},
  [134] = {.index = 450, .length = 2},
  [135] = {.index = 452, .length = 4},
  [136] = {.index = 456, .length = 4},
  [137] = {.index = 460, .length = 2},
  [138] = {.index = 462, .length = 4},
  [139] = {.index = 466, .length = 4},
  [140] = {.index = 470, .length = 6},
  [141] = {.index = 476, .length = 2},
  [142] = {.index = 478, .length = 4},
  [143] = {.index = 482, .length = 4},
  [144] = {.index = 486, .length = 4},
  [145] = {.index = 490, .length = 2},
  [146] = {.index = 492, .length = 4},
  [147] = {.index = 496, .length = 4},
  [148] = {.index = 500, .length = 2},
  [149] = {.index = 502, .length = 4},
  [150] = {.index = 506, .length = 2},
  [151] = {.index = 508, .length = 4},
  [152] = {.index = 512, .length = 4},
  [153] = {.index = 516, .length = 2},
  [154] = {.index = 518, .length = 4},
  [155] = {.index = 522, .length = 4},
  [156] = {.index = 526, .length = 4},
  [157] = {.index = 530, .length = 4},
  [158] = {.index = 534, .length = 4},
  [159] = {.index = 538, .length = 4},
  [160] = {.index = 542, .length = 4},
  [161] = {.index = 546, .length = 2},
  [162] = {.index = 548, .length = 4},
  [163] = {.index = 552, .length = 2},
  [164] = {.index = 554, .length = 4},
  [165] = {.index = 558, .length = 4},
  [166] = {.index = 562, .length = 4},
  [167] = {.index = 566, .length = 4},
  [168] = {.index = 570, .length = 4},
  [169] = {.index = 574, .length = 4},
  [170] = {.index = 578, .length = 4},
  [171] = {.index = 582, .length = 2},
  [172] = {.index = 584, .length = 4},
  [173] = {.index = 588, .length = 4},
  [174] = {.index = 592, .length = 4},
  [175] = {.index = 596, .length = 4},
  [176] = {.index = 600, .length = 4},
};

static const TSFieldMapEntry ts_field_map_entries[] = {
  [0] =
    {field_open, 0, .inherited = true},
  [1] =
    {field_meta, 0, .inherited = true},
    {field_old_meta, 0, .inherited = true},
  [3] =
    {field_meta, 0},
  [4] =
    {field_old_meta, 0},
  [5] =
    {field_close, 0, .inherited = true},
    {field_open, 0, .inherited = true},
    {field_value, 0, .inherited = true},
  [8] =
    {field_close, 0, .inherited = true},
    {field_marker, 0, .inherited = true},
    {field_open, 0, .inherited = true},
    {field_value, 0, .inherited = true},
  [12] =
    {field_marker, 0},
    {field_open, 1, .inherited = true},
    {field_value, 1},
  [15] =
    {field_open, 0},
  [16] =
    {field_marker, 0},
    {field_value, 1},
  [18] =
    {field_close, 1},
    {field_open, 0},
  [20] =
    {field_open, 0, .inherited = true},
    {field_value, 0},
  [22] =
    {field_close, 1, .inherited = true},
    {field_marker, 0},
    {field_open, 1, .inherited = true},
    {field_value, 1, .inherited = true},
  [26] =
    {field_open, 0, .inherited = true},
    {field_open, 1, .inherited = true},
  [28] =
    {field_close, 1, .inherited = true},
    {field_meta, 0, .inherited = true},
    {field_old_meta, 0, .inherited = true},
    {field_open, 1, .inherited = true},
    {field_value, 1, .inherited = true},
  [33] =
    {field_close, 1, .inherited = true},
    {field_marker, 1, .inherited = true},
    {field_meta, 0, .inherited = true},
    {field_old_meta, 0, .inherited = true},
    {field_open, 1, .inherited = true},
    {field_value, 1, .inherited = true},
  [39] =
    {field_meta, 0, .inherited = true},
    {field_meta, 1, .inherited = true},
    {field_old_meta, 0, .inherited = true},
    {field_old_meta, 1, .inherited = true},
  [43] =
    {field_marker, 0},
    {field_open, 2, .inherited = true},
    {field_value, 2},
  [46] =
    {field_open, 1},
  [47] =
    {field_close, 2},
    {field_marker, 0},
    {field_open, 1},
  [50] =
    {field_marker, 0},
    {field_value, 2},
  [52] =
    {field_close, 2},
    {field_open, 0},
  [54] =
    {field_keyword, 0},
    {field_lambda_list, 1},
  [56] =
    {field_keyword, 0},
  [57] =
    {field_close, 2},
    {field_open, 0},
    {field_open, 1, .inherited = true},
    {field_value, 1, .inherited = true},
  [61] =
    {field_open, 0, .inherited = true},
    {field_open, 1, .inherited = true},
    {field_value, 0, .inherited = true},
    {field_value, 1, .inherited = true},
  [65] =
    {field_close, 2, .inherited = true},
    {field_marker, 0},
    {field_open, 2, .inherited = true},
    {field_value, 2, .inherited = true},
  [69] =
    {field_package, 0},
    {field_symbol, 2},
  [71] =
    {field_condition, 1},
    {field_marker, 0},
    {field_open, 1, .inherited = true},
    {field_open, 2, .inherited = true},
    {field_target, 2},
  [76] =
    {field_close, 2, .inherited = true},
    {field_marker, 1},
    {field_meta, 0, .inherited = true},
    {field_old_meta, 0, .inherited = true},
    {field_open, 2, .inherited = true},
    {field_value, 2, .inherited = true},
  [82] =
    {field_marker, 1},
    {field_meta, 0, .inherited = true},
    {field_old_meta, 0, .inherited = true},
    {field_open, 2, .inherited = true},
    {field_value, 2},
  [87] =
    {field_close, 3},
    {field_marker, 0},
    {field_open, 1},
    {field_open, 2, .inherited = true},
    {field_value, 2, .inherited = true},
  [92] =
    {field_repetitions, 0},
  [93] =
    {field_numberOfArgs, 0},
  [94] =
    {field_open, 1, .inherited = true},
  [95] =
    {field_close, 3},
    {field_open, 0},
  [97] =
    {field_keyword, 0},
    {field_lambda_list, 2},
  [99] =
    {field_function_name, 1},
    {field_keyword, 0},
    {field_lambda_list, 2},
    {field_open, 1, .inherited = true},
  [103] =
    {field_close, 3},
    {field_open, 0},
    {field_open, 2, .inherited = true},
    {field_value, 2, .inherited = true},
  [107] =
    {field_condition, 1},
    {field_marker, 0},
    {field_open, 1, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_target, 3},
  [112] =
    {field_condition, 2},
    {field_marker, 0},
    {field_open, 2, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_target, 3},
  [117] =
    {field_close, 3, .inherited = true},
    {field_marker, 1},
    {field_meta, 0, .inherited = true},
    {field_old_meta, 0, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_value, 3, .inherited = true},
  [123] =
    {field_marker, 1},
    {field_meta, 0, .inherited = true},
    {field_old_meta, 0, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_value, 3},
  [128] =
    {field_condition, 2},
    {field_marker, 1},
    {field_meta, 0, .inherited = true},
    {field_old_meta, 0, .inherited = true},
    {field_open, 2, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_target, 3},
  [135] =
    {field_package, 1, .inherited = true},
    {field_symbol, 1, .inherited = true},
  [137] =
    {field_close, 4},
    {field_open, 0},
  [139] =
    {field_open, 2, .inherited = true},
  [140] =
    {field_open, 1, .inherited = true},
    {field_open, 2, .inherited = true},
    {field_variable, 1},
  [143] =
    {field_open, 1, .inherited = true},
    {field_open, 2, .inherited = true},
  [145] =
    {field_open, 1, .inherited = true},
    {field_open, 2, .inherited = true},
    {field_type, 2},
  [148] =
    {field_close, 4},
    {field_open, 0},
    {field_open, 3, .inherited = true},
    {field_value, 3, .inherited = true},
  [152] =
    {field_function_name, 1},
    {field_keyword, 0},
    {field_lambda_list, 3},
    {field_open, 1, .inherited = true},
    {field_specifier, 2},
  [157] =
    {field_function_name, 1},
    {field_keyword, 0},
    {field_lambda_list, 3},
    {field_open, 1, .inherited = true},
  [161] =
    {field_function_name, 2},
    {field_keyword, 0},
    {field_lambda_list, 3},
    {field_open, 2, .inherited = true},
  [165] =
    {field_condition, 2},
    {field_marker, 0},
    {field_open, 2, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_target, 4},
  [170] =
    {field_imaginary, 3},
    {field_marker, 0},
    {field_real, 2},
  [173] =
    {field_condition, 2},
    {field_marker, 1},
    {field_meta, 0, .inherited = true},
    {field_old_meta, 0, .inherited = true},
    {field_open, 2, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_target, 4},
  [180] =
    {field_condition, 3},
    {field_marker, 1},
    {field_meta, 0, .inherited = true},
    {field_old_meta, 0, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_target, 4},
  [187] =
    {field_close, 5},
    {field_open, 0},
  [189] =
    {field_open, 2, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_variable, 2},
  [192] =
    {field_open, 2, .inherited = true},
    {field_open, 3, .inherited = true},
  [194] =
    {field_open, 2, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_type, 3},
  [197] =
    {field_open, 3, .inherited = true},
  [198] =
    {field_open, 1, .inherited = true},
    {field_open, 2, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_type, 2},
    {field_variable, 1},
  [203] =
    {field_open, 1, .inherited = true},
    {field_open, 3, .inherited = true},
  [205] =
    {field_open, 1, .inherited = true},
    {field_open, 2, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_type, 2},
  [209] =
    {field_open, 1, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_type, 3},
  [212] =
    {field_close, 5},
    {field_open, 0},
    {field_open, 4, .inherited = true},
    {field_value, 4, .inherited = true},
  [216] =
    {field_function_name, 1},
    {field_keyword, 0},
    {field_lambda_list, 4},
    {field_open, 1, .inherited = true},
    {field_specifier, 2},
  [221] =
    {field_function_name, 1},
    {field_keyword, 0},
    {field_lambda_list, 4},
    {field_open, 1, .inherited = true},
    {field_specifier, 2},
    {field_specifier, 3},
  [227] =
    {field_function_name, 2},
    {field_keyword, 0},
    {field_lambda_list, 4},
    {field_open, 2, .inherited = true},
    {field_specifier, 3},
  [232] =
    {field_function_name, 2},
    {field_keyword, 0},
    {field_lambda_list, 4},
    {field_open, 2, .inherited = true},
  [236] =
    {field_imaginary, 4},
    {field_marker, 0},
    {field_real, 2},
  [239] =
    {field_imaginary, 4},
    {field_marker, 0},
    {field_real, 3},
  [242] =
    {field_condition, 3},
    {field_marker, 1},
    {field_meta, 0, .inherited = true},
    {field_old_meta, 0, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_target, 5},
  [249] =
    {field_imaginary, 4},
    {field_marker, 1},
    {field_meta, 0, .inherited = true},
    {field_old_meta, 0, .inherited = true},
    {field_real, 3},
  [254] =
    {field_open, 2, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_type, 3},
    {field_variable, 2},
  [259] =
    {field_open, 3, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_variable, 3},
  [262] =
    {field_open, 2, .inherited = true},
    {field_open, 4, .inherited = true},
  [264] =
    {field_open, 2, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_type, 3},
  [268] =
    {field_open, 2, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_type, 4},
  [271] =
    {field_open, 3, .inherited = true},
    {field_open, 4, .inherited = true},
  [273] =
    {field_open, 3, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_type, 4},
  [276] =
    {field_open, 4, .inherited = true},
  [277] =
    {field_open, 1, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_type, 2},
    {field_type, 3},
    {field_variable, 1},
  [283] =
    {field_open, 1, .inherited = true},
    {field_open, 4, .inherited = true},
  [285] =
    {field_open, 1, .inherited = true},
    {field_open, 2, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_type, 2},
  [289] =
    {field_open, 1, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_type, 3},
  [293] =
    {field_close, 6},
    {field_open, 0},
  [295] =
    {field_function_name, 1},
    {field_keyword, 0},
    {field_lambda_list, 5},
    {field_open, 1, .inherited = true},
    {field_specifier, 2},
    {field_specifier, 3},
  [301] =
    {field_function_name, 2},
    {field_keyword, 0},
    {field_lambda_list, 5},
    {field_open, 2, .inherited = true},
    {field_specifier, 3},
  [306] =
    {field_function_name, 2},
    {field_keyword, 0},
    {field_lambda_list, 5},
    {field_open, 2, .inherited = true},
    {field_specifier, 3},
    {field_specifier, 4},
  [312] =
    {field_imaginary, 5},
    {field_marker, 0},
    {field_real, 3},
  [315] =
    {field_imaginary, 5},
    {field_marker, 0},
    {field_real, 4},
  [318] =
    {field_imaginary, 5},
    {field_marker, 1},
    {field_meta, 0, .inherited = true},
    {field_old_meta, 0, .inherited = true},
    {field_real, 3},
  [323] =
    {field_imaginary, 5},
    {field_marker, 1},
    {field_meta, 0, .inherited = true},
    {field_old_meta, 0, .inherited = true},
    {field_real, 4},
  [328] =
    {field_open, 2, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_type, 3},
    {field_type, 4},
    {field_variable, 2},
  [334] =
    {field_open, 3, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_type, 4},
    {field_variable, 3},
  [339] =
    {field_open, 2, .inherited = true},
    {field_open, 5, .inherited = true},
  [341] =
    {field_open, 2, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_type, 3},
  [345] =
    {field_open, 2, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_type, 4},
  [349] =
    {field_open, 3, .inherited = true},
    {field_open, 5, .inherited = true},
  [351] =
    {field_open, 3, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_type, 4},
  [355] =
    {field_open, 3, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_type, 5},
  [358] =
    {field_open, 4, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_variable, 4},
  [361] =
    {field_open, 4, .inherited = true},
    {field_open, 5, .inherited = true},
  [363] =
    {field_open, 4, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_type, 5},
  [366] =
    {field_open, 1, .inherited = true},
    {field_open, 5, .inherited = true},
  [368] =
    {field_open, 1, .inherited = true},
    {field_open, 2, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_type, 2},
  [372] =
    {field_open, 1, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_type, 3},
  [376] =
    {field_function_name, 2},
    {field_keyword, 0},
    {field_lambda_list, 6},
    {field_open, 2, .inherited = true},
    {field_specifier, 3},
    {field_specifier, 4},
  [382] =
    {field_imaginary, 6},
    {field_marker, 0},
    {field_real, 4},
  [385] =
    {field_imaginary, 6},
    {field_marker, 1},
    {field_meta, 0, .inherited = true},
    {field_old_meta, 0, .inherited = true},
    {field_real, 4},
  [390] =
    {field_imaginary, 6},
    {field_marker, 1},
    {field_meta, 0, .inherited = true},
    {field_old_meta, 0, .inherited = true},
    {field_real, 5},
  [395] =
    {field_open, 3, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_open, 6, .inherited = true},
    {field_type, 4},
    {field_type, 5},
    {field_variable, 3},
  [401] =
    {field_open, 2, .inherited = true},
    {field_open, 6, .inherited = true},
  [403] =
    {field_open, 2, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_open, 6, .inherited = true},
    {field_type, 3},
  [407] =
    {field_open, 2, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_open, 6, .inherited = true},
    {field_type, 4},
  [411] =
    {field_open, 3, .inherited = true},
    {field_open, 6, .inherited = true},
  [413] =
    {field_open, 3, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_open, 6, .inherited = true},
    {field_type, 4},
  [417] =
    {field_open, 3, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_open, 6, .inherited = true},
    {field_type, 5},
  [421] =
    {field_open, 4, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_open, 6, .inherited = true},
    {field_type, 5},
    {field_variable, 4},
  [426] =
    {field_open, 4, .inherited = true},
    {field_open, 6, .inherited = true},
  [428] =
    {field_open, 4, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_open, 6, .inherited = true},
    {field_type, 5},
  [432] =
    {field_open, 4, .inherited = true},
    {field_open, 6, .inherited = true},
    {field_type, 6},
  [435] =
    {field_open, 1, .inherited = true},
    {field_open, 6, .inherited = true},
  [437] =
    {field_open, 1, .inherited = true},
    {field_open, 2, .inherited = true},
    {field_open, 6, .inherited = true},
    {field_type, 2},
  [441] =
    {field_open, 1, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_open, 6, .inherited = true},
    {field_type, 3},
  [445] =
    {field_imaginary, 7},
    {field_marker, 1},
    {field_meta, 0, .inherited = true},
    {field_old_meta, 0, .inherited = true},
    {field_real, 5},
  [450] =
    {field_open, 2, .inherited = true},
    {field_open, 7, .inherited = true},
  [452] =
    {field_open, 2, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_open, 7, .inherited = true},
    {field_type, 3},
  [456] =
    {field_open, 2, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_open, 7, .inherited = true},
    {field_type, 4},
  [460] =
    {field_open, 3, .inherited = true},
    {field_open, 7, .inherited = true},
  [462] =
    {field_open, 3, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_open, 7, .inherited = true},
    {field_type, 4},
  [466] =
    {field_open, 3, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_open, 7, .inherited = true},
    {field_type, 5},
  [470] =
    {field_open, 4, .inherited = true},
    {field_open, 6, .inherited = true},
    {field_open, 7, .inherited = true},
    {field_type, 5},
    {field_type, 6},
    {field_variable, 4},
  [476] =
    {field_open, 4, .inherited = true},
    {field_open, 7, .inherited = true},
  [478] =
    {field_open, 4, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_open, 7, .inherited = true},
    {field_type, 5},
  [482] =
    {field_open, 4, .inherited = true},
    {field_open, 6, .inherited = true},
    {field_open, 7, .inherited = true},
    {field_type, 6},
  [486] =
    {field_open, 1, .inherited = true},
    {field_open, 2, .inherited = true},
    {field_open, 7, .inherited = true},
    {field_type, 2},
  [490] =
    {field_open, 1, .inherited = true},
    {field_open, 7, .inherited = true},
  [492] =
    {field_open, 1, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_open, 7, .inherited = true},
    {field_type, 3},
  [496] =
    {field_open, 2, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_open, 8, .inherited = true},
    {field_type, 3},
  [500] =
    {field_open, 2, .inherited = true},
    {field_open, 8, .inherited = true},
  [502] =
    {field_open, 2, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_open, 8, .inherited = true},
    {field_type, 4},
  [506] =
    {field_open, 3, .inherited = true},
    {field_open, 8, .inherited = true},
  [508] =
    {field_open, 3, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_open, 8, .inherited = true},
    {field_type, 4},
  [512] =
    {field_open, 3, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_open, 8, .inherited = true},
    {field_type, 5},
  [516] =
    {field_open, 4, .inherited = true},
    {field_open, 8, .inherited = true},
  [518] =
    {field_open, 4, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_open, 8, .inherited = true},
    {field_type, 5},
  [522] =
    {field_open, 4, .inherited = true},
    {field_open, 6, .inherited = true},
    {field_open, 8, .inherited = true},
    {field_type, 6},
  [526] =
    {field_open, 1, .inherited = true},
    {field_open, 2, .inherited = true},
    {field_open, 8, .inherited = true},
    {field_type, 2},
  [530] =
    {field_open, 1, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_open, 8, .inherited = true},
    {field_type, 3},
  [534] =
    {field_open, 2, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_open, 9, .inherited = true},
    {field_type, 3},
  [538] =
    {field_open, 2, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_open, 9, .inherited = true},
    {field_type, 4},
  [542] =
    {field_open, 3, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_open, 9, .inherited = true},
    {field_type, 4},
  [546] =
    {field_open, 3, .inherited = true},
    {field_open, 9, .inherited = true},
  [548] =
    {field_open, 3, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_open, 9, .inherited = true},
    {field_type, 5},
  [552] =
    {field_open, 4, .inherited = true},
    {field_open, 9, .inherited = true},
  [554] =
    {field_open, 4, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_open, 9, .inherited = true},
    {field_type, 5},
  [558] =
    {field_open, 4, .inherited = true},
    {field_open, 6, .inherited = true},
    {field_open, 9, .inherited = true},
    {field_type, 6},
  [562] =
    {field_open, 1, .inherited = true},
    {field_open, 3, .inherited = true},
    {field_open, 9, .inherited = true},
    {field_type, 3},
  [566] =
    {field_open, 2, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_open, 10, .inherited = true},
    {field_type, 4},
  [570] =
    {field_open, 3, .inherited = true},
    {field_open, 4, .inherited = true},
    {field_open, 10, .inherited = true},
    {field_type, 4},
  [574] =
    {field_open, 3, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_open, 10, .inherited = true},
    {field_type, 5},
  [578] =
    {field_open, 4, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_open, 10, .inherited = true},
    {field_type, 5},
  [582] =
    {field_open, 4, .inherited = true},
    {field_open, 10, .inherited = true},
  [584] =
    {field_open, 4, .inherited = true},
    {field_open, 6, .inherited = true},
    {field_open, 10, .inherited = true},
    {field_type, 6},
  [588] =
    {field_open, 3, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_open, 11, .inherited = true},
    {field_type, 5},
  [592] =
    {field_open, 4, .inherited = true},
    {field_open, 5, .inherited = true},
    {field_open, 11, .inherited = true},
    {field_type, 5},
  [596] =
    {field_open, 4, .inherited = true},
    {field_open, 6, .inherited = true},
    {field_open, 11, .inherited = true},
    {field_type, 6},
  [600] =
    {field_open, 4, .inherited = true},
    {field_open, 6, .inherited = true},
    {field_open, 12, .inherited = true},
    {field_type, 6},
};

static const TSSymbol ts_alias_sequences[PRODUCTION_ID_COUNT][MAX_ALIAS_SEQUENCE_LENGTH] = {
  [0] = {0},
  [8] = {
    [0] = sym_vec_lit,
    [1] = sym_vec_lit,
  },
  [19] = {
    [1] = sym_vec_lit,
    [2] = sym_vec_lit,
  },
  [20] = {
    [0] = sym_vec_lit,
    [1] = sym_vec_lit,
    [2] = sym_vec_lit,
  },
  [22] = {
    [0] = sym_num_lit,
  },
  [34] = {
    [1] = sym_vec_lit,
    [2] = sym_vec_lit,
    [3] = sym_vec_lit,
  },
  [36] = {
    [0] = sym_sym_lit,
  },
};

static const uint16_t ts_non_terminal_alias_map[] = {
  sym_list_lit, 2,
    sym_list_lit,
    sym_vec_lit,
  aux_sym__sym_lit_without_slash_repeat1, 2,
    aux_sym__sym_lit_without_slash_repeat1,
    sym_sym_lit,
  0,
};

// Filled from tables.zlib by grammar.go, see fetchgrammar.
static TSStateId ts_primary_state_ids[STATE_COUNT];

static TSCharacterRange aux_sym_sym_lit_token1_character_set_1[] = {
  {0, 0x08}, {0x0e, 0x1b}, {'!', '!'}, {'$', '&'}, {'*', '+'}, {'-', '9'}, {'<', 'Z'}, {'_', '_'},
  {'a', 'z'}, {'|', '|'}, {'~', 0x167f}, {0x1681, 0x1fff}, {0x2007, 0x2007}, {0x200b, 0x2027}, {0x202a, 0x205e}, {0x2060, 0x2fff},
  {0x3001, 0x10ffff},
};

static TSCharacterRange aux_sym_sym_lit_token1_character_set_2[] = {
  {0, 0x08}, {0x0e, 0x1b}, {'!', '!'}, {'#', '\''}, {'*', '+'}, {'-', '9'}, {'<', 'Z'}, {'_', '_'},
  {'a', 'z'}, {'|', '|'}, {'~', 0x167f}, {0x1681, 0x1fff}, {0x2007, 0x2007}, {0x200b, 0x2027}, {0x202a, 0x205e}, {0x2060, 0x2fff},
  {0x3001, 0x10ffff},
};

static TSCharacterRange aux_sym_format_directive_type_token11_character_set_1[] = {
  {'$', '$'}, {'B', 'B'}, {'D', 'G'}, {'O', 'O'}, {'R', 'T'}, {'X', 'X'}, {'b', 'b'}, {'d', 'g'},
  {'o', 'o'}, {'r', 't'}, {'x', 'x'},
};

static TSCharacterRange aux_sym__sym_lit_without_slash_token1_character_set_1[] = {
  {0, 0x08}, {0x0e, 0x1b}, {'!', '!'}, {'$', '&'}, {'*', '+'}, {'-', '.'}, {'0', '9'}, {'<', 'Z'},
  {'_', '_'}, {'a', 'z'}, {'|', '|'}, {'~', 0x167f}, {0x1681, 0x1fff}, {0x2007, 0x2007}, {0x200b, 0x2027}, {0x202a, 0x205e},
  {0x2060, 0x2fff}, {0x3001, 0x10ffff},
};

static bool ts_lex(TSLexer *lexer, TSStateId state) {
  START_LEXER();
  eof = lexer->eof(lexer);
  switch (state) {
    case 0:
      if (eof) ADVANCE(249);
      ADVANCE_MAP(
        '\n', 305,
        '\r', 305,
        '"', 299,
        '#', 301,
        '%', 305,
        '&', 305,
        '\'', 305,
        '(', 305,
        ')', 305,
        '*', 305,
        ',', 305,
        '.', 305,
        '/', 305,
        ':', 305,
        ';', 305,
        '=', 305,
        '?', 305,
        '@', 305,
        'V', 305,
        '\\', 497,
        '^', 305,
        '_', 305,
        '`', 305,
        'v', 305,
        '{', 305,
        '|', 305,
        '}', 305,
        '~', 307,
        '<', 305,
        '>', 305,
        'A', 305,
        'a', 305,
        'C', 305,
        'c', 305,
        'I', 305,
        'i', 305,
        'L', 305,
        'l', 305,
        'P', 305,
        'p', 305,
        'W', 305,
        'w', 305,
        '[', 305,
        ']', 305,
        'D', 305,
        'F', 305,
        'S', 305,
        'd', 305,
        'f', 305,
        's', 305,
        '$', 305,
        'B', 305,
        'E', 305,
        'G', 305,
        'O', 305,
        'R', 305,
        'T', 305,
        'X', 305,
        'b', 305,
        'e', 305,
        'g', 305,
        'o', 305,
        'r', 305,
        't', 305,
        'x', 305,
      );
      if (('\t' <= lookahead && lookahead <= '\f') ||
          (0x1c <= lookahead && lookahead <= ' ') ||
          lookahead == 0x1680 ||
          (0x2000 <= lookahead && lookahead <= 0x2006) ||
          (0x2008 <= lookahead && lookahead <= 0x200a) ||
          lookahead == 0x2028 ||
          lookahead == 0x2029 ||
          lookahead == 0x205f ||
          lookahead == 0x3000) ADVANCE(305);
      if (lookahead != 0) ADVANCE(305);
      END_STATE();
    case 1:
      ADVANCE_MAP(
        '\n', 509,
        '\r', 510,
        '"', 299,
        '#', 258,
        '%', 504,
        '&', 505,
        '\'', 490,
        '*', 521,
        ',', 493,
        '/', 522,
        ':', 297,
        ';', 520,
        '?', 523,
        '@', 501,
        'N', 78,
        'V', 500,
        '^', 508,
        '_', 515,
        'v', 499,
        '|', 506,
        '~', 307,
        '+', 26,
        '-', 26,
        '<', 519,
        '>', 519,
        'A', 514,
        'a', 514,
        'C', 507,
        'c', 507,
        'I', 512,
        'i', 512,
        'P', 511,
        'p', 511,
        'W', 513,
        'w', 513,
        '[', 518,
        ']', 518,
        '{', 517,
        '}', 517,
        '(', 516,
        ')', 516,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(265);
      if (set_contains(aux_sym_format_directive_type_token11_character_set_1, 11, lookahead)) ADVANCE(525);
      END_STATE();
    case 2:
      if (lookahead == ' ') ADVANCE(104);
      END_STATE();
    case 3:
      if (lookahead == ' ') ADVANCE(110);
      END_STATE();
    case 4:
      if (lookahead == ' ') ADVANCE(110);
      if (lookahead == 's') ADVANCE(3);
      END_STATE();
    case 5:
      if (lookahead == '!') ADVANCE(252);
      if (lookahead == '?') ADVANCE(486);
      if (lookahead == '^') ADVANCE(479);
      if (lookahead == '_') ADVANCE(253);
      if (lookahead == '|') ADVANCE(229);
      END_STATE();
    case 6:
      ADVANCE_MAP(
        '!', 252,
        'C', 622,
        '^', 479,
        '_', 253,
        'c', 623,
        '|', 229,
        'B', 232,
        'b', 232,
        'O', 233,
        'o', 233,
        'X', 234,
        'x', 234,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(238);
      END_STATE();
    case 7:
      if (lookahead == '!') ADVANCE(252);
      if (lookahead == '^') ADVANCE(479);
      if (lookahead == '_') ADVANCE(253);
      if (lookahead == '|') ADVANCE(229);
      END_STATE();
    case 8:
      ADVANCE_MAP(
        '!', 252,
        '^', 479,
        '_', 253,
        '|', 229,
        'B', 232,
        'b', 232,
        'O', 233,
        'o', 233,
        'X', 234,
        'x', 234,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(238);
      END_STATE();
    case 9:
      ADVANCE_MAP(
        '"', 299,
        '#', 254,
        '\'', 490,
        '(', 480,
        ')', 481,
        ',', 494,
        '.', 261,
        ':', 296,
        ';', 252,
        '=', 561,
        '^', 478,
        '`', 491,
        'a', 377,
        'c', 374,
        'd', 412,
        'e', 378,
        'f', 364,
        'i', 349,
        'm', 312,
        'n', 323,
        'r', 332,
        's', 453,
        't', 359,
        'u', 395,
        'w', 357,
        '|', 465,
        '+', 311,
        '-', 311,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(264);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          (0x1c <= lookahead && lookahead <= ' ') ||
          lookahead == 0x1680 ||
          (0x2000 <= lookahead && lookahead <= 0x2006) ||
          (0x2008 <= lookahead && lookahead <= 0x200a) ||
          lookahead == 0x2028 ||
          lookahead == 0x2029 ||
          lookahead == 0x205f ||
          lookahead == 0x3000) ADVANCE(250);
      if (lookahead != 0 &&
          (lookahead < '[' || '^' < lookahead) &&
          (lookahead < '{' || '}' < lookahead)) ADVANCE(477);
      END_STATE();
    case 10:
      ADVANCE_MAP(
        '"', 299,
        '#', 254,
        '\'', 490,
        '(', 480,
        ')', 481,
        ',', 494,
        '.', 261,
        ':', 296,
        ';', 252,
        '=', 561,
        '^', 478,
        '`', 491,
        'a', 377,
        'c', 374,
        'd', 285,
        'e', 378,
        'f', 282,
        'i', 349,
        'm', 312,
        'n', 323,
        'r', 332,
        's', 292,
        't', 359,
        'u', 395,
        'w', 357,
        '|', 465,
        '+', 311,
        '-', 311,
        'D', 294,
        'F', 294,
        'L', 294,
        'S', 294,
        'l', 294,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(264);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          (0x1c <= lookahead && lookahead <= ' ') ||
          lookahead == 0x1680 ||
          (0x2000 <= lookahead && lookahead <= 0x2006) ||
          (0x2008 <= lookahead && lookahead <= 0x200a) ||
          lookahead == 0x2028 ||
          lookahead == 0x2029 ||
          lookahead == 0x205f ||
          lookahead == 0x3000) ADVANCE(250);
      if (lookahead != 0 &&
          (lookahead < '[' || '^' < lookahead) &&
          (lookahead < '{' || '}' < lookahead)) ADVANCE(477);
      END_STATE();
    case 11:
      ADVANCE_MAP(
        '"', 299,
        '#', 254,
        '\'', 490,
        '(', 480,
        ')', 481,
        ',', 494,
        '.', 261,
        ':', 296,
        ';', 252,
        '^', 478,
        '`', 491,
        'a', 377,
        'c', 374,
        'd', 412,
        'e', 378,
        'f', 364,
        'i', 349,
        'm', 312,
        'n', 323,
        'r', 332,
        's', 453,
        't', 359,
        'u', 395,
        'w', 357,
        '|', 465,
        '+', 311,
        '-', 311,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(264);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          (0x1c <= lookahead && lookahead <= ' ') ||
          lookahead == 0x1680 ||
          (0x2000 <= lookahead && lookahead <= 0x2006) ||
          (0x2008 <= lookahead && lookahead <= 0x200a) ||
          lookahead == 0x2028 ||
          lookahead == 0x2029 ||
          lookahead == 0x205f ||
          lookahead == 0x3000) ADVANCE(250);
      if (lookahead != 0 &&
          (lookahead < '[' || '^' < lookahead) &&
          (lookahead < '{' || '}' < lookahead)) ADVANCE(477);
      END_STATE();
    case 12:
      ADVANCE_MAP(
        '"', 299,
        '#', 254,
        '\'', 490,
        '(', 480,
        ')', 481,
        ',', 494,
        '.', 261,
        ':', 296,
        ';', 252,
        '^', 478,
        '`', 491,
        'a', 377,
        'c', 374,
        'd', 285,
        'e', 378,
        'f', 282,
        'i', 349,
        'm', 312,
        'n', 323,
        'r', 332,
        's', 292,
        't', 359,
        'u', 395,
        'w', 357,
        '|', 465,
        '+', 311,
        '-', 311,
        'D', 294,
        'F', 294,
        'L', 294,
        'S', 294,
        'l', 294,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(264);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          (0x1c <= lookahead && lookahead <= ' ') ||
          lookahead == 0x1680 ||
          (0x2000 <= lookahead && lookahead <= 0x2006) ||
          (0x2008 <= lookahead && lookahead <= 0x200a) ||
          lookahead == 0x2028 ||
          lookahead == 0x2029 ||
          lookahead == 0x205f ||
          lookahead == 0x3000) ADVANCE(250);
      if (lookahead != 0 &&
          (lookahead < '[' || '^' < lookahead) &&
          (lookahead < '{' || '}' < lookahead)) ADVANCE(477);
      END_STATE();
    case 13:
      ADVANCE_MAP(
        '"', 299,
        '#', 254,
        '\'', 490,
        '(', 480,
        ')', 481,
        ',', 494,
        '.', 261,
        ':', 296,
        ';', 252,
        '^', 478,
        '`', 491,
        'c', 375,
        'd', 281,
        'l', 280,
        'n', 363,
        '|', 465,
        '+', 311,
        '-', 311,
        'D', 294,
        'F', 294,
        'L', 294,
        'S', 294,
        'f', 294,
        's', 294,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(264);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          (0x1c <= lookahead && lookahead <= ' ') ||
          lookahead == 0x1680 ||
          (0x2000 <= lookahead && lookahead <= 0x2006) ||
          (0x2008 <= lookahead && lookahead <= 0x200a) ||
          lookahead == 0x2028 ||
          lookahead == 0x2029 ||
          lookahead == 0x205f ||
          lookahead == 0x3000) ADVANCE(250);
      if (lookahead != 0 &&
          (lookahead < '[' || '^' < lookahead) &&
          (lookahead < '{' || '}' < lookahead)) ADVANCE(477);
      END_STATE();
    case 14:
      ADVANCE_MAP(
        '"', 299,
        '#', 254,
        '\'', 490,
        '(', 480,
        ')', 481,
        ',', 494,
        '.', 261,
        ':', 296,
        ';', 252,
        '^', 478,
        '`', 491,
        'c', 375,
        'd', 337,
        'l', 318,
        'n', 363,
        '|', 465,
        '+', 311,
        '-', 311,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(264);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          (0x1c <= lookahead && lookahead <= ' ') ||
          lookahead == 0x1680 ||
          (0x2000 <= lookahead && lookahead <= 0x2006) ||
          (0x2008 <= lookahead && lookahead <= 0x200a) ||
          lookahead == 0x2028 ||
          lookahead == 0x2029 ||
          lookahead == 0x205f ||
          lookahead == 0x3000) ADVANCE(250);
      if (lookahead != 0 &&
          (lookahead < '[' || '^' < lookahead) &&
          (lookahead < '{' || '}' < lookahead)) ADVANCE(477);
      END_STATE();
    case 15:
      ADVANCE_MAP(
        '"', 299,
        '#', 254,
        '\'', 490,
        '(', 480,
        ',', 494,
        '.', 261,
        ':', 296,
        ';', 252,
        '=', 561,
        '^', 478,
        '`', 491,
        'a', 321,
        'b', 335,
        'c', 375,
        'd', 288,
        'f', 291,
        'i', 400,
        'n', 363,
        'o', 401,
        't', 361,
        'u', 427,
        '|', 465,
        '+', 311,
        '-', 311,
        'D', 294,
        'F', 294,
        'L', 294,
        'S', 294,
        'l', 294,
        's', 294,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(264);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          (0x1c <= lookahead && lookahead <= ' ') ||
          lookahead == 0x1680 ||
          (0x2000 <= lookahead && lookahead <= 0x2006) ||
          (0x2008 <= lookahead && lookahead <= 0x200a) ||
          lookahead == 0x2028 ||
          lookahead == 0x2029 ||
          lookahead == 0x205f ||
          lookahead == 0x3000) ADVANCE(250);
      if (lookahead != 0 &&
          (lookahead < '\'' || ')' < lookahead) &&
          (lookahead < '[' || '^' < lookahead) &&
          (lookahead < '{' || '}' < lookahead)) ADVANCE(477);
      END_STATE();
    case 16:
      ADVANCE_MAP(
        '"', 299,
        '#', 254,
        '\'', 490,
        '(', 480,
        ',', 494,
        '.', 261,
        ':', 296,
        ';', 252,
        '=', 561,
        '^', 478,
        '`', 491,
        'a', 321,
        'b', 335,
        'c', 375,
        'd', 416,
        'f', 433,
        'i', 400,
        'n', 363,
        'o', 401,
        't', 361,
        'u', 427,
        '|', 465,
        '+', 311,
        '-', 311,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(264);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          (0x1c <= lookahead && lookahead <= ' ') ||
          lookahead == 0x1680 ||
          (0x2000 <= lookahead && lookahead <= 0x2006) ||
          (0x2008 <= lookahead && lookahead <= 0x200a) ||
          lookahead == 0x2028 ||
          lookahead == 0x2029 ||
          lookahead == 0x205f ||
          lookahead == 0x3000) ADVANCE(250);
      if (lookahead != 0 &&
          (lookahead < '\'' || ')' < lookahead) &&
          (lookahead < '[' || '^' < lookahead) &&
          (lookahead < '{' || '}' < lookahead)) ADVANCE(477);
      END_STATE();
    case 17:
      ADVANCE_MAP(
        '"', 299,
        '#', 256,
        '\'', 490,
        '(', 480,
        ',', 494,
        '.', 261,
        ':', 296,
        '\\', 239,
        '^', 478,
        '`', 491,
        'c', 375,
        'n', 363,
        '{', 482,
        '|', 465,
        '+', 311,
        '-', 311,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(263);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 18:
      ADVANCE_MAP(
        '"', 299,
        '#', 255,
        '\'', 490,
        '(', 480,
        ',', 39,
        ':', 296,
        ';', 252,
        '^', 478,
        '{', 482,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          (0x1c <= lookahead && lookahead <= ' ') ||
          lookahead == 0x1680 ||
          (0x2000 <= lookahead && lookahead <= 0x2006) ||
          (0x2008 <= lookahead && lookahead <= 0x200a) ||
          lookahead == 0x2028 ||
          lookahead == 0x2029 ||
          lookahead == 0x205f ||
          lookahead == 0x3000) ADVANCE(250);
      if (lookahead != 0 &&
          (lookahead < '\'' || ')' < lookahead) &&
          (lookahead < '[' || '^' < lookahead) &&
          lookahead != '`' &&
          lookahead != '}') ADVANCE(477);
      END_STATE();
    case 19:
      ADVANCE_MAP(
        '"', 299,
        '#', 5,
        '(', 480,
        ',', 493,
        ':', 296,
        ';', 252,
        '^', 478,
        '{', 482,
        'D', 294,
        'F', 294,
        'L', 294,
        'S', 294,
        'd', 294,
        'f', 294,
        'l', 294,
        's', 294,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          (0x1c <= lookahead && lookahead <= ' ') ||
          lookahead == 0x1680 ||
          (0x2000 <= lookahead && lookahead <= 0x2006) ||
          (0x2008 <= lookahead && lookahead <= 0x200a) ||
          lookahead == 0x2028 ||
          lookahead == 0x2029 ||
          lookahead == 0x205f ||
          lookahead == 0x3000) ADVANCE(250);
      if (lookahead != 0 &&
          (lookahead < '\'' || ')' < lookahead) &&
          (lookahead < '[' || '^' < lookahead) &&
          lookahead != '`' &&
          lookahead != '}') ADVANCE(477);
      END_STATE();
    case 20:
      ADVANCE_MAP(
        '"', 299,
        '#', 5,
        '(', 480,
        ',', 493,
        ':', 296,
        ';', 252,
        '^', 478,
        '{', 482,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          (0x1c <= lookahead && lookahead <= ' ') ||
          lookahead == 0x1680 ||
          (0x2000 <= lookahead && lookahead <= 0x2006) ||
          (0x2008 <= lookahead && lookahead <= 0x200a) ||
          lookahead == 0x2028 ||
          lookahead == 0x2029 ||
          lookahead == 0x205f ||
          lookahead == 0x3000) ADVANCE(250);
      if (lookahead != 0 &&
          (lookahead < '\'' || ')' < lookahead) &&
          (lookahead < '[' || '^' < lookahead) &&
          lookahead != '`' &&
          lookahead != '}') ADVANCE(477);
      END_STATE();
    case 21:
      if (lookahead == '"') ADVANCE(299);
      if (lookahead == '#') ADVANCE(303);
      if (lookahead == '\\') ADVANCE(246);
      if (lookahead == '~') ADVANCE(307);
      if (lookahead != 0) ADVANCE(305);
      END_STATE();
    case 22:
      ADVANCE_MAP(
        '"', 23,
        '#', 7,
        '(', 480,
        ')', 481,
        ',', 493,
        '/', 522,
        ':', 296,
        ';', 252,
        '=', 560,
        '\\', 239,
        '^', 478,
        'a', 52,
        'b', 67,
        'c', 126,
        'd', 286,
        'e', 136,
        'f', 284,
        'i', 90,
        'm', 44,
        'n', 57,
        'o', 146,
        'r', 68,
        's', 293,
        't', 105,
        'u', 148,
        'w', 102,
        '{', 482,
        'D', 279,
        'F', 279,
        'L', 279,
        'S', 279,
        'l', 279,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(236);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          (0x1c <= lookahead && lookahead <= ' ') ||
          lookahead == 0x1680 ||
          (0x2000 <= lookahead && lookahead <= 0x2006) ||
          (0x2008 <= lookahead && lookahead <= 0x200a) ||
          lookahead == 0x2028 ||
          lookahead == 0x2029 ||
          lookahead == 0x205f ||
          lookahead == 0x3000) ADVANCE(250);
      END_STATE();
    case 23:
      if (lookahead == '"') ADVANCE(616);
      if (lookahead == '\\') ADVANCE(245);
      if (lookahead != 0) ADVANCE(23);
      END_STATE();
    case 24:
      if (lookahead == '#') ADVANCE(495);
      if (lookahead != 0) ADVANCE(229);
      END_STATE();
    case 25:
      ADVANCE_MAP(
        '#', 6,
        '(', 480,
        ')', 481,
        ',', 493,
        ':', 296,
        ';', 252,
        '^', 478,
        'a', 124,
        'c', 126,
        'd', 287,
        'e', 136,
        'f', 283,
        'i', 91,
        'm', 44,
        'n', 57,
        'r', 68,
        's', 293,
        't', 107,
        'u', 147,
        'w', 102,
        '+', 26,
        '-', 26,
        'D', 279,
        'F', 279,
        'L', 279,
        'S', 279,
        'l', 279,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(265);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          (0x1c <= lookahead && lookahead <= ' ') ||
          lookahead == 0x1680 ||
          (0x2000 <= lookahead && lookahead <= 0x2006) ||
          (0x2008 <= lookahead && lookahead <= 0x200a) ||
          lookahead == 0x2028 ||
          lookahead == 0x2029 ||
          lookahead == 0x205f ||
          lookahead == 0x3000) ADVANCE(250);
      END_STATE();
    case 26:
      if (lookahead == '#') ADVANCE(237);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(265);
      END_STATE();
    case 27:
      ADVANCE_MAP(
        '#', 255,
        '\'', 490,
        '(', 480,
        ')', 481,
        ',', 39,
        ';', 252,
        '^', 478,
        '{', 482,
        'D', 279,
        'F', 279,
        'L', 279,
        'S', 279,
        'd', 279,
        'f', 279,
        'l', 279,
        's', 279,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          (0x1c <= lookahead && lookahead <= ' ') ||
          lookahead == 0x1680 ||
          (0x2000 <= lookahead && lookahead <= 0x2006) ||
          (0x2008 <= lookahead && lookahead <= 0x200a) ||
          lookahead == 0x2028 ||
          lookahead == 0x2029 ||
          lookahead == 0x205f ||
          lookahead == 0x3000) ADVANCE(250);
      END_STATE();
    case 28:
      ADVANCE_MAP(
        '#', 255,
        '\'', 490,
        '(', 480,
        ',', 39,
        ':', 296,
        ';', 252,
        '=', 560,
        '^', 478,
        'a', 51,
        'b', 67,
        'c', 125,
        'd', 289,
        'f', 290,
        'i', 154,
        'o', 146,
        't', 108,
        'u', 182,
        '{', 482,
        'D', 279,
        'F', 279,
        'L', 279,
        'S', 279,
        'l', 279,
        's', 279,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          (0x1c <= lookahead && lookahead <= ' ') ||
          lookahead == 0x1680 ||
          (0x2000 <= lookahead && lookahead <= 0x2006) ||
          (0x2008 <= lookahead && lookahead <= 0x200a) ||
          lookahead == 0x2028 ||
          lookahead == 0x2029 ||
          lookahead == 0x205f ||
          lookahead == 0x3000) ADVANCE(250);
      END_STATE();
    case 29:
      ADVANCE_MAP(
        '#', 257,
        '\'', 490,
        '(', 480,
        ',', 39,
        '=', 561,
        '^', 478,
        'a', 321,
        'b', 335,
        'd', 416,
        'f', 433,
        'i', 400,
        'o', 401,
        't', 361,
        'u', 427,
        '{', 482,
      );
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 30:
      ADVANCE_MAP(
        '#', 8,
        '(', 480,
        ',', 493,
        ':', 295,
        ';', 252,
        '^', 478,
        'c', 125,
        'i', 163,
        '+', 26,
        '-', 26,
        'D', 279,
        'F', 279,
        'L', 279,
        'S', 279,
        'd', 279,
        'f', 279,
        'l', 279,
        's', 279,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(265);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          (0x1c <= lookahead && lookahead <= ' ') ||
          lookahead == 0x1680 ||
          (0x2000 <= lookahead && lookahead <= 0x2006) ||
          (0x2008 <= lookahead && lookahead <= 0x200a) ||
          lookahead == 0x2028 ||
          lookahead == 0x2029 ||
          lookahead == 0x205f ||
          lookahead == 0x3000) ADVANCE(250);
      END_STATE();
    case 31:
      if (lookahead == '#') ADVANCE(228);
      if (lookahead == '/') ADVANCE(522);
      if (lookahead == ':') ADVANCE(296);
      if ((!eof && set_contains(aux_sym__sym_lit_without_slash_token1_character_set_1, 18, lookahead))) ADVANCE(617);
      END_STATE();
    case 32:
      if (lookahead == '#') ADVANCE(228);
      if (lookahead == '=') ADVANCE(561);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_1, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 33:
      if (lookahead == '#') ADVANCE(228);
      if (lookahead == 'c') ADVANCE(618);
      if ((!eof && set_contains(aux_sym__sym_lit_without_slash_token1_character_set_1, 18, lookahead))) ADVANCE(617);
      END_STATE();
    case 34:
      if (lookahead == '#') ADVANCE(228);
      if (lookahead == 'd') ADVANCE(337);
      if (lookahead == 'l') ADVANCE(424);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_1, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 35:
      if (lookahead == '#') ADVANCE(228);
      if (lookahead == 'd') ADVANCE(76);
      if (lookahead == 'l') ADVANCE(180);
      END_STATE();
    case 36:
      if (lookahead == '#') ADVANCE(498);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(497);
      END_STATE();
    case 37:
      if (lookahead == '-') ADVANCE(121);
      END_STATE();
    case 38:
      if (lookahead == '-') ADVANCE(200);
      END_STATE();
    case 39:
      if (lookahead == '@') ADVANCE(492);
      END_STATE();
    case 40:
      if (lookahead == 'A') ADVANCE(484);
      if (lookahead == 'a') ADVANCE(485);
      if (lookahead == '#' ||
          lookahead == '=') ADVANCE(619);
      if (lookahead == 'R' ||
          lookahead == 'r') ADVANCE(244);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(231);
      END_STATE();
    case 41:
      if (lookahead == 'a') ADVANCE(55);
      END_STATE();
    case 42:
      if (lookahead == 'a') ADVANCE(201);
      END_STATE();
    case 43:
      if (lookahead == 'a') ADVANCE(123);
      END_STATE();
    case 44:
      if (lookahead == 'a') ADVANCE(221);
      if (lookahead == 'i') ADVANCE(158);
      END_STATE();
    case 45:
      if (lookahead == 'a') ADVANCE(226);
      END_STATE();
    case 46:
      if (lookahead == 'a') ADVANCE(58);
      if (lookahead == 'e') ADVANCE(211);
      END_STATE();
    case 47:
      if (lookahead == 'a') ADVANCE(207);
      END_STATE();
    case 48:
      if (lookahead == 'a') ADVANCE(131);
      END_STATE();
    case 49:
      if (lookahead == 'a') ADVANCE(132);
      END_STATE();
    case 50:
      if (lookahead == 'b') ADVANCE(171);
      END_STATE();
    case 51:
      if (lookahead == 'b') ADVANCE(181);
      if (lookahead == 'c') ADVANCE(189);
      END_STATE();
    case 52:
      if (lookahead == 'b') ADVANCE(181);
      if (lookahead == 'c') ADVANCE(189);
      if (lookahead == 'l') ADVANCE(218);
      if (lookahead == 'n') ADVANCE(59);
      if (lookahead == 'p') ADVANCE(184);
      if (lookahead == 's') ADVANCE(569);
      END_STATE();
    case 53:
      if (lookahead == 'c') ADVANCE(563);
      END_STATE();
    case 54:
      if (lookahead == 'c') ADVANCE(609);
      END_STATE();
    case 55:
      if (lookahead == 'c') ADVANCE(100);
      END_STATE();
    case 56:
      if (lookahead == 'c') ADVANCE(206);
      END_STATE();
    case 57:
      if (lookahead == 'c') ADVANCE(179);
      if (lookahead == 'e') ADVANCE(215);
      END_STATE();
    case 58:
      if (lookahead == 'c') ADVANCE(195);
      END_STATE();
    case 59:
      if (lookahead == 'd') ADVANCE(567);
      END_STATE();
    case 60:
      if (lookahead == 'd') ADVANCE(563);
      END_STATE();
    case 61:
      if (lookahead == 'd') ADVANCE(611);
      END_STATE();
    case 62:
      if (lookahead == 'e') ADVANCE(41);
      if (lookahead == 't') ADVANCE(99);
      END_STATE();
    case 63:
      if (lookahead == 'e') ADVANCE(2);
      END_STATE();
    case 64:
      if (lookahead == 'e') ADVANCE(144);
      END_STATE();
    case 65:
      if (lookahead == 'e') ADVANCE(222);
      END_STATE();
    case 66:
      if (lookahead == 'e') ADVANCE(4);
      END_STATE();
    case 67:
      if (lookahead == 'e') ADVANCE(114);
      if (lookahead == 'y') ADVANCE(556);
      END_STATE();
    case 68:
      if (lookahead == 'e') ADVANCE(185);
      END_STATE();
    case 69:
      if (lookahead == 'e') ADVANCE(151);
      END_STATE();
    case 70:
      if (lookahead == 'e') ADVANCE(594);
      END_STATE();
    case 71:
      if (lookahead == 'e') ADVANCE(540);
      END_STATE();
    case 72:
      if (lookahead == 'e') ADVANCE(576);
      END_STATE();
    case 73:
      if (lookahead == 'e') ADVANCE(562);
      if (lookahead == 'i') ADVANCE(160);
      END_STATE();
    case 74:
      if (lookahead == 'e') ADVANCE(524);
      END_STATE();
    case 75:
      if (lookahead == 'e') ADVANCE(202);
      END_STATE();
    case 76:
      if (lookahead == 'e') ADVANCE(92);
      END_STATE();
    case 77:
      if (lookahead == 'e') ADVANCE(56);
      END_STATE();
    case 78:
      if (lookahead == 'e') ADVANCE(220);
      END_STATE();
    case 79:
      if (lookahead == 'e') ADVANCE(187);
      END_STATE();
    case 80:
      if (lookahead == 'e') ADVANCE(47);
      END_STATE();
    case 81:
      if (lookahead == 'e') ADVANCE(190);
      END_STATE();
    case 82:
      if (lookahead == 'e') ADVANCE(116);
      END_STATE();
    case 83:
      if (lookahead == 'e') ADVANCE(152);
      if (lookahead == 'i') ADVANCE(134);
      END_STATE();
    case 84:
      if (lookahead == 'e') ADVANCE(157);
      END_STATE();
    case 85:
      if (lookahead == 'e') ADVANCE(150);
      END_STATE();
    case 86:
      if (lookahead == 'e') ADVANCE(194);
      END_STATE();
    case 87:
      if (lookahead == 'e') ADVANCE(204);
      END_STATE();
    case 88:
      if (lookahead == 'e') ADVANCE(165);
      END_STATE();
    case 89:
      if (lookahead == 'f') ADVANCE(537);
      END_STATE();
    case 90:
      if (lookahead == 'f') ADVANCE(584);
      if (lookahead == 'n') ADVANCE(529);
      END_STATE();
    case 91:
      if (lookahead == 'f') ADVANCE(584);
      if (lookahead == 'n') ADVANCE(111);
      END_STATE();
    case 92:
      if (lookahead == 'f') ADVANCE(98);
      END_STATE();
    case 93:
      if (lookahead == 'f') ADVANCE(191);
      if (lookahead == 't') ADVANCE(167);
      END_STATE();
    case 94:
      if (lookahead == 'f') ADVANCE(192);
      if (lookahead == 't') ADVANCE(168);
      END_STATE();
    case 95:
      if (lookahead == 'g') ADVANCE(533);
      END_STATE();
    case 96:
      if (lookahead == 'g') ADVANCE(535);
      END_STATE();
    case 97:
      if (lookahead == 'g') ADVANCE(562);
      END_STATE();
    case 98:
      if (lookahead == 'g') ADVANCE(88);
      if (lookahead == 'm') ADVANCE(46);
      if (lookahead == 'u') ADVANCE(155);
      END_STATE();
    case 99:
      if (lookahead == 'h') ADVANCE(63);
      END_STATE();
    case 100:
      if (lookahead == 'h') ADVANCE(2);
      END_STATE();
    case 101:
      if (lookahead == 'h') ADVANCE(37);
      END_STATE();
    case 102:
      if (lookahead == 'h') ADVANCE(83);
      if (lookahead == 'i') ADVANCE(209);
      END_STATE();
    case 103:
      if (lookahead == 'h') ADVANCE(571);
      END_STATE();
    case 104:
      if (lookahead == 'h') ADVANCE(42);
      if (lookahead == 'p') ADVANCE(188);
      END_STATE();
    case 105:
      if (lookahead == 'h') ADVANCE(69);
      if (lookahead == 'o') ADVANCE(544);
      END_STATE();
    case 106:
      if (lookahead == 'h') ADVANCE(175);
      END_STATE();
    case 107:
      if (lookahead == 'h') ADVANCE(86);
      END_STATE();
    case 108:
      if (lookahead == 'h') ADVANCE(85);
      if (lookahead == 'o') ADVANCE(544);
      END_STATE();
    case 109:
      if (lookahead == 'i') ADVANCE(227);
      END_STATE();
    case 110:
      if (lookahead == 'i') ADVANCE(145);
      if (lookahead == 'o') ADVANCE(89);
      END_STATE();
    case 111:
      if (lookahead == 'i') ADVANCE(210);
      if (lookahead == 't') ADVANCE(169);
      END_STATE();
    case 112:
      if (lookahead == 'i') ADVANCE(143);
      END_STATE();
    case 113:
      if (lookahead == 'i') ADVANCE(54);
      END_STATE();
    case 114:
      if (lookahead == 'i') ADVANCE(149);
      if (lookahead == 'l') ADVANCE(172);
      END_STATE();
    case 115:
      if (lookahead == 'i') ADVANCE(127);
      END_STATE();
    case 116:
      if (lookahead == 'i') ADVANCE(199);
      END_STATE();
    case 117:
      if (lookahead == 'i') ADVANCE(159);
      END_STATE();
    case 118:
      if (lookahead == 'i') ADVANCE(160);
      END_STATE();
    case 119:
      if (lookahead == 'i') ADVANCE(164);
      END_STATE();
    case 120:
      if (lookahead == 'i') ADVANCE(49);
      END_STATE();
    case 121:
      if (lookahead == 'k') ADVANCE(65);
      if (lookahead == 'v') ADVANCE(43);
      END_STATE();
    case 122:
      if (lookahead == 'l') ADVANCE(4);
      END_STATE();
    case 123:
      if (lookahead == 'l') ADVANCE(213);
      END_STATE();
    case 124:
      if (lookahead == 'l') ADVANCE(218);
      if (lookahead == 'n') ADVANCE(59);
      if (lookahead == 'p') ADVANCE(184);
      if (lookahead == 's') ADVANCE(569);
      END_STATE();
    case 125:
      if (lookahead == 'l') ADVANCE(526);
      END_STATE();
    case 126:
      if (lookahead == 'l') ADVANCE(526);
      if (lookahead == 'o') ADVANCE(129);
      END_STATE();
    case 127:
      if (lookahead == 'l') ADVANCE(578);
      END_STATE();
    case 128:
      if (lookahead == 'l') ADVANCE(224);
      END_STATE();
    case 129:
      if (lookahead == 'l') ADVANCE(133);
      if (lookahead == 'u') ADVANCE(161);
      END_STATE();
    case 130:
      if (lookahead == 'l') ADVANCE(225);
      END_STATE();
    case 131:
      if (lookahead == 'l') ADVANCE(128);
      END_STATE();
    case 132:
      if (lookahead == 'l') ADVANCE(130);
      END_STATE();
    case 133:
      if (lookahead == 'l') ADVANCE(77);
      END_STATE();
    case 134:
      if (lookahead == 'l') ADVANCE(72);
      END_STATE();
    case 135:
      if (lookahead == 'l') ADVANCE(87);
      if (lookahead == 't') ADVANCE(115);
      END_STATE();
    case 136:
      if (lookahead == 'l') ADVANCE(205);
      END_STATE();
    case 137:
      if (lookahead == 'l') ADVANCE(119);
      END_STATE();
    case 138:
      if (lookahead == 'm') ADVANCE(50);
      END_STATE();
    case 139:
      if (lookahead == 'm') ADVANCE(564);
      END_STATE();
    case 140:
      if (lookahead == 'm') ADVANCE(542);
      END_STATE();
    case 141:
      if (lookahead == 'm') ADVANCE(548);
      END_STATE();
    case 142:
      if (lookahead == 'm') ADVANCE(552);
      END_STATE();
    case 143:
      if (lookahead == 'm') ADVANCE(109);
      END_STATE();
    case 144:
      if (lookahead == 'n') ADVANCE(208);
      END_STATE();
    case 145:
      if (lookahead == 'n') ADVANCE(537);
      END_STATE();
    case 146:
      if (lookahead == 'n') ADVANCE(554);
      END_STATE();
    case 147:
      if (lookahead == 'n') ADVANCE(135);
      END_STATE();
    case 148:
      if (lookahead == 'n') ADVANCE(135);
      if (lookahead == 'p') ADVANCE(93);
      if (lookahead == 's') ADVANCE(117);
      END_STATE();
    case 149:
      if (lookahead == 'n') ADVANCE(95);
      END_STATE();
    case 150:
      if (lookahead == 'n') ADVANCE(558);
      END_STATE();
    case 151:
      if (lookahead == 'n') ADVANCE(558);
      if (lookahead == 'r') ADVANCE(82);
      END_STATE();
    case 152:
      if (lookahead == 'n') ADVANCE(582);
      END_STATE();
    case 153:
      if (lookahead == 'n') ADVANCE(599);
      END_STATE();
    case 154:
      if (lookahead == 'n') ADVANCE(528);
      END_STATE();
    case 155:
      if (lookahead == 'n') ADVANCE(605);
      END_STATE();
    case 156:
      if (lookahead == 'n') ADVANCE(53);
      END_STATE();
    case 157:
      if (lookahead == 'n') ADVANCE(60);
      END_STATE();
    case 158:
      if (lookahead == 'n') ADVANCE(112);
      END_STATE();
    case 159:
      if (lookahead == 'n') ADVANCE(96);
      END_STATE();
    case 160:
      if (lookahead == 'n') ADVANCE(97);
      END_STATE();
    case 161:
      if (lookahead == 'n') ADVANCE(206);
      END_STATE();
    case 162:
      if (lookahead == 'n') ADVANCE(48);
      END_STATE();
    case 163:
      if (lookahead == 'n') ADVANCE(212);
      END_STATE();
    case 164:
      if (lookahead == 'n') ADVANCE(74);
      END_STATE();
    case 165:
      if (lookahead == 'n') ADVANCE(81);
      END_STATE();
    case 166:
      if (lookahead == 'n') ADVANCE(94);
      END_STATE();
    case 167:
      if (lookahead == 'o') ADVANCE(546);
      END_STATE();
    case 168:
      if (lookahead == 'o') ADVANCE(550);
      END_STATE();
    case 169:
      if (lookahead == 'o') ADVANCE(596);
      END_STATE();
    case 170:
      if (lookahead == 'o') ADVANCE(607);
      END_STATE();
    case 171:
      if (lookahead == 'o') ADVANCE(122);
      END_STATE();
    case 172:
      if (lookahead == 'o') ADVANCE(217);
      END_STATE();
    case 173:
      if (lookahead == 'o') ADVANCE(183);
      END_STATE();
    case 174:
      if (lookahead == 'o') ADVANCE(140);
      END_STATE();
    case 175:
      if (lookahead == 'o') ADVANCE(61);
      END_STATE();
    case 176:
      if (lookahead == 'o') ADVANCE(203);
      END_STATE();
    case 177:
      if (lookahead == 'o') ADVANCE(141);
      END_STATE();
    case 178:
      if (lookahead == 'o') ADVANCE(142);
      END_STATE();
    case 179:
      if (lookahead == 'o') ADVANCE(156);
      END_STATE();
    case 180:
      if (lookahead == 'o') ADVANCE(173);
      END_STATE();
    case 181:
      if (lookahead == 'o') ADVANCE(216);
      END_STATE();
    case 182:
      if (lookahead == 'p') ADVANCE(93);
      if (lookahead == 's') ADVANCE(117);
      END_STATE();
    case 183:
      if (lookahead == 'p') ADVANCE(603);
      END_STATE();
    case 184:
      if (lookahead == 'p') ADVANCE(84);
      END_STATE();
    case 185:
      if (lookahead == 'p') ADVANCE(80);
      if (lookahead == 't') ADVANCE(214);
      END_STATE();
    case 186:
      if (lookahead == 'r') ADVANCE(565);
      END_STATE();
    case 187:
      if (lookahead == 'r') ADVANCE(592);
      END_STATE();
    case 188:
      if (lookahead == 'r') ADVANCE(75);
      END_STATE();
    case 189:
      if (lookahead == 'r') ADVANCE(176);
      END_STATE();
    case 190:
      if (lookahead == 'r') ADVANCE(113);
      END_STATE();
    case 191:
      if (lookahead == 'r') ADVANCE(177);
      END_STATE();
    case 192:
      if (lookahead == 'r') ADVANCE(178);
      END_STATE();
    case 193:
      if (lookahead == 'r') ADVANCE(153);
      END_STATE();
    case 194:
      if (lookahead == 'r') ADVANCE(82);
      END_STATE();
    case 195:
      if (lookahead == 'r') ADVANCE(170);
      END_STATE();
    case 196:
      if (lookahead == 's') ADVANCE(531);
      END_STATE();
    case 197:
      if (lookahead == 's') ADVANCE(588);
      END_STATE();
    case 198:
      if (lookahead == 's') ADVANCE(586);
      END_STATE();
    case 199:
      if (lookahead == 's') ADVANCE(590);
      END_STATE();
    case 200:
      if (lookahead == 's') ADVANCE(223);
      END_STATE();
    case 201:
      if (lookahead == 's') ADVANCE(101);
      END_STATE();
    case 202:
      if (lookahead == 's') ADVANCE(64);
      END_STATE();
    case 203:
      if (lookahead == 's') ADVANCE(196);
      END_STATE();
    case 204:
      if (lookahead == 's') ADVANCE(198);
      END_STATE();
    case 205:
      if (lookahead == 's') ADVANCE(70);
      END_STATE();
    case 206:
      if (lookahead == 't') ADVANCE(563);
      END_STATE();
    case 207:
      if (lookahead == 't') ADVANCE(580);
      END_STATE();
    case 208:
      if (lookahead == 't') ADVANCE(38);
      END_STATE();
    case 209:
      if (lookahead == 't') ADVANCE(103);
      END_STATE();
    case 210:
      if (lookahead == 't') ADVANCE(120);
      END_STATE();
    case 211:
      if (lookahead == 't') ADVANCE(106);
      END_STATE();
    case 212:
      if (lookahead == 't') ADVANCE(169);
      END_STATE();
    case 213:
      if (lookahead == 'u') ADVANCE(66);
      END_STATE();
    case 214:
      if (lookahead == 'u') ADVANCE(193);
      END_STATE();
    case 215:
      if (lookahead == 'v') ADVANCE(79);
      END_STATE();
    case 216:
      if (lookahead == 'v') ADVANCE(71);
      END_STATE();
    case 217:
      if (lookahead == 'w') ADVANCE(538);
      END_STATE();
    case 218:
      if (lookahead == 'w') ADVANCE(45);
      END_STATE();
    case 219:
      if (lookahead == 'w') ADVANCE(166);
      END_STATE();
    case 220:
      if (lookahead == 'w') ADVANCE(137);
      END_STATE();
    case 221:
      if (lookahead == 'x') ADVANCE(112);
      END_STATE();
    case 222:
      if (lookahead == 'y') ADVANCE(4);
      END_STATE();
    case 223:
      if (lookahead == 'y') ADVANCE(138);
      END_STATE();
    case 224:
      if (lookahead == 'y') ADVANCE(597);
      END_STATE();
    case 225:
      if (lookahead == 'y') ADVANCE(601);
      END_STATE();
    case 226:
      if (lookahead == 'y') ADVANCE(197);
      END_STATE();
    case 227:
      if (lookahead == 'z') ADVANCE(73);
      END_STATE();
    case 228:
      if (lookahead == '|') ADVANCE(229);
      END_STATE();
    case 229:
      if (lookahead == '|') ADVANCE(24);
      if (lookahead != 0) ADVANCE(229);
      END_STATE();
    case 230:
      if (lookahead == '|') ADVANCE(496);
      if (lookahead != 0) ADVANCE(230);
      END_STATE();
    case 231:
      if (lookahead == '#' ||
          lookahead == '=') ADVANCE(619);
      if (lookahead == 'R' ||
          lookahead == 'r') ADVANCE(244);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(231);
      END_STATE();
    case 232:
      if (lookahead == '+' ||
          lookahead == '-') ADVANCE(240);
      if (lookahead == '0' ||
          lookahead == '1') ADVANCE(266);
      END_STATE();
    case 233:
      if (lookahead == '+' ||
          lookahead == '-') ADVANCE(241);
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(268);
      END_STATE();
    case 234:
      if (lookahead == '+' ||
          lookahead == '-') ADVANCE(243);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(274);
      END_STATE();
    case 235:
      if (lookahead == '+' ||
          lookahead == '-') ADVANCE(242);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(273);
      END_STATE();
    case 236:
      if (lookahead == 'A' ||
          lookahead == 'a') ADVANCE(259);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(236);
      END_STATE();
    case 237:
      if (lookahead == 'B' ||
          lookahead == 'b') ADVANCE(232);
      if (lookahead == 'O' ||
          lookahead == 'o') ADVANCE(233);
      if (lookahead == 'X' ||
          lookahead == 'x') ADVANCE(234);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(238);
      END_STATE();
    case 238:
      if (lookahead == 'R' ||
          lookahead == 'r') ADVANCE(244);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(238);
      END_STATE();
    case 239:
      if (lookahead == '(' ||
          lookahead == ')') ADVANCE(308);
      if (lookahead != 0 &&
          lookahead != '\t' &&
          lookahead != '\n' &&
          lookahead != '\f' &&
          lookahead != '\r' &&
          lookahead != ' ') ADVANCE(309);
      END_STATE();
    case 240:
      if (lookahead == '0' ||
          lookahead == '1') ADVANCE(266);
      END_STATE();
    case 241:
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(268);
      END_STATE();
    case 242:
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(273);
      END_STATE();
    case 243:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(274);
      END_STATE();
    case 244:
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(276);
      END_STATE();
    case 245:
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(23);
      END_STATE();
    case 246:
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(306);
      END_STATE();
    case 247:
      if (eof) ADVANCE(249);
      ADVANCE_MAP(
        '"', 299,
        '#', 254,
        '\'', 490,
        '(', 480,
        ')', 481,
        ',', 494,
        '.', 261,
        ':', 296,
        ';', 252,
        '^', 478,
        '`', 491,
        'c', 375,
        'n', 363,
        '{', 482,
        '|', 465,
        '}', 483,
        '+', 311,
        '-', 311,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(264);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          (0x1c <= lookahead && lookahead <= ' ') ||
          lookahead == 0x1680 ||
          (0x2000 <= lookahead && lookahead <= 0x2006) ||
          (0x2008 <= lookahead && lookahead <= 0x200a) ||
          lookahead == 0x2028 ||
          lookahead == 0x2029 ||
          lookahead == 0x205f ||
          lookahead == 0x3000) ADVANCE(250);
      if (lookahead != 0 &&
          (lookahead < '[' || '^' < lookahead)) ADVANCE(477);
      END_STATE();
    case 248:
      if (eof) ADVANCE(249);
      ADVANCE_MAP(
        '"', 299,
        '#', 254,
        '\'', 490,
        '(', 480,
        ')', 481,
        ',', 494,
        '.', 261,
        ':', 296,
        ';', 252,
        '^', 478,
        '`', 491,
        'c', 375,
        'n', 363,
        '|', 465,
        '}', 483,
        '+', 311,
        '-', 311,
        'D', 294,
        'F', 294,
        'L', 294,
        'S', 294,
        'd', 294,
        'f', 294,
        'l', 294,
        's', 294,
      );
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(264);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          (0x1c <= lookahead && lookahead <= ' ') ||
          lookahead == 0x1680 ||
          (0x2000 <= lookahead && lookahead <= 0x2006) ||
          (0x2008 <= lookahead && lookahead <= 0x200a) ||
          lookahead == 0x2028 ||
          lookahead == 0x2029 ||
          lookahead == 0x205f ||
          lookahead == 0x3000) ADVANCE(250);
      if (lookahead != 0 &&
          (lookahead < '[' || '^' < lookahead) &&
          (lookahead < '{' || '}' < lookahead)) ADVANCE(477);
      END_STATE();
    case 249:
      ACCEPT_TOKEN(ts_builtin_sym_end);
      END_STATE();
    case 250:
      ACCEPT_TOKEN(sym__ws);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          (0x1c <= lookahead && lookahead <= ' ') ||
          lookahead == 0x1680 ||
          (0x2000 <= lookahead && lookahead <= 0x2006) ||
          (0x2008 <= lookahead && lookahead <= 0x200a) ||
          lookahead == 0x2028 ||
          lookahead == 0x2029 ||
          lookahead == 0x205f ||
          lookahead == 0x3000) ADVANCE(250);
      END_STATE();
    case 251:
      ACCEPT_TOKEN(sym_comment);
      END_STATE();
    case 252:
      ACCEPT_TOKEN(sym_comment);
      if (lookahead == '\n') ADVANCE(251);
      if (lookahead != 0) ADVANCE(252);
      END_STATE();
    case 253:
      ACCEPT_TOKEN(anon_sym_POUND_);
      END_STATE();
    case 254:
      ACCEPT_TOKEN(anon_sym_POUND);
      ADVANCE_MAP(
        '!', 252,
        '\'', 489,
        '+', 620,
        '-', 621,
        '0', 40,
        '?', 487,
        'C', 622,
        'P', 614,
        '^', 479,
        '_', 253,
        'c', 623,
        'p', 615,
        '|', 229,
        'B', 232,
        'b', 232,
        'O', 233,
        'o', 233,
        'X', 234,
        'x', 234,
      );
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(231);
      END_STATE();
    case 255:
      ACCEPT_TOKEN(anon_sym_POUND);
      ADVANCE_MAP(
        '!', 252,
        '\'', 489,
        '+', 620,
        '-', 621,
        '?', 487,
        'C', 622,
        '^', 479,
        '_', 253,
        'c', 623,
        '|', 229,
      );
      END_STATE();
    case 256:
      ACCEPT_TOKEN(anon_sym_POUND);
      ADVANCE_MAP(
        '\'', 489,
        '+', 620,
        '-', 621,
        '0', 40,
        '?', 487,
        'C', 622,
        'P', 614,
        '^', 479,
        'c', 623,
        'p', 615,
        '|', 229,
        'B', 232,
        'b', 232,
        'O', 233,
        'o', 233,
        'X', 234,
        'x', 234,
      );
      if (('1' <= lookahead && lookahead <= '9')) ADVANCE(231);
      END_STATE();
    case 257:
      ACCEPT_TOKEN(anon_sym_POUND);
      ADVANCE_MAP(
        '\'', 489,
        '+', 620,
        '-', 621,
        '?', 487,
        'C', 622,
        '^', 479,
        'c', 623,
        '|', 229,
      );
      END_STATE();
    case 258:
      ACCEPT_TOKEN(anon_sym_POUND);
      if (lookahead == '|') ADVANCE(229);
      if (lookahead == 'B' ||
          lookahead == 'b') ADVANCE(232);
      if (lookahead == 'O' ||
          lookahead == 'o') ADVANCE(233);
      if (lookahead == 'X' ||
          lookahead == 'x') ADVANCE(234);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(238);
      END_STATE();
    case 259:
      ACCEPT_TOKEN(aux_sym__form_token1);
      END_STATE();
    case 260:
      ACCEPT_TOKEN(aux_sym__form_token1);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 261:
      ACCEPT_TOKEN(anon_sym_DOT);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 262:
      ACCEPT_TOKEN(aux_sym_num_lit_token1);
      END_STATE();
    case 263:
      ACCEPT_TOKEN(aux_sym_num_lit_token1);
      if (lookahead == '.') ADVANCE(270);
      if (lookahead == '/') ADVANCE(474);
      if (lookahead == 'A' ||
          lookahead == 'a') ADVANCE(260);
      if (lookahead == 'M' ||
          lookahead == 'N') ADVANCE(278);
      if (('D' <= lookahead && lookahead <= 'F') ||
          lookahead == 'L' ||
          lookahead == 'S' ||
          ('d' <= lookahead && lookahead <= 'f') ||
          lookahead == 'l' ||
          lookahead == 's') ADVANCE(466);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(263);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 264:
      ACCEPT_TOKEN(aux_sym_num_lit_token1);
      if (lookahead == '.') ADVANCE(270);
      if (lookahead == '/') ADVANCE(474);
      if (lookahead == 'M' ||
          lookahead == 'N') ADVANCE(278);
      if (('D' <= lookahead && lookahead <= 'F') ||
          lookahead == 'L' ||
          lookahead == 'S' ||
          ('d' <= lookahead && lookahead <= 'f') ||
          lookahead == 'l' ||
          lookahead == 's') ADVANCE(466);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(264);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 265:
      ACCEPT_TOKEN(aux_sym_num_lit_token1);
      if (lookahead == '.') ADVANCE(271);
      if (lookahead == '/') ADVANCE(242);
      if (lookahead == 'M' ||
          lookahead == 'N') ADVANCE(262);
      if (('D' <= lookahead && lookahead <= 'F') ||
          lookahead == 'L' ||
          lookahead == 'S' ||
          ('d' <= lookahead && lookahead <= 'f') ||
          lookahead == 'l' ||
          lookahead == 's') ADVANCE(235);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(265);
      END_STATE();
    case 266:
      ACCEPT_TOKEN(aux_sym_num_lit_token1);
      if (lookahead == '0' ||
          lookahead == '1') ADVANCE(266);
      END_STATE();
    case 267:
      ACCEPT_TOKEN(aux_sym_num_lit_token1);
      if (lookahead == '0' ||
          lookahead == '1') ADVANCE(267);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 268:
      ACCEPT_TOKEN(aux_sym_num_lit_token1);
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(268);
      END_STATE();
    case 269:
      ACCEPT_TOKEN(aux_sym_num_lit_token1);
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(269);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 270:
      ACCEPT_TOKEN(aux_sym_num_lit_token1);
      if (('D' <= lookahead && lookahead <= 'F') ||
          lookahead == 'L' ||
          lookahead == 'S' ||
          ('d' <= lookahead && lookahead <= 'f') ||
          lookahead == 'l' ||
          lookahead == 's') ADVANCE(466);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(270);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 271:
      ACCEPT_TOKEN(aux_sym_num_lit_token1);
      if (('D' <= lookahead && lookahead <= 'F') ||
          lookahead == 'L' ||
          lookahead == 'S' ||
          ('d' <= lookahead && lookahead <= 'f') ||
          lookahead == 'l' ||
          lookahead == 's') ADVANCE(235);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(271);
      END_STATE();
    case 272:
      ACCEPT_TOKEN(aux_sym_num_lit_token1);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(272);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 273:
      ACCEPT_TOKEN(aux_sym_num_lit_token1);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(273);
      END_STATE();
    case 274:
      ACCEPT_TOKEN(aux_sym_num_lit_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(274);
      END_STATE();
    case 275:
      ACCEPT_TOKEN(aux_sym_num_lit_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(275);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 276:
      ACCEPT_TOKEN(aux_sym_num_lit_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(276);
      END_STATE();
    case 277:
      ACCEPT_TOKEN(aux_sym_num_lit_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(277);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 278:
      ACCEPT_TOKEN(aux_sym_num_lit_token1);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 279:
      ACCEPT_TOKEN(aux_sym_num_lit_token2);
      END_STATE();
    case 280:
      ACCEPT_TOKEN(aux_sym_num_lit_token2);
      if (lookahead == 'a') ADVANCE(391);
      if (lookahead == 'o') ADVANCE(417);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 281:
      ACCEPT_TOKEN(aux_sym_num_lit_token2);
      if (lookahead == 'e') ADVANCE(350);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 282:
      ACCEPT_TOKEN(aux_sym_num_lit_token2);
      if (lookahead == 'i') ADVANCE(405);
      if (lookahead == 'o') ADVANCE(431);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 283:
      ACCEPT_TOKEN(aux_sym_num_lit_token2);
      if (lookahead == 'i') ADVANCE(162);
      if (lookahead == 'o') ADVANCE(186);
      END_STATE();
    case 284:
      ACCEPT_TOKEN(aux_sym_num_lit_token2);
      if (lookahead == 'i') ADVANCE(162);
      if (lookahead == 'o') ADVANCE(186);
      if (lookahead == 'r') ADVANCE(174);
      END_STATE();
    case 285:
      ACCEPT_TOKEN(aux_sym_num_lit_token2);
      if (lookahead == 'o') ADVANCE(575);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 286:
      ACCEPT_TOKEN(aux_sym_num_lit_token2);
      if (lookahead == 'o') ADVANCE(574);
      END_STATE();
    case 287:
      ACCEPT_TOKEN(aux_sym_num_lit_token2);
      if (lookahead == 'o') ADVANCE(573);
      END_STATE();
    case 288:
      ACCEPT_TOKEN(aux_sym_num_lit_token2);
      if (lookahead == 'o') ADVANCE(459);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 289:
      ACCEPT_TOKEN(aux_sym_num_lit_token2);
      if (lookahead == 'o') ADVANCE(219);
      END_STATE();
    case 290:
      ACCEPT_TOKEN(aux_sym_num_lit_token2);
      if (lookahead == 'r') ADVANCE(174);
      END_STATE();
    case 291:
      ACCEPT_TOKEN(aux_sym_num_lit_token2);
      if (lookahead == 'r') ADVANCE(419);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 292:
      ACCEPT_TOKEN(aux_sym_num_lit_token2);
      if (lookahead == 'u') ADVANCE(392);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 293:
      ACCEPT_TOKEN(aux_sym_num_lit_token2);
      if (lookahead == 'u') ADVANCE(139);
      END_STATE();
    case 294:
      ACCEPT_TOKEN(aux_sym_num_lit_token2);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 295:
      ACCEPT_TOKEN(anon_sym_COLON);
      END_STATE();
    case 296:
      ACCEPT_TOKEN(anon_sym_COLON);
      if (lookahead == ':') ADVANCE(298);
      END_STATE();
    case 297:
      ACCEPT_TOKEN(anon_sym_COLON);
      if (lookahead == '@') ADVANCE(503);
      END_STATE();
    case 298:
      ACCEPT_TOKEN(anon_sym_COLON_COLON);
      END_STATE();
    case 299:
      ACCEPT_TOKEN(anon_sym_DQUOTE);
      END_STATE();
    case 300:
      ACCEPT_TOKEN(aux_sym_str_lit_token1);
      if (lookahead == '#') ADVANCE(305);
      if (lookahead != 0 &&
          lookahead != '"' &&
          lookahead != '#' &&
          lookahead != '\\' &&
          lookahead != '~') ADVANCE(304);
      END_STATE();
    case 301:
      ACCEPT_TOKEN(aux_sym_str_lit_token1);
      ADVANCE_MAP(
        '?', 302,
        'C', 305,
        'P', 305,
        '^', 305,
        '_', 305,
        'c', 305,
        'p', 305,
        '|', 304,
      );
      if (lookahead != 0 &&
          lookahead != '"' &&
          lookahead != '\\' &&
          lookahead != '~') ADVANCE(305);
      END_STATE();
    case 302:
      ACCEPT_TOKEN(aux_sym_str_lit_token1);
      if (lookahead == '@') ADVANCE(305);
      if (lookahead != 0 &&
          lookahead != '"' &&
          lookahead != '\\' &&
          lookahead != '~') ADVANCE(305);
      END_STATE();
    case 303:
      ACCEPT_TOKEN(aux_sym_str_lit_token1);
      if (lookahead == '|') ADVANCE(304);
      if (lookahead != 0 &&
          lookahead != '"' &&
          lookahead != '\\' &&
          lookahead != '~') ADVANCE(305);
      END_STATE();
    case 304:
      ACCEPT_TOKEN(aux_sym_str_lit_token1);
      if (lookahead == '|') ADVANCE(300);
      if (lookahead != 0 &&
          lookahead != '"' &&
          lookahead != '\\' &&
          lookahead != '~') ADVANCE(304);
      END_STATE();
    case 305:
      ACCEPT_TOKEN(aux_sym_str_lit_token1);
      if (lookahead != 0 &&
          lookahead != '"' &&
          lookahead != '\\' &&
          lookahead != '~') ADVANCE(305);
      END_STATE();
    case 306:
      ACCEPT_TOKEN(aux_sym_str_lit_token2);
      END_STATE();
    case 307:
      ACCEPT_TOKEN(anon_sym_TILDE);
      END_STATE();
    case 308:
      ACCEPT_TOKEN(aux_sym_char_lit_token1);
      END_STATE();
    case 309:
      ACCEPT_TOKEN(aux_sym_char_lit_token1);
      if (lookahead != 0 &&
          lookahead != '\t' &&
          lookahead != '\n' &&
          lookahead != '\f' &&
          lookahead != '\r' &&
          lookahead != ' ' &&
          lookahead != '(' &&
          lookahead != ')') ADVANCE(309);
      END_STATE();
    case 310:
      ACCEPT_TOKEN(sym_nil_lit);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 311:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == '#') ADVANCE(470);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(264);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 312:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'a') ADVANCE(460);
      if (lookahead == 'i') ADVANCE(407);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 313:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'a') ADVANCE(463);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 314:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'a') ADVANCE(613);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 315:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'a') ADVANCE(448);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 316:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'a') ADVANCE(326);
      if (lookahead == 'e') ADVANCE(450);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 317:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'a') ADVANCE(384);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 318:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'a') ADVANCE(391);
      if (lookahead == 'o') ADVANCE(417);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 319:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'a') ADVANCE(386);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 320:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'b') ADVANCE(329);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 321:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'b') ADVANCE(425);
      if (lookahead == 'c') ADVANCE(434);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 322:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'c') ADVANCE(610);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 323:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'c') ADVANCE(422);
      if (lookahead == 'e') ADVANCE(456);
      if (lookahead == 'i') ADVANCE(376);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 324:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'c') ADVANCE(369);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 325:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'c') ADVANCE(452);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 326:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'c') ADVANCE(440);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 327:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'd') ADVANCE(568);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 328:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'd') ADVANCE(612);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 329:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'd') ADVANCE(314);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 330:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'd') ADVANCE(369);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 331:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'e') ADVANCE(477);
      if (lookahead == 'i') ADVANCE(399);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 332:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'e') ADVANCE(430);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 333:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'e') ADVANCE(595);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 334:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'e') ADVANCE(577);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 335:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'e') ADVANCE(370);
      if (lookahead == 'y') ADVANCE(557);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 336:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'e') ADVANCE(541);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 337:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'e') ADVANCE(350);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 338:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'e') ADVANCE(446);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 339:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'e') ADVANCE(436);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 340:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'e') ADVANCE(325);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 341:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'e') ADVANCE(432);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 342:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'e') ADVANCE(315);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 343:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'e') ADVANCE(397);
      if (lookahead == 'i') ADVANCE(387);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 344:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'e') ADVANCE(367);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 345:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'e') ADVANCE(404);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 346:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'e') ADVANCE(435);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 347:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'e') ADVANCE(402);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 348:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'e') ADVANCE(411);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 349:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'f') ADVANCE(585);
      if (lookahead == 'n') ADVANCE(365);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 350:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'f') ADVANCE(356);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 351:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'f') ADVANCE(438);
      if (lookahead == 't') ADVANCE(413);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 352:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'f') ADVANCE(439);
      if (lookahead == 't') ADVANCE(414);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 353:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'g') ADVANCE(477);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 354:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'g') ADVANCE(534);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 355:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'g') ADVANCE(536);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 356:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'g') ADVANCE(348);
      if (lookahead == 'm') ADVANCE(316);
      if (lookahead == 'u') ADVANCE(403);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 357:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'h') ADVANCE(343);
      if (lookahead == 'i') ADVANCE(449);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 358:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'h') ADVANCE(572);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 359:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'h') ADVANCE(339);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 360:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'h') ADVANCE(420);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 361:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'h') ADVANCE(347);
      if (lookahead == 'o') ADVANCE(545);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 362:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'i') ADVANCE(464);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 363:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'i') ADVANCE(376);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 364:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'i') ADVANCE(405);
      if (lookahead == 'o') ADVANCE(431);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 365:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'i') ADVANCE(451);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 366:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'i') ADVANCE(322);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 367:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'i') ADVANCE(443);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 368:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'i') ADVANCE(379);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 369:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'i') ADVANCE(399);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 370:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'i') ADVANCE(406);
      if (lookahead == 'l') ADVANCE(418);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 371:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'i') ADVANCE(408);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 372:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'i') ADVANCE(394);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 373:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'i') ADVANCE(319);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 374:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'l') ADVANCE(527);
      if (lookahead == 'o') ADVANCE(382);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 375:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'l') ADVANCE(527);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 376:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'l') ADVANCE(310);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 377:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'l') ADVANCE(457);
      if (lookahead == 'n') ADVANCE(327);
      if (lookahead == 'p') ADVANCE(429);
      if (lookahead == 's') ADVANCE(570);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 378:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'l') ADVANCE(445);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 379:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'l') ADVANCE(579);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 380:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'l') ADVANCE(461);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 381:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'l') ADVANCE(462);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 382:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'l') ADVANCE(385);
      if (lookahead == 'u') ADVANCE(410);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 383:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'l') ADVANCE(338);
      if (lookahead == 't') ADVANCE(368);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 384:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'l') ADVANCE(380);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 385:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'l') ADVANCE(340);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 386:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'l') ADVANCE(381);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 387:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'l') ADVANCE(334);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 388:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'm') ADVANCE(543);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 389:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'm') ADVANCE(549);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 390:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'm') ADVANCE(553);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 391:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'm') ADVANCE(320);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 392:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'm') ADVANCE(393);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 393:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'm') ADVANCE(369);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 394:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'm') ADVANCE(362);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 395:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'n') ADVANCE(383);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 396:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'n') ADVANCE(324);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 397:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'n') ADVANCE(583);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 398:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'n') ADVANCE(600);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 399:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'n') ADVANCE(353);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 400:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'n') ADVANCE(530);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 401:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'n') ADVANCE(555);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 402:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'n') ADVANCE(559);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 403:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'n') ADVANCE(606);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 404:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'n') ADVANCE(330);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 405:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'n') ADVANCE(317);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 406:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'n') ADVANCE(354);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 407:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'n') ADVANCE(372);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 408:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'n') ADVANCE(355);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 409:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'n') ADVANCE(352);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 410:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'n') ADVANCE(452);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 411:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'n') ADVANCE(346);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 412:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'o') ADVANCE(575);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 413:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'o') ADVANCE(547);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 414:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'o') ADVANCE(551);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 415:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'o') ADVANCE(608);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 416:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'o') ADVANCE(459);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 417:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'o') ADVANCE(428);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 418:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'o') ADVANCE(458);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 419:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'o') ADVANCE(388);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 420:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'o') ADVANCE(328);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 421:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'o') ADVANCE(389);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 422:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'o') ADVANCE(396);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 423:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'o') ADVANCE(390);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 424:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'o') ADVANCE(417);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 425:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'o') ADVANCE(455);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 426:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'o') ADVANCE(447);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 427:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'p') ADVANCE(351);
      if (lookahead == 's') ADVANCE(371);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 428:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'p') ADVANCE(604);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 429:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'p') ADVANCE(345);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 430:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'p') ADVANCE(342);
      if (lookahead == 't') ADVANCE(454);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 431:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'r') ADVANCE(566);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 432:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'r') ADVANCE(593);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 433:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'r') ADVANCE(419);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 434:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'r') ADVANCE(426);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 435:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'r') ADVANCE(366);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 436:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'r') ADVANCE(344);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 437:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'r') ADVANCE(398);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 438:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'r') ADVANCE(421);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 439:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'r') ADVANCE(423);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 440:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'r') ADVANCE(415);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 441:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 's') ADVANCE(589);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 442:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 's') ADVANCE(587);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 443:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 's') ADVANCE(591);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 444:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 's') ADVANCE(532);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 445:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 's') ADVANCE(333);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 446:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 's') ADVANCE(442);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 447:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 's') ADVANCE(444);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 448:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 't') ADVANCE(581);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 449:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 't') ADVANCE(358);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 450:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 't') ADVANCE(360);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 451:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 't') ADVANCE(373);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 452:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 't') ADVANCE(369);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 453:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'u') ADVANCE(392);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 454:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'u') ADVANCE(437);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 455:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'v') ADVANCE(336);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 456:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'v') ADVANCE(341);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 457:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'w') ADVANCE(313);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 458:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'w') ADVANCE(539);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 459:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'w') ADVANCE(409);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 460:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'x') ADVANCE(372);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 461:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'y') ADVANCE(598);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 462:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'y') ADVANCE(602);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 463:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'y') ADVANCE(441);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 464:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'z') ADVANCE(331);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 465:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == '|') ADVANCE(477);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          (0x1c <= lookahead && lookahead <= ' ') ||
          lookahead == '"' ||
          lookahead == '(' ||
          lookahead == ')' ||
          lookahead == ',' ||
          lookahead == ':' ||
          lookahead == ';' ||
          ('[' <= lookahead && lookahead <= '^') ||
          lookahead == '`' ||
          ('{' <= lookahead && lookahead <= '}') ||
          lookahead == 0x1680 ||
          (0x2000 <= lookahead && lookahead <= 0x2006) ||
          (0x2008 <= lookahead && lookahead <= 0x200a) ||
          lookahead == 0x2028 ||
          lookahead == 0x2029 ||
          lookahead == 0x205f ||
          lookahead == 0x3000) ADVANCE(230);
      if (lookahead != 0) ADVANCE(465);
      END_STATE();
    case 466:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == '+' ||
          lookahead == '-') ADVANCE(474);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(272);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 467:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == '+' ||
          lookahead == '-') ADVANCE(472);
      if (lookahead == '0' ||
          lookahead == '1') ADVANCE(267);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 468:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == '+' ||
          lookahead == '-') ADVANCE(473);
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(269);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 469:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == '+' ||
          lookahead == '-') ADVANCE(475);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(275);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 470:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'B' ||
          lookahead == 'b') ADVANCE(467);
      if (lookahead == 'O' ||
          lookahead == 'o') ADVANCE(468);
      if (lookahead == 'X' ||
          lookahead == 'x') ADVANCE(469);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(471);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 471:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == 'R' ||
          lookahead == 'r') ADVANCE(476);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(471);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 472:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (lookahead == '0' ||
          lookahead == '1') ADVANCE(267);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 473:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (('0' <= lookahead && lookahead <= '7')) ADVANCE(269);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 474:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(272);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 475:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'F') ||
          ('a' <= lookahead && lookahead <= 'f')) ADVANCE(275);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 476:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(277);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 477:
      ACCEPT_TOKEN(aux_sym_sym_lit_token1);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 478:
      ACCEPT_TOKEN(anon_sym_CARET);
      END_STATE();
    case 479:
      ACCEPT_TOKEN(anon_sym_POUND_CARET);
      END_STATE();
    case 480:
      ACCEPT_TOKEN(anon_sym_LPAREN);
      END_STATE();
    case 481:
      ACCEPT_TOKEN(anon_sym_RPAREN);
      END_STATE();
    case 482:
      ACCEPT_TOKEN(anon_sym_LBRACE);
      END_STATE();
    case 483:
      ACCEPT_TOKEN(anon_sym_RBRACE);
      END_STATE();
    case 484:
      ACCEPT_TOKEN(anon_sym_POUND0A);
      END_STATE();
    case 485:
      ACCEPT_TOKEN(anon_sym_POUND0a);
      END_STATE();
    case 486:
      ACCEPT_TOKEN(anon_sym_POUND_QMARK);
      END_STATE();
    case 487:
      ACCEPT_TOKEN(anon_sym_POUND_QMARK);
      if (lookahead == '@') ADVANCE(488);
      END_STATE();
    case 488:
      ACCEPT_TOKEN(anon_sym_POUND_QMARK_AT);
      END_STATE();
    case 489:
      ACCEPT_TOKEN(anon_sym_POUND_SQUOTE);
      END_STATE();
    case 490:
      ACCEPT_TOKEN(anon_sym_SQUOTE);
      END_STATE();
    case 491:
      ACCEPT_TOKEN(anon_sym_BQUOTE);
      END_STATE();
    case 492:
      ACCEPT_TOKEN(anon_sym_COMMA_AT);
      END_STATE();
    case 493:
      ACCEPT_TOKEN(anon_sym_COMMA);
      END_STATE();
    case 494:
      ACCEPT_TOKEN(anon_sym_COMMA);
      if (lookahead == '@') ADVANCE(492);
      END_STATE();
    case 495:
      ACCEPT_TOKEN(sym_block_comment);
      END_STATE();
    case 496:
      ACCEPT_TOKEN(sym_fancy_literal);
      END_STATE();
    case 497:
      ACCEPT_TOKEN(aux_sym__format_token_token1);
      END_STATE();
    case 498:
      ACCEPT_TOKEN(aux_sym__format_token_token1);
      if (lookahead == '|') ADVANCE(229);
      END_STATE();
    case 499:
      ACCEPT_TOKEN(anon_sym_v);
      END_STATE();
    case 500:
      ACCEPT_TOKEN(anon_sym_V);
      END_STATE();
    case 501:
      ACCEPT_TOKEN(anon_sym_AT);
      if (lookahead == ':') ADVANCE(502);
      END_STATE();
    case 502:
      ACCEPT_TOKEN(anon_sym_AT_COLON);
      END_STATE();
    case 503:
      ACCEPT_TOKEN(anon_sym_COLON_AT);
      END_STATE();
    case 504:
      ACCEPT_TOKEN(anon_sym_PERCENT);
      END_STATE();
    case 505:
      ACCEPT_TOKEN(anon_sym_AMP);
      END_STATE();
    case 506:
      ACCEPT_TOKEN(anon_sym_PIPE);
      END_STATE();
    case 507:
      ACCEPT_TOKEN(aux_sym_format_directive_type_token1);
      END_STATE();
    case 508:
      ACCEPT_TOKEN(aux_sym_format_directive_type_token2);
      END_STATE();
    case 509:
      ACCEPT_TOKEN(anon_sym_LF);
      END_STATE();
    case 510:
      ACCEPT_TOKEN(anon_sym_CR);
      END_STATE();
    case 511:
      ACCEPT_TOKEN(aux_sym_format_directive_type_token3);
      END_STATE();
    case 512:
      ACCEPT_TOKEN(aux_sym_format_directive_type_token4);
      END_STATE();
    case 513:
      ACCEPT_TOKEN(aux_sym_format_directive_type_token5);
      END_STATE();
    case 514:
      ACCEPT_TOKEN(aux_sym_format_directive_type_token6);
      END_STATE();
    case 515:
      ACCEPT_TOKEN(anon_sym__);
      END_STATE();
    case 516:
      ACCEPT_TOKEN(aux_sym_format_directive_type_token7);
      END_STATE();
    case 517:
      ACCEPT_TOKEN(aux_sym_format_directive_type_token8);
      END_STATE();
    case 518:
      ACCEPT_TOKEN(aux_sym_format_directive_type_token9);
      END_STATE();
    case 519:
      ACCEPT_TOKEN(aux_sym_format_directive_type_token10);
      END_STATE();
    case 520:
      ACCEPT_TOKEN(anon_sym_SEMI);
      END_STATE();
    case 521:
      ACCEPT_TOKEN(anon_sym_STAR);
      END_STATE();
    case 522:
      ACCEPT_TOKEN(anon_sym_SLASH);
      END_STATE();
    case 523:
      ACCEPT_TOKEN(anon_sym_QMARK);
      END_STATE();
    case 524:
      ACCEPT_TOKEN(anon_sym_Newline);
      END_STATE();
    case 525:
      ACCEPT_TOKEN(aux_sym_format_directive_type_token11);
      END_STATE();
    case 526:
      ACCEPT_TOKEN(anon_sym_cl);
      END_STATE();
    case 527:
      ACCEPT_TOKEN(anon_sym_cl);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 528:
      ACCEPT_TOKEN(anon_sym_in);
      END_STATE();
    case 529:
      ACCEPT_TOKEN(anon_sym_in);
      if (lookahead == 'i') ADVANCE(210);
      END_STATE();
    case 530:
      ACCEPT_TOKEN(anon_sym_in);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 531:
      ACCEPT_TOKEN(anon_sym_across);
      END_STATE();
    case 532:
      ACCEPT_TOKEN(anon_sym_across);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 533:
      ACCEPT_TOKEN(anon_sym_being);
      if (lookahead == ' ') ADVANCE(62);
      END_STATE();
    case 534:
      ACCEPT_TOKEN(anon_sym_being);
      if (lookahead == ' ') ADVANCE(62);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 535:
      ACCEPT_TOKEN(anon_sym_using);
      END_STATE();
    case 536:
      ACCEPT_TOKEN(anon_sym_using);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 537:
      ACCEPT_TOKEN(aux_sym_for_clause_word_token1);
      END_STATE();
    case 538:
      ACCEPT_TOKEN(anon_sym_below);
      END_STATE();
    case 539:
      ACCEPT_TOKEN(anon_sym_below);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 540:
      ACCEPT_TOKEN(anon_sym_above);
      END_STATE();
    case 541:
      ACCEPT_TOKEN(anon_sym_above);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 542:
      ACCEPT_TOKEN(anon_sym_from);
      END_STATE();
    case 543:
      ACCEPT_TOKEN(anon_sym_from);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 544:
      ACCEPT_TOKEN(anon_sym_to);
      END_STATE();
    case 545:
      ACCEPT_TOKEN(anon_sym_to);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 546:
      ACCEPT_TOKEN(anon_sym_upto);
      END_STATE();
    case 547:
      ACCEPT_TOKEN(anon_sym_upto);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 548:
      ACCEPT_TOKEN(anon_sym_upfrom);
      END_STATE();
    case 549:
      ACCEPT_TOKEN(anon_sym_upfrom);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 550:
      ACCEPT_TOKEN(anon_sym_downto);
      END_STATE();
    case 551:
      ACCEPT_TOKEN(anon_sym_downto);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 552:
      ACCEPT_TOKEN(anon_sym_downfrom);
      END_STATE();
    case 553:
      ACCEPT_TOKEN(anon_sym_downfrom);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 554:
      ACCEPT_TOKEN(anon_sym_on);
      END_STATE();
    case 555:
      ACCEPT_TOKEN(anon_sym_on);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 556:
      ACCEPT_TOKEN(anon_sym_by);
      END_STATE();
    case 557:
      ACCEPT_TOKEN(anon_sym_by);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 558:
      ACCEPT_TOKEN(anon_sym_then);
      END_STATE();
    case 559:
      ACCEPT_TOKEN(anon_sym_then);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 560:
      ACCEPT_TOKEN(anon_sym_EQ);
      END_STATE();
    case 561:
      ACCEPT_TOKEN(anon_sym_EQ);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 562:
      ACCEPT_TOKEN(aux_sym_accumulation_verb_token1);
      END_STATE();
    case 563:
      ACCEPT_TOKEN(aux_sym_accumulation_verb_token1);
      if (lookahead == 'i') ADVANCE(160);
      END_STATE();
    case 564:
      ACCEPT_TOKEN(aux_sym_accumulation_verb_token1);
      if (lookahead == 'm') ADVANCE(118);
      END_STATE();
    case 565:
      ACCEPT_TOKEN(anon_sym_for);
      END_STATE();
    case 566:
      ACCEPT_TOKEN(anon_sym_for);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 567:
      ACCEPT_TOKEN(anon_sym_and);
      END_STATE();
    case 568:
      ACCEPT_TOKEN(anon_sym_and);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 569:
      ACCEPT_TOKEN(anon_sym_as);
      END_STATE();
    case 570:
      ACCEPT_TOKEN(anon_sym_as);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 571:
      ACCEPT_TOKEN(anon_sym_with);
      END_STATE();
    case 572:
      ACCEPT_TOKEN(anon_sym_with);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 573:
      ACCEPT_TOKEN(anon_sym_do);
      END_STATE();
    case 574:
      ACCEPT_TOKEN(anon_sym_do);
      if (lookahead == 'w') ADVANCE(166);
      END_STATE();
    case 575:
      ACCEPT_TOKEN(anon_sym_do);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 576:
      ACCEPT_TOKEN(anon_sym_while);
      END_STATE();
    case 577:
      ACCEPT_TOKEN(anon_sym_while);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 578:
      ACCEPT_TOKEN(anon_sym_until);
      END_STATE();
    case 579:
      ACCEPT_TOKEN(anon_sym_until);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 580:
      ACCEPT_TOKEN(anon_sym_repeat);
      END_STATE();
    case 581:
      ACCEPT_TOKEN(anon_sym_repeat);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 582:
      ACCEPT_TOKEN(anon_sym_when);
      END_STATE();
    case 583:
      ACCEPT_TOKEN(anon_sym_when);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 584:
      ACCEPT_TOKEN(anon_sym_if);
      END_STATE();
    case 585:
      ACCEPT_TOKEN(anon_sym_if);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 586:
      ACCEPT_TOKEN(anon_sym_unless);
      END_STATE();
    case 587:
      ACCEPT_TOKEN(anon_sym_unless);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 588:
      ACCEPT_TOKEN(anon_sym_always);
      END_STATE();
    case 589:
      ACCEPT_TOKEN(anon_sym_always);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 590:
      ACCEPT_TOKEN(anon_sym_thereis);
      END_STATE();
    case 591:
      ACCEPT_TOKEN(anon_sym_thereis);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 592:
      ACCEPT_TOKEN(anon_sym_never);
      END_STATE();
    case 593:
      ACCEPT_TOKEN(anon_sym_never);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 594:
      ACCEPT_TOKEN(anon_sym_else);
      END_STATE();
    case 595:
      ACCEPT_TOKEN(anon_sym_else);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 596:
      ACCEPT_TOKEN(anon_sym_into);
      END_STATE();
    case 597:
      ACCEPT_TOKEN(anon_sym_finally);
      END_STATE();
    case 598:
      ACCEPT_TOKEN(anon_sym_finally);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 599:
      ACCEPT_TOKEN(anon_sym_return);
      END_STATE();
    case 600:
      ACCEPT_TOKEN(anon_sym_return);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 601:
      ACCEPT_TOKEN(anon_sym_initially);
      END_STATE();
    case 602:
      ACCEPT_TOKEN(anon_sym_initially);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 603:
      ACCEPT_TOKEN(anon_sym_loop);
      END_STATE();
    case 604:
      ACCEPT_TOKEN(anon_sym_loop);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 605:
      ACCEPT_TOKEN(anon_sym_defun);
      END_STATE();
    case 606:
      ACCEPT_TOKEN(anon_sym_defun);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 607:
      ACCEPT_TOKEN(anon_sym_defmacro);
      END_STATE();
    case 608:
      ACCEPT_TOKEN(anon_sym_defmacro);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 609:
      ACCEPT_TOKEN(anon_sym_defgeneric);
      END_STATE();
    case 610:
      ACCEPT_TOKEN(anon_sym_defgeneric);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 611:
      ACCEPT_TOKEN(anon_sym_defmethod);
      END_STATE();
    case 612:
      ACCEPT_TOKEN(anon_sym_defmethod);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 613:
      ACCEPT_TOKEN(anon_sym_lambda);
      if ((!eof && set_contains(aux_sym_sym_lit_token1_character_set_2, 17, lookahead))) ADVANCE(477);
      END_STATE();
    case 614:
      ACCEPT_TOKEN(anon_sym_POUNDP);
      END_STATE();
    case 615:
      ACCEPT_TOKEN(anon_sym_POUNDp);
      END_STATE();
    case 616:
      ACCEPT_TOKEN(aux_sym_path_lit_token1);
      END_STATE();
    case 617:
      ACCEPT_TOKEN(aux_sym__sym_lit_without_slash_token1);
      END_STATE();
    case 618:
      ACCEPT_TOKEN(aux_sym__sym_lit_without_slash_token1);
      if (lookahead == 'l') ADVANCE(526);
      END_STATE();
    case 619:
      ACCEPT_TOKEN(sym_self_referential_reader_macro);
      END_STATE();
    case 620:
      ACCEPT_TOKEN(anon_sym_POUND_PLUS);
      END_STATE();
    case 621:
      ACCEPT_TOKEN(anon_sym_POUND_DASH);
      END_STATE();
    case 622:
      ACCEPT_TOKEN(anon_sym_POUNDC);
      END_STATE();
    case 623:
      ACCEPT_TOKEN(anon_sym_POUNDc);
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
static uint16_t ts_small_parse_table[199129];

// Filled from tables.zlib by grammar.go, see fetchgrammar.
static uint32_t ts_small_parse_table_map[3760];

// Filled from tables.zlib by grammar.go, see fetchgrammar.
static TSParseActionEntry ts_parse_actions[13230];

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

TS_PUBLIC const TSLanguage *tree_sitter_commonlisp(void) {
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
