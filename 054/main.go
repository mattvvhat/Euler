package main

import (
	"fmt"
	"sort"
	"strconv"
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
	if h[i].Rank == h[j].Rank {
		return h[i].Suit < h[j].Suit
	} else {
		return h[i].Rank < h[j].Rank
	}
}

func (h CardList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func LoadCard(v string) Card {
	//if len(v) != 2 {
	//	return Card{0, 'G'}
	//}

	var val int
	rank := v[0 : len(v)-1]

	switch rank {
	case "T":
		val = 10
	case "J":
		val = 11
	case "Q":
		val = 12
	case "K":
		val = 13
	case "A":
		val = 14
	default:
		val, _ = strconv.Atoi(rank)
	}
	suit := rune(v[len(v)-1])

	return Card{val, suit}
}

func LoadHand(v string) Hand {
	cards := strings.Split(v, " ")
	cardList := make([]Card, 5, 5)
	var hand Hand

	for i, c := range cards {
		cardList[i] = LoadCard(c)
	}

	hand = Hand{cardList}
	sort.Sort(hand.Cards)

	return hand
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
	var rank string

	switch c.Rank {
	case 10:
		rank = "T"
	case 11:
		rank = "J"
	case 12:
		rank = "Q"
	case 13:
		rank = "K"
	case 14:
		rank = "A"
	default:
		rank = fmt.Sprintf("%d", c.Rank)
	}

	return rank + string(c.Suit)
}

func (h *Hand) CardSets() map[int]int {
	set := make(map[int]int)
	return set
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

func (h Hand) CardSet() map[int]int {
	cardSet := make(map[int]int)

	for _, c := range h.Cards {
		cardSet[c.Rank] += 1
	}

	return cardSet
}

func (h Hand) CardDist() string {
	values := make([]string, 0, 5)

	for _, c := range h.CardSet() {
		values = append(values, fmt.Sprintf("%d", c))
	}

	sort.Strings(values)

	return strings.Join(values, " ")
}

func (h Hand) Score() int {
	// cardSet := h.CardSet()

	switch h.CardDist() {
	case "1 1 1 2":
		return 1
	case "1 2 2":
		return 2
	case "1 1 3":
		return 3
	case "2 3":
		return 6
	case "1 4":
		return 7
	}

	if h.IsStraight() {
		if h.IsFlush() {
			return 8
		} else {
			return 4
		}
	}

	if h.IsFlush() {
		return 5
	}

	return 0
}

func ToRank(val int) rune {
	s := map[int]rune{
		10: 'J',
		11: 'Q',
		12: 'K',
		13: 'A',
	}

	r, present := s[val]

	if present {
		return r
	} else {
		return rune(fmt.Sprintf("%d", val)[0])
	}
}

func WhoWins(lhs, rhs string) int {
	player1 := LoadHand(lhs)
	player2 := LoadHand(rhs)

	if player1.Score() == player2.Score() {
		for i, _ := range player1.Cards {
			if player1.Cards[5-i-1].Rank < player1.Cards[5-i-1].Rank {
				return +1
			} else if player1.Cards[5-i-1].Rank > player1.Cards[5-i-1].Rank {
				return -1
			}
		}
		return 0
	}

	if player1.Score() < player2.Score() {
		return 1
	} else {
		return -1
	}
}

func main() {
	fmt.Println(Card{14, 'H'}.String())
}
