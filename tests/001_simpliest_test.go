package main

import (
	"testing"

	"github.com/slonegd/go-st"
	"github.com/slonegd/go-st/compiler"
	"github.com/stretchr/testify/require"
)

func TestProgram_Execute_001(t *testing.T) {
	require := require.New(t)

	p, _ := st.NewProgram(simpliest)
	require.Equal(int64(5), p.Variables["i"].Int())
	require.Equal(int64(0), p.Variables["j"].Int())
	require.Equal(int64(-42), p.Variables["k"].Int())

	p.Execute()
	require.Equal(int64(15), p.Variables["i"].Int())
	require.Equal(int64(0), p.Variables["j"].Int())
	require.Equal(int64(-42), p.Variables["k"].Int())
}

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
