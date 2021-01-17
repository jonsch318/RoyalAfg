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
		//No lobby is found
		return m.NewLobby(class)
	}

	for i := len(res); i >= 0; i-- {
		res, err := m.Connect(res[i].LobbyID)
		if err == nil {
			return res, nil
		}
	}

	return m.NewLobby(class)
}
