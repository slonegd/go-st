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
		"", "'PROGRAM'", "'END_PROGRAM'", "'VAR'", "'END_VAR'", "':'", "';'",
		"'INT'", "':='", "'IF'", "'THEN'", "'ELSE'", "'END_IF'", "'('", "')'",
		"'*'", "'/'", "'MOD'", "'+'", "'-'", "'>'", "'>='", "'<'", "'<='", "'='",
		"'<>'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "Integer", "ID", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"program", "var_declaration_blocks", "var_declaration_block", "var_declaration",
		"type_name", "statement_list", "statement", "assignment_statement",
		"if_statement", "condition", "then_list", "else_list", "expression",
		"number",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 28, 119, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 1, 5, 1, 36, 8, 1, 10, 1, 12, 1, 39, 9, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 50, 8, 3, 10, 3, 12, 3, 53, 9,
		3, 1, 4, 1, 4, 1, 5, 5, 5, 58, 8, 5, 10, 5, 12, 5, 61, 9, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 3, 6, 68, 8, 6, 3, 6, 70, 8, 6, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 82, 8, 8, 1, 8, 1, 8, 3, 8,
		86, 8, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 101, 8, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 112, 8, 12, 10, 12, 12, 12,
		115, 9, 12, 1, 13, 1, 13, 1, 13, 0, 1, 24, 14, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 0, 3, 1, 0, 15, 17, 1, 0, 18, 19, 1, 0, 20, 25,
		116, 0, 28, 1, 0, 0, 0, 2, 37, 1, 0, 0, 0, 4, 40, 1, 0, 0, 0, 6, 51, 1,
		0, 0, 0, 8, 54, 1, 0, 0, 0, 10, 59, 1, 0, 0, 0, 12, 69, 1, 0, 0, 0, 14,
		71, 1, 0, 0, 0, 16, 75, 1, 0, 0, 0, 18, 87, 1, 0, 0, 0, 20, 89, 1, 0, 0,
		0, 22, 91, 1, 0, 0, 0, 24, 100, 1, 0, 0, 0, 26, 116, 1, 0, 0, 0, 28, 29,
		5, 1, 0, 0, 29, 30, 5, 27, 0, 0, 30, 31, 3, 2, 1, 0, 31, 32, 3, 10, 5,
		0, 32, 33, 5, 2, 0, 0, 33, 1, 1, 0, 0, 0, 34, 36, 3, 4, 2, 0, 35, 34, 1,
		0, 0, 0, 36, 39, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38,
		3, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 40, 41, 5, 3, 0, 0, 41, 42, 3, 6, 3,
		0, 42, 43, 5, 4, 0, 0, 43, 5, 1, 0, 0, 0, 44, 45, 5, 27, 0, 0, 45, 46,
		5, 5, 0, 0, 46, 47, 3, 8, 4, 0, 47, 48, 5, 6, 0, 0, 48, 50, 1, 0, 0, 0,
		49, 44, 1, 0, 0, 0, 50, 53, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 51, 52, 1,
		0, 0, 0, 52, 7, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 54, 55, 5, 7, 0, 0, 55,
		9, 1, 0, 0, 0, 56, 58, 3, 12, 6, 0, 57, 56, 1, 0, 0, 0, 58, 61, 1, 0, 0,
		0, 59, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 11, 1, 0, 0, 0, 61, 59,
		1, 0, 0, 0, 62, 63, 3, 14, 7, 0, 63, 64, 5, 6, 0, 0, 64, 70, 1, 0, 0, 0,
		65, 67, 3, 16, 8, 0, 66, 68, 5, 6, 0, 0, 67, 66, 1, 0, 0, 0, 67, 68, 1,
		0, 0, 0, 68, 70, 1, 0, 0, 0, 69, 62, 1, 0, 0, 0, 69, 65, 1, 0, 0, 0, 70,
		13, 1, 0, 0, 0, 71, 72, 5, 27, 0, 0, 72, 73, 5, 8, 0, 0, 73, 74, 3, 24,
		12, 0, 74, 15, 1, 0, 0, 0, 75, 76, 5, 9, 0, 0, 76, 77, 3, 18, 9, 0, 77,
		78, 5, 10, 0, 0, 78, 81, 3, 20, 10, 0, 79, 80, 5, 11, 0, 0, 80, 82, 3,
		22, 11, 0, 81, 79, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0,
		83, 85, 5, 12, 0, 0, 84, 86, 5, 6, 0, 0, 85, 84, 1, 0, 0, 0, 85, 86, 1,
		0, 0, 0, 86, 17, 1, 0, 0, 0, 87, 88, 3, 24, 12, 0, 88, 19, 1, 0, 0, 0,
		89, 90, 3, 10, 5, 0, 90, 21, 1, 0, 0, 0, 91, 92, 3, 10, 5, 0, 92, 23, 1,
		0, 0, 0, 93, 94, 6, 12, -1, 0, 94, 101, 3, 26, 13, 0, 95, 101, 5, 27, 0,
		0, 96, 97, 5, 13, 0, 0, 97, 98, 3, 24, 12, 0, 98, 99, 5, 14, 0, 0, 99,
		101, 1, 0, 0, 0, 100, 93, 1, 0, 0, 0, 100, 95, 1, 0, 0, 0, 100, 96, 1,
		0, 0, 0, 101, 113, 1, 0, 0, 0, 102, 103, 10, 3, 0, 0, 103, 104, 7, 0, 0,
		0, 104, 112, 3, 24, 12, 4, 105, 106, 10, 2, 0, 0, 106, 107, 7, 1, 0, 0,
		107, 112, 3, 24, 12, 3, 108, 109, 10, 1, 0, 0, 109, 110, 7, 2, 0, 0, 110,
		112, 3, 24, 12, 2, 111, 102, 1, 0, 0, 0, 111, 105, 1, 0, 0, 0, 111, 108,
		1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 114, 1, 0,
		0, 0, 114, 25, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 116, 117, 5, 26, 0, 0,
		117, 27, 1, 0, 0, 0, 10, 37, 51, 59, 67, 69, 81, 85, 100, 111, 113,
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
	STParserInteger    = 26
	STParserID         = 27
	STParserWHITESPACE = 28
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
	STParserRULE_condition              = 9
	STParserRULE_then_list              = 10
	STParserRULE_else_list              = 11
	STParserRULE_expression             = 12
	STParserRULE_number                 = 13
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
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	identifier antlr.Token
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *STParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, STParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.Match(STParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(29)

		var _m = p.Match(STParserID)

		localctx.(*ProgramContext).identifier = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(30)
		p.Var_declaration_blocks()
	}
	{
		p.SetState(31)
		p.Statement_list()
	}
	{
		p.SetState(32)
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_declaration_blocksContext() *Var_declaration_blocksContext {
	var p = new(Var_declaration_blocksContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_var_declaration_blocks
	return p
}

func InitEmptyVar_declaration_blocksContext(p *Var_declaration_blocksContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_var_declaration_blocks
}

func (*Var_declaration_blocksContext) IsVar_declaration_blocksContext() {}

func NewVar_declaration_blocksContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_declaration_blocksContext {
	var p = new(Var_declaration_blocksContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

func (s *Var_declaration_blocksContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.EnterVar_declaration_blocks(s)
	}
}

func (s *Var_declaration_blocksContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.ExitVar_declaration_blocks(s)
	}
}

func (p *STParser) Var_declaration_blocks() (localctx IVar_declaration_blocksContext) {
	localctx = NewVar_declaration_blocksContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, STParserRULE_var_declaration_blocks)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == STParserT__2 {
		{
			p.SetState(34)
			p.Var_declaration_block()
		}

		p.SetState(39)
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
	Var_declaration() IVar_declarationContext

	// IsVar_declaration_blockContext differentiates from other interfaces.
	IsVar_declaration_blockContext()
}

type Var_declaration_blockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_declaration_blockContext() *Var_declaration_blockContext {
	var p = new(Var_declaration_blockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_var_declaration_block
	return p
}

func InitEmptyVar_declaration_blockContext(p *Var_declaration_blockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_var_declaration_block
}

func (*Var_declaration_blockContext) IsVar_declaration_blockContext() {}

func NewVar_declaration_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_declaration_blockContext {
	var p = new(Var_declaration_blockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_var_declaration_block

	return p
}

func (s *Var_declaration_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_declaration_blockContext) Var_declaration() IVar_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
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

func (s *Var_declaration_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.EnterVar_declaration_block(s)
	}
}

func (s *Var_declaration_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.ExitVar_declaration_block(s)
	}
}

func (p *STParser) Var_declaration_block() (localctx IVar_declaration_blockContext) {
	localctx = NewVar_declaration_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, STParserRULE_var_declaration_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Match(STParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(41)
		p.Var_declaration()
	}
	{
		p.SetState(42)
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

	// SetType_ sets the type_ rule contexts.
	SetType_(IType_nameContext)

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllType_name() []IType_nameContext
	Type_name(i int) IType_nameContext

	// IsVar_declarationContext differentiates from other interfaces.
	IsVar_declarationContext()
}

type Var_declarationContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	identifier antlr.Token
	type_      IType_nameContext
}

func NewEmptyVar_declarationContext() *Var_declarationContext {
	var p = new(Var_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_var_declaration
	return p
}

func InitEmptyVar_declarationContext(p *Var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_var_declaration
}

func (*Var_declarationContext) IsVar_declarationContext() {}

func NewVar_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_declarationContext {
	var p = new(Var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_var_declaration

	return p
}

func (s *Var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_declarationContext) GetIdentifier() antlr.Token { return s.identifier }

func (s *Var_declarationContext) SetIdentifier(v antlr.Token) { s.identifier = v }

func (s *Var_declarationContext) GetType_() IType_nameContext { return s.type_ }

func (s *Var_declarationContext) SetType_(v IType_nameContext) { s.type_ = v }

func (s *Var_declarationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(STParserID)
}

func (s *Var_declarationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(STParserID, i)
}

func (s *Var_declarationContext) AllType_name() []IType_nameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IType_nameContext); ok {
			len++
		}
	}

	tst := make([]IType_nameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IType_nameContext); ok {
			tst[i] = t.(IType_nameContext)
			i++
		}
	}

	return tst
}

func (s *Var_declarationContext) Type_name(i int) IType_nameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_nameContext); ok {
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

	return t.(IType_nameContext)
}

func (s *Var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.EnterVar_declaration(s)
	}
}

func (s *Var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.ExitVar_declaration(s)
	}
}

func (p *STParser) Var_declaration() (localctx IVar_declarationContext) {
	localctx = NewVar_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, STParserRULE_var_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == STParserID {
		{
			p.SetState(44)

			var _m = p.Match(STParserID)

			localctx.(*Var_declarationContext).identifier = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)
			p.Match(STParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)

			var _x = p.Type_name()

			localctx.(*Var_declarationContext).type_ = _x
		}
		{
			p.SetState(47)
			p.Match(STParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(53)
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

// IType_nameContext is an interface to support dynamic dispatch.
type IType_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsType_nameContext differentiates from other interfaces.
	IsType_nameContext()
}

type Type_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_nameContext() *Type_nameContext {
	var p = new(Type_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_type_name
	return p
}

func InitEmptyType_nameContext(p *Type_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_type_name
}

func (*Type_nameContext) IsType_nameContext() {}

func NewType_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_nameContext {
	var p = new(Type_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

func (s *Type_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.EnterType_name(s)
	}
}

func (s *Type_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.ExitType_name(s)
	}
}

func (p *STParser) Type_name() (localctx IType_nameContext) {
	localctx = NewType_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, STParserRULE_type_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatement_listContext() *Statement_listContext {
	var p = new(Statement_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_statement_list
	return p
}

func InitEmptyStatement_listContext(p *Statement_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_statement_list
}

func (*Statement_listContext) IsStatement_listContext() {}

func NewStatement_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Statement_listContext {
	var p = new(Statement_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

func (s *Statement_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.EnterStatement_list(s)
	}
}

func (s *Statement_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.ExitStatement_list(s)
	}
}

func (p *STParser) Statement_list() (localctx IStatement_listContext) {
	localctx = NewStatement_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, STParserRULE_statement_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == STParserT__8 || _la == STParserID {
		{
			p.SetState(56)
			p.Statement()
		}

		p.SetState(61)
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

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *STParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, STParserRULE_statement)
	var _la int

	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case STParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.Assignment_statement()
		}
		{
			p.SetState(63)
			p.Match(STParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case STParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.If_statement()
		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == STParserT__5 {
			{
				p.SetState(66)
				p.Match(STParserT__5)
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
	left   antlr.Token
	right  IExpressionContext
}

func NewEmptyAssignment_statementContext() *Assignment_statementContext {
	var p = new(Assignment_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_assignment_statement
	return p
}

func InitEmptyAssignment_statementContext(p *Assignment_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_assignment_statement
}

func (*Assignment_statementContext) IsAssignment_statementContext() {}

func NewAssignment_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assignment_statementContext {
	var p = new(Assignment_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

func (s *Assignment_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.EnterAssignment_statement(s)
	}
}

func (s *Assignment_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.ExitAssignment_statement(s)
	}
}

func (p *STParser) Assignment_statement() (localctx IAssignment_statementContext) {
	localctx = NewAssignment_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, STParserRULE_assignment_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)

		var _m = p.Match(STParserID)

		localctx.(*Assignment_statementContext).left = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(72)
		p.Match(STParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)

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

	// GetCond returns the cond rule contexts.
	GetCond() IConditionContext

	// GetThenlist returns the thenlist rule contexts.
	GetThenlist() IThen_listContext

	// GetElselist returns the elselist rule contexts.
	GetElselist() IElse_listContext

	// SetCond sets the cond rule contexts.
	SetCond(IConditionContext)

	// SetThenlist sets the thenlist rule contexts.
	SetThenlist(IThen_listContext)

	// SetElselist sets the elselist rule contexts.
	SetElselist(IElse_listContext)

	// Getter signatures
	Condition() IConditionContext
	Then_list() IThen_listContext
	Else_list() IElse_listContext

	// IsIf_statementContext differentiates from other interfaces.
	IsIf_statementContext()
}

type If_statementContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	cond     IConditionContext
	thenlist IThen_listContext
	elselist IElse_listContext
}

func NewEmptyIf_statementContext() *If_statementContext {
	var p = new(If_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_if_statement
	return p
}

func InitEmptyIf_statementContext(p *If_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_if_statement
}

func (*If_statementContext) IsIf_statementContext() {}

func NewIf_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_statementContext {
	var p = new(If_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_if_statement

	return p
}

func (s *If_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *If_statementContext) GetCond() IConditionContext { return s.cond }

func (s *If_statementContext) GetThenlist() IThen_listContext { return s.thenlist }

func (s *If_statementContext) GetElselist() IElse_listContext { return s.elselist }

func (s *If_statementContext) SetCond(v IConditionContext) { s.cond = v }

func (s *If_statementContext) SetThenlist(v IThen_listContext) { s.thenlist = v }

func (s *If_statementContext) SetElselist(v IElse_listContext) { s.elselist = v }

func (s *If_statementContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *If_statementContext) Then_list() IThen_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IThen_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
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

func (s *If_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.EnterIf_statement(s)
	}
}

func (s *If_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.ExitIf_statement(s)
	}
}

func (p *STParser) If_statement() (localctx IIf_statementContext) {
	localctx = NewIf_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, STParserRULE_if_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(STParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)

		var _x = p.Condition()

		localctx.(*If_statementContext).cond = _x
	}
	{
		p.SetState(77)
		p.Match(STParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(78)

		var _x = p.Then_list()

		localctx.(*If_statementContext).thenlist = _x
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == STParserT__10 {
		{
			p.SetState(79)
			p.Match(STParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)

			var _x = p.Else_list()

			localctx.(*If_statementContext).elselist = _x
		}

	}
	{
		p.SetState(83)
		p.Match(STParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(84)
			p.Match(STParserT__5)
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *STParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, STParserRULE_condition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThen_listContext() *Then_listContext {
	var p = new(Then_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_then_list
	return p
}

func InitEmptyThen_listContext(p *Then_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_then_list
}

func (*Then_listContext) IsThen_listContext() {}

func NewThen_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Then_listContext {
	var p = new(Then_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

func (s *Then_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.EnterThen_list(s)
	}
}

func (s *Then_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.ExitThen_list(s)
	}
}

func (p *STParser) Then_list() (localctx IThen_listContext) {
	localctx = NewThen_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, STParserRULE_then_list)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElse_listContext() *Else_listContext {
	var p = new(Else_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_else_list
	return p
}

func InitEmptyElse_listContext(p *Else_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_else_list
}

func (*Else_listContext) IsElse_listContext() {}

func NewElse_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_listContext {
	var p = new(Else_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

func (s *Else_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.EnterElse_list(s)
	}
}

func (s *Else_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.ExitElse_list(s)
	}
}

func (p *STParser) Else_list() (localctx IElse_listContext) {
	localctx = NewElse_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, STParserRULE_else_list)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

func (s *BinaryCompareExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.EnterBinaryCompareExpr(s)
	}
}

func (s *BinaryCompareExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.ExitBinaryCompareExpr(s)
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

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.ExitConstant(s)
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

func (s *BinaryPowerExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.EnterBinaryPowerExpr(s)
	}
}

func (s *BinaryPowerExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.ExitBinaryPowerExpr(s)
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

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.ExitVariable(s)
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

func (s *BinaryPlusExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.EnterBinaryPlusExpr(s)
	}
}

func (s *BinaryPlusExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.ExitBinaryPlusExpr(s)
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

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.ExitParenExpr(s)
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
	_startState := 24
	p.EnterRecursionRule(localctx, 24, STParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case STParserInteger:
		localctx = NewConstantContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(94)
			p.Number()
		}

	case STParserID:
		localctx = NewVariableContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(95)
			p.Match(STParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case STParserT__12:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(96)
			p.Match(STParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(97)

			var _x = p.expression(0)

			localctx.(*ParenExprContext).sub = _x
		}
		{
			p.SetState(98)
			p.Match(STParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(111)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryPowerExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryPowerExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, STParserRULE_expression)
				p.SetState(102)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(103)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryPowerExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&229376) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryPowerExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(104)

					var _x = p.expression(4)

					localctx.(*BinaryPowerExprContext).right = _x
				}

			case 2:
				localctx = NewBinaryPlusExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryPlusExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, STParserRULE_expression)
				p.SetState(105)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(106)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryPlusExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == STParserT__17 || _la == STParserT__18) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryPlusExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(107)

					var _x = p.expression(3)

					localctx.(*BinaryPlusExprContext).right = _x
				}

			case 3:
				localctx = NewBinaryCompareExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryCompareExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, STParserRULE_expression)
				p.SetState(108)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(109)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryCompareExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&66060288) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryCompareExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(110)

					var _x = p.expression(2)

					localctx.(*BinaryCompareExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
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
	Integer() antlr.TerminalNode

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = STParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) Integer() antlr.TerminalNode {
	return s.GetToken(STParserInteger, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *STParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, STParserRULE_number)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(STParserInteger)
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
	case 12:
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
