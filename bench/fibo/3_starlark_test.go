package bench

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.starlark.net/lib/json"
	"go.starlark.net/starlark"
	"go.starlark.net/starlarktest"
	"go.starlark.net/syntax"
)

// в starlark нельзя изменять глобальные переменные, поэтому используется словарь
const starlark_fibo = `
def fibonacci(n):
    if n <= 1:
        return n
    a, b = 0, 1
    for _ in range(2, n+1):
        a, b = b, a + b
    return b

def init():
    # Создаём состояние с начальными значениями
    state = {"r": 0, "n": 0}
    return state

def exec(state):
    # Обновляем переданное состояние
    state["n"] = state["n"] + 1
    if state["n"] > 100:
        state["n"] = 100
    state["r"] = fibonacci(state["n"])
`

func Benchmark_starlark_fibo(b *testing.B) {
	script := starlark_fibo

	thread := &starlark.Thread{}
	globals, err := starlark.ExecFile(thread, "fibo.st", script, nil)
	if err != nil {
		b.Fatal(err)
	}

	initFn, ok := globals["init"].(*starlark.Function)
	if !ok {
		b.Fatal("init function not found")
	}
	execFn, ok := globals["exec"].(*starlark.Function)
	if !ok {
		b.Fatal("exec function not found")
	}

	// Создаём состояние один раз
	result, err := starlark.Call(thread, initFn, starlark.Tuple{}, nil)
	if err != nil {
		b.Fatal(err)
	}
	stateDict, ok := result.(*starlark.Dict)
	if !ok {
		b.Fatal("init should return dict")
	}

	b.ResetTimer()
	for b.Loop() {
		starlark.Call(thread, execFn, starlark.Tuple{stateDict}, nil)
	}
}

func Test_starlark_fibo(t *testing.T) {
	script := starlark_fibo
	starlark.Universe["json"] = json.Module
	testdata := starlarktest.DataFile("starlark", ".")
	thread := new(starlark.Thread)
	filename := filepath.Join(testdata, "fibo")
	opts := getOptions(script)

	globals, err := starlark.ExecFileOptions(opts, thread, filename, script, nil)
	assert.NoError(t, err)

	initFn, ok := globals["init"].(*starlark.Function)
	assert.True(t, ok, "init function not found")
	execFn, ok := globals["exec"].(*starlark.Function)
	assert.True(t, ok, "exec function not found")

	// Создаём состояние
	result, err := starlark.Call(thread, initFn, starlark.Tuple{}, nil)
	assert.NoError(t, err)
	stateDict, ok := result.(*starlark.Dict)
	assert.True(t, ok, "init should return dict")

	// Выполняем exec
	_, err = starlark.Call(thread, execFn, starlark.Tuple{stateDict}, nil)
	assert.NoError(t, err)

	// Проверяем результаты в state
	r, found, _ := stateDict.Get(starlark.String("r"))
	assert.True(t, found, "r not found in state")
	rInt, ok := r.(starlark.Int)
	assert.True(t, ok, "r should be int")
	rInt64, _ := rInt.Int64()
	assert.Equal(t, int64(1), rInt64)

	_, err = starlark.Call(thread, execFn, starlark.Tuple{stateDict}, nil)
	assert.NoError(t, err)
	r, found, _ = stateDict.Get(starlark.String("r"))
	assert.True(t, found, "r not found in state")
	rInt, ok = r.(starlark.Int)
	assert.True(t, ok, "r should be int")
	rInt64, _ = rInt.Int64()
	assert.Equal(t, int64(1), rInt64)

	_, err = starlark.Call(thread, execFn, starlark.Tuple{stateDict}, nil)
	assert.NoError(t, err)
	r, found, _ = stateDict.Get(starlark.String("r"))
	assert.True(t, found, "r not found in state")
	rInt, ok = r.(starlark.Int)
	assert.True(t, ok, "r should be int")
	rInt64, _ = rInt.Int64()
	assert.Equal(t, int64(2), rInt64)

}
func getOptions(src string) *syntax.FileOptions {
	return &syntax.FileOptions{
		Set:               option(src, "set"),
		While:             option(src, "while"),
		TopLevelControl:   option(src, "toplevelcontrol"),
		GlobalReassign:    option(src, "globalreassign"),
		LoadBindsGlobally: option(src, "loadbindsglobally"),
		Recursion:         option(src, "recursion"),
	}
}

func option(chunk, name string) bool {
	return strings.Contains(chunk, "option:"+name)
}
