package bank

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	moneyUtils "github.com/JohnnyS318/RoyalAfgInGo/services/poker/money"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/showdown"
)

//ConcludeRound resets the current round and adds the fair share of the to the winners wallets.
func (b *Bank) ConcludeRound(winners []showdown.WinnerInfo) []string {
	b.lock.Lock()




	if winners == nil || b.Pot.IsZero() || b.Pot.IsNegative() {
		log.Logger.Errorw("something went wrong so that the round cannot be concluded", "winners", winners, "pot", b.Pot.Display())
		return nil
	}

	//Calculate shares (most of the time 1 share)
	shares, err := b.Pot.Split(len(winners))
	if err != nil {
		log.Logger.Errorw("Could not split into shares", "winners", winners)
		return nil
	}

	//Add share to all round winners
	ret := make([]string, len(winners))
	for i, player := range winners {
		res, err := b.PlayerWallet[player.Player.ID].Add(shares[i])

		if err != nil {
			return nil
		}
		b.PlayerWallet[player.Player.ID] = res

		//Adds the winning amount to the command queue, so that it will be compressed with the expenses.
		b.AddWinEvent(player.Player.ID, shares[i])
		log.Logger.Infof("User [%v] wins share %s", player, shares[i].Display())
		ret[i] = shares[i].Display()
	}

	//Will send the compressed commands to the rabbitmq message broker, so that the bank service will transact these changes.
	//We do this this way to add resiliency, so that when this service crashes no money will be lost, because everything is compressed into one command which is published at the end of the game.
	b.ExecuteQueue()
	b.lock.Unlock()

	//Reset Bank values like pot, max bet, player bets etc...
	b.Reset()

	return ret
}

//Reset resets the state of the Bank for a new round
func (b *Bank) Reset() {
	log.Logger.Debugf("Reseting bank for a new round")
	b.lock.Lock()
	defer b.lock.Unlock()
	b.Pot = moneyUtils.Zero()
	b.MaxBet = moneyUtils.Zero()
	for id := range b.PlayerBets {
		b.PlayerBets[id] = moneyUtils.Zero()
	}
}
