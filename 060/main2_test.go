package main

import (
	"fmt"
	"testing"
)

func TestGetNewCandidates(t *testing.T) {
	sieve := NewNaiveSieve(5000)

	// ...
	primeSingletons := make([][]int, 0)
	for _, p := range sieve.Primes() {
		primeSingletons = append(primeSingletons, []int{p})
	}

	// ...
	primes := GetNewCandidates(primeSingletons, sieve.Primes())
	primes = GetNewCandidates(primes, sieve.Primes())
	primes = GetNewCandidates(primes, sieve.Primes())
	primes = GetNewCandidates(primes, sieve.Primes())

	fmt.Println(primes)

	if false {
		t.Fail()
	}
}
