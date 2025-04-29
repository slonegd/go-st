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
	require.Equal(0, p.Variables["i"])
	require.Equal(0, p.Variables["divisor"])

	p.Execute()
	log.Printf("results: %v", p.Variables)
	require.Equal(1, p.Variables["i"])
	require.Equal(1, p.Variables["divisor"], "проверка последнего elsif без else в конце")

	p.Execute()
	log.Printf("results: %v", p.Variables)
	require.Equal(2, p.Variables["i"])
	require.Equal(2, p.Variables["divisor"], "проверка внутреннего elsif")

	p.Execute()
	log.Printf("results: %v", p.Variables)
	require.Equal(3, p.Variables["i"])
	require.Equal(3, p.Variables["divisor"])

	p.Execute()
	log.Printf("results: %v", p.Variables)
	require.Equal(4, p.Variables["i"])
	require.Equal(4, p.Variables["divisor"])

	p.Execute()
	log.Printf("results: %v", p.Variables)
	require.Equal(5, p.Variables["i"])
	require.Equal(5, p.Variables["divisor"])
	require.Equal(1, p.Variables["mult"])
	require.Equal(0, p.Variables["test"], "проверка простого if в конце (ни разу не сработал)")

	p.Execute()
	log.Printf("results: %v", p.Variables)
	require.Equal(6, p.Variables["i"])
	require.Equal(3, p.Variables["divisor"])
	require.Equal(2, p.Variables["mult"])
	require.Equal(6, p.Variables["test"], "проверка простого if в конце (сработал первый раз)")

	p.Execute()
	log.Printf("results: %v", p.Variables)
	require.Equal(7, p.Variables["i"])
	require.Equal(1, p.Variables["divisor"])
	require.Equal(7, p.Variables["mult"])
	require.Equal(7, p.Variables["test"])

	p.Execute()
	log.Printf("results: %v", p.Variables)
	require.Equal(8, p.Variables["i"])
	require.Equal(4, p.Variables["divisor"])
	require.Equal(2, p.Variables["mult"])
	require.Equal(8, p.Variables["test"])

	p.Execute()
	p.Execute()
	log.Printf("results: %v", p.Variables)
	require.Equal(10, p.Variables["i"])
	require.Equal(5, p.Variables["divisor"])
	require.Equal(2, p.Variables["mult"], "проверка работы условия внутри другого условия, ветка elif")
	require.Equal(10, p.Variables["test"])

	p.Execute()
	p.Execute()
	p.Execute()
	p.Execute()
	p.Execute()
	log.Printf("results: %v", p.Variables)
	require.Equal(15, p.Variables["i"])
	require.Equal(5, p.Variables["divisor"])
	require.Equal(3, p.Variables["mult"], "проверка работы условия внутри другого условия, ветка else")
	require.Equal(15, p.Variables["test"])
}
