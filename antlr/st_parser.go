// Code generated from ./antlr/ST.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ST

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type STParser struct {
	*antlr.BaseParser
}

var STParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func stParserInit() {
	staticData := &STParserStaticData
	staticData.LiteralNames = []string{
		"", "'#'", "", "", "", "", "", "", "'BOOL'", "", "", "'VAR'", "'VAR_INPUT'",
		"'VAR_OUTPUT'", "'VAR_IN_OUT'", "'VAR_EXTERNAL'", "'VAR_GLOBAL'", "'VAR_TEMP'",
		"'END_VAR'", "'FUNCTION'", "'END_FUNCTION'", "'FUNCTION_BLOCK'", "'END_FUNCTION_BLOCK'",
		"'PROGRAM'", "'END_PROGRAM'", "'IF'", "'THEN'", "'ELSE'", "'ELSIF'",
		"'END_IF'", "'FOR'", "'TO'", "'BY'", "'DO'", "'END_FOR'", "'WHILE'",
		"'END_WHILE'", "'REPEAT'", "'UNTIL'", "'END_REPEAT'", "'CASE'", "'OF'",
		"'END_CASE'", "'RETURN'", "'NOT'", "'AND'", "'OR'", "'XOR'", "'ARRAY'",
		"'STRUCT'", "'END_STRUCT'", "'TYPE'", "'END_TYPE'", "'CONSTANT'", "'RETAIN'",
		"'NON_RETAIN'", "'CONTINUE'", "'EXIT'", "':='", "'='", "'<>'", "'<'",
		"'<='", "'>'", "'>='", "'+'", "'-'", "'*'", "'/'", "'MOD'", "'**'",
		"'.'", "','", "':'", "';'", "'('", "')'", "'['", "']'", "'..'", "",
		"", "", "", "", "", "", "", "", "", "", "", "'R_EDGE'", "'F_EDGE'",
	}
	staticData.SymbolicNames = []string{
		"", "", "INT_TYPE_NAME", "REAL_TYPE_NAME", "TIME_TYPE_NAME", "DATE_TYPE_NAME",
		"TOD_TYPE_NAME", "DT_TYPE_NAME", "BOOL_TYPE_NAME", "BIT_TYPE_NAME",
		"STRING_TYPE_NAME", "VAR", "VAR_INPUT", "VAR_OUTPUT", "VAR_IN_OUT",
		"VAR_EXTERNAL", "VAR_GLOBAL", "VAR_TEMP", "END_VAR", "FUNCTION", "END_FUNCTION",
		"FUNCTION_BLOCK", "END_FUNCTION_BLOCK", "PROGRAM", "END_PROGRAM", "IF",
		"THEN", "ELSE", "ELSIF", "END_IF", "FOR", "TO", "BY", "DO", "END_FOR",
		"WHILE", "END_WHILE", "REPEAT", "UNTIL", "END_REPEAT", "CASE", "OF",
		"END_CASE", "RETURN", "NOT", "AND", "OR", "XOR", "ARRAY", "STRUCT",
		"END_STRUCT", "TYPE", "END_TYPE", "CONSTANT", "RETAIN", "NON_RETAIN",
		"CONTINUE", "EXIT", "ASSIGN", "EQUAL", "NOT_EQUAL", "LESS", "LESS_EQ",
		"GREATER", "GREATER_EQ", "PLUS", "MINUS", "MULT", "DIV", "MOD", "POWER",
		"DOT", "COMMA", "COLON", "SEMICOLON", "LPAREN", "RPAREN", "LBRACK",
		"RBRACK", "DOTDOT", "IDENTIFIER", "INT_LITERAL", "REAL_LITERAL", "BOOL_LITERAL",
		"TIME_LITERAL", "STRING_LITERAL", "DIRECT_VARIABLE", "COMMENT", "COMMENT2",
		"LINE_COMMENT", "WS", "BIT_STRING_LITERAL", "R_EDGE", "F_EDGE",
	}
	staticData.RuleNames = []string{
		"pous", "program", "function_decl", "function_block_decl", "type_declaration",
		"type_definition", "global_var_declaration", "var_declaration_block",
		"var_decl", "data_type", "elementary_type_name", "array_type", "subrange",
		"structured_type", "statement_list", "statement", "assignment_statement",
		"if_statement", "case_statement", "case_element", "case_label", "for_statement",
		"while_statement", "repeat_statement", "function_invocation", "return_statement",
		"continue_statement", "exit_statement", "expression", "literal", "typed_literal",
		"type_name", "variable", "input_decl", "output_decl", "edge_decl", "array_conform_decl",
		"array_conformand", "variable_list", "data_type_access", "derived_type_access",
		"single_elem_type_access", "simple_type_access", "subrange_type_access",
		"enum_type_access", "array_type_access", "struct_type_access", "string_type_access",
		"namespace_name", "simple_type_name", "subrange_type_name", "enum_type_name",
		"array_type_name", "struct_type_name",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 93, 635, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 114, 8, 0, 10, 0,
		12, 0, 117, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 5, 1, 124, 8, 1, 10, 1,
		12, 1, 127, 9, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 136,
		8, 2, 1, 2, 1, 2, 5, 2, 140, 8, 2, 10, 2, 12, 2, 143, 9, 2, 1, 2, 3, 2,
		146, 8, 2, 1, 2, 1, 2, 5, 2, 150, 8, 2, 10, 2, 12, 2, 153, 9, 2, 1, 2,
		3, 2, 156, 8, 2, 1, 2, 1, 2, 5, 2, 160, 8, 2, 10, 2, 12, 2, 163, 9, 2,
		1, 2, 3, 2, 166, 8, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3,
		175, 8, 3, 10, 3, 12, 3, 178, 9, 3, 1, 3, 3, 3, 181, 8, 3, 1, 3, 1, 3,
		5, 3, 185, 8, 3, 10, 3, 12, 3, 188, 9, 3, 1, 3, 3, 3, 191, 8, 3, 1, 3,
		1, 3, 5, 3, 195, 8, 3, 10, 3, 12, 3, 198, 9, 3, 1, 3, 3, 3, 201, 8, 3,
		1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 4, 4, 208, 8, 4, 11, 4, 12, 4, 209, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 5, 6, 221, 8, 6, 10, 6,
		12, 6, 224, 9, 6, 1, 6, 1, 6, 1, 7, 1, 7, 4, 7, 230, 8, 7, 11, 7, 12, 7,
		231, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 5, 8, 239, 8, 8, 10, 8, 12, 8, 242,
		9, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 248, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 9, 1, 9, 3, 9, 256, 8, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 5, 11, 265, 8, 11, 10, 11, 12, 11, 268, 9, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 4, 13, 280, 8, 13,
		11, 13, 12, 13, 281, 1, 13, 1, 13, 1, 14, 4, 14, 287, 8, 14, 11, 14, 12,
		14, 288, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15,
		299, 8, 15, 1, 15, 1, 15, 3, 15, 303, 8, 15, 1, 15, 1, 15, 3, 15, 307,
		8, 15, 1, 15, 1, 15, 3, 15, 311, 8, 15, 1, 15, 1, 15, 3, 15, 315, 8, 15,
		1, 15, 1, 15, 3, 15, 319, 8, 15, 1, 15, 1, 15, 3, 15, 323, 8, 15, 1, 15,
		1, 15, 3, 15, 327, 8, 15, 3, 15, 329, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 344,
		8, 17, 10, 17, 12, 17, 347, 9, 17, 1, 17, 1, 17, 3, 17, 351, 8, 17, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 4, 18, 359, 8, 18, 11, 18, 12, 18,
		360, 1, 18, 1, 18, 3, 18, 365, 8, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19,
		5, 19, 372, 8, 19, 10, 19, 12, 19, 375, 9, 19, 1, 19, 1, 19, 1, 19, 1,
		20, 1, 20, 3, 20, 382, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 3, 21, 392, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 415, 8, 24, 10, 24, 12, 24, 418,
		9, 24, 3, 24, 420, 8, 24, 1, 24, 1, 24, 1, 25, 1, 25, 3, 25, 426, 8, 25,
		1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 3, 28, 442, 8, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28,
		468, 8, 28, 10, 28, 12, 28, 471, 9, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 3, 29, 479, 8, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 5, 32, 494, 8, 32, 10,
		32, 12, 32, 497, 9, 32, 1, 32, 3, 32, 500, 8, 32, 1, 33, 1, 33, 1, 33,
		3, 33, 505, 8, 33, 1, 34, 1, 34, 3, 34, 509, 8, 34, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 5, 37, 525, 8, 37, 10, 37, 12, 37, 528, 9, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 38, 1, 38, 1, 38, 5, 38, 537, 8, 38, 10, 38, 12, 38, 540, 9,
		38, 1, 39, 1, 39, 3, 39, 544, 8, 39, 1, 40, 1, 40, 1, 40, 1, 40, 3, 40,
		550, 8, 40, 1, 41, 1, 41, 1, 41, 3, 41, 555, 8, 41, 1, 42, 1, 42, 1, 42,
		5, 42, 560, 8, 42, 10, 42, 12, 42, 563, 9, 42, 1, 42, 1, 42, 1, 43, 1,
		43, 1, 43, 5, 43, 570, 8, 43, 10, 43, 12, 43, 573, 9, 43, 1, 43, 1, 43,
		1, 44, 1, 44, 1, 44, 5, 44, 580, 8, 44, 10, 44, 12, 44, 583, 9, 44, 1,
		44, 1, 44, 1, 45, 1, 45, 1, 45, 5, 45, 590, 8, 45, 10, 45, 12, 45, 593,
		9, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 5, 46, 600, 8, 46, 10, 46, 12,
		46, 603, 9, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 5, 47, 610, 8, 47, 10,
		47, 12, 47, 613, 9, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 5, 48, 620,
		8, 48, 10, 48, 12, 48, 623, 9, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1,
		51, 1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 0, 1, 56, 54, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
		48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82,
		84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 0, 10, 1, 0, 11, 17,
		2, 0, 2, 5, 8, 10, 2, 0, 44, 44, 65, 66, 1, 0, 67, 69, 1, 0, 65, 66, 1,
		0, 61, 64, 1, 0, 59, 60, 2, 0, 81, 82, 91, 91, 2, 0, 2, 3, 8, 9, 1, 0,
		92, 93, 675, 0, 115, 1, 0, 0, 0, 2, 120, 1, 0, 0, 0, 4, 131, 1, 0, 0, 0,
		6, 170, 1, 0, 0, 0, 8, 205, 1, 0, 0, 0, 10, 213, 1, 0, 0, 0, 12, 218, 1,
		0, 0, 0, 14, 227, 1, 0, 0, 0, 16, 235, 1, 0, 0, 0, 18, 255, 1, 0, 0, 0,
		20, 257, 1, 0, 0, 0, 22, 259, 1, 0, 0, 0, 24, 273, 1, 0, 0, 0, 26, 277,
		1, 0, 0, 0, 28, 286, 1, 0, 0, 0, 30, 328, 1, 0, 0, 0, 32, 330, 1, 0, 0,
		0, 34, 334, 1, 0, 0, 0, 36, 354, 1, 0, 0, 0, 38, 368, 1, 0, 0, 0, 40, 381,
		1, 0, 0, 0, 42, 383, 1, 0, 0, 0, 44, 397, 1, 0, 0, 0, 46, 403, 1, 0, 0,
		0, 48, 409, 1, 0, 0, 0, 50, 423, 1, 0, 0, 0, 52, 427, 1, 0, 0, 0, 54, 429,
		1, 0, 0, 0, 56, 441, 1, 0, 0, 0, 58, 478, 1, 0, 0, 0, 60, 480, 1, 0, 0,
		0, 62, 484, 1, 0, 0, 0, 64, 499, 1, 0, 0, 0, 66, 504, 1, 0, 0, 0, 68, 508,
		1, 0, 0, 0, 70, 510, 1, 0, 0, 0, 72, 515, 1, 0, 0, 0, 74, 519, 1, 0, 0,
		0, 76, 533, 1, 0, 0, 0, 78, 543, 1, 0, 0, 0, 80, 549, 1, 0, 0, 0, 82, 554,
		1, 0, 0, 0, 84, 561, 1, 0, 0, 0, 86, 571, 1, 0, 0, 0, 88, 581, 1, 0, 0,
		0, 90, 591, 1, 0, 0, 0, 92, 601, 1, 0, 0, 0, 94, 611, 1, 0, 0, 0, 96, 616,
		1, 0, 0, 0, 98, 624, 1, 0, 0, 0, 100, 626, 1, 0, 0, 0, 102, 628, 1, 0,
		0, 0, 104, 630, 1, 0, 0, 0, 106, 632, 1, 0, 0, 0, 108, 114, 3, 2, 1, 0,
		109, 114, 3, 4, 2, 0, 110, 114, 3, 6, 3, 0, 111, 114, 3, 8, 4, 0, 112,
		114, 3, 12, 6, 0, 113, 108, 1, 0, 0, 0, 113, 109, 1, 0, 0, 0, 113, 110,
		1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 112, 1, 0, 0, 0, 114, 117, 1, 0,
		0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 118, 1, 0, 0, 0,
		117, 115, 1, 0, 0, 0, 118, 119, 5, 0, 0, 1, 119, 1, 1, 0, 0, 0, 120, 121,
		5, 23, 0, 0, 121, 125, 5, 80, 0, 0, 122, 124, 3, 14, 7, 0, 123, 122, 1,
		0, 0, 0, 124, 127, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 125, 126, 1, 0, 0,
		0, 126, 128, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 128, 129, 3, 28, 14, 0,
		129, 130, 5, 24, 0, 0, 130, 3, 1, 0, 0, 0, 131, 132, 5, 19, 0, 0, 132,
		135, 5, 80, 0, 0, 133, 134, 5, 73, 0, 0, 134, 136, 3, 18, 9, 0, 135, 133,
		1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 145, 1, 0, 0, 0, 137, 141, 5, 12,
		0, 0, 138, 140, 3, 66, 33, 0, 139, 138, 1, 0, 0, 0, 140, 143, 1, 0, 0,
		0, 141, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 144, 1, 0, 0, 0, 143,
		141, 1, 0, 0, 0, 144, 146, 5, 18, 0, 0, 145, 137, 1, 0, 0, 0, 145, 146,
		1, 0, 0, 0, 146, 155, 1, 0, 0, 0, 147, 151, 5, 13, 0, 0, 148, 150, 3, 68,
		34, 0, 149, 148, 1, 0, 0, 0, 150, 153, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0,
		151, 152, 1, 0, 0, 0, 152, 154, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 154,
		156, 5, 18, 0, 0, 155, 147, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 165,
		1, 0, 0, 0, 157, 161, 5, 11, 0, 0, 158, 160, 3, 16, 8, 0, 159, 158, 1,
		0, 0, 0, 160, 163, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 161, 162, 1, 0, 0,
		0, 162, 164, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 164, 166, 5, 18, 0, 0, 165,
		157, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 168,
		3, 28, 14, 0, 168, 169, 5, 20, 0, 0, 169, 5, 1, 0, 0, 0, 170, 171, 5, 21,
		0, 0, 171, 180, 5, 80, 0, 0, 172, 176, 5, 12, 0, 0, 173, 175, 3, 66, 33,
		0, 174, 173, 1, 0, 0, 0, 175, 178, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 176,
		177, 1, 0, 0, 0, 177, 179, 1, 0, 0, 0, 178, 176, 1, 0, 0, 0, 179, 181,
		5, 18, 0, 0, 180, 172, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 190, 1, 0,
		0, 0, 182, 186, 5, 13, 0, 0, 183, 185, 3, 68, 34, 0, 184, 183, 1, 0, 0,
		0, 185, 188, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187,
		189, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 189, 191, 5, 18, 0, 0, 190, 182,
		1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 200, 1, 0, 0, 0, 192, 196, 5, 11,
		0, 0, 193, 195, 3, 16, 8, 0, 194, 193, 1, 0, 0, 0, 195, 198, 1, 0, 0, 0,
		196, 194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 199, 1, 0, 0, 0, 198,
		196, 1, 0, 0, 0, 199, 201, 5, 18, 0, 0, 200, 192, 1, 0, 0, 0, 200, 201,
		1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 203, 3, 28, 14, 0, 203, 204, 5,
		22, 0, 0, 204, 7, 1, 0, 0, 0, 205, 207, 5, 51, 0, 0, 206, 208, 3, 10, 5,
		0, 207, 206, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 209,
		210, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 212, 5, 52, 0, 0, 212, 9, 1,
		0, 0, 0, 213, 214, 5, 80, 0, 0, 214, 215, 5, 73, 0, 0, 215, 216, 3, 18,
		9, 0, 216, 217, 5, 74, 0, 0, 217, 11, 1, 0, 0, 0, 218, 222, 5, 16, 0, 0,
		219, 221, 3, 16, 8, 0, 220, 219, 1, 0, 0, 0, 221, 224, 1, 0, 0, 0, 222,
		220, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 225, 1, 0, 0, 0, 224, 222,
		1, 0, 0, 0, 225, 226, 5, 18, 0, 0, 226, 13, 1, 0, 0, 0, 227, 229, 7, 0,
		0, 0, 228, 230, 3, 16, 8, 0, 229, 228, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0,
		231, 229, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233,
		234, 5, 18, 0, 0, 234, 15, 1, 0, 0, 0, 235, 240, 5, 80, 0, 0, 236, 237,
		5, 72, 0, 0, 237, 239, 5, 80, 0, 0, 238, 236, 1, 0, 0, 0, 239, 242, 1,
		0, 0, 0, 240, 238, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0, 241, 243, 1, 0, 0,
		0, 242, 240, 1, 0, 0, 0, 243, 244, 5, 73, 0, 0, 244, 247, 3, 18, 9, 0,
		245, 246, 5, 58, 0, 0, 246, 248, 3, 56, 28, 0, 247, 245, 1, 0, 0, 0, 247,
		248, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 250, 5, 74, 0, 0, 250, 17,
		1, 0, 0, 0, 251, 256, 3, 20, 10, 0, 252, 256, 3, 22, 11, 0, 253, 256, 3,
		26, 13, 0, 254, 256, 5, 80, 0, 0, 255, 251, 1, 0, 0, 0, 255, 252, 1, 0,
		0, 0, 255, 253, 1, 0, 0, 0, 255, 254, 1, 0, 0, 0, 256, 19, 1, 0, 0, 0,
		257, 258, 7, 1, 0, 0, 258, 21, 1, 0, 0, 0, 259, 260, 5, 48, 0, 0, 260,
		261, 5, 77, 0, 0, 261, 266, 3, 24, 12, 0, 262, 263, 5, 72, 0, 0, 263, 265,
		3, 24, 12, 0, 264, 262, 1, 0, 0, 0, 265, 268, 1, 0, 0, 0, 266, 264, 1,
		0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 269, 1, 0, 0, 0, 268, 266, 1, 0, 0,
		0, 269, 270, 5, 78, 0, 0, 270, 271, 5, 41, 0, 0, 271, 272, 3, 18, 9, 0,
		272, 23, 1, 0, 0, 0, 273, 274, 3, 56, 28, 0, 274, 275, 5, 79, 0, 0, 275,
		276, 3, 56, 28, 0, 276, 25, 1, 0, 0, 0, 277, 279, 5, 49, 0, 0, 278, 280,
		3, 16, 8, 0, 279, 278, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 279, 1, 0,
		0, 0, 281, 282, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 284, 5, 50, 0, 0,
		284, 27, 1, 0, 0, 0, 285, 287, 3, 30, 15, 0, 286, 285, 1, 0, 0, 0, 287,
		288, 1, 0, 0, 0, 288, 286, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 29, 1,
		0, 0, 0, 290, 291, 3, 32, 16, 0, 291, 292, 5, 74, 0, 0, 292, 329, 1, 0,
		0, 0, 293, 294, 3, 48, 24, 0, 294, 295, 5, 74, 0, 0, 295, 329, 1, 0, 0,
		0, 296, 298, 3, 52, 26, 0, 297, 299, 5, 74, 0, 0, 298, 297, 1, 0, 0, 0,
		298, 299, 1, 0, 0, 0, 299, 329, 1, 0, 0, 0, 300, 302, 3, 54, 27, 0, 301,
		303, 5, 74, 0, 0, 302, 301, 1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303, 329,
		1, 0, 0, 0, 304, 306, 3, 50, 25, 0, 305, 307, 5, 74, 0, 0, 306, 305, 1,
		0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 329, 1, 0, 0, 0, 308, 310, 3, 34, 17,
		0, 309, 311, 5, 74, 0, 0, 310, 309, 1, 0, 0, 0, 310, 311, 1, 0, 0, 0, 311,
		329, 1, 0, 0, 0, 312, 314, 3, 36, 18, 0, 313, 315, 5, 74, 0, 0, 314, 313,
		1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 329, 1, 0, 0, 0, 316, 318, 3, 42,
		21, 0, 317, 319, 5, 74, 0, 0, 318, 317, 1, 0, 0, 0, 318, 319, 1, 0, 0,
		0, 319, 329, 1, 0, 0, 0, 320, 322, 3, 44, 22, 0, 321, 323, 5, 74, 0, 0,
		322, 321, 1, 0, 0, 0, 322, 323, 1, 0, 0, 0, 323, 329, 1, 0, 0, 0, 324,
		326, 3, 46, 23, 0, 325, 327, 5, 74, 0, 0, 326, 325, 1, 0, 0, 0, 326, 327,
		1, 0, 0, 0, 327, 329, 1, 0, 0, 0, 328, 290, 1, 0, 0, 0, 328, 293, 1, 0,
		0, 0, 328, 296, 1, 0, 0, 0, 328, 300, 1, 0, 0, 0, 328, 304, 1, 0, 0, 0,
		328, 308, 1, 0, 0, 0, 328, 312, 1, 0, 0, 0, 328, 316, 1, 0, 0, 0, 328,
		320, 1, 0, 0, 0, 328, 324, 1, 0, 0, 0, 329, 31, 1, 0, 0, 0, 330, 331, 3,
		64, 32, 0, 331, 332, 5, 58, 0, 0, 332, 333, 3, 56, 28, 0, 333, 33, 1, 0,
		0, 0, 334, 335, 5, 25, 0, 0, 335, 336, 3, 56, 28, 0, 336, 337, 5, 26, 0,
		0, 337, 345, 3, 28, 14, 0, 338, 339, 5, 28, 0, 0, 339, 340, 3, 56, 28,
		0, 340, 341, 5, 26, 0, 0, 341, 342, 3, 28, 14, 0, 342, 344, 1, 0, 0, 0,
		343, 338, 1, 0, 0, 0, 344, 347, 1, 0, 0, 0, 345, 343, 1, 0, 0, 0, 345,
		346, 1, 0, 0, 0, 346, 350, 1, 0, 0, 0, 347, 345, 1, 0, 0, 0, 348, 349,
		5, 27, 0, 0, 349, 351, 3, 28, 14, 0, 350, 348, 1, 0, 0, 0, 350, 351, 1,
		0, 0, 0, 351, 352, 1, 0, 0, 0, 352, 353, 5, 29, 0, 0, 353, 35, 1, 0, 0,
		0, 354, 355, 5, 40, 0, 0, 355, 356, 3, 56, 28, 0, 356, 358, 5, 41, 0, 0,
		357, 359, 3, 38, 19, 0, 358, 357, 1, 0, 0, 0, 359, 360, 1, 0, 0, 0, 360,
		358, 1, 0, 0, 0, 360, 361, 1, 0, 0, 0, 361, 364, 1, 0, 0, 0, 362, 363,
		5, 27, 0, 0, 363, 365, 3, 28, 14, 0, 364, 362, 1, 0, 0, 0, 364, 365, 1,
		0, 0, 0, 365, 366, 1, 0, 0, 0, 366, 367, 5, 42, 0, 0, 367, 37, 1, 0, 0,
		0, 368, 373, 3, 40, 20, 0, 369, 370, 5, 72, 0, 0, 370, 372, 3, 40, 20,
		0, 371, 369, 1, 0, 0, 0, 372, 375, 1, 0, 0, 0, 373, 371, 1, 0, 0, 0, 373,
		374, 1, 0, 0, 0, 374, 376, 1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 376, 377,
		5, 73, 0, 0, 377, 378, 3, 28, 14, 0, 378, 39, 1, 0, 0, 0, 379, 382, 3,
		56, 28, 0, 380, 382, 3, 24, 12, 0, 381, 379, 1, 0, 0, 0, 381, 380, 1, 0,
		0, 0, 382, 41, 1, 0, 0, 0, 383, 384, 5, 30, 0, 0, 384, 385, 5, 80, 0, 0,
		385, 386, 5, 58, 0, 0, 386, 387, 3, 56, 28, 0, 387, 388, 5, 31, 0, 0, 388,
		391, 3, 56, 28, 0, 389, 390, 5, 32, 0, 0, 390, 392, 3, 56, 28, 0, 391,
		389, 1, 0, 0, 0, 391, 392, 1, 0, 0, 0, 392, 393, 1, 0, 0, 0, 393, 394,
		5, 33, 0, 0, 394, 395, 3, 28, 14, 0, 395, 396, 5, 34, 0, 0, 396, 43, 1,
		0, 0, 0, 397, 398, 5, 35, 0, 0, 398, 399, 3, 56, 28, 0, 399, 400, 5, 33,
		0, 0, 400, 401, 3, 28, 14, 0, 401, 402, 5, 36, 0, 0, 402, 45, 1, 0, 0,
		0, 403, 404, 5, 37, 0, 0, 404, 405, 3, 28, 14, 0, 405, 406, 5, 38, 0, 0,
		406, 407, 3, 56, 28, 0, 407, 408, 5, 39, 0, 0, 408, 47, 1, 0, 0, 0, 409,
		410, 5, 80, 0, 0, 410, 419, 5, 75, 0, 0, 411, 416, 3, 56, 28, 0, 412, 413,
		5, 72, 0, 0, 413, 415, 3, 56, 28, 0, 414, 412, 1, 0, 0, 0, 415, 418, 1,
		0, 0, 0, 416, 414, 1, 0, 0, 0, 416, 417, 1, 0, 0, 0, 417, 420, 1, 0, 0,
		0, 418, 416, 1, 0, 0, 0, 419, 411, 1, 0, 0, 0, 419, 420, 1, 0, 0, 0, 420,
		421, 1, 0, 0, 0, 421, 422, 5, 76, 0, 0, 422, 49, 1, 0, 0, 0, 423, 425,
		5, 43, 0, 0, 424, 426, 3, 56, 28, 0, 425, 424, 1, 0, 0, 0, 425, 426, 1,
		0, 0, 0, 426, 51, 1, 0, 0, 0, 427, 428, 5, 56, 0, 0, 428, 53, 1, 0, 0,
		0, 429, 430, 5, 57, 0, 0, 430, 55, 1, 0, 0, 0, 431, 432, 6, 28, -1, 0,
		432, 433, 5, 75, 0, 0, 433, 434, 3, 56, 28, 0, 434, 435, 5, 76, 0, 0, 435,
		442, 1, 0, 0, 0, 436, 437, 7, 2, 0, 0, 437, 442, 3, 56, 28, 12, 438, 442,
		3, 58, 29, 0, 439, 442, 3, 64, 32, 0, 440, 442, 3, 48, 24, 0, 441, 431,
		1, 0, 0, 0, 441, 436, 1, 0, 0, 0, 441, 438, 1, 0, 0, 0, 441, 439, 1, 0,
		0, 0, 441, 440, 1, 0, 0, 0, 442, 469, 1, 0, 0, 0, 443, 444, 10, 11, 0,
		0, 444, 445, 5, 70, 0, 0, 445, 468, 3, 56, 28, 12, 446, 447, 10, 10, 0,
		0, 447, 448, 7, 3, 0, 0, 448, 468, 3, 56, 28, 11, 449, 450, 10, 9, 0, 0,
		450, 451, 7, 4, 0, 0, 451, 468, 3, 56, 28, 10, 452, 453, 10, 8, 0, 0, 453,
		454, 7, 5, 0, 0, 454, 468, 3, 56, 28, 9, 455, 456, 10, 7, 0, 0, 456, 457,
		7, 6, 0, 0, 457, 468, 3, 56, 28, 8, 458, 459, 10, 6, 0, 0, 459, 460, 5,
		45, 0, 0, 460, 468, 3, 56, 28, 7, 461, 462, 10, 5, 0, 0, 462, 463, 5, 47,
		0, 0, 463, 468, 3, 56, 28, 6, 464, 465, 10, 4, 0, 0, 465, 466, 5, 46, 0,
		0, 466, 468, 3, 56, 28, 5, 467, 443, 1, 0, 0, 0, 467, 446, 1, 0, 0, 0,
		467, 449, 1, 0, 0, 0, 467, 452, 1, 0, 0, 0, 467, 455, 1, 0, 0, 0, 467,
		458, 1, 0, 0, 0, 467, 461, 1, 0, 0, 0, 467, 464, 1, 0, 0, 0, 468, 471,
		1, 0, 0, 0, 469, 467, 1, 0, 0, 0, 469, 470, 1, 0, 0, 0, 470, 57, 1, 0,
		0, 0, 471, 469, 1, 0, 0, 0, 472, 479, 5, 81, 0, 0, 473, 479, 5, 82, 0,
		0, 474, 479, 5, 83, 0, 0, 475, 479, 5, 84, 0, 0, 476, 479, 5, 85, 0, 0,
		477, 479, 3, 60, 30, 0, 478, 472, 1, 0, 0, 0, 478, 473, 1, 0, 0, 0, 478,
		474, 1, 0, 0, 0, 478, 475, 1, 0, 0, 0, 478, 476, 1, 0, 0, 0, 478, 477,
		1, 0, 0, 0, 479, 59, 1, 0, 0, 0, 480, 481, 3, 62, 31, 0, 481, 482, 5, 1,
		0, 0, 482, 483, 7, 7, 0, 0, 483, 61, 1, 0, 0, 0, 484, 485, 7, 8, 0, 0,
		485, 63, 1, 0, 0, 0, 486, 495, 5, 80, 0, 0, 487, 488, 5, 71, 0, 0, 488,
		494, 5, 80, 0, 0, 489, 490, 5, 77, 0, 0, 490, 491, 3, 56, 28, 0, 491, 492,
		5, 78, 0, 0, 492, 494, 1, 0, 0, 0, 493, 487, 1, 0, 0, 0, 493, 489, 1, 0,
		0, 0, 494, 497, 1, 0, 0, 0, 495, 493, 1, 0, 0, 0, 495, 496, 1, 0, 0, 0,
		496, 500, 1, 0, 0, 0, 497, 495, 1, 0, 0, 0, 498, 500, 5, 86, 0, 0, 499,
		486, 1, 0, 0, 0, 499, 498, 1, 0, 0, 0, 500, 65, 1, 0, 0, 0, 501, 505, 3,
		16, 8, 0, 502, 505, 3, 70, 35, 0, 503, 505, 3, 72, 36, 0, 504, 501, 1,
		0, 0, 0, 504, 502, 1, 0, 0, 0, 504, 503, 1, 0, 0, 0, 505, 67, 1, 0, 0,
		0, 506, 509, 3, 16, 8, 0, 507, 509, 3, 72, 36, 0, 508, 506, 1, 0, 0, 0,
		508, 507, 1, 0, 0, 0, 509, 69, 1, 0, 0, 0, 510, 511, 3, 76, 38, 0, 511,
		512, 5, 73, 0, 0, 512, 513, 5, 8, 0, 0, 513, 514, 7, 9, 0, 0, 514, 71,
		1, 0, 0, 0, 515, 516, 3, 76, 38, 0, 516, 517, 5, 73, 0, 0, 517, 518, 3,
		74, 37, 0, 518, 73, 1, 0, 0, 0, 519, 520, 5, 48, 0, 0, 520, 521, 5, 77,
		0, 0, 521, 526, 5, 67, 0, 0, 522, 523, 5, 72, 0, 0, 523, 525, 5, 67, 0,
		0, 524, 522, 1, 0, 0, 0, 525, 528, 1, 0, 0, 0, 526, 524, 1, 0, 0, 0, 526,
		527, 1, 0, 0, 0, 527, 529, 1, 0, 0, 0, 528, 526, 1, 0, 0, 0, 529, 530,
		5, 78, 0, 0, 530, 531, 5, 41, 0, 0, 531, 532, 3, 78, 39, 0, 532, 75, 1,
		0, 0, 0, 533, 538, 5, 80, 0, 0, 534, 535, 5, 72, 0, 0, 535, 537, 5, 80,
		0, 0, 536, 534, 1, 0, 0, 0, 537, 540, 1, 0, 0, 0, 538, 536, 1, 0, 0, 0,
		538, 539, 1, 0, 0, 0, 539, 77, 1, 0, 0, 0, 540, 538, 1, 0, 0, 0, 541, 544,
		3, 20, 10, 0, 542, 544, 3, 80, 40, 0, 543, 541, 1, 0, 0, 0, 543, 542, 1,
		0, 0, 0, 544, 79, 1, 0, 0, 0, 545, 550, 3, 82, 41, 0, 546, 550, 3, 90,
		45, 0, 547, 550, 3, 92, 46, 0, 548, 550, 3, 94, 47, 0, 549, 545, 1, 0,
		0, 0, 549, 546, 1, 0, 0, 0, 549, 547, 1, 0, 0, 0, 549, 548, 1, 0, 0, 0,
		550, 81, 1, 0, 0, 0, 551, 555, 3, 84, 42, 0, 552, 555, 3, 86, 43, 0, 553,
		555, 3, 88, 44, 0, 554, 551, 1, 0, 0, 0, 554, 552, 1, 0, 0, 0, 554, 553,
		1, 0, 0, 0, 555, 83, 1, 0, 0, 0, 556, 557, 3, 96, 48, 0, 557, 558, 5, 71,
		0, 0, 558, 560, 1, 0, 0, 0, 559, 556, 1, 0, 0, 0, 560, 563, 1, 0, 0, 0,
		561, 559, 1, 0, 0, 0, 561, 562, 1, 0, 0, 0, 562, 564, 1, 0, 0, 0, 563,
		561, 1, 0, 0, 0, 564, 565, 3, 98, 49, 0, 565, 85, 1, 0, 0, 0, 566, 567,
		3, 96, 48, 0, 567, 568, 5, 71, 0, 0, 568, 570, 1, 0, 0, 0, 569, 566, 1,
		0, 0, 0, 570, 573, 1, 0, 0, 0, 571, 569, 1, 0, 0, 0, 571, 572, 1, 0, 0,
		0, 572, 574, 1, 0, 0, 0, 573, 571, 1, 0, 0, 0, 574, 575, 3, 100, 50, 0,
		575, 87, 1, 0, 0, 0, 576, 577, 3, 96, 48, 0, 577, 578, 5, 71, 0, 0, 578,
		580, 1, 0, 0, 0, 579, 576, 1, 0, 0, 0, 580, 583, 1, 0, 0, 0, 581, 579,
		1, 0, 0, 0, 581, 582, 1, 0, 0, 0, 582, 584, 1, 0, 0, 0, 583, 581, 1, 0,
		0, 0, 584, 585, 3, 102, 51, 0, 585, 89, 1, 0, 0, 0, 586, 587, 3, 96, 48,
		0, 587, 588, 5, 71, 0, 0, 588, 590, 1, 0, 0, 0, 589, 586, 1, 0, 0, 0, 590,
		593, 1, 0, 0, 0, 591, 589, 1, 0, 0, 0, 591, 592, 1, 0, 0, 0, 592, 594,
		1, 0, 0, 0, 593, 591, 1, 0, 0, 0, 594, 595, 3, 104, 52, 0, 595, 91, 1,
		0, 0, 0, 596, 597, 3, 96, 48, 0, 597, 598, 5, 71, 0, 0, 598, 600, 1, 0,
		0, 0, 599, 596, 1, 0, 0, 0, 600, 603, 1, 0, 0, 0, 601, 599, 1, 0, 0, 0,
		601, 602, 1, 0, 0, 0, 602, 604, 1, 0, 0, 0, 603, 601, 1, 0, 0, 0, 604,
		605, 3, 106, 53, 0, 605, 93, 1, 0, 0, 0, 606, 607, 3, 96, 48, 0, 607, 608,
		5, 71, 0, 0, 608, 610, 1, 0, 0, 0, 609, 606, 1, 0, 0, 0, 610, 613, 1, 0,
		0, 0, 611, 609, 1, 0, 0, 0, 611, 612, 1, 0, 0, 0, 612, 614, 1, 0, 0, 0,
		613, 611, 1, 0, 0, 0, 614, 615, 5, 10, 0, 0, 615, 95, 1, 0, 0, 0, 616,
		621, 5, 80, 0, 0, 617, 618, 5, 71, 0, 0, 618, 620, 5, 80, 0, 0, 619, 617,
		1, 0, 0, 0, 620, 623, 1, 0, 0, 0, 621, 619, 1, 0, 0, 0, 621, 622, 1, 0,
		0, 0, 622, 97, 1, 0, 0, 0, 623, 621, 1, 0, 0, 0, 624, 625, 5, 80, 0, 0,
		625, 99, 1, 0, 0, 0, 626, 627, 5, 80, 0, 0, 627, 101, 1, 0, 0, 0, 628,
		629, 5, 80, 0, 0, 629, 103, 1, 0, 0, 0, 630, 631, 5, 80, 0, 0, 631, 105,
		1, 0, 0, 0, 632, 633, 5, 80, 0, 0, 633, 107, 1, 0, 0, 0, 65, 113, 115,
		125, 135, 141, 145, 151, 155, 161, 165, 176, 180, 186, 190, 196, 200, 209,
		222, 231, 240, 247, 255, 266, 281, 288, 298, 302, 306, 310, 314, 318, 322,
		326, 328, 345, 350, 360, 364, 373, 381, 391, 416, 419, 425, 441, 467, 469,
		478, 493, 495, 499, 504, 508, 526, 538, 543, 549, 554, 561, 571, 581, 591,
		601, 611, 621,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// STParserInit initializes any static state used to implement STParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSTParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func STParserInit() {
	staticData := &STParserStaticData
	staticData.once.Do(stParserInit)
}

// NewSTParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSTParser(input antlr.TokenStream) *STParser {
	STParserInit()
	this := new(STParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &STParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ST.g4"

	return this
}

// STParser tokens.
const (
	STParserEOF                = antlr.TokenEOF
	STParserT__0               = 1
	STParserINT_TYPE_NAME      = 2
	STParserREAL_TYPE_NAME     = 3
	STParserTIME_TYPE_NAME     = 4
	STParserDATE_TYPE_NAME     = 5
	STParserTOD_TYPE_NAME      = 6
	STParserDT_TYPE_NAME       = 7
	STParserBOOL_TYPE_NAME     = 8
	STParserBIT_TYPE_NAME      = 9
	STParserSTRING_TYPE_NAME   = 10
	STParserVAR                = 11
	STParserVAR_INPUT          = 12
	STParserVAR_OUTPUT         = 13
	STParserVAR_IN_OUT         = 14
	STParserVAR_EXTERNAL       = 15
	STParserVAR_GLOBAL         = 16
	STParserVAR_TEMP           = 17
	STParserEND_VAR            = 18
	STParserFUNCTION           = 19
	STParserEND_FUNCTION       = 20
	STParserFUNCTION_BLOCK     = 21
	STParserEND_FUNCTION_BLOCK = 22
	STParserPROGRAM            = 23
	STParserEND_PROGRAM        = 24
	STParserIF                 = 25
	STParserTHEN               = 26
	STParserELSE               = 27
	STParserELSIF              = 28
	STParserEND_IF             = 29
	STParserFOR                = 30
	STParserTO                 = 31
	STParserBY                 = 32
	STParserDO                 = 33
	STParserEND_FOR            = 34
	STParserWHILE              = 35
	STParserEND_WHILE          = 36
	STParserREPEAT             = 37
	STParserUNTIL              = 38
	STParserEND_REPEAT         = 39
	STParserCASE               = 40
	STParserOF                 = 41
	STParserEND_CASE           = 42
	STParserRETURN             = 43
	STParserNOT                = 44
	STParserAND                = 45
	STParserOR                 = 46
	STParserXOR                = 47
	STParserARRAY              = 48
	STParserSTRUCT             = 49
	STParserEND_STRUCT         = 50
	STParserTYPE               = 51
	STParserEND_TYPE           = 52
	STParserCONSTANT           = 53
	STParserRETAIN             = 54
	STParserNON_RETAIN         = 55
	STParserCONTINUE           = 56
	STParserEXIT               = 57
	STParserASSIGN             = 58
	STParserEQUAL              = 59
	STParserNOT_EQUAL          = 60
	STParserLESS               = 61
	STParserLESS_EQ            = 62
	STParserGREATER            = 63
	STParserGREATER_EQ         = 64
	STParserPLUS               = 65
	STParserMINUS              = 66
	STParserMULT               = 67
	STParserDIV                = 68
	STParserMOD                = 69
	STParserPOWER              = 70
	STParserDOT                = 71
	STParserCOMMA              = 72
	STParserCOLON              = 73
	STParserSEMICOLON          = 74
	STParserLPAREN             = 75
	STParserRPAREN             = 76
	STParserLBRACK             = 77
	STParserRBRACK             = 78
	STParserDOTDOT             = 79
	STParserIDENTIFIER         = 80
	STParserINT_LITERAL        = 81
	STParserREAL_LITERAL       = 82
	STParserBOOL_LITERAL       = 83
	STParserTIME_LITERAL       = 84
	STParserSTRING_LITERAL     = 85
	STParserDIRECT_VARIABLE    = 86
	STParserCOMMENT            = 87
	STParserCOMMENT2           = 88
	STParserLINE_COMMENT       = 89
	STParserWS                 = 90
	STParserBIT_STRING_LITERAL = 91
	STParserR_EDGE             = 92
	STParserF_EDGE             = 93
)

// STParser rules.
const (
	STParserRULE_pous                    = 0
	STParserRULE_program                 = 1
	STParserRULE_function_decl           = 2
	STParserRULE_function_block_decl     = 3
	STParserRULE_type_declaration        = 4
	STParserRULE_type_definition         = 5
	STParserRULE_global_var_declaration  = 6
	STParserRULE_var_declaration_block   = 7
	STParserRULE_var_decl                = 8
	STParserRULE_data_type               = 9
	STParserRULE_elementary_type_name    = 10
	STParserRULE_array_type              = 11
	STParserRULE_subrange                = 12
	STParserRULE_structured_type         = 13
	STParserRULE_statement_list          = 14
	STParserRULE_statement               = 15
	STParserRULE_assignment_statement    = 16
	STParserRULE_if_statement            = 17
	STParserRULE_case_statement          = 18
	STParserRULE_case_element            = 19
	STParserRULE_case_label              = 20
	STParserRULE_for_statement           = 21
	STParserRULE_while_statement         = 22
	STParserRULE_repeat_statement        = 23
	STParserRULE_function_invocation     = 24
	STParserRULE_return_statement        = 25
	STParserRULE_continue_statement      = 26
	STParserRULE_exit_statement          = 27
	STParserRULE_expression              = 28
	STParserRULE_literal                 = 29
	STParserRULE_typed_literal           = 30
	STParserRULE_type_name               = 31
	STParserRULE_variable                = 32
	STParserRULE_input_decl              = 33
	STParserRULE_output_decl             = 34
	STParserRULE_edge_decl               = 35
	STParserRULE_array_conform_decl      = 36
	STParserRULE_array_conformand        = 37
	STParserRULE_variable_list           = 38
	STParserRULE_data_type_access        = 39
	STParserRULE_derived_type_access     = 40
	STParserRULE_single_elem_type_access = 41
	STParserRULE_simple_type_access      = 42
	STParserRULE_subrange_type_access    = 43
	STParserRULE_enum_type_access        = 44
	STParserRULE_array_type_access       = 45
	STParserRULE_struct_type_access      = 46
	STParserRULE_string_type_access      = 47
	STParserRULE_namespace_name          = 48
	STParserRULE_simple_type_name        = 49
	STParserRULE_subrange_type_name      = 50
	STParserRULE_enum_type_name          = 51
	STParserRULE_array_type_name         = 52
	STParserRULE_struct_type_name        = 53
)

// IPousContext is an interface to support dynamic dispatch.
type IPousContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllProgram() []IProgramContext
	Program(i int) IProgramContext
	AllFunction_decl() []IFunction_declContext
	Function_decl(i int) IFunction_declContext
	AllFunction_block_decl() []IFunction_block_declContext
	Function_block_decl(i int) IFunction_block_declContext
	AllType_declaration() []IType_declarationContext
	Type_declaration(i int) IType_declarationContext
	AllGlobal_var_declaration() []IGlobal_var_declarationContext
	Global_var_declaration(i int) IGlobal_var_declarationContext

	// IsPousContext differentiates from other interfaces.
	IsPousContext()
}

type PousContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyPousContext() *PousContext {
	var p = new(PousContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_pous
	return p
}

func InitEmptyPousContext(p *PousContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_pous
}

func (*PousContext) IsPousContext() {}

func NewPousContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PousContext {
	var p = new(PousContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_pous

	return p
}

func (s *PousContext) GetParser() antlr.Parser { return s.parser }

func (s *PousContext) EOF() antlr.TerminalNode {
	return s.GetToken(STParserEOF, 0)
}

func (s *PousContext) AllProgram() []IProgramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IProgramContext); ok {
			len++
		}
	}

	tst := make([]IProgramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IProgramContext); ok {
			tst[i] = t.(IProgramContext)
			i++
		}
	}

	return tst
}

func (s *PousContext) Program(i int) IProgramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgramContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgramContext)
}

func (s *PousContext) AllFunction_decl() []IFunction_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunction_declContext); ok {
			len++
		}
	}

	tst := make([]IFunction_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunction_declContext); ok {
			tst[i] = t.(IFunction_declContext)
			i++
		}
	}

	return tst
}

func (s *PousContext) Function_decl(i int) IFunction_declContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_declContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_declContext)
}

func (s *PousContext) AllFunction_block_decl() []IFunction_block_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunction_block_declContext); ok {
			len++
		}
	}

	tst := make([]IFunction_block_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunction_block_declContext); ok {
			tst[i] = t.(IFunction_block_declContext)
			i++
		}
	}

	return tst
}

func (s *PousContext) Function_block_decl(i int) IFunction_block_declContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_block_declContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_block_declContext)
}

func (s *PousContext) AllType_declaration() []IType_declarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IType_declarationContext); ok {
			len++
		}
	}

	tst := make([]IType_declarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IType_declarationContext); ok {
			tst[i] = t.(IType_declarationContext)
			i++
		}
	}

	return tst
}

func (s *PousContext) Type_declaration(i int) IType_declarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_declarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_declarationContext)
}

func (s *PousContext) AllGlobal_var_declaration() []IGlobal_var_declarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGlobal_var_declarationContext); ok {
			len++
		}
	}

	tst := make([]IGlobal_var_declarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGlobal_var_declarationContext); ok {
			tst[i] = t.(IGlobal_var_declarationContext)
			i++
		}
	}

	return tst
}

func (s *PousContext) Global_var_declaration(i int) IGlobal_var_declarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGlobal_var_declarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGlobal_var_declarationContext)
}

func (s *PousContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PousContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PousContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitPous(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Pous() (localctx IPousContext) {
	localctx = NewPousContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, STParserRULE_pous)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2251799824760832) != 0 {
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case STParserPROGRAM:
			{
				p.SetState(108)
				p.Program()
			}

		case STParserFUNCTION:
			{
				p.SetState(109)
				p.Function_decl()
			}

		case STParserFUNCTION_BLOCK:
			{
				p.SetState(110)
				p.Function_block_decl()
			}

		case STParserTYPE:
			{
				p.SetState(111)
				p.Type_declaration()
			}

		case STParserVAR_GLOBAL:
			{
				p.SetState(112)
				p.Global_var_declaration()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(118)
		p.Match(STParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// Get_var_declaration_block returns the _var_declaration_block rule contexts.
	Get_var_declaration_block() IVar_declaration_blockContext

	// GetStmts returns the stmts rule contexts.
	GetStmts() IStatement_listContext

	// Set_var_declaration_block sets the _var_declaration_block rule contexts.
	Set_var_declaration_block(IVar_declaration_blockContext)

	// SetStmts sets the stmts rule contexts.
	SetStmts(IStatement_listContext)

	// GetVars returns the vars rule context list.
	GetVars() []IVar_declaration_blockContext

	// SetVars sets the vars rule context list.
	SetVars([]IVar_declaration_blockContext)

	// Getter signatures
	PROGRAM() antlr.TerminalNode
	END_PROGRAM() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	Statement_list() IStatement_listContext
	AllVar_declaration_block() []IVar_declaration_blockContext
	Var_declaration_block(i int) IVar_declaration_blockContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*CustomContext
	parser                 antlr.Parser
	id                     antlr.Token
	_var_declaration_block IVar_declaration_blockContext
	vars                   []IVar_declaration_blockContext
	stmts                  IStatement_listContext
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) GetId() antlr.Token { return s.id }

func (s *ProgramContext) SetId(v antlr.Token) { s.id = v }

func (s *ProgramContext) Get_var_declaration_block() IVar_declaration_blockContext {
	return s._var_declaration_block
}

func (s *ProgramContext) GetStmts() IStatement_listContext { return s.stmts }

func (s *ProgramContext) Set_var_declaration_block(v IVar_declaration_blockContext) {
	s._var_declaration_block = v
}

func (s *ProgramContext) SetStmts(v IStatement_listContext) { s.stmts = v }

func (s *ProgramContext) GetVars() []IVar_declaration_blockContext { return s.vars }

func (s *ProgramContext) SetVars(v []IVar_declaration_blockContext) { s.vars = v }

func (s *ProgramContext) PROGRAM() antlr.TerminalNode {
	return s.GetToken(STParserPROGRAM, 0)
}

func (s *ProgramContext) END_PROGRAM() antlr.TerminalNode {
	return s.GetToken(STParserEND_PROGRAM, 0)
}

func (s *ProgramContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(STParserIDENTIFIER, 0)
}

func (s *ProgramContext) Statement_list() IStatement_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatement_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatement_listContext)
}

func (s *ProgramContext) AllVar_declaration_block() []IVar_declaration_blockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_declaration_blockContext); ok {
			len++
		}
	}

	tst := make([]IVar_declaration_blockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_declaration_blockContext); ok {
			tst[i] = t.(IVar_declaration_blockContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Var_declaration_block(i int) IVar_declaration_blockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declaration_blockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_declaration_blockContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, STParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Match(STParserPROGRAM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)

		var _m = p.Match(STParserIDENTIFIER)

		localctx.(*ProgramContext).id = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&260096) != 0 {
		{
			p.SetState(122)

			var _x = p.Var_declaration_block()

			localctx.(*ProgramContext)._var_declaration_block = _x
		}
		localctx.(*ProgramContext).vars = append(localctx.(*ProgramContext).vars, localctx.(*ProgramContext)._var_declaration_block)

		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(128)

		var _x = p.Statement_list()

		localctx.(*ProgramContext).stmts = _x
	}
	{
		p.SetState(129)
		p.Match(STParserEND_PROGRAM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunction_declContext is an interface to support dynamic dispatch.
type IFunction_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNCTION() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	Statement_list() IStatement_listContext
	END_FUNCTION() antlr.TerminalNode
	VAR_OUTPUT() antlr.TerminalNode
	AllEND_VAR() []antlr.TerminalNode
	END_VAR(i int) antlr.TerminalNode
	VAR() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Data_type() IData_typeContext
	VAR_INPUT() antlr.TerminalNode
	AllOutput_decl() []IOutput_declContext
	Output_decl(i int) IOutput_declContext
	AllVar_decl() []IVar_declContext
	Var_decl(i int) IVar_declContext
	AllInput_decl() []IInput_declContext
	Input_decl(i int) IInput_declContext

	// IsFunction_declContext differentiates from other interfaces.
	IsFunction_declContext()
}

type Function_declContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyFunction_declContext() *Function_declContext {
	var p = new(Function_declContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_function_decl
	return p
}

func InitEmptyFunction_declContext(p *Function_declContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_function_decl
}

func (*Function_declContext) IsFunction_declContext() {}

func NewFunction_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_declContext {
	var p = new(Function_declContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_function_decl

	return p
}

func (s *Function_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_declContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(STParserFUNCTION, 0)
}

func (s *Function_declContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(STParserIDENTIFIER, 0)
}

func (s *Function_declContext) Statement_list() IStatement_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatement_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatement_listContext)
}

func (s *Function_declContext) END_FUNCTION() antlr.TerminalNode {
	return s.GetToken(STParserEND_FUNCTION, 0)
}

func (s *Function_declContext) VAR_OUTPUT() antlr.TerminalNode {
	return s.GetToken(STParserVAR_OUTPUT, 0)
}

func (s *Function_declContext) AllEND_VAR() []antlr.TerminalNode {
	return s.GetTokens(STParserEND_VAR)
}

func (s *Function_declContext) END_VAR(i int) antlr.TerminalNode {
	return s.GetToken(STParserEND_VAR, i)
}

func (s *Function_declContext) VAR() antlr.TerminalNode {
	return s.GetToken(STParserVAR, 0)
}

func (s *Function_declContext) COLON() antlr.TerminalNode {
	return s.GetToken(STParserCOLON, 0)
}

func (s *Function_declContext) Data_type() IData_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IData_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IData_typeContext)
}

func (s *Function_declContext) VAR_INPUT() antlr.TerminalNode {
	return s.GetToken(STParserVAR_INPUT, 0)
}

func (s *Function_declContext) AllOutput_decl() []IOutput_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOutput_declContext); ok {
			len++
		}
	}

	tst := make([]IOutput_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOutput_declContext); ok {
			tst[i] = t.(IOutput_declContext)
			i++
		}
	}

	return tst
}

func (s *Function_declContext) Output_decl(i int) IOutput_declContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOutput_declContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOutput_declContext)
}

func (s *Function_declContext) AllVar_decl() []IVar_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_declContext); ok {
			len++
		}
	}

	tst := make([]IVar_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_declContext); ok {
			tst[i] = t.(IVar_declContext)
			i++
		}
	}

	return tst
}

func (s *Function_declContext) Var_decl(i int) IVar_declContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_declContext)
}

func (s *Function_declContext) AllInput_decl() []IInput_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInput_declContext); ok {
			len++
		}
	}

	tst := make([]IInput_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInput_declContext); ok {
			tst[i] = t.(IInput_declContext)
			i++
		}
	}

	return tst
}

func (s *Function_declContext) Input_decl(i int) IInput_declContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInput_declContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInput_declContext)
}

func (s *Function_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitFunction_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Function_decl() (localctx IFunction_declContext) {
	localctx = NewFunction_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, STParserRULE_function_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(STParserFUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.Match(STParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserCOLON {
		{
			p.SetState(133)
			p.Match(STParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
			p.Data_type()
		}

	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserVAR_INPUT {
		{
			p.SetState(137)
			p.Match(STParserVAR_INPUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == STParserIDENTIFIER {
			{
				p.SetState(138)
				p.Input_decl()
			}

			p.SetState(143)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(144)
			p.Match(STParserEND_VAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserVAR_OUTPUT {
		{
			p.SetState(147)
			p.Match(STParserVAR_OUTPUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == STParserIDENTIFIER {
			{
				p.SetState(148)
				p.Output_decl()
			}

			p.SetState(153)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(154)
			p.Match(STParserEND_VAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserVAR {
		{
			p.SetState(157)
			p.Match(STParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == STParserIDENTIFIER {
			{
				p.SetState(158)
				p.Var_decl()
			}

			p.SetState(163)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(164)
			p.Match(STParserEND_VAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(167)
		p.Statement_list()
	}
	{
		p.SetState(168)
		p.Match(STParserEND_FUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunction_block_declContext is an interface to support dynamic dispatch.
type IFunction_block_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNCTION_BLOCK() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	Statement_list() IStatement_listContext
	END_FUNCTION_BLOCK() antlr.TerminalNode
	VAR_INPUT() antlr.TerminalNode
	AllEND_VAR() []antlr.TerminalNode
	END_VAR(i int) antlr.TerminalNode
	VAR_OUTPUT() antlr.TerminalNode
	VAR() antlr.TerminalNode
	AllInput_decl() []IInput_declContext
	Input_decl(i int) IInput_declContext
	AllOutput_decl() []IOutput_declContext
	Output_decl(i int) IOutput_declContext
	AllVar_decl() []IVar_declContext
	Var_decl(i int) IVar_declContext

	// IsFunction_block_declContext differentiates from other interfaces.
	IsFunction_block_declContext()
}

type Function_block_declContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyFunction_block_declContext() *Function_block_declContext {
	var p = new(Function_block_declContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_function_block_decl
	return p
}

func InitEmptyFunction_block_declContext(p *Function_block_declContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_function_block_decl
}

func (*Function_block_declContext) IsFunction_block_declContext() {}

func NewFunction_block_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_block_declContext {
	var p = new(Function_block_declContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_function_block_decl

	return p
}

func (s *Function_block_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_block_declContext) FUNCTION_BLOCK() antlr.TerminalNode {
	return s.GetToken(STParserFUNCTION_BLOCK, 0)
}

func (s *Function_block_declContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(STParserIDENTIFIER, 0)
}

func (s *Function_block_declContext) Statement_list() IStatement_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatement_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatement_listContext)
}

func (s *Function_block_declContext) END_FUNCTION_BLOCK() antlr.TerminalNode {
	return s.GetToken(STParserEND_FUNCTION_BLOCK, 0)
}

func (s *Function_block_declContext) VAR_INPUT() antlr.TerminalNode {
	return s.GetToken(STParserVAR_INPUT, 0)
}

func (s *Function_block_declContext) AllEND_VAR() []antlr.TerminalNode {
	return s.GetTokens(STParserEND_VAR)
}

func (s *Function_block_declContext) END_VAR(i int) antlr.TerminalNode {
	return s.GetToken(STParserEND_VAR, i)
}

func (s *Function_block_declContext) VAR_OUTPUT() antlr.TerminalNode {
	return s.GetToken(STParserVAR_OUTPUT, 0)
}

func (s *Function_block_declContext) VAR() antlr.TerminalNode {
	return s.GetToken(STParserVAR, 0)
}

func (s *Function_block_declContext) AllInput_decl() []IInput_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInput_declContext); ok {
			len++
		}
	}

	tst := make([]IInput_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInput_declContext); ok {
			tst[i] = t.(IInput_declContext)
			i++
		}
	}

	return tst
}

func (s *Function_block_declContext) Input_decl(i int) IInput_declContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInput_declContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInput_declContext)
}

func (s *Function_block_declContext) AllOutput_decl() []IOutput_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOutput_declContext); ok {
			len++
		}
	}

	tst := make([]IOutput_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOutput_declContext); ok {
			tst[i] = t.(IOutput_declContext)
			i++
		}
	}

	return tst
}

func (s *Function_block_declContext) Output_decl(i int) IOutput_declContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOutput_declContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOutput_declContext)
}

func (s *Function_block_declContext) AllVar_decl() []IVar_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_declContext); ok {
			len++
		}
	}

	tst := make([]IVar_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_declContext); ok {
			tst[i] = t.(IVar_declContext)
			i++
		}
	}

	return tst
}

func (s *Function_block_declContext) Var_decl(i int) IVar_declContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_declContext)
}

func (s *Function_block_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_block_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_block_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitFunction_block_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Function_block_decl() (localctx IFunction_block_declContext) {
	localctx = NewFunction_block_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, STParserRULE_function_block_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(STParserFUNCTION_BLOCK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.Match(STParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserVAR_INPUT {
		{
			p.SetState(172)
			p.Match(STParserVAR_INPUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == STParserIDENTIFIER {
			{
				p.SetState(173)
				p.Input_decl()
			}

			p.SetState(178)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(179)
			p.Match(STParserEND_VAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserVAR_OUTPUT {
		{
			p.SetState(182)
			p.Match(STParserVAR_OUTPUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == STParserIDENTIFIER {
			{
				p.SetState(183)
				p.Output_decl()
			}

			p.SetState(188)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(189)
			p.Match(STParserEND_VAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserVAR {
		{
			p.SetState(192)
			p.Match(STParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == STParserIDENTIFIER {
			{
				p.SetState(193)
				p.Var_decl()
			}

			p.SetState(198)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(199)
			p.Match(STParserEND_VAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(202)
		p.Statement_list()
	}
	{
		p.SetState(203)
		p.Match(STParserEND_FUNCTION_BLOCK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IType_declarationContext is an interface to support dynamic dispatch.
type IType_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE() antlr.TerminalNode
	END_TYPE() antlr.TerminalNode
	AllType_definition() []IType_definitionContext
	Type_definition(i int) IType_definitionContext

	// IsType_declarationContext differentiates from other interfaces.
	IsType_declarationContext()
}

type Type_declarationContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyType_declarationContext() *Type_declarationContext {
	var p = new(Type_declarationContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_type_declaration
	return p
}

func InitEmptyType_declarationContext(p *Type_declarationContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_type_declaration
}

func (*Type_declarationContext) IsType_declarationContext() {}

func NewType_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_declarationContext {
	var p = new(Type_declarationContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_type_declaration

	return p
}

func (s *Type_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_declarationContext) TYPE() antlr.TerminalNode {
	return s.GetToken(STParserTYPE, 0)
}

func (s *Type_declarationContext) END_TYPE() antlr.TerminalNode {
	return s.GetToken(STParserEND_TYPE, 0)
}

func (s *Type_declarationContext) AllType_definition() []IType_definitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IType_definitionContext); ok {
			len++
		}
	}

	tst := make([]IType_definitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IType_definitionContext); ok {
			tst[i] = t.(IType_definitionContext)
			i++
		}
	}

	return tst
}

func (s *Type_declarationContext) Type_definition(i int) IType_definitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_definitionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_definitionContext)
}

func (s *Type_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitType_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Type_declaration() (localctx IType_declarationContext) {
	localctx = NewType_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, STParserRULE_type_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(STParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == STParserIDENTIFIER {
		{
			p.SetState(206)
			p.Type_definition()
		}

		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(211)
		p.Match(STParserEND_TYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IType_definitionContext is an interface to support dynamic dispatch.
type IType_definitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Data_type() IData_typeContext
	SEMICOLON() antlr.TerminalNode

	// IsType_definitionContext differentiates from other interfaces.
	IsType_definitionContext()
}

type Type_definitionContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyType_definitionContext() *Type_definitionContext {
	var p = new(Type_definitionContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_type_definition
	return p
}

func InitEmptyType_definitionContext(p *Type_definitionContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_type_definition
}

func (*Type_definitionContext) IsType_definitionContext() {}

func NewType_definitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_definitionContext {
	var p = new(Type_definitionContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_type_definition

	return p
}

func (s *Type_definitionContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_definitionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(STParserIDENTIFIER, 0)
}

func (s *Type_definitionContext) COLON() antlr.TerminalNode {
	return s.GetToken(STParserCOLON, 0)
}

func (s *Type_definitionContext) Data_type() IData_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IData_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IData_typeContext)
}

func (s *Type_definitionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(STParserSEMICOLON, 0)
}

func (s *Type_definitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_definitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_definitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitType_definition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Type_definition() (localctx IType_definitionContext) {
	localctx = NewType_definitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, STParserRULE_type_definition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.Match(STParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)
		p.Match(STParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.Data_type()
	}
	{
		p.SetState(216)
		p.Match(STParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGlobal_var_declarationContext is an interface to support dynamic dispatch.
type IGlobal_var_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR_GLOBAL() antlr.TerminalNode
	END_VAR() antlr.TerminalNode
	AllVar_decl() []IVar_declContext
	Var_decl(i int) IVar_declContext

	// IsGlobal_var_declarationContext differentiates from other interfaces.
	IsGlobal_var_declarationContext()
}

type Global_var_declarationContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyGlobal_var_declarationContext() *Global_var_declarationContext {
	var p = new(Global_var_declarationContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_global_var_declaration
	return p
}

func InitEmptyGlobal_var_declarationContext(p *Global_var_declarationContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_global_var_declaration
}

func (*Global_var_declarationContext) IsGlobal_var_declarationContext() {}

func NewGlobal_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Global_var_declarationContext {
	var p = new(Global_var_declarationContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_global_var_declaration

	return p
}

func (s *Global_var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Global_var_declarationContext) VAR_GLOBAL() antlr.TerminalNode {
	return s.GetToken(STParserVAR_GLOBAL, 0)
}

func (s *Global_var_declarationContext) END_VAR() antlr.TerminalNode {
	return s.GetToken(STParserEND_VAR, 0)
}

func (s *Global_var_declarationContext) AllVar_decl() []IVar_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_declContext); ok {
			len++
		}
	}

	tst := make([]IVar_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_declContext); ok {
			tst[i] = t.(IVar_declContext)
			i++
		}
	}

	return tst
}

func (s *Global_var_declarationContext) Var_decl(i int) IVar_declContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_declContext)
}

func (s *Global_var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Global_var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Global_var_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitGlobal_var_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Global_var_declaration() (localctx IGlobal_var_declarationContext) {
	localctx = NewGlobal_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, STParserRULE_global_var_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Match(STParserVAR_GLOBAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == STParserIDENTIFIER {
		{
			p.SetState(219)
			p.Var_decl()
		}

		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(225)
		p.Match(STParserEND_VAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_declaration_blockContext is an interface to support dynamic dispatch.
type IVar_declaration_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_var_decl returns the _var_decl rule contexts.
	Get_var_decl() IVar_declContext

	// Set_var_decl sets the _var_decl rule contexts.
	Set_var_decl(IVar_declContext)

	// GetDecls returns the decls rule context list.
	GetDecls() []IVar_declContext

	// SetDecls sets the decls rule context list.
	SetDecls([]IVar_declContext)

	// Getter signatures
	END_VAR() antlr.TerminalNode
	VAR() antlr.TerminalNode
	VAR_INPUT() antlr.TerminalNode
	VAR_OUTPUT() antlr.TerminalNode
	VAR_IN_OUT() antlr.TerminalNode
	VAR_EXTERNAL() antlr.TerminalNode
	VAR_GLOBAL() antlr.TerminalNode
	VAR_TEMP() antlr.TerminalNode
	AllVar_decl() []IVar_declContext
	Var_decl(i int) IVar_declContext

	// IsVar_declaration_blockContext differentiates from other interfaces.
	IsVar_declaration_blockContext()
}

type Var_declaration_blockContext struct {
	*CustomContext
	parser    antlr.Parser
	_var_decl IVar_declContext
	decls     []IVar_declContext
}

func NewEmptyVar_declaration_blockContext() *Var_declaration_blockContext {
	var p = new(Var_declaration_blockContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_var_declaration_block
	return p
}

func InitEmptyVar_declaration_blockContext(p *Var_declaration_blockContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_var_declaration_block
}

func (*Var_declaration_blockContext) IsVar_declaration_blockContext() {}

func NewVar_declaration_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_declaration_blockContext {
	var p = new(Var_declaration_blockContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_var_declaration_block

	return p
}

func (s *Var_declaration_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_declaration_blockContext) Get_var_decl() IVar_declContext { return s._var_decl }

func (s *Var_declaration_blockContext) Set_var_decl(v IVar_declContext) { s._var_decl = v }

func (s *Var_declaration_blockContext) GetDecls() []IVar_declContext { return s.decls }

func (s *Var_declaration_blockContext) SetDecls(v []IVar_declContext) { s.decls = v }

func (s *Var_declaration_blockContext) END_VAR() antlr.TerminalNode {
	return s.GetToken(STParserEND_VAR, 0)
}

func (s *Var_declaration_blockContext) VAR() antlr.TerminalNode {
	return s.GetToken(STParserVAR, 0)
}

func (s *Var_declaration_blockContext) VAR_INPUT() antlr.TerminalNode {
	return s.GetToken(STParserVAR_INPUT, 0)
}

func (s *Var_declaration_blockContext) VAR_OUTPUT() antlr.TerminalNode {
	return s.GetToken(STParserVAR_OUTPUT, 0)
}

func (s *Var_declaration_blockContext) VAR_IN_OUT() antlr.TerminalNode {
	return s.GetToken(STParserVAR_IN_OUT, 0)
}

func (s *Var_declaration_blockContext) VAR_EXTERNAL() antlr.TerminalNode {
	return s.GetToken(STParserVAR_EXTERNAL, 0)
}

func (s *Var_declaration_blockContext) VAR_GLOBAL() antlr.TerminalNode {
	return s.GetToken(STParserVAR_GLOBAL, 0)
}

func (s *Var_declaration_blockContext) VAR_TEMP() antlr.TerminalNode {
	return s.GetToken(STParserVAR_TEMP, 0)
}

func (s *Var_declaration_blockContext) AllVar_decl() []IVar_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_declContext); ok {
			len++
		}
	}

	tst := make([]IVar_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_declContext); ok {
			tst[i] = t.(IVar_declContext)
			i++
		}
	}

	return tst
}

func (s *Var_declaration_blockContext) Var_decl(i int) IVar_declContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_declContext)
}

func (s *Var_declaration_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_declaration_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_declaration_blockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitVar_declaration_block(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Var_declaration_block() (localctx IVar_declaration_blockContext) {
	localctx = NewVar_declaration_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, STParserRULE_var_declaration_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&260096) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == STParserIDENTIFIER {
		{
			p.SetState(228)

			var _x = p.Var_decl()

			localctx.(*Var_declaration_blockContext)._var_decl = _x
		}
		localctx.(*Var_declaration_blockContext).decls = append(localctx.(*Var_declaration_blockContext).decls, localctx.(*Var_declaration_blockContext)._var_decl)

		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(233)
		p.Match(STParserEND_VAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_declContext is an interface to support dynamic dispatch.
type IVar_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetType_ returns the type_ rule contexts.
	GetType_() IData_typeContext

	// GetExpr returns the expr rule contexts.
	GetExpr() IExpressionContext

	// SetType_ sets the type_ rule contexts.
	SetType_(IData_typeContext)

	// SetExpr sets the expr rule contexts.
	SetExpr(IExpressionContext)

	// Getter signatures
	COLON() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	Data_type() IData_typeContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsVar_declContext differentiates from other interfaces.
	IsVar_declContext()
}

type Var_declContext struct {
	*CustomContext
	parser antlr.Parser
	id     antlr.Token
	type_  IData_typeContext
	expr   IExpressionContext
}

func NewEmptyVar_declContext() *Var_declContext {
	var p = new(Var_declContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_var_decl
	return p
}

func InitEmptyVar_declContext(p *Var_declContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_var_decl
}

func (*Var_declContext) IsVar_declContext() {}

func NewVar_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_declContext {
	var p = new(Var_declContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_var_decl

	return p
}

func (s *Var_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_declContext) GetId() antlr.Token { return s.id }

func (s *Var_declContext) SetId(v antlr.Token) { s.id = v }

func (s *Var_declContext) GetType_() IData_typeContext { return s.type_ }

func (s *Var_declContext) GetExpr() IExpressionContext { return s.expr }

func (s *Var_declContext) SetType_(v IData_typeContext) { s.type_ = v }

func (s *Var_declContext) SetExpr(v IExpressionContext) { s.expr = v }

func (s *Var_declContext) COLON() antlr.TerminalNode {
	return s.GetToken(STParserCOLON, 0)
}

func (s *Var_declContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(STParserSEMICOLON, 0)
}

func (s *Var_declContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(STParserIDENTIFIER)
}

func (s *Var_declContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(STParserIDENTIFIER, i)
}

func (s *Var_declContext) Data_type() IData_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IData_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IData_typeContext)
}

func (s *Var_declContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(STParserCOMMA)
}

func (s *Var_declContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(STParserCOMMA, i)
}

func (s *Var_declContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(STParserASSIGN, 0)
}

func (s *Var_declContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Var_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitVar_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Var_decl() (localctx IVar_declContext) {
	localctx = NewVar_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, STParserRULE_var_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)

		var _m = p.Match(STParserIDENTIFIER)

		localctx.(*Var_declContext).id = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == STParserCOMMA {
		{
			p.SetState(236)
			p.Match(STParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)
			p.Match(STParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(243)
		p.Match(STParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)

		var _x = p.Data_type()

		localctx.(*Var_declContext).type_ = _x
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserASSIGN {
		{
			p.SetState(245)
			p.Match(STParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(246)

			var _x = p.expression(0)

			localctx.(*Var_declContext).expr = _x
		}

	}
	{
		p.SetState(249)
		p.Match(STParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IData_typeContext is an interface to support dynamic dispatch.
type IData_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Elementary_type_name() IElementary_type_nameContext
	Array_type() IArray_typeContext
	Structured_type() IStructured_typeContext
	IDENTIFIER() antlr.TerminalNode

	// IsData_typeContext differentiates from other interfaces.
	IsData_typeContext()
}

type Data_typeContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyData_typeContext() *Data_typeContext {
	var p = new(Data_typeContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_data_type
	return p
}

func InitEmptyData_typeContext(p *Data_typeContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_data_type
}

func (*Data_typeContext) IsData_typeContext() {}

func NewData_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Data_typeContext {
	var p = new(Data_typeContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_data_type

	return p
}

func (s *Data_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Data_typeContext) Elementary_type_name() IElementary_type_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementary_type_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElementary_type_nameContext)
}

func (s *Data_typeContext) Array_type() IArray_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_typeContext)
}

func (s *Data_typeContext) Structured_type() IStructured_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructured_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructured_typeContext)
}

func (s *Data_typeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(STParserIDENTIFIER, 0)
}

func (s *Data_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Data_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Data_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitData_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Data_type() (localctx IData_typeContext) {
	localctx = NewData_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, STParserRULE_data_type)
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case STParserINT_TYPE_NAME, STParserREAL_TYPE_NAME, STParserTIME_TYPE_NAME, STParserDATE_TYPE_NAME, STParserBOOL_TYPE_NAME, STParserBIT_TYPE_NAME, STParserSTRING_TYPE_NAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(251)
			p.Elementary_type_name()
		}

	case STParserARRAY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(252)
			p.Array_type()
		}

	case STParserSTRUCT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(253)
			p.Structured_type()
		}

	case STParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(254)
			p.Match(STParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElementary_type_nameContext is an interface to support dynamic dispatch.
type IElementary_type_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_TYPE_NAME() antlr.TerminalNode
	REAL_TYPE_NAME() antlr.TerminalNode
	BOOL_TYPE_NAME() antlr.TerminalNode
	TIME_TYPE_NAME() antlr.TerminalNode
	DATE_TYPE_NAME() antlr.TerminalNode
	STRING_TYPE_NAME() antlr.TerminalNode
	BIT_TYPE_NAME() antlr.TerminalNode

	// IsElementary_type_nameContext differentiates from other interfaces.
	IsElementary_type_nameContext()
}

type Elementary_type_nameContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyElementary_type_nameContext() *Elementary_type_nameContext {
	var p = new(Elementary_type_nameContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_elementary_type_name
	return p
}

func InitEmptyElementary_type_nameContext(p *Elementary_type_nameContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_elementary_type_name
}

func (*Elementary_type_nameContext) IsElementary_type_nameContext() {}

func NewElementary_type_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Elementary_type_nameContext {
	var p = new(Elementary_type_nameContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_elementary_type_name

	return p
}

func (s *Elementary_type_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Elementary_type_nameContext) INT_TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(STParserINT_TYPE_NAME, 0)
}

func (s *Elementary_type_nameContext) REAL_TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(STParserREAL_TYPE_NAME, 0)
}

func (s *Elementary_type_nameContext) BOOL_TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(STParserBOOL_TYPE_NAME, 0)
}

func (s *Elementary_type_nameContext) TIME_TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(STParserTIME_TYPE_NAME, 0)
}

func (s *Elementary_type_nameContext) DATE_TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(STParserDATE_TYPE_NAME, 0)
}

func (s *Elementary_type_nameContext) STRING_TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(STParserSTRING_TYPE_NAME, 0)
}

func (s *Elementary_type_nameContext) BIT_TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(STParserBIT_TYPE_NAME, 0)
}

func (s *Elementary_type_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Elementary_type_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Elementary_type_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitElementary_type_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Elementary_type_name() (localctx IElementary_type_nameContext) {
	localctx = NewElementary_type_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, STParserRULE_elementary_type_name)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1852) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArray_typeContext is an interface to support dynamic dispatch.
type IArray_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ARRAY() antlr.TerminalNode
	LBRACK() antlr.TerminalNode
	AllSubrange() []ISubrangeContext
	Subrange(i int) ISubrangeContext
	RBRACK() antlr.TerminalNode
	OF() antlr.TerminalNode
	Data_type() IData_typeContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArray_typeContext differentiates from other interfaces.
	IsArray_typeContext()
}

type Array_typeContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyArray_typeContext() *Array_typeContext {
	var p = new(Array_typeContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_array_type
	return p
}

func InitEmptyArray_typeContext(p *Array_typeContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_array_type
}

func (*Array_typeContext) IsArray_typeContext() {}

func NewArray_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_typeContext {
	var p = new(Array_typeContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_array_type

	return p
}

func (s *Array_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_typeContext) ARRAY() antlr.TerminalNode {
	return s.GetToken(STParserARRAY, 0)
}

func (s *Array_typeContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(STParserLBRACK, 0)
}

func (s *Array_typeContext) AllSubrange() []ISubrangeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISubrangeContext); ok {
			len++
		}
	}

	tst := make([]ISubrangeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISubrangeContext); ok {
			tst[i] = t.(ISubrangeContext)
			i++
		}
	}

	return tst
}

func (s *Array_typeContext) Subrange(i int) ISubrangeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubrangeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubrangeContext)
}

func (s *Array_typeContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(STParserRBRACK, 0)
}

func (s *Array_typeContext) OF() antlr.TerminalNode {
	return s.GetToken(STParserOF, 0)
}

func (s *Array_typeContext) Data_type() IData_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IData_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IData_typeContext)
}

func (s *Array_typeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(STParserCOMMA)
}

func (s *Array_typeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(STParserCOMMA, i)
}

func (s *Array_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitArray_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Array_type() (localctx IArray_typeContext) {
	localctx = NewArray_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, STParserRULE_array_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Match(STParserARRAY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(260)
		p.Match(STParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
		p.Subrange()
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == STParserCOMMA {
		{
			p.SetState(262)
			p.Match(STParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(263)
			p.Subrange()
		}

		p.SetState(268)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(269)
		p.Match(STParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(270)
		p.Match(STParserOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(271)
		p.Data_type()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISubrangeContext is an interface to support dynamic dispatch.
type ISubrangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	DOTDOT() antlr.TerminalNode

	// IsSubrangeContext differentiates from other interfaces.
	IsSubrangeContext()
}

type SubrangeContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptySubrangeContext() *SubrangeContext {
	var p = new(SubrangeContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_subrange
	return p
}

func InitEmptySubrangeContext(p *SubrangeContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_subrange
}

func (*SubrangeContext) IsSubrangeContext() {}

func NewSubrangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubrangeContext {
	var p = new(SubrangeContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_subrange

	return p
}

func (s *SubrangeContext) GetParser() antlr.Parser { return s.parser }

func (s *SubrangeContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *SubrangeContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SubrangeContext) DOTDOT() antlr.TerminalNode {
	return s.GetToken(STParserDOTDOT, 0)
}

func (s *SubrangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubrangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubrangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitSubrange(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Subrange() (localctx ISubrangeContext) {
	localctx = NewSubrangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, STParserRULE_subrange)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
		p.expression(0)
	}
	{
		p.SetState(274)
		p.Match(STParserDOTDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(275)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructured_typeContext is an interface to support dynamic dispatch.
type IStructured_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRUCT() antlr.TerminalNode
	END_STRUCT() antlr.TerminalNode
	AllVar_decl() []IVar_declContext
	Var_decl(i int) IVar_declContext

	// IsStructured_typeContext differentiates from other interfaces.
	IsStructured_typeContext()
}

type Structured_typeContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyStructured_typeContext() *Structured_typeContext {
	var p = new(Structured_typeContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_structured_type
	return p
}

func InitEmptyStructured_typeContext(p *Structured_typeContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_structured_type
}

func (*Structured_typeContext) IsStructured_typeContext() {}

func NewStructured_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Structured_typeContext {
	var p = new(Structured_typeContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_structured_type

	return p
}

func (s *Structured_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Structured_typeContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(STParserSTRUCT, 0)
}

func (s *Structured_typeContext) END_STRUCT() antlr.TerminalNode {
	return s.GetToken(STParserEND_STRUCT, 0)
}

func (s *Structured_typeContext) AllVar_decl() []IVar_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_declContext); ok {
			len++
		}
	}

	tst := make([]IVar_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_declContext); ok {
			tst[i] = t.(IVar_declContext)
			i++
		}
	}

	return tst
}

func (s *Structured_typeContext) Var_decl(i int) IVar_declContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_declContext)
}

func (s *Structured_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Structured_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Structured_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitStructured_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Structured_type() (localctx IStructured_typeContext) {
	localctx = NewStructured_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, STParserRULE_structured_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(277)
		p.Match(STParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == STParserIDENTIFIER {
		{
			p.SetState(278)
			p.Var_decl()
		}

		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(283)
		p.Match(STParserEND_STRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatement_listContext is an interface to support dynamic dispatch.
type IStatement_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsStatement_listContext differentiates from other interfaces.
	IsStatement_listContext()
}

type Statement_listContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyStatement_listContext() *Statement_listContext {
	var p = new(Statement_listContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_statement_list
	return p
}

func InitEmptyStatement_listContext(p *Statement_listContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_statement_list
}

func (*Statement_listContext) IsStatement_listContext() {}

func NewStatement_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Statement_listContext {
	var p = new(Statement_listContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_statement_list

	return p
}

func (s *Statement_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Statement_listContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *Statement_listContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *Statement_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Statement_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Statement_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitStatement_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Statement_list() (localctx IStatement_listContext) {
	localctx = NewStatement_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, STParserRULE_statement_list)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(285)
				p.Statement()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment_statement() IAssignment_statementContext
	SEMICOLON() antlr.TerminalNode
	Function_invocation() IFunction_invocationContext
	Continue_statement() IContinue_statementContext
	Exit_statement() IExit_statementContext
	Return_statement() IReturn_statementContext
	If_statement() IIf_statementContext
	Case_statement() ICase_statementContext
	For_statement() IFor_statementContext
	While_statement() IWhile_statementContext
	Repeat_statement() IRepeat_statementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Assignment_statement() IAssignment_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignment_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignment_statementContext)
}

func (s *StatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(STParserSEMICOLON, 0)
}

func (s *StatementContext) Function_invocation() IFunction_invocationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_invocationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_invocationContext)
}

func (s *StatementContext) Continue_statement() IContinue_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinue_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinue_statementContext)
}

func (s *StatementContext) Exit_statement() IExit_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExit_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExit_statementContext)
}

func (s *StatementContext) Return_statement() IReturn_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturn_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturn_statementContext)
}

func (s *StatementContext) If_statement() IIf_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_statementContext)
}

func (s *StatementContext) Case_statement() ICase_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICase_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICase_statementContext)
}

func (s *StatementContext) For_statement() IFor_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_statementContext)
}

func (s *StatementContext) While_statement() IWhile_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhile_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhile_statementContext)
}

func (s *StatementContext) Repeat_statement() IRepeat_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRepeat_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRepeat_statementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, STParserRULE_statement)
	var _la int

	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(290)
			p.Assignment_statement()
		}
		{
			p.SetState(291)
			p.Match(STParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(293)
			p.Function_invocation()
		}
		{
			p.SetState(294)
			p.Match(STParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(296)
			p.Continue_statement()
		}
		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == STParserSEMICOLON {
			{
				p.SetState(297)
				p.Match(STParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(300)
			p.Exit_statement()
		}
		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == STParserSEMICOLON {
			{
				p.SetState(301)
				p.Match(STParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(304)
			p.Return_statement()
		}
		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == STParserSEMICOLON {
			{
				p.SetState(305)
				p.Match(STParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(308)
			p.If_statement()
		}
		p.SetState(310)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == STParserSEMICOLON {
			{
				p.SetState(309)
				p.Match(STParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(312)
			p.Case_statement()
		}
		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == STParserSEMICOLON {
			{
				p.SetState(313)
				p.Match(STParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(316)
			p.For_statement()
		}
		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == STParserSEMICOLON {
			{
				p.SetState(317)
				p.Match(STParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(320)
			p.While_statement()
		}
		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == STParserSEMICOLON {
			{
				p.SetState(321)
				p.Match(STParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(324)
			p.Repeat_statement()
		}
		p.SetState(326)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == STParserSEMICOLON {
			{
				p.SetState(325)
				p.Match(STParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignment_statementContext is an interface to support dynamic dispatch.
type IAssignment_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Variable() IVariableContext
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsAssignment_statementContext differentiates from other interfaces.
	IsAssignment_statementContext()
}

type Assignment_statementContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyAssignment_statementContext() *Assignment_statementContext {
	var p = new(Assignment_statementContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_assignment_statement
	return p
}

func InitEmptyAssignment_statementContext(p *Assignment_statementContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_assignment_statement
}

func (*Assignment_statementContext) IsAssignment_statementContext() {}

func NewAssignment_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assignment_statementContext {
	var p = new(Assignment_statementContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_assignment_statement

	return p
}

func (s *Assignment_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Assignment_statementContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *Assignment_statementContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(STParserASSIGN, 0)
}

func (s *Assignment_statementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Assignment_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assignment_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assignment_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitAssignment_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Assignment_statement() (localctx IAssignment_statementContext) {
	localctx = NewAssignment_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, STParserRULE_assignment_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
		p.Variable()
	}
	{
		p.SetState(331)
		p.Match(STParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(332)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_statementContext is an interface to support dynamic dispatch.
type IIf_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_statement_list returns the _statement_list rule contexts.
	Get_statement_list() IStatement_listContext

	// GetElse_ returns the else_ rule contexts.
	GetElse_() IStatement_listContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_statement_list sets the _statement_list rule contexts.
	Set_statement_list(IStatement_listContext)

	// SetElse_ sets the else_ rule contexts.
	SetElse_(IStatement_listContext)

	// GetConds returns the conds rule context list.
	GetConds() []IExpressionContext

	// GetThens returns the thens rule context list.
	GetThens() []IStatement_listContext

	// SetConds sets the conds rule context list.
	SetConds([]IExpressionContext)

	// SetThens sets the thens rule context list.
	SetThens([]IStatement_listContext)

	// Getter signatures
	IF() antlr.TerminalNode
	AllTHEN() []antlr.TerminalNode
	THEN(i int) antlr.TerminalNode
	END_IF() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllStatement_list() []IStatement_listContext
	Statement_list(i int) IStatement_listContext
	AllELSIF() []antlr.TerminalNode
	ELSIF(i int) antlr.TerminalNode
	ELSE() antlr.TerminalNode

	// IsIf_statementContext differentiates from other interfaces.
	IsIf_statementContext()
}

type If_statementContext struct {
	*CustomContext
	parser          antlr.Parser
	_expression     IExpressionContext
	conds           []IExpressionContext
	_statement_list IStatement_listContext
	thens           []IStatement_listContext
	else_           IStatement_listContext
}

func NewEmptyIf_statementContext() *If_statementContext {
	var p = new(If_statementContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_if_statement
	return p
}

func InitEmptyIf_statementContext(p *If_statementContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_if_statement
}

func (*If_statementContext) IsIf_statementContext() {}

func NewIf_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_statementContext {
	var p = new(If_statementContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_if_statement

	return p
}

func (s *If_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *If_statementContext) Get_expression() IExpressionContext { return s._expression }

func (s *If_statementContext) Get_statement_list() IStatement_listContext { return s._statement_list }

func (s *If_statementContext) GetElse_() IStatement_listContext { return s.else_ }

func (s *If_statementContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *If_statementContext) Set_statement_list(v IStatement_listContext) { s._statement_list = v }

func (s *If_statementContext) SetElse_(v IStatement_listContext) { s.else_ = v }

func (s *If_statementContext) GetConds() []IExpressionContext { return s.conds }

func (s *If_statementContext) GetThens() []IStatement_listContext { return s.thens }

func (s *If_statementContext) SetConds(v []IExpressionContext) { s.conds = v }

func (s *If_statementContext) SetThens(v []IStatement_listContext) { s.thens = v }

func (s *If_statementContext) IF() antlr.TerminalNode {
	return s.GetToken(STParserIF, 0)
}

func (s *If_statementContext) AllTHEN() []antlr.TerminalNode {
	return s.GetTokens(STParserTHEN)
}

func (s *If_statementContext) THEN(i int) antlr.TerminalNode {
	return s.GetToken(STParserTHEN, i)
}

func (s *If_statementContext) END_IF() antlr.TerminalNode {
	return s.GetToken(STParserEND_IF, 0)
}

func (s *If_statementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *If_statementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *If_statementContext) AllStatement_list() []IStatement_listContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatement_listContext); ok {
			len++
		}
	}

	tst := make([]IStatement_listContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatement_listContext); ok {
			tst[i] = t.(IStatement_listContext)
			i++
		}
	}

	return tst
}

func (s *If_statementContext) Statement_list(i int) IStatement_listContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatement_listContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatement_listContext)
}

func (s *If_statementContext) AllELSIF() []antlr.TerminalNode {
	return s.GetTokens(STParserELSIF)
}

func (s *If_statementContext) ELSIF(i int) antlr.TerminalNode {
	return s.GetToken(STParserELSIF, i)
}

func (s *If_statementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(STParserELSE, 0)
}

func (s *If_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitIf_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) If_statement() (localctx IIf_statementContext) {
	localctx = NewIf_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, STParserRULE_if_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(334)
		p.Match(STParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(335)

		var _x = p.expression(0)

		localctx.(*If_statementContext)._expression = _x
	}
	localctx.(*If_statementContext).conds = append(localctx.(*If_statementContext).conds, localctx.(*If_statementContext)._expression)
	{
		p.SetState(336)
		p.Match(STParserTHEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(337)

		var _x = p.Statement_list()

		localctx.(*If_statementContext)._statement_list = _x
	}
	localctx.(*If_statementContext).thens = append(localctx.(*If_statementContext).thens, localctx.(*If_statementContext)._statement_list)
	p.SetState(345)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == STParserELSIF {
		{
			p.SetState(338)
			p.Match(STParserELSIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)

			var _x = p.expression(0)

			localctx.(*If_statementContext)._expression = _x
		}
		localctx.(*If_statementContext).conds = append(localctx.(*If_statementContext).conds, localctx.(*If_statementContext)._expression)
		{
			p.SetState(340)
			p.Match(STParserTHEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)

			var _x = p.Statement_list()

			localctx.(*If_statementContext)._statement_list = _x
		}
		localctx.(*If_statementContext).thens = append(localctx.(*If_statementContext).thens, localctx.(*If_statementContext)._statement_list)

		p.SetState(347)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserELSE {
		{
			p.SetState(348)
			p.Match(STParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(349)

			var _x = p.Statement_list()

			localctx.(*If_statementContext).else_ = _x
		}

	}
	{
		p.SetState(352)
		p.Match(STParserEND_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICase_statementContext is an interface to support dynamic dispatch.
type ICase_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	Expression() IExpressionContext
	OF() antlr.TerminalNode
	END_CASE() antlr.TerminalNode
	AllCase_element() []ICase_elementContext
	Case_element(i int) ICase_elementContext
	ELSE() antlr.TerminalNode
	Statement_list() IStatement_listContext

	// IsCase_statementContext differentiates from other interfaces.
	IsCase_statementContext()
}

type Case_statementContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyCase_statementContext() *Case_statementContext {
	var p = new(Case_statementContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_case_statement
	return p
}

func InitEmptyCase_statementContext(p *Case_statementContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_case_statement
}

func (*Case_statementContext) IsCase_statementContext() {}

func NewCase_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Case_statementContext {
	var p = new(Case_statementContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_case_statement

	return p
}

func (s *Case_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Case_statementContext) CASE() antlr.TerminalNode {
	return s.GetToken(STParserCASE, 0)
}

func (s *Case_statementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Case_statementContext) OF() antlr.TerminalNode {
	return s.GetToken(STParserOF, 0)
}

func (s *Case_statementContext) END_CASE() antlr.TerminalNode {
	return s.GetToken(STParserEND_CASE, 0)
}

func (s *Case_statementContext) AllCase_element() []ICase_elementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICase_elementContext); ok {
			len++
		}
	}

	tst := make([]ICase_elementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICase_elementContext); ok {
			tst[i] = t.(ICase_elementContext)
			i++
		}
	}

	return tst
}

func (s *Case_statementContext) Case_element(i int) ICase_elementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICase_elementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICase_elementContext)
}

func (s *Case_statementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(STParserELSE, 0)
}

func (s *Case_statementContext) Statement_list() IStatement_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatement_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatement_listContext)
}

func (s *Case_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Case_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Case_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitCase_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Case_statement() (localctx ICase_statementContext) {
	localctx = NewCase_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, STParserRULE_case_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(354)
		p.Match(STParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(355)
		p.expression(0)
	}
	{
		p.SetState(356)
		p.Match(STParserOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17592186045196) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&4162563) != 0) {
		{
			p.SetState(357)
			p.Case_element()
		}

		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserELSE {
		{
			p.SetState(362)
			p.Match(STParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(363)
			p.Statement_list()
		}

	}
	{
		p.SetState(366)
		p.Match(STParserEND_CASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICase_elementContext is an interface to support dynamic dispatch.
type ICase_elementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCase_label() []ICase_labelContext
	Case_label(i int) ICase_labelContext
	COLON() antlr.TerminalNode
	Statement_list() IStatement_listContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsCase_elementContext differentiates from other interfaces.
	IsCase_elementContext()
}

type Case_elementContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyCase_elementContext() *Case_elementContext {
	var p = new(Case_elementContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_case_element
	return p
}

func InitEmptyCase_elementContext(p *Case_elementContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_case_element
}

func (*Case_elementContext) IsCase_elementContext() {}

func NewCase_elementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Case_elementContext {
	var p = new(Case_elementContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_case_element

	return p
}

func (s *Case_elementContext) GetParser() antlr.Parser { return s.parser }

func (s *Case_elementContext) AllCase_label() []ICase_labelContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICase_labelContext); ok {
			len++
		}
	}

	tst := make([]ICase_labelContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICase_labelContext); ok {
			tst[i] = t.(ICase_labelContext)
			i++
		}
	}

	return tst
}

func (s *Case_elementContext) Case_label(i int) ICase_labelContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICase_labelContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICase_labelContext)
}

func (s *Case_elementContext) COLON() antlr.TerminalNode {
	return s.GetToken(STParserCOLON, 0)
}

func (s *Case_elementContext) Statement_list() IStatement_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatement_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatement_listContext)
}

func (s *Case_elementContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(STParserCOMMA)
}

func (s *Case_elementContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(STParserCOMMA, i)
}

func (s *Case_elementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Case_elementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Case_elementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitCase_element(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Case_element() (localctx ICase_elementContext) {
	localctx = NewCase_elementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, STParserRULE_case_element)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
		p.Case_label()
	}
	p.SetState(373)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == STParserCOMMA {
		{
			p.SetState(369)
			p.Match(STParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(370)
			p.Case_label()
		}

		p.SetState(375)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(376)
		p.Match(STParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(377)
		p.Statement_list()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICase_labelContext is an interface to support dynamic dispatch.
type ICase_labelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	Subrange() ISubrangeContext

	// IsCase_labelContext differentiates from other interfaces.
	IsCase_labelContext()
}

type Case_labelContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyCase_labelContext() *Case_labelContext {
	var p = new(Case_labelContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_case_label
	return p
}

func InitEmptyCase_labelContext(p *Case_labelContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_case_label
}

func (*Case_labelContext) IsCase_labelContext() {}

func NewCase_labelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Case_labelContext {
	var p = new(Case_labelContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_case_label

	return p
}

func (s *Case_labelContext) GetParser() antlr.Parser { return s.parser }

func (s *Case_labelContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Case_labelContext) Subrange() ISubrangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubrangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubrangeContext)
}

func (s *Case_labelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Case_labelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Case_labelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitCase_label(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Case_label() (localctx ICase_labelContext) {
	localctx = NewCase_labelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, STParserRULE_case_label)
	p.SetState(381)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(379)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(380)
			p.Subrange()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFor_statementContext is an interface to support dynamic dispatch.
type IFor_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	TO() antlr.TerminalNode
	DO() antlr.TerminalNode
	Statement_list() IStatement_listContext
	END_FOR() antlr.TerminalNode
	BY() antlr.TerminalNode

	// IsFor_statementContext differentiates from other interfaces.
	IsFor_statementContext()
}

type For_statementContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyFor_statementContext() *For_statementContext {
	var p = new(For_statementContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_for_statement
	return p
}

func InitEmptyFor_statementContext(p *For_statementContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_for_statement
}

func (*For_statementContext) IsFor_statementContext() {}

func NewFor_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_statementContext {
	var p = new(For_statementContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_for_statement

	return p
}

func (s *For_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *For_statementContext) FOR() antlr.TerminalNode {
	return s.GetToken(STParserFOR, 0)
}

func (s *For_statementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(STParserIDENTIFIER, 0)
}

func (s *For_statementContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(STParserASSIGN, 0)
}

func (s *For_statementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *For_statementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *For_statementContext) TO() antlr.TerminalNode {
	return s.GetToken(STParserTO, 0)
}

func (s *For_statementContext) DO() antlr.TerminalNode {
	return s.GetToken(STParserDO, 0)
}

func (s *For_statementContext) Statement_list() IStatement_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatement_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatement_listContext)
}

func (s *For_statementContext) END_FOR() antlr.TerminalNode {
	return s.GetToken(STParserEND_FOR, 0)
}

func (s *For_statementContext) BY() antlr.TerminalNode {
	return s.GetToken(STParserBY, 0)
}

func (s *For_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitFor_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) For_statement() (localctx IFor_statementContext) {
	localctx = NewFor_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, STParserRULE_for_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(383)
		p.Match(STParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(384)
		p.Match(STParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(385)
		p.Match(STParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(386)
		p.expression(0)
	}
	{
		p.SetState(387)
		p.Match(STParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(388)
		p.expression(0)
	}
	p.SetState(391)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserBY {
		{
			p.SetState(389)
			p.Match(STParserBY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(390)
			p.expression(0)
		}

	}
	{
		p.SetState(393)
		p.Match(STParserDO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(394)
		p.Statement_list()
	}
	{
		p.SetState(395)
		p.Match(STParserEND_FOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhile_statementContext is an interface to support dynamic dispatch.
type IWhile_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expression() IExpressionContext
	DO() antlr.TerminalNode
	Statement_list() IStatement_listContext
	END_WHILE() antlr.TerminalNode

	// IsWhile_statementContext differentiates from other interfaces.
	IsWhile_statementContext()
}

type While_statementContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyWhile_statementContext() *While_statementContext {
	var p = new(While_statementContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_while_statement
	return p
}

func InitEmptyWhile_statementContext(p *While_statementContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_while_statement
}

func (*While_statementContext) IsWhile_statementContext() {}

func NewWhile_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_statementContext {
	var p = new(While_statementContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_while_statement

	return p
}

func (s *While_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *While_statementContext) WHILE() antlr.TerminalNode {
	return s.GetToken(STParserWHILE, 0)
}

func (s *While_statementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *While_statementContext) DO() antlr.TerminalNode {
	return s.GetToken(STParserDO, 0)
}

func (s *While_statementContext) Statement_list() IStatement_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatement_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatement_listContext)
}

func (s *While_statementContext) END_WHILE() antlr.TerminalNode {
	return s.GetToken(STParserEND_WHILE, 0)
}

func (s *While_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *While_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitWhile_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) While_statement() (localctx IWhile_statementContext) {
	localctx = NewWhile_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, STParserRULE_while_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(397)
		p.Match(STParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(398)
		p.expression(0)
	}
	{
		p.SetState(399)
		p.Match(STParserDO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(400)
		p.Statement_list()
	}
	{
		p.SetState(401)
		p.Match(STParserEND_WHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRepeat_statementContext is an interface to support dynamic dispatch.
type IRepeat_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REPEAT() antlr.TerminalNode
	Statement_list() IStatement_listContext
	UNTIL() antlr.TerminalNode
	Expression() IExpressionContext
	END_REPEAT() antlr.TerminalNode

	// IsRepeat_statementContext differentiates from other interfaces.
	IsRepeat_statementContext()
}

type Repeat_statementContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyRepeat_statementContext() *Repeat_statementContext {
	var p = new(Repeat_statementContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_repeat_statement
	return p
}

func InitEmptyRepeat_statementContext(p *Repeat_statementContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_repeat_statement
}

func (*Repeat_statementContext) IsRepeat_statementContext() {}

func NewRepeat_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Repeat_statementContext {
	var p = new(Repeat_statementContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_repeat_statement

	return p
}

func (s *Repeat_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Repeat_statementContext) REPEAT() antlr.TerminalNode {
	return s.GetToken(STParserREPEAT, 0)
}

func (s *Repeat_statementContext) Statement_list() IStatement_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatement_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatement_listContext)
}

func (s *Repeat_statementContext) UNTIL() antlr.TerminalNode {
	return s.GetToken(STParserUNTIL, 0)
}

func (s *Repeat_statementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Repeat_statementContext) END_REPEAT() antlr.TerminalNode {
	return s.GetToken(STParserEND_REPEAT, 0)
}

func (s *Repeat_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Repeat_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Repeat_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitRepeat_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Repeat_statement() (localctx IRepeat_statementContext) {
	localctx = NewRepeat_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, STParserRULE_repeat_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(403)
		p.Match(STParserREPEAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(404)
		p.Statement_list()
	}
	{
		p.SetState(405)
		p.Match(STParserUNTIL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(406)
		p.expression(0)
	}
	{
		p.SetState(407)
		p.Match(STParserEND_REPEAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunction_invocationContext is an interface to support dynamic dispatch.
type IFunction_invocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetArgs returns the args rule context list.
	GetArgs() []IExpressionContext

	// SetArgs sets the args rule context list.
	SetArgs([]IExpressionContext)

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFunction_invocationContext differentiates from other interfaces.
	IsFunction_invocationContext()
}

type Function_invocationContext struct {
	*CustomContext
	parser      antlr.Parser
	_expression IExpressionContext
	args        []IExpressionContext
}

func NewEmptyFunction_invocationContext() *Function_invocationContext {
	var p = new(Function_invocationContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_function_invocation
	return p
}

func InitEmptyFunction_invocationContext(p *Function_invocationContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_function_invocation
}

func (*Function_invocationContext) IsFunction_invocationContext() {}

func NewFunction_invocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_invocationContext {
	var p = new(Function_invocationContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_function_invocation

	return p
}

func (s *Function_invocationContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_invocationContext) Get_expression() IExpressionContext { return s._expression }

func (s *Function_invocationContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Function_invocationContext) GetArgs() []IExpressionContext { return s.args }

func (s *Function_invocationContext) SetArgs(v []IExpressionContext) { s.args = v }

func (s *Function_invocationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(STParserIDENTIFIER, 0)
}

func (s *Function_invocationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(STParserLPAREN, 0)
}

func (s *Function_invocationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(STParserRPAREN, 0)
}

func (s *Function_invocationContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *Function_invocationContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Function_invocationContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(STParserCOMMA)
}

func (s *Function_invocationContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(STParserCOMMA, i)
}

func (s *Function_invocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_invocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_invocationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitFunction_invocation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Function_invocation() (localctx IFunction_invocationContext) {
	localctx = NewFunction_invocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, STParserRULE_function_invocation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(409)
		p.Match(STParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(410)
		p.Match(STParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17592186045196) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&4162563) != 0) {
		{
			p.SetState(411)

			var _x = p.expression(0)

			localctx.(*Function_invocationContext)._expression = _x
		}
		localctx.(*Function_invocationContext).args = append(localctx.(*Function_invocationContext).args, localctx.(*Function_invocationContext)._expression)
		p.SetState(416)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == STParserCOMMA {
			{
				p.SetState(412)
				p.Match(STParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(413)

				var _x = p.expression(0)

				localctx.(*Function_invocationContext)._expression = _x
			}
			localctx.(*Function_invocationContext).args = append(localctx.(*Function_invocationContext).args, localctx.(*Function_invocationContext)._expression)

			p.SetState(418)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(421)
		p.Match(STParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturn_statementContext is an interface to support dynamic dispatch.
type IReturn_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsReturn_statementContext differentiates from other interfaces.
	IsReturn_statementContext()
}

type Return_statementContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyReturn_statementContext() *Return_statementContext {
	var p = new(Return_statementContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_return_statement
	return p
}

func InitEmptyReturn_statementContext(p *Return_statementContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_return_statement
}

func (*Return_statementContext) IsReturn_statementContext() {}

func NewReturn_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_statementContext {
	var p = new(Return_statementContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_return_statement

	return p
}

func (s *Return_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_statementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(STParserRETURN, 0)
}

func (s *Return_statementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Return_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitReturn_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Return_statement() (localctx IReturn_statementContext) {
	localctx = NewReturn_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, STParserRULE_return_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(423)
		p.Match(STParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(425)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(424)
			p.expression(0)
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContinue_statementContext is an interface to support dynamic dispatch.
type IContinue_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTINUE() antlr.TerminalNode

	// IsContinue_statementContext differentiates from other interfaces.
	IsContinue_statementContext()
}

type Continue_statementContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyContinue_statementContext() *Continue_statementContext {
	var p = new(Continue_statementContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_continue_statement
	return p
}

func InitEmptyContinue_statementContext(p *Continue_statementContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_continue_statement
}

func (*Continue_statementContext) IsContinue_statementContext() {}

func NewContinue_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Continue_statementContext {
	var p = new(Continue_statementContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_continue_statement

	return p
}

func (s *Continue_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Continue_statementContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(STParserCONTINUE, 0)
}

func (s *Continue_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Continue_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Continue_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitContinue_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Continue_statement() (localctx IContinue_statementContext) {
	localctx = NewContinue_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, STParserRULE_continue_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(427)
		p.Match(STParserCONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExit_statementContext is an interface to support dynamic dispatch.
type IExit_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXIT() antlr.TerminalNode

	// IsExit_statementContext differentiates from other interfaces.
	IsExit_statementContext()
}

type Exit_statementContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyExit_statementContext() *Exit_statementContext {
	var p = new(Exit_statementContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_exit_statement
	return p
}

func InitEmptyExit_statementContext(p *Exit_statementContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_exit_statement
}

func (*Exit_statementContext) IsExit_statementContext() {}

func NewExit_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exit_statementContext {
	var p = new(Exit_statementContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_exit_statement

	return p
}

func (s *Exit_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Exit_statementContext) EXIT() antlr.TerminalNode {
	return s.GetToken(STParserEXIT, 0)
}

func (s *Exit_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exit_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exit_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitExit_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Exit_statement() (localctx IExit_statementContext) {
	localctx = NewExit_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, STParserRULE_exit_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(429)
		p.Match(STParserEXIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BinaryExpressionContext struct {
	ExpressionContext
	left     IExpressionContext
	operator antlr.Token
	right    IExpressionContext
}

func NewBinaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryExpressionContext {
	var p = new(BinaryExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryExpressionContext) GetOperator() antlr.Token { return s.operator }

func (s *BinaryExpressionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *BinaryExpressionContext) GetLeft() IExpressionContext { return s.left }

func (s *BinaryExpressionContext) GetRight() IExpressionContext { return s.right }

func (s *BinaryExpressionContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *BinaryExpressionContext) SetRight(v IExpressionContext) { s.right = v }

func (s *BinaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BinaryExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinaryExpressionContext) POWER() antlr.TerminalNode {
	return s.GetToken(STParserPOWER, 0)
}

func (s *BinaryExpressionContext) MULT() antlr.TerminalNode {
	return s.GetToken(STParserMULT, 0)
}

func (s *BinaryExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(STParserDIV, 0)
}

func (s *BinaryExpressionContext) MOD() antlr.TerminalNode {
	return s.GetToken(STParserMOD, 0)
}

func (s *BinaryExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(STParserPLUS, 0)
}

func (s *BinaryExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(STParserMINUS, 0)
}

func (s *BinaryExpressionContext) LESS() antlr.TerminalNode {
	return s.GetToken(STParserLESS, 0)
}

func (s *BinaryExpressionContext) LESS_EQ() antlr.TerminalNode {
	return s.GetToken(STParserLESS_EQ, 0)
}

func (s *BinaryExpressionContext) GREATER() antlr.TerminalNode {
	return s.GetToken(STParserGREATER, 0)
}

func (s *BinaryExpressionContext) GREATER_EQ() antlr.TerminalNode {
	return s.GetToken(STParserGREATER_EQ, 0)
}

func (s *BinaryExpressionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(STParserEQUAL, 0)
}

func (s *BinaryExpressionContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(STParserNOT_EQUAL, 0)
}

func (s *BinaryExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(STParserAND, 0)
}

func (s *BinaryExpressionContext) XOR() antlr.TerminalNode {
	return s.GetToken(STParserXOR, 0)
}

func (s *BinaryExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(STParserOR, 0)
}

func (s *BinaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitBinaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarExpressionContext struct {
	ExpressionContext
}

func NewVarExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarExpressionContext {
	var p = new(VarExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *VarExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarExpressionContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VarExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitVarExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExpressionContext struct {
	ExpressionContext
}

func NewParenExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExpressionContext {
	var p = new(ParenExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ParenExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(STParserLPAREN, 0)
}

func (s *ParenExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(STParserRPAREN, 0)
}

func (s *ParenExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitParenExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryExpressionContext struct {
	ExpressionContext
	operator antlr.Token
}

func NewUnaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryExpressionContext) GetOperator() antlr.Token { return s.operator }

func (s *UnaryExpressionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(STParserPLUS, 0)
}

func (s *UnaryExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(STParserMINUS, 0)
}

func (s *UnaryExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(STParserNOT, 0)
}

func (s *UnaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitUnaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralExpressionContext struct {
	ExpressionContext
}

func NewLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpressionContext {
	var p = new(LiteralExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpressionContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitLiteralExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type FuncCallExpressionContext struct {
	ExpressionContext
}

func NewFuncCallExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncCallExpressionContext {
	var p = new(FuncCallExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *FuncCallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallExpressionContext) Function_invocation() IFunction_invocationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_invocationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_invocationContext)
}

func (s *FuncCallExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitFuncCallExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *STParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 56
	p.EnterRecursionRule(localctx, 56, STParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(441)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(432)
			p.Match(STParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(433)
			p.expression(0)
		}
		{
			p.SetState(434)
			p.Match(STParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewUnaryExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(436)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryExpressionContext).operator = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64((_la-44)) & ^0x3f) == 0 && ((int64(1)<<(_la-44))&6291457) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryExpressionContext).operator = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(437)
			p.expression(12)
		}

	case 3:
		localctx = NewLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(438)
			p.Literal()
		}

	case 4:
		localctx = NewVarExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(439)
			p.Variable()
		}

	case 5:
		localctx = NewFuncCallExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(440)
			p.Function_invocation()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(469)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(467)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, STParserRULE_expression)
				p.SetState(443)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(444)

					var _m = p.Match(STParserPOWER)

					localctx.(*BinaryExpressionContext).operator = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(445)

					var _x = p.expression(12)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case 2:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, STParserRULE_expression)
				p.SetState(446)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(447)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpressionContext).operator = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-67)) & ^0x3f) == 0 && ((int64(1)<<(_la-67))&7) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpressionContext).operator = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(448)

					var _x = p.expression(11)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case 3:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, STParserRULE_expression)
				p.SetState(449)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(450)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpressionContext).operator = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == STParserPLUS || _la == STParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpressionContext).operator = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(451)

					var _x = p.expression(10)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case 4:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, STParserRULE_expression)
				p.SetState(452)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(453)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpressionContext).operator = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-61)) & ^0x3f) == 0 && ((int64(1)<<(_la-61))&15) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpressionContext).operator = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(454)

					var _x = p.expression(9)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case 5:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, STParserRULE_expression)
				p.SetState(455)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(456)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpressionContext).operator = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == STParserEQUAL || _la == STParserNOT_EQUAL) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpressionContext).operator = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(457)

					var _x = p.expression(8)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case 6:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, STParserRULE_expression)
				p.SetState(458)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(459)

					var _m = p.Match(STParserAND)

					localctx.(*BinaryExpressionContext).operator = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(460)

					var _x = p.expression(7)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case 7:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, STParserRULE_expression)
				p.SetState(461)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(462)

					var _m = p.Match(STParserXOR)

					localctx.(*BinaryExpressionContext).operator = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(463)

					var _x = p.expression(6)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case 8:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, STParserRULE_expression)
				p.SetState(464)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(465)

					var _m = p.Match(STParserOR)

					localctx.(*BinaryExpressionContext).operator = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(466)

					var _x = p.expression(5)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(471)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_LITERAL() antlr.TerminalNode
	REAL_LITERAL() antlr.TerminalNode
	BOOL_LITERAL() antlr.TerminalNode
	TIME_LITERAL() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	Typed_literal() ITyped_literalContext

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) INT_LITERAL() antlr.TerminalNode {
	return s.GetToken(STParserINT_LITERAL, 0)
}

func (s *LiteralContext) REAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(STParserREAL_LITERAL, 0)
}

func (s *LiteralContext) BOOL_LITERAL() antlr.TerminalNode {
	return s.GetToken(STParserBOOL_LITERAL, 0)
}

func (s *LiteralContext) TIME_LITERAL() antlr.TerminalNode {
	return s.GetToken(STParserTIME_LITERAL, 0)
}

func (s *LiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(STParserSTRING_LITERAL, 0)
}

func (s *LiteralContext) Typed_literal() ITyped_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITyped_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITyped_literalContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, STParserRULE_literal)
	p.SetState(478)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case STParserINT_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(472)
			p.Match(STParserINT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case STParserREAL_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(473)
			p.Match(STParserREAL_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case STParserBOOL_LITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(474)
			p.Match(STParserBOOL_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case STParserTIME_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(475)
			p.Match(STParserTIME_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case STParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(476)
			p.Match(STParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case STParserINT_TYPE_NAME, STParserREAL_TYPE_NAME, STParserBOOL_TYPE_NAME, STParserBIT_TYPE_NAME:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(477)
			p.Typed_literal()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITyped_literalContext is an interface to support dynamic dispatch.
type ITyped_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_name() IType_nameContext
	INT_LITERAL() antlr.TerminalNode
	REAL_LITERAL() antlr.TerminalNode
	BIT_STRING_LITERAL() antlr.TerminalNode

	// IsTyped_literalContext differentiates from other interfaces.
	IsTyped_literalContext()
}

type Typed_literalContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyTyped_literalContext() *Typed_literalContext {
	var p = new(Typed_literalContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_typed_literal
	return p
}

func InitEmptyTyped_literalContext(p *Typed_literalContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_typed_literal
}

func (*Typed_literalContext) IsTyped_literalContext() {}

func NewTyped_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Typed_literalContext {
	var p = new(Typed_literalContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_typed_literal

	return p
}

func (s *Typed_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Typed_literalContext) Type_name() IType_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_nameContext)
}

func (s *Typed_literalContext) INT_LITERAL() antlr.TerminalNode {
	return s.GetToken(STParserINT_LITERAL, 0)
}

func (s *Typed_literalContext) REAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(STParserREAL_LITERAL, 0)
}

func (s *Typed_literalContext) BIT_STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(STParserBIT_STRING_LITERAL, 0)
}

func (s *Typed_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Typed_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Typed_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitTyped_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Typed_literal() (localctx ITyped_literalContext) {
	localctx = NewTyped_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, STParserRULE_typed_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(480)
		p.Type_name()
	}
	{
		p.SetState(481)
		p.Match(STParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(482)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-81)) & ^0x3f) == 0 && ((int64(1)<<(_la-81))&1027) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IType_nameContext is an interface to support dynamic dispatch.
type IType_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_TYPE_NAME() antlr.TerminalNode
	REAL_TYPE_NAME() antlr.TerminalNode
	BIT_TYPE_NAME() antlr.TerminalNode
	BOOL_TYPE_NAME() antlr.TerminalNode

	// IsType_nameContext differentiates from other interfaces.
	IsType_nameContext()
}

type Type_nameContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyType_nameContext() *Type_nameContext {
	var p = new(Type_nameContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_type_name
	return p
}

func InitEmptyType_nameContext(p *Type_nameContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_type_name
}

func (*Type_nameContext) IsType_nameContext() {}

func NewType_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_nameContext {
	var p = new(Type_nameContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_type_name

	return p
}

func (s *Type_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_nameContext) INT_TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(STParserINT_TYPE_NAME, 0)
}

func (s *Type_nameContext) REAL_TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(STParserREAL_TYPE_NAME, 0)
}

func (s *Type_nameContext) BIT_TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(STParserBIT_TYPE_NAME, 0)
}

func (s *Type_nameContext) BOOL_TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(STParserBOOL_TYPE_NAME, 0)
}

func (s *Type_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitType_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Type_name() (localctx IType_nameContext) {
	localctx = NewType_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, STParserRULE_type_name)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(484)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&780) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	AllLBRACK() []antlr.TerminalNode
	LBRACK(i int) antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllRBRACK() []antlr.TerminalNode
	RBRACK(i int) antlr.TerminalNode
	DIRECT_VARIABLE() antlr.TerminalNode

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*CustomContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_variable
	return p
}

func InitEmptyVariableContext(p *VariableContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_variable
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) GetName() antlr.Token { return s.name }

func (s *VariableContext) SetName(v antlr.Token) { s.name = v }

func (s *VariableContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(STParserIDENTIFIER)
}

func (s *VariableContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(STParserIDENTIFIER, i)
}

func (s *VariableContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(STParserDOT)
}

func (s *VariableContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(STParserDOT, i)
}

func (s *VariableContext) AllLBRACK() []antlr.TerminalNode {
	return s.GetTokens(STParserLBRACK)
}

func (s *VariableContext) LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(STParserLBRACK, i)
}

func (s *VariableContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *VariableContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableContext) AllRBRACK() []antlr.TerminalNode {
	return s.GetTokens(STParserRBRACK)
}

func (s *VariableContext) RBRACK(i int) antlr.TerminalNode {
	return s.GetToken(STParserRBRACK, i)
}

func (s *VariableContext) DIRECT_VARIABLE() antlr.TerminalNode {
	return s.GetToken(STParserDIRECT_VARIABLE, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, STParserRULE_variable)
	var _alt int

	p.SetState(499)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case STParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(486)

			var _m = p.Match(STParserIDENTIFIER)

			localctx.(*VariableContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(495)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				p.SetState(493)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetTokenStream().LA(1) {
				case STParserDOT:
					{
						p.SetState(487)
						p.Match(STParserDOT)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(488)
						p.Match(STParserIDENTIFIER)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				case STParserLBRACK:
					{
						p.SetState(489)
						p.Match(STParserLBRACK)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(490)
						p.expression(0)
					}
					{
						p.SetState(491)
						p.Match(STParserRBRACK)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				default:
					p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					goto errorExit
				}

			}
			p.SetState(497)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case STParserDIRECT_VARIABLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(498)
			p.Match(STParserDIRECT_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInput_declContext is an interface to support dynamic dispatch.
type IInput_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_decl() IVar_declContext
	Edge_decl() IEdge_declContext
	Array_conform_decl() IArray_conform_declContext

	// IsInput_declContext differentiates from other interfaces.
	IsInput_declContext()
}

type Input_declContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyInput_declContext() *Input_declContext {
	var p = new(Input_declContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_input_decl
	return p
}

func InitEmptyInput_declContext(p *Input_declContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_input_decl
}

func (*Input_declContext) IsInput_declContext() {}

func NewInput_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Input_declContext {
	var p = new(Input_declContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_input_decl

	return p
}

func (s *Input_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Input_declContext) Var_decl() IVar_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_declContext)
}

func (s *Input_declContext) Edge_decl() IEdge_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEdge_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEdge_declContext)
}

func (s *Input_declContext) Array_conform_decl() IArray_conform_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_conform_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_conform_declContext)
}

func (s *Input_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Input_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Input_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitInput_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Input_decl() (localctx IInput_declContext) {
	localctx = NewInput_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, STParserRULE_input_decl)
	p.SetState(504)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(501)
			p.Var_decl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(502)
			p.Edge_decl()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(503)
			p.Array_conform_decl()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOutput_declContext is an interface to support dynamic dispatch.
type IOutput_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_decl() IVar_declContext
	Array_conform_decl() IArray_conform_declContext

	// IsOutput_declContext differentiates from other interfaces.
	IsOutput_declContext()
}

type Output_declContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyOutput_declContext() *Output_declContext {
	var p = new(Output_declContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_output_decl
	return p
}

func InitEmptyOutput_declContext(p *Output_declContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_output_decl
}

func (*Output_declContext) IsOutput_declContext() {}

func NewOutput_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Output_declContext {
	var p = new(Output_declContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_output_decl

	return p
}

func (s *Output_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Output_declContext) Var_decl() IVar_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_declContext)
}

func (s *Output_declContext) Array_conform_decl() IArray_conform_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_conform_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_conform_declContext)
}

func (s *Output_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Output_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Output_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitOutput_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Output_decl() (localctx IOutput_declContext) {
	localctx = NewOutput_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, STParserRULE_output_decl)
	p.SetState(508)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(506)
			p.Var_decl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(507)
			p.Array_conform_decl()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEdge_declContext is an interface to support dynamic dispatch.
type IEdge_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Variable_list() IVariable_listContext
	COLON() antlr.TerminalNode
	BOOL_TYPE_NAME() antlr.TerminalNode
	R_EDGE() antlr.TerminalNode
	F_EDGE() antlr.TerminalNode

	// IsEdge_declContext differentiates from other interfaces.
	IsEdge_declContext()
}

type Edge_declContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyEdge_declContext() *Edge_declContext {
	var p = new(Edge_declContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_edge_decl
	return p
}

func InitEmptyEdge_declContext(p *Edge_declContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_edge_decl
}

func (*Edge_declContext) IsEdge_declContext() {}

func NewEdge_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Edge_declContext {
	var p = new(Edge_declContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_edge_decl

	return p
}

func (s *Edge_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Edge_declContext) Variable_list() IVariable_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariable_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariable_listContext)
}

func (s *Edge_declContext) COLON() antlr.TerminalNode {
	return s.GetToken(STParserCOLON, 0)
}

func (s *Edge_declContext) BOOL_TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(STParserBOOL_TYPE_NAME, 0)
}

func (s *Edge_declContext) R_EDGE() antlr.TerminalNode {
	return s.GetToken(STParserR_EDGE, 0)
}

func (s *Edge_declContext) F_EDGE() antlr.TerminalNode {
	return s.GetToken(STParserF_EDGE, 0)
}

func (s *Edge_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Edge_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Edge_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitEdge_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Edge_decl() (localctx IEdge_declContext) {
	localctx = NewEdge_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, STParserRULE_edge_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(510)
		p.Variable_list()
	}
	{
		p.SetState(511)
		p.Match(STParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(512)
		p.Match(STParserBOOL_TYPE_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(513)
		_la = p.GetTokenStream().LA(1)

		if !(_la == STParserR_EDGE || _la == STParserF_EDGE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArray_conform_declContext is an interface to support dynamic dispatch.
type IArray_conform_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Variable_list() IVariable_listContext
	COLON() antlr.TerminalNode
	Array_conformand() IArray_conformandContext

	// IsArray_conform_declContext differentiates from other interfaces.
	IsArray_conform_declContext()
}

type Array_conform_declContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyArray_conform_declContext() *Array_conform_declContext {
	var p = new(Array_conform_declContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_array_conform_decl
	return p
}

func InitEmptyArray_conform_declContext(p *Array_conform_declContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_array_conform_decl
}

func (*Array_conform_declContext) IsArray_conform_declContext() {}

func NewArray_conform_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_conform_declContext {
	var p = new(Array_conform_declContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_array_conform_decl

	return p
}

func (s *Array_conform_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_conform_declContext) Variable_list() IVariable_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariable_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariable_listContext)
}

func (s *Array_conform_declContext) COLON() antlr.TerminalNode {
	return s.GetToken(STParserCOLON, 0)
}

func (s *Array_conform_declContext) Array_conformand() IArray_conformandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_conformandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_conformandContext)
}

func (s *Array_conform_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_conform_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_conform_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitArray_conform_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Array_conform_decl() (localctx IArray_conform_declContext) {
	localctx = NewArray_conform_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, STParserRULE_array_conform_decl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(515)
		p.Variable_list()
	}
	{
		p.SetState(516)
		p.Match(STParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(517)
		p.Array_conformand()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArray_conformandContext is an interface to support dynamic dispatch.
type IArray_conformandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ARRAY() antlr.TerminalNode
	LBRACK() antlr.TerminalNode
	AllMULT() []antlr.TerminalNode
	MULT(i int) antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	OF() antlr.TerminalNode
	Data_type_access() IData_type_accessContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArray_conformandContext differentiates from other interfaces.
	IsArray_conformandContext()
}

type Array_conformandContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyArray_conformandContext() *Array_conformandContext {
	var p = new(Array_conformandContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_array_conformand
	return p
}

func InitEmptyArray_conformandContext(p *Array_conformandContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_array_conformand
}

func (*Array_conformandContext) IsArray_conformandContext() {}

func NewArray_conformandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_conformandContext {
	var p = new(Array_conformandContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_array_conformand

	return p
}

func (s *Array_conformandContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_conformandContext) ARRAY() antlr.TerminalNode {
	return s.GetToken(STParserARRAY, 0)
}

func (s *Array_conformandContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(STParserLBRACK, 0)
}

func (s *Array_conformandContext) AllMULT() []antlr.TerminalNode {
	return s.GetTokens(STParserMULT)
}

func (s *Array_conformandContext) MULT(i int) antlr.TerminalNode {
	return s.GetToken(STParserMULT, i)
}

func (s *Array_conformandContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(STParserRBRACK, 0)
}

func (s *Array_conformandContext) OF() antlr.TerminalNode {
	return s.GetToken(STParserOF, 0)
}

func (s *Array_conformandContext) Data_type_access() IData_type_accessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IData_type_accessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IData_type_accessContext)
}

func (s *Array_conformandContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(STParserCOMMA)
}

func (s *Array_conformandContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(STParserCOMMA, i)
}

func (s *Array_conformandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_conformandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_conformandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitArray_conformand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Array_conformand() (localctx IArray_conformandContext) {
	localctx = NewArray_conformandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, STParserRULE_array_conformand)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(519)
		p.Match(STParserARRAY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(520)
		p.Match(STParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(521)
		p.Match(STParserMULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(526)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == STParserCOMMA {
		{
			p.SetState(522)
			p.Match(STParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(523)
			p.Match(STParserMULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(528)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(529)
		p.Match(STParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(530)
		p.Match(STParserOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(531)
		p.Data_type_access()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariable_listContext is an interface to support dynamic dispatch.
type IVariable_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsVariable_listContext differentiates from other interfaces.
	IsVariable_listContext()
}

type Variable_listContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyVariable_listContext() *Variable_listContext {
	var p = new(Variable_listContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_variable_list
	return p
}

func InitEmptyVariable_listContext(p *Variable_listContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_variable_list
}

func (*Variable_listContext) IsVariable_listContext() {}

func NewVariable_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Variable_listContext {
	var p = new(Variable_listContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_variable_list

	return p
}

func (s *Variable_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Variable_listContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(STParserIDENTIFIER)
}

func (s *Variable_listContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(STParserIDENTIFIER, i)
}

func (s *Variable_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(STParserCOMMA)
}

func (s *Variable_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(STParserCOMMA, i)
}

func (s *Variable_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Variable_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Variable_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitVariable_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Variable_list() (localctx IVariable_listContext) {
	localctx = NewVariable_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, STParserRULE_variable_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(533)
		p.Match(STParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(538)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == STParserCOMMA {
		{
			p.SetState(534)
			p.Match(STParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(535)
			p.Match(STParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(540)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IData_type_accessContext is an interface to support dynamic dispatch.
type IData_type_accessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Elementary_type_name() IElementary_type_nameContext
	Derived_type_access() IDerived_type_accessContext

	// IsData_type_accessContext differentiates from other interfaces.
	IsData_type_accessContext()
}

type Data_type_accessContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyData_type_accessContext() *Data_type_accessContext {
	var p = new(Data_type_accessContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_data_type_access
	return p
}

func InitEmptyData_type_accessContext(p *Data_type_accessContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_data_type_access
}

func (*Data_type_accessContext) IsData_type_accessContext() {}

func NewData_type_accessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Data_type_accessContext {
	var p = new(Data_type_accessContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_data_type_access

	return p
}

func (s *Data_type_accessContext) GetParser() antlr.Parser { return s.parser }

func (s *Data_type_accessContext) Elementary_type_name() IElementary_type_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementary_type_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElementary_type_nameContext)
}

func (s *Data_type_accessContext) Derived_type_access() IDerived_type_accessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDerived_type_accessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDerived_type_accessContext)
}

func (s *Data_type_accessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Data_type_accessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Data_type_accessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitData_type_access(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Data_type_access() (localctx IData_type_accessContext) {
	localctx = NewData_type_accessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, STParserRULE_data_type_access)
	p.SetState(543)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 55, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(541)
			p.Elementary_type_name()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(542)
			p.Derived_type_access()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDerived_type_accessContext is an interface to support dynamic dispatch.
type IDerived_type_accessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Single_elem_type_access() ISingle_elem_type_accessContext
	Array_type_access() IArray_type_accessContext
	Struct_type_access() IStruct_type_accessContext
	String_type_access() IString_type_accessContext

	// IsDerived_type_accessContext differentiates from other interfaces.
	IsDerived_type_accessContext()
}

type Derived_type_accessContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyDerived_type_accessContext() *Derived_type_accessContext {
	var p = new(Derived_type_accessContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_derived_type_access
	return p
}

func InitEmptyDerived_type_accessContext(p *Derived_type_accessContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_derived_type_access
}

func (*Derived_type_accessContext) IsDerived_type_accessContext() {}

func NewDerived_type_accessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Derived_type_accessContext {
	var p = new(Derived_type_accessContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_derived_type_access

	return p
}

func (s *Derived_type_accessContext) GetParser() antlr.Parser { return s.parser }

func (s *Derived_type_accessContext) Single_elem_type_access() ISingle_elem_type_accessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingle_elem_type_accessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingle_elem_type_accessContext)
}

func (s *Derived_type_accessContext) Array_type_access() IArray_type_accessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_type_accessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_type_accessContext)
}

func (s *Derived_type_accessContext) Struct_type_access() IStruct_type_accessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_type_accessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_type_accessContext)
}

func (s *Derived_type_accessContext) String_type_access() IString_type_accessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_type_accessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_type_accessContext)
}

func (s *Derived_type_accessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Derived_type_accessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Derived_type_accessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitDerived_type_access(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Derived_type_access() (localctx IDerived_type_accessContext) {
	localctx = NewDerived_type_accessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, STParserRULE_derived_type_access)
	p.SetState(549)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(545)
			p.Single_elem_type_access()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(546)
			p.Array_type_access()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(547)
			p.Struct_type_access()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(548)
			p.String_type_access()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingle_elem_type_accessContext is an interface to support dynamic dispatch.
type ISingle_elem_type_accessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Simple_type_access() ISimple_type_accessContext
	Subrange_type_access() ISubrange_type_accessContext
	Enum_type_access() IEnum_type_accessContext

	// IsSingle_elem_type_accessContext differentiates from other interfaces.
	IsSingle_elem_type_accessContext()
}

type Single_elem_type_accessContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptySingle_elem_type_accessContext() *Single_elem_type_accessContext {
	var p = new(Single_elem_type_accessContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_single_elem_type_access
	return p
}

func InitEmptySingle_elem_type_accessContext(p *Single_elem_type_accessContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_single_elem_type_access
}

func (*Single_elem_type_accessContext) IsSingle_elem_type_accessContext() {}

func NewSingle_elem_type_accessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Single_elem_type_accessContext {
	var p = new(Single_elem_type_accessContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_single_elem_type_access

	return p
}

func (s *Single_elem_type_accessContext) GetParser() antlr.Parser { return s.parser }

func (s *Single_elem_type_accessContext) Simple_type_access() ISimple_type_accessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_type_accessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_type_accessContext)
}

func (s *Single_elem_type_accessContext) Subrange_type_access() ISubrange_type_accessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubrange_type_accessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubrange_type_accessContext)
}

func (s *Single_elem_type_accessContext) Enum_type_access() IEnum_type_accessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnum_type_accessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnum_type_accessContext)
}

func (s *Single_elem_type_accessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Single_elem_type_accessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Single_elem_type_accessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitSingle_elem_type_access(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Single_elem_type_access() (localctx ISingle_elem_type_accessContext) {
	localctx = NewSingle_elem_type_accessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, STParserRULE_single_elem_type_access)
	p.SetState(554)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(551)
			p.Simple_type_access()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(552)
			p.Subrange_type_access()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(553)
			p.Enum_type_access()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimple_type_accessContext is an interface to support dynamic dispatch.
type ISimple_type_accessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Simple_type_name() ISimple_type_nameContext
	AllNamespace_name() []INamespace_nameContext
	Namespace_name(i int) INamespace_nameContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsSimple_type_accessContext differentiates from other interfaces.
	IsSimple_type_accessContext()
}

type Simple_type_accessContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptySimple_type_accessContext() *Simple_type_accessContext {
	var p = new(Simple_type_accessContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_simple_type_access
	return p
}

func InitEmptySimple_type_accessContext(p *Simple_type_accessContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_simple_type_access
}

func (*Simple_type_accessContext) IsSimple_type_accessContext() {}

func NewSimple_type_accessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_type_accessContext {
	var p = new(Simple_type_accessContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_simple_type_access

	return p
}

func (s *Simple_type_accessContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_type_accessContext) Simple_type_name() ISimple_type_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_type_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_type_nameContext)
}

func (s *Simple_type_accessContext) AllNamespace_name() []INamespace_nameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INamespace_nameContext); ok {
			len++
		}
	}

	tst := make([]INamespace_nameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INamespace_nameContext); ok {
			tst[i] = t.(INamespace_nameContext)
			i++
		}
	}

	return tst
}

func (s *Simple_type_accessContext) Namespace_name(i int) INamespace_nameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamespace_nameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamespace_nameContext)
}

func (s *Simple_type_accessContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(STParserDOT)
}

func (s *Simple_type_accessContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(STParserDOT, i)
}

func (s *Simple_type_accessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_type_accessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_type_accessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitSimple_type_access(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Simple_type_access() (localctx ISimple_type_accessContext) {
	localctx = NewSimple_type_accessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, STParserRULE_simple_type_access)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(561)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 58, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(556)
				p.Namespace_name()
			}
			{
				p.SetState(557)
				p.Match(STParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(563)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 58, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(564)
		p.Simple_type_name()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISubrange_type_accessContext is an interface to support dynamic dispatch.
type ISubrange_type_accessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Subrange_type_name() ISubrange_type_nameContext
	AllNamespace_name() []INamespace_nameContext
	Namespace_name(i int) INamespace_nameContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsSubrange_type_accessContext differentiates from other interfaces.
	IsSubrange_type_accessContext()
}

type Subrange_type_accessContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptySubrange_type_accessContext() *Subrange_type_accessContext {
	var p = new(Subrange_type_accessContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_subrange_type_access
	return p
}

func InitEmptySubrange_type_accessContext(p *Subrange_type_accessContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_subrange_type_access
}

func (*Subrange_type_accessContext) IsSubrange_type_accessContext() {}

func NewSubrange_type_accessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Subrange_type_accessContext {
	var p = new(Subrange_type_accessContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_subrange_type_access

	return p
}

func (s *Subrange_type_accessContext) GetParser() antlr.Parser { return s.parser }

func (s *Subrange_type_accessContext) Subrange_type_name() ISubrange_type_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubrange_type_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubrange_type_nameContext)
}

func (s *Subrange_type_accessContext) AllNamespace_name() []INamespace_nameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INamespace_nameContext); ok {
			len++
		}
	}

	tst := make([]INamespace_nameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INamespace_nameContext); ok {
			tst[i] = t.(INamespace_nameContext)
			i++
		}
	}

	return tst
}

func (s *Subrange_type_accessContext) Namespace_name(i int) INamespace_nameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamespace_nameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamespace_nameContext)
}

func (s *Subrange_type_accessContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(STParserDOT)
}

func (s *Subrange_type_accessContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(STParserDOT, i)
}

func (s *Subrange_type_accessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Subrange_type_accessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Subrange_type_accessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitSubrange_type_access(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Subrange_type_access() (localctx ISubrange_type_accessContext) {
	localctx = NewSubrange_type_accessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, STParserRULE_subrange_type_access)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(571)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(566)
				p.Namespace_name()
			}
			{
				p.SetState(567)
				p.Match(STParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(573)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(574)
		p.Subrange_type_name()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnum_type_accessContext is an interface to support dynamic dispatch.
type IEnum_type_accessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Enum_type_name() IEnum_type_nameContext
	AllNamespace_name() []INamespace_nameContext
	Namespace_name(i int) INamespace_nameContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsEnum_type_accessContext differentiates from other interfaces.
	IsEnum_type_accessContext()
}

type Enum_type_accessContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyEnum_type_accessContext() *Enum_type_accessContext {
	var p = new(Enum_type_accessContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_enum_type_access
	return p
}

func InitEmptyEnum_type_accessContext(p *Enum_type_accessContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_enum_type_access
}

func (*Enum_type_accessContext) IsEnum_type_accessContext() {}

func NewEnum_type_accessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Enum_type_accessContext {
	var p = new(Enum_type_accessContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_enum_type_access

	return p
}

func (s *Enum_type_accessContext) GetParser() antlr.Parser { return s.parser }

func (s *Enum_type_accessContext) Enum_type_name() IEnum_type_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnum_type_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnum_type_nameContext)
}

func (s *Enum_type_accessContext) AllNamespace_name() []INamespace_nameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INamespace_nameContext); ok {
			len++
		}
	}

	tst := make([]INamespace_nameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INamespace_nameContext); ok {
			tst[i] = t.(INamespace_nameContext)
			i++
		}
	}

	return tst
}

func (s *Enum_type_accessContext) Namespace_name(i int) INamespace_nameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamespace_nameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamespace_nameContext)
}

func (s *Enum_type_accessContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(STParserDOT)
}

func (s *Enum_type_accessContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(STParserDOT, i)
}

func (s *Enum_type_accessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Enum_type_accessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Enum_type_accessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitEnum_type_access(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Enum_type_access() (localctx IEnum_type_accessContext) {
	localctx = NewEnum_type_accessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, STParserRULE_enum_type_access)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(581)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(576)
				p.Namespace_name()
			}
			{
				p.SetState(577)
				p.Match(STParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(583)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(584)
		p.Enum_type_name()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArray_type_accessContext is an interface to support dynamic dispatch.
type IArray_type_accessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Array_type_name() IArray_type_nameContext
	AllNamespace_name() []INamespace_nameContext
	Namespace_name(i int) INamespace_nameContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsArray_type_accessContext differentiates from other interfaces.
	IsArray_type_accessContext()
}

type Array_type_accessContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyArray_type_accessContext() *Array_type_accessContext {
	var p = new(Array_type_accessContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_array_type_access
	return p
}

func InitEmptyArray_type_accessContext(p *Array_type_accessContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_array_type_access
}

func (*Array_type_accessContext) IsArray_type_accessContext() {}

func NewArray_type_accessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_type_accessContext {
	var p = new(Array_type_accessContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_array_type_access

	return p
}

func (s *Array_type_accessContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_type_accessContext) Array_type_name() IArray_type_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_type_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_type_nameContext)
}

func (s *Array_type_accessContext) AllNamespace_name() []INamespace_nameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INamespace_nameContext); ok {
			len++
		}
	}

	tst := make([]INamespace_nameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INamespace_nameContext); ok {
			tst[i] = t.(INamespace_nameContext)
			i++
		}
	}

	return tst
}

func (s *Array_type_accessContext) Namespace_name(i int) INamespace_nameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamespace_nameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamespace_nameContext)
}

func (s *Array_type_accessContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(STParserDOT)
}

func (s *Array_type_accessContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(STParserDOT, i)
}

func (s *Array_type_accessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_type_accessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_type_accessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitArray_type_access(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Array_type_access() (localctx IArray_type_accessContext) {
	localctx = NewArray_type_accessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, STParserRULE_array_type_access)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(591)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(586)
				p.Namespace_name()
			}
			{
				p.SetState(587)
				p.Match(STParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(593)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(594)
		p.Array_type_name()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStruct_type_accessContext is an interface to support dynamic dispatch.
type IStruct_type_accessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Struct_type_name() IStruct_type_nameContext
	AllNamespace_name() []INamespace_nameContext
	Namespace_name(i int) INamespace_nameContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsStruct_type_accessContext differentiates from other interfaces.
	IsStruct_type_accessContext()
}

type Struct_type_accessContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyStruct_type_accessContext() *Struct_type_accessContext {
	var p = new(Struct_type_accessContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_struct_type_access
	return p
}

func InitEmptyStruct_type_accessContext(p *Struct_type_accessContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_struct_type_access
}

func (*Struct_type_accessContext) IsStruct_type_accessContext() {}

func NewStruct_type_accessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_type_accessContext {
	var p = new(Struct_type_accessContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_struct_type_access

	return p
}

func (s *Struct_type_accessContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_type_accessContext) Struct_type_name() IStruct_type_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_type_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_type_nameContext)
}

func (s *Struct_type_accessContext) AllNamespace_name() []INamespace_nameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INamespace_nameContext); ok {
			len++
		}
	}

	tst := make([]INamespace_nameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INamespace_nameContext); ok {
			tst[i] = t.(INamespace_nameContext)
			i++
		}
	}

	return tst
}

func (s *Struct_type_accessContext) Namespace_name(i int) INamespace_nameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamespace_nameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamespace_nameContext)
}

func (s *Struct_type_accessContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(STParserDOT)
}

func (s *Struct_type_accessContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(STParserDOT, i)
}

func (s *Struct_type_accessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_type_accessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Struct_type_accessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitStruct_type_access(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Struct_type_access() (localctx IStruct_type_accessContext) {
	localctx = NewStruct_type_accessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, STParserRULE_struct_type_access)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(601)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 62, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(596)
				p.Namespace_name()
			}
			{
				p.SetState(597)
				p.Match(STParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(603)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 62, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(604)
		p.Struct_type_name()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IString_type_accessContext is an interface to support dynamic dispatch.
type IString_type_accessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_TYPE_NAME() antlr.TerminalNode
	AllNamespace_name() []INamespace_nameContext
	Namespace_name(i int) INamespace_nameContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsString_type_accessContext differentiates from other interfaces.
	IsString_type_accessContext()
}

type String_type_accessContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyString_type_accessContext() *String_type_accessContext {
	var p = new(String_type_accessContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_string_type_access
	return p
}

func InitEmptyString_type_accessContext(p *String_type_accessContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_string_type_access
}

func (*String_type_accessContext) IsString_type_accessContext() {}

func NewString_type_accessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_type_accessContext {
	var p = new(String_type_accessContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_string_type_access

	return p
}

func (s *String_type_accessContext) GetParser() antlr.Parser { return s.parser }

func (s *String_type_accessContext) STRING_TYPE_NAME() antlr.TerminalNode {
	return s.GetToken(STParserSTRING_TYPE_NAME, 0)
}

func (s *String_type_accessContext) AllNamespace_name() []INamespace_nameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INamespace_nameContext); ok {
			len++
		}
	}

	tst := make([]INamespace_nameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INamespace_nameContext); ok {
			tst[i] = t.(INamespace_nameContext)
			i++
		}
	}

	return tst
}

func (s *String_type_accessContext) Namespace_name(i int) INamespace_nameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamespace_nameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamespace_nameContext)
}

func (s *String_type_accessContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(STParserDOT)
}

func (s *String_type_accessContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(STParserDOT, i)
}

func (s *String_type_accessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_type_accessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_type_accessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitString_type_access(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) String_type_access() (localctx IString_type_accessContext) {
	localctx = NewString_type_accessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, STParserRULE_string_type_access)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(611)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == STParserIDENTIFIER {
		{
			p.SetState(606)
			p.Namespace_name()
		}
		{
			p.SetState(607)
			p.Match(STParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(613)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(614)
		p.Match(STParserSTRING_TYPE_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INamespace_nameContext is an interface to support dynamic dispatch.
type INamespace_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsNamespace_nameContext differentiates from other interfaces.
	IsNamespace_nameContext()
}

type Namespace_nameContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyNamespace_nameContext() *Namespace_nameContext {
	var p = new(Namespace_nameContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_namespace_name
	return p
}

func InitEmptyNamespace_nameContext(p *Namespace_nameContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_namespace_name
}

func (*Namespace_nameContext) IsNamespace_nameContext() {}

func NewNamespace_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Namespace_nameContext {
	var p = new(Namespace_nameContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_namespace_name

	return p
}

func (s *Namespace_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Namespace_nameContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(STParserIDENTIFIER)
}

func (s *Namespace_nameContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(STParserIDENTIFIER, i)
}

func (s *Namespace_nameContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(STParserDOT)
}

func (s *Namespace_nameContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(STParserDOT, i)
}

func (s *Namespace_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Namespace_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Namespace_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitNamespace_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Namespace_name() (localctx INamespace_nameContext) {
	localctx = NewNamespace_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, STParserRULE_namespace_name)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(616)
		p.Match(STParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(621)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 64, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(617)
				p.Match(STParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(618)
				p.Match(STParserIDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(623)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 64, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimple_type_nameContext is an interface to support dynamic dispatch.
type ISimple_type_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsSimple_type_nameContext differentiates from other interfaces.
	IsSimple_type_nameContext()
}

type Simple_type_nameContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptySimple_type_nameContext() *Simple_type_nameContext {
	var p = new(Simple_type_nameContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_simple_type_name
	return p
}

func InitEmptySimple_type_nameContext(p *Simple_type_nameContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_simple_type_name
}

func (*Simple_type_nameContext) IsSimple_type_nameContext() {}

func NewSimple_type_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_type_nameContext {
	var p = new(Simple_type_nameContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_simple_type_name

	return p
}

func (s *Simple_type_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_type_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(STParserIDENTIFIER, 0)
}

func (s *Simple_type_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_type_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_type_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitSimple_type_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Simple_type_name() (localctx ISimple_type_nameContext) {
	localctx = NewSimple_type_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, STParserRULE_simple_type_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(624)
		p.Match(STParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISubrange_type_nameContext is an interface to support dynamic dispatch.
type ISubrange_type_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsSubrange_type_nameContext differentiates from other interfaces.
	IsSubrange_type_nameContext()
}

type Subrange_type_nameContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptySubrange_type_nameContext() *Subrange_type_nameContext {
	var p = new(Subrange_type_nameContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_subrange_type_name
	return p
}

func InitEmptySubrange_type_nameContext(p *Subrange_type_nameContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_subrange_type_name
}

func (*Subrange_type_nameContext) IsSubrange_type_nameContext() {}

func NewSubrange_type_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Subrange_type_nameContext {
	var p = new(Subrange_type_nameContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_subrange_type_name

	return p
}

func (s *Subrange_type_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Subrange_type_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(STParserIDENTIFIER, 0)
}

func (s *Subrange_type_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Subrange_type_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Subrange_type_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitSubrange_type_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Subrange_type_name() (localctx ISubrange_type_nameContext) {
	localctx = NewSubrange_type_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, STParserRULE_subrange_type_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(626)
		p.Match(STParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnum_type_nameContext is an interface to support dynamic dispatch.
type IEnum_type_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsEnum_type_nameContext differentiates from other interfaces.
	IsEnum_type_nameContext()
}

type Enum_type_nameContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyEnum_type_nameContext() *Enum_type_nameContext {
	var p = new(Enum_type_nameContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_enum_type_name
	return p
}

func InitEmptyEnum_type_nameContext(p *Enum_type_nameContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_enum_type_name
}

func (*Enum_type_nameContext) IsEnum_type_nameContext() {}

func NewEnum_type_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Enum_type_nameContext {
	var p = new(Enum_type_nameContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_enum_type_name

	return p
}

func (s *Enum_type_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Enum_type_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(STParserIDENTIFIER, 0)
}

func (s *Enum_type_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Enum_type_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Enum_type_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitEnum_type_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Enum_type_name() (localctx IEnum_type_nameContext) {
	localctx = NewEnum_type_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, STParserRULE_enum_type_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(628)
		p.Match(STParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArray_type_nameContext is an interface to support dynamic dispatch.
type IArray_type_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsArray_type_nameContext differentiates from other interfaces.
	IsArray_type_nameContext()
}

type Array_type_nameContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyArray_type_nameContext() *Array_type_nameContext {
	var p = new(Array_type_nameContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_array_type_name
	return p
}

func InitEmptyArray_type_nameContext(p *Array_type_nameContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_array_type_name
}

func (*Array_type_nameContext) IsArray_type_nameContext() {}

func NewArray_type_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_type_nameContext {
	var p = new(Array_type_nameContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_array_type_name

	return p
}

func (s *Array_type_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_type_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(STParserIDENTIFIER, 0)
}

func (s *Array_type_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_type_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_type_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitArray_type_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Array_type_name() (localctx IArray_type_nameContext) {
	localctx = NewArray_type_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, STParserRULE_array_type_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(630)
		p.Match(STParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStruct_type_nameContext is an interface to support dynamic dispatch.
type IStruct_type_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsStruct_type_nameContext differentiates from other interfaces.
	IsStruct_type_nameContext()
}

type Struct_type_nameContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyStruct_type_nameContext() *Struct_type_nameContext {
	var p = new(Struct_type_nameContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_struct_type_name
	return p
}

func InitEmptyStruct_type_nameContext(p *Struct_type_nameContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_struct_type_name
}

func (*Struct_type_nameContext) IsStruct_type_nameContext() {}

func NewStruct_type_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_type_nameContext {
	var p = new(Struct_type_nameContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_struct_type_name

	return p
}

func (s *Struct_type_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_type_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(STParserIDENTIFIER, 0)
}

func (s *Struct_type_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_type_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Struct_type_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitStruct_type_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Struct_type_name() (localctx IStruct_type_nameContext) {
	localctx = NewStruct_type_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, STParserRULE_struct_type_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(632)
		p.Match(STParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *STParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 28:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *STParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
