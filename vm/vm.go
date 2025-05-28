package vm

import (
	"math"

	"github.com/slonegd/go-st/variant"
)

type (
	VM struct {
		// не байты, так как далее я буду в операторах делать ссылки на значения
		// но пока всё по книге, буду индексы на stack consts vars
		bytecode Bytecode
		consts   []variant.Variant // пока варианты, но потом лучше сделать просто гошные типы
		vars     []variant.Variant // тоже лучше гошные типы
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
	for i := 0; i < len(x.bytecode); {
		action, offset := x.bytecode[i:].getAction()
		jump := x.ExecuteOne(action)
		if jump != 0 {
			i = jump
		} else {
			i += offset
		}
		if offset > len(x.bytecode) {
			return
		}
	}
}

func (x *VM) ExecuteOne(a Action) (jump int) {
	switch a.Op {
	case PushConst:
		x.stack.Push(x.consts[a.Args[0]])
	case PushVar:
		x.stack.Push(x.vars[a.Args[0]])
	case PopVar:
		x.vars[a.Args[0]] = x.stack.Pop()
	case SumInt:
		left := x.stack.Pop()
		right := x.stack.Pop()
		x.stack.Push(variant.Plus(left, right)) // лучше сумму без вариантов
	}
	return 0
}

func (x Bytecode) getAction() (r Action, offset int) {
	switch Op(x[0]) {
	case PushConst:
		return Action{Op: PushConst, Args: x[1:2]}, 2
	case PushVar:
		return Action{Op: PushVar, Args: x[1:2]}, 2
	case PopVar:
		return Action{Op: PopVar, Args: x[1:2]}, 2
	case SumInt:
		return Action{Op: SumInt}, 1
	default:
		return Action{}, math.MaxInt
	}
}

const (
	PushConst Op = iota
	PushVar
	PopVar
	SumInt
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
