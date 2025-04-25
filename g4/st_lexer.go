// Code generated from ./g4/st.g4 by ANTLR 4.13.2. DO NOT EDIT.

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

type stLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var StLexerLexerStaticData struct {
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
	staticData := &StLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'PROGRAM'", "'END_PROGRAM'", "'VAR'", "'END_VAR'", "':'", "';'",
		"'INT'", "':='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "Integer", "ID", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "Integer",
		"ID", "WHITESPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 11, 87, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 5, 8,
		69, 8, 8, 10, 8, 12, 8, 72, 9, 8, 1, 9, 1, 9, 5, 9, 76, 8, 9, 10, 9, 12,
		9, 79, 9, 9, 1, 10, 4, 10, 82, 8, 10, 11, 10, 12, 10, 83, 1, 10, 1, 10,
		0, 0, 11, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 1, 0, 3, 3, 0, 65, 90, 95, 95, 97, 122, 5, 0, 36, 36, 48, 57,
		65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 89, 0, 1, 1, 0, 0,
		0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0,
		0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0,
		0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 1, 23, 1, 0, 0, 0, 3, 31, 1,
		0, 0, 0, 5, 43, 1, 0, 0, 0, 7, 47, 1, 0, 0, 0, 9, 55, 1, 0, 0, 0, 11, 57,
		1, 0, 0, 0, 13, 59, 1, 0, 0, 0, 15, 63, 1, 0, 0, 0, 17, 66, 1, 0, 0, 0,
		19, 73, 1, 0, 0, 0, 21, 81, 1, 0, 0, 0, 23, 24, 5, 80, 0, 0, 24, 25, 5,
		82, 0, 0, 25, 26, 5, 79, 0, 0, 26, 27, 5, 71, 0, 0, 27, 28, 5, 82, 0, 0,
		28, 29, 5, 65, 0, 0, 29, 30, 5, 77, 0, 0, 30, 2, 1, 0, 0, 0, 31, 32, 5,
		69, 0, 0, 32, 33, 5, 78, 0, 0, 33, 34, 5, 68, 0, 0, 34, 35, 5, 95, 0, 0,
		35, 36, 5, 80, 0, 0, 36, 37, 5, 82, 0, 0, 37, 38, 5, 79, 0, 0, 38, 39,
		5, 71, 0, 0, 39, 40, 5, 82, 0, 0, 40, 41, 5, 65, 0, 0, 41, 42, 5, 77, 0,
		0, 42, 4, 1, 0, 0, 0, 43, 44, 5, 86, 0, 0, 44, 45, 5, 65, 0, 0, 45, 46,
		5, 82, 0, 0, 46, 6, 1, 0, 0, 0, 47, 48, 5, 69, 0, 0, 48, 49, 5, 78, 0,
		0, 49, 50, 5, 68, 0, 0, 50, 51, 5, 95, 0, 0, 51, 52, 5, 86, 0, 0, 52, 53,
		5, 65, 0, 0, 53, 54, 5, 82, 0, 0, 54, 8, 1, 0, 0, 0, 55, 56, 5, 58, 0,
		0, 56, 10, 1, 0, 0, 0, 57, 58, 5, 59, 0, 0, 58, 12, 1, 0, 0, 0, 59, 60,
		5, 73, 0, 0, 60, 61, 5, 78, 0, 0, 61, 62, 5, 84, 0, 0, 62, 14, 1, 0, 0,
		0, 63, 64, 5, 58, 0, 0, 64, 65, 5, 61, 0, 0, 65, 16, 1, 0, 0, 0, 66, 70,
		2, 48, 57, 0, 67, 69, 2, 48, 57, 0, 68, 67, 1, 0, 0, 0, 69, 72, 1, 0, 0,
		0, 70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 18, 1, 0, 0, 0, 72, 70,
		1, 0, 0, 0, 73, 77, 7, 0, 0, 0, 74, 76, 7, 1, 0, 0, 75, 74, 1, 0, 0, 0,
		76, 79, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 20, 1,
		0, 0, 0, 79, 77, 1, 0, 0, 0, 80, 82, 7, 2, 0, 0, 81, 80, 1, 0, 0, 0, 82,
		83, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 85, 1, 0, 0,
		0, 85, 86, 6, 10, 0, 0, 86, 22, 1, 0, 0, 0, 4, 0, 70, 77, 83, 1, 6, 0,
		0,
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

// stLexerInit initializes any static state used to implement stLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewstLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func StLexerInit() {
	staticData := &StLexerLexerStaticData
	staticData.once.Do(stlexerLexerInit)
}

// NewstLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewstLexer(input antlr.CharStream) *stLexer {
	StLexerInit()
	l := new(stLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &StLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "st.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// stLexer tokens.
const (
	stLexerT__0       = 1
	stLexerT__1       = 2
	stLexerT__2       = 3
	stLexerT__3       = 4
	stLexerT__4       = 5
	stLexerT__5       = 6
	stLexerT__6       = 7
	stLexerT__7       = 8
	stLexerInteger    = 9
	stLexerID         = 10
	stLexerWHITESPACE = 11
)
