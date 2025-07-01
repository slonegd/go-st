package ast

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/slonegd/go-st/antlr"
	"github.com/slonegd/go-st/ops"
	"github.com/slonegd/go-st/types"
)

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
