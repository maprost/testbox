package testbox_test

import (
	"fmt"
	"testing"

	"github.com/maprost/testbox"
)

func TestSingleTest(t *testing.T) {
	testbox.SingleTest(t)

	fmt.Println("Don't skip me")
}

func TestSingleTestWithSkip(t *testing.T) {
	testbox.SingleTestWithSkip(t, 1)

	fmt.Println("skip me")
}
