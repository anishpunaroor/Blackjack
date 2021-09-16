package main

import (
	"fmt"

	"github.com/anishpunaroor/Blackjack/blackjack"
	"github.com/anishpunaroor/Blackjack/deck"
)

func main() {
	opts := blackjack.Options{
		Decks:    4,
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

func (ai *baseAI) Play(hand []deck.Card, dealer deck.Card) blackjack.Move {
	score := blackjack.Score(hand...)
	// The AI will double down when the score is 10 or 11, or the hand is soft
	if len(hand) == 2 {
		if (score == 10 || score == 11) && !blackjack.Soft(hand...) {
			return blackjack.MoveDouble
		}
	}
	dScore := blackjack.Score(dealer)
	if dScore >= 5 && dScore <= 6 {
		return blackjack.MoveStand
	}
	if score < 13 {
		return blackjack.MoveHit
	}
	return blackjack.MoveStand
}
