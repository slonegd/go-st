// Code generated from ./antlr/ST.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ST

import "github.com/antlr4-go/antlr/v4"

// BaseSTListener is a complete listener for a parse tree produced by STParser.
type BaseSTListener struct{}

var _ STListener = &BaseSTListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSTListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSTListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSTListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSTListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseSTListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseSTListener) ExitProgram(ctx *ProgramContext) {}

// EnterVar_declaration_blocks is called when production var_declaration_blocks is entered.
func (s *BaseSTListener) EnterVar_declaration_blocks(ctx *Var_declaration_blocksContext) {}

// ExitVar_declaration_blocks is called when production var_declaration_blocks is exited.
func (s *BaseSTListener) ExitVar_declaration_blocks(ctx *Var_declaration_blocksContext) {}

// EnterVar_declaration_block is called when production var_declaration_block is entered.
func (s *BaseSTListener) EnterVar_declaration_block(ctx *Var_declaration_blockContext) {}

// ExitVar_declaration_block is called when production var_declaration_block is exited.
func (s *BaseSTListener) ExitVar_declaration_block(ctx *Var_declaration_blockContext) {}

// EnterVar_declaration is called when production var_declaration is entered.
func (s *BaseSTListener) EnterVar_declaration(ctx *Var_declarationContext) {}

// ExitVar_declaration is called when production var_declaration is exited.
func (s *BaseSTListener) ExitVar_declaration(ctx *Var_declarationContext) {}

// EnterType_name is called when production type_name is entered.
func (s *BaseSTListener) EnterType_name(ctx *Type_nameContext) {}

// ExitType_name is called when production type_name is exited.
func (s *BaseSTListener) ExitType_name(ctx *Type_nameContext) {}

// EnterStatement_list is called when production statement_list is entered.
func (s *BaseSTListener) EnterStatement_list(ctx *Statement_listContext) {}

// ExitStatement_list is called when production statement_list is exited.
func (s *BaseSTListener) ExitStatement_list(ctx *Statement_listContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseSTListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseSTListener) ExitStatement(ctx *StatementContext) {}

// EnterAssignment_statement is called when production assignment_statement is entered.
func (s *BaseSTListener) EnterAssignment_statement(ctx *Assignment_statementContext) {}

// ExitAssignment_statement is called when production assignment_statement is exited.
func (s *BaseSTListener) ExitAssignment_statement(ctx *Assignment_statementContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseSTListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseSTListener) ExitConstant(ctx *ConstantContext) {}

// EnterBinaryPowerExpr is called when production binaryPowerExpr is entered.
func (s *BaseSTListener) EnterBinaryPowerExpr(ctx *BinaryPowerExprContext) {}

// ExitBinaryPowerExpr is called when production binaryPowerExpr is exited.
func (s *BaseSTListener) ExitBinaryPowerExpr(ctx *BinaryPowerExprContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseSTListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseSTListener) ExitVariable(ctx *VariableContext) {}

// EnterBinaryPlusExpr is called when production binaryPlusExpr is entered.
func (s *BaseSTListener) EnterBinaryPlusExpr(ctx *BinaryPlusExprContext) {}

// ExitBinaryPlusExpr is called when production binaryPlusExpr is exited.
func (s *BaseSTListener) ExitBinaryPlusExpr(ctx *BinaryPlusExprContext) {}

// EnterParenExpr is called when production parenExpr is entered.
func (s *BaseSTListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production parenExpr is exited.
func (s *BaseSTListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseSTListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseSTListener) ExitNumber(ctx *NumberContext) {}
