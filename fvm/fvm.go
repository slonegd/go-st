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
		vars       map[string]variant.Variant
		operators  stack.Stack[any] // any -> AnyOperator
		statements []Statement
	}
	Operator[T any] struct {
		do    func() T
		descr string
	}
	Statement       = Operator[struct{}]
	Int64Operator   = Operator[int64]
	Float64Operator = Operator[float64]
	AnyOperator     interface {
		Statement | Int64Operator | Float64Operator
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

	x.statements = make([]Statement, 0, len(x.operators))
	for _, s := range x.operators {
		x.statements = append(x.statements, s.(Statement))
	}

	return x, nil // x.err
}

func (x *FVM) GetVar(name string) variant.Variant {
	return x.vars[name]
}

func (x *FVM) Execute() {
	for _, s := range x.statements {
		s.do()
	}
}

func (s *Operator[T]) WithDescription(v string) *Operator[T] { s.descr = v; return s }

func NewConstantStep[T any](v T) Operator[T] {
	s := Operator[T]{}
	s.do = func() T {
		return v
	}
	return s
}

func NewVarStep[T any](v *T) Operator[T] {
	s := Operator[T]{}
	s.do = func() T {
		return *v
	}
	return s
}

func NewAssignStep[T any](variable *T, expr Operator[T]) Statement {
	s := Statement{}
	s.do = func() struct{} {
		v := expr.do()
		*variable = v
		return struct{}{}
	}
	return s
}

func NewBinaryPlusStep[T int64 | float64](left, right Operator[T]) Operator[T] {
	s := Operator[T]{}
	s.do = func() T {
		l := left.do()
		r := right.do()
		return l + r
	}
	return s
}

func NewBinarySubStep[T int64 | float64](left, right Operator[T]) Operator[T] {
	s := Operator[T]{}
	s.do = func() T {
		l := left.do()
		r := right.do()
		return l - r
	}
	return s
}

func NewBinaryMultStep[T int64 | float64](left, right Operator[T]) Operator[T] {
	s := Operator[T]{}
	s.do = func() T {
		l := left.do()
		r := right.do()
		return l * r
	}
	return s
}

func NewBinaryDivStep[T int64 | float64](left, right Operator[T]) Operator[T] {
	s := Operator[T]{}
	s.do = func() T {
		l := left.do()
		r := right.do()
		return l / r
	}
	return s
}

func NewBinaryModStep[T int64](left, right Operator[T]) Operator[T] {
	s := Operator[T]{}
	s.do = func() T {
		l := left.do()
		r := right.do()
		return l % r
	}
	return s
}
