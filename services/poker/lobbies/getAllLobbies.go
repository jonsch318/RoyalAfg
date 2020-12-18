package lobbies

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/dto"
)

//GetAllLobbies return the first 10 lobbies for every buy in class
func (l *LobbyManager) GetAllLobbies() [][]dto.PublicLobby {

	results := make([][]dto.PublicLobby, len(l.PublicLobbies))

	for i := 0; i < len(l.PublicLobbies); i++ {
		t := len(l.PublicLobbies[i])
		if t > 5 {
			t = 5
		}
		results[i] = l.PublicLobbies[i][:t]
	}
	return results
}
