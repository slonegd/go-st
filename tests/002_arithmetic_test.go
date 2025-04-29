package main

import (
	"testing"

	"github.com/slonegd/go-st"
	"github.com/stretchr/testify/require"
)

func TestProgram_Execute_002(t *testing.T) {
	require := require.New(t)

	p := st.NewProgram(arithmetic)
	require.Equal(0, p.Variables["i"])
	require.Equal(0, p.Variables["j"])

	p.Execute()
	require.Equal(15, p.Variables["i"])
	require.Equal(95, p.Variables["j"])

	p.Execute()
	require.Equal(110, p.Variables["i"])
	require.Equal(-3515, p.Variables["j"])
}
