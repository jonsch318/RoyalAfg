package lobby

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/jonsch318/royalafg/pkg/log"
	"github.com/jonsch318/royalafg/pkg/poker/models"
)

func (m *Manager) SearchWithClass(class int) ([]models.LobbyBase, error) {

	if m.lobbies[class] == nil {
		return nil, nil
	}

	if m.classes == nil {
		return nil, errors.New("no registered buy in classes")
	}

	gameserver, err := m.agonesClient.AgonesV1().GameServers("default").List(context.Background(), metav1.ListOptions{
		TypeMeta: metav1.TypeMeta{
			Kind:       "agones.dev/v1",
			APIVersion: "GameServer",
		},
		LabelSelector: fmt.Sprintf("game=poker, class-index=%v", class),
	})

	if err != nil {
		return nil, err
	}

	lobbies := make([]models.LobbyBase, len(gameserver.Items))
	for _, gs := range gameserver.Items {
		players, err := strconv.Atoi(gs.Labels["players"])
		log.Logger.Debugf("Lobby found: %v", models.LobbyBase{
			LobbyID:    gs.Labels["lobbyId"],
			Class:      &m.classes[class],
			ClassIndex: class,
		})
		if err != nil {
			continue
		}
		lobbies = append(lobbies, models.LobbyBase{
			LobbyID:     gs.Labels["lobbyId"],
			Class:       &m.classes[class],
			ClassIndex:  class,
			PlayerCount: players,
		})
	}

	//Sort less [2,3,4,5, etc...]
	sort.SliceStable(lobbies, func(i, j int) bool {
		return biasForX(lobbies[i].PlayerCount, 9) < biasForX(lobbies[i].PlayerCount, 9)
	})

	log.Logger.Debugf("After sort: %v", lobbies)

	//ordered list of lobbies to try
	return lobbies, nil
}

func biasForX(i, x int) int {
	if i > x {
		return 2*x - i
	}
	return i
}

func (m *Manager) SearchWithParams(min, max, blind int) ([]models.LobbyBase, error) {

	if m.classes == nil {
		return nil, errors.New("no registered buy in classes")
	}

	found := make([]int, 0)
	for i, v := range m.classes {
		if v.Blind == blind && min >= v.Min {
			found = append(found, i)
		}
	}

	smallest :=
		struct {
			i int
			m int
		}{0, 0}
	for j := 0; j < len(found); j++ {
		c := m.classes[found[j]].Max
		if smallest.m >= c {
			smallest = struct {
				i int
				m int
			}{j, c}
		}
	}

	return m.SearchWithClass(smallest.i)
}
