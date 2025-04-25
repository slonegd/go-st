// Code generated from ./g4/st.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // st

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

type stParser struct {
	*antlr.BaseParser
}

var StParserStaticData struct {
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
	staticData := &StParserStaticData
	staticData.LiteralNames = []string{
		"", "'PROGRAM'", "'END_PROGRAM'", "'VAR'", "'END_VAR'", "':'", "';'",
		"'INT'", "':='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "Integer", "ID", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"prorgamm", "var_declaration_blocks", "var_declaration_block", "var_declaration",
		"type_name", "statement_list", "statement", "assignment_statement",
		"expression", "constant", "number",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 11, 70, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 30, 8, 1, 10, 1, 12,
		1, 33, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3,
		44, 8, 3, 10, 3, 12, 3, 47, 9, 3, 1, 4, 1, 4, 1, 5, 5, 5, 52, 8, 5, 10,
		5, 12, 5, 55, 9, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 0, 0, 11, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 0, 0, 61, 0, 22, 1, 0, 0, 0, 2, 31, 1, 0, 0, 0, 4, 34, 1, 0,
		0, 0, 6, 45, 1, 0, 0, 0, 8, 48, 1, 0, 0, 0, 10, 53, 1, 0, 0, 0, 12, 56,
		1, 0, 0, 0, 14, 59, 1, 0, 0, 0, 16, 63, 1, 0, 0, 0, 18, 65, 1, 0, 0, 0,
		20, 67, 1, 0, 0, 0, 22, 23, 5, 1, 0, 0, 23, 24, 5, 10, 0, 0, 24, 25, 3,
		2, 1, 0, 25, 26, 3, 10, 5, 0, 26, 27, 5, 2, 0, 0, 27, 1, 1, 0, 0, 0, 28,
		30, 3, 4, 2, 0, 29, 28, 1, 0, 0, 0, 30, 33, 1, 0, 0, 0, 31, 29, 1, 0, 0,
		0, 31, 32, 1, 0, 0, 0, 32, 3, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 34, 35, 5,
		3, 0, 0, 35, 36, 3, 6, 3, 0, 36, 37, 5, 4, 0, 0, 37, 5, 1, 0, 0, 0, 38,
		39, 5, 10, 0, 0, 39, 40, 5, 5, 0, 0, 40, 41, 3, 8, 4, 0, 41, 42, 5, 6,
		0, 0, 42, 44, 1, 0, 0, 0, 43, 38, 1, 0, 0, 0, 44, 47, 1, 0, 0, 0, 45, 43,
		1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 7, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0,
		48, 49, 5, 7, 0, 0, 49, 9, 1, 0, 0, 0, 50, 52, 3, 12, 6, 0, 51, 50, 1,
		0, 0, 0, 52, 55, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54,
		11, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 56, 57, 3, 14, 7, 0, 57, 58, 5, 6,
		0, 0, 58, 13, 1, 0, 0, 0, 59, 60, 5, 10, 0, 0, 60, 61, 5, 8, 0, 0, 61,
		62, 3, 16, 8, 0, 62, 15, 1, 0, 0, 0, 63, 64, 3, 18, 9, 0, 64, 17, 1, 0,
		0, 0, 65, 66, 3, 20, 10, 0, 66, 19, 1, 0, 0, 0, 67, 68, 5, 9, 0, 0, 68,
		21, 1, 0, 0, 0, 3, 31, 45, 53,
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

// stParserInit initializes any static state used to implement stParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewstParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func StParserInit() {
	staticData := &StParserStaticData
	staticData.once.Do(stParserInit)
}

// NewstParser produces a new parser instance for the optional input antlr.TokenStream.
func NewstParser(input antlr.TokenStream) *stParser {
	StParserInit()
	this := new(stParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &StParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "st.g4"

	return this
}

// stParser tokens.
const (
	stParserEOF        = antlr.TokenEOF
	stParserT__0       = 1
	stParserT__1       = 2
	stParserT__2       = 3
	stParserT__3       = 4
	stParserT__4       = 5
	stParserT__5       = 6
	stParserT__6       = 7
	stParserT__7       = 8
	stParserInteger    = 9
	stParserID         = 10
	stParserWHITESPACE = 11
)

// stParser rules.
const (
	stParserRULE_prorgamm               = 0
	stParserRULE_var_declaration_blocks = 1
	stParserRULE_var_declaration_block  = 2
	stParserRULE_var_declaration        = 3
	stParserRULE_type_name              = 4
	stParserRULE_statement_list         = 5
	stParserRULE_statement              = 6
	stParserRULE_assignment_statement   = 7
	stParserRULE_expression             = 8
	stParserRULE_constant               = 9
	stParserRULE_number                 = 10
)

// IProrgammContext is an interface to support dynamic dispatch.
type IProrgammContext interface {
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

	// IsProrgammContext differentiates from other interfaces.
	IsProrgammContext()
}

type ProrgammContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	identifier antlr.Token
}

func NewEmptyProrgammContext() *ProrgammContext {
	var p = new(ProrgammContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = stParserRULE_prorgamm
	return p
}

func InitEmptyProrgammContext(p *ProrgammContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = stParserRULE_prorgamm
}

func (*ProrgammContext) IsProrgammContext() {}

func NewProrgammContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProrgammContext {
	var p = new(ProrgammContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = stParserRULE_prorgamm

	return p
}

func (s *ProrgammContext) GetParser() antlr.Parser { return s.parser }

func (s *ProrgammContext) GetIdentifier() antlr.Token { return s.identifier }

func (s *ProrgammContext) SetIdentifier(v antlr.Token) { s.identifier = v }

func (s *ProrgammContext) Var_declaration_blocks() IVar_declaration_blocksContext {
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

func (s *ProrgammContext) Statement_list() IStatement_listContext {
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

func (s *ProrgammContext) ID() antlr.TerminalNode {
	return s.GetToken(stParserID, 0)
}

func (s *ProrgammContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProrgammContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProrgammContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stListener); ok {
		listenerT.EnterProrgamm(s)
	}
}

func (s *ProrgammContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stListener); ok {
		listenerT.ExitProrgamm(s)
	}
}

func (p *stParser) Prorgamm() (localctx IProrgammContext) {
	localctx = NewProrgammContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, stParserRULE_prorgamm)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.Match(stParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(23)

		var _m = p.Match(stParserID)

		localctx.(*ProrgammContext).identifier = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(24)
		p.Var_declaration_blocks()
	}
	{
		p.SetState(25)
		p.Statement_list()
	}
	{
		p.SetState(26)
		p.Match(stParserT__1)
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
	p.RuleIndex = stParserRULE_var_declaration_blocks
	return p
}

func InitEmptyVar_declaration_blocksContext(p *Var_declaration_blocksContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = stParserRULE_var_declaration_blocks
}

func (*Var_declaration_blocksContext) IsVar_declaration_blocksContext() {}

func NewVar_declaration_blocksContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_declaration_blocksContext {
	var p = new(Var_declaration_blocksContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = stParserRULE_var_declaration_blocks

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
	if listenerT, ok := listener.(stListener); ok {
		listenerT.EnterVar_declaration_blocks(s)
	}
}

func (s *Var_declaration_blocksContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stListener); ok {
		listenerT.ExitVar_declaration_blocks(s)
	}
}

func (p *stParser) Var_declaration_blocks() (localctx IVar_declaration_blocksContext) {
	localctx = NewVar_declaration_blocksContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, stParserRULE_var_declaration_blocks)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == stParserT__2 {
		{
			p.SetState(28)
			p.Var_declaration_block()
		}

		p.SetState(33)
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
	p.RuleIndex = stParserRULE_var_declaration_block
	return p
}

func InitEmptyVar_declaration_blockContext(p *Var_declaration_blockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = stParserRULE_var_declaration_block
}

func (*Var_declaration_blockContext) IsVar_declaration_blockContext() {}

func NewVar_declaration_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_declaration_blockContext {
	var p = new(Var_declaration_blockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = stParserRULE_var_declaration_block

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
	if listenerT, ok := listener.(stListener); ok {
		listenerT.EnterVar_declaration_block(s)
	}
}

func (s *Var_declaration_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stListener); ok {
		listenerT.ExitVar_declaration_block(s)
	}
}

func (p *stParser) Var_declaration_block() (localctx IVar_declaration_blockContext) {
	localctx = NewVar_declaration_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, stParserRULE_var_declaration_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.Match(stParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(35)
		p.Var_declaration()
	}
	{
		p.SetState(36)
		p.Match(stParserT__3)
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
	p.RuleIndex = stParserRULE_var_declaration
	return p
}

func InitEmptyVar_declarationContext(p *Var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = stParserRULE_var_declaration
}

func (*Var_declarationContext) IsVar_declarationContext() {}

func NewVar_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_declarationContext {
	var p = new(Var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = stParserRULE_var_declaration

	return p
}

func (s *Var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_declarationContext) GetIdentifier() antlr.Token { return s.identifier }

func (s *Var_declarationContext) SetIdentifier(v antlr.Token) { s.identifier = v }

func (s *Var_declarationContext) GetType_() IType_nameContext { return s.type_ }

func (s *Var_declarationContext) SetType_(v IType_nameContext) { s.type_ = v }

func (s *Var_declarationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(stParserID)
}

func (s *Var_declarationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(stParserID, i)
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
	if listenerT, ok := listener.(stListener); ok {
		listenerT.EnterVar_declaration(s)
	}
}

func (s *Var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stListener); ok {
		listenerT.ExitVar_declaration(s)
	}
}

func (p *stParser) Var_declaration() (localctx IVar_declarationContext) {
	localctx = NewVar_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, stParserRULE_var_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == stParserID {
		{
			p.SetState(38)

			var _m = p.Match(stParserID)

			localctx.(*Var_declarationContext).identifier = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.Match(stParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)

			var _x = p.Type_name()

			localctx.(*Var_declarationContext).type_ = _x
		}
		{
			p.SetState(41)
			p.Match(stParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
	p.RuleIndex = stParserRULE_type_name
	return p
}

func InitEmptyType_nameContext(p *Type_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = stParserRULE_type_name
}

func (*Type_nameContext) IsType_nameContext() {}

func NewType_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_nameContext {
	var p = new(Type_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = stParserRULE_type_name

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
	if listenerT, ok := listener.(stListener); ok {
		listenerT.EnterType_name(s)
	}
}

func (s *Type_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stListener); ok {
		listenerT.ExitType_name(s)
	}
}

func (p *stParser) Type_name() (localctx IType_nameContext) {
	localctx = NewType_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, stParserRULE_type_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(stParserT__6)
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
	p.RuleIndex = stParserRULE_statement_list
	return p
}

func InitEmptyStatement_listContext(p *Statement_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = stParserRULE_statement_list
}

func (*Statement_listContext) IsStatement_listContext() {}

func NewStatement_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Statement_listContext {
	var p = new(Statement_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = stParserRULE_statement_list

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
	if listenerT, ok := listener.(stListener); ok {
		listenerT.EnterStatement_list(s)
	}
}

func (s *Statement_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stListener); ok {
		listenerT.ExitStatement_list(s)
	}
}

func (p *stParser) Statement_list() (localctx IStatement_listContext) {
	localctx = NewStatement_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, stParserRULE_statement_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == stParserID {
		{
			p.SetState(50)
			p.Statement()
		}

		p.SetState(55)
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
	p.RuleIndex = stParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = stParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = stParserRULE_statement

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

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *stParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, stParserRULE_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Assignment_statement()
	}
	{
		p.SetState(57)
		p.Match(stParserT__5)
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
	p.RuleIndex = stParserRULE_assignment_statement
	return p
}

func InitEmptyAssignment_statementContext(p *Assignment_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = stParserRULE_assignment_statement
}

func (*Assignment_statementContext) IsAssignment_statementContext() {}

func NewAssignment_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assignment_statementContext {
	var p = new(Assignment_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = stParserRULE_assignment_statement

	return p
}

func (s *Assignment_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Assignment_statementContext) GetLeft() antlr.Token { return s.left }

func (s *Assignment_statementContext) SetLeft(v antlr.Token) { s.left = v }

func (s *Assignment_statementContext) GetRight() IExpressionContext { return s.right }

func (s *Assignment_statementContext) SetRight(v IExpressionContext) { s.right = v }

func (s *Assignment_statementContext) ID() antlr.TerminalNode {
	return s.GetToken(stParserID, 0)
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
	if listenerT, ok := listener.(stListener); ok {
		listenerT.EnterAssignment_statement(s)
	}
}

func (s *Assignment_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stListener); ok {
		listenerT.ExitAssignment_statement(s)
	}
}

func (p *stParser) Assignment_statement() (localctx IAssignment_statementContext) {
	localctx = NewAssignment_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, stParserRULE_assignment_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)

		var _m = p.Match(stParserID)

		localctx.(*Assignment_statementContext).left = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Match(stParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)

		var _x = p.Expression()

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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Constant() IConstantContext

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
	p.RuleIndex = stParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = stParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = stParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *stParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, stParserRULE_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Constant()
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

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Number() INumberContext

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = stParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = stParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = stParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

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

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *stParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, stParserRULE_constant)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Number()
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
	p.RuleIndex = stParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = stParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = stParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) Integer() antlr.TerminalNode {
	return s.GetToken(stParserInteger, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(stListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *stParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, stParserRULE_number)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(stParserInteger)
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
