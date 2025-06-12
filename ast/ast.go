package ast

import (
	"github.com/antlr4-go/antlr/v4"
	parser "github.com/slonegd/go-st/antlr"
	"github.com/slonegd/go-st/variant"
)

type (
	AST struct {
		source   Source
		vars     map[string]variant.Variant
		programm Statement
		err      error
	}
)

func Parse(script string) (*AST, error) {
	// Setup the input
	is := antlr.NewInputStream(script)

	// Create the Lexer
	lexer := parser.NewSTLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewSTParser(stream)

	x := &AST{
		source: Source(script),
		vars:   map[string]variant.Variant{},
	}
	visitor := &visitor{AST: x}
	p.BaseRecognizer.AddErrorListener(visitor)
	x.programm = p.Program().Accept(visitor).(Statement)

	return x, x.err
}

func (x *AST) GetVar(name string) variant.Variant {
	return x.vars[name]
}

func (x *AST) GetVars() map[string]variant.Variant {
	return x.vars
}

func (x *AST) Execute() {
	x.programm.do()
}
