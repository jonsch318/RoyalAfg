package lobby

import "errors"

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
		m.logger.Warnw("no logger", "results", res)
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
		m.logger.Errorw("Error during search connection testing", "error", err)
	}

	return m.NewLobby(class)
}
