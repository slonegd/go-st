package ast

import (
	"fmt"
	"runtime/debug"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/slonegd/go-st/antlr"
	"github.com/slonegd/go-st/ops"
	"github.com/slonegd/go-st/source"
	"github.com/slonegd/go-st/types"
)

type (
	AST struct {
		source   source.Source
		vars     map[string]types.Variable
		programm *ops.Statement
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
		source: source.Source(script),
		vars:   map[string]types.Variable{},
	}
	visitor := &visitor{AST: x, Placeholders: ops.NewPlaceholders()}
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

	ctx := p.Program().(*parser.ProgramContext)
	ctx.Set(log{}, source.Source(script))
	x.programm = ctx.Accept(visitor).(*ops.Statement)

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
	_, s := x.programm.Do(x.programm)
	for s != nil {
		_, s = s.Do(s)
	}
}

func (x *AST) ExecuteDebug() {
	_, s := x.programm.DoDebug(x.programm)
	for s != nil {
		_, s = s.DoDebug(s)
	}
}

type log struct{}

func (log) Debug(f string, a ...any) { fmt.Printf(f, a...) }
