package main

import (
	"log"
	"testing"

	"github.com/slonegd/go-st"
	"github.com/slonegd/go-st/compiler"
	"github.com/stretchr/testify/require"
)

func TestProgram_Execute_003(t *testing.T) {
	require := require.New(t)

	p, _ := st.NewProgram(iftest)
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

func TestCompiler_Execute_003(t *testing.T) {
	require := require.New(t)

	c := compiler.New()
	err := c.Compile(iftest)
	require.NoError(err)
	require.Equal(int64(0), c.GetVar("i").Int())
	require.Equal(int64(0), c.GetVar("divisor").Int())

	c.VM.Print()

	c.Execute()
	require.Equal(int64(1), c.GetVar("i").Int())
	require.Equal(int64(1), c.GetVar("divisor").Int(), "проверка последнего elsif без else в конце")

	c.Execute()
	require.Equal(int64(2), c.GetVar("i").Int())
	require.Equal(int64(2), c.GetVar("divisor").Int(), "проверка внутреннего elsif")

	c.Execute()
	require.Equal(int64(3), c.GetVar("i").Int())
	require.Equal(int64(3), c.GetVar("divisor").Int())

	c.Execute()
	require.Equal(int64(4), c.GetVar("i").Int())
	require.Equal(int64(4), c.GetVar("divisor").Int())

	c.Execute()
	require.Equal(int64(5), c.GetVar("i").Int())
	require.Equal(int64(5), c.GetVar("divisor").Int())
	require.Equal(int64(1), c.GetVar("mult").Int())
	require.Equal(int64(0), c.GetVar("test").Int(), "проверка простого if в конце (ни разу не сработал)")

	c.Execute()
	require.Equal(int64(6), c.GetVar("i").Int())
	require.Equal(int64(3), c.GetVar("divisor").Int())
	require.Equal(int64(2), c.GetVar("mult").Int())
	require.Equal(int64(6), c.GetVar("test").Int(), "проверка простого if в конце (сработал первый раз)")

	c.Execute()
	require.Equal(int64(7), c.GetVar("i").Int())
	require.Equal(int64(1), c.GetVar("divisor").Int())
	require.Equal(int64(7), c.GetVar("mult").Int())
	require.Equal(int64(7), c.GetVar("test").Int())

	c.Execute()
	require.Equal(int64(8), c.GetVar("i").Int())
	require.Equal(int64(4), c.GetVar("divisor").Int())
	require.Equal(int64(2), c.GetVar("mult").Int())
	require.Equal(int64(8), c.GetVar("test").Int())

	c.Execute()
	c.Execute()
	require.Equal(int64(10), c.GetVar("i").Int())
	require.Equal(int64(5), c.GetVar("divisor").Int())
	require.Equal(int64(2), c.GetVar("mult").Int(), "проверка работы условия внутри другого условия, ветка elif")
	require.Equal(int64(10), c.GetVar("test").Int())

	c.Execute()
	c.Execute()
	c.Execute()
	c.Execute()
	c.Execute()
	require.Equal(int64(15), c.GetVar("i").Int())
	require.Equal(int64(5), c.GetVar("divisor").Int())
	require.Equal(int64(3), c.GetVar("mult").Int(), "проверка работы условия внутри другого условия, ветка else")
	require.Equal(int64(15), c.GetVar("test").Int())
}

// BenchmarkProgram_Execute_003-4          1000000000               0.01084 ns/op
// BenchmarkCompiler_Execute_003-4         1000000000               0.006692 ns/op
func BenchmarkProgram_Execute_003(b *testing.B) {
	p, _ := st.NewProgram(iftest)

	b.ResetTimer()
	for range 10000 {
		p.Execute()
	}
}

func BenchmarkCompiler_Execute_003(b *testing.B) {
	c := compiler.New()
	c.Compile(iftest)

	b.ResetTimer()
	for range 10000 {
		c.Execute()
	}
}
