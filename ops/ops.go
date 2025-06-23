package ops

import (
	"fmt"

	parser "github.com/slonegd/go-st/antlr"
	"github.com/slonegd/go-st/types"
	"golang.org/x/exp/constraints"
)

type (
	Op[T OpTypes | struct{}] struct {
		Do      func() T
		DoDebug func() T // тоже самое, что Do, но с выводом в логгер дебажной инфы
		// инты передаю всегда через int64
		// это поле говорит об ограничениях, например INT -> int16
		// эти ограничения надо учитывать для возможного неявного приведения
		ResultType types.Basic
		// у константы тип приводиться к аргументу не константе
		// TODO константы вообще можно не через оператор передавать, а сразу аргументом
		IsConstant bool
		// мета инфа для дебага
		ctx     *parser.CustomContext
		snippet string
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
	op.DoDebug = func() T {
		return v
	}
	return op
}

func Variable[T OpTypes](ctx *parser.CustomContext, name string, variable types.Variable) Op[T] {
	op := Op[T]{
		ResultType: variable.Type(),
		ctx:        ctx,
		snippet:    ctx.MakeSnippet(),
	}
	v := (*T)(variable.Pointer())
	op.Do = func() T {
		return *v
	}
	op.DoDebug = func() T {
		msg := fmt.Sprintf("%s:%v", name, *v)
		ctx.Debug(op.snippet, msg)
		return *v
	}
	return op
}

func Assign[T int64 | float64](ctx *parser.CustomContext, name string, v types.Variable, expr Op[T]) Statement {
	ctx.Name = name // добавляем для дебага
	ops := map[types.Basic]func(v types.Variable, expr Op[T]) Statement{
		types.SINT:  func(v types.Variable, expr Op[T]) Statement { return assign[int8](ctx, v, expr) },
		types.INT:   func(v types.Variable, expr Op[T]) Statement { return assign[int16](ctx, v, expr) },
		types.DINT:  func(v types.Variable, expr Op[T]) Statement { return assign[int32](ctx, v, expr) },
		types.LINT:  func(v types.Variable, expr Op[T]) Statement { return assign[int64](ctx, v, expr) },
		types.USINT: func(v types.Variable, expr Op[T]) Statement { return assign[uint8](ctx, v, expr) },
		types.UINT:  func(v types.Variable, expr Op[T]) Statement { return assign[uint16](ctx, v, expr) },
		types.UDINT: func(v types.Variable, expr Op[T]) Statement { return assign[uint32](ctx, v, expr) },
		types.ULINT: func(v types.Variable, expr Op[T]) Statement { return assign[uint64](ctx, v, expr) },

		types.REAL:  func(v types.Variable, expr Op[T]) Statement { return assign[float32](ctx, v, expr) },
		types.LREAL: func(v types.Variable, expr Op[T]) Statement { return assign[float64](ctx, v, expr) },
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
	op.DoDebug = func() struct{} {
		for i := range s {
			s[i].DoDebug()
		}
		return struct{}{}
	}
	return op
}

// разные функции под разные типы -> мапа не подходит
func Cast[T int64 | float64](ctx *parser.CustomContext, t types.Basic, expr Op[T]) any {
	switch t {
	case types.SINT:
		return cast[int8, int64](ctx, expr)
	case types.INT:
		return cast[int16, int64](ctx, expr)
	case types.DINT:
		return cast[int32, int64](ctx, expr)
	case types.LINT:
		return cast[int64, int64](ctx, expr)
	case types.USINT:
		return cast[uint8, int64](ctx, expr)
	case types.UINT:
		return cast[uint16, int64](ctx, expr)
	case types.UDINT:
		return cast[uint32, int64](ctx, expr)
	case types.ULINT:
		return cast[uint64, int64](ctx, expr)
	case types.REAL:
		return cast[float32, float64](ctx, expr)
	case types.LREAL:
		return cast[float64, float64](ctx, expr)
	}
	return nil
}

func Arithmetic[Result, Left, Right NumberOpTypes](
	ctx *parser.CustomContext,
	op string,
	left Op[Left],
	right Op[Right],
	resultType types.Basic,
) Op[Result] {
	type K struct {
		op         string
		resultType types.Basic
	}

	ops := map[K]func(l Op[Left], r Op[Right]) Op[Result]{
		{"+", types.SINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return plus[int8, Result](ctx, l, r) },
		{"+", types.INT}:   func(l Op[Left], r Op[Right]) Op[Result] { return plus[int16, Result](ctx, l, r) },
		{"+", types.DINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return plus[int32, Result](ctx, l, r) },
		{"+", types.LINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return plus[int64, Result](ctx, l, r) },
		{"+", types.USINT}: func(l Op[Left], r Op[Right]) Op[Result] { return plus[uint8, Result](ctx, l, r) },
		{"+", types.UINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return plus[uint16, Result](ctx, l, r) },
		{"+", types.UDINT}: func(l Op[Left], r Op[Right]) Op[Result] { return plus[uint32, Result](ctx, l, r) },
		{"+", types.ULINT}: func(l Op[Left], r Op[Right]) Op[Result] { return plus[uint64, Result](ctx, l, r) },
		{"+", types.REAL}:  func(l Op[Left], r Op[Right]) Op[Result] { return plus[float32, Result](ctx, l, r) },
		{"+", types.LREAL}: func(l Op[Left], r Op[Right]) Op[Result] { return plus[float64, Result](ctx, l, r) },

		{"-", types.SINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return sub[int8, Result](ctx, l, r) },
		{"-", types.INT}:   func(l Op[Left], r Op[Right]) Op[Result] { return sub[int16, Result](ctx, l, r) },
		{"-", types.DINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return sub[int32, Result](ctx, l, r) },
		{"-", types.LINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return sub[int64, Result](ctx, l, r) },
		{"-", types.USINT}: func(l Op[Left], r Op[Right]) Op[Result] { return sub[uint8, Result](ctx, l, r) },
		{"-", types.UINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return sub[uint16, Result](ctx, l, r) },
		{"-", types.UDINT}: func(l Op[Left], r Op[Right]) Op[Result] { return sub[uint32, Result](ctx, l, r) },
		{"-", types.ULINT}: func(l Op[Left], r Op[Right]) Op[Result] { return sub[uint64, Result](ctx, l, r) },
		{"-", types.REAL}:  func(l Op[Left], r Op[Right]) Op[Result] { return sub[float32, Result](ctx, l, r) },
		{"-", types.LREAL}: func(l Op[Left], r Op[Right]) Op[Result] { return sub[float64, Result](ctx, l, r) },

		{"*", types.SINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return mult[int8, Result](ctx, l, r) },
		{"*", types.INT}:   func(l Op[Left], r Op[Right]) Op[Result] { return mult[int16, Result](ctx, l, r) },
		{"*", types.DINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return mult[int32, Result](ctx, l, r) },
		{"*", types.LINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return mult[int64, Result](ctx, l, r) },
		{"*", types.USINT}: func(l Op[Left], r Op[Right]) Op[Result] { return mult[uint8, Result](ctx, l, r) },
		{"*", types.UINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return mult[uint16, Result](ctx, l, r) },
		{"*", types.UDINT}: func(l Op[Left], r Op[Right]) Op[Result] { return mult[uint32, Result](ctx, l, r) },
		{"*", types.ULINT}: func(l Op[Left], r Op[Right]) Op[Result] { return mult[uint64, Result](ctx, l, r) },
		{"*", types.REAL}:  func(l Op[Left], r Op[Right]) Op[Result] { return mult[float32, Result](ctx, l, r) },
		{"*", types.LREAL}: func(l Op[Left], r Op[Right]) Op[Result] { return mult[float64, Result](ctx, l, r) },

		{"/", types.SINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return div[int8, Result](ctx, l, r) },
		{"/", types.INT}:   func(l Op[Left], r Op[Right]) Op[Result] { return div[int16, Result](ctx, l, r) },
		{"/", types.DINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return div[int32, Result](ctx, l, r) },
		{"/", types.LINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return div[int64, Result](ctx, l, r) },
		{"/", types.USINT}: func(l Op[Left], r Op[Right]) Op[Result] { return div[uint8, Result](ctx, l, r) },
		{"/", types.UINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return div[uint16, Result](ctx, l, r) },
		{"/", types.UDINT}: func(l Op[Left], r Op[Right]) Op[Result] { return div[uint32, Result](ctx, l, r) },
		{"/", types.ULINT}: func(l Op[Left], r Op[Right]) Op[Result] { return div[uint64, Result](ctx, l, r) },
		{"/", types.REAL}:  func(l Op[Left], r Op[Right]) Op[Result] { return div[float32, Result](ctx, l, r) },
		{"/", types.LREAL}: func(l Op[Left], r Op[Right]) Op[Result] { return div[float64, Result](ctx, l, r) },

		{"MOD", types.SINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return mod[int8, Result](ctx, l, r) },
		{"MOD", types.INT}:   func(l Op[Left], r Op[Right]) Op[Result] { return mod[int16, Result](ctx, l, r) },
		{"MOD", types.DINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return mod[int32, Result](ctx, l, r) },
		{"MOD", types.LINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return mod[int64, Result](ctx, l, r) },
		{"MOD", types.USINT}: func(l Op[Left], r Op[Right]) Op[Result] { return mod[uint8, Result](ctx, l, r) },
		{"MOD", types.UINT}:  func(l Op[Left], r Op[Right]) Op[Result] { return mod[uint16, Result](ctx, l, r) },
		{"MOD", types.UDINT}: func(l Op[Left], r Op[Right]) Op[Result] { return mod[uint32, Result](ctx, l, r) },
		{"MOD", types.ULINT}: func(l Op[Left], r Op[Right]) Op[Result] { return mod[uint64, Result](ctx, l, r) },
	}
	return ops[K{op, resultType}](left, right)
}

func Compare[Left, Right NumberOpTypes](
	ctx *parser.CustomContext,
	op string,
	left Op[Left],
	right Op[Right],
	resultType types.Basic,
) Bool {
	type K struct {
		op         string
		resultType types.Basic
	}

	ops := map[K]func(l Op[Left], r Op[Right]) Bool{
		{">", types.SINT}:  func(l Op[Left], r Op[Right]) Bool { return greater[int8](ctx, l, r) },
		{">", types.INT}:   func(l Op[Left], r Op[Right]) Bool { return greater[int16](ctx, l, r) },
		{">", types.DINT}:  func(l Op[Left], r Op[Right]) Bool { return greater[int32](ctx, l, r) },
		{">", types.LINT}:  func(l Op[Left], r Op[Right]) Bool { return greater[int64](ctx, l, r) },
		{">", types.USINT}: func(l Op[Left], r Op[Right]) Bool { return greater[uint8](ctx, l, r) },
		{">", types.UINT}:  func(l Op[Left], r Op[Right]) Bool { return greater[uint16](ctx, l, r) },
		{">", types.UDINT}: func(l Op[Left], r Op[Right]) Bool { return greater[uint32](ctx, l, r) },
		{">", types.ULINT}: func(l Op[Left], r Op[Right]) Bool { return greater[uint64](ctx, l, r) },
		{">", types.REAL}:  func(l Op[Left], r Op[Right]) Bool { return greater[float32](ctx, l, r) },
		{">", types.LREAL}: func(l Op[Left], r Op[Right]) Bool { return greater[float64](ctx, l, r) },

		{">=", types.SINT}:  func(l Op[Left], r Op[Right]) Bool { return greaterOrEqual[int8](ctx, l, r) },
		{">=", types.INT}:   func(l Op[Left], r Op[Right]) Bool { return greaterOrEqual[int16](ctx, l, r) },
		{">=", types.DINT}:  func(l Op[Left], r Op[Right]) Bool { return greaterOrEqual[int32](ctx, l, r) },
		{">=", types.LINT}:  func(l Op[Left], r Op[Right]) Bool { return greaterOrEqual[int64](ctx, l, r) },
		{">=", types.USINT}: func(l Op[Left], r Op[Right]) Bool { return greaterOrEqual[uint8](ctx, l, r) },
		{">=", types.UINT}:  func(l Op[Left], r Op[Right]) Bool { return greaterOrEqual[uint16](ctx, l, r) },
		{">=", types.UDINT}: func(l Op[Left], r Op[Right]) Bool { return greaterOrEqual[uint32](ctx, l, r) },
		{">=", types.ULINT}: func(l Op[Left], r Op[Right]) Bool { return greaterOrEqual[uint64](ctx, l, r) },
		{">=", types.REAL}:  func(l Op[Left], r Op[Right]) Bool { return greaterOrEqual[float32](ctx, l, r) },
		{">=", types.LREAL}: func(l Op[Left], r Op[Right]) Bool { return greaterOrEqual[float64](ctx, l, r) },

		{"<", types.SINT}:  func(l Op[Left], r Op[Right]) Bool { return less[int8](ctx, l, r) },
		{"<", types.INT}:   func(l Op[Left], r Op[Right]) Bool { return less[int16](ctx, l, r) },
		{"<", types.DINT}:  func(l Op[Left], r Op[Right]) Bool { return less[int32](ctx, l, r) },
		{"<", types.LINT}:  func(l Op[Left], r Op[Right]) Bool { return less[int64](ctx, l, r) },
		{"<", types.USINT}: func(l Op[Left], r Op[Right]) Bool { return less[uint8](ctx, l, r) },
		{"<", types.UINT}:  func(l Op[Left], r Op[Right]) Bool { return less[uint16](ctx, l, r) },
		{"<", types.UDINT}: func(l Op[Left], r Op[Right]) Bool { return less[uint32](ctx, l, r) },
		{"<", types.ULINT}: func(l Op[Left], r Op[Right]) Bool { return less[uint64](ctx, l, r) },
		{"<", types.REAL}:  func(l Op[Left], r Op[Right]) Bool { return less[float32](ctx, l, r) },
		{"<", types.LREAL}: func(l Op[Left], r Op[Right]) Bool { return less[float64](ctx, l, r) },

		{"<=", types.SINT}:  func(l Op[Left], r Op[Right]) Bool { return lessOrEqual[int8](ctx, l, r) },
		{"<=", types.INT}:   func(l Op[Left], r Op[Right]) Bool { return lessOrEqual[int16](ctx, l, r) },
		{"<=", types.DINT}:  func(l Op[Left], r Op[Right]) Bool { return lessOrEqual[int32](ctx, l, r) },
		{"<=", types.LINT}:  func(l Op[Left], r Op[Right]) Bool { return lessOrEqual[int64](ctx, l, r) },
		{"<=", types.USINT}: func(l Op[Left], r Op[Right]) Bool { return lessOrEqual[uint8](ctx, l, r) },
		{"<=", types.UINT}:  func(l Op[Left], r Op[Right]) Bool { return lessOrEqual[uint16](ctx, l, r) },
		{"<=", types.UDINT}: func(l Op[Left], r Op[Right]) Bool { return lessOrEqual[uint32](ctx, l, r) },
		{"<=", types.ULINT}: func(l Op[Left], r Op[Right]) Bool { return lessOrEqual[uint64](ctx, l, r) },
		{"<=", types.REAL}:  func(l Op[Left], r Op[Right]) Bool { return lessOrEqual[float32](ctx, l, r) },
		{"<=", types.LREAL}: func(l Op[Left], r Op[Right]) Bool { return lessOrEqual[float64](ctx, l, r) },

		{"=", types.SINT}:  func(l Op[Left], r Op[Right]) Bool { return equal[int8](ctx, l, r) },
		{"=", types.INT}:   func(l Op[Left], r Op[Right]) Bool { return equal[int16](ctx, l, r) },
		{"=", types.DINT}:  func(l Op[Left], r Op[Right]) Bool { return equal[int32](ctx, l, r) },
		{"=", types.LINT}:  func(l Op[Left], r Op[Right]) Bool { return equal[int64](ctx, l, r) },
		{"=", types.USINT}: func(l Op[Left], r Op[Right]) Bool { return equal[uint8](ctx, l, r) },
		{"=", types.UINT}:  func(l Op[Left], r Op[Right]) Bool { return equal[uint16](ctx, l, r) },
		{"=", types.UDINT}: func(l Op[Left], r Op[Right]) Bool { return equal[uint32](ctx, l, r) },
		{"=", types.ULINT}: func(l Op[Left], r Op[Right]) Bool { return equal[uint64](ctx, l, r) },
		{"=", types.REAL}:  func(l Op[Left], r Op[Right]) Bool { return equal[float32](ctx, l, r) },
		{"=", types.LREAL}: func(l Op[Left], r Op[Right]) Bool { return equal[float64](ctx, l, r) },

		{"<>", types.SINT}:  func(l Op[Left], r Op[Right]) Bool { return notEqual[int8](ctx, l, r) },
		{"<>", types.INT}:   func(l Op[Left], r Op[Right]) Bool { return notEqual[int16](ctx, l, r) },
		{"<>", types.DINT}:  func(l Op[Left], r Op[Right]) Bool { return notEqual[int32](ctx, l, r) },
		{"<>", types.LINT}:  func(l Op[Left], r Op[Right]) Bool { return notEqual[int64](ctx, l, r) },
		{"<>", types.USINT}: func(l Op[Left], r Op[Right]) Bool { return notEqual[uint8](ctx, l, r) },
		{"<>", types.UINT}:  func(l Op[Left], r Op[Right]) Bool { return notEqual[uint16](ctx, l, r) },
		{"<>", types.UDINT}: func(l Op[Left], r Op[Right]) Bool { return notEqual[uint32](ctx, l, r) },
		{"<>", types.ULINT}: func(l Op[Left], r Op[Right]) Bool { return notEqual[uint64](ctx, l, r) },
		{"<>", types.REAL}:  func(l Op[Left], r Op[Right]) Bool { return notEqual[float32](ctx, l, r) },
		{"<>", types.LREAL}: func(l Op[Left], r Op[Right]) Bool { return notEqual[float64](ctx, l, r) },
	}
	return ops[K{op, resultType}](left, right)
}

func If(ctx *parser.CustomContext, cond Bool, then_ Statement) Statement {
	op := Statement{
		ctx:     ctx,
		snippet: ctx.MakeSnippet(),
	}
	op.Do = func() struct{} {
		if cond.Do() {
			then_.Do()
		}
		return struct{}{}
	}
	op.DoDebug = func() struct{} {
		v := cond.DoDebug()
		msg := fmt.Sprintf("condition = %v", v)
		ctx.Debug(op.snippet, msg)
		if v {
			then_.DoDebug()
		}
		return struct{}{}
	}
	return op
}

func IfElse(ctx *parser.CustomContext, cond Bool, then_, else_ Statement) Statement {
	op := Statement{
		ctx:     ctx,
		snippet: ctx.MakeSnippet(),
	}
	op.Do = func() struct{} {
		if cond.Do() {
			then_.Do()
		} else {
			else_.Do()
		}
		return struct{}{}
	}
	op.DoDebug = func() struct{} {
		v := cond.DoDebug()
		msg := fmt.Sprintf("condition = %v", v)
		ctx.Debug(op.snippet, msg)
		if v {
			then_.DoDebug()
		} else {
			else_.DoDebug()
		}
		return struct{}{}
	}
	return op
}

func greater[T Number, Left, Right NumberOpTypes](
	ctx *parser.CustomContext,
	left Op[Left],
	right Op[Right],
) Bool {
	op := Bool{
		ResultType: types.BOOL,
		IsConstant: left.IsConstant && right.IsConstant,
		ctx:        ctx,
		snippet:    ctx.MakeSnippet(),
	}
	op.Do = func() bool {
		l := T(left.Do())
		r := T(right.Do())
		return l > r
	}
	op.DoDebug = func() bool {
		l := T(left.DoDebug())
		r := T(right.DoDebug())
		v := l > r
		msg := fmt.Sprintf("%v > %v: %v", l, r, v)
		ctx.Debug(op.snippet, msg)
		return v
	}
	return op
}

func greaterOrEqual[T Number, Left, Right NumberOpTypes](
	ctx *parser.CustomContext,
	left Op[Left],
	right Op[Right],
) Bool {
	op := Bool{
		ResultType: types.BOOL,
		IsConstant: left.IsConstant && right.IsConstant,
		ctx:        ctx,
		snippet:    ctx.MakeSnippet(),
	}
	op.Do = func() bool {
		l := T(left.Do())
		r := T(right.Do())
		return l >= r
	}
	op.DoDebug = func() bool {
		l := T(left.DoDebug())
		r := T(right.DoDebug())
		v := l >= r
		msg := fmt.Sprintf("%v >= %v: %v", l, r, v)
		ctx.Debug(op.snippet, msg)
		return v
	}
	return op
}

func less[T Number, Left, Right NumberOpTypes](
	ctx *parser.CustomContext,
	left Op[Left],
	right Op[Right],
) Bool {
	op := Bool{
		ResultType: types.BOOL,
		IsConstant: left.IsConstant && right.IsConstant,
		ctx:        ctx,
		snippet:    ctx.MakeSnippet(),
	}
	op.Do = func() bool {
		l := T(left.Do())
		r := T(right.Do())
		return l < r
	}
	op.DoDebug = func() bool {
		l := T(left.DoDebug())
		r := T(right.DoDebug())
		v := l < r
		msg := fmt.Sprintf("%v < %v: %v", l, r, v)
		ctx.Debug(op.snippet, msg)
		return v
	}
	return op
}

func lessOrEqual[T Number, Left, Right NumberOpTypes](
	ctx *parser.CustomContext,
	left Op[Left],
	right Op[Right],
) Bool {
	op := Bool{
		ResultType: types.BOOL,
		IsConstant: left.IsConstant && right.IsConstant,
		ctx:        ctx,
		snippet:    ctx.MakeSnippet(),
	}
	op.Do = func() bool {
		l := T(left.Do())
		r := T(right.Do())
		return l <= r
	}
	op.DoDebug = func() bool {
		l := T(left.DoDebug())
		r := T(right.DoDebug())
		v := l <= r
		msg := fmt.Sprintf("%v <= %v: %v", l, r, v)
		ctx.Debug(op.snippet, msg)
		return v
	}
	return op
}

func equal[T Number, Left, Right NumberOpTypes](
	ctx *parser.CustomContext,
	left Op[Left],
	right Op[Right],
) Bool {
	op := Bool{
		ResultType: types.BOOL,
		IsConstant: left.IsConstant && right.IsConstant,
		ctx:        ctx,
		snippet:    ctx.MakeSnippet(),
	}
	op.Do = func() bool {
		l := T(left.Do())
		r := T(right.Do())
		return l == r
	}
	op.DoDebug = func() bool {
		l := T(left.DoDebug())
		r := T(right.DoDebug())
		v := l == r
		msg := fmt.Sprintf("%v = %v: %v", l, r, v)
		ctx.Debug(op.snippet, msg)
		return v
	}
	return op
}

func notEqual[T Number, Left, Right NumberOpTypes](
	ctx *parser.CustomContext,
	left Op[Left],
	right Op[Right],
) Bool {
	op := Bool{
		ResultType: types.BOOL,
		IsConstant: left.IsConstant && right.IsConstant,
		ctx:        ctx,
		snippet:    ctx.MakeSnippet(),
	}
	op.Do = func() bool {
		l := T(left.Do())
		r := T(right.Do())
		return l != r
	}
	op.DoDebug = func() bool {
		l := T(left.DoDebug())
		r := T(right.DoDebug())
		v := l != r
		msg := fmt.Sprintf("%v <> %v: %v", l, r, v)
		ctx.Debug(op.snippet, msg)
		return v
	}
	return op
}

func assign[T Number, R int64 | float64](ctx *parser.CustomContext, variable types.Variable, expr Op[R]) Statement {
	op := Statement{
		ctx:     ctx,
		snippet: ctx.MakeSnippet(),
	}
	val := (*R)(variable.Pointer())

	op.Do = func() struct{} {
		v := T(expr.Do())
		*val = R(v)
		return struct{}{}
	}
	op.DoDebug = func() struct{} {
		v := T(expr.DoDebug())
		*val = R(v)
		msg := fmt.Sprintf("%v->%s", v, ctx.Name)
		ctx.Debug(op.snippet, msg)
		return struct{}{}
	}
	return op
}

func cast[T Number, Tout, Tin int64 | float64](ctx *parser.CustomContext, expr Op[Tin]) Op[Tout] {
	op := Op[Tout]{
		IsConstant: expr.IsConstant,
		ResultType: BasicType[T](),
		ctx:        ctx,
		snippet:    ctx.MakeSnippet(),
	}
	op.Do = func() Tout {
		v := T(expr.Do())
		return Tout(v)
	}
	op.DoDebug = func() Tout {
		v := T(expr.DoDebug())
		msg := fmt.Sprintf("%T(%v)", defaultValue[Tout](), v)
		ctx.Debug(op.snippet, msg)
		return Tout(v)
	}
	return op
}

func defaultValue[T any]() T { var v T; return v }

// T для ограничения размера инта
func plus[T Number, Result, Left, Right NumberOpTypes](
	ctx *parser.CustomContext,
	left Op[Left],
	right Op[Right],
) Op[Result] {
	op := Op[Result]{
		ResultType: BasicType[T](),
		IsConstant: left.IsConstant && right.IsConstant,
		ctx:        ctx,
		snippet:    ctx.MakeSnippet(),
	}
	op.Do = func() Result {
		l := T(left.Do())
		r := T(right.Do())
		return Result(l + r)
	}
	op.DoDebug = func() Result {
		l := T(left.DoDebug())
		r := T(right.DoDebug())
		v := Result(l + r)
		msg := fmt.Sprintf("%v + %v = %v", l, r, v)
		ctx.Debug(op.snippet, msg)
		return v
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

func sub[T Number, Result, Left, Right NumberOpTypes](
	ctx *parser.CustomContext,
	left Op[Left],
	right Op[Right],
) Op[Result] {
	op := Op[Result]{
		ResultType: BasicType[T](),
		IsConstant: left.IsConstant && right.IsConstant,
		ctx:        ctx,
		snippet:    ctx.MakeSnippet(),
	}
	op.Do = func() Result {
		l := T(left.Do())
		r := T(right.Do())
		return Result(l - r)
	}
	op.DoDebug = func() Result {
		l := T(left.DoDebug())
		r := T(right.DoDebug())
		v := Result(l - r)
		msg := fmt.Sprintf("%v - %v = %v", l, r, v)
		ctx.Debug(op.snippet, msg)
		return v
	}
	return op
}

func mult[T Number, Result, Left, Right NumberOpTypes](
	ctx *parser.CustomContext,
	left Op[Left],
	right Op[Right],
) Op[Result] {
	op := Op[Result]{
		ResultType: BasicType[T](),
		IsConstant: left.IsConstant && right.IsConstant,
		ctx:        ctx,
		snippet:    ctx.MakeSnippet(),
	}
	op.Do = func() Result {
		l := T(left.Do())
		r := T(right.Do())
		return Result(l * r)
	}
	op.DoDebug = func() Result {
		l := T(left.DoDebug())
		r := T(right.DoDebug())
		v := Result(l * r)
		msg := fmt.Sprintf("%v * %v = %v", l, r, v)
		ctx.Debug(op.snippet, msg)
		return v
	}
	return op
}

func div[T Number, Result, Left, Right NumberOpTypes](
	ctx *parser.CustomContext,
	left Op[Left],
	right Op[Right],
) Op[Result] {
	op := Op[Result]{
		ResultType: BasicType[T](),
		IsConstant: left.IsConstant && right.IsConstant,
		ctx:        ctx,
		snippet:    ctx.MakeSnippet(),
	}
	op.Do = func() Result {
		l := T(left.Do())
		r := T(right.Do())
		return Result(l / r)
	}
	op.DoDebug = func() Result {
		l := T(left.DoDebug())
		r := T(right.DoDebug())
		v := Result(l / r)
		msg := fmt.Sprintf("%v / %v = %v", l, r, v)
		ctx.Debug(op.snippet, msg)
		return v
	}
	return op
}

func mod[T constraints.Integer, Result, Left, Right NumberOpTypes](
	ctx *parser.CustomContext,
	left Op[Left],
	right Op[Right],
) Op[Result] {
	op := Op[Result]{
		ResultType: BasicType[T](),
		IsConstant: left.IsConstant && right.IsConstant,
		ctx:        ctx,
		snippet:    ctx.MakeSnippet(),
	}
	op.Do = func() Result {
		l := T(left.Do())
		r := T(right.Do())
		return Result(l % r)
	}
	op.DoDebug = func() Result {
		l := T(left.DoDebug())
		r := T(right.DoDebug())
		v := Result(l % r)
		msg := fmt.Sprintf("%v MOD %v = %v", l, r, v)
		ctx.Debug(op.snippet, msg)
		return v
	}
	return op
}
