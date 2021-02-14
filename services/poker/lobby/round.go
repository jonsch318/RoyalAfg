package lobby

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
)

//RemoveAfterRound starts the recursive removal of hanging players
func (l *Lobby) RemoveAfterRound()  {
	//Lock for multithreading writes
	//we lock here so we dont need a recursive mutex lock in the next function.
	l.lock.Lock()
	defer l.lock.Unlock()

	//Remove all hanging players and update player count
	l.PlayerRemoval()
	l.SetPlayerCountLabel()

	log.Logger.Debugf("Removed Players Player Count: %v", l.PlayerCount)
}

//PrepareForRound prepares the lobby for a new Round
func (l *Lobby) PrepareForRound()  {
	//Lock for multithreading writes
	//we lock here so we dont need a recursive mutex (it's a bit of a problem) lock in the next function.
	//Could take longer for this lock to resolve
	l.lock.Lock()
	defer l.lock.Unlock()

	//Fill all slots and update player count
	l.FillLobbyPosition()
	l.SetPlayerCountLabel()

	log.Logger.Debugf("Removed Players Player Count: %v", l.PlayerCount)
}
