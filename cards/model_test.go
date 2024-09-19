package cards_test

import (
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/David-Rushton/card-collection/cards"
)

func Test_Shuffle_Generates52Cards(t *testing.T) {
	cards.Shuffle()

	expected := 52
	actual := cards.Remaining()

	if actual != expected {
		t.Errorf("Expected 52 cards in a freshly shuffled deck.  Found: %v", actual)
	}
}

func Test_Take_Returns52UniqueCardsWhenShuffled(t *testing.T) {
	cards.Shuffle()

	uniqueCards := make(map[cards.Card]int)
	hand, err := cards.Take(52)

	if err != nil {
		t.Errorf("Cannot take 52 cards from freshly shuffled deck.  Because %v", err)
	}

	for _, card := range hand {
		uniqueCards[card] = 1
	}

	if len(uniqueCards) != 52 {
		t.Errorf("Take returned duplicate cards from freshly shuffled deck.")
	}
}

func Test_Remaining_ReturnsExpectedNumber(t *testing.T) {
	testCases := []struct {
		take      int
		remaining int
	}{
		{
			take:      2,
			remaining: 50,
		},
		{
			take:      10,
			remaining: 40,
		},
		{
			take:      15,
			remaining: 25,
		},
		{
			take:      25,
			remaining: 0,
		},
	}

	cards.Shuffle()

	for _, testCase := range testCases {
		cards.Take(testCase.take)
		expected := testCase.remaining
		actual := cards.Remaining()

		if actual != expected {
			t.Errorf("Unexpected number of cards remain.  Expected %d.  Actual: %d.", expected, actual)
		}
	}
}

func Test_Sort_ReturnsExpectedResult(t *testing.T) {
	testCases := []struct {
		name     string
		hand     cards.Hand
		expected cards.Hand
	}{
		{
			name:     "odd number of cards",
			hand:     parseHand("4d 3d 2d"),
			expected: parseHand("2d 3d 4d"),
		},
		{
			name:     "even number of cards",
			hand:     parseHand("4s 7c 6h 5d"),
			expected: parseHand("4s 5d 6h 7c"),
		},
		{
			name:     "supports lot of cards",
			hand:     parseHand("7s 3s 2s As 4s Js 5s Qs Ks 6s 8s 9s"),
			expected: parseHand("2s 3s 4s 5s 6s 7s 8s 9s Js Qs Ks As"),
		},
		{
			name:     "supports single cards",
			hand:     parseHand("qh"),
			expected: parseHand("qh"),
		},
		{
			name:     "does not fail when passed an empty hand",
			hand:     cards.Hand{},
			expected: cards.Hand{},
		},
		{
			name:     "already sorted",
			hand:     parseHand("2h 3c 4c 5c 6d 7d 8h 9c Ts Jc Qh Kd As"),
			expected: parseHand("2h 3c 4c 5c 6d 7d 8h 9c Ts Jc Qh Kd As"),
		},
		{
			name:     "big gaps",
			hand:     parseHand("As 2s"),
			expected: parseHand("2s As"),
		},
	}

	for _, testCase := range testCases {
		actual := testCase.hand.Sort()
		if !slices.Equal(testCase.expected, actual) {
			t.Errorf(
				"Sort produced unexpected result.  Case: %v.  Expected: %v.  Actual: %v.",
				testCase.name,
				testCase.expected,
				actual)
		}
	}
}

func Test_Sort_AcesAreHigh(t *testing.T) {
	testCases := []struct {
		hand     cards.Hand
		expected cards.Hand
	}{
		{
			hand:     parseHand("Ad Kd Qd"),
			expected: parseHand("Qd Kd Ad"),
		},
		{
			hand:     parseHand("Ac 4d 5h 6s"),
			expected: parseHand("4d 5h 6s Ac"),
		},
		{
			hand:     parseHand("4d Ac Kh 6s 5h"),
			expected: parseHand("4d 5h 6s Kh Ac"),
		},
	}

	for _, testCase := range testCases {
		actual := testCase.hand.Sort()
		if !slices.Equal(testCase.expected, actual) {
			t.Errorf("Sort produced unexpected result.  Expected: %v.  Actual: %v.", testCase.expected, actual)
		}
	}
}

func Test_Sort_StableSortsSuits(t *testing.T) {
	testCases := []struct {
		hand     cards.Hand
		expected cards.Hand
	}{
		{
			hand:     parseHand("ac ah as ad"),
			expected: parseHand("ac ah as ad"),
		},
		{
			hand:     parseHand("4h 4c js 2d"),
			expected: parseHand("2d 4h 4c js"),
		},
		{
			hand:     parseHand("qh 4d 4s 4c 3h"),
			expected: parseHand("3h 4d 4s 4c qh"),
		},
	}

	for _, testCase := range testCases {
		actual := testCase.hand.Sort()
		if !slices.Equal(testCase.expected, actual) {
			t.Errorf("Sort produced unexpected result.  Expected: %v.  Actual: %v.", testCase.expected, actual)
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
