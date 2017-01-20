package main

import (
	"fmt"
	"testing"
)

func _(t *testing.T) {
	fmt.Println("...")
}

func TestPhonenumber(t *testing.T) {
	a := Partitions(1, 1)
	if a[0][0] != 1 {
		t.Fail()
	}

	b := Partitions(1, 5)
	if b[0][0] != 5 {
		t.Fail()
	}

	c := Partitions(2, 2)
	if c[0][0] != 1 || c[0][1] != 1 {
		fmt.Println(c)
		t.Fail()
	}

	// d := Partitions(30, 100)
	// d := Partitions(3, 100)
	// fmt.Println(len(d))
	// fmt.Println(d)
}

func TestPartitionCount(t *testing.T) {
	if PartitionCount(1, 1) != 1 {
		t.Fail()
	}

	if PartitionCount(3, 3) != 1 {
		t.Fail()
	}

	if PartitionCount(3, 2) != 1 {
		t.Fail()
	}

	/*
		fmt.Println(PartitionCount(5, 2))
		fmt.Println(PartitionCount(5, 3))
		fmt.Println(PartitionCount(5, 4))
		fmt.Println(PartitionCount(5, 5))
		fmt.Println(PartitionCount(100, 97))
	*/

	total := 0
	limit := 100
	for i := 1; i < limit; i++ {
		total += PartitionCount(limit, i)
	}

	fmt.Println(total)
}
