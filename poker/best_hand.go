package poker

import (
	"log"
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

// TODO: Review and use cards.OrderedHand over cards.Hand, where appropriate.

// Returns the best hand available.
func BestHand_old(hand cards.Hand) (HandType, cards.Hand) {
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

	// Placing the highest values cards at the start simplifies finding them later.
	// slices.Reverse(sortedHand)

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

	slices.Reverse(pairs)
	slices.Reverse(triples)
	slices.Reverse(quadruples)

	// Special case.  Aces are low and high.
	if sortedRanks[len(sortedRanks)-1] == cards.Ace {
		sortedRanks = slices.Concat([]cards.Rank{cards.Ace}, sortedRanks)
	}

	consecutiveCards := []cards.Card{}
	for i, k := range sortedRanks {
		v := cardsByRank[k]

		// Last card is one higher than the previous.
		if lastRank == 0 || k == lastRank+1 || (lastRank == cards.King && k == cards.Ace) {
			consecutiveCards = append(consecutiveCards, getCard(v, favourSuit))
			lastRank = k
			continue
		}

		// We won't find a straight at this point.
		if len(sortedRanks)-i < 5 {
			break
		}

		consecutiveCards = []cards.Card{getCard(v, favourSuit)}
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

	// TODO: Tidy this away.
	// HACK: Relies on order of from.
	// TODO : Does not support high aces 🫠.
	takeKickers := func(from cards.Hand, to cards.Hand) cards.Hand {
		for i := len(from) - 1; i >= 0; i-- {
			if len(to) >= 5 {
				break
			}

			if !slices.Contains(to, from[i]) {
				to = append(to, from[i])
			}
		}

		return to
	}

	// ## Phase 2 || Return best available hand

	// royalFlush
	if len(consecutiveCards) == 5 &&
		len(consecutiveSuits) == 1 &&
		consecutiveCards[len(consecutiveCards)-1].Rank == cards.Ace {
		return RoyalFlush, consecutiveCards
	}
	// straightFlush
	if len(consecutiveCards) == 5 && len(consecutiveSuits) == 1 {
		return StraightFlush, consecutiveCards
	}

	// fourOfAKind
	if len(quadruples) > 0 {
		return FourOfAKind, takeKickers(sortedHand, cardsByRank[quadruples[0]])
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
	if len(triples) > 0 {
		return ThreeOfAKind, takeKickers(sortedHand, cardsByRank[triples[0]])
	}

	// twoPairs
	if len(pairs) == 2 {
		return TwoPairs, takeKickers(sortedHand, slices.Concat(cardsByRank[pairs[0]], cardsByRank[pairs[1]]))
	}

	// pair
	if len(pairs) == 1 {
		return Pair, takeKickers(sortedHand, cardsByRank[pairs[0]])
	}

	// highCard
	return HighCard, takeKickers(sortedHand, cards.Hand{sortedHand[len(sortedHand)-1]})
}

// Scores a hand.
// Higher is better.
func ScoreHand(handType HandType, hand cards.Hand) int64 {
	// Each hand is assigned an integer score.
	//
	// The first two digits are hand type, where better hands are assigned a higher multiple.
	// The next two digits are the value of the first cards rank.  Ace == 13, King == 12, on so on.
	// Until we reach 2, with is worth 02.  We continue with the rest of the cards.
	//
	// Example Hand:
	// Type: Straight
	// Cards: 7 Diamonds, 8 Clubs, 9 Spades, 10 Spades and Jack of Clubs
	//
	// Straight is the 5th best hand.  This scores 04, as we index zero.  Followed by 07 for the
	// 7 of diamonds, and 08 for the 8 of clubs.  Etc.
	//
	// Result: 04 07 08 09 10 11
	// Formatted: 40,708,091,011
	//
	// In this case a straight starting with an 8 and ending with a Queen would command a better
	// score.

	if len(hand) > 5 {
		log.Fatalf("Cannot score hand.  The hand contains too many cards.  Expected up to 5.  But found %v.", len(hand))
	}

	score := int64(0)
	multiplier := int64(1)
	for i := len(hand) - 1; i >= 0; i-- {
		score += int64(rankScore(hand[i].Rank)) * multiplier
		multiplier *= 100
	}

	score += int64(handType) * multiplier

	return score
}

// Converts a rank to a score.
// Aces are always high.
func rankScore(r cards.Rank) int64 {
	switch r {
	case cards.Ace:
		return 14
	case cards.Two:
		return 2
	case cards.Three:
		return 3
	case cards.Four:
		return 4
	case cards.Five:
		return 5
	case cards.Six:
		return 6
	case cards.Seven:
		return 7
	case cards.Eight:
		return 8
	case cards.Nine:
		return 9
	case cards.Ten:
		return 10
	case cards.Jack:
		return 11
	case cards.Queen:
		return 12
	case cards.King:
		return 13
	default:
		log.Fatalf("Rank not supported: %v.  This error is fatal and cannot be recovered.  A code change is required to fix.", r)
	}

	// We will never reach this line.
	// But the compiler doesn't know that (as of Go 1.23.1) 😑.
	return 0
}
