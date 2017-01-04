package main

import (
	"fmt"
	"testing"
)

func IntArrayEquals(lhs, rhs []int) bool {
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

func StringArrayEquals(lhs, rhs []string) bool {
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

func TestIndices(t *testing.T) {
	if !IntArrayEquals(Indices("abc"), []int{0, 1, 2}) {
		t.Fail()
	}
}

func TestPermutations(t *testing.T) {
	if len(Permutations("")) != 0 {
		t.Fail()
	}

	if len(Permutations("a")) != 1 {
		t.Fail()
	}

	if !StringArrayEquals(Permutations("a"), []string{"a"}) {
		t.Fail()
	}

	if len(Permutations("abc")) != 6 {
		t.Fail()
	}
}

func TestIntegerPermutations(t *testing.T) {
	if len(IntegerPermutations(23)) != 2 {
		t.Fail()
	}

	if len(IntegerPermutations(789)) != 6 {
		t.Fail()
	}
}

func TestIsCube(t *testing.T) {
	if IsCube(2) {
		t.Fail()
	}

	if !IsCube(8) {
		t.Fail()
	}

	if !IsCube(27) {
		t.Fail()
	}

	for i := 0; i < 1000; i++ {
		cube := i * i * i
		if !IsCube(cube) {
			t.Fail()
		}
	}
}

func TestIsPermutation(t *testing.T) {
	/*
		if !IsPermutation(12, 21) {
			t.Fail()
		}

		if IsPermutation(12, 31) {
			t.Fail()
		}
	*/
}

func TestXXX(t *testing.T) {
	if 'R' != rune(int('R')) {
		t.Fail()
	}
}

func TestSortedString(t *testing.T) {
	if SortedString("") != "" {
		t.Fail()
	}

	if SortedString("a") != "a" {
		t.Fail()
	}

	if SortedString("1022") != "0122" {
		t.Fail()
	}

	if SortedString("1022") != "0122" {
		t.Fail()
	}
}

func TestCubeMapKey(t *testing.T) {
	if CubeMapKey(102) != "012" {
		t.Fail()
	}
}

func _() {
	fmt.Println("xoxo")
}
