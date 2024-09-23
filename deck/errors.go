package deck

import (
	"fmt"
)

type ErrNotEnoughCards struct {
	Requested int
	Remaining int
}

func (e ErrNotEnoughCards) Error() string {
	return fmt.Sprintf(
		"Cannot take %d cards.  Only %d remain.  Shuffle the deck and try again.",
		e.Requested,
		e.Remaining)
}
