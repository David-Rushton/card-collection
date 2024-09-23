package deck_test

import (
	"testing"

	"github.com/David-Rushton/card-collection/deck"
)

func Test_Shuffle_Returns52Cards(t *testing.T) {
	deck.Shuffle()

	actual := deck.Remaining()
	if actual != 52 {
		t.Errorf("❌ Expected: 52.  Actual: %v.", actual)
	}
}

func Test_Shuffle_Returns52UniqueCards(t *testing.T) {
	deck.Shuffle()

	cards, err := deck.Take(52)
	if err != nil {
		t.Errorf("❌ Unexpected error: %v.", err)
	}

	keys := make(map[string]int)
	for _, card := range cards {
		keys[card.String()]++
	}

	actual := len(keys)
	if actual != 52 {
		t.Errorf("❌ Expected: 52.  Actual: %v.", actual)
	}
}

func Test_Take_ReturnsError_WhenNLessThanZero(t *testing.T) {
	testCases := []int{-1_000, -10, -1}

	for _, testCase := range testCases {
		deck.Shuffle()
		if _, err := deck.Take(testCase); err == nil {
			t.Errorf("❌ Missing error when requesting %v cards.", testCase)
		}
	}
}

func Test_Take_ReturnsError_WhenNLessThanRemaining(t *testing.T) {
	testCases := []int{53, 102, 100}

	for _, testCase := range testCases {
		deck.Shuffle()
		if _, err := deck.Take(testCase); err == nil {
			t.Errorf("❌ Missing ErrNotEnoughCards when requesting %v cards.", testCase)
		}
	}
}

func Test_Remaining_ReturnsExpectedCount(t *testing.T) {
	testCases := []int{52, 47, 12, 3, 0}

	for _, testCase := range testCases {
		deck.Shuffle()
		deck.Take(testCase)

		expected := 52 - testCase
		actual := deck.Remaining()
		if actual != expected {
			t.Errorf(
				"❌ Unexpected number of remaining cards, after taking: %v.  Expected: %v.  Actual: %v",
				testCase,
				expected,
				actual)
		}
	}
}
