// Code generated from ./g4/st.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // st

import "github.com/antlr4-go/antlr/v4"

// BasestListener is a complete listener for a parse tree produced by stParser.
type BasestListener struct{}

var _ stListener = &BasestListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasestListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasestListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasestListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasestListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProrgamm is called when production prorgamm is entered.
func (s *BasestListener) EnterProrgamm(ctx *ProrgammContext) {}

// ExitProrgamm is called when production prorgamm is exited.
func (s *BasestListener) ExitProrgamm(ctx *ProrgammContext) {}

// EnterVar_declaration_blocks is called when production var_declaration_blocks is entered.
func (s *BasestListener) EnterVar_declaration_blocks(ctx *Var_declaration_blocksContext) {}

// ExitVar_declaration_blocks is called when production var_declaration_blocks is exited.
func (s *BasestListener) ExitVar_declaration_blocks(ctx *Var_declaration_blocksContext) {}

// EnterVar_declaration_block is called when production var_declaration_block is entered.
func (s *BasestListener) EnterVar_declaration_block(ctx *Var_declaration_blockContext) {}

// ExitVar_declaration_block is called when production var_declaration_block is exited.
func (s *BasestListener) ExitVar_declaration_block(ctx *Var_declaration_blockContext) {}

// EnterVar_declaration is called when production var_declaration is entered.
func (s *BasestListener) EnterVar_declaration(ctx *Var_declarationContext) {}

// ExitVar_declaration is called when production var_declaration is exited.
func (s *BasestListener) ExitVar_declaration(ctx *Var_declarationContext) {}

// EnterType_name is called when production type_name is entered.
func (s *BasestListener) EnterType_name(ctx *Type_nameContext) {}

// ExitType_name is called when production type_name is exited.
func (s *BasestListener) ExitType_name(ctx *Type_nameContext) {}

// EnterStatement_list is called when production statement_list is entered.
func (s *BasestListener) EnterStatement_list(ctx *Statement_listContext) {}

// ExitStatement_list is called when production statement_list is exited.
func (s *BasestListener) ExitStatement_list(ctx *Statement_listContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasestListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasestListener) ExitStatement(ctx *StatementContext) {}

// EnterAssignment_statement is called when production assignment_statement is entered.
func (s *BasestListener) EnterAssignment_statement(ctx *Assignment_statementContext) {}

// ExitAssignment_statement is called when production assignment_statement is exited.
func (s *BasestListener) ExitAssignment_statement(ctx *Assignment_statementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasestListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasestListener) ExitExpression(ctx *ExpressionContext) {}

// EnterConstant is called when production constant is entered.
func (s *BasestListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BasestListener) ExitConstant(ctx *ConstantContext) {}

// EnterNumber is called when production number is entered.
func (s *BasestListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BasestListener) ExitNumber(ctx *NumberContext) {}
