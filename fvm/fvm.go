package fvm

import (
	"errors"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/slonegd/go-st/antlr"
	"github.com/slonegd/go-st/stack"
	"github.com/slonegd/go-st/variant"
)

type (
	FVM struct {
		vars    map[string]variant.Variant
		runtime stack.Stack[Step]
	}
	Step struct {
		do    func() any
		descr string
	}
)

func NewProgram(script string) (*FVM, error) {
	// x.source = Source(script)

	// Setup the input
	is := antlr.NewInputStream(script)

	// Create the Lexer
	lexer := parser.NewSTLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewSTParser(stream)

	x := &FVM{vars: map[string]variant.Variant{}}
	// listener := listener{Compiler: x}
	// p.BaseRecognizer.AddErrorListener(listener)

	// Finally parse the expression (by walking the tree)
	antlr.ParseTreeWalkerDefault.Walk(x, p.Program())
	if p.SynErr != nil && p.SynErr.GetMessage() != "" {
		return nil, errors.New(p.SynErr.GetMessage())
	}

	return x, nil // x.err
}

func (x *FVM) GetVar(name string) variant.Variant {
	return x.vars[name]
}

func (x *FVM) Execute() {
	for _, s := range x.runtime {
		s.do()
	}
}

func (s *Step) WithDescription(v string) *Step { s.descr = v; return s }

func NewConstantStep(v int64) Step {
	s := Step{}
	s.do = func() any {
		return v
	}
	return s
}

func NewVarStep(v *int64) Step {
	s := Step{}
	s.do = func() any {
		return *v
	}
	return s
}

func NewAssignStep(variable *int64, expr Step) Step {
	s := Step{}
	s.do = func() any {
		v := expr.do().(int64)
		*variable = v
		return struct{}{}
	}
	return s
}

func NewBinaryPlusStep(left, right Step) Step {
	s := Step{}
	s.do = func() any {
		l := left.do().(int64)
		r := right.do().(int64)
		return l + r
	}
	return s
}

func NewBinarySubStep(left, right Step) Step {
	s := Step{}
	s.do = func() any {
		l := left.do().(int64)
		r := right.do().(int64)
		return l - r
	}
	return s
}

func NewBinaryMultStep(left, right Step) Step {
	s := Step{}
	s.do = func() any {
		l := left.do().(int64)
		r := right.do().(int64)
		return l * r
	}
	return s
}

func NewBinaryDivStep(left, right Step) Step {
	s := Step{}
	s.do = func() any {
		l := left.do().(int64)
		r := right.do().(int64)
		return l / r
	}
	return s
}

func NewBinaryModStep(left, right Step) Step {
	s := Step{}
	s.do = func() any {
		l := left.do().(int64)
		r := right.do().(int64)
		return l % r
	}
	return s
}
