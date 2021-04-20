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
	l.Players[index].Left = true
	err := l.round.Leave(l.Players[index].ID)

	if err != nil {
		return err
	}

	if !l.GetGameStarted() {
		log.Logger.Debugf("Game not started start removal")
		l.RemoveLeftPlayers()
	}

	return nil
}

//RemoveLeftPlayers starts the recursive removal of hanging players
func (l *Lobby) RemoveLeftPlayers() {

	log.Logger.Debugf("Starting Removal of players after round.")

	//Remove players that left during the round.
	for i := range l.Players {
		if l.Players[i].Left {
			l.RemovalQueue.Enqueue(RemovalRequest{
				Pos: i,
				ID:  l.Players[i].ID,
			})
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

	i := player.Pos

	if player.Pos < 0 || player.Pos >= len(l.Players) || len(l.Players) == 0 {
		log.Logger.Errorf("Player position is invalid. Player [%v]", player.ID)
		i = l.FindPlayerByID(player.ID)
	}

	if l.Players[player.Pos].ID != player.ID {
		log.Logger.Errorf("Removal Request of Player [%v] is not the player at position %v", player.ID, player.Pos)
		i = l.FindPlayerByID(player.ID)
	}

	if i == -1 {
		log.Logger.Errorf("Removal Request of Player [%v] is invalid because player does not exist", player.ID)
		l.removePlayer()
		return
	}

	public := l.PublicPlayers[i]
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
	l.Players = append(l.Players[:i], l.Players[i+1:]...)
	l.PublicPlayers = append(l.PublicPlayers[:i], l.PublicPlayers[i+1:]...)
	l.PlayerCount--

	//Send leave event
	utils.SendToAll(l.Players, events.NewPlayerLeavesEvent(&public, i, l.Count(), l.GameStarted))

	//Update gameserver label
	l.SetPlayerCountLabel()

	l.removePlayer()
}
