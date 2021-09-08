package bank

import (
	"fmt"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/Rhymond/go-money"
)

//PerformBet performs a check action. This equals the players bet to the current maximum bet.
func (b *Bank) PerformBet(id string) error {
	return b.bet(id, b.MaxBet)
}

//PerformRaise checks if it
func (b *Bank) PerformRaise(id string, amount *money.Money) (int, error) {

	//Check for regular call
	if b.amountIsCall(amount) {
		log.Logger.Debugf("Raise equals current bet. Calling")
		return 2, b.PerformBet(id)
	}

	//check for all in
	if allIn, err := b.amountIsAllIn(id, amount); err == nil && allIn {
		return 5, b.bet(id, b.allIn(id))
	}

	//check for raise
	if b.amountIsRaise(amount) {
		return 3, b.bet(id, amount)
	}

	return 3, fmt.Errorf("the specified amount is not higher than the highest bet or something else was wrong")
}

//PerformAllIn
func (b *Bank) PerformAllIn(id string) (bool, error) {
	raise, err := b.MustAllIn(id)
	if err != nil {
		return false, err
	}
	//If raise == true Bet is considered a call, because the player cannot equal the max bet without going all in.
	//Else bet is considered a raise because the player can equal the max bet without going all in.
	return raise, b.bet(id, b.allIn(id))
}

//PerformBlind is a wrapper around the core bet() function
func (b *Bank) PerformBlind(id string, blind *money.Money) error {
	return b.bet(id, blind)
}
