package main

import (
	_ "embed"
	"log"

	"github.com/antlr4-go/antlr/v4"

	"github.com/slonegd/go-st"
	parser "github.com/slonegd/go-st/antlr"
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
	//go:embed 005_cast.st
	cast string
	//go:embed 006_floats.st
	floats string
	//go:embed 007_while.st
	while string
)

func main() {
	_ = simpliest
	// example := arithmetic
	_ = iftest
	_ = parseErr
	example := iftest
	// tree(example)
	program(example)
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
	log.Printf("root node %T %+v", root, root.GetId())
	printChildren(root, "")
}

func program(example string) {
	log.Printf("\n\n\t\t program")
	program, _ := st.NewProgram(example)
	program.ExecuteDebug()
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
