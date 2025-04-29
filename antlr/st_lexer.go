// Code generated from ./antlr/ST.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type STLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var STLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func stlexerLexerInit() {
	staticData := &STLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'PROGRAM'", "'END_PROGRAM'", "'VAR'", "'END_VAR'", "':'", "';'",
		"'INT'", "':='", "'IF'", "'THEN'", "'ELSEIF'", "'ELSE'", "'END_IF'",
		"'('", "')'", "'*'", "'/'", "'MOD'", "'+'", "'-'", "'>'", "'>='", "'<'",
		"'<='", "'='", "'<>'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "Integer", "ID", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "Integer", "ID", "WHITESPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 29, 181, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 5, 26, 163, 8,
		26, 10, 26, 12, 26, 166, 9, 26, 1, 27, 1, 27, 5, 27, 170, 8, 27, 10, 27,
		12, 27, 173, 9, 27, 1, 28, 4, 28, 176, 8, 28, 11, 28, 12, 28, 177, 1, 28,
		1, 28, 0, 0, 29, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17,
		9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35,
		18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53,
		27, 55, 28, 57, 29, 1, 0, 3, 3, 0, 65, 90, 95, 95, 97, 122, 5, 0, 36, 36,
		48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 183, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0,
		0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0,
		0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0,
		0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1,
		0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55,
		1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 1, 59, 1, 0, 0, 0, 3, 67, 1, 0, 0, 0, 5,
		79, 1, 0, 0, 0, 7, 83, 1, 0, 0, 0, 9, 91, 1, 0, 0, 0, 11, 93, 1, 0, 0,
		0, 13, 95, 1, 0, 0, 0, 15, 99, 1, 0, 0, 0, 17, 102, 1, 0, 0, 0, 19, 105,
		1, 0, 0, 0, 21, 110, 1, 0, 0, 0, 23, 117, 1, 0, 0, 0, 25, 122, 1, 0, 0,
		0, 27, 129, 1, 0, 0, 0, 29, 131, 1, 0, 0, 0, 31, 133, 1, 0, 0, 0, 33, 135,
		1, 0, 0, 0, 35, 137, 1, 0, 0, 0, 37, 141, 1, 0, 0, 0, 39, 143, 1, 0, 0,
		0, 41, 145, 1, 0, 0, 0, 43, 147, 1, 0, 0, 0, 45, 150, 1, 0, 0, 0, 47, 152,
		1, 0, 0, 0, 49, 155, 1, 0, 0, 0, 51, 157, 1, 0, 0, 0, 53, 160, 1, 0, 0,
		0, 55, 167, 1, 0, 0, 0, 57, 175, 1, 0, 0, 0, 59, 60, 5, 80, 0, 0, 60, 61,
		5, 82, 0, 0, 61, 62, 5, 79, 0, 0, 62, 63, 5, 71, 0, 0, 63, 64, 5, 82, 0,
		0, 64, 65, 5, 65, 0, 0, 65, 66, 5, 77, 0, 0, 66, 2, 1, 0, 0, 0, 67, 68,
		5, 69, 0, 0, 68, 69, 5, 78, 0, 0, 69, 70, 5, 68, 0, 0, 70, 71, 5, 95, 0,
		0, 71, 72, 5, 80, 0, 0, 72, 73, 5, 82, 0, 0, 73, 74, 5, 79, 0, 0, 74, 75,
		5, 71, 0, 0, 75, 76, 5, 82, 0, 0, 76, 77, 5, 65, 0, 0, 77, 78, 5, 77, 0,
		0, 78, 4, 1, 0, 0, 0, 79, 80, 5, 86, 0, 0, 80, 81, 5, 65, 0, 0, 81, 82,
		5, 82, 0, 0, 82, 6, 1, 0, 0, 0, 83, 84, 5, 69, 0, 0, 84, 85, 5, 78, 0,
		0, 85, 86, 5, 68, 0, 0, 86, 87, 5, 95, 0, 0, 87, 88, 5, 86, 0, 0, 88, 89,
		5, 65, 0, 0, 89, 90, 5, 82, 0, 0, 90, 8, 1, 0, 0, 0, 91, 92, 5, 58, 0,
		0, 92, 10, 1, 0, 0, 0, 93, 94, 5, 59, 0, 0, 94, 12, 1, 0, 0, 0, 95, 96,
		5, 73, 0, 0, 96, 97, 5, 78, 0, 0, 97, 98, 5, 84, 0, 0, 98, 14, 1, 0, 0,
		0, 99, 100, 5, 58, 0, 0, 100, 101, 5, 61, 0, 0, 101, 16, 1, 0, 0, 0, 102,
		103, 5, 73, 0, 0, 103, 104, 5, 70, 0, 0, 104, 18, 1, 0, 0, 0, 105, 106,
		5, 84, 0, 0, 106, 107, 5, 72, 0, 0, 107, 108, 5, 69, 0, 0, 108, 109, 5,
		78, 0, 0, 109, 20, 1, 0, 0, 0, 110, 111, 5, 69, 0, 0, 111, 112, 5, 76,
		0, 0, 112, 113, 5, 83, 0, 0, 113, 114, 5, 69, 0, 0, 114, 115, 5, 73, 0,
		0, 115, 116, 5, 70, 0, 0, 116, 22, 1, 0, 0, 0, 117, 118, 5, 69, 0, 0, 118,
		119, 5, 76, 0, 0, 119, 120, 5, 83, 0, 0, 120, 121, 5, 69, 0, 0, 121, 24,
		1, 0, 0, 0, 122, 123, 5, 69, 0, 0, 123, 124, 5, 78, 0, 0, 124, 125, 5,
		68, 0, 0, 125, 126, 5, 95, 0, 0, 126, 127, 5, 73, 0, 0, 127, 128, 5, 70,
		0, 0, 128, 26, 1, 0, 0, 0, 129, 130, 5, 40, 0, 0, 130, 28, 1, 0, 0, 0,
		131, 132, 5, 41, 0, 0, 132, 30, 1, 0, 0, 0, 133, 134, 5, 42, 0, 0, 134,
		32, 1, 0, 0, 0, 135, 136, 5, 47, 0, 0, 136, 34, 1, 0, 0, 0, 137, 138, 5,
		77, 0, 0, 138, 139, 5, 79, 0, 0, 139, 140, 5, 68, 0, 0, 140, 36, 1, 0,
		0, 0, 141, 142, 5, 43, 0, 0, 142, 38, 1, 0, 0, 0, 143, 144, 5, 45, 0, 0,
		144, 40, 1, 0, 0, 0, 145, 146, 5, 62, 0, 0, 146, 42, 1, 0, 0, 0, 147, 148,
		5, 62, 0, 0, 148, 149, 5, 61, 0, 0, 149, 44, 1, 0, 0, 0, 150, 151, 5, 60,
		0, 0, 151, 46, 1, 0, 0, 0, 152, 153, 5, 60, 0, 0, 153, 154, 5, 61, 0, 0,
		154, 48, 1, 0, 0, 0, 155, 156, 5, 61, 0, 0, 156, 50, 1, 0, 0, 0, 157, 158,
		5, 60, 0, 0, 158, 159, 5, 62, 0, 0, 159, 52, 1, 0, 0, 0, 160, 164, 2, 48,
		57, 0, 161, 163, 2, 48, 57, 0, 162, 161, 1, 0, 0, 0, 163, 166, 1, 0, 0,
		0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 54, 1, 0, 0, 0, 166,
		164, 1, 0, 0, 0, 167, 171, 7, 0, 0, 0, 168, 170, 7, 1, 0, 0, 169, 168,
		1, 0, 0, 0, 170, 173, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 171, 172, 1, 0,
		0, 0, 172, 56, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 174, 176, 7, 2, 0, 0,
		175, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 177,
		178, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 180, 6, 28, 0, 0, 180, 58,
		1, 0, 0, 0, 4, 0, 164, 171, 177, 1, 6, 0, 0,
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

// STLexerInit initializes any static state used to implement STLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSTLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func STLexerInit() {
	staticData := &STLexerLexerStaticData
	staticData.once.Do(stlexerLexerInit)
}

// NewSTLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSTLexer(input antlr.CharStream) *STLexer {
	STLexerInit()
	l := new(STLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &STLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "ST.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// STLexer tokens.
const (
	STLexerT__0       = 1
	STLexerT__1       = 2
	STLexerT__2       = 3
	STLexerT__3       = 4
	STLexerT__4       = 5
	STLexerT__5       = 6
	STLexerT__6       = 7
	STLexerT__7       = 8
	STLexerT__8       = 9
	STLexerT__9       = 10
	STLexerT__10      = 11
	STLexerT__11      = 12
	STLexerT__12      = 13
	STLexerT__13      = 14
	STLexerT__14      = 15
	STLexerT__15      = 16
	STLexerT__16      = 17
	STLexerT__17      = 18
	STLexerT__18      = 19
	STLexerT__19      = 20
	STLexerT__20      = 21
	STLexerT__21      = 22
	STLexerT__22      = 23
	STLexerT__23      = 24
	STLexerT__24      = 25
	STLexerT__25      = 26
	STLexerInteger    = 27
	STLexerID         = 28
	STLexerWHITESPACE = 29
)
