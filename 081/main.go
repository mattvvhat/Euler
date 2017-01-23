package main

import (
	"fmt"
	g "github.com/whatever/math/graphs"
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
		805, 732, 524, 37, 331,
	)

	graph := g.NewWeightedUGraph()

	// Add Vertices
	graph.AddVertex("start")

	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			graph.AddVertex(fmt.Sprintf("%d,%d", i, j))
		}
	}

	// Add Edges
	graph.AddEdge("start", "0,0", float64(m.Get(0, 0)))

	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {

			current := fmt.Sprintf("%d,%d", i, j)
			weight := float64(m.Get(i, j))

			if i > 0 {
				above := fmt.Sprintf("%d,%d", i-1, j)
				graph.AddEdge(above, current, weight)
			}

			if j > 0 {
				left := fmt.Sprintf("%d,%d", i, j-1)
				graph.AddEdge(left, current, weight)
			}

			fmt.Println(fmt.Sprintf("%d,%d = %d", i, j, m.Get(i, j)))
		}
	}

	distances := graph.ShortestDistances("start")

	// fmt.Println(distances)
	// fmt.Println(distances["4,4"])
	// fmt.Println(distances["1,1"])
	fmt.Println(distances["4,4"])

	fmt.Println("...")
}
