package lobby

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"log"
)

//EnqueuePlayer enqueues a player that is added to the lobby when a position is free
func (l *Lobby) EnqueuePlayer(player *models.Player) {
	l.PlayerQueue = append(l.PlayerQueue, player)
}

//DequeuePlayer dequeues the player and adds it to the position
func (l *Lobby) DequeuePlayer(pos int) (*models.Player, bool) {
	if len(l.PlayerQueue) > 0 && len(l.Players) < 10 {
		player := l.PlayerQueue[0]
		if pos < 0 {
			l.Players = append(l.Players, *player)
		} else {
			l.Players[pos] = *player
		}

		l.PlayerQueue = l.PlayerQueue[1:]
		log.Printf("Dequeued Player [%v] into position [%v]", player, pos)
		return player, true
	}

	return nil, false
}
