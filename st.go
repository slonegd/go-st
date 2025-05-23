package st

import (
	"fmt"
	"log"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/slonegd/go-st/antlr"
	"github.com/slonegd/go-st/variant"
)

type (
	Program struct {
		source    Source
		Variables map[string]variant.Variant
		steps     []Step
		stack     Stack
		ifs       []*ifState
		err       error
	}
	ifState struct {
		lastCondition      bool
		lastConditionIndex int
		thenIndexes        []int
	}
	Stack []variant.Variant
	Step  struct {
		number   int // номер по порядку
		action   string
		callback func() int // если возврат != 0 - перейти к действию (для ветвления)
	}
)

var (
	_ parser.STListener   = (*Program)(nil)
	_ antlr.ErrorListener = (*Program)(nil)
)

func NewProgram(script string) (*Program, error) {
	result := &Program{
		source:    []rune(script),
		Variables: make(map[string]variant.Variant),
	}
	// Setup the input
	is := antlr.NewInputStream(script)

	// Create the Lexer
	lexer := parser.NewSTLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewSTParser(stream)
	p.BaseRecognizer.AddErrorListener(result)

	// Finally parse the expression (by walking the tree)
	antlr.ParseTreeWalkerDefault.Walk(result, p.Program())
	if p.SynErr != nil && p.SynErr.GetMessage() != "" {
		result.err = fmt.Errorf(p.SynErr.GetMessage())
	}

	return result, result.err
}

func (x *Program) Execute() {
	for i := 0; i < len(x.steps); {
		next := x.steps[i].callback()
		if next != 0 {
			i = next
			continue
		}
		i++
	}
}

func (x *Program) Print() {
	for _, s := range x.steps {
		log.Printf(s.action)
	}
}

func (x *Stack) Push(v variant.Variant) {
	*x = append(*x, v)
}

func (x *Stack) Pop() variant.Variant {
	size := len(*x)
	if size == 0 {
		return variant.NewAnyVariant("")
	}
	r := (*x)[size-1]
	*x = (*x)[:size-1]
	return r
}
