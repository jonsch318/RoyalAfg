package bank

import (
	"log"

	"github.com/Rhymond/go-money"
)


//AllIn gives the all in Amount the player has to bet
func (b *Bank) AllIn(id string) *money.Money {
	b.lock.RLock()
	defer b.lock.RUnlock()
	p, ok := b.PlayerWallet[id]
	if !ok {
		return nil
	}
	bet, ok := b.PlayerBets[id]
	if !ok {
		return nil
	}

	res, err := bet.Add(p)

	if err != nil {
		return nil
	}
	log.Printf("All In %v", res.Display())
	return res
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
