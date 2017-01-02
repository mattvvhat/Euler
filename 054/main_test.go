package main

import (
	"fmt"
	"testing"
)

func TestCardAsString(t *testing.T) {

	expectations := make(map[Card]string)

	expectations[Card{14, 'H'}] = "AH"
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
		"JS":  Card{11, 'S'},
		"TS":  Card{10, 'S'},
		"QD":  Card{12, 'D'},
		"KC":  Card{13, 'C'},
		"AS":  Card{14, 'S'},
	}

	for expected, result := range expectations {
		c := LoadCard(expected)
		if c.String() != result.String() {
			t.Fail()
		}
	}
}

func TestLoadHand(t *testing.T) {
	h := LoadHand("5H 5C 6S 7S KD")
	if h.String() != "5C 5H 6S 7S KD" {
		fmt.Println(h.String())
		t.Fail()
	}

	a := LoadHand("AH 5S 6D TC JS")
	if a.String() != "5S 6D TC JS AH" {
		fmt.Println(a.String())
		t.Fail()
	}
}

func areMapsEqual(lhs, rhs map[int]int) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for k, _ := range lhs {
		if lhs[k] != rhs[k] {
			return false
		}
	}
	return true
}

func TestPairs(t *testing.T) {
	expectations := map[string](map[int]int){
		"AH 2H 3H 4H 4S": map[int]int{14: 1, 2: 1, 3: 1, 4: 2},
		"AH AS AD AC 4S": map[int]int{14: 4, 4: 1},
	}

	for expected, results := range expectations {
		if !areMapsEqual(LoadHand(expected).CardSet(), results) {
			t.Fail()
		}
		// fmt.Println(LoadHand(expected).CardSet(), results)
	}

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

func TestScore(t *testing.T) {
	if LoadHand("AC 6H 7S 9S 9D").CardDist() != "1 1 1 2" {
		t.Fail()
	}

	if LoadHand("AC 6H 7S 8S 9D").CardDist() != "1 1 1 1 1" {
		t.Fail()
	}

	if LoadHand("AC 6H 7S 8S 9D").Score() != 0 {
		t.Fail()
	}
	if LoadHand("5S 5D AS TS JS").Score() != 1 {
		t.Fail()
	}
	if LoadHand("5S 5D AS AC JS").Score() != 2 {
		t.Fail()
	}
	if LoadHand("5S 5D 5S AC JS").Score() != 3 {
		t.Fail()
	}
	if LoadHand("5S 6D 7S 8C 9S").Score() != 4 {
		t.Fail()
	}
	if LoadHand("5S 6S 7S 8S TS").Score() != 5 {
		t.Fail()
	}
	if LoadHand("5S 5D 5H AC AS").Score() != 6 {
		t.Fail()
	}
	if LoadHand("5S 5D 5H 5C AS").Score() != 7 {
		t.Fail()
	}
	if LoadHand("5S 6S 7S 8S 9S").Score() != 8 {
		t.Fail()
	}

	// LoadHand("5C 6H 7S 8S 9D").Score
}

func TestWhoWins(t *testing.T) {
	switch {
	case WhoWins("5S 6S 7S 8S 9S", "5D 6D 7D 8D 2D") != -1:
		t.Fail()
	case WhoWins("5S 6S 7S 8S TS", "5D 6D 7D 8D 9D") != +1:
		t.Fail()
	case WhoWins("5S 6S 7S 8S TS", "5D 6D 7D 8D TD") != 0:
		t.Fail()
	case WhoWins("3C 7S QC QH KS", "4C 5C 5D KC AH") != -1:
		fmt.Println("??")
		t.Fail()
	}
}

func TestNothing(t *testing.T) {
	fmt.Println("")
}
