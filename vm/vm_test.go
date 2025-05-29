package vm

import (
	"testing"

	"github.com/slonegd/go-st/variant"
	"github.com/stretchr/testify/assert"
)

func TestVM_Execute(t *testing.T) {
	assert := assert.New(t)
	vm := New()
	// условно x1 := 42; x2 := 100500; x2 := x1+x2;
	vm.Bytecode = []uintptr{
		uintptr(PushConst), 0,
		uintptr(PopVar), 0, // индекс в vars
		uintptr(PushConst), 1,
		uintptr(PopVar), 1, // индекс в vars
		uintptr(PushVar), 0,
		uintptr(PushVar), 1, // стек [42,100500]
		uintptr(PlusInt),   // стек [100542]
		uintptr(PopVar), 1, /// индекс в vars
	}
	vm.Consts = []variant.Variant{variant.NewAnyVariant("42"), variant.NewAnyVariant("100500")}
	vm.Vars = []variant.Variant{nil, nil} // заполняютя программой
	vm.Execute()
	assert.Equal(int64(42), vm.Vars[0].Int())
	assert.Equal(int64(100542), vm.Vars[1].Int())
}
