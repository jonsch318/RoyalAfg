package lobby

import (
	"context"
	"fmt"
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/poker/models"
)

func (m *Manager) GetRegisteredLobbiesOfClass(index, count, class int) []models.LobbyBase {

	if class < 0 || class >= len(m.lobbies) {
		m.logger.Errorw("tried fetching registered lobbies of an unregistered class", "error", "class does not exist", "class", class)
		return nil
	}

	if index == 0 && count == 0 {
		return m.lobbies[class]
	}
	if len(m.lobbies[class]) <= 0 {
		return nil
	}

	if len(m.lobbies[class]) < count+index {
		//the query can not be fully answered, so we do our best to satisfy it as much as possible
		if len(m.lobbies[class]) <= count {
			//smaller than our result count. just use fullHistory
			return m.lobbies[class]
		}

		//return l from last.
		return m.lobbies[class][len(m.lobbies)-count:]
	}

	return m.lobbies[class][index : index+count]
}

func (m *Manager) GetRegisteredLobbies(count int) [][]models.LobbyBase {

	if m.classes == nil {
		m.logger.Errorf("No registered classes")
		return nil
	}

	list, err := m.agonesClient.AgonesV1().GameServers("default").List(context.Background(), metav1.ListOptions{
		TypeMeta: metav1.TypeMeta{
			Kind:       "agones.dev/v1",
			APIVersion: "GameServer",
		},
		LabelSelector:       fmt.Sprintf("game=poker"),
		Watch:               false,
		AllowWatchBookmarks: false,
		Limit:               int64(count),
	})

	if err != nil {
		m.logger.Errorw("error during agones lobbies", "error", err)
		return nil
	}
	m.logger.Info("Agones list decoded")

	lobbies := make([][]models.LobbyBase, len(m.classes))
	for _, gs := range list.Items {

		class, err := strconv.Atoi(gs.Labels["class-index"])
		if err != nil {
			m.logger.Errorw("error decoding class lobby", "error", err, "string", gs.Labels["class-index"])
			continue
		}

		players, err := strconv.Atoi(gs.Labels["players"])
		if err != nil {
			m.logger.Errorw("error decoding players lobby", "error", err, gs.Labels["players"])
		}

		if lobbies[class] == nil {
			lobbies[class] = make([]models.LobbyBase, 0)
		}

		lby := models.LobbyBase{
			LobbyID:     gs.Labels["lobbyId"],
			Class:       &m.classes[class],
			ClassIndex:  class,
			PlayerCount: players,
		}

		lobbies[class] = append(lobbies[class], lby)
		m.logger.Errorw("lobby [%v] %v => %v", lby.LobbyID, lby.Class, lby.PlayerCount)

	}

	//if len(m.lobbies) <= 0 {
	//	return nil
	//}
	//
	//classCount := count / len(m.classes)
	//
	////This would be extended with filter parameters and maybe divided into multiple services
	//filtered := make([][]models.LobbyBase, len(m.classes))
	//for i, lobby := range m.lobbies {
	//
	//	//resize array once not for every lobby
	//	filtered[i] = make([]models.LobbyBase, classCount)
	//
	//	for j := 0; j < classCount; j++ {
	//		if j >= len(lobby) {
	//			break
	//		}
	//		filtered[i][j] = lobby[j]
	//	}
	//}
	//
	//return filtered

	return lobbies
}
