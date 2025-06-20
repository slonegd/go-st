package types

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"unsafe"
)

type (
	Basic    int
	Variable interface {
		Type() Basic
		Bool() bool
		Int() int64
		Uint() uint64
		Float64() float64
		ToString() string // преобразование
		String() string   // для логирования
		Pointer() unsafe.Pointer
	}
)

func TypeFromString(v string) Basic {
	for t := range SIZE {
		if v == t.String() {
			return t
		}
	}
	return ANY
}

//go:generate go run .././vendor/golang.org/x/tools/cmd/stringer/stringer.go -type=Type
const (
	ANY   Basic = iota
	BOOL        // bool
	SINT        // int8
	INT         // int16
	DINT        // int32
	LINT        // int64
	USINT       // uint8
	UINT        // uint16
	UDINT       // uint32
	ULINT       // uint64
	REAL        // float32
	LREAL       // float64
	TIME        // time.Duration
	// LTIME      // time.Duration
	TOD // int секунды с начала дня
	// LTOD // int секунды с начала дня
	DT // time.Time (строковое представление в time.Local)
	// LDT // time.Time (строковое представление в time.Local)
	STRING // string
	// WSTRING // string
	CHAR  // rune
	WCHAR // rune
	BYTE  // uint8 строка бит
	WORD  // uint16 строка бит
	DWORD // uint32 строка бит
	LWORD // uint64 строка бит

	SIZE // для цикла по перечисленю
)

func (t Basic) IsInteger() bool {
	switch t {
	case SINT, INT, DINT, LINT, USINT, UINT, UDINT, ULINT:
		return true
	}
	return false
}
func (t Basic) IsFloat() bool {
	return t == REAL || t == LREAL
}

var (
	_ Variable = (*Int)(nil) // все инты, метки времени и длительность
	_ Variable = (*Bool)(nil)
	_ Variable = (*String)(nil)  // STRING, WSTRING
	_ Variable = (*Float64)(nil) // REAL, LREAL
	_ Variable = (*Any)(nil)
)

type Int struct {
	v int64
	t Basic
}

func IntVariable(v int64) *Int { return &Int{v, INT} }

func (x *Int) Type() Basic             { return x.t }
func (x *Int) Bool() bool              { return x.v != 0 }
func (x *Int) Int() int64              { return x.v }
func (x *Int) Uint() uint64            { return uint64(x.v) }
func (x *Int) Float64() float64        { return float64(x.v) }
func (x *Int) String() string          { return fmt.Sprintf("%s(%d)", x.t, x.v) }
func (x *Int) ToString() string        { return fmt.Sprintf("%d", x.v) }
func (x *Int) Pointer() unsafe.Pointer { return unsafe.Pointer(&x.v) }

type Bool struct{ v bool }

func BoolVariable(v bool) *Bool { return &Bool{v} }

func (x *Bool) Type() Basic { return BOOL }
func (x *Bool) Bool() bool  { return x.v }
func (x *Bool) Int() int64 {
	if x.v {
		return 1
	}
	return 0
}
func (x *Bool) Uint() uint64            { return uint64(x.Int()) }
func (x *Bool) Float64() float64        { return float64(x.Int()) }
func (x *Bool) String() string          { return fmt.Sprintf("BOOL(%t)", x.v) }
func (x *Bool) ToString() string        { return fmt.Sprintf("%t", x.v) }
func (x *Bool) Pointer() unsafe.Pointer { return unsafe.Pointer(&x.v) }

type String struct{ v string }

func (x *String) Type() Basic             { return STRING }
func (x *String) Bool() bool              { r, _ := strconv.ParseBool(x.v); return r }
func (x *String) Int() int64              { r, _ := strconv.ParseInt(x.v, 10, 64); return r }
func (x *String) Uint() uint64            { return uint64(x.Int()) }
func (x *String) Float64() float64        { r, _ := strconv.ParseFloat(x.v, 64); return r }
func (x *String) String() string          { return fmt.Sprintf("STRING(%q)", x.v) }
func (x *String) ToString() string        { return x.v }
func (x *String) Pointer() unsafe.Pointer { return unsafe.Pointer(&x.v) }

type Float64 struct {
	v float64
	t Basic
}

func Float64Variable(v float64) *Float64   { return &Float64{v: v, t: LREAL} }
func (x *Float64) Type() Basic             { return x.t }
func (x *Float64) Bool() bool              { return x.v != 0 }
func (x *Float64) Int() int64              { return int64(x.v) }
func (x *Float64) Uint() uint64            { return uint64(x.Int()) }
func (x *Float64) Float64() float64        { return x.v }
func (x *Float64) String() string          { return fmt.Sprintf("%s(%f)", x.t, x.v) }
func (x *Float64) ToString() string        { return fmt.Sprintf("%f", x.v) }
func (x *Float64) Pointer() unsafe.Pointer { return unsafe.Pointer(&x.v) }

type Any struct{ Variable }

func NewAnyVariable(v string) *Any {
	r := &Any{}

	i, err := strconv.ParseInt(v, 10, 64)
	if err == nil {
		r.Variable = &Int{v: i}
		return r
	}

	f, err := strconv.ParseFloat(v, 64)
	if err == nil {
		r.Variable = &Float64{v: f}
		return r
	}

	b, err := strconv.ParseBool(v)
	if err == nil {
		r.Variable = &Bool{v: b}
		return r
	}

	if v, ok := strings.CutPrefix(v, "16#"); ok {
		v = strings.ReplaceAll(v, "_", "")
		bint, ok := big.NewInt(0).SetString(v, 16)
		if ok {
			r.Variable = &Int{v: bint.Int64()}
			return r
		}
	}

	r.Variable = &String{v: v}
	return r
}

func (x *Any) Type() Basic    { return ANY }
func (x *Any) String() string { return fmt.Sprintf("ANY_%s", x.Variable) }

func SetType(v Variable, t Basic) Variable {
	switch t {
	case BOOL:
		return &Bool{v.Bool()}
	case SINT, INT, DINT, LINT, USINT, UINT, UDINT, ULINT:
		return &Int{v: v.Int(), t: t}
	case REAL, LREAL:
		return &Float64{v: v.Float64(), t: t}
	default:
		return &String{v.ToString()}
	}
}
