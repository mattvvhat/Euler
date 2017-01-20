package main

import (
	"fmt"
	primes "github.com/whatever/math/primes"
	"testing"
)

func TestGetCoprimes(t *testing.T) {
	s := primes.NewNaiveSieve(1000)
	a := GetCoprimes(8, &s)
	fmt.Println(a)
}
