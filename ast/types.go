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
	r := implicitCast[types]
	return r.Type, r.error
}

// Figure 11 – Data type conversion rules – implicit and/or explicit (Summary) 61131-3 © IEC:2013
var implicitCast = map[[2]variant.Type]struct {
	variant.Type
	error
}{
	{variant.SINT, variant.SINT}:  {variant.SINT, nil},
	{variant.SINT, variant.INT}:   {variant.INT, nil},
	{variant.SINT, variant.DINT}:  {variant.DINT, nil},
	{variant.SINT, variant.LINT}:  {variant.LINT, nil},
	{variant.SINT, variant.USINT}: {variant.ANY, errors.New("no implicit cast beetween SINT and USINT")},
	{variant.SINT, variant.UINT}:  {variant.ANY, errors.New("no implicit cast beetween SINT and UINT")},
	{variant.SINT, variant.UDINT}: {variant.ANY, errors.New("no implicit cast beetween SINT and UDINT")},
	{variant.SINT, variant.ULINT}: {variant.ANY, errors.New("no implicit cast beetween SINT and ULINT")},

	{variant.INT, variant.INT}:   {variant.INT, nil},
	{variant.INT, variant.DINT}:  {variant.DINT, nil},
	{variant.INT, variant.LINT}:  {variant.LINT, nil},
	{variant.INT, variant.USINT}: {variant.INT, nil},
	{variant.INT, variant.UINT}:  {variant.ANY, errors.New("no implicit cast beetween INT and UINT")},
	{variant.INT, variant.UDINT}: {variant.ANY, errors.New("no implicit cast beetween INT and UDINT")},
	{variant.INT, variant.ULINT}: {variant.ANY, errors.New("no implicit cast beetween INT and ULINT")},

	{variant.DINT, variant.DINT}:  {variant.DINT, nil},
	{variant.DINT, variant.LINT}:  {variant.LINT, nil},
	{variant.DINT, variant.USINT}: {variant.DINT, nil},
	{variant.DINT, variant.UINT}:  {variant.DINT, nil},
	{variant.DINT, variant.UDINT}: {variant.ANY, errors.New("no implicit cast beetween DINT and UDINT")},
	{variant.DINT, variant.ULINT}: {variant.ANY, errors.New("no implicit cast beetween DINT and ULINT")},

	{variant.LINT, variant.LINT}:  {variant.LINT, nil},
	{variant.LINT, variant.USINT}: {variant.LINT, nil},
	{variant.LINT, variant.UINT}:  {variant.LINT, nil},
	{variant.LINT, variant.UDINT}: {variant.LINT, nil},
	{variant.LINT, variant.ULINT}: {variant.ANY, errors.New("no implicit cast beetween LINT and ULINT")},

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
