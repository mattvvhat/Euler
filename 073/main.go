package main

import (
	"fmt"
	primes "github.com/whatever/math/primes"
)

func RationalString(r Rational) string {
	r.Simplify()
	return fmt.Sprintf("%d/%d", r.n, r.d)
}

func GetCoprimes(n int, s *primes.NaiveSieve) []int {
	coprimes := []int{}

	for i := 1; i < n; i++ {
		if Gcd(n, i) == 1 {
			coprimes = append(coprimes, i)
		}
	}

	return coprimes
}

func main() {
	limit := 12000
	s := primes.NewNaiveSieve(limit)
	hits := make(map[string]bool)

	for i := 2; i <= limit; i++ {
		coprimes := GetCoprimes(i, &s)
		for _, v := range coprimes {

			switch {
			case 3*v <= i:
				continue
			case 2*v >= i:
				continue
			}

			val := fmt.Sprintf("%d/%d", v, i)
			hits[val] = true
		}
	}

	fmt.Println(len(hits))
}
