package bank

import (
	"fmt"

	"github.com/Rhymond/go-money"
)

//PerformBet performs a check action. This equals the players bet to the current maximum bet.
func (b *Bank) PerformBet(id string) error {
	return b.bet(id, b.MaxBet)
}

//PerformRaise checks if it
func (b *Bank) PerformRaise(id string, amount *money.Money) error {
	if res, err := amount.GreaterThan(b.MaxBet); err != nil || !res {
		return fmt.Errorf("the specified amount is not higher than the highest bet")
	}
	return b.bet(id, amount)
}

//Perform Check
func (b *Bank) PerformAllIn(id string) (bool, error) {
	raise, err := b.MustAllIn(id)
	if err != nil {
		return false, err
	}
	//If raise == true Bet is considered a check, because the player cannot equal the max bet without going all in.
	//Else bet is considered a raise because the player can equal the max bet without going all in.
	return raise, b.bet(id, b.allIn(id))
}

//PerformBlind is a wrapper around the core bet() function
func (b *Bank) PerformBlind(id string, blind *money.Money) error {
	return b.bet(id, blind)
}
