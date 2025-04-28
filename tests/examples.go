package main

import (
	_ "embed"
	"log"

	"github.com/antlr4-go/antlr/v4"

	"github.com/slonegd/go-st"
	parser "github.com/slonegd/go-st/antlr"
	external_parser "github.com/slonegd/go-st/antlr/external"
)

var (
	//go:embed 001_simpliest.st
	simpliest string
	//go:embed 002_arithmetic.st
	arithmetic string
)

func main() {
	_ = simpliest
	example := arithmetic
	st_tokens(example)
	external_tokens(example)
	listener(example)
	tree(example)
	program(example)
}

func st_tokens(example string) {
	log.Printf("\n\n\t\t st_tokens")
	// Setup the input
	is := antlr.NewInputStream(example)

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

func external_tokens(example string) {
	log.Printf("\n\n\t\t external_tokens")
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

func listener(example string) {
	log.Printf("\n\n\t\t listener")
	// Setup the input
	is := antlr.NewInputStream(example)

	// Create the Lexer
	lexer := parser.NewSTLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewSTParser(stream)

	// Finally parse the expression (by walking the tree)
	antlr.ParseTreeWalkerDefault.Walk(&Listener{}, p.Program())
}

func tree(example string) {
	log.Printf("\n\n\t\t tree")
	// Setup the input
	is := antlr.NewInputStream(example)

	// Create the Lexer
	lexer := parser.NewSTLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewSTParser(stream)

	root := p.Program()
	log.Printf("root node %T %+v", root, root.GetIdentifier())
	printChildren(root, "")
}

func program(example string) {
	log.Printf("\n\n\t\t program")
	program := st.NewProgram(example)
	program.Run()
	log.Printf("results: %+v", program.Variables)
	program.Run()
	log.Printf("results: %+v", program.Variables)
}

func printChildren(node antlr.Tree, prefix string) {
	prefix += "  "
	for _, child := range node.GetChildren() {
		log.Printf(prefix+"%T %v", child, child.GetPayload())
		printChildren(child, prefix)
	}
}

type Listener struct{}

var _ parser.STListener = (*Listener)(nil)

func (*Listener) EnterAssignment_statement(c *parser.Assignment_statementContext) {
	log.Printf("%+v", c)
}
func (*Listener) EnterConstant(c *parser.ConstantContext)               { log.Printf("%+v", c) }
func (*Listener) EnterBinaryPlusExpr(c *parser.BinaryPlusExprContext)   { log.Printf("%+v", c) }
func (*Listener) EnterBinaryPowerExpr(c *parser.BinaryPowerExprContext) { log.Printf("%+v", c) }
func (*Listener) EnterParenExpr(c *parser.ParenExprContext)             { log.Printf("%+v", c) }
func (*Listener) EnterVariable(c *parser.VariableContext)               { log.Printf("%+v", c) }
func (*Listener) EnterEveryRule(ctx antlr.ParserRuleContext)            { log.Printf("%+v", ctx) }
func (*Listener) EnterExpression(c *parser.ExpressionContext)           { log.Printf("%+v", c) }
func (*Listener) EnterNumber(c *parser.NumberContext)                   { log.Printf("%+v", c) }
func (*Listener) EnterProgram(c *parser.ProgramContext)                 { log.Printf("%+v", c) }
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
func (*Listener) ExitBinaryPlusExpr(c *parser.BinaryPlusExprContext)   { log.Printf("%+v", c) }
func (*Listener) ExitBinaryPowerExpr(c *parser.BinaryPowerExprContext) { log.Printf("%+v", c) }
func (*Listener) ExitParenExpr(c *parser.ParenExprContext)             { log.Printf("%+v", c) }
func (*Listener) ExitVariable(c *parser.VariableContext)               { log.Printf("%+v", c) }
func (*Listener) ExitEveryRule(ctx antlr.ParserRuleContext)            { log.Printf("%+v", ctx) }
func (*Listener) ExitExpression(c *parser.ExpressionContext)           { log.Printf("%+v", c) }
func (*Listener) ExitNumber(c *parser.NumberContext)                   { log.Printf("%+v", c) }
func (*Listener) ExitProgram(c *parser.ProgramContext)                 { log.Printf("%+v", c) }
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
