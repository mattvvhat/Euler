package main

import (
	"fmt"
)

func main() {
	n := 3
	d := 7
	ayy := 3.0 / 7

	maybes := []Rational{}

	for i := 2; d*i <= 1000000; i++ {
		r := Rational{n*i - 1, d * i}
		r.Simplify()
		maybes = append(maybes, r)
	}

	smallestInd := 0
	smallestVal := maybes[smallestInd]
	smallestRat := -1.0
	cnt := 0

	for i, v := range maybes {
		if v.d >= 1000000 {
			continue
		}

		lhs := float64(v.n) / float64(v.d)
		rhs := float64(smallestVal.n) / float64(smallestVal.d)

		if lhs >= rhs {
			v.Simplify()
			smallestInd = i
			smallestVal = v
			smallestRat = lhs
			cnt++
		}
	}

	fmt.Println(smallestInd, smallestVal, smallestRat-ayy)
}
