package main

import (
	"fmt"
	"github.com/whatever/math"
	"strconv"
	"strings"
)

func Lookup(haystack map[int][]int, needle int) []int {
	matches := make([]int, 0)

	for key, values := range haystack {
		for _, maybe := range values {
			if maybe == needle {
				matches = append(matches, key)
				break
			}
		}
	}

	return matches
}

func extend(lhs, rhs []int) []int {
	result := make([]int, len(lhs)+len(rhs), len(lhs)+len(rhs))

	for i, _ := range lhs {
		result[i] = lhs[i]
	}

	for i, _ := range rhs {
		result[i+len(lhs)] = rhs[i]
	}

	return result
}

func WithinRange(f func(int) int, lowest, highest int) []int {
	results := make([]int, 0)

	i := 0
	val := f(i)

	for {
		if lowest <= val && val <= highest {
			results = append(results, val)
		} else if val > highest {
			break
		}

		// Iterate
		i += 1
		val = f(i)
	}

	return results
}

func FilterContainsZero(list []int) []int {
	results := make([]int, 0)

	for _, v := range list {
		s := fmt.Sprintf("%d", v)
		if !strings.Contains(s, "0") {
			results = append(results, v)
		}
	}

	return results
}

func CyclicMatches(shapeSets map[int][]int, value int) map[int][]int {
	newMap := make(map[int][]int)
	valueString := fmt.Sprintf("%d", value)
	valueTail := valueString[len(valueString)-2:]

	for shapeNumber, values := range shapeSets {
		hits := make([]int, 0)
		hitCount := 0
		for _, number := range values {
			if valueTail == fmt.Sprintf("%d", number)[:2] {
				hitCount += 1
				hits = append(hits, number)
			}

			if len(hits) != hitCount {
				panic("Everything is wrong")
			}
			newMap[shapeNumber] = hits
		}
	}

	return newMap
}

func RemoveKey(shapeSets map[int][]int, value int) map[int][]int {
	newMap := make(map[int][]int)
	for k, v := range shapeSets {
		if k != value {
			newMap[k] = make([]int, len(v))
			copy(newMap[k], v)
		}
	}

	return newMap
}

func CopyMap(src map[int][]int) map[int][]int {
	return nil
}

func GetIndex(list []int, value int) int {
	for i, v := range list {
		if v == value {
			return i
		}
	}
	return -1
}

func TotalMapLength(m map[int][]int) int {
	size := 0
	for _, v := range m {
		size += len(v)
	}
	return size
}

func main() {
	// Setup Mapping between side count and available 4-digit numbers

	shapeNumbers := make(map[int][]int)

	for i := 3; i <= 8; i++ {
		shapeNumbers[i] = FilterContainsZero(WithinRange(shapeFunc(i), 1000, 9999))
	}

	// Create Graph
	graph := math.NewGraph()

	nodes := make(map[int]bool)

	// Add each node
	fmt.Println("Adding nodes")
	for _, numbers := range shapeNumbers {
		for _, n := range numbers {
			if _, ok := nodes[n]; !ok {
				nodes[n] = true
				graph.AddNode(fmt.Sprintf("%d", n))
			}
		}
	}

	// Add each edge
	fmt.Println("Adding edges")
	for figurateNumber, numbers := range shapeNumbers {
		for _, n := range numbers {
			for _, neighborList := range CyclicMatches(RemoveKey(shapeNumbers, figurateNumber), n) {
				for _, neighbor := range neighborList {
					nLabel := fmt.Sprintf("%d", n)
					neighborLabel := fmt.Sprintf("%d", neighbor)
					graph.AddConnection(nLabel, neighborLabel)
				}
			}
		}
	}

	candidates := make([][]string, 0)

	// Produce cycles
	fmt.Println("Printing cycles")
	for _, n := range shapeNumbers[8] {
		nLabel := fmt.Sprintf("%d", n)
		for _, c := range graph.GetCyclesFor(nLabel, 6) {
			if len(c) == 7 {
				candidates = append(candidates, c)
			}
		}
	}

	for _, c := range candidates {
		s := 0
		for _, n := range c[1:] {
			i, _ := strconv.Atoi(n)
			fmt.Print(Lookup(shapeNumbers, i))
			s += i
		}
		fmt.Println(" =>", c, "|", s)
	}

	fmt.Println("^^^^^ Look above to find the path that contains the numbers 3-8 exactly once")
}
