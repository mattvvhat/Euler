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

func TestNothing(t *testing.T) {
	fmt.Println("k")
}
