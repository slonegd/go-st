package types

import (
	"errors"
	"slices"
)

type Expression struct {
	IsConstant bool
	Type       Basic
}

func BinaryResult(left, right Expression) (Basic, error) {
	switch {
	case left.IsConstant && right.IsConstant:
		return left.Type, nil // TODO неверно, надо брать максимальный
	case left.IsConstant:
		return right.Type, nil
	case right.IsConstant:
		return left.Type, nil
	}
	types := [2]Basic{left.Type, right.Type}
	slices.Sort(types[:])
	r := implicitCastInBinary[types]
	return r.Basic, r.error
}

func CheckAssign(variable, expression Basic) error {
	return implicitCastInAssign[assignTypes{variable, expression}]
}

// Figure 11 – Data type conversion rules – implicit and/or explicit (Summary) 61131-3 © IEC:2013
// TODO может можно сделать enum по возрастанию?
var implicitCastInBinary = map[[2]Basic]struct {
	Basic
	error
}{
	{SINT, SINT}:  {SINT, nil},
	{SINT, INT}:   {INT, nil},
	{SINT, DINT}:  {DINT, nil},
	{SINT, LINT}:  {LINT, nil},
	{SINT, USINT}: {ANY, errors.New("no implicit cast between SINT and USINT")},
	{SINT, UINT}:  {ANY, errors.New("no implicit cast between SINT and UINT")},
	{SINT, UDINT}: {ANY, errors.New("no implicit cast between SINT and UDINT")},
	{SINT, ULINT}: {ANY, errors.New("no implicit cast between SINT and ULINT")},
	{SINT, REAL}:  {REAL, nil},
	{SINT, LREAL}: {LREAL, nil},

	{INT, INT}:   {INT, nil},
	{INT, DINT}:  {DINT, nil},
	{INT, LINT}:  {LINT, nil},
	{INT, USINT}: {INT, nil},
	{INT, UINT}:  {ANY, errors.New("no implicit cast between INT and UINT")},
	{INT, UDINT}: {ANY, errors.New("no implicit cast between INT and UDINT")},
	{INT, ULINT}: {ANY, errors.New("no implicit cast between INT and ULINT")},
	{INT, REAL}:  {REAL, nil},
	{INT, LREAL}: {LREAL, nil},

	{DINT, DINT}:  {DINT, nil},
	{DINT, LINT}:  {LINT, nil},
	{DINT, USINT}: {DINT, nil},
	{DINT, UINT}:  {DINT, nil},
	{DINT, UDINT}: {ANY, errors.New("no implicit cast between DINT and UDINT")},
	{DINT, ULINT}: {ANY, errors.New("no implicit cast between DINT and ULINT")},
	{DINT, REAL}:  {ANY, errors.New("no implicit cast between DINT and REAL")},
	{DINT, LREAL}: {LREAL, nil},

	{LINT, LINT}:  {LINT, nil},
	{LINT, USINT}: {LINT, nil},
	{LINT, UINT}:  {LINT, nil},
	{LINT, UDINT}: {LINT, nil},
	{LINT, ULINT}: {ANY, errors.New("no implicit cast between LINT and ULINT")},
	{LINT, REAL}:  {ANY, errors.New("no implicit cast between LINT and REAL")},
	{LINT, LREAL}: {ANY, errors.New("no implicit cast between LINT and LREAL")},

	{USINT, USINT}: {USINT, nil},
	{USINT, UINT}:  {UINT, nil},
	{USINT, UDINT}: {UDINT, nil},
	{USINT, ULINT}: {ULINT, nil},
	{USINT, REAL}:  {REAL, nil},
	{USINT, LREAL}: {LREAL, nil},

	{UINT, UINT}:  {UINT, nil},
	{UINT, UDINT}: {UDINT, nil},
	{UINT, ULINT}: {ULINT, nil},
	{UINT, REAL}:  {REAL, nil},
	{UINT, LREAL}: {LREAL, nil},

	{UDINT, UDINT}: {UDINT, nil},
	{UDINT, ULINT}: {ULINT, nil},
	{UDINT, REAL}:  {ANY, errors.New("no implicit cast between UDINT and REAL")},
	{UDINT, LREAL}: {LREAL, nil},

	{ULINT, ULINT}: {ULINT, nil},
	{ULINT, REAL}:  {ANY, errors.New("no implicit cast between UDINT and REAL")},
	{ULINT, LREAL}: {ANY, errors.New("no implicit cast between UDINT and LREAL")},

	{REAL, REAL}:  {REAL, nil},
	{REAL, LREAL}: {LREAL, nil},

	{LREAL, LREAL}: {LREAL, nil},
}

type assignTypes struct {
	v, expr Basic
}

var implicitCastInAssign = map[assignTypes]error{
	{SINT, SINT}:  nil,
	{SINT, INT}:   errors.New("no implicit cast from INT to SINT"),
	{SINT, DINT}:  errors.New("no implicit cast from DINT to SINT"),
	{SINT, LINT}:  errors.New("no implicit cast from LINT to SINT"),
	{SINT, USINT}: errors.New("no implicit cast from USINT to SINT"),
	{SINT, UINT}:  errors.New("no implicit cast from UINT to SINT"),
	{SINT, UDINT}: errors.New("no implicit cast from UDINT to SINT"),
	{SINT, ULINT}: errors.New("no implicit cast from ULINT to SINT"),
	{SINT, REAL}:  errors.New("no implicit cast from REAL to SINT"),
	{SINT, LREAL}: errors.New("no implicit cast from LREAL to SINT"),

	{INT, SINT}:  nil,
	{INT, INT}:   nil,
	{INT, DINT}:  errors.New("no implicit cast from DINT to INT"),
	{INT, LINT}:  errors.New("no implicit cast from LINT to INT"),
	{INT, USINT}: nil,
	{INT, UINT}:  errors.New("no implicit cast from UINT to INT"),
	{INT, UDINT}: errors.New("no implicit cast from UDINT to INT"),
	{INT, ULINT}: errors.New("no implicit cast from ULINT to INT"),
	{INT, REAL}:  errors.New("no implicit cast from REAL to INT"),
	{INT, LREAL}: errors.New("no implicit cast from LREAL to INT"),

	{DINT, SINT}:  nil,
	{DINT, INT}:   nil,
	{DINT, DINT}:  nil,
	{DINT, LINT}:  errors.New("no implicit cast from LINT to DINT"),
	{DINT, USINT}: nil,
	{DINT, UINT}:  nil,
	{DINT, UDINT}: errors.New("no implicit cast from UDINT to DINT"),
	{DINT, ULINT}: errors.New("no implicit cast from ULINT to DINT"),
	{DINT, REAL}:  errors.New("no implicit cast from REAL to DINT"),
	{DINT, LREAL}: nil,
	{DINT, REAL}:  errors.New("no implicit cast from REAL to DINT"),
	{DINT, LREAL}: errors.New("no implicit cast from LREAL to DINT"),

	{LINT, SINT}:  nil,
	{LINT, INT}:   nil,
	{LINT, DINT}:  nil,
	{LINT, LINT}:  nil,
	{LINT, USINT}: nil,
	{LINT, UINT}:  nil,
	{LINT, UDINT}: nil,
	{LINT, ULINT}: errors.New("no implicit cast from ULINT to LINT"),
	{LINT, REAL}:  errors.New("no implicit cast from REAL to LINT"),
	{LINT, LREAL}: errors.New("no implicit cast from LREAL to LINT"),

	{USINT, SINT}:  errors.New("no implicit cast from SINT to USINT"),
	{USINT, INT}:   errors.New("no implicit cast from INT to USINT"),
	{USINT, DINT}:  errors.New("no implicit cast from DINT to USINT"),
	{USINT, LINT}:  errors.New("no implicit cast from LINT to USINT"),
	{USINT, USINT}: nil,
	{USINT, UINT}:  errors.New("no implicit cast from UINT to USINT"),
	{USINT, UDINT}: errors.New("no implicit cast from UDINT to USINT"),
	{USINT, ULINT}: errors.New("no implicit cast from ULINT to USINT"),
	{USINT, REAL}:  errors.New("no implicit cast from REAL to USINT"),
	{USINT, LREAL}: errors.New("no implicit cast from LREAL to USINT"),

	{UINT, SINT}:  errors.New("no implicit cast from SINT to UINT"),
	{UINT, INT}:   errors.New("no implicit cast from INT to UINT"),
	{UINT, DINT}:  errors.New("no implicit cast from DINT to UINT"),
	{UINT, LINT}:  errors.New("no implicit cast from LINT to UINT"),
	{UINT, USINT}: nil,
	{UINT, UINT}:  nil,
	{UINT, UDINT}: errors.New("no implicit cast from UDINT to UINT"),
	{UINT, ULINT}: errors.New("no implicit cast from ULINT to UINT"),
	{UINT, REAL}:  errors.New("no implicit cast from REAL to UINT"),
	{UINT, LREAL}: errors.New("no implicit cast from LREAL to UINT"),

	{UDINT, SINT}:  errors.New("no implicit cast from SINT to UDINT"),
	{UDINT, INT}:   errors.New("no implicit cast from INT to UDINT"),
	{UDINT, DINT}:  errors.New("no implicit cast from DINT to UDINT"),
	{UDINT, LINT}:  errors.New("no implicit cast from LINT to UDINT"),
	{UDINT, USINT}: nil,
	{UDINT, UINT}:  nil,
	{UDINT, UDINT}: nil,
	{UDINT, ULINT}: errors.New("no implicit cast from ULINT to UDINT"),
	{UDINT, REAL}:  errors.New("no implicit cast from REAL to UDINT"),
	{UDINT, LREAL}: errors.New("no implicit cast from LREAL to UDINT"),

	{ULINT, SINT}:  errors.New("no implicit cast from SINT to ULINT"),
	{ULINT, INT}:   errors.New("no implicit cast from INT to ULINT"),
	{ULINT, DINT}:  errors.New("no implicit cast from DINT to ULINT"),
	{ULINT, LINT}:  errors.New("no implicit cast from LINT to ULINT"),
	{ULINT, USINT}: nil,
	{ULINT, UINT}:  nil,
	{ULINT, UDINT}: nil,
	{ULINT, ULINT}: nil,
	{ULINT, REAL}:  errors.New("no implicit cast from REAL to ULINT"),
	{ULINT, LREAL}: errors.New("no implicit cast from LREAL to ULINT"),

	{REAL, SINT}:  nil,
	{REAL, INT}:   nil,
	{REAL, DINT}:  errors.New("no implicit cast from DINT to REAL"),
	{REAL, LINT}:  errors.New("no implicit cast from LINT to REAL"),
	{REAL, USINT}: nil,
	{REAL, UINT}:  nil,
	{REAL, UDINT}: errors.New("no implicit cast from UDINT to REAL"),
	{REAL, ULINT}: errors.New("no implicit cast from ULINT to REAL"),
	{REAL, REAL}:  nil,
	{REAL, LREAL}: errors.New("no implicit cast from LREAL to REAL"),

	{LREAL, SINT}:  nil,
	{LREAL, INT}:   nil,
	{LREAL, DINT}:  nil,
	{LREAL, LINT}:  errors.New("no implicit cast from LINT to LREAL"),
	{LREAL, USINT}: nil,
	{LREAL, UINT}:  nil,
	{LREAL, UDINT}: nil,
	{LREAL, ULINT}: errors.New("no implicit cast from ULINT to LREAL"),
	{LREAL, REAL}:  nil,
	{LREAL, LREAL}: nil,
}
