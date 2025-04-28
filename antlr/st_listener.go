// Code generated from ./antlr/ST.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ST

import "github.com/antlr4-go/antlr/v4"

// STListener is a complete listener for a parse tree produced by STParser.
type STListener interface {
	antlr.ParseTreeListener

	// EnterProgramm is called when entering the programm production.
	EnterProgramm(c *ProgrammContext)

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

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// ExitProgramm is called when exiting the programm production.
	ExitProgramm(c *ProgrammContext)

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

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)
}
