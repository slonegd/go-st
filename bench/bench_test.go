package bench

import (
	_ "embed"
	"testing"

	"github.com/slonegd/go-st"
)

const iterations = 1000 // при большем бенчмарк gpython зависает

//go:embed 001_arithmetic_if.st
var arithmetic_st string

// cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
// Benchmark_go_001-4          1000000000    0.0000165 ns/op (x0.08)  // pure go
// Benchmark_st_001-4          1000000000    0.0002076 ns/op (x1)     // this
// Benchmark_yaegi_001-4       1000000000    0.001653 ns/op  (x8)     // github.com/traefik/yaegi v0.16.1
// Benchmark_goscript_001-4    1000000000    0.002492 ns/op  (x12)    // github.com/japm/goScript
// Benchmark_lua_001-4         1000000000    0.002826 ns/op  (x13)    // github.com/Shopify/go-lua
// Benchmark_expr_001-4        1000000000    0.005552 ns/op  (x26)    // github.com/expr-lang/expr v1.17.5
// Benchmark_tengo_001-4       1000000000    0.01662 ns/op   (x80)    // github.com/d5/tengo/v2

// исключил gpython, он очень медленный и другие бенчмарки страдают,
// к тому же зависал, при высоком iterations, работало при iterations = 20 и был медленнее yaegi в 2222 раза
// Benchmark_gpython_001-4     1000000000    0.07707 ns/op (x17000)   // github.com/go-python/gpython v0.2.0
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
