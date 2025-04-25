package main

import (
	_ "embed"
	"log"

	"github.com/antlr4-go/antlr/v4"

	parser "github.com/slonegd/go-st/antlr"
	external_parser "github.com/slonegd/go-st/antlr/external"
)

//go:embed 001_simpliest.st
var simpliest string

func main() {
	simpliest_main()
	simpliest_external_main()
	simpliest_listener()
}

func simpliest_main() {
	log.Printf("\n\n\t\t simpliest_main")
	// Setup the input
	is := antlr.NewInputStream(simpliest)

	// Create the Lexer
	lexer := parser.NewSTLexer(is)

	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		log.Printf("%s (%q): %#v", lexer.SymbolicNames[t.GetTokenType()], t.GetText(), t)
	}
}

func simpliest_external_main() {
	log.Printf("\n\n\t\t simpliest_external_main")
	// Setup the input
	is := antlr.NewInputStream(simpliest)

	// Create the Lexer
	lexer := external_parser.NewST_lexer(is)

	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		log.Printf("%s (%q): %#v", lexer.SymbolicNames[t.GetTokenType()], t.GetText(), t)
	}
}

func simpliest_listener() {
	log.Printf("\n\n\t\t simpliest_listener")
	// Setup the input
	is := antlr.NewInputStream(simpliest)

	// Create the Lexer
	lexer := parser.NewSTLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewSTParser(stream)

	// Finally parse the expression (by walking the tree)
	antlr.ParseTreeWalkerDefault.Walk(&Listener{}, p.Prorgamm())
}

type Listener struct{}

var _ parser.STListener = (*Listener)(nil)

func (*Listener) EnterAssignment_statement(c *parser.Assignment_statementContext) {
	log.Printf("%+v", c)
}
func (*Listener) EnterConstant(c *parser.ConstantContext)               { log.Printf("%+v", c) }
func (*Listener) EnterEveryRule(ctx antlr.ParserRuleContext)            { log.Printf("%+v", ctx) }
func (*Listener) EnterExpression(c *parser.ExpressionContext)           { log.Printf("%+v", c) }
func (*Listener) EnterNumber(c *parser.NumberContext)                   { log.Printf("%+v", c) }
func (*Listener) EnterProrgamm(c *parser.ProrgammContext)               { log.Printf("%+v", c) }
func (*Listener) EnterStatement(c *parser.StatementContext)             { log.Printf("%+v", c) }
func (*Listener) EnterStatement_list(c *parser.Statement_listContext)   { log.Printf("%+v", c) }
func (*Listener) EnterType_name(c *parser.Type_nameContext)             { log.Printf("%+v", c) }
func (*Listener) EnterVar_declaration(c *parser.Var_declarationContext) { log.Printf("%+v", c) }
func (*Listener) EnterVar_declaration_block(c *parser.Var_declaration_blockContext) {
	log.Printf("%+v", c)
}
func (*Listener) EnterVar_declaration_blocks(c *parser.Var_declaration_blocksContext) {
	log.Printf("%+v", c)
}
func (*Listener) ExitAssignment_statement(c *parser.Assignment_statementContext) {
	log.Printf("%+v", c)
}
func (*Listener) ExitConstant(c *parser.ConstantContext)               { log.Printf("%+v", c) }
func (*Listener) ExitEveryRule(ctx antlr.ParserRuleContext)            { log.Printf("%+v", ctx) }
func (*Listener) ExitExpression(c *parser.ExpressionContext)           { log.Printf("%+v", c) }
func (*Listener) ExitNumber(c *parser.NumberContext)                   { log.Printf("%+v", c) }
func (*Listener) ExitProrgamm(c *parser.ProrgammContext)               { log.Printf("%+v", c) }
func (*Listener) ExitStatement(c *parser.StatementContext)             { log.Printf("%+v", c) }
func (*Listener) ExitStatement_list(c *parser.Statement_listContext)   { log.Printf("%+v", c) }
func (*Listener) ExitType_name(c *parser.Type_nameContext)             { log.Printf("%+v", c) }
func (*Listener) ExitVar_declaration(c *parser.Var_declarationContext) { log.Printf("%+v", c) }
func (*Listener) ExitVar_declaration_block(c *parser.Var_declaration_blockContext) {
	log.Printf("%+v", c)
}
func (*Listener) ExitVar_declaration_blocks(c *parser.Var_declaration_blocksContext) {
	log.Printf("%+v", c)
}
func (*Listener) VisitErrorNode(node antlr.ErrorNode)   { log.Printf("%+v", node) }
func (*Listener) VisitTerminal(node antlr.TerminalNode) { log.Printf("%+v", node) }
