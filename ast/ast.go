package ast

import (
	"fmt"
	"runtime/debug"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/slonegd/go-st/antlr"
	"github.com/slonegd/go-st/ops"
	"github.com/slonegd/go-st/types"
)

type (
	AST struct {
		source   Source
		vars     map[string]types.Variable
		programm ops.Statement
		err      error
	}
)

func Parse(script string) (x *AST, err error) {
	// Setup the input
	is := antlr.NewInputStream(script)

	// Create the Lexer
	lexer := parser.NewSTLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewSTParser(stream)

	x = &AST{
		source: Source(script),
		vars:   map[string]types.Variable{},
	}
	visitor := &visitor{AST: x}
	p.BaseRecognizer.AddErrorListener(visitor)

	// выход из обхода дерева сделан через панику, поэтому ловим
	defer func() {
		if r := recover(); r != nil {
			if x.err != nil {
				err = x.err
				return
			}
			err = fmt.Errorf("Panic: %v\n %s", r, debug.Stack())
		}
	}()

	x.programm = p.Program().Accept(visitor).(ops.Statement)

	err = x.err

	return x, err
}

func (x *AST) GetVar(name string) types.Variable {
	return x.vars[name]
}

func (x *AST) GetVars() map[string]types.Variable {
	return x.vars
}

func (x *AST) Execute() {
	x.programm.Do()
}
