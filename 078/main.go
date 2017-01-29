package main

import (
	"fmt"
)

/*
	Return list of integer
*/
func PartitionCount(n, k int) int {
	switch {
	case n <= 0:
		return 0
	case k == 1:
		return 1
	case k == n-1:
		return 1
	case n == k:
		return 1
	}

	return PartitionCount(n-1, k-1) + PartitionCount(n-k, k)
}

func CountPartitions(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += PartitionCount(n, i)
	}

	return total
}

type Lookup struct {
	table map[int]map[int]int
}

func NewLookup() Lookup {
	return Lookup{make(map[int]map[int]int)}
}

func (s *Lookup) Get(i, j int) (int, bool) {
	if _, ok := s.table[i]; !ok {
		s.table[i] = make(map[int]int)
	}

	if val, ok := s.table[i][j]; ok {
		return val, true
	}

	return 0, false
}

func (s *Lookup) Set(i, j, val int) {
	if _, ok := s.table[i]; !ok {
		s.table[i] = make(map[int]int)
	}

	s.table[i][j] = val
}

func MemoizePartitionCount(n int, seed *Lookup) int {

	// Starts with none
	total := 0

	for i := 1; i <= n; i++ {
		if v, ok := seed.Get(n, i); ok {
			total += v
		} else {
			t := PartitionCount(n, i)
			seed.Set(n, i, t)
			total += t
		}
	}

	return total
}

func main() {
	// counts := make(map[int]int)

	seed := NewLookup()

	for i := 1; i < 1000; i++ {
		v := MemoizePartitionCount(i, &seed)
		fmt.Println("Testing", i, "->", v)
		if v%1000000 == 0 {
			fmt.Println(v)
			break
		}
	}
}
