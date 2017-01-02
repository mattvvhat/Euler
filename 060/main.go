package main

import (
	"fmt"
)

func main() {
	sieve := NewNaiveSieve(1000000)
	primes := sieve.Primes()

	fmt.Println(primes)
}
