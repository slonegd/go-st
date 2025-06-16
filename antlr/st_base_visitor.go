// Code generated from ./antlr/ST.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ST

import "github.com/antlr4-go/antlr/v4"

type BaseSTVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSTVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitVar_declaration_blocks(ctx *Var_declaration_blocksContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitVar_declaration_block(ctx *Var_declaration_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitVar_declaration(ctx *Var_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitType_name(ctx *Type_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitStatement_list(ctx *Statement_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitAssignment_statement(ctx *Assignment_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitIf_statement(ctx *If_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitThen_list(ctx *Then_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitElse_list(ctx *Else_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitBinaryCompareExpr(ctx *BinaryCompareExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitBinaryPowerExpr(ctx *BinaryPowerExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitBinaryPlusExpr(ctx *BinaryPlusExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitCallExpr(ctx *CallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitInteger(ctx *IntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitSigned_integer(ctx *Signed_integerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitUnsign_integer(ctx *Unsign_integerContext) interface{} {
	return v.VisitChildren(ctx)
}
