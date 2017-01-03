package main

import (
	"fmt"
	"strconv"
)

func ConcatenateInteger(lhs, rhs int) int {
	s := fmt.Sprintf("%d%d", lhs, rhs)
	val, err := strconv.Atoi(s)

	if err == nil {
		return val
	} else {
		fmt.Println("This happened:", lhs, rhs)
		return 0
	}
}

func IsConcatablePrime(primes []int, _ NaiveSieve) bool {
	for _, v := range P(primes, 2) {
		testablePrime := ConcatenateInteger(v[0], v[1])
		if !IsNaivePrime(testablePrime) {
			return false
		}
	}

	return true
}

func main() {
	sieve := NewNaiveSieve(1000000)
	primes := sieve.Primes()
	LIMIT := 200

	for i := 1 + 0; i < LIMIT; i++ {
		fmt.Println("i ->", primes[i])
		for j := i + 1; j < LIMIT; j++ {
			for k := j + 1; k < LIMIT; k++ {
				for m := k + 1; m < LIMIT; m++ {
					for n := m + 1; n < LIMIT; n++ {

						candidates := []int{
							primes[i],
							primes[j],
							primes[k],
							primes[m],
							primes[n],
						}

						if IsConcatablePrime(candidates, sieve) {
							fmt.Println(candidates)
						}
					}
				}
			}
		}
	}
}
