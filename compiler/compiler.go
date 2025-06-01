package compiler

import (
	"errors"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/slonegd/go-st/antlr"
	"github.com/slonegd/go-st/stack"
	"github.com/slonegd/go-st/variant"
	"github.com/slonegd/go-st/vm"
)

type (
	// создаёт стековую виртуальную машину с обобщёнными операторами (без семантики типов)
	// TODO семантику типов сделать в оптимизаторе, где обощённые операторы будут заменться на типовые
	Compiler struct {
		*vm.VM
		varIndexes map[string]int
		// для простановки jump интсрукций
		conditionJumpIndexes stack.Stack[int]
		thenJumpIndexes      stack.Stack[[]int]
		source               Source
		err                  error
	}
	listener struct { // чтоб не засорять интерфейс Compiler методами от antlr
		*Compiler
	}
)

var (
	_ parser.STListener   = (*listener)(nil)
	_ antlr.ErrorListener = (*listener)(nil)
)

func New() *Compiler {
	return &Compiler{
		VM:         vm.New(),
		varIndexes: make(map[string]int),
	}
}

func (x *Compiler) Compile(script string) error {
	x.source = Source(script)

	// Setup the input
	is := antlr.NewInputStream(script)

	// Create the Lexer
	lexer := parser.NewSTLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewSTParser(stream)

	listener := listener{Compiler: x}
	p.BaseRecognizer.AddErrorListener(listener)

	// Finally parse the expression (by walking the tree)
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Program())
	if p.SynErr != nil && p.SynErr.GetMessage() != "" {
		return errors.New(p.SynErr.GetMessage())
	}

	return x.err
}

func (x *Compiler) GetVar(name string) variant.Variant {
	i := x.varIndexes[name]
	if i >= len(x.Vars) {
		return &variant.Any{}
	}
	return x.Vars[i]
}

func (x *Compiler) Variables() map[string]variant.Variant {
	r := make(map[string]variant.Variant, len(x.varIndexes))
	for name, i := range x.varIndexes {
		r[name] = x.Vars[i]
	}
	return r
}
