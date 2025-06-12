package bench

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"

	lua "github.com/Shopify/go-lua"
)

const arithmetic_lua_init = `
	a = 0
	j = 0
	k = 42
	
	i       = 0
	divisor = 0
	mult    = 0
	test    = 0
`

const arithmetic_lua = `
	a = 15 + j
	j = 5 + 2*a*(30-a)/5
	k = k + -2

	i = i+1
	if i % 5 == 0 then
		divisor = 5
		if i / 5 == 1 then
			mult = 1
		elseif i / 5 == 2 then
			mult = 2
		else
			mult = 3
		end
	elseif i % 4 == 0 then
		divisor = 4
		mult = i / 4
	elseif i % 3 == 0 then
		divisor = 3
		mult = i / 3
	elseif i % 2 == 0 then
		divisor = 2
		mult = i / 2
	elseif i % 1 == 0 then
		divisor = 1
		mult = 1
	end

	if mult ~= 1 then
		test = mult * divisor
	end
`

func Benchmark_lua_001(b *testing.B) {
	l := lua.NewState()

	lua.DoString(l, arithmetic_lua_init)
	lua.LoadString(l, arithmetic_lua)
	l.SetGlobal("exec")

	b.ResetTimer()
	for range iterations {
		l.Global("exec")
		l.Call(0, 0)
	}
}

func Test_lua_001(t *testing.T) {
	l := lua.NewState()

	err := lua.DoString(l, arithmetic_lua_init)
	assert.NoError(t, err)

	err = lua.LoadString(l, arithmetic_lua)
	assert.NoError(t, err)

	l.SetGlobal("exec")

	l.Global("exec")
	l.Call(0, 0)

	// Читаем глобальную `i` из Lua в Go
	l.Global("i")
	i, _ := l.ToInteger(-1)
	l.Pop(1) // Убираем прочитанное значение из стека

	assert.Equal(t, 1, i)

	l.Global("exec")
	l.Call(0, 0)

	// Читаем глобальную `i` из Lua в Go
	l.Global("i")
	i, _ = l.ToInteger(-1)
	l.Pop(1) // Убираем прочитанное значение из стека

	assert.Equal(t, 2, i)
}
