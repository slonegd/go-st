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
// Benchmark_go_fibo-4             16328564                69.58 ns/op (x0.02) // pure go
// Benchmark_st_fibo-4               419558              2736 ns/op    (x1)    // this
// Benchmark_yaegi_001-4              47620             24804 ns/op    (x9)    // github.com/traefik/yaegi v0.16.1
// Benchmark_starlark_fibo-4          30840             38709 ns/op    (x14)   // go.starlark.net v0.0.0-20250804182900-3c9dc17c5f2e
func Benchmark_st_fibo(b *testing.B) {
	p, _ := st.NewProgram(fibo)

	for b.Loop() {
		p.Execute()
	}
}

func Test_go_fibo(t *testing.T) {
	// [0 1 2 3 4 5 6 7  8  9  10]
	// [0 1 1 2 3 5 8 13 21 34 55]
	exec()
	assert.Equal(t, 1, r)
	exec()
	assert.Equal(t, 1, r)
	exec()
	assert.Equal(t, 2, r)
	exec()
	assert.Equal(t, 3, r)
	exec()
	assert.Equal(t, 5, r)

}

func Test_st_fibo(t *testing.T) {
	p, _ := st.NewProgram(fibo)

	p.Execute()
	assert.Equal(t, int64(1), p.GetVar("r").Int())
	p.Execute()
	assert.Equal(t, int64(1), p.GetVar("r").Int())
	p.Execute()
	assert.Equal(t, int64(2), p.GetVar("r").Int())
	p.Execute()
	assert.Equal(t, int64(3), p.GetVar("r").Int())
	p.Execute()
	assert.Equal(t, int64(5), p.GetVar("r").Int())
}
