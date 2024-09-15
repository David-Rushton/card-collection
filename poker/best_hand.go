package poker

import (
	"example.com/cards/cards"
)

// https://en.wikipedia.org/wiki/Texas_hold_%27em#Hand_values
type handValue int

const (
	highCard handValue = iota
	pair
	twoPairs
	threeOfAKind
	straight
	flush
	fullHouse
	fourOfAKind
	straightFlush
	royalFlush
)

type orderedHand []cards.Card

// Given up to 7 cards, returns the best hand available.
func BestHand(hand []cards.Card) {
	// Let's collect some stats.
	favourSuit := mostCommonSuit(hand)

}

func extractFlush(hand []cards.Card) (flush []cards.Card, ok bool) {

}

// Returns the most common suit within the hand.
// In the event of a draw, we tie break on position within the hand.  Where later outranks earlier.
// Draws won't effect how we score a hand.  You need 5 of a suit to build a flush.  As we never
// consider hands of more than 7 cards; a draw will never result in a flush.
func mostCommonSuit(hand []cards.Card) cards.Suit {
	var result cards.Suit
	maxCount := 0
	countBySuit := make(map[cards.Suit]int)

	for _, card := range hand {
		countBySuit[card.Suit]++

		// We won't find a higher count.
		if countBySuit[card.Suit] > (len(hand)/2)+1 {
			return card.Suit
		}

		// Track the most popular suit.
		// This will return the last seen suit in the event of a tie.
		maxCount = max(maxCount, countBySuit[card.Suit])
		if countBySuit[card.Suit] == maxCount {
			result = card.Suit
		}
	}

	return result
}
