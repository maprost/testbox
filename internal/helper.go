package internal

import "reflect"

// IsNil check if element is nil
func IsNil(act interface{}) bool {
	// general case
	if act == nil {
		return true
	}

	// special case for (Chan, Func, Interface, Map, Ptr, Slice)
	value := reflect.ValueOf(act)
	if value.Kind() >= reflect.Chan && value.Kind() <= reflect.Slice && value.IsNil() {
		return true
	}

	return false
}

// IsInt check if element is an int
func IsInt(act interface{}) bool {
	// general case
	if act == nil {
		return false
	}

	value := reflect.ValueOf(act)
	if value.Kind() == reflect.Int ||
		value.Kind() == reflect.Int64 ||
		value.Kind() == reflect.Int32 {
		return true
	}

	return false
}
