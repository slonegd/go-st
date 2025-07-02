package ast

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/slonegd/go-st/antlr"
	"github.com/slonegd/go-st/ops"
	"github.com/slonegd/go-st/types"
)

type (
	// обёртка, чтоб скрыть методы parser.STVisitor
	// ресиверы не по указателю специально, чтоб пробрасывать семантику вглубь дерева не создавая стек
	visitor struct {
		*AST
		// в некоторые ноды надо зайти, чтоб определить ошибку парсинга, но ничего не делать
		// такое бывает когда по всем детям проходишь, а там в конце, например TerminalNode
		lastOp any
		cycle  cycle
		*ops.Placeholders
	}
	cycle struct {
		condition *ops.Statement
		out       *ops.Statement
	}
)

// Visit implements parser.STVisitor.
func (x visitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(x)
}

// VisitArray_conform_decl implements parser.STVisitor.
func (x visitor) VisitArray_conform_decl(ctx *parser.Array_conform_declContext) any {
	panic("unimplemented")
}

// VisitArray_conformand implements parser.STVisitor.
func (x visitor) VisitArray_conformand(ctx *parser.Array_conformandContext) any {
	panic("unimplemented")
}

// VisitArray_type implements parser.STVisitor.
func (x visitor) VisitArray_type(ctx *parser.Array_typeContext) any {
	panic("unimplemented")
}

// VisitArray_type_access implements parser.STVisitor.
func (x visitor) VisitArray_type_access(ctx *parser.Array_type_accessContext) any {
	panic("unimplemented")
}

// VisitArray_type_name implements parser.STVisitor.
func (x visitor) VisitArray_type_name(ctx *parser.Array_type_nameContext) any {
	panic("unimplemented")
}

// VisitAssignment_statement implements parser.STVisitor.
func (x visitor) VisitAssignment_statement(ctx *parser.Assignment_statementContext) any {
	varName := ctx.Variable().GetText() // TODO там правило немного сложнее
	v := x.vars[varName]

	expr := ctx.Expression().Accept(x)
	switch expr := expr.(type) {
	case ops.Int64:
		if !expr.IsConstant {
			x.CheckError(ctx, types.CheckAssign(v.Type(), expr.ResultType))
		}
		return ops.Assign(ctx.CustomContext, varName, v, expr)
	case ops.Float64:
		if !expr.IsConstant {
			x.CheckError(ctx, types.CheckAssign(v.Type(), expr.ResultType))
		}
		return ops.Assign(ctx.CustomContext, varName, v, expr)
	}
	return x.CheckError(ctx, fmt.Errorf("undefined operator in assign statement: %+v", expr))
}

// VisitBinaryExpression implements parser.STVisitor.
func (x visitor) VisitBinaryExpression(ctx *parser.BinaryExpressionContext) any {
	right := ctx.GetRight().Accept(x)
	left := ctx.GetLeft().Accept(x)
	token := ctx.GetOperator()
	op := Ops[token.GetTokenType()]
	switch op {
	case ops.Plus, ops.Minus, ops.Mult, ops.Div, ops.Mod:
		switch right := right.(type) {
		case ops.Int64:
			switch left := left.(type) {
			case ops.Int64:
				resultType, err := types.BinaryResult(
					types.Expression{left.IsConstant, left.ResultType},
					types.Expression{right.IsConstant, right.ResultType},
				)
				x.CheckError(ctx, err)
				return ops.Arithmetic[int64](ctx.GetCustomContext(), op, left, right, resultType)
			case ops.Float64:
				resultType, err := types.BinaryResult(
					types.Expression{left.IsConstant, left.ResultType},
					types.Expression{right.IsConstant, right.ResultType},
				)
				x.CheckError(ctx, err)
				return ops.Arithmetic[float64](ctx.GetCustomContext(), op, left, right, resultType)
			}
		case ops.Float64:
			switch left := left.(type) {
			case ops.Int64:
				resultType, err := types.BinaryResult(
					types.Expression{left.IsConstant, left.ResultType},
					types.Expression{right.IsConstant, right.ResultType},
				)
				x.CheckError(ctx, err)
				return ops.Arithmetic[float64](ctx.GetCustomContext(), op, left, right, resultType)
			case ops.Float64:
				resultType, err := types.BinaryResult(
					types.Expression{left.IsConstant, left.ResultType},
					types.Expression{right.IsConstant, right.ResultType},
				)
				x.CheckError(ctx, err)
				return ops.Arithmetic[float64](ctx.GetCustomContext(), op, left, right, resultType)
			}
		}
	}
	return x.CheckError(ctx, fmt.Errorf("undefined operator %v (token type %d)", token, token.GetTokenType()))
}

// VisitCase_element implements parser.STVisitor.
func (x visitor) VisitCase_element(ctx *parser.Case_elementContext) any {
	panic("unimplemented")
}

// VisitCase_label implements parser.STVisitor.
func (x visitor) VisitCase_label(ctx *parser.Case_labelContext) any {
	panic("unimplemented")
}

// VisitCase_statement implements parser.STVisitor.
func (x visitor) VisitCase_statement(ctx *parser.Case_statementContext) any {
	panic("unimplemented")
}

// VisitChildren implements parser.STVisitor.
func (x visitor) VisitChildren(node antlr.RuleNode) any {
	for _, c := range node.GetChildren() {
		x.lastOp = x.Visit(c.(antlr.ParseTree))
	}
	return x.lastOp
}

// VisitContinue_statement implements parser.STVisitor.
func (x visitor) VisitContinue_statement(ctx *parser.Continue_statementContext) any {
	panic("unimplemented")
}

// VisitData_type implements parser.STVisitor.
func (x visitor) VisitData_type(ctx *parser.Data_typeContext) any {
	return ctx.Elementary_type_name().Accept(x)
}

// VisitData_type_access implements parser.STVisitor.
func (x visitor) VisitData_type_access(ctx *parser.Data_type_accessContext) any {
	panic("unimplemented")
}

// VisitDerived_type_access implements parser.STVisitor.
func (x visitor) VisitDerived_type_access(ctx *parser.Derived_type_accessContext) any {
	panic("unimplemented")
}

// VisitEdge_decl implements parser.STVisitor.
func (x visitor) VisitEdge_decl(ctx *parser.Edge_declContext) any {
	panic("unimplemented")
}

// VisitElementary_type_name implements parser.STVisitor.
func (x visitor) VisitElementary_type_name(ctx *parser.Elementary_type_nameContext) any {
	return ctx.GetText()
}

// VisitEnum_type_access implements parser.STVisitor.
func (x visitor) VisitEnum_type_access(ctx *parser.Enum_type_accessContext) any {
	panic("unimplemented")
}

// VisitEnum_type_name implements parser.STVisitor.
func (x visitor) VisitEnum_type_name(ctx *parser.Enum_type_nameContext) any {
	panic("unimplemented")
}

// VisitErrorNode implements parser.STVisitor.
func (x visitor) VisitErrorNode(node antlr.ErrorNode) any {
	// через antlr.ErrorListener ошибка информативнее
	// x.err = fmt.Errorf("error node %+v", node)
	return x.lastOp
}

// VisitExit_statement implements parser.STVisitor.
func (x visitor) VisitExit_statement(ctx *parser.Exit_statementContext) any {
	panic("unimplemented")
}

// VisitFor_statement implements parser.STVisitor.
func (x visitor) VisitFor_statement(ctx *parser.For_statementContext) any {
	panic("unimplemented")
}

// VisitFuncCallExpression implements parser.STVisitor.
func (x visitor) VisitFuncCallExpression(ctx *parser.FuncCallExpressionContext) any {
	panic("unimplemented")
}

// VisitFunction_block_decl implements parser.STVisitor.
func (x visitor) VisitFunction_block_decl(ctx *parser.Function_block_declContext) any {
	panic("unimplemented")
}

// VisitFunction_decl implements parser.STVisitor.
func (x visitor) VisitFunction_decl(ctx *parser.Function_declContext) any {
	panic("unimplemented")
}

// VisitFunction_invocation implements parser.STVisitor.
func (x visitor) VisitFunction_invocation(ctx *parser.Function_invocationContext) any {
	panic("unimplemented")
}

// VisitIf_statement implements parser.STVisitor.
func (x visitor) VisitIf_statement(ctx *parser.If_statementContext) any {
	panic("unimplemented")
}

// VisitInput_decl implements parser.STVisitor.
func (x visitor) VisitInput_decl(ctx *parser.Input_declContext) any {
	panic("unimplemented")
}

// VisitLiteral implements parser.STVisitor.
func (x visitor) VisitLiteral(ctx *parser.LiteralContext) any {
	i, err := strconv.Atoi(ctx.INT_LITERAL().GetText())
	x.CheckError(ctx, err)
	return ops.Constant(int64(i))
}

// VisitLiteralExpression implements parser.STVisitor.
func (x visitor) VisitLiteralExpression(ctx *parser.LiteralExpressionContext) any {
	return ctx.Literal().Accept(x)
}

// VisitNamespace_name implements parser.STVisitor.
func (x visitor) VisitNamespace_name(ctx *parser.Namespace_nameContext) any {
	panic("unimplemented")
}

// VisitOutput_decl implements parser.STVisitor.
func (x visitor) VisitOutput_decl(ctx *parser.Output_declContext) any {
	panic("unimplemented")
}

// VisitParenExpression implements parser.STVisitor.
func (x visitor) VisitParenExpression(ctx *parser.ParenExpressionContext) any {
	return ctx.Expression().Accept(x)
}

// VisitProgram implements parser.STVisitor.
func (x visitor) VisitProgram(ctx *parser.ProgramContext) any {
	vars := ctx.GetVars()
	if vars != nil {
		vars.Accept(x)
	}
	return ctx.GetStmts().Accept(x).(*ops.Statement)
}

// VisitRepeat_statement implements parser.STVisitor.
func (x visitor) VisitRepeat_statement(ctx *parser.Repeat_statementContext) any {
	panic("unimplemented")
}

// VisitReturn_statement implements parser.STVisitor.
func (x visitor) VisitReturn_statement(ctx *parser.Return_statementContext) any {
	panic("unimplemented")
}

// VisitSimple_type_access implements parser.STVisitor.
func (x visitor) VisitSimple_type_access(ctx *parser.Simple_type_accessContext) any {
	panic("unimplemented")
}

// VisitSimple_type_name implements parser.STVisitor.
func (x visitor) VisitSimple_type_name(ctx *parser.Simple_type_nameContext) any {
	panic("unimplemented")
}

// VisitSingle_elem_type_access implements parser.STVisitor.
func (x visitor) VisitSingle_elem_type_access(ctx *parser.Single_elem_type_accessContext) any {
	panic("unimplemented")
}

// VisitStatement implements parser.STVisitor.
func (x visitor) VisitStatement(ctx *parser.StatementContext) any {
	return x.VisitChildren(ctx)
}

// VisitStatement_list implements parser.STVisitor.
func (x visitor) VisitStatement_list(ctx *parser.Statement_listContext) any {
	var r, next *ops.Statement
	for _, s := range ctx.AllStatement() {
		tmp := s.Accept(x).(*ops.Statement)
		if r == nil {
			r = tmp
			next = tmp
			continue
		}
		next = x.SetNextStatement(next, tmp)
	}
	return r
}

// VisitString_type_access implements parser.STVisitor.
func (x visitor) VisitString_type_access(ctx *parser.String_type_accessContext) any {
	panic("unimplemented")
}

// VisitStruct_type_access implements parser.STVisitor.
func (x visitor) VisitStruct_type_access(ctx *parser.Struct_type_accessContext) any {
	panic("unimplemented")
}

// VisitStruct_type_name implements parser.STVisitor.
func (x visitor) VisitStruct_type_name(ctx *parser.Struct_type_nameContext) any {
	panic("unimplemented")
}

// VisitStructured_type implements parser.STVisitor.
func (x visitor) VisitStructured_type(ctx *parser.Structured_typeContext) any {
	panic("unimplemented")
}

// VisitSubrange implements parser.STVisitor.
func (x visitor) VisitSubrange(ctx *parser.SubrangeContext) any {
	panic("unimplemented")
}

// VisitSubrange_type_access implements parser.STVisitor.
func (x visitor) VisitSubrange_type_access(ctx *parser.Subrange_type_accessContext) any {
	panic("unimplemented")
}

// VisitSubrange_type_name implements parser.STVisitor.
func (x visitor) VisitSubrange_type_name(ctx *parser.Subrange_type_nameContext) any {
	panic("unimplemented")
}

// VisitTerminal implements parser.STVisitor.
func (x visitor) VisitTerminal(node antlr.TerminalNode) any { return x.lastOp }

// VisitType_name implements parser.STVisitor.
func (x visitor) VisitType_name(ctx *parser.Type_nameContext) any {
	panic("unimplemented")
}

// VisitTyped_literal implements parser.STVisitor.
func (x visitor) VisitTyped_literal(ctx *parser.Typed_literalContext) any {
	panic("unimplemented")
}

// VisitUnaryExpression implements parser.STVisitor.
func (x visitor) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) any {
	expr := ctx.Expression().Accept(x)
	token := ctx.GetOperator().GetTokenType()
	switch token {
	case parser.STLexerNOT, parser.STLexerPLUS:
		return x.CheckError(ctx, fmt.Errorf("unimplemented unary expression with operator %q (token: %d)", ctx.GetOperator().GetText(), token))
	case parser.STLexerMINUS:
		switch expr := expr.(type) {
		case ops.Int64:
			return ops.UnaryMinus(ctx.CustomContext, expr)
		case ops.Float64:
			return ops.UnaryMinus(ctx.CustomContext, expr)
		default:
			return x.CheckError(ctx, fmt.Errorf("unimplemented unary operator for type %T", expr))
		}
	}
	return expr
}

// VisitVarExpression implements parser.STVisitor.
func (x visitor) VisitVarExpression(ctx *parser.VarExpressionContext) any {
	return ctx.Variable().Accept(x)
}

// VisitVar_decl implements parser.STVisitor.
func (x visitor) VisitVar_decl(ctx *parser.Var_declContext) any {
	varName := ctx.GetId().GetText()
	var v types.Variable
	if ctx := ctx.GetExpr(); ctx != nil {
		switch op := ctx.Accept(x).(type) {
		case ops.Int64:
			tmp, _ := op.Do(nil)
			v = types.IntVariable(tmp)
		case ops.Float64:
			tmp, _ := op.Do(nil)
			v = types.Float64Variable(tmp)
		default:
			return x.CheckError(ctx, fmt.Errorf("unknown data type %T(%v)", op, op))
		}
	} else {
		v = types.NewAnyVariable("")
	}
	valueType := ctx.GetType_().Accept(x).(string)
	v = types.SetType(v, types.TypeFromString(valueType))
	x.vars[varName] = v
	return x.lastOp
}

// VisitVar_declaration_block implements parser.STVisitor.
func (x visitor) VisitVar_declaration_block(ctx *parser.Var_declaration_blockContext) any {
	for _, ctx := range ctx.GetDecls() {
		ctx.Accept(x)
	}
	return nil
}

// VisitVariable implements parser.STVisitor.
func (x visitor) VisitVariable(ctx *parser.VariableContext) any {
	varName := ctx.GetName().GetText() // TODO там немного сложнее со структурами/массивами
	v := x.vars[varName]
	if v.Type().IsInteger() { // TODO другие типы
		return ops.Variable[int64](ctx.CustomContext, varName, v)
	}
	if v.Type().IsFloat() {
		return ops.Variable[float64](ctx.CustomContext, varName, v)
	}
	return x.CheckError(ctx, fmt.Errorf("unknown type of variable: %s", v.Type()))
}

// VisitVariable_list implements parser.STVisitor.
func (x visitor) VisitVariable_list(ctx *parser.Variable_listContext) any {
	panic("unimplemented")
}

// VisitWhile_statement implements parser.STVisitor.
func (x visitor) VisitWhile_statement(ctx *parser.While_statementContext) any {
	panic("unimplemented")
}

var (
	_ parser.STVisitor    = visitor{}
	_ antlr.ErrorListener = (*visitor)(nil)
)
