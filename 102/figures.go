package main

type Point struct {
	x float64
	y float64
}

type Line struct {
	m float64
	o float64
}

func line(a, b Point) Line {
	m := (a.y - b.y) / (a.x - b.x)
	o := m*a.x - a.y
	return Line{m, o}
}

type Triangle struct {
	a Point
	b Point
	c Point
}

func (self *Triangle) contains(p Point) bool {
	// This is the signed value of the z-coordinate
	// of the cross product between two vectors
	sign := func(a, b, c Point) bool {
		return (a.x-c.x)*(b.y-c.y)-(b.x-c.x)*(a.y-c.y) < 0.0
	}

	in1 := sign(p, self.a, self.b)
	in2 := sign(p, self.b, self.c)
	in3 := sign(p, self.c, self.a)

	return in1 == in2 && in2 == in3
}
