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

// нет разницы работать через колбеки или через байткод
// байткод чуть более понятен + возможно взятие переменных и констант не по индексу, а по указателю
// TODO сначала проверю бэнчмарк с условиями, потом бенчмарк с указателями
// TODO ну и операторы можно сделать типозависимыми, тогда меньше преобразований
// TODO так же через операторы возможны оптимизации, например, брать аргумент прямо из констант, без стека
// cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
// BenchmarkProgram_Execute_002-4          1000000000               0.006458 ns/op
// BenchmarkCompiler_Execute_002-4         1000000000               0.005709 ns/op

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
