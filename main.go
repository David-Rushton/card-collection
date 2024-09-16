// DEBUG: We should remove this command later.
package main

import (
	"example.com/cards/cards"
	"example.com/cards/poker"
)

func main() {
	cards.Shuffle()
	hand := cards.Hand{
		cards.Card{Rank: cards.Two, Suit: cards.Hearts},
		cards.Card{Rank: cards.Three, Suit: cards.Diamonds},
		cards.Card{Rank: cards.Four, Suit: cards.Clubs},
		cards.Card{Rank: cards.Five, Suit: cards.Diamonds},
		cards.Card{Rank: cards.Six, Suit: cards.Spades},
		cards.Card{Rank: cards.Seven, Suit: cards.Diamonds},
	}

	poker.BestHand(hand)
}
