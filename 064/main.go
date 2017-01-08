package main

import (
	"fmt"
	"math"
)

func ContinuedFractionOfRootFinite(n, places int) []int {
	results := make([]int, places)

	x := math.Sqrt(float64(n))
	r0 := x
	a0 := math.Floor(r0)

	results[0] = int(a0)

	for i := 1; i < places; i++ {
		r1 := 1 / (r0 - a0)
		a1 := math.Floor(r1)

		results[i] = int(a1)

		r0 = r1
		a0 = a1
	}

	return results
}

func ContinuedFractionOfRoot(n int) []int {
	results := make([]int, 1)
	matches := make(map[float64]bool)

	x := math.Sqrt(float64(n))
	r0 := x
	a0 := math.Floor(r0)

	results[0] = int(a0)
	results[0] = int(a0)
	matches[r0] = true

	for {
		r1 := 1 / (r0 - a0)
		a1 := math.Floor(r1)

		fmt.Println(r1, a1)

		if z, exists := matches[r1]; exists {
			fmt.Println(int(r1), r1, z)
			break
		}

		results = append(results, int(a1))

		r0 = r1
		a0 = a1
	}

	return results
}

func main() {
	fmt.Println(ContinuedFractionOfRootFinite(2, 10))
	fmt.Println(ContinuedFractionOfRootFinite(23, 10))
	fmt.Println(ContinuedFractionOfRootFinite(23, 10))
}
