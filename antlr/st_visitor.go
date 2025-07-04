// Code generated from ./antlr/ST.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ST

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by STParser.
type STVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by STParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by STParser#function_decl.
	VisitFunction_decl(ctx *Function_declContext) interface{}

	// Visit a parse tree produced by STParser#function_block_decl.
	VisitFunction_block_decl(ctx *Function_block_declContext) interface{}

	// Visit a parse tree produced by STParser#var_declaration_block.
	VisitVar_declaration_block(ctx *Var_declaration_blockContext) interface{}

	// Visit a parse tree produced by STParser#var_decl.
	VisitVar_decl(ctx *Var_declContext) interface{}

	// Visit a parse tree produced by STParser#data_type.
	VisitData_type(ctx *Data_typeContext) interface{}

	// Visit a parse tree produced by STParser#elementary_type_name.
	VisitElementary_type_name(ctx *Elementary_type_nameContext) interface{}

	// Visit a parse tree produced by STParser#array_type.
	VisitArray_type(ctx *Array_typeContext) interface{}

	// Visit a parse tree produced by STParser#subrange.
	VisitSubrange(ctx *SubrangeContext) interface{}

	// Visit a parse tree produced by STParser#structured_type.
	VisitStructured_type(ctx *Structured_typeContext) interface{}

	// Visit a parse tree produced by STParser#statement_list.
	VisitStatement_list(ctx *Statement_listContext) interface{}

	// Visit a parse tree produced by STParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by STParser#assignment_statement.
	VisitAssignment_statement(ctx *Assignment_statementContext) interface{}

	// Visit a parse tree produced by STParser#if_statement.
	VisitIf_statement(ctx *If_statementContext) interface{}

	// Visit a parse tree produced by STParser#case_statement.
	VisitCase_statement(ctx *Case_statementContext) interface{}

	// Visit a parse tree produced by STParser#case_element.
	VisitCase_element(ctx *Case_elementContext) interface{}

	// Visit a parse tree produced by STParser#case_label.
	VisitCase_label(ctx *Case_labelContext) interface{}

	// Visit a parse tree produced by STParser#for_statement.
	VisitFor_statement(ctx *For_statementContext) interface{}

	// Visit a parse tree produced by STParser#while_statement.
	VisitWhile_statement(ctx *While_statementContext) interface{}

	// Visit a parse tree produced by STParser#repeat_statement.
	VisitRepeat_statement(ctx *Repeat_statementContext) interface{}

	// Visit a parse tree produced by STParser#function_invocation.
	VisitFunction_invocation(ctx *Function_invocationContext) interface{}

	// Visit a parse tree produced by STParser#return_statement.
	VisitReturn_statement(ctx *Return_statementContext) interface{}

	// Visit a parse tree produced by STParser#continue_statement.
	VisitContinue_statement(ctx *Continue_statementContext) interface{}

	// Visit a parse tree produced by STParser#exit_statement.
	VisitExit_statement(ctx *Exit_statementContext) interface{}

	// Visit a parse tree produced by STParser#binaryExpression.
	VisitBinaryExpression(ctx *BinaryExpressionContext) interface{}

	// Visit a parse tree produced by STParser#varExpression.
	VisitVarExpression(ctx *VarExpressionContext) interface{}

	// Visit a parse tree produced by STParser#parenExpression.
	VisitParenExpression(ctx *ParenExpressionContext) interface{}

	// Visit a parse tree produced by STParser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by STParser#literalExpression.
	VisitLiteralExpression(ctx *LiteralExpressionContext) interface{}

	// Visit a parse tree produced by STParser#funcCallExpression.
	VisitFuncCallExpression(ctx *FuncCallExpressionContext) interface{}

	// Visit a parse tree produced by STParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by STParser#typed_literal.
	VisitTyped_literal(ctx *Typed_literalContext) interface{}

	// Visit a parse tree produced by STParser#type_name.
	VisitType_name(ctx *Type_nameContext) interface{}

	// Visit a parse tree produced by STParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by STParser#input_decl.
	VisitInput_decl(ctx *Input_declContext) interface{}

	// Visit a parse tree produced by STParser#output_decl.
	VisitOutput_decl(ctx *Output_declContext) interface{}

	// Visit a parse tree produced by STParser#edge_decl.
	VisitEdge_decl(ctx *Edge_declContext) interface{}

	// Visit a parse tree produced by STParser#array_conform_decl.
	VisitArray_conform_decl(ctx *Array_conform_declContext) interface{}

	// Visit a parse tree produced by STParser#array_conformand.
	VisitArray_conformand(ctx *Array_conformandContext) interface{}

	// Visit a parse tree produced by STParser#variable_list.
	VisitVariable_list(ctx *Variable_listContext) interface{}

	// Visit a parse tree produced by STParser#data_type_access.
	VisitData_type_access(ctx *Data_type_accessContext) interface{}

	// Visit a parse tree produced by STParser#derived_type_access.
	VisitDerived_type_access(ctx *Derived_type_accessContext) interface{}

	// Visit a parse tree produced by STParser#single_elem_type_access.
	VisitSingle_elem_type_access(ctx *Single_elem_type_accessContext) interface{}

	// Visit a parse tree produced by STParser#simple_type_access.
	VisitSimple_type_access(ctx *Simple_type_accessContext) interface{}

	// Visit a parse tree produced by STParser#subrange_type_access.
	VisitSubrange_type_access(ctx *Subrange_type_accessContext) interface{}

	// Visit a parse tree produced by STParser#enum_type_access.
	VisitEnum_type_access(ctx *Enum_type_accessContext) interface{}

	// Visit a parse tree produced by STParser#array_type_access.
	VisitArray_type_access(ctx *Array_type_accessContext) interface{}

	// Visit a parse tree produced by STParser#struct_type_access.
	VisitStruct_type_access(ctx *Struct_type_accessContext) interface{}

	// Visit a parse tree produced by STParser#string_type_access.
	VisitString_type_access(ctx *String_type_accessContext) interface{}

	// Visit a parse tree produced by STParser#namespace_name.
	VisitNamespace_name(ctx *Namespace_nameContext) interface{}

	// Visit a parse tree produced by STParser#simple_type_name.
	VisitSimple_type_name(ctx *Simple_type_nameContext) interface{}

	// Visit a parse tree produced by STParser#subrange_type_name.
	VisitSubrange_type_name(ctx *Subrange_type_nameContext) interface{}

	// Visit a parse tree produced by STParser#enum_type_name.
	VisitEnum_type_name(ctx *Enum_type_nameContext) interface{}

	// Visit a parse tree produced by STParser#array_type_name.
	VisitArray_type_name(ctx *Array_type_nameContext) interface{}

	// Visit a parse tree produced by STParser#struct_type_name.
	VisitStruct_type_name(ctx *Struct_type_nameContext) interface{}
}
