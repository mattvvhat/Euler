package main

import (
	"fmt"
	"testing"
)

func TestCounts(t *testing.T) {

	if v := CountPartitions(0); v != 0 {
		t.Fail()
	}

	if v := CountPartitions(1); v != 1 {
		t.Fail()
	}

	if v := CountPartitions(2); v != 2 {
		t.Fail()
	}

	if v := CountPartitions(2); v != 2 {
		fmt.Println("***", v)
		t.Fail()
	}

	if v := CountPartitions(3); v != 3 {
		fmt.Println("***", v)
		t.Fail()
	}

	if v := CountPartitions(4); v != 5 {
		fmt.Println("***", v)
		t.Fail()
	}

	if v := CountPartitions(5); v != 7 {
		fmt.Println("***", v)
		t.Fail()
	}
}

func TestLookup(t *testing.T) {
	lookup := NewLookup()

	if _, ok := lookup.Get(0, 0); ok {
		t.Fail()
	}

	lookup.Set(3, 0, 10)

	if v, ok := lookup.Get(3, 0); !ok || v != 10 {
		t.Fail()
	}
}

func TestMemoizePartitionCount(t *testing.T) {
	seed := NewLookup()

	if v := MemoizePartitionCount(1, &seed); v != 1 {
		t.Fail()
	}

	if v := MemoizePartitionCount(6, &seed); v != 11 {
		t.Fail()
	}
}

func _() {
	fmt.Println("whatever")
}
