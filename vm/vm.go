package vm

import (
	"log"

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
	Bytecode []Action
	Action   [3]uintptr
)

func New() *VM {
	return &VM{VarNames: make(map[int]string)}
}

func (x *VM) Execute() {
	for i := 0; i < len(x.Bytecode); {
		action := x.Bytecode[i]
		jump := x.ExecuteOne(action)
		if jump != 0 {
			i = jump
		} else {
			i++
		}
	}
}

func (x *VM) ExecuteOne(a Action) (jump int) {
	op := Op(a[0])
	switch op {
	case PushConst:
		x.stack.Push(x.Consts[a[1]])
	case PushVar:
		x.stack.Push(x.Vars[a[1]])
	case PopVar:
		x.Vars[a[1]] = x.stack.Pop()
	case Plus:
		right := x.stack.Pop()
		left := x.stack.Pop()
		x.stack.Push(variant.Plus(left, right)) // лучше сумму без вариантов
	case Sub:
		right := x.stack.Pop()
		left := x.stack.Pop()
		x.stack.Push(variant.Sub(left, right))
	case Mult:
		right := x.stack.Pop()
		left := x.stack.Pop()
		x.stack.Push(variant.Mult(left, right))
	case Div:
		right := x.stack.Pop()
		left := x.stack.Pop()
		x.stack.Push(variant.Divide(left, right))
	case Mod:
		right := x.stack.Pop()
		left := x.stack.Pop()
		x.stack.Push(variant.Mod(left, right))
	case IfFalse:
		c := x.stack.Pop()
		if !c.Bool() {
			return int(a[1])
		}
	case Jump:
		return int(a[1])
	case Gt:
		right := x.stack.Pop()
		left := x.stack.Pop()
		x.stack.Push(variant.Greather(left, right))
	case Gte:
		right := x.stack.Pop()
		left := x.stack.Pop()
		x.stack.Push(variant.GreatherOrEqual(left, right))
	case Lt:
		right := x.stack.Pop()
		left := x.stack.Pop()
		x.stack.Push(variant.Less(left, right))
	case Lte:
		right := x.stack.Pop()
		left := x.stack.Pop()
		x.stack.Push(variant.LessOrEqual(left, right))
	case Eq:
		right := x.stack.Pop()
		left := x.stack.Pop()
		x.stack.Push(variant.Equal(left, right))
	case Neq:
		right := x.stack.Pop()
		left := x.stack.Pop()
		x.stack.Push(variant.NotEqual(left, right))
	}
	return 0
}

func (x *Bytecode) AddOp(op Op, args ...uintptr) int {
	a := [3]uintptr{uintptr(op), 0, 0}
	for i, arg := range args {
		a[i+1] = arg
	}
	i := len(*x)
	*x = append(*x, a)
	return i
}

func (x *Bytecode) ChangeOpArgs(index int, args ...uintptr) {
	for i, v := range args {
		(*x)[index][i+1] = v
	}
}

//go:generate go run .././vendor/golang.org/x/tools/cmd/stringer/stringer.go -type=Op
const (
	PushConst Op = iota // положить константу на стек
	PushVar             // положить переменную на стек
	PopVar              // забрать переменную со стека

	Plus         // сумма 2 вариантов (обобщённый тип) на стеке
	Plus_i16_p_c // сумма int16, левый по указателю на переменную, второй аргументом
	Sub          // разность...
	Mult         // переменожение...
	Div          // деление...
	Mod          // остаток от деления...

	Gt  // оператор больше для 2 интов на стеке
	Gte // оператор больше либо равно для 2 интов на стеке
	Lt  // оператор меньше для 2 интов на стеке
	Lte // оператор меньше для 2 интов на стеке
	Eq  // оператор равенства для 2 интов на стеке
	Neq // оператор неравенства для 2 интов на стеке

	IfFalse // если на стеке false перейти по метке
	Jump    // перейти по метке
	// не забыть: ExecuteOne Print
)

func (x *VM) Print() {
	for i, a := range x.Bytecode {
		op := Op(a[0])
		switch op {
		case PushConst:
			log.Printf("%04d: %s %v", i, op, x.Consts[a[1]])
		case PushVar, PopVar:
			log.Printf("%04d: %s %s", i, op, x.VarNames[int(a[1])])
		case Plus, Sub, Mult, Div, Mod, Gt, Gte, Lt, Lte, Eq, Neq:
			log.Printf("%04d: %s", i, op)
		case IfFalse, Jump:
			log.Printf("%04d: %s %v", i, op, a[1])
		}
	}
}

func (a Action) Op() Op {
	return Op(a[0])
}
