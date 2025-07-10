package bench

import (
	goScript "bench/goscript"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

const arithmetic_goscript = `
	a = 15 + j
	j = 5 + 2*a*(30-a)/5
	k = k + -2

	i = i + 1
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
`

func Benchmark_goscript_001(b *testing.B) {
	expr := &goScript.Expr{}
	expr.Prepare(arithmetic_goscript)
	ctx := map[string]any{
		"a": 0,
		"j": 0,
		"k": 42,

		"i":       0,
		"divisor": 0,
		"mult":    0,
		"test":    0,
	}

	for b.Loop() {
		expr.Eval(ctx)
	}
}

func Test_goscript_001(t *testing.T) {
	expr := &goScript.Expr{}
	err := expr.Prepare(arithmetic_goscript)
	assert.NoError(t, err)
	ctx := map[string]any{
		"a": 0,
		"j": 0,
		"k": 42,

		"i":       0,
		"divisor": 0,
		"mult":    0,
		"test":    0,
	}

	_, err = expr.Eval(ctx)
	assert.NoError(t, err)

	assert.Equal(t, int64(1), ctx["i"])
	assert.Equal(t, int64(1), ctx["divisor"])
	assert.Equal(t, int64(1), ctx["mult"])
	assert.Equal(t, 0, ctx["test"])

	_, err = expr.Eval(ctx)
	assert.Equal(t, int64(2), ctx["i"])
	assert.Equal(t, int64(2), ctx["divisor"])
	assert.Equal(t, int64(1), ctx["mult"])
	assert.Equal(t, 0, ctx["test"])

	_, err = expr.Eval(ctx)
	_, err = expr.Eval(ctx)
	_, err = expr.Eval(ctx)
	_, err = expr.Eval(ctx)
	_, err = expr.Eval(ctx)
	assert.Equal(t, int64(7), ctx["i"])
	assert.Equal(t, int64(1), ctx["divisor"])
	assert.Equal(t, int64(1), ctx["mult"])
	assert.Equal(t, int64(6), ctx["test"])
}
