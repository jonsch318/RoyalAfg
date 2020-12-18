package lobbies

import (
	"errors"
	"fmt"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/lobby"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"math/rand"
)

func (l *LobbyManager) DistributePlayer(player *models.Player, event *events.JoinEvent) (*lobby.Lobby, error) {
	ok := CheckBuyInClass(l.BuyInClasses, event.BuyInClass, event.BuyIn)
	if !ok {
		return nil, fmt.Errorf("The class does not match the buyin %d to %d", event.BuyInClass, event.BuyIn)
	}

	id := event.LobbyID
	if id == "" {
		class := l.LobbiesIndexed[event.BuyInClass]
		l.lock.RLock()
		id = find(class, l.Lobbies)
		l.lock.RUnlock()

		if id == "" {
			var err error
			id, err = l.AppendLobby(event.BuyInClass)
			if err != nil {
				if len(l.LobbiesIndexed[event.BuyInClass]) > 0 {
					i := rand.Intn(len(l.LobbiesIndexed[event.BuyInClass]))
					id = l.LobbiesIndexed[event.BuyInClass][i]
				}
			}
		}
	}

	if id == "" {
		return nil, errors.New("Something went wrong")
	}

	selectedLobby, ok := l.Lobbies[id]
	if !ok {
		return nil, errors.New("The lobby with the given id was not found")
	}
	selectedLobby.Join(player)

	return selectedLobby, nil
}
