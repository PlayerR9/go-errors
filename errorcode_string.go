// Code generated by "stringer -type=ErrorCode"; DO NOT EDIT.

package errors

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BadParameter-0]
	_ = x[InvalidUsage-1]
	_ = x[FailFix-2]
	_ = x[OperationFail-3]
	_ = x[NoSuchKey-4]
	_ = x[AssertFail-5]
}

const _ErrorCode_name = "BadParameterInvalidUsageFailFixOperationFailNoSuchKeyAssertFail"

var _ErrorCode_index = [...]uint8{0, 12, 24, 31, 44, 53, 63}

func (i ErrorCode) String() string {
	if i < 0 || i >= ErrorCode(len(_ErrorCode_index)-1) {
		return "ErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ErrorCode_name[_ErrorCode_index[i]:_ErrorCode_index[i+1]]
}
