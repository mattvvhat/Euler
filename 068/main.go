package main

import (
	"log"
	"strconv"
	"strings"
)

// Link is
type Link struct {
	Value int
	Next  *Link
}

// Ring5 is a potential 5-gon Ring
type Ring5 struct {
	Sides [5]*Link
}

// DecodeToRing5 returns a Ring-5 5-gone from a string
func DecodeToRing5(s string) Ring5 {
	results := Ring5{}

	// Network topology is a little bit easier to just make explicitly
	var values [5][3]int64

	// Make into something like a matrix of values
	for i, blob := range strings.Split(s, ";") {
		for j, v := range strings.Split(blob, ",") {
			value, err := strconv.ParseInt(v, 10, 64)
			if err == nil {
				values[i][j] = value
			} else {
				panic("value isn't right")
			}
		}
	}

	// ...
	log.Println(values)

	return results
}

// Return
func GetBlankRing5() Graph {
	g := NewGraph()

	// Add roots
	for i := 0; i < 5; i++ {
		g.AddVertex(i, i)
		g.AddVertex(i+10, i+10)
	}

	// Connect roots to first child
	for i := 0; i < 5; i++ {
		g.AddEdge(i, i+10)
	}

	// Connect internal layer
	// NOTE: This is hacky but doesn't really matter
	for i := 10; i < 15; i++ {
		if i+1 == 15 {
			g.AddEdge(i, 10)
		} else {
			g.AddEdge(i, i+1)
		}
	}

	return g
}

func GetItem(vertices map[int]*Vertex) *Vertex {
	for _, v := range vertices {
		return v
	}
	return nil
}

// Sum 3-vertex path from adjacent points
func Sum3(v *Vertex) (total int) {
	for i := 0; i < 3; i++ {
		total += v.Value
		v = GetItem(v.Neighbors)
	}
	return total
}

// main
func main() {
	g := GetBlankRing5()

	for id := 0; id < 5; id++ {
		log.Printf("%d: Sum3 = %d", id, Sum3(g.Vertices[id]))
	}
}
