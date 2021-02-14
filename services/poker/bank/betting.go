package bank

import (
	"errors"
	"fmt"
	"log"

	"github.com/Rhymond/go-money"
)

//Bet handles the betting process for a given player and Amount
func (b *Bank) Bet(id string, amount *money.Money) error {

	b.lock.Lock()
	defer b.lock.Unlock()

	playerValue, ok := b.PlayerWallet[id]
	if !ok {

		log.Printf("Player not registered in bank")
		return errors.New("player not registered in bank")
	}

	ad, err := amount.Subtract(b.PlayerBets[id])

	if err != nil {
		return err
	}

	log.Printf("Amound: %v, AD: %v, Val: %d", amount, ad, playerValue)


	if res, err := playerValue.LessThan(ad); res || err != nil {
		if err != nil {
			return err
		}
		log.Printf("The player %v does not have the capacity to bet %v [%v]", id, playerValue, amount)
		return fmt.Errorf("The player does not have the capacity to bet %v ", amount)
	}

	less, err  := playerValue.LessThan(ad)
	if err != nil {
		return err
	}
	equals, err := playerValue.Equals(amount)
	if err != nil {
		return err
	}
	if less && !equals {
		// Player bet is les than round bet and is not an all in => invalid
		return errors.New("the player has to bet more or equal the round bet or do an all in")
	}

	//player can bet Amount
	b.PlayerWallet[id], _ = playerValue.Subtract(ad)
	b.PlayerBets[id] = amount

	b.Pot, _ = b.Pot.Add(ad)
	if res, _ := amount.GreaterThan(b.MaxBet); res {
		b.MaxBet = amount
	}

	b.AddBetEvent(id, int(amount.Amount()))

	return nil
}
