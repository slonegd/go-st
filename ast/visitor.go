package ast

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/slonegd/go-st/antlr"
	"github.com/slonegd/go-st/variant"
)

// обёртка, чтоб скрыть методы parser.STVisitor
type visitor struct {
	*AST
	// в некоторые ноды надо зайти, чтоб определить ошибку парсинга, но ничего не делать
	// такое бывает когда по всем детям проходишь, а там в конце, например TerminalNode
	lastOp any
}

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
	expr := ctx.GetRight().Accept(x).(Int64Operator)
	v := x.vars[varName]
	if !expr.isConstant {
		x.CheckError(ctx, implicitCastInAssign[assignTypes{v: v.Type(), expr: expr.resultType}])
	}
	return assignOps[v.Type()](v, expr)
}

func (x visitor) VisitBinaryCompareExpr(ctx *parser.BinaryCompareExprContext) any {
	op := ctx.GetOp().GetText()
	right := ctx.GetRight().Accept(x).(Int64Operator)
	left := ctx.GetLeft().Accept(x).(Int64Operator)
	switch op { // TODO через id токена
	case ">":
		return NewGreaterOp(left, right)
	case ">=":
		return NewGreaterOrEqualOp(left, right)
	case "<":
		return NewLessOp(left, right)
	case "<=":
		return NewLessOrEqualOp(left, right)
	case "=":
		return NewEqualOp(left, right)
	case "<>":
		return NewNotEqualOp(left, right)
	}
	panic(ctx)
}

func (x visitor) VisitBinaryPlusExpr(ctx *parser.BinaryPlusExprContext) any {
	op := ctx.GetOp().GetText()
	right := ctx.GetRight().Accept(x).(Int64Operator)
	left := ctx.GetLeft().Accept(x).(Int64Operator)
	resultType, err := validateTypes(right, left)
	x.CheckError(ctx, err)
	opFunc := binaryOps[binaryOpKey{op, resultType}]
	return opFunc(left, right)
}

type binaryOpKey struct {
	op string
	t  variant.Type
}

func (x visitor) VisitBinaryPowerExpr(ctx *parser.BinaryPowerExprContext) any {
	op := ctx.GetOp().GetText()
	right := ctx.GetRight().Accept(x).(Int64Operator)
	left := ctx.GetLeft().Accept(x).(Int64Operator)
	resultType, err := validateTypes(right, left)
	x.CheckError(ctx, err)
	opFunc := binaryOps[binaryOpKey{op, resultType}]
	return opFunc(left, right)
}

func (x visitor) VisitCallExpr(ctx *parser.CallExprContext) any {
	expr := ctx.GetSub().Accept(x).(Int64Operator)
	name := ctx.GetId().GetText()
	parts := strings.Split(name, "_TO_")
	if len(parts) != 2 {
		x.CheckError(ctx, errors.New("unknown function"))
	}
	from := variant.TypeFromString(parts[0])
	if from == variant.ANY {
		x.CheckError(ctx, errors.New("unknown function"))
	}
	to := variant.TypeFromString(parts[1])
	if to == variant.ANY {
		x.CheckError(ctx, errors.New("unknown function"))
	}
	if from != expr.resultType {
		x.CheckError(ctx, fmt.Errorf("expression has %s type", expr.resultType))
	}

	return castOps[to](expr)
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
	var elseStatement *Statement // указатель для maybe
	if ctx.GetElselist() != nil {
		s := ctx.GetElselist().Accept(x).(Statement)
		elseStatement = &s
	}
	// с конца, так как последний else является аргументом предпоследнего else if
	// а тот аргументом else if перед этим
	for i := len(conditions) - 1; i >= 0; i-- {
		condition := conditions[i].Accept(x).(BoolOperator)
		then_ := thens[i].Accept(x).(Statement)
		if elseStatement == nil {
			s := NewIfOperator(condition, then_)
			elseStatement = &s
		} else {
			s := NewIfElseOperator(condition, then_, *elseStatement)
			elseStatement = &s
		}
	}
	return *elseStatement
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

func (x visitor) VisitReal(ctx *parser.RealContext) any {
	intPart := ctx.GetIntPart().Accept(x)
	fracPart, exponent := int64(0), int64(0)
	if ctx := ctx.GetFracPart(); ctx != nil {
		fracPart = ctx.Accept(x).(int64)
	}
	if ctx := ctx.GetExponent(); ctx != nil {
		exponent = ctx.Accept(x).(int64)
	}
	s := fmt.Sprintf("%d.%de%d", intPart, fracPart, exponent)
	result, err := strconv.ParseFloat(s, 64)
	x.CheckError(ctx, err)
	return result
}

func (x visitor) VisitNumber(ctx *parser.NumberContext) any {
	if ctx := ctx.Real_(); ctx != nil {
		v := ctx.Accept(x).(float64)
		return NewConstantOp(v)
	}
	if ctx := ctx.Integer(); ctx != nil {
		v := ctx.Accept(x).(int64)
		return NewConstantOp(v)
	}
	return x.VisitChildren(ctx)
}
func (x visitor) VisitParenExpr(ctx *parser.ParenExprContext) any { return ctx.GetSub().Accept(x) }

func (x visitor) VisitProgram(ctx *parser.ProgramContext) any {
	statements := make(Statements, 0)
	for _, c := range ctx.GetChildren() {
		// не всё возвращает Statement
		// есть ноды подготовки, например объявления переменных
		if s, ok := x.Visit(c.(antlr.ParseTree)).(Statement); ok {
			statements = append(statements, s)
		}
	}
	return NewStatments(statements)
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
	v := variant.NewAnyVariant(ctx.GetText())
	if v, ok := v.Variant.(*variant.Int); ok {
		return v.Int()
	}
	x.CheckError(ctx, errors.New("fail to parse integer"))
	return v
}

func (x visitor) VisitStatement(ctx *parser.StatementContext) any {
	return x.VisitChildren(ctx)
}

func (x visitor) VisitStatement_list(ctx *parser.Statement_listContext) any {
	stmts := make([]Statement, 0)
	for _, st := range ctx.AllStatement() {
		// // могут быть ErrorNode
		// if s, ok := st.Accept(x).(Statement); ok {
		// 	stmts = append(stmts, s)
		// }
		stmts = append(stmts, st.Accept(x).(Statement))

	}
	return NewStatments(stmts)
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
	var v variant.Variant
	if ctx := ctx.GetDefault_(); ctx != nil {
		switch op := ctx.Accept(x).(type) {
		case Int64Operator:
			v = variant.IntVariant(op.do())
		case Float64Operator:
			v = variant.Float64Variant(op.do())
		}
	} else {
		v = variant.NewAnyVariant("")
	}
	valueType := ctx.GetType_().Accept(x).(string)
	v = variant.SetType(v, variant.TypeFromString(valueType))
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
	return NewVarOp[int64](v)
}
