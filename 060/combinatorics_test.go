package main

import (
	"fmt"
	"testing"
)

func TestX(t *testing.T) {
	c := Combinations(5, 3)
	for val := range c {
		fmt.Println(val)
	}
}

func x() {
	fmt.Println("...")
}
