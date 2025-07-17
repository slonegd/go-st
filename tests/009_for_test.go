package main

import (
	"testing"

	"github.com/slonegd/go-st/ast"
	"github.com/stretchr/testify/require"
)

func Test_Execute_009(t *testing.T) {
	require := require.New(t)

	p, err := ast.Parse(forTest)
	require.NoError(err)

	p.ExecuteDebug()
	require.Equal(float64(45), p.GetVar("v").Float64())
}
