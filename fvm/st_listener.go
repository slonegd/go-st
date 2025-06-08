package fvm

import (
	"log"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/slonegd/go-st/antlr"
	"github.com/slonegd/go-st/variant"
)

// enters

// implements parser.STListener.
func (FVM) EnterBinaryPowerExpr(c *parser.BinaryPowerExprContext)              {}
func (x *FVM) EnterAssignment_statement(c *parser.Assignment_statementContext) {}
func (x *FVM) EnterBinaryPlusExpr(c *parser.BinaryPlusExprContext)             {}
func (x *FVM) EnterConstant(c *parser.ConstantContext)                         {}
func (x *FVM) EnterEveryRule(ctx antlr.ParserRuleContext)                      {}
func (x *FVM) EnterNumber(c *parser.NumberContext)                             {}
func (x *FVM) EnterParenExpr(c *parser.ParenExprContext)                       {}
func (x *FVM) EnterProgram(c *parser.ProgramContext)                           {}
func (x *FVM) EnterStatement(c *parser.StatementContext)                       {}
func (x *FVM) EnterStatement_list(c *parser.Statement_listContext)             {}
func (x *FVM) EnterType_name(c *parser.Type_nameContext) {
	// TODO типизировать переменные
}
func (x *FVM) EnterVar_declaration(c *parser.Var_declarationContext) {
	varName := c.GetIdentifier().GetText()
	defaultValue := ""
	if token := c.GetDefault_(); token != nil {
		defaultValue = token.GetText()
	}
	valueType := c.GetType_().GetText()
	v := variant.SetType(variant.NewAnyVariant(defaultValue), variant.TypeFromString(valueType))
	x.vars[varName] = v
}
func (FVM) EnterVar_declaration_block(c *parser.Var_declaration_blockContext)   {}
func (FVM) EnterVar_declaration_blocks(c *parser.Var_declaration_blocksContext) {}
func (FVM) EnterVariable(c *parser.VariableContext)                             {}
func (FVM) EnterBinaryCompareExpr(c *parser.BinaryCompareExprContext)           {}
func (FVM) EnterIf_statement(c *parser.If_statementContext)                     {}
func (FVM) EnterElse_list(c *parser.Else_listContext)                           {}
func (FVM) EnterThen_list(c *parser.Then_listContext)                           {}
func (FVM) EnterCondition(c *parser.ConditionContext)                           {}
func (FVM) EnterInteger(c *parser.IntegerContext)                               {}
func (FVM) EnterSigned_integer(c *parser.Signed_integerContext)                 {}

// exits

func (x *FVM) ExitAssignment_statement(c *parser.Assignment_statementContext) {
	varName := c.GetLeft().GetText()
	v := x.vars[varName]
	x.operators.Push(NewAssignStep((*int64)(v.Pointer()), x.operators.Pop().(Int64Operator)))
}
func (x *FVM) ExitBinaryPlusExpr(c *parser.BinaryPlusExprContext) {
	op := c.GetOp().GetText()
	right := x.operators.Pop().(Int64Operator)
	left := x.operators.Pop().(Int64Operator)
	switch op { // TODO через id токена
	case "+":
		x.operators.Push(NewBinaryPlusStep(left, right))
	case "-":
		x.operators.Push(NewBinarySubStep(left, right))
	}
}
func (x *FVM) ExitBinaryPowerExpr(c *parser.BinaryPowerExprContext) {
	op := c.GetOp().GetText()
	right := x.operators.Pop().(Int64Operator)
	left := x.operators.Pop().(Int64Operator)
	switch op { // TODO через id токена
	case "*":
		x.operators.Push(NewBinaryMultStep(left, right))
	case "/":
		x.operators.Push(NewBinaryDivStep(left, right))
	case "MOD":
		x.operators.Push(NewBinaryModStep(left, right))
	}
}
func (x *FVM) ExitBinaryCompareExpr(c *parser.BinaryCompareExprContext) {
	// op := c.GetOp().GetText()
	// switch op { // TODO через id токена
	// case ">":
	// 	x.Bytecode.AddOp(vm.Gt)
	// case ">=":
	// 	x.VM.Bytecode.AddOp(vm.Gte)
	// case "<":
	// 	x.Bytecode.AddOp(vm.Lt)
	// case "<=":
	// 	x.Bytecode.AddOp(vm.Lte)
	// case "=":
	// 	x.Bytecode.AddOp(vm.Eq)
	// case "<>":
	// 	x.Bytecode.AddOp(vm.Neq)
	// }
}

// ExitCondition implements parser.STListener.
func (x *FVM) ExitCondition(c *parser.ConditionContext) {
	// i := x.Bytecode.AddOp(vm.IfFalse, 0) // дозаполним после Then
	// x.conditionJumpIndexes.Push(i)
}
func (x *FVM) ExitThen_list(c *parser.Then_listContext) {
	// i := x.Bytecode.AddOp(vm.Jump, 0) // // дозаполним после Then ExitIf
	// x.thenJumpIndexes.ChangeLast(func(v []int) []int { return append(v, i) })
	// jumpIndex := x.conditionJumpIndexes.Pop()
	// x.Bytecode.ChangeOpArgs(jumpIndex, uintptr(len(x.Bytecode)))
}
func (x *FVM) ExitElse_list(c *parser.Else_listContext) {}

func (x *FVM) ExitIf_statement(c *parser.If_statementContext) {
	//	for _, jumpIndex := range x.thenJumpIndexes.Pop() {
	//		x.Bytecode.ChangeOpArgs(jumpIndex, uintptr(len(x.Bytecode)))
	//	}
}

func (x *FVM) ExitConstant(c *parser.ConstantContext) {}
func (FVM) ExitEveryRule(ctx antlr.ParserRuleContext) {}

func (x *FVM) ExitNumber(c *parser.NumberContext) {
	if _, ok := c.GetParent().(*parser.Var_declarationContext); ok {
		return // чтоб не класть на стек инициализацию
	}
	v := variant.NewAnyVariant(c.GetText())
	x.operators.Push(NewConstantStep(v.Int()))
}

func (FVM) ExitParenExpr(c *parser.ParenExprContext)                         {}
func (FVM) ExitProgram(c *parser.ProgramContext)                             {}
func (FVM) ExitStatement(c *parser.StatementContext)                         {}
func (FVM) ExitStatement_list(c *parser.Statement_listContext)               {}
func (FVM) ExitType_name(c *parser.Type_nameContext)                         {}
func (FVM) ExitVar_declaration(c *parser.Var_declarationContext)             {}
func (FVM) ExitVar_declaration_block(c *parser.Var_declaration_blockContext) {}
func (x *FVM) ExitVar_declaration_blocks(c *parser.Var_declaration_blocksContext) {
	log.Printf("variables %v", x.vars)
}
func (x *FVM) ExitVariable(c *parser.VariableContext) {
	varName := c.GetText()
	v := x.vars[varName]
	x.operators.Push(NewVarStep((*int64)(v.Pointer())))
}

func (FVM) ExitInteger(c *parser.IntegerContext)               {}
func (FVM) ExitSigned_integer(c *parser.Signed_integerContext) {}
func (FVM) VisitErrorNode(node antlr.ErrorNode)                {}
func (FVM) VisitTerminal(node antlr.TerminalNode)              {}
