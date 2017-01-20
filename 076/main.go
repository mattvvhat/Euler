package main

import (
	"fmt"
)

func _() {
	fmt.Println("...")
}

func Min(digits []int) int {
	smallest := digits[0]
	for _, v := range digits {
		if smallest < v {
			smallest = v
		}
	}
	return smallest
}

func Max(digits []int) int {
	largest := digits[0]
	for _, v := range digits {
		if largest < v {
			largest = v
		}
	}
	return largest
}

func Largest(digits [][]int) int {
	largest := digits[0][0]
	for _, m := range digits {
		n := Max(m)
		if n > largest {
			largest = n
		}
	}
	return largest
}

/*
	Return list of integer
*/
func Partitions(n, s int) [][]int {

	switch {
	case n == 1:
		return [][]int{[]int{s}}
	case s < n:
		return [][]int{}
	case s == n:
		result := [][]int{
			make([]int, n, n),
		}
		for i, _ := range result[0] {
			result[0][i] = 1
		}
		return result
	}

	results := [][]int{}

	for _, p := range Partitions(n-1, s-1) {
		results = append(results, p)
	}

	for _, p := range Partitions(n-s, s) {
		results = append(results, p)
	}

	return results
}

func PartitionCount(n, k int) int {
	switch {
	case n <= 0:
		return 0
	case k == 1:
		return 1
	case k == n-1:
		return 1
	case n == k:
		return 1
	}

	return PartitionCount(n-1, k-1) + PartitionCount(n-k, k)
}

func main() {
}
