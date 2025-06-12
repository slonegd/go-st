package bench

import (
	_ "embed"
	"testing"

	"github.com/slonegd/go-st"
)

const iterations = 20 // при большем бенчмарк gpython зависает

//go:embed 001_arithmetic_if.st
var arithmetic_st string

// cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
// Benchmark_go_001-4       1000000000    0.0000007 ns/op  // pure go
// Benchmark_st_001-4       1000000000    0.0000059 ns/op  // this
// Benchmark_yaegi_001-4    1000000000    0.0000347 ns/op  // github.com/traefik/yaegi v0.16.1
// Benchmark_lua_001-4      1000000000    0.0000566 ns/op  // github.com/Shopify/go-lua
// Benchmark_gpython_001-4  1000000000    0.07737 ns/op    // github.com/go-python/gpython v0.2.0
func Benchmark_st_001(b *testing.B) {
	p, _ := st.NewProgram(arithmetic_st)

	b.ResetTimer()
	for range iterations {
		p.Execute()
	}
}

func Benchmark_go_001(b *testing.B) {
	var i, j int
	k := 42
	f := func() {
		i = 15 + j
		j = 5 + 2*i*(30-i)/5
		k = k + -2

		i++
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
	}

	b.ResetTimer()
	for range iterations {
		f()
	}
}
