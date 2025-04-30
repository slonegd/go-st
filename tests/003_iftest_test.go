package main

import (
	"log"
	"testing"

	"github.com/slonegd/go-st"
	"github.com/stretchr/testify/require"
)

func TestProgram_Execute_003(t *testing.T) {
	require := require.New(t)

	p := st.NewProgram(iftest)
	require.Equal(int64(0), p.Variables["i"].Int())
	require.Equal(int64(0), p.Variables["divisor"].Int())

	p.Execute()
	log.Printf("results: %v", p.Variables)
	require.Equal(int64(1), p.Variables["i"].Int())
	require.Equal(int64(1), p.Variables["divisor"].Int(), "проверка последнего elsif без else в конце")

	p.Execute()
	log.Printf("results: %v", p.Variables)
	require.Equal(int64(2), p.Variables["i"].Int())
	require.Equal(int64(2), p.Variables["divisor"].Int(), "проверка внутреннего elsif")

	p.Execute()
	log.Printf("results: %v", p.Variables)
	require.Equal(int64(3), p.Variables["i"].Int())
	require.Equal(int64(3), p.Variables["divisor"].Int())

	p.Execute()
	log.Printf("results: %v", p.Variables)
	require.Equal(int64(4), p.Variables["i"].Int())
	require.Equal(int64(4), p.Variables["divisor"].Int())

	p.Execute()
	log.Printf("results: %v", p.Variables)
	require.Equal(int64(5), p.Variables["i"].Int())
	require.Equal(int64(5), p.Variables["divisor"].Int())
	require.Equal(int64(1), p.Variables["mult"].Int())
	require.Equal(int64(0), p.Variables["test"].Int(), "проверка простого if в конце (ни разу не сработал)")

	p.Execute()
	log.Printf("results: %v", p.Variables)
	require.Equal(int64(6), p.Variables["i"].Int())
	require.Equal(int64(3), p.Variables["divisor"].Int())
	require.Equal(int64(2), p.Variables["mult"].Int())
	require.Equal(int64(6), p.Variables["test"].Int(), "проверка простого if в конце (сработал первый раз)")

	p.Execute()
	log.Printf("results: %v", p.Variables)
	require.Equal(int64(7), p.Variables["i"].Int())
	require.Equal(int64(1), p.Variables["divisor"].Int())
	require.Equal(int64(7), p.Variables["mult"].Int())
	require.Equal(int64(7), p.Variables["test"].Int())

	p.Execute()
	log.Printf("results: %v", p.Variables)
	require.Equal(int64(8), p.Variables["i"].Int())
	require.Equal(int64(4), p.Variables["divisor"].Int())
	require.Equal(int64(2), p.Variables["mult"].Int())
	require.Equal(int64(8), p.Variables["test"].Int())

	p.Execute()
	p.Execute()
	log.Printf("results: %v", p.Variables)
	require.Equal(int64(10), p.Variables["i"].Int())
	require.Equal(int64(5), p.Variables["divisor"].Int())
	require.Equal(int64(2), p.Variables["mult"].Int(), "проверка работы условия внутри другого условия, ветка elif")
	require.Equal(int64(10), p.Variables["test"].Int())

	p.Execute()
	p.Execute()
	p.Execute()
	p.Execute()
	p.Execute()
	log.Printf("results: %v", p.Variables)
	require.Equal(int64(15), p.Variables["i"].Int())
	require.Equal(int64(5), p.Variables["divisor"].Int())
	require.Equal(int64(3), p.Variables["mult"].Int(), "проверка работы условия внутри другого условия, ветка else")
	require.Equal(int64(15), p.Variables["test"].Int())
}
