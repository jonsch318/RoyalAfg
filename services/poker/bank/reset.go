package bank

import "log"

//ConcludeRound resets the current round and adds the fair share of the to the winners wallets.
func (b *Bank) ConcludeRound(winners []string) int {
	b.lock.Lock()
	defer b.lock.Unlock()

	if winners == nil || b.Pot == 0 {
		return -1
	}

	share := b.Pot / len(winners)

	for _, n := range winners {
		b.PlayerWallet[n] += share
		log.Printf("User [%v] wins share %d", n, share)
	}

	b.Pot = 0
	b.MaxBet = 0
	return share
}

//Reset resets the state of the Bank for a new round
func (b *Bank) Reset() {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.Pot = 0
	b.MaxBet = 0
	for id := range b.PlayerBets {
		b.PlayerBets[id] = 0
	}
}
