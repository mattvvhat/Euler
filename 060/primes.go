package main

import (
	"math"
	"sort"
)

type NaiveSieve struct {
	sieve map[int]bool
}

func NewNaiveSieve(limit int) NaiveSieve {

	primes := make(map[int]bool)

	// Generate list of primes
	for n := 2; n < limit; n += 1 {
		if IsNaivePrime(n) {
			primes[n] = true
		}
	}

	return NaiveSieve{primes}
}

func (self *NaiveSieve) Contains(n int) bool {
	_, ok := self.sieve[n]
	return ok
}

func (self *NaiveSieve) Primes() []int {
	primes := make([]int, 0, len(self.sieve))
	for n, _ := range self.sieve {
		primes = append(primes, n)
	}
	sort.Ints(primes)
	return primes
}

func IsNaivePrime(n int) bool {
	switch {
	case n == 1:
		return false
	case n == 2:
		return true
	case n%2 == 0:
		return false
	}

	limit := int(math.Sqrt(float64(n)))

	for i := 3; i < limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
