package poker

import (
	"slices"

	"github.com/David-Rushton/card-collection/deck"
)

func BestHand(hand deck.Hand) (HandType, deck.Hand) {
	return BestHand2(deck.Hand{}, hand)
}

// Returns the best hand that can be constructed from the passed hand.
func BestHand2(board, players deck.Hand) (HandType, deck.Hand) {
	// ## Summary
	//
	// There are two phases to this method:
	//
	//  1. Analysis       || Collects data on the provided cards
	//  2. Hand Selection || Identifies the best hand available
	//
	// This method is fairly long.  It has been written in the procedural style.  This is on
	// purpose.  There are no obvious ways to break the code into sections, without passing a lot
	// of values between the sub methods.

	// ## Section 1 || Analysis
	// ------------------------

	sortedHand := board.Append(players, len(players)).Sort()
	countByRank := make(map[deck.Rank]int)
	cardsByRank := make(map[deck.Rank]deck.Hand)
	countBySuit := make(map[deck.Suit]int)
	cardsBySuit := make(map[deck.Suit]deck.Hand)
	sortedRanks := []deck.Rank{}

	// Groups and counts the cards by suit and rank.
	// This will help us find consecutive cards, pairs, trebles and quadruples later on.
	for _, card := range sortedHand {
		countBySuit[card.Suit]++
		countByRank[card.Rank]++

		if countByRank[card.Rank] == 1 {
			sortedRanks = append(sortedRanks, card.Rank)
		}

		if _, ok := cardsByRank[card.Rank]; !ok {
			cardsByRank[card.Rank] = []deck.Card{card}
		} else {
			cardsByRank[card.Rank] = append(cardsByRank[card.Rank], card)
		}

		if _, ok := cardsBySuit[card.Suit]; !ok {
			cardsBySuit[card.Suit] = []deck.Card{card}
		} else {
			cardsBySuit[card.Suit] = append(cardsBySuit[card.Suit], card)
		}
	}

	// Finds the most common suit.
	// This will help us identify flushes later on.
	var favourSuit deck.Suit
	var favourSuitCount int
	for k, v := range countBySuit {
		if v > favourSuitCount {
			favourSuit = k
			favourSuitCount = v
		}
	}

	// Find pairs, trebles and quadruples.
	// Prepend values, to ensure higher value ranks are at the start of each slice.
	// This simplifies taking later on.
	var pairs []deck.Rank
	var trebles []deck.Rank
	var quadruples []deck.Rank
	for k, v := range countByRank {
		switch v {
		case 4:
			quadruples = slices.Insert(quadruples, 0, k)
		case 3:
			trebles = slices.Insert(trebles, 0, k)
		case 2:
			pairs = slices.Insert(pairs, 0, k)
		}
	}

	// Find consecutive cards.
	// 5 or more is a straight.
	// Use the favoured suite, from earlier, to ensure straight flushes are found.

	// Aces are high and low.
	if sortedRanks[len(sortedRanks)-1] == deck.Ace {
		sortedRanks = slices.Concat([]deck.Rank{deck.Ace}, sortedRanks)
	}

	var lastRank deck.Rank
	consecutiveSuits := make(map[deck.Suit]int)
	consecutiveCards := deck.Hand{}
	for i, k := range sortedRanks {
		v := cardsByRank[k]

		// Last card is one higher than the previous.
		if lastRank == 0 || k == lastRank+1 || (lastRank == deck.King && k == deck.Ace) {
			nextCard := takeFirstFavouringSuit(v, favourSuit)
			consecutiveCards = append(consecutiveCards, nextCard)
			consecutiveSuits[nextCard.Suit]++
			lastRank = k
			continue
		}

		// Check if it is possible to find a better straight, or indeed one at all.
		if len(sortedRanks)-i < 5 {
			break
		}

		// Not consecutive.  Reset.
		consecutiveCards = deck.Hand{takeFirstFavouringSuit(v, favourSuit)}
		consecutiveSuits = map[deck.Suit]int{
			consecutiveCards[0].Suit: 1,
		}
		lastRank = k
	}

	// The final five are the highest value.
	if len(consecutiveCards) > 5 {
		consecutiveCards = consecutiveCards[len(consecutiveCards)-5:]
	}

	// Place the highest value cards at the start, making them easier to access.
	kickers := sortedHand
	slices.Reverse(kickers)

	// ## Section 2 || Identify best hand
	// ----------------------------------

	// royalFlush
	if len(consecutiveCards) == 5 &&
		len(consecutiveSuits) == 1 &&
		consecutiveCards[len(consecutiveCards)-1].Rank == deck.Ace {
		return RoyalFlush, consecutiveCards
	}
	// straightFlush
	if len(consecutiveCards) == 5 && len(consecutiveSuits) == 1 {
		return StraightFlush, consecutiveCards
	}

	// fourOfAKind
	if len(quadruples) > 0 {
		return FourOfAKind, addKickers(cardsByRank[quadruples[0]], kickers)
	}

	// fullHouse
	if len(trebles) > 0 && len(pairs) > 0 {
		return FullHouse, slices.Concat(cardsByRank[trebles[0]], cardsByRank[pairs[0]])
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
	if len(trebles) > 0 {
		return ThreeOfAKind, addKickers(cardsByRank[trebles[0]], kickers)
	}

	// twoPairs
	if len(pairs) == 2 {
		return TwoPairs, addKickers(slices.Concat(cardsByRank[pairs[0]], cardsByRank[pairs[1]]), kickers)
	}

	// pair
	if len(pairs) == 1 {
		return Pair, addKickers(cardsByRank[pairs[0]], kickers)
	}

	// highCard
	return HighCard, kickers[0:5]
}

// Takes the first card from a hand, that matches the favoured suit.
// If no match we return the first card.
func takeFirstFavouringSuit(hand deck.Hand, suit deck.Suit) deck.Card {
	for _, card := range hand {
		// Take the first card that matces the favoured suit.
		if card.Suit == suit {
			return card
		}
	}

	// There were no matches, return the first card.
	return hand[0]
}

// Pads the given hand with kickers.
func addKickers(hand deck.Hand, kickers deck.Hand) deck.Hand {
	required := 5 - len(hand)
	return hand.AppendWhen(kickers, required, func(c deck.Card) bool {
		return !slices.Contains(hand, c)
	})
}
