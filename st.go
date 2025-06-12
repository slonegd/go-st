package st

import (
	"fmt"

	"github.com/slonegd/go-st/ast"
)

func NewProgram(script string) (*ast.AST, error) {
	r, err := ast.Parse(script)
	if err != nil {
		return nil, fmt.Errorf("compile: %w", err)
	}
	return r, nil
}
