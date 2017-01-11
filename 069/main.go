package main

import (
	"fmt"
	_ "github.com/whatever/math/primes"
)

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	LIMIT := 1000000
	hits := make(map[int]int)

	for i := 1; i <= LIMIT; i++ {
		for j := 1; j < i; j++ {
			if Gcd(i, j) == 1 {
				hits[i]++
			}
		}
	}

	highestInd := 1
	highestVal := float64(hits[highestInd])

	for i, v := range hits {
		trial := float64(i) / float64(v)

		if highestVal < trial {
			highestInd = i
			highestVal = trial
		}
	}

	fmt.Println(highestInd, highestVal)
}
