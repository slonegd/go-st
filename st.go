package st

import (
	"fmt"

	"github.com/slonegd/go-st/compiler"
)

type (
	Program struct {
		*compiler.Compiler
	}
)

func NewProgram(script string) (*Program, error) {
	p := &Program{Compiler: compiler.New()}
	err := p.Compiler.Compile(script)
	if err != nil {
		return nil, fmt.Errorf("compile: %w", err)
	}
	return p, nil
}
