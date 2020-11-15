package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/peter554/gophercises/10-blackjack/deck"
)

func main() {
	rand.Seed(time.Now().Unix())

	cards := deck.New()
	cards.Shuffle()

	playerCards := deck.Cards{}
	dealerCards := deck.Cards{}

	for i := 0; i < 2; i++ {
		playerCards = append(playerCards, cards.Pop())
		dealerCards = append(dealerCards, cards.Pop())
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("\nDealer has\n")
		fmt.Println(dealerCards[0])
		fmt.Println("Your cards are")
		for _, card := range playerCards {
			fmt.Println(card)
		}
		score := getScore(playerCards)
		fmt.Printf("Score: %d\n", score)

		if score > 21 {
			fmt.Printf("Bust! You lose.\n\n")
			return
		}

		fmt.Printf("Hit or stick [h/s]?\n\n")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		if choice != "h" && choice != "s" {
			continue
		}

		if choice == "s" {
			break
		}

		playerCards = append(playerCards, cards.Pop())
	}

	for {
		if getScore(dealerCards) <= 16 {
			dealerCards = append(dealerCards, cards.Pop())
		} else {
			break
		}
	}

	dealerScore := getScore(dealerCards)
	playerScore := getScore(playerCards)
	fmt.Printf("\nDealer scores %d\n", dealerScore)
	fmt.Printf("You score %d\n", playerScore)
	if dealerScore > 21 || playerScore > dealerScore {
		fmt.Println("You win!")
	} else if playerScore < dealerScore {
		fmt.Println("You lose.")
	} else {
		fmt.Println("Draw.")
	}

}

func getScore(cards deck.Cards) int {
	rankMap := map[deck.Rank]int{
		deck.Ace:   1,
		deck.Two:   2,
		deck.Three: 3,
		deck.Four:  4,
		deck.Five:  5,
		deck.Six:   6,
		deck.Seven: 7,
		deck.Eight: 8,
		deck.Nine:  9,
		deck.Ten:   10,
		deck.Jack:  10,
		deck.Queen: 10,
		deck.King:  10,
	}
	score := 0
	for _, card := range cards {
		cardScore, _ := rankMap[card.Rank]
		score += cardScore
	}
	for _, card := range cards {
		if card.Rank == deck.Ace {
			if score+10 <= 21 {
				score += 10
			} else {
				break
			}
		}
	}
	return score
}
