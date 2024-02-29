package fit

import "strconv"

// Manually generated types: The 'Bool' type is not found in the SDK
// specification, so it won't be auto-generated, but it is also not a base
// type.

type Bool byte

const (
	BoolFalse   Bool = 0
	BoolTrue    Bool = 1
	BoolInvalid Bool = 255
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BoolFalse-0]
	_ = x[BoolTrue-1]
	_ = x[BoolInvalid-255]
}

const (
	_Bool_name_0 = "BoolFalseBoolTrue"
	_Bool_name_1 = "BoolInvalid"
)

var _Bool_index_0 = [...]uint8{0, 9, 17}

func (i Bool) String() string {
	switch {
	//lint:ignore SA4003 Check unsigned >= 0, but this matches the stringer autogen output
	case 0 <= i && i <= 1:
		return _Bool_name_0[_Bool_index_0[i]:_Bool_index_0[i+1]]
	case i == 255:
		return _Bool_name_1
	default:
		return "Bool(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
