package main

import (
	"testing"

	"github.com/slonegd/go-st/ast"
	"github.com/stretchr/testify/require"
)

func TestFVM_Execute_002(t *testing.T) {
	require := require.New(t)

	p, err := ast.Parse(arithmetic)
	require.NoError(err)

	require.Equal(int64(0), p.GetVar("i").Int())
	require.Equal(int64(0), p.GetVar("j").Int())
	require.Equal(int64(42), p.GetVar("k").Int())

	p.ExecuteDebug()
	require.Equal(int64(15), p.GetVar("i").Int())
	require.Equal(int64(95), p.GetVar("j").Int())
	require.Equal(int64(40), p.GetVar("k").Int())

	p.Execute()
	require.Equal(int64(110), p.GetVar("i").Int())
	require.Equal(int64(-3515), p.GetVar("j").Int())
	require.Equal(int64(38), p.GetVar("k").Int())
}
