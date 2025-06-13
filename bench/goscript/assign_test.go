package goScript

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAssign(t *testing.T) {
	require := require.New(t)
	expr := `
x1 := ti1001
x2 = x1 + 2.25
x1 + x2 + 5`
	ctx := map[string]any{"ti1001": 1.15}
	val, err := Eval(expr, ctx)
	require.NoError(err)
	require.Equal(9.55, val)

	prepared := &Expr{}
	err = prepared.Prepare(expr)
	require.NoError(err)

	ctx = map[string]any{"ti1001": 1.15}
	val, err = prepared.Eval(ctx)
	require.NoError(err)
	require.Equal(9.55, val)

	require.Equal(float64(3.4), ctx["x2"])
}

func TestIf(t *testing.T) {
	require := require.New(t)
	expr := `
if x1 > 1 {
	return 1.
}
x1-1
`
	ctx := func(v float64) map[string]any {
		return map[string]any{"x1": v}
	}

	val, err := Eval(expr, ctx(5))
	require.NoError(err)
	require.Equal(1., val)

	prepared := &Expr{}
	err = prepared.Prepare(expr)
	require.NoError(err)

	val, err = prepared.Eval(ctx(5))
	require.NoError(err)
	require.Equal(1., val)

	val, err = prepared.Eval(ctx(0.5))
	require.NoError(err)
	require.Equal(-0.5, val)

	expr = `
if x1 > 1 {
	return 1.
} else {
    x1 = x1-1
}
return x1
`

	prepared = &Expr{}
	err = prepared.Prepare(expr)
	require.NoError(err)

	val, err = prepared.Eval(ctx(5))
	require.NoError(err)
	require.Equal(1., val)

	val, err = prepared.Eval(ctx(0.5))
	require.NoError(err)
	require.Equal(-0.5, val)
}
