package main

import (
	"testing"

	"github.com/slonegd/go-st/ast"
	"github.com/stretchr/testify/require"
)

func Test_Execute_005(t *testing.T) {
	require := require.New(t)

	p, err := ast.Parse(implicit_int_cast)
	require.NoError(err)
	p.Execute()
	require.Equal(int64(-1), p.GetVar("i8").Int())
	require.Equal(int64(-1), p.GetVar("i16").Int())
	require.Equal(int64(-1), p.GetVar("i32").Int())
	require.Equal(int64(-1), p.GetVar("i64").Int())

	require.Equal(int64(255), p.GetVar("ui8").Int())
	require.Equal(int64(65535), p.GetVar("ui16").Int())
	require.Equal(int64(4294967295), p.GetVar("ui32").Int())
	require.Equal(uint64(18_446_744_073_709_551_615), p.GetVar("ui64").Uint())

	p.Execute()
	require.Equal(int64(127), p.GetVar("i8").Int())
	require.Equal(int64(32767), p.GetVar("i16").Int())
	require.Equal(int64(2147483647), p.GetVar("i32").Int())
	require.Equal(int64(9223372036854775807), p.GetVar("i64").Int())

	p.Execute()
	require.Equal(int64(-128), p.GetVar("i8").Int())
	require.Equal(int64(-32768), p.GetVar("i16").Int())
	require.Equal(int64(-2147483648), p.GetVar("i32").Int())
	require.Equal(int64(-9223372036854775808), p.GetVar("i64").Int())
}
