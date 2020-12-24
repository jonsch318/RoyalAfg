package lobbies

import (
	"errors"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/dto"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/lobby"
)

//AppendLobby appends a newly created lobby to the lobby list but only if the maximum lobby count is not passed yet.
func (m *LobbyManager) AppendLobby(class int) (string, error) {
	if m.LobbyCount >= m.MaxCount {
		return "", errors.New("Maximum Lobby count already passed")
	}
	l := lobby.NewLobby(m.BuyInClasses[class][0], m.BuyInClasses[class][1], m.BuyInClasses[class][2], class)
	m.LobbiesIndexed[class] = append(m.LobbiesIndexed[0], l.LobbyID)
	m.Lobbies[l.LobbyID] = l
	m.PublicLobbies[class] = append(m.PublicLobbies[class], *dto.ToPublic(l.LobbyID, l.TotalPlayerCount(), l.MinBuyIn, l.MaxBuyIn, l.SmallBlind, l.LobbyClass))
	m.LobbyCount++
	return l.LobbyID, nil
}
