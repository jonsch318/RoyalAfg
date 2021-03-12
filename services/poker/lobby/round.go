package lobby

import (
	"time"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
)

//RemoveAfterRound starts the recursive removal of hanging players
func (l *Lobby) RemoveAfterRound()  {
	//Lock for multithreading writes
	//we lock here so we dont need a recursive mutex lock in the next function.
	l.lock.Lock()
	defer l.lock.Unlock()

	//Remove players that left during the round.
	for _, player := range l.Players {
		if player.Left {
			l.RemovalQueue.Enqueue(&player)
		}
	}

	//Remove all hanging players and update player count
	l.PlayerRemoval()
	l.SetPlayerCountLabel()

	log.Logger.Debugf("Removed PlayerCount Player Count: %v", l.PlayerCount)
	if l.Count() <= 0 {
		log.Logger.Warnf("No more players... Notify agones sdk to shutdown")
		t := time.NewTimer(time.Second * 10)
		<- t.C
		err := l.sdk.Shutdown()
		if err != nil {
			log.Logger.Errorw("Error during sdk shutdown notification", "error", err.Error())
		}
	}

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

	for i := range l.Players {
		// Set player states if player can bet in the round
		if l.Bank.HasZeroWallet(l.Players[i].ID) || l.Players[i].Left {
			l.Players[i].Active = false
		}else {
			l.Players[i].Active = true
		}
	}


}
