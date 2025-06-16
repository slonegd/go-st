package main

import (
	"testing"

	"github.com/slonegd/go-st/ast"
	"github.com/stretchr/testify/require"
)

func Test_Execute_006(t *testing.T) {
	require := require.New(t)

	p, err := ast.Parse(floats)
	require.NoError(err)
	require.Equal(1., p.GetVar("f1").Float64())
	require.Equal(-0.5, p.GetVar("f2").Float64())
	require.Equal(50., p.GetVar("f3").Float64())

	p.Execute()
	require.Equal(-0.5, p.GetVar("f1").Float64())
	require.Equal(50., p.GetVar("f2").Float64())
}
