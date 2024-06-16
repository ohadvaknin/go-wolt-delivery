// Code generated by "stringer -type=Operation -trimprefix=Op"; DO NOT EDIT.

package filters

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OpInvalid-0]
	_ = x[OpNop-4294967294]
	_ = x[OpString-4294967293]
	_ = x[OpNot-4294967292]
	_ = x[OpAnd-4294967291]
	_ = x[OpOr-4294967290]
	_ = x[OpEq-4294967289]
	_ = x[OpNotEq-4294967288]
	_ = x[OpFunctionVarFunc-4294967287]
	_ = x[opLastBuiltin-4294967286]
}

const (
	_Operation_name_0 = "Invalid"
	_Operation_name_1 = "opLastBuiltinFunctionVarFuncNotEqEqOrAndNotStringNop"
)

var (
	_Operation_index_1 = [...]uint8{0, 13, 28, 33, 35, 37, 40, 43, 49, 52}
)

func (i Operation) String() string {
	switch {
	case i == 0:
		return _Operation_name_0
	case 4294967286 <= i && i <= 4294967294:
		i -= 4294967286
		return _Operation_name_1[_Operation_index_1[i]:_Operation_index_1[i+1]]
	default:
		return "Operation(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}