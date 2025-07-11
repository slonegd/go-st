// Code generated by "stringer -type=Basic"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ANY-0]
	_ = x[BOOL-1]
	_ = x[SINT-2]
	_ = x[INT-3]
	_ = x[DINT-4]
	_ = x[LINT-5]
	_ = x[USINT-6]
	_ = x[UINT-7]
	_ = x[UDINT-8]
	_ = x[ULINT-9]
	_ = x[REAL-10]
	_ = x[LREAL-11]
	_ = x[TIME-12]
	_ = x[TOD-13]
	_ = x[DT-14]
	_ = x[STRING-15]
	_ = x[WSTRING-16]
	_ = x[CHAR-17]
	_ = x[WCHAR-18]
	_ = x[BYTE-19]
	_ = x[WORD-20]
	_ = x[DWORD-21]
	_ = x[LWORD-22]
	_ = x[SIZE-23]
}

const _Basic_name = "ANYBOOLSINTINTDINTLINTUSINTUINTUDINTULINTREALLREALTIMETODDTSTRINGWSTRINGCHARWCHARBYTEWORDDWORDLWORDSIZE"

var _Basic_index = [...]uint8{0, 3, 7, 11, 14, 18, 22, 27, 31, 36, 41, 45, 50, 54, 57, 59, 65, 72, 76, 81, 85, 89, 94, 99, 103}

func (i Basic) String() string {
	if i < 0 || i >= Basic(len(_Basic_index)-1) {
		return "Basic(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Basic_name[_Basic_index[i]:_Basic_index[i+1]]
}
