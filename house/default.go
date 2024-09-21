// The place that always wins.
// Manages money.
package house

var (
	// The house account.
	// All bets are paid into the pot.
	// All winnings are paid out of the pot.
	pot = &Account{}
)

// Tracks a users balance with the house.
type Account struct {
	Balance int
}

// Returns the size of the pot.
func PotBalance() int {
	return pot.Balance
}

// Place your bets!
// Moves money from an account into the pot.
func Bet(account *Account, amount int) error {
	return transfer(account, pot, amount)
}

// Pot is shared equally between all players.
// If the pot cannot be split evenly the odd remains, to be won in later hands.
func Payout(accounts ...*Account) {
	share := pot.Balance / len(accounts)

	for _, account := range accounts {
		transfer(pot, account, share)
	}
}

func transfer(from, to *Account, amount int) error {
	if from.Balance < amount {
		return ErrInsufficientFunds
	}

	from.Balance -= amount
	to.Balance += amount

	return nil
}
