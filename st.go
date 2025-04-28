package st

import (
	"github.com/antlr4-go/antlr/v4"
	parser "github.com/slonegd/go-st/antlr"
)

type Program struct {
	Variables      map[string]int
	actions        []func()
	variablePrefix string
	stack          []int
}

var _ parser.STListener = (*Program)(nil)

func NewProgram(script string) *Program {
	result := &Program{
		Variables: make(map[string]int),
	}
	// Setup the input
	is := antlr.NewInputStream(script)

	// Create the Lexer
	lexer := parser.NewSTLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewSTParser(stream)

	// Finally parse the expression (by walking the tree)
	antlr.ParseTreeWalkerDefault.Walk(result, p.Program())
	return result
}

func (x *Program) Run() {
	for _, action := range x.actions {
		action()
	}
}
