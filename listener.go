package st

import (
	"fmt"
	"log"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/slonegd/go-st/antlr"
	"github.com/slonegd/go-st/variant"
)

// enters

// implements parser.STListener.
func (*Program) EnterBinaryPowerExpr(c *parser.BinaryPowerExprContext)             {}
func (x *Program) EnterAssignment_statement(c *parser.Assignment_statementContext) {}
func (x *Program) EnterBinaryPlusExpr(c *parser.BinaryPlusExprContext)             {}
func (x *Program) EnterConstant(c *parser.ConstantContext)                         {}
func (x *Program) EnterEveryRule(ctx antlr.ParserRuleContext)                      {}
func (x *Program) EnterNumber(c *parser.NumberContext)                             {}
func (x *Program) EnterParenExpr(c *parser.ParenExprContext)                       {}
func (x *Program) EnterProgram(c *parser.ProgramContext)                           {}
func (x *Program) EnterStatement(c *parser.StatementContext)                       {}
func (x *Program) EnterStatement_list(c *parser.Statement_listContext)             {}
func (x *Program) EnterType_name(c *parser.Type_nameContext) {
	// TODO типизировать переменные
}
func (x *Program) EnterVar_declaration(c *parser.Var_declarationContext) {
	varName := c.GetIdentifier().GetText()
	defaultValue := ""
	if token := c.GetDefault_(); token != nil {
		defaultValue = token.GetText()
	}
	valueType := c.GetType_().GetText()
	v := variant.SetType(variant.NewAnyVariant(defaultValue), variant.TypeFromString(valueType))
	x.Variables[varName] = v
}
func (*Program) EnterVar_declaration_block(c *parser.Var_declaration_blockContext)   {}
func (*Program) EnterVar_declaration_blocks(c *parser.Var_declaration_blocksContext) {}
func (*Program) EnterVariable(c *parser.VariableContext)                             {}
func (*Program) EnterBinaryCompareExpr(c *parser.BinaryCompareExprContext)           {}
func (*Program) EnterIf_statement(c *parser.If_statementContext)                     {}
func (*Program) EnterElse_list(c *parser.Else_listContext)                           {}
func (*Program) EnterThen_list(c *parser.Then_listContext)                           {}
func (*Program) EnterCondition(c *parser.ConditionContext)                           {}

// exits

func (x *Program) ExitAssignment_statement(c *parser.Assignment_statementContext) {
	varName := c.GetLeft().GetText()
	step := Step{number: len(x.steps)}
	step.action = fmt.Sprintf("step %04d: %s <- $1", step.number, varName)
	step.callback = func() int {
		log.Printf(step.action+"\t%v", x.stack)
		value := x.stack.Pop()
		x.Variables[varName] = value
		return 0
	}
	x.steps = append(x.steps, step)
}
func (x *Program) ExitBinaryPlusExpr(c *parser.BinaryPlusExprContext) {
	step := Step{number: len(x.steps)}
	op := c.GetOp().GetText()
	step.action = fmt.Sprintf("step %04d: $2 %s $1 ->", step.number, op)
	step.callback = func() int {
		log.Printf(step.action+"\t%v", x.stack)
		right := x.stack.Pop()
		left := x.stack.Pop()
		var v variant.Variant
		switch op { // TODO через id токена
		case "+":
			v = variant.Plus(left, right)
		case "-":
			v = variant.Minus(left, right)
		}
		x.stack.Push(v)
		return 0
	}
	x.steps = append(x.steps, step)
}
func (x *Program) ExitBinaryPowerExpr(c *parser.BinaryPowerExprContext) {
	step := Step{number: len(x.steps)}
	op := c.GetOp().GetText()
	step.action = fmt.Sprintf("step %04d: $2 %s $1 ->", step.number, op)
	step.callback = func() int {
		log.Printf(step.action+"\t%v", x.stack)
		right := x.stack.Pop()
		left := x.stack.Pop()
		var v variant.Variant
		switch op { // TODO через id токена или до замыкания
		case "*":
			v = variant.Mult(left, right)
		case "/":
			v = variant.Divide(left, right)
		case "MOD":
			v = variant.Mod(left, right)
		}
		x.stack.Push(v)
		return 0
	}
	x.steps = append(x.steps, step)
}
func (x *Program) ExitBinaryCompareExpr(c *parser.BinaryCompareExprContext) {
	step := Step{number: len(x.steps)}
	op := c.GetOp().GetText()
	step.action = fmt.Sprintf("step %04d: $2 %s $1 ->", step.number, op)
	step.callback = func() int {
		log.Printf(step.action+"\t%v", x.stack)
		right := x.stack.Pop()
		left := x.stack.Pop()
		var v variant.Variant
		switch op { // TODO через id токена или до замыкания
		case ">":
			v = variant.Greather(left, right)
		case ">=":
			v = variant.GreatherOrEqual(left, right)
		case "<":
			v = variant.Less(left, right)
		case "<=":
			v = variant.LessOrEqual(left, right)
		case "=":
			v = variant.Equal(left, right)
		case "<>":
			v = variant.NotEqual(left, right)
		}
		x.stack.Push(v)
		return 0
	}
	x.steps = append(x.steps, step)
}

// ExitCondition implements parser.STListener.
func (x *Program) ExitCondition(c *parser.ConditionContext) {
	if len(x.ifs) == 0 || len(x.ifs[len(x.ifs)-1].thenIndexes) == 0 {
		x.ifs = append(x.ifs, &ifState{})
	}
	ifState := x.ifs[len(x.ifs)-1]
	lastCondition := &ifState.lastCondition

	step := Step{number: len(x.steps)}
	step.action = fmt.Sprintf("step %04d: $1?", step.number)
	action := step.action
	step.callback = func() int {
		log.Printf(action+"\t\t%v", x.stack)
		*lastCondition = x.stack.Pop().Bool()
		return 0
	}
	x.steps = append(x.steps, step)
	step = Step{number: len(x.steps)}
	// место зарезервировано для перехода, который заполнится в конце then в зависимости от условия
	x.steps = append(x.steps, step)
	ifState.lastConditionIndex = step.number
}
func (x *Program) ExitThen_list(c *parser.Then_listContext) {
	step := Step{number: len(x.steps)}
	gotoIndex := step.number + 1
	ifState := x.ifs[len(x.ifs)-1]
	ifState.thenIndexes = append(ifState.thenIndexes, step.number)
	x.steps = append(x.steps, step) // место зарезервировано для перехода, который заполнится в конце if

	// надо делать копии и указатели, так как в замыкании затрётся
	lastConditionIndex := ifState.lastConditionIndex
	lastCondition := &ifState.lastCondition

	stepRef := &x.steps[lastConditionIndex]
	stepRef.action = fmt.Sprintf("step %04d: cond ? nop : goto %d", stepRef.number, gotoIndex)
	stepRef.callback = func() int {
		log.Printf(stepRef.action+", cond: %v", *lastCondition)
		if *lastCondition {
			return 0
		}
		return gotoIndex
	}
}
func (x *Program) ExitElse_list(c *parser.Else_listContext) {}

func (x *Program) ExitIf_statement(c *parser.If_statementContext) {
	gotoIndex := len(x.steps)
	ifState := x.ifs[len(x.ifs)-1]
	for _, lastThenIndex := range ifState.thenIndexes {
		lastThenIndex := lastThenIndex
		x.steps[lastThenIndex].action = fmt.Sprintf("step %04d: goto %d", x.steps[lastThenIndex].number, gotoIndex)
		x.steps[lastThenIndex].callback = func() int {
			log.Printf(x.steps[lastThenIndex].action)
			return gotoIndex
		}
	}
	x.ifs = x.ifs[:len(x.ifs)-1]
}

func (x *Program) ExitConstant(c *parser.ConstantContext)  {}
func (*Program) ExitEveryRule(ctx antlr.ParserRuleContext) {}

func (x *Program) ExitNumber(c *parser.NumberContext) {
	v := variant.NewAnyVariant(c.GetText())
	step := Step{number: len(x.steps)}
	step.action = fmt.Sprintf("step %04d: %v ->", step.number, v)
	step.callback = func() int {
		log.Printf(step.action+"\t%v", x.stack)
		x.stack.Push(v)
		return 0
	}
	x.steps = append(x.steps, step)
}

func (*Program) ExitParenExpr(c *parser.ParenExprContext)                         {}
func (*Program) ExitProgram(c *parser.ProgramContext)                             {}
func (*Program) ExitStatement(c *parser.StatementContext)                         {}
func (*Program) ExitStatement_list(c *parser.Statement_listContext)               {}
func (*Program) ExitType_name(c *parser.Type_nameContext)                         {}
func (*Program) ExitVar_declaration(c *parser.Var_declarationContext)             {}
func (*Program) ExitVar_declaration_block(c *parser.Var_declaration_blockContext) {}
func (x *Program) ExitVar_declaration_blocks(c *parser.Var_declaration_blocksContext) {
	log.Printf("variables %v", x.Variables)
}
func (x *Program) ExitVariable(c *parser.VariableContext) {
	varName := c.GetText()
	step := Step{number: len(x.steps)}
	step.action = fmt.Sprintf("step %04d: %v ->", step.number, varName)
	step.callback = func() int {
		v := x.Variables[varName]
		log.Printf(step.action+" %v\t%v", v, x.stack)
		x.stack.Push(v)
		return 0
	}
	x.steps = append(x.steps, step)
}
func (*Program) VisitErrorNode(node antlr.ErrorNode)   {}
func (*Program) VisitTerminal(node antlr.TerminalNode) {}
