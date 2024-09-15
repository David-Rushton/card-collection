package cards

import (
	"errors"
	"fmt"
	"math/rand"
)

type Suit int

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

type Rank int

const (
	Ace Rank = iota
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

type Card struct {
	Suit Suit
	Rank Rank
}

var (
	deck = []Card{}
)

func (c Card) String() string {
	var suit string
	// ♣️ ♦️ ♥️ ♠️
	switch c.Suit {
	case Clubs:
		suit = "Clubs"
	case Diamonds:
		suit = "Diamonds"
	case Hearts:
		suit = "Hearts"
	case Spades:
		suit = "Spades"
	}

	var rank string
	switch c.Rank {
	case Ace:
		rank = "Ace"
	case Jack:
		rank = "Jack"
	case Queen:
		rank = "Queen"
	case King:
		rank = "King"
	default:
		rank = fmt.Sprintf("%d", int(c.Rank))
	}
	return fmt.Sprintf("%v of %v", rank, suit)
}

func Shuffle() {
	// Reset the deck.
	deck = make([]Card, 52)
	for i := 0; i < 52; i++ {
		deck[i] = Card{Suit(i % 4), Rank(i % 13)}
	}

	// Modern Fisher-Yates shuffle.
	// https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
	for i := 51; i > 0; i-- {
		swapAt := rand.Intn(i)
		swap := deck[swapAt]
		deck[swapAt] = deck[i]
		deck[i] = swap
	}
}

type NotEnoughCardsError struct {
	Requested int
	Remaining int
}

func (e NotEnoughCardsError) Error() string {
	if e.Remaining == 1 {
		return fmt.Sprintf("Cannot take %d cards.  There is only %d remaining in the deck.  Request fewer or shuffle the deck", e.Requested, e.Remaining)
	}

	return fmt.Sprintf("Cannot take %d cards.  There are only %d remaining in the deck.  Request fewer or shuffle the deck", e.Requested, e.Remaining)
}

func Take(n int) ([]Card, error) {
	//Validate.
	if n < 1 {
		return nil, errors.New("Take cannot request less than 1 card")
	}

	if n > len(deck) {
		return nil, NotEnoughCardsError{Requested: n, Remaining: len(deck)}
	}

	// Take from deck and resize.
	result := deck[0:n]
	deck = deck[n:]

	return result, nil
}

func Remaining() int {
	return len(deck)
}
