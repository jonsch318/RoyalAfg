package bank

import (
	"errors"
	"fmt"

	"github.com/Rhymond/go-money"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
)

//bet handles the betting process for a given player and Amount
func (b *Bank) bet(id string, amount *money.Money) error {

	b.lock.Lock()
	defer b.lock.Unlock()

	playerValue, ok := b.PlayerWallet[id]
	if !ok {
		log.Logger.Errorw("Player not registered in bank", "id", id)
		return errors.New("player not registered in bank")
	}

	ad, err := amount.Subtract(b.PlayerBets[id])

	if err != nil {
		return err
	}

	log.Logger.Debugf("Amound: %v, AD: %v, Val: %s", amount, ad, playerValue.Display())

	if res, err := playerValue.LessThan(ad); res || err != nil {
		log.Logger.Warnf("The player %v does not have the capacity to bet %v [%v]", id, playerValue.Display(), amount.Display())
		return fmt.Errorf("The player does not have the capacity to bet %v ", amount)
	}

	less, err  := playerValue.LessThan(ad)
	if err != nil {
		log.Logger.Warnw("error asserting monetary value", "error", err)
		return err
	}
	equals, err := playerValue.Equals(amount)
	if err != nil {
		log.Logger.Warnw("error asserting monetary value", "error", err)
		return err
	}
	if less && !equals {
		// Player bet is les than round bet and is not an all in => invalid
		log.Logger.Warnf("the player has to bet more or equal the round bet or do an all in")
		return errors.New("the player has to bet more or equal the round bet or do an all in")
	}

	log.Logger.Debugf("All validation checks passed now transacting bet")

	//player can bet Amount
	b.PlayerWallet[id], _ = playerValue.Subtract(ad)
	b.PlayerBets[id] = amount

	b.Pot, _ = b.Pot.Add(ad)
	if res, _ := amount.GreaterThan(b.MaxBet); res {
		b.MaxBet = amount
	}

	log.Logger.Debugf("transaction done publishing event")

	//Could use b.PlayerBet but this works as well
	//b.AddBetEvent(id, ad)
	//We use b.PlayerBet for simplicity. Command Queue could include more information.
	return nil
}
