package main

import (
	"testing"
)

func TestMapInts(t *testing.T) {
	xs := []int{1, 2, 3, 4}
	ys := MapInts(
		xs,
		func(n int) int {
			return n * n
		},
	)

	if ys[0] != 1 {
		t.Fail()
	}

	if ys[1] != 4 {
		t.Fail()
	}

	if ys[3] != 16 {
		t.Fail()
	}

	if len(ys) != 4 {
		t.Fail()
	}
}

func TestDigits(t *testing.T) {
	a := Digits(4)
	if len(a) != 1 {
		t.Fail()
	}

	b := Digits(230)
	if len(b) != 3 {
		t.Fail()
	}

	if b[0] != 2 {
		t.Fail()
	}

	if b[1] != 3 {
		t.Fail()
	}
}

func TestSquareDigits(t *testing.T) {
	if SquareDigits(1) != 1 {
		t.Fail()
	}

	if SquareDigits(85) != 89 {
		t.Fail()
	}

	if SquareDigits(13) != 10 {
		t.Fail()
	}
}

func TestBecomes(t *testing.T) {
	if Becomes(1) != 1 {
		t.Fail()
	}

	if Becomes(10) != 1 {
		t.Fail()
	}

	if Becomes(16) != 89 {
		t.Fail()
	}
}
