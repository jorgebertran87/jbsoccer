package tests

import "testing"

type Test struct {
	T *testing.T
}

func (test *Test) AssertSame(v1 string, v2 string) {
	if v1 != v2 {
		test.T.Error("Expected: " + v1 + ", Actual: " + v2)
	}
}

func (test *Test) AssertTrue(condition bool) {
	if condition == false {
		test.T.Fail()
	}
}