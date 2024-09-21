package house_test

import (
	"testing"

	"github.com/David-Rushton/card-collection/house"
)

func Test_PotBalance_DefaultIsZero(t *testing.T) {
	t.Cleanup(cleanup)

	actual := house.PotBalance()
	if actual != 0 {
		t.Errorf("❌ Unexpected pot size.  Expected: 0.  Actual: %v.", actual)
	}
}

func Test_Bet_IncreasedPotSize(t *testing.T) {
	t.Cleanup(cleanup)

	house.Bet(&house.Account{Balance: 1000}, 500)

	expected := 500
	actual := house.PotBalance()

	if actual != expected {
		t.Errorf("❌ Bet resulted in expected pot size.  Expected: %v.  Actual: %v.", expected, actual)
	}
}

func Test_Bet_PotSizeIsSumOfAllBets(t *testing.T) {
	t.Cleanup(cleanup)

	// Arrange
	wager := 100
	expected := 0

	for i := 0; i < 5; i++ {
		house.Bet(&house.Account{Balance: 1000}, wager)
		expected += wager
	}

	// Act and assert
	actual := house.PotBalance()
	if actual != expected {
		t.Errorf("❌ Unexpected pot size.  Expected: %v.  Actual: %v.", expected, actual)

	}
}

func Test_Bet_ReturnsErrInsufficientFunds_WhenFundsUnavailable(t *testing.T) {
	t.Cleanup(cleanup)

	account := &house.Account{Balance: 0}
	err := house.Bet(account, account.Balance+1)

	if err != house.ErrInsufficientFunds {
		t.Errorf("❌ Unexpected error.  Expected: ErrInsufficientFunds.  Actual: %v.", err)
	}
}

func Test_Bet_ReturnsNilErr_WhenFundsAvailable(t *testing.T) {
	t.Cleanup(cleanup)

	account := &house.Account{Balance: 10}
	err := house.Bet(account, account.Balance)

	if err != nil {
		t.Errorf("❌ Unexpected error.  Expected: Nil.  Actual: %v.", err)
	}
}

func Test_Payout_ReturnsZeroPot_WhenPotDividesEqually(t *testing.T) {
	t.Cleanup(cleanup)

	// Arrange
	account1 := &house.Account{Balance: 100}
	account2 := &house.Account{Balance: 100}
	account3 := &house.Account{Balance: 100}

	house.Bet(account1, 10)
	house.Bet(account2, 10)
	house.Bet(account3, 10)

	house.Payout(account1, account2, account3)

	// Act and assert
	expected := 0
	actual := house.PotBalance()

	if actual != expected {
		t.Errorf("❌ Unexpected pot size.  Expected: %v.  Actual: %v.", expected, actual)
	}
}

func Test_Payout_ReturnsPotBalanceOne_WhenPotDoesNotDivideEqually(t *testing.T) {
	t.Cleanup(cleanup)

	// Arrange
	account1 := &house.Account{Balance: 100}
	account2 := &house.Account{Balance: 100}

	house.Bet(account1, 1)
	house.Bet(account2, 2)

	house.Payout(account1, account2)

	// Act and assert
	expected := 1
	actual := house.PotBalance()

	if actual != expected {
		t.Errorf("❌ Unexpected pot size.  Expected: %v.  Actual: %v.", expected, actual)
	}
}

func Test_Payout_CreditsAccount(t *testing.T) {
	t.Cleanup(cleanup)

	// Arrange
	account1 := &house.Account{Balance: 100}
	account2 := &house.Account{Balance: 100}

	house.Bet(account1, 10)
	house.Bet(account2, 20)

	// Act
	house.Payout(account1)

	// Assert
	expected := 120 // Balance (100) - wager1 (10) - wager2 (20) + winnings (30)
	actual := account1.Balance
	if actual != expected {
		t.Errorf("❌ Unexpected pot size.  Expected: %v.  Actual: %v.", expected, actual)
	}

}

func cleanup() {
	// Paying out to one account will always clear the entire pot.
	house.Payout(&house.Account{})
}
