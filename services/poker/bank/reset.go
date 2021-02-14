package bank

import (
	"log"

	moneyUtils "github.com/JohnnyS318/RoyalAfgInGo/services/poker/money"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/showdown"
)

//ConcludeRound resets the current round and adds the fair share of the to the winners wallets.
func (b *Bank) ConcludeRound(winners []showdown.WinnerInfo) []string {
	b.lock.Lock()
	defer b.lock.Unlock()

	if winners == nil || b.Pot.IsZero() || b.Pot.IsNegative() {
		return nil
	}

	winnersCount := len(winners)
	shares, err := b.Pot.Split(winnersCount)
	ret := make([]string, winnersCount)

	if err != nil {
		return nil
	}

	for i, player := range winners {
		res, err := b.PlayerWallet[player.Player.ID].Add(shares[i])

		if err != nil {
			return nil
		}
		b.PlayerWallet[player.Player.ID] = res

		//Adds the winning amount to the command queue, so that it will be compressed with the expenses.
		b.AddWinEvent(player.Player.ID, int(shares[i].Amount()))
		log.Printf("User [%v] wins share %d", player, shares[i])
		ret[i] = shares[i].Display()
	}

	//Will send the compressed commands to the rabbitmq message broker, so that the bank service will transact these changes.
	//We do this this way to add resiliency, so that when this service crashes no money will be lost, because everything is compressed into one command which is published at the end of the game.
	b.ExecuteQueue()

	b.Pot = moneyUtils.Zero()
	b.MaxBet = moneyUtils.Zero()
	return ret
}

//Reset resets the state of the Bank for a new round
func (b *Bank) Reset() {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.Pot = moneyUtils.Zero()
	b.MaxBet = moneyUtils.Zero()
	for id := range b.PlayerBets {
		b.PlayerBets[id] = moneyUtils.Zero()
	}
}
