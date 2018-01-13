package main

import (
	"fmt"
	"math"
)

func Gcd(a, b uint) uint {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

type Fraction struct {
	n, d uint
}

func NewFraction(n, d uint) Fraction {
	div := Gcd(n, d)
	return Fraction{n / div, d / div}
}

type ContinuedRepeatingFraction struct {
	Leading uint
	Digits  []uint
}

func ContinuedFractionOfSquare(value uint) ContinuedRepeatingFraction {
	d := float64(value)
	digits := []uint{}

	r := uint(math.Floor(math.Sqrt(d)))
	leading := r

	// Start iteration
	var a, p, q uint = r, 0, 1

	// Catch perfect squares
	if r*r == value {
		return ContinuedRepeatingFraction{r, []uint{}}
	}

	for {
		p = a*q - p
		q = (value - p*p) / q
		a = (r + p) / q

		digits = append(digits, a)

		if q == 1 {
			break
		}
	}

	return ContinuedRepeatingFraction{leading, digits}
}

func main() {
	odd_periods := 0

	for i := 2; i <= 10000; i += 1 {
		cf := ContinuedFractionOfSquare(uint(i))
		length := len(cf.Digits)
		if length%2 != 0 {
			odd_periods += 1
		}
	}

	fmt.Println(odd_periods)
}
