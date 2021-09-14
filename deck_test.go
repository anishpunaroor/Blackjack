package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Nine, Suit: Diamond})
	fmt.Println(Card{Rank: Jack, Suit: Club})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Nine of Diamonds
	// Jack of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	// 52 cards in a standard deck of cards (excluding Joker)
	if len(cards) != 52 {
		t.Error("Wrong number of cards in the deck.")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Did not sort correctly. Received:  ", cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(5))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 5 {
		t.Error("Expected 5 Jokers, got:  ", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, card := range cards {
		if card.Rank == Two || card.Rank == Three {
			t.Error("Expected twos and threes to be filtered out.")
		}
	}
}
