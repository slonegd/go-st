package variant

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"unsafe"
)

type (
	Type    int
	Variant interface {
		Type() Type
		SetValue(Variant)
		Bool() bool
		Int() int64
		Uint() uint64
		Float64() float64
		ToString() string // преобразование
		String() string   // для логирования
		Pointer() unsafe.Pointer
	}
)

func TypeFromString(v string) Type {
	for t := range SIZE {
		if v == t.String() {
			return t
		}
	}
	return ANY
}

//go:generate go run .././vendor/golang.org/x/tools/cmd/stringer/stringer.go -type=Type
const (
	ANY   Type = iota
	BOOL       // bool
	SINT       // int8
	INT        // int16
	DINT       // int32
	LINT       // int64
	USINT      // uint8
	UINT       // uint16
	UDINT      // uint32
	ULINT      // uint64
	REAL       // float32
	LREAL      // float64
	TIME       // time.Duration
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

var (
	_ Variant = (*Int)(nil) // все инты, кроме uint64, метки времени и длительность
	_ Variant = (*Bool)(nil)
	_ Variant = (*String)(nil)  // STRING, WSTRING
	_ Variant = (*Float64)(nil) // REAL, LREAL
	_ Variant = (*Any)(nil)     // для констант, подберёться наиболее подходящий из строки
)

type Int struct {
	v int64
	t Type
}

func IntVariant(v int64) *Int { return &Int{v, INT} }

func (x *Int) Type() Type              { return x.t }
func (x *Int) SetValue(v Variant)      { x.v = v.Int() }
func (x *Int) Bool() bool              { return x.v != 0 }
func (x *Int) Int() int64              { return x.v }
func (x *Int) Uint() uint64            { return uint64(x.v) }
func (x *Int) Float64() float64        { return float64(x.v) }
func (x *Int) String() string          { return fmt.Sprintf("%s(%d)", x.t, x.v) }
func (x *Int) ToString() string        { return fmt.Sprintf("%d", x.v) }
func (x *Int) Pointer() unsafe.Pointer { return unsafe.Pointer(&x.v) }

type Bool struct{ v bool }

func BoolVariant(v bool) *Bool { return &Bool{v} }

func (x *Bool) Type() Type         { return BOOL }
func (x *Bool) SetValue(v Variant) { x.v = v.Bool() }
func (x *Bool) Bool() bool         { return x.v }
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

func (x *String) Type() Type              { return STRING }
func (x *String) SetValue(v Variant)      { x.v = v.ToString() }
func (x *String) Bool() bool              { r, _ := strconv.ParseBool(x.v); return r }
func (x *String) Int() int64              { r, _ := strconv.ParseInt(x.v, 10, 64); return r }
func (x *String) Uint() uint64            { return uint64(x.Int()) }
func (x *String) Float64() float64        { r, _ := strconv.ParseFloat(x.v, 64); return r }
func (x *String) String() string          { return fmt.Sprintf("STRING(%q)", x.v) }
func (x *String) ToString() string        { return x.v }
func (x *String) Pointer() unsafe.Pointer { return unsafe.Pointer(&x.v) }

type Float64 struct{ v float64 }

func (x *Float64) Type() Type              { return LREAL }
func (x *Float64) SetValue(v Variant)      { x.v = v.Float64() }
func (x *Float64) Bool() bool              { return x.v != 0 }
func (x *Float64) Int() int64              { return int64(x.v) }
func (x *Float64) Uint() uint64            { return uint64(x.Int()) }
func (x *Float64) Float64() float64        { return x.v }
func (x *Float64) String() string          { return fmt.Sprintf("REAL(%f)", x.v) }
func (x *Float64) ToString() string        { return fmt.Sprintf("%f", x.v) }
func (x *Float64) Pointer() unsafe.Pointer { return unsafe.Pointer(&x.v) }

type Any struct{ Variant }

func NewAnyVariant(v string) *Any {
	r := &Any{}

	i, err := strconv.ParseInt(v, 10, 64)
	if err == nil {
		r.Variant = &Int{v: i}
		return r
	}

	f, err := strconv.ParseFloat(v, 64)
	if err == nil {
		r.Variant = &Float64{v: f}
		return r
	}

	b, err := strconv.ParseBool(v)
	if err == nil {
		r.Variant = &Bool{v: b}
		return r
	}

	if v, ok := strings.CutPrefix(v, "16#"); ok {
		v = strings.ReplaceAll(v, "_", "")
		bint, ok := big.NewInt(0).SetString(v, 16)
		if ok {
			r.Variant = &Int{v: bint.Int64()}
			return r
		}
	}

	r.Variant = &String{v: v}
	return r
}

func (x *Any) Type() Type     { return ANY }
func (x *Any) String() string { return fmt.Sprintf("ANY_%s", x.Variant) }

func SetType(v Variant, t Type) Variant {
	switch t {
	case BOOL:
		return &Bool{v.Bool()}
	case SINT, INT, DINT, LINT, USINT, UINT, UDINT, ULINT:
		return &Int{v: v.Int(), t: t}
	case REAL, LREAL:
		return &Float64{v: v.Float64()}
	default:
		return &String{v.ToString()}
	}
}

func Plus(left Variant, right Variant) Variant {
	t := commonType(left, right)
	switch t {
	case BOOL:
		return left // операция не поддерживается
	case SINT, INT, DINT, LINT, USINT, UINT, UDINT:
		return &Int{v: left.Int() + right.Int(), t: t}
	case REAL, LREAL:
		return &Float64{v: left.Float64() + right.Float64()}
	default:
		return &String{left.ToString() + right.ToString()} // TODO такого нет в стандарте?
	}
}

func Sub(left Variant, right Variant) Variant {
	t := commonType(left, right)
	switch t {
	case BOOL:
		return left // операция не поддерживается
	case SINT, INT, DINT, LINT, USINT, UINT, UDINT:
		return &Int{v: left.Int() - right.Int(), t: t}
	case REAL, LREAL:
		return &Float64{v: left.Float64() - right.Float64()}
	default:
		return left // операция не поддерживается
	}
}

func Mult(left Variant, right Variant) Variant {
	t := commonType(left, right)
	switch t {
	case BOOL:
		return left // операция не поддерживается
	case SINT, INT, DINT, LINT, USINT, UINT, UDINT:
		return &Int{v: left.Int() * right.Int(), t: t}
	case REAL, LREAL:
		return &Float64{v: left.Float64() * right.Float64()}
	default:
		return left // операция не поддерживается
	}
}

// TODO деление на 0
func Divide(left Variant, right Variant) Variant {
	t := commonType(left, right)
	switch t {
	case BOOL:
		return left // операция не поддерживается
	case SINT, INT, DINT, LINT, USINT, UINT, UDINT:
		return &Int{v: left.Int() / right.Int(), t: t}
	case REAL, LREAL:
		return &Float64{v: left.Float64() / right.Float64()}
	default:
		return left // операция не поддерживается
	}
}

// TODO деление на 0
func Mod(left Variant, right Variant) Variant {
	t := commonType(left, right)
	switch t {
	case BOOL:
		return left // операция не поддерживается
	case SINT, INT, DINT, LINT, USINT, UINT, UDINT:
		return &Int{v: left.Int() % right.Int(), t: t}
	case REAL, LREAL:
		return left // операция не поддерживается
	default:
		return left // операция не поддерживается
	}
}

func Equal(left Variant, right Variant) Variant {
	t := commonType(left, right)
	switch t {
	case BOOL:
		return &Bool{v: left.Bool() == right.Bool()}
	case SINT, INT, DINT, LINT, USINT, UINT, UDINT:
		return &Bool{v: left.Int() == right.Int()}
	case REAL, LREAL:
		return &Bool{v: left.Float64() == right.Float64()}
	default:
		return &Bool{v: left.ToString() == right.ToString()}
	}
}

func NotEqual(left Variant, right Variant) Variant {
	return &Bool{v: !Equal(left, right).Bool()}
}

func Greather(left Variant, right Variant) Variant {
	t := commonType(left, right)
	switch t {
	case BOOL:
		return &Bool{v: left.Int() > right.Int()} // так проще сравнить TODO а как в стандарте?
	case SINT, INT, DINT, LINT, USINT, UINT, UDINT:
		return &Bool{v: left.Int() > right.Int()}
	case REAL, LREAL:
		return &Bool{v: left.Float64() > right.Float64()}
	default:
		return &Bool{v: left.ToString() > right.ToString()}
	}
}

func GreatherOrEqual(left Variant, right Variant) Variant {
	t := commonType(left, right)
	switch t {
	case BOOL:
		return &Bool{v: left.Int() >= right.Int()} // так проще сравнить TODO а как в стандарте?
	case SINT, INT, DINT, LINT, USINT, UINT, UDINT:
		return &Bool{v: left.Int() >= right.Int()}
	case REAL, LREAL:
		return &Bool{v: left.Float64() >= right.Float64()}
	default:
		return &Bool{v: left.ToString() >= right.ToString()}
	}
}

func Less(left Variant, right Variant) Variant {
	t := commonType(left, right)
	switch t {
	case BOOL:
		return &Bool{v: left.Int() < right.Int()} // так проще сравнить TODO а как в стандарте?
	case SINT, INT, DINT, LINT, USINT, UINT, UDINT:
		return &Bool{v: left.Int() < right.Int()}
	case REAL, LREAL:
		return &Bool{v: left.Float64() < right.Float64()}
	default:
		return &Bool{v: left.ToString() < right.ToString()}
	}
}

func LessOrEqual(left Variant, right Variant) Variant {
	t := commonType(left, right)
	switch t {
	case BOOL:
		return &Bool{v: left.Int() <= right.Int()} // так проще сравнить TODO а как в стандарте?
	case SINT, INT, DINT, LINT, USINT, UINT, UDINT:
		return &Bool{v: left.Int() <= right.Int()}
	case REAL, LREAL:
		return &Bool{v: left.Float64() <= right.Float64()}
	default:
		return &Bool{v: left.ToString() <= right.ToString()}
	}
}

func commonType(left Variant, right Variant) Type {
	t := left.Type()
	if t != ANY {
		return t
	}

	t = right.Type()
	if t != ANY {
		return t
	}

	return INT // TODO определить по ANY
}
