package lobby

import (
	"errors"
	"time"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"
)

//RemovePlayerByID removes the given player identified by his id
func (l *Lobby) RemovePlayerByID(id string) error {

	i := l.FindPlayerByID(id)

	if i < 0 {
		log.Logger.Warnw("player could not be found", "id", id)
		return errors.New("the player is not in the lobby")
	}

	l.RemovalQueue.Enqueue(&l.Players[i])
	_ = l.round.Leave(id)
	if !l.GetGameStarted() {
		log.Logger.Debugf("Game not started start removal")
		l.RemoveAfterRound()
	}

	return nil
}

//RemoveAfterRound starts the recursive removal of hanging players
func (l *Lobby) RemoveAfterRound() {
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
	l.removePlayer()

	log.Logger.Debugf("Removed PlayerCount Player Count: %v", l.PlayerCount)
	if l.Count() <= 0 {
		log.Logger.Warnf("No more players... Notify agones sdk to shutdown")
		t := time.NewTimer(time.Second * 10)
		<-t.C
		err := l.sdk.Shutdown()
		if err != nil {
			log.Logger.Errorw("Error during sdk shutdown notification", "error", err.Error())
		}
	}

}

//PlayerRemoval removes all players in the removal queue.
func (l *Lobby) removePlayer() {

	player := l.RemovalQueue.Dequeue()
	if player == nil {
		//No player in queue
		return
	}

	if player.ID == "" {
		log.Logger.Debugf("player id nil")
	}
	//Get index of player
	i := l.FindPlayerByID(player.ID)
	if i < 0 {
		log.Logger.Warnf("Id [%v] not in lobby", player.ID)
		return
	}
	public := l.PublicPlayers[i]

	//Remove player from list, public list and bank
	l.Players = append(l.Players[:i], l.Players[i+1:]...)
	l.PublicPlayers = append(l.PublicPlayers[:i], l.PublicPlayers[i+1:]...)
	l.PlayerCount--

	err := l.Bank.RemovePlayer(player.ID)
	if err != nil {
		log.Logger.Errorw("error during removing player from bank", "error", err)
	}

	//Send leave event
	utils.SendToAll(l.Players, events.NewPlayerLeavesEvent(&public, i, len(l.Players), l.GameStarted))

	//Update gameserver label
	l.SetPlayerCountLabel()

	l.removePlayer()
}
