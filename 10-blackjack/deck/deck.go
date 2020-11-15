package deck

import (
	"fmt"
	"math/rand"
	"sort"
)

type Suit uint8

type Rank uint8

type Card struct {
	Suit Suit
	Rank Rank
}

type Cards []Card

const (
	Club Suit = iota
	Heart
	Spade
	Diamond
)

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

func New() Cards {
	o := Cards{}

	for _, suit := range []Suit{
		Spade,
		Diamond,
		Heart,
		Club,
	} {
		for _, flavor := range []Rank{
			Two,
			Four,
			Jack,
			Ace,
			Ten,
			Five,
			Six,
			Three,
			Seven,
			King,
			Eight,
			Nine,
			Queen,
		} {
			o = append(o, Card{Suit: suit, Rank: flavor})
		}
	}

	return o
}

func absRank(card Card) Rank {
	return Rank(card.Suit)*King + card.Rank
}

func (o Cards) Len() int {
	return len(o)
}

func (o Cards) Less(i, j int) bool {
	return absRank(o[i]) < absRank(o[j])

}

func (o Cards) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func (o Cards) Sort() {
	sort.Sort(o)
}

func (o Cards) Shuffle() {
	rand.Shuffle(len(o), o.Swap)
}

func (o *Cards) Pop() Card {
	t := (*o)[0]
	*o = (*o)[1:]
	return t
}

func (o Card) String() string {
	suitMap := map[Suit]string{
		Club:    "Club",
		Heart:   "Heart",
		Spade:   "Spade",
		Diamond: "Diamond",
	}
	rankMap := map[Rank]string{
		Ace:   "Ace",
		Two:   "Two",
		Three: "Three",
		Four:  "Four",
		Five:  "Five",
		Six:   "Six",
		Seven: "Seven",
		Eight: "Eight",
		Nine:  "Nine",
		Ten:   "Ten",
		Jack:  "Jack",
		Queen: "Queen",
		King:  "King",
	}
	suit, _ := suitMap[o.Suit]
	rank, _ := rankMap[o.Rank]
	return fmt.Sprintf("%s of %ss", rank, suit)
}
