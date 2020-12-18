package bank

import (
	"errors"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
)

//AddPlayer adds a given player to the bank
func (b *Bank) AddPlayer(player *models.Player) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.PlayerWallet[player.ID] = player.BuyIn
	b.PlayerBets[player.ID] = 0
}

//RemovePlayer removes the given player from the bank
func (b *Bank) RemovePlayer(id string) error {
	b.lock.Lock()
	defer b.lock.Unlock()
	_, ok := b.PlayerWallet[id]
	if ok {
		delete(b.PlayerWallet, id)
		return nil
	}
	return errors.New("Player not registered in bank")
}

//UpdatePublicPlayerBuyIn updates the buyIns of the public player arrays according to the current state.
func (b *Bank) UpdatePublicPlayerBuyIn(p []models.PublicPlayer) {
	b.lock.RLock()
	defer b.lock.RUnlock()
	for i := 0; i < len(p); i++ {
		t, ok := b.PlayerWallet[p[i].ID]
		if !ok {
			continue
		}
		p[i].BuyIn = float32(t) / 100
	}
}
