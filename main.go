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

type baseAI struct {
	score int
	seen  int
	decks int
}

func (ai *baseAI) Bet(shuffle bool) int {
	if shuffle {
		ai.score = 0
		ai.seen = 0
	}
	trueScore := ai.score / ((ai.decks*52 - ai.seen) / 52)
	switch {
	case trueScore >= 14:
		return 10000
	case trueScore >= 8:
		return 1000
	default:
		return 100
	}
}
