package main

import (
	"fmt"
	combo "github.com/whatever/math/combinatorics"
	"math"
	"sort"
	"strconv"
)

func Indices(s string) []int {
	indices := make([]int, len(s))
	for i, _ := range s {
		indices[i] = i
	}
	return indices
}

func IsCube(n int) bool {
	f := math.Cbrt(float64(n))
	c := int(f)

	if c*c*c == n {
		return true
	} else {
		return false
	}
}

func Permutations(s string) []string {
	results := make([]string, 0)
	indices := Indices(s)

	for _, p := range combo.P(indices, len(indices)) {
		permutationString := make([]rune, len(s))
		for i, v := range p {
			permutationString[i] = rune(s[v])
		}
		results = append(results, string(permutationString))
	}

	return results
}

func IntegerPermutations(n int) []int {
	results := make([]int, 0)

	for _, s := range Permutations(fmt.Sprintf("%d", n)) {
		n, err := strconv.Atoi(s)
		if err == nil {
			results = append(results, n)
		}
	}

	return results
}

func SortedString(s string) string {

	intArray := make([]int, len(s))

	for i, letter := range s {
		intArray[i] = int(letter)
	}

	sort.Ints(intArray)

	newString := ""

	for _, digit := range intArray {
		newString += fmt.Sprintf("%s", string(digit))
	}

	return newString
}

func CubeMapKey(n int) string {
	return SortedString(fmt.Sprintf("%d", n))
}

func IsPermutation(lhs, rhs int) bool {
	return true
}

func main() {
	cubePermutations := make(map[string][]int)

	for i := 0; i < 10000; i++ {
		cube := i * i * i
		key := CubeMapKey(cube)
		cubePermutations[key] = append(cubePermutations[key], cube)
	}

	for _, v := range cubePermutations {
		if len(v) == 5 {
			fmt.Println(v)
		}
	}

	fmt.Println("^^^ Just inspect to the smallest left side value")
}
