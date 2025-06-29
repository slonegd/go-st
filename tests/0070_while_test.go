package main

import (
	_ "embed"
	"testing"

	"github.com/slonegd/go-st/ast"
	"github.com/stretchr/testify/require"
)

//go:embed 002_arithmetic_if_while.st
var arithmetic_if_while string

func Test_Execute_0070(t *testing.T) {
	require := require.New(t)

	p, err := ast.Parse(arithmetic_if_while)
	require.NoError(err)

	p.ExecuteDebug()
	require.Equal(int64(10), p.GetVar("r").Int())
	require.Equal(int64(2), p.GetVar("i").Int())
	require.Equal(int64(9), p.GetVar("r2").Int())
	require.Equal(int64(7), p.GetVar("counter").Int())

	p.Execute()
	require.Equal(int64(19), p.GetVar("r").Int())
	require.Equal(int64(2), p.GetVar("i").Int())
}
