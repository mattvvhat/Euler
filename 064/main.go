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

type ContinuedFraction struct {
	LeadingValue   int
	RepeatingValue []int
}

func ContinuedFractionOfRoot(n int) ContinuedFraction {
	result.LeadingValue = int(math.Sqrt(float64(n)))
	return result
}

func main() {
	fmt.Println(ContinuedFractionOfRootFinite(2, 10))
	fmt.Println(ContinuedFractionOfRootFinite(23, 10))
	fmt.Println(ContinuedFractionOfRootFinite(23, 10))
}
