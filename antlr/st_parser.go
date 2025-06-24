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
		"", "'PROGRAM'", "'END_PROGRAM'", "'VAR'", "'END_VAR'", "':'", "':='",
		"';'", "'SINT'", "'INT'", "'DINT'", "'LINT'", "'USINT'", "'UINT'", "'UDINT'",
		"'ULINT'", "'REAL'", "'LREAL'", "'BOOL'", "'STRING'", "'IF'", "'THEN'",
		"'ELSIF'", "'ELSE'", "'END_IF'", "'WHILE'", "'DO'", "'END_WHILE'", "'('",
		"')'", "'*'", "'/'", "'MOD'", "'>'", "'>='", "'<'", "'<='", "'='", "'<>'",
		"", "", "", "", "", "", "'+'", "'-'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "Integer", "Hex_Int", "FLOAT", "ID", "Digit", "Hex_Digit",
		"PLUS", "MINUS", "Comment", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"program", "var_declaration_blocks", "var_declaration_block", "var_declaration",
		"type_name", "statement_list", "statement", "assignment_statement",
		"if_statement", "while_statement", "condition", "then_list", "else_list",
		"expression", "number", "integer", "signed_integer", "unsign_integer",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 48, 167, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 5,
		1, 44, 8, 1, 10, 1, 12, 1, 47, 9, 1, 1, 2, 1, 2, 5, 2, 51, 8, 2, 10, 2,
		12, 2, 54, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 63, 8,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 5, 5, 70, 8, 5, 10, 5, 12, 5, 73, 9, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 80, 8, 6, 1, 6, 1, 6, 3, 6, 84, 8,
		6, 3, 6, 86, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 101, 8, 8, 10, 8, 12, 8, 104, 9, 8, 1, 8,
		1, 8, 3, 8, 108, 8, 8, 1, 8, 1, 8, 3, 8, 112, 8, 8, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3,
		13, 138, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 5, 13, 149, 8, 13, 10, 13, 12, 13, 152, 9, 13, 1, 14, 1, 14, 3,
		14, 156, 8, 14, 1, 15, 1, 15, 3, 15, 160, 8, 15, 1, 16, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 17, 0, 1, 26, 18, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 0, 5, 1, 0, 8, 19, 1, 0, 30, 32, 1, 0, 45,
		46, 1, 0, 33, 38, 1, 0, 39, 40, 167, 0, 36, 1, 0, 0, 0, 2, 45, 1, 0, 0,
		0, 4, 48, 1, 0, 0, 0, 6, 57, 1, 0, 0, 0, 8, 66, 1, 0, 0, 0, 10, 71, 1,
		0, 0, 0, 12, 85, 1, 0, 0, 0, 14, 87, 1, 0, 0, 0, 16, 91, 1, 0, 0, 0, 18,
		113, 1, 0, 0, 0, 20, 119, 1, 0, 0, 0, 22, 121, 1, 0, 0, 0, 24, 123, 1,
		0, 0, 0, 26, 137, 1, 0, 0, 0, 28, 155, 1, 0, 0, 0, 30, 159, 1, 0, 0, 0,
		32, 161, 1, 0, 0, 0, 34, 164, 1, 0, 0, 0, 36, 37, 5, 1, 0, 0, 37, 38, 5,
		42, 0, 0, 38, 39, 3, 2, 1, 0, 39, 40, 3, 10, 5, 0, 40, 41, 5, 2, 0, 0,
		41, 1, 1, 0, 0, 0, 42, 44, 3, 4, 2, 0, 43, 42, 1, 0, 0, 0, 44, 47, 1, 0,
		0, 0, 45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 3, 1, 0, 0, 0, 47, 45,
		1, 0, 0, 0, 48, 52, 5, 3, 0, 0, 49, 51, 3, 6, 3, 0, 50, 49, 1, 0, 0, 0,
		51, 54, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 55, 1,
		0, 0, 0, 54, 52, 1, 0, 0, 0, 55, 56, 5, 4, 0, 0, 56, 5, 1, 0, 0, 0, 57,
		58, 5, 42, 0, 0, 58, 59, 5, 5, 0, 0, 59, 62, 3, 8, 4, 0, 60, 61, 5, 6,
		0, 0, 61, 63, 3, 28, 14, 0, 62, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63,
		64, 1, 0, 0, 0, 64, 65, 5, 7, 0, 0, 65, 7, 1, 0, 0, 0, 66, 67, 7, 0, 0,
		0, 67, 9, 1, 0, 0, 0, 68, 70, 3, 12, 6, 0, 69, 68, 1, 0, 0, 0, 70, 73,
		1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 11, 1, 0, 0, 0,
		73, 71, 1, 0, 0, 0, 74, 75, 3, 14, 7, 0, 75, 76, 5, 7, 0, 0, 76, 86, 1,
		0, 0, 0, 77, 79, 3, 16, 8, 0, 78, 80, 5, 7, 0, 0, 79, 78, 1, 0, 0, 0, 79,
		80, 1, 0, 0, 0, 80, 86, 1, 0, 0, 0, 81, 83, 3, 18, 9, 0, 82, 84, 5, 7,
		0, 0, 83, 82, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 86, 1, 0, 0, 0, 85, 74,
		1, 0, 0, 0, 85, 77, 1, 0, 0, 0, 85, 81, 1, 0, 0, 0, 86, 13, 1, 0, 0, 0,
		87, 88, 5, 42, 0, 0, 88, 89, 5, 6, 0, 0, 89, 90, 3, 26, 13, 0, 90, 15,
		1, 0, 0, 0, 91, 92, 5, 20, 0, 0, 92, 93, 3, 20, 10, 0, 93, 94, 5, 21, 0,
		0, 94, 102, 3, 22, 11, 0, 95, 96, 5, 22, 0, 0, 96, 97, 3, 20, 10, 0, 97,
		98, 5, 21, 0, 0, 98, 99, 3, 22, 11, 0, 99, 101, 1, 0, 0, 0, 100, 95, 1,
		0, 0, 0, 101, 104, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0,
		0, 103, 107, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 105, 106, 5, 23, 0, 0, 106,
		108, 3, 24, 12, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 109,
		1, 0, 0, 0, 109, 111, 5, 24, 0, 0, 110, 112, 5, 7, 0, 0, 111, 110, 1, 0,
		0, 0, 111, 112, 1, 0, 0, 0, 112, 17, 1, 0, 0, 0, 113, 114, 5, 25, 0, 0,
		114, 115, 3, 26, 13, 0, 115, 116, 5, 26, 0, 0, 116, 117, 3, 10, 5, 0, 117,
		118, 5, 27, 0, 0, 118, 19, 1, 0, 0, 0, 119, 120, 3, 26, 13, 0, 120, 21,
		1, 0, 0, 0, 121, 122, 3, 10, 5, 0, 122, 23, 1, 0, 0, 0, 123, 124, 3, 10,
		5, 0, 124, 25, 1, 0, 0, 0, 125, 126, 6, 13, -1, 0, 126, 138, 3, 28, 14,
		0, 127, 128, 5, 42, 0, 0, 128, 129, 5, 28, 0, 0, 129, 130, 3, 26, 13, 0,
		130, 131, 5, 29, 0, 0, 131, 138, 1, 0, 0, 0, 132, 138, 5, 42, 0, 0, 133,
		134, 5, 28, 0, 0, 134, 135, 3, 26, 13, 0, 135, 136, 5, 29, 0, 0, 136, 138,
		1, 0, 0, 0, 137, 125, 1, 0, 0, 0, 137, 127, 1, 0, 0, 0, 137, 132, 1, 0,
		0, 0, 137, 133, 1, 0, 0, 0, 138, 150, 1, 0, 0, 0, 139, 140, 10, 3, 0, 0,
		140, 141, 7, 1, 0, 0, 141, 149, 3, 26, 13, 4, 142, 143, 10, 2, 0, 0, 143,
		144, 7, 2, 0, 0, 144, 149, 3, 26, 13, 3, 145, 146, 10, 1, 0, 0, 146, 147,
		7, 3, 0, 0, 147, 149, 3, 26, 13, 2, 148, 139, 1, 0, 0, 0, 148, 142, 1,
		0, 0, 0, 148, 145, 1, 0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148, 1, 0, 0,
		0, 150, 151, 1, 0, 0, 0, 151, 27, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 153,
		156, 5, 41, 0, 0, 154, 156, 3, 30, 15, 0, 155, 153, 1, 0, 0, 0, 155, 154,
		1, 0, 0, 0, 156, 29, 1, 0, 0, 0, 157, 160, 3, 32, 16, 0, 158, 160, 3, 34,
		17, 0, 159, 157, 1, 0, 0, 0, 159, 158, 1, 0, 0, 0, 160, 31, 1, 0, 0, 0,
		161, 162, 7, 2, 0, 0, 162, 163, 3, 34, 17, 0, 163, 33, 1, 0, 0, 0, 164,
		165, 7, 4, 0, 0, 165, 35, 1, 0, 0, 0, 15, 45, 52, 62, 71, 79, 83, 85, 102,
		107, 111, 137, 148, 150, 155, 159,
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
	STParserEOF        = antlr.TokenEOF
	STParserT__0       = 1
	STParserT__1       = 2
	STParserT__2       = 3
	STParserT__3       = 4
	STParserT__4       = 5
	STParserT__5       = 6
	STParserT__6       = 7
	STParserT__7       = 8
	STParserT__8       = 9
	STParserT__9       = 10
	STParserT__10      = 11
	STParserT__11      = 12
	STParserT__12      = 13
	STParserT__13      = 14
	STParserT__14      = 15
	STParserT__15      = 16
	STParserT__16      = 17
	STParserT__17      = 18
	STParserT__18      = 19
	STParserT__19      = 20
	STParserT__20      = 21
	STParserT__21      = 22
	STParserT__22      = 23
	STParserT__23      = 24
	STParserT__24      = 25
	STParserT__25      = 26
	STParserT__26      = 27
	STParserT__27      = 28
	STParserT__28      = 29
	STParserT__29      = 30
	STParserT__30      = 31
	STParserT__31      = 32
	STParserT__32      = 33
	STParserT__33      = 34
	STParserT__34      = 35
	STParserT__35      = 36
	STParserT__36      = 37
	STParserT__37      = 38
	STParserInteger    = 39
	STParserHex_Int    = 40
	STParserFLOAT      = 41
	STParserID         = 42
	STParserDigit      = 43
	STParserHex_Digit  = 44
	STParserPLUS       = 45
	STParserMINUS      = 46
	STParserComment    = 47
	STParserWHITESPACE = 48
)

// STParser rules.
const (
	STParserRULE_program                = 0
	STParserRULE_var_declaration_blocks = 1
	STParserRULE_var_declaration_block  = 2
	STParserRULE_var_declaration        = 3
	STParserRULE_type_name              = 4
	STParserRULE_statement_list         = 5
	STParserRULE_statement              = 6
	STParserRULE_assignment_statement   = 7
	STParserRULE_if_statement           = 8
	STParserRULE_while_statement        = 9
	STParserRULE_condition              = 10
	STParserRULE_then_list              = 11
	STParserRULE_else_list              = 12
	STParserRULE_expression             = 13
	STParserRULE_number                 = 14
	STParserRULE_integer                = 15
	STParserRULE_signed_integer         = 16
	STParserRULE_unsign_integer         = 17
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIdentifier returns the identifier token.
	GetIdentifier() antlr.Token

	// SetIdentifier sets the identifier token.
	SetIdentifier(antlr.Token)

	// Getter signatures
	Var_declaration_blocks() IVar_declaration_blocksContext
	Statement_list() IStatement_listContext
	ID() antlr.TerminalNode

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*CustomContext
	parser     antlr.Parser
	identifier antlr.Token
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

func (s *ProgramContext) GetIdentifier() antlr.Token { return s.identifier }

func (s *ProgramContext) SetIdentifier(v antlr.Token) { s.identifier = v }

func (s *ProgramContext) Var_declaration_blocks() IVar_declaration_blocksContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declaration_blocksContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_declaration_blocksContext)
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

func (s *ProgramContext) ID() antlr.TerminalNode {
	return s.GetToken(STParserID, 0)
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		p.Match(STParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(37)

		var _m = p.Match(STParserID)

		localctx.(*ProgramContext).identifier = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(38)
		p.Var_declaration_blocks()
	}
	{
		p.SetState(39)
		p.Statement_list()
	}
	{
		p.SetState(40)
		p.Match(STParserT__1)
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

// IVar_declaration_blocksContext is an interface to support dynamic dispatch.
type IVar_declaration_blocksContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVar_declaration_block() []IVar_declaration_blockContext
	Var_declaration_block(i int) IVar_declaration_blockContext

	// IsVar_declaration_blocksContext differentiates from other interfaces.
	IsVar_declaration_blocksContext()
}

type Var_declaration_blocksContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyVar_declaration_blocksContext() *Var_declaration_blocksContext {
	var p = new(Var_declaration_blocksContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_var_declaration_blocks
	return p
}

func InitEmptyVar_declaration_blocksContext(p *Var_declaration_blocksContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_var_declaration_blocks
}

func (*Var_declaration_blocksContext) IsVar_declaration_blocksContext() {}

func NewVar_declaration_blocksContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_declaration_blocksContext {
	var p = new(Var_declaration_blocksContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_var_declaration_blocks

	return p
}

func (s *Var_declaration_blocksContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_declaration_blocksContext) AllVar_declaration_block() []IVar_declaration_blockContext {
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

func (s *Var_declaration_blocksContext) Var_declaration_block(i int) IVar_declaration_blockContext {
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

func (s *Var_declaration_blocksContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_declaration_blocksContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_declaration_blocksContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitVar_declaration_blocks(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Var_declaration_blocks() (localctx IVar_declaration_blocksContext) {
	localctx = NewVar_declaration_blocksContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, STParserRULE_var_declaration_blocks)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == STParserT__2 {
		{
			p.SetState(42)
			p.Var_declaration_block()
		}

		p.SetState(47)
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

// IVar_declaration_blockContext is an interface to support dynamic dispatch.
type IVar_declaration_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVar_declaration() []IVar_declarationContext
	Var_declaration(i int) IVar_declarationContext

	// IsVar_declaration_blockContext differentiates from other interfaces.
	IsVar_declaration_blockContext()
}

type Var_declaration_blockContext struct {
	*CustomContext
	parser antlr.Parser
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

func (s *Var_declaration_blockContext) AllVar_declaration() []IVar_declarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_declarationContext); ok {
			len++
		}
	}

	tst := make([]IVar_declarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_declarationContext); ok {
			tst[i] = t.(IVar_declarationContext)
			i++
		}
	}

	return tst
}

func (s *Var_declaration_blockContext) Var_declaration(i int) IVar_declarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declarationContext); ok {
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

	return t.(IVar_declarationContext)
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
	p.EnterRule(localctx, 4, STParserRULE_var_declaration_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(STParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == STParserID {
		{
			p.SetState(49)
			p.Var_declaration()
		}

		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(55)
		p.Match(STParserT__3)
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

// IVar_declarationContext is an interface to support dynamic dispatch.
type IVar_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIdentifier returns the identifier token.
	GetIdentifier() antlr.Token

	// SetIdentifier sets the identifier token.
	SetIdentifier(antlr.Token)

	// GetType_ returns the type_ rule contexts.
	GetType_() IType_nameContext

	// GetDefault_ returns the default_ rule contexts.
	GetDefault_() INumberContext

	// SetType_ sets the type_ rule contexts.
	SetType_(IType_nameContext)

	// SetDefault_ sets the default_ rule contexts.
	SetDefault_(INumberContext)

	// Getter signatures
	ID() antlr.TerminalNode
	Type_name() IType_nameContext
	Number() INumberContext

	// IsVar_declarationContext differentiates from other interfaces.
	IsVar_declarationContext()
}

type Var_declarationContext struct {
	*CustomContext
	parser     antlr.Parser
	identifier antlr.Token
	type_      IType_nameContext
	default_   INumberContext
}

func NewEmptyVar_declarationContext() *Var_declarationContext {
	var p = new(Var_declarationContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_var_declaration
	return p
}

func InitEmptyVar_declarationContext(p *Var_declarationContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_var_declaration
}

func (*Var_declarationContext) IsVar_declarationContext() {}

func NewVar_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_declarationContext {
	var p = new(Var_declarationContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_var_declaration

	return p
}

func (s *Var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_declarationContext) GetIdentifier() antlr.Token { return s.identifier }

func (s *Var_declarationContext) SetIdentifier(v antlr.Token) { s.identifier = v }

func (s *Var_declarationContext) GetType_() IType_nameContext { return s.type_ }

func (s *Var_declarationContext) GetDefault_() INumberContext { return s.default_ }

func (s *Var_declarationContext) SetType_(v IType_nameContext) { s.type_ = v }

func (s *Var_declarationContext) SetDefault_(v INumberContext) { s.default_ = v }

func (s *Var_declarationContext) ID() antlr.TerminalNode {
	return s.GetToken(STParserID, 0)
}

func (s *Var_declarationContext) Type_name() IType_nameContext {
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

func (s *Var_declarationContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *Var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitVar_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Var_declaration() (localctx IVar_declarationContext) {
	localctx = NewVar_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, STParserRULE_var_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)

		var _m = p.Match(STParserID)

		localctx.(*Var_declarationContext).identifier = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
		p.Match(STParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(59)

		var _x = p.Type_name()

		localctx.(*Var_declarationContext).type_ = _x
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserT__5 {
		{
			p.SetState(60)
			p.Match(STParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)

			var _x = p.Number()

			localctx.(*Var_declarationContext).default_ = _x
		}

	}
	{
		p.SetState(64)
		p.Match(STParserT__6)
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

// IType_nameContext is an interface to support dynamic dispatch.
type IType_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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
	p.EnterRule(localctx, 8, STParserRULE_type_name)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1048320) != 0) {
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
	p.EnterRule(localctx, 10, STParserRULE_statement_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4398081114112) != 0 {
		{
			p.SetState(68)
			p.Statement()
		}

		p.SetState(73)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment_statement() IAssignment_statementContext
	If_statement() IIf_statementContext
	While_statement() IWhile_statementContext

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
	p.EnterRule(localctx, 12, STParserRULE_statement)
	var _la int

	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case STParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.Assignment_statement()
		}
		{
			p.SetState(75)
			p.Match(STParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case STParserT__19:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.If_statement()
		}
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == STParserT__6 {
			{
				p.SetState(78)
				p.Match(STParserT__6)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case STParserT__24:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(81)
			p.While_statement()
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == STParserT__6 {
			{
				p.SetState(82)
				p.Match(STParserT__6)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
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

// IAssignment_statementContext is an interface to support dynamic dispatch.
type IAssignment_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left token.
	GetLeft() antlr.Token

	// SetLeft sets the left token.
	SetLeft(antlr.Token)

	// GetRight returns the right rule contexts.
	GetRight() IExpressionContext

	// SetRight sets the right rule contexts.
	SetRight(IExpressionContext)

	// Getter signatures
	ID() antlr.TerminalNode
	Expression() IExpressionContext

	// IsAssignment_statementContext differentiates from other interfaces.
	IsAssignment_statementContext()
}

type Assignment_statementContext struct {
	*CustomContext
	parser antlr.Parser
	left   antlr.Token
	right  IExpressionContext
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

func (s *Assignment_statementContext) GetLeft() antlr.Token { return s.left }

func (s *Assignment_statementContext) SetLeft(v antlr.Token) { s.left = v }

func (s *Assignment_statementContext) GetRight() IExpressionContext { return s.right }

func (s *Assignment_statementContext) SetRight(v IExpressionContext) { s.right = v }

func (s *Assignment_statementContext) ID() antlr.TerminalNode {
	return s.GetToken(STParserID, 0)
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
	p.EnterRule(localctx, 14, STParserRULE_assignment_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)

		var _m = p.Match(STParserID)

		localctx.(*Assignment_statementContext).left = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(88)
		p.Match(STParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)

		var _x = p.expression(0)

		localctx.(*Assignment_statementContext).right = _x
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

	// Get_condition returns the _condition rule contexts.
	Get_condition() IConditionContext

	// Get_then_list returns the _then_list rule contexts.
	Get_then_list() IThen_listContext

	// GetElselist returns the elselist rule contexts.
	GetElselist() IElse_listContext

	// Set_condition sets the _condition rule contexts.
	Set_condition(IConditionContext)

	// Set_then_list sets the _then_list rule contexts.
	Set_then_list(IThen_listContext)

	// SetElselist sets the elselist rule contexts.
	SetElselist(IElse_listContext)

	// GetCond returns the cond rule context list.
	GetCond() []IConditionContext

	// GetThenlist returns the thenlist rule context list.
	GetThenlist() []IThen_listContext

	// SetCond sets the cond rule context list.
	SetCond([]IConditionContext)

	// SetThenlist sets the thenlist rule context list.
	SetThenlist([]IThen_listContext)

	// Getter signatures
	AllCondition() []IConditionContext
	Condition(i int) IConditionContext
	AllThen_list() []IThen_listContext
	Then_list(i int) IThen_listContext
	Else_list() IElse_listContext

	// IsIf_statementContext differentiates from other interfaces.
	IsIf_statementContext()
}

type If_statementContext struct {
	*CustomContext
	parser     antlr.Parser
	_condition IConditionContext
	cond       []IConditionContext
	_then_list IThen_listContext
	thenlist   []IThen_listContext
	elselist   IElse_listContext
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

func (s *If_statementContext) Get_condition() IConditionContext { return s._condition }

func (s *If_statementContext) Get_then_list() IThen_listContext { return s._then_list }

func (s *If_statementContext) GetElselist() IElse_listContext { return s.elselist }

func (s *If_statementContext) Set_condition(v IConditionContext) { s._condition = v }

func (s *If_statementContext) Set_then_list(v IThen_listContext) { s._then_list = v }

func (s *If_statementContext) SetElselist(v IElse_listContext) { s.elselist = v }

func (s *If_statementContext) GetCond() []IConditionContext { return s.cond }

func (s *If_statementContext) GetThenlist() []IThen_listContext { return s.thenlist }

func (s *If_statementContext) SetCond(v []IConditionContext) { s.cond = v }

func (s *If_statementContext) SetThenlist(v []IThen_listContext) { s.thenlist = v }

func (s *If_statementContext) AllCondition() []IConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionContext); ok {
			len++
		}
	}

	tst := make([]IConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionContext); ok {
			tst[i] = t.(IConditionContext)
			i++
		}
	}

	return tst
}

func (s *If_statementContext) Condition(i int) IConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
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

	return t.(IConditionContext)
}

func (s *If_statementContext) AllThen_list() []IThen_listContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IThen_listContext); ok {
			len++
		}
	}

	tst := make([]IThen_listContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IThen_listContext); ok {
			tst[i] = t.(IThen_listContext)
			i++
		}
	}

	return tst
}

func (s *If_statementContext) Then_list(i int) IThen_listContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IThen_listContext); ok {
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

	return t.(IThen_listContext)
}

func (s *If_statementContext) Else_list() IElse_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElse_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElse_listContext)
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
	p.EnterRule(localctx, 16, STParserRULE_if_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(STParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(92)

		var _x = p.Condition()

		localctx.(*If_statementContext)._condition = _x
	}
	localctx.(*If_statementContext).cond = append(localctx.(*If_statementContext).cond, localctx.(*If_statementContext)._condition)
	{
		p.SetState(93)
		p.Match(STParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)

		var _x = p.Then_list()

		localctx.(*If_statementContext)._then_list = _x
	}
	localctx.(*If_statementContext).thenlist = append(localctx.(*If_statementContext).thenlist, localctx.(*If_statementContext)._then_list)
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == STParserT__21 {
		{
			p.SetState(95)
			p.Match(STParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)

			var _x = p.Condition()

			localctx.(*If_statementContext)._condition = _x
		}
		localctx.(*If_statementContext).cond = append(localctx.(*If_statementContext).cond, localctx.(*If_statementContext)._condition)
		{
			p.SetState(97)
			p.Match(STParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)

			var _x = p.Then_list()

			localctx.(*If_statementContext)._then_list = _x
		}
		localctx.(*If_statementContext).thenlist = append(localctx.(*If_statementContext).thenlist, localctx.(*If_statementContext)._then_list)

		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserT__22 {
		{
			p.SetState(105)
			p.Match(STParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)

			var _x = p.Else_list()

			localctx.(*If_statementContext).elselist = _x
		}

	}
	{
		p.SetState(109)
		p.Match(STParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(110)
			p.Match(STParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IWhile_statementContext is an interface to support dynamic dispatch.
type IWhile_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCond returns the cond rule contexts.
	GetCond() IExpressionContext

	// GetBody returns the body rule contexts.
	GetBody() IStatement_listContext

	// SetCond sets the cond rule contexts.
	SetCond(IExpressionContext)

	// SetBody sets the body rule contexts.
	SetBody(IStatement_listContext)

	// Getter signatures
	Expression() IExpressionContext
	Statement_list() IStatement_listContext

	// IsWhile_statementContext differentiates from other interfaces.
	IsWhile_statementContext()
}

type While_statementContext struct {
	*CustomContext
	parser antlr.Parser
	cond   IExpressionContext
	body   IStatement_listContext
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

func (s *While_statementContext) GetCond() IExpressionContext { return s.cond }

func (s *While_statementContext) GetBody() IStatement_listContext { return s.body }

func (s *While_statementContext) SetCond(v IExpressionContext) { s.cond = v }

func (s *While_statementContext) SetBody(v IStatement_listContext) { s.body = v }

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
	p.EnterRule(localctx, 18, STParserRULE_while_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.Match(STParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)

		var _x = p.expression(0)

		localctx.(*While_statementContext).cond = _x
	}
	{
		p.SetState(115)
		p.Match(STParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)

		var _x = p.Statement_list()

		localctx.(*While_statementContext).body = _x
	}
	{
		p.SetState(117)
		p.Match(STParserT__26)
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

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) Expression() IExpressionContext {
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

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, STParserRULE_condition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
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

// IThen_listContext is an interface to support dynamic dispatch.
type IThen_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Statement_list() IStatement_listContext

	// IsThen_listContext differentiates from other interfaces.
	IsThen_listContext()
}

type Then_listContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyThen_listContext() *Then_listContext {
	var p = new(Then_listContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_then_list
	return p
}

func InitEmptyThen_listContext(p *Then_listContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_then_list
}

func (*Then_listContext) IsThen_listContext() {}

func NewThen_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Then_listContext {
	var p = new(Then_listContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_then_list

	return p
}

func (s *Then_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Then_listContext) Statement_list() IStatement_listContext {
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

func (s *Then_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Then_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Then_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitThen_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Then_list() (localctx IThen_listContext) {
	localctx = NewThen_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, STParserRULE_then_list)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
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

// IElse_listContext is an interface to support dynamic dispatch.
type IElse_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Statement_list() IStatement_listContext

	// IsElse_listContext differentiates from other interfaces.
	IsElse_listContext()
}

type Else_listContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyElse_listContext() *Else_listContext {
	var p = new(Else_listContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_else_list
	return p
}

func InitEmptyElse_listContext(p *Else_listContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_else_list
}

func (*Else_listContext) IsElse_listContext() {}

func NewElse_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_listContext {
	var p = new(Else_listContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_else_list

	return p
}

func (s *Else_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_listContext) Statement_list() IStatement_listContext {
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

func (s *Else_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Else_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitElse_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Else_list() (localctx IElse_listContext) {
	localctx = NewElse_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, STParserRULE_else_list)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
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

type BinaryCompareExprContext struct {
	ExpressionContext
	left  IExpressionContext
	op    antlr.Token
	right IExpressionContext
}

func NewBinaryCompareExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryCompareExprContext {
	var p = new(BinaryCompareExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryCompareExprContext) GetOp() antlr.Token { return s.op }

func (s *BinaryCompareExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinaryCompareExprContext) GetLeft() IExpressionContext { return s.left }

func (s *BinaryCompareExprContext) GetRight() IExpressionContext { return s.right }

func (s *BinaryCompareExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *BinaryCompareExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *BinaryCompareExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryCompareExprContext) AllExpression() []IExpressionContext {
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

func (s *BinaryCompareExprContext) Expression(i int) IExpressionContext {
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

func (s *BinaryCompareExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitBinaryCompareExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConstantContext struct {
	ExpressionContext
}

func NewConstantContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstantContext {
	var p = new(ConstantContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinaryPowerExprContext struct {
	ExpressionContext
	left  IExpressionContext
	op    antlr.Token
	right IExpressionContext
}

func NewBinaryPowerExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryPowerExprContext {
	var p = new(BinaryPowerExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryPowerExprContext) GetOp() antlr.Token { return s.op }

func (s *BinaryPowerExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinaryPowerExprContext) GetLeft() IExpressionContext { return s.left }

func (s *BinaryPowerExprContext) GetRight() IExpressionContext { return s.right }

func (s *BinaryPowerExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *BinaryPowerExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *BinaryPowerExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryPowerExprContext) AllExpression() []IExpressionContext {
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

func (s *BinaryPowerExprContext) Expression(i int) IExpressionContext {
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

func (s *BinaryPowerExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitBinaryPowerExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type VariableContext struct {
	ExpressionContext
}

func NewVariableContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableContext {
	var p = new(VariableContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ID() antlr.TerminalNode {
	return s.GetToken(STParserID, 0)
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinaryPlusExprContext struct {
	ExpressionContext
	left  IExpressionContext
	op    antlr.Token
	right IExpressionContext
}

func NewBinaryPlusExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryPlusExprContext {
	var p = new(BinaryPlusExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryPlusExprContext) GetOp() antlr.Token { return s.op }

func (s *BinaryPlusExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinaryPlusExprContext) GetLeft() IExpressionContext { return s.left }

func (s *BinaryPlusExprContext) GetRight() IExpressionContext { return s.right }

func (s *BinaryPlusExprContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *BinaryPlusExprContext) SetRight(v IExpressionContext) { s.right = v }

func (s *BinaryPlusExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryPlusExprContext) AllExpression() []IExpressionContext {
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

func (s *BinaryPlusExprContext) Expression(i int) IExpressionContext {
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

func (s *BinaryPlusExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(STParserPLUS, 0)
}

func (s *BinaryPlusExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(STParserMINUS, 0)
}

func (s *BinaryPlusExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitBinaryPlusExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallExprContext struct {
	ExpressionContext
	id  antlr.Token
	sub IExpressionContext
}

func NewCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExprContext {
	var p = new(CallExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CallExprContext) GetId() antlr.Token { return s.id }

func (s *CallExprContext) SetId(v antlr.Token) { s.id = v }

func (s *CallExprContext) GetSub() IExpressionContext { return s.sub }

func (s *CallExprContext) SetSub(v IExpressionContext) { s.sub = v }

func (s *CallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExprContext) ID() antlr.TerminalNode {
	return s.GetToken(STParserID, 0)
}

func (s *CallExprContext) Expression() IExpressionContext {
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

func (s *CallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExprContext struct {
	ExpressionContext
	sub IExpressionContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ParenExprContext) GetSub() IExpressionContext { return s.sub }

func (s *ParenExprContext) SetSub(v IExpressionContext) { s.sub = v }

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) Expression() IExpressionContext {
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

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitParenExpr(s)

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
	_startState := 26
	p.EnterRecursionRule(localctx, 26, STParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewConstantContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(126)
			p.Number()
		}

	case 2:
		localctx = NewCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(127)

			var _m = p.Match(STParserID)

			localctx.(*CallExprContext).id = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.Match(STParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)

			var _x = p.expression(0)

			localctx.(*CallExprContext).sub = _x
		}
		{
			p.SetState(130)
			p.Match(STParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewVariableContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(132)
			p.Match(STParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(133)
			p.Match(STParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)

			var _x = p.expression(0)

			localctx.(*ParenExprContext).sub = _x
		}
		{
			p.SetState(135)
			p.Match(STParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(148)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryPowerExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryPowerExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, STParserRULE_expression)
				p.SetState(139)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(140)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryPowerExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7516192768) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryPowerExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(141)

					var _x = p.expression(4)

					localctx.(*BinaryPowerExprContext).right = _x
				}

			case 2:
				localctx = NewBinaryPlusExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryPlusExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, STParserRULE_expression)
				p.SetState(142)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(143)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryPlusExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == STParserPLUS || _la == STParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryPlusExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(144)

					var _x = p.expression(3)

					localctx.(*BinaryPlusExprContext).right = _x
				}

			case 3:
				localctx = NewBinaryCompareExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryCompareExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, STParserRULE_expression)
				p.SetState(145)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(146)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryCompareExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&541165879296) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryCompareExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(147)

					var _x = p.expression(2)

					localctx.(*BinaryCompareExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
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

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FLOAT() antlr.TerminalNode
	Integer() IIntegerContext

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(STParserFLOAT, 0)
}

func (s *NumberContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, STParserRULE_number)
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case STParserFLOAT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.Match(STParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case STParserInteger, STParserHex_Int, STParserPLUS, STParserMINUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)
			p.Integer()
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

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Signed_integer() ISigned_integerContext
	Unsign_integer() IUnsign_integerContext

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_integer
	return p
}

func InitEmptyIntegerContext(p *IntegerContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_integer
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) Signed_integer() ISigned_integerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISigned_integerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISigned_integerContext)
}

func (s *IntegerContext) Unsign_integer() IUnsign_integerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnsign_integerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnsign_integerContext)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Integer() (localctx IIntegerContext) {
	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, STParserRULE_integer)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case STParserPLUS, STParserMINUS:
		{
			p.SetState(157)
			p.Signed_integer()
		}

	case STParserInteger, STParserHex_Int:
		{
			p.SetState(158)
			p.Unsign_integer()
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

// ISigned_integerContext is an interface to support dynamic dispatch.
type ISigned_integerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSign returns the sign token.
	GetSign() antlr.Token

	// SetSign sets the sign token.
	SetSign(antlr.Token)

	// Getter signatures
	Unsign_integer() IUnsign_integerContext
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode

	// IsSigned_integerContext differentiates from other interfaces.
	IsSigned_integerContext()
}

type Signed_integerContext struct {
	*CustomContext
	parser antlr.Parser
	sign   antlr.Token
}

func NewEmptySigned_integerContext() *Signed_integerContext {
	var p = new(Signed_integerContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_signed_integer
	return p
}

func InitEmptySigned_integerContext(p *Signed_integerContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_signed_integer
}

func (*Signed_integerContext) IsSigned_integerContext() {}

func NewSigned_integerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Signed_integerContext {
	var p = new(Signed_integerContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_signed_integer

	return p
}

func (s *Signed_integerContext) GetParser() antlr.Parser { return s.parser }

func (s *Signed_integerContext) GetSign() antlr.Token { return s.sign }

func (s *Signed_integerContext) SetSign(v antlr.Token) { s.sign = v }

func (s *Signed_integerContext) Unsign_integer() IUnsign_integerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnsign_integerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnsign_integerContext)
}

func (s *Signed_integerContext) PLUS() antlr.TerminalNode {
	return s.GetToken(STParserPLUS, 0)
}

func (s *Signed_integerContext) MINUS() antlr.TerminalNode {
	return s.GetToken(STParserMINUS, 0)
}

func (s *Signed_integerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Signed_integerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Signed_integerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitSigned_integer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Signed_integer() (localctx ISigned_integerContext) {
	localctx = NewSigned_integerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, STParserRULE_signed_integer)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Signed_integerContext).sign = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == STParserPLUS || _la == STParserMINUS) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Signed_integerContext).sign = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(162)
		p.Unsign_integer()
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

// IUnsign_integerContext is an interface to support dynamic dispatch.
type IUnsign_integerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Integer() antlr.TerminalNode
	Hex_Int() antlr.TerminalNode

	// IsUnsign_integerContext differentiates from other interfaces.
	IsUnsign_integerContext()
}

type Unsign_integerContext struct {
	*CustomContext
	parser antlr.Parser
}

func NewEmptyUnsign_integerContext() *Unsign_integerContext {
	var p = new(Unsign_integerContext)
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_unsign_integer
	return p
}

func InitEmptyUnsign_integerContext(p *Unsign_integerContext) {
	p.CustomContext = NewCustomContext(nil, -1) // Jim super
	p.RuleIndex = STParserRULE_unsign_integer
}

func (*Unsign_integerContext) IsUnsign_integerContext() {}

func NewUnsign_integerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unsign_integerContext {
	var p = new(Unsign_integerContext)

	p.CustomContext = NewCustomContext(parent, invokingState)
	p.parser = parser
	p.RuleIndex = STParserRULE_unsign_integer

	return p
}

func (s *Unsign_integerContext) GetParser() antlr.Parser { return s.parser }

func (s *Unsign_integerContext) Integer() antlr.TerminalNode {
	return s.GetToken(STParserInteger, 0)
}

func (s *Unsign_integerContext) Hex_Int() antlr.TerminalNode {
	return s.GetToken(STParserHex_Int, 0)
}

func (s *Unsign_integerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unsign_integerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unsign_integerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STVisitor:
		return t.VisitUnsign_integer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Unsign_integer() (localctx IUnsign_integerContext) {
	localctx = NewUnsign_integerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, STParserRULE_unsign_integer)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		_la = p.GetTokenStream().LA(1)

		if !(_la == STParserInteger || _la == STParserHex_Int) {
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

func (p *STParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 13:
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
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
