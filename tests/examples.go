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
	//go:embed 003_iftest.st
	iftest string
	//go:embed 004_parse_err.st
	parseErr string
	//go:embed 005_implicit_int_cast.st
	implicit_int_cast string
)

func main() {
	_ = simpliest
	_ = arithmetic
	_ = iftest
	_ = parseErr
	example := iftest
	// st_tokens(example)
	// external_tokens(example)
	// tree(example)
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
	program, _ := st.NewProgram(example)
	program.Execute()
	log.Printf("results: %+v", program.GetVars())
	log.Printf("results: %+v", program.GetVars())
}

func printChildren(node antlr.Tree, prefix string) {
	prefix += "  "
	for _, child := range node.GetChildren() {
		log.Printf(prefix+"%T %v", child, child.GetPayload())
		printChildren(child, prefix)
	}
}
