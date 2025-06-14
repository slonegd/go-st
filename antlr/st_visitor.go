// Code generated from ./antlr/ST.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ST

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by STParser.
type STVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by STParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by STParser#var_declaration_blocks.
	VisitVar_declaration_blocks(ctx *Var_declaration_blocksContext) interface{}

	// Visit a parse tree produced by STParser#var_declaration_block.
	VisitVar_declaration_block(ctx *Var_declaration_blockContext) interface{}

	// Visit a parse tree produced by STParser#var_declaration.
	VisitVar_declaration(ctx *Var_declarationContext) interface{}

	// Visit a parse tree produced by STParser#type_name.
	VisitType_name(ctx *Type_nameContext) interface{}

	// Visit a parse tree produced by STParser#statement_list.
	VisitStatement_list(ctx *Statement_listContext) interface{}

	// Visit a parse tree produced by STParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by STParser#assignment_statement.
	VisitAssignment_statement(ctx *Assignment_statementContext) interface{}

	// Visit a parse tree produced by STParser#if_statement.
	VisitIf_statement(ctx *If_statementContext) interface{}

	// Visit a parse tree produced by STParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by STParser#then_list.
	VisitThen_list(ctx *Then_listContext) interface{}

	// Visit a parse tree produced by STParser#else_list.
	VisitElse_list(ctx *Else_listContext) interface{}

	// Visit a parse tree produced by STParser#binaryCompareExpr.
	VisitBinaryCompareExpr(ctx *BinaryCompareExprContext) interface{}

	// Visit a parse tree produced by STParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by STParser#binaryPowerExpr.
	VisitBinaryPowerExpr(ctx *BinaryPowerExprContext) interface{}

	// Visit a parse tree produced by STParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by STParser#binaryPlusExpr.
	VisitBinaryPlusExpr(ctx *BinaryPlusExprContext) interface{}

	// Visit a parse tree produced by STParser#callExpr.
	VisitCallExpr(ctx *CallExprContext) interface{}

	// Visit a parse tree produced by STParser#parenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by STParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by STParser#real.
	VisitReal(ctx *RealContext) interface{}

	// Visit a parse tree produced by STParser#integer.
	VisitInteger(ctx *IntegerContext) interface{}

	// Visit a parse tree produced by STParser#signed_integer.
	VisitSigned_integer(ctx *Signed_integerContext) interface{}

	// Visit a parse tree produced by STParser#unsign_integer.
	VisitUnsign_integer(ctx *Unsign_integerContext) interface{}
}
