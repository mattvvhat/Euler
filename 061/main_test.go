package main

import (
	"fmt"
	"strings"
	"testing"
)

func ToString(list []int) string {
	resultArray := make([]string, len(list))

	for i, v := range list {
		resultArray[i] = fmt.Sprintf("%d", v)
	}

	return strings.Join(resultArray, ",")
}

func TestToString(t *testing.T) {
	switch {
	case ToString([]int{1}) != "1":
		t.Fail()
	case ToString([]int{1, 2, 3, 5}) != "1,2,3,5":
		t.Fail()
	}
}

func TestWithinRange(t *testing.T) {
	if ToString(WithinRange(tri, 1, 10)) != "1,3,6,10" {
		t.Fail()
	}
}

func TestSimple(t *testing.T) {
	candidates := map[int][]int{
		3: []int{3344},
		4: []int{4455},
		5: []int{5566},
		6: []int{6677},
		7: []int{7788},
		8: []int{8833},
	}

	values := RecursiveSearch(candidates)

	for _, path := range values {
		if len(path) == 6 {
			fmt.Println(path)
		}
	}
	fmt.Println(values)
}
