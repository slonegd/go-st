package ast

type (
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
