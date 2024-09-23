package deck

import "fmt"

type Suit int

const (
	Clubs Suit = iota + 1
	Diamonds
	Hearts
	Spades
)

type Rank int

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Rank Rank
	Suit Suit
}

func (c *Card) String() string {
	var suit string
	// ♣️ ♦️ ♥️ ♠️
	switch c.Suit {
	case Clubs:
		suit = "Clubs"
	case Diamonds:
		suit = "Diamonds"
	case Hearts:
		suit = "Hearts"
	case Spades:
		suit = "Spades"
	}

	var rank string
	switch c.Rank {
	case Ace:
		rank = "Ace"
	case Jack:
		rank = "Jack"
	case Queen:
		rank = "Queen"
	case King:
		rank = "King"
	default:
		rank = fmt.Sprintf("%d", int(c.Rank))
	}

	return fmt.Sprintf("%v of %v", rank, suit)
}
