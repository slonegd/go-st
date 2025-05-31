package main

import (
	"testing"

	"github.com/slonegd/go-st"
	"github.com/slonegd/go-st/compiler"
	"github.com/stretchr/testify/require"
)

func TestProgram_Execute_002(t *testing.T) {
	require := require.New(t)

	p, _ := st.NewProgram(arithmetic)
	require.Equal(int64(0), p.Variables["i"].Int())
	require.Equal(int64(0), p.Variables["j"].Int())
	require.Equal(int64(42), p.Variables["k"].Int())

	p.Execute()
	require.Equal(int64(15), p.Variables["i"].Int())
	require.Equal(int64(95), p.Variables["j"].Int())
	require.Equal(int64(40), p.Variables["k"].Int())

	p.Execute()
	require.Equal(int64(110), p.Variables["i"].Int())
	require.Equal(int64(-3515), p.Variables["j"].Int())
	require.Equal(int64(38), p.Variables["k"].Int())
}

func TestCompiler_Execute_002(t *testing.T) {
	require := require.New(t)

	c := compiler.New()
	err := c.Compile(arithmetic)
	require.NoError(err)

	require.Equal(int64(0), c.GetVar("i").Int())
	require.Equal(int64(0), c.GetVar("j").Int())
	require.Equal(int64(42), c.GetVar("k").Int())

	c.Execute()
	require.Equal(int64(15), c.GetVar("i").Int())
	require.Equal(int64(95), c.GetVar("j").Int())
	require.Equal(int64(40), c.GetVar("k").Int())

	c.Execute()
	require.Equal(int64(110), c.GetVar("i").Int())
	require.Equal(int64(-3515), c.GetVar("j").Int())
	require.Equal(int64(38), c.GetVar("k").Int())
}

// пока байткод оказался лучше колбэков
// UPDATE использование фиксированного размера оператора с аргументами дало +20% прирост скорости
// UPDATE использование типизированных операторов (пока только INT) дало +15% прирост скорости
// TODO можно на стек класть не варианты, а значения, типизированные операторы знают к чему привести
// TODO для переменных на стеке можно использовать указатели на значения, но тогда надо больше операторов
// чтоб отличить указатель от значения.
// cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
// BenchmarkProgram_Execute_002-4          1000000000               0.006594 ns/op
// BenchmarkCompiler_Execute_002-4         1000000000               0.004587 ns/op
// BenchmarkGo_Execute_002-4               1000000000               0.0000611 ns/op
func BenchmarkProgram_Execute_002(b *testing.B) {
	p, _ := st.NewProgram(arithmetic)

	b.ResetTimer()
	for range 10000 {
		p.Execute()
	}
}

func BenchmarkCompiler_Execute_002(b *testing.B) {
	c := compiler.New()
	c.Compile(arithmetic)

	b.ResetTimer()
	for range 10000 {
		c.Execute()
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
