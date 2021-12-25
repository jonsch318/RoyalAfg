package bank

import (
	"errors"
	"fmt"

	"github.com/Rhymond/go-money"
)

//amountIsRaise checks if the specified amount qualifies as a raise action. It has to be greater than the original maximum bet
func (b *Bank) amountIsRaise(amount *money.Money) bool {
	//Error can neglected because it will return false if an error occurs.
	res, _ := amount.GreaterThan(b.MaxBet)
	return res
}

//amountIsCall checks if the specified amount qualifies as a call action. It has to equal than the original maximum bet
func (b *Bank) amountIsCall(amount *money.Money) bool {
	//Error can neglected because it will return false if an error occurs.
	res, _ := amount.Equals(b.MaxBet)
	return res
}

//amountIsAllIn checks if the specified amount qualifies as a all in for the given player action. It has to be greater or equal the player all in amount
func (b *Bank) amountIsAllIn(id string, amount *money.Money) (bool, error) {
	_, ok := b.PlayerWallet[id]
	_, ok2 := b.PlayerBets[id]
	if !ok || ok2 {
		return false, fmt.Errorf("player [%v] is not registered in the bank", id)
	}
	//Error can neglected because it will return false if an error occurs.
	res, err := amount.GreaterThanOrEqual(b.allIn(id))
	if err != nil {
		return false, err
	}

	return res && !b.IsAllIn(id), nil
}

//MustAllIn determines whether a player has to bet everything in because the maximum bet is already past his wallet amount
func (b *Bank) MustAllIn(id string) (bool, error) {
	b.lock.RLock()
	defer b.lock.RUnlock()
	p, ok := b.PlayerWallet[id]
	if !ok {
		return false, errors.New("the player was not found")
	}
	bet, ok := b.PlayerBets[id]
	if !ok {
		return false, errors.New("the player was not found")
	}
	add, err := bet.Add(p)
	if err != nil {
		return false, err
	}
	return b.MaxBet.GreaterThanOrEqual(add)
}

//IsAllIn determines whether a given player has already placed all his wallet. He can be excluded from the blocking list
func (b *Bank) IsAllIn(id string) bool {
	b.lock.RLock()
	defer b.lock.RUnlock()
	w, ok := b.PlayerWallet[id]
	if !ok {
		return true
	}

	bet, ok := b.PlayerBets[id]
	if !ok {
		return true
	}

	return w.IsZero() && bet.IsPositive()
}
