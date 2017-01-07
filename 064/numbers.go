package main

import (
	"fmt"
)

/**
 * Return Greatest Common Division of two integers
 */

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

/**
 * Rational Number
 */

type Rational struct {
	n int
	d int
}

func (self *Rational) Simplify() *Rational {
	gcd := Gcd(self.n, self.d)
	self.n /= gcd
	self.d /= gcd

	switch {
	case self.n < 0 && self.d < 0:
		self.n *= -1
		self.d *= -1
	case self.n > 0 && self.d < 0:
		self.n *= -1
		self.d *= -1
	case self.d == 0:
		panic("Divide by zero error!")
	}
	return self
}

func (lhs Rational) Add(rhs Rational) Rational {
	num := lhs.n*rhs.d + rhs.n*lhs.d
	den := lhs.d * rhs.d
	result := Rational{num, den}
	result.Simplify()
	return result
}

func (lhs Rational) Mul(rhs Rational) Rational {
	result := Rational{lhs.n * rhs.n, lhs.d * rhs.d}
	result.Simplify()
	return result
}

func (self Rational) Inverse() Rational {
	return Rational{self.d, self.n}
}

/**
 * Rational number + square number
 */

type SquareNumber struct {
	Rat Rational
	Sqr int
}

func _() {
	fmt.Println("_")
}
