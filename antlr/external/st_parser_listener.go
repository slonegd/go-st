// Code generated from ./antlr/external/ST_parser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ST_parser

import "github.com/antlr4-go/antlr/v4"

// ST_parserListener is a complete listener for a parse tree produced by ST_parser.
type ST_parserListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterNamespace_declaration is called when entering the namespace_declaration production.
	EnterNamespace_declaration(c *Namespace_declarationContext)

	// EnterNamespace_elements is called when entering the namespace_elements production.
	EnterNamespace_elements(c *Namespace_elementsContext)

	// EnterFull_qualified_identifier is called when entering the full_qualified_identifier production.
	EnterFull_qualified_identifier(c *Full_qualified_identifierContext)

	// EnterUsing_directive is called when entering the using_directive production.
	EnterUsing_directive(c *Using_directiveContext)

	// EnterLibrary_element_declaration is called when entering the library_element_declaration production.
	EnterLibrary_element_declaration(c *Library_element_declarationContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterCast is called when entering the cast production.
	EnterCast(c *CastContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// EnterBits is called when entering the bits production.
	EnterBits(c *BitsContext)

	// EnterReal is called when entering the real production.
	EnterReal(c *RealContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterTime is called when entering the time production.
	EnterTime(c *TimeContext)

	// EnterTimeofday is called when entering the timeofday production.
	EnterTimeofday(c *TimeofdayContext)

	// EnterDate is called when entering the date production.
	EnterDate(c *DateContext)

	// EnterDatetime is called when entering the datetime production.
	EnterDatetime(c *DatetimeContext)

	// EnterRef_null is called when entering the ref_null production.
	EnterRef_null(c *Ref_nullContext)

	// EnterData_type_name is called when entering the data_type_name production.
	EnterData_type_name(c *Data_type_nameContext)

	// EnterNon_generic_type_name is called when entering the non_generic_type_name production.
	EnterNon_generic_type_name(c *Non_generic_type_nameContext)

	// EnterElementary_type_name is called when entering the elementary_type_name production.
	EnterElementary_type_name(c *Elementary_type_nameContext)

	// EnterNumeric_type_name is called when entering the numeric_type_name production.
	EnterNumeric_type_name(c *Numeric_type_nameContext)

	// EnterInteger_type_name is called when entering the integer_type_name production.
	EnterInteger_type_name(c *Integer_type_nameContext)

	// EnterSigned_integer_type_name is called when entering the signed_integer_type_name production.
	EnterSigned_integer_type_name(c *Signed_integer_type_nameContext)

	// EnterUnsigned_integer_type_name is called when entering the unsigned_integer_type_name production.
	EnterUnsigned_integer_type_name(c *Unsigned_integer_type_nameContext)

	// EnterReal_type_name is called when entering the real_type_name production.
	EnterReal_type_name(c *Real_type_nameContext)

	// EnterDate_type_name is called when entering the date_type_name production.
	EnterDate_type_name(c *Date_type_nameContext)

	// EnterBit_string_type_name is called when entering the bit_string_type_name production.
	EnterBit_string_type_name(c *Bit_string_type_nameContext)

	// EnterGeneric_type_name is called when entering the generic_type_name production.
	EnterGeneric_type_name(c *Generic_type_nameContext)

	// EnterData_type_declaration is called when entering the data_type_declaration production.
	EnterData_type_declaration(c *Data_type_declarationContext)

	// EnterType_declaration is called when entering the type_declaration production.
	EnterType_declaration(c *Type_declarationContext)

	// EnterInitializations_constant is called when entering the initializations_constant production.
	EnterInitializations_constant(c *Initializations_constantContext)

	// EnterInitializations_identifier is called when entering the initializations_identifier production.
	EnterInitializations_identifier(c *Initializations_identifierContext)

	// EnterInitializations_array_initialization is called when entering the initializations_array_initialization production.
	EnterInitializations_array_initialization(c *Initializations_array_initializationContext)

	// EnterInitializations_structure_initialization is called when entering the initializations_structure_initialization production.
	EnterInitializations_structure_initialization(c *Initializations_structure_initializationContext)

	// EnterSubrange_spec_init is called when entering the subrange_spec_init production.
	EnterSubrange_spec_init(c *Subrange_spec_initContext)

	// EnterSubrange is called when entering the subrange production.
	EnterSubrange(c *SubrangeContext)

	// EnterEnumerated_specification is called when entering the enumerated_specification production.
	EnterEnumerated_specification(c *Enumerated_specificationContext)

	// EnterArray_specification is called when entering the array_specification production.
	EnterArray_specification(c *Array_specificationContext)

	// EnterArray_initialization is called when entering the array_initialization production.
	EnterArray_initialization(c *Array_initializationContext)

	// EnterArray_initial_elements is called when entering the array_initial_elements production.
	EnterArray_initial_elements(c *Array_initial_elementsContext)

	// EnterArray_initial_element is called when entering the array_initial_element production.
	EnterArray_initial_element(c *Array_initial_elementContext)

	// EnterStructure_declaration is called when entering the structure_declaration production.
	EnterStructure_declaration(c *Structure_declarationContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterStructure_initialization is called when entering the structure_initialization production.
	EnterStructure_initialization(c *Structure_initializationContext)

	// EnterString_type_declaration is called when entering the string_type_declaration production.
	EnterString_type_declaration(c *String_type_declarationContext)

	// EnterReference_specification is called when entering the reference_specification production.
	EnterReference_specification(c *Reference_specificationContext)

	// EnterReference_value is called when entering the reference_value production.
	EnterReference_value(c *Reference_valueContext)

	// EnterIdentifier_list is called when entering the identifier_list production.
	EnterIdentifier_list(c *Identifier_listContext)

	// EnterFunction_declaration is called when entering the function_declaration production.
	EnterFunction_declaration(c *Function_declarationContext)

	// EnterVar_decls is called when entering the var_decls production.
	EnterVar_decls(c *Var_declsContext)

	// EnterVar_decl is called when entering the var_decl production.
	EnterVar_decl(c *Var_declContext)

	// EnterVar_decl_inner is called when entering the var_decl_inner production.
	EnterVar_decl_inner(c *Var_decl_innerContext)

	// EnterVariable_keyword is called when entering the variable_keyword production.
	EnterVariable_keyword(c *Variable_keywordContext)

	// EnterAccess_specifier is called when entering the access_specifier production.
	EnterAccess_specifier(c *Access_specifierContext)

	// EnterFunction_block_declaration is called when entering the function_block_declaration production.
	EnterFunction_block_declaration(c *Function_block_declarationContext)

	// EnterBody is called when entering the body production.
	EnterBody(c *BodyContext)

	// EnterFuncBody is called when entering the funcBody production.
	EnterFuncBody(c *FuncBodyContext)

	// EnterInterface_declaration is called when entering the interface_declaration production.
	EnterInterface_declaration(c *Interface_declarationContext)

	// EnterClass_declaration is called when entering the class_declaration production.
	EnterClass_declaration(c *Class_declarationContext)

	// EnterMethods is called when entering the methods production.
	EnterMethods(c *MethodsContext)

	// EnterMethod is called when entering the method production.
	EnterMethod(c *MethodContext)

	// EnterProgram_declaration is called when entering the program_declaration production.
	EnterProgram_declaration(c *Program_declarationContext)

	// EnterGlobal_variable_list_declaration is called when entering the global_variable_list_declaration production.
	EnterGlobal_variable_list_declaration(c *Global_variable_list_declarationContext)

	// EnterStl_list is called when entering the stl_list production.
	EnterStl_list(c *Stl_listContext)

	// EnterStl_expression is called when entering the stl_expression production.
	EnterStl_expression(c *Stl_expressionContext)

	// EnterStl_call is called when entering the stl_call production.
	EnterStl_call(c *Stl_callContext)

	// EnterUnaryNegateExpr is called when entering the unaryNegateExpr production.
	EnterUnaryNegateExpr(c *UnaryNegateExprContext)

	// EnterBinaryOrExpr is called when entering the binaryOrExpr production.
	EnterBinaryOrExpr(c *BinaryOrExprContext)

	// EnterBinaryCmpExpr is called when entering the binaryCmpExpr production.
	EnterBinaryCmpExpr(c *BinaryCmpExprContext)

	// EnterBinaryModDivExpr is called when entering the binaryModDivExpr production.
	EnterBinaryModDivExpr(c *BinaryModDivExprContext)

	// EnterParenExpr is called when entering the parenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterBinaryXORExpr is called when entering the binaryXORExpr production.
	EnterBinaryXORExpr(c *BinaryXORExprContext)

	// EnterUnaryMinusExpr is called when entering the unaryMinusExpr production.
	EnterUnaryMinusExpr(c *UnaryMinusExprContext)

	// EnterPrimaryExpr is called when entering the primaryExpr production.
	EnterPrimaryExpr(c *PrimaryExprContext)

	// EnterBinaryPowerExpr is called when entering the binaryPowerExpr production.
	EnterBinaryPowerExpr(c *BinaryPowerExprContext)

	// EnterBinaryMultExpr is called when entering the binaryMultExpr production.
	EnterBinaryMultExpr(c *BinaryMultExprContext)

	// EnterBinaryPlusMinusExpr is called when entering the binaryPlusMinusExpr production.
	EnterBinaryPlusMinusExpr(c *BinaryPlusMinusExprContext)

	// EnterBinaryEqExpr is called when entering the binaryEqExpr production.
	EnterBinaryEqExpr(c *BinaryEqExprContext)

	// EnterBinaryAndExpr is called when entering the binaryAndExpr production.
	EnterBinaryAndExpr(c *BinaryAndExprContext)

	// EnterPrimary_expression is called when entering the primary_expression production.
	EnterPrimary_expression(c *Primary_expressionContext)

	// EnterInvocation is called when entering the invocation production.
	EnterInvocation(c *InvocationContext)

	// EnterStatement_list is called when entering the statement_list production.
	EnterStatement_list(c *Statement_listContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterEmpty_statement is called when entering the empty_statement production.
	EnterEmpty_statement(c *Empty_statementContext)

	// EnterJump_statement is called when entering the jump_statement production.
	EnterJump_statement(c *Jump_statementContext)

	// EnterLabel_statement is called when entering the label_statement production.
	EnterLabel_statement(c *Label_statementContext)

	// EnterAssignment_statement is called when entering the assignment_statement production.
	EnterAssignment_statement(c *Assignment_statementContext)

	// EnterMult_assignment_statement is called when entering the mult_assignment_statement production.
	EnterMult_assignment_statement(c *Mult_assignment_statementContext)

	// EnterInvocation_statement is called when entering the invocation_statement production.
	EnterInvocation_statement(c *Invocation_statementContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterVariable_names is called when entering the variable_names production.
	EnterVariable_names(c *Variable_namesContext)

	// EnterSymbolic_variable is called when entering the symbolic_variable production.
	EnterSymbolic_variable(c *Symbolic_variableContext)

	// EnterSubscript_list is called when entering the subscript_list production.
	EnterSubscript_list(c *Subscript_listContext)

	// EnterDirect_variable is called when entering the direct_variable production.
	EnterDirect_variable(c *Direct_variableContext)

	// EnterReturn_statement is called when entering the return_statement production.
	EnterReturn_statement(c *Return_statementContext)

	// EnterParam_assignment is called when entering the param_assignment production.
	EnterParam_assignment(c *Param_assignmentContext)

	// EnterIf_statement is called when entering the if_statement production.
	EnterIf_statement(c *If_statementContext)

	// EnterCase_statement is called when entering the case_statement production.
	EnterCase_statement(c *Case_statementContext)

	// EnterCase_entry is called when entering the case_entry production.
	EnterCase_entry(c *Case_entryContext)

	// EnterCase_condition is called when entering the case_condition production.
	EnterCase_condition(c *Case_conditionContext)

	// EnterFor_statement is called when entering the for_statement production.
	EnterFor_statement(c *For_statementContext)

	// EnterWhile_statement is called when entering the while_statement production.
	EnterWhile_statement(c *While_statementContext)

	// EnterRepeat_statement is called when entering the repeat_statement production.
	EnterRepeat_statement(c *Repeat_statementContext)

	// EnterExit_statement is called when entering the exit_statement production.
	EnterExit_statement(c *Exit_statementContext)

	// EnterContinue_statement is called when entering the continue_statement production.
	EnterContinue_statement(c *Continue_statementContext)

	// EnterSfc is called when entering the sfc production.
	EnterSfc(c *SfcContext)

	// EnterSfc_network is called when entering the sfc_network production.
	EnterSfc_network(c *Sfc_networkContext)

	// EnterInit_step is called when entering the init_step production.
	EnterInit_step(c *Init_stepContext)

	// EnterStep is called when entering the step production.
	EnterStep(c *StepContext)

	// EnterAction_association is called when entering the action_association production.
	EnterAction_association(c *Action_associationContext)

	// EnterActionQualifier is called when entering the actionQualifier production.
	EnterActionQualifier(c *ActionQualifierContext)

	// EnterTransition is called when entering the transition production.
	EnterTransition(c *TransitionContext)

	// EnterSteps is called when entering the steps production.
	EnterSteps(c *StepsContext)

	// EnterTransitionCond is called when entering the transitionCond production.
	EnterTransitionCond(c *TransitionCondContext)

	// EnterAction is called when entering the action production.
	EnterAction(c *ActionContext)

	// EnterIlBody is called when entering the ilBody production.
	EnterIlBody(c *IlBodyContext)

	// EnterIlInstruction is called when entering the ilInstruction production.
	EnterIlInstruction(c *IlInstructionContext)

	// EnterIlSInstr is called when entering the ilSInstr production.
	EnterIlSInstr(c *IlSInstrContext)

	// EnterIlInstr is called when entering the ilInstr production.
	EnterIlInstr(c *IlInstrContext)

	// EnterIlSInstrList is called when entering the ilSInstrList production.
	EnterIlSInstrList(c *IlSInstrListContext)

	// EnterIlSimple is called when entering the ilSimple production.
	EnterIlSimple(c *IlSimpleContext)

	// EnterIlExpr is called when entering the ilExpr production.
	EnterIlExpr(c *IlExprContext)

	// EnterIlFunctionCall is called when entering the ilFunctionCall production.
	EnterIlFunctionCall(c *IlFunctionCallContext)

	// EnterIlFormalFunctionCall is called when entering the ilFormalFunctionCall production.
	EnterIlFormalFunctionCall(c *IlFormalFunctionCallContext)

	// EnterIlJump is called when entering the ilJump production.
	EnterIlJump(c *IlJumpContext)

	// EnterIlCall is called when entering the ilCall production.
	EnterIlCall(c *IlCallContext)

	// EnterIlOperand is called when entering the ilOperand production.
	EnterIlOperand(c *IlOperandContext)

	// EnterJump_op is called when entering the jump_op production.
	EnterJump_op(c *Jump_opContext)

	// EnterCall_op is called when entering the call_op production.
	EnterCall_op(c *Call_opContext)

	// EnterSimple_op is called when entering the simple_op production.
	EnterSimple_op(c *Simple_opContext)

	// EnterExprOperator is called when entering the exprOperator production.
	EnterExprOperator(c *ExprOperatorContext)

	// EnterIl_param_assignment is called when entering the il_param_assignment production.
	EnterIl_param_assignment(c *Il_param_assignmentContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitNamespace_declaration is called when exiting the namespace_declaration production.
	ExitNamespace_declaration(c *Namespace_declarationContext)

	// ExitNamespace_elements is called when exiting the namespace_elements production.
	ExitNamespace_elements(c *Namespace_elementsContext)

	// ExitFull_qualified_identifier is called when exiting the full_qualified_identifier production.
	ExitFull_qualified_identifier(c *Full_qualified_identifierContext)

	// ExitUsing_directive is called when exiting the using_directive production.
	ExitUsing_directive(c *Using_directiveContext)

	// ExitLibrary_element_declaration is called when exiting the library_element_declaration production.
	ExitLibrary_element_declaration(c *Library_element_declarationContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitCast is called when exiting the cast production.
	ExitCast(c *CastContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)

	// ExitBits is called when exiting the bits production.
	ExitBits(c *BitsContext)

	// ExitReal is called when exiting the real production.
	ExitReal(c *RealContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitTime is called when exiting the time production.
	ExitTime(c *TimeContext)

	// ExitTimeofday is called when exiting the timeofday production.
	ExitTimeofday(c *TimeofdayContext)

	// ExitDate is called when exiting the date production.
	ExitDate(c *DateContext)

	// ExitDatetime is called when exiting the datetime production.
	ExitDatetime(c *DatetimeContext)

	// ExitRef_null is called when exiting the ref_null production.
	ExitRef_null(c *Ref_nullContext)

	// ExitData_type_name is called when exiting the data_type_name production.
	ExitData_type_name(c *Data_type_nameContext)

	// ExitNon_generic_type_name is called when exiting the non_generic_type_name production.
	ExitNon_generic_type_name(c *Non_generic_type_nameContext)

	// ExitElementary_type_name is called when exiting the elementary_type_name production.
	ExitElementary_type_name(c *Elementary_type_nameContext)

	// ExitNumeric_type_name is called when exiting the numeric_type_name production.
	ExitNumeric_type_name(c *Numeric_type_nameContext)

	// ExitInteger_type_name is called when exiting the integer_type_name production.
	ExitInteger_type_name(c *Integer_type_nameContext)

	// ExitSigned_integer_type_name is called when exiting the signed_integer_type_name production.
	ExitSigned_integer_type_name(c *Signed_integer_type_nameContext)

	// ExitUnsigned_integer_type_name is called when exiting the unsigned_integer_type_name production.
	ExitUnsigned_integer_type_name(c *Unsigned_integer_type_nameContext)

	// ExitReal_type_name is called when exiting the real_type_name production.
	ExitReal_type_name(c *Real_type_nameContext)

	// ExitDate_type_name is called when exiting the date_type_name production.
	ExitDate_type_name(c *Date_type_nameContext)

	// ExitBit_string_type_name is called when exiting the bit_string_type_name production.
	ExitBit_string_type_name(c *Bit_string_type_nameContext)

	// ExitGeneric_type_name is called when exiting the generic_type_name production.
	ExitGeneric_type_name(c *Generic_type_nameContext)

	// ExitData_type_declaration is called when exiting the data_type_declaration production.
	ExitData_type_declaration(c *Data_type_declarationContext)

	// ExitType_declaration is called when exiting the type_declaration production.
	ExitType_declaration(c *Type_declarationContext)

	// ExitInitializations_constant is called when exiting the initializations_constant production.
	ExitInitializations_constant(c *Initializations_constantContext)

	// ExitInitializations_identifier is called when exiting the initializations_identifier production.
	ExitInitializations_identifier(c *Initializations_identifierContext)

	// ExitInitializations_array_initialization is called when exiting the initializations_array_initialization production.
	ExitInitializations_array_initialization(c *Initializations_array_initializationContext)

	// ExitInitializations_structure_initialization is called when exiting the initializations_structure_initialization production.
	ExitInitializations_structure_initialization(c *Initializations_structure_initializationContext)

	// ExitSubrange_spec_init is called when exiting the subrange_spec_init production.
	ExitSubrange_spec_init(c *Subrange_spec_initContext)

	// ExitSubrange is called when exiting the subrange production.
	ExitSubrange(c *SubrangeContext)

	// ExitEnumerated_specification is called when exiting the enumerated_specification production.
	ExitEnumerated_specification(c *Enumerated_specificationContext)

	// ExitArray_specification is called when exiting the array_specification production.
	ExitArray_specification(c *Array_specificationContext)

	// ExitArray_initialization is called when exiting the array_initialization production.
	ExitArray_initialization(c *Array_initializationContext)

	// ExitArray_initial_elements is called when exiting the array_initial_elements production.
	ExitArray_initial_elements(c *Array_initial_elementsContext)

	// ExitArray_initial_element is called when exiting the array_initial_element production.
	ExitArray_initial_element(c *Array_initial_elementContext)

	// ExitStructure_declaration is called when exiting the structure_declaration production.
	ExitStructure_declaration(c *Structure_declarationContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitStructure_initialization is called when exiting the structure_initialization production.
	ExitStructure_initialization(c *Structure_initializationContext)

	// ExitString_type_declaration is called when exiting the string_type_declaration production.
	ExitString_type_declaration(c *String_type_declarationContext)

	// ExitReference_specification is called when exiting the reference_specification production.
	ExitReference_specification(c *Reference_specificationContext)

	// ExitReference_value is called when exiting the reference_value production.
	ExitReference_value(c *Reference_valueContext)

	// ExitIdentifier_list is called when exiting the identifier_list production.
	ExitIdentifier_list(c *Identifier_listContext)

	// ExitFunction_declaration is called when exiting the function_declaration production.
	ExitFunction_declaration(c *Function_declarationContext)

	// ExitVar_decls is called when exiting the var_decls production.
	ExitVar_decls(c *Var_declsContext)

	// ExitVar_decl is called when exiting the var_decl production.
	ExitVar_decl(c *Var_declContext)

	// ExitVar_decl_inner is called when exiting the var_decl_inner production.
	ExitVar_decl_inner(c *Var_decl_innerContext)

	// ExitVariable_keyword is called when exiting the variable_keyword production.
	ExitVariable_keyword(c *Variable_keywordContext)

	// ExitAccess_specifier is called when exiting the access_specifier production.
	ExitAccess_specifier(c *Access_specifierContext)

	// ExitFunction_block_declaration is called when exiting the function_block_declaration production.
	ExitFunction_block_declaration(c *Function_block_declarationContext)

	// ExitBody is called when exiting the body production.
	ExitBody(c *BodyContext)

	// ExitFuncBody is called when exiting the funcBody production.
	ExitFuncBody(c *FuncBodyContext)

	// ExitInterface_declaration is called when exiting the interface_declaration production.
	ExitInterface_declaration(c *Interface_declarationContext)

	// ExitClass_declaration is called when exiting the class_declaration production.
	ExitClass_declaration(c *Class_declarationContext)

	// ExitMethods is called when exiting the methods production.
	ExitMethods(c *MethodsContext)

	// ExitMethod is called when exiting the method production.
	ExitMethod(c *MethodContext)

	// ExitProgram_declaration is called when exiting the program_declaration production.
	ExitProgram_declaration(c *Program_declarationContext)

	// ExitGlobal_variable_list_declaration is called when exiting the global_variable_list_declaration production.
	ExitGlobal_variable_list_declaration(c *Global_variable_list_declarationContext)

	// ExitStl_list is called when exiting the stl_list production.
	ExitStl_list(c *Stl_listContext)

	// ExitStl_expression is called when exiting the stl_expression production.
	ExitStl_expression(c *Stl_expressionContext)

	// ExitStl_call is called when exiting the stl_call production.
	ExitStl_call(c *Stl_callContext)

	// ExitUnaryNegateExpr is called when exiting the unaryNegateExpr production.
	ExitUnaryNegateExpr(c *UnaryNegateExprContext)

	// ExitBinaryOrExpr is called when exiting the binaryOrExpr production.
	ExitBinaryOrExpr(c *BinaryOrExprContext)

	// ExitBinaryCmpExpr is called when exiting the binaryCmpExpr production.
	ExitBinaryCmpExpr(c *BinaryCmpExprContext)

	// ExitBinaryModDivExpr is called when exiting the binaryModDivExpr production.
	ExitBinaryModDivExpr(c *BinaryModDivExprContext)

	// ExitParenExpr is called when exiting the parenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitBinaryXORExpr is called when exiting the binaryXORExpr production.
	ExitBinaryXORExpr(c *BinaryXORExprContext)

	// ExitUnaryMinusExpr is called when exiting the unaryMinusExpr production.
	ExitUnaryMinusExpr(c *UnaryMinusExprContext)

	// ExitPrimaryExpr is called when exiting the primaryExpr production.
	ExitPrimaryExpr(c *PrimaryExprContext)

	// ExitBinaryPowerExpr is called when exiting the binaryPowerExpr production.
	ExitBinaryPowerExpr(c *BinaryPowerExprContext)

	// ExitBinaryMultExpr is called when exiting the binaryMultExpr production.
	ExitBinaryMultExpr(c *BinaryMultExprContext)

	// ExitBinaryPlusMinusExpr is called when exiting the binaryPlusMinusExpr production.
	ExitBinaryPlusMinusExpr(c *BinaryPlusMinusExprContext)

	// ExitBinaryEqExpr is called when exiting the binaryEqExpr production.
	ExitBinaryEqExpr(c *BinaryEqExprContext)

	// ExitBinaryAndExpr is called when exiting the binaryAndExpr production.
	ExitBinaryAndExpr(c *BinaryAndExprContext)

	// ExitPrimary_expression is called when exiting the primary_expression production.
	ExitPrimary_expression(c *Primary_expressionContext)

	// ExitInvocation is called when exiting the invocation production.
	ExitInvocation(c *InvocationContext)

	// ExitStatement_list is called when exiting the statement_list production.
	ExitStatement_list(c *Statement_listContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitEmpty_statement is called when exiting the empty_statement production.
	ExitEmpty_statement(c *Empty_statementContext)

	// ExitJump_statement is called when exiting the jump_statement production.
	ExitJump_statement(c *Jump_statementContext)

	// ExitLabel_statement is called when exiting the label_statement production.
	ExitLabel_statement(c *Label_statementContext)

	// ExitAssignment_statement is called when exiting the assignment_statement production.
	ExitAssignment_statement(c *Assignment_statementContext)

	// ExitMult_assignment_statement is called when exiting the mult_assignment_statement production.
	ExitMult_assignment_statement(c *Mult_assignment_statementContext)

	// ExitInvocation_statement is called when exiting the invocation_statement production.
	ExitInvocation_statement(c *Invocation_statementContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitVariable_names is called when exiting the variable_names production.
	ExitVariable_names(c *Variable_namesContext)

	// ExitSymbolic_variable is called when exiting the symbolic_variable production.
	ExitSymbolic_variable(c *Symbolic_variableContext)

	// ExitSubscript_list is called when exiting the subscript_list production.
	ExitSubscript_list(c *Subscript_listContext)

	// ExitDirect_variable is called when exiting the direct_variable production.
	ExitDirect_variable(c *Direct_variableContext)

	// ExitReturn_statement is called when exiting the return_statement production.
	ExitReturn_statement(c *Return_statementContext)

	// ExitParam_assignment is called when exiting the param_assignment production.
	ExitParam_assignment(c *Param_assignmentContext)

	// ExitIf_statement is called when exiting the if_statement production.
	ExitIf_statement(c *If_statementContext)

	// ExitCase_statement is called when exiting the case_statement production.
	ExitCase_statement(c *Case_statementContext)

	// ExitCase_entry is called when exiting the case_entry production.
	ExitCase_entry(c *Case_entryContext)

	// ExitCase_condition is called when exiting the case_condition production.
	ExitCase_condition(c *Case_conditionContext)

	// ExitFor_statement is called when exiting the for_statement production.
	ExitFor_statement(c *For_statementContext)

	// ExitWhile_statement is called when exiting the while_statement production.
	ExitWhile_statement(c *While_statementContext)

	// ExitRepeat_statement is called when exiting the repeat_statement production.
	ExitRepeat_statement(c *Repeat_statementContext)

	// ExitExit_statement is called when exiting the exit_statement production.
	ExitExit_statement(c *Exit_statementContext)

	// ExitContinue_statement is called when exiting the continue_statement production.
	ExitContinue_statement(c *Continue_statementContext)

	// ExitSfc is called when exiting the sfc production.
	ExitSfc(c *SfcContext)

	// ExitSfc_network is called when exiting the sfc_network production.
	ExitSfc_network(c *Sfc_networkContext)

	// ExitInit_step is called when exiting the init_step production.
	ExitInit_step(c *Init_stepContext)

	// ExitStep is called when exiting the step production.
	ExitStep(c *StepContext)

	// ExitAction_association is called when exiting the action_association production.
	ExitAction_association(c *Action_associationContext)

	// ExitActionQualifier is called when exiting the actionQualifier production.
	ExitActionQualifier(c *ActionQualifierContext)

	// ExitTransition is called when exiting the transition production.
	ExitTransition(c *TransitionContext)

	// ExitSteps is called when exiting the steps production.
	ExitSteps(c *StepsContext)

	// ExitTransitionCond is called when exiting the transitionCond production.
	ExitTransitionCond(c *TransitionCondContext)

	// ExitAction is called when exiting the action production.
	ExitAction(c *ActionContext)

	// ExitIlBody is called when exiting the ilBody production.
	ExitIlBody(c *IlBodyContext)

	// ExitIlInstruction is called when exiting the ilInstruction production.
	ExitIlInstruction(c *IlInstructionContext)

	// ExitIlSInstr is called when exiting the ilSInstr production.
	ExitIlSInstr(c *IlSInstrContext)

	// ExitIlInstr is called when exiting the ilInstr production.
	ExitIlInstr(c *IlInstrContext)

	// ExitIlSInstrList is called when exiting the ilSInstrList production.
	ExitIlSInstrList(c *IlSInstrListContext)

	// ExitIlSimple is called when exiting the ilSimple production.
	ExitIlSimple(c *IlSimpleContext)

	// ExitIlExpr is called when exiting the ilExpr production.
	ExitIlExpr(c *IlExprContext)

	// ExitIlFunctionCall is called when exiting the ilFunctionCall production.
	ExitIlFunctionCall(c *IlFunctionCallContext)

	// ExitIlFormalFunctionCall is called when exiting the ilFormalFunctionCall production.
	ExitIlFormalFunctionCall(c *IlFormalFunctionCallContext)

	// ExitIlJump is called when exiting the ilJump production.
	ExitIlJump(c *IlJumpContext)

	// ExitIlCall is called when exiting the ilCall production.
	ExitIlCall(c *IlCallContext)

	// ExitIlOperand is called when exiting the ilOperand production.
	ExitIlOperand(c *IlOperandContext)

	// ExitJump_op is called when exiting the jump_op production.
	ExitJump_op(c *Jump_opContext)

	// ExitCall_op is called when exiting the call_op production.
	ExitCall_op(c *Call_opContext)

	// ExitSimple_op is called when exiting the simple_op production.
	ExitSimple_op(c *Simple_opContext)

	// ExitExprOperator is called when exiting the exprOperator production.
	ExitExprOperator(c *ExprOperatorContext)

	// ExitIl_param_assignment is called when exiting the il_param_assignment production.
	ExitIl_param_assignment(c *Il_param_assignmentContext)
}
