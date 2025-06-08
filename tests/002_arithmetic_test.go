package main

import (
	"testing"

	"github.com/slonegd/go-st"
	"github.com/slonegd/go-st/fvm"
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

func TestFVM_Execute_002(t *testing.T) {
	require := require.New(t)

	p, err := fvm.NewProgram(arithmetic)
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

// на примере арифметики с int64 бенчмарк показал, что последовательный вызов функций быстрее байткода
// cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz

// с прошлых экспериментов
// BenchmarkProgram_Execute_002-4   1000000000    0.006594 ns/op // 1 вариант на колбеках и вариантом
// BenchmarkCompiler_Execute_002-4  1000000000    0.004587 ns/op // байткод c variant стеком

// эксперимент с функциональным рантаймом с int64 арифметикой
// BenchmarkProgram_Execute_002-4   1000000000    0.004498 ns/op // байткод c any стеком
// BenchmarkFVM_Execute_002-4       1000000000    0.002985 ns/op // функции с any (пробовал uintptr - медленнее)
// BenchmarkFVM_Execute_002-4       1000000000    0.0005666 ns/op // дженерик функции
// BenchmarkGo_Execute_002-4        1000000000    0.0000611 ns/op // чистый го
func BenchmarkProgram_Execute_002(b *testing.B) {
	p, _ := st.NewProgram(arithmetic)

	b.ResetTimer()
	for range 10000 {
		p.Execute()
	}
}

func BenchmarkFVM_Execute_002(b *testing.B) {
	p, _ := fvm.NewProgram(arithmetic)

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
