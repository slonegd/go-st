package bench

import (
	_ "embed"
	"testing"

	"github.com/d5/tengo/v2"
	"github.com/stretchr/testify/assert"
)

const arithmetic_tengo = `
	a = 15 + j
	j = 5 + 2*a*(30-a)/5
	k = k + -2

	i = i + 1
	if i%5 == 0 {
		divisor = 5
		if i/5 == 1 {
			mult = 1
		} else if i/5 == 2 {
			mult = 2
		} else {
			mult = 3
		}
	} else if i%4 == 0 {
		divisor = 4
		mult = i / 4
	} else if i%3 == 0 {
		divisor = 3
		mult = i / 3
	} else if i%2 == 0 {
		divisor = 2
		mult = i / 2
	} else if i%1 == 0 {
		divisor = 1
		mult = 1
	}

	if mult != 1 {
		test = mult * divisor
	}
`

func Benchmark_tengo_001(b *testing.B) {
	script := tengo.NewScript([]byte(arithmetic_tengo))
	_ = script.Add("a", 0)
	_ = script.Add("j", 0)
	_ = script.Add("k", 42)
	_ = script.Add("i", 0)
	_ = script.Add("divisor", 0)
	_ = script.Add("mult", 0)
	_ = script.Add("test", 0)
	compiled, _ := script.Compile()

	b.ResetTimer()
	for range iterations {
		compiled.Run()
	}
}

func Test_tengo_001(t *testing.T) {
	script := tengo.NewScript([]byte(arithmetic_tengo))
	_ = script.Add("a", 0)
	_ = script.Add("j", 0)
	_ = script.Add("k", 42)
	_ = script.Add("i", 0)
	_ = script.Add("divisor", 0)
	_ = script.Add("mult", 0)
	_ = script.Add("test", 0)

	compiled, err := script.Compile()
	assert.NoError(t, err)

	err = compiled.Run()
	assert.NoError(t, err)
	assert.Equal(t, 1, compiled.Get("i").Int())

	err = compiled.Run()
	assert.Equal(t, 2, compiled.Get("i").Int())

	err = compiled.Run()
	err = compiled.Run()
	err = compiled.Run()
	err = compiled.Run()
	err = compiled.Run()
	assert.Equal(t, 7, compiled.Get("i").Int())
	assert.Equal(t, 1, compiled.Get("divisor").Int())
	assert.Equal(t, 1, compiled.Get("mult").Int())
	assert.Equal(t, 6, compiled.Get("test").Int())

}
