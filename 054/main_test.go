package main

import (
	"fmt"
	"testing"
)

func TestCardAsString(t *testing.T) {

	expectations := make(map[Card]string)

	expectations[Card{14, 'H'}] = "14H"
	expectations[Card{9, 'D'}] = "9D"
	expectations[Card{1, 'S'}] = "1S"
	expectations[Card{8, 'C'}] = "8C"

	for card, cardString := range expectations {
		if card.String() != cardString {
			t.Fail()
		}
	}
}

func TestHandString(t *testing.T) {
	v := MakeHand(
		Card{1, 'A'},
		Card{1, 'A'},
		Card{1, 'A'},
		Card{1, 'A'},
		Card{1, 'A'},
	)
	if v.String() != "1A 1A 1A 1A 1A" {
		t.Fail()
	}

	if !v.IsFlush() {
		t.Fail()
	}

	b := MakeHand(
		Card{1, 'H'},
		Card{2, 'H'},
		Card{3, 'H'},
		Card{4, 'H'},
		Card{5, 'S'},
	)

	if b.IsFlush() {
		t.Fail()
	}

	c := MakeHand(
		Card{2, 'H'},
		Card{3, 'H'},
		Card{4, 'H'},
		Card{5, 'S'},
		Card{1, 'H'},
	)

	if c.String() != "1H 2H 3H 4H 5S" {
		t.Fail()
	}

	if !c.IsStraight() {
		t.Fail()
	}
}

func TestLoad(t *testing.T) {
	expectations := map[string]Card{
		"2S":  Card{2, 'S'},
		"12H": Card{12, 'H'},
	}

	for expected, result := range expectations {
		c := LoadCard(expected)
		if c.String() != result.String() {
			t.Fail()
		}
	}
}

func TestPairs(t *testing.T) {
	//expectations := map[string](map[int]int){
	//	"1H 2H 3H 4H 4S": map[int]int{1: 1, 2: 2, 3: 3, 4: 2},
	//}

	//for expected, results := range expectations {
	//	fmt.Println(expected, results)
	//}

}

func TestToRank(t *testing.T) {
	expectations := map[int]rune{
		11: 'Q',
		12: 'K',
		13: 'A',
		3:  '3',
	}

	for expected, results := range expectations {
		if ToRank(expected) != results {
			t.Fail()
		}
	}
}

func TestNothing(t *testing.T) {
	fmt.Println("")
}
