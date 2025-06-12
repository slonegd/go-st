package fvm

import (
	"github.com/antlr4-go/antlr/v4"
	parser "github.com/slonegd/go-st/antlr"
	"github.com/slonegd/go-st/variant"
)

type (
	FVM struct {
		vars     map[string]variant.Variant
		programm Statement
	}
	Operator[T any] struct {
		do    func() T
		descr string
	}
	Statement       = Operator[struct{}]
	Statements      []Statement
	Int64Operator   = Operator[int64]
	BoolOperator    = Operator[bool]
	Float64Operator = Operator[float64]
	AnyOperator     interface {
		Statement | Int64Operator | Float64Operator | BoolOperator
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

	p.BuildParseTrees = true

	x := &FVM{vars: map[string]variant.Variant{}}
	x.programm = p.Program().Accept(x).(Statement)
	// listener := listener{Compiler: x}
	// p.BaseRecognizer.AddErrorListener(listener)

	// Finally parse the expression (by walking the tree)
	// antlr.ParseTreeWalkerDefault.Walk(x, p.Program())
	// if p.SynErr != nil && p.SynErr.GetMessage() != "" {
	// 	return nil, errors.New(p.SynErr.GetMessage())
	// }

	return x, nil // x.err
}

func (x *FVM) GetVar(name string) variant.Variant {
	return x.vars[name]
}

func (x *FVM) Execute() {
	x.programm.do()
}

func (s *Operator[T]) WithDescription(v string) *Operator[T] { s.descr = v; return s }

func NewConstantOp[T any](v T) Operator[T] {
	op := Operator[T]{}
	op.do = func() T {
		return v
	}
	return op
}

func NewVarOp[T any](v *T) Operator[T] {
	op := Operator[T]{}
	op.do = func() T {
		return *v
	}
	return op
}

func NewAssignOp[T any](variable *T, expr Operator[T]) Statement {
	op := Statement{}
	op.do = func() struct{} {
		v := expr.do()
		*variable = v
		return struct{}{}
	}
	return op
}

func NewStatments(s Statements) Statement {
	op := Statement{}
	op.do = func() struct{} {
		for i := range s {
			s[i].do()
		}
		return struct{}{}
	}
	return op
}

func NewPlusOp[T int64 | float64](left, right Operator[T]) Operator[T] {
	op := Operator[T]{}
	op.do = func() T {
		l := left.do()
		r := right.do()
		return l + r
	}
	return op
}

func NewSubOp[T int64 | float64](left, right Operator[T]) Operator[T] {
	op := Operator[T]{}
	op.do = func() T {
		l := left.do()
		r := right.do()
		return l - r
	}
	return op
}

func NewMultOp[T int64 | float64](left, right Operator[T]) Operator[T] {
	op := Operator[T]{}
	op.do = func() T {
		l := left.do()
		r := right.do()
		return l * r
	}
	return op
}

func NewDivOp[T int64 | float64](left, right Operator[T]) Operator[T] {
	op := Operator[T]{}
	op.do = func() T {
		l := left.do()
		r := right.do()
		return l / r
	}
	return op
}

func NewModOp[T int64](left, right Operator[T]) Operator[T] {
	op := Operator[T]{}
	op.do = func() T {
		l := left.do()
		r := right.do()
		return l % r
	}
	return op
}

func NewGreaterOp[T int64 | float64](left, right Operator[T]) BoolOperator {
	op := BoolOperator{}
	op.do = func() bool {
		l := left.do()
		r := right.do()
		return l > r
	}
	return op
}

func NewGreaterOrEqualOp[T int64 | float64](left, right Operator[T]) BoolOperator {
	op := BoolOperator{}
	op.do = func() bool {
		l := left.do()
		r := right.do()
		return l >= r
	}
	return op
}

func NewLessOp[T int64 | float64](left, right Operator[T]) BoolOperator {
	op := BoolOperator{}
	op.do = func() bool {
		l := left.do()
		r := right.do()
		return l < r
	}
	return op
}

func NewLessOrEqualOp[T int64 | float64](left, right Operator[T]) BoolOperator {
	op := BoolOperator{}
	op.do = func() bool {
		l := left.do()
		r := right.do()
		return l <= r
	}
	return op
}

func NewEqualOp[T int64 | float64](left, right Operator[T]) BoolOperator {
	op := BoolOperator{}
	op.do = func() bool {
		l := left.do()
		r := right.do()
		return l == r
	}
	return op
}

func NewNotEqualOp[T int64 | float64](left, right Operator[T]) BoolOperator {
	op := BoolOperator{}
	op.do = func() bool {
		l := left.do()
		r := right.do()
		return l != r
	}
	return op
}

func NewIfOperator(cond BoolOperator, then_ Statement) Statement {
	op := Statement{}
	op.do = func() struct{} {
		if cond.do() {
			then_.do()
		}
		return struct{}{}
	}
	return op
}

func NewIfElseOperator(cond BoolOperator, then_, else_ Statement) Statement {
	op := Statement{}
	op.do = func() struct{} {
		if cond.do() {
			then_.do()
		} else {
			else_.do()
		}
		return struct{}{}
	}
	return op
}
