package main

import (
	"fmt"
	"math"
	"sort"
)

type PrimeSieve struct {
	sieve map[int]bool
}

func NewPrimeSieve(limit int) PrimeSieve {

	primes := make(map[int]bool)

	// Generate list of primes
	for n := 2; n < limit; n += 1 {
		if IsNaivePrime(n) {
			primes[n] = true
		}
	}

	return PrimeSieve{primes}
}

func (self *PrimeSieve) Contains(n int) bool {
	_, ok := self.sieve[n]
	return ok
}

func (self *PrimeSieve) Primes() []int {
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

func main() {
	sieve := NewPrimeSieve(1000000)
	primes := sieve.Primes()

	for i := 0; i < len(primes); i += 1 {
		for j := 0; j < len(primes); j += 1 {
			for k := 0; k < len(primes); k += 1 {
				for m := 0; m < len(primes); m += 1 {
					for n := 0; m < len(primes); m += 1 {
						// ...
					}
				}
			}
		}
	}

	fmt.Println(primes)
}
