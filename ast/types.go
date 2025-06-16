package ast

import (
	"errors"
	"slices"

	"github.com/slonegd/go-st/variant"
)

func validateTypes(op1, op2 Int64Operator) (variant.Type, error) {
	switch {
	case op1.isConstant && op2.isConstant:
		return op1.resultType, nil
	case op1.isConstant:
		return op2.resultType, nil
	case op2.isConstant:
		return op1.resultType, nil
	}
	types := [2]variant.Type{op1.resultType, op2.resultType}
	slices.Sort(types[:])
	r := implicitCastInBinary[types]
	return r.Type, r.error
}

// Figure 11 – Data type conversion rules – implicit and/or explicit (Summary) 61131-3 © IEC:2013
var implicitCastInBinary = map[[2]variant.Type]struct {
	variant.Type
	error
}{
	{variant.SINT, variant.SINT}:  {variant.SINT, nil},
	{variant.SINT, variant.INT}:   {variant.INT, nil},
	{variant.SINT, variant.DINT}:  {variant.DINT, nil},
	{variant.SINT, variant.LINT}:  {variant.LINT, nil},
	{variant.SINT, variant.USINT}: {variant.ANY, errors.New("no implicit cast between SINT and USINT")},
	{variant.SINT, variant.UINT}:  {variant.ANY, errors.New("no implicit cast between SINT and UINT")},
	{variant.SINT, variant.UDINT}: {variant.ANY, errors.New("no implicit cast between SINT and UDINT")},
	{variant.SINT, variant.ULINT}: {variant.ANY, errors.New("no implicit cast between SINT and ULINT")},

	{variant.INT, variant.INT}:   {variant.INT, nil},
	{variant.INT, variant.DINT}:  {variant.DINT, nil},
	{variant.INT, variant.LINT}:  {variant.LINT, nil},
	{variant.INT, variant.USINT}: {variant.INT, nil},
	{variant.INT, variant.UINT}:  {variant.ANY, errors.New("no implicit cast between INT and UINT")},
	{variant.INT, variant.UDINT}: {variant.ANY, errors.New("no implicit cast between INT and UDINT")},
	{variant.INT, variant.ULINT}: {variant.ANY, errors.New("no implicit cast between INT and ULINT")},

	{variant.DINT, variant.DINT}:  {variant.DINT, nil},
	{variant.DINT, variant.LINT}:  {variant.LINT, nil},
	{variant.DINT, variant.USINT}: {variant.DINT, nil},
	{variant.DINT, variant.UINT}:  {variant.DINT, nil},
	{variant.DINT, variant.UDINT}: {variant.ANY, errors.New("no implicit cast between DINT and UDINT")},
	{variant.DINT, variant.ULINT}: {variant.ANY, errors.New("no implicit cast between DINT and ULINT")},

	{variant.LINT, variant.LINT}:  {variant.LINT, nil},
	{variant.LINT, variant.USINT}: {variant.LINT, nil},
	{variant.LINT, variant.UINT}:  {variant.LINT, nil},
	{variant.LINT, variant.UDINT}: {variant.LINT, nil},
	{variant.LINT, variant.ULINT}: {variant.ANY, errors.New("no implicit cast between LINT and ULINT")},

	{variant.USINT, variant.USINT}: {variant.USINT, nil},
	{variant.USINT, variant.UINT}:  {variant.UINT, nil},
	{variant.USINT, variant.UDINT}: {variant.UDINT, nil},
	{variant.USINT, variant.ULINT}: {variant.ULINT, nil},

	{variant.UINT, variant.UINT}:  {variant.UINT, nil},
	{variant.UINT, variant.UDINT}: {variant.UDINT, nil},
	{variant.UINT, variant.ULINT}: {variant.ULINT, nil},

	{variant.UDINT, variant.UDINT}: {variant.UDINT, nil},
	{variant.UDINT, variant.ULINT}: {variant.ULINT, nil},

	{variant.ULINT, variant.ULINT}: {variant.ULINT, nil},
}

type assignTypes struct {
	v, expr variant.Type
}

var implicitCastInAssign = map[assignTypes]error{
	{variant.SINT, variant.SINT}:  nil,
	{variant.SINT, variant.INT}:   errors.New("no implicit cast from INT to SINT"),
	{variant.SINT, variant.DINT}:  errors.New("no implicit cast from DINT to SINT"),
	{variant.SINT, variant.LINT}:  errors.New("no implicit cast from LINT to SINT"),
	{variant.SINT, variant.USINT}: errors.New("no implicit cast from USINT to SINT"),
	{variant.SINT, variant.UINT}:  errors.New("no implicit cast from UINT to SINT"),
	{variant.SINT, variant.UDINT}: errors.New("no implicit cast from UDINT to SINT"),
	{variant.SINT, variant.ULINT}: errors.New("no implicit cast from ULINT to SINT"),
	{variant.SINT, variant.REAL}:  errors.New("no implicit cast from REAL to SINT"),
	{variant.SINT, variant.LREAL}: errors.New("no implicit cast from LREAL to SINT"),

	{variant.INT, variant.SINT}:  nil,
	{variant.INT, variant.INT}:   nil,
	{variant.INT, variant.DINT}:  errors.New("no implicit cast from DINT to INT"),
	{variant.INT, variant.LINT}:  errors.New("no implicit cast from LINT to INT"),
	{variant.INT, variant.USINT}: nil,
	{variant.INT, variant.UINT}:  errors.New("no implicit cast from UINT to INT"),
	{variant.INT, variant.UDINT}: errors.New("no implicit cast from UDINT to INT"),
	{variant.INT, variant.ULINT}: errors.New("no implicit cast from ULINT to INT"),
	{variant.INT, variant.REAL}:  errors.New("no implicit cast from REAL to INT"),
	{variant.INT, variant.LREAL}: errors.New("no implicit cast from LREAL to INT"),

	{variant.DINT, variant.SINT}:  nil,
	{variant.DINT, variant.INT}:   nil,
	{variant.DINT, variant.DINT}:  nil,
	{variant.DINT, variant.LINT}:  errors.New("no implicit cast from LINT to DINT"),
	{variant.DINT, variant.USINT}: nil,
	{variant.DINT, variant.UINT}:  nil,
	{variant.DINT, variant.UDINT}: errors.New("no implicit cast from UDINT to DINT"),
	{variant.DINT, variant.ULINT}: errors.New("no implicit cast from ULINT to DINT"),
	{variant.DINT, variant.REAL}:  errors.New("no implicit cast from REAL to DINT"),
	{variant.DINT, variant.LREAL}: nil,
	{variant.DINT, variant.REAL}:  errors.New("no implicit cast from REAL to DINT"),
	{variant.DINT, variant.LREAL}: errors.New("no implicit cast from LREAL to DINT"),

	{variant.LINT, variant.SINT}:  nil,
	{variant.LINT, variant.INT}:   nil,
	{variant.LINT, variant.DINT}:  nil,
	{variant.LINT, variant.LINT}:  nil,
	{variant.LINT, variant.USINT}: nil,
	{variant.LINT, variant.UINT}:  nil,
	{variant.LINT, variant.UDINT}: nil,
	{variant.LINT, variant.ULINT}: errors.New("no implicit cast from ULINT to LINT"),
	{variant.LINT, variant.REAL}:  errors.New("no implicit cast from REAL to LINT"),
	{variant.LINT, variant.LREAL}: errors.New("no implicit cast from LREAL to LINT"),

	{variant.USINT, variant.SINT}:  errors.New("no implicit cast from SINT to USINT"),
	{variant.USINT, variant.INT}:   errors.New("no implicit cast from INT to USINT"),
	{variant.USINT, variant.DINT}:  errors.New("no implicit cast from DINT to USINT"),
	{variant.USINT, variant.LINT}:  errors.New("no implicit cast from LINT to USINT"),
	{variant.USINT, variant.USINT}: nil,
	{variant.USINT, variant.UINT}:  errors.New("no implicit cast from UINT to USINT"),
	{variant.USINT, variant.UDINT}: errors.New("no implicit cast from UDINT to USINT"),
	{variant.USINT, variant.ULINT}: errors.New("no implicit cast from ULINT to USINT"),
	{variant.USINT, variant.REAL}:  errors.New("no implicit cast from REAL to USINT"),
	{variant.USINT, variant.LREAL}: errors.New("no implicit cast from LREAL to USINT"),

	{variant.UINT, variant.SINT}:  errors.New("no implicit cast from SINT to UINT"),
	{variant.UINT, variant.INT}:   errors.New("no implicit cast from INT to UINT"),
	{variant.UINT, variant.DINT}:  errors.New("no implicit cast from DINT to UINT"),
	{variant.UINT, variant.LINT}:  errors.New("no implicit cast from LINT to UINT"),
	{variant.UINT, variant.USINT}: nil,
	{variant.UINT, variant.UINT}:  nil,
	{variant.UINT, variant.UDINT}: errors.New("no implicit cast from UDINT to UINT"),
	{variant.UINT, variant.ULINT}: errors.New("no implicit cast from ULINT to UINT"),
	{variant.UINT, variant.REAL}:  errors.New("no implicit cast from REAL to UINT"),
	{variant.UINT, variant.LREAL}: errors.New("no implicit cast from LREAL to UINT"),

	{variant.UDINT, variant.SINT}:  errors.New("no implicit cast from SINT to UDINT"),
	{variant.UDINT, variant.INT}:   errors.New("no implicit cast from INT to UDINT"),
	{variant.UDINT, variant.DINT}:  errors.New("no implicit cast from DINT to UDINT"),
	{variant.UDINT, variant.LINT}:  errors.New("no implicit cast from LINT to UDINT"),
	{variant.UDINT, variant.USINT}: nil,
	{variant.UDINT, variant.UINT}:  nil,
	{variant.UDINT, variant.UDINT}: nil,
	{variant.UDINT, variant.ULINT}: errors.New("no implicit cast from ULINT to UDINT"),
	{variant.UDINT, variant.REAL}:  errors.New("no implicit cast from REAL to UDINT"),
	{variant.UDINT, variant.LREAL}: errors.New("no implicit cast from LREAL to UDINT"),

	{variant.ULINT, variant.SINT}:  errors.New("no implicit cast from SINT to ULINT"),
	{variant.ULINT, variant.INT}:   errors.New("no implicit cast from INT to ULINT"),
	{variant.ULINT, variant.DINT}:  errors.New("no implicit cast from DINT to ULINT"),
	{variant.ULINT, variant.LINT}:  errors.New("no implicit cast from LINT to ULINT"),
	{variant.ULINT, variant.USINT}: nil,
	{variant.ULINT, variant.UINT}:  nil,
	{variant.ULINT, variant.UDINT}: nil,
	{variant.ULINT, variant.ULINT}: nil,
	{variant.ULINT, variant.REAL}:  errors.New("no implicit cast from REAL to ULINT"),
	{variant.ULINT, variant.LREAL}: errors.New("no implicit cast from LREAL to ULINT"),

	{variant.REAL, variant.SINT}:  nil,
	{variant.REAL, variant.INT}:   nil,
	{variant.REAL, variant.DINT}:  errors.New("no implicit cast from DINT to REAL"),
	{variant.REAL, variant.LINT}:  errors.New("no implicit cast from LINT to REAL"),
	{variant.REAL, variant.USINT}: nil,
	{variant.REAL, variant.UINT}:  nil,
	{variant.REAL, variant.UDINT}: errors.New("no implicit cast from UDINT to REAL"),
	{variant.REAL, variant.ULINT}: errors.New("no implicit cast from ULINT to REAL"),
	{variant.REAL, variant.REAL}:  nil,
	{variant.REAL, variant.LREAL}: errors.New("no implicit cast from LREAL to REAL"),

	{variant.LREAL, variant.SINT}:  nil,
	{variant.LREAL, variant.INT}:   nil,
	{variant.LREAL, variant.DINT}:  nil,
	{variant.LREAL, variant.LINT}:  errors.New("no implicit cast from LINT to LREAL"),
	{variant.LREAL, variant.USINT}: nil,
	{variant.LREAL, variant.UINT}:  nil,
	{variant.LREAL, variant.UDINT}: nil,
	{variant.LREAL, variant.ULINT}: errors.New("no implicit cast from ULINT to LREAL"),
	{variant.LREAL, variant.REAL}:  nil,
	{variant.LREAL, variant.LREAL}: nil,
}
