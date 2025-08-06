package bench

import (
	"testing"

	"github.com/Shopify/go-lua"
	"github.com/stretchr/testify/assert"
)

const luaScript = `
function init()
    n = 0
    r = 0
end

function fibonacci(n)
    if n <= 1 then
        return n
    end
    local a, b = 0, 1
    for _ = 2, n do
        a, b = b, a + b
    end
    return b
end

function exec()
    n = n + 1
    if n > 100 then
        n = 100
    end
    r = fibonacci(n)
end
`

func Benchmark_go_lua(b *testing.B) {
	l := lua.NewState()

	// Загружаем скрипт
	assert.NoError(b, lua.DoString(l, luaScript))

	// Инициализируем
	l.Global("init")
	l.Call(0, 0)

	b.ResetTimer()
	for b.Loop() {
		// Вызываем exec()
		l.Global("exec")
		l.Call(0, 0)
	}
}

func Test_go_lua(t *testing.T) {
	l := lua.NewState()

	// Загружаем скрипт
	assert.NoError(t, lua.DoString(l, luaScript))

	// Инициализируем
	l.Global("init")
	l.Call(0, 0)

	// Выполняем exec и проверяем результаты
	execAndCheck := func(expectedR int) {
		l.Global("exec")
		l.Call(0, 0)

		// Читаем глобальную `r` из Lua в Go
		l.Global("r")
		r, _ := l.ToInteger(-1)
		l.Pop(1) // Убираем прочитанное значение из стека

		assert.Equal(t, expectedR, r)
	}

	execAndCheck(1)
	execAndCheck(1)
	execAndCheck(2)
	execAndCheck(3)
	execAndCheck(5)
}
