package bank

import (
	"fmt"

	"github.com/Rhymond/go-money"
)

//PerformBet performs a check action. This equals the players bet to the current maximum bet.
func (b *Bank) PerformBet(playerId string) error {
	return b.Bet(playerId, b.MaxBet)
}

//PerformRaise checks if it
func (b *Bank) PerformRaise(playerId string, amount *money.Money) error {
	if !b.IsRaise(amount) {
		return fmt.Errorf("the specified amount is not higher than the highest bet")
	}
	return b.Bet(playerId, amount)
}

//Perform Check
func (b *Bank) PerformAllIn(playerId string) (bool, error) {
	raise, err := b.MustAllIn(playerId)
	if err != nil {
		return false, err
	}
	//If raise == true Bet is considered a check, because the player cannot equal the max bet without going all in.
	//Else bet is considered a raise because the player can equal the max bet without going all in.
	return raise, b.Bet(playerId, b.AllIn(playerId))
}
