package ast

import (
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
	_ parser.STVisitor    = (*visitor)(nil)
	_ antlr.ErrorListener = (*visitor)(nil)
)

// implements parser.STVisitor.
func (x visitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(x)
}

func (x visitor) VisitAssignment_statement(ctx *parser.Assignment_statementContext) any {
	varName := ctx.GetLeft().GetText()
	expr := ctx.GetRight().Accept(x).(Int64Operator)
	variable := x.vars[varName]
	return NewAssignOp((*int64)(variable.Pointer()), expr)
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
	if err != nil {
		// TODO прервать с ошибкой
	}
	opFunc := binaryOps[binaryOpKey{op, resultType}]
	return opFunc(left, right)
}

type binaryOpKey struct {
	op string
	t  variant.Type
}

var binaryOps = map[binaryOpKey]func(left, right Int64Operator) any{
	{"+", variant.SINT}:  func(left, right Int64Operator) any { return NewPlusOp[int8](left, right) },
	{"+", variant.INT}:   func(left, right Int64Operator) any { return NewPlusOp[int16](left, right) },
	{"+", variant.DINT}:  func(left, right Int64Operator) any { return NewPlusOp[int32](left, right) },
	{"+", variant.LINT}:  func(left, right Int64Operator) any { return NewPlusOp[int64](left, right) },
	{"+", variant.USINT}: func(left, right Int64Operator) any { return NewPlusOp[uint8](left, right) },
	{"+", variant.UINT}:  func(left, right Int64Operator) any { return NewPlusOp[uint16](left, right) },
	{"+", variant.UDINT}: func(left, right Int64Operator) any { return NewPlusOp[uint32](left, right) },
	{"+", variant.ULINT}: func(left, right Int64Operator) any { return NewPlusOp[uint64](left, right) },

	{"-", variant.SINT}:  func(left, right Int64Operator) any { return NewSubOp[int8](left, right) },
	{"-", variant.INT}:   func(left, right Int64Operator) any { return NewSubOp[int16](left, right) },
	{"-", variant.DINT}:  func(left, right Int64Operator) any { return NewSubOp[int32](left, right) },
	{"-", variant.LINT}:  func(left, right Int64Operator) any { return NewSubOp[int64](left, right) },
	{"-", variant.USINT}: func(left, right Int64Operator) any { return NewSubOp[uint8](left, right) },
	{"-", variant.UINT}:  func(left, right Int64Operator) any { return NewSubOp[uint16](left, right) },
	{"-", variant.UDINT}: func(left, right Int64Operator) any { return NewSubOp[uint32](left, right) },
	{"-", variant.ULINT}: func(left, right Int64Operator) any { return NewSubOp[uint64](left, right) },

	{"*", variant.SINT}:  func(left, right Int64Operator) any { return NewMultOp[int8](left, right) },
	{"*", variant.INT}:   func(left, right Int64Operator) any { return NewMultOp[int16](left, right) },
	{"*", variant.DINT}:  func(left, right Int64Operator) any { return NewMultOp[int32](left, right) },
	{"*", variant.LINT}:  func(left, right Int64Operator) any { return NewMultOp[int64](left, right) },
	{"*", variant.USINT}: func(left, right Int64Operator) any { return NewMultOp[uint8](left, right) },
	{"*", variant.UINT}:  func(left, right Int64Operator) any { return NewMultOp[uint16](left, right) },
	{"*", variant.UDINT}: func(left, right Int64Operator) any { return NewMultOp[uint32](left, right) },
	{"*", variant.ULINT}: func(left, right Int64Operator) any { return NewMultOp[uint64](left, right) },

	{"/", variant.SINT}:  func(left, right Int64Operator) any { return NewDivOp[int8](left, right) },
	{"/", variant.INT}:   func(left, right Int64Operator) any { return NewDivOp[int16](left, right) },
	{"/", variant.DINT}:  func(left, right Int64Operator) any { return NewDivOp[int32](left, right) },
	{"/", variant.LINT}:  func(left, right Int64Operator) any { return NewDivOp[int64](left, right) },
	{"/", variant.USINT}: func(left, right Int64Operator) any { return NewDivOp[uint8](left, right) },
	{"/", variant.UINT}:  func(left, right Int64Operator) any { return NewDivOp[uint16](left, right) },
	{"/", variant.UDINT}: func(left, right Int64Operator) any { return NewDivOp[uint32](left, right) },
	{"/", variant.ULINT}: func(left, right Int64Operator) any { return NewDivOp[uint64](left, right) },

	{"MOD", variant.SINT}:  func(left, right Int64Operator) any { return NewModOp[int8](left, right) },
	{"MOD", variant.INT}:   func(left, right Int64Operator) any { return NewModOp[int16](left, right) },
	{"MOD", variant.DINT}:  func(left, right Int64Operator) any { return NewModOp[int32](left, right) },
	{"MOD", variant.LINT}:  func(left, right Int64Operator) any { return NewModOp[int64](left, right) },
	{"MOD", variant.USINT}: func(left, right Int64Operator) any { return NewModOp[uint8](left, right) },
	{"MOD", variant.UINT}:  func(left, right Int64Operator) any { return NewModOp[uint16](left, right) },
	{"MOD", variant.UDINT}: func(left, right Int64Operator) any { return NewModOp[uint32](left, right) },
	{"MOD", variant.ULINT}: func(left, right Int64Operator) any { return NewModOp[uint64](left, right) },
}

func (x visitor) VisitBinaryPowerExpr(ctx *parser.BinaryPowerExprContext) any {
	op := ctx.GetOp().GetText()
	right := ctx.GetRight().Accept(x).(Int64Operator)
	left := ctx.GetLeft().Accept(x).(Int64Operator)
	resultType, err := validateTypes(right, left)
	if err != nil {
		// TODO прервать с ошибкой
	}
	opFunc := binaryOps[binaryOpKey{op, resultType}]
	return opFunc(left, right)
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
	v := variant.NewAnyVariant(ctx.GetText())
	return NewConstantOp(v.Int())
}
func (x visitor) VisitNumber(ctx *parser.NumberContext) any {
	if ctx := ctx.Signed_integer(); ctx != nil {
		return ctx.Accept(x)
	}
	if ctx := ctx.Integer(); ctx != nil {
		return ctx.Accept(x)
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
	v := variant.NewAnyVariant(ctx.GetText())
	return NewConstantOp(v.Int())
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
	defaultValue := ""
	if token := ctx.GetDefault_(); token != nil { // TODO надо через Accept, чтоб найти ошибку паринга (как тип ниже)
		defaultValue = token.GetText()
	}
	valueType := ctx.GetType_().Accept(x).(string)
	v := variant.SetType(variant.NewAnyVariant(defaultValue), variant.TypeFromString(valueType))
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
