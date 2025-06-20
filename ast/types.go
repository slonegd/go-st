package ast

import (
	"errors"
	"slices"

	"github.com/slonegd/go-st/ops"
	"github.com/slonegd/go-st/types"
)

func validateTypes[Left, Right ops.NumberOpTypes](op1 ops.Op[Left], op2 ops.Op[Right]) (types.Basic, error) {
	switch {
	case op1.IsConstant && op2.IsConstant:
		return op1.ResultType, nil
	case op1.IsConstant:
		return op2.ResultType, nil
	case op2.IsConstant:
		return op1.ResultType, nil
	}
	types := [2]types.Basic{op1.ResultType, op2.ResultType}
	slices.Sort(types[:])
	r := implicitCastInBinary[types]
	return r.Basic, r.error
}

// Figure 11 – Data type conversion rules – implicit and/or explicit (Summary) 61131-3 © IEC:2013
var implicitCastInBinary = map[[2]types.Basic]struct {
	types.Basic
	error
}{
	{types.SINT, types.SINT}:  {types.SINT, nil},
	{types.SINT, types.INT}:   {types.INT, nil},
	{types.SINT, types.DINT}:  {types.DINT, nil},
	{types.SINT, types.LINT}:  {types.LINT, nil},
	{types.SINT, types.USINT}: {types.ANY, errors.New("no implicit cast between SINT and USINT")},
	{types.SINT, types.UINT}:  {types.ANY, errors.New("no implicit cast between SINT and UINT")},
	{types.SINT, types.UDINT}: {types.ANY, errors.New("no implicit cast between SINT and UDINT")},
	{types.SINT, types.ULINT}: {types.ANY, errors.New("no implicit cast between SINT and ULINT")},
	{types.SINT, types.REAL}:  {types.REAL, nil},
	{types.SINT, types.LREAL}: {types.LREAL, nil},

	{types.INT, types.INT}:   {types.INT, nil},
	{types.INT, types.DINT}:  {types.DINT, nil},
	{types.INT, types.LINT}:  {types.LINT, nil},
	{types.INT, types.USINT}: {types.INT, nil},
	{types.INT, types.UINT}:  {types.ANY, errors.New("no implicit cast between INT and UINT")},
	{types.INT, types.UDINT}: {types.ANY, errors.New("no implicit cast between INT and UDINT")},
	{types.INT, types.ULINT}: {types.ANY, errors.New("no implicit cast between INT and ULINT")},
	{types.INT, types.REAL}:  {types.REAL, nil},
	{types.INT, types.LREAL}: {types.LREAL, nil},

	{types.DINT, types.DINT}:  {types.DINT, nil},
	{types.DINT, types.LINT}:  {types.LINT, nil},
	{types.DINT, types.USINT}: {types.DINT, nil},
	{types.DINT, types.UINT}:  {types.DINT, nil},
	{types.DINT, types.UDINT}: {types.ANY, errors.New("no implicit cast between DINT and UDINT")},
	{types.DINT, types.ULINT}: {types.ANY, errors.New("no implicit cast between DINT and ULINT")},
	{types.DINT, types.REAL}:  {types.ANY, errors.New("no implicit cast between DINT and REAL")},
	{types.DINT, types.LREAL}: {types.LREAL, nil},

	{types.LINT, types.LINT}:  {types.LINT, nil},
	{types.LINT, types.USINT}: {types.LINT, nil},
	{types.LINT, types.UINT}:  {types.LINT, nil},
	{types.LINT, types.UDINT}: {types.LINT, nil},
	{types.LINT, types.ULINT}: {types.ANY, errors.New("no implicit cast between LINT and ULINT")},
	{types.LINT, types.REAL}:  {types.ANY, errors.New("no implicit cast between LINT and REAL")},
	{types.LINT, types.LREAL}: {types.ANY, errors.New("no implicit cast between LINT and LREAL")},

	{types.USINT, types.USINT}: {types.USINT, nil},
	{types.USINT, types.UINT}:  {types.UINT, nil},
	{types.USINT, types.UDINT}: {types.UDINT, nil},
	{types.USINT, types.ULINT}: {types.ULINT, nil},
	{types.USINT, types.REAL}:  {types.REAL, nil},
	{types.USINT, types.LREAL}: {types.LREAL, nil},

	{types.UINT, types.UINT}:  {types.UINT, nil},
	{types.UINT, types.UDINT}: {types.UDINT, nil},
	{types.UINT, types.ULINT}: {types.ULINT, nil},
	{types.UINT, types.REAL}:  {types.REAL, nil},
	{types.UINT, types.LREAL}: {types.LREAL, nil},

	{types.UDINT, types.UDINT}: {types.UDINT, nil},
	{types.UDINT, types.ULINT}: {types.ULINT, nil},
	{types.UDINT, types.REAL}:  {types.ANY, errors.New("no implicit cast between UDINT and REAL")},
	{types.UDINT, types.LREAL}: {types.LREAL, nil},

	{types.ULINT, types.ULINT}: {types.ULINT, nil},
	{types.ULINT, types.REAL}:  {types.ANY, errors.New("no implicit cast between UDINT and REAL")},
	{types.ULINT, types.LREAL}: {types.ANY, errors.New("no implicit cast between UDINT and LREAL")},

	{types.REAL, types.REAL}:  {types.REAL, nil},
	{types.REAL, types.LREAL}: {types.LREAL, nil},

	{types.LREAL, types.LREAL}: {types.LREAL, nil},
}

type assignTypes struct {
	v, expr types.Basic
}

var implicitCastInAssign = map[assignTypes]error{
	{types.SINT, types.SINT}:  nil,
	{types.SINT, types.INT}:   errors.New("no implicit cast from INT to SINT"),
	{types.SINT, types.DINT}:  errors.New("no implicit cast from DINT to SINT"),
	{types.SINT, types.LINT}:  errors.New("no implicit cast from LINT to SINT"),
	{types.SINT, types.USINT}: errors.New("no implicit cast from USINT to SINT"),
	{types.SINT, types.UINT}:  errors.New("no implicit cast from UINT to SINT"),
	{types.SINT, types.UDINT}: errors.New("no implicit cast from UDINT to SINT"),
	{types.SINT, types.ULINT}: errors.New("no implicit cast from ULINT to SINT"),
	{types.SINT, types.REAL}:  errors.New("no implicit cast from REAL to SINT"),
	{types.SINT, types.LREAL}: errors.New("no implicit cast from LREAL to SINT"),

	{types.INT, types.SINT}:  nil,
	{types.INT, types.INT}:   nil,
	{types.INT, types.DINT}:  errors.New("no implicit cast from DINT to INT"),
	{types.INT, types.LINT}:  errors.New("no implicit cast from LINT to INT"),
	{types.INT, types.USINT}: nil,
	{types.INT, types.UINT}:  errors.New("no implicit cast from UINT to INT"),
	{types.INT, types.UDINT}: errors.New("no implicit cast from UDINT to INT"),
	{types.INT, types.ULINT}: errors.New("no implicit cast from ULINT to INT"),
	{types.INT, types.REAL}:  errors.New("no implicit cast from REAL to INT"),
	{types.INT, types.LREAL}: errors.New("no implicit cast from LREAL to INT"),

	{types.DINT, types.SINT}:  nil,
	{types.DINT, types.INT}:   nil,
	{types.DINT, types.DINT}:  nil,
	{types.DINT, types.LINT}:  errors.New("no implicit cast from LINT to DINT"),
	{types.DINT, types.USINT}: nil,
	{types.DINT, types.UINT}:  nil,
	{types.DINT, types.UDINT}: errors.New("no implicit cast from UDINT to DINT"),
	{types.DINT, types.ULINT}: errors.New("no implicit cast from ULINT to DINT"),
	{types.DINT, types.REAL}:  errors.New("no implicit cast from REAL to DINT"),
	{types.DINT, types.LREAL}: nil,
	{types.DINT, types.REAL}:  errors.New("no implicit cast from REAL to DINT"),
	{types.DINT, types.LREAL}: errors.New("no implicit cast from LREAL to DINT"),

	{types.LINT, types.SINT}:  nil,
	{types.LINT, types.INT}:   nil,
	{types.LINT, types.DINT}:  nil,
	{types.LINT, types.LINT}:  nil,
	{types.LINT, types.USINT}: nil,
	{types.LINT, types.UINT}:  nil,
	{types.LINT, types.UDINT}: nil,
	{types.LINT, types.ULINT}: errors.New("no implicit cast from ULINT to LINT"),
	{types.LINT, types.REAL}:  errors.New("no implicit cast from REAL to LINT"),
	{types.LINT, types.LREAL}: errors.New("no implicit cast from LREAL to LINT"),

	{types.USINT, types.SINT}:  errors.New("no implicit cast from SINT to USINT"),
	{types.USINT, types.INT}:   errors.New("no implicit cast from INT to USINT"),
	{types.USINT, types.DINT}:  errors.New("no implicit cast from DINT to USINT"),
	{types.USINT, types.LINT}:  errors.New("no implicit cast from LINT to USINT"),
	{types.USINT, types.USINT}: nil,
	{types.USINT, types.UINT}:  errors.New("no implicit cast from UINT to USINT"),
	{types.USINT, types.UDINT}: errors.New("no implicit cast from UDINT to USINT"),
	{types.USINT, types.ULINT}: errors.New("no implicit cast from ULINT to USINT"),
	{types.USINT, types.REAL}:  errors.New("no implicit cast from REAL to USINT"),
	{types.USINT, types.LREAL}: errors.New("no implicit cast from LREAL to USINT"),

	{types.UINT, types.SINT}:  errors.New("no implicit cast from SINT to UINT"),
	{types.UINT, types.INT}:   errors.New("no implicit cast from INT to UINT"),
	{types.UINT, types.DINT}:  errors.New("no implicit cast from DINT to UINT"),
	{types.UINT, types.LINT}:  errors.New("no implicit cast from LINT to UINT"),
	{types.UINT, types.USINT}: nil,
	{types.UINT, types.UINT}:  nil,
	{types.UINT, types.UDINT}: errors.New("no implicit cast from UDINT to UINT"),
	{types.UINT, types.ULINT}: errors.New("no implicit cast from ULINT to UINT"),
	{types.UINT, types.REAL}:  errors.New("no implicit cast from REAL to UINT"),
	{types.UINT, types.LREAL}: errors.New("no implicit cast from LREAL to UINT"),

	{types.UDINT, types.SINT}:  errors.New("no implicit cast from SINT to UDINT"),
	{types.UDINT, types.INT}:   errors.New("no implicit cast from INT to UDINT"),
	{types.UDINT, types.DINT}:  errors.New("no implicit cast from DINT to UDINT"),
	{types.UDINT, types.LINT}:  errors.New("no implicit cast from LINT to UDINT"),
	{types.UDINT, types.USINT}: nil,
	{types.UDINT, types.UINT}:  nil,
	{types.UDINT, types.UDINT}: nil,
	{types.UDINT, types.ULINT}: errors.New("no implicit cast from ULINT to UDINT"),
	{types.UDINT, types.REAL}:  errors.New("no implicit cast from REAL to UDINT"),
	{types.UDINT, types.LREAL}: errors.New("no implicit cast from LREAL to UDINT"),

	{types.ULINT, types.SINT}:  errors.New("no implicit cast from SINT to ULINT"),
	{types.ULINT, types.INT}:   errors.New("no implicit cast from INT to ULINT"),
	{types.ULINT, types.DINT}:  errors.New("no implicit cast from DINT to ULINT"),
	{types.ULINT, types.LINT}:  errors.New("no implicit cast from LINT to ULINT"),
	{types.ULINT, types.USINT}: nil,
	{types.ULINT, types.UINT}:  nil,
	{types.ULINT, types.UDINT}: nil,
	{types.ULINT, types.ULINT}: nil,
	{types.ULINT, types.REAL}:  errors.New("no implicit cast from REAL to ULINT"),
	{types.ULINT, types.LREAL}: errors.New("no implicit cast from LREAL to ULINT"),

	{types.REAL, types.SINT}:  nil,
	{types.REAL, types.INT}:   nil,
	{types.REAL, types.DINT}:  errors.New("no implicit cast from DINT to REAL"),
	{types.REAL, types.LINT}:  errors.New("no implicit cast from LINT to REAL"),
	{types.REAL, types.USINT}: nil,
	{types.REAL, types.UINT}:  nil,
	{types.REAL, types.UDINT}: errors.New("no implicit cast from UDINT to REAL"),
	{types.REAL, types.ULINT}: errors.New("no implicit cast from ULINT to REAL"),
	{types.REAL, types.REAL}:  nil,
	{types.REAL, types.LREAL}: errors.New("no implicit cast from LREAL to REAL"),

	{types.LREAL, types.SINT}:  nil,
	{types.LREAL, types.INT}:   nil,
	{types.LREAL, types.DINT}:  nil,
	{types.LREAL, types.LINT}:  errors.New("no implicit cast from LINT to LREAL"),
	{types.LREAL, types.USINT}: nil,
	{types.LREAL, types.UINT}:  nil,
	{types.LREAL, types.UDINT}: nil,
	{types.LREAL, types.ULINT}: errors.New("no implicit cast from ULINT to LREAL"),
	{types.LREAL, types.REAL}:  nil,
	{types.LREAL, types.LREAL}: nil,
}
