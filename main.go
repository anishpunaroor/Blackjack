package main

import (
	"fmt"

	"github.com/anishpunaroor/Blackjack/blackjack"
)

func main() {
	opts := blackjack.Options{
		Decks:    3,
		Hands:    2,
		BJPayout: 1.5,
	}
	game := blackjack.New(opts)
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)

}
