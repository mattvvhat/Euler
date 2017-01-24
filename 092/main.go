package main

import (
	"fmt"
	"strconv"
)

func MapInts(xs []int, f func(int) int) []int {
	ys := make([]int, len(xs))

	for i, v := range xs {
		ys[i] = f(v)
	}

	return ys
}

func Digits(n int) []int {
	d := fmt.Sprintf("%d", n)
	digits := make([]int, len(d))

	for i, c := range d {
		digits[i], _ = strconv.Atoi(string(c))
	}

	return digits
}

func SquareDigits(n int) int {
	// Get squares
	digits := MapInts(
		Digits(n),
		func(x int) int {
			return x * x
		},
	)

	sum := 0

	for _, v := range digits {
		sum += v
	}

	return sum
}

func Becomes(n int) int {
	if n < 1 {
		return -1
	}

	for {
		switch n {
		case 89:
			return 89
		case 1:
			return 1
		}
		n = SquareDigits(n)
	}
}

func main() {

	limit := 10000000

	arrivedAt89 := 0

	for i := 0; i < limit; i++ {
		if Becomes(i) == 89 {
			arrivedAt89++
		}
	}

	fmt.Println(arrivedAt89)
}
