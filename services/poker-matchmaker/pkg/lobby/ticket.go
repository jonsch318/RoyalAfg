package lobby

import (
	"errors"
)

type TicketRequestResult struct {
	Address string
	LobbyId string
}

//RequestTicket requests a ticket of a available game server from the manger.
func (m *Manager) RequestTicket(class int) (*TicketRequestResult, error) {
	if class < 0 || class >= len(m.classes) {
		return nil, errors.New("must be a valid class")
	}

	res, err := m.SearchWithClass(class)

	if err != nil {
		return nil, err
	}

	if res == nil {
		m.logger.Warnw("no lobbies registered")
		//No lobby is found
		return m.NewLobby(class)
	}

	m.logger.Infow("Search result ranking", "results", res)

	for i := len(res) - 1; i >= 0; i-- {
		ret, err := m.Connect(res[i].LobbyID)
		if err == nil {
			m.logger.Infof("Connection success [%v]", res[i].LobbyID)
			return ret, nil
		}
		m.Remove(class, res[i].LobbyID)
		m.logger.Errorw("Error during search connection testing", "error", err)
	}

	return m.NewLobby(class)
}

func (m *Manager) Remove(class int, lobbyId string){
	i := m.GetIndex(class, lobbyId)
	if i < 0 {
		return
	}
	m.RemoveLobby(class, i, lobbyId)
}

func (m *Manager) RemoveLobby(class int, i int, lobbyId string,) {

	m.lobbies[class][i] = m.lobbies[class][len(m.lobbies[class])-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	m.lobbies[class] = m.lobbies[class][:len(m.lobbies[class])-1]
}

func (m *Manager) GetIndex(class int, id string) int {
	for i, base := range m.lobbies[class] {
		if base.LobbyID == id {
			return i
		}
	}
	return -1
}