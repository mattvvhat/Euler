package main

import (
	"fmt"
	primes "github.com/whatever/math/primes"
	_ "math"
	"sort"
)

func Factorial(n int) int {
	result := 1
	for n > 0 {
		result *= n
		n--
	}
	return result
}

func SortedString(s string) string {

	intArray := make([]int, len(s))

	for i, letter := range s {
		intArray[i] = int(letter)
	}

	sort.Ints(intArray)

	newString := ""

	for _, digit := range intArray {
		newString += fmt.Sprintf("%s", string(digit))
	}

	return newString
}

func IsPermutation(lhs, rhs int) bool {
	return SortedString(fmt.Sprintf("%d", lhs)) == SortedString(fmt.Sprintf("%d", rhs))
}

func main() {
	limit := 1000000
	s := primes.NewNaiveSieve(limit)
	hits := make(map[int]float64)

	minIndex := 200000000
	minValue := float64(limit + 1)

	for i := 2; i <= limit; i++ {
		if i%(limit/100) == 0 {
			fmt.Println(i)
		}

		totient := s.Totient(i)
		t := float64(i) / float64(totient)
		hits[i] = t

		if !IsPermutation(totient, i) {
			continue
		}

		if t < minValue {
			minValue = t
			minIndex = i
			fmt.Println(minIndex, minValue, i, totient)
		}
	}

	fmt.Println(minIndex, minValue)
}
