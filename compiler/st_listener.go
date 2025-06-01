package compiler

import (
	"log"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/slonegd/go-st/antlr"
	"github.com/slonegd/go-st/variant"
	"github.com/slonegd/go-st/vm"
)

// enters

// implements parser.STListener.
func (listener) EnterBinaryPowerExpr(c *parser.BinaryPowerExprContext)             {}
func (x listener) EnterAssignment_statement(c *parser.Assignment_statementContext) {}
func (x listener) EnterBinaryPlusExpr(c *parser.BinaryPlusExprContext)             {}
func (x listener) EnterConstant(c *parser.ConstantContext)                         {}
func (x listener) EnterEveryRule(ctx antlr.ParserRuleContext)                      {}
func (x listener) EnterNumber(c *parser.NumberContext)                             {}
func (x listener) EnterParenExpr(c *parser.ParenExprContext)                       {}
func (x listener) EnterProgram(c *parser.ProgramContext)                           {}
func (x listener) EnterStatement(c *parser.StatementContext)                       {}
func (x listener) EnterStatement_list(c *parser.Statement_listContext)             {}
func (x listener) EnterType_name(c *parser.Type_nameContext) {
	// TODO типизировать переменные
}
func (x listener) EnterVar_declaration(c *parser.Var_declarationContext) {
	varName := c.GetIdentifier().GetText()
	defaultValue := ""
	if token := c.GetDefault_(); token != nil {
		defaultValue = token.GetText()
	}
	valueType := c.GetType_().GetText()
	v := variant.SetType(variant.NewAnyVariant(defaultValue), variant.TypeFromString(valueType))
	x.varIndexes[varName] = len(x.Vars)
	x.VarNames[len(x.Vars)] = varName
	x.Vars = append(x.VM.Vars, v)
}
func (listener) EnterVar_declaration_block(c *parser.Var_declaration_blockContext)   {}
func (listener) EnterVar_declaration_blocks(c *parser.Var_declaration_blocksContext) {}
func (listener) EnterVariable(c *parser.VariableContext)                             {}
func (listener) EnterBinaryCompareExpr(c *parser.BinaryCompareExprContext)           {}
func (listener) EnterIf_statement(c *parser.If_statementContext)                     {}
func (listener) EnterElse_list(c *parser.Else_listContext)                           {}
func (listener) EnterThen_list(c *parser.Then_listContext)                           {}
func (listener) EnterCondition(c *parser.ConditionContext)                           {}
func (listener) EnterInteger(c *parser.IntegerContext)                               {}
func (listener) EnterSigned_integer(c *parser.Signed_integerContext)                 {}

// exits

func (x listener) ExitAssignment_statement(c *parser.Assignment_statementContext) {
	varName := c.GetLeft().GetText()
	i := x.varIndexes[varName]
	x.Bytecode.AddOp(vm.PopVar, uintptr(i))

}
func (x listener) ExitBinaryPlusExpr(c *parser.BinaryPlusExprContext) {
	op := c.GetOp().GetText()
	switch op { // TODO через id токена
	case "+":
		x.Bytecode.AddOp(vm.Plus)
	case "-":
		x.Bytecode.AddOp(vm.Sub)
	}
}
func (x listener) ExitBinaryPowerExpr(c *parser.BinaryPowerExprContext) {
	op := c.GetOp().GetText()
	switch op { // TODO через id токена
	case "*":
		x.Bytecode.AddOp(vm.Mult)
	case "/":
		x.VM.Bytecode.AddOp(vm.Div)
	case "MOD":
		x.Bytecode.AddOp(vm.Mod)
	}
}
func (x listener) ExitBinaryCompareExpr(c *parser.BinaryCompareExprContext) {
	op := c.GetOp().GetText()
	switch op { // TODO через id токена
	case ">":
		x.Bytecode.AddOp(vm.Gt)
	case ">=":
		x.VM.Bytecode.AddOp(vm.Gte)
	case "<":
		x.Bytecode.AddOp(vm.Lt)
	case "<=":
		x.Bytecode.AddOp(vm.Lte)
	case "=":
		x.Bytecode.AddOp(vm.Eq)
	case "<>":
		x.Bytecode.AddOp(vm.Neq)
	}
}

// ExitCondition implements parser.STListener.
func (x listener) ExitCondition(c *parser.ConditionContext) {
	i := x.Bytecode.AddOp(vm.IfFalse, 0) // дозаполним после Then
	x.conditionJumpIndexes.Push(i)
}
func (x listener) ExitThen_list(c *parser.Then_listContext) {
	i := x.Bytecode.AddOp(vm.Jump, 0) // // дозаполним после Then ExitIf
	x.thenJumpIndexes.ChangeLast(func(v []int) []int { return append(v, i) })
	jumpIndex := x.conditionJumpIndexes.Pop()
	x.Bytecode.ChangeOpArgs(jumpIndex, uintptr(len(x.Bytecode)))
}
func (x listener) ExitElse_list(c *parser.Else_listContext) {}

func (x listener) ExitIf_statement(c *parser.If_statementContext) {
	for _, jumpIndex := range x.thenJumpIndexes.Pop() {
		x.Bytecode.ChangeOpArgs(jumpIndex, uintptr(len(x.Bytecode)))
	}
}

func (x listener) ExitConstant(c *parser.ConstantContext)  {}
func (listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

func (x listener) ExitNumber(c *parser.NumberContext) {
	if _, ok := c.GetParent().(*parser.Var_declarationContext); ok {
		return // чтоб не класть на стек инициализацию
	}
	v := variant.NewAnyVariant(c.GetText())
	x.Bytecode.AddOp(vm.PushConst, uintptr(len(x.Consts)))
	x.Consts = append(x.Consts, v)
}

func (listener) ExitParenExpr(c *parser.ParenExprContext)                         {}
func (listener) ExitProgram(c *parser.ProgramContext)                             {}
func (listener) ExitStatement(c *parser.StatementContext)                         {}
func (listener) ExitStatement_list(c *parser.Statement_listContext)               {}
func (listener) ExitType_name(c *parser.Type_nameContext)                         {}
func (listener) ExitVar_declaration(c *parser.Var_declarationContext)             {}
func (listener) ExitVar_declaration_block(c *parser.Var_declaration_blockContext) {}
func (x listener) ExitVar_declaration_blocks(c *parser.Var_declaration_blocksContext) {
	log.Printf("variables %v, %v", x.varIndexes, x.Vars)
}
func (x listener) ExitVariable(c *parser.VariableContext) {
	varName := c.GetText()
	i := x.varIndexes[varName]
	x.Bytecode.AddOp(vm.PushVar, uintptr(i))
}

func (listener) ExitInteger(c *parser.IntegerContext)               {}
func (listener) ExitSigned_integer(c *parser.Signed_integerContext) {}
func (listener) VisitErrorNode(node antlr.ErrorNode)                {}
func (listener) VisitTerminal(node antlr.TerminalNode)              {}
