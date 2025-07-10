package bench

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/traefik/yaegi/interp"
)

func Benchmark_yaegi_001(b *testing.B) {
	i := interp.New(interp.Options{})

	code := arithmetic_go
	i.Eval(code)

	cb, _ := i.Eval("bench.exec")
	exec := cb.Interface().(func())

	for b.Loop() {
		exec()
	}
}

//go:embed 002_arithmetic_if_while.go
var arithmetic_go string

func Test_yaegi_001(t *testing.T) {
	i := interp.New(interp.Options{})

	code := arithmetic_go
	_, err := i.Eval(code)
	assert.NoError(t, err)

	cb, err := i.Eval("bench.exec")
	assert.NoError(t, err)
	exec := cb.Interface().(func())

	exec()
	v, _ := i.Eval("bench.a")
	assert.Equal(t, 15, v.Interface().(int))
	exec()

	v, _ = i.Eval("bench.i")
	assert.Equal(t, 2, v.Interface().(int))

	exec()
	assert.Equal(t, 3, v.Interface().(int))
}
