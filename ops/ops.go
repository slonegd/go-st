package ops

import (
	"fmt"

	"github.com/slonegd/go-st/types"
	"golang.org/x/exp/constraints"
)

type (
	Op[T OpTypes | struct{}] struct {
		Do func() T
		// инты передаю всегда через int64
		// это поле говорит об ограничениях, например INT -> int16
		// эти ограничения надо учитывать для возможного неявного приведения
		ResultType types.Basic
		// у константы тип приводиться к аргументу не константе
		// TODO константы вообще можно не через оператор передавать, а сразу аргументом
		IsConstant bool
	}
	Statement = Op[struct{}]
	Int64     = Op[int64]
	Bool      = Op[bool]
	Float64   = Op[float64]
	Any       interface {
		Statement | Int64 | Float64 | Bool
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

func Constant[T OpTypes](v T) Op[T] {
	op := Op[T]{
		IsConstant: true,
	}
	var tmp any = v
	switch tmp.(type) {
	case int64:
		op.ResultType = types.LINT
	case float64:
		op.ResultType = types.LREAL
	case string:
		op.ResultType = types.STRING
	case bool:
		op.ResultType = types.BOOL
	}
	op.Do = func() T {
		return v
	}
	return op
}

func Variable[T OpTypes](variable types.Variable) Op[T] {
	op := Op[T]{
		ResultType: variable.Type(),
	}
	v := (*T)(variable.Pointer())
	op.Do = func() T {
		return *v
	}
	return op
}

func Assign[T int64 | float64](v types.Variable, expr Op[T]) Statement {
	ops := map[types.Basic]func(v types.Variable, expr Op[T]) Statement{
		types.SINT:  func(v types.Variable, expr Op[T]) Statement { return assign[int8](v, expr) },
		types.INT:   func(v types.Variable, expr Op[T]) Statement { return assign[int16](v, expr) },
		types.DINT:  func(v types.Variable, expr Op[T]) Statement { return assign[int32](v, expr) },
		types.LINT:  func(v types.Variable, expr Op[T]) Statement { return assign[int64](v, expr) },
		types.USINT: func(v types.Variable, expr Op[T]) Statement { return assign[uint8](v, expr) },
		types.UINT:  func(v types.Variable, expr Op[T]) Statement { return assign[uint16](v, expr) },
		types.UDINT: func(v types.Variable, expr Op[T]) Statement { return assign[uint32](v, expr) },
		types.ULINT: func(v types.Variable, expr Op[T]) Statement { return assign[uint64](v, expr) },

		types.REAL:  func(v types.Variable, expr Op[T]) Statement { return assign[float32](v, expr) },
		types.LREAL: func(v types.Variable, expr Op[T]) Statement { return assign[float64](v, expr) },
	}
	op := ops[v.Type()]
	return op(v, expr)
}

func Statements(s []Statement) Statement {
	op := Statement{}
	op.Do = func() struct{} {
		for i := range s {
			s[i].Do()
		}
		return struct{}{}
	}
	return op
}

// разные функции под разные типы -> мапа не подходит
func Cast[T int64 | float64](t types.Basic, expr Op[T]) any {
	switch t {
	case types.SINT:
		return cast[int8, int64](expr)
	case types.INT:
		return cast[int16, int64](expr)
	case types.DINT:
		return cast[int32, int64](expr)
	case types.LINT:
		return cast[int64, int64](expr)
	case types.USINT:
		return cast[uint8, int64](expr)
	case types.UINT:
		return cast[uint16, int64](expr)
	case types.UDINT:
		return cast[uint32, int64](expr)
	case types.ULINT:
		return cast[uint64, int64](expr)
	case types.REAL:
		return cast[float32, float64](expr)
	case types.LREAL:
		return cast[float64, float64](expr)
	}
	return nil
}

func Binary[Result, Left, Right NumberOpTypes](op string, left Op[Left], right Op[Right], resultType types.Basic) Op[Result] {
	type K struct {
		op         string
		resultType types.Basic
	}

	ops := map[K]func(l Op[Left], r Op[Right]) Op[Result]{
		{"+", types.SINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return plus[int8, Result](l, r) },
		{"+", types.INT}:   func(l Op[Left], r Op[Right]) Op[Result] { return plus[int16, Result](l, r) },
		{"+", types.DINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return plus[int32, Result](l, r) },
		{"+", types.LINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return plus[int64, Result](l, r) },
		{"+", types.USINT}: func(l Op[Left], r Op[Right]) Op[Result] { return plus[uint8, Result](l, r) },
		{"+", types.UINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return plus[uint16, Result](l, r) },
		{"+", types.UDINT}: func(l Op[Left], r Op[Right]) Op[Result] { return plus[uint32, Result](l, r) },
		{"+", types.ULINT}: func(l Op[Left], r Op[Right]) Op[Result] { return plus[uint64, Result](l, r) },
		{"+", types.REAL}:  func(l Op[Left], r Op[Right]) Op[Result] { return plus[float32, Result](l, r) },
		{"+", types.LREAL}: func(l Op[Left], r Op[Right]) Op[Result] { return plus[float64, Result](l, r) },

		{"-", types.SINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return sub[int8, Result](l, r) },
		{"-", types.INT}:   func(l Op[Left], r Op[Right]) Op[Result] { return sub[int16, Result](l, r) },
		{"-", types.DINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return sub[int32, Result](l, r) },
		{"-", types.LINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return sub[int64, Result](l, r) },
		{"-", types.USINT}: func(l Op[Left], r Op[Right]) Op[Result] { return sub[uint8, Result](l, r) },
		{"-", types.UINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return sub[uint16, Result](l, r) },
		{"-", types.UDINT}: func(l Op[Left], r Op[Right]) Op[Result] { return sub[uint32, Result](l, r) },
		{"-", types.ULINT}: func(l Op[Left], r Op[Right]) Op[Result] { return sub[uint64, Result](l, r) },
		{"-", types.REAL}:  func(l Op[Left], r Op[Right]) Op[Result] { return sub[float32, Result](l, r) },
		{"-", types.LREAL}: func(l Op[Left], r Op[Right]) Op[Result] { return sub[float64, Result](l, r) },

		{"*", types.SINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return mult[int8, Result](l, r) },
		{"*", types.INT}:   func(l Op[Left], r Op[Right]) Op[Result] { return mult[int16, Result](l, r) },
		{"*", types.DINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return mult[int32, Result](l, r) },
		{"*", types.LINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return mult[int64, Result](l, r) },
		{"*", types.USINT}: func(l Op[Left], r Op[Right]) Op[Result] { return mult[uint8, Result](l, r) },
		{"*", types.UINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return mult[uint16, Result](l, r) },
		{"*", types.UDINT}: func(l Op[Left], r Op[Right]) Op[Result] { return mult[uint32, Result](l, r) },
		{"*", types.ULINT}: func(l Op[Left], r Op[Right]) Op[Result] { return mult[uint64, Result](l, r) },
		{"*", types.REAL}:  func(l Op[Left], r Op[Right]) Op[Result] { return mult[float32, Result](l, r) },
		{"*", types.LREAL}: func(l Op[Left], r Op[Right]) Op[Result] { return mult[float64, Result](l, r) },

		{"/", types.SINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return div[int8, Result](l, r) },
		{"/", types.INT}:   func(l Op[Left], r Op[Right]) Op[Result] { return div[int16, Result](l, r) },
		{"/", types.DINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return div[int32, Result](l, r) },
		{"/", types.LINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return div[int64, Result](l, r) },
		{"/", types.USINT}: func(l Op[Left], r Op[Right]) Op[Result] { return div[uint8, Result](l, r) },
		{"/", types.UINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return div[uint16, Result](l, r) },
		{"/", types.UDINT}: func(l Op[Left], r Op[Right]) Op[Result] { return div[uint32, Result](l, r) },
		{"/", types.ULINT}: func(l Op[Left], r Op[Right]) Op[Result] { return div[uint64, Result](l, r) },
		{"/", types.REAL}:  func(l Op[Left], r Op[Right]) Op[Result] { return div[float32, Result](l, r) },
		{"/", types.LREAL}: func(l Op[Left], r Op[Right]) Op[Result] { return div[float64, Result](l, r) },

		{"MOD", types.SINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return mod[int8, Result](l, r) },
		{"MOD", types.INT}:   func(l Op[Left], r Op[Right]) Op[Result] { return mod[int16, Result](l, r) },
		{"MOD", types.DINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return mod[int32, Result](l, r) },
		{"MOD", types.LINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return mod[int64, Result](l, r) },
		{"MOD", types.USINT}: func(l Op[Left], r Op[Right]) Op[Result] { return mod[uint8, Result](l, r) },
		{"MOD", types.UINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return mod[uint16, Result](l, r) },
		{"MOD", types.UDINT}: func(l Op[Left], r Op[Right]) Op[Result] { return mod[uint32, Result](l, r) },
		{"MOD", types.ULINT}: func(l Op[Left], r Op[Right]) Op[Result] { return mod[uint64, Result](l, r) },
	}
	return ops[K{op, resultType}](left, right)
}

func Greater[T int64 | float64](left, right Op[T]) Bool {
	op := Bool{}
	op.Do = func() bool {
		l := left.Do()
		r := right.Do()
		return l > r
	}
	return op
}

func GreaterOrEqual[T int64 | float64](left, right Op[T]) Bool {
	op := Bool{}
	op.Do = func() bool {
		l := left.Do()
		r := right.Do()
		return l >= r
	}
	return op
}

func Less[T int64 | float64](left, right Op[T]) Bool {
	op := Bool{}
	op.Do = func() bool {
		l := left.Do()
		r := right.Do()
		return l < r
	}
	return op
}

func LessOrEqual[T int64 | float64](left, right Op[T]) Bool {
	op := Bool{}
	op.Do = func() bool {
		l := left.Do()
		r := right.Do()
		return l <= r
	}
	return op
}

func Equal[T int64 | float64](left, right Op[T]) Bool {
	op := Bool{}
	op.Do = func() bool {
		l := left.Do()
		r := right.Do()
		return l == r
	}
	return op
}

func NotEqual[T int64 | float64](left, right Op[T]) Bool {
	op := Bool{}
	op.Do = func() bool {
		l := left.Do()
		r := right.Do()
		return l != r
	}
	return op
}

func If(cond Bool, then_ Statement) Statement {
	op := Statement{}
	op.Do = func() struct{} {
		if cond.Do() {
			then_.Do()
		}
		return struct{}{}
	}
	return op
}

func IfElse(cond Bool, then_, else_ Statement) Statement {
	op := Statement{}
	op.Do = func() struct{} {
		if cond.Do() {
			then_.Do()
		} else {
			else_.Do()
		}
		return struct{}{}
	}
	return op
}

func assign[T Number, R int64 | float64](variable types.Variable, expr Op[R]) Statement {
	op := Statement{}
	val := (*R)(variable.Pointer())

	op.Do = func() struct{} {
		v := T(expr.Do())
		*val = R(v)
		return struct{}{}
	}
	return op
}

func cast[T Number, Tout, Tin int64 | float64](expr Op[Tin]) Op[Tout] {
	op := Op[Tout]{
		IsConstant: expr.IsConstant,
		ResultType: BasicType[T](),
	}
	op.Do = func() Tout {
		v := T(expr.Do())
		return Tout(v)
	}
	return op
}

// T для ограничения размера инта
func plus[T Number, Result, Left, Right NumberOpTypes](left Op[Left], right Op[Right]) Op[Result] {
	op := Op[Result]{
		ResultType: BasicType[T](),
		IsConstant: left.IsConstant && right.IsConstant,
	}
	op.Do = func() Result {
		l := T(left.Do())
		r := T(right.Do())
		return Result(l + r)
	}
	return op
}

func BasicType[T Number]() types.Basic {
	var v T
	var tmp any = v
	switch tmp.(type) {
	case int8:
		return types.SINT
	case int16:
		return types.INT
	case int32:
		return types.DINT
	case int64:
		return types.LINT
	case uint8:
		return types.USINT
	case uint16:
		return types.UINT
	case uint32:
		return types.UDINT
	case uint64:
		return types.ULINT
	case float32:
		return types.REAL
	case float64:
		return types.LREAL
	default:
		panic(fmt.Sprintf("cant find type from %T", v))
	}
}

func sub[T Number, Result, Left, Right NumberOpTypes](left Op[Left], right Op[Right]) Op[Result] {
	op := Op[Result]{
		ResultType: BasicType[T](),
		IsConstant: left.IsConstant && right.IsConstant,
	}
	op.Do = func() Result {
		l := T(left.Do())
		r := T(right.Do())
		return Result(l - r)
	}
	return op
}

func mult[T Number, Result, Left, Right NumberOpTypes](left Op[Left], right Op[Right]) Op[Result] {
	op := Op[Result]{
		ResultType: BasicType[T](),
		IsConstant: left.IsConstant && right.IsConstant,
	}
	op.Do = func() Result {
		l := T(left.Do())
		r := T(right.Do())
		return Result(l * r)
	}
	return op
}

func div[T Number, Result, Left, Right NumberOpTypes](left Op[Left], right Op[Right]) Op[Result] {
	op := Op[Result]{
		ResultType: BasicType[T](),
		IsConstant: left.IsConstant && right.IsConstant,
	}
	op.Do = func() Result {
		l := T(left.Do())
		r := T(right.Do())
		return Result(l / r)
	}
	return op
}

func mod[T constraints.Integer, Result, Left, Right NumberOpTypes](left Op[Left], right Op[Right]) Op[Result] {
	op := Op[Result]{
		ResultType: BasicType[T](),
		IsConstant: left.IsConstant && right.IsConstant,
	}
	op.Do = func() Result {
		l := T(left.Do())
		r := T(right.Do())
		return Result(l % r)
	}
	return op
}
