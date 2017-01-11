package main

import (
	"fmt"
	primes "github.com/whatever/math/primes"
	"testing"
)

func TestTotient(t *testing.T) {

	_ = primes.NewNaiveSieve(1000)

	expected := map[int]int{
		2:  1,
		3:  1,
		4:  2,
		5:  4,
		6:  2,
		7:  6,
		8:  4,
		9:  6,
		10: 4,
	}

	for n, totient := range expected {
		v := Totient(n)
		if v != totient {
			fmt.Printf("T(%d) = %d != %d\n", n, v, totient)
			t.Fail()
		}
	}
}
