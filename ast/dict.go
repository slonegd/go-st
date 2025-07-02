package ast

import (
	parser "github.com/slonegd/go-st/antlr"
	"github.com/slonegd/go-st/ops"
)

type LexerTokenType = int

var Ops = map[LexerTokenType]ops.Operator{
	parser.STLexerPLUS:  ops.Plus,
	parser.STLexerMINUS: ops.Minus,
	parser.STLexerMULT:  ops.Mult,
	parser.STLexerDIV:   ops.Div,
	parser.STLexerMOD:   ops.Mod,
}
