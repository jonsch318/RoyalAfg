package lobby

//FindPlayerByID searches a player based on the given id linearly.
func (l *Lobby) FindPlayerByID(id string) int {
	for i, n := range l.Players {
		if n.ID == id {
			return i
		}
	}
	return -1
}
