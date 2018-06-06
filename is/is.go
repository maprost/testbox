package is

import (
	"reflect"
	"strings"

	"github.com/maprost/testbox/internal"
)

// Equal checks if 'act' == 'exp'
func Equal(act interface{}, exp interface{}) bool {
	ok, _ := Equalf(act, exp)
	return ok
}

// Equalf checks if 'act' == 'exp'
func Equalf(act interface{}, exp interface{}, msgArgs ...interface{}) (bool, error) {
	if !isEqual(act, exp) {
		return false, internal.Errorf(msgArgs, "Not equal:", internal.TypeValues(act, exp))
	}

	return true, nil
}

// NotEqual checks if 'act' != 'exp'
func NotEqual(act interface{}, exp interface{}) bool {
	ok, _ := NotEqualf(act, exp)
	return ok
}

// NotEqualf checks if 'act' != 'exp'
func NotEqualf(act interface{}, exp interface{}, msgArgs ...interface{}) (bool, error) {
	if isEqual(act, exp) {
		return false, internal.Errorf(msgArgs, "Should not equal:", internal.TypeValues(act, exp))
	}

	return true, nil
}

// EqualValue checks if 'act' == 'exp' without type check
func EqualValue(act interface{}, exp interface{}) bool {
	ok, _ := EqualValuef(act, exp)
	return ok
}

// EqualValuef checks if 'act' == 'exp'' without type check
func EqualValuef(act interface{}, exp interface{}, msgArgs ...interface{}) (bool, error) {
	if !isEqualValue(act, exp) {
		return false, internal.Errorf(msgArgs, "Not equal:", internal.Values(act, exp))
	}

	return true, nil
}

// NotEqualValue checks if 'act' != 'exp'' without type check
func NotEqualValue(act interface{}, exp interface{}) bool {
	ok, _ := NotEqualf(act, exp)
	return ok
}

// NotEqualValuef checks if 'act' != 'exp'' without type check
func NotEqualValuef(act interface{}, exp interface{}, msgArgs ...interface{}) (bool, error) {
	if isEqualValue(act, exp) {
		return false, internal.Errorf(msgArgs, "Should not equal:", internal.Values(act, exp))
	}

	return true, nil
}

// True checks if 'act' == true
func True(act bool) bool {
	ok, _ := Truef(act)
	return ok
}

// Truef checks if 'act' == true
func Truef(act bool, msgArgs ...interface{}) (bool, error) {
	if !act {
		return false, internal.Errorf(msgArgs, "Should be true: ", internal.Value(act))
	}

	return true, nil
}

// False checks if 'act' == false
func False(act bool) bool {
	ok, _ := Falsef(act)
	return ok
}

// Falsef checks if 'act' == false
func Falsef(act bool, msgArgs ...interface{}) (bool, error) {
	if act {
		return false, internal.Errorf(msgArgs, "Should be false: ", internal.Value(act))
	}

	return true, nil
}

// Length checks if len(col) == 'exp'
func Length(col interface{}, len int) bool {
	ok, _ := Lengthf(col, len)
	return ok
}

// Lengthf checks if len(col) == 'exp'
func Lengthf(col interface{}, len int, msgArgs ...interface{}) (bool, error) {
	ok, err := isLength(col, len)
	if err != nil {
		return false, err
	}

	if !ok {
		return false, internal.Errorf(msgArgs, "Wrong length:", internal.Values(col, len))
	}
	return true, nil
}

// Empty checks if len(col) == 0
func Empty(col interface{}) bool {
	ok, _ := Emptyf(col)
	return ok
}

// Emptyf checks if len(col) == 0
func Emptyf(col interface{}, msgArgs ...interface{}) (bool, error) {
	ok, err := isLength(col, 0)
	if err != nil {
		return false, err
	}

	if !ok {
		return false, internal.Errorf(msgArgs, "Should be empty:", internal.Value(col))
	}
	return true, nil
}

// NotEmpty checks if len(col) != 0
func NotEmpty(col interface{}) bool {
	ok, _ := NotEmptyf(col)
	return ok
}

// NotEmptyf checks if len(col) != 0
func NotEmptyf(col interface{}, msgArgs ...interface{}) (bool, error) {
	ok, err := isLength(col, 0)
	if err != nil {
		return false, err
	}

	if ok {
		return false, internal.Errorf(msgArgs, "Should not be empty:", internal.Value(col))
	}
	return true, nil
}

// Nil checks if 'act' == nil
func Nil(act interface{}) bool {
	ok, _ := Nilf(act)
	return ok
}

// Nilf checks if 'act' == nil
func Nilf(act interface{}, msgArgs ...interface{}) (bool, error) {
	if isNil(act) == false {
		return false, internal.Errorf(msgArgs, "Is not nil: ", internal.Value(act))
	}

	return true, nil
}

// NotNil checks if 'act' != nil
func NotNil(act interface{}) bool {
	ok, _ := NotNilf(act)
	return ok
}

// NotNilf checks if 'act' != nil
func NotNilf(act interface{}, msgArgs ...interface{}) (bool, error) {
	if isNil(act) {
		return false, internal.Errorf(msgArgs, "Is nil!", "")
	}

	return true, nil
}

// Error checks if 'err' != nil
func Error(err error) bool {
	ok, _ := Errorf(err)
	return ok
}

// Errorf checks if 'err' != nil
func Errorf(err error, msgArgs ...interface{}) (bool, error) {
	if isNil(err) {
		return false, internal.Errorf(msgArgs, "There should be an error.", "")
	}

	return true, nil
}

// Error checks if 'err' != nil
func NoError(err error) bool {
	ok, _ := NoErrorf(err)
	return ok
}

// Errorf checks if 'err' != nil
func NoErrorf(err error, msgArgs ...interface{}) (bool, error) {
	if isNil(err) == false {
		return false, internal.Errorf(msgArgs, "There should be no error.", internal.Value(err))
	}

	return true, nil
}

// Containing checks if the collection(array/slice/map/string) 'col' contains the given elements 'exp'.
// if 'col' is a map, it will check if the map have a value that is equal with 'exp'
func Containing(col interface{}, exp interface{}) bool {
	ok, _ := Containingf(col, exp)
	return ok
}

// Containingf checks if the collection(array/slice/map/string) 'col' contains the given elements 'exp'.
// if 'col' is a map, it will check if the map have a value that is equal with 'exp'
func Containingf(col interface{}, exp interface{}, msgArgs ...interface{}) (bool, error) {
	switch reflect.TypeOf(col).Kind() {
	case reflect.Slice, reflect.Array:
		if isInList(col, exp) == false {
			return false, internal.Errorf(msgArgs, "Element is not in array/slice:", internal.Collection(col, exp))
		}
	case reflect.Map:
		if isInMap(col, exp) == false {
			return false, internal.Errorf(msgArgs, "Element is not in map:", internal.Collection(col, exp))
		}
	case reflect.String:
		colString := reflect.ValueOf(col).String()
		if in := strings.Contains(colString, reflect.ValueOf(exp).String()); !in {
			return false, internal.Errorf(msgArgs, "Element is not in string:", internal.Collection(col, exp))
		}
	default:
		return false, internal.TypeError("Wrong type, should be a slice, array, map or string.", col)
	}

	return true, nil
}

// NotContaining checks if the collection 'col' contains not the given element 'exp'.
// if 'col' is a map, it will check if the map have not a value that is equal with 'exp'
func NotContaining(col interface{}, exp interface{}) bool {
	ok, _ := NotContainingf(col, exp)
	return ok
}

// NotContainingf checks if the collection 'col' contains not the given element 'exp'.
// if 'col' is a map, it will check if the map have not a value that is equal with 'exp'
func NotContainingf(col interface{}, exp interface{}, msgArgs ...interface{}) (bool, error) {
	switch reflect.TypeOf(col).Kind() {
	case reflect.Slice, reflect.Array:
		if isInList(col, exp) {
			return false, internal.Errorf(msgArgs, "Element is in array/slice:", internal.Collection(col, exp))
		}
	case reflect.Map:
		if isInMap(col, exp) {
			return false, internal.Errorf(msgArgs, "Element is in map:", internal.Collection(col, exp))
		}
	case reflect.String:
		colString := reflect.ValueOf(col).String()
		if strings.Contains(colString, reflect.ValueOf(exp).String()) {
			return false, internal.Errorf(msgArgs, "Element is in string:", internal.Collection(col, exp))
		}
	default:
		return false, internal.TypeError("Wrong type, should be a slice, array, map or string.", col)
	}

	return true, nil
}

// Similar checks if two arrays/slices contains the same items.
func Similar(act interface{}, exp interface{}) bool {
	ok, _ := Similarf(act, exp)
	return ok
}

// Similarf checks if two arrays/slices contains the same items.
func Similarf(act interface{}, exp interface{}, msgArgs ...interface{}) (bool, error) {
	aKind := reflect.TypeOf(act).Kind()
	eKind := reflect.TypeOf(exp).Kind()

	if (aKind == reflect.Array || aKind == reflect.Slice) &&
		(eKind == reflect.Array || eKind == reflect.Slice) {
		actList := reflect.ValueOf(act)
		expList := reflect.ValueOf(exp)

		// check first the length
		if actList.Len() != expList.Len() {
			return false, internal.Errorf(msgArgs, "Not similar, collections doesn't have the same length:",
				internal.Values(actList.Len(), expList.Len()))
		}

		// check if every element of 'exp' is in 'act'
		for i := 0; i < expList.Len(); i++ {
			eValue := expList.Index(i).Interface()
			if isInList(act, eValue) == false {
				return false, internal.Errorf(msgArgs, "Not similar:", internal.TypeValues(act, exp))
			}
		}
	} else {
		return false, internal.TypeError("Wrong type, should be a slice or array.", act)
	}

	return true, nil
}

// NotSimilar checks if two arrays/slices contains at least one different item.
func NotSimilar(act interface{}, exp interface{}) bool {
	ok, _ := NotSimilarf(act, exp)
	return ok
}

// NotSimilarf checks if two arrays/slices contains at least one different item.
func NotSimilarf(act interface{}, exp interface{}, msgArgs ...interface{}) (bool, error) {
	aKind := reflect.TypeOf(act).Kind()
	eKind := reflect.TypeOf(exp).Kind()

	if (aKind == reflect.Array || aKind == reflect.Slice) &&
		(eKind == reflect.Array || eKind == reflect.Slice) {
		aList := reflect.ValueOf(act)
		eList := reflect.ValueOf(exp)

		// check first the length
		if aList.Len() != eList.Len() {
			// not similar
			return true, nil
		}

		// check if every element of 'exp' is in 'act'
		for i := 0; i < eList.Len(); i++ {
			eValue := eList.Index(i).Interface()
			if isInList(act, eValue) == false {
				// no similar
				return true, nil
			}
		}
		// all element are in -> similar -> fail!
		return false, internal.Errorf(msgArgs, "Similar:", internal.TypeValues(act, exp))
	} else {
		return false, internal.TypeError("Wrong type, should be a slice or array.", act)
	}

	return true, nil
}

// OneOf check if the 'act' element one of the element inside the 'exp' array/slice.
func OneOf(act interface{}, exp interface{}) bool {
	ok, _ := OneOff(act, exp)
	return ok
}

// OneOff check if the 'act' element one of the element inside the 'exp' array/slice.
func OneOff(act interface{}, exp interface{}, msgArgs ...interface{}) (bool, error) {
	switch reflect.TypeOf(exp).Kind() {
	case reflect.Slice, reflect.Array:
		if isInList(exp, act) == false {
			return false, internal.Errorf(msgArgs, "Actual element is not in expected array/slice:", internal.Collection(exp, act))
		}
	default:
		return false, internal.TypeError("Wrong type, 'exp' should be a slice or array.", exp)
	}

	return true, nil
}

// ====================== Helper ===============================

func isLength(col interface{}, len int) (bool, error) {
	switch reflect.TypeOf(col).Kind() {
	case reflect.Slice, reflect.Array, reflect.Map, reflect.String:
		a := reflect.ValueOf(col).Len()
		return a == len, nil

	default:
		return false, internal.TypeError("Wrong type, should be a slice, array, map or string.", col)
	}
}

// check if the given array/slice contains the element
func isInList(col interface{}, elem interface{}) bool {
	list := reflect.ValueOf(col)
	for i := 0; i < list.Len(); i++ {
		e := list.Index(i).Interface()
		if isEqual(e, elem) {
			return true
		}
	}
	return false
}

// check if the given map contains the element as value
func isInMap(col interface{}, elem interface{}) bool {
	mp := reflect.ValueOf(col)
	for _, key := range mp.MapKeys() {
		e := mp.MapIndex(key).Interface()
		if isEqual(e, elem) {
			return true
		}
	}
	return false
}

// check if element is nil
func isNil(act interface{}) bool {
	return internal.IsNil(act)
}

func isEqual(act interface{}, exp interface{}) bool {
	if isNil(act) && isNil(exp) {
		return true
	}

	return reflect.DeepEqual(act, exp)
}

func isEqualValue(act interface{}, exp interface{}) bool {
	if isNil(act) && isNil(exp) {
		return true
	}

	actType := reflect.TypeOf(act)
	if actType == nil {
		return false
	}

	expValue := reflect.ValueOf(exp)
	if expValue.IsValid() && expValue.Type().ConvertibleTo(actType) {
		return reflect.DeepEqual(act, expValue.Convert(actType).Interface())
	}

	return false
}
