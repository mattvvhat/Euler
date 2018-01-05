package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Return Triangle from Text Line
func LoadTriangleFromLine(line string) Triangle {
	pieces := strings.Split(line, ",")
	numbers := [6]float64{}

	for i, val := range pieces {
		num, _ := strconv.ParseFloat(val, 64)
		numbers[i] = num
	}

	return Triangle{
		Point{numbers[0], numbers[1]},
		Point{numbers[2], numbers[3]},
		Point{numbers[4], numbers[5]},
	}
}

func main() {
	triangleFile, _ := os.Open("p102_triangles.txt")
	scanner := bufio.NewScanner(triangleFile)
	scanner.Split(bufio.ScanLines)

	count := 0

	for scanner.Scan() {
		tri := LoadTriangleFromLine(scanner.Text())
		if tri.contains(Point{0, 0}) {
			count += 1
		}
	}

	fmt.Println("count =", count)
}
