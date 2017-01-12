package main

import (
	"testing"
)

func TestIsPermutation(t *testing.T) {
	if !IsPermutation(23, 32) {
		t.Fail()
	}

	if IsPermutation(3, 9) {
		t.Fail()
	}
}
