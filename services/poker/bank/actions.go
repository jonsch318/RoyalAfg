package bank

//PerformBet performs a check action. This equals the players bet to the current maximum bet.
func (b *Bank) PerformBet(playerId string) error {
	return b.Bet(playerId, b.MaxBet)
}

//Perform Check
func (b *Bank) PerformAllIn(playerId string) (bool, error) {
	raise, err := b.MustAllIn(playerId)
	if err != nil {
		return false, err
	}
	//If raise == true Bet is considered a check, because the player cannot equal the max bet without going all in.
	//Else bet is considered a raise because the player can equal the max bet without going all in.
	return raise, b.Bet(playerId, b.AllIn(playerId))
}
