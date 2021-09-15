package main

import (
	"fmt"

	"github.com/anishpunaroor/Blackjack/blackjack"
)

func main() {
	game := blackjack.New(blackjack.Options{})
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)

}
