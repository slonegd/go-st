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
	// только присваивания меняют состояние программы,
	// всё остальное считается на функциях внутри Statement
	// а функции определяются через визитор
	x.statements = append(x.statements, NewAssignStep((*int64)(variable.Pointer()), expr))
	return nil
}

func (x *FVM) VisitBinaryCompareExpr(ctx *parser.BinaryCompareExprContext) any { return nil }

func (x *FVM) VisitBinaryPlusExpr(ctx *parser.BinaryPlusExprContext) any {
	op := ctx.GetOp().GetText()
	right := ctx.GetRight().Accept(x).(Int64Operator)
	left := ctx.GetLeft().Accept(x).(Int64Operator)
	switch op { // TODO через id токена
	case "+":
		return NewBinaryPlusStep(left, right)
	case "-":
		return NewBinarySubStep(left, right)
	}
	return nil
}

func (x *FVM) VisitBinaryPowerExpr(ctx *parser.BinaryPowerExprContext) any {
	op := ctx.GetOp().GetText()
	right := ctx.GetRight().Accept(x).(Int64Operator) // x.operators.Pop().(Int64Operator)
	left := ctx.GetLeft().Accept(x).(Int64Operator)   //  x.operators.Pop().(Int64Operator)
	switch op {                                       // TODO через id токена
	case "*":
		return NewBinaryMultStep(left, right)
	case "/":
		return NewBinaryDivStep(left, right)
	case "MOD":
		return NewBinaryModStep(left, right)
	}
	return nil
}

func (x *FVM) VisitChildren(node antlr.RuleNode) any           { return nil }
func (x *FVM) VisitCondition(ctx *parser.ConditionContext) any { return nil }

func (x *FVM) VisitConstant(ctx *parser.ConstantContext) any {
	v := variant.NewAnyVariant(ctx.GetText())
	return NewConstantStep(v.Int())
}

func (x *FVM) VisitElse_list(ctx *parser.Else_listContext) any       { return nil }
func (x *FVM) VisitErrorNode(node antlr.ErrorNode) any               { return nil }
func (x *FVM) VisitIf_statement(ctx *parser.If_statementContext) any { return nil }
func (x *FVM) VisitInteger(ctx *parser.IntegerContext) any           { return nil }
func (x *FVM) VisitNumber(ctx *parser.NumberContext) any             { return nil }
func (x *FVM) VisitParenExpr(ctx *parser.ParenExprContext) any       { return ctx.GetSub().Accept(x) }

func (x *FVM) VisitProgram(ctx *parser.ProgramContext) any {
	for _, c := range ctx.GetChildren() {
		x.Visit(c.(antlr.ParseTree))
	}
	return nil
}

func (x *FVM) VisitSigned_integer(ctx *parser.Signed_integerContext) any { return nil }

func (x *FVM) VisitStatement(ctx *parser.StatementContext) any {
	return ctx.Assignment_statement().Accept(x)
}

func (x *FVM) VisitStatement_list(ctx *parser.Statement_listContext) any {
	for _, st := range ctx.AllStatement() {
		st.Accept(x)
	}
	return nil
}

func (x *FVM) VisitTerminal(node antlr.TerminalNode) any                               { return nil }
func (x *FVM) VisitThen_list(ctx *parser.Then_listContext) any                         { return nil }
func (x *FVM) VisitType_name(ctx *parser.Type_nameContext) any                         { return nil }
func (x *FVM) VisitVar_declaration(ctx *parser.Var_declarationContext) any             { return nil }
func (x *FVM) VisitVar_declaration_block(ctx *parser.Var_declaration_blockContext) any { return nil }

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
	return NewVarStep((*int64)(v.Pointer()))
}
