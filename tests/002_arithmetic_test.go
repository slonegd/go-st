package main

import (
	"testing"

	"github.com/slonegd/go-st"
	"github.com/stretchr/testify/require"
)

func TestCompiler_Execute_002(t *testing.T) {
	require := require.New(t)

	p, err := st.NewProgram(arithmetic)
	require.NoError(err)

	require.Equal(int64(0), p.GetVar("i").Int())
	require.Equal(int64(0), p.GetVar("j").Int())
	require.Equal(int64(42), p.GetVar("k").Int())

	p.Execute()
	require.Equal(int64(15), p.GetVar("i").Int())
	require.Equal(int64(95), p.GetVar("j").Int())
	require.Equal(int64(40), p.GetVar("k").Int())

	p.Execute()
	require.Equal(int64(110), p.GetVar("i").Int())
	require.Equal(int64(-3515), p.GetVar("j").Int())
	require.Equal(int64(38), p.GetVar("k").Int())
}

// пока байткод оказался лучше колбэков
// UPDATE использование фиксированного размера оператора с аргументами дало +20% прирост скорости
// UPDATE использование типизированных операторов (пока только INT) дало +15% прирост скорости, пока убрал, будет в оптимизаторе
// TODO можно на стек класть не варианты, а значения, типизированные операторы знают к чему привести
// TODO для переменных на стеке можно использовать указатели на значения, но тогда надо больше операторов
// чтоб отличить указатель от значения.
// cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
// BenchmarkProgram_Execute_002-4          1000000000               0.006594 ns/op // 1 вариант на колбеках
// BenchmarkCompiler_Execute_002-4         1000000000               0.004587 ns/op // текущий вариант
// BenchmarkGo_Execute_002-4               1000000000               0.0000611 ns/op
func BenchmarkProgram_Execute_002(b *testing.B) {
	p, _ := st.NewProgram(arithmetic)

	b.ResetTimer()
	for range 10000 {
		p.Execute()
	}
}

func BenchmarkGo_Execute_002(b *testing.B) {
	var i, j int
	k := 42
	f := func() {
		i = 15 + j
		j = 5 + 2*i*(30-i)/5
		k = k + -2
	}

	b.ResetTimer()
	for range 10000 {
		f()
	}
	b.StopTimer()
	b.Logf("%v %v %v", i, j, k)
}
