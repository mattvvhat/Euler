package main

import (
	"fmt"
)

/*
 * Return the maximum value of a list of integers
 */
func Maximum(list []int) (int, int) {
	maxIndex, maxValue := 0, list[0]

	for i, v := range list {
		if v > maxValue {
			maxIndex, maxValue = i, v
		}
	}

	return maxIndex, maxValue
}

func GreaterThan(lis []int, lowestValue int) []int {
	results := make([]int, 0)
	for _, v := range lis {
		if v > lowestValue {
			results = append(results, v)
		}
	}
	return results
}

func GetNewCandidates(candidates [][]int, primes []int, show bool) [][]int {
	results := make([][]int, 0)

	for _, c := range candidates {
		_, maxVal := Maximum(c)
		_ = maxVal
		for _, p := range GreaterThan(primes, maxVal) {
			newP := []int{p}
			candidate := extend(c, newP)

			if IsConcatablePrime(candidate) {
				if show {
					fmt.Println(candidate)
				}
				results = append(results, candidate)
			}
		}
	}

	return results
}

func main() {
	sieve := NewNaiveSieve(1000)
	PRIMES := sieve.Primes()

	// ...
	primeSingletons := make([][]int, 0)
	for _, p := range PRIMES {
		primeSingletons = append(primeSingletons, []int{p})
	}

	// ...
	primes := GetNewCandidates(primeSingletons, PRIMES, false)
	primes = GetNewCandidates(primes, PRIMES, false)
	primes = GetNewCandidates(primes, PRIMES, false)
	// primes = GetNewCandidates(primes, PRIMES, true)

	fmt.Println(primes)
}
