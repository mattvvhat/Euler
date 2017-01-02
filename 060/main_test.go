package main

import (
	"fmt"
	"testing"
)

func TestWhatever(t *testing.T) {
	sieve := NewPrimeSieve(1000000)

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
