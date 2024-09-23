package poker_test

import (
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/David-Rushton/card-collection/deck"
	"github.com/David-Rushton/card-collection/poker"
)

func Test_BestHand_ReturnsHighCard(t *testing.T) {
	testCases := []struct {
		hand         deck.Hand
		expectedHand deck.Hand
		expectedType poker.HandType
	}{
		{
			hand:         parseHand("3s 5d 7c 9s Th"),
			expectedHand: parseHand("Th 9s 7c 5d 3s"),
			expectedType: poker.HighCard,
		},
	}

	for _, testCase := range testCases {
		actualType, actualHand := poker.BestHand(testCase.hand)

		if actualType != testCase.expectedType {
			t.Errorf("BestHand did not return expected type.  Expected: %v.  Actual: %v.", testCase.expectedType, actualType)
		}

		if !slices.Equal(testCase.expectedHand, actualHand) {
			t.Errorf("BestHand did not return expected cards.  Expected: %v.  Actual: %v.", testCase.expectedHand, actualHand)
		}
	}
}

func Test_BestHand_ReturnsPair(t *testing.T) {
	testCases := []struct {
		hand         deck.Hand
		expectedHand deck.Hand
		expectedType poker.HandType
	}{
		{
			hand:         parseHand("3h 3s 7d qc kc"),
			expectedHand: parseHand("3h 3s kc qc 7d"),
			expectedType: poker.Pair,
		},
	}

	for _, testCase := range testCases {
		actualType, actualHand := poker.BestHand(testCase.hand)

		if actualType != testCase.expectedType {
			t.Errorf("BestHand did not return expected type.  Expected: %v.  Actual: %v.", testCase.expectedType, actualType)
		}

		if !slices.Equal(testCase.expectedHand, actualHand) {
			t.Errorf("BestHand did not return expected cards.  Expected: %v.  Actual: %v.", testCase.expectedHand, actualHand)
		}
	}
}

func Test_BestHand_ReturnsTwoPairs(t *testing.T) {
	testCases := []struct {
		hand         deck.Hand
		expectedHand deck.Hand
		expectedType poker.HandType
	}{
		{
			hand:         parseHand("2h 2d 7c 3c 3s"),
			expectedHand: parseHand("3c 3s 2h 2d 7c"),
			expectedType: poker.TwoPairs,
		},
	}

	for _, testCase := range testCases {
		actualType, actualHand := poker.BestHand(testCase.hand)

		if actualType != testCase.expectedType {
			t.Errorf("BestHand did not return expected type.  Expected: %v.  Actual: %v.", testCase.expectedType, actualType)
		}

		if !slices.Equal(testCase.expectedHand, actualHand) {
			t.Errorf("BestHand did not return expected cards.  Expected: %v.  Actual: %v.", testCase.expectedHand, actualHand)
		}
	}
}

func Test_BestHand_ReturnsThreeOfAKind(t *testing.T) {
	testCases := []struct {
		hand         deck.Hand
		expectedHand deck.Hand
		expectedType poker.HandType
	}{
		{
			hand:         parseHand("2d 2h 2s qs 9c"),
			expectedHand: parseHand("2d 2h 2s qs 9c"),
			expectedType: poker.ThreeOfAKind,
		},
	}

	for _, testCase := range testCases {
		actualType, actualHand := poker.BestHand(testCase.hand)

		if actualType != testCase.expectedType {
			t.Errorf("BestHand did not return expected type.  Expected: %v.  Actual: %v.", testCase.expectedType, actualType)
		}

		if !slices.Equal(testCase.expectedHand, actualHand) {
			t.Errorf("BestHand did not return expected cards.  Expected: %v.  Actual: %v.", testCase.expectedHand, actualHand)
		}
	}
}

func Test_BestHand_ReturnsStraight(t *testing.T) {
	testCases := []struct {
		hand         deck.Hand
		expectedHand deck.Hand
		expectedType poker.HandType
	}{
		{
			hand:         parseHand("6h 7s 8h 9c tc"),
			expectedHand: parseHand("6h 7s 8h 9c tc"),
			expectedType: poker.Straight,
		},
		// Aces low.
		{
			hand:         parseHand("5s 4h 2h 3h Ad"),
			expectedHand: parseHand("Ad 2h 3h 4h 5s"),
			expectedType: poker.Straight,
		},
		// Aces high.
		{
			hand:         parseHand("Ac Tc Qs Jc Kd"),
			expectedHand: parseHand("Tc Jc Qs Kd Ac"),
			expectedType: poker.Straight,
		},
		// Straight at start of available cards.
		{
			hand:         parseHand("5h 3h 2s Ac 4s 8c 9d"),
			expectedHand: parseHand("Ac 2s 3h 4s 5h"),
			expectedType: poker.Straight,
		},
		// Multiple possible straights
		{
			hand:         parseHand("2s 3s 4d 5c 6s 7d 8d 9h"),
			expectedHand: parseHand("5c 6s 7d 8d 9h"),
			expectedType: poker.Straight,
		},
	}

	for _, testCase := range testCases {
		actualType, actualHand := poker.BestHand(testCase.hand)

		if actualType != testCase.expectedType {
			t.Errorf("BestHand did not return expected type.  Expected: %v.  Actual: %v.", testCase.expectedType, actualType)
		}

		if !slices.Equal(testCase.expectedHand, actualHand) {
			t.Errorf("BestHand did not return expected cards.  Expected: %v.  Actual: %v.", testCase.expectedHand, actualHand)
		}
	}
}

func Test_BestHand_ReturnsFlush(t *testing.T) {
	testCases := []struct {
		hand         deck.Hand
		expectedHand deck.Hand
		expectedType poker.HandType
	}{
		{
			hand:         parseHand("3c 7c 9c tc kc"),
			expectedHand: parseHand("3c 7c 9c tc kc"),
			expectedType: poker.Flush,
		},
	}

	for _, testCase := range testCases {
		actualType, actualHand := poker.BestHand(testCase.hand)

		if actualType != testCase.expectedType {
			t.Errorf("BestHand did not return expected type.  Expected: %v.  Actual: %v.", testCase.expectedType, actualType)
		}

		if !slices.Equal(testCase.expectedHand, actualHand) {
			t.Errorf("BestHand did not return expected cards.  Expected: %v.  Actual: %v.", testCase.expectedHand, actualHand)
		}
	}
}

func Test_BestHand_ReturnsFullHouse(t *testing.T) {
	testCases := []struct {
		hand         deck.Hand
		expectedHand deck.Hand
		expectedType poker.HandType
	}{
		{
			hand:         parseHand("ad ah as Ts Td"),
			expectedHand: parseHand("ad ah as Ts Td"),
			expectedType: poker.FullHouse,
		},
	}

	for _, testCase := range testCases {
		actualType, actualHand := poker.BestHand(testCase.hand)

		if actualType != testCase.expectedType {
			t.Errorf("BestHand did not return expected type.  Expected: %v.  Actual: %v.", testCase.expectedType, actualType)
		}

		if !slices.Equal(testCase.expectedHand, actualHand) {
			t.Errorf("BestHand did not return expected cards.  Expected: %v.  Actual: %v.", testCase.expectedHand, actualHand)
		}
	}
}

func Test_BestHand_ReturnsFourOfAKind(t *testing.T) {
	testCases := []struct {
		hand         deck.Hand
		expectedHand deck.Hand
		expectedType poker.HandType
	}{
		{
			hand:         parseHand("4c 4d 4s 4h 5c"),
			expectedHand: parseHand("4c 4d 4s 4h 5c"),
			expectedType: poker.FourOfAKind,
		},
	}

	for _, testCase := range testCases {
		actualType, actualHand := poker.BestHand(testCase.hand)

		if actualType != testCase.expectedType {
			t.Errorf("BestHand did not return expected type.  Expected: %v.  Actual: %v.", testCase.expectedType, actualType)
		}

		if !slices.Equal(testCase.expectedHand, actualHand) {
			t.Errorf("BestHand did not return expected cards.  Expected: %v.  Actual: %v.", testCase.expectedHand, actualHand)
		}
	}
}

func Test_BestHand_ReturnsStraightFlush(t *testing.T) {
	testCases := []struct {
		hand         deck.Hand
		expectedHand deck.Hand
		expectedType poker.HandType
	}{
		{
			hand:         parseHand("6s 7s 8s 9s Ts"),
			expectedHand: parseHand("6s 7s 8s 9s Ts"),
			expectedType: poker.StraightFlush,
		},
	}

	for _, testCase := range testCases {
		actualType, actualHand := poker.BestHand(testCase.hand)

		if actualType != testCase.expectedType {
			t.Errorf("BestHand did not return expected type.  Expected: %v.  Actual: %v.", testCase.expectedType, actualType)
		}

		if !slices.Equal(testCase.expectedHand, actualHand) {
			t.Errorf("BestHand did not return expected cards.  Expected: %v.  Actual: %v.", testCase.expectedHand, actualHand)
		}
	}
}

func Test_BestHand_ReturnsRoyalFlush(t *testing.T) {
	testCases := []struct {
		hand         deck.Hand
		expectedHand deck.Hand
		expectedType poker.HandType
	}{
		{
			hand:         parseHand("Th jh qh kh ah"),
			expectedHand: parseHand("Th jh qh kh ah"),
			expectedType: poker.RoyalFlush,
		},
	}

	for _, testCase := range testCases {
		actualType, actualHand := poker.BestHand(testCase.hand)

		if actualType != testCase.expectedType {
			t.Errorf("BestHand did not return expected type.  Expected: %v.  Actual: %v.", testCase.expectedType, actualType)
		}

		if !slices.Equal(testCase.expectedHand, actualHand) {
			t.Errorf("BestHand did not return expected cards.  Expected: %v.  Actual: %v.", testCase.expectedHand, actualHand)
		}
	}
}

// func Test_ScoreHand_CorrectlyRanksHands(t *testing.T) {
// 	testCases := []struct {
// 		better      deck.Hand
// 		worse       deck.Hand
// 		description string
// 	}{
// 		{
// 			better:      parseHand("3d 6d 9h Jh As"),
// 			worse:       parseHand("3d 6d 9h Jh Ks"),
// 			description: "High card should outrank lower high card",
// 		},
// 		{
// 			better:      parseHand("3c 3h 9h 4s 7d"),
// 			worse:       parseHand("as 4c 7s 9d Qc"),
// 			description: "Pair should outrank high card",
// 		},
// 		{
// 			better:      parseHand("as ac 7s 7d Qc"),
// 			worse:       parseHand("3c 3h 9h 4s 7d"),
// 			description: "Two pair should outrank pair",
// 		},
// 		{
// 			better:      parseHand("Qc Qs Qh Th 9h"),
// 			worse:       parseHand("as ac 7s 7d Qc"),
// 			description: "Three-of-a-kind should outrank two pair",
// 		},
// 		{
// 			better:      parseHand("5c 6s 7s 8s 9s"),
// 			worse:       parseHand("Qc Qs Qh Th 9h"),
// 			description: "Straight should outrank three-of-a-kind",
// 		},
// 		{
// 			better:      parseHand("3s 4s 7s Ks Qs"),
// 			worse:       parseHand("5c 6s 7s 8s 9s"),
// 			description: "Flush should outrank straight",
// 		},
// 		{
// 			better:      parseHand("Kh Ks Kd Ad As"),
// 			worse:       parseHand("3s 4s 7s Ks Qs"),
// 			description: "Full house should outrank flush",
// 		},
// 		{
// 			better:      parseHand("4c 4h 4s 4d 5s"),
// 			worse:       parseHand("Kh Ks Kd Ad As"),
// 			description: "Four-of-a-kind house should outrank full house",
// 		},
// 		{
// 			better:      parseHand("8h 9h Th Jh Qh"),
// 			worse:       parseHand("4c 4h 4s 4d 5s"),
// 			description: "Straight flush house should outrank four-of-a-kind",
// 		},
// 		{
// 			better:      parseHand("Td Jd Qd Kd Ad"),
// 			worse:       parseHand("8h 9h Th Jh Qh"),
// 			description: "Royal flush house should outrank straight flush",
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		worseType, worseHand := poker.BestHand(testCase.worse)
// 		betterType, betterHand := poker.BestHand(testCase.better)
// 		if poker.ScoreHand(worseType, worseHand) > poker.ScoreHand(betterType, betterHand) {
// 			t.Errorf(
// 				"Score hand failed: %v.  Better hand: %v.  Worse hand: %v.",
// 				testCase.description,
// 				betterHand,
// 				worseHand)
// 		}
// 	}
// }

// Generates a hand from a string.
// Example: "2d 3c th ks" returns a slice of four cards:
//   - Card{Rank: Two, Suit: Diamonds}
//   - Card{Rank: Three, Suit: Clubs}
//   - Card{Rank: Ten, Suit: Hearts}
//   - Card{Rank: King, Suit: Spades}
//
// Helper util.  Use to make test setup code less verbose and easier to read.
func parseHand(hand string) deck.Hand {
	var result deck.Hand

	const space = " "
	elements := strings.Split(hand, space)

	for _, element := range elements {
		var rank deck.Rank
		switch element[0] {
		case 't', 'T':
			rank = deck.Ten
		case 'j', 'J':
			rank = deck.Jack
		case 'q', 'Q':
			rank = deck.Queen
		case 'k', 'K':
			rank = deck.King
		case 'a', 'A':
			rank = deck.Ace
		default:
			i, _ := strconv.ParseInt(string(element[0]), 10, 64)
			rank = deck.Rank(i)
		}

		var suit deck.Suit
		switch element[1] {
		case 'c', 'C':
			suit = deck.Clubs
		case 'd', 'D':
			suit = deck.Diamonds
		case 'h', 'H':
			suit = deck.Hearts
		case 's', 'S':
			suit = deck.Spades
		}

		result = append(result, deck.Card{Rank: rank, Suit: suit})
	}

	return deck.Hand(result)
}
