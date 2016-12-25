package main

import (
	"fmt"
)

type Card struct {
	Value int
	Suit  rune
}

func (c Card) String() string {
	return fmt.Sprintf("%d", c.Value) + string(c.Suit)
}

func main() {
	fmt.Println(Card{14, 'H'}.String())
}
