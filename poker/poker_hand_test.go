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
		expectedName poker.HandName
	}{
		{
			hand:         parseHand("3s 5d 7c 9s Th"),
			expectedHand: parseHand("Th 9s 7c 5d 3s"),
			expectedName: poker.HighCard,
		},
	}

	for _, testCase := range testCases {
		actual := poker.BestHand(testCase.hand)

		if actual.Name != testCase.expectedName {
			t.Errorf("❌ BestHand did not return expected type.  Expected: %v.  Actual: %v.", testCase.expectedName, actual.Name)
		}

		if !slices.Equal(testCase.expectedHand, actual.Hand) {
			t.Errorf("❌ BestHand did not return expected cards.  Expected: %v.  Actual: %v.", testCase.expectedHand, actual.Hand)
		}
	}
}

func Test_BestHand_ReturnsPair(t *testing.T) {
	testCases := []struct {
		hand         deck.Hand
		expectedHand deck.Hand
		expectedName poker.HandName
	}{
		{
			hand:         parseHand("3h 3s 7d qc kc"),
			expectedHand: parseHand("3h 3s kc qc 7d"),
			expectedName: poker.Pair,
		},
	}

	for _, testCase := range testCases {
		actual := poker.BestHand(testCase.hand)

		if actual.Name != testCase.expectedName {
			t.Errorf("❌ BestHand did not return expected type.  Expected: %v.  Actual: %v.", testCase.expectedName, actual.Name)
		}

		if !slices.Equal(testCase.expectedHand, actual.Hand) {
			t.Errorf("❌ BestHand did not return expected cards.  Expected: %v.  Actual: %v.", testCase.expectedHand, actual.Hand)
		}
	}
}

func Test_BestHand_ReturnsTwoPairs(t *testing.T) {
	testCases := []struct {
		hand         deck.Hand
		expectedHand deck.Hand
		expectedName poker.HandName
	}{
		{
			hand:         parseHand("2h 2d 7c 3c 3s"),
			expectedHand: parseHand("3c 3s 2h 2d 7c"),
			expectedName: poker.TwoPairs,
		},
	}

	for _, testCase := range testCases {
		actual := poker.BestHand(testCase.hand)

		if actual.Name != testCase.expectedName {
			t.Errorf("❌ BestHand did not return expected type.  Expected: %v.  Actual: %v.", testCase.expectedName, actual.Name)
		}

		if !slices.Equal(testCase.expectedHand, actual.Hand) {
			t.Errorf("❌ BestHand did not return expected cards.  Expected: %v.  Actual: %v.", testCase.expectedHand, actual.Hand)
		}
	}
}

func Test_BestHand_ReturnsThreeOfAKind(t *testing.T) {
	testCases := []struct {
		hand         deck.Hand
		expectedHand deck.Hand
		expectedName poker.HandName
	}{
		{
			hand:         parseHand("2d 2h 2s qs 9c"),
			expectedHand: parseHand("2d 2h 2s qs 9c"),
			expectedName: poker.ThreeOfAKind,
		},
	}

	for _, testCase := range testCases {
		actual := poker.BestHand(testCase.hand)

		if actual.Name != testCase.expectedName {
			t.Errorf("❌ BestHand did not return expected type.  Expected: %v.  Actual: %v.", testCase.expectedName, actual.Name)
		}

		if !slices.Equal(testCase.expectedHand, actual.Hand) {
			t.Errorf("❌ BestHand did not return expected cards.  Expected: %v.  Actual: %v.", testCase.expectedHand, actual.Hand)
		}
	}
}

func Test_BestHand_ReturnsStraight(t *testing.T) {
	testCases := []struct {
		hand         deck.Hand
		expectedHand deck.Hand
		expectedName poker.HandName
	}{
		{
			hand:         parseHand("6h 7s 8h 9c tc"),
			expectedHand: parseHand("6h 7s 8h 9c tc"),
			expectedName: poker.Straight,
		},
		// Aces low.
		{
			hand:         parseHand("5s 4h 2h 3h Ad"),
			expectedHand: parseHand("Ad 2h 3h 4h 5s"),
			expectedName: poker.Straight,
		},
		// Aces high.
		{
			hand:         parseHand("Ac Tc Qs Jc Kd"),
			expectedHand: parseHand("Tc Jc Qs Kd Ac"),
			expectedName: poker.Straight,
		},
		// Straight at start of available cards.
		{
			hand:         parseHand("5h 3h 2s Ac 4s 8c 9d"),
			expectedHand: parseHand("Ac 2s 3h 4s 5h"),
			expectedName: poker.Straight,
		},
		// Multiple possible straights
		{
			hand:         parseHand("2s 3s 4d 5c 6s 7d 8d 9h"),
			expectedHand: parseHand("5c 6s 7d 8d 9h"),
			expectedName: poker.Straight,
		},
	}

	for _, testCase := range testCases {
		actual := poker.BestHand(testCase.hand)

		if actual.Name != testCase.expectedName {
			t.Errorf("❌ BestHand did not return expected type.  Expected: %v.  Actual: %v.", testCase.expectedName, actual.Name)
		}

		if !slices.Equal(testCase.expectedHand, actual.Hand) {
			t.Errorf("❌ BestHand did not return expected cards.  Expected: %v.  Actual: %v.", testCase.expectedHand, actual.Hand)
		}
	}
}

func Test_BestHand_ReturnsFlush(t *testing.T) {
	testCases := []struct {
		hand         deck.Hand
		expectedHand deck.Hand
		expectedName poker.HandName
	}{
		{
			hand:         parseHand("3c 7c 9c tc kc"),
			expectedHand: parseHand("3c 7c 9c tc kc"),
			expectedName: poker.Flush,
		},
	}

	for _, testCase := range testCases {
		actual := poker.BestHand(testCase.hand)

		if actual.Name != testCase.expectedName {
			t.Errorf("❌ BestHand did not return expected type.  Expected: %v.  Actual: %v.", testCase.expectedName, actual.Name)
		}

		if !slices.Equal(testCase.expectedHand, actual.Hand) {
			t.Errorf("❌ BestHand did not return expected cards.  Expected: %v.  Actual: %v.", testCase.expectedHand, actual.Hand)
		}
	}
}

func Test_BestHand_ReturnsFullHouse(t *testing.T) {
	testCases := []struct {
		hand         deck.Hand
		expectedHand deck.Hand
		expectedName poker.HandName
	}{
		{
			hand:         parseHand("ad ah as Ts Td"),
			expectedHand: parseHand("ad ah as Ts Td"),
			expectedName: poker.FullHouse,
		},
	}

	for _, testCase := range testCases {
		actual := poker.BestHand(testCase.hand)

		if actual.Name != testCase.expectedName {
			t.Errorf("❌ BestHand did not return expected type.  Expected: %v.  Actual: %v.", testCase.expectedName, actual.Name)
		}

		if !slices.Equal(testCase.expectedHand, actual.Hand) {
			t.Errorf("❌ BestHand did not return expected cards.  Expected: %v.  Actual: %v.", testCase.expectedHand, actual.Hand)
		}
	}
}

func Test_BestHand_ReturnsFourOfAKind(t *testing.T) {
	testCases := []struct {
		hand         deck.Hand
		expectedHand deck.Hand
		expectedName poker.HandName
	}{
		{
			hand:         parseHand("4c 4d 4s 4h 5c"),
			expectedHand: parseHand("4c 4d 4s 4h 5c"),
			expectedName: poker.FourOfAKind,
		},
	}

	for _, testCase := range testCases {
		actual := poker.BestHand(testCase.hand)

		if actual.Name != testCase.expectedName {
			t.Errorf("❌ BestHand did not return expected type.  Expected: %v.  Actual: %v.", testCase.expectedName, actual.Name)
		}

		if !slices.Equal(testCase.expectedHand, actual.Hand) {
			t.Errorf("❌ BestHand did not return expected cards.  Expected: %v.  Actual: %v.", testCase.expectedHand, actual.Hand)
		}
	}
}

func Test_BestHand_ReturnsStraightFlush(t *testing.T) {
	testCases := []struct {
		hand         deck.Hand
		expectedHand deck.Hand
		expectedName poker.HandName
	}{
		{
			hand:         parseHand("6s 7s 8s 9s Ts"),
			expectedHand: parseHand("6s 7s 8s 9s Ts"),
			expectedName: poker.StraightFlush,
		},
	}

	for _, testCase := range testCases {
		actual := poker.BestHand(testCase.hand)

		if actual.Name != testCase.expectedName {
			t.Errorf("❌ BestHand did not return expected type.  Expected: %v.  Actual: %v.", testCase.expectedName, actual.Name)
		}

		if !slices.Equal(testCase.expectedHand, actual.Hand) {
			t.Errorf("❌ BestHand did not return expected cards.  Expected: %v.  Actual: %v.", testCase.expectedHand, actual.Hand)
		}
	}
}

func Test_BestHand_ReturnsRoyalFlush(t *testing.T) {
	testCases := []struct {
		hand         deck.Hand
		expectedHand deck.Hand
		expectedName poker.HandName
	}{
		{
			hand:         parseHand("Th jh qh kh ah"),
			expectedHand: parseHand("Th jh qh kh ah"),
			expectedName: poker.RoyalFlush,
		},
	}

	for _, testCase := range testCases {
		actual := poker.BestHand(testCase.hand)

		if actual.Name != testCase.expectedName {
			t.Errorf("❌ BestHand did not return expected type.  Expected: %v.  Actual: %v.", testCase.expectedName, actual.Name)
		}

		if !slices.Equal(testCase.expectedHand, actual.Hand) {
			t.Errorf("❌ BestHand did not return expected cards.  Expected: %v.  Actual: %v.", testCase.expectedHand, actual.Hand)
		}
	}
}

func Test_ScoreHand_CorrectlyRanksHands(t *testing.T) {
	testCases := []struct {
		best        deck.Hand
		worst       deck.Hand
		description string
	}{
		{
			best:        parseHand("3d 6d 9h Jh As"),
			worst:       parseHand("3d 6d 9h Jh Ks"),
			description: "High card should outrank lower high card",
		},
		{
			best:        parseHand("3c 3h 9h 4s 7d"),
			worst:       parseHand("as 4c 7s 9d Qc"),
			description: "Pair should outrank high card",
		},
		{
			best:        parseHand("7s 7s Qd 6s Th"),
			worst:       parseHand("5s 5s Qd 6s Th"),
			description: "Pair should outrank worse pair",
		},
		{
			best:        parseHand("as ac 7s 7d Qc"),
			worst:       parseHand("3c 3h 9h 4s 7d"),
			description: "Two pair should outrank pair",
		},
		{
			best:        parseHand("4c 4d 5d 5h Ad"),
			worst:       parseHand("3c 3s 4c 4d 5d"),
			description: "Two pair should outrank worse two pair",
		},
		{
			best:        parseHand("Qc Qs Qh Th 9h"),
			worst:       parseHand("as ac 7s 7d Qc"),
			description: "Three-of-a-kind should outrank two pair",
		},
		{
			best:        parseHand("Qs Qd Qh 7d 6c"),
			worst:       parseHand("4s 4d 4h 7d 6c"),
			description: "Three-of-a-kind should outrank worst three-of-a-kind",
		},
		{
			best:        parseHand("5c 6s 7s 8s 9s"),
			worst:       parseHand("Qc Qs Qh Th 9h"),
			description: "Straight should outrank three-of-a-kind",
		},
		{
			best:        parseHand("5c 6s 7s 8s 9s"),
			worst:       parseHand("4s 5c 6s 7s 8s"),
			description: "Straight should outrank worse Straight",
		},
		{
			best:        parseHand("3s 4s 7s Ks Qs"),
			worst:       parseHand("5c 6s 7s 8s 9s"),
			description: "Flush should outrank straight",
		},
		{
			best:        parseHand("3s 4s 7s Ks As"),
			worst:       parseHand("3s 4s 7s Ks Qs"),
			description: "Flush should outrank worse flush",
		},
		{
			best:        parseHand("Kh Ks Kd Ad As"),
			worst:       parseHand("3s 4s 7s Ks Qs"),
			description: "Full house should outrank flush",
		},
		{
			best:        parseHand("Ah As Ad Kd Ks"),
			worst:       parseHand("Kh Ks Kd Ad As"),
			description: "Full house should outrank worse full house",
		},
		{
			best:        parseHand("4c 4h 4s 4d 5s"),
			worst:       parseHand("Kh Ks Kd Ad As"),
			description: "Four-of-a-kind should outrank full house",
		},
		{
			best:        parseHand("Tc Th Ts Td 5s"),
			worst:       parseHand("4c 4h 4s 4d 5s"),
			description: "Four-of-a-kind should outrank worse four-of-a-kind house",
		},
		{
			best:        parseHand("8h 9h Th Jh Qh"),
			worst:       parseHand("4c 4h 4s 4d 5s"),
			description: "Straight flush house should outrank four-of-a-kind",
		},
		{
			best:        parseHand("9h Th Jh Qh Kh"),
			worst:       parseHand("8h 9h Th Jh Qh"),
			description: "Straight flush house should outrank worse straight flush house",
		},
		{
			best:        parseHand("Td Jd Qd Kd Ad"),
			worst:       parseHand("9d Td Jd Qd Kd"),
			description: "Royal flush house should outrank straight flush",
		},
	}

	for _, testCase := range testCases {
		best := poker.BestHand(testCase.best)
		worst := poker.BestHand(testCase.worst)
		if worst.Score > best.Score {
			t.Errorf(
				"Score hand failed: %v.  Best hand: %v.  Worst hand: %v.",
				testCase.description,
				best.Score,
				worst.Score)
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
