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

	//Get player rest wallet. (The full wallet is wallet + bet)
	rest, ok := b.PlayerWallet[id]
	if !ok {
		log.Logger.Errorw("Player not registered in bank", "id", id)
		return errors.New("player not registered in bank")
	}

	//Calculate difference of the amount and the bet.
	newAmount, err := amount.Subtract(b.PlayerBets[id])
	if err != nil {
		return err
	}


	log.Logger.Debugf("Amound: %v, AD: %v, Val: %s", amount, newAmount, rest.Display())


	//We have to validate the amount correctly so that no invalid bet can occurs.
	//First we check if the player can bet the specified amount
	if res, err2 := rest.LessThan(newAmount); res || err2 != nil {
		log.Logger.Warnf("The player %v does not have the capacity to bet %v [%v]", id, rest.Display(), amount.Display())
		return fmt.Errorf("The player does not have the capacity to bet %v ", amount)
	}

	//Check if bet is lower than the round bet and the player is not all in
	lessThanRoundBet, err := amount.LessThan(b.MaxBet)
	if err != nil {
		log.Logger.Errorw("error during money comparison", "error", err)
		return err
	}

	//We check if the player bet is allIn to check if the player is all in
	allIn, err := rest.Equals(newAmount)
	if err != nil {
		log.Logger.Errorw("error during money comparison", "error", err)
		return err
	}

	if lessThanRoundBet &&  !allIn {
		log.Logger.Errorf("Bet attemted that is not all in and lower than the current round bet %v < %v", amount.Display(), b.MaxBet.Display())
		return fmt.Errorf("the player has to bet more or equal the round bet or go all in")
	}

	return b.betTransact(id, amount, rest, newAmount)
}

func (b *Bank) betTransact(id string, amount *money.Money, rest *money.Money, newAmount *money.Money) error {
	//player can bet Amount
	newWallet, err := rest.Subtract(newAmount)
	if err != nil {
		log.Logger.Errorw("error during bet transaction", "error", err)
		return err
	}
	b.PlayerWallet[id] = newWallet

	b.PlayerBets[id] = amount

	newPot, err := b.Pot.Add(newAmount)
	if err != nil {
		log.Logger.Errorw("error during bet transaction", "error", err)
		return err
	}
	b.Pot = newPot

	if res, _ := amount.GreaterThan(b.MaxBet); res {
		b.MaxBet = amount
	}

	log.Logger.Debugf("Bet transacted")

	return nil
}