package main

import (
	"testing"

	"github.com/slonegd/go-st/ast"
	"github.com/stretchr/testify/require"
)

func Test_Execute_008(t *testing.T) {
	require := require.New(t)

	_, err := ast.Parse(function)
	require.Error(err)

	// p.ExecuteDebug()
	// require.Equal(int64(10), p.GetVar("r").Int())
}
