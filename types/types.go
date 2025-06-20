package types

type Basic int

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
