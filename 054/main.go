package main

import (
	"fmt"
	"sort"
	"strings"
)

type Card struct {
	Rank int
	Suit rune
}

type CardList []Card

type Hand struct {
	Cards CardList
}

func (h CardList) Len() int {
	return len(h)
}

func (h CardList) Less(i, j int) bool {
	return h[i].String() < h[j].String()
}

func (h CardList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func MakeHand(a, b, c, d, e Card) Hand {
	hand := Hand{[]Card{a, b, c, d, e}}
	sort.Sort(hand.Cards)
	return hand
}

func (h Hand) String() string {
	values := make([]string, 5, 5)
	for i, v := range h.Cards {
		values[i] = v.String()
	}
	return strings.Join(values, " ")
}

func (c Card) String() string {
	return fmt.Sprintf("%d", c.Rank) + string(c.Suit)
}

func (h *Hand) IsStraight() bool {
	prev := h.Cards[0].Rank - 1

	for _, v := range h.Cards {
		if prev+1 != v.Rank {
			return false
		}
		prev = v.Rank
	}

	return true
}

func (h *Hand) IsFlush() bool {
	suit := h.Cards[0].Suit

	for _, v := range h.Cards {
		if suit != v.Suit {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(Card{14, 'H'}.String())
}
