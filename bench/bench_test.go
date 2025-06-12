package bench

import (
	_ "embed"
	"testing"

	"github.com/slonegd/go-st"
)

//go:embed 001_arithmetic_if.st
var arithmetic_st string

// cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
// Benchmark_st_001-4              1000000000               0.002118 ns/op
// Benchmark_yaegi_001-4           1000000000               0.01586 ns/op
func Benchmark_st_001(b *testing.B) {
	p, _ := st.NewProgram(arithmetic_st)

	b.ResetTimer()
	for range 10000 {
		p.Execute()
	}
}
