package st

import (
	"log"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/slonegd/go-st/antlr"
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
	for _, id := range c.GetIdentifier() {
		varName := id.GetText()
		x.Variables[varName] = 0 // TODO значения по умолчанию
	}
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
	step := len(x.actions)
	x.actions = append(x.actions, func() int {
		value := x.stack[len(x.stack)-1]
		x.stack = x.stack[:len(x.stack)-1]
		x.Variables[varName] = value
		log.Printf("%d:\t%v\t<-\t%v\t\tstack: %v", step, varName, value, x.stack)
		return 0
	})
}
func (x *Program) ExitBinaryPlusExpr(c *parser.BinaryPlusExprContext) {
	step := len(x.actions)
	x.actions = append(x.actions, func() int {
		right := x.stack[len(x.stack)-1]
		left := x.stack[len(x.stack)-2]
		x.stack = x.stack[:len(x.stack)-2]
		op := c.GetOp().GetText()
		switch op { // TODO через id токена
		case "+":
			x.stack = append(x.stack, left+right)
		case "-":
			x.stack = append(x.stack, left-right)
		}
		log.Printf("%d:\t%v\t<-\t%v %v %v\t\tstack: %v", step, x.stack[len(x.stack)-1], left, op, right, x.stack)
		return 0
	})
}
func (x *Program) ExitBinaryPowerExpr(c *parser.BinaryPowerExprContext) {
	step := len(x.actions)
	x.actions = append(x.actions, func() int {
		right := x.stack[len(x.stack)-1]
		left := x.stack[len(x.stack)-2]
		x.stack = x.stack[:len(x.stack)-2]
		op := c.GetOp().GetText()
		switch op { // TODO через id токена или до замыкания
		case "*":
			x.stack = append(x.stack, left*right)
		case "/":
			x.stack = append(x.stack, left/right)
		case "MOD":
			x.stack = append(x.stack, left%right)
		}
		log.Printf("%d:\t%v\t<-\t%v %v %v\t\tstack: %v", step, x.stack[len(x.stack)-1], left, op, right, x.stack)
		return 0
	})
}
func (x *Program) ExitBinaryCompareExpr(c *parser.BinaryCompareExprContext) {
	step := len(x.actions)
	x.actions = append(x.actions, func() int {
		right := x.stack[len(x.stack)-1]
		left := x.stack[len(x.stack)-2]
		x.stack = x.stack[:len(x.stack)-2]
		op := c.GetOp().GetText()
		switch op { // TODO через id токена или до замыкания
		case ">":
			if left > right {
				x.stack = append(x.stack, 1) // TODO когда будут типы тут bool
			} else {
				x.stack = append(x.stack, 0)
			}
		case ">=":
			if left >= right {
				x.stack = append(x.stack, 1)
			} else {
				x.stack = append(x.stack, 0)
			}
		case "<":
			if left < right {
				x.stack = append(x.stack, 1)
			} else {
				x.stack = append(x.stack, 0)
			}
		case "<=":
			if left <= right {
				x.stack = append(x.stack, 1)
			} else {
				x.stack = append(x.stack, 0)
			}
		case "=":
			if left == right {
				x.stack = append(x.stack, 1)
			} else {
				x.stack = append(x.stack, 0)
			}
		case "<>":
			if left != right {
				x.stack = append(x.stack, 1)
			} else {
				x.stack = append(x.stack, 0)
			}
		}
		log.Printf("%d:\t%v\t<-\t%v %v %v\t\tstack: %v", step, x.stack[len(x.stack)-1], left, op, right, x.stack)
		return 0
	})
}

// ExitCondition implements parser.STListener.
func (x *Program) ExitCondition(c *parser.ConditionContext) {
	if len(x.ifs) == 0 || len(x.ifs[len(x.ifs)-1].thenIndexes) == 0 {
		x.ifs = append(x.ifs, &ifState{})
	}
	ifState := x.ifs[len(x.ifs)-1]
	lastCondition := &ifState.lastCondition
	x.actions = append(x.actions, func() int {
		*lastCondition = x.stack[len(x.stack)-1]
		x.stack = x.stack[:len(x.stack)-1]
		return 0
	})
	ifState.lastConditionIndex = len(x.actions)
	// место зарезервировано для перехода, который заполнится в конце then в зависимости от условия
	x.actions = append(x.actions, nil)
}
func (x *Program) ExitThen_list(c *parser.Then_listContext) {
	gotoIndex := len(x.actions)
	lastThenIndex := gotoIndex
	ifState := x.ifs[len(x.ifs)-1]
	ifState.thenIndexes = append(ifState.thenIndexes, lastThenIndex)
	x.actions = append(x.actions, nil) // место зарезервировано для перехода, который заполнится в конце if

	// надо делать копии и указатели, так как в замыкании затрётся
	lastConditionIndex := ifState.lastConditionIndex
	lastCondition := &ifState.lastCondition
	x.actions[lastConditionIndex] = func() int {
		if *lastCondition != 0 {
			return 0
		}
		log.Printf("%d:\tgoto %v\t\tstack: %v", lastConditionIndex, gotoIndex+1, x.stack)
		return gotoIndex + 1
	}
}
func (x *Program) ExitElse_list(c *parser.Else_listContext) {}

func (x *Program) ExitIf_statement(c *parser.If_statementContext) {
	gotoIndex := len(x.actions)
	ifState := x.ifs[len(x.ifs)-1]
	for _, lastThenIndex := range ifState.thenIndexes {
		lastThenIndex := lastThenIndex
		x.actions[lastThenIndex] = func() int {
			log.Printf("%d:\tgoto %v\t\tstack: %v", lastThenIndex, gotoIndex, x.stack)
			return gotoIndex
		}
	}
	x.ifs = x.ifs[:len(x.ifs)-1]
}

func (x *Program) ExitConstant(c *parser.ConstantContext)  {}
func (*Program) ExitEveryRule(ctx antlr.ParserRuleContext) {}

func (x *Program) ExitNumber(c *parser.NumberContext) {
	step := len(x.actions)
	i, _ := strconv.Atoi(c.GetText())
	x.actions = append(x.actions, func() int {
		x.stack = append(x.stack, i)
		log.Printf("%d:\t%v\t<-\t%v\t\tstack: %v", step, x.stack[len(x.stack)-1], x.stack[len(x.stack)-1], x.stack)
		return 0
	})
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
	step := len(x.actions)
	varName := c.GetText()
	x.actions = append(x.actions, func() int {
		x.stack = append(x.stack, x.Variables[varName])
		log.Printf("%d:\t%v\t<-\t%v\t\tstack: %v", step, x.stack[len(x.stack)-1], varName, x.stack)
		return 0
	})
}
func (*Program) VisitErrorNode(node antlr.ErrorNode)   {}
func (*Program) VisitTerminal(node antlr.TerminalNode) {}
