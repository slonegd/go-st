package main

import (
	"testing"

	"github.com/slonegd/go-st/ast"
	"github.com/stretchr/testify/require"
)

func Test_Execute_008(t *testing.T) {
	require := require.New(t)

	p, err := ast.Parse(function)
	require.NoError(err)

	p.ExecuteDebug()
	require.Equal(int64(25), p.GetVar("square").Int())
	require.Equal(int64(120), p.GetVar("factorial").Int())
	require.Equal(int64(720), p.GetVar("factorial6").Int())

	p.Execute() // проверяем, что в функциях временные переменные сбрасываются
	require.Equal(int64(25), p.GetVar("square").Int())
	require.Equal(int64(120), p.GetVar("factorial").Int())
	require.Equal(int64(720), p.GetVar("factorial6").Int())
}
