package cards_test

import (
	"testing"

	"example.com/cards/cards"
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
		t.Errorf("Cannot take 52 cards from freshly shuffled deck.  Beause %v", err)
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
