package bench

import (
	_ "embed"
	"testing"

	"github.com/slonegd/go-st"
	"github.com/stretchr/testify/assert"
)

func Benchmark_go_fibo(b *testing.B) {
	for b.Loop() {
		exec()
	}
}

//go:embed fibo.st
var fibo string

// cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
// Benchmark_go_fibo-4             99592340                12.06 ns/op (x0.02) // pure go
// Benchmark_st_fibo-4              2409584               491.5 ns/op  (x1)    // this
// Benchmark_yaegi_001-4             219374              5023 ns/op    (x10)   // github.com/traefik/yaegi v0.16.1
func Benchmark_st_fibo(b *testing.B) {
	p, _ := st.NewProgram(fibo)

	for b.Loop() {
		p.Execute()
	}
}

func Test_go_fibo(t *testing.T) {
	exec()
	// [0 1 2 3 4 5 6 7  8  9  10]
	// [0 1 1 2 3 5 8 13 21 34 55]
	assert.Equal(t, 5, fibo5)
	assert.Equal(t, 55, fibo10)
}

func Test_st_fibo(t *testing.T) {
	p, _ := st.NewProgram(fibo)
	p.Execute()

	assert.Equal(t, int64(5), p.GetVar("fibo5").Int())
	assert.Equal(t, int64(55), p.GetVar("fibo10").Int())
}
