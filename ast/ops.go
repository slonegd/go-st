package ast

import (
	"fmt"

	"github.com/slonegd/go-st/variant"
	"golang.org/x/exp/constraints"
)

type (
	Operator[T any] struct {
		do    func() T
		descr string
		// инты передаю всегда через int64
		// это поле говорит об ограничениях, например INT -> int16
		// эти ограничения надо учитывать для возможного неявного приведения
		resultType variant.Type
		// у константы тип приводиться к аргументу не константе
		// TODO константы вообще можно не через оператор передавать, а сразу аргументом
		isConstant bool
	}
	Statement       = Operator[struct{}]
	Statements      []Statement
	Int64Operator   = Operator[int64]
	BoolOperator    = Operator[bool]
	Float64Operator = Operator[float64]
	AnyOperator     interface {
		Statement | Int64Operator | Float64Operator | BoolOperator
	}
	// только такие типы могут быть результатом оператора
	// специально ограничил, чтобы было меньше ветвлений
	// с типами меньшего размера буду приводить при вычислении
	OpTypes interface {
		int64 | float64 | string | bool
	}
)

func (s *Operator[T]) WithDescription(v string) *Operator[T] { s.descr = v; return s }

func NewConstantOp[T OpTypes](v T) Operator[T] {
	op := Operator[T]{
		isConstant: true,
	}
	var tmp any = v
	switch tmp.(type) {
	case int64:
		op.resultType = variant.LINT
	case float64:
		op.resultType = variant.LREAL
	case string:
		op.resultType = variant.STRING
	case bool:
		op.resultType = variant.BOOL
	}
	op.do = func() T {
		return v
	}
	return op
}

func NewVarOp[T any](variable variant.Variant) Operator[T] {
	op := Operator[T]{
		resultType: variable.Type(),
	}
	v := (*T)(variable.Pointer())
	op.do = func() T {
		return *v
	}
	return op
}

func NewAssignOp[T any](variable *T, expr Operator[T]) Statement {
	// TODO проверить тип
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

// T для ограничения размера инта
func NewPlusOp[T constraints.Integer, R int64 | float64](left, right Operator[R]) Operator[R] {
	op := Operator[R]{
		resultType: VariantType[T](),
		isConstant: left.isConstant && right.isConstant,
	}
	op.do = func() R {
		l := T(left.do())
		r := T(right.do())
		return R(l + r)
	}
	return op
}

func VariantType[T constraints.Integer]() variant.Type {
	var v T
	var tmp any = v
	switch tmp.(type) {
	case int8:
		return variant.SINT
	case int16:
		return variant.INT
	case int32:
		return variant.DINT
	case int64:
		return variant.LINT
	case uint8:
		return variant.USINT
	case uint16:
		return variant.UINT
	case uint32:
		return variant.UDINT
	case uint64:
		return variant.ULINT
	default:
		panic(fmt.Sprintf("cant find type from %T", v))
	}
}

func NewSubOp[T constraints.Integer, R int64 | float64](left, right Operator[R]) Operator[R] {
	op := Operator[R]{
		resultType: VariantType[T](),
		isConstant: left.isConstant && right.isConstant,
	}
	op.do = func() R {
		l := T(left.do())
		r := T(right.do())
		return R(l - r)
	}
	return op
}

func NewMultOp[T constraints.Integer, R int64 | float64](left, right Operator[R]) Operator[R] {
	op := Operator[R]{
		resultType: VariantType[T](),
		isConstant: left.isConstant && right.isConstant,
	}
	op.do = func() R {
		l := T(left.do())
		r := T(right.do())
		return R(l * r)
	}
	return op
}

func NewDivOp[T constraints.Integer, R int64 | float64](left, right Operator[R]) Operator[R] {
	op := Operator[R]{
		resultType: VariantType[T](),
		isConstant: left.isConstant && right.isConstant,
	}
	op.do = func() R {
		l := T(left.do())
		r := T(right.do())
		return R(l / r)
	}
	return op
}

func NewModOp[T constraints.Integer, R int64](left, right Operator[R]) Operator[R] {
	op := Operator[R]{
		resultType: VariantType[T](),
		isConstant: left.isConstant && right.isConstant,
	}
	op.do = func() R {
		l := T(left.do())
		r := T(right.do())
		return R(l % r)
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
