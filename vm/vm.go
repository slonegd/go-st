package vm

import (
	"math"

	"github.com/slonegd/go-st/variant"
)

type (
	VM struct {
		// не байты, так как далее я буду в операторах делать ссылки на значения
		// но пока всё по книге, буду индексы на stack consts vars
		Bytecode Bytecode
		Consts   []variant.Variant // пока варианты, но потом лучше сделать просто гошные типы
		Vars     []variant.Variant // тоже лучше гошные типы
		stack    Stack
	}
	Op       uintptr
	Bytecode []uintptr
	Action   struct {
		Op
		Args []uintptr
	}
	Stack []variant.Variant
)

func New() *VM {
	return &VM{}
}

func (x *VM) Execute() {
	for i := 0; i < len(x.Bytecode); {
		action, offset := x.Bytecode[i:].getAction()
		jump := x.ExecuteOne(action)
		if jump != 0 {
			i = jump
		} else {
			i += offset
		}
		if offset > len(x.Bytecode) {
			return
		}
	}
}

func (x *VM) ExecuteOne(a Action) (jump int) {
	switch a.Op {
	case PushConst:
		x.stack.Push(x.Consts[a.Args[0]])
	case PushVar:
		x.stack.Push(x.Vars[a.Args[0]])
	case PopVar:
		x.Vars[a.Args[0]] = x.stack.Pop()
	case PlusInt:
		right := x.stack.Pop()
		left := x.stack.Pop()
		x.stack.Push(variant.Plus(left, right)) // лучше сумму без вариантов
	case MinusInt:
		right := x.stack.Pop()
		left := x.stack.Pop()
		x.stack.Push(variant.Minus(left, right))
	case MultInt:
		right := x.stack.Pop()
		left := x.stack.Pop()
		x.stack.Push(variant.Mult(left, right))
	case DivInt:
		right := x.stack.Pop()
		left := x.stack.Pop()
		x.stack.Push(variant.Divide(left, right))
	case ModInt:
		right := x.stack.Pop()
		left := x.stack.Pop()
		x.stack.Push(variant.Mod(left, right))
	}
	return 0
}

// TODO подумать о том, чтобы зафиксировать количество аргументов, тогда offset не нужен
// через бенчмарк
func (x Bytecode) getAction() (r Action, offset int) {
	switch Op(x[0]) {
	case PushConst:
		return Action{Op: PushConst, Args: x[1:2]}, 2
	case PushVar:
		return Action{Op: PushVar, Args: x[1:2]}, 2
	case PopVar:
		return Action{Op: PopVar, Args: x[1:2]}, 2
	case PlusInt:
		return Action{Op: PlusInt}, 1
	case MinusInt:
		return Action{Op: MinusInt}, 1
	case MultInt:
		return Action{Op: MultInt}, 1
	case DivInt:
		return Action{Op: DivInt}, 1
	case ModInt:
		return Action{Op: DivInt}, 1
	default:
		return Action{}, math.MaxInt
	}
}

func (x *Bytecode) AddOp(op Op, args ...uintptr) {
	*x = append(*x, uintptr(op))
	for _, a := range args {
		*x = append(*x, a)
	}
}

const (
	PushConst Op = iota
	PushVar
	PopVar
	PlusInt
	MinusInt
	MultInt
	DivInt
	ModInt
	// не забыть: getAction ExecuteOne
)

func (x *Stack) Push(v variant.Variant) {
	*x = append(*x, v)
}

func (x *Stack) Pop() variant.Variant {
	// TODO наверное лучше через указатель на стек (проверить бенчмарком!)
	size := len(*x)
	r := (*x)[size-1]
	*x = (*x)[:size-1]
	return r
}
