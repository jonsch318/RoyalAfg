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

