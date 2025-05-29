package compiler

import (
	"fmt"
	"log"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/slonegd/go-st/antlr"
	"github.com/slonegd/go-st/variant"
	"github.com/slonegd/go-st/vm"
)

type Compiler struct {
	*vm.VM
	varIndexes map[string]int
}

func New() *Compiler {
	return &Compiler{VM: vm.New(), varIndexes: make(map[string]int)}
}

func (x *Compiler) Compile(script string) error {
	// Setup the input
	is := antlr.NewInputStream(script)

	// Create the Lexer
	lexer := parser.NewSTLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewSTParser(stream)
	// p.BaseRecognizer.AddErrorListener(result) // TODO

	// Finally parse the expression (by walking the tree)
	antlr.ParseTreeWalkerDefault.Walk(x, p.Program())
	if p.SynErr != nil && p.SynErr.GetMessage() != "" {
		return fmt.Errorf(p.SynErr.GetMessage())
	}

	return nil
}

func (x *Compiler) GetVar(name string) variant.Variant {
	i := x.varIndexes[name]
	if i >= len(x.Vars) {
		return &variant.Any{}
	}
	return x.Vars[i]
}

// enters

// implements parser.STListener.
func (*Compiler) EnterBinaryPowerExpr(c *parser.BinaryPowerExprContext)             {}
func (x *Compiler) EnterAssignment_statement(c *parser.Assignment_statementContext) {}
func (x *Compiler) EnterBinaryPlusExpr(c *parser.BinaryPlusExprContext)             {}
func (x *Compiler) EnterConstant(c *parser.ConstantContext)                         {}
func (x *Compiler) EnterEveryRule(ctx antlr.ParserRuleContext)                      {}
func (x *Compiler) EnterNumber(c *parser.NumberContext)                             {}
func (x *Compiler) EnterParenExpr(c *parser.ParenExprContext)                       {}
func (x *Compiler) EnterProgram(c *parser.ProgramContext)                           {}
func (x *Compiler) EnterStatement(c *parser.StatementContext)                       {}
func (x *Compiler) EnterStatement_list(c *parser.Statement_listContext)             {}
func (x *Compiler) EnterType_name(c *parser.Type_nameContext) {
	// TODO типизировать переменные
}
func (x *Compiler) EnterVar_declaration(c *parser.Var_declarationContext) {
	varName := c.GetIdentifier().GetText()
	defaultValue := ""
	if token := c.GetDefault_(); token != nil {
		defaultValue = token.GetText()
	}
	valueType := c.GetType_().GetText()
	v := variant.SetType(variant.NewAnyVariant(defaultValue), variant.TypeFromString(valueType))
	x.varIndexes[varName] = len(x.Vars)
	x.Vars = append(x.VM.Vars, v)
}
func (*Compiler) EnterVar_declaration_block(c *parser.Var_declaration_blockContext)   {}
func (*Compiler) EnterVar_declaration_blocks(c *parser.Var_declaration_blocksContext) {}
func (*Compiler) EnterVariable(c *parser.VariableContext)                             {}
func (*Compiler) EnterBinaryCompareExpr(c *parser.BinaryCompareExprContext)           {}
func (*Compiler) EnterIf_statement(c *parser.If_statementContext)                     {}
func (*Compiler) EnterElse_list(c *parser.Else_listContext)                           {}
func (*Compiler) EnterThen_list(c *parser.Then_listContext)                           {}
func (*Compiler) EnterCondition(c *parser.ConditionContext)                           {}
func (*Compiler) EnterInteger(c *parser.IntegerContext)                               {}
func (*Compiler) EnterSigned_integer(c *parser.Signed_integerContext)                 {}

// exits

func (x *Compiler) ExitAssignment_statement(c *parser.Assignment_statementContext) {
	varName := c.GetLeft().GetText()
	i := x.varIndexes[varName]
	x.Bytecode.AddOp(vm.PopVar, uintptr(i))

}
func (x *Compiler) ExitBinaryPlusExpr(c *parser.BinaryPlusExprContext) {
	op := c.GetOp().GetText()
	switch op { // TODO через id токена
	case "+":
		x.Bytecode.AddOp(vm.PlusInt)
	case "-":
		x.Bytecode.AddOp(vm.MinusInt)
	}
}
func (x *Compiler) ExitBinaryPowerExpr(c *parser.BinaryPowerExprContext) {
	op := c.GetOp().GetText()
	switch op { // TODO через id токена
	case "*":
		x.Bytecode.AddOp(vm.MultInt)
	case "/":
		x.VM.Bytecode.AddOp(vm.DivInt)
	case "MOD":
		x.Bytecode.AddOp(vm.ModInt)
	}
}
func (x *Compiler) ExitBinaryCompareExpr(c *parser.BinaryCompareExprContext) {
}

// ExitCondition implements parser.STListener.
func (x *Compiler) ExitCondition(c *parser.ConditionContext) {
}
func (x *Compiler) ExitThen_list(c *parser.Then_listContext) {
}
func (x *Compiler) ExitElse_list(c *parser.Else_listContext) {}

func (x *Compiler) ExitIf_statement(c *parser.If_statementContext) {
}

func (x *Compiler) ExitConstant(c *parser.ConstantContext)  {}
func (*Compiler) ExitEveryRule(ctx antlr.ParserRuleContext) {}

func (x *Compiler) ExitNumber(c *parser.NumberContext) {
	if _, ok := c.GetParent().(*parser.Var_declarationContext); ok {
		return // чтоб не класть на стек инициализацию
	}
	v := variant.NewAnyVariant(c.GetText())
	x.Bytecode.AddOp(vm.PushConst, uintptr(len(x.Consts)))
	x.Consts = append(x.Consts, v)
}

func (*Compiler) ExitParenExpr(c *parser.ParenExprContext)                         {}
func (*Compiler) ExitProgram(c *parser.ProgramContext)                             {}
func (*Compiler) ExitStatement(c *parser.StatementContext)                         {}
func (*Compiler) ExitStatement_list(c *parser.Statement_listContext)               {}
func (*Compiler) ExitType_name(c *parser.Type_nameContext)                         {}
func (*Compiler) ExitVar_declaration(c *parser.Var_declarationContext)             {}
func (*Compiler) ExitVar_declaration_block(c *parser.Var_declaration_blockContext) {}
func (x *Compiler) ExitVar_declaration_blocks(c *parser.Var_declaration_blocksContext) {
	log.Printf("variables %v, %v", x.varIndexes, x.Vars)
}
func (x *Compiler) ExitVariable(c *parser.VariableContext) {
	varName := c.GetText()
	i := x.varIndexes[varName]
	x.Bytecode.AddOp(vm.PushVar, uintptr(i))
}

func (*Compiler) ExitInteger(c *parser.IntegerContext)               {}
func (*Compiler) ExitSigned_integer(c *parser.Signed_integerContext) {}
func (*Compiler) VisitErrorNode(node antlr.ErrorNode)                {}
func (*Compiler) VisitTerminal(node antlr.TerminalNode)              {}
