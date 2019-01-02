package main

import (
	"fmt"
)

// Vertex represents a vertex on a graph and its connctions to its neihbors
type Vertex struct {
	Id        int
	Value     int
	Neighbors map[int]*Vertex
}

// AddNeighbor
func (self *Vertex) AddNeighbor(dst *Vertex) {

	if self.Neighbors == nil {
		self.Neighbors = make(map[int]*Vertex)
	}

	if _, found := self.Neighbors[dst.Id]; found {
		return
	}

	// fmt.Printf("!!!!! %d %d\n", self.Id, dst.Id)
	self.Neighbors[dst.Id] = dst
}

// Graph represents the connections between vertices
type Graph struct {
	Vertices map[int]*Vertex
}

// NewGraphs constructs an empty graph
func NewGraph() Graph {
	return Graph{make(map[int]*Vertex)}
}

// AddVertex adds a node with matching id and value
func (self *Graph) AddVertex(id, value int) {
	_, found := self.Vertices[id]
	if !found {
		v := Vertex{id, value, make(map[int]*Vertex)}
		self.Vertices[id] = &v
	} else {
		panic("we already have that node")
	}
}

// AddEdge connects 2 existings edges
func (self *Graph) AddEdge(lhs, rhs int) {
	src, found_src := self.Vertices[lhs]
	dst, found_dst := self.Vertices[rhs]

	if !found_src {
		panic(fmt.Sprintf("Missing %d", lhs))
	}

	if !found_dst {
		panic(fmt.Sprintf("Missing %d", lhs))
	}

	src.AddNeighbor(dst)
}
