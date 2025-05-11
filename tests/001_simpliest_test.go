package main

import (
	"testing"

	"github.com/slonegd/go-st"
	"github.com/stretchr/testify/require"
)

func TestProgram_Execute_001(t *testing.T) {
	require := require.New(t)

	p := st.NewProgram(simpliest)
	require.Equal(int64(5), p.Variables["i"].Int())
	require.Equal(int64(0), p.Variables["j"].Int())
	require.Equal(int64(42), p.Variables["k"].Int())

	p.Execute()
	require.Equal(int64(15), p.Variables["i"].Int())
	require.Equal(int64(0), p.Variables["j"].Int())
	require.Equal(int64(42), p.Variables["k"].Int())
}
