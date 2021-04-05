package lobby

//PrepareForRound prepares the lobby for a new Round
func (l *Lobby) PrepareForRound() {
	//Lock for multithreading writes
	//we lock here so we dont need a recursive mutex (it's a bit of a problem) lock in the next function.
	//Could take longer for this lock to resolve
	l.lock.Lock()
	defer l.lock.Unlock()

	//Fill all slots and update player count
	l.FillLobbyPosition()
	l.SetPlayerCountLabel()

	for i := range l.Players {
		// Set player states if player can bet in the round
		if l.Bank.HasZeroWallet(l.Players[i].ID) || l.Players[i].Left {
			l.Players[i].Active = false
		} else {
			l.Players[i].Active = true
		}
	}

}
