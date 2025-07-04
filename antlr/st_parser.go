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
		"program", "function_decl", "function_block_decl", "var_declaration_block",
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
		4, 1, 93, 595, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0,
		106, 8, 0, 10, 0, 12, 0, 109, 9, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		1, 1, 3, 1, 118, 8, 1, 1, 1, 1, 1, 5, 1, 122, 8, 1, 10, 1, 12, 1, 125,
		9, 1, 1, 1, 3, 1, 128, 8, 1, 1, 1, 1, 1, 5, 1, 132, 8, 1, 10, 1, 12, 1,
		135, 9, 1, 1, 1, 3, 1, 138, 8, 1, 1, 1, 1, 1, 5, 1, 142, 8, 1, 10, 1, 12,
		1, 145, 9, 1, 1, 1, 3, 1, 148, 8, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 5, 2, 157, 8, 2, 10, 2, 12, 2, 160, 9, 2, 1, 2, 3, 2, 163, 8, 2,
		1, 2, 1, 2, 5, 2, 167, 8, 2, 10, 2, 12, 2, 170, 9, 2, 1, 2, 3, 2, 173,
		8, 2, 1, 2, 1, 2, 5, 2, 177, 8, 2, 10, 2, 12, 2, 180, 9, 2, 1, 2, 3, 2,
		183, 8, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 4, 3, 190, 8, 3, 11, 3, 12, 3,
		191, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 5, 4, 199, 8, 4, 10, 4, 12, 4, 202,
		9, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 208, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 3, 5, 216, 8, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		5, 7, 225, 8, 7, 10, 7, 12, 7, 228, 9, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 4, 9, 240, 8, 9, 11, 9, 12, 9, 241, 1, 9,
		1, 9, 1, 10, 4, 10, 247, 8, 10, 11, 10, 12, 10, 248, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 259, 8, 11, 1, 11, 1, 11, 3,
		11, 263, 8, 11, 1, 11, 1, 11, 3, 11, 267, 8, 11, 1, 11, 1, 11, 3, 11, 271,
		8, 11, 1, 11, 1, 11, 3, 11, 275, 8, 11, 1, 11, 1, 11, 3, 11, 279, 8, 11,
		1, 11, 1, 11, 3, 11, 283, 8, 11, 1, 11, 1, 11, 3, 11, 287, 8, 11, 3, 11,
		289, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 304, 8, 13, 10, 13, 12, 13, 307,
		9, 13, 1, 13, 1, 13, 3, 13, 311, 8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 4, 14, 319, 8, 14, 11, 14, 12, 14, 320, 1, 14, 1, 14, 3, 14,
		325, 8, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 5, 15, 332, 8, 15, 10, 15,
		12, 15, 335, 9, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 3, 16, 342, 8, 16,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 352, 8,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 5, 20, 375, 8, 20, 10, 20, 12, 20, 378, 9, 20, 3, 20, 380, 8, 20, 1,
		20, 1, 20, 1, 21, 1, 21, 3, 21, 386, 8, 21, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3,
		24, 402, 8, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 428, 8, 24, 10, 24, 12, 24,
		431, 9, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 439, 8, 25,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 5, 28, 454, 8, 28, 10, 28, 12, 28, 457, 9, 28, 1, 28,
		3, 28, 460, 8, 28, 1, 29, 1, 29, 1, 29, 3, 29, 465, 8, 29, 1, 30, 1, 30,
		3, 30, 469, 8, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 5, 33, 485, 8, 33, 10, 33,
		12, 33, 488, 9, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 5,
		34, 497, 8, 34, 10, 34, 12, 34, 500, 9, 34, 1, 35, 1, 35, 3, 35, 504, 8,
		35, 1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 510, 8, 36, 1, 37, 1, 37, 1, 37,
		3, 37, 515, 8, 37, 1, 38, 1, 38, 1, 38, 5, 38, 520, 8, 38, 10, 38, 12,
		38, 523, 9, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 5, 39, 530, 8, 39, 10,
		39, 12, 39, 533, 9, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 5, 40, 540,
		8, 40, 10, 40, 12, 40, 543, 9, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 5,
		41, 550, 8, 41, 10, 41, 12, 41, 553, 9, 41, 1, 41, 1, 41, 1, 42, 1, 42,
		1, 42, 5, 42, 560, 8, 42, 10, 42, 12, 42, 563, 9, 42, 1, 42, 1, 42, 1,
		43, 1, 43, 1, 43, 5, 43, 570, 8, 43, 10, 43, 12, 43, 573, 9, 43, 1, 43,
		1, 43, 1, 44, 1, 44, 1, 44, 5, 44, 580, 8, 44, 10, 44, 12, 44, 583, 9,
		44, 1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 49, 1, 49,
		1, 49, 0, 1, 48, 50, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62,
		64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98,
		0, 10, 1, 0, 11, 17, 2, 0, 2, 5, 8, 10, 2, 0, 44, 44, 65, 66, 1, 0, 67,
		69, 1, 0, 65, 66, 1, 0, 61, 64, 1, 0, 59, 60, 2, 0, 81, 82, 91, 91, 2,
		0, 2, 3, 8, 9, 1, 0, 92, 93, 634, 0, 100, 1, 0, 0, 0, 2, 113, 1, 0, 0,
		0, 4, 152, 1, 0, 0, 0, 6, 187, 1, 0, 0, 0, 8, 195, 1, 0, 0, 0, 10, 215,
		1, 0, 0, 0, 12, 217, 1, 0, 0, 0, 14, 219, 1, 0, 0, 0, 16, 233, 1, 0, 0,
		0, 18, 237, 1, 0, 0, 0, 20, 246, 1, 0, 0, 0, 22, 288, 1, 0, 0, 0, 24, 290,
		1, 0, 0, 0, 26, 294, 1, 0, 0, 0, 28, 314, 1, 0, 0, 0, 30, 328, 1, 0, 0,
		0, 32, 341, 1, 0, 0, 0, 34, 343, 1, 0, 0, 0, 36, 357, 1, 0, 0, 0, 38, 363,
		1, 0, 0, 0, 40, 369, 1, 0, 0, 0, 42, 383, 1, 0, 0, 0, 44, 387, 1, 0, 0,
		0, 46, 389, 1, 0, 0, 0, 48, 401, 1, 0, 0, 0, 50, 438, 1, 0, 0, 0, 52, 440,
		1, 0, 0, 0, 54, 444, 1, 0, 0, 0, 56, 459, 1, 0, 0, 0, 58, 464, 1, 0, 0,
		0, 60, 468, 1, 0, 0, 0, 62, 470, 1, 0, 0, 0, 64, 475, 1, 0, 0, 0, 66, 479,
		1, 0, 0, 0, 68, 493, 1, 0, 0, 0, 70, 503, 1, 0, 0, 0, 72, 509, 1, 0, 0,
		0, 74, 514, 1, 0, 0, 0, 76, 521, 1, 0, 0, 0, 78, 531, 1, 0, 0, 0, 80, 541,
		1, 0, 0, 0, 82, 551, 1, 0, 0, 0, 84, 561, 1, 0, 0, 0, 86, 571, 1, 0, 0,
		0, 88, 576, 1, 0, 0, 0, 90, 584, 1, 0, 0, 0, 92, 586, 1, 0, 0, 0, 94, 588,
		1, 0, 0, 0, 96, 590, 1, 0, 0, 0, 98, 592, 1, 0, 0, 0, 100, 101, 5, 23,
		0, 0, 101, 107, 5, 80, 0, 0, 102, 106, 3, 6, 3, 0, 103, 106, 3, 2, 1, 0,
		104, 106, 3, 4, 2, 0, 105, 102, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 105,
		104, 1, 0, 0, 0, 106, 109, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 108,
		1, 0, 0, 0, 108, 110, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 110, 111, 3, 20,
		10, 0, 111, 112, 5, 24, 0, 0, 112, 1, 1, 0, 0, 0, 113, 114, 5, 19, 0, 0,
		114, 117, 5, 80, 0, 0, 115, 116, 5, 73, 0, 0, 116, 118, 3, 10, 5, 0, 117,
		115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 127, 1, 0, 0, 0, 119, 123,
		5, 12, 0, 0, 120, 122, 3, 58, 29, 0, 121, 120, 1, 0, 0, 0, 122, 125, 1,
		0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 126, 1, 0, 0,
		0, 125, 123, 1, 0, 0, 0, 126, 128, 5, 18, 0, 0, 127, 119, 1, 0, 0, 0, 127,
		128, 1, 0, 0, 0, 128, 137, 1, 0, 0, 0, 129, 133, 5, 13, 0, 0, 130, 132,
		3, 60, 30, 0, 131, 130, 1, 0, 0, 0, 132, 135, 1, 0, 0, 0, 133, 131, 1,
		0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 136, 1, 0, 0, 0, 135, 133, 1, 0, 0,
		0, 136, 138, 5, 18, 0, 0, 137, 129, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138,
		147, 1, 0, 0, 0, 139, 143, 5, 11, 0, 0, 140, 142, 3, 8, 4, 0, 141, 140,
		1, 0, 0, 0, 142, 145, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 143, 144, 1, 0,
		0, 0, 144, 146, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 146, 148, 5, 18, 0, 0,
		147, 139, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149,
		150, 3, 20, 10, 0, 150, 151, 5, 20, 0, 0, 151, 3, 1, 0, 0, 0, 152, 153,
		5, 21, 0, 0, 153, 162, 5, 80, 0, 0, 154, 158, 5, 12, 0, 0, 155, 157, 3,
		58, 29, 0, 156, 155, 1, 0, 0, 0, 157, 160, 1, 0, 0, 0, 158, 156, 1, 0,
		0, 0, 158, 159, 1, 0, 0, 0, 159, 161, 1, 0, 0, 0, 160, 158, 1, 0, 0, 0,
		161, 163, 5, 18, 0, 0, 162, 154, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163,
		172, 1, 0, 0, 0, 164, 168, 5, 13, 0, 0, 165, 167, 3, 60, 30, 0, 166, 165,
		1, 0, 0, 0, 167, 170, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 168, 169, 1, 0,
		0, 0, 169, 171, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 171, 173, 5, 18, 0, 0,
		172, 164, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 182, 1, 0, 0, 0, 174,
		178, 5, 11, 0, 0, 175, 177, 3, 8, 4, 0, 176, 175, 1, 0, 0, 0, 177, 180,
		1, 0, 0, 0, 178, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 181, 1, 0,
		0, 0, 180, 178, 1, 0, 0, 0, 181, 183, 5, 18, 0, 0, 182, 174, 1, 0, 0, 0,
		182, 183, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 185, 3, 20, 10, 0, 185,
		186, 5, 22, 0, 0, 186, 5, 1, 0, 0, 0, 187, 189, 7, 0, 0, 0, 188, 190, 3,
		8, 4, 0, 189, 188, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 189, 1, 0, 0,
		0, 191, 192, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 194, 5, 18, 0, 0, 194,
		7, 1, 0, 0, 0, 195, 200, 5, 80, 0, 0, 196, 197, 5, 72, 0, 0, 197, 199,
		5, 80, 0, 0, 198, 196, 1, 0, 0, 0, 199, 202, 1, 0, 0, 0, 200, 198, 1, 0,
		0, 0, 200, 201, 1, 0, 0, 0, 201, 203, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0,
		203, 204, 5, 73, 0, 0, 204, 207, 3, 10, 5, 0, 205, 206, 5, 58, 0, 0, 206,
		208, 3, 48, 24, 0, 207, 205, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 209,
		1, 0, 0, 0, 209, 210, 5, 74, 0, 0, 210, 9, 1, 0, 0, 0, 211, 216, 3, 12,
		6, 0, 212, 216, 3, 14, 7, 0, 213, 216, 3, 18, 9, 0, 214, 216, 5, 80, 0,
		0, 215, 211, 1, 0, 0, 0, 215, 212, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 215,
		214, 1, 0, 0, 0, 216, 11, 1, 0, 0, 0, 217, 218, 7, 1, 0, 0, 218, 13, 1,
		0, 0, 0, 219, 220, 5, 48, 0, 0, 220, 221, 5, 77, 0, 0, 221, 226, 3, 16,
		8, 0, 222, 223, 5, 72, 0, 0, 223, 225, 3, 16, 8, 0, 224, 222, 1, 0, 0,
		0, 225, 228, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227,
		229, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 229, 230, 5, 78, 0, 0, 230, 231,
		5, 41, 0, 0, 231, 232, 3, 10, 5, 0, 232, 15, 1, 0, 0, 0, 233, 234, 3, 48,
		24, 0, 234, 235, 5, 79, 0, 0, 235, 236, 3, 48, 24, 0, 236, 17, 1, 0, 0,
		0, 237, 239, 5, 49, 0, 0, 238, 240, 3, 8, 4, 0, 239, 238, 1, 0, 0, 0, 240,
		241, 1, 0, 0, 0, 241, 239, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 243,
		1, 0, 0, 0, 243, 244, 5, 50, 0, 0, 244, 19, 1, 0, 0, 0, 245, 247, 3, 22,
		11, 0, 246, 245, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 246, 1, 0, 0, 0,
		248, 249, 1, 0, 0, 0, 249, 21, 1, 0, 0, 0, 250, 251, 3, 24, 12, 0, 251,
		252, 5, 74, 0, 0, 252, 289, 1, 0, 0, 0, 253, 254, 3, 40, 20, 0, 254, 255,
		5, 74, 0, 0, 255, 289, 1, 0, 0, 0, 256, 258, 3, 44, 22, 0, 257, 259, 5,
		74, 0, 0, 258, 257, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 289, 1, 0, 0,
		0, 260, 262, 3, 46, 23, 0, 261, 263, 5, 74, 0, 0, 262, 261, 1, 0, 0, 0,
		262, 263, 1, 0, 0, 0, 263, 289, 1, 0, 0, 0, 264, 266, 3, 42, 21, 0, 265,
		267, 5, 74, 0, 0, 266, 265, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 289,
		1, 0, 0, 0, 268, 270, 3, 26, 13, 0, 269, 271, 5, 74, 0, 0, 270, 269, 1,
		0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 289, 1, 0, 0, 0, 272, 274, 3, 28, 14,
		0, 273, 275, 5, 74, 0, 0, 274, 273, 1, 0, 0, 0, 274, 275, 1, 0, 0, 0, 275,
		289, 1, 0, 0, 0, 276, 278, 3, 34, 17, 0, 277, 279, 5, 74, 0, 0, 278, 277,
		1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 289, 1, 0, 0, 0, 280, 282, 3, 36,
		18, 0, 281, 283, 5, 74, 0, 0, 282, 281, 1, 0, 0, 0, 282, 283, 1, 0, 0,
		0, 283, 289, 1, 0, 0, 0, 284, 286, 3, 38, 19, 0, 285, 287, 5, 74, 0, 0,
		286, 285, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 289, 1, 0, 0, 0, 288,
		250, 1, 0, 0, 0, 288, 253, 1, 0, 0, 0, 288, 256, 1, 0, 0, 0, 288, 260,
		1, 0, 0, 0, 288, 264, 1, 0, 0, 0, 288, 268, 1, 0, 0, 0, 288, 272, 1, 0,
		0, 0, 288, 276, 1, 0, 0, 0, 288, 280, 1, 0, 0, 0, 288, 284, 1, 0, 0, 0,
		289, 23, 1, 0, 0, 0, 290, 291, 3, 56, 28, 0, 291, 292, 5, 58, 0, 0, 292,
		293, 3, 48, 24, 0, 293, 25, 1, 0, 0, 0, 294, 295, 5, 25, 0, 0, 295, 296,
		3, 48, 24, 0, 296, 297, 5, 26, 0, 0, 297, 305, 3, 20, 10, 0, 298, 299,
		5, 28, 0, 0, 299, 300, 3, 48, 24, 0, 300, 301, 5, 26, 0, 0, 301, 302, 3,
		20, 10, 0, 302, 304, 1, 0, 0, 0, 303, 298, 1, 0, 0, 0, 304, 307, 1, 0,
		0, 0, 305, 303, 1, 0, 0, 0, 305, 306, 1, 0, 0, 0, 306, 310, 1, 0, 0, 0,
		307, 305, 1, 0, 0, 0, 308, 309, 5, 27, 0, 0, 309, 311, 3, 20, 10, 0, 310,
		308, 1, 0, 0, 0, 310, 311, 1, 0, 0, 0, 311, 312, 1, 0, 0, 0, 312, 313,
		5, 29, 0, 0, 313, 27, 1, 0, 0, 0, 314, 315, 5, 40, 0, 0, 315, 316, 3, 48,
		24, 0, 316, 318, 5, 41, 0, 0, 317, 319, 3, 30, 15, 0, 318, 317, 1, 0, 0,
		0, 319, 320, 1, 0, 0, 0, 320, 318, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0, 321,
		324, 1, 0, 0, 0, 322, 323, 5, 27, 0, 0, 323, 325, 3, 20, 10, 0, 324, 322,
		1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 327, 5, 42,
		0, 0, 327, 29, 1, 0, 0, 0, 328, 333, 3, 32, 16, 0, 329, 330, 5, 72, 0,
		0, 330, 332, 3, 32, 16, 0, 331, 329, 1, 0, 0, 0, 332, 335, 1, 0, 0, 0,
		333, 331, 1, 0, 0, 0, 333, 334, 1, 0, 0, 0, 334, 336, 1, 0, 0, 0, 335,
		333, 1, 0, 0, 0, 336, 337, 5, 73, 0, 0, 337, 338, 3, 20, 10, 0, 338, 31,
		1, 0, 0, 0, 339, 342, 3, 48, 24, 0, 340, 342, 3, 16, 8, 0, 341, 339, 1,
		0, 0, 0, 341, 340, 1, 0, 0, 0, 342, 33, 1, 0, 0, 0, 343, 344, 5, 30, 0,
		0, 344, 345, 5, 80, 0, 0, 345, 346, 5, 58, 0, 0, 346, 347, 3, 48, 24, 0,
		347, 348, 5, 31, 0, 0, 348, 351, 3, 48, 24, 0, 349, 350, 5, 32, 0, 0, 350,
		352, 3, 48, 24, 0, 351, 349, 1, 0, 0, 0, 351, 352, 1, 0, 0, 0, 352, 353,
		1, 0, 0, 0, 353, 354, 5, 33, 0, 0, 354, 355, 3, 20, 10, 0, 355, 356, 5,
		34, 0, 0, 356, 35, 1, 0, 0, 0, 357, 358, 5, 35, 0, 0, 358, 359, 3, 48,
		24, 0, 359, 360, 5, 33, 0, 0, 360, 361, 3, 20, 10, 0, 361, 362, 5, 36,
		0, 0, 362, 37, 1, 0, 0, 0, 363, 364, 5, 37, 0, 0, 364, 365, 3, 20, 10,
		0, 365, 366, 5, 38, 0, 0, 366, 367, 3, 48, 24, 0, 367, 368, 5, 39, 0, 0,
		368, 39, 1, 0, 0, 0, 369, 370, 5, 80, 0, 0, 370, 379, 5, 75, 0, 0, 371,
		376, 3, 48, 24, 0, 372, 373, 5, 72, 0, 0, 373, 375, 3, 48, 24, 0, 374,
		372, 1, 0, 0, 0, 375, 378, 1, 0, 0, 0, 376, 374, 1, 0, 0, 0, 376, 377,
		1, 0, 0, 0, 377, 380, 1, 0, 0, 0, 378, 376, 1, 0, 0, 0, 379, 371, 1, 0,
		0, 0, 379, 380, 1, 0, 0, 0, 380, 381, 1, 0, 0, 0, 381, 382, 5, 76, 0, 0,
		382, 41, 1, 0, 0, 0, 383, 385, 5, 43, 0, 0, 384, 386, 3, 48, 24, 0, 385,
		384, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386, 43, 1, 0, 0, 0, 387, 388, 5,
		56, 0, 0, 388, 45, 1, 0, 0, 0, 389, 390, 5, 57, 0, 0, 390, 47, 1, 0, 0,
		0, 391, 392, 6, 24, -1, 0, 392, 393, 5, 75, 0, 0, 393, 394, 3, 48, 24,
		0, 394, 395, 5, 76, 0, 0, 395, 402, 1, 0, 0, 0, 396, 397, 7, 2, 0, 0, 397,
		402, 3, 48, 24, 12, 398, 402, 3, 50, 25, 0, 399, 402, 3, 56, 28, 0, 400,
		402, 3, 40, 20, 0, 401, 391, 1, 0, 0, 0, 401, 396, 1, 0, 0, 0, 401, 398,
		1, 0, 0, 0, 401, 399, 1, 0, 0, 0, 401, 400, 1, 0, 0, 0, 402, 429, 1, 0,
		0, 0, 403, 404, 10, 11, 0, 0, 404, 405, 5, 70, 0, 0, 405, 428, 3, 48, 24,
		12, 406, 407, 10, 10, 0, 0, 407, 408, 7, 3, 0, 0, 408, 428, 3, 48, 24,
		11, 409, 410, 10, 9, 0, 0, 410, 411, 7, 4, 0, 0, 411, 428, 3, 48, 24, 10,
		412, 413, 10, 8, 0, 0, 413, 414, 7, 5, 0, 0, 414, 428, 3, 48, 24, 9, 415,
		416, 10, 7, 0, 0, 416, 417, 7, 6, 0, 0, 417, 428, 3, 48, 24, 8, 418, 419,
		10, 6, 0, 0, 419, 420, 5, 45, 0, 0, 420, 428, 3, 48, 24, 7, 421, 422, 10,
		5, 0, 0, 422, 423, 5, 47, 0, 0, 423, 428, 3, 48, 24, 6, 424, 425, 10, 4,
		0, 0, 425, 426, 5, 46, 0, 0, 426, 428, 3, 48, 24, 5, 427, 403, 1, 0, 0,
		0, 427, 406, 1, 0, 0, 0, 427, 409, 1, 0, 0, 0, 427, 412, 1, 0, 0, 0, 427,
		415, 1, 0, 0, 0, 427, 418, 1, 0, 0, 0, 427, 421, 1, 0, 0, 0, 427, 424,
		1, 0, 0, 0, 428, 431, 1, 0, 0, 0, 429, 427, 1, 0, 0, 0, 429, 430, 1, 0,
		0, 0, 430, 49, 1, 0, 0, 0, 431, 429, 1, 0, 0, 0, 432, 439, 5, 81, 0, 0,
		433, 439, 5, 82, 0, 0, 434, 439, 5, 83, 0, 0, 435, 439, 5, 84, 0, 0, 436,
		439, 5, 85, 0, 0, 437, 439, 3, 52, 26, 0, 438, 432, 1, 0, 0, 0, 438, 433,
		1, 0, 0, 0, 438, 434, 1, 0, 0, 0, 438, 435, 1, 0, 0, 0, 438, 436, 1, 0,
		0, 0, 438, 437, 1, 0, 0, 0, 439, 51, 1, 0, 0, 0, 440, 441, 3, 54, 27, 0,
		441, 442, 5, 1, 0, 0, 442, 443, 7, 7, 0, 0, 443, 53, 1, 0, 0, 0, 444, 445,
		7, 8, 0, 0, 445, 55, 1, 0, 0, 0, 446, 455, 5, 80, 0, 0, 447, 448, 5, 71,
		0, 0, 448, 454, 5, 80, 0, 0, 449, 450, 5, 77, 0, 0, 450, 451, 3, 48, 24,
		0, 451, 452, 5, 78, 0, 0, 452, 454, 1, 0, 0, 0, 453, 447, 1, 0, 0, 0, 453,
		449, 1, 0, 0, 0, 454, 457, 1, 0, 0, 0, 455, 453, 1, 0, 0, 0, 455, 456,
		1, 0, 0, 0, 456, 460, 1, 0, 0, 0, 457, 455, 1, 0, 0, 0, 458, 460, 5, 86,
		0, 0, 459, 446, 1, 0, 0, 0, 459, 458, 1, 0, 0, 0, 460, 57, 1, 0, 0, 0,
		461, 465, 3, 8, 4, 0, 462, 465, 3, 62, 31, 0, 463, 465, 3, 64, 32, 0, 464,
		461, 1, 0, 0, 0, 464, 462, 1, 0, 0, 0, 464, 463, 1, 0, 0, 0, 465, 59, 1,
		0, 0, 0, 466, 469, 3, 8, 4, 0, 467, 469, 3, 64, 32, 0, 468, 466, 1, 0,
		0, 0, 468, 467, 1, 0, 0, 0, 469, 61, 1, 0, 0, 0, 470, 471, 3, 68, 34, 0,
		471, 472, 5, 73, 0, 0, 472, 473, 5, 8, 0, 0, 473, 474, 7, 9, 0, 0, 474,
		63, 1, 0, 0, 0, 475, 476, 3, 68, 34, 0, 476, 477, 5, 73, 0, 0, 477, 478,
		3, 66, 33, 0, 478, 65, 1, 0, 0, 0, 479, 480, 5, 48, 0, 0, 480, 481, 5,
		77, 0, 0, 481, 486, 5, 67, 0, 0, 482, 483, 5, 72, 0, 0, 483, 485, 5, 67,
		0, 0, 484, 482, 1, 0, 0, 0, 485, 488, 1, 0, 0, 0, 486, 484, 1, 0, 0, 0,
		486, 487, 1, 0, 0, 0, 487, 489, 1, 0, 0, 0, 488, 486, 1, 0, 0, 0, 489,
		490, 5, 78, 0, 0, 490, 491, 5, 41, 0, 0, 491, 492, 3, 70, 35, 0, 492, 67,
		1, 0, 0, 0, 493, 498, 5, 80, 0, 0, 494, 495, 5, 72, 0, 0, 495, 497, 5,
		80, 0, 0, 496, 494, 1, 0, 0, 0, 497, 500, 1, 0, 0, 0, 498, 496, 1, 0, 0,
		0, 498, 499, 1, 0, 0, 0, 499, 69, 1, 0, 0, 0, 500, 498, 1, 0, 0, 0, 501,
		504, 3, 12, 6, 0, 502, 504, 3, 72, 36, 0, 503, 501, 1, 0, 0, 0, 503, 502,
		1, 0, 0, 0, 504, 71, 1, 0, 0, 0, 505, 510, 3, 74, 37, 0, 506, 510, 3, 82,
		41, 0, 507, 510, 3, 84, 42, 0, 508, 510, 3, 86, 43, 0, 509, 505, 1, 0,
		0, 0, 509, 506, 1, 0, 0, 0, 509, 507, 1, 0, 0, 0, 509, 508, 1, 0, 0, 0,
		510, 73, 1, 0, 0, 0, 511, 515, 3, 76, 38, 0, 512, 515, 3, 78, 39, 0, 513,
		515, 3, 80, 40, 0, 514, 511, 1, 0, 0, 0, 514, 512, 1, 0, 0, 0, 514, 513,
		1, 0, 0, 0, 515, 75, 1, 0, 0, 0, 516, 517, 3, 88, 44, 0, 517, 518, 5, 71,
		0, 0, 518, 520, 1, 0, 0, 0, 519, 516, 1, 0, 0, 0, 520, 523, 1, 0, 0, 0,
		521, 519, 1, 0, 0, 0, 521, 522, 1, 0, 0, 0, 522, 524, 1, 0, 0, 0, 523,
		521, 1, 0, 0, 0, 524, 525, 3, 90, 45, 0, 525, 77, 1, 0, 0, 0, 526, 527,
		3, 88, 44, 0, 527, 528, 5, 71, 0, 0, 528, 530, 1, 0, 0, 0, 529, 526, 1,
		0, 0, 0, 530, 533, 1, 0, 0, 0, 531, 529, 1, 0, 0, 0, 531, 532, 1, 0, 0,
		0, 532, 534, 1, 0, 0, 0, 533, 531, 1, 0, 0, 0, 534, 535, 3, 92, 46, 0,
		535, 79, 1, 0, 0, 0, 536, 537, 3, 88, 44, 0, 537, 538, 5, 71, 0, 0, 538,
		540, 1, 0, 0, 0, 539, 536, 1, 0, 0, 0, 540, 543, 1, 0, 0, 0, 541, 539,
		1, 0, 0, 0, 541, 542, 1, 0, 0, 0, 542, 544, 1, 0, 0, 0, 543, 541, 1, 0,
		0, 0, 544, 545, 3, 94, 47, 0, 545, 81, 1, 0, 0, 0, 546, 547, 3, 88, 44,
		0, 547, 548, 5, 71, 0, 0, 548, 550, 1, 0, 0, 0, 549, 546, 1, 0, 0, 0, 550,
		553, 1, 0, 0, 0, 551, 549, 1, 0, 0, 0, 551, 552, 1, 0, 0, 0, 552, 554,
		1, 0, 0, 0, 553, 551, 1, 0, 0, 0, 554, 555, 3, 96, 48, 0, 555, 83, 1, 0,
		0, 0, 556, 557, 3, 88, 44, 0, 557, 558, 5, 71, 0, 0, 558, 560, 1, 0, 0,
		0, 559, 556, 1, 0, 0, 0, 560, 563, 1, 0, 0, 0, 561, 559, 1, 0, 0, 0, 561,
		562, 1, 0, 0, 0, 562, 564, 1, 0, 0, 0, 563, 561, 1, 0, 0, 0, 564, 565,
		3, 98, 49, 0, 565, 85, 1, 0, 0, 0, 566, 567, 3, 88, 44, 0, 567, 568, 5,
		71, 0, 0, 568, 570, 1, 0, 0, 0, 569, 566, 1, 0, 0, 0, 570, 573, 1, 0, 0,
		0, 571, 569, 1, 0, 0, 0, 571, 572, 1, 0, 0, 0, 572, 574, 1, 0, 0, 0, 573,
		571, 1, 0, 0, 0, 574, 575, 5, 10, 0, 0, 575, 87, 1, 0, 0, 0, 576, 581,
		5, 80, 0, 0, 577, 578, 5, 71, 0, 0, 578, 580, 5, 80, 0, 0, 579, 577, 1,
		0, 0, 0, 580, 583, 1, 0, 0, 0, 581, 579, 1, 0, 0, 0, 581, 582, 1, 0, 0,
		0, 582, 89, 1, 0, 0, 0, 583, 581, 1, 0, 0, 0, 584, 585, 5, 80, 0, 0, 585,
		91, 1, 0, 0, 0, 586, 587, 5, 80, 0, 0, 587, 93, 1, 0, 0, 0, 588, 589, 5,
		80, 0, 0, 589, 95, 1, 0, 0, 0, 590, 591, 5, 80, 0, 0, 591, 97, 1, 0, 0,
		0, 592, 593, 5, 80, 0, 0, 593, 99, 1, 0, 0, 0, 62, 105, 107, 117, 123,
		127, 133, 137, 143, 147, 158, 162, 168, 172, 178, 182, 191, 200, 207, 215,
		226, 241, 248, 258, 262, 266, 270, 274, 278, 282, 286, 288, 305, 310, 320,
		324, 333, 341, 351, 376, 379, 385, 401, 427, 429, 438, 453, 455, 459, 464,
		468, 486, 498, 503, 509, 514, 521, 531, 541, 551, 561, 571, 581,
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
	STParserRULE_program                 = 0
	STParserRULE_function_decl           = 1
	STParserRULE_function_block_decl     = 2
	STParserRULE_var_declaration_block   = 3
	STParserRULE_var_decl                = 4
	STParserRULE_data_type               = 5
	STParserRULE_elementary_type_name    = 6
	STParserRULE_array_type              = 7
	STParserRULE_subrange                = 8
	STParserRULE_structured_type         = 9
	STParserRULE_statement_list          = 10
	STParserRULE_statement               = 11
	STParserRULE_assignment_statement    = 12
	STParserRULE_if_statement            = 13
	STParserRULE_case_statement          = 14
	STParserRULE_case_element            = 15
	STParserRULE_case_label              = 16
	STParserRULE_for_statement           = 17
	STParserRULE_while_statement         = 18
	STParserRULE_repeat_statement        = 19
	STParserRULE_function_invocation     = 20
	STParserRULE_return_statement        = 21
	STParserRULE_continue_statement      = 22
	STParserRULE_exit_statement          = 23
	STParserRULE_expression              = 24
	STParserRULE_literal                 = 25
	STParserRULE_typed_literal           = 26
	STParserRULE_type_name               = 27
	STParserRULE_variable                = 28
	STParserRULE_input_decl              = 29
	STParserRULE_output_decl             = 30
	STParserRULE_edge_decl               = 31
	STParserRULE_array_conform_decl      = 32
	STParserRULE_array_conformand        = 33
	STParserRULE_variable_list           = 34
	STParserRULE_data_type_access        = 35
	STParserRULE_derived_type_access     = 36
	STParserRULE_single_elem_type_access = 37
	STParserRULE_simple_type_access      = 38
	STParserRULE_subrange_type_access    = 39
	STParserRULE_enum_type_access        = 40
	STParserRULE_array_type_access       = 41
	STParserRULE_struct_type_access      = 42
	STParserRULE_string_type_access      = 43
	STParserRULE_namespace_name          = 44
	STParserRULE_simple_type_name        = 45
	STParserRULE_subrange_type_name      = 46
	STParserRULE_enum_type_name          = 47
	STParserRULE_array_type_name         = 48
	STParserRULE_struct_type_name        = 49
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetVars returns the vars rule contexts.
	GetVars() IVar_declaration_blockContext

	// GetStmts returns the stmts rule contexts.
	GetStmts() IStatement_listContext

	// SetVars sets the vars rule contexts.
	SetVars(IVar_declaration_blockContext)

	// SetStmts sets the stmts rule contexts.
	SetStmts(IStatement_listContext)

	// Getter signatures
	PROGRAM() antlr.TerminalNode
	END_PROGRAM() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	Statement_list() IStatement_listContext
	AllFunction_decl() []IFunction_declContext
	Function_decl(i int) IFunction_declContext
	AllFunction_block_decl() []IFunction_block_declContext
	Function_block_decl(i int) IFunction_block_declContext
	AllVar_declaration_block() []IVar_declaration_blockContext
	Var_declaration_block(i int) IVar_declaration_blockContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*CustomContext
	parser antlr.Parser
	id     antlr.Token
	vars   IVar_declaration_blockContext
	stmts  IStatement_listContext
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

func (s *ProgramContext) GetVars() IVar_declaration_blockContext { return s.vars }

func (s *ProgramContext) GetStmts() IStatement_listContext { return s.stmts }

func (s *ProgramContext) SetVars(v IVar_declaration_blockContext) { s.vars = v }

func (s *ProgramContext) SetStmts(v IStatement_listContext) { s.stmts = v }

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

func (s *ProgramContext) AllFunction_decl() []IFunction_declContext {
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

func (s *ProgramContext) Function_decl(i int) IFunction_declContext {
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

func (s *ProgramContext) AllFunction_block_decl() []IFunction_block_declContext {
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

func (s *ProgramContext) Function_block_decl(i int) IFunction_block_declContext {
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
	p.EnterRule(localctx, 0, STParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(STParserPROGRAM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)

		var _m = p.Match(STParserIDENTIFIER)

		localctx.(*ProgramContext).id = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2881536) != 0 {
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case STParserVAR, STParserVAR_INPUT, STParserVAR_OUTPUT, STParserVAR_IN_OUT, STParserVAR_EXTERNAL, STParserVAR_GLOBAL, STParserVAR_TEMP:
			{
				p.SetState(102)

				var _x = p.Var_declaration_block()

				localctx.(*ProgramContext).vars = _x
			}

		case STParserFUNCTION:
			{
				p.SetState(103)
				p.Function_decl()
			}

		case STParserFUNCTION_BLOCK:
			{
				p.SetState(104)
				p.Function_block_decl()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(110)

		var _x = p.Statement_list()

		localctx.(*ProgramContext).stmts = _x
	}
	{
		p.SetState(111)
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
	p.EnterRule(localctx, 2, STParserRULE_function_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.Match(STParserFUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)
		p.Match(STParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserCOLON {
		{
			p.SetState(115)
			p.Match(STParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.Data_type()
		}

	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserVAR_INPUT {
		{
			p.SetState(119)
			p.Match(STParserVAR_INPUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == STParserIDENTIFIER {
			{
				p.SetState(120)
				p.Input_decl()
			}

			p.SetState(125)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(126)
			p.Match(STParserEND_VAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserVAR_OUTPUT {
		{
			p.SetState(129)
			p.Match(STParserVAR_OUTPUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == STParserIDENTIFIER {
			{
				p.SetState(130)
				p.Output_decl()
			}

			p.SetState(135)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(136)
			p.Match(STParserEND_VAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserVAR {
		{
			p.SetState(139)
			p.Match(STParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == STParserIDENTIFIER {
			{
				p.SetState(140)
				p.Var_decl()
			}

			p.SetState(145)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(146)
			p.Match(STParserEND_VAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(149)
		p.Statement_list()
	}
	{
		p.SetState(150)
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
	p.EnterRule(localctx, 4, STParserRULE_function_block_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Match(STParserFUNCTION_BLOCK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(STParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserVAR_INPUT {
		{
			p.SetState(154)
			p.Match(STParserVAR_INPUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == STParserIDENTIFIER {
			{
				p.SetState(155)
				p.Input_decl()
			}

			p.SetState(160)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(161)
			p.Match(STParserEND_VAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserVAR_OUTPUT {
		{
			p.SetState(164)
			p.Match(STParserVAR_OUTPUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == STParserIDENTIFIER {
			{
				p.SetState(165)
				p.Output_decl()
			}

			p.SetState(170)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(171)
			p.Match(STParserEND_VAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserVAR {
		{
			p.SetState(174)
			p.Match(STParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == STParserIDENTIFIER {
			{
				p.SetState(175)
				p.Var_decl()
			}

			p.SetState(180)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(181)
			p.Match(STParserEND_VAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(184)
		p.Statement_list()
	}
	{
		p.SetState(185)
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
	p.EnterRule(localctx, 6, STParserRULE_var_declaration_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&260096) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == STParserIDENTIFIER {
		{
			p.SetState(188)

			var _x = p.Var_decl()

			localctx.(*Var_declaration_blockContext)._var_decl = _x
		}
		localctx.(*Var_declaration_blockContext).decls = append(localctx.(*Var_declaration_blockContext).decls, localctx.(*Var_declaration_blockContext)._var_decl)

		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(193)
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
	p.EnterRule(localctx, 8, STParserRULE_var_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)

		var _m = p.Match(STParserIDENTIFIER)

		localctx.(*Var_declContext).id = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == STParserCOMMA {
		{
			p.SetState(196)
			p.Match(STParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.Match(STParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(203)
		p.Match(STParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(204)

		var _x = p.Data_type()

		localctx.(*Var_declContext).type_ = _x
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserASSIGN {
		{
			p.SetState(205)
			p.Match(STParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(206)

			var _x = p.expression(0)

			localctx.(*Var_declContext).expr = _x
		}

	}
	{
		p.SetState(209)
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
	p.EnterRule(localctx, 10, STParserRULE_data_type)
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case STParserINT_TYPE_NAME, STParserREAL_TYPE_NAME, STParserTIME_TYPE_NAME, STParserDATE_TYPE_NAME, STParserBOOL_TYPE_NAME, STParserBIT_TYPE_NAME, STParserSTRING_TYPE_NAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)
			p.Elementary_type_name()
		}

	case STParserARRAY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(212)
			p.Array_type()
		}

	case STParserSTRUCT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(213)
			p.Structured_type()
		}

	case STParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(214)
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
	p.EnterRule(localctx, 12, STParserRULE_elementary_type_name)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)
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
	p.EnterRule(localctx, 14, STParserRULE_array_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(219)
		p.Match(STParserARRAY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(220)
		p.Match(STParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(221)
		p.Subrange()
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == STParserCOMMA {
		{
			p.SetState(222)
			p.Match(STParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
			p.Subrange()
		}

		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(229)
		p.Match(STParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
		p.Match(STParserOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(231)
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
	p.EnterRule(localctx, 16, STParserRULE_subrange)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		p.expression(0)
	}
	{
		p.SetState(234)
		p.Match(STParserDOTDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(235)
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
	p.EnterRule(localctx, 18, STParserRULE_structured_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.Match(STParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == STParserIDENTIFIER {
		{
			p.SetState(238)
			p.Var_decl()
		}

		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(243)
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
	p.EnterRule(localctx, 20, STParserRULE_statement_list)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(245)
				p.Statement()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 22, STParserRULE_statement)
	var _la int

	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(250)
			p.Assignment_statement()
		}
		{
			p.SetState(251)
			p.Match(STParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(253)
			p.Function_invocation()
		}
		{
			p.SetState(254)
			p.Match(STParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(256)
			p.Continue_statement()
		}
		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == STParserSEMICOLON {
			{
				p.SetState(257)
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
			p.SetState(260)
			p.Exit_statement()
		}
		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == STParserSEMICOLON {
			{
				p.SetState(261)
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
			p.SetState(264)
			p.Return_statement()
		}
		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == STParserSEMICOLON {
			{
				p.SetState(265)
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
			p.SetState(268)
			p.If_statement()
		}
		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == STParserSEMICOLON {
			{
				p.SetState(269)
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
			p.SetState(272)
			p.Case_statement()
		}
		p.SetState(274)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == STParserSEMICOLON {
			{
				p.SetState(273)
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
			p.SetState(276)
			p.For_statement()
		}
		p.SetState(278)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == STParserSEMICOLON {
			{
				p.SetState(277)
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
			p.SetState(280)
			p.While_statement()
		}
		p.SetState(282)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == STParserSEMICOLON {
			{
				p.SetState(281)
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
			p.SetState(284)
			p.Repeat_statement()
		}
		p.SetState(286)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == STParserSEMICOLON {
			{
				p.SetState(285)
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
	p.EnterRule(localctx, 24, STParserRULE_assignment_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.Variable()
	}
	{
		p.SetState(291)
		p.Match(STParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(292)
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
	p.EnterRule(localctx, 26, STParserRULE_if_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(294)
		p.Match(STParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(295)

		var _x = p.expression(0)

		localctx.(*If_statementContext)._expression = _x
	}
	localctx.(*If_statementContext).conds = append(localctx.(*If_statementContext).conds, localctx.(*If_statementContext)._expression)
	{
		p.SetState(296)
		p.Match(STParserTHEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(297)

		var _x = p.Statement_list()

		localctx.(*If_statementContext)._statement_list = _x
	}
	localctx.(*If_statementContext).thens = append(localctx.(*If_statementContext).thens, localctx.(*If_statementContext)._statement_list)
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == STParserELSIF {
		{
			p.SetState(298)
			p.Match(STParserELSIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)

			var _x = p.expression(0)

			localctx.(*If_statementContext)._expression = _x
		}
		localctx.(*If_statementContext).conds = append(localctx.(*If_statementContext).conds, localctx.(*If_statementContext)._expression)
		{
			p.SetState(300)
			p.Match(STParserTHEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(301)

			var _x = p.Statement_list()

			localctx.(*If_statementContext)._statement_list = _x
		}
		localctx.(*If_statementContext).thens = append(localctx.(*If_statementContext).thens, localctx.(*If_statementContext)._statement_list)

		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserELSE {
		{
			p.SetState(308)
			p.Match(STParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(309)

			var _x = p.Statement_list()

			localctx.(*If_statementContext).else_ = _x
		}

	}
	{
		p.SetState(312)
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
	p.EnterRule(localctx, 28, STParserRULE_case_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.Match(STParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(315)
		p.expression(0)
	}
	{
		p.SetState(316)
		p.Match(STParserOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17592186045196) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&4162563) != 0) {
		{
			p.SetState(317)
			p.Case_element()
		}

		p.SetState(320)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserELSE {
		{
			p.SetState(322)
			p.Match(STParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)
			p.Statement_list()
		}

	}
	{
		p.SetState(326)
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
	p.EnterRule(localctx, 30, STParserRULE_case_element)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)
		p.Case_label()
	}
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == STParserCOMMA {
		{
			p.SetState(329)
			p.Match(STParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(330)
			p.Case_label()
		}

		p.SetState(335)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(336)
		p.Match(STParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(337)
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
	p.EnterRule(localctx, 32, STParserRULE_case_label)
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(339)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(340)
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
	p.EnterRule(localctx, 34, STParserRULE_for_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(343)
		p.Match(STParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(344)
		p.Match(STParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(345)
		p.Match(STParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(346)
		p.expression(0)
	}
	{
		p.SetState(347)
		p.Match(STParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(348)
		p.expression(0)
	}
	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserBY {
		{
			p.SetState(349)
			p.Match(STParserBY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(350)
			p.expression(0)
		}

	}
	{
		p.SetState(353)
		p.Match(STParserDO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(354)
		p.Statement_list()
	}
	{
		p.SetState(355)
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
	p.EnterRule(localctx, 36, STParserRULE_while_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.Match(STParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(358)
		p.expression(0)
	}
	{
		p.SetState(359)
		p.Match(STParserDO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(360)
		p.Statement_list()
	}
	{
		p.SetState(361)
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
	p.EnterRule(localctx, 38, STParserRULE_repeat_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.Match(STParserREPEAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(364)
		p.Statement_list()
	}
	{
		p.SetState(365)
		p.Match(STParserUNTIL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(366)
		p.expression(0)
	}
	{
		p.SetState(367)
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
	p.EnterRule(localctx, 40, STParserRULE_function_invocation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(369)
		p.Match(STParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(370)
		p.Match(STParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17592186045196) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&4162563) != 0) {
		{
			p.SetState(371)

			var _x = p.expression(0)

			localctx.(*Function_invocationContext)._expression = _x
		}
		localctx.(*Function_invocationContext).args = append(localctx.(*Function_invocationContext).args, localctx.(*Function_invocationContext)._expression)
		p.SetState(376)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == STParserCOMMA {
			{
				p.SetState(372)
				p.Match(STParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(373)

				var _x = p.expression(0)

				localctx.(*Function_invocationContext)._expression = _x
			}
			localctx.(*Function_invocationContext).args = append(localctx.(*Function_invocationContext).args, localctx.(*Function_invocationContext)._expression)

			p.SetState(378)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(381)
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
	p.EnterRule(localctx, 42, STParserRULE_return_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(383)
		p.Match(STParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(385)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(384)
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
	p.EnterRule(localctx, 44, STParserRULE_continue_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(387)
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
	p.EnterRule(localctx, 46, STParserRULE_exit_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(389)
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
	_startState := 48
	p.EnterRecursionRule(localctx, 48, STParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(392)
			p.Match(STParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(393)
			p.expression(0)
		}
		{
			p.SetState(394)
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
			p.SetState(396)

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
			p.SetState(397)
			p.expression(12)
		}

	case 3:
		localctx = NewLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(398)
			p.Literal()
		}

	case 4:
		localctx = NewVarExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(399)
			p.Variable()
		}

	case 5:
		localctx = NewFuncCallExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(400)
			p.Function_invocation()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(429)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(427)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, STParserRULE_expression)
				p.SetState(403)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(404)

					var _m = p.Match(STParserPOWER)

					localctx.(*BinaryExpressionContext).operator = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(405)

					var _x = p.expression(12)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case 2:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, STParserRULE_expression)
				p.SetState(406)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(407)

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
					p.SetState(408)

					var _x = p.expression(11)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case 3:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, STParserRULE_expression)
				p.SetState(409)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(410)

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
					p.SetState(411)

					var _x = p.expression(10)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case 4:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, STParserRULE_expression)
				p.SetState(412)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(413)

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
					p.SetState(414)

					var _x = p.expression(9)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case 5:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, STParserRULE_expression)
				p.SetState(415)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(416)

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
					p.SetState(417)

					var _x = p.expression(8)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case 6:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, STParserRULE_expression)
				p.SetState(418)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(419)

					var _m = p.Match(STParserAND)

					localctx.(*BinaryExpressionContext).operator = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(420)

					var _x = p.expression(7)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case 7:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, STParserRULE_expression)
				p.SetState(421)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(422)

					var _m = p.Match(STParserXOR)

					localctx.(*BinaryExpressionContext).operator = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(423)

					var _x = p.expression(6)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case 8:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, STParserRULE_expression)
				p.SetState(424)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(425)

					var _m = p.Match(STParserOR)

					localctx.(*BinaryExpressionContext).operator = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(426)

					var _x = p.expression(5)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(431)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 50, STParserRULE_literal)
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case STParserINT_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(432)
			p.Match(STParserINT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case STParserREAL_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(433)
			p.Match(STParserREAL_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case STParserBOOL_LITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(434)
			p.Match(STParserBOOL_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case STParserTIME_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(435)
			p.Match(STParserTIME_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case STParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(436)
			p.Match(STParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case STParserINT_TYPE_NAME, STParserREAL_TYPE_NAME, STParserBOOL_TYPE_NAME, STParserBIT_TYPE_NAME:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(437)
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
	p.EnterRule(localctx, 52, STParserRULE_typed_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(440)
		p.Type_name()
	}
	{
		p.SetState(441)
		p.Match(STParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(442)
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
	p.EnterRule(localctx, 54, STParserRULE_type_name)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(444)
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
	p.EnterRule(localctx, 56, STParserRULE_variable)
	var _alt int

	p.SetState(459)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case STParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(446)

			var _m = p.Match(STParserIDENTIFIER)

			localctx.(*VariableContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(455)
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
				p.SetState(453)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetTokenStream().LA(1) {
				case STParserDOT:
					{
						p.SetState(447)
						p.Match(STParserDOT)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(448)
						p.Match(STParserIDENTIFIER)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				case STParserLBRACK:
					{
						p.SetState(449)
						p.Match(STParserLBRACK)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(450)
						p.expression(0)
					}
					{
						p.SetState(451)
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
			p.SetState(457)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case STParserDIRECT_VARIABLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(458)
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
	p.EnterRule(localctx, 58, STParserRULE_input_decl)
	p.SetState(464)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(461)
			p.Var_decl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(462)
			p.Edge_decl()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(463)
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
	p.EnterRule(localctx, 60, STParserRULE_output_decl)
	p.SetState(468)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(466)
			p.Var_decl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(467)
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
	p.EnterRule(localctx, 62, STParserRULE_edge_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(470)
		p.Variable_list()
	}
	{
		p.SetState(471)
		p.Match(STParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(472)
		p.Match(STParserBOOL_TYPE_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(473)
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
	p.EnterRule(localctx, 64, STParserRULE_array_conform_decl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(475)
		p.Variable_list()
	}
	{
		p.SetState(476)
		p.Match(STParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(477)
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
	p.EnterRule(localctx, 66, STParserRULE_array_conformand)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(479)
		p.Match(STParserARRAY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(480)
		p.Match(STParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(481)
		p.Match(STParserMULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(486)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == STParserCOMMA {
		{
			p.SetState(482)
			p.Match(STParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(483)
			p.Match(STParserMULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(488)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(489)
		p.Match(STParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(490)
		p.Match(STParserOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(491)
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
	p.EnterRule(localctx, 68, STParserRULE_variable_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(493)
		p.Match(STParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(498)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == STParserCOMMA {
		{
			p.SetState(494)
			p.Match(STParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(495)
			p.Match(STParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(500)
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
	p.EnterRule(localctx, 70, STParserRULE_data_type_access)
	p.SetState(503)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(501)
			p.Elementary_type_name()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(502)
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
	p.EnterRule(localctx, 72, STParserRULE_derived_type_access)
	p.SetState(509)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(505)
			p.Single_elem_type_access()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(506)
			p.Array_type_access()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(507)
			p.Struct_type_access()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(508)
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
	p.EnterRule(localctx, 74, STParserRULE_single_elem_type_access)
	p.SetState(514)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 54, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(511)
			p.Simple_type_access()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(512)
			p.Subrange_type_access()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(513)
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
	p.EnterRule(localctx, 76, STParserRULE_simple_type_access)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(521)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 55, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(516)
				p.Namespace_name()
			}
			{
				p.SetState(517)
				p.Match(STParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(523)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 55, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(524)
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
	p.EnterRule(localctx, 78, STParserRULE_subrange_type_access)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(531)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(526)
				p.Namespace_name()
			}
			{
				p.SetState(527)
				p.Match(STParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(533)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(534)
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
	p.EnterRule(localctx, 80, STParserRULE_enum_type_access)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(541)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(536)
				p.Namespace_name()
			}
			{
				p.SetState(537)
				p.Match(STParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(543)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(544)
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
	p.EnterRule(localctx, 82, STParserRULE_array_type_access)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(551)
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
				p.SetState(546)
				p.Namespace_name()
			}
			{
				p.SetState(547)
				p.Match(STParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(553)
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
		p.SetState(554)
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
	p.EnterRule(localctx, 84, STParserRULE_struct_type_access)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(561)
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
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(564)
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
	p.EnterRule(localctx, 86, STParserRULE_string_type_access)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(571)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == STParserIDENTIFIER {
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

		p.SetState(573)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(574)
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
	p.EnterRule(localctx, 88, STParserRULE_namespace_name)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(576)
		p.Match(STParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(581)
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
				p.SetState(577)
				p.Match(STParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(578)
				p.Match(STParserIDENTIFIER)
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
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 90, STParserRULE_simple_type_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(584)
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
	p.EnterRule(localctx, 92, STParserRULE_subrange_type_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(586)
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
	p.EnterRule(localctx, 94, STParserRULE_enum_type_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(588)
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
	p.EnterRule(localctx, 96, STParserRULE_array_type_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(590)
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
	p.EnterRule(localctx, 98, STParserRULE_struct_type_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(592)
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
	case 24:
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
