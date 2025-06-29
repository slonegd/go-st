package ast

import (
	"errors"
	"fmt"
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
	}
	cycle struct {
		condition *ops.Statement
		out       *ops.Statement
	}
)

var (
	_ parser.STVisitor    = visitor{}
	_ antlr.ErrorListener = (*visitor)(nil)
)

// implements parser.STVisitor.
func (x visitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(x)
}

func (x visitor) VisitAssignment_statement(ctx *parser.Assignment_statementContext) any {
	varName := ctx.GetLeft().GetText()
	v := x.vars[varName]

	expr := ctx.GetRight().Accept(x)
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

func (x visitor) VisitBinaryCompareExpr(ctx *parser.BinaryCompareExprContext) any {
	op := ctx.GetOp().GetText()
	right := ctx.GetRight().Accept(x)
	left := ctx.GetLeft().Accept(x)
	switch right := right.(type) {
	case ops.Int64:
		switch left := left.(type) {
		case ops.Int64:
			resultType, err := types.BinaryResult(
				types.Expression{left.IsConstant, left.ResultType},
				types.Expression{right.IsConstant, right.ResultType},
			)
			x.CheckError(ctx, err)
			return ops.Compare(ctx.GetCustomContext(), op, left, right, resultType)
		case ops.Float64:
			resultType, err := types.BinaryResult(
				types.Expression{left.IsConstant, left.ResultType},
				types.Expression{right.IsConstant, right.ResultType},
			)
			x.CheckError(ctx, err)
			return ops.Compare(ctx.GetCustomContext(), op, left, right, resultType)
		}
	case ops.Float64:
		switch left := left.(type) {
		case ops.Int64:
			resultType, err := types.BinaryResult(
				types.Expression{left.IsConstant, left.ResultType},
				types.Expression{right.IsConstant, right.ResultType},
			)
			x.CheckError(ctx, err)
			return ops.Compare(ctx.GetCustomContext(), op, left, right, resultType)
		case ops.Float64:
			resultType, err := types.BinaryResult(
				types.Expression{left.IsConstant, left.ResultType},
				types.Expression{right.IsConstant, right.ResultType},
			)
			x.CheckError(ctx, err)
			return ops.Compare(ctx.GetCustomContext(), op, left, right, resultType)
		}
	}
	x.CheckError(ctx, fmt.Errorf("unsupported %T %s %T", left, op, right))
	return ops.Bool{}
}

func (x visitor) VisitBinaryPlusExpr(ctx *parser.BinaryPlusExprContext) any {
	return x.visitBinaryExpr(ctx)
}

type binaryOpKey struct {
	op string
	t  types.Basic
}

func (x visitor) VisitBinaryPowerExpr(ctx *parser.BinaryPowerExprContext) any {
	return x.visitBinaryExpr(ctx)
}

type BinaryContext interface {
	GetCustomContext() *parser.CustomContext
	GetRight() parser.IExpressionContext
	GetLeft() parser.IExpressionContext
	GetOp() antlr.Token
	antlr.ParserRuleContext
}

func (x visitor) visitBinaryExpr(ctx BinaryContext) any {
	op := ctx.GetOp().GetText()
	right := ctx.GetRight().Accept(x)
	left := ctx.GetLeft().Accept(x)
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
	return x.CheckError(ctx, fmt.Errorf("unsupported %T %s %T", left, op, right))
}

func (x visitor) VisitCallExpr(ctx *parser.CallExprContext) any {
	name := ctx.GetId().GetText()
	parts := strings.Split(name, "_TO_")
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

	switch expr := ctx.GetSub().Accept(x).(type) {
	case ops.Int64:
		if from != expr.ResultType {
			x.CheckError(ctx, fmt.Errorf("expression has %s type", expr.ResultType))
		}
		return ops.Cast(ctx.CustomContext, to, expr)
	case ops.Float64:
		if from != expr.ResultType {
			x.CheckError(ctx, fmt.Errorf("expression has %s type", expr.ResultType))
		}
		return ops.Cast(ctx.CustomContext, to, expr)
	default:
		x.CheckError(ctx, fmt.Errorf("expression operator has unknown type: %T", expr))
		return nil
	}

}

func (x visitor) VisitChildren(node antlr.RuleNode) any {
	for _, c := range node.GetChildren() {
		x.lastOp = x.Visit(c.(antlr.ParseTree))
	}
	return x.lastOp
}
func (x visitor) VisitCondition(ctx *parser.ConditionContext) any {
	return ctx.Expression().Accept(x)
}

func (x visitor) VisitConstant(ctx *parser.ConstantContext) any {
	return ctx.Number().Accept(x)
}

func (x visitor) VisitElse_list(ctx *parser.Else_listContext) any {
	return ctx.Statement_list().Accept(x)
}
func (x visitor) VisitErrorNode(node antlr.ErrorNode) any {
	// через antlr.ErrorListener ошибка информативнее
	// x.err = fmt.Errorf("error node %+v", node)
	return x.lastOp
}
func (x visitor) VisitIf_statement(ctx *parser.If_statementContext) any {
	conditions := ctx.GetCond()
	thens := ctx.GetThenlist()
	var firstIf, lastIf *ops.Statement
	out := ops.Placeholder() // для выхода из всех body
	for i := range conditions {
		cond := conditions[i].Accept(x).(ops.Bool)
		body := thens[i].Accept(x).(*ops.Statement)
		x.SetNextStatement(body, out)
		prevIf := lastIf
		lastIf = ops.IfTrue(ctx.CustomContext, cond, body)
		if prevIf == nil {
			firstIf = lastIf
			continue
		}
		x.SetNextStatement(prevIf, lastIf)
	}
	if ctx.GetElselist() != nil {
		// else добаляется в конец последнего if
		s := ctx.GetElselist().Accept(x).(*ops.Statement)
		x.SetNextStatement(lastIf, s)
		x.SetNextStatement(s, out)
	} else {
		x.SetNextStatement(lastIf, out)
	}
	return firstIf
}

func (x visitor) VisitWhile_statement(ctx *parser.While_statementContext) any {
	condition := ctx.GetCond().Accept(x).(ops.Bool)
	body := ops.Placeholder()
	condOp := ops.IfTrue(ctx.CustomContext, condition, body)
	x.cycle.condition = condOp
	x.cycle.out = x.SetNextStatement(condOp, ops.Placeholder())
	// Accept выполняется после, так как ему нужен x.cycle
	x.SetNextStatement(body, ctx.GetBody().Accept(x).(*ops.Statement))
	x.SetNextStatement(body, condOp) // зацикливание на условии
	return condOp
}

func (x visitor) VisitContinue_statement(ctx *parser.Continue_statementContext) any {
	if x.cycle.condition == nil {
		return x.CheckError(ctx, errors.New("no cycle body there"))
	}
	return ops.Jump(ctx.CustomContext, x.cycle.condition, "CONTINUE")
}

func (x visitor) VisitExit_statement(ctx *parser.Exit_statementContext) any {
	if x.cycle.out == nil {
		return x.CheckError(ctx, errors.New("no cycle body there"))
	}
	return ops.Jump(ctx.CustomContext, x.cycle.out, "EXIT")
}

func (x visitor) VisitInteger(ctx *parser.IntegerContext) any {
	if ctx := ctx.Unsign_integer(); ctx != nil {
		v := ctx.Accept(x).(int64)
		return v

	}
	if ctx := ctx.Signed_integer(); ctx != nil {
		v := ctx.Accept(x).(int64)
		return v
	}
	return x.VisitChildren(ctx)
}

func (x visitor) VisitNumber(ctx *parser.NumberContext) any {
	if f := ctx.FLOAT(); f != nil {
		result, err := strconv.ParseFloat(f.GetText(), 64)
		x.CheckError(ctx, err)
		return ops.Constant(result)
	}
	if ctx := ctx.Integer(); ctx != nil {
		v := ctx.Accept(x).(int64)
		return ops.Constant(v)
	}
	return x.VisitChildren(ctx)
}
func (x visitor) VisitParenExpr(ctx *parser.ParenExprContext) any { return ctx.GetSub().Accept(x) }

func (x visitor) VisitProgram(ctx *parser.ProgramContext) any {
	var r, next *ops.Statement
	// TODO через statement list?
	for _, c := range ctx.GetChildren() {
		// не всё возвращает Statement
		// есть ноды подготовки, например объявления переменных
		if s, ok := x.Visit(c.(antlr.ParseTree)).(*ops.Statement); ok {
			if r == nil {
				r = s
				next = s
			} else {
				x.SetNextStatement(next, s)
				next = s
			}
		}
	}
	return r
}

func (x visitor) VisitSigned_integer(ctx *parser.Signed_integerContext) any {
	v := ctx.Unsign_integer().Accept(x).(int64)
	if ctx.GetSign().GetText() == "-" {
		v = -v
	}
	return v
}

func (x visitor) VisitUnsign_integer(ctx *parser.Unsign_integerContext) any {
	// TODO тут не надо возиться с вариантом
	v := types.NewAnyVariable(ctx.GetText())
	if v, ok := v.Variable.(*types.Int); ok {
		return v.Int()
	}
	x.CheckError(ctx, errors.New("fail to parse integer"))
	return v
}

func (x visitor) VisitStatement(ctx *parser.StatementContext) any {
	return x.VisitChildren(ctx)
}

func (x visitor) VisitStatement_list(ctx *parser.Statement_listContext) any {
	var r, next *ops.Statement
	for _, s := range ctx.AllStatement() { // TODO проверить если пустой список и может ли такое быть?
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

func (x visitor) VisitTerminal(node antlr.TerminalNode) any { return x.lastOp }
func (x visitor) VisitThen_list(ctx *parser.Then_listContext) any {
	return ctx.Statement_list().Accept(x)
}
func (x visitor) VisitType_name(ctx *parser.Type_nameContext) any {
	return ctx.GetText()
}
func (x visitor) VisitVar_declaration(ctx *parser.Var_declarationContext) any {
	varName := ctx.GetIdentifier().GetText()
	var v types.Variable
	if ctx := ctx.GetDefault_(); ctx != nil {
		switch op := ctx.Accept(x).(type) {
		case ops.Int64:
			tmp, _ := op.Do(nil)
			v = types.IntVariable(tmp)
		case ops.Float64:
			tmp, _ := op.Do(nil)
			v = types.Float64Variable(tmp)
		}
	} else {
		v = types.NewAnyVariable("")
	}
	valueType := ctx.GetType_().Accept(x).(string)
	v = types.SetType(v, types.TypeFromString(valueType))
	x.vars[varName] = v
	return x.lastOp
}
func (x visitor) VisitVar_declaration_block(ctx *parser.Var_declaration_blockContext) any {
	var r any
	for _, decl := range ctx.AllVar_declaration() {
		r = decl.Accept(x)
	}
	return r
}

func (x visitor) VisitVar_declaration_blocks(ctx *parser.Var_declaration_blocksContext) any {
	var r any
	for _, decls := range ctx.AllVar_declaration_block() {
		r = decls.Accept(x)
	}
	return r
}

func (x visitor) VisitVariable(ctx *parser.VariableContext) any {
	varName := ctx.GetText()
	v := x.vars[varName]

	if v.Type().IsInteger() {
		return ops.Variable[int64](ctx.CustomContext, varName, v)
	}
	if v.Type().IsFloat() {
		return ops.Variable[float64](ctx.CustomContext, varName, v)
	}
	return x.CheckError(ctx, fmt.Errorf("unknown type of variable: %s", v.Type()))
}
