package vm

import (
	"log"
	"math"

	"github.com/slonegd/go-st/stack"
	"github.com/slonegd/go-st/variant"
)

type (
	VM struct {
		// не байты, так как далее я буду в операторах делать ссылки на значения
		// но пока всё по книге, буду индексы на stack consts vars
		Bytecode Bytecode
		Consts   []variant.Variant // пока варианты, но потом лучше сделать просто гошные типы
		Vars     []variant.Variant // тоже лучше гошные типы
		VarNames map[int]string
		stack    stack.Stack[variant.Variant]
	}
	Op       uintptr
	Bytecode []uintptr
	Action   struct {
		Op
		Args []uintptr
	}
)

func New() *VM {
	return &VM{VarNames: make(map[int]string)}
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
	case IfFalse:
		c := x.stack.Pop()
		if !c.Bool() {
			return int(a.Args[0])
		}
	case Jump:
		return int(a.Args[0])
	case GtInt:
		right := x.stack.Pop()
		left := x.stack.Pop()
		x.stack.Push(variant.Greather(left, right))
	case GteInt:
		right := x.stack.Pop()
		left := x.stack.Pop()
		x.stack.Push(variant.GreatherOrEqual(left, right))
	case LtInt:
		right := x.stack.Pop()
		left := x.stack.Pop()
		x.stack.Push(variant.Less(left, right))
	case LteInt:
		right := x.stack.Pop()
		left := x.stack.Pop()
		x.stack.Push(variant.LessOrEqual(left, right))
	case EqInt:
		right := x.stack.Pop()
		left := x.stack.Pop()
		x.stack.Push(variant.Equal(left, right))
	case NeqInt:
		right := x.stack.Pop()
		left := x.stack.Pop()
		x.stack.Push(variant.NotEqual(left, right))
	}
	return 0
}

// TODO подумать о том, чтобы зафиксировать количество аргументов, тогда offset не нужен
// через бенчмарк
func (x Bytecode) getAction() (r Action, offset int) {
	op := Op(x[0])
	switch op {
	case PushConst:
		return Action{Op: op, Args: x[1:2]}, 2
	case PushVar:
		return Action{Op: op, Args: x[1:2]}, 2
	case PopVar:
		return Action{Op: op, Args: x[1:2]}, 2
	case PlusInt:
		return Action{Op: op}, 1
	case MinusInt:
		return Action{Op: op}, 1
	case MultInt:
		return Action{Op: op}, 1
	case DivInt:
		return Action{Op: op}, 1
	case ModInt:
		return Action{Op: op}, 1
	case IfFalse:
		return Action{Op: op, Args: x[1:2]}, 2
	case Jump:
		return Action{Op: op, Args: x[1:2]}, 2
	case GtInt, GteInt, LtInt, LteInt, EqInt, NeqInt:
		return Action{Op: op}, 1
	default:
		return Action{}, math.MaxInt
	}
}

func (x *Bytecode) AddOp(op Op, args ...uintptr) int {
	i := len(*x)
	*x = append(*x, uintptr(op))
	for _, a := range args {
		*x = append(*x, a)
	}
	return i
}

func (x *Bytecode) ChangeOpArgs(index int, args ...uintptr) {
	for i, v := range args {
		(*x)[index+1+i] = v
	}
}

//go:generate go run .././vendor/golang.org/x/tools/cmd/stringer/stringer.go -type=Op
const (
	PushConst Op = iota // положить константу на стек
	PushVar             // положить переменную на стек
	PopVar              // забрать переменную со стека
	PlusInt             // сумма 2 интов на стеке
	MinusInt            // разность...
	MultInt             // переменожение...
	DivInt              // деление...
	ModInt              // остаток от деления...
	IfFalse             // если на стеке false перейти по метке
	Jump                // перейти по метке
	GtInt               // оператор больше для 2 интов на стеке
	GteInt              // оператор больше либо равно для 2 интов на стеке
	LtInt               // оператор меньше для 2 интов на стеке
	LteInt              // оператор меньше для 2 интов на стеке
	EqInt               // оператор равенства для 2 интов на стеке
	NeqInt              // оператор неравенства для 2 интов на стеке
	// не забыть: getAction ExecuteOne Print
)

func (x *VM) Print() {
	for i := 0; i < len(x.Bytecode); {
		a, offset := x.Bytecode[i:].getAction()
		switch a.Op {
		case PushConst:
			log.Printf("%04d: %s %v", i, a.Op, x.Consts[a.Args[0]])
		case PushVar, PopVar:
			log.Printf("%04d: %s %s", i, a.Op, x.VarNames[int(a.Args[0])])
		case PlusInt, MinusInt, MultInt, DivInt, ModInt, GtInt, GteInt, LtInt, LteInt, EqInt, NeqInt:
			log.Printf("%04d: %s", i, a.Op)
		case IfFalse, Jump:
			log.Printf("%04d: %s %v", i, a.Op, a.Args[0])
		}
		i += offset
	}
}
