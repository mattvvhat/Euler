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

func IsConcatablePrime(primes []int) bool {
	for _, v := range P(primes, 2) {
		testablePrime := ConcatenateInteger(v[0], v[1])
		if !IsNaivePrime(testablePrime) {
			return false
		}
	}

	return true
}

func Unique(list []int) []int {
	resultMap := make(map[int]bool)

	for _, v := range list {
		resultMap[v] = true
	}

	results := make([]int, 0, len(resultMap))
	for v, _ := range resultMap {
		results = append(results, v)
	}

	return results
}

func FlattenAndUnique(lists [][]int) []int {

	size := 0
	for _, v := range lists {
		size += len(v)
	}

	flattened := make([]int, 0, size)

	for _, v := range lists {
		for _, u := range v {
			flattened = append(flattened, u)
		}
	}

	return Unique(flattened)
}
