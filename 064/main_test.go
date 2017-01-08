package main

import (
	"testing"
)

func ArrayEquals(lhs, rhs []int) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for i, _ := range lhs {
		if lhs[i] != rhs[i] {
			return false
		}
	}
	return true
}

func TestContinuedFractionOfRootFinite(t *testing.T) {
	a := ContinuedFractionOfRootFinite(2, 3)

	if len(a) != 3 {
		t.Fail()
	}

	if !ArrayEquals(a, []int{1, 2, 2}) {
		t.Fail()
	}
}

func TestContinuedFractionOfRoot(t *testing.T) {
	a := ContinuedFractionOfRoot(2)

	if len(a) != 2 {
		t.Fail()
	}
}
