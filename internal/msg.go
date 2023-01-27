package internal

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

const (
	spacerSize = 12
)

var (
	succIcon = TxtColor(Green, "✔")
	failIcon = TxtColor(Red, "✘")
	errIcon  = TxtColor(IRed, "🔥")
	actTxt   = TxtColor(Yellow, Spacer("actual:"))
	expTxt   = TxtColor(Yellow, Spacer("expected:"))
	colTxt   = TxtColor(Yellow, Spacer("collection:"))
	elemTxt  = TxtColor(Yellow, Spacer("element:"))
	valTxt   = TxtColor(Yellow, Spacer("value:"))
	typeTxt  = TxtColor(Yellow, Spacer("type:"))
	stackTxt = TxtColor(Yellow, Spacer("stack:"))
)

func Spacer(txt string) string {
	for i := len(txt); i < spacerSize; i++ {
		txt = " " + txt
	}

	return txt
}

// ------------------------ Message ---------------------------

func MsgArgs(msgArgs []interface{}, defaultMsg string) string {
	msg := defaultMsg
	if len(msgArgs) > 0 {
		_, ok := (msgArgs[0]).(string)
		if ok { // first part is a string
			msg = fmt.Sprintf(msgArgs[0].(string), msgArgs[1:]...)
		}
	}

	return msg
}

func Errorf(msgArgs []interface{}, defaultMsg string, valueCheck string) error {
	msg := MsgArgs(msgArgs, defaultMsg)
	return fmt.Errorf("%s %s\n%s", failIcon, TxtColor(Yellow, msg), valueCheck)
}

func TypeError(msg string, typ interface{}) error {
	return fmt.Errorf("%s %s\n%s", errIcon, TxtColor(Yellow, msg), Type(typ))
}

func StackTrace(offset int) string {
	stack := stackTrace(offset + 1) // +1 this method call
	if len(stack) > 1 {
		return "\n" + stackTxt + "\n" + strings.Join(stack, "\n")
	}

	return ""
}

func stackTrace(offset int) []string {
	var result []string

	for i := 2 + offset; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		f := runtime.FuncForPC(pc)
		if f == nil {
			break
		}

		if f.Name() == "testing.tRunner" {
			break
		}

		result = append(result, fmt.Sprintf("%s:%d", file, line))
	}

	return result
}

func Success() string {
	return succIcon
}

// --------------------- valueCheck -----------------------------

func Value(a interface{}) string {
	return fmt.Sprintf("%s %s", valTxt, toString(a))
}

func Values(a interface{}, b interface{}) string {
	return fmt.Sprintf(
		"%s %s\n%s %s", actTxt, toString(a), expTxt, toString(b))
}

func TypeValues(a interface{}, b interface{}) string {
	return fmt.Sprintf(
		"%s %s (%T)\n%s %s (%T)", actTxt, toString(a), a, expTxt, toString(b), b)
}

func Type(a interface{}) string {
	return fmt.Sprintf("%s %T", typeTxt, a)
}

// Collection: (Slice, Array, Map, String) return the value as string
func Collection(col interface{}, elem interface{}) string {
	return fmt.Sprintf(
		"%s %s (%T)\n%s %v (%T)", colTxt, toString(col), col, elemTxt, toString(elem), elem)
}

func toString(val interface{}) string {
	return valueToString(val, 1)
}

func valueToString(val interface{}, depth int) string {
	if IsNil(val) {
		return ""
	}

	vType := reflect.TypeOf(val).Kind()
	switch vType {
	case reflect.Slice, reflect.Array:
		return arrayToString(val, 1)
	case reflect.Map:
		return mapToString(val, 1)
	case reflect.String:
		return fmt.Sprint("\"", elemToString(val), "\"")
	case reflect.Ptr:
		return pointerToString(val, depth)
	}

	return elemToString(val)
}

func arrayToString(arr interface{}, depth int) (res string) {
	if depth == 0 {
		return elemToString(arr)
	}

	a := reflect.ValueOf(arr)
	for i := 0; i < a.Len(); i++ {
		e := a.Index(i)

		if len(res) > 0 {
			res += ", "
			if isBiggerStructure(e) {
				res += "\n\t"
			}
		}

		res += valueToString(e.Interface(), depth-1)
	}
	return "[" + res + "]"
}

func mapToString(mp interface{}, depth int) (res string) {
	if depth == 0 {
		return elemToString(mp)
	}

	m := reflect.ValueOf(mp)
	for _, key := range m.MapKeys() {
		if len(res) > 0 {
			res += ",\n\t"
		}

		e := m.MapIndex(key)
		res += fmt.Sprint(key) + ":" + valueToString(e.Interface(), depth-1)
	}
	return "{" + res + "}"
}

func pointerToString(ptr interface{}, depth int) (res string) {
	p := reflect.ValueOf(ptr)
	p = p.Elem()
	return "&" + valueToString(p.Interface(), depth)
}

func elemToString(elem interface{}) (res string) {
	return fmt.Sprintf("%+v", elem)
}

func isBiggerStructure(e reflect.Value) bool {
	kind := e.Kind()
	return kind == reflect.Array ||
		kind == reflect.Slice ||
		kind == reflect.Map ||
		kind == reflect.Struct ||
		kind == reflect.Interface
}
