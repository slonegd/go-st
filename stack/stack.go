package stack

// дженерики не замедляют, проверил бенчмарком
type Stack[T any] []T

func (x *Stack[T]) Push(v T) {
	*x = append(*x, v)
}

func (x *Stack[T]) Pop() T {
	size := len(*x)
	r := (*x)[size-1]
	*x = (*x)[:size-1]
	return r
}

func (x *Stack[T]) ChangeLast(cb func(T) T) {
	if len(*x) == 0 {
		x.Push(defaultValue[T]())
	}
	i := len(*x) - 1
	(*x)[i] = cb((*x)[i])
}

func defaultValue[T any]() T { var v T; return v }
