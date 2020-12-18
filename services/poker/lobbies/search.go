package lobbies

func (l *LobbyManager) Search() []string {

	a := make([]string, 0)

	for _, v := range l.Lobbies {
		if len(v.Players) < 10 {
			a = append(a, v.LobbyID)
		}
	}
	return a
}
