package poker_test

import (
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/David-Rushton/card-collection/cards"
	"github.com/David-Rushton/card-collection/poker"
)

func Test_BestHand_ReturnsHighCard(t *testing.T) {
	testCases := []struct {
		hand         cards.Hand
		expectedHand cards.Hand
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
		hand         cards.Hand
		expectedHand cards.Hand
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
		hand         cards.Hand
		expectedHand cards.Hand
		expectedType poker.HandType
	}{
		{
			hand:         parseHand("2h 2d 3c 3s 7c"),
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
		hand         cards.Hand
		expectedHand cards.Hand
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
		hand         cards.Hand
		expectedHand cards.Hand
		expectedType poker.HandType
	}{
		{
			hand:         parseHand("6h 7s 8h 9c tc"),
			expectedHand: parseHand("6h 7s 8h 9c tc"),
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
		hand         cards.Hand
		expectedHand cards.Hand
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
		hand         cards.Hand
		expectedHand cards.Hand
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
		hand         cards.Hand
		expectedHand cards.Hand
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
		hand         cards.Hand
		expectedHand cards.Hand
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
		hand         cards.Hand
		expectedHand cards.Hand
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

// Generates a hand from a string.
// Example: "2d 3c th ks" returns a slice of four cards:
//   - Card{Rank: Two, Suit: Diamonds}
//   - Card{Rank: Three, Suit: Clubs}
//   - Card{Rank: Ten, Suit: Hearts}
//   - Card{Rank: King, Suit: Spades}
//
// Helper util.  Use to make test setup code less verbose and easier to read.
func parseHand(hand string) cards.Hand {
	var result []cards.Card

	const space = " "
	elements := strings.Split(hand, space)

	for _, element := range elements {
		var rank cards.Rank
		switch element[0] {
		case 't', 'T':
			rank = cards.Ten
		case 'j', 'J':
			rank = cards.Jack
		case 'q', 'Q':
			rank = cards.Queen
		case 'k', 'K':
			rank = cards.King
		case 'a', 'A':
			rank = cards.Ace
		default:
			i, _ := strconv.ParseInt(string(element[0]), 10, 64)
			rank = cards.Rank(i)
		}

		var suit cards.Suit
		switch element[1] {
		case 'c', 'C':
			suit = cards.Clubs
		case 'd', 'D':
			suit = cards.Diamonds
		case 'h', 'H':
			suit = cards.Hearts
		case 's', 'S':
			suit = cards.Spades
		}

		result = append(result, cards.Card{Rank: rank, Suit: suit})
	}

	return cards.Hand(result)
}
