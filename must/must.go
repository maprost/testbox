package must

import (
	"testing"

	"github.com/maprost/testbox/internal"
	"github.com/maprost/testbox/is"
)

// BeEqual checks if 'act' == 'exp'
func BeEqual(t testing.TB, act interface{}, exp interface{}, msgArgs ...interface{}) {
	t.Helper()
	if ok, err := is.Equalf(act, exp, msgArgs); !ok {
		fail(t, err)
	} else {
		success(t)
	}
}

// NotBeEqual checks if 'act' != 'exp'
func NotBeEqual(t testing.TB, act interface{}, exp interface{}, msgArgs ...interface{}) {
	t.Helper()
	if ok, err := is.NotEqualf(act, exp, msgArgs); !ok {
		fail(t, err)
	} else {
		success(t)
	}
}

// BeTrue checks if 'act' == true
func BeTrue(t testing.TB, act bool, msgArgs ...interface{}) {
	t.Helper()
	if ok, err := is.Truef(act, msgArgs); !ok {
		fail(t, err)
	} else {
		success(t)
	}
}

// BeFalse checks if 'act' == false
func BeFalse(t testing.TB, act bool, msgArgs ...interface{}) {
	t.Helper()
	if ok, err := is.Falsef(act, msgArgs); !ok {
		fail(t, err)
	} else {
		success(t)
	}
}

// HaveLength checks if len(col) == 'exp'
func HaveLength(t testing.TB, col interface{}, len int, msgArgs ...interface{}) {
	t.Helper()
	if ok, err := is.Lengthf(col, len, msgArgs); !ok {
		fail(t, err)
	} else {
		success(t)
	}
}

// BeEmpty checks if len(col) == 0
func BeEmpty(t testing.TB, col interface{}, msgArgs ...interface{}) {
	t.Helper()
	if ok, err := is.Emptyf(col, msgArgs); !ok {
		fail(t, err)
	} else {
		success(t)
	}
}

// NotBeEmpty checks if len(col) != 0
func NotBeEmpty(t testing.TB, col interface{}, msgArgs ...interface{}) {
	t.Helper()
	if ok, err := is.NotEmptyf(col, msgArgs); !ok {
		fail(t, err)
	} else {
		success(t)
	}
}

// BeNil checks if 'act' == nil
func BeNil(t testing.TB, act interface{}, msgArgs ...interface{}) {
	t.Helper()
	if ok, err := is.Nilf(act, msgArgs); !ok {
		fail(t, err)
	} else {
		success(t)
	}
}

// NotBeNil checks if 'act' != nil
func NotBeNil(t testing.TB, act interface{}, msgArgs ...interface{}) {
	t.Helper()
	if ok, err := is.NotNilf(act, msgArgs); !ok {
		fail(t, err)
	} else {
		success(t)
	}
}

// HaveError checks if 'err' != nil
func HaveError(t testing.TB, e error, msgArgs ...interface{}) {
	t.Helper()
	if ok, err := is.Errorf(e, msgArgs); !ok {
		fail(t, err)
	} else {
		success(t)
	}
}

// HaveNoError checks if 'err' == nil
func HaveNoError(t testing.TB, e error, msgArgs ...interface{}) {
	t.Helper()
	if ok, err := is.NoErrorf(e, msgArgs); !ok {
		fail(t, err)
	} else {
		success(t)
	}
}

// Contain checks if the collection(array/slice/map/string) 'col' contains the given elements 'exp'.
// if 'col' is a map, it will check if the map have a value that is equal with 'exp'
func Contain(t testing.TB, col interface{}, exp interface{}, msgArgs ...interface{}) {
	t.Helper()
	if ok, err := is.Containingf(col, exp, msgArgs); !ok {
		fail(t, err)
	} else {
		success(t)
	}
}

// NotContain checks if the collection 'col' contains not the given element 'exp'.
// if 'col' is a map, it will check if the map have not a value that is equal with 'exp'
func NotContain(t testing.TB, col interface{}, exp interface{}, msgArgs ...interface{}) {
	t.Helper()
	if ok, err := is.NotContainingf(col, exp, msgArgs); !ok {
		fail(t, err)
	} else {
		success(t)
	}
}

// BeSimilar checks if two arrays/slices contains the same items.
func BeSimilar(t testing.TB, act interface{}, exp interface{}, msgArgs ...interface{}) {
	t.Helper()
	if ok, err := is.Similarf(act, exp, msgArgs); !ok {
		fail(t, err)
	} else {
		success(t)
	}
}

// NotBeSimilar checks if two arrays/slices contains at least one different item.
func NotBeSimilar(t testing.TB, act interface{}, exp interface{}, msgArgs ...interface{}) {
	t.Helper()
	if ok, err := is.NotSimilarf(act, exp, msgArgs); !ok {
		fail(t, err)
	} else {
		success(t)
	}
}

// BeOneOf check if the 'act' element one of the element inside the 'exp' array/slice.
func BeOneOf(t testing.TB, act interface{}, exp interface{}, msgArgs ...interface{}) {
	t.Helper()
	if ok, err := is.OneOff(act, exp, msgArgs); !ok {
		fail(t, err)
	} else {
		success(t)
	}
}

// Fail with message
func Fail(t testing.TB, msgArgs ...interface{}) {
	t.Helper()
	fail(t, internal.Errorf(msgArgs, "Something failed", ""))
}

// ====================== Helper ===============================

func fail(t testing.TB, err error) {
	t.Helper()
	stack := internal.StackTrace(1)
	t.Fatal(err.Error(), stack, "\n")
}

func success(t testing.TB) {
	t.Helper()
	t.Log(internal.Success())
}
