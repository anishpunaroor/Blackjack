package blackjack

import (
	"fmt"

	"github.com/anishpunaroor/Blackjack/deck"
)

type state int8

const (
	stateBet state = iota
	statePlayerTurn
	stateDealerTurn
	stateHandOver
)

type Options struct {
	Decks    int
	Hands    int
	BJPayout float64
}

func New(opts Options) Game {
	g := Game{
		state:    statePlayerTurn,
		dealerAI: dealerAI{},
		balance:  0,
	}
	if opts.Decks == 0 {
		opts.Decks = 3
	}
	if opts.Hands == 0 {
		opts.Hands = 100
	}
	if opts.BJPayout == 0 {
		opts.BJPayout = 1.5
	}
	g.nDecks = opts.Hands
	g.nHands = opts.Decks
	g.BJPayout = opts.BJPayout
	return Game{
		state:    statePlayerTurn,
		dealerAI: dealerAI{},
		balance:  0,
	}
}

type Game struct {

	// Hidden fields
	deck     []deck.Card
	nDecks   int
	state    state
	nHands   int
	player   []deck.Card
	dealer   []deck.Card
	dealerAI AI
	balance  int
	BJPayout float64
}

// Determine the current hand in the game
func (g *Game) CurrentHand() *[]deck.Card {
	switch g.state {
	case statePlayerTurn:
		return &g.player
	case stateDealerTurn:
		return &g.dealer
	default:
		panic("It isn't currently any player's turn.")
	}
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

// Deal out the cards to the player and dealer in the game
func Deal(g *Game) {
	g.player = make([]deck.Card, 0, 5)
	g.dealer = make([]deck.Card, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, g.deck = draw(g.deck)
		g.player = append(g.player, card)
		card, g.deck = draw(g.deck)
		g.dealer = append(g.dealer, card)
	}
	g.state = statePlayerTurn
}

// Play the game
func (g *Game) Play(ai AI) int {
	g.deck = deck.New(deck.NumDeck(3), deck.Shuffle)

	for i := 0; i < 2; i++ {
		Deal(g)

		for g.state == statePlayerTurn {
			hand := make([]deck.Card, len(g.player))
			copy(hand, g.player)
			move := ai.Play(hand, g.dealer[0])
			move(g)
		}
		for g.state == stateDealerTurn {
			hand := make([]deck.Card, len(g.dealer))
			copy(hand, g.dealer)
			move := g.dealerAI.Play(hand, g.dealer[0])
			move(g)
		}
		EndHand(g, ai)
	}
	return g.balance
}

type Move func(*Game)

// Hit in blackjack, adding a card to the current hand
func MoveHit(g *Game) {
	hand := g.CurrentHand()
	var card deck.Card
	card, g.deck = draw(g.deck)
	*hand = append(*hand, card)
	if Score(*hand...) > 21 {
		MoveStand(g)
	}
}

func MoveStand(g *Game) {
	g.state++
}

// End the blackjack game and display the results
func EndHand(g *Game, ai AI) {
	pScore, dScore := Score(g.player...), Score(g.dealer...)

	switch {
	case pScore > 21:
		fmt.Println("You busted")
		g.balance--
	case dScore > 21:
		fmt.Println("Dealer busted")
		g.balance++
	case pScore > dScore:
		fmt.Println("You win!")
		g.balance++
	case dScore > pScore:
		fmt.Println("You lose.")
		g.balance--
	case pScore == dScore:
		fmt.Print("It's a draw.")
	}
	fmt.Println()
	ai.Results([][]deck.Card{g.player}, g.dealer)
	g.player = nil
	g.dealer = nil
}

// Determine the score of the hand, accounting for Ace's special case
func Score(hand ...deck.Card) int {
	minScore := MinScore(hand...)
	if minScore > 11 {
		return minScore
	}
	for _, c := range hand {
		if c.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}

// Returns true if the score of a hand has an ace being counted as 11 points
func Soft(hand ...deck.Card) bool {
	minScore := MinScore(hand...)
	score := Score(hand...)
	return minScore != score
}

// Determine the minimum score of a hand, counting ace as 1 point
func MinScore(hand ...deck.Card) int {
	score := 0
	for _, c := range hand {
		score += min(int(c.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
