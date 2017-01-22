package main

import (
	"fmt"
)

type Matrix struct {
	Height int
	Width  int
	Values []int
}

func NewMatrix(height, width int) Matrix {
	return Matrix{height, width, make([]int, width*height)}
}

func (m *Matrix) Apply(values ...int) {
	if len(values) != len(m.Values) {
		panic("Err")
	}

	for i, v := range values {
		m.Values[i] = v
	}
}

func (m *Matrix) Get(i, j int) int {
	return m.Values[i*m.Width+j]
}

func ShortestDistances(m *Matrix) [][]int {
	distances := make([][]int, m.Width*m.Height)
	return distances
}

func main() {
	m := NewMatrix(5, 5)
	m.Apply(
		131, 673, 234, 103, 18,
		201, 96, 342, 965, 150,
		630, 803, 746, 422, 111,
		537, 699, 497, 121, 956,
		807, 732, 524, 37, 331,
	)

	fmt.Println("...", m)
	fmt.Println(m.Get(2, 0))
	fmt.Println(m.Get(4, 4))
}
