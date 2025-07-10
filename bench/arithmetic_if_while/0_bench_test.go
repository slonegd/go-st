package bench

import (
	_ "embed"
	"testing"

	"github.com/slonegd/go-st"
)

func Benchmark_go_001(b *testing.B) {
	var i, j int
	k = 42
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

		for counter1 < 10 {
			r += counter1
			counter1++
		}
		counter1 -= 7

		for counter2 < 10 {
			counter2++
			if counter2%2 == 0 {
				continue
			}
			r2 += counter2
		}
		counter2 = 0

		for counter3 < 10 {
			counter3++
			if counter3 > 7 {
				break
			}
			r3 += counter3
		}
		counter3 = 0
	}

	for b.Loop() {
		f()
	}
}

//go:embed 002_arithmetic_if_while.st
var arithmetic_st_while string

// cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
// Benchmark_go_001-4              24658555                48 ns/op (x0.04) // pure go
// Benchmark_st_001-4               1000000              1083 ns/op (x1)    // this
// Benchmark_yaegi_001-4             300710              3832 ns/op (x3)     // github.com/traefik/yaegi v0.16.1
// Benchmark_lua_001-4                61252             19402 ns/op (x17)    // github.com/Shopify/go-lua

// cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
// Benchmark_go_001-4          1000000000    0.0000487 ns/op (x0.04)  // pure go
// Benchmark_st_001-4          1000000000    0.0011470 ns/op (x1)     // this
// Benchmark_yaegi_001-4       1000000000    0.0036320 ns/op (x3)     // github.com/traefik/yaegi v0.16.1
// Benchmark_lua_001-4         1000000000    0.0207000 ns/op (x18)    // github.com/Shopify/go-lua
func Benchmark_st_001(b *testing.B) {
	p, err := st.NewProgram(arithmetic_st_while)
	if err != nil {
		b.Log(err)
		b.Fail()
	}

	for b.Loop() {
		p.Execute()
	}
}
