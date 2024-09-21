package house

import (
	"errors"
)

var (
	// Returned if a player tries to bet any amount that is greater than their balance.
	ErrInsufficientFunds = errors.New("cannot place bet, due to insufficient funds")
)
