package is_test

import (
	"fmt"
	"testing"

	"github.com/maprost/testbox/is"
	"github.com/maprost/testbox/should"
)

type TestStruct struct {
	Id  int64
	Msg string
}

var (
	ts1 = TestStruct{Id: 12, Msg: "New"}
	ts2 = TestStruct{Id: 12, Msg: "New"}
	ts3 = TestStruct{Id: 12, Msg: "Old"}
)

type TestErrorStruct struct{}

func (TestErrorStruct) Error() string { return "" }

func TestTrue(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		should.BeTrue(t, is.True(true))
	})

	t.Run("negative", func(t *testing.T) {
		should.BeFalse(t, is.True(false))
	})
}

func TestFalse(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		should.BeTrue(t, is.False(false))
	})

	t.Run("negative", func(t *testing.T) {
		should.BeFalse(t, is.False(true))
	})
}

func TestEqual(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		should.BeTrue(t, is.Equal(1, 1))
		should.BeTrue(t, is.Equal(1.1, 1.1))
		should.BeTrue(t, is.Equal("Hello", "Hello"))
		should.BeTrue(t, is.Equal([3]int{1, 2, 3}, [3]int{1, 2, 3}))
		should.BeTrue(t, is.Equal(map[int]string{1: "1", 2: "2", 3: "3"}, map[int]string{1: "1", 2: "2", 3: "3"}))
		should.BeTrue(t, is.Equal(ts1, ts1))
		should.BeTrue(t, is.Equal(ts1, ts2))
		should.BeTrue(t, is.Equal(&ts1, &ts2))
		should.BeTrue(t, is.Equal([]TestStruct{ts1}, []TestStruct{ts2}))
		should.BeTrue(t, is.Equal(nil, nil))

		var zeroNilPtr *int
		should.BeTrue(t, is.Equal(nil, zeroNilPtr))
	})

	t.Run("negative", func(t *testing.T) {
		should.BeFalse(t, is.Equal(1.0, 1))
		should.BeFalse(t, is.Equal("1", 1))
		should.BeFalse(t, is.Equal(ts1, ts3))
		should.BeFalse(t, is.Equal(&ts1, &ts3))
		should.BeFalse(t, is.Equal([]TestStruct{ts1}, []TestStruct{ts3}))
		should.BeFalse(t, is.Equal(nil, 1))
		should.BeFalse(t, is.Equal("hello", nil))
	})
}

func TestEqualMessage(t *testing.T) {
	t.Run("empty error", func(t *testing.T) {
		ok, emptyErr := is.Equalf(1, 1)
		should.BeTrue(t, ok)
		should.BeNil(t, emptyErr)
	})

	t.Run("error", func(t *testing.T) {
		ok, err := is.Equalf(1, 2)
		should.BeFalse(t, ok)
		should.Contain(t, err.Error(), "Not equal:")
	})

	t.Run("custom error", func(t *testing.T) {
		ok, customErr := is.Equalf(1, 2, "My error:")
		should.BeFalse(t, ok)
		should.Contain(t, customErr.Error(), "My error:")
	})

	t.Run("long custom error", func(t *testing.T) {
		ok, customLongErr := is.Equalf(1, 2, "My error %s %d:", "super", 10)
		should.BeFalse(t, ok)
		should.Contain(t, customLongErr.Error(), "My error super 10")
	})
}

func TestNotEqual(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		should.BeTrue(t, is.NotEqual(1.0, 1))
		should.BeTrue(t, is.NotEqual("1", 1))
		should.BeTrue(t, is.NotEqual(ts1, ts3))
		should.BeTrue(t, is.NotEqual(&ts1, &ts3))
		should.BeTrue(t, is.NotEqual(nil, 1))
		should.BeTrue(t, is.NotEqual("hello", nil))
	})

	t.Run("negative", func(t *testing.T) {
		should.BeFalse(t, is.NotEqual(1, 1))
		should.BeFalse(t, is.NotEqual(1.1, 1.1))
		should.BeFalse(t, is.NotEqual("Hello", "Hello"))
		should.BeFalse(t, is.NotEqual([3]int{1, 2, 3}, [3]int{1, 2, 3}))
		should.BeFalse(t, is.NotEqual(map[int]string{1: "1", 2: "2", 3: "3"}, map[int]string{1: "1", 2: "2", 3: "3"}))
		should.BeFalse(t, is.NotEqual(ts1, ts1))
		should.BeFalse(t, is.NotEqual(ts1, ts2))
		should.BeFalse(t, is.NotEqual(&ts1, &ts2))
		should.BeFalse(t, is.NotEqual(nil, nil))

		var zeroNilPtr *int
		should.BeFalse(t, is.NotEqual(nil, zeroNilPtr))
	})
}

func TestEqualValue(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		should.BeTrue(t, is.EqualValue(int(1), int64(1)))
		should.BeTrue(t, is.EqualValue(uint8(1), uint64(1)))
		should.BeTrue(t, is.EqualValue(1, 1.0))
	})

	t.Run("negative", func(t *testing.T) {

	})
}

func TestLength(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		should.BeTrue(t, is.Length([]string{}, 0))
		should.BeTrue(t, is.Length([]string{"hello", "world"}, 2))

		var strArr [1]string
		strArr[0] = "hello"
		should.BeTrue(t, is.Length(strArr, 1))

		should.BeTrue(t, is.Length("Welcome", 7))

		should.BeTrue(t, is.Length(map[string]int{"1": 1}, 1))
	})

	t.Run("negative", func(t *testing.T) {
		should.BeFalse(t, is.Length([]string{}, 1))

		var strArr [1]string
		should.BeFalse(t, is.Length(strArr, 2))

		should.BeFalse(t, is.Length(" ", 2))

		should.BeFalse(t, is.Length(map[string]int{"1": 1}, 0))
	})

	t.Run("wrong type", func(t *testing.T) {
		should.BeFalse(t, is.Length(1, 1))
	})
}

func TestEmpty(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		should.BeTrue(t, is.Empty([]string{}))

		var strArr [0]string
		should.BeTrue(t, is.Empty(strArr))

		should.BeTrue(t, is.Empty(""))

		should.BeTrue(t, is.Empty(map[string]int{}))
	})

	t.Run("negative", func(t *testing.T) {
		should.BeFalse(t, is.Empty([]string{""}))

		var strArr [1]string
		should.BeFalse(t, is.Empty(strArr))

		should.BeFalse(t, is.Empty(" "))

		should.BeFalse(t, is.Empty(map[string]int{"1": 1}))
	})

	t.Run("wrong type", func(t *testing.T) {
		should.BeFalse(t, is.Empty(1))
	})
}

func TestNotEmpty(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		should.BeTrue(t, is.NotEmpty([]string{""}))

		var strArr [1]string
		should.BeTrue(t, is.NotEmpty(strArr))

		should.BeTrue(t, is.NotEmpty(" "))

		should.BeTrue(t, is.NotEmpty(map[string]int{"1": 1}))
	})

	t.Run("negative", func(t *testing.T) {
		should.BeFalse(t, is.NotEmpty([]string{}))

		var strArr [0]string
		should.BeFalse(t, is.NotEmpty(strArr))

		should.BeFalse(t, is.NotEmpty(""))

		should.BeFalse(t, is.NotEmpty(map[string]int{}))
	})

	t.Run("wrong type", func(t *testing.T) {
		should.BeFalse(t, is.NotEmpty(1))
	})
}

func TestNil(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		should.BeTrue(t, is.Nil(nil))

		var zeroNilPtr *int
		should.BeTrue(t, is.Nil(zeroNilPtr))
	})

	t.Run("negative", func(t *testing.T) {
		should.BeFalse(t, is.Nil([]string{}))

		should.BeFalse(t, is.Nil(1))

		should.BeFalse(t, is.Nil(""))

		i := 4
		var iPtr *int
		iPtr = &i
		should.BeFalse(t, is.Nil(iPtr))
	})
}

func TestNotNil(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		should.BeTrue(t, is.NotNil([]string{}))

		should.BeTrue(t, is.NotNil(1))

		should.BeTrue(t, is.NotNil(""))

		i := 4
		var iPtr *int
		iPtr = &i
		should.BeTrue(t, is.NotNil(iPtr))
	})

	t.Run("negative", func(t *testing.T) {
		should.BeFalse(t, is.NotNil(nil))

		var zeroNilPtr *int
		should.BeFalse(t, is.NotNil(zeroNilPtr))
	})
}

func TestError(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		should.BeTrue(t, is.Error(fmt.Errorf("fail")))

		should.BeTrue(t, is.Error(TestErrorStruct{}))
	})

	t.Run("negative", func(t *testing.T) {
		should.BeFalse(t, is.Error(nil))

		var err *TestErrorStruct
		should.BeFalse(t, is.Error(err))
	})
}

func TestNoError(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		should.BeTrue(t, is.NoError(nil))

		var err *TestErrorStruct
		should.BeTrue(t, is.NoError(err))
	})

	t.Run("negative", func(t *testing.T) {
		should.BeFalse(t, is.NoError(fmt.Errorf("fail")))

		should.BeFalse(t, is.NoError(TestErrorStruct{}))
	})
}

func TestContaining(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		should.BeTrue(t, is.Containing([]int{1, 2, 3}, 2))
		should.BeTrue(t, is.Containing([][]int{{1, 2}, {2, 3}}, []int{2, 3}))

		var strArr [1]string
		should.BeTrue(t, is.Containing(strArr, ""))

		should.BeTrue(t, is.Containing("hello world", "or"))

		should.BeTrue(t, is.Containing(map[string]int{"1": 1, "2": 32, "4": 42, "5": 89}, 42))
	})

	t.Run("negative", func(t *testing.T) {
		should.BeFalse(t, is.Containing([]int{1, 2, 3}, 4))

		var strArr [1]string
		strArr[0] = "hello world"
		should.BeFalse(t, is.Containing(strArr, "hello"))

		should.BeFalse(t, is.Containing("hello world", "and"))

		should.BeFalse(t, is.Containing(map[string]int{"1": 1, "2": 32, "4": 42, "5": 89}, 21))
	})

	t.Run("wrong type", func(t *testing.T) {
		should.BeFalse(t, is.Containing(1, 1))
	})
}

func TestNotContaining(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		should.BeTrue(t, is.NotContaining([]int{1, 2, 3}, 4))

		var strArr [1]string
		strArr[0] = "hello world"
		should.BeTrue(t, is.NotContaining(strArr, "hello"))

		should.BeTrue(t, is.NotContaining("hello world", "and"))

		should.BeTrue(t, is.NotContaining(map[string]int{"1": 1, "2": 32, "4": 42, "5": 89}, 21))
	})

	t.Run("negative", func(t *testing.T) {
		should.BeFalse(t, is.NotContaining([]int{1, 2, 3}, 2))
		should.BeFalse(t, is.NotContaining([][]int{{1, 2}, {2, 3}}, []int{2, 3}))

		var strArr [1]string
		should.BeFalse(t, is.NotContaining(strArr, ""))

		should.BeFalse(t, is.NotContaining("hello world", "or"))

		should.BeFalse(t, is.NotContaining(map[string]int{"1": 1, "2": 32, "4": 42, "5": 89}, 42))
	})

	t.Run("wrong type", func(t *testing.T) {
		should.BeFalse(t, is.NotContaining(1, 1))
	})
}

func TestSimilar(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		should.BeTrue(t, is.Similar([]int{1, 2, 3}, []int{3, 2, 1}))
		should.BeTrue(t, is.Similar([]int{1, 2, 3}, [3]int{3, 2, 1}))
	})

	t.Run("negative", func(t *testing.T) {
		should.BeFalse(t, is.Similar([]int{1, 2, 3}, [3]int{3, 2, 4}))
		should.BeFalse(t, is.Similar([]int{1, 2, 3}, []int{2, 1}))
	})

	t.Run("wrong type", func(t *testing.T) {
		should.BeFalse(t, is.Similar(1, 1))
	})
}

func TestNotSimilar(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		should.BeTrue(t, is.NotSimilar([]int{1, 2, 3}, [3]int{3, 2, 4}))
		should.BeTrue(t, is.NotSimilar([]int{1, 2, 3}, []int{2, 1}))
	})

	t.Run("negative", func(t *testing.T) {
		should.BeFalse(t, is.NotSimilar([]int{1, 2, 3}, []int{3, 2, 1}))
		should.BeFalse(t, is.NotSimilar([]int{1, 2, 3}, [3]int{3, 2, 1}))
	})

	t.Run("wrong type", func(t *testing.T) {
		should.BeFalse(t, is.NotSimilar(1, 1))
	})
}

func TestDataStructures(t *testing.T) {
	//is.Length([]int{1, 2, 3}, 3)
	//is.Length([3]int{1, 2, 3}, 3)
	//is.Length(map[int]int{1: 1, 2: 3, 3: 5}, 3)
	//is.BeEqual([]int{1, 2, 3}, []int{1, 2, 3})
	//is.NotBeEqual(t, []int{1, 2, 3}, []int{3, 2, 1})
	//is.BeSimilar(t, []int{1, 2, 3}, []int{3, 2, 1})
	//is.BeSimilar(t, []int{1, 2, 3}, [3]int{3, 2, 1})
	//is.NotBeSimilar(t, []int{1, 2, 3}, [3]int{3, 2, 4})
	//is.NotBeSimilar(t, []int{1, 2, 3}, []int{2, 3, 4, 1})
	//is.NotBeEqual(t, []int{1, 2, 3}, [3]int{1, 2, 3})
	//is.BeEqual(t, [3]int{1, 2, 3}, [3]int{1, 2, 3})
	//is.Contain(t, []int{1, 2, 3}, 2)
	//is.NotContain(t, []int{1, 2, 3}, 4)
	//is.Contain(t, [3]int{1, 2, 3}, 3)
	//is.NotContain(t, [3]int{1, 2, 3}, 4)
	//is.BeEqual(t, map[int]string{1: "1", 2: "2", 3: "3"}, map[int]string{1: "1", 2: "2", 3: "3"})
	//is.NotBeEqual(t, map[int]string{1: "1", 3: "3"}, map[int]string{1: "1", 4: "4"})
	//is.BeEqual(t, map[int]string{1: "1", 3: "3"}, map[int]string{3: "3", 1: "1"})
	//is.Contain(t, map[int]string{1: "11", 3: "33"}, "33")
	//is.NotContain(t, map[int]string{1: "11", 3: "33"}, "55")
}

func TestStructs(t *testing.T) {
	//type Post struct {
	//	Id  int64
	//	Msg string
	//}
	//
	//p1 := Post{Id: 12, Msg: "New"}
	//p2 := Post{Id: 12, Msg: "New"}
	//p3 := Post{Id: 12, Msg: "Old"}
	//
	//
	//is.BeEqual(t, []Post{p1}, []Post{p2})
	//is.NotBeEqual(t, []Post{p1}, []Post{p3})
	//is.Contain(t, []Post{p1, p2, p3}, p3)
	//is.Contain(t, []*Post{&p1, &p2, &p3}, &p3)
	//is.Contain(t, []*Post{&p1, &p3}, &p2)
	//is.NotContain(t, []Post{p1, p2}, p3)
	//is.NotContain(t, []Post{p1}, 22)
	//is.NotContain(t, []Post{}, p1)
	//is.NotContain(t, []*Post{&p1, &p2}, &p3)
	//is.NotContain(t, []*Post{&p1, &p2}, p2)
	//
	//is.BeEqual(t, [1]Post{p1}, [1]Post{p2})
	//is.NotBeEqual(t, [1]Post{p1}, [1]Post{p3})
	//is.Contain(t, [2]Post{p1, p2}, p2)
	//is.Contain(t, [3]*Post{&p1, &p2, &p3}, &p3)
	//is.Contain(t, [2]*Post{&p1, &p3}, &p2)
	//is.NotContain(t, [2]Post{p1, p2}, p3)
	//is.NotContain(t, [1]Post{p1}, "blob")
	//is.NotContain(t, [0]Post{}, p1)
	//is.NotContain(t, [2]*Post{&p1, &p2}, &p3)
	//is.NotContain(t, [2]*Post{&p1, &p2}, p1)
	//
	//is.BeEqual(t, map[int]Post{1: p1}, map[int]Post{1: p2})
	//is.NotBeEqual(t, map[int]Post{1: p1}, map[int]Post{1: p3})
	//is.NotBeEqual(t, map[int]Post{1: p1}, map[int]Post{2: p1})
	//is.Contain(t, map[int]Post{1: p1, 2: p2}, p2)
	//is.Contain(t, map[int]*Post{1: &p1, 2: &p2}, &p2)
	//is.NotContain(t, map[int]Post{1: p1, 2: p2}, p3)
	//is.NotContain(t, map[int]Post{1: p1}, "blob")
	//is.NotContain(t, map[int]Post{}, p1)
	//is.NotContain(t, map[int]*Post{1: &p1}, p1)
}

func TestNilWithMapDefault(t *testing.T) {
	//m := make(map[int]*int)
	//
	//is.BeNil(t, m[4])
}

func TestOneOf(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		should.BeTrue(t, is.OneOf("hello", []string{"world", "hello"}))
	})

	t.Run("negative", func(t *testing.T) {
		should.BeFalse(t, is.OneOf("hello", []string{"world", "blob"}))
	})

	t.Run("wrong type", func(t *testing.T) {
		should.BeFalse(t, is.OneOf("hello", 1))
	})
}
