package should_test

import (
	"fmt"
	"testing"

	. "github.com/maprost/testbox/should"
)

type mockTest struct {
	*testing.T
	FailedTest bool
}

func (t *mockTest) Error(args ...interface{}) {
	t.FailedTest = true
}

func testToFail(t *testing.T, testFunc func(t testing.TB)) {
	t.Helper()
	mt := mockTest{t, false}
	testFunc(&mt)
	BeTrue(t, mt.FailedTest)
}

func TestSimple(t *testing.T) {
	t.Run("equal", func(t *testing.T) {
		BeEqual(t, 1, 1)

		testToFail(t, func(t testing.TB) {
			BeEqual(t, 1, 2)
		})
	})

	t.Run("not equal", func(t *testing.T) {
		NotBeEqual(t, 1, 2)

		testToFail(t, func(t testing.TB) {
			NotBeEqual(t, 1, 1)
		})
	})

	t.Run("nil", func(t *testing.T) {
		BeNil(t, nil)

		testToFail(t, func(t testing.TB) {
			BeNil(t, 1)
		})
		testToFail(t, func(t testing.TB) {
			BeNil(t, &mockTest{nil, false})
		})
	})

	t.Run("not nil", func(t *testing.T) {
		NotBeNil(t, 1)

		NotBeNil(t, &mockTest{nil, false})

		testToFail(t, func(t testing.TB) {
			NotBeNil(t, nil)
		})
	})

	t.Run("error", func(t *testing.T) {
		BeError(t, fmt.Errorf("blob"))

		testToFail(t, func(t testing.TB) {
			BeError(t, nil)
		})
	})

	t.Run("not error", func(t *testing.T) {
		BeNoError(t, nil)

		testToFail(t, func(t testing.TB) {
			BeNoError(t, fmt.Errorf("blob"))
		})
	})

	t.Run("true", func(t *testing.T) {
		BeTrue(t, true)

		testToFail(t, func(t testing.TB) {
			BeTrue(t, false)
		})
	})

	t.Run("false", func(t *testing.T) {
		BeFalse(t, false)

		testToFail(t, func(t testing.TB) {
			BeFalse(t, true)
		})
	})

	t.Run("have length", func(t *testing.T) {
		HaveLength(t, []int{1, 2, 3}, 3)

		testToFail(t, func(t testing.TB) {
			HaveLength(t, []int{1, 2, 3}, 2)
		})
	})

	t.Run("is empty", func(t *testing.T) {
		BeEmpty(t, []int32{})

		testToFail(t, func(t testing.TB) {
			BeEmpty(t, []int{1})
		})
	})

	t.Run("is not empty", func(t *testing.T) {
		NotBeEmpty(t, []int32{int32(1)})

		testToFail(t, func(t testing.TB) {
			NotBeEmpty(t, []int{})
		})
	})

	t.Run("contain", func(t *testing.T) {
		Contain(t, [3]int{1, 2, 3}, 3)

		testToFail(t, func(t testing.TB) {
			Contain(t, [3]int{1, 2, 3}, 4)
		})
	})

	t.Run("not contain", func(t *testing.T) {
		NotContain(t, [3]int{1, 2, 3}, 4)

		testToFail(t, func(t testing.TB) {
			NotContain(t, [3]int{1, 2, 3}, 3)
		})
	})

	t.Run("similar", func(t *testing.T) {
		BeSimilar(t, []int{1, 2, 3}, []int{3, 2, 1})

		testToFail(t, func(t testing.TB) {
			BeSimilar(t, []int{1, 2, 3}, [3]int{3, 2, 4})
		})
	})

	t.Run("not similar", func(t *testing.T) {
		NotBeSimilar(t, []int{1, 2, 3}, [3]int{3, 2, 4})

		testToFail(t, func(t testing.TB) {
			NotBeSimilar(t, []int{1, 2, 3}, []int{3, 2, 1})
		})
	})

	t.Run("one of", func(t *testing.T) {
		BeOneOf(t, 2, [3]int{1, 2, 3})

		testToFail(t, func(t testing.TB) {
			BeOneOf(t, 4, []int{1, 2, 3})
		})
	})

	t.Run("fail", func(t *testing.T) {
		testToFail(t, func(t testing.TB) {
			Fail(t)
		})
	})

	t.Run("be equal struct", func(t *testing.T) {
		type r struct {
			Blob    string
			Drop    bool
			private int
		}

		BeEqualStructField(t,
			r{Blob: "hello", Drop: true, private: 1},
			r{Blob: "hello", Drop: true, private: 1})
	})

	t.Run("be equal struct", func(t *testing.T) {
		type r struct {
			Blob string
			Drop bool
		}
		testToFail(t, func(t testing.TB) {
			BeEqualStructField(t,
				r{Blob: "hello", Drop: true},
				r{Blob: "world", Drop: true})
		})
	})
}
