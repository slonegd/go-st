package st

import (
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/slonegd/go-st/antlr"
)

// implements parser.STListener.
// EnterBinaryPowerExpr implements parser.STListener.
func (*Program) EnterBinaryPowerExpr(c *parser.BinaryPowerExprContext)             {}
func (x *Program) EnterAssignment_statement(c *parser.Assignment_statementContext) {}
func (x *Program) EnterBinaryPlusExpr(c *parser.BinaryPlusExprContext)             {}
func (x *Program) EnterConstant(c *parser.ConstantContext)                         {}
func (x *Program) EnterEveryRule(ctx antlr.ParserRuleContext)                      {}
func (x *Program) EnterNumber(c *parser.NumberContext)                             {}
func (x *Program) EnterParenExpr(c *parser.ParenExprContext)                       {}
func (x *Program) EnterProgram(c *parser.ProgramContext) {
	x.variablePrefix = c.GetIdentifier().GetText()
}
func (x *Program) EnterStatement(c *parser.StatementContext)           {}
func (x *Program) EnterStatement_list(c *parser.Statement_listContext) {}
func (x *Program) EnterType_name(c *parser.Type_nameContext) {
	// TODO типизировать переменные
}
func (x *Program) EnterVar_declaration(c *parser.Var_declarationContext) {
	varName := x.variablePrefix + "." + c.GetIdentifier().GetText()
	x.Variables[varName] = 0 // TODO значения по умолчанию
}
func (*Program) EnterVar_declaration_block(c *parser.Var_declaration_blockContext)   {}
func (*Program) EnterVar_declaration_blocks(c *parser.Var_declaration_blocksContext) {}
func (*Program) EnterVariable(c *parser.VariableContext)                             {}

func (x *Program) ExitAssignment_statement(c *parser.Assignment_statementContext) {
	varName := x.variablePrefix + "." + c.GetLeft().GetText()
	x.actions = append(x.actions, func() {
		x.Variables[varName] = x.stack[len(x.stack)-1]
		x.stack = x.stack[:len(x.stack)-1]
	})
}
func (x *Program) ExitBinaryPlusExpr(c *parser.BinaryPlusExprContext) {
	x.actions = append(x.actions, func() {
		right := x.stack[len(x.stack)-1]
		left := x.stack[len(x.stack)-2]
		x.stack = x.stack[:len(x.stack)-2]
		switch c.GetOp().GetText() { // TODO через id токена
		case "+":
			x.stack = append(x.stack, left+right)
		case "-":
			x.stack = append(x.stack, left-right)
		}
	})
}
func (x *Program) ExitBinaryPowerExpr(c *parser.BinaryPowerExprContext) {
	x.actions = append(x.actions, func() {
		right := x.stack[len(x.stack)-1]
		left := x.stack[len(x.stack)-2]
		x.stack = x.stack[:len(x.stack)-2]
		switch c.GetOp().GetText() { // TODO через id токена или до замыкания
		case "*":
			x.stack = append(x.stack, left*right)
		case "/":
			x.stack = append(x.stack, left/right)
		}
	})
}

func (x *Program) ExitConstant(c *parser.ConstantContext)  {}
func (*Program) ExitEveryRule(ctx antlr.ParserRuleContext) {}

func (x *Program) ExitNumber(c *parser.NumberContext) {
	i, _ := strconv.Atoi(c.GetText())
	x.actions = append(x.actions, func() {
		x.stack = append(x.stack, i)
	})
}

func (*Program) ExitParenExpr(c *parser.ParenExprContext)                           {}
func (*Program) ExitProgram(c *parser.ProgramContext)                               {}
func (*Program) ExitStatement(c *parser.StatementContext)                           {}
func (*Program) ExitStatement_list(c *parser.Statement_listContext)                 {}
func (*Program) ExitType_name(c *parser.Type_nameContext)                           {}
func (*Program) ExitVar_declaration(c *parser.Var_declarationContext)               {}
func (*Program) ExitVar_declaration_block(c *parser.Var_declaration_blockContext)   {}
func (*Program) ExitVar_declaration_blocks(c *parser.Var_declaration_blocksContext) {}
func (x *Program) ExitVariable(c *parser.VariableContext) {
	varName := x.variablePrefix + "." + c.GetText()
	x.actions = append(x.actions, func() {
		x.stack = append(x.stack, x.Variables[varName])
	})
}
func (*Program) VisitErrorNode(node antlr.ErrorNode)   {}
func (*Program) VisitTerminal(node antlr.TerminalNode) {}
