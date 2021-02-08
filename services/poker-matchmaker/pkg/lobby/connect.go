package lobby

import (
	"context"
	"encoding/json"
	fmt "fmt"
	"io/ioutil"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (m *Manager) Connect(id string) (*TicketRequestResult, error) {

	return m.AgonesConnect(id)

	addr, err := m.rdg.Get(context.Background(), id).Result()
	if err != nil {
		return nil, err
	}
	m.logger.Infof("Address lookup: %v", addr)

	running := false
	running, err = m.PingHealth(addr, id)

	if err != nil || !running {
		m.rdg.Del(context.Background(), id)
		return nil, fmt.Errorf("error during ping: %s", err)
	}

	//GameServer pointed by addr is healthy
	return &TicketRequestResult{Address: addr, LobbyId: id}, err
}

func (m *Manager) AgonesConnect(id string) (*TicketRequestResult, error) {
	list, err := m.agonesClient.AgonesV1().GameServers("default").List(metav1.ListOptions{
		TypeMeta: metav1.TypeMeta{
			Kind:       "agones.dev/v1",
			APIVersion: "GameServer",
		},
		LabelSelector:       fmt.Sprintf("game=poker,lobbyId=%v", id),
		Watch:               false,
		AllowWatchBookmarks: false,
	})
	if err != nil {
		m.logger.Errorw("err during connect", "error", err)
		return nil, err
	}
	m.logger.Infow("Agones list results", "id", id, "list", list)

	if len(list.Items) < 1 {
		return nil, fmt.Errorf("no lobbies found with the given id")
	}
	return &TicketRequestResult{Address: fmt.Sprintf("%s:%v", list.Items[0].Status.Address, list.Items[0].Status.Ports[0].Port), LobbyId: id}, nil
}

type HealthPingResponse struct {
	Count   int    `json:"count"`
	Running bool   `json:"running"`
	LobbyID string `json:"lobbyId"`
}

func (m *Manager) PingHealth(addr string, id string) (bool, error) {
	res, err := http.Get(fmt.Sprintf("http://%v/api/poker/health", addr))

	if err != nil {
		return false, err
	}

	js, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		return false, err
	}

	var ping HealthPingResponse
	err = json.Unmarshal(js, &ping)
	if err != nil {
		return false, err
	}

	return ping.Running && ping.LobbyID == id, nil
}
