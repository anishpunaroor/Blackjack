package blackjack

import (
	"fmt"

	"github.com/anishpunaroor/Blackjack/deck"
)

type AI interface {
	Bet(shuffle bool) int
	Play(hand []deck.Card, dealer deck.Card) Move
	Results(hand [][]deck.Card, dealer []deck.Card)
}

type dealerAI struct{}

func (ai dealerAI) Bet(shuffle bool) int {
	return 1
}

func (ai dealerAI) Play(hand []deck.Card, dealer deck.Card) Move {
	dScore := Score(hand...)
	if dScore <= 14 || (dScore == 17 && Soft(hand...)) {
		return MoveHit
	}
	return MoveStand
}

func (ai dealerAI) Results(hand [][]deck.Card, dealer []deck.Card) {

}

type humanAI struct{}

func HumanAI() AI {
	return humanAI{}
}

//
func (ai humanAI) Bet(shuffle bool) int {
	if shuffle {
		fmt.Println("The deck was shuffled.")
	}
	fmt.Println("What would you like to bet?")
	var bet int
	fmt.Scanf("%d\n", &bet)
	return bet
}

func (ai humanAI) Play(hand []deck.Card, dealer deck.Card) Move {
	for {
		fmt.Println("\nPlayer: ", hand)
		fmt.Println("Dealer: ", dealer)
		fmt.Println("What will you do? (h)it, (s)tand, (d)ouble")
		var input string
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		case "d":
			return MoveDouble
		default:
			fmt.Println("Enter a valid option.")
		}
	}
}

func (ai humanAI) Results(hands [][]deck.Card, dealer []deck.Card) {
	fmt.Println("--FINAL HANDS--")
	fmt.Println("Player:")
	for _, h := range hands {
		fmt.Println(" ", h)
	}
	fmt.Println("Dealer:", dealer)
}
