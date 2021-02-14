package bank

import (
	"log"

	moneyUtils "github.com/JohnnyS318/RoyalAfgInGo/services/poker/money"
)

//ConcludeRound resets the current round and adds the fair share of the to the winners wallets.
func (b *Bank) ConcludeRound(winners []string) []string {
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

	for i, n := range winners {
		res, err := b.PlayerWallet[n].Add(shares[i])

		if err != nil {
			return nil
		}
		b.PlayerWallet[n] = res

		b.AddWinEvent(n, int(shares[i].Amount()))
		log.Printf("User [%v] wins share %d", n, shares[i])
		ret[i] = shares[i].Display()
	}

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
