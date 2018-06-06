package internal

import "reflect"

// check if element is nil
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
