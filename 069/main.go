package main

import (
	"fmt"
	primes "github.com/whatever/math/primes"
)

func Totient(number int, s primes.NaiveSieve) int {
	result := 1
	n := 1
	d := 1

	for p, _ := range s.PrimeFactorize(number) {
		n *= p - 1
		d *= p
	}
	result = number * n / d
	return result
}

func Max(v map[int]float64) (int, float64) {
	maxIndex := -1
	maxValue := -1.0

	for i, v := range v {
		if v > maxValue {
			maxValue = v
			maxIndex = i
		}
	}

	return maxIndex, maxValue
}

func main() {
	limit := 1000000
	s := primes.NewNaiveSieve(limit)

	totients := make(map[int]float64)

	for i := 2; i <= limit; i++ {
		if i%1000 == 0 {
			fmt.Println(i)
		}
		v := Totient(i, s)
		totients[i] = float64(i) / float64(v)
	}

	largestIndex, largestValue := Max(totients)

	fmt.Println(largestIndex, largestValue)
}
