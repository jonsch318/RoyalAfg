package lobby

import "errors"

//RequestTicket requests a ticket of a available game server from the manger.
func (m *Manager) RequestTicket(class int) (string, error) {
	if class < 0 || class >= len(m.classes) {
		return "", errors.New("must be a valid class")
	}

	res, err := m.SearchWithClass(class)

	if err != nil {
		return "", err
	}

	if res == nil {
		//No lobby is found
		return m.NewLobby(class)
	}

	for i := len(res); i >= 0; i-- {
		addr := ""
		addr, err = m.Connect(res[i].LobbyID)
		if err == nil {
			return addr, nil
		}
	}

	return m.NewLobby(class)
}
