package bench

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/traefik/yaegi/interp"
)

func Benchmark_yaegi_001(b *testing.B) {
	i := interp.New(interp.Options{})

	code := fibo_code
	i.Eval(code)

	cb, _ := i.Eval("bench.exec")
	exec := cb.Interface().(func())

	for b.Loop() {
		exec()
	}
}

//go:embed fibo.go
var fibo_code string

func Test_yaegi_fibo(t *testing.T) {
	i := interp.New(interp.Options{})

	code := fibo_code
	_, err := i.Eval(code)
	assert.NoError(t, err)

	cb, err := i.Eval("bench.exec")
	assert.NoError(t, err)
	exec := cb.Interface().(func())

	exec()
	v, _ := i.Eval("bench.fibo5")
	assert.Equal(t, 5, v.Interface().(int))
	v, _ = i.Eval("bench.fibo10")
	assert.Equal(t, 55, v.Interface().(int))
}
