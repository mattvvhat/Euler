package main

import (
	"fmt"
)

func factorial(n int) int {
	result := 1
	for i := 1; i < n; i += 1 {
		result *= i
	}
	return result
}

func combinations() {
}

/**
 * Return a channel yielding array of indices
 */
func Combinations(n, k int) chan []int {
	results := make(chan []int, 10)
	results <- []int{1, 2, 3}
	close(results)
	return results
}

func _() {
	fmt.Println("shutup")
}
