package main

import (
	"fmt"
	"strings"

	"github.com/anishpunaroor/Blackjack/deck"
)

func main() {
	var gs GameState
	gs = Shuffle(gs)

	for i := 0; i < 10; i++ {
		gs = Deal(gs)

		var input string
		for gs.State == StatePlayerTurn {
			fmt.Println("\nPlayer: ", gs.Player)
			fmt.Println("Dealer: ", gs.Dealer.DealerString())
			fmt.Println("What will you do? (h)it, (s)tand")
			fmt.Scanf("%s\n", &input)
			switch input {
			case "h":
				gs = Hit(gs)
			case "s":
				gs = Stand(gs)
			default:
				fmt.Println("Enter a valid option.")
			}
		}

		// If dealer's score is less than 16, we hit. If dealer has a soft 17, then we hit.
		for gs.State == StateDealerTurn {
			if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
				gs = Hit(gs)
			} else {
				gs = Stand(gs)
			}
		}
		gs = EndHand(gs)
	}

}

func Shuffle(gs GameState) GameState {
	ret := clone(gs)
	ret.Deck = deck.New(deck.NumDeck(3), deck.Shuffle)
	return ret
}

func Stand(gs GameState) GameState {
	ret := clone(gs)
	ret.State++
	return ret
}

func Hit(gs GameState) GameState {
	ret := clone(gs)
	hand := ret.CurrentPlayer()
	var card deck.Card
	card, ret.Deck = draw(ret.Deck)
	*hand = append(*hand, card)
	if hand.Score() > 21 {
		return Stand(ret)
	}
	return ret
}

// Print out the suit and rank of all cards in the hand
func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

// Display the dealer first card in the hand
func (h Hand) DealerString() string {
	return h[0].String() + ", HIDDEN"
}

func clone(gs GameState) GameState {
	ret := GameState{
		Deck:   make([]deck.Card, len(gs.Deck)),
		State:  gs.State,
		Player: make(Hand, len(gs.Player)),
		Dealer: make(Hand, len(gs.Dealer)),
	}
	copy(ret.Deck, gs.Deck)
	copy(ret.Player, gs.Player)
	copy(ret.Dealer, gs.Dealer)
	return ret
}
