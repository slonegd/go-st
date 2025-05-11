package main

import (
	"testing"

	"github.com/slonegd/go-st"
	"github.com/stretchr/testify/require"
)

func TestProgram_Execute_002(t *testing.T) {
	require := require.New(t)

	p := st.NewProgram(arithmetic)
	require.Equal(int64(0), p.Variables["i"].Int())
	require.Equal(int64(0), p.Variables["j"].Int())
	require.Equal(int64(42), p.Variables["k"].Int())

	p.Execute()
	require.Equal(int64(15), p.Variables["i"].Int())
	require.Equal(int64(95), p.Variables["j"].Int())
	require.Equal(int64(40), p.Variables["k"].Int())

	p.Execute()
	require.Equal(int64(110), p.Variables["i"].Int())
	require.Equal(int64(-3515), p.Variables["j"].Int())
	require.Equal(int64(38), p.Variables["k"].Int())
}
