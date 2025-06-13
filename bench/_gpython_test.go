package bench

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-python/gpython/py"
	_ "github.com/go-python/gpython/stdlib" // там надо вызвать init
)

const arithmetic_gpython = `
a = 15 + j
j = 5 + 2 * a * (30 - a) // 5
k = k + -2

i += 1

if i % 5 == 0:
    divisor = 5
    if i // 5 == 1:
        mult = 1
    elif i // 5 == 2:
        mult = 2
    else:
        mult = 3
elif i % 4 == 0:
    divisor = 4
    mult = i // 4
elif i % 3 == 0:
    divisor = 3
    mult = i // 3
elif i % 2 == 0:
    divisor = 2
    mult = i // 2
elif i % 1 == 0:
    divisor = 1
    mult = 1

if mult != 1:
    test = mult * divisor
`

func Benchmark_gpython_001(b *testing.B) {
	// 1. Создаём пустое глобальное окружение
	globals := py.NewStringDict()

	// 2. Инициализируем переменные
	globals["a"] = py.Int(0)
	globals["j"] = py.Int(0)
	globals["k"] = py.Int(42)

	globals["i"] = py.Int(0)
	globals["divisor"] = py.Int(0)
	globals["mult"] = py.Int(0)
	globals["test"] = py.Int(0)

	// 3. Компилируем скрипт один раз
	script := arithmetic_gpython
	code, _ := py.Compile(script, "<script>", py.ExecMode, 0, false)
	ctx := py.NewContext(py.DefaultContextOpts())

	b.ResetTimer()
	for range iterations {
		ctx.RunCode(code, globals, globals, nil)
	}
}

func Test_gpython_001(t *testing.T) {
	// 1. Создаём пустое глобальное окружение
	globals := py.NewStringDict()

	// 2. Инициализируем переменные
	globals["a"] = py.Int(0)
	globals["j"] = py.Int(0)
	globals["k"] = py.Int(42)

	globals["i"] = py.Int(0)
	globals["divisor"] = py.Int(0)
	globals["mult"] = py.Int(0)
	globals["test"] = py.Int(0)

	// 3. Компилируем скрипт один раз
	script := arithmetic_gpython
	code, err := py.Compile(script, "<script>", py.ExecMode, 0, false)
	assert.NoError(t, err)

	ctx := py.NewContext(py.DefaultContextOpts())
	_, err = ctx.RunCode(code, globals, globals, nil)
	assert.NoError(t, err)

	i := globals["i"].(py.Int)
	assert.Equal(t, 1, int(i))

	_, err = ctx.RunCode(code, globals, globals, nil)
	assert.NoError(t, err)

	i = globals["i"].(py.Int)
	assert.Equal(t, 2, int(i))
}
