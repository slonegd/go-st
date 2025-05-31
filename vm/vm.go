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
	Action   [2]uintptr
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
			return int(a[1])
		}
	case Jump:
		return int(a[1])
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

func (x *Bytecode) AddOp(op Op, args ...uintptr) int {
	a := [2]uintptr{uintptr(op), 0}
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
	for i, a := range x.Bytecode {
		op := Op(a[0])
		switch op {
		case PushConst:
			log.Printf("%04d: %s %v", i, op, x.Consts[a[1]])
		case PushVar, PopVar:
			log.Printf("%04d: %s %s", i, op, x.VarNames[int(a[1])])
		case PlusInt, MinusInt, MultInt, DivInt, ModInt, GtInt, GteInt, LtInt, LteInt, EqInt, NeqInt:
			log.Printf("%04d: %s", i, op)
		case IfFalse, Jump:
			log.Printf("%04d: %s %v", i, op, a[1])
		}
	}
}
