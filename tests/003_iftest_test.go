package main

import (
	"testing"

	"github.com/slonegd/go-st/ast"
	"github.com/stretchr/testify/require"
)

func TestCompiler_Execute_003(t *testing.T) {
	require := require.New(t)

	p, err := ast.Parse(iftest)
	require.NoError(err)
	require.Equal(int64(0), p.GetVar("i").Int())
	require.Equal(int64(0), p.GetVar("divisor").Int())

	p.Execute()
	require.Equal(int64(1), p.GetVar("i").Int())
	require.Equal(int64(1), p.GetVar("divisor").Int(), "проверка последнего elsif без else в конце")

	p.Execute()
	require.Equal(int64(2), p.GetVar("i").Int())
	require.Equal(int64(2), p.GetVar("divisor").Int(), "проверка внутреннего elsif")

	p.Execute()
	require.Equal(int64(3), p.GetVar("i").Int())
	require.Equal(int64(3), p.GetVar("divisor").Int())

	p.Execute()
	require.Equal(int64(4), p.GetVar("i").Int())
	require.Equal(int64(4), p.GetVar("divisor").Int())

	p.Execute()
	require.Equal(int64(5), p.GetVar("i").Int())
	require.Equal(int64(5), p.GetVar("divisor").Int())
	require.Equal(int64(1), p.GetVar("mult").Int())
	require.Equal(int64(0), p.GetVar("test").Int(), "проверка простого if в конце (ни разу не сработал)")

	p.Execute()
	require.Equal(int64(6), p.GetVar("i").Int())
	require.Equal(int64(3), p.GetVar("divisor").Int())
	require.Equal(int64(2), p.GetVar("mult").Int())
	require.Equal(int64(6), p.GetVar("test").Int(), "проверка простого if в конце (сработал первый раз)")

	p.Execute()
	require.Equal(int64(7), p.GetVar("i").Int())
	require.Equal(int64(1), p.GetVar("divisor").Int())
	require.Equal(int64(7), p.GetVar("mult").Int())
	require.Equal(int64(7), p.GetVar("test").Int())

	p.Execute()
	require.Equal(int64(8), p.GetVar("i").Int())
	require.Equal(int64(4), p.GetVar("divisor").Int())
	require.Equal(int64(2), p.GetVar("mult").Int())
	require.Equal(int64(8), p.GetVar("test").Int())

	p.Execute()
	p.Execute()
	require.Equal(int64(10), p.GetVar("i").Int())
	require.Equal(int64(5), p.GetVar("divisor").Int())
	require.Equal(int64(2), p.GetVar("mult").Int(), "проверка работы условия внутри другого условия, ветка elif")
	require.Equal(int64(10), p.GetVar("test").Int())

	p.Execute()
	p.Execute()
	p.Execute()
	p.Execute()
	p.Execute()
	require.Equal(int64(15), p.GetVar("i").Int())
	require.Equal(int64(5), p.GetVar("divisor").Int())
	require.Equal(int64(3), p.GetVar("mult").Int(), "проверка работы условия внутри другого условия, ветка else")
	require.Equal(int64(15), p.GetVar("test").Int())
}
