package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	_ Rank = iota
	Ace
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

const (
	minRank = Ace
	maxRank = King
)

// Cards belong to a suit and a rank
type Card struct {
	Suit
	Rank
}

// String representation of a card
func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

type NewOpts struct {
	Shuffle bool
}

// Initialize a deck of cards, and shuffle/sort them based on options
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

// Sort cards by absolute rank
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

// Absolute Rank comparator
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

// Determine the absolute rank of a card (like an array index)
func absRank(c Card) int {
	return int(c.Suit) * int(maxRank) * int(c.Rank)
}

// Shuffle the deck of cards using randomized permutations
func Shuffle(cards []Card) []Card {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	shuf := make([]Card, len(cards))
	perm := r.Perm(len(cards))
	for i, randInd := range perm {
		shuf[i] = cards[randInd]
	}
	return shuf
}

// Add a varying amount of Jokers to the deck
func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Rank: Rank(i),
				Suit: Joker,
			})
		}
		return cards
	}
}

// Filter the deck and remove any cards as specified
func Filter(f func(card Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var fil []Card
		for _, card := range cards {
			if !f(card) {
				fil = append(fil, card)
			}
		}
		return fil
	}
}

// Add multiple copies of a standard deck
func NumDeck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var deck []Card
		for i := 0; i < n; i++ {
			deck = append(deck, cards...)
		}
		return deck
	}
}
