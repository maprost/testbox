package internal_test

import (
	"fmt"
	"strings"
	"testing"

	"errors"
	"github.com/maprost/testbox/internal"
	"github.com/maprost/testbox/should"
)

func clearErrColor(err error) error {
	return fmt.Errorf(clearStrColor(err.Error()))
}

func clearStrColor(s string) string {
	s = strings.Replace(s, internal.Color_Off, "", -1)
	s = strings.Replace(s, internal.Yellow, "", -1)
	s = strings.Replace(s, internal.Red, "", -1)
	s = strings.Replace(s, internal.IRed, "", -1)
	s = strings.Replace(s, internal.Green, "", -1)
	return s
}

// ------------------------ Message ---------------------------

func TestErrorf(t *testing.T) {
	should.BeEqual(t, clearErrColor(internal.Errorf([]interface{}{"jo"}, "no", "so")),
		fmt.Errorf("✘ jo\nso"))

	should.BeEqual(t, clearErrColor(internal.Errorf([]interface{}{}, "no", "so")),
		fmt.Errorf("✘ no\nso"))
}

func TestTypeError(t *testing.T) {
	var a map[*int]*string
	should.BeEqual(t, clearErrColor(internal.TypeError("no", a)),
		fmt.Errorf("🔥 no\n       type: map[*int]*string"))
}

func TestStackTrace(t *testing.T) {
	t.Run("stack with 3 elements", func(t *testing.T) {
		stackTrace := internal.StackTrace(-3)

		should.HaveLength(t, strings.Split(stackTrace, "\n"), 5)
		should.Contain(t, stackTrace, "src/github.com/maprost/testbox/internal/msg_test.go")
	})

	t.Run("stack with 1 element", func(t *testing.T) {
		should.BeEmpty(t, internal.StackTrace(-1))
	})
}

func TestSuccess(t *testing.T) {
	should.BeEqual(t, clearStrColor(internal.Success()), "✔")
}

// --------------------- valueCheck -----------------------------

func TestValue(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		str := "hello"
		should.BeEqual(t, clearStrColor(internal.Value(str)), "      value: \"hello\"")
	})

	t.Run("string pointer", func(t *testing.T) {
		str := "hello"
		ptr := &str
		should.BeEqual(t, clearStrColor(internal.Value(ptr)), "      value: &\"hello\"")
	})

	t.Run("int", func(t *testing.T) {
		i := 1
		should.BeEqual(t, clearStrColor(internal.Value(i)), "      value: 1")
	})

	t.Run("int pointer", func(t *testing.T) {
		i := 1
		ptr := &i
		should.BeEqual(t, clearStrColor(internal.Value(ptr)), "      value: &1")
	})

	t.Run("array", func(t *testing.T) {
		a := []int{1, 2, 3}
		should.BeEqual(t, clearStrColor(internal.Value(a)), "      value: [1, 2, 3]")
	})

	t.Run("array pointer", func(t *testing.T) {
		a := []int{1, 2, 3}
		ptr := &a
		should.BeEqual(t, clearStrColor(internal.Value(ptr)), "      value: &[1, 2, 3]")
	})

	t.Run("array with structs", func(t *testing.T) {
		a := []error{errors.New("1"), errors.New("2"), errors.New("3")}
		should.BeEqual(t, clearStrColor(internal.Value(a)), "      value: [1, 2, 3]")
	})

	t.Run("map", func(t *testing.T) {
		mp := map[int]string{1: "hello", 2: "welt"}
		should.BeOneOf(t, clearStrColor(internal.Value(mp)), []string{
			"      value: {1:\"hello\",\n\t2:\"welt\"}",
			"      value: {2:\"welt\",\n\t1:\"hello\"}",
		})
	})

	t.Run("map with pointer value", func(t *testing.T) {
		str := "hello"
		ptr := map[int]*string{1: &str}
		should.BeEqual(t, clearStrColor(internal.Value(ptr)), "      value: {1:&\"hello\"}")
	})

	t.Run("pointer map", func(t *testing.T) {
		mpPtr := &(map[int]string{1: "hello", 2: "welt"})
		should.BeOneOf(t, clearStrColor(internal.Value(mpPtr)), []string{
			"      value: &{1:\"hello\",\n\t2:\"welt\"}",
			"      value: &{2:\"welt\",\n\t1:\"hello\"}",
		})
	})
}

func TestType(t *testing.T) {
	var a map[*int]*string
	should.BeEqual(t, clearStrColor(internal.Type(a)), "       type: map[*int]*string")
}

func TestCollection(t *testing.T) {
	//type Drop struct {
	//	Hidden string
	//}
	//
	//drop := &Drop{"secret"}
	//should.BeEqual(t, internal.Collection([]*Drop{drop, drop}, drop), `
	//collection: [0xc420052470] ([]*msg_test.Drop)
	//   element: &{secret} (*msg_test.Drop)`)
}
