// Code generated from ./antlr/ST.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ST

import "github.com/antlr4-go/antlr/v4"

// STListener is a complete listener for a parse tree produced by STParser.
type STListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterVar_declaration_blocks is called when entering the var_declaration_blocks production.
	EnterVar_declaration_blocks(c *Var_declaration_blocksContext)

	// EnterVar_declaration_block is called when entering the var_declaration_block production.
	EnterVar_declaration_block(c *Var_declaration_blockContext)

	// EnterVar_declaration is called when entering the var_declaration production.
	EnterVar_declaration(c *Var_declarationContext)

	// EnterType_name is called when entering the type_name production.
	EnterType_name(c *Type_nameContext)

	// EnterStatement_list is called when entering the statement_list production.
	EnterStatement_list(c *Statement_listContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAssignment_statement is called when entering the assignment_statement production.
	EnterAssignment_statement(c *Assignment_statementContext)

	// EnterIf_statement is called when entering the if_statement production.
	EnterIf_statement(c *If_statementContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterThen_list is called when entering the then_list production.
	EnterThen_list(c *Then_listContext)

	// EnterElse_list is called when entering the else_list production.
	EnterElse_list(c *Else_listContext)

	// EnterBinaryCompareExpr is called when entering the binaryCompareExpr production.
	EnterBinaryCompareExpr(c *BinaryCompareExprContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterBinaryPowerExpr is called when entering the binaryPowerExpr production.
	EnterBinaryPowerExpr(c *BinaryPowerExprContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterBinaryPlusExpr is called when entering the binaryPlusExpr production.
	EnterBinaryPlusExpr(c *BinaryPlusExprContext)

	// EnterParenExpr is called when entering the parenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitVar_declaration_blocks is called when exiting the var_declaration_blocks production.
	ExitVar_declaration_blocks(c *Var_declaration_blocksContext)

	// ExitVar_declaration_block is called when exiting the var_declaration_block production.
	ExitVar_declaration_block(c *Var_declaration_blockContext)

	// ExitVar_declaration is called when exiting the var_declaration production.
	ExitVar_declaration(c *Var_declarationContext)

	// ExitType_name is called when exiting the type_name production.
	ExitType_name(c *Type_nameContext)

	// ExitStatement_list is called when exiting the statement_list production.
	ExitStatement_list(c *Statement_listContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAssignment_statement is called when exiting the assignment_statement production.
	ExitAssignment_statement(c *Assignment_statementContext)

	// ExitIf_statement is called when exiting the if_statement production.
	ExitIf_statement(c *If_statementContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitThen_list is called when exiting the then_list production.
	ExitThen_list(c *Then_listContext)

	// ExitElse_list is called when exiting the else_list production.
	ExitElse_list(c *Else_listContext)

	// ExitBinaryCompareExpr is called when exiting the binaryCompareExpr production.
	ExitBinaryCompareExpr(c *BinaryCompareExprContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitBinaryPowerExpr is called when exiting the binaryPowerExpr production.
	ExitBinaryPowerExpr(c *BinaryPowerExprContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitBinaryPlusExpr is called when exiting the binaryPlusExpr production.
	ExitBinaryPlusExpr(c *BinaryPlusExprContext)

	// ExitParenExpr is called when exiting the parenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)
}
