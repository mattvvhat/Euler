package main

import (
	"fmt"
)

func extend(lhs, rhs []int) []int {
	result := make([]int, len(lhs)+len(rhs), len(lhs)+len(rhs))

	for i, _ := range lhs {
		result[i] = lhs[i]
	}

	for i, _ := range rhs {
		result[i+len(lhs)] = rhs[i]
	}

	return result
}

func factorial(n int) int {
	result := 1
	for i := 1; i < n; i += 1 {
		result *= i
	}
	return result
}

func C(lis []int, k int) [][]int {
	var result [][]int

	switch {
	case k == 0:
		return make([][]int, 0, 0)
	case k == 1:
		result = make([][]int, len(lis), len(lis))

		for i, v := range lis {
			u := []int{v}
			result[i] = make([]int, 1, 1)
			copy(result[i], u)
		}

		return result
	}
	return nil
}

// Recursive
func combinations(n int, k int, c *chan []int) {
	if k == 0 {
		return
	}

	if n == 0 {
		return
	}

	for i := 0; i < n; i++ {
		x := extend([]int{}, []int{1, 2, 3})
		*c <- x
	}
}

// Return a channel yielding array of indices
func Combinations(n, k int) chan []int {
	length := factorial(n) / factorial(n-k) / factorial(k)
	results := make(chan []int, length)

	go func() {
		combinations(n, k, &results)
		close(results)
	}()

	return results
}

func _() {
	fmt.Println("shutup")
}
