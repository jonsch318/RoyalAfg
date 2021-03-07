package lobby

import (
	"strconv"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
)

func IsInRange(players []models.Player, i int) bool {
	return i > 0 && i < len(players)
}

func (l *Lobby) Count() int {
	return l.PlayerCount + l.PlayerQueue.Length()
}

func (l *Lobby) SetPlayerCountLabel() {
	log.Logger.Debugf("PlayerCount %v", l.Count())
	err := l.sdk.SetLabel("players", strconv.Itoa(l.Count()))
	if err != nil {
		log.Logger.Warnw("error during player label set %v", "error", err.Error())
	}
}

//FindPlayerByID searches a player based on the given id linearly.
func (l *Lobby) FindPlayerByID(id string) int {
	//A map structure or hashtable like map[string(id)]int(index) would be more efficient, though it is not necessary for the maximum of 10 players in this list.
	for i, n := range l.Players {
		if n.ID == id {
			return i
		}
	}
	return -1
}
