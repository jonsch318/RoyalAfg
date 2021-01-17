package bank

import (
	"errors"
	"fmt"
	"log"
)

//Bet handles the betting process for a given player and Amount
func (b *Bank) Bet(id string, amount int) error {

	b.lock.Lock()
	defer b.lock.Unlock()

	playerValue, ok := b.PlayerWallet[id]

	ad := amount - b.PlayerBets[id]

	log.Printf("Amound: %v, AD: %v, Val: %d", amount, ad, playerValue)

	if !ok {
		log.Printf("Player not registered in bank")
		return errors.New("player not registered in bank")
	}

	if playerValue < ad {
		log.Printf("The player %v does not have the capacity to bet %v [%v]", id, playerValue, amount)
		return fmt.Errorf("The player does not have the capacity to bet %v ", amount)
	}

	if amount < b.MaxBet && playerValue != amount {
		// Player bet is les than round bet and is not an all in => invalid
		return errors.New("the player has to bet more or equal the round bet or do an all in")
	}

	//player can bet Amount
	b.PlayerWallet[id] = playerValue - ad
	b.PlayerBets[id] = amount

	b.Pot += ad
	if amount > b.MaxBet {
		b.MaxBet = amount
	}

	b.AddBetEvent(id, amount)

	return nil
}
