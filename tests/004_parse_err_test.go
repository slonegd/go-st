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
	require.Equal(`compile: 
  3|    i : INT := 5;
  4|  END_VAR
  5|  i := 1x5;
   |........^    mismatched input 'x5' expecting 'END_PROGRAM'
  6|END_PROGRAM
`, err.Error())
}
