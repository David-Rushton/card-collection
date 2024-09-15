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
	Ace Rank = iota + 1
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
	Rank Rank
	Suit Suit
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

type Hand []Card
type OrderedHand []Card

// Sorts the hands by rank.
// Aces are considered low, for the purposes of sorting.
// Suit is not considered.  The outcome is stable.
func (h Hand) Sort() Hand {
	if len(h) > 1 {
		return mergeSort(h)
	}

	return h
}

// [Merge sorts] a hand of cards, by rank.
// Suits are not considered.  The outcome is stable.
//
// [Merge sorts]: https://en.wikipedia.org/wiki/Merge_sort
func mergeSort(h Hand) Hand {
	// Nothing to sort, exit early.
	if len(h) == 1 {
		return h
	}

	var left Hand
	var right Hand

	for i, card := range h {
		if i < len(h)/2 {
			left = append(left, card)
		} else {
			right = append(right, card)
		}
	}

	left = left.Sort()
	right = right.Sort()

	return merge(left, right)
}

// Merges left and right.
// Taking the lowest rank from the head of left/right on each iteration.
// The result is a stable sorted merge.
func merge(left, right Hand) Hand {
	var result Hand

	// Iterate until either left or right is depleted.
	// Always take lower of the two and append to result.
	for len(left) > 0 && len(right) > 0 {
		if left[0].Rank <= right[0].Rank {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	// Consume any remaining elements.
	// At most only one of these conditions will be true.
	for len(left) > 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) > 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}
