// A deck of 52 playing cards, in the classic [french-suited style].
//
// [french-suited style]: https://en.wikipedia.org/wiki/French-suited_playing_cards
package deck

import (
	"errors"
	"math/rand"
)

var (
	deck []Card
)

// Shuffles the deck.
// Each card is moved to a random location.
func Shuffle() {
	// Reset the deck.
	deck = make([]Card, 52)
	for i := 0; i < 52; i++ {
		deck[i] = Card{Rank(i % 13), Suit(i % 4)}
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

// Takes the top n cards from the deck.
// If there are not enough cards returns ErrNotEnoughCards.
func Take(n int) (Hand, error) {
	// Validate.
	if n < 0 {
		return nil, errors.New("cannot take less than 0 card")
	}

	if len(deck) < n {
		return nil, ErrNotEnoughCards{Requested: n, Remaining: len(deck)}
	}

	// Take.
	result := deck[0:n]
	deck = deck[n:]

	return result, nil
}

// Returns the number cards in the deck.
func Remaining() int {
	return len(deck)
}
