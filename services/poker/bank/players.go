package bank

import (
	"errors"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	moneyUtils "github.com/JohnnyS318/RoyalAfgInGo/services/poker/money"
)

//AddPlayer adds a given player to the bank
func (b *Bank) AddPlayer(player *models.Player) {
	b.lock.Lock()
	defer b.lock.Unlock()

	log.Logger.Debugf("BANK: Adding player[%v] with BuyIn: %v", player.ID, player.BuyIn.Display())
	b.PlayerWallet[player.ID] = player.BuyIn
	b.PlayerBets[player.ID] = moneyUtils.Zero()
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
	log.Logger.Infof("Tried removing player [%v] which is not registered in the bank", id)
	return errors.New("player not registered in bank")
}

//UpdatePublicPlayerBuyIn updates the buyIns of the public player arrays according to the current state.
func (b *Bank) UpdatePublicPlayerBuyIn(p []models.PublicPlayer) {
	b.lock.RLock()
	defer b.lock.RUnlock()
	for i := range p {
		t, ok := b.PlayerWallet[p[i].ID]
		if !ok {
			log.Logger.Warnf("Player [%v] has no wallet", p[i].ID)
			continue
		}
		p[i].SetBuyIn(t.Display(), t.AsMajorUnits())
		log.Logger.Debugf("Player [%v] wallet is %v ", p[i].Username, p[i].BuyIn)
	}
}
