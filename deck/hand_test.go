package deck_test

import (
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/David-Rushton/card-collection/deck"
)

func Test_Append_JoinsTwoHands(t *testing.T) {
	testCases := []struct {
		hand1Len    int
		hand2Len    int
		take        int
		expectedLen int
	}{
		// take == hand2 size
		{
			hand1Len:    1,
			hand2Len:    1,
			take:        1,
			expectedLen: 2,
		},
		{
			hand1Len:    5,
			hand2Len:    5,
			take:        5,
			expectedLen: 10,
		},
		{
			hand1Len:    10,
			hand2Len:    11,
			take:        11,
			expectedLen: 21,
		},
		{
			hand1Len:    1,
			hand2Len:    0,
			take:        0,
			expectedLen: 1,
		},
		{
			hand1Len:    0,
			hand2Len:    0,
			take:        0,
			expectedLen: 0,
		},
		// take < hand2 size
		{
			hand1Len:    0,
			hand2Len:    0,
			take:        10,
			expectedLen: 0,
		},
		{
			hand1Len:    1,
			hand2Len:    5,
			take:        10,
			expectedLen: 6,
		},
		// take > hand2 size
		{
			hand1Len:    5,
			hand2Len:    5,
			take:        10,
			expectedLen: 10,
		},
		{
			hand1Len:    0,
			hand2Len:    5,
			take:        100,
			expectedLen: 5,
		},
		{
			hand1Len:    5,
			hand2Len:    0,
			take:        1_000,
			expectedLen: 5,
		},
	}

	for _, testCase := range testCases {
		deck.Shuffle()

		hand1, _ := deck.Take(testCase.hand1Len)
		hand2, _ := deck.Take(testCase.hand2Len)
		result := hand1.Append(hand2, testCase.hand2Len)

		expected := len(hand1) + len(hand2)
		actual := len(result)

		if actual != expected {
			t.Errorf("❌ Append expected: %v.  Actual: %v.", expected, actual)
		}
	}
}

func Test_AppendWhen(t *testing.T) {
	testCases := []struct {
		hand1     string
		hand2     string
		expected  string
		take      int
		predicate func(deck.Card) bool
	}{
		// Take all
		{
			hand1:    "aH 2c",
			hand2:    "3d 4s",
			expected: "aH 2c 3d 4s",
			take:     2,
			predicate: func(c deck.Card) bool {
				return true
			},
		},
		// Take none
		{
			hand1:    "Th 7s",
			hand2:    "9c 6d",
			expected: "Th 7s",
			take:     2,
			predicate: func(c deck.Card) bool {
				return false
			},
		},
		// n < 0
		{
			hand1:    "aH 2c",
			hand2:    "3d 4s",
			expected: "aH 2c",
			take:     -1,
			predicate: func(c deck.Card) bool {
				return true
			},
		},
		// Take requested
		{
			hand1:    "Ac Ad",
			hand2:    "Ah As",
			expected: "Ac Ad Ah",
			take:     2,
			predicate: func(c deck.Card) bool {
				return c.Rank == deck.Ace && c.Suit == deck.Hearts
			},
		},
	}

	for _, testCase := range testCases {
		hand1 := parseHand(testCase.hand1)
		hand2 := parseHand(testCase.hand2)
		expected := parseHand(testCase.expected)
		actual := hand1.AppendWhen(hand2, testCase.take, testCase.predicate)

		if !slices.Equal(actual, expected) {
			t.Errorf("❌ Append when failed.  Expected: %v.  Actual: %v.", expected, actual)
		}
	}
}

func Test_Take(t *testing.T) {
	testCases := []struct {
		startingHand deck.Hand
		take         int
		expectedHand deck.Hand
	}{
		{
			startingHand: parseHand("5c Ts Ad 7s 7d Kc"),
			take:         4,
			expectedHand: parseHand("5c Ts Ad 7s"),
		},
		// n <= 0 == none
		{
			startingHand: parseHand("Ks Qc 2h 3h"),
			take:         -1,
			expectedHand: deck.Hand{},
		},
		{
			startingHand: parseHand("7s 4c 2h"),
			take:         0,
			expectedHand: deck.Hand{},
		},
		// n > len == all
		{
			startingHand: parseHand("3s 7d Ad"),
			take:         100_000,
			expectedHand: parseHand("3s 7d Ad"),
		},
	}

	for _, testCase := range testCases {
		actual := testCase.startingHand.Take(testCase.take)
		if !slices.Equal(actual, testCase.expectedHand) {
			t.Errorf("❌ Take failed.  Expected: %v.  Actual: %v.", testCase.expectedHand, actual)
		}
	}
}

func Test_TakeWhen(t *testing.T) {
	testCases := []struct {
		startingHand deck.Hand
		expectedHand deck.Hand
		take         int
		predicate    func(deck.Card) bool
	}{
		// Take anything
		{
			startingHand: parseHand("3h Ts 2c Ks"),
			expectedHand: parseHand("3h Ts 2c Ks"),
			take:         100,
			predicate: func(c deck.Card) bool {
				return true
			},
		},
		// Take nothing
		{
			startingHand: parseHand("Ac 6c 3s"),
			expectedHand: deck.Hand{},
			take:         200,
			predicate: func(c deck.Card) bool {
				return false
			},
		},
		// Take first n matches
		{
			startingHand: parseHand("Th 4h Ah 3h"),
			expectedHand: parseHand("Th 4h"),
			take:         2,
			predicate: func(c deck.Card) bool {
				return c.Suit == deck.Hearts
			},
		},
		// n < 0
		{
			startingHand: parseHand("3c 7c Ts 4c 4d Ah"),
			expectedHand: deck.Hand{},
			take:         0,
			predicate: func(c deck.Card) bool {
				return true
			},
		},
	}

	for _, testCase := range testCases {
		actual := testCase.startingHand.TakeWhen(testCase.take, testCase.predicate)
		if !slices.Equal(actual, testCase.expectedHand) {
			t.Errorf("❌ TakeWhen failed.  Expected: %v.  Actual: %v.", testCase.expectedHand, actual)
		}
	}
}

func Test_Sort(t *testing.T) {
	testCases := []struct {
		unorderedHand deck.Hand
		expectedHand  deck.Hand
	}{
		{
			unorderedHand: parseHand("6h 2c 9s 7s"),
			expectedHand:  parseHand("2c 6h 7s 9s"),
		},
		// suit is stable sorted
		{
			unorderedHand: parseHand("Tc 7c Th 2d Td Ts"),
			expectedHand:  parseHand("2d 7c Tc Th Td Ts"),
		},
		// aces are high
		{
			unorderedHand: parseHand("Ah 3s Qd"),
			expectedHand:  parseHand("3s Qd Ah"),
		},
		// no cards
		{
			unorderedHand: deck.Hand{},
			expectedHand:  deck.Hand{},
		},
		// 1 card
		{
			unorderedHand: parseHand("6h"),
			expectedHand:  parseHand("6h"),
		},
		// lots of cards
		{
			unorderedHand: parseHand("5c Qd Kh 3c Ah 4d 9d 6s Jd Th 2c 7s 8h"),
			expectedHand:  parseHand("2c 3c 4d 5c 6s 7s 8h 9d Th Jd Qd Kh Ah"),
		},
	}

	for _, testCase := range testCases {
		actual := testCase.unorderedHand.Sort()
		if !slices.Equal(actual, testCase.expectedHand) {
			t.Errorf("❌ Sort failed.  Expected: %v.  Actual: %v.", testCase.expectedHand, actual)
		}
	}
}

// Returns a hand parsed from short codes.
// Example: 9d Tc Js Qh == 9 of diamonds, ten of clubs, jack of spades and queen of clubs.
func parseHand(requested string) deck.Hand {
	var result deck.Hand

	cards := strings.Split(requested, " ")
	for _, c := range cards {
		var rank deck.Rank
		switch c[0] {
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
			num, _ := strconv.ParseInt(string(c[0]), 10, 64)
			rank = deck.Rank(num)
		}

		var suit deck.Suit
		switch c[1] {
		case 'c', 'C':
			suit = deck.Clubs
		case 'd', 'D':
			suit = deck.Diamonds
		case 'h', 'H':
			suit = deck.Hearts
		case 's', 'S':
			suit = deck.Spades
		}

		result = append(result, deck.Card{rank, suit})
	}

	return result
}
