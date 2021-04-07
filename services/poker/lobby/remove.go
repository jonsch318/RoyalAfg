package lobby

import (
	"errors"
	"time"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"
)

type RemovalRequest struct {
	Pos int
	ID  string
}

//RemovePlayerByID removes the given player identified by his id. (LOCKS)
func (l *Lobby) RemovePlayerByID(id string) error {
	l.lock.RLock()
	i := l.FindPlayerByID(id)
	l.lock.RUnlock()
	if i < 0 {
		log.Logger.Warnw("player could not be found", "id", id)
		return errors.New("the player is not in the lobby")
	}

	return l.RemovePlayer(i)
}

func (l *Lobby) RemovePlayer(index int) error {
	l.RemovalQueue.Enqueue(RemovalRequest{
		Pos: index,
		ID:  l.Players[index].ID,
	})

	_ = l.round.Leave(l.Players[index].ID)
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

	log.Logger.Debugf("Starting Removal of players after round.")

	//Remove players that left during the round.
	for _, player := range l.Players {
		if player.Left {
			l.RemovalQueue.Enqueue(&player)
		}
	}

	//Remove all hanging players and update player count
	l.removePlayer()

	log.Logger.Debugf("Updated: Player Count: %v", l.PlayerCount)
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

	r := l.RemovalQueue.Dequeue()
	if r == nil {
		//No player in queue
		return
	}

	player := r.(RemovalRequest)
	log.Logger.Warnf("REMOVING Player [%v]", player.ID)

	if player.Pos < 0 || player.Pos > len(l.Players) || len(l.Players) == 0 {
		log.Logger.Errorf("Player position is invalid. Player [%v]", player.ID)
		l.removePlayer()
		return
	}

	if l.Players[player.Pos].ID != player.ID {
		log.Logger.Errorf("Removal Request of Player [%v] is not the player at position %v", player.ID, player.Pos)
		l.removePlayer()
		return
	}

	public := l.PublicPlayers[player.Pos]
	if public.ID != player.ID {
		log.Logger.Errorf("Public Playerlist is not syncronised with Playerlist. [%v] != [%v]", public.ID, player.ID)
		l.removePlayer()
		return
	}

	err := l.Bank.RemovePlayer(player.ID)
	if err != nil {
		log.Logger.Errorw("error during removing player from bank", "error", err)
		l.removePlayer()
		return
	}

	//Remove player from list, public list and bank
	l.Players = append(l.Players[:player.Pos], l.Players[player.Pos+1:]...)
	l.PublicPlayers = append(l.PublicPlayers[:player.Pos], l.PublicPlayers[player.Pos+1:]...)
	l.PlayerCount--

	//Send leave event
	utils.SendToAll(l.Players, events.NewPlayerLeavesEvent(&public, player.Pos, l.Count(), l.GameStarted))

	//Update gameserver label
	l.SetPlayerCountLabel()

	l.removePlayer()
}
