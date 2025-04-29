package main

import (
	"testing"

	"github.com/slonegd/go-st"
	"github.com/stretchr/testify/require"
)

func TestProgram_Execute_003(t *testing.T) {
	require := require.New(t)

	p := st.NewProgram(iftest)
	require.Equal(0, p.Variables["i"])
	require.Equal(0, p.Variables["divisor"])

	p.Execute()
	require.Equal(1, p.Variables["i"])
	require.Equal(1, p.Variables["divisor"])

	p.Execute()
	require.Equal(2, p.Variables["i"])
	require.Equal(2, p.Variables["divisor"])

	p.Execute()
	require.Equal(3, p.Variables["i"])
	require.Equal(3, p.Variables["divisor"])

	p.Execute()
	require.Equal(4, p.Variables["i"])
	require.Equal(4, p.Variables["divisor"])

	p.Execute()
	require.Equal(5, p.Variables["i"])
	require.Equal(5, p.Variables["divisor"])

	p.Execute()
	require.Equal(6, p.Variables["i"])
	require.Equal(3, p.Variables["divisor"])

	p.Execute()
	require.Equal(7, p.Variables["i"])
	require.Equal(1, p.Variables["divisor"])

	p.Execute()
	require.Equal(8, p.Variables["i"])
	require.Equal(4, p.Variables["divisor"])
}
