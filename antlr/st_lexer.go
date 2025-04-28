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
		"'INT'", "':='", "'('", "')'", "'*'", "'/'", "'+'", "'-'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Integer",
		"ID", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "Integer", "ID", "WHITESPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 17, 111, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 14, 1, 14, 5, 14, 93, 8, 14, 10, 14, 12, 14, 96, 9, 14, 1, 15, 1, 15,
		5, 15, 100, 8, 15, 10, 15, 12, 15, 103, 9, 15, 1, 16, 4, 16, 106, 8, 16,
		11, 16, 12, 16, 107, 1, 16, 1, 16, 0, 0, 17, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 1, 0, 3, 3, 0, 65, 90, 95, 95, 97, 122, 5, 0, 36,
		36, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 113,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 1, 35, 1, 0, 0, 0, 3, 43, 1, 0, 0, 0, 5, 55,
		1, 0, 0, 0, 7, 59, 1, 0, 0, 0, 9, 67, 1, 0, 0, 0, 11, 69, 1, 0, 0, 0, 13,
		71, 1, 0, 0, 0, 15, 75, 1, 0, 0, 0, 17, 78, 1, 0, 0, 0, 19, 80, 1, 0, 0,
		0, 21, 82, 1, 0, 0, 0, 23, 84, 1, 0, 0, 0, 25, 86, 1, 0, 0, 0, 27, 88,
		1, 0, 0, 0, 29, 90, 1, 0, 0, 0, 31, 97, 1, 0, 0, 0, 33, 105, 1, 0, 0, 0,
		35, 36, 5, 80, 0, 0, 36, 37, 5, 82, 0, 0, 37, 38, 5, 79, 0, 0, 38, 39,
		5, 71, 0, 0, 39, 40, 5, 82, 0, 0, 40, 41, 5, 65, 0, 0, 41, 42, 5, 77, 0,
		0, 42, 2, 1, 0, 0, 0, 43, 44, 5, 69, 0, 0, 44, 45, 5, 78, 0, 0, 45, 46,
		5, 68, 0, 0, 46, 47, 5, 95, 0, 0, 47, 48, 5, 80, 0, 0, 48, 49, 5, 82, 0,
		0, 49, 50, 5, 79, 0, 0, 50, 51, 5, 71, 0, 0, 51, 52, 5, 82, 0, 0, 52, 53,
		5, 65, 0, 0, 53, 54, 5, 77, 0, 0, 54, 4, 1, 0, 0, 0, 55, 56, 5, 86, 0,
		0, 56, 57, 5, 65, 0, 0, 57, 58, 5, 82, 0, 0, 58, 6, 1, 0, 0, 0, 59, 60,
		5, 69, 0, 0, 60, 61, 5, 78, 0, 0, 61, 62, 5, 68, 0, 0, 62, 63, 5, 95, 0,
		0, 63, 64, 5, 86, 0, 0, 64, 65, 5, 65, 0, 0, 65, 66, 5, 82, 0, 0, 66, 8,
		1, 0, 0, 0, 67, 68, 5, 58, 0, 0, 68, 10, 1, 0, 0, 0, 69, 70, 5, 59, 0,
		0, 70, 12, 1, 0, 0, 0, 71, 72, 5, 73, 0, 0, 72, 73, 5, 78, 0, 0, 73, 74,
		5, 84, 0, 0, 74, 14, 1, 0, 0, 0, 75, 76, 5, 58, 0, 0, 76, 77, 5, 61, 0,
		0, 77, 16, 1, 0, 0, 0, 78, 79, 5, 40, 0, 0, 79, 18, 1, 0, 0, 0, 80, 81,
		5, 41, 0, 0, 81, 20, 1, 0, 0, 0, 82, 83, 5, 42, 0, 0, 83, 22, 1, 0, 0,
		0, 84, 85, 5, 47, 0, 0, 85, 24, 1, 0, 0, 0, 86, 87, 5, 43, 0, 0, 87, 26,
		1, 0, 0, 0, 88, 89, 5, 45, 0, 0, 89, 28, 1, 0, 0, 0, 90, 94, 2, 48, 57,
		0, 91, 93, 2, 48, 57, 0, 92, 91, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94, 92,
		1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 30, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0,
		97, 101, 7, 0, 0, 0, 98, 100, 7, 1, 0, 0, 99, 98, 1, 0, 0, 0, 100, 103,
		1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 32, 1, 0, 0,
		0, 103, 101, 1, 0, 0, 0, 104, 106, 7, 2, 0, 0, 105, 104, 1, 0, 0, 0, 106,
		107, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 109,
		1, 0, 0, 0, 109, 110, 6, 16, 0, 0, 110, 34, 1, 0, 0, 0, 4, 0, 94, 101,
		107, 1, 6, 0, 0,
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
	STLexerInteger    = 15
	STLexerID         = 16
	STLexerWHITESPACE = 17
)
