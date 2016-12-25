package main

import (
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
