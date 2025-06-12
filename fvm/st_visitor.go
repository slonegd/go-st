package fvm

import (
	"github.com/antlr4-go/antlr/v4"
	parser "github.com/slonegd/go-st/antlr"
	"github.com/slonegd/go-st/variant"
)

var _ parser.STVisitor = (*FVM)(nil)

// implements parser.STVisitor.
func (x *FVM) Visit(tree antlr.ParseTree) any {
	return tree.Accept(x)
}

func (x *FVM) VisitAssignment_statement(ctx *parser.Assignment_statementContext) any {
	varName := ctx.GetLeft().GetText()
	expr := ctx.GetRight().Accept(x).(Int64Operator)
	variable := x.vars[varName]
	return NewAssignOp((*int64)(variable.Pointer()), expr)
}

func (x *FVM) VisitBinaryCompareExpr(ctx *parser.BinaryCompareExprContext) any {
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

func (x *FVM) VisitBinaryPlusExpr(ctx *parser.BinaryPlusExprContext) any {
	op := ctx.GetOp().GetText()
	right := ctx.GetRight().Accept(x).(Int64Operator)
	left := ctx.GetLeft().Accept(x).(Int64Operator)
	switch op { // TODO через id токена
	case "+":
		return NewPlusOp(left, right)
	case "-":
		return NewSubOp(left, right)
	}
	panic(ctx)
}

func (x *FVM) VisitBinaryPowerExpr(ctx *parser.BinaryPowerExprContext) any {
	op := ctx.GetOp().GetText()
	right := ctx.GetRight().Accept(x).(Int64Operator) // x.operators.Pop().(Int64Operator)
	left := ctx.GetLeft().Accept(x).(Int64Operator)   //  x.operators.Pop().(Int64Operator)
	switch op {                                       // TODO через id токена
	case "*":
		return NewMultOp(left, right)
	case "/":
		return NewDivOp(left, right)
	case "MOD":
		return NewModOp(left, right)
	}
	panic(ctx)
}

func (x *FVM) VisitChildren(node antlr.RuleNode) any { panic(node) }
func (x *FVM) VisitCondition(ctx *parser.ConditionContext) any {
	return ctx.Expression().Accept(x)
}

func (x *FVM) VisitConstant(ctx *parser.ConstantContext) any {
	v := variant.NewAnyVariant(ctx.GetText())
	return NewConstantOp(v.Int())
}

func (x *FVM) VisitElse_list(ctx *parser.Else_listContext) any {
	return ctx.Statement_list().Accept(x)
}
func (x *FVM) VisitErrorNode(node antlr.ErrorNode) any { panic(node) }
func (x *FVM) VisitIf_statement(ctx *parser.If_statementContext) any {
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
func (x *FVM) VisitInteger(ctx *parser.IntegerContext) any     { panic(ctx) }
func (x *FVM) VisitNumber(ctx *parser.NumberContext) any       { panic(ctx) }
func (x *FVM) VisitParenExpr(ctx *parser.ParenExprContext) any { return ctx.GetSub().Accept(x) }

func (x *FVM) VisitProgram(ctx *parser.ProgramContext) any {
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

func (x *FVM) VisitSigned_integer(ctx *parser.Signed_integerContext) any { panic(ctx) }

func (x *FVM) VisitStatement(ctx *parser.StatementContext) any {
	if ctx := ctx.Assignment_statement(); ctx != nil {
		return ctx.Accept(x)
	}
	if ctx := ctx.If_statement(); ctx != nil {
		return ctx.Accept(x)
	}
	panic(ctx)
}

func (x *FVM) VisitStatement_list(ctx *parser.Statement_listContext) any {
	stmts := make([]Statement, 0)
	for _, st := range ctx.AllStatement() {
		s := st.Accept(x).(Statement)
		stmts = append(stmts, s)
	}
	return NewStatments(stmts)
}

func (x *FVM) VisitTerminal(node antlr.TerminalNode) any { return nil }
func (x *FVM) VisitThen_list(ctx *parser.Then_listContext) any {
	return ctx.Statement_list().Accept(x)
}
func (x *FVM) VisitType_name(ctx *parser.Type_nameContext) any                         { panic(ctx) }
func (x *FVM) VisitVar_declaration(ctx *parser.Var_declarationContext) any             { panic(ctx) }
func (x *FVM) VisitVar_declaration_block(ctx *parser.Var_declaration_blockContext) any { panic(ctx) }

func (x *FVM) VisitVar_declaration_blocks(ctx *parser.Var_declaration_blocksContext) any {
	for _, decls := range ctx.AllVar_declaration_block() {
		for _, decl := range decls.AllVar_declaration() {
			varName := decl.GetIdentifier().GetText()
			defaultValue := ""
			if token := decl.GetDefault_(); token != nil {
				defaultValue = token.GetText()
			}
			valueType := decl.GetType_().GetText()
			v := variant.SetType(variant.NewAnyVariant(defaultValue), variant.TypeFromString(valueType))
			x.vars[varName] = v
		}
	}
	return nil
}

func (x *FVM) VisitVariable(ctx *parser.VariableContext) any {
	varName := ctx.GetText()
	v := x.vars[varName]
	return NewVarOp((*int64)(v.Pointer()))
}
