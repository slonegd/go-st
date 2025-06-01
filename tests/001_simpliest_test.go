package main

import (
	"testing"

	"github.com/slonegd/go-st/compiler"
	"github.com/stretchr/testify/require"
)

func TestCompiler_Execute_001(t *testing.T) {
	require := require.New(t)

	c := compiler.New()
	err := c.Compile(simpliest)
	require.NoError(err)

	require.Equal(int64(5), c.GetVar("i").Int())
	require.Equal(int64(0), c.GetVar("j").Int())
	require.Equal(int64(-42), c.GetVar("k").Int())

	c.Execute() // константы из инициализации остаются на стеке
	require.Equal(int64(15), c.GetVar("i").Int())
	require.Equal(int64(0), c.GetVar("j").Int())
	require.Equal(int64(-42), c.GetVar("k").Int())
}
