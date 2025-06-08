package optimizer

import (
	"testing"

	"github.com/slonegd/go-st/variant"
	"github.com/slonegd/go-st/vm"
	"github.com/stretchr/testify/assert"
)

func Test_plus(t *testing.T) {
	x := Optimizer{in: &vm.VM{
		Consts: []variant.Variant{variant.NewAnyVariant("42")},
		Vars:   []variant.Variant{variant.IntVariant(7)},
	}}

	tests := map[string]struct {
		left       stackValue
		right      stackValue
		wantAction vm.Action
		wantValue  stackValue
	}{
		"INT var + any const": {
			left:       stackValue{variant.INT, fromVar, 0},
			right:      stackValue{variant.ANY, fromConst, 0},
			wantAction: vm.Action{uintptr(vm.Plus_i16_p_c), x.in.Vars[0].Pointer(), 42},
			wantValue:  stackValue{variant.INT, fromStack, 0},
		},
	}
	for name, tt := range tests {
		gotAction, gotValue := x.plus(tt.left, tt.right)
		assert.Equal(t, tt.wantAction, gotAction, name)
		assert.Equal(t, tt.wantValue, gotValue, name)
	}
}
