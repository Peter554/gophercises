package main

import (
	"fmt"

	"github.com/peter554/gophercises/09-deck-of-cards/deck"
)

func main() {
	cards := deck.New()
	fmt.Println(cards)

	cards.Sort()
	fmt.Println(cards)

	cards.Shuffle()
	fmt.Println(cards)

	cards.Shuffle()
	fmt.Println(cards)
}
