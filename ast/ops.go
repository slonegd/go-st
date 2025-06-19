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
	NumberOpTypes interface {
		int64 | float64
	}
	Number interface {
		constraints.Integer | constraints.Float
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

func NewAssignOp[T Number, R int64 | float64](variable variant.Variant, expr Operator[R]) Statement {
	op := Statement{}
	val := (*R)(variable.Pointer())

	// TODO T from variable.Type

	op.do = func() struct{} {
		v := T(expr.do())
		*val = R(v)
		return struct{}{}
	}
	return op
}

func assignOps[T int64 | float64](v variant.Variant, expr Operator[T]) Statement {
	ops := map[variant.Type]func(v variant.Variant, expr Operator[T]) Statement{
		variant.SINT:  func(v variant.Variant, expr Operator[T]) Statement { return NewAssignOp[int8](v, expr) },
		variant.INT:   func(v variant.Variant, expr Operator[T]) Statement { return NewAssignOp[int16](v, expr) },
		variant.DINT:  func(v variant.Variant, expr Operator[T]) Statement { return NewAssignOp[int32](v, expr) },
		variant.LINT:  func(v variant.Variant, expr Operator[T]) Statement { return NewAssignOp[int64](v, expr) },
		variant.USINT: func(v variant.Variant, expr Operator[T]) Statement { return NewAssignOp[uint8](v, expr) },
		variant.UINT:  func(v variant.Variant, expr Operator[T]) Statement { return NewAssignOp[uint16](v, expr) },
		variant.UDINT: func(v variant.Variant, expr Operator[T]) Statement { return NewAssignOp[uint32](v, expr) },
		variant.ULINT: func(v variant.Variant, expr Operator[T]) Statement { return NewAssignOp[uint64](v, expr) },

		variant.REAL:  func(v variant.Variant, expr Operator[T]) Statement { return NewAssignOp[float32](v, expr) },
		variant.LREAL: func(v variant.Variant, expr Operator[T]) Statement { return NewAssignOp[float64](v, expr) },
	}
	op := ops[v.Type()]
	return op(v, expr)
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

func NewCastOp[T Number, Tout, Tin int64 | float64](expr Operator[Tin]) Operator[Tout] {
	op := Operator[Tout]{
		isConstant: expr.isConstant,
		resultType: VariantType[T](),
	}
	op.do = func() Tout {
		v := T(expr.do())
		return Tout(v)
	}
	return op
}

// разные функции под разные типы -> мапа не подходит
func castOps[T int64 | float64](t variant.Type, expr Operator[T]) any {
	switch t {
	case variant.SINT:
		return NewCastOp[int8, int64](expr)
	case variant.INT:
		return NewCastOp[int16, int64](expr)
	case variant.DINT:
		return NewCastOp[int32, int64](expr)
	case variant.LINT:
		return NewCastOp[int64, int64](expr)
	case variant.USINT:
		return NewCastOp[uint8, int64](expr)
	case variant.UINT:
		return NewCastOp[uint16, int64](expr)
	case variant.UDINT:
		return NewCastOp[uint32, int64](expr)
	case variant.ULINT:
		return NewCastOp[uint64, int64](expr)
	case variant.REAL:
		return NewCastOp[float32, float64](expr)
	case variant.LREAL:
		return NewCastOp[float64, float64](expr)
	}
	return nil
}

// T для ограничения размера инта
func NewPlusOp[T Number, Result, Left, Right NumberOpTypes](left Operator[Left], right Operator[Right]) Operator[Result] {
	op := Operator[Result]{
		resultType: VariantType[T](),
		isConstant: left.isConstant && right.isConstant,
	}
	op.do = func() Result {
		l := T(left.do())
		r := T(right.do())
		return Result(l + r)
	}
	return op
}

func VariantType[T Number]() variant.Type {
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
	case float32:
		return variant.REAL
	case float64:
		return variant.LREAL
	default:
		panic(fmt.Sprintf("cant find type from %T", v))
	}
}

func binaryOps[Result, Left, Right NumberOpTypes](op string, left Operator[Left], right Operator[Right], resultType variant.Type) Operator[Result] {
	type K struct {
		op         string
		resultType variant.Type
	}

	ops := map[K]func(l Operator[Left], r Operator[Right]) Operator[Result]{
		{"+", variant.SINT}:  func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewPlusOp[int8, Result](l, r) },
		{"+", variant.INT}:   func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewPlusOp[int16, Result](l, r) },
		{"+", variant.DINT}:  func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewPlusOp[int32, Result](l, r) },
		{"+", variant.LINT}:  func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewPlusOp[int64, Result](l, r) },
		{"+", variant.USINT}: func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewPlusOp[uint8, Result](l, r) },
		{"+", variant.UINT}:  func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewPlusOp[uint16, Result](l, r) },
		{"+", variant.UDINT}: func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewPlusOp[uint32, Result](l, r) },
		{"+", variant.ULINT}: func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewPlusOp[uint64, Result](l, r) },
		{"+", variant.REAL}:  func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewPlusOp[float32, Result](l, r) },
		{"+", variant.LREAL}: func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewPlusOp[float64, Result](l, r) },

		{"-", variant.SINT}:  func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewSubOp[int8, Result](l, r) },
		{"-", variant.INT}:   func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewSubOp[int16, Result](l, r) },
		{"-", variant.DINT}:  func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewSubOp[int32, Result](l, r) },
		{"-", variant.LINT}:  func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewSubOp[int64, Result](l, r) },
		{"-", variant.USINT}: func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewSubOp[uint8, Result](l, r) },
		{"-", variant.UINT}:  func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewSubOp[uint16, Result](l, r) },
		{"-", variant.UDINT}: func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewSubOp[uint32, Result](l, r) },
		{"-", variant.ULINT}: func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewSubOp[uint64, Result](l, r) },
		{"-", variant.REAL}:  func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewSubOp[float32, Result](l, r) },
		{"-", variant.LREAL}: func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewSubOp[float64, Result](l, r) },

		{"*", variant.SINT}:  func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewMultOp[int8, Result](l, r) },
		{"*", variant.INT}:   func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewMultOp[int16, Result](l, r) },
		{"*", variant.DINT}:  func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewMultOp[int32, Result](l, r) },
		{"*", variant.LINT}:  func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewMultOp[int64, Result](l, r) },
		{"*", variant.USINT}: func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewMultOp[uint8, Result](l, r) },
		{"*", variant.UINT}:  func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewMultOp[uint16, Result](l, r) },
		{"*", variant.UDINT}: func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewMultOp[uint32, Result](l, r) },
		{"*", variant.ULINT}: func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewMultOp[uint64, Result](l, r) },
		{"*", variant.REAL}:  func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewMultOp[float32, Result](l, r) },
		{"*", variant.LREAL}: func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewMultOp[float64, Result](l, r) },

		{"/", variant.SINT}:  func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewDivOp[int8, Result](l, r) },
		{"/", variant.INT}:   func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewDivOp[int16, Result](l, r) },
		{"/", variant.DINT}:  func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewDivOp[int32, Result](l, r) },
		{"/", variant.LINT}:  func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewDivOp[int64, Result](l, r) },
		{"/", variant.USINT}: func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewDivOp[uint8, Result](l, r) },
		{"/", variant.UINT}:  func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewDivOp[uint16, Result](l, r) },
		{"/", variant.UDINT}: func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewDivOp[uint32, Result](l, r) },
		{"/", variant.ULINT}: func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewDivOp[uint64, Result](l, r) },
		{"/", variant.REAL}:  func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewDivOp[float32, Result](l, r) },
		{"/", variant.LREAL}: func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewDivOp[float64, Result](l, r) },

		{"MOD", variant.SINT}:  func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewModOp[int8, Result](l, r) },
		{"MOD", variant.INT}:   func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewModOp[int16, Result](l, r) },
		{"MOD", variant.DINT}:  func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewModOp[int32, Result](l, r) },
		{"MOD", variant.LINT}:  func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewModOp[int64, Result](l, r) },
		{"MOD", variant.USINT}: func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewModOp[uint8, Result](l, r) },
		{"MOD", variant.UINT}:  func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewModOp[uint16, Result](l, r) },
		{"MOD", variant.UDINT}: func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewModOp[uint32, Result](l, r) },
		{"MOD", variant.ULINT}: func(l Operator[Left], r Operator[Right]) Operator[Result] { return NewModOp[uint64, Result](l, r) },
	}
	return ops[K{op, resultType}](left, right)
}

func NewSubOp[T Number, Result, Left, Right NumberOpTypes](left Operator[Left], right Operator[Right]) Operator[Result] {
	op := Operator[Result]{
		resultType: VariantType[T](),
		isConstant: left.isConstant && right.isConstant,
	}
	op.do = func() Result {
		l := T(left.do())
		r := T(right.do())
		return Result(l - r)
	}
	return op
}

func NewMultOp[T Number, Result, Left, Right NumberOpTypes](left Operator[Left], right Operator[Right]) Operator[Result] {
	op := Operator[Result]{
		resultType: VariantType[T](),
		isConstant: left.isConstant && right.isConstant,
	}
	op.do = func() Result {
		l := T(left.do())
		r := T(right.do())
		return Result(l * r)
	}
	return op
}

func NewDivOp[T Number, Result, Left, Right NumberOpTypes](left Operator[Left], right Operator[Right]) Operator[Result] {
	op := Operator[Result]{
		resultType: VariantType[T](),
		isConstant: left.isConstant && right.isConstant,
	}
	op.do = func() Result {
		l := T(left.do())
		r := T(right.do())
		return Result(l / r)
	}
	return op
}

func NewModOp[T constraints.Integer, Result, Left, Right NumberOpTypes](left Operator[Left], right Operator[Right]) Operator[Result] {
	op := Operator[Result]{
		resultType: VariantType[T](),
		isConstant: left.isConstant && right.isConstant,
	}
	op.do = func() Result {
		l := T(left.do())
		r := T(right.do())
		return Result(l % r)
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
