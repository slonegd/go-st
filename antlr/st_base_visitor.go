// Code generated from ./antlr/ST.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ST

import "github.com/antlr4-go/antlr/v4"

type BaseSTVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSTVisitor) VisitPous(ctx *PousContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitFunction_decl(ctx *Function_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitFunction_block_decl(ctx *Function_block_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitType_declaration(ctx *Type_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitType_definition(ctx *Type_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitGlobal_var_declaration(ctx *Global_var_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitVar_declaration_block(ctx *Var_declaration_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitVar_decl(ctx *Var_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitData_type(ctx *Data_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitElementary_type_name(ctx *Elementary_type_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitArray_type(ctx *Array_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitSubrange(ctx *SubrangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitStructured_type(ctx *Structured_typeContext) interface{} {
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

func (v *BaseSTVisitor) VisitCase_statement(ctx *Case_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitCase_element(ctx *Case_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitCase_label(ctx *Case_labelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitFor_statement(ctx *For_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitWhile_statement(ctx *While_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitRepeat_statement(ctx *Repeat_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitFunction_invocation(ctx *Function_invocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitReturn_statement(ctx *Return_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitContinue_statement(ctx *Continue_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitExit_statement(ctx *Exit_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitBinaryExpression(ctx *BinaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitVarExpression(ctx *VarExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitParenExpression(ctx *ParenExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitLiteralExpression(ctx *LiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitFuncCallExpression(ctx *FuncCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitTyped_literal(ctx *Typed_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitType_name(ctx *Type_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitInput_decl(ctx *Input_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitOutput_decl(ctx *Output_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitEdge_decl(ctx *Edge_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitArray_conform_decl(ctx *Array_conform_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitArray_conformand(ctx *Array_conformandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitVariable_list(ctx *Variable_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitData_type_access(ctx *Data_type_accessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitDerived_type_access(ctx *Derived_type_accessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitSingle_elem_type_access(ctx *Single_elem_type_accessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitSimple_type_access(ctx *Simple_type_accessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitSubrange_type_access(ctx *Subrange_type_accessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitEnum_type_access(ctx *Enum_type_accessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitArray_type_access(ctx *Array_type_accessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitStruct_type_access(ctx *Struct_type_accessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitString_type_access(ctx *String_type_accessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitNamespace_name(ctx *Namespace_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitSimple_type_name(ctx *Simple_type_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitSubrange_type_name(ctx *Subrange_type_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitEnum_type_name(ctx *Enum_type_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitArray_type_name(ctx *Array_type_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSTVisitor) VisitStruct_type_name(ctx *Struct_type_nameContext) interface{} {
	return v.VisitChildren(ctx)
}
