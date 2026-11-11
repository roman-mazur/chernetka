#include "tree_sitter/parser.h"

#if defined(__GNUC__) || defined(__clang__)
#pragma GCC diagnostic ignored "-Wmissing-field-initializers"
#endif

#define LANGUAGE_VERSION 14
#define STATE_COUNT 1286
#define LARGE_STATE_COUNT 8
#define SYMBOL_COUNT 177
#define ALIAS_COUNT 5
#define TOKEN_COUNT 111
#define EXTERNAL_TOKEN_COUNT 0
#define FIELD_COUNT 24
#define MAX_ALIAS_SEQUENCE_LENGTH 9
#define PRODUCTION_ID_COUNT 80

enum ts_symbol_identifiers {
  sym_word = 1,
  aux_sym__thing_token1 = 2,
  anon_sym_COLON = 3,
  anon_sym_AMP_COLON = 4,
  anon_sym_COLON_COLON = 5,
  aux_sym__ordinary_rule_token1 = 6,
  anon_sym_PIPE = 7,
  anon_sym_SEMI = 8,
  anon_sym_AT = 9,
  anon_sym_DASH = 10,
  anon_sym_PLUS = 11,
  anon_sym_VPATH = 12,
  anon_sym_EQ = 13,
  anon_sym_COLON_EQ = 14,
  anon_sym_COLON_COLON_EQ = 15,
  anon_sym_QMARK_EQ = 16,
  anon_sym_PLUS_EQ = 17,
  anon_sym_DOTRECIPEPREFIX = 18,
  anon_sym_BANG_EQ = 19,
  anon_sym_define = 20,
  anon_sym_endef = 21,
  anon_sym_include = 22,
  anon_sym_sinclude = 23,
  anon_sym_DASHinclude = 24,
  anon_sym_vpath = 25,
  anon_sym_export = 26,
  anon_sym_unexport = 27,
  anon_sym_override = 28,
  anon_sym_undefine = 29,
  anon_sym_private = 30,
  anon_sym_endif = 31,
  anon_sym_else = 32,
  anon_sym_ifeq = 33,
  anon_sym_ifneq = 34,
  anon_sym_ifdef = 35,
  anon_sym_ifndef = 36,
  anon_sym_LPAREN = 37,
  anon_sym_COMMA = 38,
  anon_sym_RPAREN = 39,
  anon_sym_DOLLAR = 40,
  anon_sym_DOLLAR_DOLLAR = 41,
  anon_sym_LPAREN2 = 42,
  anon_sym_LBRACE = 43,
  anon_sym_RBRACE = 44,
  aux_sym_variable_reference_token1 = 45,
  anon_sym_AT2 = 46,
  anon_sym_PERCENT = 47,
  anon_sym_LT = 48,
  anon_sym_QMARK = 49,
  anon_sym_CARET = 50,
  anon_sym_PLUS2 = 51,
  anon_sym_SLASH = 52,
  anon_sym_STAR = 53,
  anon_sym_PERCENT2 = 54,
  anon_sym_LT2 = 55,
  anon_sym_QMARK2 = 56,
  anon_sym_CARET2 = 57,
  anon_sym_SLASH2 = 58,
  anon_sym_STAR2 = 59,
  anon_sym_D = 60,
  anon_sym_F = 61,
  anon_sym_subst = 62,
  anon_sym_patsubst = 63,
  anon_sym_strip = 64,
  anon_sym_findstring = 65,
  anon_sym_filter = 66,
  anon_sym_filter_DASHout = 67,
  anon_sym_sort = 68,
  anon_sym_word = 69,
  anon_sym_words = 70,
  anon_sym_wordlist = 71,
  anon_sym_firstword = 72,
  anon_sym_lastword = 73,
  anon_sym_dir = 74,
  anon_sym_notdir = 75,
  anon_sym_suffix = 76,
  anon_sym_basename = 77,
  anon_sym_addsuffix = 78,
  anon_sym_addprefix = 79,
  anon_sym_join = 80,
  anon_sym_wildcard = 81,
  anon_sym_realpath = 82,
  anon_sym_abspath = 83,
  anon_sym_error = 84,
  anon_sym_warning = 85,
  anon_sym_info = 86,
  anon_sym_origin = 87,
  anon_sym_flavor = 88,
  anon_sym_foreach = 89,
  anon_sym_if = 90,
  anon_sym_or = 91,
  anon_sym_and = 92,
  anon_sym_call = 93,
  anon_sym_eval = 94,
  anon_sym_file = 95,
  anon_sym_value = 96,
  anon_sym_shell = 97,
  aux_sym_list_token1 = 98,
  anon_sym_COLON2 = 99,
  anon_sym_SEMI2 = 100,
  anon_sym_DQUOTE = 101,
  anon_sym_SQUOTE = 102,
  aux_sym__string_token1 = 103,
  anon_sym_RPAREN2 = 104,
  anon_sym_TAB = 105,
  sym__rawline = 106,
  aux_sym__shell_text_without_split_token1 = 107,
  anon_sym_SLASH_SLASH = 108,
  aux_sym_text_token1 = 109,
  sym_comment = 110,
  sym_makefile = 111,
  sym__thing = 112,
  sym_rule = 113,
  sym__ordinary_rule = 114,
  sym__static_pattern_rule = 115,
  sym__normal_prerequisites = 116,
  sym_recipe = 117,
  sym__attached_recipe_line = 118,
  sym__prefixed_recipe_line = 119,
  sym_recipe_line = 120,
  sym__variable_definition = 121,
  sym_VPATH_assignment = 122,
  sym_RECIPEPREFIX_assignment = 123,
  sym_variable_assignment = 124,
  sym_shell_assignment = 125,
  sym_define_directive = 126,
  sym__directive = 127,
  sym_include_directive = 128,
  sym_vpath_directive = 129,
  sym_export_directive = 130,
  sym_unexport_directive = 131,
  sym_override_directive = 132,
  sym_undefine_directive = 133,
  sym_private_directive = 134,
  sym_conditional = 135,
  sym_elsif_directive = 136,
  sym_else_directive = 137,
  sym__conditional_directives = 138,
  aux_sym__conditional_consequence = 139,
  sym_ifeq_directive = 140,
  sym_ifneq_directive = 141,
  sym_ifdef_directive = 142,
  sym_ifndef_directive = 143,
  sym__conditional_args_cmp = 144,
  sym__variable = 145,
  sym_variable_reference = 146,
  sym_substitution_reference = 147,
  sym_automatic_variable = 148,
  sym__function = 149,
  sym_function_call = 150,
  sym_arguments = 151,
  sym_shell_function = 152,
  sym_list = 153,
  sym_paths = 154,
  sym_concatenation = 155,
  sym_string = 156,
  aux_sym__string = 157,
  sym_archive = 158,
  sym__recipeprefix = 159,
  sym__shell_text_without_split = 160,
  sym_shell_text_with_split = 161,
  sym__shell_command = 162,
  sym_text = 163,
  aux_sym_makefile_repeat1 = 164,
  aux_sym_recipe_repeat1 = 165,
  aux_sym_recipe_line_repeat1 = 166,
  aux_sym_define_directive_repeat1 = 167,
  aux_sym_conditional_repeat1 = 168,
  aux_sym_arguments_repeat1 = 169,
  aux_sym_list_repeat1 = 170,
  aux_sym_paths_repeat1 = 171,
  aux_sym_concatenation_repeat1 = 172,
  aux_sym__shell_text_without_split_repeat1 = 173,
  aux_sym__shell_text_without_split_repeat2 = 174,
  aux_sym_text_repeat1 = 175,
  aux_sym_text_repeat2 = 176,
  alias_sym_pattern_list = 177,
  alias_sym_prerequisites = 178,
  alias_sym_raw_text = 179,
  alias_sym_shell_command = 180,
  alias_sym_targets = 181,
};

static const char * const ts_symbol_names[] = {
  [ts_builtin_sym_end] = "end",
  [sym_word] = "word",
  [aux_sym__thing_token1] = "_thing_token1",
  [anon_sym_COLON] = ":",
  [anon_sym_AMP_COLON] = "&:",
  [anon_sym_COLON_COLON] = "::",
  [aux_sym__ordinary_rule_token1] = "_ordinary_rule_token1",
  [anon_sym_PIPE] = "|",
  [anon_sym_SEMI] = ";",
  [anon_sym_AT] = "@",
  [anon_sym_DASH] = "-",
  [anon_sym_PLUS] = "+",
  [anon_sym_VPATH] = "VPATH",
  [anon_sym_EQ] = "=",
  [anon_sym_COLON_EQ] = ":=",
  [anon_sym_COLON_COLON_EQ] = "::=",
  [anon_sym_QMARK_EQ] = "\?=",
  [anon_sym_PLUS_EQ] = "+=",
  [anon_sym_DOTRECIPEPREFIX] = ".RECIPEPREFIX",
  [anon_sym_BANG_EQ] = "!=",
  [anon_sym_define] = "define",
  [anon_sym_endef] = "endef",
  [anon_sym_include] = "include",
  [anon_sym_sinclude] = "sinclude",
  [anon_sym_DASHinclude] = "-include",
  [anon_sym_vpath] = "vpath",
  [anon_sym_export] = "export",
  [anon_sym_unexport] = "unexport",
  [anon_sym_override] = "override",
  [anon_sym_undefine] = "undefine",
  [anon_sym_private] = "private",
  [anon_sym_endif] = "endif",
  [anon_sym_else] = "else",
  [anon_sym_ifeq] = "ifeq",
  [anon_sym_ifneq] = "ifneq",
  [anon_sym_ifdef] = "ifdef",
  [anon_sym_ifndef] = "ifndef",
  [anon_sym_LPAREN] = "(",
  [anon_sym_COMMA] = ",",
  [anon_sym_RPAREN] = ")",
  [anon_sym_DOLLAR] = "$",
  [anon_sym_DOLLAR_DOLLAR] = "$$",
  [anon_sym_LPAREN2] = "(",
  [anon_sym_LBRACE] = "{",
  [anon_sym_RBRACE] = "}",
  [aux_sym_variable_reference_token1] = "word",
  [anon_sym_AT2] = "@",
  [anon_sym_PERCENT] = "%",
  [anon_sym_LT] = "<",
  [anon_sym_QMARK] = "\?",
  [anon_sym_CARET] = "^",
  [anon_sym_PLUS2] = "+",
  [anon_sym_SLASH] = "/",
  [anon_sym_STAR] = "*",
  [anon_sym_PERCENT2] = "%",
  [anon_sym_LT2] = "<",
  [anon_sym_QMARK2] = "\?",
  [anon_sym_CARET2] = "^",
  [anon_sym_SLASH2] = "/",
  [anon_sym_STAR2] = "*",
  [anon_sym_D] = "D",
  [anon_sym_F] = "F",
  [anon_sym_subst] = "subst",
  [anon_sym_patsubst] = "patsubst",
  [anon_sym_strip] = "strip",
  [anon_sym_findstring] = "findstring",
  [anon_sym_filter] = "filter",
  [anon_sym_filter_DASHout] = "filter-out",
  [anon_sym_sort] = "sort",
  [anon_sym_word] = "word",
  [anon_sym_words] = "words",
  [anon_sym_wordlist] = "wordlist",
  [anon_sym_firstword] = "firstword",
  [anon_sym_lastword] = "lastword",
  [anon_sym_dir] = "dir",
  [anon_sym_notdir] = "notdir",
  [anon_sym_suffix] = "suffix",
  [anon_sym_basename] = "basename",
  [anon_sym_addsuffix] = "addsuffix",
  [anon_sym_addprefix] = "addprefix",
  [anon_sym_join] = "join",
  [anon_sym_wildcard] = "wildcard",
  [anon_sym_realpath] = "realpath",
  [anon_sym_abspath] = "abspath",
  [anon_sym_error] = "error",
  [anon_sym_warning] = "warning",
  [anon_sym_info] = "info",
  [anon_sym_origin] = "origin",
  [anon_sym_flavor] = "flavor",
  [anon_sym_foreach] = "foreach",
  [anon_sym_if] = "if",
  [anon_sym_or] = "or",
  [anon_sym_and] = "and",
  [anon_sym_call] = "call",
  [anon_sym_eval] = "eval",
  [anon_sym_file] = "file",
  [anon_sym_value] = "value",
  [anon_sym_shell] = "shell",
  [aux_sym_list_token1] = "\\",
  [anon_sym_COLON2] = ":",
  [anon_sym_SEMI2] = ";",
  [anon_sym_DQUOTE] = "\"",
  [anon_sym_SQUOTE] = "'",
  [aux_sym__string_token1] = "_string_token1",
  [anon_sym_RPAREN2] = ")",
  [anon_sym_TAB] = "\t",
  [sym__rawline] = "_rawline",
  [aux_sym__shell_text_without_split_token1] = "_shell_text_without_split_token1",
  [anon_sym_SLASH_SLASH] = "escape",
  [aux_sym_text_token1] = "text_token1",
  [sym_comment] = "comment",
  [sym_makefile] = "makefile",
  [sym__thing] = "_thing",
  [sym_rule] = "rule",
  [sym__ordinary_rule] = "_ordinary_rule",
  [sym__static_pattern_rule] = "_static_pattern_rule",
  [sym__normal_prerequisites] = "_normal_prerequisites",
  [sym_recipe] = "recipe",
  [sym__attached_recipe_line] = "_attached_recipe_line",
  [sym__prefixed_recipe_line] = "_prefixed_recipe_line",
  [sym_recipe_line] = "recipe_line",
  [sym__variable_definition] = "_variable_definition",
  [sym_VPATH_assignment] = "VPATH_assignment",
  [sym_RECIPEPREFIX_assignment] = "RECIPEPREFIX_assignment",
  [sym_variable_assignment] = "variable_assignment",
  [sym_shell_assignment] = "shell_assignment",
  [sym_define_directive] = "define_directive",
  [sym__directive] = "_directive",
  [sym_include_directive] = "include_directive",
  [sym_vpath_directive] = "vpath_directive",
  [sym_export_directive] = "export_directive",
  [sym_unexport_directive] = "unexport_directive",
  [sym_override_directive] = "override_directive",
  [sym_undefine_directive] = "undefine_directive",
  [sym_private_directive] = "private_directive",
  [sym_conditional] = "conditional",
  [sym_elsif_directive] = "elsif_directive",
  [sym_else_directive] = "else_directive",
  [sym__conditional_directives] = "_conditional_directives",
  [aux_sym__conditional_consequence] = "_conditional_consequence",
  [sym_ifeq_directive] = "ifeq_directive",
  [sym_ifneq_directive] = "ifneq_directive",
  [sym_ifdef_directive] = "ifdef_directive",
  [sym_ifndef_directive] = "ifndef_directive",
  [sym__conditional_args_cmp] = "_conditional_args_cmp",
  [sym__variable] = "_variable",
  [sym_variable_reference] = "variable_reference",
  [sym_substitution_reference] = "substitution_reference",
  [sym_automatic_variable] = "automatic_variable",
  [sym__function] = "_function",
  [sym_function_call] = "function_call",
  [sym_arguments] = "arguments",
  [sym_shell_function] = "shell_function",
  [sym_list] = "list",
  [sym_paths] = "paths",
  [sym_concatenation] = "concatenation",
  [sym_string] = "string",
  [aux_sym__string] = "_string",
  [sym_archive] = "archive",
  [sym__recipeprefix] = "_recipeprefix",
  [sym__shell_text_without_split] = "_shell_text_without_split",
  [sym_shell_text_with_split] = "shell_text",
  [sym__shell_command] = "_shell_command",
  [sym_text] = "text",
  [aux_sym_makefile_repeat1] = "makefile_repeat1",
  [aux_sym_recipe_repeat1] = "recipe_repeat1",
  [aux_sym_recipe_line_repeat1] = "recipe_line_repeat1",
  [aux_sym_define_directive_repeat1] = "define_directive_repeat1",
  [aux_sym_conditional_repeat1] = "conditional_repeat1",
  [aux_sym_arguments_repeat1] = "arguments_repeat1",
  [aux_sym_list_repeat1] = "list_repeat1",
  [aux_sym_paths_repeat1] = "paths_repeat1",
  [aux_sym_concatenation_repeat1] = "concatenation_repeat1",
  [aux_sym__shell_text_without_split_repeat1] = "_shell_text_without_split_repeat1",
  [aux_sym__shell_text_without_split_repeat2] = "_shell_text_without_split_repeat2",
  [aux_sym_text_repeat1] = "text_repeat1",
  [aux_sym_text_repeat2] = "text_repeat2",
  [alias_sym_pattern_list] = "pattern_list",
  [alias_sym_prerequisites] = "prerequisites",
  [alias_sym_raw_text] = "raw_text",
  [alias_sym_shell_command] = "shell_command",
  [alias_sym_targets] = "targets",
};

static const TSSymbol ts_symbol_map[] = {
  [ts_builtin_sym_end] = ts_builtin_sym_end,
  [sym_word] = sym_word,
  [aux_sym__thing_token1] = aux_sym__thing_token1,
  [anon_sym_COLON] = anon_sym_COLON,
  [anon_sym_AMP_COLON] = anon_sym_AMP_COLON,
  [anon_sym_COLON_COLON] = anon_sym_COLON_COLON,
  [aux_sym__ordinary_rule_token1] = aux_sym__ordinary_rule_token1,
  [anon_sym_PIPE] = anon_sym_PIPE,
  [anon_sym_SEMI] = anon_sym_SEMI,
  [anon_sym_AT] = anon_sym_AT,
  [anon_sym_DASH] = anon_sym_DASH,
  [anon_sym_PLUS] = anon_sym_PLUS,
  [anon_sym_VPATH] = anon_sym_VPATH,
  [anon_sym_EQ] = anon_sym_EQ,
  [anon_sym_COLON_EQ] = anon_sym_COLON_EQ,
  [anon_sym_COLON_COLON_EQ] = anon_sym_COLON_COLON_EQ,
  [anon_sym_QMARK_EQ] = anon_sym_QMARK_EQ,
  [anon_sym_PLUS_EQ] = anon_sym_PLUS_EQ,
  [anon_sym_DOTRECIPEPREFIX] = anon_sym_DOTRECIPEPREFIX,
  [anon_sym_BANG_EQ] = anon_sym_BANG_EQ,
  [anon_sym_define] = anon_sym_define,
  [anon_sym_endef] = anon_sym_endef,
  [anon_sym_include] = anon_sym_include,
  [anon_sym_sinclude] = anon_sym_sinclude,
  [anon_sym_DASHinclude] = anon_sym_DASHinclude,
  [anon_sym_vpath] = anon_sym_vpath,
  [anon_sym_export] = anon_sym_export,
  [anon_sym_unexport] = anon_sym_unexport,
  [anon_sym_override] = anon_sym_override,
  [anon_sym_undefine] = anon_sym_undefine,
  [anon_sym_private] = anon_sym_private,
  [anon_sym_endif] = anon_sym_endif,
  [anon_sym_else] = anon_sym_else,
  [anon_sym_ifeq] = anon_sym_ifeq,
  [anon_sym_ifneq] = anon_sym_ifneq,
  [anon_sym_ifdef] = anon_sym_ifdef,
  [anon_sym_ifndef] = anon_sym_ifndef,
  [anon_sym_LPAREN] = anon_sym_LPAREN,
  [anon_sym_COMMA] = anon_sym_COMMA,
  [anon_sym_RPAREN] = anon_sym_RPAREN,
  [anon_sym_DOLLAR] = anon_sym_DOLLAR,
  [anon_sym_DOLLAR_DOLLAR] = anon_sym_DOLLAR_DOLLAR,
  [anon_sym_LPAREN2] = anon_sym_LPAREN,
  [anon_sym_LBRACE] = anon_sym_LBRACE,
  [anon_sym_RBRACE] = anon_sym_RBRACE,
  [aux_sym_variable_reference_token1] = sym_word,
  [anon_sym_AT2] = anon_sym_AT,
  [anon_sym_PERCENT] = anon_sym_PERCENT,
  [anon_sym_LT] = anon_sym_LT,
  [anon_sym_QMARK] = anon_sym_QMARK,
  [anon_sym_CARET] = anon_sym_CARET,
  [anon_sym_PLUS2] = anon_sym_PLUS,
  [anon_sym_SLASH] = anon_sym_SLASH,
  [anon_sym_STAR] = anon_sym_STAR,
  [anon_sym_PERCENT2] = anon_sym_PERCENT,
  [anon_sym_LT2] = anon_sym_LT,
  [anon_sym_QMARK2] = anon_sym_QMARK,
  [anon_sym_CARET2] = anon_sym_CARET,
  [anon_sym_SLASH2] = anon_sym_SLASH,
  [anon_sym_STAR2] = anon_sym_STAR,
  [anon_sym_D] = anon_sym_D,
  [anon_sym_F] = anon_sym_F,
  [anon_sym_subst] = anon_sym_subst,
  [anon_sym_patsubst] = anon_sym_patsubst,
  [anon_sym_strip] = anon_sym_strip,
  [anon_sym_findstring] = anon_sym_findstring,
  [anon_sym_filter] = anon_sym_filter,
  [anon_sym_filter_DASHout] = anon_sym_filter_DASHout,
  [anon_sym_sort] = anon_sym_sort,
  [anon_sym_word] = anon_sym_word,
  [anon_sym_words] = anon_sym_words,
  [anon_sym_wordlist] = anon_sym_wordlist,
  [anon_sym_firstword] = anon_sym_firstword,
  [anon_sym_lastword] = anon_sym_lastword,
  [anon_sym_dir] = anon_sym_dir,
  [anon_sym_notdir] = anon_sym_notdir,
  [anon_sym_suffix] = anon_sym_suffix,
  [anon_sym_basename] = anon_sym_basename,
  [anon_sym_addsuffix] = anon_sym_addsuffix,
  [anon_sym_addprefix] = anon_sym_addprefix,
  [anon_sym_join] = anon_sym_join,
  [anon_sym_wildcard] = anon_sym_wildcard,
  [anon_sym_realpath] = anon_sym_realpath,
  [anon_sym_abspath] = anon_sym_abspath,
  [anon_sym_error] = anon_sym_error,
  [anon_sym_warning] = anon_sym_warning,
  [anon_sym_info] = anon_sym_info,
  [anon_sym_origin] = anon_sym_origin,
  [anon_sym_flavor] = anon_sym_flavor,
  [anon_sym_foreach] = anon_sym_foreach,
  [anon_sym_if] = anon_sym_if,
  [anon_sym_or] = anon_sym_or,
  [anon_sym_and] = anon_sym_and,
  [anon_sym_call] = anon_sym_call,
  [anon_sym_eval] = anon_sym_eval,
  [anon_sym_file] = anon_sym_file,
  [anon_sym_value] = anon_sym_value,
  [anon_sym_shell] = anon_sym_shell,
  [aux_sym_list_token1] = aux_sym_list_token1,
  [anon_sym_COLON2] = anon_sym_COLON,
  [anon_sym_SEMI2] = anon_sym_SEMI,
  [anon_sym_DQUOTE] = anon_sym_DQUOTE,
  [anon_sym_SQUOTE] = anon_sym_SQUOTE,
  [aux_sym__string_token1] = aux_sym__string_token1,
  [anon_sym_RPAREN2] = anon_sym_RPAREN,
  [anon_sym_TAB] = anon_sym_TAB,
  [sym__rawline] = sym__rawline,
  [aux_sym__shell_text_without_split_token1] = aux_sym__shell_text_without_split_token1,
  [anon_sym_SLASH_SLASH] = anon_sym_SLASH_SLASH,
  [aux_sym_text_token1] = aux_sym_text_token1,
  [sym_comment] = sym_comment,
  [sym_makefile] = sym_makefile,
  [sym__thing] = sym__thing,
  [sym_rule] = sym_rule,
  [sym__ordinary_rule] = sym__ordinary_rule,
  [sym__static_pattern_rule] = sym__static_pattern_rule,
  [sym__normal_prerequisites] = sym__normal_prerequisites,
  [sym_recipe] = sym_recipe,
  [sym__attached_recipe_line] = sym__attached_recipe_line,
  [sym__prefixed_recipe_line] = sym__prefixed_recipe_line,
  [sym_recipe_line] = sym_recipe_line,
  [sym__variable_definition] = sym__variable_definition,
  [sym_VPATH_assignment] = sym_VPATH_assignment,
  [sym_RECIPEPREFIX_assignment] = sym_RECIPEPREFIX_assignment,
  [sym_variable_assignment] = sym_variable_assignment,
  [sym_shell_assignment] = sym_shell_assignment,
  [sym_define_directive] = sym_define_directive,
  [sym__directive] = sym__directive,
  [sym_include_directive] = sym_include_directive,
  [sym_vpath_directive] = sym_vpath_directive,
  [sym_export_directive] = sym_export_directive,
  [sym_unexport_directive] = sym_unexport_directive,
  [sym_override_directive] = sym_override_directive,
  [sym_undefine_directive] = sym_undefine_directive,
  [sym_private_directive] = sym_private_directive,
  [sym_conditional] = sym_conditional,
  [sym_elsif_directive] = sym_elsif_directive,
  [sym_else_directive] = sym_else_directive,
  [sym__conditional_directives] = sym__conditional_directives,
  [aux_sym__conditional_consequence] = aux_sym__conditional_consequence,
  [sym_ifeq_directive] = sym_ifeq_directive,
  [sym_ifneq_directive] = sym_ifneq_directive,
  [sym_ifdef_directive] = sym_ifdef_directive,
  [sym_ifndef_directive] = sym_ifndef_directive,
  [sym__conditional_args_cmp] = sym__conditional_args_cmp,
  [sym__variable] = sym__variable,
  [sym_variable_reference] = sym_variable_reference,
  [sym_substitution_reference] = sym_substitution_reference,
  [sym_automatic_variable] = sym_automatic_variable,
  [sym__function] = sym__function,
  [sym_function_call] = sym_function_call,
  [sym_arguments] = sym_arguments,
  [sym_shell_function] = sym_shell_function,
  [sym_list] = sym_list,
  [sym_paths] = sym_paths,
  [sym_concatenation] = sym_concatenation,
  [sym_string] = sym_string,
  [aux_sym__string] = aux_sym__string,
  [sym_archive] = sym_archive,
  [sym__recipeprefix] = sym__recipeprefix,
  [sym__shell_text_without_split] = sym__shell_text_without_split,
  [sym_shell_text_with_split] = sym_shell_text_with_split,
  [sym__shell_command] = sym__shell_command,
  [sym_text] = sym_text,
  [aux_sym_makefile_repeat1] = aux_sym_makefile_repeat1,
  [aux_sym_recipe_repeat1] = aux_sym_recipe_repeat1,
  [aux_sym_recipe_line_repeat1] = aux_sym_recipe_line_repeat1,
  [aux_sym_define_directive_repeat1] = aux_sym_define_directive_repeat1,
  [aux_sym_conditional_repeat1] = aux_sym_conditional_repeat1,
  [aux_sym_arguments_repeat1] = aux_sym_arguments_repeat1,
  [aux_sym_list_repeat1] = aux_sym_list_repeat1,
  [aux_sym_paths_repeat1] = aux_sym_paths_repeat1,
  [aux_sym_concatenation_repeat1] = aux_sym_concatenation_repeat1,
  [aux_sym__shell_text_without_split_repeat1] = aux_sym__shell_text_without_split_repeat1,
  [aux_sym__shell_text_without_split_repeat2] = aux_sym__shell_text_without_split_repeat2,
  [aux_sym_text_repeat1] = aux_sym_text_repeat1,
  [aux_sym_text_repeat2] = aux_sym_text_repeat2,
  [alias_sym_pattern_list] = alias_sym_pattern_list,
  [alias_sym_prerequisites] = alias_sym_prerequisites,
  [alias_sym_raw_text] = alias_sym_raw_text,
  [alias_sym_shell_command] = alias_sym_shell_command,
  [alias_sym_targets] = alias_sym_targets,
};

static const TSSymbolMetadata ts_symbol_metadata[] = {
  [ts_builtin_sym_end] = {
    .visible = false,
    .named = true,
  },
  [sym_word] = {
    .visible = true,
    .named = true,
  },
  [aux_sym__thing_token1] = {
    .visible = false,
    .named = false,
  },
  [anon_sym_COLON] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_AMP_COLON] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_COLON_COLON] = {
    .visible = true,
    .named = false,
  },
  [aux_sym__ordinary_rule_token1] = {
    .visible = false,
    .named = false,
  },
  [anon_sym_PIPE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_SEMI] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_AT] = {
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
  [anon_sym_VPATH] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_COLON_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_COLON_COLON_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_QMARK_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PLUS_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DOTRECIPEPREFIX] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_BANG_EQ] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_define] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_endef] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_include] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_sinclude] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DASHinclude] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_vpath] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_export] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_unexport] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_override] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_undefine] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_private] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_endif] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_else] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_ifeq] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_ifneq] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_ifdef] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_ifndef] = {
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
  [anon_sym_DOLLAR] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DOLLAR_DOLLAR] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LPAREN2] = {
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
  [aux_sym_variable_reference_token1] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_AT2] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PERCENT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_QMARK] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_CARET] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PLUS2] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_SLASH] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_STAR] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_PERCENT2] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LT2] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_QMARK2] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_CARET2] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_SLASH2] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_STAR2] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_D] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_F] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_subst] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_patsubst] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_strip] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_findstring] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_filter] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_filter_DASHout] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_sort] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_word] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_words] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_wordlist] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_firstword] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_lastword] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_dir] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_notdir] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_suffix] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_basename] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_addsuffix] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_addprefix] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_join] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_wildcard] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_realpath] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_abspath] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_error] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_warning] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_info] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_origin] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_flavor] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_foreach] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_if] = {
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
  [anon_sym_call] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_eval] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_file] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_value] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_shell] = {
    .visible = true,
    .named = false,
  },
  [aux_sym_list_token1] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_COLON2] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_SEMI2] = {
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
  [aux_sym__string_token1] = {
    .visible = false,
    .named = false,
  },
  [anon_sym_RPAREN2] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_TAB] = {
    .visible = true,
    .named = false,
  },
  [sym__rawline] = {
    .visible = false,
    .named = true,
  },
  [aux_sym__shell_text_without_split_token1] = {
    .visible = false,
    .named = false,
  },
  [anon_sym_SLASH_SLASH] = {
    .visible = true,
    .named = true,
  },
  [aux_sym_text_token1] = {
    .visible = false,
    .named = false,
  },
  [sym_comment] = {
    .visible = true,
    .named = true,
  },
  [sym_makefile] = {
    .visible = true,
    .named = true,
  },
  [sym__thing] = {
    .visible = false,
    .named = true,
  },
  [sym_rule] = {
    .visible = true,
    .named = true,
  },
  [sym__ordinary_rule] = {
    .visible = false,
    .named = true,
  },
  [sym__static_pattern_rule] = {
    .visible = false,
    .named = true,
  },
  [sym__normal_prerequisites] = {
    .visible = false,
    .named = true,
  },
  [sym_recipe] = {
    .visible = true,
    .named = true,
  },
  [sym__attached_recipe_line] = {
    .visible = false,
    .named = true,
  },
  [sym__prefixed_recipe_line] = {
    .visible = false,
    .named = true,
  },
  [sym_recipe_line] = {
    .visible = true,
    .named = true,
  },
  [sym__variable_definition] = {
    .visible = false,
    .named = true,
  },
  [sym_VPATH_assignment] = {
    .visible = true,
    .named = true,
  },
  [sym_RECIPEPREFIX_assignment] = {
    .visible = true,
    .named = true,
  },
  [sym_variable_assignment] = {
    .visible = true,
    .named = true,
  },
  [sym_shell_assignment] = {
    .visible = true,
    .named = true,
  },
  [sym_define_directive] = {
    .visible = true,
    .named = true,
  },
  [sym__directive] = {
    .visible = false,
    .named = true,
  },
  [sym_include_directive] = {
    .visible = true,
    .named = true,
  },
  [sym_vpath_directive] = {
    .visible = true,
    .named = true,
  },
  [sym_export_directive] = {
    .visible = true,
    .named = true,
  },
  [sym_unexport_directive] = {
    .visible = true,
    .named = true,
  },
  [sym_override_directive] = {
    .visible = true,
    .named = true,
  },
  [sym_undefine_directive] = {
    .visible = true,
    .named = true,
  },
  [sym_private_directive] = {
    .visible = true,
    .named = true,
  },
  [sym_conditional] = {
    .visible = true,
    .named = true,
  },
  [sym_elsif_directive] = {
    .visible = true,
    .named = true,
  },
  [sym_else_directive] = {
    .visible = true,
    .named = true,
  },
  [sym__conditional_directives] = {
    .visible = false,
    .named = true,
  },
  [aux_sym__conditional_consequence] = {
    .visible = false,
    .named = false,
  },
  [sym_ifeq_directive] = {
    .visible = true,
    .named = true,
  },
  [sym_ifneq_directive] = {
    .visible = true,
    .named = true,
  },
  [sym_ifdef_directive] = {
    .visible = true,
    .named = true,
  },
  [sym_ifndef_directive] = {
    .visible = true,
    .named = true,
  },
  [sym__conditional_args_cmp] = {
    .visible = false,
    .named = true,
  },
  [sym__variable] = {
    .visible = false,
    .named = true,
  },
  [sym_variable_reference] = {
    .visible = true,
    .named = true,
  },
  [sym_substitution_reference] = {
    .visible = true,
    .named = true,
  },
  [sym_automatic_variable] = {
    .visible = true,
    .named = true,
  },
  [sym__function] = {
    .visible = false,
    .named = true,
  },
  [sym_function_call] = {
    .visible = true,
    .named = true,
  },
  [sym_arguments] = {
    .visible = true,
    .named = true,
  },
  [sym_shell_function] = {
    .visible = true,
    .named = true,
  },
  [sym_list] = {
    .visible = true,
    .named = true,
  },
  [sym_paths] = {
    .visible = true,
    .named = true,
  },
  [sym_concatenation] = {
    .visible = true,
    .named = true,
  },
  [sym_string] = {
    .visible = true,
    .named = true,
  },
  [aux_sym__string] = {
    .visible = false,
    .named = false,
  },
  [sym_archive] = {
    .visible = true,
    .named = true,
  },
  [sym__recipeprefix] = {
    .visible = false,
    .named = true,
  },
  [sym__shell_text_without_split] = {
    .visible = false,
    .named = true,
  },
  [sym_shell_text_with_split] = {
    .visible = true,
    .named = true,
  },
  [sym__shell_command] = {
    .visible = false,
    .named = true,
  },
  [sym_text] = {
    .visible = true,
    .named = true,
  },
  [aux_sym_makefile_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_recipe_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_recipe_line_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_define_directive_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_conditional_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_arguments_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_list_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_paths_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_concatenation_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__shell_text_without_split_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym__shell_text_without_split_repeat2] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_text_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_text_repeat2] = {
    .visible = false,
    .named = false,
  },
  [alias_sym_pattern_list] = {
    .visible = true,
    .named = true,
  },
  [alias_sym_prerequisites] = {
    .visible = true,
    .named = true,
  },
  [alias_sym_raw_text] = {
    .visible = true,
    .named = true,
  },
  [alias_sym_shell_command] = {
    .visible = true,
    .named = true,
  },
  [alias_sym_targets] = {
    .visible = true,
    .named = true,
  },
};

enum ts_field_identifiers {
  field_archive = 1,
  field_arg0 = 2,
  field_arg1 = 3,
  field_argument = 4,
  field_condition = 5,
  field_consequence = 6,
  field_directories = 7,
  field_filenames = 8,
  field_function = 9,
  field_members = 10,
  field_name = 11,
  field_normal = 12,
  field_operator = 13,
  field_order_only = 14,
  field_pattern = 15,
  field_prerequisite = 16,
  field_replacement = 17,
  field_string = 18,
  field_target = 19,
  field_target_or_pattern = 20,
  field_text = 21,
  field_value = 22,
  field_variable = 23,
  field_variables = 24,
};

static const char * const ts_field_names[] = {
  [0] = NULL,
  [field_archive] = "archive",
  [field_arg0] = "arg0",
  [field_arg1] = "arg1",
  [field_argument] = "argument",
  [field_condition] = "condition",
  [field_consequence] = "consequence",
  [field_directories] = "directories",
  [field_filenames] = "filenames",
  [field_function] = "function",
  [field_members] = "members",
  [field_name] = "name",
  [field_normal] = "normal",
  [field_operator] = "operator",
  [field_order_only] = "order_only",
  [field_pattern] = "pattern",
  [field_prerequisite] = "prerequisite",
  [field_replacement] = "replacement",
  [field_string] = "string",
  [field_target] = "target",
  [field_target_or_pattern] = "target_or_pattern",
  [field_text] = "text",
  [field_value] = "value",
  [field_variable] = "variable",
  [field_variables] = "variables",
};

static const TSFieldMapSlice ts_field_map_slices[PRODUCTION_ID_COUNT] = {
  [1] = {.index = 0, .length = 2},
  [2] = {.index = 2, .length = 2},
  [3] = {.index = 4, .length = 2},
  [5] = {.index = 6, .length = 1},
  [6] = {.index = 7, .length = 1},
  [7] = {.index = 8, .length = 1},
  [8] = {.index = 9, .length = 1},
  [9] = {.index = 10, .length = 2},
  [10] = {.index = 12, .length = 2},
  [11] = {.index = 14, .length = 3},
  [12] = {.index = 17, .length = 2},
  [14] = {.index = 19, .length = 1},
  [15] = {.index = 20, .length = 1},
  [18] = {.index = 21, .length = 1},
  [19] = {.index = 22, .length = 3},
  [20] = {.index = 25, .length = 2},
  [21] = {.index = 27, .length = 1},
  [22] = {.index = 28, .length = 2},
  [23] = {.index = 30, .length = 2},
  [24] = {.index = 32, .length = 1},
  [25] = {.index = 33, .length = 2},
  [26] = {.index = 35, .length = 2},
  [29] = {.index = 37, .length = 1},
  [30] = {.index = 38, .length = 3},
  [32] = {.index = 41, .length = 1},
  [33] = {.index = 42, .length = 1},
  [34] = {.index = 43, .length = 1},
  [35] = {.index = 44, .length = 1},
  [36] = {.index = 45, .length = 2},
  [37] = {.index = 47, .length = 3},
  [40] = {.index = 50, .length = 1},
  [41] = {.index = 51, .length = 1},
  [42] = {.index = 52, .length = 3},
  [43] = {.index = 55, .length = 1},
  [44] = {.index = 56, .length = 2},
  [45] = {.index = 58, .length = 2},
  [46] = {.index = 60, .length = 2},
  [47] = {.index = 62, .length = 1},
  [48] = {.index = 63, .length = 2},
  [49] = {.index = 65, .length = 3},
  [52] = {.index = 68, .length = 1},
  [53] = {.index = 69, .length = 3},
  [54] = {.index = 72, .length = 1},
  [55] = {.index = 73, .length = 3},
  [56] = {.index = 76, .length = 4},
  [57] = {.index = 80, .length = 2},
  [58] = {.index = 82, .length = 2},
  [59] = {.index = 84, .length = 2},
  [60] = {.index = 86, .length = 2},
  [61] = {.index = 88, .length = 3},
  [63] = {.index = 91, .length = 3},
  [64] = {.index = 94, .length = 4},
  [65] = {.index = 98, .length = 2},
  [66] = {.index = 100, .length = 2},
  [67] = {.index = 102, .length = 4},
  [68] = {.index = 106, .length = 4},
  [69] = {.index = 110, .length = 2},
  [70] = {.index = 112, .length = 2},
  [71] = {.index = 114, .length = 3},
  [72] = {.index = 117, .length = 3},
  [73] = {.index = 120, .length = 3},
  [74] = {.index = 123, .length = 4},
  [75] = {.index = 127, .length = 4},
  [76] = {.index = 131, .length = 2},
  [77] = {.index = 133, .length = 4},
  [78] = {.index = 137, .length = 3},
  [79] = {.index = 140, .length = 4},
};

static const TSFieldMapEntry ts_field_map_entries[] = {
  [0] =
    {field_normal, 0, .inherited = true},
    {field_order_only, 0, .inherited = true},
  [2] =
    {field_prerequisite, 0, .inherited = true},
    {field_target, 0, .inherited = true},
  [4] =
    {field_string, 0},
    {field_string, 1},
  [6] =
    {field_filenames, 1},
  [7] =
    {field_pattern, 1},
  [8] =
    {field_variables, 1},
  [9] =
    {field_variable, 1},
  [10] =
    {field_arg0, 0},
    {field_arg1, 1},
  [12] =
    {field_arg0, 1, .inherited = true},
    {field_arg1, 1, .inherited = true},
  [14] =
    {field_string, 0},
    {field_string, 1},
    {field_string, 2},
  [17] =
    {field_name, 0},
    {field_operator, 1},
  [19] =
    {field_condition, 0},
  [20] =
    {field_condition, 1},
  [21] =
    {field_normal, 0},
  [22] =
    {field_name, 0},
    {field_operator, 1},
    {field_value, 2},
  [25] =
    {field_directories, 2},
    {field_pattern, 1},
  [27] =
    {field_argument, 0},
  [28] =
    {field_name, 0},
    {field_operator, 2},
  [30] =
    {field_archive, 0},
    {field_members, 2},
  [32] =
    {field_consequence, 2},
  [33] =
    {field_condition, 1},
    {field_consequence, 2},
  [35] =
    {field_condition, 0},
    {field_consequence, 1},
  [37] =
    {field_normal, 2, .inherited = true},
  [38] =
    {field_name, 0},
    {field_operator, 2},
    {field_value, 3},
  [41] =
    {field_name, 1},
  [42] =
    {field_arg1, 2},
  [43] =
    {field_arg0, 1},
  [44] =
    {field_function, 2},
  [45] =
    {field_argument, 0},
    {field_argument, 1, .inherited = true},
  [47] =
    {field_name, 0},
    {field_operator, 1},
    {field_value, 3},
  [50] =
    {field_normal, 3, .inherited = true},
  [51] =
    {field_order_only, 3},
  [52] =
    {field_name, 2},
    {field_operator, 3},
    {field_target_or_pattern, 0},
  [55] =
    {field_target, 2},
  [56] =
    {field_name, 1},
    {field_value, 3},
  [58] =
    {field_name, 1},
    {field_operator, 2},
  [60] =
    {field_arg0, 1},
    {field_arg1, 3},
  [62] =
    {field_argument, 1},
  [63] =
    {field_argument, 0, .inherited = true},
    {field_argument, 1, .inherited = true},
  [65] =
    {field_name, 0},
    {field_operator, 2},
    {field_value, 4},
  [68] =
    {field_order_only, 4},
  [69] =
    {field_name, 3},
    {field_operator, 4},
    {field_target_or_pattern, 0},
  [72] =
    {field_target, 3},
  [73] =
    {field_name, 2},
    {field_operator, 4},
    {field_target_or_pattern, 0},
  [76] =
    {field_name, 2},
    {field_operator, 3},
    {field_target_or_pattern, 0},
    {field_value, 4},
  [80] =
    {field_normal, 2, .inherited = true},
    {field_order_only, 4},
  [82] =
    {field_prerequisite, 4},
    {field_target, 2},
  [84] =
    {field_name, 1},
    {field_value, 4},
  [86] =
    {field_name, 1},
    {field_operator, 3},
  [88] =
    {field_name, 1},
    {field_operator, 2},
    {field_value, 4},
  [91] =
    {field_name, 3},
    {field_operator, 5},
    {field_target_or_pattern, 0},
  [94] =
    {field_name, 3},
    {field_operator, 4},
    {field_target_or_pattern, 0},
    {field_value, 5},
  [98] =
    {field_normal, 3, .inherited = true},
    {field_order_only, 5},
  [100] =
    {field_prerequisite, 5},
    {field_target, 3},
  [102] =
    {field_name, 2},
    {field_operator, 4},
    {field_target_or_pattern, 0},
    {field_value, 5},
  [106] =
    {field_name, 2},
    {field_operator, 3},
    {field_target_or_pattern, 0},
    {field_value, 5},
  [110] =
    {field_prerequisite, 5},
    {field_target, 2},
  [112] =
    {field_name, 1},
    {field_value, 5},
  [114] =
    {field_name, 1},
    {field_operator, 3},
    {field_value, 5},
  [117] =
    {field_name, 1},
    {field_operator, 2},
    {field_value, 5},
  [120] =
    {field_pattern, 4},
    {field_replacement, 6},
    {field_text, 2},
  [123] =
    {field_name, 3},
    {field_operator, 5},
    {field_target_or_pattern, 0},
    {field_value, 6},
  [127] =
    {field_name, 3},
    {field_operator, 4},
    {field_target_or_pattern, 0},
    {field_value, 6},
  [131] =
    {field_prerequisite, 6},
    {field_target, 3},
  [133] =
    {field_name, 2},
    {field_operator, 4},
    {field_target_or_pattern, 0},
    {field_value, 6},
  [137] =
    {field_name, 1},
    {field_operator, 3},
    {field_value, 6},
  [140] =
    {field_name, 3},
    {field_operator, 5},
    {field_target_or_pattern, 0},
    {field_value, 7},
};

static const TSSymbol ts_alias_sequences[PRODUCTION_ID_COUNT][MAX_ALIAS_SEQUENCE_LENGTH] = {
  [0] = {0},
  [4] = {
    [0] = anon_sym_SLASH_SLASH,
  },
  [13] = {
    [0] = alias_sym_shell_command,
  },
  [16] = {
    [0] = sym_shell_text_with_split,
  },
  [17] = {
    [0] = alias_sym_targets,
  },
  [18] = {
    [0] = alias_sym_prerequisites,
  },
  [27] = {
    [1] = sym_shell_text_with_split,
  },
  [28] = {
    [0] = sym_shell_text_with_split,
    [1] = sym_shell_text_with_split,
  },
  [29] = {
    [0] = alias_sym_targets,
  },
  [31] = {
    [1] = anon_sym_SLASH_SLASH,
  },
  [38] = {
    [1] = sym_shell_text_with_split,
    [2] = sym_shell_text_with_split,
  },
  [39] = {
    [0] = sym_shell_text_with_split,
    [2] = sym_shell_text_with_split,
  },
  [40] = {
    [0] = alias_sym_targets,
  },
  [41] = {
    [0] = alias_sym_targets,
    [3] = alias_sym_prerequisites,
  },
  [43] = {
    [0] = alias_sym_targets,
    [2] = alias_sym_pattern_list,
  },
  [44] = {
    [3] = alias_sym_raw_text,
  },
  [50] = {
    [1] = sym_shell_text_with_split,
    [3] = sym_shell_text_with_split,
  },
  [51] = {
    [0] = sym_shell_text_with_split,
    [3] = sym_shell_text_with_split,
  },
  [52] = {
    [0] = alias_sym_targets,
    [4] = alias_sym_prerequisites,
  },
  [54] = {
    [0] = alias_sym_targets,
    [3] = alias_sym_pattern_list,
  },
  [57] = {
    [0] = alias_sym_targets,
    [4] = alias_sym_prerequisites,
  },
  [58] = {
    [0] = alias_sym_targets,
    [2] = alias_sym_pattern_list,
    [4] = alias_sym_pattern_list,
  },
  [59] = {
    [4] = alias_sym_raw_text,
  },
  [61] = {
    [4] = alias_sym_raw_text,
  },
  [62] = {
    [1] = sym_shell_text_with_split,
    [4] = sym_shell_text_with_split,
  },
  [65] = {
    [0] = alias_sym_targets,
    [5] = alias_sym_prerequisites,
  },
  [66] = {
    [0] = alias_sym_targets,
    [3] = alias_sym_pattern_list,
    [5] = alias_sym_pattern_list,
  },
  [69] = {
    [0] = alias_sym_targets,
    [2] = alias_sym_pattern_list,
    [5] = alias_sym_pattern_list,
  },
  [70] = {
    [5] = alias_sym_raw_text,
  },
  [71] = {
    [5] = alias_sym_raw_text,
  },
  [72] = {
    [5] = alias_sym_raw_text,
  },
  [76] = {
    [0] = alias_sym_targets,
    [3] = alias_sym_pattern_list,
    [6] = alias_sym_pattern_list,
  },
  [78] = {
    [6] = alias_sym_raw_text,
  },
};

static const uint16_t ts_non_terminal_alias_map[] = {
  sym_list, 4,
    sym_list,
    alias_sym_pattern_list,
    alias_sym_prerequisites,
    alias_sym_targets,
  sym__shell_text_without_split, 2,
    sym__shell_text_without_split,
    sym_shell_text_with_split,
  sym_text, 2,
    sym_text,
    alias_sym_shell_command,
  aux_sym_define_directive_repeat1, 2,
    aux_sym_define_directive_repeat1,
    alias_sym_raw_text,
  0,
};

// Filled from tables.zlib by grammar.go, see fetchgrammar.
static TSStateId ts_primary_state_ids[STATE_COUNT];

static TSCharacterRange sym_word_character_set_1[] = {
  {'!', '$'}, {'&', '*'}, {',', ','}, {'0', '9'}, {';', '<'}, {'>', '?'}, {'E', 'E'}, {'[', '^'},
  {'`', 'b'}, {'f', 'f'}, {'n', 'n'}, {'r', 'r'}, {'t', 't'}, {'v', 'v'}, {'{', '~'},
};

static bool ts_lex(TSLexer *lexer, TSStateId state) {
  START_LEXER();
  eof = lexer->eof(lexer);
  switch (state) {
    case 0:
      if (eof) ADVANCE(107);
      ADVANCE_MAP(
        '!', 95,
        '"', 168,
        '#', 233,
        '$', 136,
        '%', 150,
        '&', 93,
        '\'', 169,
        '(', 138,
        ')', 200,
        '*', 156,
        '+', 154,
        ',', 133,
        '-', 121,
        '.', 185,
        '/', 155,
        ':', 166,
        ';', 167,
        '<', 151,
        '=', 123,
        '?', 152,
        '@', 149,
        '\\', 3,
        '^', 153,
        'e', 196,
        '{', 140,
        '|', 118,
        '}', 143,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(105);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 1:
      ADVANCE_MAP(
        '\t', 201,
        '"', 168,
        '#', 233,
        '$', 136,
        '\'', 169,
        '-', 194,
        '.', 185,
        '\\', 5,
      );
      if (('\n' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(1);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('/' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 2:
      if (lookahead == '\t') ADVANCE(202);
      if (lookahead == '#') ADVANCE(212);
      if (lookahead == '$') ADVANCE(136);
      if (lookahead == '/') ADVANCE(210);
      if (lookahead == '\\') ADVANCE(28);
      if (lookahead == '\n' ||
          lookahead == '\r') SKIP(2);
      if (lookahead == 0x0b ||
          lookahead == '\f' ||
          lookahead == ' ') ADVANCE(207);
      if (lookahead != 0) ADVANCE(214);
      END_STATE();
    case 3:
      if (lookahead == '\n') SKIP(45);
      if (lookahead == '\r') SKIP(45);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(101);
      if (set_contains(sym_word_character_set_1, 15, lookahead)) ADVANCE(199);
      END_STATE();
    case 4:
      if (lookahead == '\n') SKIP(54);
      if (lookahead == '\r') SKIP(54);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(101);
      if (set_contains(sym_word_character_set_1, 15, lookahead)) ADVANCE(199);
      END_STATE();
    case 5:
      if (lookahead == '\n') SKIP(1);
      if (lookahead == '\r') SKIP(1);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(101);
      if (set_contains(sym_word_character_set_1, 15, lookahead)) ADVANCE(199);
      END_STATE();
    case 6:
      if (lookahead == '\n') SKIP(49);
      if (lookahead == '\r') SKIP(49);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(101);
      if (set_contains(sym_word_character_set_1, 15, lookahead)) ADVANCE(199);
      END_STATE();
    case 7:
      if (lookahead == '\n') SKIP(47);
      if (lookahead == '\r') SKIP(47);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(101);
      if (set_contains(sym_word_character_set_1, 15, lookahead)) ADVANCE(199);
      END_STATE();
    case 8:
      if (lookahead == '\n') ADVANCE(163);
      if (lookahead == '\r') ADVANCE(164);
      END_STATE();
    case 9:
      if (lookahead == '\n') ADVANCE(163);
      if (lookahead == '\r') ADVANCE(164);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(101);
      if (set_contains(sym_word_character_set_1, 15, lookahead)) ADVANCE(199);
      END_STATE();
    case 10:
      if (lookahead == '\n') ADVANCE(163);
      if (lookahead == '\r') ADVANCE(164);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(213);
      if (lookahead != 0) ADVANCE(214);
      END_STATE();
    case 11:
      ADVANCE_MAP(
        '\n', 109,
        '\r', 109,
        '#', 145,
        '$', 136,
        '%', 150,
        '(', 138,
        ')', 144,
        '*', 156,
        '+', 154,
        '/', 155,
        '<', 151,
        '?', 152,
        '@', 149,
        '\\', 144,
        '^', 153,
        '{', 141,
        '\t', 144,
        0x0b, 144,
        '\f', 144,
        ' ', 144,
      );
      if (lookahead != 0) ADVANCE(146);
      END_STATE();
    case 12:
      ADVANCE_MAP(
        '\n', 109,
        '\r', 109,
        '#', 147,
        '$', 136,
        '%', 150,
        '(', 139,
        '*', 156,
        '+', 154,
        '/', 155,
        '<', 151,
        '?', 152,
        '@', 149,
        '\\', 144,
        '^', 153,
        '{', 142,
        '\t', 144,
        0x0b, 144,
        '\f', 144,
        ' ', 144,
      );
      if (lookahead != 0) ADVANCE(148);
      END_STATE();
    case 13:
      if (lookahead == '\n') SKIP(63);
      if (lookahead == '\r') SKIP(63);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(101);
      if (set_contains(sym_word_character_set_1, 15, lookahead)) ADVANCE(199);
      END_STATE();
    case 14:
      if (lookahead == '\n') SKIP(80);
      if (lookahead == '\r') ADVANCE(144);
      if (lookahead == '#') ADVANCE(145);
      if (lookahead == '$') ADVANCE(136);
      if (lookahead == '%') ADVANCE(150);
      if (lookahead == '(') ADVANCE(138);
      if (lookahead == ')') ADVANCE(135);
      if (lookahead == '*') ADVANCE(156);
      if (lookahead == '+') ADVANCE(154);
      if (lookahead == ',') ADVANCE(134);
      if (lookahead == '/') ADVANCE(155);
      if (lookahead == '<') ADVANCE(151);
      if (lookahead == '?') ADVANCE(152);
      if (lookahead == '@') ADVANCE(149);
      if (lookahead == '\\') ADVANCE(144);
      if (lookahead == '^') ADVANCE(153);
      if (lookahead == '{') ADVANCE(141);
      if (('\t' <= lookahead && lookahead <= '\f') ||
          lookahead == ' ') ADVANCE(144);
      if (lookahead != 0) ADVANCE(146);
      END_STATE();
    case 15:
      if (lookahead == '\n') ADVANCE(219);
      if (lookahead == '\r') ADVANCE(220);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(226);
      if (lookahead != 0) ADVANCE(227);
      END_STATE();
    case 16:
      if (lookahead == '\n') SKIP(88);
      if (lookahead == '\r') ADVANCE(144);
      if (lookahead == '#') ADVANCE(147);
      if (lookahead == '$') ADVANCE(136);
      if (lookahead == '%') ADVANCE(150);
      if (lookahead == '(') ADVANCE(139);
      if (lookahead == '*') ADVANCE(156);
      if (lookahead == '+') ADVANCE(154);
      if (lookahead == '/') ADVANCE(155);
      if (lookahead == '<') ADVANCE(151);
      if (lookahead == '?') ADVANCE(152);
      if (lookahead == '@') ADVANCE(149);
      if (lookahead == '\\') ADVANCE(144);
      if (lookahead == '^') ADVANCE(153);
      if (lookahead == '{') ADVANCE(142);
      if (('\t' <= lookahead && lookahead <= '\f') ||
          lookahead == ' ') ADVANCE(144);
      if (lookahead != 0) ADVANCE(148);
      END_STATE();
    case 17:
      if (lookahead == '\n') SKIP(88);
      if (lookahead == '\r') SKIP(88);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(213);
      if (lookahead != 0) ADVANCE(214);
      END_STATE();
    case 18:
      if (lookahead == '\n') SKIP(81);
      if (lookahead == '\r') ADVANCE(144);
      if (lookahead == '#') ADVANCE(145);
      if (lookahead == '$') ADVANCE(136);
      if (lookahead == '%') ADVANCE(150);
      if (lookahead == '(') ADVANCE(138);
      if (lookahead == ')') ADVANCE(135);
      if (lookahead == '*') ADVANCE(156);
      if (lookahead == '+') ADVANCE(154);
      if (lookahead == '/') ADVANCE(155);
      if (lookahead == '<') ADVANCE(151);
      if (lookahead == '?') ADVANCE(152);
      if (lookahead == '@') ADVANCE(149);
      if (lookahead == '\\') ADVANCE(144);
      if (lookahead == '^') ADVANCE(153);
      if (lookahead == '{') ADVANCE(141);
      if (('\t' <= lookahead && lookahead <= '\f') ||
          lookahead == ' ') ADVANCE(144);
      if (lookahead != 0) ADVANCE(146);
      END_STATE();
    case 19:
      if (lookahead == '\n') ADVANCE(221);
      if (lookahead == '\r') ADVANCE(222);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(226);
      if (lookahead != 0) ADVANCE(227);
      END_STATE();
    case 20:
      if (lookahead == '\n') SKIP(56);
      if (lookahead == '\r') SKIP(56);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(101);
      if (set_contains(sym_word_character_set_1, 15, lookahead)) ADVANCE(199);
      END_STATE();
    case 21:
      if (lookahead == '\n') SKIP(55);
      if (lookahead == '\r') SKIP(55);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(101);
      if (set_contains(sym_word_character_set_1, 15, lookahead)) ADVANCE(199);
      END_STATE();
    case 22:
      if (lookahead == '\n') SKIP(66);
      if (lookahead == '\r') SKIP(66);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(101);
      if (set_contains(sym_word_character_set_1, 15, lookahead)) ADVANCE(199);
      END_STATE();
    case 23:
      if (lookahead == '\n') SKIP(64);
      if (lookahead == '\r') SKIP(64);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(101);
      if (set_contains(sym_word_character_set_1, 15, lookahead)) ADVANCE(199);
      END_STATE();
    case 24:
      if (lookahead == '\n') SKIP(61);
      if (lookahead == '\r') SKIP(61);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(101);
      if (set_contains(sym_word_character_set_1, 15, lookahead)) ADVANCE(199);
      END_STATE();
    case 25:
      if (lookahead == '\n') SKIP(67);
      if (lookahead == '\r') SKIP(67);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(101);
      if (set_contains(sym_word_character_set_1, 15, lookahead)) ADVANCE(199);
      END_STATE();
    case 26:
      if (lookahead == '\n') SKIP(53);
      if (lookahead == '\r') SKIP(53);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(101);
      if (set_contains(sym_word_character_set_1, 15, lookahead)) ADVANCE(199);
      END_STATE();
    case 27:
      if (lookahead == '\n') SKIP(87);
      if (lookahead == '\r') SKIP(87);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(213);
      if (lookahead != 0) ADVANCE(214);
      END_STATE();
    case 28:
      if (lookahead == '\n') SKIP(2);
      if (lookahead == '\r') SKIP(2);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(213);
      if (lookahead != 0) ADVANCE(214);
      END_STATE();
    case 29:
      if (lookahead == '\n') SKIP(69);
      if (lookahead == '\r') SKIP(69);
      END_STATE();
    case 30:
      if (lookahead == '\n') SKIP(73);
      if (lookahead == '\r') SKIP(73);
      END_STATE();
    case 31:
      if (lookahead == '\n') SKIP(68);
      if (lookahead == '\r') SKIP(68);
      if (lookahead != 0) ADVANCE(173);
      END_STATE();
    case 32:
      if (lookahead == '\n') SKIP(91);
      if (lookahead == '\r') SKIP(91);
      if (lookahead != 0) ADVANCE(173);
      END_STATE();
    case 33:
      if (lookahead == '\n') SKIP(77);
      if (lookahead == '\r') SKIP(77);
      END_STATE();
    case 34:
      if (lookahead == '\n') SKIP(77);
      if (lookahead == '#') ADVANCE(144);
      if (lookahead == '%') ADVANCE(150);
      if (lookahead == '(') ADVANCE(138);
      if (lookahead == '*') ADVANCE(156);
      if (lookahead == '+') ADVANCE(154);
      if (lookahead == '/') ADVANCE(155);
      if (lookahead == '<') ADVANCE(151);
      if (lookahead == '?') ADVANCE(152);
      if (lookahead == '@') ADVANCE(149);
      if (lookahead == '\\') ADVANCE(144);
      if (lookahead == '^') ADVANCE(153);
      if (lookahead == '{') ADVANCE(140);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') ADVANCE(144);
      if (lookahead != 0) ADVANCE(144);
      END_STATE();
    case 35:
      if (lookahead == '\n') SKIP(76);
      if (lookahead == '\r') SKIP(76);
      END_STATE();
    case 36:
      ADVANCE_MAP(
        '\n', 203,
        '\r', 203,
        '#', 229,
        '\\', 37,
        'e', 41,
        '\t', 36,
        0x0b, 36,
        '\f', 36,
        ' ', 36,
      );
      if (lookahead != 0) ADVANCE(42);
      END_STATE();
    case 37:
      if (lookahead == '\n') ADVANCE(203);
      if (lookahead == '\r') ADVANCE(203);
      if (lookahead != 0) ADVANCE(42);
      END_STATE();
    case 38:
      if (lookahead == '\n') ADVANCE(206);
      if (lookahead == '\r') ADVANCE(205);
      if (lookahead == 'd') ADVANCE(39);
      if (lookahead != 0) ADVANCE(42);
      END_STATE();
    case 39:
      if (lookahead == '\n') ADVANCE(206);
      if (lookahead == '\r') ADVANCE(205);
      if (lookahead == 'e') ADVANCE(40);
      if (lookahead != 0) ADVANCE(42);
      END_STATE();
    case 40:
      if (lookahead == '\n') ADVANCE(206);
      if (lookahead == '\r') ADVANCE(205);
      if (lookahead == 'f') ADVANCE(130);
      if (lookahead != 0) ADVANCE(42);
      END_STATE();
    case 41:
      if (lookahead == '\n') ADVANCE(206);
      if (lookahead == '\r') ADVANCE(205);
      if (lookahead == 'n') ADVANCE(38);
      if (lookahead != 0) ADVANCE(42);
      END_STATE();
    case 42:
      if (lookahead == '\n') ADVANCE(206);
      if (lookahead == '\r') ADVANCE(205);
      if (lookahead != 0) ADVANCE(42);
      END_STATE();
    case 43:
      if (lookahead == '\n') SKIP(79);
      if (lookahead == '\r') SKIP(79);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(101);
      if (set_contains(sym_word_character_set_1, 15, lookahead)) ADVANCE(199);
      END_STATE();
    case 44:
      if (lookahead == '\r') ADVANCE(228);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(226);
      if (lookahead != 0) ADVANCE(227);
      END_STATE();
    case 45:
      ADVANCE_MAP(
        '!', 95,
        '"', 168,
        '#', 233,
        '$', 136,
        '%', 157,
        '&', 93,
        '\'', 169,
        '(', 132,
        ')', 135,
        '*', 162,
        '+', 122,
        ',', 133,
        '-', 121,
        '.', 185,
        '/', 161,
        ':', 111,
        ';', 119,
        '<', 158,
        '=', 123,
        '?', 159,
        '@', 120,
        '\\', 3,
        '^', 160,
        'e', 196,
        '|', 118,
        '}', 143,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(45);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 46:
      ADVANCE_MAP(
        '!', 95,
        '"', 168,
        '#', 233,
        '$', 136,
        '&', 93,
        '\'', 169,
        '(', 138,
        ')', 200,
        '+', 174,
        ':', 111,
        '=', 123,
        '?', 175,
        '\\', 9,
        '\t', 117,
        ' ', 117,
      );
      if (('\n' <= lookahead && lookahead <= '\r')) SKIP(47);
      if (('%' <= lookahead && lookahead <= '*') ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('@' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 47:
      ADVANCE_MAP(
        '!', 95,
        '"', 168,
        '#', 233,
        '$', 136,
        '&', 93,
        '\'', 169,
        '+', 174,
        ':', 111,
        '=', 123,
        '?', 175,
        '\\', 7,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(47);
      if (lookahead == '%' ||
          lookahead == '*' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('@' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 48:
      ADVANCE_MAP(
        '"', 168,
        '#', 233,
        '$', 136,
        '%', 157,
        '\'', 169,
        ')', 200,
        '*', 162,
        '+', 122,
        '/', 161,
        '<', 158,
        '?', 159,
        '@', 120,
        '\\', 6,
        '^', 160,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(49);
      if (('-' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 49:
      ADVANCE_MAP(
        '"', 168,
        '#', 233,
        '$', 136,
        '%', 157,
        '\'', 169,
        '*', 162,
        '+', 122,
        '/', 161,
        '<', 158,
        '?', 159,
        '@', 120,
        '\\', 6,
        '^', 160,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(49);
      if (('-' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 50:
      ADVANCE_MAP(
        '"', 168,
        '#', 233,
        '$', 136,
        '&', 93,
        '\'', 169,
        '(', 138,
        ')', 200,
        ':', 112,
        '\\', 9,
        '\t', 117,
        ' ', 117,
      );
      if (('\n' <= lookahead && lookahead <= '\r')) SKIP(56);
      if (('%' <= lookahead && lookahead <= '+') ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 51:
      ADVANCE_MAP(
        '"', 168,
        '#', 233,
        '$', 136,
        '&', 93,
        '\'', 169,
        '(', 138,
        ':', 112,
        ';', 119,
        '\\', 9,
        '|', 118,
        '\t', 117,
        ' ', 117,
        '\n', 109,
        '\r', 109,
      );
      if (lookahead == 0x0b ||
          lookahead == '\f') SKIP(55);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 52:
      ADVANCE_MAP(
        '"', 168,
        '#', 233,
        '$', 136,
        '&', 93,
        '\'', 169,
        '(', 132,
        ')', 200,
        ':', 112,
        '\\', 26,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(53);
      if (('%' <= lookahead && lookahead <= '+') ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 53:
      ADVANCE_MAP(
        '"', 168,
        '#', 233,
        '$', 136,
        '&', 93,
        '\'', 169,
        '(', 132,
        ':', 112,
        '\\', 26,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(53);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 54:
      ADVANCE_MAP(
        '"', 168,
        '#', 233,
        '$', 136,
        '&', 93,
        '\'', 169,
        ')', 135,
        ',', 133,
        '-', 194,
        '.', 185,
        ':', 112,
        '\\', 4,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(54);
      if (lookahead == '%' ||
          ('*' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 55:
      ADVANCE_MAP(
        '"', 168,
        '#', 233,
        '$', 136,
        '&', 93,
        '\'', 169,
        ':', 112,
        ';', 119,
        '\\', 21,
        '|', 118,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(55);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 56:
      if (lookahead == '"') ADVANCE(168);
      if (lookahead == '#') ADVANCE(233);
      if (lookahead == '$') ADVANCE(136);
      if (lookahead == '&') ADVANCE(93);
      if (lookahead == '\'') ADVANCE(169);
      if (lookahead == ':') ADVANCE(112);
      if (lookahead == '\\') ADVANCE(20);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(56);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 57:
      ADVANCE_MAP(
        '"', 168,
        '#', 233,
        '$', 136,
        '\'', 169,
        '(', 138,
        ')', 135,
        ',', 133,
        ':', 110,
        '=', 123,
        '\\', 24,
        '}', 143,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(61);
      if (lookahead == '%' ||
          ('*' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 58:
      ADVANCE_MAP(
        '"', 168,
        '#', 233,
        '$', 136,
        '\'', 169,
        '(', 138,
        '+', 174,
        ':', 113,
        ';', 119,
        '=', 123,
        '?', 175,
        '\\', 9,
        '|', 118,
        '\t', 117,
        ' ', 117,
        '\n', 109,
        '\r', 109,
      );
      if (lookahead == 0x0b ||
          lookahead == '\f') SKIP(63);
      if (lookahead == '%' ||
          lookahead == '*' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('@' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 59:
      ADVANCE_MAP(
        '"', 168,
        '#', 233,
        '$', 136,
        '\'', 169,
        '(', 138,
        ':', 110,
        ';', 119,
        '\\', 23,
        '|', 118,
        '\n', 109,
        '\r', 109,
      );
      if (('\t' <= lookahead && lookahead <= '\f') ||
          lookahead == ' ') SKIP(64);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 60:
      ADVANCE_MAP(
        '"', 168,
        '#', 233,
        '$', 136,
        '\'', 169,
        '(', 138,
        ':', 165,
        ';', 167,
        '\\', 25,
        '\n', 109,
        '\r', 109,
      );
      if (('\t' <= lookahead && lookahead <= '\f') ||
          lookahead == ' ') SKIP(67);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 61:
      ADVANCE_MAP(
        '"', 168,
        '#', 233,
        '$', 136,
        '\'', 169,
        ')', 135,
        ',', 133,
        ':', 110,
        '=', 123,
        '\\', 24,
        '}', 143,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(61);
      if (lookahead == '%' ||
          ('*' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 62:
      ADVANCE_MAP(
        '"', 168,
        '#', 233,
        '$', 136,
        '\'', 169,
        '+', 174,
        ':', 113,
        ';', 119,
        '=', 123,
        '?', 175,
        '\\', 13,
        '|', 118,
        '\n', 109,
        '\r', 109,
      );
      if (('\t' <= lookahead && lookahead <= '\f') ||
          lookahead == ' ') SKIP(63);
      if (lookahead == '%' ||
          lookahead == '*' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('@' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 63:
      ADVANCE_MAP(
        '"', 168,
        '#', 233,
        '$', 136,
        '\'', 169,
        '+', 174,
        ':', 113,
        ';', 119,
        '=', 123,
        '?', 175,
        '\\', 13,
        '|', 118,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(63);
      if (lookahead == '%' ||
          lookahead == '*' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('@' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 64:
      ADVANCE_MAP(
        '"', 168,
        '#', 233,
        '$', 136,
        '\'', 169,
        ':', 110,
        ';', 119,
        '\\', 23,
        '|', 118,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(64);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 65:
      ADVANCE_MAP(
        '"', 168,
        '#', 233,
        '$', 136,
        '\'', 169,
        ';', 119,
        '\\', 22,
        '|', 118,
        '\t', 117,
        ' ', 117,
        '\n', 109,
        '\r', 109,
      );
      if (lookahead == 0x0b ||
          lookahead == '\f') SKIP(66);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 66:
      if (lookahead == '"') ADVANCE(168);
      if (lookahead == '#') ADVANCE(233);
      if (lookahead == '$') ADVANCE(136);
      if (lookahead == '\'') ADVANCE(169);
      if (lookahead == ';') ADVANCE(119);
      if (lookahead == '\\') ADVANCE(22);
      if (lookahead == '|') ADVANCE(118);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(66);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 67:
      if (lookahead == '"') ADVANCE(168);
      if (lookahead == '#') ADVANCE(233);
      if (lookahead == '$') ADVANCE(136);
      if (lookahead == '\'') ADVANCE(169);
      if (lookahead == '\\') ADVANCE(25);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(67);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 68:
      if (lookahead == '"') ADVANCE(168);
      if (lookahead == '#') ADVANCE(172);
      if (lookahead == '$') ADVANCE(136);
      if (lookahead == '\\') ADVANCE(31);
      if (lookahead == '\n' ||
          lookahead == '\r') SKIP(68);
      if (('\t' <= lookahead && lookahead <= '\f') ||
          lookahead == ' ') ADVANCE(170);
      if (lookahead != 0 &&
          lookahead != '\'') ADVANCE(173);
      END_STATE();
    case 69:
      if (lookahead == '#') ADVANCE(233);
      if (lookahead == '$') ADVANCE(136);
      if (lookahead == ')') ADVANCE(135);
      if (lookahead == ',') ADVANCE(133);
      if (lookahead == '/') ADVANCE(92);
      if (lookahead == '\\') SKIP(29);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(69);
      END_STATE();
    case 70:
      if (lookahead == '#') ADVANCE(233);
      if (lookahead == '$') ADVANCE(136);
      if (lookahead == '/') ADVANCE(92);
      if (lookahead == '\\') ADVANCE(8);
      if (lookahead == '\n' ||
          lookahead == '\r') ADVANCE(109);
      if (('\t' <= lookahead && lookahead <= '\f') ||
          lookahead == ' ') SKIP(73);
      END_STATE();
    case 71:
      if (lookahead == '#') ADVANCE(233);
      if (lookahead == '$') ADVANCE(136);
      if (lookahead == '/') ADVANCE(92);
      if (lookahead == '\\') ADVANCE(8);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(73);
      END_STATE();
    case 72:
      if (lookahead == '#') ADVANCE(233);
      if (lookahead == '$') ADVANCE(136);
      if (lookahead == '/') ADVANCE(92);
      if (lookahead == '\\') SKIP(30);
      if (lookahead == '\n' ||
          lookahead == '\r') ADVANCE(109);
      if (('\t' <= lookahead && lookahead <= '\f') ||
          lookahead == ' ') SKIP(73);
      END_STATE();
    case 73:
      if (lookahead == '#') ADVANCE(233);
      if (lookahead == '$') ADVANCE(136);
      if (lookahead == '/') ADVANCE(92);
      if (lookahead == '\\') SKIP(30);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(73);
      END_STATE();
    case 74:
      if (lookahead == '#') ADVANCE(233);
      if (lookahead == '+') ADVANCE(97);
      if (lookahead == ':') ADVANCE(94);
      if (lookahead == '=') ADVANCE(123);
      if (lookahead == '?') ADVANCE(98);
      if (lookahead == '\\') SKIP(35);
      if (lookahead == '\t' ||
          lookahead == ' ') ADVANCE(117);
      if (lookahead == '\n' ||
          lookahead == '\r') ADVANCE(109);
      if (lookahead == 0x0b ||
          lookahead == '\f') SKIP(76);
      END_STATE();
    case 75:
      if (lookahead == '#') ADVANCE(233);
      if (lookahead == '+') ADVANCE(97);
      if (lookahead == ':') ADVANCE(94);
      if (lookahead == '=') ADVANCE(123);
      if (lookahead == '?') ADVANCE(98);
      if (lookahead == '\\') SKIP(35);
      if (lookahead == '\t' ||
          lookahead == ' ') ADVANCE(117);
      if (('\n' <= lookahead && lookahead <= '\r')) SKIP(76);
      END_STATE();
    case 76:
      if (lookahead == '#') ADVANCE(233);
      if (lookahead == '+') ADVANCE(97);
      if (lookahead == ':') ADVANCE(94);
      if (lookahead == '=') ADVANCE(123);
      if (lookahead == '?') ADVANCE(98);
      if (lookahead == '\\') SKIP(35);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(76);
      END_STATE();
    case 77:
      if (lookahead == '#') ADVANCE(233);
      if (lookahead == '\\') SKIP(33);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(77);
      END_STATE();
    case 78:
      if (lookahead == '#') ADVANCE(233);
      if (lookahead == '\\') ADVANCE(43);
      if (lookahead == '\t' ||
          lookahead == ' ') ADVANCE(117);
      if (('\n' <= lookahead && lookahead <= '\r')) SKIP(79);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 79:
      if (lookahead == '#') ADVANCE(233);
      if (lookahead == '\\') ADVANCE(43);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(79);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 80:
      if (lookahead == '#') ADVANCE(225);
      if (lookahead == '$') ADVANCE(136);
      if (lookahead == ')') ADVANCE(135);
      if (lookahead == ',') ADVANCE(134);
      if (lookahead == '/') ADVANCE(223);
      if (lookahead == '\\') ADVANCE(15);
      if (lookahead == '\n' ||
          lookahead == '\r') SKIP(80);
      if (('\t' <= lookahead && lookahead <= '\f') ||
          lookahead == ' ') ADVANCE(219);
      if (lookahead != 0 &&
          lookahead != '(' &&
          lookahead != ')') ADVANCE(227);
      END_STATE();
    case 81:
      if (lookahead == '#') ADVANCE(225);
      if (lookahead == '$') ADVANCE(136);
      if (lookahead == ')') ADVANCE(135);
      if (lookahead == '/') ADVANCE(223);
      if (lookahead == '\\') ADVANCE(19);
      if (lookahead == '\n' ||
          lookahead == '\r') SKIP(81);
      if (('\t' <= lookahead && lookahead <= '\f') ||
          lookahead == ' ') ADVANCE(221);
      if (lookahead != 0 &&
          lookahead != '(' &&
          lookahead != ')') ADVANCE(227);
      END_STATE();
    case 82:
      ADVANCE_MAP(
        '#', 225,
        '$', 136,
        '/', 223,
        '\\', 19,
        '\t', 117,
        ' ', 117,
        '\n', 109,
        '\r', 109,
        0x0b, 221,
        '\f', 221,
      );
      if (lookahead != 0 &&
          lookahead != '(' &&
          lookahead != ')') ADVANCE(227);
      END_STATE();
    case 83:
      if (lookahead == '#') ADVANCE(225);
      if (lookahead == '$') ADVANCE(136);
      if (lookahead == '/') ADVANCE(223);
      if (lookahead == '\\') ADVANCE(19);
      if (lookahead == '\t' ||
          lookahead == ' ') ADVANCE(117);
      if (lookahead == '\n' ||
          lookahead == '\r') SKIP(85);
      if (lookahead == 0x0b ||
          lookahead == '\f') ADVANCE(221);
      if (lookahead != 0 &&
          lookahead != '(' &&
          lookahead != ')') ADVANCE(227);
      END_STATE();
    case 84:
      ADVANCE_MAP(
        '#', 225,
        '$', 136,
        '/', 223,
        '\\', 19,
        '\n', 109,
        '\r', 109,
        '\t', 221,
        0x0b, 221,
        '\f', 221,
        ' ', 221,
      );
      if (lookahead != 0 &&
          lookahead != '(' &&
          lookahead != ')') ADVANCE(227);
      END_STATE();
    case 85:
      if (lookahead == '#') ADVANCE(225);
      if (lookahead == '$') ADVANCE(136);
      if (lookahead == '/') ADVANCE(223);
      if (lookahead == '\\') ADVANCE(19);
      if (lookahead == '\n' ||
          lookahead == '\r') SKIP(85);
      if (('\t' <= lookahead && lookahead <= '\f') ||
          lookahead == ' ') ADVANCE(221);
      if (lookahead != 0 &&
          lookahead != '(' &&
          lookahead != ')') ADVANCE(227);
      END_STATE();
    case 86:
      ADVANCE_MAP(
        '#', 212,
        '$', 136,
        '+', 122,
        '-', 121,
        '/', 210,
        '@', 120,
        '\\', 27,
        '\n', 108,
        '\r', 108,
        '\t', 208,
        0x0b, 208,
        '\f', 208,
        ' ', 208,
      );
      if (lookahead != 0) ADVANCE(214);
      END_STATE();
    case 87:
      if (lookahead == '#') ADVANCE(212);
      if (lookahead == '$') ADVANCE(136);
      if (lookahead == '+') ADVANCE(122);
      if (lookahead == '-') ADVANCE(121);
      if (lookahead == '/') ADVANCE(210);
      if (lookahead == '@') ADVANCE(120);
      if (lookahead == '\\') ADVANCE(27);
      if (lookahead == '\n' ||
          lookahead == '\r') SKIP(87);
      if (('\t' <= lookahead && lookahead <= '\f') ||
          lookahead == ' ') ADVANCE(208);
      if (lookahead != 0) ADVANCE(214);
      END_STATE();
    case 88:
      if (lookahead == '#') ADVANCE(212);
      if (lookahead == '$') ADVANCE(136);
      if (lookahead == '/') ADVANCE(210);
      if (lookahead == '\\') ADVANCE(17);
      if (lookahead == '\n' ||
          lookahead == '\r') SKIP(88);
      if (('\t' <= lookahead && lookahead <= '\f') ||
          lookahead == ' ') ADVANCE(209);
      if (lookahead != 0) ADVANCE(214);
      END_STATE();
    case 89:
      ADVANCE_MAP(
        '#', 212,
        '$', 136,
        '/', 210,
        '\\', 10,
        '\n', 109,
        '\r', 109,
        '\t', 209,
        0x0b, 209,
        '\f', 209,
        ' ', 209,
      );
      if (lookahead != 0) ADVANCE(214);
      END_STATE();
    case 90:
      if (lookahead == '#') ADVANCE(212);
      if (lookahead == '$') ADVANCE(136);
      if (lookahead == '/') ADVANCE(210);
      if (lookahead == '\\') ADVANCE(10);
      if (lookahead == '\n' ||
          lookahead == '\r') SKIP(88);
      if (('\t' <= lookahead && lookahead <= '\f') ||
          lookahead == ' ') ADVANCE(209);
      if (lookahead != 0) ADVANCE(214);
      END_STATE();
    case 91:
      if (lookahead == '#') ADVANCE(172);
      if (lookahead == '$') ADVANCE(136);
      if (lookahead == '\'') ADVANCE(169);
      if (lookahead == '\\') ADVANCE(32);
      if (lookahead == '\n' ||
          lookahead == '\r') SKIP(91);
      if (('\t' <= lookahead && lookahead <= '\f') ||
          lookahead == ' ') ADVANCE(171);
      if (lookahead != 0 &&
          (lookahead < '"' || '$' < lookahead)) ADVANCE(173);
      END_STATE();
    case 92:
      if (lookahead == '/') ADVANCE(215);
      END_STATE();
    case 93:
      if (lookahead == ':') ADVANCE(114);
      END_STATE();
    case 94:
      if (lookahead == ':') ADVANCE(96);
      if (lookahead == '=') ADVANCE(124);
      END_STATE();
    case 95:
      if (lookahead == '=') ADVANCE(129);
      END_STATE();
    case 96:
      if (lookahead == '=') ADVANCE(125);
      END_STATE();
    case 97:
      if (lookahead == '=') ADVANCE(127);
      END_STATE();
    case 98:
      if (lookahead == '=') ADVANCE(126);
      END_STATE();
    case 99:
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(199);
      END_STATE();
    case 100:
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(101);
      if (set_contains(sym_word_character_set_1, 15, lookahead)) ADVANCE(199);
      END_STATE();
    case 101:
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(99);
      END_STATE();
    case 102:
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(213);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r') ADVANCE(214);
      END_STATE();
    case 103:
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r') ADVANCE(173);
      END_STATE();
    case 104:
      if (eof) ADVANCE(107);
      ADVANCE_MAP(
        '\t', 201,
        '"', 168,
        '#', 233,
        '$', 136,
        '\'', 169,
        '-', 194,
        '.', 185,
        '\\', 5,
      );
      if (('\n' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(104);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('/' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 105:
      if (eof) ADVANCE(107);
      ADVANCE_MAP(
        '!', 95,
        '"', 168,
        '#', 233,
        '$', 136,
        '%', 157,
        '&', 93,
        '\'', 169,
        '(', 132,
        ')', 135,
        '*', 162,
        '+', 122,
        ',', 133,
        '-', 121,
        '.', 185,
        '/', 161,
        ':', 111,
        ';', 119,
        '<', 158,
        '=', 123,
        '?', 159,
        '@', 120,
        '\\', 3,
        '^', 160,
        'e', 196,
        '|', 118,
        '}', 143,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(105);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 106:
      if (eof) ADVANCE(107);
      ADVANCE_MAP(
        '"', 168,
        '#', 233,
        '$', 136,
        '&', 93,
        '\'', 169,
        ')', 135,
        ',', 133,
        '-', 194,
        '.', 185,
        ':', 112,
        '\\', 4,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(106);
      if (lookahead == '%' ||
          ('*' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 107:
      ACCEPT_TOKEN(ts_builtin_sym_end);
      END_STATE();
    case 108:
      ACCEPT_TOKEN(aux_sym__thing_token1);
      if (lookahead == '+') ADVANCE(122);
      if (lookahead == '-') ADVANCE(121);
      if (lookahead == '@') ADVANCE(120);
      if (lookahead == '\n' ||
          lookahead == '\r') ADVANCE(108);
      END_STATE();
    case 109:
      ACCEPT_TOKEN(aux_sym__thing_token1);
      if (lookahead == '\n' ||
          lookahead == '\r') ADVANCE(109);
      END_STATE();
    case 110:
      ACCEPT_TOKEN(anon_sym_COLON);
      END_STATE();
    case 111:
      ACCEPT_TOKEN(anon_sym_COLON);
      if (lookahead == ':') ADVANCE(116);
      if (lookahead == '=') ADVANCE(124);
      END_STATE();
    case 112:
      ACCEPT_TOKEN(anon_sym_COLON);
      if (lookahead == ':') ADVANCE(115);
      END_STATE();
    case 113:
      ACCEPT_TOKEN(anon_sym_COLON);
      if (lookahead == ':') ADVANCE(96);
      if (lookahead == '=') ADVANCE(124);
      END_STATE();
    case 114:
      ACCEPT_TOKEN(anon_sym_AMP_COLON);
      END_STATE();
    case 115:
      ACCEPT_TOKEN(anon_sym_COLON_COLON);
      END_STATE();
    case 116:
      ACCEPT_TOKEN(anon_sym_COLON_COLON);
      if (lookahead == '=') ADVANCE(125);
      END_STATE();
    case 117:
      ACCEPT_TOKEN(aux_sym__ordinary_rule_token1);
      if (lookahead == '\t' ||
          lookahead == ' ') ADVANCE(117);
      END_STATE();
    case 118:
      ACCEPT_TOKEN(anon_sym_PIPE);
      END_STATE();
    case 119:
      ACCEPT_TOKEN(anon_sym_SEMI);
      END_STATE();
    case 120:
      ACCEPT_TOKEN(anon_sym_AT);
      END_STATE();
    case 121:
      ACCEPT_TOKEN(anon_sym_DASH);
      END_STATE();
    case 122:
      ACCEPT_TOKEN(anon_sym_PLUS);
      END_STATE();
    case 123:
      ACCEPT_TOKEN(anon_sym_EQ);
      END_STATE();
    case 124:
      ACCEPT_TOKEN(anon_sym_COLON_EQ);
      END_STATE();
    case 125:
      ACCEPT_TOKEN(anon_sym_COLON_COLON_EQ);
      END_STATE();
    case 126:
      ACCEPT_TOKEN(anon_sym_QMARK_EQ);
      END_STATE();
    case 127:
      ACCEPT_TOKEN(anon_sym_PLUS_EQ);
      END_STATE();
    case 128:
      ACCEPT_TOKEN(anon_sym_DOTRECIPEPREFIX);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 129:
      ACCEPT_TOKEN(anon_sym_BANG_EQ);
      END_STATE();
    case 130:
      ACCEPT_TOKEN(anon_sym_endef);
      END_STATE();
    case 131:
      ACCEPT_TOKEN(anon_sym_DASHinclude);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 132:
      ACCEPT_TOKEN(anon_sym_LPAREN);
      END_STATE();
    case 133:
      ACCEPT_TOKEN(anon_sym_COMMA);
      END_STATE();
    case 134:
      ACCEPT_TOKEN(anon_sym_COMMA);
      if (lookahead == '\\') ADVANCE(44);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '$' &&
          lookahead != '(' &&
          lookahead != ')') ADVANCE(227);
      END_STATE();
    case 135:
      ACCEPT_TOKEN(anon_sym_RPAREN);
      END_STATE();
    case 136:
      ACCEPT_TOKEN(anon_sym_DOLLAR);
      if (lookahead == '$') ADVANCE(137);
      END_STATE();
    case 137:
      ACCEPT_TOKEN(anon_sym_DOLLAR_DOLLAR);
      END_STATE();
    case 138:
      ACCEPT_TOKEN(anon_sym_LPAREN2);
      END_STATE();
    case 139:
      ACCEPT_TOKEN(anon_sym_LPAREN2);
      if (lookahead == '\\') ADVANCE(102);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '$') ADVANCE(214);
      END_STATE();
    case 140:
      ACCEPT_TOKEN(anon_sym_LBRACE);
      END_STATE();
    case 141:
      ACCEPT_TOKEN(anon_sym_LBRACE);
      if (lookahead == '\\') ADVANCE(44);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '$' &&
          lookahead != '(' &&
          lookahead != ')') ADVANCE(227);
      END_STATE();
    case 142:
      ACCEPT_TOKEN(anon_sym_LBRACE);
      if (lookahead == '\\') ADVANCE(102);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '$') ADVANCE(214);
      END_STATE();
    case 143:
      ACCEPT_TOKEN(anon_sym_RBRACE);
      END_STATE();
    case 144:
      ACCEPT_TOKEN(aux_sym_variable_reference_token1);
      END_STATE();
    case 145:
      ACCEPT_TOKEN(aux_sym_variable_reference_token1);
      if (lookahead == '\\') ADVANCE(230);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '$' &&
          lookahead != '(' &&
          lookahead != ')') ADVANCE(225);
      END_STATE();
    case 146:
      ACCEPT_TOKEN(aux_sym_variable_reference_token1);
      if (lookahead == '\\') ADVANCE(44);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '$' &&
          lookahead != '(' &&
          lookahead != ')') ADVANCE(227);
      END_STATE();
    case 147:
      ACCEPT_TOKEN(aux_sym_variable_reference_token1);
      if (lookahead == '\\') ADVANCE(231);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '$') ADVANCE(212);
      END_STATE();
    case 148:
      ACCEPT_TOKEN(aux_sym_variable_reference_token1);
      if (lookahead == '\\') ADVANCE(102);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '$') ADVANCE(214);
      END_STATE();
    case 149:
      ACCEPT_TOKEN(anon_sym_AT2);
      END_STATE();
    case 150:
      ACCEPT_TOKEN(anon_sym_PERCENT);
      END_STATE();
    case 151:
      ACCEPT_TOKEN(anon_sym_LT);
      END_STATE();
    case 152:
      ACCEPT_TOKEN(anon_sym_QMARK);
      END_STATE();
    case 153:
      ACCEPT_TOKEN(anon_sym_CARET);
      END_STATE();
    case 154:
      ACCEPT_TOKEN(anon_sym_PLUS2);
      END_STATE();
    case 155:
      ACCEPT_TOKEN(anon_sym_SLASH);
      END_STATE();
    case 156:
      ACCEPT_TOKEN(anon_sym_STAR);
      END_STATE();
    case 157:
      ACCEPT_TOKEN(anon_sym_PERCENT2);
      END_STATE();
    case 158:
      ACCEPT_TOKEN(anon_sym_LT2);
      END_STATE();
    case 159:
      ACCEPT_TOKEN(anon_sym_QMARK2);
      END_STATE();
    case 160:
      ACCEPT_TOKEN(anon_sym_CARET2);
      END_STATE();
    case 161:
      ACCEPT_TOKEN(anon_sym_SLASH2);
      END_STATE();
    case 162:
      ACCEPT_TOKEN(anon_sym_STAR2);
      END_STATE();
    case 163:
      ACCEPT_TOKEN(aux_sym_list_token1);
      END_STATE();
    case 164:
      ACCEPT_TOKEN(aux_sym_list_token1);
      if (lookahead == '\n') ADVANCE(163);
      END_STATE();
    case 165:
      ACCEPT_TOKEN(anon_sym_COLON2);
      END_STATE();
    case 166:
      ACCEPT_TOKEN(anon_sym_COLON2);
      if (lookahead == ':') ADVANCE(116);
      if (lookahead == '=') ADVANCE(124);
      END_STATE();
    case 167:
      ACCEPT_TOKEN(anon_sym_SEMI2);
      END_STATE();
    case 168:
      ACCEPT_TOKEN(anon_sym_DQUOTE);
      END_STATE();
    case 169:
      ACCEPT_TOKEN(anon_sym_SQUOTE);
      END_STATE();
    case 170:
      ACCEPT_TOKEN(aux_sym__string_token1);
      if (lookahead == '"') ADVANCE(168);
      if (lookahead == '#') ADVANCE(172);
      if (lookahead == '$') ADVANCE(136);
      if (lookahead == '\\') ADVANCE(31);
      if (lookahead == '\n' ||
          lookahead == '\r') SKIP(68);
      if (('\t' <= lookahead && lookahead <= '\f') ||
          lookahead == ' ') ADVANCE(170);
      if (lookahead != 0 &&
          lookahead != '\'') ADVANCE(173);
      END_STATE();
    case 171:
      ACCEPT_TOKEN(aux_sym__string_token1);
      if (lookahead == '#') ADVANCE(172);
      if (lookahead == '$') ADVANCE(136);
      if (lookahead == '\'') ADVANCE(169);
      if (lookahead == '\\') ADVANCE(32);
      if (lookahead == '\n' ||
          lookahead == '\r') SKIP(91);
      if (('\t' <= lookahead && lookahead <= '\f') ||
          lookahead == ' ') ADVANCE(171);
      if (lookahead != 0 &&
          (lookahead < '"' || '$' < lookahead)) ADVANCE(173);
      END_STATE();
    case 172:
      ACCEPT_TOKEN(aux_sym__string_token1);
      if (lookahead == '\\') ADVANCE(232);
      if (lookahead == '\r' ||
          lookahead == '"' ||
          lookahead == '$' ||
          lookahead == '\'') ADVANCE(233);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(172);
      END_STATE();
    case 173:
      ACCEPT_TOKEN(aux_sym__string_token1);
      if (lookahead == '\\') ADVANCE(103);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '"' &&
          lookahead != '$' &&
          lookahead != '\'') ADVANCE(173);
      END_STATE();
    case 174:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == '=') ADVANCE(127);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 175:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == '=') ADVANCE(126);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 176:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == 'C') ADVANCE(181);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 177:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == 'E') ADVANCE(176);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 178:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == 'E') ADVANCE(180);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 179:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == 'E') ADVANCE(184);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 180:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == 'F') ADVANCE(182);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 181:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == 'I') ADVANCE(183);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 182:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == 'I') ADVANCE(187);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 183:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == 'P') ADVANCE(179);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 184:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == 'P') ADVANCE(186);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 185:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == 'R') ADVANCE(177);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 186:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == 'R') ADVANCE(178);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 187:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == 'X') ADVANCE(128);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 188:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == 'c') ADVANCE(195);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 189:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == 'd') ADVANCE(191);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 190:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == 'd') ADVANCE(192);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 191:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == 'e') ADVANCE(193);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 192:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == 'e') ADVANCE(131);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 193:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == 'f') ADVANCE(130);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 194:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == 'i') ADVANCE(197);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 195:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == 'l') ADVANCE(198);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 196:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == 'n') ADVANCE(189);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 197:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == 'n') ADVANCE(188);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 198:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == 'u') ADVANCE(190);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 199:
      ACCEPT_TOKEN(sym_word);
      if (lookahead == '\\') ADVANCE(100);
      if (lookahead == '%' ||
          lookahead == '*' ||
          lookahead == '+' ||
          ('-' <= lookahead && lookahead <= '9') ||
          ('?' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(199);
      END_STATE();
    case 200:
      ACCEPT_TOKEN(anon_sym_RPAREN2);
      END_STATE();
    case 201:
      ACCEPT_TOKEN(anon_sym_TAB);
      if (lookahead == '\t') ADVANCE(201);
      if (lookahead == '\\') ADVANCE(5);
      END_STATE();
    case 202:
      ACCEPT_TOKEN(anon_sym_TAB);
      if (lookahead == '\t') ADVANCE(202);
      if (lookahead == '\\') ADVANCE(28);
      if (lookahead == 0x0b ||
          lookahead == '\f' ||
          lookahead == ' ') ADVANCE(207);
      END_STATE();
    case 203:
      ACCEPT_TOKEN(sym__rawline);
      ADVANCE_MAP(
        '\n', 203,
        '\r', 203,
        '#', 229,
        '\\', 37,
        'e', 41,
        '\t', 36,
        0x0b, 36,
        '\f', 36,
        ' ', 36,
      );
      if (lookahead != 0) ADVANCE(42);
      END_STATE();
    case 204:
      ACCEPT_TOKEN(sym__rawline);
      if (lookahead == '\n') ADVANCE(206);
      if (lookahead == '\r') ADVANCE(204);
      if (lookahead != 0) ADVANCE(229);
      END_STATE();
    case 205:
      ACCEPT_TOKEN(sym__rawline);
      if (lookahead == '\n') ADVANCE(206);
      if (lookahead == '\r') ADVANCE(205);
      if (lookahead != 0) ADVANCE(42);
      END_STATE();
    case 206:
      ACCEPT_TOKEN(sym__rawline);
      if (lookahead == '\n' ||
          lookahead == '\r') ADVANCE(206);
      END_STATE();
    case 207:
      ACCEPT_TOKEN(aux_sym__shell_text_without_split_token1);
      if (lookahead == '\t') ADVANCE(202);
      if (lookahead == '#') ADVANCE(212);
      if (lookahead == '/') ADVANCE(210);
      if (lookahead == '\\') ADVANCE(28);
      if (lookahead == 0x0b ||
          lookahead == '\f' ||
          lookahead == ' ') ADVANCE(207);
      if (lookahead != 0 &&
          (lookahead < '\t' || '\r' < lookahead) &&
          lookahead != '#' &&
          lookahead != '$') ADVANCE(214);
      END_STATE();
    case 208:
      ACCEPT_TOKEN(aux_sym__shell_text_without_split_token1);
      ADVANCE_MAP(
        '#', 212,
        '+', 122,
        '-', 121,
        '/', 210,
        '@', 120,
        '\\', 27,
        '\t', 208,
        0x0b, 208,
        '\f', 208,
        ' ', 208,
      );
      if (lookahead != 0 &&
          (lookahead < '\t' || '\r' < lookahead) &&
          lookahead != '#' &&
          lookahead != '$') ADVANCE(214);
      END_STATE();
    case 209:
      ACCEPT_TOKEN(aux_sym__shell_text_without_split_token1);
      if (lookahead == '#') ADVANCE(212);
      if (lookahead == '/') ADVANCE(210);
      if (lookahead == '\\') ADVANCE(17);
      if (lookahead == '\t' ||
          lookahead == 0x0b ||
          lookahead == '\f' ||
          lookahead == ' ') ADVANCE(209);
      if (lookahead != 0 &&
          (lookahead < '\t' || '\r' < lookahead) &&
          lookahead != '#' &&
          lookahead != '$') ADVANCE(214);
      END_STATE();
    case 210:
      ACCEPT_TOKEN(aux_sym__shell_text_without_split_token1);
      if (lookahead == '/') ADVANCE(217);
      if (lookahead == '\\') ADVANCE(102);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '$') ADVANCE(214);
      END_STATE();
    case 211:
      ACCEPT_TOKEN(aux_sym__shell_text_without_split_token1);
      if (lookahead == '\\') ADVANCE(231);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(212);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '$') ADVANCE(212);
      END_STATE();
    case 212:
      ACCEPT_TOKEN(aux_sym__shell_text_without_split_token1);
      if (lookahead == '\\') ADVANCE(231);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '$') ADVANCE(212);
      END_STATE();
    case 213:
      ACCEPT_TOKEN(aux_sym__shell_text_without_split_token1);
      if (lookahead == '\\') ADVANCE(102);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(214);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '$') ADVANCE(214);
      END_STATE();
    case 214:
      ACCEPT_TOKEN(aux_sym__shell_text_without_split_token1);
      if (lookahead == '\\') ADVANCE(102);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '$') ADVANCE(214);
      END_STATE();
    case 215:
      ACCEPT_TOKEN(anon_sym_SLASH_SLASH);
      END_STATE();
    case 216:
      ACCEPT_TOKEN(anon_sym_SLASH_SLASH);
      if (lookahead == '\\') ADVANCE(44);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '$' &&
          lookahead != '(' &&
          lookahead != ')') ADVANCE(227);
      END_STATE();
    case 217:
      ACCEPT_TOKEN(anon_sym_SLASH_SLASH);
      if (lookahead == '\\') ADVANCE(102);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '$') ADVANCE(214);
      END_STATE();
    case 218:
      ACCEPT_TOKEN(aux_sym_text_token1);
      if (lookahead == '\n') ADVANCE(227);
      if (lookahead == '\\') ADVANCE(230);
      if (lookahead != 0 &&
          lookahead != '\r' &&
          lookahead != '$' &&
          lookahead != '(' &&
          lookahead != ')') ADVANCE(225);
      END_STATE();
    case 219:
      ACCEPT_TOKEN(aux_sym_text_token1);
      if (lookahead == '#') ADVANCE(225);
      if (lookahead == ',') ADVANCE(134);
      if (lookahead == '/') ADVANCE(223);
      if (lookahead == '\\') ADVANCE(15);
      if (lookahead == '\t' ||
          lookahead == 0x0b ||
          lookahead == '\f' ||
          lookahead == ' ') ADVANCE(219);
      if (lookahead != 0 &&
          (lookahead < '\t' || '\r' < lookahead) &&
          lookahead != '#' &&
          lookahead != '$' &&
          lookahead != '(' &&
          lookahead != ')') ADVANCE(227);
      END_STATE();
    case 220:
      ACCEPT_TOKEN(aux_sym_text_token1);
      if (lookahead == '#') ADVANCE(225);
      if (lookahead == ',') ADVANCE(134);
      if (lookahead == '/') ADVANCE(223);
      if (lookahead == '\\') ADVANCE(15);
      if (('\t' <= lookahead && lookahead <= '\f') ||
          lookahead == ' ') ADVANCE(219);
      if (lookahead != 0 &&
          (lookahead < '\t' || '\r' < lookahead) &&
          lookahead != '#' &&
          lookahead != '$' &&
          lookahead != '(' &&
          lookahead != ')') ADVANCE(227);
      END_STATE();
    case 221:
      ACCEPT_TOKEN(aux_sym_text_token1);
      if (lookahead == '#') ADVANCE(225);
      if (lookahead == '/') ADVANCE(223);
      if (lookahead == '\\') ADVANCE(19);
      if (lookahead == '\t' ||
          lookahead == 0x0b ||
          lookahead == '\f' ||
          lookahead == ' ') ADVANCE(221);
      if (lookahead != 0 &&
          (lookahead < '\t' || '\r' < lookahead) &&
          lookahead != '#' &&
          lookahead != '$' &&
          lookahead != '(' &&
          lookahead != ')') ADVANCE(227);
      END_STATE();
    case 222:
      ACCEPT_TOKEN(aux_sym_text_token1);
      if (lookahead == '#') ADVANCE(225);
      if (lookahead == '/') ADVANCE(223);
      if (lookahead == '\\') ADVANCE(19);
      if (('\t' <= lookahead && lookahead <= '\f') ||
          lookahead == ' ') ADVANCE(221);
      if (lookahead != 0 &&
          (lookahead < '\t' || '\r' < lookahead) &&
          lookahead != '#' &&
          lookahead != '$' &&
          lookahead != '(' &&
          lookahead != ')') ADVANCE(227);
      END_STATE();
    case 223:
      ACCEPT_TOKEN(aux_sym_text_token1);
      if (lookahead == '/') ADVANCE(216);
      if (lookahead == '\\') ADVANCE(44);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '$' &&
          lookahead != '(' &&
          lookahead != ')') ADVANCE(227);
      END_STATE();
    case 224:
      ACCEPT_TOKEN(aux_sym_text_token1);
      if (lookahead == '\\') ADVANCE(230);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(225);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '$' &&
          lookahead != '(' &&
          lookahead != ')') ADVANCE(225);
      END_STATE();
    case 225:
      ACCEPT_TOKEN(aux_sym_text_token1);
      if (lookahead == '\\') ADVANCE(230);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '$' &&
          lookahead != '(' &&
          lookahead != ')') ADVANCE(225);
      END_STATE();
    case 226:
      ACCEPT_TOKEN(aux_sym_text_token1);
      if (lookahead == '\\') ADVANCE(44);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(227);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '$' &&
          lookahead != '(' &&
          lookahead != ')') ADVANCE(227);
      END_STATE();
    case 227:
      ACCEPT_TOKEN(aux_sym_text_token1);
      if (lookahead == '\\') ADVANCE(44);
      if (lookahead != 0 &&
          lookahead != '\n' &&
          lookahead != '\r' &&
          lookahead != '$' &&
          lookahead != '(' &&
          lookahead != ')') ADVANCE(227);
      END_STATE();
    case 228:
      ACCEPT_TOKEN(aux_sym_text_token1);
      if (lookahead == '\\') ADVANCE(44);
      if (lookahead != 0 &&
          lookahead != '\r' &&
          lookahead != '$' &&
          lookahead != '(' &&
          lookahead != ')') ADVANCE(227);
      END_STATE();
    case 229:
      ACCEPT_TOKEN(sym_comment);
      if (lookahead == '\n') ADVANCE(206);
      if (lookahead == '\r') ADVANCE(204);
      if (lookahead != 0) ADVANCE(229);
      END_STATE();
    case 230:
      ACCEPT_TOKEN(sym_comment);
      if (lookahead == '\n') ADVANCE(227);
      if (lookahead == '\r') ADVANCE(218);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(224);
      if (lookahead != 0) ADVANCE(225);
      END_STATE();
    case 231:
      ACCEPT_TOKEN(sym_comment);
      if (lookahead == '\r') ADVANCE(233);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(211);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(212);
      END_STATE();
    case 232:
      ACCEPT_TOKEN(sym_comment);
      if (lookahead == '\r') ADVANCE(233);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(172);
      END_STATE();
    case 233:
      ACCEPT_TOKEN(sym_comment);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(233);
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
      if (lookahead == 'D') ADVANCE(1);
      if (lookahead == 'F') ADVANCE(2);
      if (lookahead == 'V') ADVANCE(3);
      if (lookahead == '\\') SKIP(4);
      if (lookahead == 'a') ADVANCE(5);
      if (lookahead == 'b') ADVANCE(6);
      if (lookahead == 'c') ADVANCE(7);
      if (lookahead == 'd') ADVANCE(8);
      if (lookahead == 'e') ADVANCE(9);
      if (lookahead == 'f') ADVANCE(10);
      if (lookahead == 'i') ADVANCE(11);
      if (lookahead == 'j') ADVANCE(12);
      if (lookahead == 'l') ADVANCE(13);
      if (lookahead == 'n') ADVANCE(14);
      if (lookahead == 'o') ADVANCE(15);
      if (lookahead == 'p') ADVANCE(16);
      if (lookahead == 'r') ADVANCE(17);
      if (lookahead == 's') ADVANCE(18);
      if (lookahead == 'u') ADVANCE(19);
      if (lookahead == 'v') ADVANCE(20);
      if (lookahead == 'w') ADVANCE(21);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(22);
      END_STATE();
    case 1:
      ACCEPT_TOKEN(anon_sym_D);
      END_STATE();
    case 2:
      ACCEPT_TOKEN(anon_sym_F);
      END_STATE();
    case 3:
      if (lookahead == 'P') ADVANCE(23);
      END_STATE();
    case 4:
      if (lookahead == '\n') SKIP(22);
      if (lookahead == '\r') SKIP(24);
      END_STATE();
    case 5:
      if (lookahead == 'b') ADVANCE(25);
      if (lookahead == 'd') ADVANCE(26);
      if (lookahead == 'n') ADVANCE(27);
      END_STATE();
    case 6:
      if (lookahead == 'a') ADVANCE(28);
      END_STATE();
    case 7:
      if (lookahead == 'a') ADVANCE(29);
      END_STATE();
    case 8:
      if (lookahead == 'e') ADVANCE(30);
      if (lookahead == 'i') ADVANCE(31);
      END_STATE();
    case 9:
      if (lookahead == 'l') ADVANCE(32);
      if (lookahead == 'n') ADVANCE(33);
      if (lookahead == 'r') ADVANCE(34);
      if (lookahead == 'v') ADVANCE(35);
      if (lookahead == 'x') ADVANCE(36);
      END_STATE();
    case 10:
      if (lookahead == 'i') ADVANCE(37);
      if (lookahead == 'l') ADVANCE(38);
      if (lookahead == 'o') ADVANCE(39);
      END_STATE();
    case 11:
      if (lookahead == 'f') ADVANCE(40);
      if (lookahead == 'n') ADVANCE(41);
      END_STATE();
    case 12:
      if (lookahead == 'o') ADVANCE(42);
      END_STATE();
    case 13:
      if (lookahead == 'a') ADVANCE(43);
      END_STATE();
    case 14:
      if (lookahead == 'o') ADVANCE(44);
      END_STATE();
    case 15:
      if (lookahead == 'r') ADVANCE(45);
      if (lookahead == 'v') ADVANCE(46);
      END_STATE();
    case 16:
      if (lookahead == 'a') ADVANCE(47);
      if (lookahead == 'r') ADVANCE(48);
      END_STATE();
    case 17:
      if (lookahead == 'e') ADVANCE(49);
      END_STATE();
    case 18:
      if (lookahead == 'h') ADVANCE(50);
      if (lookahead == 'i') ADVANCE(51);
      if (lookahead == 'o') ADVANCE(52);
      if (lookahead == 't') ADVANCE(53);
      if (lookahead == 'u') ADVANCE(54);
      END_STATE();
    case 19:
      if (lookahead == 'n') ADVANCE(55);
      END_STATE();
    case 20:
      if (lookahead == 'a') ADVANCE(56);
      if (lookahead == 'p') ADVANCE(57);
      END_STATE();
    case 21:
      if (lookahead == 'a') ADVANCE(58);
      if (lookahead == 'i') ADVANCE(59);
      if (lookahead == 'o') ADVANCE(60);
      END_STATE();
    case 22:
      if (lookahead == 'V') ADVANCE(3);
      if (lookahead == '\\') SKIP(4);
      if (lookahead == 'd') ADVANCE(61);
      if (lookahead == 'e') ADVANCE(62);
      if (lookahead == 'i') ADVANCE(63);
      if (lookahead == 'o') ADVANCE(64);
      if (lookahead == 'p') ADVANCE(65);
      if (lookahead == 's') ADVANCE(66);
      if (lookahead == 'u') ADVANCE(19);
      if (lookahead == 'v') ADVANCE(67);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(22);
      END_STATE();
    case 23:
      if (lookahead == 'A') ADVANCE(68);
      END_STATE();
    case 24:
      if (lookahead == 'V') ADVANCE(3);
      if (lookahead == '\\') SKIP(4);
      if (lookahead == 'd') ADVANCE(61);
      if (lookahead == 'e') ADVANCE(62);
      if (lookahead == 'i') ADVANCE(63);
      if (lookahead == 'o') ADVANCE(64);
      if (lookahead == 'p') ADVANCE(65);
      if (lookahead == 's') ADVANCE(66);
      if (lookahead == 'u') ADVANCE(19);
      if (lookahead == 'v') ADVANCE(67);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(22);
      END_STATE();
    case 25:
      if (lookahead == 's') ADVANCE(69);
      END_STATE();
    case 26:
      if (lookahead == 'd') ADVANCE(70);
      END_STATE();
    case 27:
      if (lookahead == 'd') ADVANCE(71);
      END_STATE();
    case 28:
      if (lookahead == 's') ADVANCE(72);
      END_STATE();
    case 29:
      if (lookahead == 'l') ADVANCE(73);
      END_STATE();
    case 30:
      if (lookahead == 'f') ADVANCE(74);
      END_STATE();
    case 31:
      if (lookahead == 'r') ADVANCE(75);
      END_STATE();
    case 32:
      if (lookahead == 's') ADVANCE(76);
      END_STATE();
    case 33:
      if (lookahead == 'd') ADVANCE(77);
      END_STATE();
    case 34:
      if (lookahead == 'r') ADVANCE(78);
      END_STATE();
    case 35:
      if (lookahead == 'a') ADVANCE(79);
      END_STATE();
    case 36:
      if (lookahead == 'p') ADVANCE(80);
      END_STATE();
    case 37:
      if (lookahead == 'l') ADVANCE(81);
      if (lookahead == 'n') ADVANCE(82);
      if (lookahead == 'r') ADVANCE(83);
      END_STATE();
    case 38:
      if (lookahead == 'a') ADVANCE(84);
      END_STATE();
    case 39:
      if (lookahead == 'r') ADVANCE(85);
      END_STATE();
    case 40:
      ACCEPT_TOKEN(anon_sym_if);
      if (lookahead == 'd') ADVANCE(86);
      if (lookahead == 'e') ADVANCE(87);
      if (lookahead == 'n') ADVANCE(88);
      END_STATE();
    case 41:
      if (lookahead == 'c') ADVANCE(89);
      if (lookahead == 'f') ADVANCE(90);
      END_STATE();
    case 42:
      if (lookahead == 'i') ADVANCE(91);
      END_STATE();
    case 43:
      if (lookahead == 's') ADVANCE(92);
      END_STATE();
    case 44:
      if (lookahead == 't') ADVANCE(93);
      END_STATE();
    case 45:
      ACCEPT_TOKEN(anon_sym_or);
      if (lookahead == 'i') ADVANCE(94);
      END_STATE();
    case 46:
      if (lookahead == 'e') ADVANCE(95);
      END_STATE();
    case 47:
      if (lookahead == 't') ADVANCE(96);
      END_STATE();
    case 48:
      if (lookahead == 'i') ADVANCE(97);
      END_STATE();
    case 49:
      if (lookahead == 'a') ADVANCE(98);
      END_STATE();
    case 50:
      if (lookahead == 'e') ADVANCE(99);
      END_STATE();
    case 51:
      if (lookahead == 'n') ADVANCE(100);
      END_STATE();
    case 52:
      if (lookahead == 'r') ADVANCE(101);
      END_STATE();
    case 53:
      if (lookahead == 'r') ADVANCE(102);
      END_STATE();
    case 54:
      if (lookahead == 'b') ADVANCE(103);
      if (lookahead == 'f') ADVANCE(104);
      END_STATE();
    case 55:
      if (lookahead == 'd') ADVANCE(105);
      if (lookahead == 'e') ADVANCE(106);
      END_STATE();
    case 56:
      if (lookahead == 'l') ADVANCE(107);
      END_STATE();
    case 57:
      if (lookahead == 'a') ADVANCE(108);
      END_STATE();
    case 58:
      if (lookahead == 'r') ADVANCE(109);
      END_STATE();
    case 59:
      if (lookahead == 'l') ADVANCE(110);
      END_STATE();
    case 60:
      if (lookahead == 'r') ADVANCE(111);
      END_STATE();
    case 61:
      if (lookahead == 'e') ADVANCE(30);
      END_STATE();
    case 62:
      if (lookahead == 'l') ADVANCE(32);
      if (lookahead == 'n') ADVANCE(33);
      if (lookahead == 'x') ADVANCE(36);
      END_STATE();
    case 63:
      if (lookahead == 'f') ADVANCE(112);
      if (lookahead == 'n') ADVANCE(113);
      END_STATE();
    case 64:
      if (lookahead == 'v') ADVANCE(46);
      END_STATE();
    case 65:
      if (lookahead == 'r') ADVANCE(48);
      END_STATE();
    case 66:
      if (lookahead == 'h') ADVANCE(50);
      if (lookahead == 'i') ADVANCE(51);
      END_STATE();
    case 67:
      if (lookahead == 'p') ADVANCE(57);
      END_STATE();
    case 68:
      if (lookahead == 'T') ADVANCE(114);
      END_STATE();
    case 69:
      if (lookahead == 'p') ADVANCE(115);
      END_STATE();
    case 70:
      if (lookahead == 'p') ADVANCE(116);
      if (lookahead == 's') ADVANCE(117);
      END_STATE();
    case 71:
      ACCEPT_TOKEN(anon_sym_and);
      END_STATE();
    case 72:
      if (lookahead == 'e') ADVANCE(118);
      END_STATE();
    case 73:
      if (lookahead == 'l') ADVANCE(119);
      END_STATE();
    case 74:
      if (lookahead == 'i') ADVANCE(120);
      END_STATE();
    case 75:
      ACCEPT_TOKEN(anon_sym_dir);
      END_STATE();
    case 76:
      if (lookahead == 'e') ADVANCE(121);
      END_STATE();
    case 77:
      if (lookahead == 'i') ADVANCE(122);
      END_STATE();
    case 78:
      if (lookahead == 'o') ADVANCE(123);
      END_STATE();
    case 79:
      if (lookahead == 'l') ADVANCE(124);
      END_STATE();
    case 80:
      if (lookahead == 'o') ADVANCE(125);
      END_STATE();
    case 81:
      if (lookahead == 'e') ADVANCE(126);
      if (lookahead == 't') ADVANCE(127);
      END_STATE();
    case 82:
      if (lookahead == 'd') ADVANCE(128);
      END_STATE();
    case 83:
      if (lookahead == 's') ADVANCE(129);
      END_STATE();
    case 84:
      if (lookahead == 'v') ADVANCE(130);
      END_STATE();
    case 85:
      if (lookahead == 'e') ADVANCE(131);
      END_STATE();
    case 86:
      if (lookahead == 'e') ADVANCE(132);
      END_STATE();
    case 87:
      if (lookahead == 'q') ADVANCE(133);
      END_STATE();
    case 88:
      if (lookahead == 'd') ADVANCE(134);
      if (lookahead == 'e') ADVANCE(135);
      END_STATE();
    case 89:
      if (lookahead == 'l') ADVANCE(136);
      END_STATE();
    case 90:
      if (lookahead == 'o') ADVANCE(137);
      END_STATE();
    case 91:
      if (lookahead == 'n') ADVANCE(138);
      END_STATE();
    case 92:
      if (lookahead == 't') ADVANCE(139);
      END_STATE();
    case 93:
      if (lookahead == 'd') ADVANCE(140);
      END_STATE();
    case 94:
      if (lookahead == 'g') ADVANCE(141);
      END_STATE();
    case 95:
      if (lookahead == 'r') ADVANCE(142);
      END_STATE();
    case 96:
      if (lookahead == 's') ADVANCE(143);
      END_STATE();
    case 97:
      if (lookahead == 'v') ADVANCE(144);
      END_STATE();
    case 98:
      if (lookahead == 'l') ADVANCE(145);
      END_STATE();
    case 99:
      if (lookahead == 'l') ADVANCE(146);
      END_STATE();
    case 100:
      if (lookahead == 'c') ADVANCE(147);
      END_STATE();
    case 101:
      if (lookahead == 't') ADVANCE(148);
      END_STATE();
    case 102:
      if (lookahead == 'i') ADVANCE(149);
      END_STATE();
    case 103:
      if (lookahead == 's') ADVANCE(150);
      END_STATE();
    case 104:
      if (lookahead == 'f') ADVANCE(151);
      END_STATE();
    case 105:
      if (lookahead == 'e') ADVANCE(152);
      END_STATE();
    case 106:
      if (lookahead == 'x') ADVANCE(153);
      END_STATE();
    case 107:
      if (lookahead == 'u') ADVANCE(154);
      END_STATE();
    case 108:
      if (lookahead == 't') ADVANCE(155);
      END_STATE();
    case 109:
      if (lookahead == 'n') ADVANCE(156);
      END_STATE();
    case 110:
      if (lookahead == 'd') ADVANCE(157);
      END_STATE();
    case 111:
      if (lookahead == 'd') ADVANCE(158);
      END_STATE();
    case 112:
      if (lookahead == 'd') ADVANCE(86);
      if (lookahead == 'e') ADVANCE(87);
      if (lookahead == 'n') ADVANCE(88);
      END_STATE();
    case 113:
      if (lookahead == 'c') ADVANCE(89);
      END_STATE();
    case 114:
      if (lookahead == 'H') ADVANCE(159);
      END_STATE();
    case 115:
      if (lookahead == 'a') ADVANCE(160);
      END_STATE();
    case 116:
      if (lookahead == 'r') ADVANCE(161);
      END_STATE();
    case 117:
      if (lookahead == 'u') ADVANCE(162);
      END_STATE();
    case 118:
      if (lookahead == 'n') ADVANCE(163);
      END_STATE();
    case 119:
      ACCEPT_TOKEN(anon_sym_call);
      END_STATE();
    case 120:
      if (lookahead == 'n') ADVANCE(164);
      END_STATE();
    case 121:
      ACCEPT_TOKEN(anon_sym_else);
      END_STATE();
    case 122:
      if (lookahead == 'f') ADVANCE(165);
      END_STATE();
    case 123:
      if (lookahead == 'r') ADVANCE(166);
      END_STATE();
    case 124:
      ACCEPT_TOKEN(anon_sym_eval);
      END_STATE();
    case 125:
      if (lookahead == 'r') ADVANCE(167);
      END_STATE();
    case 126:
      ACCEPT_TOKEN(anon_sym_file);
      END_STATE();
    case 127:
      if (lookahead == 'e') ADVANCE(168);
      END_STATE();
    case 128:
      if (lookahead == 's') ADVANCE(169);
      END_STATE();
    case 129:
      if (lookahead == 't') ADVANCE(170);
      END_STATE();
    case 130:
      if (lookahead == 'o') ADVANCE(171);
      END_STATE();
    case 131:
      if (lookahead == 'a') ADVANCE(172);
      END_STATE();
    case 132:
      if (lookahead == 'f') ADVANCE(173);
      END_STATE();
    case 133:
      ACCEPT_TOKEN(anon_sym_ifeq);
      END_STATE();
    case 134:
      if (lookahead == 'e') ADVANCE(174);
      END_STATE();
    case 135:
      if (lookahead == 'q') ADVANCE(175);
      END_STATE();
    case 136:
      if (lookahead == 'u') ADVANCE(176);
      END_STATE();
    case 137:
      ACCEPT_TOKEN(anon_sym_info);
      END_STATE();
    case 138:
      ACCEPT_TOKEN(anon_sym_join);
      END_STATE();
    case 139:
      if (lookahead == 'w') ADVANCE(177);
      END_STATE();
    case 140:
      if (lookahead == 'i') ADVANCE(178);
      END_STATE();
    case 141:
      if (lookahead == 'i') ADVANCE(179);
      END_STATE();
    case 142:
      if (lookahead == 'r') ADVANCE(180);
      END_STATE();
    case 143:
      if (lookahead == 'u') ADVANCE(181);
      END_STATE();
    case 144:
      if (lookahead == 'a') ADVANCE(182);
      END_STATE();
    case 145:
      if (lookahead == 'p') ADVANCE(183);
      END_STATE();
    case 146:
      if (lookahead == 'l') ADVANCE(184);
      END_STATE();
    case 147:
      if (lookahead == 'l') ADVANCE(185);
      END_STATE();
    case 148:
      ACCEPT_TOKEN(anon_sym_sort);
      END_STATE();
    case 149:
      if (lookahead == 'p') ADVANCE(186);
      END_STATE();
    case 150:
      if (lookahead == 't') ADVANCE(187);
      END_STATE();
    case 151:
      if (lookahead == 'i') ADVANCE(188);
      END_STATE();
    case 152:
      if (lookahead == 'f') ADVANCE(189);
      END_STATE();
    case 153:
      if (lookahead == 'p') ADVANCE(190);
      END_STATE();
    case 154:
      if (lookahead == 'e') ADVANCE(191);
      END_STATE();
    case 155:
      if (lookahead == 'h') ADVANCE(192);
      END_STATE();
    case 156:
      if (lookahead == 'i') ADVANCE(193);
      END_STATE();
    case 157:
      if (lookahead == 'c') ADVANCE(194);
      END_STATE();
    case 158:
      ACCEPT_TOKEN(anon_sym_word);
      if (lookahead == 'l') ADVANCE(195);
      if (lookahead == 's') ADVANCE(196);
      END_STATE();
    case 159:
      ACCEPT_TOKEN(anon_sym_VPATH);
      END_STATE();
    case 160:
      if (lookahead == 't') ADVANCE(197);
      END_STATE();
    case 161:
      if (lookahead == 'e') ADVANCE(198);
      END_STATE();
    case 162:
      if (lookahead == 'f') ADVANCE(199);
      END_STATE();
    case 163:
      if (lookahead == 'a') ADVANCE(200);
      END_STATE();
    case 164:
      if (lookahead == 'e') ADVANCE(201);
      END_STATE();
    case 165:
      ACCEPT_TOKEN(anon_sym_endif);
      END_STATE();
    case 166:
      ACCEPT_TOKEN(anon_sym_error);
      END_STATE();
    case 167:
      if (lookahead == 't') ADVANCE(202);
      END_STATE();
    case 168:
      if (lookahead == 'r') ADVANCE(203);
      END_STATE();
    case 169:
      if (lookahead == 't') ADVANCE(204);
      END_STATE();
    case 170:
      if (lookahead == 'w') ADVANCE(205);
      END_STATE();
    case 171:
      if (lookahead == 'r') ADVANCE(206);
      END_STATE();
    case 172:
      if (lookahead == 'c') ADVANCE(207);
      END_STATE();
    case 173:
      ACCEPT_TOKEN(anon_sym_ifdef);
      END_STATE();
    case 174:
      if (lookahead == 'f') ADVANCE(208);
      END_STATE();
    case 175:
      ACCEPT_TOKEN(anon_sym_ifneq);
      END_STATE();
    case 176:
      if (lookahead == 'd') ADVANCE(209);
      END_STATE();
    case 177:
      if (lookahead == 'o') ADVANCE(210);
      END_STATE();
    case 178:
      if (lookahead == 'r') ADVANCE(211);
      END_STATE();
    case 179:
      if (lookahead == 'n') ADVANCE(212);
      END_STATE();
    case 180:
      if (lookahead == 'i') ADVANCE(213);
      END_STATE();
    case 181:
      if (lookahead == 'b') ADVANCE(214);
      END_STATE();
    case 182:
      if (lookahead == 't') ADVANCE(215);
      END_STATE();
    case 183:
      if (lookahead == 'a') ADVANCE(216);
      END_STATE();
    case 184:
      ACCEPT_TOKEN(anon_sym_shell);
      END_STATE();
    case 185:
      if (lookahead == 'u') ADVANCE(217);
      END_STATE();
    case 186:
      ACCEPT_TOKEN(anon_sym_strip);
      END_STATE();
    case 187:
      ACCEPT_TOKEN(anon_sym_subst);
      END_STATE();
    case 188:
      if (lookahead == 'x') ADVANCE(218);
      END_STATE();
    case 189:
      if (lookahead == 'i') ADVANCE(219);
      END_STATE();
    case 190:
      if (lookahead == 'o') ADVANCE(220);
      END_STATE();
    case 191:
      ACCEPT_TOKEN(anon_sym_value);
      END_STATE();
    case 192:
      ACCEPT_TOKEN(anon_sym_vpath);
      END_STATE();
    case 193:
      if (lookahead == 'n') ADVANCE(221);
      END_STATE();
    case 194:
      if (lookahead == 'a') ADVANCE(222);
      END_STATE();
    case 195:
      if (lookahead == 'i') ADVANCE(223);
      END_STATE();
    case 196:
      ACCEPT_TOKEN(anon_sym_words);
      END_STATE();
    case 197:
      if (lookahead == 'h') ADVANCE(224);
      END_STATE();
    case 198:
      if (lookahead == 'f') ADVANCE(225);
      END_STATE();
    case 199:
      if (lookahead == 'f') ADVANCE(226);
      END_STATE();
    case 200:
      if (lookahead == 'm') ADVANCE(227);
      END_STATE();
    case 201:
      ACCEPT_TOKEN(anon_sym_define);
      END_STATE();
    case 202:
      ACCEPT_TOKEN(anon_sym_export);
      END_STATE();
    case 203:
      ACCEPT_TOKEN(anon_sym_filter);
      if (lookahead == '-') ADVANCE(228);
      END_STATE();
    case 204:
      if (lookahead == 'r') ADVANCE(229);
      END_STATE();
    case 205:
      if (lookahead == 'o') ADVANCE(230);
      END_STATE();
    case 206:
      ACCEPT_TOKEN(anon_sym_flavor);
      END_STATE();
    case 207:
      if (lookahead == 'h') ADVANCE(231);
      END_STATE();
    case 208:
      ACCEPT_TOKEN(anon_sym_ifndef);
      END_STATE();
    case 209:
      if (lookahead == 'e') ADVANCE(232);
      END_STATE();
    case 210:
      if (lookahead == 'r') ADVANCE(233);
      END_STATE();
    case 211:
      ACCEPT_TOKEN(anon_sym_notdir);
      END_STATE();
    case 212:
      ACCEPT_TOKEN(anon_sym_origin);
      END_STATE();
    case 213:
      if (lookahead == 'd') ADVANCE(234);
      END_STATE();
    case 214:
      if (lookahead == 's') ADVANCE(235);
      END_STATE();
    case 215:
      if (lookahead == 'e') ADVANCE(236);
      END_STATE();
    case 216:
      if (lookahead == 't') ADVANCE(237);
      END_STATE();
    case 217:
      if (lookahead == 'd') ADVANCE(238);
      END_STATE();
    case 218:
      ACCEPT_TOKEN(anon_sym_suffix);
      END_STATE();
    case 219:
      if (lookahead == 'n') ADVANCE(239);
      END_STATE();
    case 220:
      if (lookahead == 'r') ADVANCE(240);
      END_STATE();
    case 221:
      if (lookahead == 'g') ADVANCE(241);
      END_STATE();
    case 222:
      if (lookahead == 'r') ADVANCE(242);
      END_STATE();
    case 223:
      if (lookahead == 's') ADVANCE(243);
      END_STATE();
    case 224:
      ACCEPT_TOKEN(anon_sym_abspath);
      END_STATE();
    case 225:
      if (lookahead == 'i') ADVANCE(244);
      END_STATE();
    case 226:
      if (lookahead == 'i') ADVANCE(245);
      END_STATE();
    case 227:
      if (lookahead == 'e') ADVANCE(246);
      END_STATE();
    case 228:
      if (lookahead == 'o') ADVANCE(247);
      END_STATE();
    case 229:
      if (lookahead == 'i') ADVANCE(248);
      END_STATE();
    case 230:
      if (lookahead == 'r') ADVANCE(249);
      END_STATE();
    case 231:
      ACCEPT_TOKEN(anon_sym_foreach);
      END_STATE();
    case 232:
      ACCEPT_TOKEN(anon_sym_include);
      END_STATE();
    case 233:
      if (lookahead == 'd') ADVANCE(250);
      END_STATE();
    case 234:
      if (lookahead == 'e') ADVANCE(251);
      END_STATE();
    case 235:
      if (lookahead == 't') ADVANCE(252);
      END_STATE();
    case 236:
      ACCEPT_TOKEN(anon_sym_private);
      END_STATE();
    case 237:
      if (lookahead == 'h') ADVANCE(253);
      END_STATE();
    case 238:
      if (lookahead == 'e') ADVANCE(254);
      END_STATE();
    case 239:
      if (lookahead == 'e') ADVANCE(255);
      END_STATE();
    case 240:
      if (lookahead == 't') ADVANCE(256);
      END_STATE();
    case 241:
      ACCEPT_TOKEN(anon_sym_warning);
      END_STATE();
    case 242:
      if (lookahead == 'd') ADVANCE(257);
      END_STATE();
    case 243:
      if (lookahead == 't') ADVANCE(258);
      END_STATE();
    case 244:
      if (lookahead == 'x') ADVANCE(259);
      END_STATE();
    case 245:
      if (lookahead == 'x') ADVANCE(260);
      END_STATE();
    case 246:
      ACCEPT_TOKEN(anon_sym_basename);
      END_STATE();
    case 247:
      if (lookahead == 'u') ADVANCE(261);
      END_STATE();
    case 248:
      if (lookahead == 'n') ADVANCE(262);
      END_STATE();
    case 249:
      if (lookahead == 'd') ADVANCE(263);
      END_STATE();
    case 250:
      ACCEPT_TOKEN(anon_sym_lastword);
      END_STATE();
    case 251:
      ACCEPT_TOKEN(anon_sym_override);
      END_STATE();
    case 252:
      ACCEPT_TOKEN(anon_sym_patsubst);
      END_STATE();
    case 253:
      ACCEPT_TOKEN(anon_sym_realpath);
      END_STATE();
    case 254:
      ACCEPT_TOKEN(anon_sym_sinclude);
      END_STATE();
    case 255:
      ACCEPT_TOKEN(anon_sym_undefine);
      END_STATE();
    case 256:
      ACCEPT_TOKEN(anon_sym_unexport);
      END_STATE();
    case 257:
      ACCEPT_TOKEN(anon_sym_wildcard);
      END_STATE();
    case 258:
      ACCEPT_TOKEN(anon_sym_wordlist);
      END_STATE();
    case 259:
      ACCEPT_TOKEN(anon_sym_addprefix);
      END_STATE();
    case 260:
      ACCEPT_TOKEN(anon_sym_addsuffix);
      END_STATE();
    case 261:
      if (lookahead == 't') ADVANCE(264);
      END_STATE();
    case 262:
      if (lookahead == 'g') ADVANCE(265);
      END_STATE();
    case 263:
      ACCEPT_TOKEN(anon_sym_firstword);
      END_STATE();
    case 264:
      ACCEPT_TOKEN(anon_sym_filter_DASHout);
      END_STATE();
    case 265:
      ACCEPT_TOKEN(anon_sym_findstring);
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
static uint16_t ts_small_parse_table[33874];

// Filled from tables.zlib by grammar.go, see fetchgrammar.
static uint32_t ts_small_parse_table_map[1278];

// Filled from tables.zlib by grammar.go, see fetchgrammar.
static TSParseActionEntry ts_parse_actions[2798];

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

TS_PUBLIC const TSLanguage *tree_sitter_make(void) {
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
    .keyword_capture_token = sym_word,
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
