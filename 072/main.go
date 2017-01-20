package main

import (
	"fmt"
	primes "github.com/whatever/math/primes"
)

func main() {
	limit := 1000000
	s := primes.NewNaiveSieve(limit)
	total := 0

	for i := 2; i <= limit; i++ {
		total += s.Totient(i)
	}

	fmt.Println(total)
}
