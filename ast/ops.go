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

func NewAssignOp[T constraints.Integer, R int64](variable variant.Variant, expr Operator[R]) Statement {
	op := Statement{}
	val := (*R)(variable.Pointer())

	op.do = func() struct{} {
		v := T(expr.do())
		*val = R(v)
		return struct{}{}
	}
	return op
}

var assignOps = map[variant.Type]func(v variant.Variant, expr Int64Operator) any{
	variant.SINT:  func(v variant.Variant, expr Int64Operator) any { return NewAssignOp[int8](v, expr) },
	variant.INT:   func(v variant.Variant, expr Int64Operator) any { return NewAssignOp[int16](v, expr) },
	variant.DINT:  func(v variant.Variant, expr Int64Operator) any { return NewAssignOp[int32](v, expr) },
	variant.LINT:  func(v variant.Variant, expr Int64Operator) any { return NewAssignOp[int64](v, expr) },
	variant.USINT: func(v variant.Variant, expr Int64Operator) any { return NewAssignOp[uint8](v, expr) },
	variant.UINT:  func(v variant.Variant, expr Int64Operator) any { return NewAssignOp[uint16](v, expr) },
	variant.UDINT: func(v variant.Variant, expr Int64Operator) any { return NewAssignOp[uint32](v, expr) },
	variant.ULINT: func(v variant.Variant, expr Int64Operator) any { return NewAssignOp[uint64](v, expr) },
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

func NewCastOp[T constraints.Integer, R int64](expr Operator[R]) Operator[R] {
	op := Operator[R]{
		isConstant: expr.isConstant,
		resultType: VariantType[T](),
	}
	op.do = func() R {
		v := T(expr.do())
		return R(v)
	}
	return op
}

var castOps = map[variant.Type]func(expr Int64Operator) any{
	variant.SINT:  func(expr Int64Operator) any { return NewCastOp[int8](expr) },
	variant.INT:   func(expr Int64Operator) any { return NewCastOp[int16](expr) },
	variant.DINT:  func(expr Int64Operator) any { return NewCastOp[int32](expr) },
	variant.LINT:  func(expr Int64Operator) any { return NewCastOp[int64](expr) },
	variant.USINT: func(expr Int64Operator) any { return NewCastOp[uint8](expr) },
	variant.UINT:  func(expr Int64Operator) any { return NewCastOp[uint16](expr) },
	variant.UDINT: func(expr Int64Operator) any { return NewCastOp[uint32](expr) },
	variant.ULINT: func(expr Int64Operator) any { return NewCastOp[uint64](expr) },
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

var binaryOps = map[binaryOpKey]func(left, right Int64Operator) any{
	{"+", variant.SINT}:  func(left, right Int64Operator) any { return NewPlusOp[int8](left, right) },
	{"+", variant.INT}:   func(left, right Int64Operator) any { return NewPlusOp[int16](left, right) },
	{"+", variant.DINT}:  func(left, right Int64Operator) any { return NewPlusOp[int32](left, right) },
	{"+", variant.LINT}:  func(left, right Int64Operator) any { return NewPlusOp[int64](left, right) },
	{"+", variant.USINT}: func(left, right Int64Operator) any { return NewPlusOp[uint8](left, right) },
	{"+", variant.UINT}:  func(left, right Int64Operator) any { return NewPlusOp[uint16](left, right) },
	{"+", variant.UDINT}: func(left, right Int64Operator) any { return NewPlusOp[uint32](left, right) },
	{"+", variant.ULINT}: func(left, right Int64Operator) any { return NewPlusOp[uint64](left, right) },

	{"-", variant.SINT}:  func(left, right Int64Operator) any { return NewSubOp[int8](left, right) },
	{"-", variant.INT}:   func(left, right Int64Operator) any { return NewSubOp[int16](left, right) },
	{"-", variant.DINT}:  func(left, right Int64Operator) any { return NewSubOp[int32](left, right) },
	{"-", variant.LINT}:  func(left, right Int64Operator) any { return NewSubOp[int64](left, right) },
	{"-", variant.USINT}: func(left, right Int64Operator) any { return NewSubOp[uint8](left, right) },
	{"-", variant.UINT}:  func(left, right Int64Operator) any { return NewSubOp[uint16](left, right) },
	{"-", variant.UDINT}: func(left, right Int64Operator) any { return NewSubOp[uint32](left, right) },
	{"-", variant.ULINT}: func(left, right Int64Operator) any { return NewSubOp[uint64](left, right) },

	{"*", variant.SINT}:  func(left, right Int64Operator) any { return NewMultOp[int8](left, right) },
	{"*", variant.INT}:   func(left, right Int64Operator) any { return NewMultOp[int16](left, right) },
	{"*", variant.DINT}:  func(left, right Int64Operator) any { return NewMultOp[int32](left, right) },
	{"*", variant.LINT}:  func(left, right Int64Operator) any { return NewMultOp[int64](left, right) },
	{"*", variant.USINT}: func(left, right Int64Operator) any { return NewMultOp[uint8](left, right) },
	{"*", variant.UINT}:  func(left, right Int64Operator) any { return NewMultOp[uint16](left, right) },
	{"*", variant.UDINT}: func(left, right Int64Operator) any { return NewMultOp[uint32](left, right) },
	{"*", variant.ULINT}: func(left, right Int64Operator) any { return NewMultOp[uint64](left, right) },

	{"/", variant.SINT}:  func(left, right Int64Operator) any { return NewDivOp[int8](left, right) },
	{"/", variant.INT}:   func(left, right Int64Operator) any { return NewDivOp[int16](left, right) },
	{"/", variant.DINT}:  func(left, right Int64Operator) any { return NewDivOp[int32](left, right) },
	{"/", variant.LINT}:  func(left, right Int64Operator) any { return NewDivOp[int64](left, right) },
	{"/", variant.USINT}: func(left, right Int64Operator) any { return NewDivOp[uint8](left, right) },
	{"/", variant.UINT}:  func(left, right Int64Operator) any { return NewDivOp[uint16](left, right) },
	{"/", variant.UDINT}: func(left, right Int64Operator) any { return NewDivOp[uint32](left, right) },
	{"/", variant.ULINT}: func(left, right Int64Operator) any { return NewDivOp[uint64](left, right) },

	{"MOD", variant.SINT}:  func(left, right Int64Operator) any { return NewModOp[int8](left, right) },
	{"MOD", variant.INT}:   func(left, right Int64Operator) any { return NewModOp[int16](left, right) },
	{"MOD", variant.DINT}:  func(left, right Int64Operator) any { return NewModOp[int32](left, right) },
	{"MOD", variant.LINT}:  func(left, right Int64Operator) any { return NewModOp[int64](left, right) },
	{"MOD", variant.USINT}: func(left, right Int64Operator) any { return NewModOp[uint8](left, right) },
	{"MOD", variant.UINT}:  func(left, right Int64Operator) any { return NewModOp[uint16](left, right) },
	{"MOD", variant.UDINT}: func(left, right Int64Operator) any { return NewModOp[uint32](left, right) },
	{"MOD", variant.ULINT}: func(left, right Int64Operator) any { return NewModOp[uint64](left, right) },
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
