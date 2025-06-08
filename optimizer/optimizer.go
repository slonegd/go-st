package optimizer

import (
	"github.com/slonegd/go-st/stack"
	"github.com/slonegd/go-st/variant"
	"github.com/slonegd/go-st/vm"
)

type Optimizer struct {
	in *vm.VM
}

func New(in *vm.VM) *Optimizer {
	return &Optimizer{in: in}
}

// идея в том, чтобы в байткоде заменить обощённые операторы (например Plus) на типизированные Plus_i16_p_c
// при этом происходит проверка типов, что можно над ними выполнять эти действия
func (x *Optimizer) ReplaceVariantOps() {
	var stack stack.Stack[stackValue]
	out := &vm.VM{}
	*out = *x.in // копия
	out.Bytecode = nil
	for i := 0; i < len(x.in.Bytecode); {
		action := x.in.Bytecode[i]
		switch action.Op() {
		// кладём всё на стек, чтоб определить типы
		case vm.PushConst:
			stack.Push(stackValue{x.in.Consts[action[1]].Type(), fromConst, int(action[1])})
			out.Bytecode = append(out.Bytecode, action)
		case vm.PushVar:
			stack.Push(stackValue{x.in.Vars[action[1]].Type(), fromVar, int(action[1])})
			out.Bytecode = append(out.Bytecode, action)

		// тестирую подход пока на одном операторе
		case vm.Plus:
			right := stack.Pop()
			left := stack.Pop()
			// TODO понять, что тут
			// у переменных тип фиксирован
			// остальное должно приводиться
			// если оба константы, то вооще можно заменить на расчёт,
			// но пока тип определяется по результату (минимальный размер)
			// TODO возможно в стандарте типы констант заданы явно, посмотреть
			// на стеке и в константах лежат int64, но типом  ограничено
			// если на стек попалась сумма, где была переменная, то тип уже зафиксирован
			// примеры операторов:
			// Plus_i16_p_c - сумма типов int16(INT), левое берётся из аргумента (const), второе по указателю (var)
			// Plus_i32_c_s - сумма типов int32(DINT), левое лежит на uintptr стеке, второе константа
			// проритеты p c s суффиксов в операторе:
			//  - p всегда первое, если есть
			//  - s всегда последнее, так как для него не надо ничего в аргумент
			a, v := x.plus(left, right)
			stack.Push(v)
			out.Bytecode = append(out.Bytecode, a) // TODO удалить те опретаторы, что положили на стек

		default:
			out.Bytecode = append(out.Bytecode, action)
		}
	}
}

type (
	stackValue struct {
		variant.Type
		source source
		index  int
	}
	source int
)

const (
	fromStack source = iota
	fromConst
	fromVar
)

func (x *Optimizer) plus(left, right stackValue) (vm.Action, stackValue) {
	var (
		r                   stackValue
		pointers, constants []uintptr
	)
	for _, arg := range []stackValue{left, right} {
		switch arg.source {
		case fromVar:
			v := x.in.Vars[arg.index]
			pointers = append(pointers, v.Pointer())
			r.Type = arg.Type // TODO проверка типа
		case fromConst:
			v := x.in.Consts[arg.index]
			constants = append(constants, uintptr(v.Int()))
		}
	}
	// TODO аналогично и для других типов
	switch {
	case r.Type == variant.INT && len(pointers) == 2:
		return vm.Action{uintptr(vm.Plus_i16_p_p), pointers[0], pointers[1]}, r
	case r.Type == variant.INT && len(pointers) == 1 && len(constants) == 1:
		return vm.Action{uintptr(vm.Plus_i16_p_c), pointers[0], constants[0]}, r
	case r.Type == variant.INT && len(constants) == 2:
		return vm.Action{uintptr(vm.Plus_i16_c_c), constants[0], constants[1]}, r
	case r.Type == variant.INT && len(pointers) == 1:
		return vm.Action{uintptr(vm.Plus_i16_p_s), pointers[0], 0}, r
	case r.Type == variant.INT && len(constants) == 1:
		return vm.Action{uintptr(vm.Plus_i16_c_s), 0, 0}, r
	default:
		return vm.Action{uintptr(vm.Plus_i16_s_s), 0, 0}, r
	}
}
