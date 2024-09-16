package poker

import (
	"slices"

	"github.com/David-Rushton/card-collection/cards"
)

// https://en.wikipedia.org/wiki/Texas_hold_%27em#Hand_values
type HandType int

const (
	HighCard HandType = iota
	Pair
	TwoPairs
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

// Returns the best hand available.
func BestHand(hand cards.Hand) (HandType, cards.Hand) {
	// # Summary
	//
	// There are two phases to this method.
	//
	// In the first phase we collection information and generate a series of indexes over the hand.
	//
	// The second phase uses the collected info and indexes to return the best hand available.

	// TODO: Add support for kickers.

	// ## Phase 1 || Analyse.

	// Create subgroups of cards:
	// - by rank
	// - by suit
	// - consecutive cards (built from by rank, favouring suit _x_)

	sortedHand := hand.Sort()
	countByRank := make(map[cards.Rank]int)
	cardsByRank := make(map[cards.Rank][]cards.Card)
	countBySuit := make(map[cards.Suit]int)
	cardsBySuit := make(map[cards.Suit][]cards.Card)
	sortedRanks := []cards.Rank{}

	for _, card := range sortedHand {
		countBySuit[card.Suit]++
		countByRank[card.Rank]++

		if countByRank[card.Rank] == 1 {
			sortedRanks = append(sortedRanks, card.Rank)
		}

		if _, ok := cardsByRank[card.Rank]; !ok {
			cardsByRank[card.Rank] = []cards.Card{card}
		} else {
			cardsByRank[card.Rank] = append(cardsByRank[card.Rank], card)
		}

		if _, ok := cardsBySuit[card.Suit]; !ok {
			cardsBySuit[card.Suit] = []cards.Card{card}
		} else {
			cardsBySuit[card.Suit] = append(cardsBySuit[card.Suit], card)
		}
	}

	var favourSuit cards.Suit
	var favourSuitCount int
	for k, v := range countBySuit {
		if v > favourSuitCount {
			favourSuit = k
			favourSuitCount = v
		}
	}

	getCard := func(hand []cards.Card, favourSuit cards.Suit) cards.Card {
		for _, card := range hand {
			if card.Suit == favourSuit {
				return card
			}
		}

		return hand[0]
	}

	var pairs []cards.Rank
	var triples []cards.Rank
	var quadruples []cards.Rank
	var lastRank cards.Rank
	for k, v := range countByRank {
		switch v {
		case 4:
			quadruples = append(quadruples, k)
		case 3:
			triples = append(triples, k)
		case 2:
			pairs = append(pairs, k)
		}
	}

	var consecutiveCards []cards.Card
	for _, k := range sortedRanks {
		v := cardsByRank[k]
		if lastRank == 0 {
			lastRank = k
			consecutiveCards = append(consecutiveCards, getCard(v, favourSuit))
			continue
		}

		if k == lastRank+1 {
			consecutiveCards = append(consecutiveCards, getCard(v, favourSuit))
		} else {
			consecutiveCards = []cards.Card{}
		}

		lastRank = k
	}

	// The final five are the highest value
	if len(consecutiveCards) > 5 {
		consecutiveCards = consecutiveCards[len(consecutiveCards)-5:]
	}

	consecutiveSuits := make(map[cards.Suit]int)
	for _, c := range consecutiveCards {
		consecutiveSuits[c.Suit]++
	}

	// ## Phase 2 || Return best available hand

	// royalFlush
	// TODO: Support high aces.

	// straightFlush
	if len(consecutiveCards) == 5 && len(consecutiveSuits) == 1 {
		return StraightFlush, consecutiveCards
	}

	// fourOfAKind
	if len(quadruples) > 0 {
		return FourOfAKind, cardsByRank[quadruples[0]]
	}

	// fullHouse
	if len(triples) > 0 && len(pairs) > 0 {
		return FullHouse, slices.Concat(cardsByRank[triples[0]], cardsByRank[pairs[0]])
	}

	// flush
	if countBySuit[favourSuit] >= 5 {
		return Flush, cardsBySuit[favourSuit]
	}

	// straight
	if len(consecutiveCards) == 5 {
		return Straight, consecutiveCards
	}

	// threeOfAKind
	if len(triples) > 1 {
		return ThreeOfAKind, cardsByRank[triples[0]]
	}

	// twoPairs
	if len(pairs) == 2 {
		return TwoPairs, slices.Concat(cardsByRank[pairs[0]], cardsByRank[pairs[1]])
	}

	// pair
	if len(pairs) == 1 {
		return Pair, cardsByRank[pairs[0]]
	}

	// highCard
	return HighCard, cards.Hand{sortedHand[len(sortedHand)-1]}
}

// func extractFlush(hand []cards.Card) (flush []cards.Card, ok bool) {

// }

// Returns the most common suit within the hand.
// In the event of a draw, we tie break on position within the hand.  Where later outranks earlier.
// Draws won't effect how we score a hand.  You need 5 of a suit to build a flush.  As we never
// consider hands of more than 7 cards; a draw will never result in a flush.
// func mostCommonSuit(hand []cards.Card) cards.Suit {
// 	var result cards.Suit
// 	maxCount := 0
// 	countBySuit := make(map[cards.Suit]int)

// 	for _, card := range hand {
// 		countBySuit[card.Suit]++

// 		// We won't find a higher count.
// 		if countBySuit[card.Suit] > (len(hand)/2)+1 {
// 			return card.Suit
// 		}

// 		// Track the most popular suit.
// 		// This will return the last seen suit in the event of a tie.
// 		maxCount = max(maxCount, countBySuit[card.Suit])
// 		if countBySuit[card.Suit] == maxCount {
// 			result = card.Suit
// 		}
// 	}

// 	return result
// }
