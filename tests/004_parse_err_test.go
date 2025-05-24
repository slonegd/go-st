package main

import (
	"testing"

	"github.com/slonegd/go-st"
	"github.com/stretchr/testify/require"
)

func TestProgram_Execute_004(t *testing.T) {
	require := require.New(t)

	_, err := st.NewProgram(parseErr)
	require.Error(err)
	require.Equal(`
03|    i : INT := 5;
04|  END_VAR
05|  i := 1x5;
00|........^    extraneous input 'x5' expecting ';'
06|END_PROGRAM
`, err.Error())
}
