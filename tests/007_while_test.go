package main

import (
	"testing"

	"github.com/slonegd/go-st/ast"
	"github.com/stretchr/testify/require"
)

func Test_Execute_007(t *testing.T) {
	require := require.New(t)

	p, err := ast.Parse(while)
	require.NoError(err)

	p.Execute()
	require.Equal(int64(10), p.GetVar("r").Int())
	require.Equal(int64(2), p.GetVar("i").Int())

	p.Execute()
	require.Equal(int64(19), p.GetVar("r").Int())
	require.Equal(int64(2), p.GetVar("i").Int())
}
