package metrics

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Created-1]
	_ = x[Updated-2]
	_ = x[Removed-3]
}

const _eventType_name = "CreatedUpdatedDeleted"

var _eventType_index = [...]uint8{0, 7, 14, 21}

func (i eventType) String() string {
	i -= 1
	if i >= eventType(len(_eventType_index)-1) {
		return "eventType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _eventType_name[_eventType_index[i]:_eventType_index[i+1]]
}
