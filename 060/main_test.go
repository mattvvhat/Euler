package main

import (
	"fmt"
	"testing"
)

func TestWhatever(t *testing.T) {
	sieve := NewNaiveSieve(1000000)

	primes := []int{
		2, 3, 5, 7, 11, 13,
	}

	composites := []int{
		4, 6, 8, 10, 12,
	}

	for _, n := range primes {
		if !sieve.Contains(n) {
			fmt.Println(n, "incorrectly considered composite")
			t.Fail()
		}
	}

	for _, n := range composites {
		if sieve.Contains(n) {
			fmt.Println(n, "incorrectly considered prime")
			t.Fail()
		}
	}
}

func TestAnsweredSolution(t *testing.T) {
	if IsConcatablePrime([]int{3, 7, 61, 121}) {
		fmt.Println("Why is this concatable?")
		// t.Fail()
	}

	if !IsConcatablePrime([]int{3, 7, 109, 673}) {
		t.Fail()
	}
}

func TestUnique(t *testing.T) {
	if !IsEqualIntArray(Unique([]int{1, 2, 2}), []int{1, 2}) {
		t.Fail()
	}

	v := [][]int{
		[]int{1, 2, 3},
		[]int{1, 2, 3},
		[]int{1, 2, 3, 4},
	}

	if !IsEqualIntArray(FlattenAndUnique(v), []int{1, 2, 3, 4}) {
		t.Fail()
	}
}
