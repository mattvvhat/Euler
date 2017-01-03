package main

import (
	"fmt"
	"sort"
)

func IntArrayString(values []int) string {
	var s string

	for _, v := range values {
		s += fmt.Sprintf("%d", v)
		s += ","
	}

	return s
}

func UniqueIntArrayArray(values [][]int) [][]int {
	matches := make(map[string][]int)

	result := make([][]int, 0)

	for _, array := range values {
		sort.Ints(array)
		matches[IntArrayString(array)] = array
	}

	for _, value := range matches {
		result = append(result, value)
	}

	return result
}

func minSum(values [][]int) (int, []int) {
	minVal := values[0]
	minSum := 10000000000
	for _, v := range values {
		sum := 0
		for _, n := range v {
			sum += n
		}

		if sum < minSum {
			minSum = sum
			minVal = v
		}
	}

	return minSum, minVal
}

func main2() {
	sieve := NewNaiveSieve(5000)
	primes := sieve.Primes()

	// maps number of values to set of matches
	candidatesN := make([][][]int, 6)
	candidatesN[0] = make([][]int, 0)
	candidatesN[1] = make([][]int, 0)
	candidatesN[2] = make([][]int, 0)
	candidatesN[3] = make([][]int, 0)
	candidatesN[4] = make([][]int, 0)

	for _, p := range primes {
		candidatesN[0] = append(candidatesN[0], []int{p})
	}

	for _, v := range candidatesN[0] {
		for _, p := range primes {
			trial := append(v, p)
			if IsConcatablePrime(trial) {
				candidatesN[1] = append(candidatesN[1], trial)
			}
		}
	}

	candidatesN[1] = UniqueIntArrayArray(candidatesN[1])

	for _, v := range candidatesN[1] {
		for _, p := range primes {
			trial := append(v, p)
			if IsConcatablePrime(trial) {
				candidatesN[2] = append(candidatesN[2], trial)
			}
		}
	}

	candidatesN[2] = UniqueIntArrayArray(candidatesN[2])

	for _, v := range candidatesN[2] {
		for _, p := range primes {
			trial := make([]int, 1+len(v))
			copy(trial, v)
			trial[len(trial)-1] = p
			if IsConcatablePrime(trial) {
				candidatesN[3] = append(candidatesN[3], trial)
			}
		}
	}

	candidatesN[3] = UniqueIntArrayArray(candidatesN[3])

	for _, v := range candidatesN[3] {
		for _, p := range primes {
			trial := make([]int, 1+len(v))
			copy(trial, v)
			trial[len(trial)-1] = p
			if IsConcatablePrime(trial) {
				candidatesN[4] = append(candidatesN[4], trial)
			}
		}
	}

	fmt.Println(candidatesN[4])
	fmt.Println(len(candidatesN[4]))
	fmt.Println(minSum(candidatesN[4]))
}
