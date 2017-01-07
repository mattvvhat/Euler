package main

import (
	"testing"
)

func TestGcd(t *testing.T) {
	if Gcd(3, 7) != 1 && Gcd(7, 3) != 1 {
		t.Fail()
	}
	if Gcd(4, 2) != 2 && Gcd(2, 4) != 2 {
		t.Fail()
	}
	if Gcd(13, 169) != 13 {
		t.Fail()
	}
	if Gcd(3*7*13, 89*13) != 13 {
		t.Fail()
	}
	if Gcd(0, 1) != 1 || Gcd(1, 0) != 1 {
		t.Fail()
	}

}

func TestRational(t *testing.T) {
	a := Rational{1, 2}
	if a.n != 1 || a.d != 2 {
		t.Fail()
	}

	b := Rational{1, 2}
	c := a.Add(b)
	if c.n != 1 || c.d != 1 {
		t.Fail()
	}

	d := Rational{12, 20}
	d.Simplify()
	if d.n != 3 && d.d != 5 {
		t.Fail()
	}

	e := Rational{-3, -4}
	e.Simplify()
	if e.n != 3 && e.d != 4 {
		t.Fail()
	}

	f := Rational{30, -40}
	f.Simplify()
	if f.n != -3 || f.d != 4 {
		t.Fail()
	}

	g := e.Mul(e)
	if g.n != 9 || g.d != 16 {
		t.Fail()
	}

	h := Rational{1, -3}.Mul(Rational{3, 1})
	if h.n != -1 || h.d != 1 {
		t.Fail()
	}

	if (Rational{1, 3}.Inverse().n) != 3 {
		t.Fail()
	}
}
