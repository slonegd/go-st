package ast

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

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
		state string // для добавления имени функции/ФБ к имени их переменных
	}
	cycle struct {
		condition *ops.Statement
		out       *ops.Statement
	}
)

// VisitGlobal_var_declaration implements parser.STVisitor.
func (x visitor) VisitGlobal_var_declaration(ctx *parser.Global_var_declarationContext) any {
	panic("unimplemented")
}

// VisitPou implements parser.STVisitor.
func (x visitor) VisitPous(ctx *parser.PousContext) any {
	// TODO пока не рассматриваю вариант, когда в функции вызывается другая функция
	for _, p := range ctx.AllFunction_decl() {
		p.Accept(x)

	}
	for _, p := range ctx.AllProgram() {
		p.Accept(x)
	}
	return nil
}

// VisitType_declaration implements parser.STVisitor.
func (x visitor) VisitType_declaration(ctx *parser.Type_declarationContext) any {
	panic("unimplemented")
}

// VisitType_definition implements parser.STVisitor.
func (x visitor) VisitType_definition(ctx *parser.Type_definitionContext) any {
	panic("unimplemented")
}

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
	varName := x.state + ctx.Variable().GetText() // TODO там правило немного сложнее
	v, ok := x.vars[varName]
	if !ok {
		x.CheckError(ctx.Variable(), fmt.Errorf("no variable with name: %q", varName))

	}

	expr, err := ops.MakeExprNumber(ctx.Expression().Accept(x))
	x.CheckError(ctx.Expression(), err)
	if !expr.IsConstant() {
		x.CheckError(ctx.Expression(), types.CheckAssign(v.Type(), expr.Type()))
	}
	return ops.Assign(ctx.CustomContext, varName, v, expr)
}

// VisitBinaryExpression implements parser.STVisitor.
func (x visitor) VisitBinaryExpression(ctx *parser.BinaryExpressionContext) any {
	right := ctx.GetRight().Accept(x)
	left := ctx.GetLeft().Accept(x)
	r, err := ops.MakeExprNumber(right)
	x.CheckError(ctx.GetRight(), err)
	l, err := ops.MakeExprNumber(left)
	x.CheckError(ctx.GetLeft(), err)
	castType, err := types.BinaryResult(
		types.Expression{IsConstant: l.IsConstant(), Type: l.Type()},
		types.Expression{IsConstant: r.IsConstant(), Type: r.Type()},
	)
	token := ctx.GetOperator()
	op := Ops[token.GetTokenType()]
	x.CheckError(ctx, err)
	res := ops.Binary(ctx.GetCustomContext(), op, l, r, castType)
	if res == nil {
		return x.CheckError(ctx, fmt.Errorf("undefined operator %v (token type %d)", token, token.GetTokenType()))
	}
	return res
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
	if x.cycle.condition == nil {
		return x.CheckError(ctx, errors.New("no cycle body there"))
	}
	return ops.Jump(ctx.CustomContext, x.cycle.condition, "CONTINUE")
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
	if x.cycle.out == nil {
		return x.CheckError(ctx, errors.New("no cycle body there"))
	}
	return ops.Jump(ctx.CustomContext, x.cycle.out, "EXIT")
}

// VisitFor_statement implements parser.STVisitor.
func (x visitor) VisitFor_statement(ctx *parser.For_statementContext) any {
	result := ops.Placeholder()

	// стартовое состояние
	name := x.state + ctx.IDENTIFIER().GetText()
	v, ok := x.vars[name]
	if !ok {
		x.CheckError(ctx, fmt.Errorf("no variable with name: %q", name))
	}
	start, err := ops.MakeExprNumber(ctx.GetStart_().Accept(x))
	x.CheckError(ctx.GetStart_(), err)
	x.AddToStatementChain(result, ops.Assign(ctx.CustomContext, name, v, start))

	// конечное значение
	endV := types.SetType(types.IntVariable(0), v.Type())
	endName := "_end_" + name
	end, err := ops.MakeExprNumber(ctx.GetEnd().Accept(x))
	x.CheckError(ctx.GetEnd(), err)
	if !end.IsConstant() {
		x.CheckError(ctx.GetEnd(), types.CheckAssign(endV.Type(), end.Type()))
	}
	// TODO не тот контекст (проверить в дебаге)
	x.AddToStatementChain(result, ops.Assign(ctx.CustomContext, endName, endV, end))

	// шаг
	stepV := types.SetType(types.IntVariable(1), v.Type())
	stepName := "_step_" + name
	if stepCtx := ctx.GetStep(); stepCtx != nil {
		step, err := ops.MakeExprNumber(stepCtx.Accept(x))
		x.CheckError(stepCtx, err)
		if !step.IsConstant() {
			x.CheckError(stepCtx, types.CheckAssign(stepV.Type(), step.Type()))
		}
		x.AddToStatementChain(result, ops.Assign(ctx.CustomContext, stepName, stepV, step))
	} else {
		step := ops.Constant(types.IntVariable(1)).(ops.ExprInt64)
		x.AddToStatementChain(result, ops.Assign(ctx.CustomContext, stepName, stepV, step))
	}

	// само условие выхода зависит от того положительный шаг или нет
	plusCondition := ops.Binary(
		ctx.CustomContext,
		ops.Lte,
		ops.Variable(ctx.CustomContext, name, v).(ops.ExprNumber),
		ops.Variable(ctx.CustomContext, endName, endV).(ops.ExprNumber),
		v.Type(),
	).(ops.ExprBool)

	minusCondition := ops.Binary(
		ctx.CustomContext,
		ops.Gte,
		ops.Variable(ctx.CustomContext, name, v).(ops.ExprNumber),
		ops.Variable(ctx.CustomContext, endName, endV).(ops.ExprNumber),
		v.Type(),
	).(ops.ExprBool)

	// сам цикл
	bodies := [...]*ops.Statement{ops.Placeholder(), ops.Placeholder()}
	plusCondOp := ops.IfTrue(ctx.CustomContext, plusCondition, bodies[0])
	minusCondOp := ops.IfTrue(ctx.CustomContext, minusCondition, bodies[1])
	zero := ops.Constant(types.IntVariable(0)).(ops.ExprInt64)
	chooseCond := ops.Binary(
		ctx.CustomContext,
		ops.Gt,
		ops.Variable(ctx.CustomContext, stepName, stepV).(ops.ExprNumber),
		zero,
		stepV.Type(),
	).(ops.ExprBool)
	chooseCondOp := ops.IfTrue(ctx.CustomContext, chooseCond, plusCondOp)
	x.AddToStatementChain(chooseCondOp, minusCondOp)

	x.AddToStatementChain(result, chooseCondOp)

	// TODO тут развилка в зависимости положительный или отрицательный шаг
	// но при этом x.cycle.out общий
	x.cycle.out = ops.Placeholder()
	for i, body := range bodies {
		if i == 0 {
			x.AddToStatementChain(plusCondOp, x.cycle.out)
			x.cycle.condition = plusCondOp
		} else {
			x.AddToStatementChain(minusCondOp, x.cycle.out)
			x.cycle.condition = minusCondOp
		}
		// Accept выполняется после, так как ему нужен x.cycle
		x.AddToStatementChain(body, ctx.Statement_list().Accept(x).(*ops.Statement))
		stepOp := ops.Binary(
			ctx.CustomContext,
			ops.Plus,
			ops.Variable(ctx.CustomContext, name, v).(ops.ExprNumber),
			ops.Variable(ctx.CustomContext, stepName, stepV).(ops.ExprNumber),
			v.Type(),
		).(ops.ExprNumber)
		x.AddToStatementChain(body, ops.Assign(
			ctx.CustomContext,
			name,
			v,
			stepOp,
		))
		x.AddToStatementChain(body, x.cycle.condition) // зацикливание на условии
	}

	return result
}

// VisitFuncCallExpression implements parser.STVisitor.
func (x visitor) VisitFuncCallExpression(ctx *parser.FuncCallExpressionContext) any {
	return ctx.Function_invocation().Accept(x)
}

// VisitFunction_block_decl implements parser.STVisitor.
func (x visitor) VisitFunction_block_decl(ctx *parser.Function_block_declContext) any {
	panic("unimplemented")
}

// VisitFunction_decl implements parser.STVisitor.
func (x visitor) VisitFunction_decl(ctx *parser.Function_declContext) any {
	funcName := ctx.GetId().GetText() // TODO общий код с VisitVar_decl
	valueType := ctx.GetResType().Accept(x).(string)
	v := types.SetType(types.NewAnyVariable(""), types.TypeFromString(valueType))
	x.vars[funcName] = v
	x.vars[funcName+"."+funcName] = v // TODO пока проще так (обращение к результату по имени функции внутри функции)
	f := function{
		name: funcName,
		args: nil,
	}

	x.state = funcName + "."
	for _, a := range ctx.AllInput_decl() {
		v := a.Accept(x).(variable)
		f.args = append(f.args, v)
	}

	body := ops.Placeholder()
	for _, a := range ctx.AllVar_decl() {
		v := a.Accept(x).(variable)
		expr := ops.Constant(v.defaultV)
		x.AddToStatementChain(body, ops.Assign(ctx.CustomContext, v.name, x.vars[v.name], expr))
	}

	x.AddToStatementChain(body, ctx.Statement_list().Accept(x).(*ops.Statement))
	f.body = body

	x.funcs[funcName] = f
	return nil
}

// VisitFunction_invocation implements parser.STVisitor.
func (x visitor) VisitFunction_invocation(ctx *parser.Function_invocationContext) any {
	funcName := ctx.IDENTIFIER().GetText()
	body := ops.Placeholder()
	// собрать аргументы и положить в перменные
	if f, ok := x.funcs[funcName]; ok {
		// присвоить аргументы функции
		// TODO проверка на количество
		for i, a := range ctx.GetArgs() {
			arg := f.args[i]
			varName := arg.name
			v, ok := x.vars[varName]
			if !ok {
				x.CheckError(a, fmt.Errorf("no variable with name: %q", varName))
			}
			expr, err := ops.MakeExprNumber(a.Accept(x))
			x.CheckError(a, err)
			if !expr.IsConstant() {
				x.CheckError(a, types.CheckAssign(v.Type(), expr.Type()))
			}
			x.AddToStatementChain(body, ops.Assign(ctx.CustomContext, varName, v, expr))
		}

		// вызвать тело функции, в ней результат будет лежать уже
		x.AddToStatementChain(body, f.body)
		// вернуть переменную
		v, ok := x.vars[funcName] // TODO вынести в common вместе с доступом к переменной?
		if !ok {
			x.CheckError(ctx, fmt.Errorf("unknown function name: %q", funcName))
		}
		if v.Type().IsInteger() { // TODO другие типы
			return ops.StatementVariable[int64](ctx.CustomContext, body, funcName, v)
		}
		if v.Type().IsFloat() {
			return ops.StatementVariable[float64](ctx.CustomContext, body, funcName, v)
		}
		return x.CheckError(ctx, fmt.Errorf("unimplemented type of variable in function call result: %v", v.Type()))
	}

	parts := strings.Split(funcName, "_TO_")
	if len(parts) != 2 {
		x.CheckError(ctx, errors.New("unknown function"))
	}
	from := types.TypeFromString(parts[0])
	if from == types.ANY {
		x.CheckError(ctx, errors.New("unknown function"))
	}
	to := types.TypeFromString(parts[1])
	if to == types.ANY {
		x.CheckError(ctx, errors.New("unknown function"))
	}

	args := ctx.GetArgs()
	if len(args) != 1 {
		x.CheckError(ctx, errors.New("cast must have only one argument"))
	}

	expr := args[0].Accept(x).(ops.ExprNumber)
	if from != expr.Type() {
		x.CheckError(ctx, fmt.Errorf("expression has %s type", expr.Type()))
	}
	return ops.Cast(ctx.CustomContext, to, expr)
}

// VisitIf_statement implements parser.STVisitor.
func (x visitor) VisitIf_statement(ctx *parser.If_statementContext) any {
	conditions := ctx.GetConds()
	thens := ctx.GetThens()
	var enterOp, lastIf *ops.Statement
	out := ops.Placeholder() // для выхода из всех body
	for i := range conditions {
		cond := conditions[i].Accept(x).(ops.ExprBool)
		body := thens[i].Accept(x).(*ops.Statement)
		x.AddToStatementChain(body, out)
		prevIf := lastIf
		lastIf = ops.IfTrue(ctx.CustomContext, cond, body)
		if prevIf == nil {
			enterOp = lastIf
			continue
		}
		x.AddToStatementChain(prevIf, lastIf)
	}
	if ctx.GetElse_() != nil {
		// else добаляется в конец последнего if
		s := ctx.GetElse_().Accept(x).(*ops.Statement)
		x.AddToStatementChain(lastIf, s)
		x.AddToStatementChain(s, out)
	} else {
		x.AddToStatementChain(lastIf, out)
	}
	return enterOp
}

// VisitInput_decl implements parser.STVisitor.
func (x visitor) VisitInput_decl(ctx *parser.Input_declContext) any {
	return ctx.Var_decl().Accept(x)
}

// VisitLiteral implements parser.STVisitor.
func (x visitor) VisitLiteral(ctx *parser.LiteralContext) any {
	switch {
	case ctx.INT_LITERAL() != nil:
		v := ctx.INT_LITERAL().GetText()
		if v, ok := strings.CutPrefix(v, "16#"); ok {
			v = strings.ReplaceAll(v, "_", "")
			bint, ok := big.NewInt(0).SetString(v, 16)
			if ok {
				return ops.Constant(types.IntVariable(bint.Int64()))
			}
			x.CheckError(ctx, fmt.Errorf("fail parse int from %q", v))
		}
		i, err := strconv.Atoi(ctx.INT_LITERAL().GetText())
		x.CheckError(ctx, err)
		return ops.Constant(types.IntVariable(int64(i)))

	case ctx.REAL_LITERAL() != nil:
		v := ctx.REAL_LITERAL()
		result, err := strconv.ParseFloat(v.GetText(), 64)
		x.CheckError(ctx, err)
		return ops.Constant(types.Float64Variable(result))
	}
	return x.CheckError(ctx, fmt.Errorf("unknown literal: %q", ctx.GetText()))
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
	for _, v := range ctx.GetVars() {
		v.Accept(x)
	}
	x.programs = append(x.programs, ctx.GetStmts().Accept(x).(*ops.Statement))
	return nil
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
		next = x.AddToStatementChain(next, tmp)
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
	expr, err := ops.MakeExprNumber(ctx.Expression().Accept(x))
	x.CheckError(ctx.Expression(), err)
	token := ctx.GetOperator().GetTokenType()
	switch token {
	case parser.STLexerNOT, parser.STLexerPLUS:
		return x.CheckError(ctx, fmt.Errorf("unimplemented unary expression with operator %q (token: %d)", ctx.GetOperator().GetText(), token))
	case parser.STLexerMINUS:
		return ops.UnaryMinus(ctx.CustomContext, expr)
	}
	return expr
}

// VisitVarExpression implements parser.STVisitor.
func (x visitor) VisitVarExpression(ctx *parser.VarExpressionContext) any {
	return ctx.Variable().Accept(x)
}

// VisitVar_decl implements parser.STVisitor.
func (x visitor) VisitVar_decl(ctx *parser.Var_declContext) any {
	varName := x.state + ctx.GetId().GetText()
	variable := variable{
		name:     varName,
		defaultV: types.NewAnyVariable(""),
	}
	var v types.Variable
	if ctx := ctx.GetExpr(); ctx != nil {
		switch op := ctx.Accept(x).(type) {
		case ops.ExprInt64:
			tmp, _ := op.Do(nil)
			v = types.IntVariable(tmp)
			variable.defaultV = v
		case ops.ExprFloat64:
			tmp, _ := op.Do(nil)
			v = types.Float64Variable(tmp)
			variable.defaultV = v
		default:
			return x.CheckError(ctx, fmt.Errorf("unknown data type %T(%v)", op, op))
		}
	} else {
		v = types.NewAnyVariable("")
	}
	tmp := ctx.GetType_().Accept(x).(string)
	valueType := types.TypeFromString(tmp)
	variable.type_ = valueType
	variable.defaultV = types.SetType(variable.defaultV, valueType)
	x.vars[varName] = types.SetType(variable.defaultV, valueType)
	return variable
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
	varName := x.state + ctx.GetName().GetText() // TODO там немного сложнее со структурами/массивами
	v, ok := x.vars[varName]
	if !ok {
		x.CheckError(ctx, fmt.Errorf("unknown variable: %q", varName))
	}
	res := ops.Variable(ctx.CustomContext, varName, v)
	if res == nil {
		return x.CheckError(ctx, fmt.Errorf("unknown type of variable: %s", v.Type()))
	}
	return res
}

// VisitVariable_list implements parser.STVisitor.
func (x visitor) VisitVariable_list(ctx *parser.Variable_listContext) any {
	panic("unimplemented")
}

// VisitWhile_statement implements parser.STVisitor.
func (x visitor) VisitWhile_statement(ctx *parser.While_statementContext) any {
	condition := ctx.Expression().Accept(x).(ops.ExprBool)
	body := ops.Placeholder()
	condOp := ops.IfTrue(ctx.CustomContext, condition, body)
	x.cycle.condition = condOp
	x.cycle.out = x.AddToStatementChain(condOp, ops.Placeholder())
	// Accept выполняется после, так как ему нужен x.cycle
	x.AddToStatementChain(body, ctx.Statement_list().Accept(x).(*ops.Statement))
	x.AddToStatementChain(body, condOp) // зацикливание на условии
	return condOp
}

var (
	_ parser.STVisitor    = visitor{}
	_ antlr.ErrorListener = (*visitor)(nil)
)
