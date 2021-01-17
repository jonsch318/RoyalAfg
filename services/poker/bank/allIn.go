package bank

import (
	"errors"
	"log"
)

//MustAllIn determins whether a player has to bet all the buyin because the maximum bet is already past his Amount
func (b *Bank) MustAllIn(id string) (bool, error) {
	b.lock.RLock()
	defer b.lock.RUnlock()
	p, ok := b.PlayerWallet[id]
	if !ok {
		return false, errors.New("The player was not found")
	}
	bet, ok := b.PlayerBets[id]
	if !ok {
		return false, errors.New("The player was not found")
	}
	return b.MaxBet >= (p + bet), nil
}

//AllIn gives the all in Amount the player has to bet
func (b *Bank) AllIn(id string) int {
	b.lock.RLock()
	defer b.lock.RUnlock()
	p, ok := b.PlayerWallet[id]
	if !ok {
		return -1
	}
	bet, ok := b.PlayerBets[id]
	if !ok {
		return -1
	}
	log.Printf("All In %v", bet+p)
	return bet + p
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

	return w == 0 && bet > 0
}
